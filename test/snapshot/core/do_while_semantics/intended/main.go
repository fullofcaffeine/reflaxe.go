package main

import "snapshot/hxrt"

func main() {
	once := 0
	hx_do_first_1 := true
	for hx_do_first_1 || false {
		hx_do_first_1 = false
		once = int(int32((once + 1)))
	}
	i := 0
	hit := 0
	hx_do_first_2 := true
	for hx_do_first_2 || (i < 3) {
		hx_do_first_2 = false
		i = int(int32((i + 1)))
		if i < 3 {
			continue
		}
		hit = i
	}
	var v any = any(hxrt.StdString(once))
	hxrt.Println(v)
	var v_1 any = any(hxrt.StdString(hit))
	hxrt.Println(v_1)
}
