package hxrt

import (
	"bytes"
	"compress/flate"
	"compress/zlib"
	"errors"
	"io"
	"sync"
)

// Zip flush values mirror haxe.zip.FlushMode. Go's public compression API
// exposes NO, SYNC, and FINISH semantics, but not zlib's dictionary-resetting
// FULL flush or block-boundary BLOCK mode.
const (
	ZipFlushNo = iota
	ZipFlushSync
	ZipFlushFull
	ZipFlushFinish
	ZipFlushBlock
)

// ZipCodecStep carries one bounded codec result without depending on generated
// haxe.io.Bytes layouts.
type ZipCodecStep struct {
	Values []int
	Read   int
	Done   bool
}

// ZipDeflateHandle retains one native zlib writer and any output that did not
// fit in the caller's previous destination.
type ZipDeflateHandle struct {
	output   bytes.Buffer
	writer   *zlib.Writer
	finished bool
	closed   bool
}

// ZipInflateHandle retains one live inflater, a caller-supplied input fragment,
// and decoded output bounded by the current or retained destination allowance.
// The handle itself implements flate.Reader so the decoder can pause at input
// and output boundaries without treating them as a terminal io.EOF.
type ZipInflateHandle struct {
	raw bool

	mu              sync.Mutex
	condition       *sync.Cond
	input           []byte
	inputPosition   int
	feeding         bool
	waitingForInput bool
	outputBudget    int
	outputInFlight  int
	output          bytes.Buffer
	finished        bool
	failure         error
	closed          bool
	stop            bool
	decoderStopped  chan struct{}
}

// ZipDeflateCreate starts one progressive zlib stream.
func ZipDeflateCreate(level int) *ZipDeflateHandle {
	if level < -1 || level > 9 {
		zipCodecThrow("invalid zlib compression level")
		return nil
	}
	handle := &ZipDeflateHandle{}
	writer, err := zlib.NewWriterLevel(&handle.output, level)
	if err != nil {
		zipCodecThrow(err.Error())
		return nil
	}
	handle.writer = writer
	return handle
}

// ZipDeflateExecute consumes source values and returns at most outputLimit
// bytes. Pending output is drained before more source is accepted, keeping the
// caller's read count safe for repeated execute calls.
func ZipDeflateExecute(handle *ZipDeflateHandle, values []int, outputLimit int, flushMode int) *ZipCodecStep {
	zipValidateDeflateHandle(handle)
	zipValidateOutputLimit(outputLimit)
	zipValidateFlushMode(flushMode)
	if outputLimit == 0 {
		return &ZipCodecStep{Values: []int{}}
	}
	if handle.output.Len() > 0 {
		return zipDrainDeflate(handle, 0, outputLimit)
	}
	if handle.finished {
		if len(values) > 0 {
			zipCodecThrow("zlib compressor is finished")
		}
		return &ZipCodecStep{Values: []int{}, Done: true}
	}

	read := 0
	if len(values) > 0 {
		var err error
		read, err = handle.writer.Write(zipValuesToBytes(values))
		if err != nil {
			zipCodecThrow(err.Error())
			return &ZipCodecStep{Values: []int{}}
		}
	}

	switch flushMode {
	case ZipFlushNo:
	case ZipFlushSync:
		if err := handle.writer.Flush(); err != nil {
			zipCodecThrow(err.Error())
			return &ZipCodecStep{Values: []int{}}
		}
	case ZipFlushFinish:
		if err := handle.writer.Close(); err != nil {
			zipCodecThrow(err.Error())
			return &ZipCodecStep{Values: []int{}}
		}
		handle.writer = nil
		handle.finished = true
	}

	return zipDrainDeflate(handle, read, outputLimit)
}

// ZipDeflateClose releases one progressive compressor. Closing an already
// closed handle is intentionally harmless; executing it afterward is not.
func ZipDeflateClose(handle *ZipDeflateHandle) {
	if handle == nil || handle.closed {
		return
	}
	handle.closed = true
	var err error
	if handle.writer != nil {
		err = handle.writer.Close()
		handle.writer = nil
	}
	handle.output.Reset()
	if err != nil {
		zipCodecThrow(err.Error())
	}
}

// ZipInflateCreate starts one progressive zlib or raw-DEFLATE stream.
func ZipInflateCreate(raw bool) *ZipInflateHandle {
	handle := &ZipInflateHandle{
		raw:            raw,
		decoderStopped: make(chan struct{}),
	}
	handle.condition = sync.NewCond(&handle.mu)
	go handle.runInflater()
	return handle
}

// ZipInflateExecute accepts partial compressed input and returns at most
// outputLimit decoded bytes. The native inflater pauses when the supplied
// fragment is exhausted and resumes when a later call installs more input.
func ZipInflateExecute(handle *ZipInflateHandle, values []int, outputLimit int, flushMode int) *ZipCodecStep {
	if handle == nil {
		zipCodecThrow("zlib inflater handle is null")
		return &ZipCodecStep{Values: []int{}}
	}
	zipValidateOutputLimit(outputLimit)
	zipValidateFlushMode(flushMode)

	handle.mu.Lock()
	if handle.closed {
		handle.mu.Unlock()
		zipCodecThrow("zlib inflater is closed")
		return &ZipCodecStep{Values: []int{}}
	}
	if outputLimit == 0 {
		handle.mu.Unlock()
		return &ZipCodecStep{Values: []int{}}
	}
	if handle.failure != nil {
		failure := handle.failure
		handle.mu.Unlock()
		zipCodecThrow(failure.Error())
		return &ZipCodecStep{Values: []int{}}
	}
	if handle.output.Len() > 0 {
		step := zipDrainInflateLocked(handle, 0, outputLimit)
		handle.mu.Unlock()
		return step
	}
	if handle.finished {
		handle.mu.Unlock()
		return &ZipCodecStep{Values: []int{}, Done: true}
	}
	if handle.feeding {
		handle.mu.Unlock()
		zipCodecThrow("concurrent zlib inflater execution is not supported")
		return &ZipCodecStep{Values: []int{}}
	}

	if len(values) > 0 {
		handle.input = zipValuesToBytes(values)
		handle.inputPosition = 0
		handle.feeding = true
		handle.waitingForInput = false
	}
	if handle.outputBudget == 0 && handle.outputInFlight == 0 {
		handle.outputBudget = outputLimit
	}
	handle.condition.Broadcast()
	for !handle.waitingForInput && !handle.finished && handle.failure == nil &&
		(handle.outputBudget > 0 || handle.outputInFlight > 0) {
		handle.condition.Wait()
	}
	read := 0
	if handle.feeding {
		read = handle.inputPosition
		handle.input = nil
		handle.inputPosition = 0
		handle.feeding = false
	}
	if handle.failure != nil {
		failure := handle.failure
		handle.mu.Unlock()
		zipCodecThrow(failure.Error())
		return &ZipCodecStep{Values: []int{}}
	}
	step := zipDrainInflateLocked(handle, read, outputLimit)
	handle.mu.Unlock()
	return step
}

// ZipInflateClose releases retained compressed and decoded buffers. It is
// idempotent; subsequent execute calls fail deterministically.
func ZipInflateClose(handle *ZipInflateHandle) {
	if handle == nil {
		return
	}
	handle.mu.Lock()
	if handle.closed {
		handle.mu.Unlock()
		return
	}
	handle.closed = true
	handle.stop = true
	handle.condition.Broadcast()
	decoderStopped := handle.decoderStopped
	handle.mu.Unlock()

	<-decoderStopped

	handle.mu.Lock()
	handle.input = nil
	handle.output.Reset()
	handle.mu.Unlock()
}

// Read lets compress/flate request bytes from the current execute fragment.
// When the fragment is exhausted it acknowledges the caller and waits for the
// next one instead of returning a stream-ending io.EOF.
func (handle *ZipInflateHandle) Read(values []byte) (int, error) {
	if len(values) == 0 {
		return 0, nil
	}
	handle.mu.Lock()
	defer handle.mu.Unlock()
	if err := handle.waitForInflaterInputLocked(); err != nil {
		return 0, err
	}
	read := copy(values, handle.input[handle.inputPosition:])
	handle.inputPosition += read
	return read, nil
}

// ReadByte keeps the handle on the exact-read path used by compress/flate and
// compress/zlib, preventing decoder lookahead from consuming trailing bytes.
func (handle *ZipInflateHandle) ReadByte() (byte, error) {
	handle.mu.Lock()
	defer handle.mu.Unlock()
	if err := handle.waitForInflaterInputLocked(); err != nil {
		return 0, err
	}
	value := handle.input[handle.inputPosition]
	handle.inputPosition++
	return value, nil
}

func (handle *ZipInflateHandle) waitForInflaterInputLocked() error {
	for handle.inputPosition >= len(handle.input) {
		handle.waitingForInput = true
		handle.condition.Broadcast()
		if handle.stop {
			return io.EOF
		}
		handle.condition.Wait()
	}
	handle.waitingForInput = false
	return nil
}

func (handle *ZipInflateHandle) runInflater() {
	defer close(handle.decoderStopped)

	var reader io.ReadCloser
	if handle.raw {
		reader = flate.NewReader(handle)
	} else {
		resolved, err := zlib.NewReader(handle)
		if err != nil {
			handle.finishInflater(err)
			return
		}
		reader = resolved
	}

	buffer := make([]byte, 32*1024)
	for {
		handle.mu.Lock()
		for handle.outputBudget == 0 && !handle.stop {
			handle.condition.Wait()
		}
		if handle.stop {
			handle.mu.Unlock()
			_ = reader.Close()
			return
		}
		request := min(len(buffer), handle.outputBudget)
		handle.outputBudget -= request
		handle.outputInFlight += request
		handle.mu.Unlock()

		read, err := reader.Read(buffer[:request])

		handle.mu.Lock()
		handle.outputInFlight -= request
		handle.outputBudget += request - read
		if read > 0 {
			_, _ = handle.output.Write(buffer[:read])
		}
		if err != nil {
			if err == io.EOF {
				handle.finished = true
			} else {
				handle.failure = err
			}
		}
		handle.condition.Broadcast()
		if err != nil {
			handle.mu.Unlock()
			_ = reader.Close()
			return
		}
		handle.mu.Unlock()
	}
}

func (handle *ZipInflateHandle) finishInflater(err error) {
	handle.mu.Lock()
	handle.failure = err
	handle.condition.Broadcast()
	handle.mu.Unlock()
}

func zipDrainDeflate(handle *ZipDeflateHandle, read int, outputLimit int) *ZipCodecStep {
	values := zipBytesToValues(handle.output.Next(min(outputLimit, handle.output.Len())))
	return &ZipCodecStep{
		Values: values,
		Read:   read,
		Done:   handle.finished && handle.output.Len() == 0,
	}
}

func zipDrainInflateLocked(handle *ZipInflateHandle, read int, outputLimit int) *ZipCodecStep {
	values := zipBytesToValues(handle.output.Next(min(outputLimit, handle.output.Len())))
	return &ZipCodecStep{
		Values: values,
		Read:   read,
		Done:   handle.finished && handle.output.Len() == 0,
	}
}

func zipValidateDeflateHandle(handle *ZipDeflateHandle) {
	if handle == nil {
		zipCodecThrow("zlib compressor handle is null")
		return
	}
	if handle.closed {
		zipCodecThrow("zlib compressor is closed")
	}
}

func zipValidateOutputLimit(outputLimit int) {
	if outputLimit < 0 {
		zipCodecThrow("zlib output limit must not be negative")
	}
}

func zipValidateFlushMode(flushMode int) {
	switch flushMode {
	case ZipFlushNo, ZipFlushSync, ZipFlushFinish:
		return
	case ZipFlushFull:
		zipCodecThrow("haxe.zip.FlushMode.FULL is not supported by Go's standard compression API")
	case ZipFlushBlock:
		zipCodecThrow("haxe.zip.FlushMode.BLOCK is not supported by Go's standard compression API")
	default:
		zipCodecThrow("invalid haxe.zip flush mode")
	}
}

func zipCodecThrow(message string) {
	Throw(StringFromLiteral(message))
}

// ZipCompress performs one complete zlib compression operation for staged
// haxe.zip.Compress. Public level policy and Haxe Bytes conversion stay in Haxe.
func ZipCompress(values []int, level int) []int {
	if level < -1 || level > 9 {
		Throw(errors.New("invalid zlib compression level"))
		return []int{}
	}

	var output bytes.Buffer
	writer, err := zlib.NewWriterLevel(&output, level)
	if err != nil {
		Throw(err)
		return []int{}
	}
	if _, err := writer.Write(zipValuesToBytes(values)); err != nil {
		_ = writer.Close()
		Throw(err)
		return []int{}
	}
	if err := writer.Close(); err != nil {
		Throw(err)
		return []int{}
	}
	return zipBytesToValues(output.Bytes())
}

// ZipUncompress performs one complete zlib or raw-DEFLATE expansion. The
// caller supplies its buffer-size policy; no generated Haxe Bytes layout crosses
// this runtime package boundary.
func ZipUncompress(values []int, raw bool, bufferSize int) []int {
	if bufferSize <= 0 {
		Throw(errors.New("zlib buffer size must be positive"))
		return []int{}
	}

	input := bytes.NewReader(zipValuesToBytes(values))
	var reader io.ReadCloser
	if raw {
		reader = flate.NewReader(input)
	} else {
		resolved, err := zlib.NewReader(input)
		if err != nil {
			Throw(err)
			return []int{}
		}
		reader = resolved
	}

	buffer := make([]byte, bufferSize)
	var output bytes.Buffer
	for {
		read, err := reader.Read(buffer)
		if read > 0 {
			_, _ = output.Write(buffer[:read])
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			_ = reader.Close()
			Throw(err)
			return []int{}
		}
	}
	if err := reader.Close(); err != nil {
		Throw(err)
		return []int{}
	}
	return zipBytesToValues(output.Bytes())
}

func zipValuesToBytes(values []int) []byte {
	converted := make([]byte, len(values))
	for index, value := range values {
		converted[index] = byte(value)
	}
	return converted
}

func zipBytesToValues(values []byte) []int {
	converted := make([]int, len(values))
	for index, value := range values {
		converted[index] = int(value)
	}
	return converted
}
