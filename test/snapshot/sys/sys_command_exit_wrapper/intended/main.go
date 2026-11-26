package main

import "snapshot/hxrt"

func main() {
	code := Sys_command(hxrt.StringFromLiteral("sh"), hxrt.NewArray(hxrt.StringFromLiteral("-c"), hxrt.StringFromLiteral("printf 'wrapper-out\\n'; exit 7")))
	hxrt.SysExit(code)
}
