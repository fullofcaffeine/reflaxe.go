package main

import "snapshot/hxrt"

type I_Base interface {
}

type Base struct {
	__hx_this I_Base
}

func (self *Base) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_Base() *Base {
	return hx_init_Base(hx_alloc_Base(nil))
}

func hx_alloc_Base(owner any) *Base {
	self := &Base{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_Base)
	}
	return self
}

func hx_init_Base(self *Base) *Base {
	return self
}

func (self *Base) HxClassName() string {
	return "Base"
}

func main() {
	var other any = New_Other()
	hxrt.TryCatch(func() {
		func(hx_cast_6 any) *Mid {
			if !hxrt.AnyEqualsNull(hx_cast_6) && !hx_is_nominal_Mid(hx_cast_6) {
				hxrt.Throw(haxe__Exception_thrown(hxrt.StringFromLiteral("Class cast error")))
			}
			return hx_recover_class_Mid(hx_cast_6)
		}(other)
		hxrt.Println(any(hxrt.StringFromLiteral("not reached")))
	}, func(hx_caught_1 any) {
		hx_exception_2 := haxe__Exception_caught(hx_caught_1)
		hx_payload_3 := haxe__Exception_catchValue(hx_caught_1)
		_ = hx_exception_2
		if func(hx_value any) bool {
			switch hx_value.(type) {
			case *string:
				return true
			default:
				return false
			}
		}(hx_payload_3) {
			var message *string = func(hx_value_4 any) *string {
				if hx_value_4 == nil {
					var hx_zero_5 *string
					return hx_zero_5
				}
				return hx_value_4.(*string)
			}(hx_payload_3)
			hxrt.Println(any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("caught: "), message)))
		} else {
			hxrt.Throw(hx_caught_1)
		}
	})
	mid := New_Mid()
	base := mid
	var v any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("upcast: "), hxrt.StdString(hxrt.HaxeEqual(base, mid))))
	hxrt.Println(v)
}

type I_Mid interface {
}

type Mid struct {
	*Base
	__hx_this I_Mid
}

func (self *Mid) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_Mid() *Mid {
	return hx_init_Mid(hx_alloc_Mid(nil))
}

func hx_alloc_Mid(owner any) *Mid {
	self := &Mid{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_Mid)
	}
	self.Base = hx_alloc_Base(self.__hx_this)
	return self
}

func hx_init_Mid(self *Mid) *Mid {
	hx_init_Base(self.Base)
	return self
}

func (self *Mid) HxClassName() string {
	return "Mid"
}

type I_Other interface {
}

type Other struct {
	__hx_this I_Other
}

func (self *Other) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_Other() *Other {
	return hx_init_Other(hx_alloc_Other(nil))
}

func hx_alloc_Other(owner any) *Other {
	self := &Other{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_Other)
	}
	return self
}

func hx_init_Other(self *Other) *Other {
	return self
}

func (self *Other) HxClassName() string {
	return "Other"
}

func hx_recover_class_Mid(value any) *Mid {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *Mid:
		return source
	default:
		return value.(*Mid)
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

func hx_is_nominal_Mid(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *Mid:
		return (source != nil)
	default:
		return false
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
