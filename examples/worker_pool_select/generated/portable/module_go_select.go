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
	if func(hx_value_3 any) bool {
		if hx_value_3 == nil {
			var hx_zero_4 bool
			return hx_zero_4
		}
		return hx_value_3.(bool)
	}(firstRecv.isOk()) {
		return go___SelectRecv2_First(func(hx_value_1 any) *string {
			if hx_value_1 == nil {
				var hx_zero_2 *string
				return hx_zero_2
			}
			return hx_value_1.(*string)
		}(firstRecv.unwrap()))
	}
	secondRecv := go__concurrency_tryRecv___string_f613ccd0(second.__hx_native)
	if func(hx_value_7 any) bool {
		if hx_value_7 == nil {
			var hx_zero_8 bool
			return hx_zero_8
		}
		return hx_value_7.(bool)
	}(secondRecv.isOk()) {
		return go___SelectRecv2_Second(func(hx_value_5 any) *string {
			if hx_value_5 == nil {
				var hx_zero_6 *string
				return hx_zero_6
			}
			return hx_value_5.(*string)
		}(secondRecv.unwrap()))
	}
	return go___SelectRecv2_Defaulted
}

func go___Select_recv_Int(channel *go___Chan) *go___SelectRecv {
	received := go__concurrency_tryRecv__int_95e97e5e(channel.__hx_native)
	if func(hx_value_11 any) bool {
		if hx_value_11 == nil {
			var hx_zero_12 bool
			return hx_zero_12
		}
		return hx_value_11.(bool)
	}(received.isOk()) {
		return go___SelectRecv_Received(func(hx_value_9 any) int {
			if hx_value_9 == nil {
				var hx_zero_10 int
				return hx_zero_10
			}
			return hx_value_9.(int)
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
