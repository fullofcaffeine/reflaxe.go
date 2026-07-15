package hxrt

import (
	"io"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

type FileInput struct {
	file *os.File
}

type FileOutput struct {
	file *os.File
}

func SysGetCwd() *string {
	cwd, err := os.Getwd()
	if err != nil {
		return StringFromLiteral("")
	}
	return StringFromLiteral(cwd)
}

func SysArgs() []*string {
	args := os.Args
	if len(args) <= 1 {
		return []*string{}
	}
	out := make([]*string, 0, len(args)-1)
	for _, arg := range args[1:] {
		out = append(out, StringFromLiteral(arg))
	}
	return out
}

func SysGetEnv(key *string) *string {
	if key == nil {
		return nil
	}
	value, ok := os.LookupEnv(*key)
	if !ok {
		return nil
	}
	return StringFromLiteral(value)
}

// SysPutEnv changes one environment entry and reports the native OS failure.
//
// Portable Sys.putEnv deliberately ignores this error because the upstream
// Haxe 4.3.7 eval contract exposes a Void operation that does not throw for an
// invalid key. Keeping the error in this helper lets native facades preserve Go
// os.Setenv/os.Unsetenv behavior instead of inheriting that portable contract.
func SysPutEnv(key *string, value *string) error {
	if key == nil {
		return nil
	}
	if value == nil {
		return os.Unsetenv(*key)
	}
	return os.Setenv(*key, *value)
}

func SysCommand(command *string, args []*string) int {
	if command == nil {
		return -1
	}
	cmd := exec.Command(*StdString(command), StringSlice(args)...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			return exitErr.ExitCode()
		}
		return -1
	}
	return 0
}

func SysExit(code int) {
	os.Exit(code)
}

func SysEnvironment() map[string]string {
	out := map[string]string{}
	for _, entry := range os.Environ() {
		key, value, ok := strings.Cut(entry, "=")
		if !ok {
			out[entry] = ""
			continue
		}
		out[key] = value
	}
	return out
}

func SysSystemName() *string {
	switch runtime.GOOS {
	case "darwin":
		return StringFromLiteral("Mac")
	case "linux":
		return StringFromLiteral("Linux")
	case "windows":
		return StringFromLiteral("Windows")
	case "freebsd", "openbsd", "netbsd", "dragonfly":
		return StringFromLiteral("BSD")
	default:
		return StringFromLiteral(runtime.GOOS)
	}
}

// SysSleep suspends the current goroutine for a Haxe seconds-based duration.
//
// What: Implements the portable Sys.sleep(Float) blocking contract.
// Why: The mainstream Haxe declaration lowers to a target-owned Sys_sleep
// symbol, while Go's time.Sleep accepts a nanosecond Duration instead of
// seconds. Keeping conversion here avoids library behavior in compiler GoRaw.
// How: Clamp non-positive and NaN inputs to an immediate return, convert
// positive seconds through time.Second, and delegate scheduling to time.Sleep.
func SysSleep(seconds float64) {
	if !(seconds > 0) {
		return
	}
	time.Sleep(time.Duration(seconds * float64(time.Second)))
}

// FileSaveContent stores text without collapsing write failures into success.
func FileSaveContent(path *string, content *string) error {
	return os.WriteFile(*StdString(path), []byte(*StdString(content)), 0o644)
}

// FileGetContent returns text and preserves missing, permission, and I/O errors.
func FileGetContent(path *string) (*string, error) {
	raw, err := os.ReadFile(*StdString(path))
	if err != nil {
		return nil, err
	}
	return StringFromLiteral(string(raw)), nil
}

func OpenFileInput(path *string) (*FileInput, error) {
	file, err := os.Open(*StdString(path))
	if err != nil {
		return nil, err
	}
	return &FileInput{file: file}, nil
}

func openFileOutput(path *string, flags int) (*FileOutput, error) {
	file, err := os.OpenFile(*StdString(path), flags, 0o644)
	if err != nil {
		return nil, err
	}
	return &FileOutput{file: file}, nil
}

func OpenFileWriteOutput(path *string) (*FileOutput, error) {
	return openFileOutput(path, os.O_CREATE|os.O_WRONLY|os.O_TRUNC)
}

func OpenFileAppendOutput(path *string) (*FileOutput, error) {
	return openFileOutput(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND)
}

func OpenFileUpdateOutput(path *string) (*FileOutput, error) {
	return openFileOutput(path, os.O_CREATE|os.O_RDWR)
}

func FileGetBytes(path *string) ([]byte, error) {
	return os.ReadFile(*StdString(path))
}

func FileSaveBytes(path *string, raw []byte) error {
	return os.WriteFile(*StdString(path), raw, 0o644)
}

func FileCopy(srcPath *string, dstPath *string) error {
	raw, err := os.ReadFile(*StdString(srcPath))
	if err != nil {
		return err
	}
	return os.WriteFile(*StdString(dstPath), raw, 0o644)
}

func (self *FileInput) ReadByte() (int, bool, error) {
	if self == nil || self.file == nil {
		return 0, true, io.EOF
	}
	var one [1]byte
	n, err := self.file.Read(one[:])
	if err != nil {
		if err == io.EOF {
			return 0, true, nil
		}
		return 0, false, err
	}
	if n == 0 {
		return 0, true, nil
	}
	return int(one[0]), false, nil
}

func (self *FileInput) Tell() (int, error) {
	if self == nil || self.file == nil {
		return 0, os.ErrInvalid
	}
	offset, err := self.file.Seek(0, io.SeekCurrent)
	return int(offset), err
}

func (self *FileInput) Seek(offset int, whence int) error {
	if self == nil || self.file == nil {
		return os.ErrInvalid
	}
	_, err := self.file.Seek(int64(offset), whence)
	return err
}

func (self *FileInput) Eof() (bool, error) {
	if self == nil || self.file == nil {
		return true, os.ErrInvalid
	}
	offset, err := self.file.Seek(0, io.SeekCurrent)
	if err != nil {
		return true, err
	}
	info, err := self.file.Stat()
	if err != nil {
		return true, err
	}
	return offset >= info.Size(), nil
}

func (self *FileInput) Close() error {
	if self == nil || self.file == nil {
		return nil
	}
	err := self.file.Close()
	self.file = nil
	return err
}

func (self *FileOutput) WriteByte(value int) error {
	if self == nil || self.file == nil {
		return os.ErrInvalid
	}
	_, err := self.file.Write([]byte{byte(value)})
	return err
}

func (self *FileOutput) Tell() (int, error) {
	if self == nil || self.file == nil {
		return 0, os.ErrInvalid
	}
	offset, err := self.file.Seek(0, io.SeekCurrent)
	return int(offset), err
}

func (self *FileOutput) Seek(offset int, whence int) error {
	if self == nil || self.file == nil {
		return os.ErrInvalid
	}
	_, err := self.file.Seek(int64(offset), whence)
	return err
}

func (self *FileOutput) Flush() error {
	if self == nil || self.file == nil {
		return nil
	}
	return self.file.Sync()
}

func (self *FileOutput) Close() error {
	if self == nil || self.file == nil {
		return nil
	}
	err := self.file.Close()
	self.file = nil
	return err
}
