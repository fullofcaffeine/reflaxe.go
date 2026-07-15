package main

import "snapshot/hxrt"

type Expr struct {
	tag    int
	params []any
}

func Expr_Lit(value int) *Expr {
	enumValue := &Expr{tag: 0}
	enumValue.params = []any{value}
	return enumValue
}

func Expr_Pair(left int, right int) *Expr {
	enumValue := &Expr{tag: 1}
	enumValue.params = []any{left, right}
	return enumValue
}

func eval(expr *Expr) int {
	var hx_switch_1 int
	switch expr.tag {
	case 0:
		_g := expr.params[0].(int)
		value := _g
		hx_switch_1 = value
	case 1:
		_g_1 := expr.params[0].(int)
		_g1 := expr.params[1].(int)
		left := _g_1
		right := _g1
		hx_switch_1 = int(int32((hxrt.Int32Wrap(left) + hxrt.Int32Wrap(right))))
	}
	return hx_switch_1
}

func main() {
	writeOnly()
	var v any = any(eval(Expr_Lit(3)))
	hxrt.Println(v)
	var v_1 any = any(eval(Expr_Pair(2, 5)))
	hxrt.Println(v_1)
}

func writeOnly() {
	x := 0
	_ = x
	x = 1
}
