package hxrt

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"runtime"
	"sync"
)

// ProcessOutput is one readable child-process stream.
//
// A buffered byte reader is used instead of bufio.Scanner so an ordinary line
// is never reclassified as EOF merely because it exceeds Scanner's token cap.
type ProcessOutput struct {
	reader    *bufio.Reader
	pipe      io.ReadCloser
	closeOnce sync.Once
	closeErr  error
}

func newProcessOutput(pipe io.ReadCloser) *ProcessOutput {
	return &ProcessOutput{reader: bufio.NewReader(pipe), pipe: pipe}
}

// ReadByte preserves a normal EOF separately from other stream failures.
func (self *ProcessOutput) ReadByte() (int, bool, error) {
	if self == nil || self.reader == nil {
		return 0, false, os.ErrInvalid
	}
	value, err := self.reader.ReadByte()
	if errors.Is(err, io.EOF) {
		return 0, true, nil
	}
	if err != nil {
		return 0, false, err
	}
	return int(value), false, nil
}

func (self *ProcessOutput) Close() error {
	if self == nil || self.pipe == nil {
		return nil
	}
	self.closeOnce.Do(func() {
		self.closeErr = self.pipe.Close()
		if errors.Is(self.closeErr, os.ErrClosed) {
			self.closeErr = nil
		}
	})
	return self.closeErr
}

// ProcessInput is the writable child stdin stream.
type ProcessInput struct {
	mu        sync.Mutex
	pipe      io.WriteCloser
	closed    bool
	closeOnce sync.Once
	closeErr  error
}

func newProcessInput(pipe io.WriteCloser) *ProcessInput {
	return &ProcessInput{pipe: pipe}
}

func (self *ProcessInput) WriteByte(value int) error {
	if self == nil {
		return os.ErrInvalid
	}
	self.mu.Lock()
	defer self.mu.Unlock()
	if self.closed || self.pipe == nil {
		return io.ErrClosedPipe
	}
	_, err := self.pipe.Write([]byte{byte(value)})
	return err
}

// Flush is intentionally a no-op because os/exec exposes an unbuffered pipe.
func (self *ProcessInput) Flush() error {
	return nil
}

func (self *ProcessInput) Close() error {
	if self == nil || self.pipe == nil {
		return nil
	}
	self.closeOnce.Do(func() {
		self.mu.Lock()
		self.closed = true
		self.closeErr = self.pipe.Close()
		self.mu.Unlock()
		if errors.Is(self.closeErr, os.ErrClosed) {
			self.closeErr = nil
		}
	})
	return self.closeErr
}

// Process owns a started command and its three portable Haxe streams.
// Waiting is started lazily so stdout/stderr can be drained before exec.Cmd
// closes its pipes, while close still starts a background reaper.
type Process struct {
	cmd       *exec.Cmd
	stdout    *ProcessOutput
	stderr    *ProcessOutput
	stdin     *ProcessInput
	waitOnce  sync.Once
	waitDone  chan struct{}
	waitErr   error
	closeOnce sync.Once
	closeErr  error
}

// ProcessExitStatus is the typed nonblocking status returned to staged Process.
//
// What: Carries an exit code together with whether that code is available.
// Why: Haxe exposes Null<Int>, but using any at the native boundary would erase
// the distinction between a running process and a completed process with code 0.
// How: Keep availability explicit in Go; staged Haxe performs the final Null<Int>
// conversion at the public API boundary.
type ProcessExitStatus struct {
	Code      int
	Available bool
}

func processCommand(command string, args []*string) *exec.Cmd {
	if args != nil {
		return exec.Command(command, StringSlice(args)...)
	}
	if runtime.GOOS == "windows" {
		return exec.Command("cmd", "/C", command)
	}
	return exec.Command("/bin/sh", "-c", command)
}

// NewProcess creates all pipes before starting the child and returns the first
// native failure instead of a partially initialized process.
func NewProcess(command *string, args []*string) (*Process, error) {
	if command == nil {
		return nil, fmt.Errorf("process command is nil")
	}

	cmd := processCommand(*command, args)
	stdoutPipe, err := cmd.StdoutPipe()
	if err != nil {
		return nil, fmt.Errorf("create process stdout pipe: %w", err)
	}
	stderrPipe, err := cmd.StderrPipe()
	if err != nil {
		_ = stdoutPipe.Close()
		return nil, fmt.Errorf("create process stderr pipe: %w", err)
	}
	stdinPipe, err := cmd.StdinPipe()
	if err != nil {
		_ = stdoutPipe.Close()
		_ = stderrPipe.Close()
		return nil, fmt.Errorf("create process stdin pipe: %w", err)
	}

	if err := cmd.Start(); err != nil {
		_ = stdinPipe.Close()
		_ = stdoutPipe.Close()
		_ = stderrPipe.Close()
		return nil, fmt.Errorf("start process %q: %w", *command, err)
	}

	return &Process{
		cmd:      cmd,
		stdout:   newProcessOutput(stdoutPipe),
		stderr:   newProcessOutput(stderrPipe),
		stdin:    newProcessInput(stdinPipe),
		waitDone: make(chan struct{}),
	}, nil
}

func (self *Process) Stdout() *ProcessOutput {
	if self == nil {
		return nil
	}
	return self.stdout
}

func (self *Process) Stderr() *ProcessOutput {
	if self == nil {
		return nil
	}
	return self.stderr
}

func (self *Process) Stdin() *ProcessInput {
	if self == nil {
		return nil
	}
	return self.stdin
}

func (self *Process) GetPid() (int, error) {
	if self == nil || self.cmd == nil || self.cmd.Process == nil {
		return 0, os.ErrInvalid
	}
	return self.cmd.Process.Pid, nil
}

func (self *Process) ensureWait() {
	if self == nil || self.cmd == nil || self.waitDone == nil {
		return
	}
	self.waitOnce.Do(func() {
		go func() {
			self.waitErr = self.cmd.Wait()
			close(self.waitDone)
		}()
	})
}

// ExitCode returns (code, available, error). A nonzero child exit is a normal
// available code; only failures outside exec.ExitError occupy error.
func (self *Process) ExitCode(block bool) (int, bool, error) {
	if self == nil || self.cmd == nil || self.waitDone == nil {
		return 0, false, os.ErrInvalid
	}
	self.ensureWait()
	if block {
		<-self.waitDone
	} else {
		select {
		case <-self.waitDone:
		default:
			return 0, false, nil
		}
	}

	if self.waitErr == nil {
		return 0, true, nil
	}
	var exitErr *exec.ExitError
	if errors.As(self.waitErr, &exitErr) {
		return exitErr.ExitCode(), true, nil
	}
	return 0, true, self.waitErr
}

func (self *Process) Kill() error {
	if self == nil || self.cmd == nil || self.cmd.Process == nil {
		return os.ErrInvalid
	}
	return self.cmd.Process.Kill()
}

// Close releases streams and starts a background reap without killing the
// child. Process termination remains the explicit responsibility of Kill.
func (self *Process) Close() error {
	if self == nil {
		return nil
	}
	self.closeOnce.Do(func() {
		var closeErrors []error
		if err := self.stdin.Close(); err != nil {
			closeErrors = append(closeErrors, err)
		}
		if err := self.stdout.Close(); err != nil {
			closeErrors = append(closeErrors, err)
		}
		if err := self.stderr.Close(); err != nil {
			closeErrors = append(closeErrors, err)
		}
		self.ensureWait()
		self.closeErr = errors.Join(closeErrors...)
	})
	return self.closeErr
}

// ProcessCreate starts one child process for the staged Haxe Process wrapper.
//
// What: Exposes process construction as a typed capability.
// Why: Native startup and pipe failures belong in hxrt, while the public Haxe
// constructor and stream objects belong in canonical staged source.
// How: Delegate to NewProcess and translate its typed Go error into the portable
// Haxe exception carrier before returning the opaque process handle.
func ProcessCreate(command *string, args []*string) *Process {
	process, err := NewProcess(command, args)
	if err != nil {
		Throw(StringFromLiteral(err.Error()))
		return nil
	}
	return process
}

// ProcessStdout returns the typed stdout pipe owned by process.
func ProcessStdout(process *Process) *ProcessOutput {
	if process == nil {
		return nil
	}
	return process.Stdout()
}

// ProcessStderr returns the typed stderr pipe owned by process.
func ProcessStderr(process *Process) *ProcessOutput {
	if process == nil {
		return nil
	}
	return process.Stderr()
}

// ProcessStdin returns the typed stdin pipe owned by process.
func ProcessStdin(process *Process) *ProcessInput {
	if process == nil {
		return nil
	}
	return process.Stdin()
}

// ProcessOutputReadByteValue reads one byte and uses -1 only for ordinary EOF.
func ProcessOutputReadByteValue(output *ProcessOutput) int {
	if output == nil {
		Throw(StringFromLiteral("Process output is closed"))
		return 0
	}
	value, eof, err := output.ReadByte()
	if err != nil {
		Throw(StringFromLiteral(err.Error()))
		return 0
	}
	if eof {
		return -1
	}
	return value
}

// ProcessOutputReadValues reads up to length byte values and returns an empty
// slice only when the stream has reached ordinary EOF.
func ProcessOutputReadValues(output *ProcessOutput, length int) []int {
	if length <= 0 {
		return []int{}
	}
	values := make([]int, 0, length)
	for len(values) < length {
		value := ProcessOutputReadByteValue(output)
		if value < 0 {
			break
		}
		values = append(values, value)
	}
	return values
}

// ProcessOutputClose releases one readable process pipe.
func ProcessOutputClose(output *ProcessOutput) {
	if output == nil {
		return
	}
	if err := output.Close(); err != nil {
		Throw(StringFromLiteral(err.Error()))
	}
}

// ProcessInputWriteByteValue writes one byte and reports whether the staged
// wrapper should preserve success or translate the native failure to Eof.
func ProcessInputWriteByteValue(input *ProcessInput, value int) bool {
	return input != nil && input.WriteByte(value) == nil
}

// ProcessInputWriteValues writes all values or reports one pipe failure.
func ProcessInputWriteValues(input *ProcessInput, values []int) bool {
	if input == nil {
		return false
	}
	for _, value := range values {
		if err := input.WriteByte(value); err != nil {
			return false
		}
	}
	return true
}

// ProcessInputFlush flushes the writable process pipe.
func ProcessInputFlush(input *ProcessInput) {
	if input == nil {
		return
	}
	if err := input.Flush(); err != nil {
		Throw(StringFromLiteral(err.Error()))
	}
}

// ProcessInputClose releases one writable process pipe.
func ProcessInputClose(input *ProcessInput) {
	if input == nil {
		return
	}
	if err := input.Close(); err != nil {
		Throw(StringFromLiteral(err.Error()))
	}
}

// ProcessPid returns the native child process identifier.
func ProcessPid(process *Process) int {
	pid, err := process.GetPid()
	if err != nil {
		Throw(StringFromLiteral(err.Error()))
		return 0
	}
	return pid
}

// ProcessExitStatusValue waits or polls for one typed child exit status.
func ProcessExitStatusValue(process *Process, block bool) *ProcessExitStatus {
	code, available, err := process.ExitCode(block)
	if err != nil {
		Throw(StringFromLiteral(err.Error()))
		return &ProcessExitStatus{}
	}
	return &ProcessExitStatus{Code: code, Available: available}
}

// ProcessKill terminates one child process and preserves native failure details.
func ProcessKill(process *Process) {
	if err := process.Kill(); err != nil {
		Throw(StringFromLiteral(err.Error()))
	}
}

// ProcessClose releases process resources without killing the child.
func ProcessClose(process *Process) {
	if process == nil {
		return
	}
	if err := process.Close(); err != nil {
		Throw(StringFromLiteral(err.Error()))
	}
}
