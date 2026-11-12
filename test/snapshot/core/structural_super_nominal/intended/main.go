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
	return hx_init_Base(hx_alloc_Base(nil), it)
}

func hx_alloc_Base(owner any) *Base {
	self := &Base{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_Base)
	}
	return self
}

func hx_init_Base(self *Base, it hxrt.Anonymous) *Base {
	self.it = it
	return self
}

func (self *Base) read() int {
	return func() int {
		var hx_callable_1 *hxrt.Function[func() int] = hx_recover_callable_0(hxrt.AnonymousGet(self.it, "next"))
		return hxrt.FunctionCall(hx_callable_1)()
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
	return hx_init_Child(hx_alloc_Child(nil))
}

func hx_alloc_Child(owner any) *Child {
	self := &Child{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_Child)
	}
	self.Base = hx_alloc_Base(self.__hx_this)
	return self
}

func hx_init_Child(self *Child) *Child {
	hx_init_Base(self.Base, func(hx_structural_source_2 *Cursor) hxrt.Anonymous {
		return hxrt.NewAnonymousView(hx_structural_source_2, func(name string) hxrt.AnonymousField {
			switch name {
			case "hasNext":
				return hxrt.AnonymousField{Value: func() *hxrt.Function[func() bool] {
					hx_receiver_3 := hx_structural_source_2
					return hxrt.BoundFunction(func() bool {
						return hx_receiver_3.__hx_this.hasNext()
					}, func() any {
						return hx_receiver_3.__hx_this
					}, "hasNext")
				}(), Found: true}
			case "next":
				return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
					hx_receiver_4 := hx_structural_source_2
					return hxrt.BoundFunction(func() int {
						return hx_receiver_4.__hx_this.next()
					}, func() any {
						return hx_receiver_4.__hx_this
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
	return hx_init_Cursor(hx_alloc_Cursor(nil))
}

func hx_alloc_Cursor(owner any) *Cursor {
	self := &Cursor{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_Cursor)
	}
	return self
}

func hx_init_Cursor(self *Cursor) *Cursor {
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

func hx_recover_callable_0(value any) *hxrt.Function[func() int] {
	if value == nil {
		return nil
	}
	switch exact := value.(type) {
	case *hxrt.Function[func() int]:
		return exact
	}
	origin := hxrt.FunctionOrigin(value)
	if origin == nil {
		return nil
	}
	switch source := origin.(type) {
	case *hxrt.Function[func() bool]:
		return hxrt.ProjectFunction(source, func() int {
			return func(hx_value_1 any) int {
				if hx_value_1 == nil {
					var hx_zero_2 int
					return hx_zero_2
				}
				return hx_value_1.(int)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() int]:
		return source
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}
