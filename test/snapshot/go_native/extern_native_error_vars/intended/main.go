package main

import (
	"context"
	"errors"
	"image"
	"io"
	"snapshot/hxrt"
	"time"
)

type I_CancelPair interface {
}

type CancelPair struct {
	__hx_this I_CancelPair
	context   context.Context
	cancel    *hxrt.Function[func(*hxrt.NativeErrorValue)]
}

func (self *CancelPair) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_CancelPair(context_local context.Context, cancel *hxrt.Function[func(*hxrt.NativeErrorValue)]) *CancelPair {
	return hx_init_CancelPair(hx_alloc_CancelPair(nil), context_local, cancel)
}

func hx_alloc_CancelPair(owner any) *CancelPair {
	self := &CancelPair{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_CancelPair)
	}
	return self
}

func hx_init_CancelPair(self *CancelPair, context_local context.Context, cancel *hxrt.Function[func(*hxrt.NativeErrorValue)]) *CancelPair {
	self.context = context_local
	self.cancel = cancel
	return self
}

func (self *CancelPair) HxClassName() string {
	return "CancelPair"
}

type I_Holder interface {
}

type Holder struct {
	__hx_this I_Holder
	error     *hxrt.NativeErrorValue
}

func (self *Holder) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_Holder(error *hxrt.NativeErrorValue) *Holder {
	return hx_init_Holder(hx_alloc_Holder(nil), error)
}

func hx_alloc_Holder(owner any) *Holder {
	self := &Holder{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_Holder)
	}
	return self
}

func hx_init_Holder(self *Holder, error *hxrt.NativeErrorValue) *Holder {
	self.error = error
	return self
}

func (self *Holder) HxClassName() string {
	return "Holder"
}

func check(ok bool, label *string) {
	if !ok {
		hxrt.Throw(haxe__Exception_thrown(label))
	}
}

func describe(error *hxrt.NativeErrorValue) *string {
	var hx_if_1 *string
	if error == nil {
		hx_if_1 = hxrt.StringFromLiteral("nil")
	} else {
		hx_if_1 = hxrt.StdString(error.Text())
	}
	return hx_if_1
}

var last *hxrt.NativeErrorValue = nil

func main() {
	eof := hxrt.ImportNativeError(io.EOF)
	var v any = any(hxrt.StdString(eof.Text()))
	hxrt.Println(v)
	var v_1 any = any(errors.Is(hxrt.ExportNativeError(eof), io.EOF))
	hxrt.Println(v_1)
	var v_2 any = any(hxrt.StdString(hxrt.ImportNativeError(io.EOF).Text()))
	hxrt.Println(v_2)
	check(errors.Is(io.EOF, io.EOF), hxrt.StringFromLiteral("direct identity"))
	check(!errors.Is(io.EOF, context.Canceled), hxrt.StringFromLiteral("distinct variables"))
	check(((hxrt.ImportNativeError(io.EOF) != nil) && (hxrt.ImportNativeError(io.EOF) != nil)), hxrt.StringFromLiteral("non-null read"))
	var v_3 any = any(describe(hxrt.ImportNativeError(context.Canceled)))
	hxrt.Println(v_3)
	holder := New_Holder(hxrt.ImportNativeError(io.EOF))
	check(errors.Is(hxrt.ExportNativeError(holder.error), io.EOF), hxrt.StringFromLiteral("field storage"))
	holder.error = hxrt.ImportNativeError(context.Canceled)
	check(errors.Is(hxrt.ExportNativeError(holder.error), context.Canceled), hxrt.StringFromLiteral("field write"))
	last = hxrt.ImportNativeError(io.EOF)
	check(errors.Is(hxrt.ExportNativeError(last), hxrt.ExportNativeError(eof)), hxrt.StringFromLiteral("static storage"))
	stored := hxrt.NewArray(hxrt.ImportNativeError(io.EOF), hxrt.ImportNativeError(context.Canceled))
	stored.Push(hxrt.ImportNativeError(io.EOF))
	check(((((stored.Len() == 3) && errors.Is(hxrt.ExportNativeError(func(hx_value_3 any) *hxrt.NativeErrorValue {
		if hx_value_3 == nil {
			var hx_zero_4 *hxrt.NativeErrorValue
			return hx_zero_4
		}
		return hx_value_3.(*hxrt.NativeErrorValue)
	}(stored.Get(0))), io.EOF)) && errors.Is(hxrt.ExportNativeError(func(hx_value_5 any) *hxrt.NativeErrorValue {
		if hx_value_5 == nil {
			var hx_zero_6 *hxrt.NativeErrorValue
			return hx_zero_6
		}
		return hx_value_5.(*hxrt.NativeErrorValue)
	}(stored.Get(1))), context.Canceled)) && errors.Is(hxrt.ExportNativeError(func(hx_value_7 any) *hxrt.NativeErrorValue {
		if hx_value_7 == nil {
			var hx_zero_8 *hxrt.NativeErrorValue
			return hx_zero_8
		}
		return hx_value_7.(*hxrt.NativeErrorValue)
	}(stored.Get(2))), hxrt.ExportNativeError(eof))), hxrt.StringFromLiteral("array storage"))
	joined := hxrt.ImportNativeError(errors.Join([]hxrt.RawError{io.EOF, context.Canceled}...))
	check((errors.Is(hxrt.ExportNativeError(joined), io.EOF) && errors.Is(hxrt.ExportNativeError(joined), context.Canceled)), hxrt.StringFromLiteral("variadic native argument"))
	pair := func() *CancelPair {
		hx_tuple_9, hx_tuple_10 := context.WithCancelCause(context.Background())
		return New_CancelPair(hx_tuple_9, func() *hxrt.Function[func(*hxrt.NativeErrorValue)] {
			hx_extern_method_11 := hx_tuple_10
			if hx_extern_method_11 == nil {
				return nil
			}
			return hxrt.NewFunction(func(hx_extern_arg_12 *hxrt.NativeErrorValue) {
				hx_extern_method_11(hxrt.ExportNativeError(hx_extern_arg_12))
			})
		}())
	}()
	func() {
		var hx_callable_13 *hxrt.Function[func(*hxrt.NativeErrorValue)] = pair.cancel
		var hx_argument_14 *hxrt.NativeErrorValue = nil
		hxrt.FunctionCall(hx_callable_13)(hx_argument_14)
	}()
	check(errors.Is(context.Cause(pair.context), context.Canceled), hxrt.StringFromLiteral("native result matches variable"))
	original := hxrt.ImportNativeError(errors.ErrUnsupported)
	replacement := hxrt.ImportNativeError(errors.New(*hxrt.StdString(hxrt.StringFromLiteral("replacement"))))
	errors.ErrUnsupported = hxrt.ExportNativeError(replacement)
	check(errors.Is(errors.ErrUnsupported, hxrt.ExportNativeError(replacement)), hxrt.StringFromLiteral("variable write"))
	restored := func() *hxrt.NativeErrorValue {
		errors.ErrUnsupported = hxrt.ExportNativeError(original)
		return hxrt.ImportNativeError(errors.ErrUnsupported)
	}()
	check((errors.Is(hxrt.ExportNativeError(restored), hxrt.ExportNativeError(original)) && errors.Is(errors.ErrUnsupported, hxrt.ExportNativeError(original))), hxrt.StringFromLiteral("variable write result"))
	errors.ErrUnsupported = hxrt.ExportNativeError(nil)
	check((hxrt.ImportNativeError(errors.ErrUnsupported) == nil), hxrt.StringFromLiteral("nil variable write"))
	errors.ErrUnsupported = hxrt.ExportNativeError(original)
	origin := func() *image.Point {
		hx_native_field_15 := image.ZP
		return &hx_native_field_15
	}()
	origin.X = 7
	check((((origin.X == 7) && (func() *image.Point {
		hx_native_field_16 := image.ZP
		return &hx_native_field_16
	}().X == 0)) && (sum(func() *image.Point {
		hx_native_field_17 := image.ZP
		return &hx_native_field_17
	}()) == 0)), hxrt.StringFromLiteral("struct value variable copy"))
	moved := &image.Point{}
	moved.Y = 5
	image.ZP = *moved
	moved.Y = 9
	check((sum(func() *image.Point {
		hx_native_field_18 := image.ZP
		return &hx_native_field_18
	}()) == 5), hxrt.StringFromLiteral("struct value variable write"))
	image.ZP = *&image.Point{}
	hxrt.Println(any(hxrt.StdString(time.RFC3339)))
	check((measure(hxrt.StdString(time.RFC3339)) == 25), hxrt.StringFromLiteral("string variable"))
	hxrt.Println(any(hxrt.StringFromLiteral("native-error-vars:ok")))
}

func measure(text *string) int {
	return hxrt.StringLengthStringPtr(text)
}

func sum(point *image.Point) int {
	return int((hxrt.Int32Wrap(point.X) + hxrt.Int32Wrap(point.Y)))
}

func hx_recover_class_haxe__Exception(value any) *haxe__Exception {
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

func hx_recover_class_haxe__ValueException(value any) *haxe__ValueException {
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

func hx_is_nominal_haxe__Exception(value any) bool {
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

func hx_is_nominal_haxe__ValueException(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *haxe__ValueException:
		return (source != nil)
	default:
		return false
	}
}
