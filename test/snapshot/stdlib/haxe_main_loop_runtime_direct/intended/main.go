package main

import "snapshot/hxrt"

func main() {
	defer hxrt.ThreadWaitForAll()
	ran := false
	var event *haxe__MainEvent = nil
	event = haxe__MainLoop_add(hxrt.NewFunction(func() {
		ran = true
		hxrt.Println(any(hxrt.StringFromLiteral("mainloop.add=ran")))
		event.__hx_this.stop()
	}), 0)
	haxe__EntryPoint_run()
	var v any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("mainloop.after="), hxrt.StdString(ran)))
	hxrt.Println(v)
	haxe__Timer_delay(hxrt.NewFunction(func() {
		hxrt.Println(any(hxrt.StringFromLiteral("timer.delay=ran")))
	}), 1)
	haxe__EntryPoint_run()
}

func hx_recover_class_0(value any) *haxe__MainEvent {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *haxe__MainEvent:
		return source
	default:
		return value.(*haxe__MainEvent)
	}
}

func hx_recover_class_1(value any) *haxe__ValueException {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *haxe__ValueException:
		return source
	default:
		return value.(*haxe__ValueException)
	}
}

func hx_recover_class_2(value any) *haxe__Exception {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *sys__thread__NoEventLoopException:
		return source.haxe__Exception
	case *haxe__Exception:
		return source
	case *haxe__ValueException:
		return source.haxe__Exception
	default:
		return value.(*haxe__Exception)
	}
}
