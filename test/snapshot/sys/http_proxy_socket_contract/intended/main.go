package main

import "snapshot/hxrt"

func main() {
	var v any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("proxy0="), sys__Http_hxrt_proxyDescriptor()))
	hxrt.Println(v)
	sys__Http_PROXY = func() hxrt.Anonymous {
		hx_obj_1 := map[string]any{}
		hx_obj_1["host"] = hxrt.StringFromLiteral("proxy.local")
		hx_obj_1["port"] = 3128
		hx_obj_2 := map[string]any{}
		hx_obj_2["user"] = hxrt.StringFromLiteral("scott")
		hx_obj_2["pass"] = hxrt.StringFromLiteral("tiger")
		hx_obj_1["auth"] = hxrt.AnonymousFromMap(hx_obj_2)
		return hxrt.AnonymousFromMap(hx_obj_1)
	}()
	var v_1 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("proxy1="), sys__Http_hxrt_proxyDescriptor()))
	hxrt.Println(v_1)
	sys__Http_PROXY = func() hxrt.Anonymous {
		hx_obj_3 := map[string]any{}
		hx_obj_3["host"] = hxrt.StringFromLiteral("proxy.local")
		hx_obj_3["port"] = 80
		hx_obj_4 := map[string]any{}
		hx_obj_4["user"] = hxrt.StringFromLiteral("scott")
		hx_obj_4["pass"] = nil
		hx_obj_3["auth"] = hxrt.AnonymousFromMap(hx_obj_4)
		return hxrt.AnonymousFromMap(hx_obj_3)
	}()
	var v_2 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("proxy2="), sys__Http_hxrt_proxyDescriptor()))
	hxrt.Println(v_2)
	sys__Http_PROXY = func() hxrt.Anonymous {
		hx_obj_5 := map[string]any{}
		hx_obj_5["host"] = hxrt.StringFromLiteral("proxy.local:9000")
		hx_obj_5["port"] = 3128
		hx_obj_5["auth"] = nil
		return hxrt.AnonymousFromMap(hx_obj_5)
	}()
	var v_3 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("proxy3="), sys__Http_hxrt_proxyDescriptor()))
	hxrt.Println(v_3)
	sys__Http_PROXY = func() hxrt.Anonymous {
		hx_obj_6 := map[string]any{}
		hx_obj_6["host"] = nil
		hx_obj_6["port"] = 3128
		hx_obj_6["auth"] = nil
		return hxrt.AnonymousFromMap(hx_obj_6)
	}()
	var v_4 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("proxy4="), sys__Http_hxrt_proxyDescriptor()))
	hxrt.Println(v_4)
	sys__Http_PROXY = nil
	http := New_sys__Http(hxrt.StringFromLiteral("data:text/plain,body"))
	http.__hx_this.setPostData(hxrt.StringFromLiteral("payload"))
	sink := New_haxe__io__BytesOutput()
	http.__hx_this.customRequest(false, func(hx_value *haxe__io__BytesOutput) *haxe__io__Output {
		if hx_value == nil {
			return nil
		}
		return hx_value.haxe__io__Output
	}(sink), New_sys__net__Socket(), hxrt.StringFromLiteral("pAtCh"))
	var v_5 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("methodSock="), sink.__hx_this.getBytes().__hx_this.toString()))
	hxrt.Println(v_5)
}

func hx_structural_recover(value any) hxrt.Anonymous {
	switch source := value.(type) {
	case *haxe___Int64_____Int64:
		return func(hx_structural_source_1 *haxe___Int64_____Int64) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_1, func(name string) hxrt.AnonymousField {
				switch name {
				case "high":
					return hxrt.AnonymousField{Value: hx_structural_source_1.high, Found: true}
				case "low":
					return hxrt.AnonymousField{Value: hx_structural_source_1.low, Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "high":
					hx_structural_source_1.high = func(hx_value_2 any) int {
						if hx_value_2 == nil {
							var hx_zero_3 int
							return hx_zero_3
						}
						return hx_value_2.(int)
					}(incoming)
					return true
				case "low":
					hx_structural_source_1.low = func(hx_value_4 any) int {
						if hx_value_4 == nil {
							var hx_zero_5 int
							return hx_zero_5
						}
						return hx_value_4.(int)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *sys__net__SocketOutput:
		return func(hx_structural_source_6 *sys__net__SocketOutput) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_6, func(name string) hxrt.AnonymousField {
				switch name {
				case "handle":
					return hxrt.AnonymousField{Value: hx_structural_source_6.handle, Found: true}
				case "writeByte":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_9 := hx_structural_source_6
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_9.__hx_this.writeByte(value)
						}, func() any {
							return hx_receiver_9.__hx_this
						}, "writeByte")
					}(), Found: true}
				case "writeBytes":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes, int, int) int] {
						hx_receiver_10 := hx_structural_source_6
						return hxrt.BoundFunction(func(bytes *haxe__io__Bytes, pos int, length int) int {
							return hx_receiver_10.__hx_this.writeBytes(bytes, pos, length)
						}, func() any {
							return hx_receiver_10.__hx_this
						}, "writeBytes")
					}(), Found: true}
				case "flush":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func()] {
						hx_receiver_11 := hx_structural_source_6
						return hxrt.BoundFunction(func() {
							hx_receiver_11.__hx_this.flush()
						}, func() any {
							return hx_receiver_11.__hx_this
						}, "flush")
					}(), Found: true}
				case "close":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func()] {
						hx_receiver_12 := hx_structural_source_6
						return hxrt.BoundFunction(func() {
							hx_receiver_12.__hx_this.close()
						}, func() any {
							return hx_receiver_12.__hx_this
						}, "close")
					}(), Found: true}
				case "bigEndian":
					return hxrt.AnonymousField{Value: hx_structural_source_6.bigEndian, Found: true}
				case "set_bigEndian":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(bool) bool] {
						hx_receiver_13 := hx_structural_source_6
						return hxrt.BoundFunction(func(value bool) bool {
							return hx_receiver_13.__hx_this.set_bigEndian(value)
						}, func() any {
							return hx_receiver_13.__hx_this
						}, "set_bigEndian")
					}(), Found: true}
				case "write":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes)] {
						hx_receiver_14 := hx_structural_source_6
						return hxrt.BoundFunction(func(bytes *haxe__io__Bytes) {
							hx_receiver_14.__hx_this.write(bytes)
						}, func() any {
							return hx_receiver_14.__hx_this
						}, "write")
					}(), Found: true}
				case "writeFullBytes":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes, int, int)] {
						hx_receiver_15 := hx_structural_source_6
						return hxrt.BoundFunction(func(bytes *haxe__io__Bytes, pos int, len_local int) {
							hx_receiver_15.__hx_this.writeFullBytes(bytes, pos, len_local)
						}, func() any {
							return hx_receiver_15.__hx_this
						}, "writeFullBytes")
					}(), Found: true}
				case "writeFloat":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(float64)] {
						hx_receiver_16 := hx_structural_source_6
						return hxrt.BoundFunction(func(value float64) {
							hx_receiver_16.__hx_this.writeFloat(value)
						}, func() any {
							return hx_receiver_16.__hx_this
						}, "writeFloat")
					}(), Found: true}
				case "writeDouble":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(float64)] {
						hx_receiver_17 := hx_structural_source_6
						return hxrt.BoundFunction(func(value float64) {
							hx_receiver_17.__hx_this.writeDouble(value)
						}, func() any {
							return hx_receiver_17.__hx_this
						}, "writeDouble")
					}(), Found: true}
				case "writeInt8":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_18 := hx_structural_source_6
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_18.__hx_this.writeInt8(value)
						}, func() any {
							return hx_receiver_18.__hx_this
						}, "writeInt8")
					}(), Found: true}
				case "writeInt16":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_19 := hx_structural_source_6
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_19.__hx_this.writeInt16(value)
						}, func() any {
							return hx_receiver_19.__hx_this
						}, "writeInt16")
					}(), Found: true}
				case "writeUInt16":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_20 := hx_structural_source_6
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_20.__hx_this.writeUInt16(value)
						}, func() any {
							return hx_receiver_20.__hx_this
						}, "writeUInt16")
					}(), Found: true}
				case "writeInt24":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_21 := hx_structural_source_6
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_21.__hx_this.writeInt24(value)
						}, func() any {
							return hx_receiver_21.__hx_this
						}, "writeInt24")
					}(), Found: true}
				case "writeUInt24":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_22 := hx_structural_source_6
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_22.__hx_this.writeUInt24(value)
						}, func() any {
							return hx_receiver_22.__hx_this
						}, "writeUInt24")
					}(), Found: true}
				case "writeInt32":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_23 := hx_structural_source_6
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_23.__hx_this.writeInt32(value)
						}, func() any {
							return hx_receiver_23.__hx_this
						}, "writeInt32")
					}(), Found: true}
				case "prepare":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_24 := hx_structural_source_6
						return hxrt.BoundFunction(func(nbytes int) {
							hx_receiver_24.__hx_this.prepare(nbytes)
						}, func() any {
							return hx_receiver_24.__hx_this
						}, "prepare")
					}(), Found: true}
				case "writeInput":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Input, any)] {
						hx_receiver_25 := hx_structural_source_6
						return hxrt.BoundFunction(func(input *haxe__io__Input, bufsize any) {
							hx_receiver_25.__hx_this.writeInput(input, bufsize)
						}, func() any {
							return hx_receiver_25.__hx_this
						}, "writeInput")
					}(), Found: true}
				case "writeString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string, *haxe__io__Encoding)] {
						hx_receiver_26 := hx_structural_source_6
						return hxrt.BoundFunction(func(value *string, encoding *haxe__io__Encoding) {
							hx_receiver_26.__hx_this.writeString(value, encoding)
						}, func() any {
							return hx_receiver_26.__hx_this
						}, "writeString")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "handle":
					hx_structural_source_6.handle = func(hx_value_7 any) *hxrt.SocketHandle {
						if hx_value_7 == nil {
							var hx_zero_8 *hxrt.SocketHandle
							return hx_zero_8
						}
						return hx_value_7.(*hxrt.SocketHandle)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *sys__net__SocketInput:
		return func(hx_structural_source_27 *sys__net__SocketInput) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_27, func(name string) hxrt.AnonymousField {
				switch name {
				case "handle":
					return hxrt.AnonymousField{Value: hx_structural_source_27.handle, Found: true}
				case "readByte":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_30 := hx_structural_source_27
						return hxrt.BoundFunction(func() int {
							return hx_receiver_30.__hx_this.readByte()
						}, func() any {
							return hx_receiver_30.__hx_this
						}, "readByte")
					}(), Found: true}
				case "readBytes":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes, int, int) int] {
						hx_receiver_31 := hx_structural_source_27
						return hxrt.BoundFunction(func(bytes *haxe__io__Bytes, pos int, length int) int {
							return hx_receiver_31.__hx_this.readBytes(bytes, pos, length)
						}, func() any {
							return hx_receiver_31.__hx_this
						}, "readBytes")
					}(), Found: true}
				case "close":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func()] {
						hx_receiver_32 := hx_structural_source_27
						return hxrt.BoundFunction(func() {
							hx_receiver_32.__hx_this.close()
						}, func() any {
							return hx_receiver_32.__hx_this
						}, "close")
					}(), Found: true}
				case "bigEndian":
					return hxrt.AnonymousField{Value: hx_structural_source_27.bigEndian, Found: true}
				case "set_bigEndian":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(bool) bool] {
						hx_receiver_33 := hx_structural_source_27
						return hxrt.BoundFunction(func(value bool) bool {
							return hx_receiver_33.__hx_this.set_bigEndian(value)
						}, func() any {
							return hx_receiver_33.__hx_this
						}, "set_bigEndian")
					}(), Found: true}
				case "readAll":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(any) *haxe__io__Bytes] {
						hx_receiver_34 := hx_structural_source_27
						return hxrt.BoundFunction(func(bufsize any) *haxe__io__Bytes {
							return hx_receiver_34.__hx_this.readAll(bufsize)
						}, func() any {
							return hx_receiver_34.__hx_this
						}, "readAll")
					}(), Found: true}
				case "readFullBytes":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes, int, int)] {
						hx_receiver_35 := hx_structural_source_27
						return hxrt.BoundFunction(func(bytes *haxe__io__Bytes, pos int, len_local int) {
							hx_receiver_35.__hx_this.readFullBytes(bytes, pos, len_local)
						}, func() any {
							return hx_receiver_35.__hx_this
						}, "readFullBytes")
					}(), Found: true}
				case "read":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int) *haxe__io__Bytes] {
						hx_receiver_36 := hx_structural_source_27
						return hxrt.BoundFunction(func(nbytes int) *haxe__io__Bytes {
							return hx_receiver_36.__hx_this.read(nbytes)
						}, func() any {
							return hx_receiver_36.__hx_this
						}, "read")
					}(), Found: true}
				case "readUntil":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int) *string] {
						hx_receiver_37 := hx_structural_source_27
						return hxrt.BoundFunction(func(end int) *string {
							return hx_receiver_37.__hx_this.readUntil(end)
						}, func() any {
							return hx_receiver_37.__hx_this
						}, "readUntil")
					}(), Found: true}
				case "readLine":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_38 := hx_structural_source_27
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_38.__hx_this.readLine()
						}, func() any {
							return hx_receiver_38.__hx_this
						}, "readLine")
					}(), Found: true}
				case "readFloat":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() float64] {
						hx_receiver_39 := hx_structural_source_27
						return hxrt.BoundFunction(func() float64 {
							return hx_receiver_39.__hx_this.readFloat()
						}, func() any {
							return hx_receiver_39.__hx_this
						}, "readFloat")
					}(), Found: true}
				case "readDouble":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() float64] {
						hx_receiver_40 := hx_structural_source_27
						return hxrt.BoundFunction(func() float64 {
							return hx_receiver_40.__hx_this.readDouble()
						}, func() any {
							return hx_receiver_40.__hx_this
						}, "readDouble")
					}(), Found: true}
				case "readInt8":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_41 := hx_structural_source_27
						return hxrt.BoundFunction(func() int {
							return hx_receiver_41.__hx_this.readInt8()
						}, func() any {
							return hx_receiver_41.__hx_this
						}, "readInt8")
					}(), Found: true}
				case "readInt16":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_42 := hx_structural_source_27
						return hxrt.BoundFunction(func() int {
							return hx_receiver_42.__hx_this.readInt16()
						}, func() any {
							return hx_receiver_42.__hx_this
						}, "readInt16")
					}(), Found: true}
				case "readUInt16":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_43 := hx_structural_source_27
						return hxrt.BoundFunction(func() int {
							return hx_receiver_43.__hx_this.readUInt16()
						}, func() any {
							return hx_receiver_43.__hx_this
						}, "readUInt16")
					}(), Found: true}
				case "readInt24":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_44 := hx_structural_source_27
						return hxrt.BoundFunction(func() int {
							return hx_receiver_44.__hx_this.readInt24()
						}, func() any {
							return hx_receiver_44.__hx_this
						}, "readInt24")
					}(), Found: true}
				case "readUInt24":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_45 := hx_structural_source_27
						return hxrt.BoundFunction(func() int {
							return hx_receiver_45.__hx_this.readUInt24()
						}, func() any {
							return hx_receiver_45.__hx_this
						}, "readUInt24")
					}(), Found: true}
				case "readInt32":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_46 := hx_structural_source_27
						return hxrt.BoundFunction(func() int {
							return hx_receiver_46.__hx_this.readInt32()
						}, func() any {
							return hx_receiver_46.__hx_this
						}, "readInt32")
					}(), Found: true}
				case "readString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, *haxe__io__Encoding) *string] {
						hx_receiver_47 := hx_structural_source_27
						return hxrt.BoundFunction(func(len_local int, encoding *haxe__io__Encoding) *string {
							return hx_receiver_47.__hx_this.readString(len_local, encoding)
						}, func() any {
							return hx_receiver_47.__hx_this
						}, "readString")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "handle":
					hx_structural_source_27.handle = func(hx_value_28 any) *hxrt.SocketHandle {
						if hx_value_28 == nil {
							var hx_zero_29 *hxrt.SocketHandle
							return hx_zero_29
						}
						return hx_value_28.(*hxrt.SocketHandle)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *sys__net__Socket:
		return func(hx_structural_source_48 *sys__net__Socket) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_48, func(name string) hxrt.AnonymousField {
				switch name {
				case "input":
					return hxrt.AnonymousField{Value: hx_structural_source_48.input, Found: true}
				case "output":
					return hxrt.AnonymousField{Value: hx_structural_source_48.output, Found: true}
				case "custom":
					return hxrt.AnonymousField{Value: hx_structural_source_48.custom, Found: true}
				case "handle":
					return hxrt.AnonymousField{Value: hx_structural_source_48.handle, Found: true}
				case "close":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func()] {
						hx_receiver_51 := hx_structural_source_48
						return hxrt.BoundFunction(func() {
							hx_receiver_51.__hx_this.close()
						}, func() any {
							return hx_receiver_51.__hx_this
						}, "close")
					}(), Found: true}
				case "read":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_52 := hx_structural_source_48
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_52.__hx_this.read()
						}, func() any {
							return hx_receiver_52.__hx_this
						}, "read")
					}(), Found: true}
				case "write":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string)] {
						hx_receiver_53 := hx_structural_source_48
						return hxrt.BoundFunction(func(content *string) {
							hx_receiver_53.__hx_this.write(content)
						}, func() any {
							return hx_receiver_53.__hx_this
						}, "write")
					}(), Found: true}
				case "connect":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*sys__net__Host, int)] {
						hx_receiver_54 := hx_structural_source_48
						return hxrt.BoundFunction(func(host *sys__net__Host, port int) {
							hx_receiver_54.__hx_this.connect(host, port)
						}, func() any {
							return hx_receiver_54.__hx_this
						}, "connect")
					}(), Found: true}
				case "listen":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_55 := hx_structural_source_48
						return hxrt.BoundFunction(func(connections int) {
							hx_receiver_55.__hx_this.listen(connections)
						}, func() any {
							return hx_receiver_55.__hx_this
						}, "listen")
					}(), Found: true}
				case "shutdown":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(bool, bool)] {
						hx_receiver_56 := hx_structural_source_48
						return hxrt.BoundFunction(func(read bool, write bool) {
							hx_receiver_56.__hx_this.shutdown(read, write)
						}, func() any {
							return hx_receiver_56.__hx_this
						}, "shutdown")
					}(), Found: true}
				case "bind":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*sys__net__Host, int)] {
						hx_receiver_57 := hx_structural_source_48
						return hxrt.BoundFunction(func(host *sys__net__Host, port int) {
							hx_receiver_57.__hx_this.bind(host, port)
						}, func() any {
							return hx_receiver_57.__hx_this
						}, "bind")
					}(), Found: true}
				case "accept":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *sys__net__Socket] {
						hx_receiver_58 := hx_structural_source_48
						return hxrt.BoundFunction(func() *sys__net__Socket {
							return hx_receiver_58.__hx_this.accept()
						}, func() any {
							return hx_receiver_58.__hx_this
						}, "accept")
					}(), Found: true}
				case "peer":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() hxrt.Anonymous] {
						hx_receiver_59 := hx_structural_source_48
						return hxrt.BoundFunction(func() hxrt.Anonymous {
							return hx_receiver_59.__hx_this.peer()
						}, func() any {
							return hx_receiver_59.__hx_this
						}, "peer")
					}(), Found: true}
				case "host":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() hxrt.Anonymous] {
						hx_receiver_60 := hx_structural_source_48
						return hxrt.BoundFunction(func() hxrt.Anonymous {
							return hx_receiver_60.__hx_this.host()
						}, func() any {
							return hx_receiver_60.__hx_this
						}, "host")
					}(), Found: true}
				case "setTimeout":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(float64)] {
						hx_receiver_61 := hx_structural_source_48
						return hxrt.BoundFunction(func(timeout float64) {
							hx_receiver_61.__hx_this.setTimeout(timeout)
						}, func() any {
							return hx_receiver_61.__hx_this
						}, "setTimeout")
					}(), Found: true}
				case "waitForRead":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func()] {
						hx_receiver_62 := hx_structural_source_48
						return hxrt.BoundFunction(func() {
							hx_receiver_62.__hx_this.waitForRead()
						}, func() any {
							return hx_receiver_62.__hx_this
						}, "waitForRead")
					}(), Found: true}
				case "setBlocking":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(bool)] {
						hx_receiver_63 := hx_structural_source_48
						return hxrt.BoundFunction(func(blocking bool) {
							hx_receiver_63.__hx_this.setBlocking(blocking)
						}, func() any {
							return hx_receiver_63.__hx_this
						}, "setBlocking")
					}(), Found: true}
				case "setFastSend":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(bool)] {
						hx_receiver_64 := hx_structural_source_48
						return hxrt.BoundFunction(func(fastSend bool) {
							hx_receiver_64.__hx_this.setFastSend(fastSend)
						}, func() any {
							return hx_receiver_64.__hx_this
						}, "setFastSend")
					}(), Found: true}
				case "replaceHandle":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*hxrt.SocketHandle)] {
						hx_receiver_65 := hx_structural_source_48
						return hxrt.BoundFunction(func(next *hxrt.SocketHandle) {
							hx_receiver_65.__hx_this.replaceHandle(next)
						}, func() any {
							return hx_receiver_65.__hx_this
						}, "replaceHandle")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "custom":
					hx_structural_source_48.custom = incoming
					return true
				case "handle":
					hx_structural_source_48.handle = func(hx_value_49 any) *hxrt.SocketHandle {
						if hx_value_49 == nil {
							var hx_zero_50 *hxrt.SocketHandle
							return hx_zero_50
						}
						return hx_value_49.(*hxrt.SocketHandle)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *sys__net__Host:
		return func(hx_structural_source_66 *sys__net__Host) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_66, func(name string) hxrt.AnonymousField {
				switch name {
				case "host":
					return hxrt.AnonymousField{Value: hx_structural_source_66.host, Found: true}
				case "ip":
					return hxrt.AnonymousField{Value: hx_structural_source_66.ip, Found: true}
				case "toString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_67 := hx_structural_source_66
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_67.__hx_this.toString()
						}, func() any {
							return hx_receiver_67.__hx_this
						}, "toString")
					}(), Found: true}
				case "reverse":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_68 := hx_structural_source_66
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_68.__hx_this.reverse()
						}, func() any {
							return hx_receiver_68.__hx_this
						}, "reverse")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				}
				return false
			})
		}(source)
	case *sys__Http:
		return func(hx_structural_source_69 *sys__Http) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_69, func(name string) hxrt.AnonymousField {
				switch name {
				case "noShutdown":
					return hxrt.AnonymousField{Value: hx_structural_source_69.noShutdown, Found: true}
				case "cnxTimeout":
					return hxrt.AnonymousField{Value: hx_structural_source_69.cnxTimeout, Found: true}
				case "responseHeaders":
					return hxrt.AnonymousField{Value: hx_structural_source_69.responseHeaders, Found: true}
				case "responseHeadersSameKey":
					return hxrt.AnonymousField{Value: hx_structural_source_69.responseHeadersSameKey, Found: true}
				case "file":
					return hxrt.AnonymousField{Value: hx_structural_source_69.file, Found: true}
				case "request":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(any)] {
						hx_receiver_74 := hx_structural_source_69
						return hxrt.BoundFunction(func(post any) {
							hx_receiver_74.__hx_this.request(post)
						}, func() any {
							return hx_receiver_74.__hx_this
						}, "request")
					}(), Found: true}
				case "fileTransfert":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string, *string, *haxe__io__Input, int, *string)] {
						hx_receiver_75 := hx_structural_source_69
						return hxrt.BoundFunction(func(argname *string, filename *string, file *haxe__io__Input, size int, mimeType *string) {
							if mimeType == nil {
								mimeType = hxrt.StringFromLiteral("application/octet-stream")
							}
							hx_receiver_75.__hx_this.fileTransfert(argname, filename, file, size, mimeType)
						}, func() any {
							return hx_receiver_75.__hx_this
						}, "fileTransfert")
					}(), Found: true}
				case "fileTransfer":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string, *string, *haxe__io__Input, int, *string)] {
						hx_receiver_76 := hx_structural_source_69
						return hxrt.BoundFunction(func(argname *string, filename *string, file *haxe__io__Input, size int, mimeType *string) {
							if mimeType == nil {
								mimeType = hxrt.StringFromLiteral("application/octet-stream")
							}
							hx_receiver_76.__hx_this.fileTransfer(argname, filename, file, size, mimeType)
						}, func() any {
							return hx_receiver_76.__hx_this
						}, "fileTransfer")
					}(), Found: true}
				case "customRequest":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(bool, *haxe__io__Output, *sys__net__Socket, *string)] {
						hx_receiver_77 := hx_structural_source_69
						return hxrt.BoundFunction(func(post bool, api *haxe__io__Output, sock *sys__net__Socket, method *string) {
							hx_receiver_77.__hx_this.customRequest(post, api, sock, method)
						}, func() any {
							return hx_receiver_77.__hx_this
						}, "customRequest")
					}(), Found: true}
				case "getResponseHeaderValues":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string) *hxrt.Array] {
						hx_receiver_78 := hx_structural_source_69
						return hxrt.BoundFunction(func(key *string) *hxrt.Array {
							return hx_receiver_78.__hx_this.getResponseHeaderValues(key)
						}, func() any {
							return hx_receiver_78.__hx_this
						}, "getResponseHeaderValues")
					}(), Found: true}
				case "requestWith":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(bool, *haxe__io__Output, *sys__net__Socket, *string)] {
						hx_receiver_79 := hx_structural_source_69
						return hxrt.BoundFunction(func(post bool, api *haxe__io__Output, sock *sys__net__Socket, method *string) {
							hx_receiver_79.__hx_this.requestWith(post, api, sock, method)
						}, func() any {
							return hx_receiver_79.__hx_this
						}, "requestWith")
					}(), Found: true}
				case "pumpUpload":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*hxrt.HttpExchange, hxrt.Anonymous) hxrt.Anonymous] {
						hx_receiver_80 := hx_structural_source_69
						return hxrt.BoundFunction(func(exchange *hxrt.HttpExchange, upload hxrt.Anonymous) hxrt.Anonymous {
							return hx_receiver_80.__hx_this.pumpUpload(exchange, upload)
						}, func() any {
							return hx_receiver_80.__hx_this
						}, "pumpUpload")
					}(), Found: true}
				case "handleDataRequest":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(bool, *haxe__io__Output, *string)] {
						hx_receiver_81 := hx_structural_source_69
						return hxrt.BoundFunction(func(post bool, api *haxe__io__Output, method *string) {
							hx_receiver_81.__hx_this.handleDataRequest(post, api, method)
						}, func() any {
							return hx_receiver_81.__hx_this
						}, "handleDataRequest")
					}(), Found: true}
				case "recordResponseHeaders":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*hxrt.HttpExchange)] {
						hx_receiver_82 := hx_structural_source_69
						return hxrt.BoundFunction(func(exchange *hxrt.HttpExchange) {
							hx_receiver_82.__hx_this.recordResponseHeaders(exchange)
						}, func() any {
							return hx_receiver_82.__hx_this
						}, "recordResponseHeaders")
					}(), Found: true}
				case "resetResponseHeaders":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func()] {
						hx_receiver_83 := hx_structural_source_69
						return hxrt.BoundFunction(func() {
							hx_receiver_83.__hx_this.resetResponseHeaders()
						}, func() any {
							return hx_receiver_83.__hx_this
						}, "resetResponseHeaders")
					}(), Found: true}
				case "encodedParameters":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_84 := hx_structural_source_69
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_84.__hx_this.encodedParameters()
						}, func() any {
							return hx_receiver_84.__hx_this
						}, "encodedParameters")
					}(), Found: true}
				case "url":
					return hxrt.AnonymousField{Value: hx_structural_source_69.url, Found: true}
				case "responseBytes":
					return hxrt.AnonymousField{Value: hx_structural_source_69.responseBytes, Found: true}
				case "responseAsString":
					return hxrt.AnonymousField{Value: hx_structural_source_69.responseAsString, Found: true}
				case "postData":
					return hxrt.AnonymousField{Value: hx_structural_source_69.postData, Found: true}
				case "postBytes":
					return hxrt.AnonymousField{Value: hx_structural_source_69.postBytes, Found: true}
				case "headers":
					return hxrt.AnonymousField{Value: hx_structural_source_69.headers, Found: true}
				case "params":
					return hxrt.AnonymousField{Value: hx_structural_source_69.params, Found: true}
				case "emptyOnData":
					return hxrt.AnonymousField{Value: hx_structural_source_69.emptyOnData, Found: true}
				case "setHeader":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string, *string)] {
						hx_receiver_95 := hx_structural_source_69
						return hxrt.BoundFunction(func(name *string, value *string) {
							hx_receiver_95.__hx_this.setHeader(name, value)
						}, func() any {
							return hx_receiver_95.__hx_this
						}, "setHeader")
					}(), Found: true}
				case "addHeader":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string, *string)] {
						hx_receiver_96 := hx_structural_source_69
						return hxrt.BoundFunction(func(header *string, value *string) {
							hx_receiver_96.__hx_this.addHeader(header, value)
						}, func() any {
							return hx_receiver_96.__hx_this
						}, "addHeader")
					}(), Found: true}
				case "setParameter":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string, *string)] {
						hx_receiver_97 := hx_structural_source_69
						return hxrt.BoundFunction(func(name *string, value *string) {
							hx_receiver_97.__hx_this.setParameter(name, value)
						}, func() any {
							return hx_receiver_97.__hx_this
						}, "setParameter")
					}(), Found: true}
				case "addParameter":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string, *string)] {
						hx_receiver_98 := hx_structural_source_69
						return hxrt.BoundFunction(func(name *string, value *string) {
							hx_receiver_98.__hx_this.addParameter(name, value)
						}, func() any {
							return hx_receiver_98.__hx_this
						}, "addParameter")
					}(), Found: true}
				case "setPostData":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string)] {
						hx_receiver_99 := hx_structural_source_69
						return hxrt.BoundFunction(func(data *string) {
							hx_receiver_99.__hx_this.setPostData(data)
						}, func() any {
							return hx_receiver_99.__hx_this
						}, "setPostData")
					}(), Found: true}
				case "setPostBytes":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes)] {
						hx_receiver_100 := hx_structural_source_69
						return hxrt.BoundFunction(func(data *haxe__io__Bytes) {
							hx_receiver_100.__hx_this.setPostBytes(data)
						}, func() any {
							return hx_receiver_100.__hx_this
						}, "setPostBytes")
					}(), Found: true}
				case "onData":
					return hxrt.AnonymousField{Value: hxrt.ReadDynamicFunction(hx_structural_source_69.onData), Found: true}
				case "onBytes":
					return hxrt.AnonymousField{Value: hxrt.ReadDynamicFunction(hx_structural_source_69.onBytes), Found: true}
				case "onError":
					return hxrt.AnonymousField{Value: hxrt.ReadDynamicFunction(hx_structural_source_69.onError), Found: true}
				case "onStatus":
					return hxrt.AnonymousField{Value: hxrt.ReadDynamicFunction(hx_structural_source_69.onStatus), Found: true}
				case "hasOnData":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() bool] {
						hx_receiver_101 := hx_structural_source_69
						return hxrt.BoundFunction(func() bool {
							return hx_receiver_101.__hx_this.hasOnData()
						}, func() any {
							return hx_receiver_101.__hx_this
						}, "hasOnData")
					}(), Found: true}
				case "success":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes)] {
						hx_receiver_102 := hx_structural_source_69
						return hxrt.BoundFunction(func(data *haxe__io__Bytes) {
							hx_receiver_102.__hx_this.success(data)
						}, func() any {
							return hx_receiver_102.__hx_this
						}, "success")
					}(), Found: true}
				case "get_responseData":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_103 := hx_structural_source_69
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_103.__hx_this.get_responseData()
						}, func() any {
							return hx_receiver_103.__hx_this
						}, "get_responseData")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "noShutdown":
					hx_structural_source_69.noShutdown = func(hx_value_70 any) bool {
						if hx_value_70 == nil {
							var hx_zero_71 bool
							return hx_zero_71
						}
						return hx_value_70.(bool)
					}(incoming)
					return true
				case "cnxTimeout":
					hx_structural_source_69.cnxTimeout = hxrt.DynamicToFloatSlot(incoming)
					return true
				case "responseHeaders":
					hx_structural_source_69.responseHeaders = func(hx_value_72 any) *haxe__ds__StringMap {
						if hx_value_72 == nil {
							var hx_zero_73 *haxe__ds__StringMap
							return hx_zero_73
						}
						return hx_value_72.(*haxe__ds__StringMap)
					}(incoming)
					return true
				case "responseHeadersSameKey":
					hx_structural_source_69.responseHeadersSameKey = hx_recover_class_haxe__ds__StringMap(incoming)
					return true
				case "file":
					hx_structural_source_69.file = hx_structural_recover(incoming)
					return true
				case "url":
					hx_structural_source_69.url = func(hx_value_85 any) *string {
						if hx_value_85 == nil {
							var hx_zero_86 *string
							return hx_zero_86
						}
						return hx_value_85.(*string)
					}(incoming)
					return true
				case "responseAsString":
					hx_structural_source_69.responseAsString = func(hx_value_87 any) *string {
						if hx_value_87 == nil {
							var hx_zero_88 *string
							return hx_zero_88
						}
						return hx_value_87.(*string)
					}(incoming)
					return true
				case "postData":
					hx_structural_source_69.postData = func(hx_value_89 any) *string {
						if hx_value_89 == nil {
							var hx_zero_90 *string
							return hx_zero_90
						}
						return hx_value_89.(*string)
					}(incoming)
					return true
				case "postBytes":
					hx_structural_source_69.postBytes = hx_recover_class_haxe__io__Bytes(incoming)
					return true
				case "headers":
					hx_structural_source_69.headers = func(hx_value_91 any) *hxrt.Array {
						if hx_value_91 == nil {
							var hx_zero_92 *hxrt.Array
							return hx_zero_92
						}
						return hx_value_91.(*hxrt.Array)
					}(incoming)
					return true
				case "params":
					hx_structural_source_69.params = func(hx_value_93 any) *hxrt.Array {
						if hx_value_93 == nil {
							var hx_zero_94 *hxrt.Array
							return hx_zero_94
						}
						return hx_value_93.(*hxrt.Array)
					}(incoming)
					return true
				case "onData":
					hx_structural_source_69.onData = hx_recover_callable_9da3ad96(incoming)
					return true
				case "onBytes":
					hx_structural_source_69.onBytes = hx_recover_callable_c5f1b802(incoming)
					return true
				case "onError":
					hx_structural_source_69.onError = hx_recover_callable_9da3ad96(incoming)
					return true
				case "onStatus":
					hx_structural_source_69.onStatus = hx_recover_callable_d8be58e9(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe__io__Output:
		return func(hx_structural_source_104 *haxe__io__Output) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_104, func(name string) hxrt.AnonymousField {
				switch name {
				case "bigEndian":
					return hxrt.AnonymousField{Value: hx_structural_source_104.bigEndian, Found: true}
				case "writeByte":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_105 := hx_structural_source_104
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_105.__hx_this.writeByte(value)
						}, func() any {
							return hx_receiver_105.__hx_this
						}, "writeByte")
					}(), Found: true}
				case "writeBytes":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes, int, int) int] {
						hx_receiver_106 := hx_structural_source_104
						return hxrt.BoundFunction(func(bytes *haxe__io__Bytes, pos int, len_local int) int {
							return hx_receiver_106.__hx_this.writeBytes(bytes, pos, len_local)
						}, func() any {
							return hx_receiver_106.__hx_this
						}, "writeBytes")
					}(), Found: true}
				case "flush":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func()] {
						hx_receiver_107 := hx_structural_source_104
						return hxrt.BoundFunction(func() {
							hx_receiver_107.__hx_this.flush()
						}, func() any {
							return hx_receiver_107.__hx_this
						}, "flush")
					}(), Found: true}
				case "close":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func()] {
						hx_receiver_108 := hx_structural_source_104
						return hxrt.BoundFunction(func() {
							hx_receiver_108.__hx_this.close()
						}, func() any {
							return hx_receiver_108.__hx_this
						}, "close")
					}(), Found: true}
				case "set_bigEndian":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(bool) bool] {
						hx_receiver_109 := hx_structural_source_104
						return hxrt.BoundFunction(func(value bool) bool {
							return hx_receiver_109.__hx_this.set_bigEndian(value)
						}, func() any {
							return hx_receiver_109.__hx_this
						}, "set_bigEndian")
					}(), Found: true}
				case "write":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes)] {
						hx_receiver_110 := hx_structural_source_104
						return hxrt.BoundFunction(func(bytes *haxe__io__Bytes) {
							hx_receiver_110.__hx_this.write(bytes)
						}, func() any {
							return hx_receiver_110.__hx_this
						}, "write")
					}(), Found: true}
				case "writeFullBytes":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes, int, int)] {
						hx_receiver_111 := hx_structural_source_104
						return hxrt.BoundFunction(func(bytes *haxe__io__Bytes, pos int, len_local int) {
							hx_receiver_111.__hx_this.writeFullBytes(bytes, pos, len_local)
						}, func() any {
							return hx_receiver_111.__hx_this
						}, "writeFullBytes")
					}(), Found: true}
				case "writeFloat":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(float64)] {
						hx_receiver_112 := hx_structural_source_104
						return hxrt.BoundFunction(func(value float64) {
							hx_receiver_112.__hx_this.writeFloat(value)
						}, func() any {
							return hx_receiver_112.__hx_this
						}, "writeFloat")
					}(), Found: true}
				case "writeDouble":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(float64)] {
						hx_receiver_113 := hx_structural_source_104
						return hxrt.BoundFunction(func(value float64) {
							hx_receiver_113.__hx_this.writeDouble(value)
						}, func() any {
							return hx_receiver_113.__hx_this
						}, "writeDouble")
					}(), Found: true}
				case "writeInt8":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_114 := hx_structural_source_104
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_114.__hx_this.writeInt8(value)
						}, func() any {
							return hx_receiver_114.__hx_this
						}, "writeInt8")
					}(), Found: true}
				case "writeInt16":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_115 := hx_structural_source_104
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_115.__hx_this.writeInt16(value)
						}, func() any {
							return hx_receiver_115.__hx_this
						}, "writeInt16")
					}(), Found: true}
				case "writeUInt16":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_116 := hx_structural_source_104
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_116.__hx_this.writeUInt16(value)
						}, func() any {
							return hx_receiver_116.__hx_this
						}, "writeUInt16")
					}(), Found: true}
				case "writeInt24":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_117 := hx_structural_source_104
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_117.__hx_this.writeInt24(value)
						}, func() any {
							return hx_receiver_117.__hx_this
						}, "writeInt24")
					}(), Found: true}
				case "writeUInt24":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_118 := hx_structural_source_104
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_118.__hx_this.writeUInt24(value)
						}, func() any {
							return hx_receiver_118.__hx_this
						}, "writeUInt24")
					}(), Found: true}
				case "writeInt32":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_119 := hx_structural_source_104
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_119.__hx_this.writeInt32(value)
						}, func() any {
							return hx_receiver_119.__hx_this
						}, "writeInt32")
					}(), Found: true}
				case "prepare":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_120 := hx_structural_source_104
						return hxrt.BoundFunction(func(nbytes int) {
							hx_receiver_120.__hx_this.prepare(nbytes)
						}, func() any {
							return hx_receiver_120.__hx_this
						}, "prepare")
					}(), Found: true}
				case "writeInput":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Input, any)] {
						hx_receiver_121 := hx_structural_source_104
						return hxrt.BoundFunction(func(input *haxe__io__Input, bufsize any) {
							hx_receiver_121.__hx_this.writeInput(input, bufsize)
						}, func() any {
							return hx_receiver_121.__hx_this
						}, "writeInput")
					}(), Found: true}
				case "writeString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string, *haxe__io__Encoding)] {
						hx_receiver_122 := hx_structural_source_104
						return hxrt.BoundFunction(func(value *string, encoding *haxe__io__Encoding) {
							hx_receiver_122.__hx_this.writeString(value, encoding)
						}, func() any {
							return hx_receiver_122.__hx_this
						}, "writeString")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				}
				return false
			})
		}(source)
	case *haxe__io__Input:
		return func(hx_structural_source_123 *haxe__io__Input) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_123, func(name string) hxrt.AnonymousField {
				switch name {
				case "bigEndian":
					return hxrt.AnonymousField{Value: hx_structural_source_123.bigEndian, Found: true}
				case "readByte":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_124 := hx_structural_source_123
						return hxrt.BoundFunction(func() int {
							return hx_receiver_124.__hx_this.readByte()
						}, func() any {
							return hx_receiver_124.__hx_this
						}, "readByte")
					}(), Found: true}
				case "readBytes":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes, int, int) int] {
						hx_receiver_125 := hx_structural_source_123
						return hxrt.BoundFunction(func(bytes *haxe__io__Bytes, pos int, len_local int) int {
							return hx_receiver_125.__hx_this.readBytes(bytes, pos, len_local)
						}, func() any {
							return hx_receiver_125.__hx_this
						}, "readBytes")
					}(), Found: true}
				case "close":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func()] {
						hx_receiver_126 := hx_structural_source_123
						return hxrt.BoundFunction(func() {
							hx_receiver_126.__hx_this.close()
						}, func() any {
							return hx_receiver_126.__hx_this
						}, "close")
					}(), Found: true}
				case "set_bigEndian":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(bool) bool] {
						hx_receiver_127 := hx_structural_source_123
						return hxrt.BoundFunction(func(value bool) bool {
							return hx_receiver_127.__hx_this.set_bigEndian(value)
						}, func() any {
							return hx_receiver_127.__hx_this
						}, "set_bigEndian")
					}(), Found: true}
				case "readAll":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(any) *haxe__io__Bytes] {
						hx_receiver_128 := hx_structural_source_123
						return hxrt.BoundFunction(func(bufsize any) *haxe__io__Bytes {
							return hx_receiver_128.__hx_this.readAll(bufsize)
						}, func() any {
							return hx_receiver_128.__hx_this
						}, "readAll")
					}(), Found: true}
				case "readFullBytes":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes, int, int)] {
						hx_receiver_129 := hx_structural_source_123
						return hxrt.BoundFunction(func(bytes *haxe__io__Bytes, pos int, len_local int) {
							hx_receiver_129.__hx_this.readFullBytes(bytes, pos, len_local)
						}, func() any {
							return hx_receiver_129.__hx_this
						}, "readFullBytes")
					}(), Found: true}
				case "read":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int) *haxe__io__Bytes] {
						hx_receiver_130 := hx_structural_source_123
						return hxrt.BoundFunction(func(nbytes int) *haxe__io__Bytes {
							return hx_receiver_130.__hx_this.read(nbytes)
						}, func() any {
							return hx_receiver_130.__hx_this
						}, "read")
					}(), Found: true}
				case "readUntil":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int) *string] {
						hx_receiver_131 := hx_structural_source_123
						return hxrt.BoundFunction(func(end int) *string {
							return hx_receiver_131.__hx_this.readUntil(end)
						}, func() any {
							return hx_receiver_131.__hx_this
						}, "readUntil")
					}(), Found: true}
				case "readLine":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_132 := hx_structural_source_123
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_132.__hx_this.readLine()
						}, func() any {
							return hx_receiver_132.__hx_this
						}, "readLine")
					}(), Found: true}
				case "readFloat":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() float64] {
						hx_receiver_133 := hx_structural_source_123
						return hxrt.BoundFunction(func() float64 {
							return hx_receiver_133.__hx_this.readFloat()
						}, func() any {
							return hx_receiver_133.__hx_this
						}, "readFloat")
					}(), Found: true}
				case "readDouble":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() float64] {
						hx_receiver_134 := hx_structural_source_123
						return hxrt.BoundFunction(func() float64 {
							return hx_receiver_134.__hx_this.readDouble()
						}, func() any {
							return hx_receiver_134.__hx_this
						}, "readDouble")
					}(), Found: true}
				case "readInt8":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_135 := hx_structural_source_123
						return hxrt.BoundFunction(func() int {
							return hx_receiver_135.__hx_this.readInt8()
						}, func() any {
							return hx_receiver_135.__hx_this
						}, "readInt8")
					}(), Found: true}
				case "readInt16":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_136 := hx_structural_source_123
						return hxrt.BoundFunction(func() int {
							return hx_receiver_136.__hx_this.readInt16()
						}, func() any {
							return hx_receiver_136.__hx_this
						}, "readInt16")
					}(), Found: true}
				case "readUInt16":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_137 := hx_structural_source_123
						return hxrt.BoundFunction(func() int {
							return hx_receiver_137.__hx_this.readUInt16()
						}, func() any {
							return hx_receiver_137.__hx_this
						}, "readUInt16")
					}(), Found: true}
				case "readInt24":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_138 := hx_structural_source_123
						return hxrt.BoundFunction(func() int {
							return hx_receiver_138.__hx_this.readInt24()
						}, func() any {
							return hx_receiver_138.__hx_this
						}, "readInt24")
					}(), Found: true}
				case "readUInt24":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_139 := hx_structural_source_123
						return hxrt.BoundFunction(func() int {
							return hx_receiver_139.__hx_this.readUInt24()
						}, func() any {
							return hx_receiver_139.__hx_this
						}, "readUInt24")
					}(), Found: true}
				case "readInt32":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_140 := hx_structural_source_123
						return hxrt.BoundFunction(func() int {
							return hx_receiver_140.__hx_this.readInt32()
						}, func() any {
							return hx_receiver_140.__hx_this
						}, "readInt32")
					}(), Found: true}
				case "readString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, *haxe__io__Encoding) *string] {
						hx_receiver_141 := hx_structural_source_123
						return hxrt.BoundFunction(func(len_local int, encoding *haxe__io__Encoding) *string {
							return hx_receiver_141.__hx_this.readString(len_local, encoding)
						}, func() any {
							return hx_receiver_141.__hx_this
						}, "readString")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				}
				return false
			})
		}(source)
	case *haxe__io__Eof:
		return func(hx_structural_source_142 *haxe__io__Eof) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_142, func(name string) hxrt.AnonymousField {
				switch name {
				case "toString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_143 := hx_structural_source_142
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_143.__hx_this.toString()
						}, func() any {
							return hx_receiver_143.__hx_this
						}, "toString")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				}
				return false
			})
		}(source)
	case *haxe__io__BytesOutput:
		return func(hx_structural_source_144 *haxe__io__BytesOutput) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_144, func(name string) hxrt.AnonymousField {
				switch name {
				case "b":
					return hxrt.AnonymousField{Value: hx_structural_source_144.b, Found: true}
				case "get_length":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_145 := hx_structural_source_144
						return hxrt.BoundFunction(func() int {
							return hx_receiver_145.__hx_this.get_length()
						}, func() any {
							return hx_receiver_145.__hx_this
						}, "get_length")
					}(), Found: true}
				case "writeByte":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_146 := hx_structural_source_144
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_146.__hx_this.writeByte(value)
						}, func() any {
							return hx_receiver_146.__hx_this
						}, "writeByte")
					}(), Found: true}
				case "writeBytes":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes, int, int) int] {
						hx_receiver_147 := hx_structural_source_144
						return hxrt.BoundFunction(func(bytes *haxe__io__Bytes, pos int, len_local int) int {
							return hx_receiver_147.__hx_this.writeBytes(bytes, pos, len_local)
						}, func() any {
							return hx_receiver_147.__hx_this
						}, "writeBytes")
					}(), Found: true}
				case "getBytes":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *haxe__io__Bytes] {
						hx_receiver_148 := hx_structural_source_144
						return hxrt.BoundFunction(func() *haxe__io__Bytes {
							return hx_receiver_148.__hx_this.getBytes()
						}, func() any {
							return hx_receiver_148.__hx_this
						}, "getBytes")
					}(), Found: true}
				case "bigEndian":
					return hxrt.AnonymousField{Value: hx_structural_source_144.bigEndian, Found: true}
				case "flush":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func()] {
						hx_receiver_149 := hx_structural_source_144
						return hxrt.BoundFunction(func() {
							hx_receiver_149.__hx_this.flush()
						}, func() any {
							return hx_receiver_149.__hx_this
						}, "flush")
					}(), Found: true}
				case "close":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func()] {
						hx_receiver_150 := hx_structural_source_144
						return hxrt.BoundFunction(func() {
							hx_receiver_150.__hx_this.close()
						}, func() any {
							return hx_receiver_150.__hx_this
						}, "close")
					}(), Found: true}
				case "set_bigEndian":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(bool) bool] {
						hx_receiver_151 := hx_structural_source_144
						return hxrt.BoundFunction(func(value bool) bool {
							return hx_receiver_151.__hx_this.set_bigEndian(value)
						}, func() any {
							return hx_receiver_151.__hx_this
						}, "set_bigEndian")
					}(), Found: true}
				case "write":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes)] {
						hx_receiver_152 := hx_structural_source_144
						return hxrt.BoundFunction(func(bytes *haxe__io__Bytes) {
							hx_receiver_152.__hx_this.write(bytes)
						}, func() any {
							return hx_receiver_152.__hx_this
						}, "write")
					}(), Found: true}
				case "writeFullBytes":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes, int, int)] {
						hx_receiver_153 := hx_structural_source_144
						return hxrt.BoundFunction(func(bytes *haxe__io__Bytes, pos int, len_local int) {
							hx_receiver_153.__hx_this.writeFullBytes(bytes, pos, len_local)
						}, func() any {
							return hx_receiver_153.__hx_this
						}, "writeFullBytes")
					}(), Found: true}
				case "writeFloat":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(float64)] {
						hx_receiver_154 := hx_structural_source_144
						return hxrt.BoundFunction(func(value float64) {
							hx_receiver_154.__hx_this.writeFloat(value)
						}, func() any {
							return hx_receiver_154.__hx_this
						}, "writeFloat")
					}(), Found: true}
				case "writeDouble":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(float64)] {
						hx_receiver_155 := hx_structural_source_144
						return hxrt.BoundFunction(func(value float64) {
							hx_receiver_155.__hx_this.writeDouble(value)
						}, func() any {
							return hx_receiver_155.__hx_this
						}, "writeDouble")
					}(), Found: true}
				case "writeInt8":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_156 := hx_structural_source_144
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_156.__hx_this.writeInt8(value)
						}, func() any {
							return hx_receiver_156.__hx_this
						}, "writeInt8")
					}(), Found: true}
				case "writeInt16":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_157 := hx_structural_source_144
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_157.__hx_this.writeInt16(value)
						}, func() any {
							return hx_receiver_157.__hx_this
						}, "writeInt16")
					}(), Found: true}
				case "writeUInt16":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_158 := hx_structural_source_144
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_158.__hx_this.writeUInt16(value)
						}, func() any {
							return hx_receiver_158.__hx_this
						}, "writeUInt16")
					}(), Found: true}
				case "writeInt24":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_159 := hx_structural_source_144
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_159.__hx_this.writeInt24(value)
						}, func() any {
							return hx_receiver_159.__hx_this
						}, "writeInt24")
					}(), Found: true}
				case "writeUInt24":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_160 := hx_structural_source_144
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_160.__hx_this.writeUInt24(value)
						}, func() any {
							return hx_receiver_160.__hx_this
						}, "writeUInt24")
					}(), Found: true}
				case "writeInt32":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_161 := hx_structural_source_144
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_161.__hx_this.writeInt32(value)
						}, func() any {
							return hx_receiver_161.__hx_this
						}, "writeInt32")
					}(), Found: true}
				case "prepare":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_162 := hx_structural_source_144
						return hxrt.BoundFunction(func(nbytes int) {
							hx_receiver_162.__hx_this.prepare(nbytes)
						}, func() any {
							return hx_receiver_162.__hx_this
						}, "prepare")
					}(), Found: true}
				case "writeInput":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Input, any)] {
						hx_receiver_163 := hx_structural_source_144
						return hxrt.BoundFunction(func(input *haxe__io__Input, bufsize any) {
							hx_receiver_163.__hx_this.writeInput(input, bufsize)
						}, func() any {
							return hx_receiver_163.__hx_this
						}, "writeInput")
					}(), Found: true}
				case "writeString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string, *haxe__io__Encoding)] {
						hx_receiver_164 := hx_structural_source_144
						return hxrt.BoundFunction(func(value *string, encoding *haxe__io__Encoding) {
							hx_receiver_164.__hx_this.writeString(value, encoding)
						}, func() any {
							return hx_receiver_164.__hx_this
						}, "writeString")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "b":
					hx_structural_source_144.b = hx_recover_class_haxe__io__BytesBuffer(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe__io__BytesBuffer:
		return func(hx_structural_source_165 *haxe__io__BytesBuffer) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_165, func(name string) hxrt.AnonymousField {
				switch name {
				case "b":
					return hxrt.AnonymousField{Value: hx_structural_source_165.b, Found: true}
				case "get_length":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_168 := hx_structural_source_165
						return hxrt.BoundFunction(func() int {
							return hx_receiver_168.__hx_this.get_length()
						}, func() any {
							return hx_receiver_168.__hx_this
						}, "get_length")
					}(), Found: true}
				case "addByte":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_169 := hx_structural_source_165
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_169.__hx_this.addByte(value)
						}, func() any {
							return hx_receiver_169.__hx_this
						}, "addByte")
					}(), Found: true}
				case "add":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes)] {
						hx_receiver_170 := hx_structural_source_165
						return hxrt.BoundFunction(func(source *haxe__io__Bytes) {
							hx_receiver_170.__hx_this.add(source)
						}, func() any {
							return hx_receiver_170.__hx_this
						}, "add")
					}(), Found: true}
				case "addString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string, *haxe__io__Encoding)] {
						hx_receiver_171 := hx_structural_source_165
						return hxrt.BoundFunction(func(value *string, encoding *haxe__io__Encoding) {
							hx_receiver_171.__hx_this.addString(value, encoding)
						}, func() any {
							return hx_receiver_171.__hx_this
						}, "addString")
					}(), Found: true}
				case "addInt32":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_172 := hx_structural_source_165
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_172.__hx_this.addInt32(value)
						}, func() any {
							return hx_receiver_172.__hx_this
						}, "addInt32")
					}(), Found: true}
				case "addInt64":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe___Int64_____Int64)] {
						hx_receiver_173 := hx_structural_source_165
						return hxrt.BoundFunction(func(value *haxe___Int64_____Int64) {
							hx_receiver_173.__hx_this.addInt64(value)
						}, func() any {
							return hx_receiver_173.__hx_this
						}, "addInt64")
					}(), Found: true}
				case "addFloat":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(float64)] {
						hx_receiver_174 := hx_structural_source_165
						return hxrt.BoundFunction(func(value float64) {
							hx_receiver_174.__hx_this.addFloat(value)
						}, func() any {
							return hx_receiver_174.__hx_this
						}, "addFloat")
					}(), Found: true}
				case "addDouble":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(float64)] {
						hx_receiver_175 := hx_structural_source_165
						return hxrt.BoundFunction(func(value float64) {
							hx_receiver_175.__hx_this.addDouble(value)
						}, func() any {
							return hx_receiver_175.__hx_this
						}, "addDouble")
					}(), Found: true}
				case "addBytes":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes, int, int)] {
						hx_receiver_176 := hx_structural_source_165
						return hxrt.BoundFunction(func(source *haxe__io__Bytes, pos int, len_local int) {
							hx_receiver_176.__hx_this.addBytes(source, pos, len_local)
						}, func() any {
							return hx_receiver_176.__hx_this
						}, "addBytes")
					}(), Found: true}
				case "getBytes":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *haxe__io__Bytes] {
						hx_receiver_177 := hx_structural_source_165
						return hxrt.BoundFunction(func() *haxe__io__Bytes {
							return hx_receiver_177.__hx_this.getBytes()
						}, func() any {
							return hx_receiver_177.__hx_this
						}, "getBytes")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "b":
					hx_structural_source_165.b = func(hx_value_166 any) []int {
						if hx_value_166 == nil {
							var hx_zero_167 []int
							return hx_zero_167
						}
						return hx_value_166.([]int)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe__io__Bytes:
		return func(hx_structural_source_178 *haxe__io__Bytes) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_178, func(name string) hxrt.AnonymousField {
				switch name {
				case "length":
					return hxrt.AnonymousField{Value: hx_structural_source_178.length, Found: true}
				case "b":
					return hxrt.AnonymousField{Value: hx_structural_source_178.b, Found: true}
				case "__hx_raw":
					return hxrt.AnonymousField{Value: hx_structural_source_178.__hx_raw, Found: true}
				case "__hx_rawValid":
					return hxrt.AnonymousField{Value: hx_structural_source_178.__hx_rawValid, Found: true}
				case "__hx_dataExposed":
					return hxrt.AnonymousField{Value: hx_structural_source_178.__hx_dataExposed, Found: true}
				case "get":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int) int] {
						hx_receiver_187 := hx_structural_source_178
						return hxrt.BoundFunction(func(pos int) int {
							return hx_receiver_187.__hx_this.get(pos)
						}, func() any {
							return hx_receiver_187.__hx_this
						}, "get")
					}(), Found: true}
				case "set":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, int)] {
						hx_receiver_188 := hx_structural_source_178
						return hxrt.BoundFunction(func(pos int, value int) {
							hx_receiver_188.__hx_this.set(pos, value)
						}, func() any {
							return hx_receiver_188.__hx_this
						}, "set")
					}(), Found: true}
				case "blit":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, *haxe__io__Bytes, int, int)] {
						hx_receiver_189 := hx_structural_source_178
						return hxrt.BoundFunction(func(pos int, src *haxe__io__Bytes, srcpos int, len_local int) {
							hx_receiver_189.__hx_this.blit(pos, src, srcpos, len_local)
						}, func() any {
							return hx_receiver_189.__hx_this
						}, "blit")
					}(), Found: true}
				case "fill":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, int, int)] {
						hx_receiver_190 := hx_structural_source_178
						return hxrt.BoundFunction(func(pos int, len_local int, value int) {
							hx_receiver_190.__hx_this.fill(pos, len_local, value)
						}, func() any {
							return hx_receiver_190.__hx_this
						}, "fill")
					}(), Found: true}
				case "sub":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, int) *haxe__io__Bytes] {
						hx_receiver_191 := hx_structural_source_178
						return hxrt.BoundFunction(func(pos int, len_local int) *haxe__io__Bytes {
							return hx_receiver_191.__hx_this.sub(pos, len_local)
						}, func() any {
							return hx_receiver_191.__hx_this
						}, "sub")
					}(), Found: true}
				case "compare":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes) int] {
						hx_receiver_192 := hx_structural_source_178
						return hxrt.BoundFunction(func(other *haxe__io__Bytes) int {
							return hx_receiver_192.__hx_this.compare(other)
						}, func() any {
							return hx_receiver_192.__hx_this
						}, "compare")
					}(), Found: true}
				case "getDouble":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int) float64] {
						hx_receiver_193 := hx_structural_source_178
						return hxrt.BoundFunction(func(pos int) float64 {
							return hx_receiver_193.__hx_this.getDouble(pos)
						}, func() any {
							return hx_receiver_193.__hx_this
						}, "getDouble")
					}(), Found: true}
				case "getFloat":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int) float64] {
						hx_receiver_194 := hx_structural_source_178
						return hxrt.BoundFunction(func(pos int) float64 {
							return hx_receiver_194.__hx_this.getFloat(pos)
						}, func() any {
							return hx_receiver_194.__hx_this
						}, "getFloat")
					}(), Found: true}
				case "setDouble":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, float64)] {
						hx_receiver_195 := hx_structural_source_178
						return hxrt.BoundFunction(func(pos int, value float64) {
							hx_receiver_195.__hx_this.setDouble(pos, value)
						}, func() any {
							return hx_receiver_195.__hx_this
						}, "setDouble")
					}(), Found: true}
				case "setFloat":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, float64)] {
						hx_receiver_196 := hx_structural_source_178
						return hxrt.BoundFunction(func(pos int, value float64) {
							hx_receiver_196.__hx_this.setFloat(pos, value)
						}, func() any {
							return hx_receiver_196.__hx_this
						}, "setFloat")
					}(), Found: true}
				case "getUInt16":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int) int] {
						hx_receiver_197 := hx_structural_source_178
						return hxrt.BoundFunction(func(pos int) int {
							return hx_receiver_197.__hx_this.getUInt16(pos)
						}, func() any {
							return hx_receiver_197.__hx_this
						}, "getUInt16")
					}(), Found: true}
				case "setUInt16":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, int)] {
						hx_receiver_198 := hx_structural_source_178
						return hxrt.BoundFunction(func(pos int, value int) {
							hx_receiver_198.__hx_this.setUInt16(pos, value)
						}, func() any {
							return hx_receiver_198.__hx_this
						}, "setUInt16")
					}(), Found: true}
				case "getInt32":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int) int] {
						hx_receiver_199 := hx_structural_source_178
						return hxrt.BoundFunction(func(pos int) int {
							return hx_receiver_199.__hx_this.getInt32(pos)
						}, func() any {
							return hx_receiver_199.__hx_this
						}, "getInt32")
					}(), Found: true}
				case "getInt64":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int) *haxe___Int64_____Int64] {
						hx_receiver_200 := hx_structural_source_178
						return hxrt.BoundFunction(func(pos int) *haxe___Int64_____Int64 {
							return hx_receiver_200.__hx_this.getInt64(pos)
						}, func() any {
							return hx_receiver_200.__hx_this
						}, "getInt64")
					}(), Found: true}
				case "setInt32":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, int)] {
						hx_receiver_201 := hx_structural_source_178
						return hxrt.BoundFunction(func(pos int, value int) {
							hx_receiver_201.__hx_this.setInt32(pos, value)
						}, func() any {
							return hx_receiver_201.__hx_this
						}, "setInt32")
					}(), Found: true}
				case "setInt64":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, *haxe___Int64_____Int64)] {
						hx_receiver_202 := hx_structural_source_178
						return hxrt.BoundFunction(func(pos int, value *haxe___Int64_____Int64) {
							hx_receiver_202.__hx_this.setInt64(pos, value)
						}, func() any {
							return hx_receiver_202.__hx_this
						}, "setInt64")
					}(), Found: true}
				case "getString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, int, *haxe__io__Encoding) *string] {
						hx_receiver_203 := hx_structural_source_178
						return hxrt.BoundFunction(func(pos int, len_local int, encoding *haxe__io__Encoding) *string {
							return hx_receiver_203.__hx_this.getString(pos, len_local, encoding)
						}, func() any {
							return hx_receiver_203.__hx_this
						}, "getString")
					}(), Found: true}
				case "readString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, int) *string] {
						hx_receiver_204 := hx_structural_source_178
						return hxrt.BoundFunction(func(pos int, len_local int) *string {
							return hx_receiver_204.__hx_this.readString(pos, len_local)
						}, func() any {
							return hx_receiver_204.__hx_this
						}, "readString")
					}(), Found: true}
				case "toString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_205 := hx_structural_source_178
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_205.__hx_this.toString()
						}, func() any {
							return hx_receiver_205.__hx_this
						}, "toString")
					}(), Found: true}
				case "toHex":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_206 := hx_structural_source_178
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_206.__hx_this.toHex()
						}, func() any {
							return hx_receiver_206.__hx_this
						}, "toHex")
					}(), Found: true}
				case "getData":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() []int] {
						hx_receiver_207 := hx_structural_source_178
						return hxrt.BoundFunction(func() []int {
							return hx_receiver_207.__hx_this.getData()
						}, func() any {
							return hx_receiver_207.__hx_this
						}, "getData")
					}(), Found: true}
				case "__hx_nativeView":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *hxrt.ByteView] {
						hx_receiver_208 := hx_structural_source_178
						return hxrt.BoundFunction(func() *hxrt.ByteView {
							return hx_receiver_208.__hx_this.__hx_nativeView()
						}, func() any {
							return hx_receiver_208.__hx_this
						}, "__hx_nativeView")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "b":
					hx_structural_source_178.b = func(hx_value_179 any) []int {
						if hx_value_179 == nil {
							var hx_zero_180 []int
							return hx_zero_180
						}
						return hx_value_179.([]int)
					}(incoming)
					return true
				case "__hx_raw":
					hx_structural_source_178.__hx_raw = func(hx_value_181 any) *hxrt.ByteView {
						if hx_value_181 == nil {
							var hx_zero_182 *hxrt.ByteView
							return hx_zero_182
						}
						return hx_value_181.(*hxrt.ByteView)
					}(incoming)
					return true
				case "__hx_rawValid":
					hx_structural_source_178.__hx_rawValid = func(hx_value_183 any) bool {
						if hx_value_183 == nil {
							var hx_zero_184 bool
							return hx_zero_184
						}
						return hx_value_183.(bool)
					}(incoming)
					return true
				case "__hx_dataExposed":
					hx_structural_source_178.__hx_dataExposed = func(hx_value_185 any) bool {
						if hx_value_185 == nil {
							var hx_zero_186 bool
							return hx_zero_186
						}
						return hx_value_185.(bool)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe__http__HttpBase:
		return func(hx_structural_source_209 *haxe__http__HttpBase) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_209, func(name string) hxrt.AnonymousField {
				switch name {
				case "url":
					return hxrt.AnonymousField{Value: hx_structural_source_209.url, Found: true}
				case "responseBytes":
					return hxrt.AnonymousField{Value: hx_structural_source_209.responseBytes, Found: true}
				case "responseAsString":
					return hxrt.AnonymousField{Value: hx_structural_source_209.responseAsString, Found: true}
				case "postData":
					return hxrt.AnonymousField{Value: hx_structural_source_209.postData, Found: true}
				case "postBytes":
					return hxrt.AnonymousField{Value: hx_structural_source_209.postBytes, Found: true}
				case "headers":
					return hxrt.AnonymousField{Value: hx_structural_source_209.headers, Found: true}
				case "params":
					return hxrt.AnonymousField{Value: hx_structural_source_209.params, Found: true}
				case "emptyOnData":
					return hxrt.AnonymousField{Value: hx_structural_source_209.emptyOnData, Found: true}
				case "setHeader":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string, *string)] {
						hx_receiver_220 := hx_structural_source_209
						return hxrt.BoundFunction(func(name *string, value *string) {
							hx_receiver_220.__hx_this.setHeader(name, value)
						}, func() any {
							return hx_receiver_220.__hx_this
						}, "setHeader")
					}(), Found: true}
				case "addHeader":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string, *string)] {
						hx_receiver_221 := hx_structural_source_209
						return hxrt.BoundFunction(func(header *string, value *string) {
							hx_receiver_221.__hx_this.addHeader(header, value)
						}, func() any {
							return hx_receiver_221.__hx_this
						}, "addHeader")
					}(), Found: true}
				case "setParameter":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string, *string)] {
						hx_receiver_222 := hx_structural_source_209
						return hxrt.BoundFunction(func(name *string, value *string) {
							hx_receiver_222.__hx_this.setParameter(name, value)
						}, func() any {
							return hx_receiver_222.__hx_this
						}, "setParameter")
					}(), Found: true}
				case "addParameter":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string, *string)] {
						hx_receiver_223 := hx_structural_source_209
						return hxrt.BoundFunction(func(name *string, value *string) {
							hx_receiver_223.__hx_this.addParameter(name, value)
						}, func() any {
							return hx_receiver_223.__hx_this
						}, "addParameter")
					}(), Found: true}
				case "setPostData":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string)] {
						hx_receiver_224 := hx_structural_source_209
						return hxrt.BoundFunction(func(data *string) {
							hx_receiver_224.__hx_this.setPostData(data)
						}, func() any {
							return hx_receiver_224.__hx_this
						}, "setPostData")
					}(), Found: true}
				case "setPostBytes":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes)] {
						hx_receiver_225 := hx_structural_source_209
						return hxrt.BoundFunction(func(data *haxe__io__Bytes) {
							hx_receiver_225.__hx_this.setPostBytes(data)
						}, func() any {
							return hx_receiver_225.__hx_this
						}, "setPostBytes")
					}(), Found: true}
				case "request":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(any)] {
						hx_receiver_226 := hx_structural_source_209
						return hxrt.BoundFunction(func(post any) {
							hx_receiver_226.__hx_this.request(post)
						}, func() any {
							return hx_receiver_226.__hx_this
						}, "request")
					}(), Found: true}
				case "onData":
					return hxrt.AnonymousField{Value: hxrt.ReadDynamicFunction(hx_structural_source_209.onData), Found: true}
				case "onBytes":
					return hxrt.AnonymousField{Value: hxrt.ReadDynamicFunction(hx_structural_source_209.onBytes), Found: true}
				case "onError":
					return hxrt.AnonymousField{Value: hxrt.ReadDynamicFunction(hx_structural_source_209.onError), Found: true}
				case "onStatus":
					return hxrt.AnonymousField{Value: hxrt.ReadDynamicFunction(hx_structural_source_209.onStatus), Found: true}
				case "hasOnData":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() bool] {
						hx_receiver_227 := hx_structural_source_209
						return hxrt.BoundFunction(func() bool {
							return hx_receiver_227.__hx_this.hasOnData()
						}, func() any {
							return hx_receiver_227.__hx_this
						}, "hasOnData")
					}(), Found: true}
				case "success":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes)] {
						hx_receiver_228 := hx_structural_source_209
						return hxrt.BoundFunction(func(data *haxe__io__Bytes) {
							hx_receiver_228.__hx_this.success(data)
						}, func() any {
							return hx_receiver_228.__hx_this
						}, "success")
					}(), Found: true}
				case "get_responseData":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_229 := hx_structural_source_209
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_229.__hx_this.get_responseData()
						}, func() any {
							return hx_receiver_229.__hx_this
						}, "get_responseData")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "url":
					hx_structural_source_209.url = func(hx_value_210 any) *string {
						if hx_value_210 == nil {
							var hx_zero_211 *string
							return hx_zero_211
						}
						return hx_value_210.(*string)
					}(incoming)
					return true
				case "responseAsString":
					hx_structural_source_209.responseAsString = func(hx_value_212 any) *string {
						if hx_value_212 == nil {
							var hx_zero_213 *string
							return hx_zero_213
						}
						return hx_value_212.(*string)
					}(incoming)
					return true
				case "postData":
					hx_structural_source_209.postData = func(hx_value_214 any) *string {
						if hx_value_214 == nil {
							var hx_zero_215 *string
							return hx_zero_215
						}
						return hx_value_214.(*string)
					}(incoming)
					return true
				case "postBytes":
					hx_structural_source_209.postBytes = hx_recover_class_haxe__io__Bytes(incoming)
					return true
				case "headers":
					hx_structural_source_209.headers = func(hx_value_216 any) *hxrt.Array {
						if hx_value_216 == nil {
							var hx_zero_217 *hxrt.Array
							return hx_zero_217
						}
						return hx_value_216.(*hxrt.Array)
					}(incoming)
					return true
				case "params":
					hx_structural_source_209.params = func(hx_value_218 any) *hxrt.Array {
						if hx_value_218 == nil {
							var hx_zero_219 *hxrt.Array
							return hx_zero_219
						}
						return hx_value_218.(*hxrt.Array)
					}(incoming)
					return true
				case "onData":
					hx_structural_source_209.onData = hx_recover_callable_9da3ad96(incoming)
					return true
				case "onBytes":
					hx_structural_source_209.onBytes = hx_recover_callable_c5f1b802(incoming)
					return true
				case "onError":
					hx_structural_source_209.onError = hx_recover_callable_9da3ad96(incoming)
					return true
				case "onStatus":
					hx_structural_source_209.onStatus = hx_recover_callable_d8be58e9(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe__exceptions__NotImplementedException:
		return func(hx_structural_source_230 *haxe__exceptions__NotImplementedException) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_230, func(name string) hxrt.AnonymousField {
				switch name {
				case "posInfos":
					return hxrt.AnonymousField{Value: hx_structural_source_230.posInfos, Found: true}
				case "toString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_231 := hx_structural_source_230
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_231.__hx_this.toString()
						}, func() any {
							return hx_receiver_231.__hx_this
						}, "toString")
					}(), Found: true}
				case "savedMessage":
					return hxrt.AnonymousField{Value: hx_structural_source_230.savedMessage, Found: true}
				case "savedStack":
					return hxrt.AnonymousField{Value: hx_structural_source_230.savedStack, Found: true}
				case "savedPrevious":
					return hxrt.AnonymousField{Value: hx_structural_source_230.savedPrevious, Found: true}
				case "savedNative":
					return hxrt.AnonymousField{Value: hx_structural_source_230.savedNative, Found: true}
				case "get_message":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_234 := hx_structural_source_230
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_234.__hx_this.get_message()
						}, func() any {
							return hx_receiver_234.__hx_this
						}, "get_message")
					}(), Found: true}
				case "get_stack":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_235 := hx_structural_source_230
						return hxrt.BoundFunction(func() any {
							return hx_receiver_235.__hx_this.get_stack()
						}, func() any {
							return hx_receiver_235.__hx_this
						}, "get_stack")
					}(), Found: true}
				case "get_previous":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *haxe__Exception] {
						hx_receiver_236 := hx_structural_source_230
						return hxrt.BoundFunction(func() *haxe__Exception {
							return hx_receiver_236.__hx_this.get_previous()
						}, func() any {
							return hx_receiver_236.__hx_this
						}, "get_previous")
					}(), Found: true}
				case "get_native":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_237 := hx_structural_source_230
						return hxrt.BoundFunction(func() any {
							return hx_receiver_237.__hx_this.get_native()
						}, func() any {
							return hx_receiver_237.__hx_this
						}, "get_native")
					}(), Found: true}
				case "unwrap":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_238 := hx_structural_source_230
						return hxrt.BoundFunction(func() any {
							return hx_receiver_238.__hx_this.unwrap()
						}, func() any {
							return hx_receiver_238.__hx_this
						}, "unwrap")
					}(), Found: true}
				case "details":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_239 := hx_structural_source_230
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_239.__hx_this.details()
						}, func() any {
							return hx_receiver_239.__hx_this
						}, "details")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "savedMessage":
					hx_structural_source_230.savedMessage = func(hx_value_232 any) *string {
						if hx_value_232 == nil {
							var hx_zero_233 *string
							return hx_zero_233
						}
						return hx_value_232.(*string)
					}(incoming)
					return true
				case "savedStack":
					hx_structural_source_230.savedStack = incoming
					return true
				case "savedPrevious":
					hx_structural_source_230.savedPrevious = hx_recover_class_haxe__Exception(incoming)
					return true
				case "savedNative":
					hx_structural_source_230.savedNative = incoming
					return true
				}
				return false
			})
		}(source)
	case *haxe__ds__StringMap:
		return func(hx_structural_source_240 *haxe__ds__StringMap) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_240, func(name string) hxrt.AnonymousField {
				switch name {
				case "h":
					return hxrt.AnonymousField{Value: hx_structural_source_240.h, Found: true}
				case "set":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string, any)] {
						hx_receiver_243 := hx_structural_source_240
						return hxrt.BoundFunction(func(key *string, value any) {
							hx_receiver_243.__hx_this.set(key, value)
						}, func() any {
							return hx_receiver_243.__hx_this
						}, "set")
					}(), Found: true}
				case "get":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string) any] {
						hx_receiver_244 := hx_structural_source_240
						return hxrt.BoundFunction(func(key *string) any {
							return hx_receiver_244.__hx_this.get(key)
						}, func() any {
							return hx_receiver_244.__hx_this
						}, "get")
					}(), Found: true}
				case "exists":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string) bool] {
						hx_receiver_245 := hx_structural_source_240
						return hxrt.BoundFunction(func(key *string) bool {
							return hx_receiver_245.__hx_this.exists(key)
						}, func() any {
							return hx_receiver_245.__hx_this
						}, "exists")
					}(), Found: true}
				case "remove":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string) bool] {
						hx_receiver_246 := hx_structural_source_240
						return hxrt.BoundFunction(func(key *string) bool {
							return hx_receiver_246.__hx_this.remove(key)
						}, func() any {
							return hx_receiver_246.__hx_this
						}, "remove")
					}(), Found: true}
				case "keys":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() hxrt.Anonymous] {
						hx_receiver_247 := hx_structural_source_240
						return hxrt.BoundFunction(func() hxrt.Anonymous {
							return hx_receiver_247.__hx_this.keys()
						}, func() any {
							return hx_receiver_247.__hx_this
						}, "keys")
					}(), Found: true}
				case "iterator":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() hxrt.Anonymous] {
						hx_receiver_248 := hx_structural_source_240
						return hxrt.BoundFunction(func() hxrt.Anonymous {
							return hx_receiver_248.__hx_this.iterator()
						}, func() any {
							return hx_receiver_248.__hx_this
						}, "iterator")
					}(), Found: true}
				case "keyValueIterator":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() hxrt.Anonymous] {
						hx_receiver_249 := hx_structural_source_240
						return hxrt.BoundFunction(func() hxrt.Anonymous {
							return hx_receiver_249.__hx_this.keyValueIterator()
						}, func() any {
							return hx_receiver_249.__hx_this
						}, "keyValueIterator")
					}(), Found: true}
				case "getIMap":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(any) any] {
						hx_receiver_250 := hx_structural_source_240
						return hxrt.BoundFunction(func(key any) any {
							return hx_receiver_250.__hx_this.getIMap(key)
						}, func() any {
							return hx_receiver_250.__hx_this
						}, "getIMap")
					}(), Found: true}
				case "setIMap":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(any, any)] {
						hx_receiver_251 := hx_structural_source_240
						return hxrt.BoundFunction(func(key any, value any) {
							hx_receiver_251.__hx_this.setIMap(key, value)
						}, func() any {
							return hx_receiver_251.__hx_this
						}, "setIMap")
					}(), Found: true}
				case "existsIMap":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(any) bool] {
						hx_receiver_252 := hx_structural_source_240
						return hxrt.BoundFunction(func(key any) bool {
							return hx_receiver_252.__hx_this.existsIMap(key)
						}, func() any {
							return hx_receiver_252.__hx_this
						}, "existsIMap")
					}(), Found: true}
				case "removeIMap":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(any) bool] {
						hx_receiver_253 := hx_structural_source_240
						return hxrt.BoundFunction(func(key any) bool {
							return hx_receiver_253.__hx_this.removeIMap(key)
						}, func() any {
							return hx_receiver_253.__hx_this
						}, "removeIMap")
					}(), Found: true}
				case "copyIMap":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() haxe__IMap] {
						hx_receiver_254 := hx_structural_source_240
						return hxrt.BoundFunction(func() haxe__IMap {
							return hx_receiver_254.__hx_this.copyIMap()
						}, func() any {
							return hx_receiver_254.__hx_this
						}, "copyIMap")
					}(), Found: true}
				case "copy":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *haxe__ds__StringMap] {
						hx_receiver_255 := hx_structural_source_240
						return hxrt.BoundFunction(func() *haxe__ds__StringMap {
							return hx_receiver_255.__hx_this.copy()
						}, func() any {
							return hx_receiver_255.__hx_this
						}, "copy")
					}(), Found: true}
				case "toString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_256 := hx_structural_source_240
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_256.__hx_this.toString()
						}, func() any {
							return hx_receiver_256.__hx_this
						}, "toString")
					}(), Found: true}
				case "clear":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func()] {
						hx_receiver_257 := hx_structural_source_240
						return hxrt.BoundFunction(func() {
							hx_receiver_257.__hx_this.clear()
						}, func() any {
							return hx_receiver_257.__hx_this
						}, "clear")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "h":
					hx_structural_source_240.h = func(hx_value_241 any) *hxrt.StringMapCell {
						if hx_value_241 == nil {
							var hx_zero_242 *hxrt.StringMapCell
							return hx_zero_242
						}
						return hx_value_241.(*hxrt.StringMapCell)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe__Exception:
		return func(hx_structural_source_258 *haxe__Exception) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_258, func(name string) hxrt.AnonymousField {
				switch name {
				case "savedMessage":
					return hxrt.AnonymousField{Value: hx_structural_source_258.savedMessage, Found: true}
				case "savedStack":
					return hxrt.AnonymousField{Value: hx_structural_source_258.savedStack, Found: true}
				case "savedPrevious":
					return hxrt.AnonymousField{Value: hx_structural_source_258.savedPrevious, Found: true}
				case "savedNative":
					return hxrt.AnonymousField{Value: hx_structural_source_258.savedNative, Found: true}
				case "get_message":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_261 := hx_structural_source_258
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_261.__hx_this.get_message()
						}, func() any {
							return hx_receiver_261.__hx_this
						}, "get_message")
					}(), Found: true}
				case "get_stack":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_262 := hx_structural_source_258
						return hxrt.BoundFunction(func() any {
							return hx_receiver_262.__hx_this.get_stack()
						}, func() any {
							return hx_receiver_262.__hx_this
						}, "get_stack")
					}(), Found: true}
				case "get_previous":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *haxe__Exception] {
						hx_receiver_263 := hx_structural_source_258
						return hxrt.BoundFunction(func() *haxe__Exception {
							return hx_receiver_263.__hx_this.get_previous()
						}, func() any {
							return hx_receiver_263.__hx_this
						}, "get_previous")
					}(), Found: true}
				case "get_native":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_264 := hx_structural_source_258
						return hxrt.BoundFunction(func() any {
							return hx_receiver_264.__hx_this.get_native()
						}, func() any {
							return hx_receiver_264.__hx_this
						}, "get_native")
					}(), Found: true}
				case "unwrap":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_265 := hx_structural_source_258
						return hxrt.BoundFunction(func() any {
							return hx_receiver_265.__hx_this.unwrap()
						}, func() any {
							return hx_receiver_265.__hx_this
						}, "unwrap")
					}(), Found: true}
				case "toString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_266 := hx_structural_source_258
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_266.__hx_this.toString()
						}, func() any {
							return hx_receiver_266.__hx_this
						}, "toString")
					}(), Found: true}
				case "details":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_267 := hx_structural_source_258
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_267.__hx_this.details()
						}, func() any {
							return hx_receiver_267.__hx_this
						}, "details")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "savedMessage":
					hx_structural_source_258.savedMessage = func(hx_value_259 any) *string {
						if hx_value_259 == nil {
							var hx_zero_260 *string
							return hx_zero_260
						}
						return hx_value_259.(*string)
					}(incoming)
					return true
				case "savedStack":
					hx_structural_source_258.savedStack = incoming
					return true
				case "savedPrevious":
					hx_structural_source_258.savedPrevious = hx_recover_class_haxe__Exception(incoming)
					return true
				case "savedNative":
					hx_structural_source_258.savedNative = incoming
					return true
				}
				return false
			})
		}(source)
	case *haxe__iterators__StringKeyValueIterator:
		return func(hx_structural_source_268 *haxe__iterators__StringKeyValueIterator) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_268, func(name string) hxrt.AnonymousField {
				switch name {
				case "offset":
					return hxrt.AnonymousField{Value: hx_structural_source_268.offset, Found: true}
				case "s":
					return hxrt.AnonymousField{Value: hx_structural_source_268.s, Found: true}
				case "hasNext":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() bool] {
						hx_receiver_273 := hx_structural_source_268
						return hxrt.BoundFunction(func() bool {
							return hx_receiver_273.__hx_this.hasNext()
						}, func() any {
							return hx_receiver_273.__hx_this
						}, "hasNext")
					}(), Found: true}
				case "next":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() hxrt.Anonymous] {
						hx_receiver_274 := hx_structural_source_268
						return hxrt.BoundFunction(func() hxrt.Anonymous {
							return hx_receiver_274.__hx_this.next()
						}, func() any {
							return hx_receiver_274.__hx_this
						}, "next")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "offset":
					hx_structural_source_268.offset = func(hx_value_269 any) int {
						if hx_value_269 == nil {
							var hx_zero_270 int
							return hx_zero_270
						}
						return hx_value_269.(int)
					}(incoming)
					return true
				case "s":
					hx_structural_source_268.s = func(hx_value_271 any) *string {
						if hx_value_271 == nil {
							var hx_zero_272 *string
							return hx_zero_272
						}
						return hx_value_271.(*string)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe__iterators__StringIterator:
		return func(hx_structural_source_275 *haxe__iterators__StringIterator) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_275, func(name string) hxrt.AnonymousField {
				switch name {
				case "offset":
					return hxrt.AnonymousField{Value: hx_structural_source_275.offset, Found: true}
				case "s":
					return hxrt.AnonymousField{Value: hx_structural_source_275.s, Found: true}
				case "hasNext":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() bool] {
						hx_receiver_280 := hx_structural_source_275
						return hxrt.BoundFunction(func() bool {
							return hx_receiver_280.__hx_this.hasNext()
						}, func() any {
							return hx_receiver_280.__hx_this
						}, "hasNext")
					}(), Found: true}
				case "next":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_281 := hx_structural_source_275
						return hxrt.BoundFunction(func() int {
							return hx_receiver_281.__hx_this.next()
						}, func() any {
							return hx_receiver_281.__hx_this
						}, "next")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "offset":
					hx_structural_source_275.offset = func(hx_value_276 any) int {
						if hx_value_276 == nil {
							var hx_zero_277 int
							return hx_zero_277
						}
						return hx_value_276.(int)
					}(incoming)
					return true
				case "s":
					hx_structural_source_275.s = func(hx_value_278 any) *string {
						if hx_value_278 == nil {
							var hx_zero_279 *string
							return hx_zero_279
						}
						return hx_value_278.(*string)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe__iterators__MapKeyValueIterator:
		return func(hx_structural_source_282 *haxe__iterators__MapKeyValueIterator) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_282, func(name string) hxrt.AnonymousField {
				switch name {
				case "map":
					return hxrt.AnonymousField{Value: hx_structural_source_282.map_, Found: true}
				case "keys":
					return hxrt.AnonymousField{Value: hx_structural_source_282.keys, Found: true}
				case "hasNext":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() bool] {
						hx_receiver_285 := hx_structural_source_282
						return hxrt.BoundFunction(func() bool {
							return hx_receiver_285.__hx_this.hasNext()
						}, func() any {
							return hx_receiver_285.__hx_this
						}, "hasNext")
					}(), Found: true}
				case "next":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() hxrt.Anonymous] {
						hx_receiver_286 := hx_structural_source_282
						return hxrt.BoundFunction(func() hxrt.Anonymous {
							return hx_receiver_286.__hx_this.next()
						}, func() any {
							return hx_receiver_286.__hx_this
						}, "next")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "map":
					hx_structural_source_282.map_ = func(hx_value_283 any) haxe__IMap {
						if hx_value_283 == nil {
							var hx_zero_284 haxe__IMap
							return hx_zero_284
						}
						return hx_value_283.(haxe__IMap)
					}(hxrt.ObjectSource(incoming))
					return true
				case "keys":
					hx_structural_source_282.keys = hx_structural_recover(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe__ValueException:
		return func(hx_structural_source_287 *haxe__ValueException) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_287, func(name string) hxrt.AnonymousField {
				switch name {
				case "value":
					return hxrt.AnonymousField{Value: hx_structural_source_287.value, Found: true}
				case "unwrap":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_288 := hx_structural_source_287
						return hxrt.BoundFunction(func() any {
							return hx_receiver_288.__hx_this.unwrap()
						}, func() any {
							return hx_receiver_288.__hx_this
						}, "unwrap")
					}(), Found: true}
				case "savedMessage":
					return hxrt.AnonymousField{Value: hx_structural_source_287.savedMessage, Found: true}
				case "savedStack":
					return hxrt.AnonymousField{Value: hx_structural_source_287.savedStack, Found: true}
				case "savedPrevious":
					return hxrt.AnonymousField{Value: hx_structural_source_287.savedPrevious, Found: true}
				case "savedNative":
					return hxrt.AnonymousField{Value: hx_structural_source_287.savedNative, Found: true}
				case "get_message":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_291 := hx_structural_source_287
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_291.__hx_this.get_message()
						}, func() any {
							return hx_receiver_291.__hx_this
						}, "get_message")
					}(), Found: true}
				case "get_stack":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_292 := hx_structural_source_287
						return hxrt.BoundFunction(func() any {
							return hx_receiver_292.__hx_this.get_stack()
						}, func() any {
							return hx_receiver_292.__hx_this
						}, "get_stack")
					}(), Found: true}
				case "get_previous":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *haxe__Exception] {
						hx_receiver_293 := hx_structural_source_287
						return hxrt.BoundFunction(func() *haxe__Exception {
							return hx_receiver_293.__hx_this.get_previous()
						}, func() any {
							return hx_receiver_293.__hx_this
						}, "get_previous")
					}(), Found: true}
				case "get_native":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_294 := hx_structural_source_287
						return hxrt.BoundFunction(func() any {
							return hx_receiver_294.__hx_this.get_native()
						}, func() any {
							return hx_receiver_294.__hx_this
						}, "get_native")
					}(), Found: true}
				case "toString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_295 := hx_structural_source_287
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_295.__hx_this.toString()
						}, func() any {
							return hx_receiver_295.__hx_this
						}, "toString")
					}(), Found: true}
				case "details":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_296 := hx_structural_source_287
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_296.__hx_this.details()
						}, func() any {
							return hx_receiver_296.__hx_this
						}, "details")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "savedMessage":
					hx_structural_source_287.savedMessage = func(hx_value_289 any) *string {
						if hx_value_289 == nil {
							var hx_zero_290 *string
							return hx_zero_290
						}
						return hx_value_289.(*string)
					}(incoming)
					return true
				case "savedStack":
					hx_structural_source_287.savedStack = incoming
					return true
				case "savedPrevious":
					hx_structural_source_287.savedPrevious = hx_recover_class_haxe__Exception(incoming)
					return true
				case "savedNative":
					hx_structural_source_287.savedNative = incoming
					return true
				}
				return false
			})
		}(source)
	case *haxe__exceptions__PosException:
		return func(hx_structural_source_297 *haxe__exceptions__PosException) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_297, func(name string) hxrt.AnonymousField {
				switch name {
				case "posInfos":
					return hxrt.AnonymousField{Value: hx_structural_source_297.posInfos, Found: true}
				case "toString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_298 := hx_structural_source_297
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_298.__hx_this.toString()
						}, func() any {
							return hx_receiver_298.__hx_this
						}, "toString")
					}(), Found: true}
				case "savedMessage":
					return hxrt.AnonymousField{Value: hx_structural_source_297.savedMessage, Found: true}
				case "savedStack":
					return hxrt.AnonymousField{Value: hx_structural_source_297.savedStack, Found: true}
				case "savedPrevious":
					return hxrt.AnonymousField{Value: hx_structural_source_297.savedPrevious, Found: true}
				case "savedNative":
					return hxrt.AnonymousField{Value: hx_structural_source_297.savedNative, Found: true}
				case "get_message":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_301 := hx_structural_source_297
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_301.__hx_this.get_message()
						}, func() any {
							return hx_receiver_301.__hx_this
						}, "get_message")
					}(), Found: true}
				case "get_stack":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_302 := hx_structural_source_297
						return hxrt.BoundFunction(func() any {
							return hx_receiver_302.__hx_this.get_stack()
						}, func() any {
							return hx_receiver_302.__hx_this
						}, "get_stack")
					}(), Found: true}
				case "get_previous":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *haxe__Exception] {
						hx_receiver_303 := hx_structural_source_297
						return hxrt.BoundFunction(func() *haxe__Exception {
							return hx_receiver_303.__hx_this.get_previous()
						}, func() any {
							return hx_receiver_303.__hx_this
						}, "get_previous")
					}(), Found: true}
				case "get_native":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_304 := hx_structural_source_297
						return hxrt.BoundFunction(func() any {
							return hx_receiver_304.__hx_this.get_native()
						}, func() any {
							return hx_receiver_304.__hx_this
						}, "get_native")
					}(), Found: true}
				case "unwrap":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_305 := hx_structural_source_297
						return hxrt.BoundFunction(func() any {
							return hx_receiver_305.__hx_this.unwrap()
						}, func() any {
							return hx_receiver_305.__hx_this
						}, "unwrap")
					}(), Found: true}
				case "details":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_306 := hx_structural_source_297
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_306.__hx_this.details()
						}, func() any {
							return hx_receiver_306.__hx_this
						}, "details")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "savedMessage":
					hx_structural_source_297.savedMessage = func(hx_value_299 any) *string {
						if hx_value_299 == nil {
							var hx_zero_300 *string
							return hx_zero_300
						}
						return hx_value_299.(*string)
					}(incoming)
					return true
				case "savedStack":
					hx_structural_source_297.savedStack = incoming
					return true
				case "savedPrevious":
					hx_structural_source_297.savedPrevious = hx_recover_class_haxe__Exception(incoming)
					return true
				case "savedNative":
					hx_structural_source_297.savedNative = incoming
					return true
				}
				return false
			})
		}(source)
	default:
		return hxrt.AnonymousFromValue(value)
	}
}

func hx_recover_callable_9da3ad96(value any) *hxrt.Function[func(*string)] {
	if value == nil {
		return nil
	}
	switch exact := value.(type) {
	case *hxrt.Function[func(*string)]:
		return exact
	}
	origin := hxrt.FunctionOrigin(value)
	if origin == nil {
		return nil
	}
	switch source := origin.(type) {
	case *hxrt.Function[func(*string)]:
		return source
	case *hxrt.Function[func(*haxe__io__Bytes)]:
		return hxrt.ProjectFunction(source, func(arg_0 *string) {
			hxrt.FunctionCall(source)(hx_recover_class_haxe__io__Bytes(arg_0))
		})
	case *hxrt.Function[func(int)]:
		return hxrt.ProjectFunction(source, func(arg_0 *string) {
			hxrt.FunctionCall(source)(func(hx_value_1 any) int {
				if hx_value_1 == nil {
					var hx_zero_2 int
					return hx_zero_2
				}
				return hx_value_1.(int)
			}(arg_0))
		})
	case *hxrt.Function[func(bool) bool]:
		return hxrt.ProjectFunction(source, func(arg_0 *string) {
			hxrt.FunctionCall(source)(func(hx_value_3 any) bool {
				if hx_value_3 == nil {
					var hx_zero_4 bool
					return hx_zero_4
				}
				return hx_value_3.(bool)
			}(arg_0))
		})
	case *hxrt.Function[func(float64)]:
		return hxrt.ProjectFunction(source, func(arg_0 *string) {
			hxrt.FunctionCall(source)(func(hx_value_5 any) float64 {
				if hx_value_5 == nil {
					var hx_zero_6 float64
					return hx_zero_6
				}
				return hx_value_5.(float64)
			}(arg_0))
		})
	case *hxrt.Function[func(any) *haxe__io__Bytes]:
		return hxrt.ProjectFunction(source, func(arg_0 *string) {
			hxrt.FunctionCall(source)(arg_0)
		})
	case *hxrt.Function[func(int) *haxe__io__Bytes]:
		return hxrt.ProjectFunction(source, func(arg_0 *string) {
			hxrt.FunctionCall(source)(func(hx_value_7 any) int {
				if hx_value_7 == nil {
					var hx_zero_8 int
					return hx_zero_8
				}
				return hx_value_7.(int)
			}(arg_0))
		})
	case *hxrt.Function[func(int) *string]:
		return hxrt.ProjectFunction(source, func(arg_0 *string) {
			hxrt.FunctionCall(source)(func(hx_value_9 any) int {
				if hx_value_9 == nil {
					var hx_zero_10 int
					return hx_zero_10
				}
				return hx_value_9.(int)
			}(arg_0))
		})
	case *hxrt.Function[func(bool)]:
		return hxrt.ProjectFunction(source, func(arg_0 *string) {
			hxrt.FunctionCall(source)(func(hx_value_11 any) bool {
				if hx_value_11 == nil {
					var hx_zero_12 bool
					return hx_zero_12
				}
				return hx_value_11.(bool)
			}(arg_0))
		})
	case *hxrt.Function[func(*hxrt.SocketHandle)]:
		return hxrt.ProjectFunction(source, func(arg_0 *string) {
			hxrt.FunctionCall(source)(func(hx_value_13 any) *hxrt.SocketHandle {
				if hx_value_13 == nil {
					var hx_zero_14 *hxrt.SocketHandle
					return hx_zero_14
				}
				return hx_value_13.(*hxrt.SocketHandle)
			}(arg_0))
		})
	case *hxrt.Function[func(any)]:
		return hxrt.ProjectFunction(source, func(arg_0 *string) {
			hxrt.FunctionCall(source)(arg_0)
		})
	case *hxrt.Function[func(*string) *hxrt.Array]:
		return hxrt.ProjectFunction(source, func(arg_0 *string) {
			hxrt.FunctionCall(source)(arg_0)
		})
	case *hxrt.Function[func(*hxrt.HttpExchange)]:
		return hxrt.ProjectFunction(source, func(arg_0 *string) {
			hxrt.FunctionCall(source)(func(hx_value_15 any) *hxrt.HttpExchange {
				if hx_value_15 == nil {
					var hx_zero_16 *hxrt.HttpExchange
					return hx_zero_16
				}
				return hx_value_15.(*hxrt.HttpExchange)
			}(arg_0))
		})
	case *hxrt.Function[func(*haxe___Int64_____Int64)]:
		return hxrt.ProjectFunction(source, func(arg_0 *string) {
			hxrt.FunctionCall(source)(hx_recover_class_haxe___Int64_____Int64(arg_0))
		})
	case *hxrt.Function[func(int) int]:
		return hxrt.ProjectFunction(source, func(arg_0 *string) {
			hxrt.FunctionCall(source)(func(hx_value_17 any) int {
				if hx_value_17 == nil {
					var hx_zero_18 int
					return hx_zero_18
				}
				return hx_value_17.(int)
			}(arg_0))
		})
	case *hxrt.Function[func(*haxe__io__Bytes) int]:
		return hxrt.ProjectFunction(source, func(arg_0 *string) {
			hxrt.FunctionCall(source)(hx_recover_class_haxe__io__Bytes(arg_0))
		})
	case *hxrt.Function[func(int) float64]:
		return hxrt.ProjectFunction(source, func(arg_0 *string) {
			hxrt.FunctionCall(source)(func(hx_value_19 any) int {
				if hx_value_19 == nil {
					var hx_zero_20 int
					return hx_zero_20
				}
				return hx_value_19.(int)
			}(arg_0))
		})
	case *hxrt.Function[func(int) *haxe___Int64_____Int64]:
		return hxrt.ProjectFunction(source, func(arg_0 *string) {
			hxrt.FunctionCall(source)(func(hx_value_21 any) int {
				if hx_value_21 == nil {
					var hx_zero_22 int
					return hx_zero_22
				}
				return hx_value_21.(int)
			}(arg_0))
		})
	case *hxrt.Function[func(*string) any]:
		return hxrt.ProjectFunction(source, func(arg_0 *string) {
			hxrt.FunctionCall(source)(arg_0)
		})
	case *hxrt.Function[func(*string) bool]:
		return hxrt.ProjectFunction(source, func(arg_0 *string) {
			hxrt.FunctionCall(source)(arg_0)
		})
	case *hxrt.Function[func(any) any]:
		return hxrt.ProjectFunction(source, func(arg_0 *string) {
			hxrt.FunctionCall(source)(arg_0)
		})
	case *hxrt.Function[func(any) bool]:
		return hxrt.ProjectFunction(source, func(arg_0 *string) {
			hxrt.FunctionCall(source)(arg_0)
		})
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}

func hx_recover_callable_b9d17ba0(value any) *hxrt.Function[func() *string] {
	if value == nil {
		return nil
	}
	switch exact := value.(type) {
	case *hxrt.Function[func() *string]:
		return exact
	}
	origin := hxrt.FunctionOrigin(value)
	if origin == nil {
		return nil
	}
	switch source := origin.(type) {
	case *hxrt.Function[func() bool]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_1 any) *string {
				if hx_value_1 == nil {
					var hx_zero_2 *string
					return hx_zero_2
				}
				return hx_value_1.(*string)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *string]:
		return source
	case *hxrt.Function[func() any]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_3 any) *string {
				if hx_value_3 == nil {
					var hx_zero_4 *string
					return hx_zero_4
				}
				return hx_value_3.(*string)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() hxrt.Anonymous]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_5 any) *string {
				if hx_value_5 == nil {
					var hx_zero_6 *string
					return hx_zero_6
				}
				return hx_value_5.(*string)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func()]:
		return hxrt.ProjectFunction(source, func() *string {
			hxrt.FunctionCall(source)()
			return func(hx_value_7 any) *string {
				if hx_value_7 == nil {
					var hx_zero_8 *string
					return hx_zero_8
				}
				return hx_value_7.(*string)
			}(nil)
		})
	case *hxrt.Function[func() int]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_9 any) *string {
				if hx_value_9 == nil {
					var hx_zero_10 *string
					return hx_zero_10
				}
				return hx_value_9.(*string)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() float64]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_11 any) *string {
				if hx_value_11 == nil {
					var hx_zero_12 *string
					return hx_zero_12
				}
				return hx_value_11.(*string)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *sys__net__Socket]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_13 any) *string {
				if hx_value_13 == nil {
					var hx_zero_14 *string
					return hx_zero_14
				}
				return hx_value_13.(*string)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *haxe__io__Bytes]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_15 any) *string {
				if hx_value_15 == nil {
					var hx_zero_16 *string
					return hx_zero_16
				}
				return hx_value_15.(*string)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() []int]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_17 any) *string {
				if hx_value_17 == nil {
					var hx_zero_18 *string
					return hx_zero_18
				}
				return hx_value_17.(*string)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *hxrt.ByteView]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_19 any) *string {
				if hx_value_19 == nil {
					var hx_zero_20 *string
					return hx_zero_20
				}
				return hx_value_19.(*string)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *haxe__Exception]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_21 any) *string {
				if hx_value_21 == nil {
					var hx_zero_22 *string
					return hx_zero_22
				}
				return hx_value_21.(*string)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() haxe__IMap]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_23 any) *string {
				if hx_value_23 == nil {
					var hx_zero_24 *string
					return hx_zero_24
				}
				return hx_value_23.(*string)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *haxe__ds__StringMap]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_25 any) *string {
				if hx_value_25 == nil {
					var hx_zero_26 *string
					return hx_zero_26
				}
				return hx_value_25.(*string)
			}(hxrt.FunctionCall(source)())
		})
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}

func hx_recover_callable_c5309b43(value any) *hxrt.Function[func() any] {
	if value == nil {
		return nil
	}
	switch exact := value.(type) {
	case *hxrt.Function[func() any]:
		return exact
	}
	origin := hxrt.FunctionOrigin(value)
	if origin == nil {
		return nil
	}
	switch source := origin.(type) {
	case *hxrt.Function[func() bool]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func() *string]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func() any]:
		return source
	case *hxrt.Function[func() hxrt.Anonymous]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func()]:
		return hxrt.ProjectFunction(source, func() any {
			hxrt.FunctionCall(source)()
			return nil
		})
	case *hxrt.Function[func() int]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func() float64]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func() *sys__net__Socket]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func() *haxe__io__Bytes]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func() []int]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func() *hxrt.ByteView]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func() *haxe__Exception]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func() haxe__IMap]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func() *haxe__ds__StringMap]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}

func hx_recover_callable_c5f1b802(value any) *hxrt.Function[func(*haxe__io__Bytes)] {
	if value == nil {
		return nil
	}
	switch exact := value.(type) {
	case *hxrt.Function[func(*haxe__io__Bytes)]:
		return exact
	}
	origin := hxrt.FunctionOrigin(value)
	if origin == nil {
		return nil
	}
	switch source := origin.(type) {
	case *hxrt.Function[func(*string)]:
		return hxrt.ProjectFunction(source, func(arg_0 *haxe__io__Bytes) {
			hxrt.FunctionCall(source)(func(hx_value_1 any) *string {
				if hx_value_1 == nil {
					var hx_zero_2 *string
					return hx_zero_2
				}
				return hx_value_1.(*string)
			}(arg_0))
		})
	case *hxrt.Function[func(*haxe__io__Bytes)]:
		return source
	case *hxrt.Function[func(int)]:
		return hxrt.ProjectFunction(source, func(arg_0 *haxe__io__Bytes) {
			hxrt.FunctionCall(source)(func(hx_value_3 any) int {
				if hx_value_3 == nil {
					var hx_zero_4 int
					return hx_zero_4
				}
				return hx_value_3.(int)
			}(arg_0))
		})
	case *hxrt.Function[func(bool) bool]:
		return hxrt.ProjectFunction(source, func(arg_0 *haxe__io__Bytes) {
			hxrt.FunctionCall(source)(func(hx_value_5 any) bool {
				if hx_value_5 == nil {
					var hx_zero_6 bool
					return hx_zero_6
				}
				return hx_value_5.(bool)
			}(arg_0))
		})
	case *hxrt.Function[func(float64)]:
		return hxrt.ProjectFunction(source, func(arg_0 *haxe__io__Bytes) {
			hxrt.FunctionCall(source)(func(hx_value_7 any) float64 {
				if hx_value_7 == nil {
					var hx_zero_8 float64
					return hx_zero_8
				}
				return hx_value_7.(float64)
			}(arg_0))
		})
	case *hxrt.Function[func(any) *haxe__io__Bytes]:
		return hxrt.ProjectFunction(source, func(arg_0 *haxe__io__Bytes) {
			hxrt.FunctionCall(source)(arg_0)
		})
	case *hxrt.Function[func(int) *haxe__io__Bytes]:
		return hxrt.ProjectFunction(source, func(arg_0 *haxe__io__Bytes) {
			hxrt.FunctionCall(source)(func(hx_value_9 any) int {
				if hx_value_9 == nil {
					var hx_zero_10 int
					return hx_zero_10
				}
				return hx_value_9.(int)
			}(arg_0))
		})
	case *hxrt.Function[func(int) *string]:
		return hxrt.ProjectFunction(source, func(arg_0 *haxe__io__Bytes) {
			hxrt.FunctionCall(source)(func(hx_value_11 any) int {
				if hx_value_11 == nil {
					var hx_zero_12 int
					return hx_zero_12
				}
				return hx_value_11.(int)
			}(arg_0))
		})
	case *hxrt.Function[func(bool)]:
		return hxrt.ProjectFunction(source, func(arg_0 *haxe__io__Bytes) {
			hxrt.FunctionCall(source)(func(hx_value_13 any) bool {
				if hx_value_13 == nil {
					var hx_zero_14 bool
					return hx_zero_14
				}
				return hx_value_13.(bool)
			}(arg_0))
		})
	case *hxrt.Function[func(*hxrt.SocketHandle)]:
		return hxrt.ProjectFunction(source, func(arg_0 *haxe__io__Bytes) {
			hxrt.FunctionCall(source)(func(hx_value_15 any) *hxrt.SocketHandle {
				if hx_value_15 == nil {
					var hx_zero_16 *hxrt.SocketHandle
					return hx_zero_16
				}
				return hx_value_15.(*hxrt.SocketHandle)
			}(arg_0))
		})
	case *hxrt.Function[func(any)]:
		return hxrt.ProjectFunction(source, func(arg_0 *haxe__io__Bytes) {
			hxrt.FunctionCall(source)(arg_0)
		})
	case *hxrt.Function[func(*string) *hxrt.Array]:
		return hxrt.ProjectFunction(source, func(arg_0 *haxe__io__Bytes) {
			hxrt.FunctionCall(source)(func(hx_value_17 any) *string {
				if hx_value_17 == nil {
					var hx_zero_18 *string
					return hx_zero_18
				}
				return hx_value_17.(*string)
			}(arg_0))
		})
	case *hxrt.Function[func(*hxrt.HttpExchange)]:
		return hxrt.ProjectFunction(source, func(arg_0 *haxe__io__Bytes) {
			hxrt.FunctionCall(source)(func(hx_value_19 any) *hxrt.HttpExchange {
				if hx_value_19 == nil {
					var hx_zero_20 *hxrt.HttpExchange
					return hx_zero_20
				}
				return hx_value_19.(*hxrt.HttpExchange)
			}(arg_0))
		})
	case *hxrt.Function[func(*haxe___Int64_____Int64)]:
		return hxrt.ProjectFunction(source, func(arg_0 *haxe__io__Bytes) {
			hxrt.FunctionCall(source)(hx_recover_class_haxe___Int64_____Int64(arg_0))
		})
	case *hxrt.Function[func(int) int]:
		return hxrt.ProjectFunction(source, func(arg_0 *haxe__io__Bytes) {
			hxrt.FunctionCall(source)(func(hx_value_21 any) int {
				if hx_value_21 == nil {
					var hx_zero_22 int
					return hx_zero_22
				}
				return hx_value_21.(int)
			}(arg_0))
		})
	case *hxrt.Function[func(*haxe__io__Bytes) int]:
		return hxrt.ProjectFunction(source, func(arg_0 *haxe__io__Bytes) {
			hxrt.FunctionCall(source)(arg_0)
		})
	case *hxrt.Function[func(int) float64]:
		return hxrt.ProjectFunction(source, func(arg_0 *haxe__io__Bytes) {
			hxrt.FunctionCall(source)(func(hx_value_23 any) int {
				if hx_value_23 == nil {
					var hx_zero_24 int
					return hx_zero_24
				}
				return hx_value_23.(int)
			}(arg_0))
		})
	case *hxrt.Function[func(int) *haxe___Int64_____Int64]:
		return hxrt.ProjectFunction(source, func(arg_0 *haxe__io__Bytes) {
			hxrt.FunctionCall(source)(func(hx_value_25 any) int {
				if hx_value_25 == nil {
					var hx_zero_26 int
					return hx_zero_26
				}
				return hx_value_25.(int)
			}(arg_0))
		})
	case *hxrt.Function[func(*string) any]:
		return hxrt.ProjectFunction(source, func(arg_0 *haxe__io__Bytes) {
			hxrt.FunctionCall(source)(func(hx_value_27 any) *string {
				if hx_value_27 == nil {
					var hx_zero_28 *string
					return hx_zero_28
				}
				return hx_value_27.(*string)
			}(arg_0))
		})
	case *hxrt.Function[func(*string) bool]:
		return hxrt.ProjectFunction(source, func(arg_0 *haxe__io__Bytes) {
			hxrt.FunctionCall(source)(func(hx_value_29 any) *string {
				if hx_value_29 == nil {
					var hx_zero_30 *string
					return hx_zero_30
				}
				return hx_value_29.(*string)
			}(arg_0))
		})
	case *hxrt.Function[func(any) any]:
		return hxrt.ProjectFunction(source, func(arg_0 *haxe__io__Bytes) {
			hxrt.FunctionCall(source)(arg_0)
		})
	case *hxrt.Function[func(any) bool]:
		return hxrt.ProjectFunction(source, func(arg_0 *haxe__io__Bytes) {
			hxrt.FunctionCall(source)(arg_0)
		})
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}

func hx_recover_callable_cf6928b1(value any) *hxrt.Function[func() bool] {
	if value == nil {
		return nil
	}
	switch exact := value.(type) {
	case *hxrt.Function[func() bool]:
		return exact
	}
	origin := hxrt.FunctionOrigin(value)
	if origin == nil {
		return nil
	}
	switch source := origin.(type) {
	case *hxrt.Function[func() bool]:
		return source
	case *hxrt.Function[func() *string]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_1 any) bool {
				if hx_value_1 == nil {
					var hx_zero_2 bool
					return hx_zero_2
				}
				return hx_value_1.(bool)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() any]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_3 any) bool {
				if hx_value_3 == nil {
					var hx_zero_4 bool
					return hx_zero_4
				}
				return hx_value_3.(bool)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() hxrt.Anonymous]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_5 any) bool {
				if hx_value_5 == nil {
					var hx_zero_6 bool
					return hx_zero_6
				}
				return hx_value_5.(bool)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func()]:
		return hxrt.ProjectFunction(source, func() bool {
			hxrt.FunctionCall(source)()
			return func(hx_value_7 any) bool {
				if hx_value_7 == nil {
					var hx_zero_8 bool
					return hx_zero_8
				}
				return hx_value_7.(bool)
			}(nil)
		})
	case *hxrt.Function[func() int]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_9 any) bool {
				if hx_value_9 == nil {
					var hx_zero_10 bool
					return hx_zero_10
				}
				return hx_value_9.(bool)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() float64]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_11 any) bool {
				if hx_value_11 == nil {
					var hx_zero_12 bool
					return hx_zero_12
				}
				return hx_value_11.(bool)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *sys__net__Socket]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_13 any) bool {
				if hx_value_13 == nil {
					var hx_zero_14 bool
					return hx_zero_14
				}
				return hx_value_13.(bool)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *haxe__io__Bytes]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_15 any) bool {
				if hx_value_15 == nil {
					var hx_zero_16 bool
					return hx_zero_16
				}
				return hx_value_15.(bool)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() []int]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_17 any) bool {
				if hx_value_17 == nil {
					var hx_zero_18 bool
					return hx_zero_18
				}
				return hx_value_17.(bool)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *hxrt.ByteView]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_19 any) bool {
				if hx_value_19 == nil {
					var hx_zero_20 bool
					return hx_zero_20
				}
				return hx_value_19.(bool)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *haxe__Exception]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_21 any) bool {
				if hx_value_21 == nil {
					var hx_zero_22 bool
					return hx_zero_22
				}
				return hx_value_21.(bool)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() haxe__IMap]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_23 any) bool {
				if hx_value_23 == nil {
					var hx_zero_24 bool
					return hx_zero_24
				}
				return hx_value_23.(bool)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *haxe__ds__StringMap]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_25 any) bool {
				if hx_value_25 == nil {
					var hx_zero_26 bool
					return hx_zero_26
				}
				return hx_value_25.(bool)
			}(hxrt.FunctionCall(source)())
		})
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}

func hx_recover_callable_d8be58e9(value any) *hxrt.Function[func(int)] {
	if value == nil {
		return nil
	}
	switch exact := value.(type) {
	case *hxrt.Function[func(int)]:
		return exact
	}
	origin := hxrt.FunctionOrigin(value)
	if origin == nil {
		return nil
	}
	switch source := origin.(type) {
	case *hxrt.Function[func(*string)]:
		return hxrt.ProjectFunction(source, func(arg_0 int) {
			hxrt.FunctionCall(source)(func(hx_value_1 any) *string {
				if hx_value_1 == nil {
					var hx_zero_2 *string
					return hx_zero_2
				}
				return hx_value_1.(*string)
			}(arg_0))
		})
	case *hxrt.Function[func(*haxe__io__Bytes)]:
		return hxrt.ProjectFunction(source, func(arg_0 int) {
			hxrt.FunctionCall(source)(hx_recover_class_haxe__io__Bytes(arg_0))
		})
	case *hxrt.Function[func(int)]:
		return source
	case *hxrt.Function[func(bool) bool]:
		return hxrt.ProjectFunction(source, func(arg_0 int) {
			hxrt.FunctionCall(source)(func(hx_value_3 any) bool {
				if hx_value_3 == nil {
					var hx_zero_4 bool
					return hx_zero_4
				}
				return hx_value_3.(bool)
			}(arg_0))
		})
	case *hxrt.Function[func(float64)]:
		return hxrt.ProjectFunction(source, func(arg_0 int) {
			hxrt.FunctionCall(source)(float64(arg_0))
		})
	case *hxrt.Function[func(any) *haxe__io__Bytes]:
		return hxrt.ProjectFunction(source, func(arg_0 int) {
			hxrt.FunctionCall(source)(arg_0)
		})
	case *hxrt.Function[func(int) *haxe__io__Bytes]:
		return hxrt.ProjectFunction(source, func(arg_0 int) {
			hxrt.FunctionCall(source)(arg_0)
		})
	case *hxrt.Function[func(int) *string]:
		return hxrt.ProjectFunction(source, func(arg_0 int) {
			hxrt.FunctionCall(source)(arg_0)
		})
	case *hxrt.Function[func(bool)]:
		return hxrt.ProjectFunction(source, func(arg_0 int) {
			hxrt.FunctionCall(source)(func(hx_value_5 any) bool {
				if hx_value_5 == nil {
					var hx_zero_6 bool
					return hx_zero_6
				}
				return hx_value_5.(bool)
			}(arg_0))
		})
	case *hxrt.Function[func(*hxrt.SocketHandle)]:
		return hxrt.ProjectFunction(source, func(arg_0 int) {
			hxrt.FunctionCall(source)(func(hx_value_7 any) *hxrt.SocketHandle {
				if hx_value_7 == nil {
					var hx_zero_8 *hxrt.SocketHandle
					return hx_zero_8
				}
				return hx_value_7.(*hxrt.SocketHandle)
			}(arg_0))
		})
	case *hxrt.Function[func(any)]:
		return hxrt.ProjectFunction(source, func(arg_0 int) {
			hxrt.FunctionCall(source)(arg_0)
		})
	case *hxrt.Function[func(*string) *hxrt.Array]:
		return hxrt.ProjectFunction(source, func(arg_0 int) {
			hxrt.FunctionCall(source)(func(hx_value_9 any) *string {
				if hx_value_9 == nil {
					var hx_zero_10 *string
					return hx_zero_10
				}
				return hx_value_9.(*string)
			}(arg_0))
		})
	case *hxrt.Function[func(*hxrt.HttpExchange)]:
		return hxrt.ProjectFunction(source, func(arg_0 int) {
			hxrt.FunctionCall(source)(func(hx_value_11 any) *hxrt.HttpExchange {
				if hx_value_11 == nil {
					var hx_zero_12 *hxrt.HttpExchange
					return hx_zero_12
				}
				return hx_value_11.(*hxrt.HttpExchange)
			}(arg_0))
		})
	case *hxrt.Function[func(*haxe___Int64_____Int64)]:
		return hxrt.ProjectFunction(source, func(arg_0 int) {
			hxrt.FunctionCall(source)(hx_recover_class_haxe___Int64_____Int64(arg_0))
		})
	case *hxrt.Function[func(int) int]:
		return hxrt.ProjectFunction(source, func(arg_0 int) {
			hxrt.FunctionCall(source)(arg_0)
		})
	case *hxrt.Function[func(*haxe__io__Bytes) int]:
		return hxrt.ProjectFunction(source, func(arg_0 int) {
			hxrt.FunctionCall(source)(hx_recover_class_haxe__io__Bytes(arg_0))
		})
	case *hxrt.Function[func(int) float64]:
		return hxrt.ProjectFunction(source, func(arg_0 int) {
			hxrt.FunctionCall(source)(arg_0)
		})
	case *hxrt.Function[func(int) *haxe___Int64_____Int64]:
		return hxrt.ProjectFunction(source, func(arg_0 int) {
			hxrt.FunctionCall(source)(arg_0)
		})
	case *hxrt.Function[func(*string) any]:
		return hxrt.ProjectFunction(source, func(arg_0 int) {
			hxrt.FunctionCall(source)(func(hx_value_13 any) *string {
				if hx_value_13 == nil {
					var hx_zero_14 *string
					return hx_zero_14
				}
				return hx_value_13.(*string)
			}(arg_0))
		})
	case *hxrt.Function[func(*string) bool]:
		return hxrt.ProjectFunction(source, func(arg_0 int) {
			hxrt.FunctionCall(source)(func(hx_value_15 any) *string {
				if hx_value_15 == nil {
					var hx_zero_16 *string
					return hx_zero_16
				}
				return hx_value_15.(*string)
			}(arg_0))
		})
	case *hxrt.Function[func(any) any]:
		return hxrt.ProjectFunction(source, func(arg_0 int) {
			hxrt.FunctionCall(source)(arg_0)
		})
	case *hxrt.Function[func(any) bool]:
		return hxrt.ProjectFunction(source, func(arg_0 int) {
			hxrt.FunctionCall(source)(arg_0)
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
	case *haxe__exceptions__NotImplementedException:
		return source.haxe__exceptions__PosException.haxe__Exception
	case *haxe__Exception:
		return source
	case *haxe__ValueException:
		return source.haxe__Exception
	case *haxe__exceptions__PosException:
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

func hx_recover_class_haxe___Int64_____Int64(value any) *haxe___Int64_____Int64 {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *haxe___Int64_____Int64:
		return source
	default:
		return value.(*haxe___Int64_____Int64)
	}
}

func hx_recover_class_haxe__ds__StringMap(value any) *haxe__ds__StringMap {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *haxe__ds__StringMap:
		return source
	default:
		return value.(*haxe__ds__StringMap)
	}
}

func hx_recover_class_haxe__io__Bytes(value any) *haxe__io__Bytes {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *haxe__io__Bytes:
		return source
	default:
		return value.(*haxe__io__Bytes)
	}
}

func hx_recover_class_haxe__io__BytesBuffer(value any) *haxe__io__BytesBuffer {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *haxe__io__BytesBuffer:
		return source
	default:
		return value.(*haxe__io__BytesBuffer)
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

func hx_recover_class_haxe__io__Input(value any) *haxe__io__Input {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *sys__net__SocketInput:
		return source.haxe__io__Input
	case *haxe__io__Input:
		return source
	default:
		return value.(*haxe__io__Input)
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
	case *haxe__ValueException:
		return (source != nil)
	case *haxe__exceptions__PosException:
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

func hx_is_nominal_haxe__io__Eof(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *haxe__io__Eof:
		return (source != nil)
	default:
		return false
	}
}
