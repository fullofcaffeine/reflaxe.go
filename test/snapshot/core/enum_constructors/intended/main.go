package main

import "snapshot/hxrt"

type Color struct {
	tag    int
	params []any
}

var Color_Red *Color = &Color{tag: 0}

func Color_RGB(r int, g int, b int) *Color {
	value := &Color{tag: 1}
	value.params = []any{r, g, b}
	return value
}

func isSome(value *Color) int {
	if hxrt.StringEqualAny(value, nil) {
		return 0
	}
	return 1
}

func main() {
	red := Color_Red
	rgb := Color_RGB(1, 2, 3)
	hxrt.Println(isSome(red))
	hxrt.Println(isSome(rgb))
}
