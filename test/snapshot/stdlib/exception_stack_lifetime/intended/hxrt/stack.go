package hxrt

import "runtime"

// StackFrame is the small target-owned carrier used by the optional Go-native
// diagnostic stack capture path. It intentionally stores Go runtime frame data;
// staged std code decides how much of it becomes Haxe CallStack data.
type StackFrame struct {
	Function *string
	File     *string
	Line     int
}

// NativeStackCapture returns best-effort Go runtime frames for diagnostics.
// `skip` is counted after this helper, so callers can hide their own wrapper
// frames without depending on runtime.Callers internals.
func NativeStackCapture(skip int) []*StackFrame {
	if skip < 0 {
		skip = 0
	}
	pcs := make([]uintptr, 64)
	n := runtime.Callers(skip+2, pcs)
	if n <= 0 {
		return []*StackFrame{}
	}
	frames := runtime.CallersFrames(pcs[:n])
	out := make([]*StackFrame, 0, n)
	for {
		frame, more := frames.Next()
		out = append(out, &StackFrame{
			Function: StringFromLiteral(frame.Function),
			File:     StringFromLiteral(frame.File),
			Line:     frame.Line,
		})
		if !more {
			break
		}
	}
	return out
}

// NativeStackIsFrameSlice lets staged std validate the opaque NativeStackTrace
// carrier before casting it back to Go frames.
func NativeStackIsFrameSlice(value any) bool {
	_, ok := value.([]*StackFrame)
	return ok
}
