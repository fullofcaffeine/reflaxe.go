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
