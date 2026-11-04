package main

import "snapshot/hxrt"

func main() {
	cwd := hxrt.StdString(hxrt.SysGetCwd())
	if hxrt.StringLengthStringPtr(cwd) == 0 {
		hxrt.Throw(haxe__Exception_thrown(hxrt.StringFromLiteral("expected a current working directory")))
	}
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
	case *haxe__Exception:
		return source
	case *haxe__ValueException:
		return source.haxe__Exception
	default:
		return value.(*haxe__Exception)
	}
}
