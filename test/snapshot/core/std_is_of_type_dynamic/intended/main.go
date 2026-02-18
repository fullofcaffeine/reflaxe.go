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
	_ = d
	hxrt.Println(func(hx_value any) bool {
		switch hx_type := hx_value.(type) {
		case *Base:
			_ = hx_type
			return true
		case *Child:
			_ = hx_type
			return true
		default:
			_ = hx_type
			return false
		}
	}(any(d)))
	hxrt.Println(func(hx_value any) bool {
		switch hx_type := hx_value.(type) {
		case *Child:
			_ = hx_type
			return true
		default:
			_ = hx_type
			return false
		}
	}(any(d)))
	d = New_Base()
	hxrt.Println(func(hx_value any) bool {
		switch hx_type := hx_value.(type) {
		case *Child:
			_ = hx_type
			return true
		default:
			_ = hx_type
			return false
		}
	}(any(d)))
	d = []any{1, 2}
	hxrt.Println(func(hx_value any) bool {
		switch hx_type := hx_value.(type) {
		case []*Base:
			_ = hx_type
			return true
		case []*Child:
			_ = hx_type
			return true
		case []*string:
			_ = hx_type
			return true
		case []any:
			_ = hx_type
			return true
		case []bool:
			_ = hx_type
			return true
		case []float64:
			_ = hx_type
			return true
		case []int:
			_ = hx_type
			return true
		default:
			_ = hx_type
			return false
		}
	}(any(d)))
	d = 1
	hxrt.Println(func(hx_value any) bool {
		switch hx_type := hx_value.(type) {
		case []*Base:
			_ = hx_type
			return true
		case []*Child:
			_ = hx_type
			return true
		case []*string:
			_ = hx_type
			return true
		case []any:
			_ = hx_type
			return true
		case []bool:
			_ = hx_type
			return true
		case []float64:
			_ = hx_type
			return true
		case []int:
			_ = hx_type
			return true
		default:
			_ = hx_type
			return false
		}
	}(any(d)))
	d = nil
	hxrt.Println((d != nil))
}
