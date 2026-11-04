package main

import "snapshot/hxrt"

type Choice struct {
	tag    int
	params []any
}

var Choice_One *Choice = &Choice{tag: 0}

type I_Base interface {
}

type Base struct {
	__hx_this I_Base
	value     int
}

func (self *Base) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_Base() *Base {
	return hx_init_Base(hx_alloc_Base(nil))
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

func hx_init_Base(self *Base) *Base {
	self.value = 3
	return self
}

type I_Child interface {
	name() *string
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
	hx_init_Base(self.Base)
	return self
}

func (self *Child) name() *string {
	return hxrt.StringFromLiteral("child")
}

func main() {
	var erased any = New_Child()
	base := hx_recover_class_0(erased)
	var v any = any(base.value)
	hxrt.Println(v)
	var named Named = func(hx_value_1 any) Named {
		if hx_value_1 == nil {
			var hx_zero_2 Named
			return hx_zero_2
		}
		return hx_value_1.(Named)
	}(hxrt.ObjectSource(erased))
	var v_1 any = any(named.name())
	hxrt.Println(v_1)
	base.value = 9
	child := hx_recover_class_1(erased)
	var v_2 any = any(child.value)
	hxrt.Println(v_2)
	var absent any = nil
	missing := hx_recover_class_0(absent)
	hxrt.Println(any((missing == nil)))
	var missingNamed Named = func(hx_value_3 any) Named {
		if hx_value_3 == nil {
			var hx_zero_4 Named
			return hx_zero_4
		}
		return hx_value_3.(Named)
	}(hxrt.ObjectSource(absent))
	hxrt.Println(any((missingNamed == nil)))
	var choiceValue any = Choice_One
	choice := func(hx_value_5 any) *Choice {
		if hx_value_5 == nil {
			var hx_zero_6 *Choice
			return hx_zero_6
		}
		return hx_value_5.(*Choice)
	}(any(choiceValue))
	hxrt.Println(any((choice == Choice_One)))
	var arrayValue any = hxrt.NewArray(4, 5)
	values := func(hx_value_7 any) *hxrt.Array {
		if hx_value_7 == nil {
			var hx_zero_8 *hxrt.Array
			return hx_zero_8
		}
		return hx_value_7.(*hxrt.Array)
	}(any(arrayValue))
	hxrt.Println(values.Get(0))
}

type Named interface {
	name() *string
}

func hx_recover_class_0(value any) *Base {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *Base:
		return source
	case *Child:
		return source.Base
	default:
		return value.(*Base)
	}
}

func hx_recover_class_1(value any) *Child {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *Child:
		return source
	default:
		return value.(*Child)
	}
}
