package hxrt

import (
	"bytes"
	"compress/flate"
	"compress/zlib"
	"errors"
	"io"
)

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
