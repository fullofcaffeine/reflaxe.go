package main

import "snapshot/hxrt"

func main() {
	host := New_sys__net__Host(hxrt.StringFromLiteral("127.0.0.1"))
	addr := New_sys__net__Address()
	addr.host = host.ip
	addr.port = 3210
	var v any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("snapshot.host="), addr.getHost().toString()))
	hxrt.Println(v)
	var v_1 any = any(hxrt.StringConcatAny(hxrt.StringFromLiteral("snapshot.compare="), addr.compare(addr.clone())))
	hxrt.Println(v_1)
	hxrt.Println(any(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("snapshot.alg="), any(hxrt.StringFromLiteral("SHA224"))), hxrt.StringFromLiteral(",")), any(hxrt.StringFromLiteral("SHA384"))), hxrt.StringFromLiteral(",")), any(hxrt.StringFromLiteral("RIPEMD160")))))
}
