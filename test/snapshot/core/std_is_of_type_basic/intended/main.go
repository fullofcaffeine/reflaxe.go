package main

import "snapshot/hxrt"

type I_Base interface {
}

type Base struct {
	__hx_this I_Base
}

func New_Base() *Base {
	self := &Base{}
	self.__hx_this = self
	return self
}

type I_Child interface {
}

type Child struct {
	*Base
	__hx_this I_Child
}

func New_Child() *Child {
	self := &Child{}
	self.Base = New_Base()
	self.Base.__hx_this = self
	self.__hx_this = self
	return self
}

func main() {
	child := New_Child().Base
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
	}(nil))
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
