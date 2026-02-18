package main

import "snapshot/hxrt"

type Color struct {
	tag    int
	params []any
}

var Color_Red *Color = &Color{tag: 0}

func Color_RGB(r int, g int, b int) *Color {
	enumValue := &Color{tag: 1}
	enumValue.params = []any{r, g, b}
	return enumValue
}

func isSome(value *Color) int {
	if hxrt.StringEqualAny(value, nil) {
		return 0
	}
	return 1
}

func main() {
	red := Color_Red
	_ = red
	rgb := Color_RGB(1, 2, 3)
	_ = rgb
	hxrt.Println(isSome(red))
	hxrt.Println(isSome(rgb))
}
