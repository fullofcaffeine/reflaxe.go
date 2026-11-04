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
	self := &Cell{}
	self.__hx_this = self
	self.value = value
	return self
}

func main() {
	source := New_Cell(nil)
	view := func(hx_structural_source_1 *Cell) hxrt.Anonymous {
		return hxrt.NewAnonymousView(hx_structural_source_1, func(name string) hxrt.AnonymousField {
			switch name {
			case "value":
				return hxrt.AnonymousField{Value: func(hx_value_2 any) *string {
					if hx_value_2 == nil {
						var hx_zero_3 *string
						return hx_zero_3
					}
					return hx_value_2.(*string)
				}(hx_structural_source_1.value), Found: true}
			}
			return hxrt.AnonymousField{}
		}, func(name string, incoming any) bool {
			switch name {
			case "value":
				hx_structural_source_1.value = func(hx_value_4 any) *string {
					if hx_value_4 == nil {
						var hx_zero_5 *string
						return hx_zero_5
					}
					return hx_value_4.(*string)
				}(incoming)
				return true
			}
			return false
		})
	}(source)
	var observed any = func(hx_obj_6 hxrt.Anonymous) *string {
		hx_field_7 := hxrt.AnonymousGet(hx_obj_6, "value")
		if hx_field_7 == nil {
			var hx_zero_8 *string
			return hx_zero_8
		}
		return hx_field_7.(*string)
	}(view)
	hxrt.Println(any(hxrt.AnyEqualsNull(observed)))
}
