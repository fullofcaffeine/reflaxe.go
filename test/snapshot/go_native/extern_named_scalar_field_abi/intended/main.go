package main

import (
	"net"
	"net/http"
	"net/url"
	"snapshot/hxrt"
	"time"
)

type I_DurationResult interface {
}

type DurationResult struct {
	__hx_this I_DurationResult
	value1    *time.Duration
	value2    *hxrt.NativeErrorValue
}

func (self *DurationResult) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_DurationResult(value1 *time.Duration, value2 *hxrt.NativeErrorValue) *DurationResult {
	return hx_init_DurationResult(hx_alloc_DurationResult(nil), value1, value2)
}

func hx_alloc_DurationResult(owner any) *DurationResult {
	self := &DurationResult{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_DurationResult)
	}
	return self
}

func hx_init_DurationResult(self *DurationResult, value1 *time.Duration, value2 *hxrt.NativeErrorValue) *DurationResult {
	self.value1 = value1
	self.value2 = value2
	return self
}

func main() {
	parsed := func() *DurationResult {
		hx_tuple_1, hx_tuple_2 := time.ParseDuration(*hxrt.StdString(hxrt.StringFromLiteral("750ms")))
		return New_DurationResult(&hx_tuple_1, hxrt.ImportNativeError(hx_tuple_2))
	}()
	if parsed.value2 != nil {
		hxrt.Throw(haxe__Exception_thrown(hxrt.StringFromLiteral("duration parse failed")))
	}
	dialer := &net.Dialer{}
	dialer.Timeout = *parsed.value1
	roundTrip := func() *time.Duration {
		hx_native_field_3 := observed(dialer).Timeout
		return &hx_native_field_3
	}()
	if !hxrt.StringEqualStringPtr(hxrt.StdString(roundTrip.String()), hxrt.StringFromLiteral("750ms")) {
		hxrt.Throw(haxe__Exception_thrown(hxrt.StringFromLiteral("named scalar field changed value")))
	}
	if reads != 1 {
		hxrt.Throw(haxe__Exception_thrown(hxrt.StringFromLiteral("native field receiver evaluated more than once")))
	}
	reads = 0
	order = hxrt.StringFromLiteral("")
	assigned := func() *time.Duration {
		hx_native_target_4 := observed(dialer)
		hx_native_target_4.Timeout = *supplied(parsed.value1)
		return func() *time.Duration {
			hx_native_field_5 := hx_native_target_4.Timeout
			return &hx_native_field_5
		}()
	}()
	if (reads != 1) || !hxrt.StringEqualStringPtr(order, hxrt.StringFromLiteral("LR")) {
		hxrt.Throw(haxe__Exception_thrown(hxrt.StringFromLiteral("native field assignment must evaluate receiver once before right side")))
	}
	second := func() *DurationResult {
		hx_tuple_6, hx_tuple_7 := time.ParseDuration(*hxrt.StdString(hxrt.StringFromLiteral("1s")))
		return New_DurationResult(&hx_tuple_6, hxrt.ImportNativeError(hx_tuple_7))
	}()
	if second.value2 != nil {
		hxrt.Throw(haxe__Exception_thrown(hxrt.StringFromLiteral("second duration parse failed")))
	}
	dialer.Timeout = *second.value1
	if (!hxrt.StringEqualStringPtr(hxrt.StdString(roundTrip.String()), hxrt.StringFromLiteral("750ms")) || !hxrt.StringEqualStringPtr(hxrt.StdString(assigned.String()), hxrt.StringFromLiteral("750ms"))) || !hxrt.StringEqualStringPtr(hxrt.StdString(func() *time.Duration {
		hx_native_field_8 := dialer.Timeout
		return &hx_native_field_8
	}().String()), hxrt.StringFromLiteral("1s")) {
		hxrt.Throw(haxe__Exception_thrown(hxrt.StringFromLiteral("native value read or assignment result aliased field storage")))
	}
	request := &http.Request{}
	url := &url.URL{}
	url.Path = *hxrt.StdString(hxrt.StringFromLiteral("before"))
	request.URL = url
	pointerRead := request.URL
	url.Path = *hxrt.StdString(hxrt.StringFromLiteral("after"))
	if !hxrt.StringEqualStringPtr(hxrt.StdString(pointerRead.Path), hxrt.StringFromLiteral("after")) || !hxrt.StringEqualStringPtr(hxrt.StdString(request.URL.Path), hxrt.StringFromLiteral("after")) {
		hxrt.Throw(haxe__Exception_thrown(hxrt.StringFromLiteral("unmarked pointer field lost native alias")))
	}
	hxrt.Println(any(hxrt.StringFromLiteral("named scalar field roundtrip: PASS")))
}

func observed(dialer *net.Dialer) *net.Dialer {
	reads = int(int32((reads + 1)))
	order = hxrt.StringConcatStringPtr(order, hxrt.StringFromLiteral("L"))
	return dialer
}

var order *string = hxrt.StringFromLiteral("")

var reads int = 0

func supplied(value *time.Duration) *time.Duration {
	order = hxrt.StringConcatStringPtr(order, hxrt.StringFromLiteral("R"))
	return value
}

func hx_recover_class_0(value any) *haxe__ValueException {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *haxe__ValueException:
		return source
	default:
		return value.(*haxe__ValueException)
	}
}

func hx_recover_class_1(value any) *haxe__Exception {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *haxe__Exception:
		return source
	case *haxe__ValueException:
		return source.haxe__Exception
	default:
		return value.(*haxe__Exception)
	}
}

func hx_is_nominal_0(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *haxe__ValueException:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_1(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *haxe__Exception:
		return (source != nil)
	case *haxe__ValueException:
		return (source != nil)
	default:
		return false
	}
}
