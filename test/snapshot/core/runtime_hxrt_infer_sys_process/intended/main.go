package main

import "snapshot/hxrt"

func main() {
	cwd := hxrt.StdString(hxrt.SysGetCwd())
	process := New_sys__io__Process(hxrt.StringFromLiteral("echo"), hxrt.NewArray(hxrt.StringFromLiteral("ok")), false)
	process.__hx_this.close()
	hxrt.Println(any(cwd))
}
