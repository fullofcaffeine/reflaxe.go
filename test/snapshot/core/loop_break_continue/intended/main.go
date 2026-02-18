package main

import "snapshot/hxrt"

func main() {
	i := 0
	_ = i
	sum := 0
	_ = sum
	for i < 10 {
		i = (i + 1)
		if (i % 2) == 0 {
			continue
		}
		if i > 7 {
			break
		}
		sum = (sum + i)
	}
	hxrt.Println(hxrt.StdString(sum))
	hxrt.Println(hxrt.StdString(i))
}
