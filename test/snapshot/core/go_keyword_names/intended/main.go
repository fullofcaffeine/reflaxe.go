package main

import "snapshot/hxrt"

type hxrt__TypeClassValue struct {
	name *string
}

type hxrt__TypeEnumValue struct {
	name *string
}

type Shape struct {
	tag    int
	params []any
}

func (self *Shape) HxEnumValue() (string, []any) {
	return [1]string{"Box"}[self.tag], self.params
}

func Shape_Box(struct_ int, const_ int, goto_ int) *Shape {
	enumValue := &Shape{tag: 0}
	enumValue.params = []any{struct_, const_, goto_}
	return enumValue
}

type I_Fields interface {
}

type Fields struct {
	__hx_this I_Fields
	struct_   int
	const_    int
	goto_     int
}

func (self *Fields) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_Fields(struct_ int, const_ int, goto_ int) *Fields {
	return hx_init_Fields(hx_alloc_Fields(nil), struct_, const_, goto_)
}

func hx_alloc_Fields(owner any) *Fields {
	self := &Fields{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_Fields)
	}
	return self
}

func hx_init_Fields(self *Fields, struct_ int, const_ int, goto_ int) *Fields {
	self.struct_ = struct_
	self.const_ = const_
	self.goto_ = goto_
	return self
}

func (self *Fields) HxClassName() string {
	return "Fields"
}

var const_ int = 2

func goto_(struct_ int, const_ int, goto_ any) int {
	return int((hxrt.Int32Wrap(int((hxrt.Int32Wrap(struct_) + hxrt.Int32Wrap(const_)))) + hxrt.Int32Wrap(hxrt.IntFromNullableAny(hxrt.IntFromNullableAny(func() any {
		var hx_if_1 any
		if goto_ == nil {
			hx_if_1 = 0
		} else {
			hx_if_1 = goto_.(int)
		}
		return hx_if_1
	}())))))
}

func locals() {
	struct_ := 3
	const_ := 4
	goto_ := 5
	hxrt.Println(any(int((hxrt.Int32Wrap(int((hxrt.Int32Wrap(struct_) + hxrt.Int32Wrap(const_)))) + hxrt.Int32Wrap(goto_)))))
	twice := hxrt.NewFunction(func(struct__1 int) int {
		return int((hxrt.Int32Wrap(struct__1) * hxrt.Int32Wrap(2)))
	})
	var v any = any(func() int {
		var hx_callable_2 *hxrt.Function[func(int) int] = twice
		var hx_argument_3 int = 21
		return hxrt.FunctionCall(hx_callable_2)(hx_argument_3)
	}())
	hxrt.Println(v)
	const__1 := hxrt.NewFunction(func(goto__1 int) int {
		return int((hxrt.Int32Wrap(goto__1) + hxrt.Int32Wrap(100)))
	})
	var v_1 any = any(func() int {
		var hx_callable_4 *hxrt.Function[func(int) int] = const__1
		var hx_argument_5 int = 1
		return hxrt.FunctionCall(hx_callable_4)(hx_argument_5)
	}())
	hxrt.Println(v_1)
	total := 0
	total = int((hxrt.Int32Wrap(total) + hxrt.Int32Wrap(0)))
	total = int((hxrt.Int32Wrap(total) + hxrt.Int32Wrap(1)))
	total = int((hxrt.Int32Wrap(total) + hxrt.Int32Wrap(2)))
	total = int((hxrt.Int32Wrap(total) + hxrt.Int32Wrap(3)))
	hxrt.Println(any(total))
	hxrt.TryCatch(func() {
		hxrt.Throw(haxe__Exception_thrown(hxrt.StringFromLiteral("boom")))
	}, func(hx_caught_6 any) {
		hx_exception_7 := haxe__Exception_caught(hx_caught_6)
		hx_payload_8 := haxe__Exception_catchValue(hx_caught_6)
		_ = hx_exception_7
		if func(hx_value any) bool {
			switch hx_value.(type) {
			case *string:
				return true
			default:
				return false
			}
		}(hx_payload_8) {
			var struct__1 *string = func(hx_value_9 any) *string {
				if hx_value_9 == nil {
					var hx_zero_10 *string
					return hx_zero_10
				}
				return hx_value_9.(*string)
			}(hx_payload_8)
			hxrt.Println(any(struct__1))
		} else {
			hxrt.Throw(hx_caught_6)
		}
	})
	_g := Shape_Box(1, 2, 3)
	_g1 := _g.params[0].(int)
	_g2 := _g.params[1].(int)
	_g_1 := _g.params[2].(int)
	struct__2 := _g1
	const__2 := _g2
	goto__1 := _g_1
	hxrt.Println(any(int((hxrt.Int32Wrap(int((hxrt.Int32Wrap(int((hxrt.Int32Wrap(struct__2) * hxrt.Int32Wrap(100)))) + hxrt.Int32Wrap(int((hxrt.Int32Wrap(const__2) * hxrt.Int32Wrap(10))))))) + hxrt.Int32Wrap(goto__1)))))
}

func main() {
	statics()
	locals()
	members()
}

func members() {
	fields := New_Fields(6, 7, 8)
	fields.goto_ = int((hxrt.Int32Wrap(fields.goto_) + hxrt.Int32Wrap(1)))
	var v any = any(int((hxrt.Int32Wrap(int((hxrt.Int32Wrap(fields.struct_) + hxrt.Int32Wrap(fields.const_)))) + hxrt.Int32Wrap(fields.goto_))))
	hxrt.Println(v)
	var v_1 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StdString(Reflect_field(fields, hxrt.StringFromLiteral("struct"))), hxrt.StringFromLiteral(",")), hxrt.StdString(Reflect_field(fields, hxrt.StringFromLiteral("const")))), hxrt.StringFromLiteral(",")), hxrt.StdString(Reflect_field(fields, hxrt.StringFromLiteral("goto")))))
	hxrt.Println(v_1)
	methods := New_Methods()
	var v_2 any = any(int((hxrt.Int32Wrap(int((hxrt.Int32Wrap(methods.struct_(1)) + hxrt.Int32Wrap(methods.const_(2))))) + hxrt.Int32Wrap(methods.goto_(10)))))
	hxrt.Println(v_2)
	bound := func() *hxrt.Function[func(int) int] {
		hx_receiver_11 := methods
		return hxrt.BoundFunction(func(value int) int {
			return hx_receiver_11.__hx_this.const_(value)
		}, func() any {
			return hx_receiver_11.__hx_this
		}, "const")
	}()
	var v_3 any = any(func() int {
		var hx_callable_12 *hxrt.Function[func(int) int] = bound
		var hx_argument_13 int = 5
		return hxrt.FunctionCall(hx_callable_12)(hx_argument_13)
	}())
	hxrt.Println(v_3)
	var dynamicMethods any = methods
	var v_4 any = Reflect_callMethod(dynamicMethods, Reflect_field(dynamicMethods, hxrt.StringFromLiteral("goto")), hxrt.NewArray(20))
	hxrt.Println(v_4)
	var record_struct int
	var record_goto int
	var record_const int
	record_struct = 7
	record_const = 8
	record_goto = 9
	var v_5 any = any(int((hxrt.Int32Wrap(int((hxrt.Int32Wrap(record_struct) + hxrt.Int32Wrap(record_const)))) + hxrt.Int32Wrap(record_goto))))
	hxrt.Println(v_5)
	hx_obj_14 := map[string]any{}
	hx_obj_14["break"] = 1
	hx_obj_14["continue"] = 2
	hx_obj_14["import"] = 3
	hx_obj_14["interface"] = 4
	var quoted any = hxrt.AnonymousFromMap(hx_obj_14)
	var v_6 any = hxrt.DynamicAdd(hxrt.DynamicAdd(hxrt.DynamicAdd(Reflect_field(quoted, hxrt.StringFromLiteral("break")), Reflect_field(quoted, hxrt.StringFromLiteral("continue"))), Reflect_field(quoted, hxrt.StringFromLiteral("import"))), Reflect_field(quoted, hxrt.StringFromLiteral("interface")))
	hxrt.Println(v_6)
}

func statics() {
	var v any = any(int((hxrt.Int32Wrap(int((hxrt.Int32Wrap(int((hxrt.Int32Wrap(struct_) + hxrt.Int32Wrap(const_)))) + hxrt.Int32Wrap(goto_(1, 2, nil))))) + hxrt.Int32Wrap(goto_(1, 2, 3)))))
	hxrt.Println(v)
	var v_1 any = any(int((hxrt.Int32Wrap(int((hxrt.Int32Wrap(Statics_struct_()) + hxrt.Int32Wrap(Statics_const_())))) + hxrt.Int32Wrap(Statics_goto_))))
	hxrt.Println(v_1)
}

var struct_ int = 1

type I_Methods interface {
	struct_(value int) int
	const_(value int) int
	goto_(value int) int
}

type Methods struct {
	__hx_this I_Methods
}

func (self *Methods) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_Methods() *Methods {
	return hx_init_Methods(hx_alloc_Methods(nil))
}

func hx_alloc_Methods(owner any) *Methods {
	self := &Methods{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_Methods)
	}
	return self
}

func hx_init_Methods(self *Methods) *Methods {
	return self
}

func (self *Methods) struct_(value int) int {
	return int((hxrt.Int32Wrap(value) + hxrt.Int32Wrap(1)))
}

func (self *Methods) const_(value int) int {
	return int((hxrt.Int32Wrap(value) * hxrt.Int32Wrap(2)))
}

func (self *Methods) goto_(value int) int {
	return int((hxrt.Int32Wrap(value) - hxrt.Int32Wrap(3)))
}

func (self *Methods) HxClassName() string {
	return "Methods"
}

func Statics_const_() int {
	return 20
}

var Statics_goto_ int = 40

func Statics_struct_() int {
	return 10
}

func hxrt__generated_method_field(obj any, key string) any {
	var receiver any
	switch value := hxrt.ObjectSource(obj).(type) {
	case *Methods:
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
	case *Methods:
		return hxrt__generated_method_field__Methods(value, key)
	case *haxe__Exception:
		return hxrt__generated_method_field__haxe__Exception(value, key)
	case *haxe__ValueException:
		return hxrt__generated_method_field__haxe__ValueException(value, key)
	default:
		return nil
	}
}

func hxrt__generated_method_field__Methods(value *Methods, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "const":
		return func() *hxrt.Function[func(int) int] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(value int) int {
				return hx_receiver_1.__hx_this.const_(value)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "const")
		}()
	case "goto":
		return func() *hxrt.Function[func(int) int] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(value int) int {
				return hx_receiver_1.__hx_this.goto_(value)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "goto")
		}()
	case "struct":
		return func() *hxrt.Function[func(int) int] {
			hx_receiver_1 := value
			return hxrt.BoundFunction(func(value int) int {
				return hx_receiver_1.__hx_this.struct_(value)
			}, func() any {
				return hx_receiver_1.__hx_this
			}, "struct")
		}()
	}
	return nil
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
	case *Fields:
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
	case *Fields:
		return hxrt__generated_field_lookup__Fields(value, key)
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
	case *Fields:
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
	case *Fields:
		return hxrt__generated_field_has__Fields(value, key)
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
	case *Fields:
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
	case *Fields:
		return hxrt__generated_field_set__Fields(value, key, incoming)
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
	case *Fields:
		if (value == nil) || (value.__hx_this == nil) {
			return nil
		}
		receiver = value.__hx_this
	case *Methods:
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
	case *Fields:
		return hxrt.NewArray(hxrt.StringFromLiteral("struct"), hxrt.StringFromLiteral("const"), hxrt.StringFromLiteral("goto"))
	case *Methods:
		return hxrt.NewArray()
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

func hxrt__generated_field_lookup__Fields(value *Fields, key string) any {
	if value == nil {
		return nil
	}
	switch key {
	case "const":
		return hxrt.ReadStoredField(value.const_)
	case "goto":
		return hxrt.ReadStoredField(value.goto_)
	case "struct":
		return hxrt.ReadStoredField(value.struct_)
	}
	return nil
}

func hxrt__generated_field_has__Fields(value *Fields, key string) bool {
	if value == nil {
		return false
	}
	switch key {
	case "const":
		return true
	case "goto":
		return true
	case "struct":
		return true
	}
	return false
}

func hxrt__generated_field_set__Fields(value *Fields, key string, incoming any) bool {
	if value == nil {
		return false
	}
	switch key {
	case "const":
		if incoming == nil {
			var zero int
			value.const_ = zero
			return true
		}
		switch typed := incoming.(type) {
		case int:
			value.const_ = typed
			return true
		default:
			return false
		}
	case "goto":
		if incoming == nil {
			var zero int
			value.goto_ = zero
			return true
		}
		switch typed := incoming.(type) {
		case int:
			value.goto_ = typed
			return true
		default:
			return false
		}
	case "struct":
		if incoming == nil {
			var zero int
			value.struct_ = zero
			return true
		}
		switch typed := incoming.(type) {
		case int:
			value.struct_ = typed
			return true
		default:
			return false
		}
	}
	return false
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
	case *Shape:
		return (enumValue != nil)
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
	case *hxrt.Function[func() any]:
		return hxrt.FunctionCall(source)()
	case *hxrt.Function[func() *haxe__Exception]:
		return hxrt.FunctionCall(source)()
	default:
		return hxrt.ReflectCallMethod(value, arguments)
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
