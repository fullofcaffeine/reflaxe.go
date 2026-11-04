package main

import "snapshot/hxrt"

func main() {
	var named Named = New_Person()
	printNamed(named)
	printNamed(New_Person())
}

func printNamed(value Named) {
	var v any = any(value.name())
	hxrt.Println(v)
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

func (self *Person) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_Person() *Person {
	return hx_init_Person(hx_alloc_Person(nil))
}

func hx_alloc_Person(owner any) *Person {
	self := &Person{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_Person)
	}
	return self
}

func hx_init_Person(self *Person) *Person {
	return self
}

func (self *Person) name() *string {
	return hxrt.StringFromLiteral("person")
}
