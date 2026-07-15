package main

import "snapshot/hxrt"

func main() {
	code := hxrt.SysCommand(hxrt.StringFromLiteral("sh"), []*string{hxrt.StringFromLiteral("-c"), hxrt.StringFromLiteral("printf 'wrapper-out\\n'; exit 7")})
	hxrt.SysExit(code)
}
