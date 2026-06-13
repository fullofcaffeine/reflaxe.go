package main

type I_app__http__HttpResponse interface {
}

type app__http__HttpResponse struct {
	__hx_this I_app__http__HttpResponse
	status    int
	body      *string
}

func New_app__http__HttpResponse(status int, body *string) *app__http__HttpResponse {
	self := &app__http__HttpResponse{}
	self.__hx_this = self
	self.status = status
	self.body = body
	return self
}

func app__http__HttpResponse_json(status int, body *string) *app__http__HttpResponse {
	return New_app__http__HttpResponse(status, body)
}
