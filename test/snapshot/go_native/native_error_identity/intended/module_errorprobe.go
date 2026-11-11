//go:build go1.26

package main

import (
	"context"
	"errors"
	"fmt"
	"io/fs"
	"path/filepath"
	"snapshot/hxrt"
	"sync"
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

func New_CancelPair(context context.Context, cancel *hxrt.Function[func(*hxrt.NativeErrorValue)]) *CancelPair {
	return hx_init_CancelPair(hx_alloc_CancelPair(nil), context, cancel)
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

func hx_init_CancelPair(self *CancelPair, context context.Context, cancel *hxrt.Function[func(*hxrt.NativeErrorValue)]) *CancelPair {
	self.context = context
	self.cancel = cancel
	return self
}

func ErrorProbe_check(ok bool, label *string) {
	if !ok {
		hxrt.Throw(haxe__Exception_thrown(label))
	}
}

func ErrorProbe_run() {
	root := hxrt.ImportNativeError(errors.New(*hxrt.StdString(hxrt.StringFromLiteral("root"))))
	ErrorProbe_check(hxrt.StringEqualStringPtr(hxrt.StdString(root.Text()), hxrt.StringFromLiteral("root")), hxrt.StringFromLiteral("explicit text"))
	ErrorProbe_check(errors.Is(hxrt.ExportNativeError(root), hxrt.ExportNativeError(root)), hxrt.StringFromLiteral("direct identity"))
	ErrorProbe_check(!errors.Is(hxrt.ExportNativeError(root), hxrt.ExportNativeError(hxrt.ImportNativeError(errors.New(*hxrt.StdString(hxrt.StringFromLiteral("root")))))), hxrt.StringFromLiteral("distinct equal text"))
	ErrorProbe_check(errors.Is(hxrt.ExportNativeError(nil), hxrt.ExportNativeError(nil)), hxrt.StringFromLiteral("nil export"))
	create := func() *hxrt.Function[func(*string) *hxrt.NativeErrorValue] {
		hx_extern_method_1 := errors.New
		if hx_extern_method_1 == nil {
			return nil
		}
		return hxrt.NewFunction(func(hx_extern_arg_2 *string) *hxrt.NativeErrorValue {
			return hxrt.ImportNativeError(hx_extern_method_1(*hxrt.StdString(hx_extern_arg_2)))
		})
	}()
	matches := func() *hxrt.Function[func(*hxrt.NativeErrorValue, *hxrt.NativeErrorValue) bool] {
		hx_extern_method_3 := errors.Is
		if hx_extern_method_3 == nil {
			return nil
		}
		return hxrt.NewFunction(func(hx_extern_arg_4 *hxrt.NativeErrorValue, hx_extern_arg_5 *hxrt.NativeErrorValue) bool {
			return hx_extern_method_3(hxrt.ExportNativeError(hx_extern_arg_4), hxrt.ExportNativeError(hx_extern_arg_5))
		})
	}()
	stored := func() *hxrt.NativeErrorValue {
		var hx_callable_6 *hxrt.Function[func(*string) *hxrt.NativeErrorValue] = create
		var hx_argument_7 *string = hxrt.StringFromLiteral("stored")
		return hxrt.FunctionCall(hx_callable_6)(hx_argument_7)
	}()
	ErrorProbe_check(func() bool {
		var hx_callable_8 *hxrt.Function[func(*hxrt.NativeErrorValue, *hxrt.NativeErrorValue) bool] = matches
		var hx_argument_9 *hxrt.NativeErrorValue = stored
		var hx_argument_10 *hxrt.NativeErrorValue = stored
		return hxrt.FunctionCall(hx_callable_8)(hx_argument_9, hx_argument_10)
	}(), hxrt.StringFromLiteral("stored native function"))
	pair := func() *CancelPair {
		hx_tuple_11, hx_tuple_12 := context.WithCancelCause(context.Background())
		return New_CancelPair(hx_tuple_11, func() *hxrt.Function[func(*hxrt.NativeErrorValue)] {
			hx_extern_method_13 := hx_tuple_12
			if hx_extern_method_13 == nil {
				return nil
			}
			return hxrt.NewFunction(func(hx_extern_arg_14 *hxrt.NativeErrorValue) {
				hx_extern_method_13(hxrt.ExportNativeError(hx_extern_arg_14))
			})
		}())
	}()
	ErrorProbe_check((hxrt.ImportNativeError(context.Cause(pair.context)) == nil), hxrt.StringFromLiteral("nil scalar result"))
	func() {
		var hx_callable_15 *hxrt.Function[func(*hxrt.NativeErrorValue)] = pair.cancel
		var hx_argument_16 *hxrt.NativeErrorValue = root
		hxrt.FunctionCall(hx_callable_15)(hx_argument_16)
	}()
	ErrorProbe_check(func() bool {
		var hx_callable_17 *hxrt.Function[func(*hxrt.NativeErrorValue, *hxrt.NativeErrorValue) bool] = matches
		var hx_argument_18 *hxrt.NativeErrorValue = hxrt.ImportNativeError(context.Cause(pair.context))
		var hx_argument_19 *hxrt.NativeErrorValue = root
		return hxrt.FunctionCall(hx_callable_17)(hx_argument_18, hx_argument_19)
	}(), hxrt.StringFromLiteral("function inside native tuple"))
	once := func() *hxrt.Function[func() *hxrt.NativeErrorValue] {
		hx_extern_method_21 := sync.OnceValue(func(hx_native_callback_20 *hxrt.Function[func() *hxrt.NativeErrorValue]) func() hxrt.RawError {
			if hx_native_callback_20 == nil {
				return nil
			}
			return func() hxrt.RawError {
				return hxrt.ExportNativeError(hxrt.FunctionCall(hx_native_callback_20)())
			}
		}(hxrt.NewFunction(func() *hxrt.NativeErrorValue {
			return root
		})))
		if hx_extern_method_21 == nil {
			return nil
		}
		return hxrt.NewFunction(func() *hxrt.NativeErrorValue {
			return hxrt.ImportNativeError(hx_extern_method_21())
		})
	}()
	ErrorProbe_check(func() bool {
		var hx_callable_23 *hxrt.Function[func(*hxrt.NativeErrorValue, *hxrt.NativeErrorValue) bool] = matches
		var hx_argument_24 *hxrt.NativeErrorValue = func() *hxrt.NativeErrorValue {
			var hx_callable_22 *hxrt.Function[func() *hxrt.NativeErrorValue] = once
			return hxrt.FunctionCall(hx_callable_22)()
		}()
		var hx_argument_25 *hxrt.NativeErrorValue = root
		return hxrt.FunctionCall(hx_callable_23)(hx_argument_24, hx_argument_25)
	}(), hxrt.StringFromLiteral("returned generic callback"))
	var observed *hxrt.NativeErrorValue = nil
	failure := hxrt.ImportNativeError(filepath.Walk(*hxrt.StdString(hxrt.StringFromLiteral("\x00")), func(hx_native_callback_26 *hxrt.Function[func(*string, fs.FileInfo, *hxrt.NativeErrorValue) *hxrt.NativeErrorValue]) func(string, fs.FileInfo, hxrt.RawError) hxrt.RawError {
		if hx_native_callback_26 == nil {
			return nil
		}
		return func(hx_native_arg_27 string, hx_native_arg_28 fs.FileInfo, hx_native_arg_29 hxrt.RawError) hxrt.RawError {
			return hxrt.ExportNativeError(hxrt.FunctionCall(hx_native_callback_26)(hxrt.StdString(hx_native_arg_27), hx_native_arg_28, hxrt.ImportNativeError(hx_native_arg_29)))
		}
	}(hxrt.NewFunction(func(hx_tmp *string, _1 fs.FileInfo, error *hxrt.NativeErrorValue) *hxrt.NativeErrorValue {
		observed = error
		return error
	}))))
	ErrorProbe_check(((observed != nil) && func() bool {
		var hx_callable_30 *hxrt.Function[func(*hxrt.NativeErrorValue, *hxrt.NativeErrorValue) bool] = matches
		var hx_argument_31 *hxrt.NativeErrorValue = failure
		var hx_argument_32 *hxrt.NativeErrorValue = observed
		return hxrt.FunctionCall(hx_callable_30)(hx_argument_31, hx_argument_32)
	}()), hxrt.StringFromLiteral("callback arguments and return"))
	typed := func() *Match {
		hx_tuple_33, hx_tuple_34 := errors.AsType[*fs.PathError](hxrt.ExportNativeError(failure))
		return New_Match(hx_tuple_33, hx_tuple_34)
	}()
	ErrorProbe_check((typed.matched && !hxrt.AnyEqualsNull(func(hx_value_35 any) *fs.PathError {
		if hx_value_35 == nil {
			var hx_zero_36 *fs.PathError
			return hx_zero_36
		}
		return hx_value_35.(*fs.PathError)
	}(typed.value))), hxrt.StringFromLiteral("concrete matching"))
	func(hx_value_37 any) *fs.PathError {
		if hx_value_37 == nil {
			var hx_zero_38 *fs.PathError
			return hx_zero_38
		}
		return hx_value_37.(*fs.PathError)
	}(typed.value).Err = hxrt.ExportNativeError(root)
	ErrorProbe_check((func() bool {
		var hx_callable_41 *hxrt.Function[func(*hxrt.NativeErrorValue, *hxrt.NativeErrorValue) bool] = matches
		var hx_argument_42 *hxrt.NativeErrorValue = hxrt.ImportNativeError(func(hx_value_39 any) *fs.PathError {
			if hx_value_39 == nil {
				var hx_zero_40 *fs.PathError
				return hx_zero_40
			}
			return hx_value_39.(*fs.PathError)
		}(typed.value).Err)
		var hx_argument_43 *hxrt.NativeErrorValue = root
		return hxrt.FunctionCall(hx_callable_41)(hx_argument_42, hx_argument_43)
	}() && func() bool {
		var hx_callable_44 *hxrt.Function[func(*hxrt.NativeErrorValue, *hxrt.NativeErrorValue) bool] = matches
		var hx_argument_45 *hxrt.NativeErrorValue = failure
		var hx_argument_46 *hxrt.NativeErrorValue = root
		return hxrt.FunctionCall(hx_callable_44)(hx_argument_45, hx_argument_46)
	}()), hxrt.StringFromLiteral("native field read/write"))
	generic := func() *Match {
		hx_tuple_47, hx_tuple_48 := errors.AsType[hxrt.RawError](hxrt.ExportNativeError(root))
		return New_Match(hxrt.ImportNativeError(hx_tuple_47), hx_tuple_48)
	}()
	ErrorProbe_check((generic.matched && func() bool {
		var hx_callable_51 *hxrt.Function[func(*hxrt.NativeErrorValue, *hxrt.NativeErrorValue) bool] = matches
		var hx_argument_52 *hxrt.NativeErrorValue = func(hx_value_49 any) *hxrt.NativeErrorValue {
			if hx_value_49 == nil {
				var hx_zero_50 *hxrt.NativeErrorValue
				return hx_zero_50
			}
			return hx_value_49.(*hxrt.NativeErrorValue)
		}(generic.value)
		var hx_argument_53 *hxrt.NativeErrorValue = root
		return hxrt.FunctionCall(hx_callable_51)(hx_argument_52, hx_argument_53)
	}()), hxrt.StringFromLiteral("native generic type argument and tuple slot"))
	asNative := hxrt.NewFunction(func(hx_extern_arg_55 *hxrt.NativeErrorValue) *Match {
		return func() *Match {
			hx_tuple_56, hx_tuple_57 := errors.AsType[hxrt.RawError](hxrt.ExportNativeError(hx_extern_arg_55))
			return New_Match(hxrt.ImportNativeError(hx_tuple_56), hx_tuple_57)
		}()
	})
	storedMatch := func() *Match {
		var hx_callable_58 *hxrt.Function[func(*hxrt.NativeErrorValue) *Match] = asNative
		var hx_argument_59 *hxrt.NativeErrorValue = root
		return hxrt.FunctionCall(hx_callable_58)(hx_argument_59)
	}()
	ErrorProbe_check((storedMatch.matched && func() bool {
		var hx_callable_62 *hxrt.Function[func(*hxrt.NativeErrorValue, *hxrt.NativeErrorValue) bool] = matches
		var hx_argument_63 *hxrt.NativeErrorValue = func(hx_value_60 any) *hxrt.NativeErrorValue {
			if hx_value_60 == nil {
				var hx_zero_61 *hxrt.NativeErrorValue
				return hx_zero_61
			}
			return hx_value_60.(*hxrt.NativeErrorValue)
		}(storedMatch.value)
		var hx_argument_64 *hxrt.NativeErrorValue = root
		return hxrt.FunctionCall(hx_callable_62)(hx_argument_63, hx_argument_64)
	}()), hxrt.StringFromLiteral("stored generic tuple function"))
	withCause := func() *hxrt.Function[func(context.Context) *CancelPair] {
		hx_extern_method_65 := context.WithCancelCause
		if hx_extern_method_65 == nil {
			return nil
		}
		return hxrt.NewFunction(func(hx_extern_arg_66 context.Context) *CancelPair {
			return func() *CancelPair {
				hx_tuple_67, hx_tuple_68 := hx_extern_method_65(hx_extern_arg_66)
				return New_CancelPair(hx_tuple_67, func() *hxrt.Function[func(*hxrt.NativeErrorValue)] {
					hx_extern_method_69 := hx_tuple_68
					if hx_extern_method_69 == nil {
						return nil
					}
					return hxrt.NewFunction(func(hx_extern_arg_70 *hxrt.NativeErrorValue) {
						hx_extern_method_69(hxrt.ExportNativeError(hx_extern_arg_70))
					})
				}())
			}()
		})
	}()
	storedPair := func() *CancelPair {
		var hx_callable_71 *hxrt.Function[func(context.Context) *CancelPair] = withCause
		var hx_argument_72 context.Context = context.Background()
		return hxrt.FunctionCall(hx_callable_71)(hx_argument_72)
	}()
	func() {
		var hx_callable_73 *hxrt.Function[func(*hxrt.NativeErrorValue)] = storedPair.cancel
		var hx_argument_74 *hxrt.NativeErrorValue = root
		hxrt.FunctionCall(hx_callable_73)(hx_argument_74)
	}()
	ErrorProbe_check(func() bool {
		var hx_callable_75 *hxrt.Function[func(*hxrt.NativeErrorValue, *hxrt.NativeErrorValue) bool] = matches
		var hx_argument_76 *hxrt.NativeErrorValue = hxrt.ImportNativeError(context.Cause(storedPair.context))
		var hx_argument_77 *hxrt.NativeErrorValue = root
		return hxrt.FunctionCall(hx_callable_75)(hx_argument_76, hx_argument_77)
	}(), hxrt.StringFromLiteral("stored tuple function"))
	var nilPointer *fs.PathError = nil
	wrappedNil := hxrt.ImportNativeError(fmt.Errorf(*hxrt.StdString(hxrt.StringFromLiteral("typed nil: %w")), []any{nilPointer}...))
	typedNil := hxrt.ImportNativeError(errors.Unwrap(hxrt.ExportNativeError(wrappedNil)))
	nilMatch := func() *Match {
		hx_tuple_78, hx_tuple_79 := errors.AsType[*fs.PathError](hxrt.ExportNativeError(typedNil))
		return New_Match(hx_tuple_78, hx_tuple_79)
	}()
	ErrorProbe_check((nilMatch.matched && hxrt.AnyEqualsNull(func(hx_value_80 any) *fs.PathError {
		if hx_value_80 == nil {
			var hx_zero_81 *fs.PathError
			return hx_zero_81
		}
		return hx_value_80.(*fs.PathError)
	}(nilMatch.value))), hxrt.StringFromLiteral("typed nil matches independently of value"))
	ErrorProbe_check((((typedNil != nil) && !hxrt.AnyEqualsNull(func(hx_value_82 any) *hxrt.NativeErrorValue {
		if hx_value_82 == nil {
			var hx_zero_83 *hxrt.NativeErrorValue
			return hx_zero_83
		}
		return hx_value_82.(*hxrt.NativeErrorValue)
	}(New_Holder(typedNil).value))) && hxrt.AnyEqualsNull(func(hx_value_84 any) *fs.PathError {
		if hx_value_84 == nil {
			var hx_zero_85 *fs.PathError
			return hx_zero_85
		}
		return hx_value_84.(*fs.PathError)
	}(New_Holder(nilPointer).value))), hxrt.StringFromLiteral("typed nil through generic storage"))
	holder := New_Holder(root)
	ErrorProbe_check((!hxrt.AnyEqualsNull(func(hx_value_86 any) *hxrt.NativeErrorValue {
		if hx_value_86 == nil {
			var hx_zero_87 *hxrt.NativeErrorValue
			return hx_zero_87
		}
		return hx_value_86.(*hxrt.NativeErrorValue)
	}(holder.value)) && func() bool {
		var hx_callable_90 *hxrt.Function[func(*hxrt.NativeErrorValue, *hxrt.NativeErrorValue) bool] = matches
		var hx_argument_91 *hxrt.NativeErrorValue = func(hx_value_88 any) *hxrt.NativeErrorValue {
			if hx_value_88 == nil {
				var hx_zero_89 *hxrt.NativeErrorValue
				return hx_zero_89
			}
			return hx_value_88.(*hxrt.NativeErrorValue)
		}(holder.value)
		var hx_argument_92 *hxrt.NativeErrorValue = root
		return hxrt.FunctionCall(hx_callable_90)(hx_argument_91, hx_argument_92)
	}()), hxrt.StringFromLiteral("generic storage"))
	legacy := New_go___Error(hxrt.StringFromLiteral("legacy"))
	ErrorProbe_check((hxrt.StringEqualStringPtr(legacy.message, hxrt.StringFromLiteral("legacy")) && hxrt.StringEqualStringPtr(legacy.toString(), hxrt.StringFromLiteral("legacy"))), hxrt.StringFromLiteral("legacy message"))
	ErrorProbe_check(hxrt.StringEqualStringPtr(func(hx_value_93 any) *string {
		if hx_value_93 == nil {
			var hx_zero_94 *string
			return hx_zero_94
		}
		return hx_value_93.(*string)
	}(go___Result_failure(hxrt.StringFromLiteral("old")).error()), hxrt.StringFromLiteral("old")), hxrt.StringFromLiteral("legacy result"))
	hxrt.Println(any(hxrt.StringFromLiteral("native-error:ok")))
}

type I_Holder interface {
}

type Holder struct {
	__hx_this I_Holder
	value     any
}

func (self *Holder) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_Holder(value any) *Holder {
	return hx_init_Holder(hx_alloc_Holder(nil), value)
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

func hx_init_Holder(self *Holder, value any) *Holder {
	self.value = value
	return self
}

type I_Match interface {
}

type Match struct {
	__hx_this I_Match
	value     any
	matched   bool
}

func (self *Match) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_Match(value any, matched bool) *Match {
	return hx_init_Match(hx_alloc_Match(nil), value, matched)
}

func hx_alloc_Match(owner any) *Match {
	self := &Match{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_Match)
	}
	return self
}

func hx_init_Match(self *Match, value any, matched bool) *Match {
	self.value = value
	self.matched = matched
	return self
}
