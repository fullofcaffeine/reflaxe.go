package main

import "snapshot/hxrt"

type I_Item interface {
}

type Item struct {
	__hx_this I_Item
	n         int
}

func (self *Item) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_Item() *Item {
	return hx_init_Item(hx_alloc_Item(nil))
}

func hx_alloc_Item(owner any) *Item {
	self := &Item{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_Item)
	}
	return self
}

func hx_init_Item(self *Item) *Item {
	self.n = 4
	return self
}

func identity(value any) any {
	return value
}

func main() {
	item := New_Item()
	var v any = any(same(hx_recover_class_0(identity(item)), item))
	hxrt.Println(v)
	var v_1 any = any(hx_recover_class_0(identity(item)).n)
	hxrt.Println(v_1)
	hx_obj_1 := map[string]any{}
	hx_obj_1["n"] = 7
	object := hxrt.AnonymousFromMap(hx_obj_1)
	var v_2 any = any(same(hx_structural_recover(identity(object)), object))
	hxrt.Println(v_2)
	var v_3 any = any(func(hx_value_2 any) int {
		if hx_value_2 == nil {
			var hx_zero_3 int
			return hx_zero_3
		}
		return hx_value_2.(int)
	}(hxrt.AnonymousGet(hx_structural_recover(identity(object)), "n")))
	hxrt.Println(v_3)
}

func same(a any, b any) bool {
	return hxrt.HaxeEqual(a, b)
}

func hx_structural_recover(value any) hxrt.Anonymous {
	switch source := value.(type) {
	case *Item:
		return func(hx_structural_source_1 *Item) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_1, func(name string) hxrt.AnonymousField {
				switch name {
				case "n":
					return hxrt.AnonymousField{Value: hx_structural_source_1.n, Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "n":
					hx_structural_source_1.n = func(hx_value_2 any) int {
						if hx_value_2 == nil {
							var hx_zero_3 int
							return hx_zero_3
						}
						return hx_value_2.(int)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe___Int64_____Int64:
		return func(hx_structural_source_4 *haxe___Int64_____Int64) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_4, func(name string) hxrt.AnonymousField {
				switch name {
				case "high":
					return hxrt.AnonymousField{Value: hx_structural_source_4.high, Found: true}
				case "low":
					return hxrt.AnonymousField{Value: hx_structural_source_4.low, Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "high":
					hx_structural_source_4.high = func(hx_value_5 any) int {
						if hx_value_5 == nil {
							var hx_zero_6 int
							return hx_zero_6
						}
						return hx_value_5.(int)
					}(incoming)
					return true
				case "low":
					hx_structural_source_4.low = func(hx_value_7 any) int {
						if hx_value_7 == nil {
							var hx_zero_8 int
							return hx_zero_8
						}
						return hx_value_7.(int)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	default:
		return hxrt.AnonymousFromValue(value)
	}
}

func hx_recover_class_0(value any) *Item {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *Item:
		return source
	default:
		return value.(*Item)
	}
}
