package main

import "snapshot/hxrt"

func isBlocked(operation func()) bool {
	hx_try_return_1 := false
	var hx_try_value_2 bool
	hxrt.TryCatch(func() {
		operation()
		hx_try_value_2 = false
		hx_try_return_1 = true
		return
	}, func(hx_caught_3 any) {
		switch hx_typed_4 := hx_caught_3.(type) {
		case *haxe__io__Error:
			error := hx_typed_4
			var hx_if_6 bool
			if error.tag == 0 {
				hx_if_6 = true
			} else {
				hx_if_6 = func() bool {
					hxrt.Throw(error)
					var hx_throw_zero_5 bool
					return hx_throw_zero_5
				}()
			}
			hx_try_value_2 = hx_if_6
			hx_try_return_1 = true
			return
		default:
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
		client.__hx_this.connect(host, func(hx_obj_9 map[string]any) int {
			hx_field_10 := hx_obj_9["port"]
			if hx_field_10 == nil {
				var hx_zero_11 int
				return hx_zero_11
			}
			return hx_field_10.(int)
		}(server.__hx_this.host()))
		accepted = server.__hx_this.accept()
		client.custom = hxrt.StringFromLiteral("client")
		idle := sys__net__Socket_select_(hxrt.NewArray(client), hxrt.NewArray(), hxrt.NewArray(unconnected), 0.0)
		var v any = any(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("idle="), func(hx_obj_12 map[string]any) *hxrt.Array {
			hx_field_13 := hx_obj_12["read"]
			if hx_field_13 == nil {
				var hx_zero_14 *hxrt.Array
				return hx_zero_14
			}
			return hx_field_13.(*hxrt.Array)
		}(idle).Len()), hxrt.StringFromLiteral(":")), func(hx_obj_15 map[string]any) *hxrt.Array {
			hx_field_16 := hx_obj_15["others"]
			if hx_field_16 == nil {
				var hx_zero_17 *hxrt.Array
				return hx_zero_17
			}
			return hx_field_16.(*hxrt.Array)
		}(idle).Len()))
		hxrt.Println(v)
		client.__hx_this.setBlocking(false)
		var v_1 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("readBlocked="), hxrt.StdString(isBlocked(func() {
			client.input.__hx_this.readByte()
		}))))
		hxrt.Println(v_1)
		client.output.__hx_this.writeByte(110)
		var v_2 any = any(hxrt.StringConcatAny(hxrt.StringFromLiteral("nonblockingWrite="), accepted.input.__hx_this.readByte()))
		hxrt.Println(v_2)
		client.__hx_this.setBlocking(true)
		server.__hx_this.setBlocking(false)
		var v_3 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("acceptBlocked="), hxrt.StdString(isBlocked(func() {
			server.__hx_this.accept()
		}))))
		hxrt.Println(v_3)
		server.__hx_this.setBlocking(true)
		accepted.output.__hx_this.writeString(hxrt.StringFromLiteral("xy"), nil)
		accepted.output.__hx_this.flush()
		duplicate := sys__net__Socket_select_(hxrt.NewArray(client, client), hxrt.NewArray(), hxrt.NewArray(), 1.0)
		var v_4 any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("duplicates="), func(hx_obj_18 map[string]any) *hxrt.Array {
			hx_field_19 := hx_obj_18["read"]
			if hx_field_19 == nil {
				var hx_zero_20 *hxrt.Array
				return hx_zero_20
			}
			return hx_field_19.(*hxrt.Array)
		}(duplicate).Len()), hxrt.StringFromLiteral(":")), hxrt.StdString(hxrt.HaxeEqual(func(hx_obj_21 map[string]any) *hxrt.Array {
			hx_field_22 := hx_obj_21["read"]
			if hx_field_22 == nil {
				var hx_zero_23 *hxrt.Array
				return hx_zero_23
			}
			return hx_field_22.(*hxrt.Array)
		}(duplicate).Get(0), client))), hxrt.StringFromLiteral(":")), hxrt.StdString(func(hx_value_27 any) *sys__net__Socket {
			if hx_value_27 == nil {
				var hx_zero_28 *sys__net__Socket
				return hx_zero_28
			}
			return hx_value_27.(*sys__net__Socket)
		}(func(hx_obj_24 map[string]any) *hxrt.Array {
			hx_field_25 := hx_obj_24["read"]
			if hx_field_25 == nil {
				var hx_zero_26 *hxrt.Array
				return hx_zero_26
			}
			return hx_field_25.(*hxrt.Array)
		}(duplicate).Get(1)).custom)))
		hxrt.Println(v_4)
		var v_5 any = any(hxrt.StringConcatAny(hxrt.StringFromLiteral("first="), client.input.__hx_this.readByte()))
		hxrt.Println(v_5)
		buffered := sys__net__Socket_select_(hxrt.NewArray(client), hxrt.NewArray(), hxrt.NewArray(), 0.0)
		var v_6 any = any(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("buffered="), func(hx_obj_29 map[string]any) *hxrt.Array {
			hx_field_30 := hx_obj_29["read"]
			if hx_field_30 == nil {
				var hx_zero_31 *hxrt.Array
				return hx_zero_31
			}
			return hx_field_30.(*hxrt.Array)
		}(buffered).Len()), hxrt.StringFromLiteral(":")), client.input.__hx_this.readByte()))
		hxrt.Println(v_6)
		writable := sys__net__Socket_select_(hxrt.NewArray(), hxrt.NewArray(client), hxrt.NewArray(), 0.0)
		var v_7 any = any(hxrt.StringConcatAny(hxrt.StringFromLiteral("writable="), func(hx_obj_32 map[string]any) *hxrt.Array {
			hx_field_33 := hx_obj_32["write"]
			if hx_field_33 == nil {
				var hx_zero_34 *hxrt.Array
				return hx_zero_34
			}
			return hx_field_33.(*hxrt.Array)
		}(writable).Len()))
		hxrt.Println(v_7)
	}, func(hx_caught_7 any) {
		error := hx_caught_7
		failed = error
	})
	safeClose(unconnected)
	safeClose(accepted)
	safeClose(client)
	safeClose(server)
	if !hxrt.AnyEqualsNull(failed) {
		hxrt.Throw(failed)
	}
}

func safeClose(socket *sys__net__Socket) {
	if socket == nil {
		return
	}
	hxrt.TryCatch(func() {
		socket.__hx_this.close()
	}, func(hx_caught_35 any) {
		hx_tmp := hx_caught_35
		_ = hx_tmp
	})
}
