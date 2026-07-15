package main

import "snapshot/hxrt"

func main() {
	var v any = any(!hxrt.StringEqualStringPtr(hxrt.StdString(hxrt.SysGetCwd()), nil))
	hxrt.Println(v)
	var v_1 any = any(hxrt.StdString(hxrt.SysSystemName()))
	hxrt.Println(v_1)
}
