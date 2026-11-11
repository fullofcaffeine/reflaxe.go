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
}

func (self *Leaf) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_Leaf() *Leaf {
	return hx_init_Leaf(hx_alloc_Leaf(nil))
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

func hx_init_Leaf(self *Leaf) *Leaf {
	return self
}

func main() {
	var absent *Leaf = nil
	empty := New_Box(absent)
	present := New_Box(New_Leaf())
	var v any = any(hxrt.AnyEqualsNull(hx_recover_class_0(empty.value)))
	hxrt.Println(v)
	var v_1 any = any(hxrt.AnyEqualsNull(hx_recover_class_0(empty.value)))
	hxrt.Println(v_1)
	var v_2 any = any(!hxrt.AnyEqualsNull(hx_recover_class_0(empty.value)))
	hxrt.Println(v_2)
	var v_3 any = any(!hxrt.AnyEqualsNull(hx_recover_class_0(empty.value)))
	hxrt.Println(v_3)
	var v_4 any = any(hxrt.AnyEqualsNull(hx_recover_class_0(present.value)))
	hxrt.Println(v_4)
	var v_5 any = any(!hxrt.AnyEqualsNull(hx_recover_class_0(present.value)))
	hxrt.Println(v_5)
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
