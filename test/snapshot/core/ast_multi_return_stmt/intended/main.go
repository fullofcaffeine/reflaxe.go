package main

import (
	"snapshot/hxrt"
	"strconv"
)

func main() {
	hxrt.Println(any(hxrt.StringFromLiteral("ok")))
}

func hx_test_return_pair() (string, bool) {
	value := 42
	return strconv.Itoa(value), true
}

func hx_test_return_nil_pair() (any, bool) {
	return nil, true
}

func init() {
	value, matched := hx_test_return_pair()
	if !matched || (value != "42") {
		panic("multiple return lost value or match")
	}
	empty, found := hx_test_return_nil_pair()
	if !found || (empty != nil) {
		panic("successful nil result was lost")
	}
}
