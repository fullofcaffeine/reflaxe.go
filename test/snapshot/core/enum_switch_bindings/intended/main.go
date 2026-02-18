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
		_ = _g
		value := _g
		_ = value
		hx_switch_1 = value
	case 1:
		_g_1 := expr.params[0].(int)
		_ = _g_1
		_g1 := expr.params[1].(int)
		_ = _g1
		left := _g_1
		_ = left
		right := _g1
		_ = right
		hx_switch_1 = (left + right)
	}
	return hx_switch_1
}

func main() {
	hxrt.Println(eval(Expr_Lit(3)))
	hxrt.Println(eval(Expr_Pair(2, 5)))
}
