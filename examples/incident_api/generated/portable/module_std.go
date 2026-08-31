package main

import (
	"examples_incident_api_portable/hxrt"
	"math/rand"
)

func Std_digitValue(code int, hexadecimal bool) int {
	if (code >= 48) && (code <= 57) {
		return int((hxrt.Int32Wrap(code) - hxrt.Int32Wrap(48)))
	}
	if (hexadecimal && (code >= 97)) && (code <= 102) {
		return int((hxrt.Int32Wrap(int((hxrt.Int32Wrap(code) - hxrt.Int32Wrap(97)))) + hxrt.Int32Wrap(10)))
	}
	if (hexadecimal && (code >= 65)) && (code <= 70) {
		return int((hxrt.Int32Wrap(int((hxrt.Int32Wrap(code) - hxrt.Int32Wrap(65)))) + hxrt.Int32Wrap(10)))
	}
	return -1
}

func Std_downcast(value any, c any) any {
	var hx_if_1 any
	if func(hx_value any, hx_type any) bool {
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
			case "Date":
				return func(hx_value any) bool {
					switch hx_carrier := hx_value.(type) {
					case *Date:
						if hx_carrier == nil {
							return false
						}
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
			case "app.core.Incident":
				return func(hx_value any) bool {
					switch hx_carrier := hx_value.(type) {
					case *app__core__Incident:
						if hx_carrier == nil {
							return false
						}
						return true
					default:
						return false
					}
				}(hx_value)
			case "app.core.IncidentApi":
				return func(hx_value any) bool {
					switch hx_carrier := hx_value.(type) {
					case *app__core__IncidentApi:
						if hx_carrier == nil {
							return false
						}
						return true
					default:
						return false
					}
				}(hx_value)
			case "app.core.IncidentConfig":
				return func(hx_value any) bool {
					switch hx_carrier := hx_value.(type) {
					case *app__core__IncidentConfig:
						if hx_carrier == nil {
							return false
						}
						return true
					default:
						return false
					}
				}(hx_value)
			case "app.core.IncidentRequestException":
				return func(hx_value any) bool {
					switch hx_carrier := hx_value.(type) {
					case *app__core__IncidentRequestException:
						if hx_carrier == nil {
							return false
						}
						return true
					default:
						return false
					}
				}(hx_value)
			case "app.core.IncidentStore":
				return func(hx_value any) bool {
					switch hx_carrier := hx_value.(type) {
					case *app__core__IncidentStore:
						if hx_carrier == nil {
							return false
						}
						return true
					default:
						return false
					}
				}(hx_value)
			case "app.http.HttpRequest":
				return func(hx_value any) bool {
					switch hx_carrier := hx_value.(type) {
					case *app__http__HttpRequest:
						if hx_carrier == nil {
							return false
						}
						return true
					default:
						return false
					}
				}(hx_value)
			case "app.http.HttpResponse":
				return func(hx_value any) bool {
					switch hx_carrier := hx_value.(type) {
					case *app__http__HttpResponse:
						if hx_carrier == nil {
							return false
						}
						return true
					default:
						return false
					}
				}(hx_value)
			case "app.http.TinyHttpServer":
				return func(hx_value any) bool {
					switch hx_carrier := hx_value.(type) {
					case *app__http__TinyHttpServer:
						if hx_carrier == nil {
							return false
						}
						return true
					default:
						return false
					}
				}(hx_value)
			case "haxe._Int64.___Int64":
				return func(hx_value any) bool {
					switch hx_carrier := hx_value.(type) {
					case *haxe___Int64_____Int64:
						if hx_carrier == nil {
							return false
						}
						return true
					default:
						return false
					}
				}(hx_value)
			case "haxe.exceptions.NotImplementedException":
				return func(hx_value any) bool {
					switch hx_carrier := hx_value.(type) {
					case *haxe__exceptions__NotImplementedException:
						if hx_carrier == nil {
							return false
						}
						return true
					case *haxe__exceptions__PosException:
						if hx_carrier == nil {
							return false
						}
						_, hx_ok := hx_carrier.__hx_this.(*haxe__exceptions__NotImplementedException)
						return hx_ok
					default:
						return false
					}
				}(hx_value)
			case "haxe.exceptions.PosException":
				return func(hx_value any) bool {
					switch hx_carrier := hx_value.(type) {
					case *haxe__exceptions__NotImplementedException:
						if hx_carrier == nil {
							return false
						}
						return true
					case *haxe__exceptions__PosException:
						if hx_carrier == nil {
							return false
						}
						return true
					default:
						return false
					}
				}(hx_value)
			case "haxe.io.Bytes":
				return func(hx_value any) bool {
					switch hx_carrier := hx_value.(type) {
					case *haxe__io__Bytes:
						if hx_carrier == nil {
							return false
						}
						return true
					default:
						return false
					}
				}(hx_value)
			case "haxe.io.BytesBuffer":
				return func(hx_value any) bool {
					switch hx_carrier := hx_value.(type) {
					case *haxe__io__BytesBuffer:
						if hx_carrier == nil {
							return false
						}
						return true
					default:
						return false
					}
				}(hx_value)
			case "haxe.io.Eof":
				return func(hx_value any) bool {
					switch hx_carrier := hx_value.(type) {
					case *haxe__io__Eof:
						if hx_carrier == nil {
							return false
						}
						return true
					default:
						return false
					}
				}(hx_value)
			case "haxe.io.Input":
				return func(hx_value any) bool {
					switch hx_carrier := hx_value.(type) {
					case *haxe__io__Input:
						if hx_carrier == nil {
							return false
						}
						return true
					case *sys__io__FileInput:
						if hx_carrier == nil {
							return false
						}
						return true
					case *sys__net__SocketInput:
						if hx_carrier == nil {
							return false
						}
						return true
					default:
						return false
					}
				}(hx_value)
			case "haxe.io.Output":
				return func(hx_value any) bool {
					switch hx_carrier := hx_value.(type) {
					case *haxe__io__Output:
						if hx_carrier == nil {
							return false
						}
						return true
					case *sys__io__FileOutput:
						if hx_carrier == nil {
							return false
						}
						return true
					case *sys__net__SocketOutput:
						if hx_carrier == nil {
							return false
						}
						return true
					default:
						return false
					}
				}(hx_value)
			case "haxe.iterators.StringIterator":
				return func(hx_value any) bool {
					switch hx_carrier := hx_value.(type) {
					case *haxe__iterators__StringIterator:
						if hx_carrier == nil {
							return false
						}
						return true
					default:
						return false
					}
				}(hx_value)
			case "haxe.iterators.StringKeyValueIterator":
				return func(hx_value any) bool {
					switch hx_carrier := hx_value.(type) {
					case *haxe__iterators__StringKeyValueIterator:
						if hx_carrier == nil {
							return false
						}
						return true
					default:
						return false
					}
				}(hx_value)
			case "sys.io.FileInput":
				return func(hx_value any) bool {
					switch hx_carrier := hx_value.(type) {
					case *haxe__io__Input:
						if hx_carrier == nil {
							return false
						}
						_, hx_ok := hx_carrier.__hx_this.(*sys__io__FileInput)
						return hx_ok
					case *sys__io__FileInput:
						if hx_carrier == nil {
							return false
						}
						return true
					default:
						return false
					}
				}(hx_value)
			case "sys.io.FileOutput":
				return func(hx_value any) bool {
					switch hx_carrier := hx_value.(type) {
					case *haxe__io__Output:
						if hx_carrier == nil {
							return false
						}
						_, hx_ok := hx_carrier.__hx_this.(*sys__io__FileOutput)
						return hx_ok
					case *sys__io__FileOutput:
						if hx_carrier == nil {
							return false
						}
						return true
					default:
						return false
					}
				}(hx_value)
			case "sys.net.Host":
				return func(hx_value any) bool {
					switch hx_carrier := hx_value.(type) {
					case *sys__net__Host:
						if hx_carrier == nil {
							return false
						}
						return true
					default:
						return false
					}
				}(hx_value)
			case "sys.net.Socket":
				return func(hx_value any) bool {
					switch hx_carrier := hx_value.(type) {
					case *sys__net__Socket:
						if hx_carrier == nil {
							return false
						}
						return true
					default:
						return false
					}
				}(hx_value)
			case "sys.net.SocketInput":
				return func(hx_value any) bool {
					switch hx_carrier := hx_value.(type) {
					case *haxe__io__Input:
						if hx_carrier == nil {
							return false
						}
						_, hx_ok := hx_carrier.__hx_this.(*sys__net__SocketInput)
						return hx_ok
					case *sys__net__SocketInput:
						if hx_carrier == nil {
							return false
						}
						return true
					default:
						return false
					}
				}(hx_value)
			case "sys.net.SocketOutput":
				return func(hx_value any) bool {
					switch hx_carrier := hx_value.(type) {
					case *haxe__io__Output:
						if hx_carrier == nil {
							return false
						}
						_, hx_ok := hx_carrier.__hx_this.(*sys__net__SocketOutput)
						return hx_ok
					case *sys__net__SocketOutput:
						if hx_carrier == nil {
							return false
						}
						return true
					default:
						return false
					}
				}(hx_value)
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
			default:
				return false
			}
		default:
			return false
		}
	}(value, c) {
		hx_if_1 = value
	} else {
		hx_if_1 = nil
	}
	return hx_if_1
}

func Std_instance(value any, c any) any {
	return func() any {
		var hx_if_2 any
		if func(hx_value any, hx_type any) bool {
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
				case "Date":
					return func(hx_value any) bool {
						switch hx_carrier := hx_value.(type) {
						case *Date:
							if hx_carrier == nil {
								return false
							}
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
				case "app.core.Incident":
					return func(hx_value any) bool {
						switch hx_carrier := hx_value.(type) {
						case *app__core__Incident:
							if hx_carrier == nil {
								return false
							}
							return true
						default:
							return false
						}
					}(hx_value)
				case "app.core.IncidentApi":
					return func(hx_value any) bool {
						switch hx_carrier := hx_value.(type) {
						case *app__core__IncidentApi:
							if hx_carrier == nil {
								return false
							}
							return true
						default:
							return false
						}
					}(hx_value)
				case "app.core.IncidentConfig":
					return func(hx_value any) bool {
						switch hx_carrier := hx_value.(type) {
						case *app__core__IncidentConfig:
							if hx_carrier == nil {
								return false
							}
							return true
						default:
							return false
						}
					}(hx_value)
				case "app.core.IncidentRequestException":
					return func(hx_value any) bool {
						switch hx_carrier := hx_value.(type) {
						case *app__core__IncidentRequestException:
							if hx_carrier == nil {
								return false
							}
							return true
						default:
							return false
						}
					}(hx_value)
				case "app.core.IncidentStore":
					return func(hx_value any) bool {
						switch hx_carrier := hx_value.(type) {
						case *app__core__IncidentStore:
							if hx_carrier == nil {
								return false
							}
							return true
						default:
							return false
						}
					}(hx_value)
				case "app.http.HttpRequest":
					return func(hx_value any) bool {
						switch hx_carrier := hx_value.(type) {
						case *app__http__HttpRequest:
							if hx_carrier == nil {
								return false
							}
							return true
						default:
							return false
						}
					}(hx_value)
				case "app.http.HttpResponse":
					return func(hx_value any) bool {
						switch hx_carrier := hx_value.(type) {
						case *app__http__HttpResponse:
							if hx_carrier == nil {
								return false
							}
							return true
						default:
							return false
						}
					}(hx_value)
				case "app.http.TinyHttpServer":
					return func(hx_value any) bool {
						switch hx_carrier := hx_value.(type) {
						case *app__http__TinyHttpServer:
							if hx_carrier == nil {
								return false
							}
							return true
						default:
							return false
						}
					}(hx_value)
				case "haxe._Int64.___Int64":
					return func(hx_value any) bool {
						switch hx_carrier := hx_value.(type) {
						case *haxe___Int64_____Int64:
							if hx_carrier == nil {
								return false
							}
							return true
						default:
							return false
						}
					}(hx_value)
				case "haxe.exceptions.NotImplementedException":
					return func(hx_value any) bool {
						switch hx_carrier := hx_value.(type) {
						case *haxe__exceptions__NotImplementedException:
							if hx_carrier == nil {
								return false
							}
							return true
						case *haxe__exceptions__PosException:
							if hx_carrier == nil {
								return false
							}
							_, hx_ok := hx_carrier.__hx_this.(*haxe__exceptions__NotImplementedException)
							return hx_ok
						default:
							return false
						}
					}(hx_value)
				case "haxe.exceptions.PosException":
					return func(hx_value any) bool {
						switch hx_carrier := hx_value.(type) {
						case *haxe__exceptions__NotImplementedException:
							if hx_carrier == nil {
								return false
							}
							return true
						case *haxe__exceptions__PosException:
							if hx_carrier == nil {
								return false
							}
							return true
						default:
							return false
						}
					}(hx_value)
				case "haxe.io.Bytes":
					return func(hx_value any) bool {
						switch hx_carrier := hx_value.(type) {
						case *haxe__io__Bytes:
							if hx_carrier == nil {
								return false
							}
							return true
						default:
							return false
						}
					}(hx_value)
				case "haxe.io.BytesBuffer":
					return func(hx_value any) bool {
						switch hx_carrier := hx_value.(type) {
						case *haxe__io__BytesBuffer:
							if hx_carrier == nil {
								return false
							}
							return true
						default:
							return false
						}
					}(hx_value)
				case "haxe.io.Eof":
					return func(hx_value any) bool {
						switch hx_carrier := hx_value.(type) {
						case *haxe__io__Eof:
							if hx_carrier == nil {
								return false
							}
							return true
						default:
							return false
						}
					}(hx_value)
				case "haxe.io.Input":
					return func(hx_value any) bool {
						switch hx_carrier := hx_value.(type) {
						case *haxe__io__Input:
							if hx_carrier == nil {
								return false
							}
							return true
						case *sys__io__FileInput:
							if hx_carrier == nil {
								return false
							}
							return true
						case *sys__net__SocketInput:
							if hx_carrier == nil {
								return false
							}
							return true
						default:
							return false
						}
					}(hx_value)
				case "haxe.io.Output":
					return func(hx_value any) bool {
						switch hx_carrier := hx_value.(type) {
						case *haxe__io__Output:
							if hx_carrier == nil {
								return false
							}
							return true
						case *sys__io__FileOutput:
							if hx_carrier == nil {
								return false
							}
							return true
						case *sys__net__SocketOutput:
							if hx_carrier == nil {
								return false
							}
							return true
						default:
							return false
						}
					}(hx_value)
				case "haxe.iterators.StringIterator":
					return func(hx_value any) bool {
						switch hx_carrier := hx_value.(type) {
						case *haxe__iterators__StringIterator:
							if hx_carrier == nil {
								return false
							}
							return true
						default:
							return false
						}
					}(hx_value)
				case "haxe.iterators.StringKeyValueIterator":
					return func(hx_value any) bool {
						switch hx_carrier := hx_value.(type) {
						case *haxe__iterators__StringKeyValueIterator:
							if hx_carrier == nil {
								return false
							}
							return true
						default:
							return false
						}
					}(hx_value)
				case "sys.io.FileInput":
					return func(hx_value any) bool {
						switch hx_carrier := hx_value.(type) {
						case *haxe__io__Input:
							if hx_carrier == nil {
								return false
							}
							_, hx_ok := hx_carrier.__hx_this.(*sys__io__FileInput)
							return hx_ok
						case *sys__io__FileInput:
							if hx_carrier == nil {
								return false
							}
							return true
						default:
							return false
						}
					}(hx_value)
				case "sys.io.FileOutput":
					return func(hx_value any) bool {
						switch hx_carrier := hx_value.(type) {
						case *haxe__io__Output:
							if hx_carrier == nil {
								return false
							}
							_, hx_ok := hx_carrier.__hx_this.(*sys__io__FileOutput)
							return hx_ok
						case *sys__io__FileOutput:
							if hx_carrier == nil {
								return false
							}
							return true
						default:
							return false
						}
					}(hx_value)
				case "sys.net.Host":
					return func(hx_value any) bool {
						switch hx_carrier := hx_value.(type) {
						case *sys__net__Host:
							if hx_carrier == nil {
								return false
							}
							return true
						default:
							return false
						}
					}(hx_value)
				case "sys.net.Socket":
					return func(hx_value any) bool {
						switch hx_carrier := hx_value.(type) {
						case *sys__net__Socket:
							if hx_carrier == nil {
								return false
							}
							return true
						default:
							return false
						}
					}(hx_value)
				case "sys.net.SocketInput":
					return func(hx_value any) bool {
						switch hx_carrier := hx_value.(type) {
						case *haxe__io__Input:
							if hx_carrier == nil {
								return false
							}
							_, hx_ok := hx_carrier.__hx_this.(*sys__net__SocketInput)
							return hx_ok
						case *sys__net__SocketInput:
							if hx_carrier == nil {
								return false
							}
							return true
						default:
							return false
						}
					}(hx_value)
				case "sys.net.SocketOutput":
					return func(hx_value any) bool {
						switch hx_carrier := hx_value.(type) {
						case *haxe__io__Output:
							if hx_carrier == nil {
								return false
							}
							_, hx_ok := hx_carrier.__hx_this.(*sys__net__SocketOutput)
							return hx_ok
						case *sys__net__SocketOutput:
							if hx_carrier == nil {
								return false
							}
							return true
						default:
							return false
						}
					}(hx_value)
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
				default:
					return false
				}
			default:
				return false
			}
		}(value, c) {
			hx_if_2 = value
		} else {
			hx_if_2 = nil
		}
		return hx_if_2
	}()
}

func Std_int(x float64) int {
	return hxrt.MathTruncInt(x)
}

func Std_invalidFloat() float64 {
	return hxrt.StringParseFloatExact(hxrt.StringFromLiteral(""))
}

func Std_is(v any, t any) bool {
	return func(hx_value any, hx_type any) bool {
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
			case "Date":
				return func(hx_value any) bool {
					switch hx_carrier := hx_value.(type) {
					case *Date:
						if hx_carrier == nil {
							return false
						}
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
			case "app.core.Incident":
				return func(hx_value any) bool {
					switch hx_carrier := hx_value.(type) {
					case *app__core__Incident:
						if hx_carrier == nil {
							return false
						}
						return true
					default:
						return false
					}
				}(hx_value)
			case "app.core.IncidentApi":
				return func(hx_value any) bool {
					switch hx_carrier := hx_value.(type) {
					case *app__core__IncidentApi:
						if hx_carrier == nil {
							return false
						}
						return true
					default:
						return false
					}
				}(hx_value)
			case "app.core.IncidentConfig":
				return func(hx_value any) bool {
					switch hx_carrier := hx_value.(type) {
					case *app__core__IncidentConfig:
						if hx_carrier == nil {
							return false
						}
						return true
					default:
						return false
					}
				}(hx_value)
			case "app.core.IncidentRequestException":
				return func(hx_value any) bool {
					switch hx_carrier := hx_value.(type) {
					case *app__core__IncidentRequestException:
						if hx_carrier == nil {
							return false
						}
						return true
					default:
						return false
					}
				}(hx_value)
			case "app.core.IncidentStore":
				return func(hx_value any) bool {
					switch hx_carrier := hx_value.(type) {
					case *app__core__IncidentStore:
						if hx_carrier == nil {
							return false
						}
						return true
					default:
						return false
					}
				}(hx_value)
			case "app.http.HttpRequest":
				return func(hx_value any) bool {
					switch hx_carrier := hx_value.(type) {
					case *app__http__HttpRequest:
						if hx_carrier == nil {
							return false
						}
						return true
					default:
						return false
					}
				}(hx_value)
			case "app.http.HttpResponse":
				return func(hx_value any) bool {
					switch hx_carrier := hx_value.(type) {
					case *app__http__HttpResponse:
						if hx_carrier == nil {
							return false
						}
						return true
					default:
						return false
					}
				}(hx_value)
			case "app.http.TinyHttpServer":
				return func(hx_value any) bool {
					switch hx_carrier := hx_value.(type) {
					case *app__http__TinyHttpServer:
						if hx_carrier == nil {
							return false
						}
						return true
					default:
						return false
					}
				}(hx_value)
			case "haxe._Int64.___Int64":
				return func(hx_value any) bool {
					switch hx_carrier := hx_value.(type) {
					case *haxe___Int64_____Int64:
						if hx_carrier == nil {
							return false
						}
						return true
					default:
						return false
					}
				}(hx_value)
			case "haxe.exceptions.NotImplementedException":
				return func(hx_value any) bool {
					switch hx_carrier := hx_value.(type) {
					case *haxe__exceptions__NotImplementedException:
						if hx_carrier == nil {
							return false
						}
						return true
					case *haxe__exceptions__PosException:
						if hx_carrier == nil {
							return false
						}
						_, hx_ok := hx_carrier.__hx_this.(*haxe__exceptions__NotImplementedException)
						return hx_ok
					default:
						return false
					}
				}(hx_value)
			case "haxe.exceptions.PosException":
				return func(hx_value any) bool {
					switch hx_carrier := hx_value.(type) {
					case *haxe__exceptions__NotImplementedException:
						if hx_carrier == nil {
							return false
						}
						return true
					case *haxe__exceptions__PosException:
						if hx_carrier == nil {
							return false
						}
						return true
					default:
						return false
					}
				}(hx_value)
			case "haxe.io.Bytes":
				return func(hx_value any) bool {
					switch hx_carrier := hx_value.(type) {
					case *haxe__io__Bytes:
						if hx_carrier == nil {
							return false
						}
						return true
					default:
						return false
					}
				}(hx_value)
			case "haxe.io.BytesBuffer":
				return func(hx_value any) bool {
					switch hx_carrier := hx_value.(type) {
					case *haxe__io__BytesBuffer:
						if hx_carrier == nil {
							return false
						}
						return true
					default:
						return false
					}
				}(hx_value)
			case "haxe.io.Eof":
				return func(hx_value any) bool {
					switch hx_carrier := hx_value.(type) {
					case *haxe__io__Eof:
						if hx_carrier == nil {
							return false
						}
						return true
					default:
						return false
					}
				}(hx_value)
			case "haxe.io.Input":
				return func(hx_value any) bool {
					switch hx_carrier := hx_value.(type) {
					case *haxe__io__Input:
						if hx_carrier == nil {
							return false
						}
						return true
					case *sys__io__FileInput:
						if hx_carrier == nil {
							return false
						}
						return true
					case *sys__net__SocketInput:
						if hx_carrier == nil {
							return false
						}
						return true
					default:
						return false
					}
				}(hx_value)
			case "haxe.io.Output":
				return func(hx_value any) bool {
					switch hx_carrier := hx_value.(type) {
					case *haxe__io__Output:
						if hx_carrier == nil {
							return false
						}
						return true
					case *sys__io__FileOutput:
						if hx_carrier == nil {
							return false
						}
						return true
					case *sys__net__SocketOutput:
						if hx_carrier == nil {
							return false
						}
						return true
					default:
						return false
					}
				}(hx_value)
			case "haxe.iterators.StringIterator":
				return func(hx_value any) bool {
					switch hx_carrier := hx_value.(type) {
					case *haxe__iterators__StringIterator:
						if hx_carrier == nil {
							return false
						}
						return true
					default:
						return false
					}
				}(hx_value)
			case "haxe.iterators.StringKeyValueIterator":
				return func(hx_value any) bool {
					switch hx_carrier := hx_value.(type) {
					case *haxe__iterators__StringKeyValueIterator:
						if hx_carrier == nil {
							return false
						}
						return true
					default:
						return false
					}
				}(hx_value)
			case "sys.io.FileInput":
				return func(hx_value any) bool {
					switch hx_carrier := hx_value.(type) {
					case *haxe__io__Input:
						if hx_carrier == nil {
							return false
						}
						_, hx_ok := hx_carrier.__hx_this.(*sys__io__FileInput)
						return hx_ok
					case *sys__io__FileInput:
						if hx_carrier == nil {
							return false
						}
						return true
					default:
						return false
					}
				}(hx_value)
			case "sys.io.FileOutput":
				return func(hx_value any) bool {
					switch hx_carrier := hx_value.(type) {
					case *haxe__io__Output:
						if hx_carrier == nil {
							return false
						}
						_, hx_ok := hx_carrier.__hx_this.(*sys__io__FileOutput)
						return hx_ok
					case *sys__io__FileOutput:
						if hx_carrier == nil {
							return false
						}
						return true
					default:
						return false
					}
				}(hx_value)
			case "sys.net.Host":
				return func(hx_value any) bool {
					switch hx_carrier := hx_value.(type) {
					case *sys__net__Host:
						if hx_carrier == nil {
							return false
						}
						return true
					default:
						return false
					}
				}(hx_value)
			case "sys.net.Socket":
				return func(hx_value any) bool {
					switch hx_carrier := hx_value.(type) {
					case *sys__net__Socket:
						if hx_carrier == nil {
							return false
						}
						return true
					default:
						return false
					}
				}(hx_value)
			case "sys.net.SocketInput":
				return func(hx_value any) bool {
					switch hx_carrier := hx_value.(type) {
					case *haxe__io__Input:
						if hx_carrier == nil {
							return false
						}
						_, hx_ok := hx_carrier.__hx_this.(*sys__net__SocketInput)
						return hx_ok
					case *sys__net__SocketInput:
						if hx_carrier == nil {
							return false
						}
						return true
					default:
						return false
					}
				}(hx_value)
			case "sys.net.SocketOutput":
				return func(hx_value any) bool {
					switch hx_carrier := hx_value.(type) {
					case *haxe__io__Output:
						if hx_carrier == nil {
							return false
						}
						_, hx_ok := hx_carrier.__hx_this.(*sys__net__SocketOutput)
						return hx_ok
					case *sys__net__SocketOutput:
						if hx_carrier == nil {
							return false
						}
						return true
					default:
						return false
					}
				}(hx_value)
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
			default:
				return false
			}
		default:
			return false
		}
	}(v, t)
}

func Std_isDecimalDigit(code int) bool {
	return ((code >= 48) && (code <= 57))
}

func Std_isSpaceCode(code int) bool {
	return (((code >= 9) && (code <= 13)) || (code == 32))
}

func Std_parseFloat(x *string) float64 {
	if hxrt.StringEqualStringPtr(x, nil) {
		return hxrt.StringParseFloatExact(hxrt.StringFromLiteral(""))
	}
	index := 0
	length := hxrt.StringLengthStringPtr(x)
	for (index < length) && func() bool {
		code := hxrt.IntFromNullableAny(hxrt.StringCharCodeAtAnyStringPtr(x, index))
		return (((code >= 9) && (code <= 13)) || (code == 32))
	}() {
		index = int(int32((index + 1)))
	}
	start := index
	if index < length {
		var sign any = hxrt.StringCharCodeAtAnyStringPtr(x, index)
		if (sign == 45) || (sign == 43) {
			index = int(int32((index + 1)))
		}
	}
	digits := 0
	for (index < length) && func() bool {
		code_1 := hxrt.IntFromNullableAny(hxrt.StringCharCodeAtAnyStringPtr(x, index))
		return ((code_1 >= 48) && (code_1 <= 57))
	}() {
		digits = int(int32((digits + 1)))
		index = int(int32((index + 1)))
	}
	if (index < length) && (hxrt.StringCharCodeAtAnyStringPtr(x, index) == 46) {
		index = int(int32((index + 1)))
		for (index < length) && func() bool {
			code_2 := hxrt.IntFromNullableAny(hxrt.StringCharCodeAtAnyStringPtr(x, index))
			return ((code_2 >= 48) && (code_2 <= 57))
		}() {
			digits = int(int32((digits + 1)))
			index = int(int32((index + 1)))
		}
	}
	if digits == 0 {
		return hxrt.StringParseFloatExact(hxrt.StringFromLiteral(""))
	}
	if (index < length) && ((hxrt.StringCharCodeAtAnyStringPtr(x, index) == 101) || (hxrt.StringCharCodeAtAnyStringPtr(x, index) == 69)) {
		index = int(int32((index + 1)))
		if (index < length) && ((hxrt.StringCharCodeAtAnyStringPtr(x, index) == 45) || (hxrt.StringCharCodeAtAnyStringPtr(x, index) == 43)) {
			index = int(int32((index + 1)))
		}
		exponentStart := index
		for (index < length) && func() bool {
			code_3 := hxrt.IntFromNullableAny(hxrt.StringCharCodeAtAnyStringPtr(x, index))
			return ((code_3 >= 48) && (code_3 <= 57))
		}() {
			index = int(int32((index + 1)))
		}
		if index == exponentStart {
			return hxrt.StringParseFloatExact(hxrt.StringFromLiteral(""))
		}
	}
	return hxrt.StringParseFloatExact(hxrt.StringSubstringStringPtr(x, start, index))
}

func Std_parseInt(x *string) any {
	if hxrt.StringEqualStringPtr(x, nil) {
		return nil
	}
	index := 0
	length := hxrt.StringLengthStringPtr(x)
	for (index < length) && func() bool {
		code := hxrt.IntFromNullableAny(hxrt.StringCharCodeAtAnyStringPtr(x, index))
		return (((code >= 9) && (code <= 13)) || (code == 32))
	}() {
		index = int(int32((index + 1)))
	}
	negative := false
	if index < length {
		var sign any = hxrt.StringCharCodeAtAnyStringPtr(x, index)
		if (sign == 45) || (sign == 43) {
			negative = (sign == 45)
			index = int(int32((index + 1)))
		}
	}
	hexadecimal := (((int((hxrt.Int32Wrap(index) + hxrt.Int32Wrap(1))) < length) && (hxrt.StringCharCodeAtAnyStringPtr(x, index) == 48)) && ((hxrt.StringCharCodeAtAnyStringPtr(x, int((hxrt.Int32Wrap(index)+hxrt.Int32Wrap(1)))) == 120) || (hxrt.StringCharCodeAtAnyStringPtr(x, int((hxrt.Int32Wrap(index)+hxrt.Int32Wrap(1)))) == 88)))
	if hexadecimal {
		index = int((hxrt.Int32Wrap(index) + hxrt.Int32Wrap(2)))
	}
	digitStart := index
	if hexadecimal {
		value := 0
		significantDigits := 0
		sawNonZero := false
		for index < length {
			code_1 := hxrt.IntFromNullableAny(hxrt.StringCharCodeAtAnyStringPtr(x, index))
			var hx_if_5 int
			if (code_1 >= 48) && (code_1 <= 57) {
				hx_if_5 = int((hxrt.Int32Wrap(code_1) - hxrt.Int32Wrap(48)))
			} else {
				var hx_if_4 int
				if (code_1 >= 97) && (code_1 <= 102) {
					hx_if_4 = int((hxrt.Int32Wrap(int((hxrt.Int32Wrap(code_1) - hxrt.Int32Wrap(97)))) + hxrt.Int32Wrap(10)))
				} else {
					var hx_if_3 int
					if (code_1 >= 65) && (code_1 <= 70) {
						hx_if_3 = int((hxrt.Int32Wrap(int((hxrt.Int32Wrap(code_1) - hxrt.Int32Wrap(65)))) + hxrt.Int32Wrap(10)))
					} else {
						hx_if_3 = -1
					}
					hx_if_4 = hx_if_3
				}
				hx_if_5 = hx_if_4
			}
			digit := hx_if_5
			if digit < 0 {
				break
			}
			if (digit != 0) || sawNonZero {
				sawNonZero = true
				significantDigits = int(int32((significantDigits + 1)))
				if significantDigits > 8 {
					return nil
				}
			}
			value = int((hxrt.Int32Wrap(int((hxrt.Int32Wrap(value) << uint(4)))) | hxrt.Int32Wrap(digit)))
			index = int(int32((index + 1)))
		}
		if index == digitStart {
			return nil
		}
		var hx_if_6 any
		if negative {
			hx_if_6 = int(-int32(value))
		} else {
			hx_if_6 = value
		}
		return hx_if_6
	}
	result := 0
	for index < length {
		code_2 := hxrt.IntFromNullableAny(hxrt.StringCharCodeAtAnyStringPtr(x, index))
		var hx_if_7 int
		if (code_2 >= 48) && (code_2 <= 57) {
			hx_if_7 = int((hxrt.Int32Wrap(code_2) - hxrt.Int32Wrap(48)))
		} else {
			hx_if_7 = -1
		}
		digit_1 := hx_if_7
		if digit_1 < 0 {
			break
		}
		var hx_if_8 int
		if negative {
			hx_if_8 = 8
		} else {
			hx_if_8 = 7
		}
		lastAllowedDigit := hx_if_8
		if (result < -214748364) || ((result == -214748364) && (digit_1 > lastAllowedDigit)) {
			return nil
		}
		result = int((hxrt.Int32Wrap(int((hxrt.Int32Wrap(result) * hxrt.Int32Wrap(10)))) - hxrt.Int32Wrap(digit_1)))
		index = int(int32((index + 1)))
	}
	if index == digitStart {
		return nil
	}
	var hx_if_9 any
	if negative {
		hx_if_9 = result
	} else {
		hx_if_9 = int(-int32(result))
	}
	return hx_if_9
}

func Std_random(x int) int {
	var hx_if_10 int
	if x <= 1 {
		hx_if_10 = 0
	} else {
		hx_if_10 = rand.Intn(x)
	}
	return hx_if_10
}
