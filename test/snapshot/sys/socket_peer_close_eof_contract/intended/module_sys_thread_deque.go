package main

type I_sys__thread__Deque interface {
	add(i any)
	push(i any)
	pop(block bool) any
}

type sys__thread__Deque struct {
	__hx_this   I_sys__thread__Deque
	__mutex     *sys__thread__Mutex
	__available *sys__thread__Lock
	__items     *haxe__ds__List
}

func New_sys__thread__Deque() *sys__thread__Deque {
	self := &sys__thread__Deque{}
	self.__hx_this = self
	self.__items = New_haxe__ds__List()
	self.__available = New_sys__thread__Lock()
	self.__mutex = New_sys__thread__Mutex()
	return self
}

func (self *sys__thread__Deque) add(i any) {
	self.__mutex.__hx_this.acquire()
	self.__items.__hx_this.add(i)
	self.__mutex.__hx_this.release()
	self.__available.__hx_this.release()
}

func (self *sys__thread__Deque) push(i any) {
	self.__mutex.__hx_this.acquire()
	self.__items.__hx_this.push(i)
	self.__mutex.__hx_this.release()
	self.__available.__hx_this.release()
}

func (self *sys__thread__Deque) pop(block bool) any {
	if block {
		self.__available.__hx_this.wait(nil)
	} else {
		if !self.__available.__hx_this.wait(0.0) {
			return nil
		}
	}
	self.__mutex.__hx_this.acquire()
	var value any = self.__items.__hx_this.pop()
	self.__mutex.__hx_this.release()
	return value
}
