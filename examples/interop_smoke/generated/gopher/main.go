package main

import (
	"context"
	"examples_interop_smoke_gopher/hxrt"
	"fmt"
	"net/http"
	"time"
)

func main() {
	now := time.Now()
	unixDirect := now.Unix()
	_ = unixDirect
	unixReceiver := now.Unix()
	_ = unixReceiver
	var ctx context.Context = context.Background()
	_ = ctx
	statusOk := hxrt.StringEqualStringPtr(hxrt.StdString(http.StatusText(200)), hxrt.StringFromLiteral("OK"))
	ok := ((((unixDirect == unixReceiver) && (unixDirect > 0)) && !hxrt.StringEqualAny(ctx, nil)) && statusOk)
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
