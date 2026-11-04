//go:build exception_probe

package main

import "snapshot/hxrt"

type I_TaggedError interface {
	get_message() *string
	get_stack() any
	get_previous() *haxe__Exception
	get_native() any
	unwrap() any
	toString() *string
	details() *string
	code() int
}

type TaggedError struct {
	*haxe__Exception
	__hx_this I_TaggedError
}

func (self *TaggedError) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_TaggedError() *TaggedError {
	return hx_init_TaggedError(hx_alloc_TaggedError(nil))
}

func hx_alloc_TaggedError(owner any) *TaggedError {
	self := &TaggedError{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_TaggedError)
	}
	self.haxe__Exception = hx_alloc_haxe__Exception(self.__hx_this)
	return self
}

func hx_init_TaggedError(self *TaggedError) *TaggedError {
	hx_init_haxe__Exception(self.haxe__Exception, hxrt.StringFromLiteral("tagged"), nil, nil)
	return self
}

func (self *TaggedError) code() int {
	return 7
}

func (self *TaggedError) String() string {
	return *self.toString()
}

func TaggedError_check() bool {
	original := New_TaggedError()
	base := func(hx_value *TaggedError) *haxe__Exception {
		if hx_value == nil {
			return nil
		}
		return hx_value.haxe__Exception
	}(original)
	var erased any = base
	if !hx_is_nominal_1(erased) || !hx_is_nominal_0(erased) {
		hxrt.Throw(haxe__Exception_thrown(hxrt.StringFromLiteral("conditional ancestry was not recognized")))
	}
	var view Tag = original
	if !hx_is_nominal_2(view) {
		hxrt.Throw(haxe__Exception_thrown(hxrt.StringFromLiteral("interface view lost conditional class identity")))
	}
	var absent *TaggedError = nil
	if hx_is_nominal_0(absent) || hx_is_nominal_1(absent) {
		hxrt.Throw(haxe__Exception_thrown(hxrt.StringFromLiteral("null matched nominal membership")))
	}
	hxrt.TryCatch(func() {
		hxrt.Throw(haxe__Exception_thrown(base))
	}, func(hx_caught_1 any) {
		hx_exception_2 := haxe__Exception_caught(hx_caught_1)
		hx_payload_3 := haxe__Exception_catchValue(hx_caught_1)
		if hx_is_nominal_2(hx_exception_2) {
			var caught *TaggedError = hx_recover_class_0(hx_exception_2)
			if caught != original {
				hxrt.Throw(haxe__Exception_thrown(hxrt.StringFromLiteral("conditional catch identity changed")))
			}
		} else {
			var hx_tmp any = hx_payload_3
			_ = hx_tmp
			hxrt.Throw(haxe__Exception_thrown(hxrt.StringFromLiteral("conditional subclass catch missed")))
		}
	})
	var hx_try_4 bool
	hxrt.TryCatch(func() {
		hxrt.Throw(haxe__Exception_thrown(base))
		hx_try_4 = false
	}, func(hx_caught_5 any) {
		hx_exception_6 := haxe__Exception_caught(hx_caught_5)
		hx_payload_7 := haxe__Exception_catchValue(hx_caught_5)
		if hx_is_nominal_2(hx_exception_6) {
			var caught_1 *TaggedError = hx_recover_class_0(hx_exception_6)
			hx_try_4 = (caught_1 == original)
		} else {
			var hx_tmp_1 any = hx_payload_7
			_ = hx_tmp_1
			hx_try_4 = false
		}
	})
	same := hx_try_4
	if !same {
		hxrt.Throw(haxe__Exception_thrown(hxrt.StringFromLiteral("conditional expression catch identity changed")))
	}
	hxrt.TryCatch(func() {
		hxrt.TryCatch(func() {
			hxrt.Throw(haxe__Exception_thrown(original))
		}, func(hx_caught_13 any) {
			hx_exception_14 := haxe__Exception_caught(hx_caught_13)
			hx_payload_15 := haxe__Exception_catchValue(hx_caught_13)
			_ = hx_payload_15
			var caught_3 *haxe__Exception = hx_recover_class_1(hx_exception_14)
			hxrt.Throw(haxe__Exception_thrown(caught_3))
		})
	}, func(hx_caught_8 any) {
		hx_exception_9 := haxe__Exception_caught(hx_caught_8)
		hx_payload_10 := haxe__Exception_catchValue(hx_caught_8)
		_ = hx_exception_9
		if hx_is_nominal_0(hx_payload_10) {
			var caught_2 Tag = func(hx_value_11 any) Tag {
				if hx_value_11 == nil {
					var hx_zero_12 Tag
					return hx_zero_12
				}
				return hx_value_11.(Tag)
			}(hxrt.ObjectSource(hx_payload_10))
			if caught_2.code() != 7 {
				hxrt.Throw(haxe__Exception_thrown(hxrt.StringFromLiteral("conditional interface catch changed")))
			}
		} else {
			var hx_tmp_2 any = hx_payload_10
			_ = hx_tmp_2
			hxrt.Throw(haxe__Exception_thrown(hxrt.StringFromLiteral("conditional interface catch missed")))
		}
	})
	return Registry_install()
}

var TaggedError_installed bool = TaggedError_check()
