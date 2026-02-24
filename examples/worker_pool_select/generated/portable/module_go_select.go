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

func go___Select_recv2_String_String(first *go___Chan, second *go___Chan) *go___SelectRecv2 {
	firstRecv := func(hx_value_18 any) *go___Result {
		if hx_value_18 == nil {
			var hx_zero_19 *go___Result
			return hx_zero_19
		}
		return hx_value_18.(*go___Result)
	}(first.__hx_this.tryRecv())
	_ = firstRecv
	if firstRecv.__hx_this.isOk() {
		return go___SelectRecv2_First(firstRecv.__hx_this.unwrap())
	}
	secondRecv := func(hx_value_20 any) *go___Result {
		if hx_value_20 == nil {
			var hx_zero_21 *go___Result
			return hx_zero_21
		}
		return hx_value_20.(*go___Result)
	}(second.__hx_this.tryRecv())
	_ = secondRecv
	if secondRecv.__hx_this.isOk() {
		return go___SelectRecv2_Second(secondRecv.__hx_this.unwrap())
	}
	return go___SelectRecv2_Defaulted
}

func go___Select_recv_Int(channel *go___Chan) *go___SelectRecv {
	received := func(hx_value_22 any) *go___Result {
		if hx_value_22 == nil {
			var hx_zero_23 *go___Result
			return hx_zero_23
		}
		return hx_value_22.(*go___Result)
	}(channel.__hx_this.tryRecv())
	_ = received
	if received.__hx_this.isOk() {
		return go___SelectRecv_Received(received.__hx_this.unwrap())
	}
	return go___SelectRecv_Defaulted
}

func go___Select_send2_Int_Int(first *go___Chan, firstValue int, second *go___Chan, secondValue int) *go___SelectSend2 {
	if first.__hx_this.trySend(firstValue) {
		return go___SelectSend2_FirstSent
	}
	if second.__hx_this.trySend(secondValue) {
		return go___SelectSend2_SecondSent
	}
	return go___SelectSend2_Defaulted
}

func go___Select_send_Int(channel *go___Chan, value int) *go___SelectSend {
	if channel.__hx_this.trySend(value) {
		return go___SelectSend_Sent
	}
	return go___SelectSend_Defaulted
}
