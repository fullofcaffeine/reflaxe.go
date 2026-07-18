package main

import "snapshot/hxrt"

func main() {
	var v any = any(hxrt.StdString(hxrt.HttpProxyDescriptor(hxrt.StringFromLiteral("proxy.local"), 3128, nil, nil)))
	hxrt.Println(v)
}
