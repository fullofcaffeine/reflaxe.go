package hxrt

import (
	"io"
	"os"
)

// FileInput is the opaque native handle behind staged sys.io.FileInput.
//
// What: Owns one readable Go file and its close policy.
// Why: A live *os.File cannot be represented by portable Haxe source, while the
// public stream behavior and Haxe exceptions do not belong in GoCompiler.
// How: Staged Haxe stores this pointer through a typed extern and calls the
// exported capability functions below.
type FileInput struct {
	file     *os.File
	ownsFile bool
}

// FileOutput is the opaque native handle behind staged sys.io.FileOutput.
//
// What: Owns one writable Go file plus its flush and close policy.
// Why: Standard streams are non-owning and unbuffered, while ordinary file
// handles must close and sync; this distinction is native resource state.
// How: Constructors set the policy once and staged Haxe invokes typed helpers.
type FileOutput struct {
	file        *os.File
	ownsFile    bool
	syncOnFlush bool
}

// SysStdin returns a non-owning runtime view of the process standard input.
func SysStdin() *FileInput {
	return &FileInput{file: os.Stdin}
}

// SysStdout returns a non-owning, unbuffered view of standard output.
func SysStdout() *FileOutput {
	return &FileOutput{file: os.Stdout}
}

// SysStderr returns a non-owning, unbuffered view of standard error.
func SysStderr() *FileOutput {
	return &FileOutput{file: os.Stderr}
}

// FileSaveContent stores text without collapsing write failures into success.
func FileSaveContent(path *string, content *string) error {
	return os.WriteFile(*StdString(path), []byte(*StdString(content)), 0o644) //nolint:gosec // Haxe file APIs create ordinary user-visible files and honor the process umask.
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
	return &FileInput{file: file, ownsFile: true}, nil
}

func openFileOutput(path *string, flags int) (*FileOutput, error) {
	file, err := os.OpenFile(*StdString(path), flags, 0o644) //nolint:gosec // Haxe file APIs create ordinary user-visible files and honor the process umask.
	if err != nil {
		return nil, err
	}
	return &FileOutput{file: file, ownsFile: true, syncOnFlush: true}, nil
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
	return os.WriteFile(*StdString(path), raw, 0o644) //nolint:gosec // Haxe file APIs create ordinary user-visible files and honor the process umask.
}

func FileCopy(srcPath *string, dstPath *string) error {
	raw, err := os.ReadFile(*StdString(srcPath))
	if err != nil {
		return err
	}
	return os.WriteFile(*StdString(dstPath), raw, 0o644) //nolint:gosec // Haxe file APIs create ordinary user-visible files and honor the process umask.
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
	if !self.ownsFile {
		self.file = nil
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
	if !self.syncOnFlush {
		return nil
	}
	return self.file.Sync()
}

func (self *FileOutput) Close() error {
	if self == nil || self.file == nil {
		return nil
	}
	if !self.ownsFile {
		self.file = nil
		return nil
	}
	err := self.file.Close()
	self.file = nil
	return err
}

func fileThrow(err error) {
	if err != nil {
		Throw(StringFromLiteral(err.Error()))
	}
}

func bytesToValues(raw []byte) []int {
	values := make([]int, len(raw))
	for index, value := range raw {
		values[index] = int(value)
	}
	return values
}

func valuesToBytes(values []int) []byte {
	raw := make([]byte, len(values))
	for index, value := range values {
		raw[index] = byte(value)
	}
	return raw
}

// FileReadContent is the Haxe-shaped typed text-read capability.
func FileReadContent(path *string) *string {
	content, err := FileGetContent(path)
	if err != nil {
		fileThrow(err)
		return StringFromLiteral("")
	}
	return content
}

// FileWriteContent is the Haxe-shaped typed text-write capability.
func FileWriteContent(path *string, content *string) {
	fileThrow(FileSaveContent(path, content))
}

// FileReadByteValues returns arbitrary file bytes as portable integer values.
func FileReadByteValues(path *string) []int {
	raw, err := FileGetBytes(path)
	if err != nil {
		fileThrow(err)
		return []int{}
	}
	return bytesToValues(raw)
}

// FileWriteByteValues stores portable integer byte values without text conversion.
func FileWriteByteValues(path *string, values []int) {
	fileThrow(FileSaveBytes(path, valuesToBytes(values)))
}

// FileCopyContents copies one file and preserves native failures.
func FileCopyContents(srcPath *string, dstPath *string) {
	fileThrow(FileCopy(srcPath, dstPath))
}

func FileOpenInput(path *string) *FileInput {
	handle, err := OpenFileInput(path)
	if err != nil {
		fileThrow(err)
		return &FileInput{}
	}
	return handle
}

func FileOpenWrite(path *string) *FileOutput {
	handle, err := OpenFileWriteOutput(path)
	if err != nil {
		fileThrow(err)
		return &FileOutput{}
	}
	return handle
}

func FileOpenAppend(path *string) *FileOutput {
	handle, err := OpenFileAppendOutput(path)
	if err != nil {
		fileThrow(err)
		return &FileOutput{}
	}
	return handle
}

func FileOpenUpdate(path *string) *FileOutput {
	handle, err := OpenFileUpdateOutput(path)
	if err != nil {
		fileThrow(err)
		return &FileOutput{}
	}
	return handle
}

// FileInputReadByteValue returns -1 for EOF and throws native failures.
func FileInputReadByteValue(handle *FileInput) int {
	value, eof, err := handle.ReadByte()
	if err != nil {
		fileThrow(err)
		return -1
	}
	if eof {
		return -1
	}
	return value
}

// FileInputReadValues reads at most length bytes and returns an empty slice at EOF.
func FileInputReadValues(handle *FileInput, length int) []int {
	if length <= 0 {
		return []int{}
	}
	if handle == nil || handle.file == nil {
		fileThrow(os.ErrInvalid)
		return []int{}
	}
	raw := make([]byte, length)
	count, err := handle.file.Read(raw)
	if err != nil && err != io.EOF {
		fileThrow(err)
		return []int{}
	}
	return bytesToValues(raw[:count])
}

func FileInputTell(handle *FileInput) int {
	position, err := handle.Tell()
	if err != nil {
		fileThrow(err)
		return 0
	}
	return position
}

func FileInputSeek(handle *FileInput, offset int, whence int) {
	fileThrow(handle.Seek(offset, whence))
}

func FileInputEof(handle *FileInput) bool {
	eof, err := handle.Eof()
	if err != nil {
		fileThrow(err)
		return true
	}
	return eof
}

func FileInputClose(handle *FileInput) {
	fileThrow(handle.Close())
}

func FileOutputWriteByteValue(handle *FileOutput, value int) {
	fileThrow(handle.WriteByte(value))
}

// FileOutputWriteValues writes exactly one validated Haxe byte range.
func FileOutputWriteValues(handle *FileOutput, values []int, position int, length int) int {
	if handle == nil || handle.file == nil || position < 0 || length < 0 || position+length > len(values) {
		fileThrow(os.ErrInvalid)
		return 0
	}
	raw := valuesToBytes(values[position : position+length])
	written := 0
	for written < len(raw) {
		count, err := handle.file.Write(raw[written:])
		if err != nil {
			fileThrow(err)
			return written
		}
		if count == 0 {
			fileThrow(io.ErrShortWrite)
			return written
		}
		written += count
	}
	return written
}

func FileOutputTell(handle *FileOutput) int {
	position, err := handle.Tell()
	if err != nil {
		fileThrow(err)
		return 0
	}
	return position
}

func FileOutputSeek(handle *FileOutput, offset int, whence int) {
	fileThrow(handle.Seek(offset, whence))
}

func FileOutputFlush(handle *FileOutput) {
	fileThrow(handle.Flush())
}

func FileOutputClose(handle *FileOutput) {
	fileThrow(handle.Close())
}
