package main

import (
	"context"
	"examples_interop_smoke_portable/hxrt"
	"fmt"
	"net/http"
	"time"
)

func main() {
	now := time.Now()
	_ = now
	unixDirect := now.Unix()
	_ = unixDirect
	unixReceiver := now.Unix()
	_ = unixReceiver
	var ctx context.Context = context.Background()
	_ = ctx
	statusOk := hxrt.StringEqualAny(hxrt.StdString(http.StatusText(200)), hxrt.StringFromLiteral("OK"))
	_ = statusOk
	ok := ((((unixDirect == unixReceiver) && (unixDirect > 0)) && !hxrt.StringEqualAny(ctx, nil)) && statusOk)
	_ = ok
	fmt.Println(func() int {
		var hx_if_1 int
		if ok {
			hx_if_1 = 1
		} else {
			hx_if_1 = 0
		}
		return hx_if_1
	}())
}
