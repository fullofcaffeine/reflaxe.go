package main

import "snapshot/hxrt"

type I_haxe__Exception interface {
	get_message() *string
	get_stack() any
	get_previous() *haxe__Exception
	get_native() any
	unwrap() any
	toString() *string
	details() *string
}

type haxe__Exception struct {
	__hx_this     I_haxe__Exception
	message       *string
	stack         any
	previous      *haxe__Exception
	native        any
	savedMessage  *string
	savedStack    any
	savedPrevious *haxe__Exception
	savedNative   any
}

func (self *haxe__Exception) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_haxe__Exception(message *string, previous *haxe__Exception, native any) *haxe__Exception {
	return hx_init_haxe__Exception(hx_alloc_haxe__Exception(nil), message, previous, native)
}

func hx_alloc_haxe__Exception(owner any) *haxe__Exception {
	self := &haxe__Exception{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_haxe__Exception)
	}
	return self
}

func hx_init_haxe__Exception(self *haxe__Exception, message *string, previous *haxe__Exception, native any) *haxe__Exception {
	self.savedMessage = message
	self.savedPrevious = previous
	var hx_if_1 any
	if hxrt.AnyEqualsNull(native) {
		hx_if_1 = self
	} else {
		hx_if_1 = native
	}
	self.savedNative = hx_if_1
	self.savedStack = haxe___CallStack__CallStack_Impl__callStack()
	return self
}

func (self *haxe__Exception) get_message() *string {
	return self.savedMessage
}

func (self *haxe__Exception) get_stack() any {
	return self.savedStack
}

func (self *haxe__Exception) get_previous() *haxe__Exception {
	return self.savedPrevious
}

func (self *haxe__Exception) get_native() any {
	return self.savedNative
}

func (self *haxe__Exception) unwrap() any {
	return self.savedNative
}

func (self *haxe__Exception) toString() *string {
	return self.__hx_this.get_message()
}

func (self *haxe__Exception) details() *string {
	return haxe___CallStack__CallStack_Impl__exceptionToString(self)
}

func (self *haxe__Exception) String() string {
	return *self.__hx_this.toString()
}

func haxe__Exception_catchValue(value any) any {
	if hx_is_nominal_2(value) {
		wrapped := hx_recover_class_2(hx_recover_class_2(value))
		return wrapped.value
	}
	return value
}

func haxe__Exception_caught(value any) *haxe__Exception {
	if hx_is_nominal_1(value) {
		return hx_recover_class_3(value)
	}
	return func(hx_value *haxe__ValueException) *haxe__Exception {
		if hx_value == nil {
			return nil
		}
		return hx_value.haxe__Exception
	}(New_haxe__ValueException(value, nil, value))
}

func haxe__Exception_thrown(value any) any {
	if hx_is_nominal_1(value) {
		return value
	}
	return New_haxe__ValueException(value, nil, nil)
}
