package main

import "snapshot/hxrt"

type hxrt__TypeClassValue struct {
	name *string
}

type hxrt__TypeEnumValue struct {
	name *string
}

type I_Animal interface {
	sound() *string
}

type Animal struct {
	__hx_this I_Animal
}

func (self *Animal) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_Animal() *Animal {
	return hx_init_Animal(hx_alloc_Animal(nil))
}

func hx_alloc_Animal(owner any) *Animal {
	self := &Animal{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_Animal)
	}
	return self
}

func hx_init_Animal(self *Animal) *Animal {
	return self
}

func (self *Animal) sound() *string {
	return hxrt.StringFromLiteral("animal")
}

func (self *Animal) HxClassName() string {
	return "Animal"
}

type I_Dog interface {
	sound() *string
}

type Dog struct {
	*Animal
	__hx_this I_Dog
}

func (self *Dog) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_Dog() *Dog {
	return hx_init_Dog(hx_alloc_Dog(nil))
}

func hx_alloc_Dog(owner any) *Dog {
	self := &Dog{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_Dog)
	}
	self.Animal = hx_alloc_Animal(self.__hx_this)
	return self
}

func hx_init_Dog(self *Dog) *Dog {
	hx_init_Animal(self.Animal)
	return self
}

func (self *Dog) sound() *string {
	return hxrt.StringFromLiteral("dog")
}

func (self *Dog) HxClassName() string {
	return "Dog"
}

func main() {
	n := _Main__Integer_Impl__fromString(hxrt.StringFromLiteral("5"))
	hxrt.Println(any(hxrt.StringConcatAny(hxrt.StringFromLiteral("Integer from String: "), int((hxrt.Int32Wrap(n) + hxrt.Int32Wrap(1))))))
	var half any = 2.5
	var v any = any(hxrt.StringConcatAny(hxrt.StringFromLiteral("Dynamic as Real: "), (func(hx_cast_1 any) float64 {
		if hxrt.AnyEqualsNull(hx_cast_1) {
			var hx_zero_2 float64
			return hx_zero_2
		}
		if !func(hx_value any) bool {
			switch hx_value.(type) {
			case int:
				return true
			case int8:
				return true
			case int16:
				return true
			case int32:
				return true
			case int64:
				return true
			case uint:
				return true
			case uint8:
				return true
			case uint16:
				return true
			case uint32:
				return true
			case uint64:
				return true
			case uintptr:
				return true
			case float32:
				return true
			case float64:
				return true
			default:
				return false
			}
		}(hx_cast_1) {
			hxrt.Throw(haxe__Exception_thrown(hxrt.StringFromLiteral("Class cast error")))
		}
		return hxrt.DynamicToFloatSlot(hx_cast_1)
	}(half) * float64(2))))
	hxrt.Println(v)
	var seven any = 7
	var v_1 any = any(hxrt.StringConcatAny(hxrt.StringFromLiteral("Dynamic Int as Real: "), (func(hx_cast_3 any) float64 {
		if hxrt.AnyEqualsNull(hx_cast_3) {
			var hx_zero_4 float64
			return hx_zero_4
		}
		if !func(hx_value any) bool {
			switch hx_value.(type) {
			case int:
				return true
			case int8:
				return true
			case int16:
				return true
			case int32:
				return true
			case int64:
				return true
			case uint:
				return true
			case uint8:
				return true
			case uint16:
				return true
			case uint32:
				return true
			case uint64:
				return true
			case uintptr:
				return true
			case float32:
				return true
			case float64:
				return true
			default:
				return false
			}
		}(hx_cast_3) {
			hxrt.Throw(haxe__Exception_thrown(hxrt.StringFromLiteral("Class cast error")))
		}
		return hxrt.DynamicToFloatSlot(hx_cast_3)
	}(seven) / float64(2))))
	hxrt.Println(v_1)
	var dog any = New_Dog()
	var v_2 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Dynamic as Pet: "), hx_recover_class_Animal(func(hx_cast_5 any) any {
		if !hxrt.AnyEqualsNull(hx_cast_5) && !hx_is_nominal_Animal(hx_cast_5) {
			hxrt.Throw(haxe__Exception_thrown(hxrt.StringFromLiteral("Class cast error")))
		}
		return hx_cast_5
	}(dog)).__hx_this.sound()))
	hxrt.Println(v_2)
	var one any = 1
	var v_3 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("Dynamic as Color: "), hxrt.StdString((func(hx_cast_6 any) int {
		if hxrt.AnyEqualsNull(hx_cast_6) {
			var hx_zero_7 int
			return hx_zero_7
		}
		if !func(hx_value any) bool {
			switch hx_value.(type) {
			case int:
				return true
			case int8:
				return true
			case int16:
				return true
			case int32:
				return true
			case int64:
				return true
			case uint:
				return true
			case uint8:
				return true
			case uint16:
				return true
			case uint32:
				return true
			case uint64:
				return true
			case uintptr:
				return true
			default:
				return false
			}
		}(hx_cast_6) {
			hxrt.Throw(haxe__Exception_thrown(hxrt.StringFromLiteral("Class cast error")))
		}
		return hxrt.IntFromNullableAny(hx_cast_6)
	}(one) == 1))))
	hxrt.Println(v_3)
	var text any = hxrt.StringFromLiteral("text")
	hxrt.TryCatch(func() {
		func(hx_cast_13 any) int {
			if hxrt.AnyEqualsNull(hx_cast_13) {
				var hx_zero_14 int
				return hx_zero_14
			}
			if !func(hx_value any) bool {
				switch hx_value.(type) {
				case int:
					return true
				case int8:
					return true
				case int16:
					return true
				case int32:
					return true
				case int64:
					return true
				case uint:
					return true
				case uint8:
					return true
				case uint16:
					return true
				case uint32:
					return true
				case uint64:
					return true
				case uintptr:
					return true
				default:
					return false
				}
			}(hx_cast_13) {
				hxrt.Throw(haxe__Exception_thrown(hxrt.StringFromLiteral("Class cast error")))
			}
			return hxrt.IntFromNullableAny(hx_cast_13)
		}(text)
		hxrt.Println(any(hxrt.StringFromLiteral("not reached")))
	}, func(hx_caught_8 any) {
		hx_exception_9 := haxe__Exception_caught(hx_caught_8)
		hx_payload_10 := haxe__Exception_catchValue(hx_caught_8)
		_ = hx_exception_9
		if func(hx_value any) bool {
			switch hx_value.(type) {
			case *string:
				return true
			default:
				return false
			}
		}(hx_payload_10) {
			var message *string = func(hx_value_11 any) *string {
				if hx_value_11 == nil {
					var hx_zero_12 *string
					return hx_zero_12
				}
				return hx_value_11.(*string)
			}(hx_payload_10)
			hxrt.Println(any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("caught: "), message)))
		} else {
			hxrt.Throw(hx_caught_8)
		}
	})
}

func _Main__Integer_Impl__fromString(s *string) int {
	var n any = Std_parseInt(s)
	return func(hx_value_15 any) int {
		if hx_value_15 == nil {
			var hx_zero_16 int
			return hx_zero_16
		}
		return hx_value_15.(int)
	}(n)
}

func hx_is_of_type_token(hx_value any, hx_type any) bool {
	switch hx_type_marker := hx_type.(type) {
	case *hxrt__TypeClassValue:
		if hx_type_marker == nil {
			return false
		}
		if hx_type_marker.name == nil {
			return false
		}
		switch *hx_type_marker.name {
		case "Animal":
			return hx_is_nominal_Animal(hx_value)
		case "Array":
			return func(hx_value any) bool {
				switch hx_value.(type) {
				case *hxrt.Array:
					return true
				default:
					return false
				}
			}(hx_value)
		case "Bool":
			return func(hx_value any) bool {
				switch hx_value.(type) {
				case bool:
					return true
				default:
					return false
				}
			}(hx_value)
		case "Class":
			return func(hx_value any) bool {
				switch hx_value.(type) {
				case *hxrt__TypeClassValue:
					return true
				default:
					return false
				}
			}(hx_value)
		case "Dog":
			return hx_is_nominal_Dog(hx_value)
		case "Dynamic":
			return (hx_value != nil)
		case "Enum":
			return func(hx_value any) bool {
				switch hx_value.(type) {
				case *hxrt__TypeEnumValue:
					return true
				default:
					return false
				}
			}(hx_value)
		case "Float":
			return func(hx_value any) bool {
				switch hx_value.(type) {
				case int:
					return true
				case int8:
					return true
				case int16:
					return true
				case int32:
					return true
				case int64:
					return true
				case uint:
					return true
				case uint8:
					return true
				case uint16:
					return true
				case uint32:
					return true
				case uint64:
					return true
				case uintptr:
					return true
				case float32:
					return true
				case float64:
					return true
				default:
					return false
				}
			}(hx_value)
		case "Int":
			return func(hx_value any) bool {
				switch hx_value.(type) {
				case int:
					return true
				case int8:
					return true
				case int16:
					return true
				case int32:
					return true
				case int64:
					return true
				case uint:
					return true
				case uint8:
					return true
				case uint16:
					return true
				case uint32:
					return true
				case uint64:
					return true
				case uintptr:
					return true
				default:
					return false
				}
			}(hx_value)
		case "String":
			return func(hx_value any) bool {
				switch hx_value.(type) {
				case *string:
					return true
				case string:
					return true
				default:
					return false
				}
			}(hx_value)
		case "haxe.Exception":
			return hx_is_nominal_haxe__Exception(hx_value)
		case "haxe.ValueException":
			return hx_is_nominal_haxe__ValueException(hx_value)
		case "haxe._Int64.___Int64":
			return hx_is_nominal_haxe___Int64_____Int64(hx_value)
		default:
			return false
		}
	case *hxrt__TypeEnumValue:
		if hx_type_marker == nil {
			return false
		}
		if hx_type_marker.name == nil {
			return false
		}
		switch *hx_type_marker.name {
		case "haxe.StackItem":
			return func(hx_value any) bool {
				switch hx_value.(type) {
				case *haxe__StackItem:
					return true
				default:
					return false
				}
			}(hx_value)
		default:
			return false
		}
	default:
		return false
	}
}

func hx_recover_class_Animal(value any) *Animal {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *Animal:
		return source
	case *Dog:
		return source.Animal
	default:
		return value.(*Animal)
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

func hx_is_nominal_Animal(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *Animal:
		return (source != nil)
	case *Dog:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_Dog(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *Dog:
		return (source != nil)
	default:
		return false
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

func hx_is_nominal_haxe___Int64_____Int64(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *haxe___Int64_____Int64:
		return (source != nil)
	default:
		return false
	}
}
