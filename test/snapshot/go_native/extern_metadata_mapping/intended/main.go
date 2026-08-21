package main

import (
	"fmt"
	"net/http"
	"os"
	"snapshot/hxrt"
	"time"
)

func main() {
	now := time.Now()
	direct := now.Unix()
	viaReceiver := now.Unix()
	statusOk := hxrt.StringEqualStringPtr(hxrt.StdString(http.StatusText(200)), hxrt.StringFromLiteral("OK"))
	if ((direct == viaReceiver) && (direct > 0)) && statusOk {
		os.Setenv(*hxrt.StdString(hxrt.StringFromLiteral("HAXE_GO_EXTERN_STRING")), *hxrt.StdString(hxrt.StringFromLiteral("converted")))
		fmt.Println(321)
	} else {
		hxrt.Println(any(-1))
	}
}
