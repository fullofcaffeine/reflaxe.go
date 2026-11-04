package main

import "snapshot/hxrt"

type I_FixedException interface {
	get_message() *string
	get_stack() any
	get_previous() *haxe__Exception
	get_native() any
	unwrap() any
	toString() *string
	details() *string
}

type FixedException struct {
	*haxe__Exception
	__hx_this I_FixedException
	frames    any
}

func (self *FixedException) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_FixedException(message *string, frames any, previous *haxe__Exception) *FixedException {
	return hx_init_FixedException(hx_alloc_FixedException(nil), message, frames, previous)
}

func hx_alloc_FixedException(owner any) *FixedException {
	self := &FixedException{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_FixedException)
	}
	self.haxe__Exception = hx_alloc_haxe__Exception(self.__hx_this)
	return self
}

func hx_init_FixedException(self *FixedException, message *string, frames any, previous *haxe__Exception) *FixedException {
	hx_init_haxe__Exception(self.haxe__Exception, message, previous, nil)
	self.frames = frames
	return self
}

func (self *FixedException) get_stack() any {
	return self.frames
}

func (self *FixedException) String() string {
	return *self.toString()
}

func main() {
	shared := haxe__StackItem_Method(hxrt.StringFromLiteral("Shared"), hxrt.StringFromLiteral("tail"))
	var first any = hxrt.NewArray(haxe__StackItem_Method(hxrt.StringFromLiteral("Cause"), hxrt.StringFromLiteral("start")), shared)
	var second any = hxrt.NewArray(haxe__StackItem_Method(hxrt.StringFromLiteral("Outer"), hxrt.StringFromLiteral("start")), shared)
	var v any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("subtract="), haxe___CallStack__CallStack_Impl__toString(haxe___CallStack__CallStack_Impl__subtract(func(hx_value_1 any) *hxrt.Array {
		if hx_value_1 == nil {
			var hx_zero_2 *hxrt.Array
			return hx_zero_2
		}
		return hx_value_1.(*hxrt.Array)
	}(any(first)), second))))
	hxrt.Println(v)
	cause := New_FixedException(hxrt.StringFromLiteral("cause"), first, nil)
	outer := New_FixedException(hxrt.StringFromLiteral("outer"), second, func(hx_value *FixedException) *haxe__Exception {
		if hx_value == nil {
			return nil
		}
		return hx_value.haxe__Exception
	}(cause))
	var v_1 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("details="), outer.details()))
	hxrt.Println(v_1)
	var identical any = hxrt.NewArray(haxe__StackItem_FilePos(haxe__StackItem_Method(hxrt.StringFromLiteral("Same"), hxrt.StringFromLiteral("f")), hxrt.StringFromLiteral("sample.hx"), 7, 3))
	var equal any = hxrt.NewArray(haxe__StackItem_FilePos(haxe__StackItem_Method(hxrt.StringFromLiteral("Same"), hxrt.StringFromLiteral("f")), hxrt.StringFromLiteral("sample.hx"), 7, 3))
	var v_2 any = any(hxrt.StringConcatAny(hxrt.StringFromLiteral("identical="), func() int {
		this1 := func(hx_value_5 any) *hxrt.Array {
			if hx_value_5 == nil {
				var hx_zero_6 *hxrt.Array
				return hx_zero_6
			}
			return hx_value_5.(*hxrt.Array)
		}(any(haxe___CallStack__CallStack_Impl__subtract(func(hx_value_3 any) *hxrt.Array {
			if hx_value_3 == nil {
				var hx_zero_4 *hxrt.Array
				return hx_zero_4
			}
			return hx_value_3.(*hxrt.Array)
		}(any(identical)), equal)))
		return this1.Len()
	}()))
	hxrt.Println(v_2)
	var different any = hxrt.NewArray(haxe__StackItem_Module(hxrt.StringFromLiteral("elsewhere")))
	var v_3 any = any(hxrt.StringConcatAny(hxrt.StringFromLiteral("distinct="), func() int {
		this1_1 := func(hx_value_9 any) *hxrt.Array {
			if hx_value_9 == nil {
				var hx_zero_10 *hxrt.Array
				return hx_zero_10
			}
			return hx_value_9.(*hxrt.Array)
		}(any(haxe___CallStack__CallStack_Impl__subtract(func(hx_value_7 any) *hxrt.Array {
			if hx_value_7 == nil {
				var hx_zero_8 *hxrt.Array
				return hx_zero_8
			}
			return hx_value_7.(*hxrt.Array)
		}(any(first)), different)))
		return this1_1.Len()
	}()))
	hxrt.Println(v_3)
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
	case *FixedException:
		return source.haxe__Exception
	case *haxe__Exception:
		return source
	case *haxe__ValueException:
		return source.haxe__Exception
	default:
		return value.(*haxe__Exception)
	}
}
