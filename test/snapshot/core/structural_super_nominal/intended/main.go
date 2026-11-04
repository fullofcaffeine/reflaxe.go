package main

import "snapshot/hxrt"

type I_Base interface {
	read() int
}

type Base struct {
	__hx_this I_Base
	it        hxrt.Anonymous
}

func (self *Base) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_Base(it hxrt.Anonymous) *Base {
	self := &Base{}
	self.__hx_this = self
	self.it = it
	return self
}

func (self *Base) read() int {
	return func() int {
		var hx_callable_4 *hxrt.Function[func() int] = func(hx_obj_1 hxrt.Anonymous) *hxrt.Function[func() int] {
			hx_field_2 := hxrt.AnonymousGet(hx_obj_1, "next")
			if hx_field_2 == nil {
				var hx_zero_3 *hxrt.Function[func() int]
				return hx_zero_3
			}
			return hx_field_2.(*hxrt.Function[func() int])
		}(self.it)
		return hxrt.FunctionCall(hx_callable_4)()
	}()
}

type I_Child interface {
	read() int
}

type Child struct {
	*Base
	__hx_this I_Child
}

func (self *Child) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_Child() *Child {
	self := &Child{}
	self.Base = New_Base(func(hx_structural_source_5 *Cursor) hxrt.Anonymous {
		return hxrt.NewAnonymousView(hx_structural_source_5, func(name string) hxrt.AnonymousField {
			switch name {
			case "hasNext":
				return hxrt.AnonymousField{Value: func() *hxrt.Function[func() bool] {
					hx_receiver_6 := hx_structural_source_5
					return hxrt.BoundFunction(func() bool {
						return hx_receiver_6.__hx_this.hasNext()
					}, func() any {
						return hx_receiver_6.__hx_this
					}, "hasNext")
				}(), Found: true}
			case "next":
				return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
					hx_receiver_7 := hx_structural_source_5
					return hxrt.BoundFunction(func() int {
						return hx_receiver_7.__hx_this.next()
					}, func() any {
						return hx_receiver_7.__hx_this
					}, "next")
				}(), Found: true}
			}
			return hxrt.AnonymousField{}
		}, func(name string, incoming any) bool {
			switch name {
			}
			return false
		})
	}(New_Cursor()))
	self.Base.__hx_this = self
	self.__hx_this = self
	return self
}

type I_Cursor interface {
	hasNext() bool
	next() int
}

type Cursor struct {
	__hx_this I_Cursor
}

func (self *Cursor) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_Cursor() *Cursor {
	self := &Cursor{}
	self.__hx_this = self
	return self
}

func (self *Cursor) hasNext() bool {
	return true
}

func (self *Cursor) next() int {
	return 3
}

func main() {
	var v any = any(New_Child().read())
	hxrt.Println(v)
}
