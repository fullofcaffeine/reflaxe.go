package main

import "snapshot/hxrt"

func main() {
	var named Named = New_Person()
	_ = named
	printNamed(named)
	printNamed(New_Person())
}

func printNamed(value Named) {
	hxrt.Println(value.name())
}

type Named interface {
	name() *string
}

type I_Person interface {
	name() *string
}

type Person struct {
	__hx_this I_Person
}

func New_Person() *Person {
	self := &Person{}
	self.__hx_this = self
	return self
}

func (self *Person) name() *string {
	return hxrt.StringFromLiteral("person")
}
