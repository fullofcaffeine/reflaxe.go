package main

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
	return func(hx_value_3 any) *go___Result {
		if hx_value_3 == nil {
			var hx_zero_4 *go___Result
			return hx_zero_4
		}
		return hx_value_3.(*go___Result)
	}(nil)
}

func go___Go___chanTrySend(channel any, value any) bool {
	return false
}

func go___Go___goSpawn(fn func()) {
}

func go___Go_fail(message *string) *go___Result {
	return New_go___Result(nil, New_go___Error(message))
}

func go___Go_newChan(buffer int) *go___Chan {
	channel := New_go___Chan()
	if buffer > 0 {
		channel.__hx_setBuffer(buffer)
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
	return New_go___Result(value, nil)
}

func go___Go_spawn(fn func()) {
	go__concurrency_spawn(fn)
}
