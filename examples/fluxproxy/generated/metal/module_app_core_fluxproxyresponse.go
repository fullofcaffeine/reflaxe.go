package main

type I_app__core__FluxProxyResponse interface {
}

type app__core__FluxProxyResponse struct {
	__hx_this I_app__core__FluxProxyResponse
	requestId int
	route     *string
	upstream  *string
	status    int
	latencyMs int
	attempts  int
	success   bool
}

func New_app__core__FluxProxyResponse(requestId int, route *string, upstream *string, status int, latencyMs int, attempts int, success bool) *app__core__FluxProxyResponse {
	self := &app__core__FluxProxyResponse{}
	self.__hx_this = self
	self.requestId = requestId
	self.route = route
	self.upstream = upstream
	self.status = status
	self.latencyMs = latencyMs
	self.attempts = attempts
	self.success = success
	return self
}
