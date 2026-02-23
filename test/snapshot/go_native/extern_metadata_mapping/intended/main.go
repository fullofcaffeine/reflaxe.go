package main

import (
	"fmt"
	"snapshot/hxrt"
	"time"
)

func main() {
	now := time.Now()
	_ = now
	direct := now.Unix()
	_ = direct
	viaReceiver := now.Unix()
	_ = viaReceiver
	if (direct == viaReceiver) && (direct > 0) {
		fmt.Println(321)
	} else {
		hxrt.Println(-1)
	}
}
