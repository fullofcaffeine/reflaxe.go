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
