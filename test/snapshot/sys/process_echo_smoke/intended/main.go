package main

import "snapshot/hxrt"

func main() {
	p := New_sys__io__Process(hxrt.StringFromLiteral("echo"), hxrt.NewArray(hxrt.StringFromLiteral("hi")), false)
	line := p.stdout.__hx_this.readLine()
	hxrt.Println(any(line))
	p.__hx_this.close()
}
