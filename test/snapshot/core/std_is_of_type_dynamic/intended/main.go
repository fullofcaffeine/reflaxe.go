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
		switch hx_value.(type) {
		case *Base:
			return true
		case *Child:
			return true
		default:
			return false
		}
	}(any(d)))
	hxrt.Println(v)
	var v_1 any = any(func(hx_value any) bool {
		switch hx_value.(type) {
		case *Child:
			return true
		default:
			return false
		}
	}(any(d)))
	hxrt.Println(v_1)
	d = New_Base()
	var v_2 any = any(func(hx_value any) bool {
		switch hx_value.(type) {
		case *Child:
			return true
		default:
			return false
		}
	}(any(d)))
	hxrt.Println(v_2)
	d = []any{1, 2}
	var v_3 any = any(func(hx_value any) bool {
		switch hx_value.(type) {
		case []*Base:
			return true
		case []*Child:
			return true
		case []*haxe___Int64_____Int64:
			return true
		case []*string:
			return true
		case []any:
			return true
		case []bool:
			return true
		case []float64:
			return true
		case []int:
			return true
		default:
			return false
		}
	}(any(d)))
	hxrt.Println(v_3)
	d = 1
	var v_4 any = any(func(hx_value any) bool {
		switch hx_value.(type) {
		case []*Base:
			return true
		case []*Child:
			return true
		case []*haxe___Int64_____Int64:
			return true
		case []*string:
			return true
		case []any:
			return true
		case []bool:
			return true
		case []float64:
			return true
		case []int:
			return true
		default:
			return false
		}
	}(any(d)))
	hxrt.Println(v_4)
	d = nil
	var v_5 any = any((d != nil))
	hxrt.Println(v_5)
}
