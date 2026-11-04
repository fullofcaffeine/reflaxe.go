package main

type I_app__http__HttpRequest interface {
}

type app__http__HttpRequest struct {
	__hx_this I_app__http__HttpRequest
	method    *string
	path      *string
	body      *string
}

func (self *app__http__HttpRequest) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_app__http__HttpRequest(method *string, path *string, body *string) *app__http__HttpRequest {
	return hx_init_app__http__HttpRequest(hx_alloc_app__http__HttpRequest(nil), method, path, body)
}

func hx_alloc_app__http__HttpRequest(owner any) *app__http__HttpRequest {
	self := &app__http__HttpRequest{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_app__http__HttpRequest)
	}
	return self
}

func hx_init_app__http__HttpRequest(self *app__http__HttpRequest, method *string, path *string, body *string) *app__http__HttpRequest {
	self.method = method
	self.path = path
	self.body = body
	return self
}
