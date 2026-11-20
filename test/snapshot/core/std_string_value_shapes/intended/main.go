package main

import "snapshot/hxrt"

type Color struct {
	tag    int
	params []any
}

func (self *Color) HxEnumValue() (string, []any) {
	return [2]string{"Red", "RGB"}[self.tag], self.params
}

var Color_Red *Color = &Color{tag: 0}

func Color_RGB(r int, g int, b int) *Color {
	enumValue := &Color{tag: 1}
	enumValue.params = []any{r, g, b}
	return enumValue
}

type I_Described interface {
	toString() *string
}

type Described struct {
	__hx_this I_Described
}

func (self *Described) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_Described() *Described {
	return hx_init_Described(hx_alloc_Described(nil))
}

func hx_alloc_Described(owner any) *Described {
	self := &Described{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_Described)
	}
	return self
}

func hx_init_Described(self *Described) *Described {
	return self
}

func (self *Described) toString() *string {
	return hxrt.StringFromLiteral("described")
}

func (self *Described) String() string {
	return *self.toString()
}

func main() {
	var rgb any = Color_RGB(1, 2, 3)
	var v any = any(hxrt.StdString(Color_Red))
	hxrt.Println(v)
	var v_1 any = any(hxrt.StdString(rgb))
	hxrt.Println(v_1)
	var v_2 any = any(hxrt.StdString(New_Plain()))
	hxrt.Println(v_2)
	var v_3 any = any(hxrt.StdString(New_Described()))
	hxrt.Println(v_3)
	var v_4 any = any(hxrt.StdString(func() hxrt.Anonymous {
		hx_obj_1 := map[string]any{}
		hx_obj_1["color"] = Color_Red
		hx_obj_1["plain"] = New_Plain()
		return hxrt.AnonymousFromMap(hx_obj_1)
	}()))
	hxrt.Println(v_4)
	var v_5 any = any(hxrt.StdString(hxrt.NewArray(Color_Red, Color_RGB(4, 5, 6))))
	hxrt.Println(v_5)
}

type I_Plain interface {
}

type Plain struct {
	__hx_this I_Plain
}

func (self *Plain) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_Plain() *Plain {
	return hx_init_Plain(hx_alloc_Plain(nil))
}

func hx_alloc_Plain(owner any) *Plain {
	self := &Plain{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_Plain)
	}
	return self
}

func hx_init_Plain(self *Plain) *Plain {
	return self
}

func (self *Plain) HxClassName() string {
	return "Plain"
}
