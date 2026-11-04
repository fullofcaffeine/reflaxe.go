package main

import "snapshot/hxrt"

func main() {
	hxrt.TryCatch(func() {
		New_haxe__ds__WeakMap()
		hxrt.Println(any(hxrt.StringFromLiteral("constructed")))
	}, func(hx_caught_1 any) {
		hx_exception_2 := haxe__Exception_caught(hx_caught_1)
		hx_payload_3 := haxe__Exception_catchValue(hx_caught_1)
		_ = hx_payload_3
		if hx_is_nominal_0(hx_exception_2) {
			var err *haxe__exceptions__NotImplementedException = hx_recover_class_0(hx_exception_2)
			var v any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("not_impl="), err.__hx_this.get_message()))
			hxrt.Println(v)
			var v_1 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("typed="), hxrt.StdString(hx_is_nominal_0(err))))
			hxrt.Println(v_1)
		} else {
			hxrt.Throw(hx_caught_1)
		}
	})
}

type I__Main__WeakMapKey interface {
}

type _Main__WeakMapKey struct {
	__hx_this I__Main__WeakMapKey
}

func (self *_Main__WeakMapKey) HxFunctionReceiver() any {
	return self.__hx_this
}

func New__Main__WeakMapKey() *_Main__WeakMapKey {
	return hx_init__Main__WeakMapKey(hx_alloc__Main__WeakMapKey(nil))
}

func hx_alloc__Main__WeakMapKey(owner any) *_Main__WeakMapKey {
	self := &_Main__WeakMapKey{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I__Main__WeakMapKey)
	}
	return self
}

func hx_init__Main__WeakMapKey(self *_Main__WeakMapKey) *_Main__WeakMapKey {
	return self
}

func hx_recover_class_0(value any) *haxe__exceptions__NotImplementedException {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *haxe__exceptions__NotImplementedException:
		return source
	default:
		return value.(*haxe__exceptions__NotImplementedException)
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
	case *haxe__exceptions__PosException:
		return source.haxe__Exception
	case *haxe__exceptions__NotImplementedException:
		return source.haxe__exceptions__PosException.haxe__Exception
	case *haxe__ValueException:
		return source.haxe__Exception
	case *haxe__Exception:
		return source
	default:
		return value.(*haxe__Exception)
	}
}

func hx_is_nominal_0(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *haxe__exceptions__NotImplementedException:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_1(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *haxe__ValueException:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_2(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *haxe__exceptions__PosException:
		return (source != nil)
	case *haxe__exceptions__NotImplementedException:
		return (source != nil)
	case *haxe__ValueException:
		return (source != nil)
	case *haxe__Exception:
		return (source != nil)
	default:
		return false
	}
}
