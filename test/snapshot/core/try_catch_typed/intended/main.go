package main

import "snapshot/hxrt"

func handle(v int) {
	hxrt.TryCatch(func() {
		raise(v)
	}, func(hx_caught_1 any) {
		switch hx_typed_2 := hx_caught_1.(type) {
		case *string:
			s := hx_typed_2
			_ = s
			hxrt.Println(hxrt.StringConcatAny(hxrt.StringFromLiteral("S:"), s))
		case int:
			i := hx_typed_2
			_ = i
			hxrt.Println(i)
		default:
			e := hx_caught_1
			_ = e
			hxrt.Println(hxrt.StringFromLiteral("D"))
		}
	})
}

func main() {
	handle(0)
	handle(1)
	handle(2)
}

func raise(v int) {
	if v == 0 {
		hxrt.Throw(hxrt.StringFromLiteral("text"))
	}
	if v == 1 {
		hxrt.Throw(11)
	}
	hxrt.Throw(true)
}
