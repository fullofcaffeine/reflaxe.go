package main

import "snapshot/hxrt"

func main() {
	values := hxrt.NewArray()
	values.Push(4)
	values.Push(9)
	values.Pop()
	pushLen := func() int {
		return values.Push(12)
	}()
	var removed any = func() any {
		return values.Pop()
	}()
	var v any = any(values.Len())
	hxrt.Println(v)
	hxrt.Println(any(values.Get(0)))
	hxrt.Println(any(pushLen))
	hxrt.Println(any(removed))
}
