package main

type I_sys__thread__NoEventLoopException interface {
	get_message() *string
	get_stack() any
	get_previous() *haxe__Exception
	get_native() any
	unwrap() any
	toString() *string
	details() *string
}

type sys__thread__NoEventLoopException struct {
	*haxe__Exception
	__hx_this I_sys__thread__NoEventLoopException
}

func (self *sys__thread__NoEventLoopException) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_sys__thread__NoEventLoopException(msg *string, previous *haxe__Exception) *sys__thread__NoEventLoopException {
	return hx_init_sys__thread__NoEventLoopException(hx_alloc_sys__thread__NoEventLoopException(nil), msg, previous)
}

func hx_alloc_sys__thread__NoEventLoopException(owner any) *sys__thread__NoEventLoopException {
	self := &sys__thread__NoEventLoopException{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_sys__thread__NoEventLoopException)
	}
	self.haxe__Exception = hx_alloc_haxe__Exception(self.__hx_this)
	return self
}

func hx_init_sys__thread__NoEventLoopException(self *sys__thread__NoEventLoopException, msg *string, previous *haxe__Exception) *sys__thread__NoEventLoopException {
	hx_init_haxe__Exception(self.haxe__Exception, msg, previous, nil)
	return self
}

func (self *sys__thread__NoEventLoopException) String() string {
	return *self.__hx_this.toString()
}
