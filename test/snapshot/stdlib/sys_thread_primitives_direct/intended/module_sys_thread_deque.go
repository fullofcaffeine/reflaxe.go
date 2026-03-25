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
	self.__mutex.acquire()
	self.__items.add(i)
	self.__mutex.release()
	self.__available.release()
}

func (self *sys__thread__Deque) push(i any) {
	self.__mutex.acquire()
	self.__items.push(i)
	self.__mutex.release()
	self.__available.release()
}

func (self *sys__thread__Deque) pop(block bool) any {
	if block {
		self.__available.wait(nil)
	} else {
		if !self.__available.wait(0.0) {
			return nil
		}
	}
	self.__mutex.acquire()
	var value any = self.__items.pop()
	self.__mutex.release()
	return value
}
