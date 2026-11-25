package main

import "snapshot/hxrt"

func ACaster_toLeaf(value *Base) *Leaf {
	return hx_recover_class_Leaf(value)
}

func ACaster_toMid(value *Base) *Mid {
	return hx_recover_class_Mid(value)
}

type I_Base interface {
	describe() *string
}

type Base struct {
	__hx_this I_Base
	tag       *string
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
	self.tag = hxrt.StringFromLiteral("base")
	return self
}

func (self *Base) describe() *string {
	return hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Base("), self.tag), hxrt.StringFromLiteral(")"))
}

func (self *Base) HxClassName() string {
	return "Base"
}

type I_Leaf interface {
	describe() *string
	midName() *string
}

type Leaf struct {
	*Mid
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
	self.Mid = hx_alloc_Mid(self.__hx_this)
	return self
}

func hx_init_Leaf(self *Leaf) *Leaf {
	hx_init_Mid(self.Mid)
	self.tag = hxrt.StringFromLiteral("leaf")
	return self
}

func (self *Leaf) describe() *string {
	return hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Leaf("), self.tag), hxrt.StringFromLiteral(")"))
}

func (self *Leaf) HxClassName() string {
	return "Leaf"
}

func main() {
	leafAsBase := func(hx_value *Leaf) *Base {
		if hx_value == nil {
			return nil
		}
		return hx_value.Mid.Base
	}(New_Leaf())
	mid := hx_recover_class_Mid(leafAsBase)
	var v any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("leaf as Mid tag="), mid.tag))
	hxrt.Println(v)
	var v_1 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("leaf as Mid describe="), mid.__hx_this.describe()))
	hxrt.Println(v_1)
	var v_2 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("leaf as Mid midName="), mid.__hx_this.midName()))
	hxrt.Println(v_2)
	var v_3 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("leaf as Mid isOfType Leaf="), hxrt.StdString(hx_is_nominal_Leaf(mid))))
	hxrt.Println(v_3)
	midAsBase := func(hx_value *Mid) *Base {
		if hx_value == nil {
			return nil
		}
		return hx_value.Base
	}(mid)
	var v_4 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("leaf as Mid keeps identity="), hxrt.StdString((midAsBase == leafAsBase))))
	hxrt.Println(v_4)
	mid.tag = hxrt.StringFromLiteral("changed")
	var v_5 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("write through Mid seen by Base="), leafAsBase.tag))
	hxrt.Println(v_5)
	tipAsBase := func(hx_value *Tip) *Base {
		if hx_value == nil {
			return nil
		}
		return hx_value.Leaf.Mid.Base
	}(New_Tip())
	tipAsMid := hx_recover_class_Mid(tipAsBase)
	tipAsLeaf := hx_recover_class_Leaf(tipAsMid)
	var v_6 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("tip as Mid="), tipAsMid.__hx_this.describe()))
	hxrt.Println(v_6)
	var v_7 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("tip as Leaf from Mid="), tipAsLeaf.__hx_this.describe()))
	hxrt.Println(v_7)
	exactMid := func(hx_value *Mid) *Base {
		if hx_value == nil {
			return nil
		}
		return hx_value.Base
	}(New_Mid())
	exact := hx_recover_class_Mid(exactMid)
	var v_8 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("exact Mid="), exact.__hx_this.describe()))
	hxrt.Println(v_8)
	var v_9 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("caster toMid="), ACaster_toMid(func(hx_value *Tip) *Base {
		if hx_value == nil {
			return nil
		}
		return hx_value.Leaf.Mid.Base
	}(New_Tip())).__hx_this.describe()))
	hxrt.Println(v_9)
	var v_10 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("caster toLeaf="), ACaster_toLeaf(func(hx_value *Tip) *Base {
		if hx_value == nil {
			return nil
		}
		return hx_value.Leaf.Mid.Base
	}(New_Tip())).__hx_this.describe()))
	hxrt.Println(v_10)
	var none *Base = nil
	noneMid := hx_recover_class_Mid(none)
	var v_11 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("null stays null="), hxrt.StdString((noneMid == nil))))
	hxrt.Println(v_11)
	error_local := func(hx_value *haxe__exceptions__NotImplementedException) *haxe__Exception {
		if hx_value == nil {
			return nil
		}
		return hx_value.haxe__exceptions__PosException.haxe__Exception
	}(New_haxe__exceptions__NotImplementedException(hxrt.StringFromLiteral("late"), nil, func() hxrt.Anonymous {
		hx_obj_1 := map[string]any{}
		hx_obj_1["fileName"] = hxrt.StringFromLiteral("Main.hx")
		hx_obj_1["lineNumber"] = 99
		hx_obj_1["className"] = hxrt.StringFromLiteral("Main")
		hx_obj_1["methodName"] = hxrt.StringFromLiteral("main")
		return hxrt.AnonymousFromMap(hx_obj_1)
	}()))
	pos := hx_recover_class_haxe__exceptions__PosException(error_local)
	var v_12 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("std exception as PosException="), pos.__hx_this.get_message()))
	hxrt.Println(v_12)
}

type I_Mid interface {
	describe() *string
	midName() *string
}

type Mid struct {
	*Base
	__hx_this I_Mid
}

func (self *Mid) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_Mid() *Mid {
	return hx_init_Mid(hx_alloc_Mid(nil))
}

func hx_alloc_Mid(owner any) *Mid {
	self := &Mid{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_Mid)
	}
	self.Base = hx_alloc_Base(self.__hx_this)
	return self
}

func hx_init_Mid(self *Mid) *Mid {
	hx_init_Base(self.Base)
	self.tag = hxrt.StringFromLiteral("mid")
	return self
}

func (self *Mid) describe() *string {
	return hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Mid("), self.tag), hxrt.StringFromLiteral(")"))
}

func (self *Mid) midName() *string {
	return hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("midName:"), self.tag)
}

func (self *Mid) HxClassName() string {
	return "Mid"
}

type I_Tip interface {
	describe() *string
	midName() *string
}

type Tip struct {
	*Leaf
	__hx_this I_Tip
}

func (self *Tip) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_Tip() *Tip {
	return hx_init_Tip(hx_alloc_Tip(nil))
}

func hx_alloc_Tip(owner any) *Tip {
	self := &Tip{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_Tip)
	}
	self.Leaf = hx_alloc_Leaf(self.__hx_this)
	return self
}

func hx_init_Tip(self *Tip) *Tip {
	hx_init_Leaf(self.Leaf)
	self.tag = hxrt.StringFromLiteral("tip")
	return self
}

func (self *Tip) HxClassName() string {
	return "Tip"
}

func hx_recover_class_Leaf(value any) *Leaf {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *Leaf:
		return source
	case *Tip:
		return source.Leaf
	default:
		return value.(*Leaf)
	}
}

func hx_recover_class_Mid(value any) *Mid {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *Leaf:
		return source.Mid
	case *Mid:
		return source
	case *Tip:
		return source.Leaf.Mid
	default:
		return value.(*Mid)
	}
}

func hx_recover_class_haxe__Exception(value any) *haxe__Exception {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *haxe__exceptions__PosException:
		return source.haxe__Exception
	case *haxe__exceptions__NotImplementedException:
		return source.haxe__exceptions__PosException.haxe__Exception
	case *haxe__ValueException:
		return source.haxe__Exception
	case *haxe__Exception:
		return source
	default:
		return value.(*haxe__Exception)
	}
}

func hx_recover_class_haxe__ValueException(value any) *haxe__ValueException {
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

func hx_recover_class_haxe__exceptions__PosException(value any) *haxe__exceptions__PosException {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *haxe__exceptions__PosException:
		return source
	case *haxe__exceptions__NotImplementedException:
		return source.haxe__exceptions__PosException
	default:
		return value.(*haxe__exceptions__PosException)
	}
}

func hx_is_nominal_Leaf(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *Leaf:
		return (source != nil)
	case *Tip:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_haxe__Exception(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *haxe__exceptions__PosException:
		return (source != nil)
	case *haxe__exceptions__NotImplementedException:
		return (source != nil)
	case *haxe__ValueException:
		return (source != nil)
	case *haxe__Exception:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_haxe__ValueException(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *haxe__ValueException:
		return (source != nil)
	default:
		return false
	}
}
