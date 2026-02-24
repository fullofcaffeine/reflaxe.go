package main

import "examples_worker_pool_select_portable/hxrt"

func go___Go___chanClose(channel any) {
}

func go___Go___chanMake(buffer int) any {
	return nil
}

func go___Go___chanRecv(channel any) any {
	return nil
}

func go___Go___chanRecvOr(channel any, defaultValue any) any {
	return defaultValue
}

func go___Go___chanSend(channel any, value any) {
}

func go___Go___chanTryRecv(channel any) *go___Result {
	return go___Result_failure(hxrt.StringFromLiteral("empty"))
}

func go___Go___chanTrySend(channel any, value any) bool {
	return false
}

func go___Go___goSpawn(fn func()) {
}

func go___Go_fail(message *string) *go___Result {
	return go___Result_failure(message)
}

func go___Go_newChan(buffer int) *go___Chan {
	channel := New_go___Chan()
	_ = channel
	if buffer > 0 {
		channel.__hx_this.__hx_setBuffer(buffer)
	}
	return channel
}

func go___Go_newMap() *go___Map {
	return New_go___Map()
}

func go___Go_newSlice() *go___Slice {
	return New_go___Slice()
}

func go___Go_ok(value any) *go___Result {
	return go___Result_ok(value)
}

func go___Go_spawn(fn func()) {
	go__concurrency_spawn(fn)
}
