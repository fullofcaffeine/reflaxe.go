package main

import "snapshot/hxrt"

type I_DomainError interface {
}

type DomainError struct {
	__hx_this      I_DomainError
	code           *string
	__hx_exception *hxrt.ExceptionValue
}

func New_DomainError(code *string, message *string) *DomainError {
	self := &DomainError{}
	self.__hx_exception = hxrt.BindException(self, message, nil, nil)
	self.__hx_this = self
	self.code = code
	return self
}

func (self *DomainError) HxExceptionValue() *hxrt.ExceptionValue {
	return self.__hx_exception
}

func main() {
	hxrt.TryCatch(func() {
		raiseDomain()
	}, func(hx_caught_1 any) {
		switch hx_typed_2 := hx_caught_1.(type) {
		case *DomainError:
			error := hx_typed_2
			hxrt.Println(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("typed:"), error.code))
		default:
			error_1 := hxrt.ExceptionCaught(hx_caught_1)
			hxrt.Println(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("base:"), hxrt.ExceptionMessage(error_1)))
		}
	})
	hxrt.TryCatch(func() {
		raisePlain()
	}, func(hx_caught_3 any) {
		switch hx_typed_4 := hx_caught_3.(type) {
		case *DomainError:
			error_2 := hx_typed_4
			hxrt.Println(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("typed:"), error_2.code))
		default:
			error_3 := hxrt.ExceptionCaught(hx_caught_3)
			hxrt.Println(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("base:"), hxrt.ExceptionMessage(error_3)))
		}
	})
}

func raiseDomain() {
	hxrt.Throw(New_DomainError(hxrt.StringFromLiteral("E42"), hxrt.StringFromLiteral("typed domain failure")))
}

func raisePlain() {
	hxrt.Throw(hxrt.StringFromLiteral("plain"))
}
