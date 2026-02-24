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

func New_app__core__FluxRequest(id int, route *string, latencyMs int, status int) *app__core__FluxRequest {
	self := &app__core__FluxRequest{}
	self.__hx_this = self
	self.id = id
	self.route = route
	self.latencyMs = latencyMs
	self.status = status
	return self
}
