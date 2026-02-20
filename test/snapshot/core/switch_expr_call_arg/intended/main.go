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
			_ = _g
			value := _g
			_ = value
			hx_switch_1 = int(int32((hxrt.Int32Wrap(value) + hxrt.Int32Wrap(1))))
		}
		return hx_switch_1
	}())
}

func id(value int) int {
	return value
}

func main() {
	hxrt.Println(hxrt.StdString(asInt(Kind_A)))
	hxrt.Println(hxrt.StdString(asInt(Kind_B(6))))
}
