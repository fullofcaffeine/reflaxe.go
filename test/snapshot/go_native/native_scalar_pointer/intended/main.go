package main

import (
	"fmt"
	"snapshot/hxrt"
)

func main() {
	original := 3
	number := func(hx_native_pointer_value_1 int) *int {
		return &hx_native_pointer_value_1
	}(original)
	independent := func(hx_native_pointer_value_2 int) *int {
		return &hx_native_pointer_value_2
	}(*number)
	word := func(hx_native_pointer_value_3 string) *string {
		return &hx_native_pointer_value_3
	}(*hxrt.StdString(hxrt.StringFromLiteral("before")))
	flag := func(hx_native_pointer_value_4 bool) *bool {
		return &hx_native_pointer_value_4
	}(false)
	fmt.Sscan(*hxrt.StdString(hxrt.StringFromLiteral("7 after true")), []any{number, word, flag}...)
	var v any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("mutation="), hxrt.StdString((((*number == 7) && hxrt.StringEqualStringPtr(hxrt.StringFromLiteral(*word), hxrt.StringFromLiteral("after"))) && *flag))))
	hxrt.Println(v)
	var v_1 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("independent="), hxrt.StdString(((original == 3) && (*independent == 3)))))
	hxrt.Println(v_1)
	alias := number
	fmt.Sscan(*hxrt.StdString(hxrt.StringFromLiteral("9")), []any{alias}...)
	var v_2 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("identity="), hxrt.StdString((*number == 9))))
	hxrt.Println(v_2)
	var v_3 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("typed-parameter="), hxrt.StdString((readNumber(number) == 9))))
	hxrt.Println(v_3)
	captured := hxrt.StringFromLiteral(*word)
	fmt.Sscan(*hxrt.StdString(hxrt.StringFromLiteral("latest")), []any{word}...)
	var v_4 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("string-snapshot="), hxrt.StdString((hxrt.StringEqualStringPtr(captured, hxrt.StringFromLiteral("after")) && hxrt.StringEqualStringPtr(hxrt.StringFromLiteral(*word), hxrt.StringFromLiteral("latest"))))))
	hxrt.Println(v_4)
	inputCalls := 0
	evaluated := func(hx_native_pointer_value_5 int) *int {
		return &hx_native_pointer_value_5
	}(func() int {
		inputCalls = int(int32((inputCalls + 1)))
		return inputCalls
	}())
	receiverCalls := 0
	read := *func(hx_native_pointer_value_6 int) *int {
		return &hx_native_pointer_value_6
	}(func() int {
		receiverCalls = int(int32((receiverCalls + 1)))
		return receiverCalls
	}())
	var v_5 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("evaluation="), hxrt.StdString(((((inputCalls == 1) && (*evaluated == 1)) && (receiverCalls == 1)) && (read == 1)))))
	hxrt.Println(v_5)
}

func readNumber(pointer *int) int {
	return *pointer
}
