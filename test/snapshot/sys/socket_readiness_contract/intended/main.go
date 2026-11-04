package main

import "snapshot/hxrt"

func isBlocked(operation *hxrt.Function[func()]) bool {
	hx_try_return_1 := false
	var hx_try_value_2 bool
	hxrt.TryCatch(func() {
		func() {
			var hx_callable_10 *hxrt.Function[func()] = operation
			hxrt.FunctionCall(hx_callable_10)()
		}()
		hx_try_value_2 = false
		hx_try_return_1 = true
		return
	}, func(hx_caught_3 any) {
		hx_exception_6 := haxe__Exception_caught(hx_caught_3)
		hx_payload_7 := haxe__Exception_catchValue(hx_caught_3)
		_ = hx_exception_6
		if func(hx_value any) bool {
			switch hx_value.(type) {
			case *haxe__io__Error:
				return true
			default:
				return false
			}
		}(hx_payload_7) {
			var error *haxe__io__Error = func(hx_value_8 any) *haxe__io__Error {
				if hx_value_8 == nil {
					var hx_zero_9 *haxe__io__Error
					return hx_zero_9
				}
				return hx_value_8.(*haxe__io__Error)
			}(hx_payload_7)
			var hx_if_5 bool
			if error.tag == 0 {
				hx_if_5 = true
			} else {
				hx_if_5 = func() bool {
					hxrt.Throw(haxe__Exception_thrown(error))
					var hx_throw_zero_4 bool
					return hx_throw_zero_4
				}()
			}
			hx_try_value_2 = hx_if_5
			hx_try_return_1 = true
			return
		} else {
			hxrt.Throw(hx_caught_3)
		}
	})
	if hx_try_return_1 {
		return hx_try_value_2
	}
	return false
}

func main() {
	host := New_sys__net__Host(hxrt.StringFromLiteral("127.0.0.1"))
	server := New_sys__net__Socket()
	client := New_sys__net__Socket()
	var accepted *sys__net__Socket = nil
	unconnected := New_sys__net__Socket()
	var failed any = nil
	hxrt.TryCatch(func() {
		server.__hx_this.bind(host, 0)
		server.__hx_this.listen(1)
		client.__hx_this.connect(host, func(hx_value_14 any) int {
			if hx_value_14 == nil {
				var hx_zero_15 int
				return hx_zero_15
			}
			return hx_value_14.(int)
		}(hxrt.AnonymousGet(server.__hx_this.host(), "port")))
		accepted = server.__hx_this.accept()
		client.custom = hxrt.StringFromLiteral("client")
		idle := sys__net__Socket_select_(hxrt.NewArray(client), hxrt.NewArray(), hxrt.NewArray(unconnected), 0.0)
		var v any = any(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("idle="), func(hx_value_16 any) *hxrt.Array {
			if hx_value_16 == nil {
				var hx_zero_17 *hxrt.Array
				return hx_zero_17
			}
			return hx_value_16.(*hxrt.Array)
		}(hxrt.AnonymousGet(idle, "read")).Len()), hxrt.StringFromLiteral(":")), func(hx_value_18 any) *hxrt.Array {
			if hx_value_18 == nil {
				var hx_zero_19 *hxrt.Array
				return hx_zero_19
			}
			return hx_value_18.(*hxrt.Array)
		}(hxrt.AnonymousGet(idle, "others")).Len()))
		hxrt.Println(v)
		client.__hx_this.setBlocking(false)
		var v_1 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("readBlocked="), hxrt.StdString(isBlocked(hxrt.NewFunction(func() {
			client.input.__hx_this.readByte()
		})))))
		hxrt.Println(v_1)
		client.output.__hx_this.writeByte(110)
		var v_2 any = any(hxrt.StringConcatAny(hxrt.StringFromLiteral("nonblockingWrite="), accepted.input.__hx_this.readByte()))
		hxrt.Println(v_2)
		client.__hx_this.setBlocking(true)
		server.__hx_this.setBlocking(false)
		var v_3 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("acceptBlocked="), hxrt.StdString(isBlocked(hxrt.NewFunction(func() {
			server.__hx_this.accept()
		})))))
		hxrt.Println(v_3)
		server.__hx_this.setBlocking(true)
		accepted.output.__hx_this.writeString(hxrt.StringFromLiteral("xy"), nil)
		accepted.output.__hx_this.flush()
		duplicate := sys__net__Socket_select_(hxrt.NewArray(client, client), hxrt.NewArray(), hxrt.NewArray(), 1.0)
		var v_4 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("duplicates="), func(hx_value_20 any) *hxrt.Array {
			if hx_value_20 == nil {
				var hx_zero_21 *hxrt.Array
				return hx_zero_21
			}
			return hx_value_20.(*hxrt.Array)
		}(hxrt.AnonymousGet(duplicate, "read")).Len()), hxrt.StringFromLiteral(":")), hxrt.StdString(hxrt.HaxeEqual(func(hx_value_22 any) *hxrt.Array {
			if hx_value_22 == nil {
				var hx_zero_23 *hxrt.Array
				return hx_zero_23
			}
			return hx_value_22.(*hxrt.Array)
		}(hxrt.AnonymousGet(duplicate, "read")).Get(0), client))), hxrt.StringFromLiteral(":")), hxrt.StdString(hx_recover_class_0(func(hx_value_24 any) *hxrt.Array {
			if hx_value_24 == nil {
				var hx_zero_25 *hxrt.Array
				return hx_zero_25
			}
			return hx_value_24.(*hxrt.Array)
		}(hxrt.AnonymousGet(duplicate, "read")).Get(1)).custom)))
		hxrt.Println(v_4)
		var v_5 any = any(hxrt.StringConcatAny(hxrt.StringFromLiteral("first="), client.input.__hx_this.readByte()))
		hxrt.Println(v_5)
		buffered := sys__net__Socket_select_(hxrt.NewArray(client), hxrt.NewArray(), hxrt.NewArray(), 0.0)
		var v_6 any = any(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("buffered="), func(hx_value_26 any) *hxrt.Array {
			if hx_value_26 == nil {
				var hx_zero_27 *hxrt.Array
				return hx_zero_27
			}
			return hx_value_26.(*hxrt.Array)
		}(hxrt.AnonymousGet(buffered, "read")).Len()), hxrt.StringFromLiteral(":")), client.input.__hx_this.readByte()))
		hxrt.Println(v_6)
		writable := sys__net__Socket_select_(hxrt.NewArray(), hxrt.NewArray(client), hxrt.NewArray(), 0.0)
		var v_7 any = any(hxrt.StringConcatAny(hxrt.StringFromLiteral("writable="), func(hx_value_28 any) *hxrt.Array {
			if hx_value_28 == nil {
				var hx_zero_29 *hxrt.Array
				return hx_zero_29
			}
			return hx_value_28.(*hxrt.Array)
		}(hxrt.AnonymousGet(writable, "write")).Len()))
		hxrt.Println(v_7)
	}, func(hx_caught_11 any) {
		hx_exception_12 := haxe__Exception_caught(hx_caught_11)
		hx_payload_13 := haxe__Exception_catchValue(hx_caught_11)
		_ = hx_exception_12
		var error any = hx_payload_13
		failed = error
	})
	safeClose(unconnected)
	safeClose(accepted)
	safeClose(client)
	safeClose(server)
	if !hxrt.AnyEqualsNull(failed) {
		hxrt.Throw(haxe__Exception_thrown(failed))
	}
}

func safeClose(socket *sys__net__Socket) {
	if socket == nil {
		return
	}
	hxrt.TryCatch(func() {
		socket.__hx_this.close()
	}, func(hx_caught_30 any) {
		hx_exception_31 := haxe__Exception_caught(hx_caught_30)
		hx_payload_32 := haxe__Exception_catchValue(hx_caught_30)
		_ = hx_exception_31
		var hx_tmp any = hx_payload_32
		_ = hx_tmp
	})
}

func hx_recover_class_0(value any) *sys__net__Socket {
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

func hx_recover_class_1(value any) *haxe__io__Eof {
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

func hx_recover_class_2(value any) *haxe__ValueException {
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

func hx_recover_class_3(value any) *haxe__Exception {
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
