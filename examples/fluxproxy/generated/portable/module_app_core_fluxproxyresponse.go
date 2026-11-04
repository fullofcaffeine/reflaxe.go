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

func (self *app__core__FluxProxyResponse) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_app__core__FluxProxyResponse(requestId int, route *string, upstream *string, status int, latencyMs int, attempts int, success bool) *app__core__FluxProxyResponse {
	return hx_init_app__core__FluxProxyResponse(hx_alloc_app__core__FluxProxyResponse(nil), requestId, route, upstream, status, latencyMs, attempts, success)
}

func hx_alloc_app__core__FluxProxyResponse(owner any) *app__core__FluxProxyResponse {
	self := &app__core__FluxProxyResponse{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_app__core__FluxProxyResponse)
	}
	return self
}

func hx_init_app__core__FluxProxyResponse(self *app__core__FluxProxyResponse, requestId int, route *string, upstream *string, status int, latencyMs int, attempts int, success bool) *app__core__FluxProxyResponse {
	self.requestId = requestId
	self.route = route
	self.upstream = upstream
	self.status = status
	self.latencyMs = latencyMs
	self.attempts = attempts
	self.success = success
	return self
}
