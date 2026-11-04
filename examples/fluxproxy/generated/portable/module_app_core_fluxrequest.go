package main

type I_app__core__FluxRequest interface {
}

type app__core__FluxRequest struct {
	__hx_this I_app__core__FluxRequest
	id        int
	route     *string
	latencyMs int
	status    int
}

func (self *app__core__FluxRequest) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_app__core__FluxRequest(id int, route *string, latencyMs int, status int) *app__core__FluxRequest {
	return hx_init_app__core__FluxRequest(hx_alloc_app__core__FluxRequest(nil), id, route, latencyMs, status)
}

func hx_alloc_app__core__FluxRequest(owner any) *app__core__FluxRequest {
	self := &app__core__FluxRequest{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_app__core__FluxRequest)
	}
	return self
}

func hx_init_app__core__FluxRequest(self *app__core__FluxRequest, id int, route *string, latencyMs int, status int) *app__core__FluxRequest {
	self.id = id
	self.route = route
	self.latencyMs = latencyMs
	self.status = status
	return self
}
