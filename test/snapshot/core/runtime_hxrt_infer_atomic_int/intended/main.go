package main

import "snapshot/hxrt"

func main() {
	var this1 *hxrt.AtomicIntCell
	this1 = hxrt.AtomicIntNew(1)
	count := this1
	hxrt.AtomicIntAdd(count, 2)
	var this1_1 *hxrt.AtomicIntCell
	var this1_2 *hxrt.AtomicIntCell
	this1_2 = hxrt.AtomicIntNew(0)
	this1_1 = this1_2
	flag := this1_1
	v := hxrt.AtomicIntExchange(flag, 1)
	_ = (v == 1)
}
