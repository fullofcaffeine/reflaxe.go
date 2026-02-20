package main

import "snapshot/hxrt"

func main() {
	ok := go___Result_ok(7)
	_ = ok
	hxrt.Println(ok.__hx_this.isOk())
	hxrt.Println(ok.__hx_this.unwrap())
	err := go___Result_failure(hxrt.StringFromLiteral("boom"))
	_ = err
	hxrt.Println(err.__hx_this.isErr())
	hxrt.Println(err.__hx_this.error())
}

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

type I_go___Result interface {
	isOk() bool
	isErr() bool
	unwrap() any
	error() *string
}

type go___Result struct {
	__hx_this  I_go___Result
	value      any
	errorValue *go___Error
}

func New_go___Result(value any, errorValue *go___Error) *go___Result {
	self := &go___Result{}
	self.__hx_this = self
	self.value = value
	self.errorValue = errorValue
	return self
}

func (self *go___Result) isOk() bool {
	return hxrt.StringEqualAny(self.errorValue, nil)
}

func (self *go___Result) isErr() bool {
	return !hxrt.StringEqualAny(self.errorValue, nil)
}

func (self *go___Result) unwrap() any {
	if !hxrt.StringEqualAny(self.errorValue, nil) {
		hxrt.Throw(self.errorValue.__hx_this.toString())
		var hx_throw_zero_1 any
		return hx_throw_zero_1
	}
	return self.value
}

func (self *go___Result) error() *string {
	if hxrt.StringEqualAny(self.errorValue, nil) {
		return nil
	}
	return self.errorValue.__hx_this.toString()
}

func go___Result_failure(message *string) *go___Result {
	return New_go___Result(nil, New_go___Error(message))
}

func go___Result_ok(value any) *go___Result {
	return New_go___Result(value, nil)
}
