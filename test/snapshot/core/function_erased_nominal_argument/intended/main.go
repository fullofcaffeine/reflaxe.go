package main

import "snapshot/hxrt"

type hxrt__TypeClassValue struct {
	name *string
}

type hxrt__TypeEnumValue struct {
	name *string
}

type I_Base interface {
}

type Base struct {
	__hx_this I_Base
	n         int
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
	self.n = 3
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

type I_Grandchild interface {
	name() *string
}

type Grandchild struct {
	*Child
	__hx_this I_Grandchild
}

func (self *Grandchild) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_Grandchild() *Grandchild {
	return hx_init_Grandchild(hx_alloc_Grandchild(nil))
}

func hx_alloc_Grandchild(owner any) *Grandchild {
	self := &Grandchild{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_Grandchild)
	}
	self.Child = hx_alloc_Child(self.__hx_this)
	return self
}

func hx_init_Grandchild(self *Grandchild) *Grandchild {
	hx_init_Child(self.Child)
	return self
}

func call(callback *hxrt.Function[func(any) int], value any) int {
	return func() int {
		var hx_callable_1 *hxrt.Function[func(any) int] = callback
		var hx_argument_2 any = value
		return hxrt.FunctionCall(hx_callable_1)(hx_argument_2)
	}()
}

func main() {
	original := hxrt.NewFunction(func(value *Base) int {
		return value.n
	})
	var v any = any(call(hx_recover_callable_0(original), func(hx_value *Child) *Base {
		if hx_value == nil {
			return nil
		}
		return hx_value.Base
	}(New_Child())))
	hxrt.Println(v)
	var v_1 any = Reflect_callMethod(nil, original, hxrt.NewArray(New_Child()))
	hxrt.Println(v_1)
	child := New_Grandchild()
	base := func(hx_value *Grandchild) *Base {
		if hx_value == nil {
			return nil
		}
		return hx_value.Child.Base
	}(child)
	read := hxrt.NewFunction(func(value *Base) int {
		var hx_if_3 int
		if value == nil {
			hx_if_3 = -1
		} else {
			hx_if_3 = value.n
		}
		return hx_if_3
	})
	var v_2 any = Reflect_callMethod(nil, read, hxrt.NewArray(child))
	hxrt.Println(v_2)
	var v_3 any = Reflect_callMethod(nil, read, hxrt.NewArray(base))
	hxrt.Println(v_3)
	var v_4 any = Reflect_callMethod(nil, read, hxrt.NewArray(nil))
	hxrt.Println(v_4)
	name := hxrt.NewFunction(func(value Named) *string {
		return value.name()
	})
	var v_5 any = Reflect_callMethod(nil, name, hxrt.NewArray(base))
	hxrt.Println(v_5)
	create := hxrt.NewFunction(func() *Base {
		return func(hx_value *Grandchild) *Base {
			if hx_value == nil {
				return nil
			}
			return hx_value.Child.Base
		}(child)
	})
	var v_6 any = any(hx_recover_class_0(Reflect_callMethod(nil, create, hxrt.NewArray())).n)
	hxrt.Println(v_6)
}

type Named interface {
	name() *string
}

func hxrt__generated_method_field(obj any, key string) any {
	var receiver any
	switch value := hxrt.ObjectSource(obj).(type) {
	case *Base:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *Child:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *Grandchild:
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
	case *Child:
		return hxrt__generated_method_field__Child(value, key)
	case *Grandchild:
		return hxrt__generated_method_field__Grandchild(value, key)
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
	return nil
}

func hxrt__generated_method_field__Child(value *Child, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "name":
		return func() *hxrt.Function[func() *string] {
			hx_receiver_18 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_18.__hx_this.name()
			}, func() any {
				return hx_receiver_18.__hx_this
			}, "name")
		}()
	}
	if value.Base == nil {
		return nil
	}
	return hxrt__generated_method_field__Base(value.Base, key)
}

func hxrt__generated_method_field__Grandchild(value *Grandchild, key string) any {
	if value == nil {
		return nil
	}
	if value.Child == nil {
		return nil
	}
	return hxrt__generated_method_field__Child(value.Child, key)
}

func hxrt__generated_method_field__haxe__Exception(value *haxe__Exception, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "details":
		return func() *hxrt.Function[func() *string] {
			hx_receiver_19 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_19.__hx_this.details()
			}, func() any {
				return hx_receiver_19.__hx_this
			}, "details")
		}()
	case "get_message":
		return func() *hxrt.Function[func() *string] {
			hx_receiver_20 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_20.__hx_this.get_message()
			}, func() any {
				return hx_receiver_20.__hx_this
			}, "get_message")
		}()
	case "get_native":
		return func() *hxrt.Function[func() any] {
			hx_receiver_21 := value
			return hxrt.BoundFunction(func() any {
				return hx_receiver_21.__hx_this.get_native()
			}, func() any {
				return hx_receiver_21.__hx_this
			}, "get_native")
		}()
	case "get_previous":
		return func() *hxrt.Function[func() *haxe__Exception] {
			hx_receiver_22 := value
			return hxrt.BoundFunction(func() *haxe__Exception {
				return hx_receiver_22.__hx_this.get_previous()
			}, func() any {
				return hx_receiver_22.__hx_this
			}, "get_previous")
		}()
	case "get_stack":
		return func() *hxrt.Function[func() any] {
			hx_receiver_23 := value
			return hxrt.BoundFunction(func() any {
				return hx_receiver_23.__hx_this.get_stack()
			}, func() any {
				return hx_receiver_23.__hx_this
			}, "get_stack")
		}()
	case "toString":
		return func() *hxrt.Function[func() *string] {
			hx_receiver_24 := value
			return hxrt.BoundFunction(func() *string {
				return hx_receiver_24.__hx_this.toString()
			}, func() any {
				return hx_receiver_24.__hx_this
			}, "toString")
		}()
	case "unwrap":
		return func() *hxrt.Function[func() any] {
			hx_receiver_25 := value
			return hxrt.BoundFunction(func() any {
				return hx_receiver_25.__hx_this.unwrap()
			}, func() any {
				return hx_receiver_25.__hx_this
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
			hx_receiver_26 := value
			return hxrt.BoundFunction(func() any {
				return hx_receiver_26.__hx_this.unwrap()
			}, func() any {
				return hx_receiver_26.__hx_this
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
	case *Child:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *Grandchild:
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
	case *Child:
		return hxrt__generated_field_lookup__Child(value, key)
	case *Grandchild:
		return hxrt__generated_field_lookup__Grandchild(value, key)
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
	case *Child:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *Grandchild:
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
	case *Child:
		return hxrt__generated_field_has__Child(value, key)
	case *Grandchild:
		return hxrt__generated_field_has__Grandchild(value, key)
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
	case *Child:
		if (value == nil) || (value.__hx_this == nil) {
			return false
		}
		receiver = value.__hx_this
	case *Grandchild:
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
	case *Child:
		return hxrt__generated_field_set__Child(value, key, incoming)
	case *Grandchild:
		return hxrt__generated_field_set__Grandchild(value, key, incoming)
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
	case *Child:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *Grandchild:
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
		return hxrt.NewArray(hxrt.StringFromLiteral("n"))
	case *Child:
		return hxrt.NewArray(hxrt.StringFromLiteral("n"))
	case *Grandchild:
		return hxrt.NewArray(hxrt.StringFromLiteral("n"))
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
	case "n":
		return hxrt.ReadStoredField(value.n)
	}
	return nil
}

func hxrt__generated_field_has__Base(value *Base, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "n":
		return true
	}
	return false
}

func hxrt__generated_field_set__Base(value *Base, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "n":
		if incoming == nil {
			var zero int
			value.n = zero
			return true
		}
		switch typed := incoming.(type) {
		case int:
			value.n = typed
			return true
		default:
			return false
		}
	}
	return false
}

func hxrt__generated_field_lookup__Child(value *Child, key string) any {
	if value == nil {
		return nil
	}
	if value.Base == nil {
		return nil
	}
	return hxrt__generated_field_lookup__Base(value.Base, key)
}

func hxrt__generated_field_has__Child(value *Child, key string) bool {
	if value == nil {
		return false
	}
	if value.Base == nil {
		return false
	}
	return hxrt__generated_field_has__Base(value.Base, key)
}

func hxrt__generated_field_set__Child(value *Child, key string, incoming any) bool {
	if value == nil {
		return false
	}
	if value.Base == nil {
		return false
	}
	return hxrt__generated_field_set__Base(value.Base, key, incoming)
}

func hxrt__generated_field_lookup__Grandchild(value *Grandchild, key string) any {
	if value == nil {
		return nil
	}
	if value.Child == nil {
		return nil
	}
	return hxrt__generated_field_lookup__Child(value.Child, key)
}

func hxrt__generated_field_has__Grandchild(value *Grandchild, key string) bool {
	if value == nil {
		return false
	}
	if value.Child == nil {
		return false
	}
	return hxrt__generated_field_has__Child(value.Child, key)
}

func hxrt__generated_field_set__Grandchild(value *Grandchild, key string, incoming any) bool {
	if value == nil {
		return false
	}
	if value.Child == nil {
		return false
	}
	return hxrt__generated_field_set__Child(value.Child, key, incoming)
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
	case *hxrt.Function[func(*Base) int]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(hx_recover_class_0(argument_0))
	case *hxrt.Function[func(Named) *string]:
		var argument_0 any = nil
		if 0 < len(arguments) {
			argument_0 = arguments[0]
		}
		return hxrt.FunctionCall(source)(func(hx_value_27 any) Named {
			if hx_value_27 == nil {
				var hx_zero_28 Named
				return hx_zero_28
			}
			return hx_value_27.(Named)
		}(hxrt.ObjectSource(argument_0)))
	case *hxrt.Function[func() *Base]:
		return hxrt.FunctionCall(source)()
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
	case *hxrt.Function[func() any]:
		return hxrt.FunctionCall(source)()
	case *hxrt.Function[func() *haxe__Exception]:
		return hxrt.FunctionCall(source)()
	default:
		return hxrt.ReflectCallMethod(value, arguments)
	}
}

func hx_recover_callable_0(value any) *hxrt.Function[func(any) int] {
	if value == nil {
		return nil
	}
	switch exact := value.(type) {
	case *hxrt.Function[func(any) int]:
		return exact
	}
	origin := hxrt.FunctionOrigin(value)
	if origin == nil {
		return nil
	}
	switch source := origin.(type) {
	case *hxrt.Function[func(*Base) int]:
		return hxrt.ProjectFunction(source, func(arg_0 any) int {
			return hxrt.FunctionCall(source)(hx_recover_class_0(arg_0))
		})
	case *hxrt.Function[func(Named) *string]:
		return hxrt.ProjectFunction(source, func(arg_0 any) int {
			return func(hx_value_31 any) int {
				if hx_value_31 == nil {
					var hx_zero_32 int
					return hx_zero_32
				}
				return hx_value_31.(int)
			}(any(hxrt.FunctionCall(source)(func(hx_value_29 any) Named {
				if hx_value_29 == nil {
					var hx_zero_30 Named
					return hx_zero_30
				}
				return hx_value_29.(Named)
			}(hxrt.ObjectSource(arg_0)))))
		})
	case *hxrt.Function[func(...any) any]:
		return hxrt.ProjectFunction(source, func(arg_0 any) int {
			return func(hx_value_33 any) int {
				if hx_value_33 == nil {
					var hx_zero_34 int
					return hx_zero_34
				}
				return hx_value_33.(int)
			}(any(hxrt.FunctionCall(source)(arg_0)))
		})
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
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
	case *Grandchild:
		return source.Child.Base
	default:
		return value.(*Base)
	}
}

func hx_recover_class_1(value any) *haxe__ValueException {
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

func hx_recover_class_2(value any) *haxe__Exception {
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
