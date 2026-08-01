package main

import "snapshot/hxrt"

type DefaultedValue interface {
	get(value any) int
	label(value *string) *string
}

type I_DefaultedValueBase interface {
	get(value any) int
}

type DefaultedValueBase struct {
	__hx_this I_DefaultedValueBase
}

func New_DefaultedValueBase() *DefaultedValueBase {
	self := &DefaultedValueBase{}
	self.__hx_this = self
	return self
}

func (self *DefaultedValueBase) get(value any) int {
	if value == nil {
		value = 7
	}
	return value.(int)
}

type I_DefaultedValueImpl interface {
	get(value any) int
	label(value *string) *string
}

type DefaultedValueImpl struct {
	*DefaultedValueBase
	__hx_this I_DefaultedValueImpl
}

func New_DefaultedValueImpl() *DefaultedValueImpl {
	self := &DefaultedValueImpl{}
	self.DefaultedValueBase = New_DefaultedValueBase()
	self.DefaultedValueBase.__hx_this = self
	self.__hx_this = self
	return self
}

func (self *DefaultedValueImpl) label(value *string) *string {
	if value == nil {
		value = hxrt.StringFromLiteral("implementation")
	}
	return value
}

func main() {
	var value DefaultedValue = New_DefaultedValueImpl()
	var v any = any(value.get(nil))
	hxrt.Println(v)
	var v_1 any = any(value.get(9))
	hxrt.Println(v_1)
	var v_2 any = any(value.label(nil))
	hxrt.Println(v_2)
	var v_3 any = any(value.label(hxrt.StringFromLiteral("explicit")))
	hxrt.Println(v_3)
}
