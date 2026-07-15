package main

import (
	"fmt"
	"net/http"
	"snapshot/hxrt"
	"time"
)

func main() {
	now := time.Now()
	direct := now.Unix()
	viaReceiver := now.Unix()
	statusOk := hxrt.StringEqualStringPtr(hxrt.StdString(http.StatusText(200)), hxrt.StringFromLiteral("OK"))
	if ((direct == viaReceiver) && (direct > 0)) && statusOk {
		fmt.Println(321)
	} else {
		hxrt.Println(any(-1))
	}
}
