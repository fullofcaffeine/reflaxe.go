package main

type I_go___Chan interface {
	__hx_setBuffer(buffer int)
	send(value any)
	recv() any
	trySend(value any) bool
	tryRecv() *go___Result
	recvOr(defaultValue any) any
	close()
}

type go___Chan struct {
	__hx_this   I_go___Chan
	__hx_native any
}

func (self *go___Chan) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_go___Chan() *go___Chan {
	return hx_init_go___Chan(hx_alloc_go___Chan(nil))
}

func hx_alloc_go___Chan(owner any) *go___Chan {
	self := &go___Chan{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_go___Chan)
	}
	return self
}

func hx_init_go___Chan(self *go___Chan) *go___Chan {
	self.__hx_native = go__concurrency_makeChan(0)
	return self
}

func (self *go___Chan) __hx_setBuffer(buffer int) {
	self.__hx_native = go__concurrency_makeChan(buffer)
}

func (self *go___Chan) send(value any) {
	go__concurrency_send(self.__hx_native, value)
}

func (self *go___Chan) recv() any {
	return go__concurrency_recv(self.__hx_native)
}

func (self *go___Chan) trySend(value any) bool {
	return go__concurrency_trySend(self.__hx_native, value)
}

func (self *go___Chan) tryRecv() *go___Result {
	return go__concurrency_tryRecv(self.__hx_native)
}

func (self *go___Chan) recvOr(defaultValue any) any {
	return go__concurrency_recvOr(self.__hx_native, defaultValue)
}

func (self *go___Chan) close() {
	go__concurrency_close(self.__hx_native)
}
