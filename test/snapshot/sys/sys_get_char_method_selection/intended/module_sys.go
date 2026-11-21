package main

import "snapshot/hxrt"

func Sys_environment() *haxe__ds__StringMap {
	environment := New_haxe__ds__StringMap()
	_g := 0
	_g1 := hxrt.SysEnvironmentEntries()
	for _g < len(_g1) {
		entry := _g1[_g]
		_g = int(int32((_g + 1)))
		environment.__hx_this.set(entry.Key, entry.Value)
	}
	return environment
}

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
