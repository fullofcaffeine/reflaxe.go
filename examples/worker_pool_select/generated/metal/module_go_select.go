package main

type go___SelectRecv struct {
	tag    int
	params []any
}

func go___SelectRecv_Received(value any) *go___SelectRecv {
	enumValue := &go___SelectRecv{tag: 0}
	enumValue.params = []any{value}
	return enumValue
}

var go___SelectRecv_Defaulted *go___SelectRecv = &go___SelectRecv{tag: 1}

type go___SelectRecv2 struct {
	tag    int
	params []any
}

func go___SelectRecv2_First(value any) *go___SelectRecv2 {
	enumValue := &go___SelectRecv2{tag: 0}
	enumValue.params = []any{value}
	return enumValue
}

func go___SelectRecv2_Second(value any) *go___SelectRecv2 {
	enumValue := &go___SelectRecv2{tag: 1}
	enumValue.params = []any{value}
	return enumValue
}

var go___SelectRecv2_Defaulted *go___SelectRecv2 = &go___SelectRecv2{tag: 2}

type go___SelectSend struct {
	tag    int
	params []any
}

var go___SelectSend_Sent *go___SelectSend = &go___SelectSend{tag: 0}

var go___SelectSend_Defaulted *go___SelectSend = &go___SelectSend{tag: 1}

type go___SelectSend2 struct {
	tag    int
	params []any
}

var go___SelectSend2_FirstSent *go___SelectSend2 = &go___SelectSend2{tag: 0}

var go___SelectSend2_SecondSent *go___SelectSend2 = &go___SelectSend2{tag: 1}

var go___SelectSend2_Defaulted *go___SelectSend2 = &go___SelectSend2{tag: 2}

func go___Select_recv(channel *go___Chan) *go___SelectRecv {
	received := func(hx_value_10 any) *go___Result {
		if hx_value_10 == nil {
			var hx_zero_11 *go___Result
			return hx_zero_11
		}
		return hx_value_10.(*go___Result)
	}(channel.__hx_this.tryRecv())
	if received.__hx_this.isOk() {
		return go___SelectRecv_Received(received.__hx_this.unwrap())
	}
	return go___SelectRecv_Defaulted
}

func go___Select_recv2(first *go___Chan, second *go___Chan) *go___SelectRecv2 {
	firstRecv := func(hx_value_12 any) *go___Result {
		if hx_value_12 == nil {
			var hx_zero_13 *go___Result
			return hx_zero_13
		}
		return hx_value_12.(*go___Result)
	}(first.__hx_this.tryRecv())
	if firstRecv.__hx_this.isOk() {
		return go___SelectRecv2_First(firstRecv.__hx_this.unwrap())
	}
	secondRecv := func(hx_value_14 any) *go___Result {
		if hx_value_14 == nil {
			var hx_zero_15 *go___Result
			return hx_zero_15
		}
		return hx_value_14.(*go___Result)
	}(second.__hx_this.tryRecv())
	if secondRecv.__hx_this.isOk() {
		return go___SelectRecv2_Second(secondRecv.__hx_this.unwrap())
	}
	return go___SelectRecv2_Defaulted
}

func go___Select_send(channel *go___Chan, value any) *go___SelectSend {
	if channel.__hx_this.trySend(value) {
		return go___SelectSend_Sent
	}
	return go___SelectSend_Defaulted
}

func go___Select_send2(first *go___Chan, firstValue any, second *go___Chan, secondValue any) *go___SelectSend2 {
	if first.__hx_this.trySend(firstValue) {
		return go___SelectSend2_FirstSent
	}
	if second.__hx_this.trySend(secondValue) {
		return go___SelectSend2_SecondSent
	}
	return go___SelectSend2_Defaulted
}
