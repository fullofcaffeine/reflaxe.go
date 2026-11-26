package main

import "examples_pulseforge_metal/hxrt"

func hasArg(flag *string) bool {
	_g := 0
	_g1 := Sys_args()
	for _g < _g1.Len() {
		arg := func(hx_value_1 any) *string {
			if hx_value_1 == nil {
				var hx_zero_2 *string
				return hx_zero_2
			}
			return hx_value_1.(*string)
		}(_g1.Get(_g))
		_g = int(int32((_g + 1)))
		if hxrt.StringEqualStringPtr(arg, flag) {
			return true
		}
	}
	return false
}

func main() {
	var runtime app__runtime__PulseRuntime = app__runtime__RuntimeFactory_create()
	if hasArg(hxrt.StringFromLiteral("--scripted")) {
		var v any = any(Harness_run(runtime))
		hxrt.Println(v)
	} else {
		InteractiveCli_run(runtime)
	}
}

func hx_recover_class_app__core__PulseAlert(value any) *app__core__PulseAlert {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *app__core__PulseAlert:
		return source
	default:
		return value.(*app__core__PulseAlert)
	}
}

func hx_recover_class_app__core__PulseEnrichedEvent(value any) *app__core__PulseEnrichedEvent {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *app__core__PulseEnrichedEvent:
		return source
	default:
		return value.(*app__core__PulseEnrichedEvent)
	}
}

func hx_recover_class_app__core__PulseEvent(value any) *app__core__PulseEvent {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *app__core__PulseEvent:
		return source
	default:
		return value.(*app__core__PulseEvent)
	}
}

func hx_recover_class_app__core__PulseIngressFrame(value any) *app__core__PulseIngressFrame {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *app__core__PulseIngressFrame:
		return source
	default:
		return value.(*app__core__PulseIngressFrame)
	}
}

func hx_recover_class_app__core__PulseSourceAggregate(value any) *app__core__PulseSourceAggregate {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *app__core__PulseSourceAggregate:
		return source
	default:
		return value.(*app__core__PulseSourceAggregate)
	}
}

func hx_recover_class_haxe__Exception(value any) *haxe__Exception {
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

func hx_recover_class_haxe__ValueException(value any) *haxe__ValueException {
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

func hx_is_nominal_haxe__Exception(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *haxe__Exception:
		return (source != nil)
	case *haxe__ValueException:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_haxe__ValueException(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *haxe__ValueException:
		return (source != nil)
	default:
		return false
	}
}
