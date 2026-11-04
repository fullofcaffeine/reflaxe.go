package main

import "snapshot/hxrt"

func main() {
	defer hxrt.ThreadWaitForAll()
	started := New_sys__thread__Lock()
	sys__thread__Thread_create(hxrt.NewFunction(func() {
		started.__hx_this.release()
		New_sys__thread__Lock().__hx_this.wait(0.02)
		hxrt.Println(any(hxrt.StringFromLiteral("child-before-throw")))
		hxrt.Throw(haxe__Exception_thrown(hxrt.StringFromLiteral("child-failure")))
	}))
	started.__hx_this.wait(nil)
	hxrt.Println(any(hxrt.StringFromLiteral("main-survived")))
}

func hx_recover_class_0(value any) *haxe__ValueException {
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

func hx_recover_class_1(value any) *haxe__Exception {
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

func hx_is_nominal_0(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *haxe__ValueException:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_1(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *sys__thread__NoEventLoopException:
		return (source != nil)
	case *haxe__Exception:
		return (source != nil)
	case *haxe__ValueException:
		return (source != nil)
	default:
		return false
	}
}
