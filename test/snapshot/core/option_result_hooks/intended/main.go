package main

import "snapshot/hxrt"

type Maybe struct {
	tag    int
	params []any
}

var Maybe_None *Maybe = &Maybe{tag: 0}

func Maybe_Some(value any) *Maybe {
	enumValue := &Maybe{tag: 1}
	enumValue.params = []any{value}
	return enumValue
}

type Res struct {
	tag    int
	params []any
}

func Res_Ok(value any) *Res {
	enumValue := &Res{tag: 0}
	enumValue.params = []any{value}
	return enumValue
}

func Res_Err(error any) *Res {
	enumValue := &Res{tag: 1}
	enumValue.params = []any{error}
	return enumValue
}

func main() {
	hxrt.Println(unwrapOr(Maybe_Some(7), 0))
	hxrt.Println(unwrapOr(Maybe_None, 5))
	hxrt.Println(render(Res_Ok(9)))
	hxrt.Println(render(Res_Err(hxrt.StringFromLiteral("bad"))))
}

func render(res *Res) *string {
	var hx_switch_1 *string
	switch res.tag {
	case 0:
		_g := res.params[0].(int)
		_ = _g
		v := _g
		_ = v
		hx_switch_1 = hxrt.StdString(v)
	case 1:
		_g_1 := res.params[0].(*string)
		_ = _g_1
		e := _g_1
		_ = e
		hx_switch_1 = e
	}
	return hx_switch_1
}

func unwrapOr(value *Maybe, fallback int) int {
	var hx_switch_2 int
	switch value.tag {
	case 0:
		hx_switch_2 = fallback
	case 1:
		_g := value.params[0].(int)
		_ = _g
		v := _g
		_ = v
		hx_switch_2 = v
	}
	return hx_switch_2
}
