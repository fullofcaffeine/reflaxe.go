package main

import (
	"snapshot/hxrt"
	"time"
)

func main() {
	now := time.Now()
	unix := now.Unix
	var v any = any((unix() > 0))
	hxrt.Println(v)
}
