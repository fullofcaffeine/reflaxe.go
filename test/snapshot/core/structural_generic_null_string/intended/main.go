package main

import "snapshot/hxrt"

type I_Cell interface {
}

type Cell struct {
	__hx_this I_Cell
	value     any
}

func (self *Cell) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_Cell(value any) *Cell {
	return hx_init_Cell(hx_alloc_Cell(nil), value)
}

func hx_alloc_Cell(owner any) *Cell {
	self := &Cell{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_Cell)
	}
	return self
}

func hx_init_Cell(self *Cell, value any) *Cell {
	self.value = value
	return self
}

func main() {
	source := New_Cell(nil)
	view := func(hx_structural_source_1 *Cell) hxrt.Anonymous {
		return hxrt.NewAnonymousView(hx_structural_source_1, func(name string) hxrt.AnonymousField {
			switch name {
			case "value":
				return hxrt.AnonymousField{Value: hx_structural_source_1.value, Found: true}
			}
			return hxrt.AnonymousField{}
		}, func(name string, incoming any) bool {
			switch name {
			case "value":
				hx_structural_source_1.value = incoming
				return true
			}
			return false
		})
	}(source)
	var observed any = func(hx_value_2 any) *string {
		if hx_value_2 == nil {
			var hx_zero_3 *string
			return hx_zero_3
		}
		return hx_value_2.(*string)
	}(hxrt.AnonymousGet(view, "value"))
	hxrt.Println(any(hxrt.AnyEqualsNull(observed)))
}
