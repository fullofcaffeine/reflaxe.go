package main

import "snapshot/hxrt"

type I_Base interface {
	name() *string
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

func (self *Base) name() *string {
	return hxrt.StringFromLiteral("base")
}

func (self *Base) HxClassName() string {
	return "Base"
}

func main() {
	var value any = New_Mid()
	mid := hx_recover_class_Mid(value)
	var v any = any(mid.name())
	hxrt.Println(v)
	hxrt.TryCatch(func() {
		hxrt.Throw(haxe__Exception_thrown(hxrt.StringFromLiteral("boom")))
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
			hxrt.Println(any(message))
		} else {
			hxrt.Throw(hx_caught_1)
		}
	})
}

type I_Mid interface {
	name() *string
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

func (self *Mid) name() *string {
	return hxrt.StringFromLiteral("mid")
}

func (self *Mid) HxClassName() string {
	return "Mid"
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
