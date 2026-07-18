package main

type I_go___Error interface {
	toString() *string
}

type go___Error struct {
	__hx_this I_go___Error
	message   *string
}

func New_go___Error(message *string) *go___Error {
	self := &go___Error{}
	self.__hx_this = self
	self.message = message
	return self
}

func (self *go___Error) toString() *string {
	return self.message
}

func (self *go___Error) String() string {
	return *self.toString()
}
