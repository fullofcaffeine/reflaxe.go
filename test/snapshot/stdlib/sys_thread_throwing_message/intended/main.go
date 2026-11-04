package main

import "snapshot/hxrt"

type I_BadMessage interface {
	get_message() *string
	get_stack() any
	get_previous() *haxe__Exception
	get_native() any
	unwrap() any
	toString() *string
	details() *string
}

type BadMessage struct {
	*haxe__Exception
	__hx_this I_BadMessage
}

func (self *BadMessage) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_BadMessage() *BadMessage {
	return hx_init_BadMessage(hx_alloc_BadMessage(nil))
}

func hx_alloc_BadMessage(owner any) *BadMessage {
	self := &BadMessage{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_BadMessage)
	}
	self.haxe__Exception = hx_alloc_haxe__Exception(self.__hx_this)
	return self
}

func hx_init_BadMessage(self *BadMessage) *BadMessage {
	hx_init_haxe__Exception(self.haxe__Exception, hxrt.StringFromLiteral("original"), nil, nil)
	return self
}

func (self *BadMessage) get_message() *string {
	hxrt.Throw(haxe__Exception_thrown(hxrt.StringFromLiteral("message failed")))
	var hx_throw_zero_1 *string
	return hx_throw_zero_1
}

func (self *BadMessage) String() string {
	return *self.toString()
}

func main() {
	defer hxrt.ThreadWaitForAll()
	sys__thread__Thread_create(hxrt.NewFunction(func() {
		hxrt.Throw(haxe__Exception_thrown(New_BadMessage()))
	}))
	sys__thread__Thread_createWithEventLoop(hxrt.NewFunction(func() {
		hxrt.Throw(haxe__Exception_thrown(New_BadMessage()))
	}))
	hxrt.Println(any(hxrt.StringFromLiteral("main-survived")))
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
	case *BadMessage:
		return source.haxe__Exception
	case *haxe__ValueException:
		return source.haxe__Exception
	case *haxe__Exception:
		return source
	case *sys__thread__NoEventLoopException:
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
	case *BadMessage:
		return (source != nil)
	case *haxe__ValueException:
		return (source != nil)
	case *haxe__Exception:
		return (source != nil)
	case *sys__thread__NoEventLoopException:
		return (source != nil)
	default:
		return false
	}
}
