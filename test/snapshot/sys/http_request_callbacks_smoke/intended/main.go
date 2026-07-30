package main

import "snapshot/hxrt"

func main() {
	http := New_sys__Http(hxrt.StringFromLiteral("data:text/plain,hello%20from%20haxe.go"))
	dataLog := hxrt.StringFromLiteral("")
	statusLog := -1
	byteCount := -1
	errLog := hxrt.StringFromLiteral("")
	http.onData = func(data *string) {
		dataLog = data
	}
	http.onStatus = func(status int) {
		statusLog = status
	}
	http.onBytes = func(bytes *haxe__io__Bytes) {
		byteCount = bytes.length
	}
	http.onError = func(msg *string) {
		errLog = msg
	}
	http.__hx_this.setHeader(hxrt.StringFromLiteral("X-Test"), hxrt.StringFromLiteral("one"))
	http.__hx_this.setHeader(hxrt.StringFromLiteral("X-Test"), hxrt.StringFromLiteral("two"))
	http.__hx_this.addHeader(hxrt.StringFromLiteral("X-Trace"), hxrt.StringFromLiteral("ok"))
	http.__hx_this.setParameter(hxrt.StringFromLiteral("a"), hxrt.StringFromLiteral("1"))
	http.__hx_this.addParameter(hxrt.StringFromLiteral("b"), hxrt.StringFromLiteral("2"))
	http.__hx_this.request(nil)
	hxrt.Println(any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("data="), dataLog)))
	hxrt.Println(any(hxrt.StringConcatAny(hxrt.StringFromLiteral("bytes="), byteCount)))
	hxrt.Println(any(hxrt.StringConcatAny(hxrt.StringFromLiteral("status="), statusLog)))
	var v any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("response="), http.__hx_this.get_responseData()))
	hxrt.Println(v)
	hxrt.Println(any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("error="), errLog)))
	post := New_sys__Http(hxrt.StringFromLiteral("data:text/plain,ignored"))
	post.__hx_this.setPostData(hxrt.StringFromLiteral("post-body"))
	postData := hxrt.StringFromLiteral("")
	post.onData = func(data *string) {
		postData = data
	}
	post.__hx_this.request(true)
	hxrt.Println(any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("post="), postData)))
	bad := New_sys__Http(hxrt.StringFromLiteral("://bad"))
	badErr := hxrt.StringFromLiteral("")
	bad.onError = func(msg *string) {
		badErr = msg
	}
	bad.__hx_this.request(nil)
	hxrt.Println(any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("bad="), badErr)))
	var v_1 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("direct="), sys__Http_requestUrl(hxrt.StringFromLiteral("data:text/plain,direct%20ok"))))
	hxrt.Println(v_1)
}
