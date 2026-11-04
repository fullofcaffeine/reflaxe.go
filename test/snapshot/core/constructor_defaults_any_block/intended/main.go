package main

import "snapshot/hxrt"

type I_Base interface {
}

type Base struct {
	__hx_this I_Base
	state     *string
}

func (self *Base) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_Base(label *string, number any, suffix *string) *Base {
	return hx_init_Base(hx_alloc_Base(nil), label, number, suffix)
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

func hx_init_Base(self *Base, label *string, number any, suffix *string) *Base {
	self.state = hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(label, hxrt.StringFromLiteral(":")), hxrt.StdString(number.(int))), hxrt.StringFromLiteral(":")), hxrt.StdString(suffix))
	return self
}

type I_Child interface {
}

type Child struct {
	*Base
	__hx_this I_Child
}

func (self *Child) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_Child() *Child {
	return hx_init_Child(hx_alloc_Child(nil))
}

func hx_alloc_Child(owner any) *Child {
	self := &Child{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_Child)
	}
	self.Base = hx_alloc_Base(self.__hx_this)
	return self
}

func hx_init_Child(self *Child) *Child {
	hx_init_Base(self.Base, hxrt.StringFromLiteral("x"), 4, nil)
	return self
}

func erasedValue(value *Child) any {
	return value
}

func main() {
	value := New_Child()
	var erased any = value
	var v any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("state="), value.state))
	hxrt.Println(v)
	var v_1 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("identity="), hxrt.StdString((hx_recover_class_0(erased) == value))))
	hxrt.Println(v_1)
	var v_2 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("call-identity="), hxrt.StdString((func() *Child {
		var this1 any = erasedValue(value)
		return hx_recover_class_0(this1)
	}() == value))))
	hxrt.Println(v_2)
	base := func(hx_value *Child) *Base {
		if hx_value == nil {
			return nil
		}
		return hx_value.Base
	}(value)
	var v_3 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("base-identity="), hxrt.StdString(hxrt.HaxeEqual(base, value))))
	hxrt.Println(v_3)
}

func hx_recover_class_0(value any) *Child {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *Child:
		return source
	default:
		return value.(*Child)
	}
}
