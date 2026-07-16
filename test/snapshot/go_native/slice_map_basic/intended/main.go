package main

import "snapshot/hxrt"

func main() {
	s := New_go___Slice()
	s.push(1)
	s.push(2)
	s.push(3)
	s.set(1, 7)
	var v any = any(func(hx_value_1 any) int {
		if hx_value_1 == nil {
			var hx_zero_2 int
			return hx_zero_2
		}
		return hx_value_1.(int)
	}(s.get_length()))
	hxrt.Println(v)
	var v_1 any = any(func(hx_value_3 any) int {
		if hx_value_3 == nil {
			var hx_zero_4 int
			return hx_zero_4
		}
		return hx_value_3.(int)
	}(s.get(1)))
	hxrt.Println(v_1)
	m := New_go___Map()
	m.set(42, hxrt.StringFromLiteral("answer"))
	var v_2 any = any(func(hx_value_5 any) bool {
		if hx_value_5 == nil {
			var hx_zero_6 bool
			return hx_zero_6
		}
		return hx_value_5.(bool)
	}(m.exists(42)))
	hxrt.Println(v_2)
	var v_3 any = any(func(hx_value_7 any) *string {
		if hx_value_7 == nil {
			var hx_zero_8 *string
			return hx_zero_8
		}
		return hx_value_7.(*string)
	}(m.get(42)))
	hxrt.Println(v_3)
}
