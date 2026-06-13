package main

type I_app__http__HttpRequest interface {
}

type app__http__HttpRequest struct {
	__hx_this I_app__http__HttpRequest
	method    *string
	path      *string
	body      *string
}

func New_app__http__HttpRequest(method *string, path *string, body *string) *app__http__HttpRequest {
	self := &app__http__HttpRequest{}
	self.__hx_this = self
	self.method = method
	self.path = path
	self.body = body
	return self
}
