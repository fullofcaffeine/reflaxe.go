package main

import "snapshot/hxrt"

type Parcel struct {
	tag    int
	params []any
}

func Parcel_Item(value int) *Parcel {
	enumValue := &Parcel{tag: 0}
	enumValue.params = []any{value}
	return enumValue
}

func main() {
	make = hxrt.NamedFunction(Parcel_Item, "Parcel.Item")
	alias := make
	_g := func() *Parcel {
		var hx_callable_1 *hxrt.Function[func(int) *Parcel] = alias
		var hx_argument_2 int = 17
		return hxrt.FunctionCall(hx_callable_1)(hx_argument_2)
	}()
	_g_1 := _g.params[0].(int)
	value := _g_1
	hxrt.Println(any(value))
	var v any = any(func() bool {
		var f1 any = any(make)
		return hxrt.ReflectCompareMethods(f1, any(alias))
	}())
	hxrt.Println(v)
	var v_1 any = any(hxrt.ReflectCompareMethods(any(hxrt.NamedFunction(Parcel_Item, "Parcel.Item")), any(hxrt.NamedFunction(Parcel_Item, "Parcel.Item"))))
	hxrt.Println(v_1)
}

var make *hxrt.Function[func(int) *Parcel]
