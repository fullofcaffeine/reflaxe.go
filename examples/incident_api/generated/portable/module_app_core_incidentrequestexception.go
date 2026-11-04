package main

type I_app__core__IncidentRequestException interface {
	get_message() *string
	get_stack() any
	get_previous() *haxe__Exception
	get_native() any
	unwrap() any
	toString() *string
	details() *string
}

type app__core__IncidentRequestException struct {
	*haxe__Exception
	__hx_this I_app__core__IncidentRequestException
	code      *string
}

func (self *app__core__IncidentRequestException) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_app__core__IncidentRequestException(code *string) *app__core__IncidentRequestException {
	return hx_init_app__core__IncidentRequestException(hx_alloc_app__core__IncidentRequestException(nil), code)
}

func hx_alloc_app__core__IncidentRequestException(owner any) *app__core__IncidentRequestException {
	self := &app__core__IncidentRequestException{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_app__core__IncidentRequestException)
	}
	self.haxe__Exception = hx_alloc_haxe__Exception(self.__hx_this)
	return self
}

func hx_init_app__core__IncidentRequestException(self *app__core__IncidentRequestException, code *string) *app__core__IncidentRequestException {
	hx_init_haxe__Exception(self.haxe__Exception, code, nil, nil)
	self.code = code
	return self
}

func (self *app__core__IncidentRequestException) String() string {
	return *self.toString()
}
