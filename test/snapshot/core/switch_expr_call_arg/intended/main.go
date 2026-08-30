package main

import "snapshot/hxrt"

type Kind struct {
	tag    int
	params []any
}

var Kind_A *Kind = &Kind{tag: 0}

func Kind_B(value int) *Kind {
	enumValue := &Kind{tag: 1}
	enumValue.params = []any{value}
	return enumValue
}

func asInt(kind *Kind) int {
	return id(func() int {
		var hx_switch_1 int
		switch kind.tag {
		case 0:
			hx_switch_1 = 1
		case 1:
			_g := kind.params[0].(int)
			value := _g
			hx_switch_1 = int((hxrt.Int32Wrap(value) + hxrt.Int32Wrap(1)))
		}
		return hx_switch_1
	}())
}

func id(value int) int {
	return value
}

func main() {
	var v any = any(hxrt.StdString(asInt(Kind_A)))
	hxrt.Println(v)
	var v_1 any = any(hxrt.StdString(asInt(Kind_B(6))))
	hxrt.Println(v_1)
}
