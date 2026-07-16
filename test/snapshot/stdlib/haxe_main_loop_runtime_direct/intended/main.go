package main

import "snapshot/hxrt"

func main() {
	defer hxrt.ThreadWaitForAll()
	ran := false
	var event *haxe__MainEvent = nil
	event = haxe__MainLoop_add(func() {
		ran = true
		hxrt.Println(any(hxrt.StringFromLiteral("mainloop.add=ran")))
		event.stop()
	}, 0)
	haxe__EntryPoint_run()
	var v any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("mainloop.after="), hxrt.StdString(ran)))
	hxrt.Println(v)
	haxe__Timer_delay(func() {
		hxrt.Println(any(hxrt.StringFromLiteral("timer.delay=ran")))
	}, 1)
	haxe__EntryPoint_run()
}
