package main

import "snapshot/hxrt"

type I_Box interface {
}

type Box struct {
	__hx_this I_Box
	value     any
}

func (self *Box) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_Box(value any) *Box {
	return hx_init_Box(hx_alloc_Box(nil), value)
}

func hx_alloc_Box(owner any) *Box {
	self := &Box{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_Box)
	}
	return self
}

func hx_init_Box(self *Box, value any) *Box {
	self.value = value
	return self
}

type I_Leaf interface {
}

type Leaf struct {
	__hx_this I_Leaf
	id        int
}

func (self *Leaf) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_Leaf(id int) *Leaf {
	return hx_init_Leaf(hx_alloc_Leaf(nil), id)
}

func hx_alloc_Leaf(owner any) *Leaf {
	self := &Leaf{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_Leaf)
	}
	return self
}

func hx_init_Leaf(self *Leaf, id int) *Leaf {
	self.id = id
	return self
}

func main() {
	reference := New_Box(New_Leaf(7))
	var v any = any(hx_recover_class_0(reference.value).id)
	hxrt.Println(v)
	number := New_Box(4)
	var v_1 any = any(int((hxrt.Int32Wrap(func(hx_value_1 any) int {
		if hx_value_1 == nil {
			var hx_zero_2 int
			return hx_zero_2
		}
		return hx_value_1.(int)
	}(number.value)) + hxrt.Int32Wrap(1))))
	hxrt.Println(v_1)
	text := New_Box(hxrt.StringFromLiteral("word"))
	var v_2 any = any(hxrt.StringToUpperCaseStringPtr(func(hx_value_3 any) *string {
		if hx_value_3 == nil {
			var hx_zero_4 *string
			return hx_zero_4
		}
		return hx_value_3.(*string)
	}(text.value)))
	hxrt.Println(v_2)
	missing := New_Box(nil)
	loaded := hx_recover_class_0(missing.value)
	hxrt.Println(any((loaded == nil)))
	callable := New_Box(hxrt.NewFunction(func(value int) int {
		return int((hxrt.Int32Wrap(value) + hxrt.Int32Wrap(1)))
	}))
	var v_3 any = any(func() int {
		var hx_callable_5 *hxrt.Function[func(int) int] = hx_recover_callable_0(callable.value)
		var hx_argument_6 int = 8
		return hxrt.FunctionCall(hx_callable_5)(hx_argument_6)
	}())
	hxrt.Println(v_3)
}

func hx_recover_callable_0(value any) *hxrt.Function[func(int) int] {
	if value == nil {
		return nil
	}
	switch exact := value.(type) {
	case *hxrt.Function[func(int) int]:
		return exact
	}
	origin := hxrt.FunctionOrigin(value)
	if origin == nil {
		return nil
	}
	switch source := origin.(type) {
	case *hxrt.Function[func(int) int]:
		return source
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}

func hx_recover_class_0(value any) *Leaf {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *Leaf:
		return source
	default:
		return value.(*Leaf)
	}
}
