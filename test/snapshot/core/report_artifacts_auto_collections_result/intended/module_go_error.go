package main

type I_go___Error interface {
	toString() *string
}

type go___Error struct {
	__hx_this I_go___Error
	message   *string
}

func (self *go___Error) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_go___Error(message *string) *go___Error {
	return hx_init_go___Error(hx_alloc_go___Error(nil), message)
}

func hx_alloc_go___Error(owner any) *go___Error {
	self := &go___Error{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_go___Error)
	}
	return self
}

func hx_init_go___Error(self *go___Error, message *string) *go___Error {
	self.message = message
	return self
}

func (self *go___Error) toString() *string {
	return self.message
}

func (self *go___Error) String() string {
	return *self.toString()
}
