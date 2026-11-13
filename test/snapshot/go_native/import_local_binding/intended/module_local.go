package main

import "snapshot/hxrt"

type I_Local interface {
}

type Local struct {
	__hx_this I_Local
	value     int
}

func (self *Local) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_Local() *Local {
	return hx_init_Local(hx_alloc_Local(nil))
}

func hx_alloc_Local(owner any) *Local {
	self := &Local{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_Local)
	}
	return self
}

func hx_init_Local(self *Local) *Local {
	self.value = 7
	return self
}

func Local_argument(bytes_local *Local) int {
	return bytes_local.value
}

func Local_catchValue() *string {
	hx_try_return_1 := false
	var hx_try_value_2 *string
	hxrt.TryCatch(func() {
		hxrt.Throw(haxe__Exception_thrown(hxrt.StringFromLiteral("caught")))
	}, func(hx_caught_3 any) {
		hx_exception_4 := haxe__Exception_caught(hx_caught_3)
		hx_payload_5 := haxe__Exception_catchValue(hx_caught_3)
		_ = hx_payload_5
		var bytes_local *haxe__Exception = hx_recover_class_0(hx_exception_4)
		hx_try_value_2 = bytes_local.__hx_this.get_message()
		hx_try_return_1 = true
		return
	})
	if hx_try_return_1 {
		return hx_try_value_2
	}
	var hx_try_zero_6 *string
	return hx_try_zero_6
}

func Local_closure() int {
	strings_local := 7
	fn := hxrt.NewFunction(func() int {
		strings_local_1 := 3
		return int((hxrt.Int32Wrap(int((hxrt.Int32Wrap(strings_local) * hxrt.Int32Wrap(10)))) + hxrt.Int32Wrap(strings_local_1)))
	})
	return func() int {
		var hx_callable_7 *hxrt.Function[func() int] = fn
		return hxrt.FunctionCall(hx_callable_7)()
	}()
}

func Local_loops() int {
	sum := 0
	bytes_local := 1
	sum = int((hxrt.Int32Wrap(sum) + hxrt.Int32Wrap(bytes_local)))
	bytes_local_1 := 2
	sum = int((hxrt.Int32Wrap(sum) + hxrt.Int32Wrap(bytes_local_1)))
	bytes_local_2 := 3
	sum = int((hxrt.Int32Wrap(sum) + hxrt.Int32Wrap(bytes_local_2)))
	return sum
}

func Local_read() int {
	bytes_local := New_Local()
	strings_local := New_Local()
	return int((hxrt.Int32Wrap(int((hxrt.Int32Wrap(bytes_local.value) + hxrt.Int32Wrap(strings_local.value)))) - hxrt.Int32Wrap(7)))
}

func Local_runtimeAlias() int {
	hxrt_local := New_Local()
	return int((hxrt.Int32Wrap(hxrt_local.value) + hxrt.Int32Wrap(hxrt.StringLengthStringPtr(hxrt.StringFromLiteral("abc")))))
}

func Local_suffix() int {
	total := 0
	bytes_local := New_Local()
	total = int((hxrt.Int32Wrap(total) + hxrt.Int32Wrap(bytes_local.value)))
	bytes_local_1 := New_Local()
	total = int((hxrt.Int32Wrap(total) + hxrt.Int32Wrap(bytes_local_1.value)))
	bytes_local_2 := 1
	total = int((hxrt.Int32Wrap(total) + hxrt.Int32Wrap(bytes_local_2)))
	return total
}
