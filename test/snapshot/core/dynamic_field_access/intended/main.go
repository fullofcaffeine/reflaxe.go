package main

import "snapshot/hxrt"

type hxrt__TypeClassValue struct {
	name *string
}

type hxrt__TypeEnumValue struct {
	name *string
}

type I_Base interface {
	describe() *string
}

type Base struct {
	__hx_this I_Base
	inherited *string
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
	self.inherited = hxrt.StringFromLiteral("base")
	return self
}

func (self *Base) describe() *string {
	return hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("describe "), self.inherited)
}

func (self *Base) HxClassName() string {
	return "Base"
}

type I_Box interface {
	describe() *string
	get_computed() int
	add(n int) int
}

type Box struct {
	*Base
	__hx_this I_Box
	value     int
	computed  int
	hidden    *string
	scale     *hxrt.Function[func(int) int]
}

func (self *Box) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_Box(value int) *Box {
	return hx_init_Box(hx_alloc_Box(nil), value)
}

func hx_alloc_Box(owner any) *Box {
	self := &Box{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_Box)
	}
	self.Base = hx_alloc_Base(self.__hx_this)
	return self
}

func hx_init_Box(self *Box, value int) *Box {
	hx_init_Base(self.Base)
	self.scale = hxrt.DefaultDynamicFunction(hxrt.NewFunction(func(n int) int {
		return int((hxrt.Int32Wrap(n) * hxrt.Int32Wrap(self.value)))
	}), func() any {
		return self.__hx_this
	}, "scale")
	self.hidden = hxrt.StringFromLiteral("hidden")
	self.value = value
	return self
}

func (self *Box) get_computed() int {
	return int((hxrt.Int32Wrap(self.value) * hxrt.Int32Wrap(2)))
}

func (self *Box) add(n int) int {
	return int((hxrt.Int32Wrap(self.value) + hxrt.Int32Wrap(n)))
}

func (self *Box) HxClassName() string {
	return "Box"
}

func anonymousReads() {
	hx_obj_1 := map[string]any{}
	hx_obj_1["to"] = 2
	hx_obj_1["name"] = hxrt.StringFromLiteral("anon")
	hx_obj_1["flag"] = true
	var o any = hxrt.AnonymousFromMap(hx_obj_1)
	var v any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("anon "), hxrt.StdString(Reflect_field(o, hxrt.StringFromLiteral("to")))), hxrt.StringFromLiteral(" ")), hxrt.StdString(Reflect_field(o, hxrt.StringFromLiteral("name")))))
	hxrt.Println(v)
	var v_1 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("arg "), hxrt.StdString(echo(Reflect_field(o, hxrt.StringFromLiteral("name"))))))
	hxrt.Println(v_1)
	if hxrt.DynamicIsTrue(Reflect_field(o, hxrt.StringFromLiteral("flag"))) {
		hxrt.Println(any(hxrt.StringFromLiteral("cond yes")))
	}
	var v_2 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("ternary "), func() *string {
		var hx_if_2 *string
		if hxrt.DynamicIsTrue(Reflect_field(o, hxrt.StringFromLiteral("flag"))) {
			hx_if_2 = hxrt.StringFromLiteral("on")
		} else {
			hx_if_2 = hxrt.StringFromLiteral("off")
		}
		return hx_if_2
	}()))
	hxrt.Println(v_2)
	var v_3 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("interp to="), hxrt.StdString(Reflect_field(o, hxrt.StringFromLiteral("to")))), hxrt.StringFromLiteral(" name=")), hxrt.StdString(Reflect_field(o, hxrt.StringFromLiteral("name")))))
	hxrt.Println(v_3)
	var v_4 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("eq "), hxrt.StdString(hxrt.HaxeEqual(Reflect_field(o, hxrt.StringFromLiteral("to")), 2))), hxrt.StringFromLiteral(" ")), hxrt.StdString(hxrt.StringEqualAny(Reflect_field(o, hxrt.StringFromLiteral("name")), hxrt.StringFromLiteral("anon")))), hxrt.StringFromLiteral(" ")), hxrt.StdString(!hxrt.HaxeEqual(Reflect_field(o, hxrt.StringFromLiteral("to")), 3))))
	hxrt.Println(v_4)
	typed := hxrt.IntFromNullableAny(Reflect_field(o, hxrt.StringFromLiteral("to")))
	hxrt.Println(any(hxrt.StringConcatAny(hxrt.StringFromLiteral("typed "), int((hxrt.Int32Wrap(typed) + hxrt.Int32Wrap(1))))))
	text := hxrt.DynamicToString(Reflect_field(o, hxrt.StringFromLiteral("name")))
	var v_5 any = any(hxrt.StringConcatAny(hxrt.StringFromLiteral("text "), hxrt.StringLengthStringPtr(text)))
	hxrt.Println(v_5)
	var v_6 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("missing "), hxrt.StdString(Reflect_field(o, hxrt.StringFromLiteral("missing")))), hxrt.StringFromLiteral(" ")), hxrt.StdString(hxrt.AnyEqualsNull(Reflect_field(o, hxrt.StringFromLiteral("missing"))))))
	hxrt.Println(v_6)
	var nothing any = nil
	var v_7 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("null receiver "), hxrt.StdString(Reflect_field(nothing, hxrt.StringFromLiteral("field")))))
	hxrt.Println(v_7)
}

func builtAndTypedObjects() {
	hx_obj_3 := map[string]any{}
	var built any = hxrt.AnonymousFromMap(hx_obj_3)
	Reflect_setField(built, hxrt.StringFromLiteral("a"), 1)
	var v any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("reflect built "), hxrt.StdString(Reflect_field(built, hxrt.StringFromLiteral("a")))))
	hxrt.Println(v)
	hx_obj_4 := map[string]any{}
	hx_obj_4["x"] = 1
	hx_obj_4["label"] = hxrt.StringFromLiteral("p")
	point := hxrt.AnonymousFromMap(hx_obj_4)
	var view any = point
	var v_1 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("typedef "), hxrt.StdString(Reflect_field(view, hxrt.StringFromLiteral("x")))), hxrt.StringFromLiteral(" ")), hxrt.StdString(Reflect_field(view, hxrt.StringFromLiteral("label")))))
	hxrt.Println(v_1)
	func() any {
		hx_dynamic_target_5 := view
		hx_dynamic_value_7 := 5
		Reflect_setField(hx_dynamic_target_5, hxrt.StringFromLiteral("x"), hx_dynamic_value_7)
		return hx_dynamic_value_7
	}()
	var v_2 any = any(hxrt.StringConcatAny(hxrt.StringFromLiteral("typedef shared "), func(hx_value_8 any) int {
		if hx_value_8 == nil {
			var hx_zero_9 int
			return hx_zero_9
		}
		return hx_value_8.(int)
	}(hxrt.AnonymousGet(point, "x"))))
	hxrt.Println(v_2)
}

func calls() {
	hx_obj_10 := map[string]any{}
	hx_obj_10["f"] = hxrt.NewFunction(func(x int) int {
		return int((hxrt.Int32Wrap(x) + hxrt.Int32Wrap(100)))
	})
	var functions any = hxrt.AnonymousFromMap(hx_obj_10)
	var v any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("anon call "), hxrt.StdString(reflaxe__go___internal__CompilerReflect_callMethod(Reflect_field(functions, hxrt.StringFromLiteral("f")), []any{1}))))
	hxrt.Println(v)
	var box any = New_Box(3)
	var v_1 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("method call "), hxrt.StdString(reflaxe__go___internal__CompilerReflect_callMethod(Reflect_field(box, hxrt.StringFromLiteral("add")), []any{2}))), hxrt.StringFromLiteral(" ")), hxrt.StdString(reflaxe__go___internal__CompilerReflect_callMethod(Reflect_field(box, hxrt.StringFromLiteral("describe")), []any{}))))
	hxrt.Println(v_1)
	var v_2 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("dynamic method "), hxrt.StdString(reflaxe__go___internal__CompilerReflect_callMethod(Reflect_field(box, hxrt.StringFromLiteral("scale")), []any{5}))))
	hxrt.Println(v_2)
}

func classInstanceReads() {
	box := New_Box(4)
	var dynamicBox any = box
	var v any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("field "), hxrt.StdString(Reflect_field(dynamicBox, hxrt.StringFromLiteral("value")))), hxrt.StringFromLiteral(" inherited ")), hxrt.StdString(Reflect_field(dynamicBox, hxrt.StringFromLiteral("inherited")))), hxrt.StringFromLiteral(" private ")), hxrt.StdString(Reflect_field(dynamicBox, hxrt.StringFromLiteral("hidden")))))
	hxrt.Println(v)
	var v_1 any = any(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("property without storage "), hxrt.StdString(Reflect_field(dynamicBox, hxrt.StringFromLiteral("computed")))), hxrt.StringFromLiteral(" typed ")), box.get_computed()))
	hxrt.Println(v_1)
	var v_2 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("class missing "), hxrt.StdString(Reflect_field(dynamicBox, hxrt.StringFromLiteral("missing")))))
	hxrt.Println(v_2)
	bound := hx_recover_callable_408c8b7a(Reflect_field(dynamicBox, hxrt.StringFromLiteral("add")))
	first := func() int {
		var hx_callable_11 *hxrt.Function[func(int) int] = bound
		var hx_argument_12 int = 1
		return hxrt.FunctionCall(hx_callable_11)(hx_argument_12)
	}()
	var dynamicBound any = Reflect_field(dynamicBox, hxrt.StringFromLiteral("add"))
	box.value = 10
	second := func() int {
		var hx_callable_13 *hxrt.Function[func(int) int] = bound
		var hx_argument_14 int = 1
		return hxrt.FunctionCall(hx_callable_13)(hx_argument_14)
	}()
	var v_3 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("bound "), first), hxrt.StringFromLiteral(" live ")), second), hxrt.StringFromLiteral(" dynamic ")), hxrt.StdString(reflaxe__go___internal__CompilerReflect_callMethod(dynamicBound, []any{2}))))
	hxrt.Println(v_3)
}

func echo(value any) any {
	return value
}

func evaluationOrder() {
	hx_obj_15 := map[string]any{}
	hx_obj_15["name"] = hxrt.StringFromLiteral("n")
	hx_obj_15["count"] = 0
	var o any = hxrt.AnonymousFromMap(hx_obj_15)
	func() any {
		hx_dynamic_target_16 := step(hxrt.StringFromLiteral("t"), o)
		hx_dynamic_value_18 := step(hxrt.StringFromLiteral("v"), hxrt.StringFromLiteral("w"))
		Reflect_setField(hx_dynamic_target_16, hxrt.StringFromLiteral("name"), hx_dynamic_value_18)
		return hx_dynamic_value_18
	}()
	var v any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("write order "), order), hxrt.StringFromLiteral(" ")), hxrt.StdString(Reflect_field(o, hxrt.StringFromLiteral("name")))))
	hxrt.Println(v)
	order = hxrt.StringFromLiteral("")
	var fh any = step(hxrt.StringFromLiteral("t"), o)
	func() any {
		hx_dynamic_target_19 := fh
		hx_dynamic_value_21 := hxrt.StringConcatStringPtr(hxrt.StdString(Reflect_field(fh, hxrt.StringFromLiteral("name"))), stepText(hxrt.StringFromLiteral("v"), hxrt.StringFromLiteral("+")))
		Reflect_setField(hx_dynamic_target_19, hxrt.StringFromLiteral("name"), hx_dynamic_value_21)
		return hx_dynamic_value_21
	}()
	var v_1 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("concat order "), order), hxrt.StringFromLiteral(" ")), hxrt.StdString(Reflect_field(o, hxrt.StringFromLiteral("name")))))
	hxrt.Println(v_1)
	order = hxrt.StringFromLiteral("")
	var v_2 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("read order "), hxrt.StdString(Reflect_field(step(hxrt.StringFromLiteral("t"), o), hxrt.StringFromLiteral("name")))), hxrt.StdString(step(hxrt.StringFromLiteral("v"), hxrt.StringFromLiteral("")))), hxrt.StringFromLiteral(" ")), order))
	hxrt.Println(v_2)
}

func main() {
	anonymousReads()
	builtAndTypedObjects()
	nestedReads()
	classInstanceReads()
	writes()
	calls()
	evaluationOrder()
}

func nestedReads() {
	hx_obj_22 := map[string]any{}
	hx_obj_23 := map[string]any{}
	hx_obj_23["b"] = 3
	hx_obj_22["a"] = hxrt.AnonymousFromMap(hx_obj_23)
	hx_obj_22["list"] = hxrt.NewArray(1, 2)
	var tree any = hxrt.AnonymousFromMap(hx_obj_22)
	var v any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("nested "), hxrt.StdString(Reflect_field(Reflect_field(tree, hxrt.StringFromLiteral("a")), hxrt.StringFromLiteral("b")))), hxrt.StringFromLiteral(" ")), hxrt.StdString(Reflect_field(Reflect_field(tree, hxrt.StringFromLiteral("a")), hxrt.StringFromLiteral("missing")))), hxrt.StringFromLiteral(" ")), hxrt.StdString(Reflect_field(tree, hxrt.StringFromLiteral("list")))))
	hxrt.Println(v)
	func() any {
		hx_dynamic_target_24 := Reflect_field(tree, hxrt.StringFromLiteral("a"))
		hx_dynamic_value_26 := 4
		Reflect_setField(hx_dynamic_target_24, hxrt.StringFromLiteral("b"), hx_dynamic_value_26)
		return hx_dynamic_value_26
	}()
	var v_1 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("nested write "), hxrt.StdString(Reflect_field(Reflect_field(tree, hxrt.StringFromLiteral("a")), hxrt.StringFromLiteral("b")))))
	hxrt.Println(v_1)
}

var order *string = hxrt.StringFromLiteral("")

func step(name *string, value any) any {
	order = hxrt.StringConcatStringPtr(order, name)
	return value
}

func stepText(name *string, value *string) *string {
	order = hxrt.StringConcatStringPtr(order, name)
	return value
}

func writes() {
	hx_obj_27 := map[string]any{}
	hx_obj_27["to"] = 2
	hx_obj_27["name"] = hxrt.StringFromLiteral("anon")
	var o any = hxrt.AnonymousFromMap(hx_obj_27)
	func() any {
		hx_dynamic_target_28 := o
		hx_dynamic_value_30 := 9
		Reflect_setField(hx_dynamic_target_28, hxrt.StringFromLiteral("to"), hx_dynamic_value_30)
		return hx_dynamic_value_30
	}()
	func() any {
		hx_dynamic_target_31 := o
		hx_dynamic_value_33 := hxrt.StringFromLiteral("added")
		Reflect_setField(hx_dynamic_target_31, hxrt.StringFromLiteral("extra"), hx_dynamic_value_33)
		return hx_dynamic_value_33
	}()
	var v any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("anon write "), hxrt.StdString(Reflect_field(o, hxrt.StringFromLiteral("to")))), hxrt.StringFromLiteral(" ")), hxrt.StdString(Reflect_field(o, hxrt.StringFromLiteral("extra")))), hxrt.StringFromLiteral(" ")), hxrt.StdString(Reflect_hasField(o, hxrt.StringFromLiteral("extra")))))
	hxrt.Println(v)
	func() any {
		hx_dynamic_target_34 := o
		hx_dynamic_value_36 := hxrt.StringConcatStringPtr(hxrt.StdString(Reflect_field(o, hxrt.StringFromLiteral("name"))), hxrt.StringFromLiteral("!"))
		Reflect_setField(hx_dynamic_target_34, hxrt.StringFromLiteral("name"), hx_dynamic_value_36)
		return hx_dynamic_value_36
	}()
	var v_1 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("concat "), hxrt.StdString(Reflect_field(o, hxrt.StringFromLiteral("name")))))
	hxrt.Println(v_1)
	assigned := hxrt.DynamicToString(func() any {
		hx_dynamic_target_37 := o
		hx_dynamic_value_39 := 11
		Reflect_setField(hx_dynamic_target_37, hxrt.StringFromLiteral("to"), hx_dynamic_value_39)
		return hx_dynamic_value_39
	}())
	var v_2 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("assigned "), assigned), hxrt.StringFromLiteral(" ")), hxrt.StdString(Reflect_field(o, hxrt.StringFromLiteral("to")))))
	hxrt.Println(v_2)
	box := New_Box(1)
	var dynamicBox any = box
	func() any {
		hx_dynamic_target_40 := dynamicBox
		hx_dynamic_value_42 := 8
		Reflect_setField(hx_dynamic_target_40, hxrt.StringFromLiteral("value"), hx_dynamic_value_42)
		return hx_dynamic_value_42
	}()
	func() any {
		hx_dynamic_target_43 := dynamicBox
		hx_dynamic_value_45 := hxrt.StringFromLiteral("changed")
		Reflect_setField(hx_dynamic_target_43, hxrt.StringFromLiteral("inherited"), hx_dynamic_value_45)
		return hx_dynamic_value_45
	}()
	var v_3 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("class write "), box.value), hxrt.StringFromLiteral(" ")), box.inherited))
	hxrt.Println(v_3)
}

func hxrt__generated_method_field(obj any, key string) any {
	var receiver any
	switch value := hxrt.ObjectSource(obj).(type) {
	case *Base:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *Box:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__Exception:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__ValueException:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	default:
		return nil
	}
	switch value := receiver.(type) {
	case *Base:
		return hxrt__generated_method_field__Base(value, key)
	case *Box:
		return hxrt__generated_method_field__Box(value, key)
	case *haxe__Exception:
		return hxrt__generated_method_field__haxe__Exception(value, key)
	case *haxe__ValueException:
		return hxrt__generated_method_field__haxe__ValueException(value, key)
	default:
		return nil
	}
}

func hxrt__generated_method_field__Base(value *Base, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "describe":
		return func() *hxrt.Function[func() *string] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_1.__hx_this.describe()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "describe")
		}()
	}
	return nil
}

func hxrt__generated_method_field__Box(value *Box, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "add":
		return func() *hxrt.Function[func(int) int] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(n int) int {
				return hx_receiver_1.__hx_this.add(n)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "add")
		}()
	case "get_computed":
		return func() *hxrt.Function[func() int] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() int {
				return hx_receiver_1.__hx_this.get_computed()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "get_computed")
		}()
	}
	if value.Base == nil {
		return nil
	}
	return hxrt__generated_method_field__Base(value.Base, key)
}

func hxrt__generated_method_field__haxe__Exception(value *haxe__Exception, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "details":
		return func() *hxrt.Function[func() *string] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_1.__hx_this.details()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "details")
		}()
	case "get_message":
		return func() *hxrt.Function[func() *string] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_1.__hx_this.get_message()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "get_message")
		}()
	case "get_native":
		return func() *hxrt.Function[func() any] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() any {
				return hx_receiver_1.__hx_this.get_native()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "get_native")
		}()
	case "get_previous":
		return func() *hxrt.Function[func() *haxe__Exception] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() *haxe__Exception {
				return hx_receiver_1.__hx_this.get_previous()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "get_previous")
		}()
	case "get_stack":
		return func() *hxrt.Function[func() any] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() any {
				return hx_receiver_1.__hx_this.get_stack()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "get_stack")
		}()
	case "toString":
		return func() *hxrt.Function[func() *string] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_1.__hx_this.toString()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "toString")
		}()
	case "unwrap":
		return func() *hxrt.Function[func() any] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() any {
				return hx_receiver_1.__hx_this.unwrap()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "unwrap")
		}()
	}
	return nil
}

func hxrt__generated_method_field__haxe__ValueException(value *haxe__ValueException, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "unwrap":
		return func() *hxrt.Function[func() any] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func() any {
				return hx_receiver_1.__hx_this.unwrap()
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "unwrap")
		}()
	}
	if value.haxe__Exception == nil {
		return nil
	}
	return hxrt__generated_method_field__haxe__Exception(value.haxe__Exception, key)
}

func hxrt_typeClassMetadataField(value any, key string) (any, bool) {
	classValue, ok := value.(*hxrt__TypeClassValue)
	if !ok || classValue == nil {
		return nil, false
	}
	className := *hxrt.StdString(classValue.name)
	switch className {
	default:
		return nil, false
	}
}

func reflaxe__go___internal__CompilerReflect_generatedField(object any, field *string) any {
	key := *hxrt.StdString(field)
	var receiver any
	switch value := hxrt.ObjectSource(object).(type) {
	case *Base:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *Box:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__Exception:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__ValueException:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe___Int64_____Int64:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	default:
		return nil
	}
	switch value := receiver.(type) {
	case *Base:
		return hxrt__generated_field_lookup__Base(value, key)
	case *Box:
		return hxrt__generated_field_lookup__Box(value, key)
	case *haxe__Exception:
		return hxrt__generated_field_lookup__haxe__Exception(value, key)
	case *haxe__ValueException:
		return hxrt__generated_field_lookup__haxe__ValueException(value, key)
	case *haxe___Int64_____Int64:
		return hxrt__generated_field_lookup__haxe___Int64_____Int64(value, key)
	default:
		return nil
	}
}

func reflaxe__go___internal__CompilerReflect_hasGeneratedField(object any, field *string) bool {
	key := *hxrt.StdString(field)
	var receiver any
	switch value := hxrt.ObjectSource(object).(type) {
	case *Base:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *Box:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__Exception:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__ValueException:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe___Int64_____Int64:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	default:
		return false
	}
	switch value := receiver.(type) {
	case *Base:
		return hxrt__generated_field_has__Base(value, key)
	case *Box:
		return hxrt__generated_field_has__Box(value, key)
	case *haxe__Exception:
		return hxrt__generated_field_has__haxe__Exception(value, key)
	case *haxe__ValueException:
		return hxrt__generated_field_has__haxe__ValueException(value, key)
	case *haxe___Int64_____Int64:
		return hxrt__generated_field_has__haxe___Int64_____Int64(value, key)
	default:
		return false
	}
}

func reflaxe__go___internal__CompilerReflect_setGeneratedField(object any, field *string, incoming any) bool {
	key := *hxrt.StdString(field)
	var receiver any
	switch value := hxrt.ObjectSource(object).(type) {
	case *Base:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *Box:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__Exception:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe__ValueException:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *haxe___Int64_____Int64:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	default:
		return false
	}
	switch value := receiver.(type) {
	case *Base:
		return hxrt__generated_field_set__Base(value, key, incoming)
	case *Box:
		return hxrt__generated_field_set__Box(value, key, incoming)
	case *haxe__Exception:
		return hxrt__generated_field_set__haxe__Exception(value, key, incoming)
	case *haxe__ValueException:
		return hxrt__generated_field_set__haxe__ValueException(value, key, incoming)
	case *haxe___Int64_____Int64:
		return hxrt__generated_field_set__haxe___Int64_____Int64(value, key, incoming)
	default:
		return false
	}
}

func reflaxe__go___internal__CompilerReflect_generatedFields(object any) *hxrt.Array {
	var receiver any
	switch value := hxrt.ObjectSource(object).(type) {
	case *Base:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *Box:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__Exception:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe__ValueException:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *haxe___Int64_____Int64:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	default:
		return nil
	}
	switch receiver.(type) {
	case *Base:
		return hxrt.NewArray(hxrt.StringFromLiteral("inherited"))
	case *Box:
		return hxrt.NewArray(hxrt.StringFromLiteral("inherited"), hxrt.StringFromLiteral("value"), hxrt.StringFromLiteral("hidden"), hxrt.StringFromLiteral("scale"))
	case *haxe__Exception:
		return hxrt.NewArray(hxrt.StringFromLiteral("savedMessage"), hxrt.StringFromLiteral("savedStack"), hxrt.StringFromLiteral("savedPrevious"), hxrt.StringFromLiteral("savedNative"))
	case *haxe__ValueException:
		return hxrt.NewArray(hxrt.StringFromLiteral("savedMessage"), hxrt.StringFromLiteral("savedStack"), hxrt.StringFromLiteral("savedPrevious"), hxrt.StringFromLiteral("savedNative"), hxrt.StringFromLiteral("value"))
	case *haxe___Int64_____Int64:
		return hxrt.NewArray(hxrt.StringFromLiteral("high"), hxrt.StringFromLiteral("low"))
	default:
		return nil
	}
}

func hxrt__generated_field_lookup__Base(value *Base, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "inherited":
		return hxrt.ReadStoredField(value.inherited)
	}
	return nil
}

func hxrt__generated_field_has__Base(value *Base, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "inherited":
		return true
	}
	return false
}

func hxrt__generated_field_set__Base(value *Base, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "inherited":
		if incoming == nil {
			var zero *string
			value.inherited = zero
			return true
		}
		switch typed := incoming.(type) {
		case *string:
			value.inherited = typed
			return true
		default:
			return false
		}
	}
	return false
}

func hxrt__generated_field_lookup__Box(value *Box, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "hidden":
		return hxrt.ReadStoredField(value.hidden)
	case "scale":
		return hxrt.ReadStoredField(value.scale)
	case "value":
		return hxrt.ReadStoredField(value.value)
	}
	if value.Base == nil {
		return nil
	}
	return hxrt__generated_field_lookup__Base(value.Base, key)
}

func hxrt__generated_field_has__Box(value *Box, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "hidden":
		return true
	case "scale":
		return true
	case "value":
		return true
	}
	if value.Base == nil {
		return false
	}
	return hxrt__generated_field_has__Base(value.Base, key)
}

func hxrt__generated_field_set__Box(value *Box, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "hidden":
		if incoming == nil {
			var zero *string
			value.hidden = zero
			return true
		}
		switch typed := incoming.(type) {
		case *string:
			value.hidden = typed
			return true
		default:
			return false
		}
	case "scale":
		if (hxrt.FunctionOrigin(incoming) == nil) || hxrt.ReflectIsFunction(incoming) {
			value.scale = hx_recover_callable_408c8b7a(incoming)
			return true
		}
		return false
	case "value":
		if incoming == nil {
			var zero int
			value.value = zero
			return true
		}
		switch typed := incoming.(type) {
		case int:
			value.value = typed
			return true
		default:
			return false
		}
	}
	if value.Base == nil {
		return false
	}
	return hxrt__generated_field_set__Base(value.Base, key, incoming)
}

func hxrt__generated_field_lookup__haxe__Exception(value *haxe__Exception, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "savedMessage":
		return hxrt.ReadStoredField(value.savedMessage)
	case "savedNative":
		return hxrt.ReadStoredField(value.savedNative)
	case "savedPrevious":
		return hxrt.ReadStoredField(value.savedPrevious)
	case "savedStack":
		return hxrt.ReadStoredField(value.savedStack)
	}
	return nil
}

func hxrt__generated_field_has__haxe__Exception(value *haxe__Exception, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "savedMessage":
		return true
	case "savedNative":
		return true
	case "savedPrevious":
		return true
	case "savedStack":
		return true
	}
	return false
}

func hxrt__generated_field_set__haxe__Exception(value *haxe__Exception, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "savedMessage":
		if incoming == nil {
			var zero *string
			value.savedMessage = zero
			return true
		}
		switch typed := incoming.(type) {
		case *string:
			value.savedMessage = typed
			return true
		default:
			return false
		}
	case "savedNative":
		if incoming == nil {
			var zero any
			value.savedNative = zero
			return true
		}
		switch typed := incoming.(type) {
		case any:
			value.savedNative = typed
			return true
		default:
			return false
		}
	case "savedPrevious":
		if incoming == nil {
			var zero *haxe__Exception
			value.savedPrevious = zero
			return true
		}
		switch typed := incoming.(type) {
		case *haxe__Exception:
			value.savedPrevious = typed
			return true
		default:
			return false
		}
	case "savedStack":
		if incoming == nil {
			var zero any
			value.savedStack = zero
			return true
		}
		switch typed := incoming.(type) {
		case any:
			value.savedStack = typed
			return true
		default:
			return false
		}
	}
	return false
}

func hxrt__generated_field_lookup__haxe__ValueException(value *haxe__ValueException, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "value":
		return hxrt.ReadStoredField(value.value)
	}
	if value.haxe__Exception == nil {
		return nil
	}
	return hxrt__generated_field_lookup__haxe__Exception(value.haxe__Exception, key)
}

func hxrt__generated_field_has__haxe__ValueException(value *haxe__ValueException, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "value":
		return true
	}
	if value.haxe__Exception == nil {
		return false
	}
	return hxrt__generated_field_has__haxe__Exception(value.haxe__Exception, key)
}

func hxrt__generated_field_set__haxe__ValueException(value *haxe__ValueException, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "value":
		if incoming == nil {
			var zero any
			value.value = zero
			return true
		}
		switch typed := incoming.(type) {
		case any:
			value.value = typed
			return true
		default:
			return false
		}
	}
	if value.haxe__Exception == nil {
		return false
	}
	return hxrt__generated_field_set__haxe__Exception(value.haxe__Exception, key, incoming)
}

func hxrt__generated_field_lookup__haxe___Int64_____Int64(value *haxe___Int64_____Int64, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "high":
		return hxrt.ReadStoredField(value.high)
	case "low":
		return hxrt.ReadStoredField(value.low)
	}
	return nil
}

func hxrt__generated_field_has__haxe___Int64_____Int64(value *haxe___Int64_____Int64, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "high":
		return true
	case "low":
		return true
	}
	return false
}

func hxrt__generated_field_set__haxe___Int64_____Int64(value *haxe___Int64_____Int64, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "high":
		if incoming == nil {
			var zero int
			value.high = zero
			return true
		}
		switch typed := incoming.(type) {
		case int:
			value.high = typed
			return true
		default:
			return false
		}
	case "low":
		if incoming == nil {
			var zero int
			value.low = zero
			return true
		}
		switch typed := incoming.(type) {
		case int:
			value.low = typed
			return true
		default:
			return false
		}
	}
	return false
}

func reflaxe__go___internal__CompilerReflect_typeField(object any, field *string) any {
	key := *hxrt.StdString(field)
	value, found := hxrt_typeClassMetadataField(object, key)
	if !found {
		return nil
	}
	return value
}

func reflaxe__go___internal__CompilerReflect_hasTypeField(object any, field *string) bool {
	key := *hxrt.StdString(field)
	_, found := hxrt_typeClassMetadataField(object, key)
	return found
}

func reflaxe__go___internal__CompilerReflect_generatedMethod(object any, field *string) any {
	key := *hxrt.StdString(field)
	return hxrt__generated_method_field(object, key)
}

func reflaxe__go___internal__CompilerReflect_isEnumValue(value any) bool {
	switch enumValue := value.(type) {
	case *haxe__StackItem:
		return (enumValue != nil)
	default:
		return false
	}
}

func reflaxe__go___internal__CompilerReflect_callMethod(value any, arguments []any) any {
	origin := hxrt.FunctionOrigin(value)
	if origin == nil {
		hxrt.Throw(hxrt.StringFromLiteral("Cannot call null"))
		return nil
	}
	switch source := origin.(type) {
	case *hxrt.Function[func(int) int]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(func(hx_value_1 any) int {
			if hx_value_1 == nil {
				var hx_zero_2 int
				return hx_zero_2
			}
			return hx_value_1.(int)
		}(argument_0))
	case *hxrt.Function[func(...any) any]:
		var argument_0 []any = nil
		for index, argument := range arguments {
			if index >= 0 {
				argument_0 = append(argument_0, argument)
			}
		}
		return hxrt.FunctionCall(source)(argument_0...)
	case *hxrt.Function[func() *string]:
		return hxrt.FunctionCall(source)()
	case *hxrt.Function[func() int]:
		return hxrt.FunctionCall(source)()
	case *hxrt.Function[func() any]:
		return hxrt.FunctionCall(source)()
	case *hxrt.Function[func() *haxe__Exception]:
		return hxrt.FunctionCall(source)()
	default:
		return hxrt.ReflectCallMethod(value, arguments)
	}
}

func hx_recover_callable_408c8b7a(value any) *hxrt.Function[func(int) int] {
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
	case *hxrt.Function[func(...any) any]:
		return hxrt.ProjectFunction(source, func(arg_0 int) int {
			return func(hx_value_1 any) int {
				if hx_value_1 == nil {
					var hx_zero_2 int
					return hx_zero_2
				}
				return hx_value_1.(int)
			}(hxrt.FunctionCall(source)(arg_0))
		})
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}

func hx_recover_class_haxe__Exception(value any) *haxe__Exception {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *haxe__Exception:
		return source
	case *haxe__ValueException:
		return source.haxe__Exception
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

func hx_is_nominal_haxe__Exception(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *haxe__Exception:
		return (source != nil)
	case *haxe__ValueException:
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
