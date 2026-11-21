package main

import "snapshot/hxrt"

func Sys_getChar(echo bool) int {
	value := hxrt.SysReadCharValue()
	if value < 0 {
		hxrt.Throw(haxe__Exception_thrown(New_haxe__io__Eof()))
	}
	if echo {
		New_sys__io__FileOutput(hxrt.SysStdout()).__hx_this.writeByte(value)
	}
	return value
}
