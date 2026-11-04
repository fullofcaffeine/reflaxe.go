package main

import "snapshot/hxrt"

type I_Child interface {
	get_message() *string
	get_stack() any
	get_previous() *haxe__Exception
	get_native() any
	unwrap() any
	toString() *string
	details() *string
}

type Child struct {
	*Parent
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
	self.Parent = hx_alloc_Parent(self.__hx_this)
	return self
}

func hx_init_Child(self *Child) *Child {
	hx_init_Parent(self.Parent)
	return self
}

func (self *Child) String() string {
	return *self.toString()
}

func absent() *Child {
	calls = int(int32((calls + 1)))
	return nil
}

var calls int = 0

func main() {
	child := absent()
	parent := func(hx_value *Child) *Parent {
		if hx_value == nil {
			return nil
		}
		return hx_value.Parent
	}(child)
	base := func(hx_value *Child) *haxe__Exception {
		if hx_value == nil {
			return nil
		}
		return hx_value.Parent.haxe__Exception
	}(absent())
	var v any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("parent-null="), hxrt.StdString((parent == nil))))
	hxrt.Println(v)
	var v_1 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("base-null="), hxrt.StdString((base == nil))))
	hxrt.Println(v_1)
	var v_2 any = any(hxrt.StringConcatAny(hxrt.StringFromLiteral("calls="), calls))
	hxrt.Println(v_2)
}

type I_Parent interface {
	get_message() *string
	get_stack() any
	get_previous() *haxe__Exception
	get_native() any
	unwrap() any
	toString() *string
	details() *string
}

type Parent struct {
	*haxe__Exception
	__hx_this I_Parent
}

func (self *Parent) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_Parent() *Parent {
	return hx_init_Parent(hx_alloc_Parent(nil))
}

func hx_alloc_Parent(owner any) *Parent {
	self := &Parent{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_Parent)
	}
	self.haxe__Exception = hx_alloc_haxe__Exception(self.__hx_this)
	return self
}

func hx_init_Parent(self *Parent) *Parent {
	hx_init_haxe__Exception(self.haxe__Exception, hxrt.StringFromLiteral("parent"), nil, nil)
	return self
}

func (self *Parent) String() string {
	return *self.__hx_this.toString()
}

func hx_recover_class_0(value any) *haxe__ValueException {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *haxe__ValueException:
		return source
	default:
		return value.(*haxe__ValueException)
	}
}

func hx_recover_class_1(value any) *haxe__Exception {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *Child:
		return source.Parent.haxe__Exception
	case *Parent:
		return source.haxe__Exception
	case *haxe__Exception:
		return source
	case *haxe__ValueException:
		return source.haxe__Exception
	default:
		return value.(*haxe__Exception)
	}
}
