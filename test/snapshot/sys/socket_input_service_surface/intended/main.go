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
