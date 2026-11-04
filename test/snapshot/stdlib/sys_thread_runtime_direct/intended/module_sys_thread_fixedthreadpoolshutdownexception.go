package main

type I_sys__thread__FixedThreadPoolShutdownException interface {
	get_message() *string
	get_stack() any
	get_previous() *haxe__Exception
	get_native() any
	unwrap() any
	toString() *string
	details() *string
}

type sys__thread__FixedThreadPoolShutdownException struct {
	*haxe__Exception
	__hx_this I_sys__thread__FixedThreadPoolShutdownException
}

func (self *sys__thread__FixedThreadPoolShutdownException) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_sys__thread__FixedThreadPoolShutdownException(message *string, previous *haxe__Exception, native any) *sys__thread__FixedThreadPoolShutdownException {
	return hx_init_sys__thread__FixedThreadPoolShutdownException(hx_alloc_sys__thread__FixedThreadPoolShutdownException(nil), message, previous, native)
}

func hx_alloc_sys__thread__FixedThreadPoolShutdownException(owner any) *sys__thread__FixedThreadPoolShutdownException {
	self := &sys__thread__FixedThreadPoolShutdownException{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_sys__thread__FixedThreadPoolShutdownException)
	}
	self.haxe__Exception = hx_alloc_haxe__Exception(self.__hx_this)
	return self
}

func hx_init_sys__thread__FixedThreadPoolShutdownException(self *sys__thread__FixedThreadPoolShutdownException, message *string, previous *haxe__Exception, native any) *sys__thread__FixedThreadPoolShutdownException {
	hx_init_haxe__Exception(self.haxe__Exception, message, previous, native)
	return self
}

func (self *sys__thread__FixedThreadPoolShutdownException) String() string {
	return *self.__hx_this.toString()
}
