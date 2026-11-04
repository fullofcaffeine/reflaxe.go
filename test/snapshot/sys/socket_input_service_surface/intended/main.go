package main

import "snapshot/hxrt"

type hxrt__TypeClassValue struct {
	name *string
}

type hxrt__TypeEnumValue struct {
	name *string
}

func consume(socket *sys__net__Socket) *string {
	line := socket.input.__hx_this.readLine()
	lower := hxrt.StringToLowerCaseStringPtr(line)
	var parsed any = Std_parseInt(hxrt.StringFromLiteral("42"))
	first := socket.input.__hx_this.readByte()
	buffer := haxe__io__Bytes_alloc(4)
	read := socket.input.__hx_this.readBytes(buffer, 0, buffer.length)
	rest := socket.input.__hx_this.readAll(nil).__hx_this.toString()
	return hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(lower, hxrt.StringFromLiteral(":")), parsed), hxrt.StringFromLiteral(":")), first), hxrt.StringFromLiteral(":")), read), hxrt.StringFromLiteral(":")), rest)
}

func main() {
	hxrt.Println(any(hxrt.StringFromLiteral("socket-input-service-surface")))
}

func hx_recover_class_0(value any) *sys__net__Socket {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *sys__net__Socket:
		return source
	default:
		return value.(*sys__net__Socket)
	}
}

func hx_recover_class_1(value any) *haxe__io__Eof {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *haxe__io__Eof:
		return source
	default:
		return value.(*haxe__io__Eof)
	}
}

func hx_recover_class_2(value any) *haxe__ValueException {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *haxe__ValueException:
		return source
	default:
		return value.(*haxe__ValueException)
	}
}

func hx_recover_class_3(value any) *haxe__Exception {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *haxe__exceptions__NotImplementedException:
		return source.haxe__exceptions__PosException.haxe__Exception
	case *haxe__Exception:
		return source
	case *haxe__exceptions__PosException:
		return source.haxe__Exception
	case *haxe__ValueException:
		return source.haxe__Exception
	default:
		return value.(*haxe__Exception)
	}
}

func hx_is_nominal_0(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *haxe__io__Eof:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_1(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *haxe__exceptions__NotImplementedException:
		return (source != nil)
	case *haxe__Exception:
		return (source != nil)
	case *haxe__exceptions__PosException:
		return (source != nil)
	case *haxe__ValueException:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_2(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *haxe__ValueException:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_3(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *haxe___Int64_____Int64:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_4(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *haxe__exceptions__NotImplementedException:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_5(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *haxe__exceptions__NotImplementedException:
		return (source != nil)
	case *haxe__exceptions__PosException:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_6(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *haxe__io__Bytes:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_7(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *haxe__io__BytesBuffer:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_8(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *sys__net__SocketInput:
		return (source != nil)
	case *haxe__io__Input:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_9(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *sys__net__SocketOutput:
		return (source != nil)
	case *haxe__io__Output:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_10(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *sys__net__Host:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_11(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *sys__net__Socket:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_12(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *sys__net__SocketInput:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_13(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *sys__net__SocketOutput:
		return (source != nil)
	default:
		return false
	}
}
