package main

import "snapshot/hxrt"

func main() {
	cwd := hxrt.StdString(hxrt.SysGetCwd())
	if hxrt.StringLengthStringPtr(cwd) == 0 {
		hxrt.Throw(hxrt.StringFromLiteral("expected a current working directory"))
	}
}
