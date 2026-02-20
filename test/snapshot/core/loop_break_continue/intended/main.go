package main

import "snapshot/hxrt"

func main() {
	i := 0
	_ = i
	sum := 0
	_ = sum
	for i < 10 {
		i = int(int32((i + 1)))
		if int(int32((hxrt.Int32Wrap(i) % hxrt.Int32Wrap(2)))) == 0 {
			continue
		}
		if i > 7 {
			break
		}
		sum = int(int32((hxrt.Int32Wrap(sum) + hxrt.Int32Wrap(i))))
	}
	hxrt.Println(hxrt.StdString(sum))
	hxrt.Println(hxrt.StdString(i))
}
