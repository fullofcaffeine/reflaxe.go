package main

import "snapshot/hxrt"

type hxrt__TypeClassValue struct {
	name *string
}

type hxrt__TypeEnumValue struct {
	name *string
}

func consume(socket *sys__net__Socket) *string {
	line := socket.input.__hx_this.readLine()
	lower := hxrt.StringToLowerCaseStringPtr(line)
	var parsed any = Std_parseInt(hxrt.StringFromLiteral("42"))
	first := socket.input.__hx_this.readByte()
	buffer := haxe__io__Bytes_alloc(4)
	read := socket.input.__hx_this.readBytes(buffer, 0, buffer.length)
	rest := socket.input.__hx_this.readAll(nil).__hx_this.toString()
	return hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(lower, hxrt.StringFromLiteral(":")), parsed), hxrt.StringFromLiteral(":")), first), hxrt.StringFromLiteral(":")), read), hxrt.StringFromLiteral(":")), rest)
}

func main() {
	hxrt.Println(any(hxrt.StringFromLiteral("socket-input-service-surface")))
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
		case "haxe.io.Bytes":
			return hx_is_nominal_haxe__io__Bytes(hx_value)
		case "haxe.io.BytesBuffer":
			return hx_is_nominal_haxe__io__BytesBuffer(hx_value)
		case "haxe.io.Eof":
			return hx_is_nominal_haxe__io__Eof(hx_value)
		case "haxe.io.Input":
			return hx_is_nominal_haxe__io__Input(hx_value)
		case "haxe.io.Output":
			return hx_is_nominal_haxe__io__Output(hx_value)
		case "sys.net.Host":
			return hx_is_nominal_sys__net__Host(hx_value)
		case "sys.net.Socket":
			return hx_is_nominal_sys__net__Socket(hx_value)
		case "sys.net.SocketInput":
			return hx_is_nominal_sys__net__SocketInput(hx_value)
		case "sys.net.SocketOutput":
			return hx_is_nominal_sys__net__SocketOutput(hx_value)
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
		case "haxe.io.Encoding":
			return func(hx_value any) bool {
				switch hx_value.(type) {
				case *haxe__io__Encoding:
					return true
				default:
					return false
				}
			}(hx_value)
		case "haxe.io.Error":
			return func(hx_value any) bool {
				switch hx_value.(type) {
				case *haxe__io__Error:
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
	case *haxe__exceptions__NotImplementedException:
		return source.haxe__exceptions__PosException.haxe__Exception
	case *haxe__Exception:
		return source
	case *haxe__exceptions__PosException:
		return source.haxe__Exception
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

func hx_recover_class_haxe__io__Eof(value any) *haxe__io__Eof {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *haxe__io__Eof:
		return source
	default:
		return value.(*haxe__io__Eof)
	}
}

func hx_recover_class_sys__net__Socket(value any) *sys__net__Socket {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *sys__net__Socket:
		return source
	default:
		return value.(*sys__net__Socket)
	}
}

func hx_is_nominal_haxe__Exception(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *haxe__exceptions__NotImplementedException:
		return (source != nil)
	case *haxe__Exception:
		return (source != nil)
	case *haxe__exceptions__PosException:
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
	case *haxe__exceptions__NotImplementedException:
		return (source != nil)
	case *haxe__exceptions__PosException:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_haxe__io__Bytes(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *haxe__io__Bytes:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_haxe__io__BytesBuffer(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *haxe__io__BytesBuffer:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_haxe__io__Eof(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *haxe__io__Eof:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_haxe__io__Input(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *sys__net__SocketInput:
		return (source != nil)
	case *haxe__io__Input:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_haxe__io__Output(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *sys__net__SocketOutput:
		return (source != nil)
	case *haxe__io__Output:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_sys__net__Host(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *sys__net__Host:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_sys__net__Socket(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *sys__net__Socket:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_sys__net__SocketInput(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *sys__net__SocketInput:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_sys__net__SocketOutput(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *sys__net__SocketOutput:
		return (source != nil)
	default:
		return false
	}
}
