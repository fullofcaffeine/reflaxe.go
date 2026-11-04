package main

type I_app__http__HttpResponse interface {
}

type app__http__HttpResponse struct {
	__hx_this I_app__http__HttpResponse
	status    int
	body      *string
}

func (self *app__http__HttpResponse) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_app__http__HttpResponse(status int, body *string) *app__http__HttpResponse {
	return hx_init_app__http__HttpResponse(hx_alloc_app__http__HttpResponse(nil), status, body)
}

func hx_alloc_app__http__HttpResponse(owner any) *app__http__HttpResponse {
	self := &app__http__HttpResponse{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_app__http__HttpResponse)
	}
	return self
}

func hx_init_app__http__HttpResponse(self *app__http__HttpResponse, status int, body *string) *app__http__HttpResponse {
	self.status = status
	self.body = body
	return self
}

func app__http__HttpResponse_json(status int, body *string) *app__http__HttpResponse {
	return New_app__http__HttpResponse(status, body)
}
