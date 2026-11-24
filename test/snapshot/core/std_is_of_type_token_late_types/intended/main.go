package main

import "snapshot/hxrt"

type hxrt__TypeClassValue struct {
	name *string
}

type hxrt__TypeEnumValue struct {
	name *string
}

func ClassTokenCheck_isClass(value any, token any) bool {
	return hx_is_of_type_token(value, token)
}

func ClassTokenCheck_isEnum(value any, token any) bool {
	return hx_is_of_type_token(value, token)
}

func fail() {
	hxrt.Throw(haxe__Exception_thrown(New_haxe__exceptions__NotImplementedException(nil, nil, func() hxrt.Anonymous {
		hx_obj_1 := map[string]any{}
		hx_obj_1["fileName"] = hxrt.StringFromLiteral("Main.hx")
		hx_obj_1["lineNumber"] = 31
		hx_obj_1["className"] = hxrt.StringFromLiteral("Main")
		hx_obj_1["methodName"] = hxrt.StringFromLiteral("fail")
		return hxrt.AnonymousFromMap(hx_obj_1)
	}())))
}

func main() {
	var error any = New_haxe__exceptions__NotImplementedException(nil, nil, func() hxrt.Anonymous {
		hx_obj_2 := map[string]any{}
		hx_obj_2["fileName"] = hxrt.StringFromLiteral("Main.hx")
		hx_obj_2["lineNumber"] = 35
		hx_obj_2["className"] = hxrt.StringFromLiteral("Main")
		hx_obj_2["methodName"] = hxrt.StringFromLiteral("main")
		return hxrt.AnonymousFromMap(hx_obj_2)
	}())
	var v any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("class token NotImplementedException="), hxrt.StdString(ClassTokenCheck_isClass(error, &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe.exceptions.NotImplementedException")}))))
	hxrt.Println(v)
	var v_1 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("class token PosException="), hxrt.StdString(ClassTokenCheck_isClass(error, &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe.exceptions.PosException")}))))
	hxrt.Println(v_1)
	var v_2 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("class token Exception="), hxrt.StdString(ClassTokenCheck_isClass(error, &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe.Exception")}))))
	hxrt.Println(v_2)
	var v_3 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("class token String="), hxrt.StdString(ClassTokenCheck_isClass(error, &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("String")}))))
	hxrt.Println(v_3)
	var v_4 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("class token on string="), hxrt.StdString(ClassTokenCheck_isClass(hxrt.StringFromLiteral("text"), &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("haxe.exceptions.PosException")}))))
	hxrt.Println(v_4)
	var option any = haxe__ds__Option_Some(1)
	var v_5 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("enum token Option="), hxrt.StdString(ClassTokenCheck_isEnum(option, &hxrt__TypeEnumValue{name: hxrt.StringFromLiteral("haxe.ds.Option")}))))
	hxrt.Println(v_5)
	var v_6 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("enum token on error="), hxrt.StdString(ClassTokenCheck_isEnum(error, &hxrt__TypeEnumValue{name: hxrt.StringFromLiteral("haxe.ds.Option")}))))
	hxrt.Println(v_6)
	hxrt.TryCatch(func() {
		fail()
	}, func(hx_caught_3 any) {
		hx_exception_4 := haxe__Exception_caught(hx_caught_3)
		hx_payload_5 := haxe__Exception_catchValue(hx_caught_3)
		_ = hx_payload_5
		if hx_is_nominal_haxe__exceptions__PosException(hx_exception_4) {
			var caught *haxe__exceptions__PosException = hx_recover_class_haxe__exceptions__PosException(hx_exception_4)
			var v_7 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("caught as PosException="), hxrt.StdString(hx_is_nominal_haxe__exceptions__NotImplementedException(caught))))
			hxrt.Println(v_7)
		} else {
			hxrt.Throw(hx_caught_3)
		}
	})
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
		case "haxe.exceptions.NotImplementedException":
			return hx_is_nominal_haxe__exceptions__NotImplementedException(hx_value)
		case "haxe.exceptions.PosException":
			return hx_is_nominal_haxe__exceptions__PosException(hx_value)
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
		case "haxe.ds.Option":
			return func(hx_value any) bool {
				switch hx_value.(type) {
				case *haxe__ds__Option:
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

func hx_is_nominal_haxe___Int64_____Int64(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *haxe___Int64_____Int64:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_haxe__exceptions__NotImplementedException(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *haxe__exceptions__NotImplementedException:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_haxe__exceptions__PosException(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *haxe__exceptions__PosException:
		return (source != nil)
	case *haxe__exceptions__NotImplementedException:
		return (source != nil)
	default:
		return false
	}
}
