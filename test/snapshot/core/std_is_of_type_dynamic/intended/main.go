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
	var d any = New_Child()
	var v any = any(func(hx_value any) bool {
		switch hx_carrier := hx_value.(type) {
		case *Base:
			if hx_carrier == nil {
				return false
			}
			return true
		case *Child:
			if hx_carrier == nil {
				return false
			}
			return true
		default:
			return false
		}
	}(d))
	hxrt.Println(v)
	var v_1 any = any(func(hx_value any) bool {
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
	}(d))
	hxrt.Println(v_1)
	d = New_Base()
	var v_2 any = any(func(hx_value any) bool {
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
	}(d))
	hxrt.Println(v_2)
	d = hxrt.NewArray(1, 2)
	var v_3 any = any(func(hx_value any) bool {
		switch hx_value.(type) {
		case *hxrt.Array:
			return true
		default:
			return false
		}
	}(d))
	hxrt.Println(v_3)
	d = 1
	var v_4 any = any(func(hx_value any) bool {
		switch hx_value.(type) {
		case *hxrt.Array:
			return true
		default:
			return false
		}
	}(d))
	hxrt.Println(v_4)
	d = nil
	var v_5 any = any((d != nil))
	hxrt.Println(v_5)
}
