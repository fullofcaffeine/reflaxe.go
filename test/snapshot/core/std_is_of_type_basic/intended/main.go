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
	hx_init_Base(self.Base)
	return self
}

func main() {
	child := func(hx_value *Child) *Base {
		if hx_value == nil {
			return nil
		}
		return hx_value.Base
	}(New_Child())
	base := New_Base()
	var v any = any(func(hx_value *Base) bool {
		if hx_value == nil {
			return false
		}
		_, ok := hx_value.__hx_this.(*Child)
		return ok
	}(child))
	hxrt.Println(v)
	var v_1 any = any((child != nil))
	hxrt.Println(v_1)
	var v_2 any = any(func(hx_value *Base) bool {
		if hx_value == nil {
			return false
		}
		_, ok := hx_value.__hx_this.(*Child)
		return ok
	}(base))
	hxrt.Println(v_2)
	var v_3 any = any(func(hx_value any) bool {
		switch hx_carrier := hx_value.(type) {
		case *Base:
			if hx_carrier == nil {
				return false
			}
			_, hx_ok := hx_carrier.__hx_this.(*Child)
			return hx_ok
		case *Child:
			if hx_carrier == nil {
				return false
			}
			return true
		default:
			return false
		}
	}(hxrt.ObjectSource(hxrt.ObjectSource(nil))))
	hxrt.Println(v_3)
	var v_4 any = any(true)
	hxrt.Println(v_4)
	var v_5 any = any(true)
	hxrt.Println(v_5)
	var v_6 any = any(false)
	hxrt.Println(v_6)
	var v_7 any = any((hxrt.StringFromLiteral("x") != nil))
	hxrt.Println(v_7)
	var v_8 any = any(true)
	hxrt.Println(v_8)
	var v_9 any = any(false)
	hxrt.Println(v_9)
}
