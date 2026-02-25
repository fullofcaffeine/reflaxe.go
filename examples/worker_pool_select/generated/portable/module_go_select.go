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
	firstRecv := func(hx_value_20 any) *go___Result {
		if hx_value_20 == nil {
			var hx_zero_21 *go___Result
			return hx_zero_21
		}
		return hx_value_20.(*go___Result)
	}(first.tryRecv())
	if func(hx_value_22 any) bool {
		if hx_value_22 == nil {
			var hx_zero_23 bool
			return hx_zero_23
		}
		return hx_value_22.(bool)
	}(firstRecv.isOk()) {
		return go___SelectRecv2_First(func(hx_value_24 any) *string {
			if hx_value_24 == nil {
				var hx_zero_25 *string
				return hx_zero_25
			}
			return hx_value_24.(*string)
		}(firstRecv.unwrap()))
	}
	secondRecv := func(hx_value_26 any) *go___Result {
		if hx_value_26 == nil {
			var hx_zero_27 *go___Result
			return hx_zero_27
		}
		return hx_value_26.(*go___Result)
	}(second.tryRecv())
	if func(hx_value_28 any) bool {
		if hx_value_28 == nil {
			var hx_zero_29 bool
			return hx_zero_29
		}
		return hx_value_28.(bool)
	}(secondRecv.isOk()) {
		return go___SelectRecv2_Second(func(hx_value_30 any) *string {
			if hx_value_30 == nil {
				var hx_zero_31 *string
				return hx_zero_31
			}
			return hx_value_30.(*string)
		}(secondRecv.unwrap()))
	}
	return go___SelectRecv2_Defaulted
}

func go___Select_recv_Int(channel *go___Chan) *go___SelectRecv {
	received := func(hx_value_32 any) *go___Result {
		if hx_value_32 == nil {
			var hx_zero_33 *go___Result
			return hx_zero_33
		}
		return hx_value_32.(*go___Result)
	}(channel.tryRecv())
	if func(hx_value_34 any) bool {
		if hx_value_34 == nil {
			var hx_zero_35 bool
			return hx_zero_35
		}
		return hx_value_34.(bool)
	}(received.isOk()) {
		return go___SelectRecv_Received(func(hx_value_36 any) int {
			if hx_value_36 == nil {
				var hx_zero_37 int
				return hx_zero_37
			}
			return hx_value_36.(int)
		}(received.unwrap()))
	}
	return go___SelectRecv_Defaulted
}

func go___Select_send2_Int_Int(first *go___Chan, firstValue int, second *go___Chan, secondValue int) *go___SelectSend2 {
	if func(hx_value_38 any) bool {
		if hx_value_38 == nil {
			var hx_zero_39 bool
			return hx_zero_39
		}
		return hx_value_38.(bool)
	}(first.trySend(firstValue)) {
		return go___SelectSend2_FirstSent
	}
	if func(hx_value_40 any) bool {
		if hx_value_40 == nil {
			var hx_zero_41 bool
			return hx_zero_41
		}
		return hx_value_40.(bool)
	}(second.trySend(secondValue)) {
		return go___SelectSend2_SecondSent
	}
	return go___SelectSend2_Defaulted
}

func go___Select_send_Int(channel *go___Chan, value int) *go___SelectSend {
	if func(hx_value_42 any) bool {
		if hx_value_42 == nil {
			var hx_zero_43 bool
			return hx_zero_43
		}
		return hx_value_42.(bool)
	}(channel.trySend(value)) {
		return go___SelectSend_Sent
	}
	return go___SelectSend_Defaulted
}
