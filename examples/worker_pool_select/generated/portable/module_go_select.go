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
	firstRecv := go__concurrency_tryRecv___string_f613ccd0(first.__hx_native)
	if func(hx_value_15 any) bool {
		if hx_value_15 == nil {
			var hx_zero_16 bool
			return hx_zero_16
		}
		return hx_value_15.(bool)
	}(firstRecv.isOk()) {
		return go___SelectRecv2_First(func(hx_value_13 any) *string {
			if hx_value_13 == nil {
				var hx_zero_14 *string
				return hx_zero_14
			}
			return hx_value_13.(*string)
		}(firstRecv.unwrap()))
	}
	secondRecv := go__concurrency_tryRecv___string_f613ccd0(second.__hx_native)
	if func(hx_value_19 any) bool {
		if hx_value_19 == nil {
			var hx_zero_20 bool
			return hx_zero_20
		}
		return hx_value_19.(bool)
	}(secondRecv.isOk()) {
		return go___SelectRecv2_Second(func(hx_value_17 any) *string {
			if hx_value_17 == nil {
				var hx_zero_18 *string
				return hx_zero_18
			}
			return hx_value_17.(*string)
		}(secondRecv.unwrap()))
	}
	return go___SelectRecv2_Defaulted
}

func go___Select_recv_Int(channel *go___Chan) *go___SelectRecv {
	received := go__concurrency_tryRecv__int_95e97e5e(channel.__hx_native)
	if func(hx_value_23 any) bool {
		if hx_value_23 == nil {
			var hx_zero_24 bool
			return hx_zero_24
		}
		return hx_value_23.(bool)
	}(received.isOk()) {
		return go___SelectRecv_Received(func(hx_value_21 any) int {
			if hx_value_21 == nil {
				var hx_zero_22 int
				return hx_zero_22
			}
			return hx_value_21.(int)
		}(received.unwrap()))
	}
	return go___SelectRecv_Defaulted
}

func go___Select_send2_Int_Int(first *go___Chan, firstValue int, second *go___Chan, secondValue int) *go___SelectSend2 {
	if go__concurrency_trySend__int_95e97e5e(first.__hx_native, firstValue) {
		return go___SelectSend2_FirstSent
	}
	if go__concurrency_trySend__int_95e97e5e(second.__hx_native, secondValue) {
		return go___SelectSend2_SecondSent
	}
	return go___SelectSend2_Defaulted
}

func go___Select_send_Int(channel *go___Chan, value int) *go___SelectSend {
	if go__concurrency_trySend__int_95e97e5e(channel.__hx_native, value) {
		return go___SelectSend_Sent
	}
	return go___SelectSend_Defaulted
}
