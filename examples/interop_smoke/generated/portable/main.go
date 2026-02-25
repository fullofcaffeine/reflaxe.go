package main

import (
	"context"
	"examples_interop_smoke_portable/hxrt"
	"fmt"
	"net/http"
	"time"
)

func main() {
	wrappedNow := time.Now()
	wrappedUnixDirect := wrappedNow.Unix()
	_ = wrappedUnixDirect
	wrappedUnixReceiver := wrappedNow.Unix()
	_ = wrappedUnixReceiver
	var wrappedCtx context.Context = context.Background()
	_ = wrappedCtx
	wrappedStatusOk := hxrt.StringEqualStringPtr(hxrt.StdString(http.StatusText(200)), hxrt.StringFromLiteral("OK"))
	wrappedOk := ((((wrappedUnixDirect == wrappedUnixReceiver) && (wrappedUnixDirect > 0)) && (wrappedCtx != nil)) && wrappedStatusOk)
	_ = wrappedOk
	externNow := time.Now()
	externUnixDirect := externNow.Unix()
	_ = externUnixDirect
	externUnixReceiver := externNow.Unix()
	_ = externUnixReceiver
	var externCtx context.Context = context.Background()
	_ = externCtx
	externStatusOk := hxrt.StringEqualStringPtr(hxrt.StdString(http.StatusText(200)), hxrt.StringFromLiteral("OK"))
	externOk := ((((externUnixDirect == externUnixReceiver) && (externUnixDirect > 0)) && (externCtx != nil)) && externStatusOk)
	fmt.Println(func() int {
		var hx_if_1 int
		if wrappedOk && externOk {
			hx_if_1 = 1
		} else {
			hx_if_1 = 0
		}
		return hx_if_1
	}())
}
