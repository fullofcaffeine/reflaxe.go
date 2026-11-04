package main

import "snapshot/hxrt"

func main() {
	http := New_sys__Http(hxrt.StringFromLiteral("data:text/plain,hello%20from%20haxe.go"))
	dataLog := hxrt.StringFromLiteral("")
	statusLog := -1
	byteCount := -1
	errLog := hxrt.StringFromLiteral("")
	http.onData = hxrt.NewFunction(func(data *string) {
		dataLog = data
	})
	http.onStatus = hxrt.NewFunction(func(status int) {
		statusLog = status
	})
	http.onBytes = hxrt.NewFunction(func(bytes *haxe__io__Bytes) {
		byteCount = bytes.length
	})
	http.onError = hxrt.NewFunction(func(msg *string) {
		errLog = msg
	})
	http.__hx_this.setHeader(hxrt.StringFromLiteral("X-Test"), hxrt.StringFromLiteral("one"))
	http.__hx_this.setHeader(hxrt.StringFromLiteral("X-Test"), hxrt.StringFromLiteral("two"))
	http.__hx_this.addHeader(hxrt.StringFromLiteral("X-Trace"), hxrt.StringFromLiteral("ok"))
	http.__hx_this.setParameter(hxrt.StringFromLiteral("a"), hxrt.StringFromLiteral("1"))
	http.__hx_this.addParameter(hxrt.StringFromLiteral("b"), hxrt.StringFromLiteral("2"))
	http.__hx_this.request(nil)
	hxrt.Println(any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("data="), dataLog)))
	hxrt.Println(any(hxrt.StringConcatAny(hxrt.StringFromLiteral("bytes="), byteCount)))
	hxrt.Println(any(hxrt.StringConcatAny(hxrt.StringFromLiteral("status="), statusLog)))
	var v any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("response="), http.__hx_this.get_responseData()))
	hxrt.Println(v)
	hxrt.Println(any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("error="), errLog)))
	post := New_sys__Http(hxrt.StringFromLiteral("data:text/plain,ignored"))
	post.__hx_this.setPostData(hxrt.StringFromLiteral("post-body"))
	postData := hxrt.StringFromLiteral("")
	post.onData = hxrt.NewFunction(func(data *string) {
		postData = data
	})
	post.__hx_this.request(true)
	hxrt.Println(any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("post="), postData)))
	bad := New_sys__Http(hxrt.StringFromLiteral("://bad"))
	badErr := hxrt.StringFromLiteral("")
	bad.onError = hxrt.NewFunction(func(msg *string) {
		badErr = msg
	})
	bad.__hx_this.request(nil)
	hxrt.Println(any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("bad="), badErr)))
	var v_1 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("direct="), sys__Http_requestUrl(hxrt.StringFromLiteral("data:text/plain,direct%20ok"))))
	hxrt.Println(v_1)
}

func hx_structural_recover(value any) hxrt.Anonymous {
	switch source := value.(type) {
	case *haxe___Int64_____Int64:
		return func(hx_structural_source_429 *haxe___Int64_____Int64) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_429, func(name string) hxrt.AnonymousField {
				switch name {
				case "high":
					return hxrt.AnonymousField{Value: hx_structural_source_429.high, Found: true}
				case "low":
					return hxrt.AnonymousField{Value: hx_structural_source_429.low, Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "high":
					hx_structural_source_429.high = func(hx_value_430 any) int {
						if hx_value_430 == nil {
							var hx_zero_431 int
							return hx_zero_431
						}
						return hx_value_430.(int)
					}(incoming)
					return true
				case "low":
					hx_structural_source_429.low = func(hx_value_432 any) int {
						if hx_value_432 == nil {
							var hx_zero_433 int
							return hx_zero_433
						}
						return hx_value_432.(int)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *sys__net__SocketOutput:
		return func(hx_structural_source_434 *sys__net__SocketOutput) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_434, func(name string) hxrt.AnonymousField {
				switch name {
				case "handle":
					return hxrt.AnonymousField{Value: hx_structural_source_434.handle, Found: true}
				case "writeByte":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_437 := hx_structural_source_434
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_437.__hx_this.writeByte(value)
						}, func() any {
							return hx_receiver_437.__hx_this
						}, "writeByte")
					}(), Found: true}
				case "writeBytes":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes, int, int) int] {
						hx_receiver_438 := hx_structural_source_434
						return hxrt.BoundFunction(func(bytes *haxe__io__Bytes, pos int, length int) int {
							return hx_receiver_438.__hx_this.writeBytes(bytes, pos, length)
						}, func() any {
							return hx_receiver_438.__hx_this
						}, "writeBytes")
					}(), Found: true}
				case "flush":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func()] {
						hx_receiver_439 := hx_structural_source_434
						return hxrt.BoundFunction(func() {
							hx_receiver_439.__hx_this.flush()
						}, func() any {
							return hx_receiver_439.__hx_this
						}, "flush")
					}(), Found: true}
				case "close":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func()] {
						hx_receiver_440 := hx_structural_source_434
						return hxrt.BoundFunction(func() {
							hx_receiver_440.__hx_this.close()
						}, func() any {
							return hx_receiver_440.__hx_this
						}, "close")
					}(), Found: true}
				case "bigEndian":
					return hxrt.AnonymousField{Value: hx_structural_source_434.bigEndian, Found: true}
				case "set_bigEndian":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(bool) bool] {
						hx_receiver_441 := hx_structural_source_434
						return hxrt.BoundFunction(func(value bool) bool {
							return hx_receiver_441.__hx_this.set_bigEndian(value)
						}, func() any {
							return hx_receiver_441.__hx_this
						}, "set_bigEndian")
					}(), Found: true}
				case "write":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes)] {
						hx_receiver_442 := hx_structural_source_434
						return hxrt.BoundFunction(func(bytes *haxe__io__Bytes) {
							hx_receiver_442.__hx_this.write(bytes)
						}, func() any {
							return hx_receiver_442.__hx_this
						}, "write")
					}(), Found: true}
				case "writeFullBytes":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes, int, int)] {
						hx_receiver_443 := hx_structural_source_434
						return hxrt.BoundFunction(func(bytes *haxe__io__Bytes, pos int, len int) {
							hx_receiver_443.__hx_this.writeFullBytes(bytes, pos, len)
						}, func() any {
							return hx_receiver_443.__hx_this
						}, "writeFullBytes")
					}(), Found: true}
				case "writeFloat":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(float64)] {
						hx_receiver_444 := hx_structural_source_434
						return hxrt.BoundFunction(func(value float64) {
							hx_receiver_444.__hx_this.writeFloat(value)
						}, func() any {
							return hx_receiver_444.__hx_this
						}, "writeFloat")
					}(), Found: true}
				case "writeDouble":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(float64)] {
						hx_receiver_445 := hx_structural_source_434
						return hxrt.BoundFunction(func(value float64) {
							hx_receiver_445.__hx_this.writeDouble(value)
						}, func() any {
							return hx_receiver_445.__hx_this
						}, "writeDouble")
					}(), Found: true}
				case "writeInt8":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_446 := hx_structural_source_434
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_446.__hx_this.writeInt8(value)
						}, func() any {
							return hx_receiver_446.__hx_this
						}, "writeInt8")
					}(), Found: true}
				case "writeInt16":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_447 := hx_structural_source_434
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_447.__hx_this.writeInt16(value)
						}, func() any {
							return hx_receiver_447.__hx_this
						}, "writeInt16")
					}(), Found: true}
				case "writeUInt16":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_448 := hx_structural_source_434
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_448.__hx_this.writeUInt16(value)
						}, func() any {
							return hx_receiver_448.__hx_this
						}, "writeUInt16")
					}(), Found: true}
				case "writeInt24":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_449 := hx_structural_source_434
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_449.__hx_this.writeInt24(value)
						}, func() any {
							return hx_receiver_449.__hx_this
						}, "writeInt24")
					}(), Found: true}
				case "writeUInt24":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_450 := hx_structural_source_434
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_450.__hx_this.writeUInt24(value)
						}, func() any {
							return hx_receiver_450.__hx_this
						}, "writeUInt24")
					}(), Found: true}
				case "writeInt32":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_451 := hx_structural_source_434
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_451.__hx_this.writeInt32(value)
						}, func() any {
							return hx_receiver_451.__hx_this
						}, "writeInt32")
					}(), Found: true}
				case "prepare":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_452 := hx_structural_source_434
						return hxrt.BoundFunction(func(nbytes int) {
							hx_receiver_452.__hx_this.prepare(nbytes)
						}, func() any {
							return hx_receiver_452.__hx_this
						}, "prepare")
					}(), Found: true}
				case "writeInput":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Input, any)] {
						hx_receiver_453 := hx_structural_source_434
						return hxrt.BoundFunction(func(input *haxe__io__Input, bufsize any) {
							hx_receiver_453.__hx_this.writeInput(input, bufsize)
						}, func() any {
							return hx_receiver_453.__hx_this
						}, "writeInput")
					}(), Found: true}
				case "writeString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string, *haxe__io__Encoding)] {
						hx_receiver_454 := hx_structural_source_434
						return hxrt.BoundFunction(func(value *string, encoding *haxe__io__Encoding) {
							hx_receiver_454.__hx_this.writeString(value, encoding)
						}, func() any {
							return hx_receiver_454.__hx_this
						}, "writeString")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "handle":
					hx_structural_source_434.handle = func(hx_value_435 any) *hxrt.SocketHandle {
						if hx_value_435 == nil {
							var hx_zero_436 *hxrt.SocketHandle
							return hx_zero_436
						}
						return hx_value_435.(*hxrt.SocketHandle)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *sys__net__SocketInput:
		return func(hx_structural_source_455 *sys__net__SocketInput) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_455, func(name string) hxrt.AnonymousField {
				switch name {
				case "handle":
					return hxrt.AnonymousField{Value: hx_structural_source_455.handle, Found: true}
				case "readByte":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_458 := hx_structural_source_455
						return hxrt.BoundFunction(func() int {
							return hx_receiver_458.__hx_this.readByte()
						}, func() any {
							return hx_receiver_458.__hx_this
						}, "readByte")
					}(), Found: true}
				case "readBytes":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes, int, int) int] {
						hx_receiver_459 := hx_structural_source_455
						return hxrt.BoundFunction(func(bytes *haxe__io__Bytes, pos int, length int) int {
							return hx_receiver_459.__hx_this.readBytes(bytes, pos, length)
						}, func() any {
							return hx_receiver_459.__hx_this
						}, "readBytes")
					}(), Found: true}
				case "close":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func()] {
						hx_receiver_460 := hx_structural_source_455
						return hxrt.BoundFunction(func() {
							hx_receiver_460.__hx_this.close()
						}, func() any {
							return hx_receiver_460.__hx_this
						}, "close")
					}(), Found: true}
				case "bigEndian":
					return hxrt.AnonymousField{Value: hx_structural_source_455.bigEndian, Found: true}
				case "set_bigEndian":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(bool) bool] {
						hx_receiver_461 := hx_structural_source_455
						return hxrt.BoundFunction(func(value bool) bool {
							return hx_receiver_461.__hx_this.set_bigEndian(value)
						}, func() any {
							return hx_receiver_461.__hx_this
						}, "set_bigEndian")
					}(), Found: true}
				case "readAll":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(any) *haxe__io__Bytes] {
						hx_receiver_462 := hx_structural_source_455
						return hxrt.BoundFunction(func(bufsize any) *haxe__io__Bytes {
							return hx_receiver_462.__hx_this.readAll(bufsize)
						}, func() any {
							return hx_receiver_462.__hx_this
						}, "readAll")
					}(), Found: true}
				case "readFullBytes":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes, int, int)] {
						hx_receiver_463 := hx_structural_source_455
						return hxrt.BoundFunction(func(bytes *haxe__io__Bytes, pos int, len int) {
							hx_receiver_463.__hx_this.readFullBytes(bytes, pos, len)
						}, func() any {
							return hx_receiver_463.__hx_this
						}, "readFullBytes")
					}(), Found: true}
				case "read":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int) *haxe__io__Bytes] {
						hx_receiver_464 := hx_structural_source_455
						return hxrt.BoundFunction(func(nbytes int) *haxe__io__Bytes {
							return hx_receiver_464.__hx_this.read(nbytes)
						}, func() any {
							return hx_receiver_464.__hx_this
						}, "read")
					}(), Found: true}
				case "readUntil":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int) *string] {
						hx_receiver_465 := hx_structural_source_455
						return hxrt.BoundFunction(func(end int) *string {
							return hx_receiver_465.__hx_this.readUntil(end)
						}, func() any {
							return hx_receiver_465.__hx_this
						}, "readUntil")
					}(), Found: true}
				case "readLine":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_466 := hx_structural_source_455
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_466.__hx_this.readLine()
						}, func() any {
							return hx_receiver_466.__hx_this
						}, "readLine")
					}(), Found: true}
				case "readFloat":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() float64] {
						hx_receiver_467 := hx_structural_source_455
						return hxrt.BoundFunction(func() float64 {
							return hx_receiver_467.__hx_this.readFloat()
						}, func() any {
							return hx_receiver_467.__hx_this
						}, "readFloat")
					}(), Found: true}
				case "readDouble":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() float64] {
						hx_receiver_468 := hx_structural_source_455
						return hxrt.BoundFunction(func() float64 {
							return hx_receiver_468.__hx_this.readDouble()
						}, func() any {
							return hx_receiver_468.__hx_this
						}, "readDouble")
					}(), Found: true}
				case "readInt8":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_469 := hx_structural_source_455
						return hxrt.BoundFunction(func() int {
							return hx_receiver_469.__hx_this.readInt8()
						}, func() any {
							return hx_receiver_469.__hx_this
						}, "readInt8")
					}(), Found: true}
				case "readInt16":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_470 := hx_structural_source_455
						return hxrt.BoundFunction(func() int {
							return hx_receiver_470.__hx_this.readInt16()
						}, func() any {
							return hx_receiver_470.__hx_this
						}, "readInt16")
					}(), Found: true}
				case "readUInt16":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_471 := hx_structural_source_455
						return hxrt.BoundFunction(func() int {
							return hx_receiver_471.__hx_this.readUInt16()
						}, func() any {
							return hx_receiver_471.__hx_this
						}, "readUInt16")
					}(), Found: true}
				case "readInt24":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_472 := hx_structural_source_455
						return hxrt.BoundFunction(func() int {
							return hx_receiver_472.__hx_this.readInt24()
						}, func() any {
							return hx_receiver_472.__hx_this
						}, "readInt24")
					}(), Found: true}
				case "readUInt24":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_473 := hx_structural_source_455
						return hxrt.BoundFunction(func() int {
							return hx_receiver_473.__hx_this.readUInt24()
						}, func() any {
							return hx_receiver_473.__hx_this
						}, "readUInt24")
					}(), Found: true}
				case "readInt32":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_474 := hx_structural_source_455
						return hxrt.BoundFunction(func() int {
							return hx_receiver_474.__hx_this.readInt32()
						}, func() any {
							return hx_receiver_474.__hx_this
						}, "readInt32")
					}(), Found: true}
				case "readString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, *haxe__io__Encoding) *string] {
						hx_receiver_475 := hx_structural_source_455
						return hxrt.BoundFunction(func(len int, encoding *haxe__io__Encoding) *string {
							return hx_receiver_475.__hx_this.readString(len, encoding)
						}, func() any {
							return hx_receiver_475.__hx_this
						}, "readString")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "handle":
					hx_structural_source_455.handle = func(hx_value_456 any) *hxrt.SocketHandle {
						if hx_value_456 == nil {
							var hx_zero_457 *hxrt.SocketHandle
							return hx_zero_457
						}
						return hx_value_456.(*hxrt.SocketHandle)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *sys__net__Socket:
		return func(hx_structural_source_476 *sys__net__Socket) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_476, func(name string) hxrt.AnonymousField {
				switch name {
				case "input":
					return hxrt.AnonymousField{Value: hx_structural_source_476.input, Found: true}
				case "output":
					return hxrt.AnonymousField{Value: hx_structural_source_476.output, Found: true}
				case "custom":
					return hxrt.AnonymousField{Value: hx_structural_source_476.custom, Found: true}
				case "handle":
					return hxrt.AnonymousField{Value: hx_structural_source_476.handle, Found: true}
				case "close":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func()] {
						hx_receiver_479 := hx_structural_source_476
						return hxrt.BoundFunction(func() {
							hx_receiver_479.__hx_this.close()
						}, func() any {
							return hx_receiver_479.__hx_this
						}, "close")
					}(), Found: true}
				case "read":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_480 := hx_structural_source_476
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_480.__hx_this.read()
						}, func() any {
							return hx_receiver_480.__hx_this
						}, "read")
					}(), Found: true}
				case "write":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string)] {
						hx_receiver_481 := hx_structural_source_476
						return hxrt.BoundFunction(func(content *string) {
							hx_receiver_481.__hx_this.write(content)
						}, func() any {
							return hx_receiver_481.__hx_this
						}, "write")
					}(), Found: true}
				case "connect":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*sys__net__Host, int)] {
						hx_receiver_482 := hx_structural_source_476
						return hxrt.BoundFunction(func(host *sys__net__Host, port int) {
							hx_receiver_482.__hx_this.connect(host, port)
						}, func() any {
							return hx_receiver_482.__hx_this
						}, "connect")
					}(), Found: true}
				case "listen":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_483 := hx_structural_source_476
						return hxrt.BoundFunction(func(connections int) {
							hx_receiver_483.__hx_this.listen(connections)
						}, func() any {
							return hx_receiver_483.__hx_this
						}, "listen")
					}(), Found: true}
				case "shutdown":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(bool, bool)] {
						hx_receiver_484 := hx_structural_source_476
						return hxrt.BoundFunction(func(read bool, write bool) {
							hx_receiver_484.__hx_this.shutdown(read, write)
						}, func() any {
							return hx_receiver_484.__hx_this
						}, "shutdown")
					}(), Found: true}
				case "bind":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*sys__net__Host, int)] {
						hx_receiver_485 := hx_structural_source_476
						return hxrt.BoundFunction(func(host *sys__net__Host, port int) {
							hx_receiver_485.__hx_this.bind(host, port)
						}, func() any {
							return hx_receiver_485.__hx_this
						}, "bind")
					}(), Found: true}
				case "accept":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *sys__net__Socket] {
						hx_receiver_486 := hx_structural_source_476
						return hxrt.BoundFunction(func() *sys__net__Socket {
							return hx_receiver_486.__hx_this.accept()
						}, func() any {
							return hx_receiver_486.__hx_this
						}, "accept")
					}(), Found: true}
				case "peer":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() hxrt.Anonymous] {
						hx_receiver_487 := hx_structural_source_476
						return hxrt.BoundFunction(func() hxrt.Anonymous {
							return hx_receiver_487.__hx_this.peer()
						}, func() any {
							return hx_receiver_487.__hx_this
						}, "peer")
					}(), Found: true}
				case "host":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() hxrt.Anonymous] {
						hx_receiver_488 := hx_structural_source_476
						return hxrt.BoundFunction(func() hxrt.Anonymous {
							return hx_receiver_488.__hx_this.host()
						}, func() any {
							return hx_receiver_488.__hx_this
						}, "host")
					}(), Found: true}
				case "setTimeout":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(float64)] {
						hx_receiver_489 := hx_structural_source_476
						return hxrt.BoundFunction(func(timeout float64) {
							hx_receiver_489.__hx_this.setTimeout(timeout)
						}, func() any {
							return hx_receiver_489.__hx_this
						}, "setTimeout")
					}(), Found: true}
				case "waitForRead":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func()] {
						hx_receiver_490 := hx_structural_source_476
						return hxrt.BoundFunction(func() {
							hx_receiver_490.__hx_this.waitForRead()
						}, func() any {
							return hx_receiver_490.__hx_this
						}, "waitForRead")
					}(), Found: true}
				case "setBlocking":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(bool)] {
						hx_receiver_491 := hx_structural_source_476
						return hxrt.BoundFunction(func(blocking bool) {
							hx_receiver_491.__hx_this.setBlocking(blocking)
						}, func() any {
							return hx_receiver_491.__hx_this
						}, "setBlocking")
					}(), Found: true}
				case "setFastSend":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(bool)] {
						hx_receiver_492 := hx_structural_source_476
						return hxrt.BoundFunction(func(fastSend bool) {
							hx_receiver_492.__hx_this.setFastSend(fastSend)
						}, func() any {
							return hx_receiver_492.__hx_this
						}, "setFastSend")
					}(), Found: true}
				case "replaceHandle":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*hxrt.SocketHandle)] {
						hx_receiver_493 := hx_structural_source_476
						return hxrt.BoundFunction(func(next *hxrt.SocketHandle) {
							hx_receiver_493.__hx_this.replaceHandle(next)
						}, func() any {
							return hx_receiver_493.__hx_this
						}, "replaceHandle")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "custom":
					hx_structural_source_476.custom = incoming
					return true
				case "handle":
					hx_structural_source_476.handle = func(hx_value_477 any) *hxrt.SocketHandle {
						if hx_value_477 == nil {
							var hx_zero_478 *hxrt.SocketHandle
							return hx_zero_478
						}
						return hx_value_477.(*hxrt.SocketHandle)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *sys__net__Host:
		return func(hx_structural_source_494 *sys__net__Host) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_494, func(name string) hxrt.AnonymousField {
				switch name {
				case "host":
					return hxrt.AnonymousField{Value: hx_structural_source_494.host, Found: true}
				case "ip":
					return hxrt.AnonymousField{Value: hx_structural_source_494.ip, Found: true}
				case "toString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_495 := hx_structural_source_494
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_495.__hx_this.toString()
						}, func() any {
							return hx_receiver_495.__hx_this
						}, "toString")
					}(), Found: true}
				case "reverse":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_496 := hx_structural_source_494
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_496.__hx_this.reverse()
						}, func() any {
							return hx_receiver_496.__hx_this
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
		return func(hx_structural_source_497 *sys__Http) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_497, func(name string) hxrt.AnonymousField {
				switch name {
				case "noShutdown":
					return hxrt.AnonymousField{Value: hx_structural_source_497.noShutdown, Found: true}
				case "cnxTimeout":
					return hxrt.AnonymousField{Value: hx_structural_source_497.cnxTimeout, Found: true}
				case "responseHeaders":
					return hxrt.AnonymousField{Value: hx_structural_source_497.responseHeaders, Found: true}
				case "responseHeadersSameKey":
					return hxrt.AnonymousField{Value: hx_structural_source_497.responseHeadersSameKey, Found: true}
				case "file":
					return hxrt.AnonymousField{Value: hx_structural_source_497.file, Found: true}
				case "request":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(any)] {
						hx_receiver_504 := hx_structural_source_497
						return hxrt.BoundFunction(func(post any) {
							hx_receiver_504.__hx_this.request(post)
						}, func() any {
							return hx_receiver_504.__hx_this
						}, "request")
					}(), Found: true}
				case "fileTransfert":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string, *string, *haxe__io__Input, int, *string)] {
						hx_receiver_505 := hx_structural_source_497
						return hxrt.BoundFunction(func(argname *string, filename *string, file *haxe__io__Input, size int, mimeType *string) {
							if mimeType == nil {
								mimeType = hxrt.StringFromLiteral("application/octet-stream")
							}
							hx_receiver_505.__hx_this.fileTransfert(argname, filename, file, size, mimeType)
						}, func() any {
							return hx_receiver_505.__hx_this
						}, "fileTransfert")
					}(), Found: true}
				case "fileTransfer":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string, *string, *haxe__io__Input, int, *string)] {
						hx_receiver_506 := hx_structural_source_497
						return hxrt.BoundFunction(func(argname *string, filename *string, file *haxe__io__Input, size int, mimeType *string) {
							if mimeType == nil {
								mimeType = hxrt.StringFromLiteral("application/octet-stream")
							}
							hx_receiver_506.__hx_this.fileTransfer(argname, filename, file, size, mimeType)
						}, func() any {
							return hx_receiver_506.__hx_this
						}, "fileTransfer")
					}(), Found: true}
				case "customRequest":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(bool, *haxe__io__Output, *sys__net__Socket, *string)] {
						hx_receiver_507 := hx_structural_source_497
						return hxrt.BoundFunction(func(post bool, api *haxe__io__Output, sock *sys__net__Socket, method *string) {
							hx_receiver_507.__hx_this.customRequest(post, api, sock, method)
						}, func() any {
							return hx_receiver_507.__hx_this
						}, "customRequest")
					}(), Found: true}
				case "getResponseHeaderValues":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string) *hxrt.Array] {
						hx_receiver_508 := hx_structural_source_497
						return hxrt.BoundFunction(func(key *string) *hxrt.Array {
							return hx_receiver_508.__hx_this.getResponseHeaderValues(key)
						}, func() any {
							return hx_receiver_508.__hx_this
						}, "getResponseHeaderValues")
					}(), Found: true}
				case "requestWith":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(bool, *haxe__io__Output, *sys__net__Socket, *string)] {
						hx_receiver_509 := hx_structural_source_497
						return hxrt.BoundFunction(func(post bool, api *haxe__io__Output, sock *sys__net__Socket, method *string) {
							hx_receiver_509.__hx_this.requestWith(post, api, sock, method)
						}, func() any {
							return hx_receiver_509.__hx_this
						}, "requestWith")
					}(), Found: true}
				case "pumpUpload":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*hxrt.HttpExchange, hxrt.Anonymous) hxrt.Anonymous] {
						hx_receiver_510 := hx_structural_source_497
						return hxrt.BoundFunction(func(exchange *hxrt.HttpExchange, upload hxrt.Anonymous) hxrt.Anonymous {
							return hx_receiver_510.__hx_this.pumpUpload(exchange, upload)
						}, func() any {
							return hx_receiver_510.__hx_this
						}, "pumpUpload")
					}(), Found: true}
				case "handleDataRequest":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(bool, *haxe__io__Output, *string)] {
						hx_receiver_511 := hx_structural_source_497
						return hxrt.BoundFunction(func(post bool, api *haxe__io__Output, method *string) {
							hx_receiver_511.__hx_this.handleDataRequest(post, api, method)
						}, func() any {
							return hx_receiver_511.__hx_this
						}, "handleDataRequest")
					}(), Found: true}
				case "recordResponseHeaders":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*hxrt.HttpExchange)] {
						hx_receiver_512 := hx_structural_source_497
						return hxrt.BoundFunction(func(exchange *hxrt.HttpExchange) {
							hx_receiver_512.__hx_this.recordResponseHeaders(exchange)
						}, func() any {
							return hx_receiver_512.__hx_this
						}, "recordResponseHeaders")
					}(), Found: true}
				case "resetResponseHeaders":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func()] {
						hx_receiver_513 := hx_structural_source_497
						return hxrt.BoundFunction(func() {
							hx_receiver_513.__hx_this.resetResponseHeaders()
						}, func() any {
							return hx_receiver_513.__hx_this
						}, "resetResponseHeaders")
					}(), Found: true}
				case "encodedParameters":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_514 := hx_structural_source_497
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_514.__hx_this.encodedParameters()
						}, func() any {
							return hx_receiver_514.__hx_this
						}, "encodedParameters")
					}(), Found: true}
				case "url":
					return hxrt.AnonymousField{Value: hx_structural_source_497.url, Found: true}
				case "responseBytes":
					return hxrt.AnonymousField{Value: hx_structural_source_497.responseBytes, Found: true}
				case "responseAsString":
					return hxrt.AnonymousField{Value: hx_structural_source_497.responseAsString, Found: true}
				case "postData":
					return hxrt.AnonymousField{Value: hx_structural_source_497.postData, Found: true}
				case "postBytes":
					return hxrt.AnonymousField{Value: hx_structural_source_497.postBytes, Found: true}
				case "headers":
					return hxrt.AnonymousField{Value: hx_structural_source_497.headers, Found: true}
				case "params":
					return hxrt.AnonymousField{Value: hx_structural_source_497.params, Found: true}
				case "emptyOnData":
					return hxrt.AnonymousField{Value: hx_structural_source_497.emptyOnData, Found: true}
				case "setHeader":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string, *string)] {
						hx_receiver_525 := hx_structural_source_497
						return hxrt.BoundFunction(func(name *string, value *string) {
							hx_receiver_525.__hx_this.setHeader(name, value)
						}, func() any {
							return hx_receiver_525.__hx_this
						}, "setHeader")
					}(), Found: true}
				case "addHeader":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string, *string)] {
						hx_receiver_526 := hx_structural_source_497
						return hxrt.BoundFunction(func(header *string, value *string) {
							hx_receiver_526.__hx_this.addHeader(header, value)
						}, func() any {
							return hx_receiver_526.__hx_this
						}, "addHeader")
					}(), Found: true}
				case "setParameter":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string, *string)] {
						hx_receiver_527 := hx_structural_source_497
						return hxrt.BoundFunction(func(name *string, value *string) {
							hx_receiver_527.__hx_this.setParameter(name, value)
						}, func() any {
							return hx_receiver_527.__hx_this
						}, "setParameter")
					}(), Found: true}
				case "addParameter":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string, *string)] {
						hx_receiver_528 := hx_structural_source_497
						return hxrt.BoundFunction(func(name *string, value *string) {
							hx_receiver_528.__hx_this.addParameter(name, value)
						}, func() any {
							return hx_receiver_528.__hx_this
						}, "addParameter")
					}(), Found: true}
				case "setPostData":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string)] {
						hx_receiver_529 := hx_structural_source_497
						return hxrt.BoundFunction(func(data *string) {
							hx_receiver_529.__hx_this.setPostData(data)
						}, func() any {
							return hx_receiver_529.__hx_this
						}, "setPostData")
					}(), Found: true}
				case "setPostBytes":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes)] {
						hx_receiver_530 := hx_structural_source_497
						return hxrt.BoundFunction(func(data *haxe__io__Bytes) {
							hx_receiver_530.__hx_this.setPostBytes(data)
						}, func() any {
							return hx_receiver_530.__hx_this
						}, "setPostBytes")
					}(), Found: true}
				case "onData":
					return hxrt.AnonymousField{Value: hxrt.ReadDynamicFunction(hx_structural_source_497.onData), Found: true}
				case "onBytes":
					return hxrt.AnonymousField{Value: hxrt.ReadDynamicFunction(hx_structural_source_497.onBytes), Found: true}
				case "onError":
					return hxrt.AnonymousField{Value: hxrt.ReadDynamicFunction(hx_structural_source_497.onError), Found: true}
				case "onStatus":
					return hxrt.AnonymousField{Value: hxrt.ReadDynamicFunction(hx_structural_source_497.onStatus), Found: true}
				case "hasOnData":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() bool] {
						hx_receiver_531 := hx_structural_source_497
						return hxrt.BoundFunction(func() bool {
							return hx_receiver_531.__hx_this.hasOnData()
						}, func() any {
							return hx_receiver_531.__hx_this
						}, "hasOnData")
					}(), Found: true}
				case "success":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes)] {
						hx_receiver_532 := hx_structural_source_497
						return hxrt.BoundFunction(func(data *haxe__io__Bytes) {
							hx_receiver_532.__hx_this.success(data)
						}, func() any {
							return hx_receiver_532.__hx_this
						}, "success")
					}(), Found: true}
				case "get_responseData":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_533 := hx_structural_source_497
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_533.__hx_this.get_responseData()
						}, func() any {
							return hx_receiver_533.__hx_this
						}, "get_responseData")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "noShutdown":
					hx_structural_source_497.noShutdown = func(hx_value_498 any) bool {
						if hx_value_498 == nil {
							var hx_zero_499 bool
							return hx_zero_499
						}
						return hx_value_498.(bool)
					}(incoming)
					return true
				case "cnxTimeout":
					hx_structural_source_497.cnxTimeout = func(hx_value_500 any) float64 {
						if hx_value_500 == nil {
							var hx_zero_501 float64
							return hx_zero_501
						}
						return hx_value_500.(float64)
					}(incoming)
					return true
				case "responseHeaders":
					hx_structural_source_497.responseHeaders = func(hx_value_502 any) *haxe__ds__StringMap {
						if hx_value_502 == nil {
							var hx_zero_503 *haxe__ds__StringMap
							return hx_zero_503
						}
						return hx_value_502.(*haxe__ds__StringMap)
					}(incoming)
					return true
				case "responseHeadersSameKey":
					hx_structural_source_497.responseHeadersSameKey = hx_recover_class_4(incoming)
					return true
				case "file":
					hx_structural_source_497.file = hx_structural_recover(incoming)
					return true
				case "url":
					hx_structural_source_497.url = func(hx_value_515 any) *string {
						if hx_value_515 == nil {
							var hx_zero_516 *string
							return hx_zero_516
						}
						return hx_value_515.(*string)
					}(incoming)
					return true
				case "responseAsString":
					hx_structural_source_497.responseAsString = func(hx_value_517 any) *string {
						if hx_value_517 == nil {
							var hx_zero_518 *string
							return hx_zero_518
						}
						return hx_value_517.(*string)
					}(incoming)
					return true
				case "postData":
					hx_structural_source_497.postData = func(hx_value_519 any) *string {
						if hx_value_519 == nil {
							var hx_zero_520 *string
							return hx_zero_520
						}
						return hx_value_519.(*string)
					}(incoming)
					return true
				case "postBytes":
					hx_structural_source_497.postBytes = hx_recover_class_6(incoming)
					return true
				case "headers":
					hx_structural_source_497.headers = func(hx_value_521 any) *hxrt.Array {
						if hx_value_521 == nil {
							var hx_zero_522 *hxrt.Array
							return hx_zero_522
						}
						return hx_value_521.(*hxrt.Array)
					}(incoming)
					return true
				case "params":
					hx_structural_source_497.params = func(hx_value_523 any) *hxrt.Array {
						if hx_value_523 == nil {
							var hx_zero_524 *hxrt.Array
							return hx_zero_524
						}
						return hx_value_523.(*hxrt.Array)
					}(incoming)
					return true
				case "onData":
					hx_structural_source_497.onData = hx_recover_callable_3(incoming)
					return true
				case "onBytes":
					hx_structural_source_497.onBytes = hx_recover_callable_4(incoming)
					return true
				case "onError":
					hx_structural_source_497.onError = hx_recover_callable_3(incoming)
					return true
				case "onStatus":
					hx_structural_source_497.onStatus = hx_recover_callable_5(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe__io__Output:
		return func(hx_structural_source_534 *haxe__io__Output) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_534, func(name string) hxrt.AnonymousField {
				switch name {
				case "bigEndian":
					return hxrt.AnonymousField{Value: hx_structural_source_534.bigEndian, Found: true}
				case "writeByte":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_535 := hx_structural_source_534
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_535.__hx_this.writeByte(value)
						}, func() any {
							return hx_receiver_535.__hx_this
						}, "writeByte")
					}(), Found: true}
				case "writeBytes":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes, int, int) int] {
						hx_receiver_536 := hx_structural_source_534
						return hxrt.BoundFunction(func(bytes *haxe__io__Bytes, pos int, len int) int {
							return hx_receiver_536.__hx_this.writeBytes(bytes, pos, len)
						}, func() any {
							return hx_receiver_536.__hx_this
						}, "writeBytes")
					}(), Found: true}
				case "flush":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func()] {
						hx_receiver_537 := hx_structural_source_534
						return hxrt.BoundFunction(func() {
							hx_receiver_537.__hx_this.flush()
						}, func() any {
							return hx_receiver_537.__hx_this
						}, "flush")
					}(), Found: true}
				case "close":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func()] {
						hx_receiver_538 := hx_structural_source_534
						return hxrt.BoundFunction(func() {
							hx_receiver_538.__hx_this.close()
						}, func() any {
							return hx_receiver_538.__hx_this
						}, "close")
					}(), Found: true}
				case "set_bigEndian":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(bool) bool] {
						hx_receiver_539 := hx_structural_source_534
						return hxrt.BoundFunction(func(value bool) bool {
							return hx_receiver_539.__hx_this.set_bigEndian(value)
						}, func() any {
							return hx_receiver_539.__hx_this
						}, "set_bigEndian")
					}(), Found: true}
				case "write":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes)] {
						hx_receiver_540 := hx_structural_source_534
						return hxrt.BoundFunction(func(bytes *haxe__io__Bytes) {
							hx_receiver_540.__hx_this.write(bytes)
						}, func() any {
							return hx_receiver_540.__hx_this
						}, "write")
					}(), Found: true}
				case "writeFullBytes":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes, int, int)] {
						hx_receiver_541 := hx_structural_source_534
						return hxrt.BoundFunction(func(bytes *haxe__io__Bytes, pos int, len int) {
							hx_receiver_541.__hx_this.writeFullBytes(bytes, pos, len)
						}, func() any {
							return hx_receiver_541.__hx_this
						}, "writeFullBytes")
					}(), Found: true}
				case "writeFloat":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(float64)] {
						hx_receiver_542 := hx_structural_source_534
						return hxrt.BoundFunction(func(value float64) {
							hx_receiver_542.__hx_this.writeFloat(value)
						}, func() any {
							return hx_receiver_542.__hx_this
						}, "writeFloat")
					}(), Found: true}
				case "writeDouble":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(float64)] {
						hx_receiver_543 := hx_structural_source_534
						return hxrt.BoundFunction(func(value float64) {
							hx_receiver_543.__hx_this.writeDouble(value)
						}, func() any {
							return hx_receiver_543.__hx_this
						}, "writeDouble")
					}(), Found: true}
				case "writeInt8":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_544 := hx_structural_source_534
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_544.__hx_this.writeInt8(value)
						}, func() any {
							return hx_receiver_544.__hx_this
						}, "writeInt8")
					}(), Found: true}
				case "writeInt16":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_545 := hx_structural_source_534
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_545.__hx_this.writeInt16(value)
						}, func() any {
							return hx_receiver_545.__hx_this
						}, "writeInt16")
					}(), Found: true}
				case "writeUInt16":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_546 := hx_structural_source_534
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_546.__hx_this.writeUInt16(value)
						}, func() any {
							return hx_receiver_546.__hx_this
						}, "writeUInt16")
					}(), Found: true}
				case "writeInt24":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_547 := hx_structural_source_534
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_547.__hx_this.writeInt24(value)
						}, func() any {
							return hx_receiver_547.__hx_this
						}, "writeInt24")
					}(), Found: true}
				case "writeUInt24":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_548 := hx_structural_source_534
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_548.__hx_this.writeUInt24(value)
						}, func() any {
							return hx_receiver_548.__hx_this
						}, "writeUInt24")
					}(), Found: true}
				case "writeInt32":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_549 := hx_structural_source_534
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_549.__hx_this.writeInt32(value)
						}, func() any {
							return hx_receiver_549.__hx_this
						}, "writeInt32")
					}(), Found: true}
				case "prepare":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_550 := hx_structural_source_534
						return hxrt.BoundFunction(func(nbytes int) {
							hx_receiver_550.__hx_this.prepare(nbytes)
						}, func() any {
							return hx_receiver_550.__hx_this
						}, "prepare")
					}(), Found: true}
				case "writeInput":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Input, any)] {
						hx_receiver_551 := hx_structural_source_534
						return hxrt.BoundFunction(func(input *haxe__io__Input, bufsize any) {
							hx_receiver_551.__hx_this.writeInput(input, bufsize)
						}, func() any {
							return hx_receiver_551.__hx_this
						}, "writeInput")
					}(), Found: true}
				case "writeString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string, *haxe__io__Encoding)] {
						hx_receiver_552 := hx_structural_source_534
						return hxrt.BoundFunction(func(value *string, encoding *haxe__io__Encoding) {
							hx_receiver_552.__hx_this.writeString(value, encoding)
						}, func() any {
							return hx_receiver_552.__hx_this
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
		return func(hx_structural_source_553 *haxe__io__Input) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_553, func(name string) hxrt.AnonymousField {
				switch name {
				case "bigEndian":
					return hxrt.AnonymousField{Value: hx_structural_source_553.bigEndian, Found: true}
				case "readByte":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_554 := hx_structural_source_553
						return hxrt.BoundFunction(func() int {
							return hx_receiver_554.__hx_this.readByte()
						}, func() any {
							return hx_receiver_554.__hx_this
						}, "readByte")
					}(), Found: true}
				case "readBytes":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes, int, int) int] {
						hx_receiver_555 := hx_structural_source_553
						return hxrt.BoundFunction(func(bytes *haxe__io__Bytes, pos int, len int) int {
							return hx_receiver_555.__hx_this.readBytes(bytes, pos, len)
						}, func() any {
							return hx_receiver_555.__hx_this
						}, "readBytes")
					}(), Found: true}
				case "close":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func()] {
						hx_receiver_556 := hx_structural_source_553
						return hxrt.BoundFunction(func() {
							hx_receiver_556.__hx_this.close()
						}, func() any {
							return hx_receiver_556.__hx_this
						}, "close")
					}(), Found: true}
				case "set_bigEndian":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(bool) bool] {
						hx_receiver_557 := hx_structural_source_553
						return hxrt.BoundFunction(func(value bool) bool {
							return hx_receiver_557.__hx_this.set_bigEndian(value)
						}, func() any {
							return hx_receiver_557.__hx_this
						}, "set_bigEndian")
					}(), Found: true}
				case "readAll":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(any) *haxe__io__Bytes] {
						hx_receiver_558 := hx_structural_source_553
						return hxrt.BoundFunction(func(bufsize any) *haxe__io__Bytes {
							return hx_receiver_558.__hx_this.readAll(bufsize)
						}, func() any {
							return hx_receiver_558.__hx_this
						}, "readAll")
					}(), Found: true}
				case "readFullBytes":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes, int, int)] {
						hx_receiver_559 := hx_structural_source_553
						return hxrt.BoundFunction(func(bytes *haxe__io__Bytes, pos int, len int) {
							hx_receiver_559.__hx_this.readFullBytes(bytes, pos, len)
						}, func() any {
							return hx_receiver_559.__hx_this
						}, "readFullBytes")
					}(), Found: true}
				case "read":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int) *haxe__io__Bytes] {
						hx_receiver_560 := hx_structural_source_553
						return hxrt.BoundFunction(func(nbytes int) *haxe__io__Bytes {
							return hx_receiver_560.__hx_this.read(nbytes)
						}, func() any {
							return hx_receiver_560.__hx_this
						}, "read")
					}(), Found: true}
				case "readUntil":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int) *string] {
						hx_receiver_561 := hx_structural_source_553
						return hxrt.BoundFunction(func(end int) *string {
							return hx_receiver_561.__hx_this.readUntil(end)
						}, func() any {
							return hx_receiver_561.__hx_this
						}, "readUntil")
					}(), Found: true}
				case "readLine":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_562 := hx_structural_source_553
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_562.__hx_this.readLine()
						}, func() any {
							return hx_receiver_562.__hx_this
						}, "readLine")
					}(), Found: true}
				case "readFloat":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() float64] {
						hx_receiver_563 := hx_structural_source_553
						return hxrt.BoundFunction(func() float64 {
							return hx_receiver_563.__hx_this.readFloat()
						}, func() any {
							return hx_receiver_563.__hx_this
						}, "readFloat")
					}(), Found: true}
				case "readDouble":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() float64] {
						hx_receiver_564 := hx_structural_source_553
						return hxrt.BoundFunction(func() float64 {
							return hx_receiver_564.__hx_this.readDouble()
						}, func() any {
							return hx_receiver_564.__hx_this
						}, "readDouble")
					}(), Found: true}
				case "readInt8":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_565 := hx_structural_source_553
						return hxrt.BoundFunction(func() int {
							return hx_receiver_565.__hx_this.readInt8()
						}, func() any {
							return hx_receiver_565.__hx_this
						}, "readInt8")
					}(), Found: true}
				case "readInt16":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_566 := hx_structural_source_553
						return hxrt.BoundFunction(func() int {
							return hx_receiver_566.__hx_this.readInt16()
						}, func() any {
							return hx_receiver_566.__hx_this
						}, "readInt16")
					}(), Found: true}
				case "readUInt16":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_567 := hx_structural_source_553
						return hxrt.BoundFunction(func() int {
							return hx_receiver_567.__hx_this.readUInt16()
						}, func() any {
							return hx_receiver_567.__hx_this
						}, "readUInt16")
					}(), Found: true}
				case "readInt24":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_568 := hx_structural_source_553
						return hxrt.BoundFunction(func() int {
							return hx_receiver_568.__hx_this.readInt24()
						}, func() any {
							return hx_receiver_568.__hx_this
						}, "readInt24")
					}(), Found: true}
				case "readUInt24":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_569 := hx_structural_source_553
						return hxrt.BoundFunction(func() int {
							return hx_receiver_569.__hx_this.readUInt24()
						}, func() any {
							return hx_receiver_569.__hx_this
						}, "readUInt24")
					}(), Found: true}
				case "readInt32":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_570 := hx_structural_source_553
						return hxrt.BoundFunction(func() int {
							return hx_receiver_570.__hx_this.readInt32()
						}, func() any {
							return hx_receiver_570.__hx_this
						}, "readInt32")
					}(), Found: true}
				case "readString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, *haxe__io__Encoding) *string] {
						hx_receiver_571 := hx_structural_source_553
						return hxrt.BoundFunction(func(len int, encoding *haxe__io__Encoding) *string {
							return hx_receiver_571.__hx_this.readString(len, encoding)
						}, func() any {
							return hx_receiver_571.__hx_this
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
	case *haxe__io__Bytes:
		return func(hx_structural_source_572 *haxe__io__Bytes) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_572, func(name string) hxrt.AnonymousField {
				switch name {
				case "length":
					return hxrt.AnonymousField{Value: hx_structural_source_572.length, Found: true}
				case "b":
					return hxrt.AnonymousField{Value: hx_structural_source_572.b, Found: true}
				case "__hx_raw":
					return hxrt.AnonymousField{Value: hx_structural_source_572.__hx_raw, Found: true}
				case "__hx_rawValid":
					return hxrt.AnonymousField{Value: hx_structural_source_572.__hx_rawValid, Found: true}
				case "__hx_dataExposed":
					return hxrt.AnonymousField{Value: hx_structural_source_572.__hx_dataExposed, Found: true}
				case "get":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int) int] {
						hx_receiver_581 := hx_structural_source_572
						return hxrt.BoundFunction(func(pos int) int {
							return hx_receiver_581.__hx_this.get(pos)
						}, func() any {
							return hx_receiver_581.__hx_this
						}, "get")
					}(), Found: true}
				case "set":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, int)] {
						hx_receiver_582 := hx_structural_source_572
						return hxrt.BoundFunction(func(pos int, value int) {
							hx_receiver_582.__hx_this.set(pos, value)
						}, func() any {
							return hx_receiver_582.__hx_this
						}, "set")
					}(), Found: true}
				case "blit":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, *haxe__io__Bytes, int, int)] {
						hx_receiver_583 := hx_structural_source_572
						return hxrt.BoundFunction(func(pos int, src *haxe__io__Bytes, srcpos int, len int) {
							hx_receiver_583.__hx_this.blit(pos, src, srcpos, len)
						}, func() any {
							return hx_receiver_583.__hx_this
						}, "blit")
					}(), Found: true}
				case "fill":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, int, int)] {
						hx_receiver_584 := hx_structural_source_572
						return hxrt.BoundFunction(func(pos int, len int, value int) {
							hx_receiver_584.__hx_this.fill(pos, len, value)
						}, func() any {
							return hx_receiver_584.__hx_this
						}, "fill")
					}(), Found: true}
				case "sub":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, int) *haxe__io__Bytes] {
						hx_receiver_585 := hx_structural_source_572
						return hxrt.BoundFunction(func(pos int, len int) *haxe__io__Bytes {
							return hx_receiver_585.__hx_this.sub(pos, len)
						}, func() any {
							return hx_receiver_585.__hx_this
						}, "sub")
					}(), Found: true}
				case "compare":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes) int] {
						hx_receiver_586 := hx_structural_source_572
						return hxrt.BoundFunction(func(other *haxe__io__Bytes) int {
							return hx_receiver_586.__hx_this.compare(other)
						}, func() any {
							return hx_receiver_586.__hx_this
						}, "compare")
					}(), Found: true}
				case "getDouble":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int) float64] {
						hx_receiver_587 := hx_structural_source_572
						return hxrt.BoundFunction(func(pos int) float64 {
							return hx_receiver_587.__hx_this.getDouble(pos)
						}, func() any {
							return hx_receiver_587.__hx_this
						}, "getDouble")
					}(), Found: true}
				case "getFloat":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int) float64] {
						hx_receiver_588 := hx_structural_source_572
						return hxrt.BoundFunction(func(pos int) float64 {
							return hx_receiver_588.__hx_this.getFloat(pos)
						}, func() any {
							return hx_receiver_588.__hx_this
						}, "getFloat")
					}(), Found: true}
				case "setDouble":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, float64)] {
						hx_receiver_589 := hx_structural_source_572
						return hxrt.BoundFunction(func(pos int, value float64) {
							hx_receiver_589.__hx_this.setDouble(pos, value)
						}, func() any {
							return hx_receiver_589.__hx_this
						}, "setDouble")
					}(), Found: true}
				case "setFloat":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, float64)] {
						hx_receiver_590 := hx_structural_source_572
						return hxrt.BoundFunction(func(pos int, value float64) {
							hx_receiver_590.__hx_this.setFloat(pos, value)
						}, func() any {
							return hx_receiver_590.__hx_this
						}, "setFloat")
					}(), Found: true}
				case "getUInt16":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int) int] {
						hx_receiver_591 := hx_structural_source_572
						return hxrt.BoundFunction(func(pos int) int {
							return hx_receiver_591.__hx_this.getUInt16(pos)
						}, func() any {
							return hx_receiver_591.__hx_this
						}, "getUInt16")
					}(), Found: true}
				case "setUInt16":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, int)] {
						hx_receiver_592 := hx_structural_source_572
						return hxrt.BoundFunction(func(pos int, value int) {
							hx_receiver_592.__hx_this.setUInt16(pos, value)
						}, func() any {
							return hx_receiver_592.__hx_this
						}, "setUInt16")
					}(), Found: true}
				case "getInt32":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int) int] {
						hx_receiver_593 := hx_structural_source_572
						return hxrt.BoundFunction(func(pos int) int {
							return hx_receiver_593.__hx_this.getInt32(pos)
						}, func() any {
							return hx_receiver_593.__hx_this
						}, "getInt32")
					}(), Found: true}
				case "getInt64":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int) *haxe___Int64_____Int64] {
						hx_receiver_594 := hx_structural_source_572
						return hxrt.BoundFunction(func(pos int) *haxe___Int64_____Int64 {
							return hx_receiver_594.__hx_this.getInt64(pos)
						}, func() any {
							return hx_receiver_594.__hx_this
						}, "getInt64")
					}(), Found: true}
				case "setInt32":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, int)] {
						hx_receiver_595 := hx_structural_source_572
						return hxrt.BoundFunction(func(pos int, value int) {
							hx_receiver_595.__hx_this.setInt32(pos, value)
						}, func() any {
							return hx_receiver_595.__hx_this
						}, "setInt32")
					}(), Found: true}
				case "setInt64":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, *haxe___Int64_____Int64)] {
						hx_receiver_596 := hx_structural_source_572
						return hxrt.BoundFunction(func(pos int, value *haxe___Int64_____Int64) {
							hx_receiver_596.__hx_this.setInt64(pos, value)
						}, func() any {
							return hx_receiver_596.__hx_this
						}, "setInt64")
					}(), Found: true}
				case "getString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, int, *haxe__io__Encoding) *string] {
						hx_receiver_597 := hx_structural_source_572
						return hxrt.BoundFunction(func(pos int, len int, encoding *haxe__io__Encoding) *string {
							return hx_receiver_597.__hx_this.getString(pos, len, encoding)
						}, func() any {
							return hx_receiver_597.__hx_this
						}, "getString")
					}(), Found: true}
				case "readString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, int) *string] {
						hx_receiver_598 := hx_structural_source_572
						return hxrt.BoundFunction(func(pos int, len int) *string {
							return hx_receiver_598.__hx_this.readString(pos, len)
						}, func() any {
							return hx_receiver_598.__hx_this
						}, "readString")
					}(), Found: true}
				case "toString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_599 := hx_structural_source_572
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_599.__hx_this.toString()
						}, func() any {
							return hx_receiver_599.__hx_this
						}, "toString")
					}(), Found: true}
				case "toHex":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_600 := hx_structural_source_572
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_600.__hx_this.toHex()
						}, func() any {
							return hx_receiver_600.__hx_this
						}, "toHex")
					}(), Found: true}
				case "getData":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() []int] {
						hx_receiver_601 := hx_structural_source_572
						return hxrt.BoundFunction(func() []int {
							return hx_receiver_601.__hx_this.getData()
						}, func() any {
							return hx_receiver_601.__hx_this
						}, "getData")
					}(), Found: true}
				case "__hx_nativeView":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *hxrt.ByteView] {
						hx_receiver_602 := hx_structural_source_572
						return hxrt.BoundFunction(func() *hxrt.ByteView {
							return hx_receiver_602.__hx_this.__hx_nativeView()
						}, func() any {
							return hx_receiver_602.__hx_this
						}, "__hx_nativeView")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "b":
					hx_structural_source_572.b = func(hx_value_573 any) []int {
						if hx_value_573 == nil {
							var hx_zero_574 []int
							return hx_zero_574
						}
						return hx_value_573.([]int)
					}(incoming)
					return true
				case "__hx_raw":
					hx_structural_source_572.__hx_raw = func(hx_value_575 any) *hxrt.ByteView {
						if hx_value_575 == nil {
							var hx_zero_576 *hxrt.ByteView
							return hx_zero_576
						}
						return hx_value_575.(*hxrt.ByteView)
					}(incoming)
					return true
				case "__hx_rawValid":
					hx_structural_source_572.__hx_rawValid = func(hx_value_577 any) bool {
						if hx_value_577 == nil {
							var hx_zero_578 bool
							return hx_zero_578
						}
						return hx_value_577.(bool)
					}(incoming)
					return true
				case "__hx_dataExposed":
					hx_structural_source_572.__hx_dataExposed = func(hx_value_579 any) bool {
						if hx_value_579 == nil {
							var hx_zero_580 bool
							return hx_zero_580
						}
						return hx_value_579.(bool)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe__http__HttpBase:
		return func(hx_structural_source_603 *haxe__http__HttpBase) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_603, func(name string) hxrt.AnonymousField {
				switch name {
				case "url":
					return hxrt.AnonymousField{Value: hx_structural_source_603.url, Found: true}
				case "responseBytes":
					return hxrt.AnonymousField{Value: hx_structural_source_603.responseBytes, Found: true}
				case "responseAsString":
					return hxrt.AnonymousField{Value: hx_structural_source_603.responseAsString, Found: true}
				case "postData":
					return hxrt.AnonymousField{Value: hx_structural_source_603.postData, Found: true}
				case "postBytes":
					return hxrt.AnonymousField{Value: hx_structural_source_603.postBytes, Found: true}
				case "headers":
					return hxrt.AnonymousField{Value: hx_structural_source_603.headers, Found: true}
				case "params":
					return hxrt.AnonymousField{Value: hx_structural_source_603.params, Found: true}
				case "emptyOnData":
					return hxrt.AnonymousField{Value: hx_structural_source_603.emptyOnData, Found: true}
				case "setHeader":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string, *string)] {
						hx_receiver_614 := hx_structural_source_603
						return hxrt.BoundFunction(func(name *string, value *string) {
							hx_receiver_614.__hx_this.setHeader(name, value)
						}, func() any {
							return hx_receiver_614.__hx_this
						}, "setHeader")
					}(), Found: true}
				case "addHeader":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string, *string)] {
						hx_receiver_615 := hx_structural_source_603
						return hxrt.BoundFunction(func(header *string, value *string) {
							hx_receiver_615.__hx_this.addHeader(header, value)
						}, func() any {
							return hx_receiver_615.__hx_this
						}, "addHeader")
					}(), Found: true}
				case "setParameter":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string, *string)] {
						hx_receiver_616 := hx_structural_source_603
						return hxrt.BoundFunction(func(name *string, value *string) {
							hx_receiver_616.__hx_this.setParameter(name, value)
						}, func() any {
							return hx_receiver_616.__hx_this
						}, "setParameter")
					}(), Found: true}
				case "addParameter":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string, *string)] {
						hx_receiver_617 := hx_structural_source_603
						return hxrt.BoundFunction(func(name *string, value *string) {
							hx_receiver_617.__hx_this.addParameter(name, value)
						}, func() any {
							return hx_receiver_617.__hx_this
						}, "addParameter")
					}(), Found: true}
				case "setPostData":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string)] {
						hx_receiver_618 := hx_structural_source_603
						return hxrt.BoundFunction(func(data *string) {
							hx_receiver_618.__hx_this.setPostData(data)
						}, func() any {
							return hx_receiver_618.__hx_this
						}, "setPostData")
					}(), Found: true}
				case "setPostBytes":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes)] {
						hx_receiver_619 := hx_structural_source_603
						return hxrt.BoundFunction(func(data *haxe__io__Bytes) {
							hx_receiver_619.__hx_this.setPostBytes(data)
						}, func() any {
							return hx_receiver_619.__hx_this
						}, "setPostBytes")
					}(), Found: true}
				case "request":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(any)] {
						hx_receiver_620 := hx_structural_source_603
						return hxrt.BoundFunction(func(post any) {
							hx_receiver_620.__hx_this.request(post)
						}, func() any {
							return hx_receiver_620.__hx_this
						}, "request")
					}(), Found: true}
				case "onData":
					return hxrt.AnonymousField{Value: hxrt.ReadDynamicFunction(hx_structural_source_603.onData), Found: true}
				case "onBytes":
					return hxrt.AnonymousField{Value: hxrt.ReadDynamicFunction(hx_structural_source_603.onBytes), Found: true}
				case "onError":
					return hxrt.AnonymousField{Value: hxrt.ReadDynamicFunction(hx_structural_source_603.onError), Found: true}
				case "onStatus":
					return hxrt.AnonymousField{Value: hxrt.ReadDynamicFunction(hx_structural_source_603.onStatus), Found: true}
				case "hasOnData":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() bool] {
						hx_receiver_621 := hx_structural_source_603
						return hxrt.BoundFunction(func() bool {
							return hx_receiver_621.__hx_this.hasOnData()
						}, func() any {
							return hx_receiver_621.__hx_this
						}, "hasOnData")
					}(), Found: true}
				case "success":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes)] {
						hx_receiver_622 := hx_structural_source_603
						return hxrt.BoundFunction(func(data *haxe__io__Bytes) {
							hx_receiver_622.__hx_this.success(data)
						}, func() any {
							return hx_receiver_622.__hx_this
						}, "success")
					}(), Found: true}
				case "get_responseData":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_623 := hx_structural_source_603
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_623.__hx_this.get_responseData()
						}, func() any {
							return hx_receiver_623.__hx_this
						}, "get_responseData")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "url":
					hx_structural_source_603.url = func(hx_value_604 any) *string {
						if hx_value_604 == nil {
							var hx_zero_605 *string
							return hx_zero_605
						}
						return hx_value_604.(*string)
					}(incoming)
					return true
				case "responseAsString":
					hx_structural_source_603.responseAsString = func(hx_value_606 any) *string {
						if hx_value_606 == nil {
							var hx_zero_607 *string
							return hx_zero_607
						}
						return hx_value_606.(*string)
					}(incoming)
					return true
				case "postData":
					hx_structural_source_603.postData = func(hx_value_608 any) *string {
						if hx_value_608 == nil {
							var hx_zero_609 *string
							return hx_zero_609
						}
						return hx_value_608.(*string)
					}(incoming)
					return true
				case "postBytes":
					hx_structural_source_603.postBytes = hx_recover_class_6(incoming)
					return true
				case "headers":
					hx_structural_source_603.headers = func(hx_value_610 any) *hxrt.Array {
						if hx_value_610 == nil {
							var hx_zero_611 *hxrt.Array
							return hx_zero_611
						}
						return hx_value_610.(*hxrt.Array)
					}(incoming)
					return true
				case "params":
					hx_structural_source_603.params = func(hx_value_612 any) *hxrt.Array {
						if hx_value_612 == nil {
							var hx_zero_613 *hxrt.Array
							return hx_zero_613
						}
						return hx_value_612.(*hxrt.Array)
					}(incoming)
					return true
				case "onData":
					hx_structural_source_603.onData = hx_recover_callable_3(incoming)
					return true
				case "onBytes":
					hx_structural_source_603.onBytes = hx_recover_callable_4(incoming)
					return true
				case "onError":
					hx_structural_source_603.onError = hx_recover_callable_3(incoming)
					return true
				case "onStatus":
					hx_structural_source_603.onStatus = hx_recover_callable_5(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe__ds__StringMap:
		return func(hx_structural_source_624 *haxe__ds__StringMap) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_624, func(name string) hxrt.AnonymousField {
				switch name {
				case "h":
					return hxrt.AnonymousField{Value: hx_structural_source_624.h, Found: true}
				case "set":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string, any)] {
						hx_receiver_627 := hx_structural_source_624
						return hxrt.BoundFunction(func(key *string, value any) {
							hx_receiver_627.__hx_this.set(key, value)
						}, func() any {
							return hx_receiver_627.__hx_this
						}, "set")
					}(), Found: true}
				case "get":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string) any] {
						hx_receiver_628 := hx_structural_source_624
						return hxrt.BoundFunction(func(key *string) any {
							return hx_receiver_628.__hx_this.get(key)
						}, func() any {
							return hx_receiver_628.__hx_this
						}, "get")
					}(), Found: true}
				case "exists":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string) bool] {
						hx_receiver_629 := hx_structural_source_624
						return hxrt.BoundFunction(func(key *string) bool {
							return hx_receiver_629.__hx_this.exists(key)
						}, func() any {
							return hx_receiver_629.__hx_this
						}, "exists")
					}(), Found: true}
				case "remove":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string) bool] {
						hx_receiver_630 := hx_structural_source_624
						return hxrt.BoundFunction(func(key *string) bool {
							return hx_receiver_630.__hx_this.remove(key)
						}, func() any {
							return hx_receiver_630.__hx_this
						}, "remove")
					}(), Found: true}
				case "keys":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() hxrt.Anonymous] {
						hx_receiver_631 := hx_structural_source_624
						return hxrt.BoundFunction(func() hxrt.Anonymous {
							return hx_receiver_631.__hx_this.keys()
						}, func() any {
							return hx_receiver_631.__hx_this
						}, "keys")
					}(), Found: true}
				case "iterator":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() hxrt.Anonymous] {
						hx_receiver_632 := hx_structural_source_624
						return hxrt.BoundFunction(func() hxrt.Anonymous {
							return hx_receiver_632.__hx_this.iterator()
						}, func() any {
							return hx_receiver_632.__hx_this
						}, "iterator")
					}(), Found: true}
				case "keyValueIterator":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() hxrt.Anonymous] {
						hx_receiver_633 := hx_structural_source_624
						return hxrt.BoundFunction(func() hxrt.Anonymous {
							return hx_receiver_633.__hx_this.keyValueIterator()
						}, func() any {
							return hx_receiver_633.__hx_this
						}, "keyValueIterator")
					}(), Found: true}
				case "getIMap":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(any) any] {
						hx_receiver_634 := hx_structural_source_624
						return hxrt.BoundFunction(func(key any) any {
							return hx_receiver_634.__hx_this.getIMap(key)
						}, func() any {
							return hx_receiver_634.__hx_this
						}, "getIMap")
					}(), Found: true}
				case "setIMap":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(any, any)] {
						hx_receiver_635 := hx_structural_source_624
						return hxrt.BoundFunction(func(key any, value any) {
							hx_receiver_635.__hx_this.setIMap(key, value)
						}, func() any {
							return hx_receiver_635.__hx_this
						}, "setIMap")
					}(), Found: true}
				case "existsIMap":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(any) bool] {
						hx_receiver_636 := hx_structural_source_624
						return hxrt.BoundFunction(func(key any) bool {
							return hx_receiver_636.__hx_this.existsIMap(key)
						}, func() any {
							return hx_receiver_636.__hx_this
						}, "existsIMap")
					}(), Found: true}
				case "removeIMap":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(any) bool] {
						hx_receiver_637 := hx_structural_source_624
						return hxrt.BoundFunction(func(key any) bool {
							return hx_receiver_637.__hx_this.removeIMap(key)
						}, func() any {
							return hx_receiver_637.__hx_this
						}, "removeIMap")
					}(), Found: true}
				case "copyIMap":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() haxe__IMap] {
						hx_receiver_638 := hx_structural_source_624
						return hxrt.BoundFunction(func() haxe__IMap {
							return hx_receiver_638.__hx_this.copyIMap()
						}, func() any {
							return hx_receiver_638.__hx_this
						}, "copyIMap")
					}(), Found: true}
				case "copy":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *haxe__ds__StringMap] {
						hx_receiver_639 := hx_structural_source_624
						return hxrt.BoundFunction(func() *haxe__ds__StringMap {
							return hx_receiver_639.__hx_this.copy()
						}, func() any {
							return hx_receiver_639.__hx_this
						}, "copy")
					}(), Found: true}
				case "toString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_640 := hx_structural_source_624
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_640.__hx_this.toString()
						}, func() any {
							return hx_receiver_640.__hx_this
						}, "toString")
					}(), Found: true}
				case "clear":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func()] {
						hx_receiver_641 := hx_structural_source_624
						return hxrt.BoundFunction(func() {
							hx_receiver_641.__hx_this.clear()
						}, func() any {
							return hx_receiver_641.__hx_this
						}, "clear")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "h":
					hx_structural_source_624.h = func(hx_value_625 any) *hxrt.StringMapCell {
						if hx_value_625 == nil {
							var hx_zero_626 *hxrt.StringMapCell
							return hx_zero_626
						}
						return hx_value_625.(*hxrt.StringMapCell)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe__io__Eof:
		return func(hx_structural_source_642 *haxe__io__Eof) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_642, func(name string) hxrt.AnonymousField {
				switch name {
				case "toString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_643 := hx_structural_source_642
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_643.__hx_this.toString()
						}, func() any {
							return hx_receiver_643.__hx_this
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
	case *haxe__io__BytesBuffer:
		return func(hx_structural_source_644 *haxe__io__BytesBuffer) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_644, func(name string) hxrt.AnonymousField {
				switch name {
				case "b":
					return hxrt.AnonymousField{Value: hx_structural_source_644.b, Found: true}
				case "get_length":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_647 := hx_structural_source_644
						return hxrt.BoundFunction(func() int {
							return hx_receiver_647.__hx_this.get_length()
						}, func() any {
							return hx_receiver_647.__hx_this
						}, "get_length")
					}(), Found: true}
				case "addByte":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_648 := hx_structural_source_644
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_648.__hx_this.addByte(value)
						}, func() any {
							return hx_receiver_648.__hx_this
						}, "addByte")
					}(), Found: true}
				case "add":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes)] {
						hx_receiver_649 := hx_structural_source_644
						return hxrt.BoundFunction(func(source *haxe__io__Bytes) {
							hx_receiver_649.__hx_this.add(source)
						}, func() any {
							return hx_receiver_649.__hx_this
						}, "add")
					}(), Found: true}
				case "addString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string, *haxe__io__Encoding)] {
						hx_receiver_650 := hx_structural_source_644
						return hxrt.BoundFunction(func(value *string, encoding *haxe__io__Encoding) {
							hx_receiver_650.__hx_this.addString(value, encoding)
						}, func() any {
							return hx_receiver_650.__hx_this
						}, "addString")
					}(), Found: true}
				case "addInt32":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_651 := hx_structural_source_644
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_651.__hx_this.addInt32(value)
						}, func() any {
							return hx_receiver_651.__hx_this
						}, "addInt32")
					}(), Found: true}
				case "addInt64":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe___Int64_____Int64)] {
						hx_receiver_652 := hx_structural_source_644
						return hxrt.BoundFunction(func(value *haxe___Int64_____Int64) {
							hx_receiver_652.__hx_this.addInt64(value)
						}, func() any {
							return hx_receiver_652.__hx_this
						}, "addInt64")
					}(), Found: true}
				case "addFloat":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(float64)] {
						hx_receiver_653 := hx_structural_source_644
						return hxrt.BoundFunction(func(value float64) {
							hx_receiver_653.__hx_this.addFloat(value)
						}, func() any {
							return hx_receiver_653.__hx_this
						}, "addFloat")
					}(), Found: true}
				case "addDouble":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(float64)] {
						hx_receiver_654 := hx_structural_source_644
						return hxrt.BoundFunction(func(value float64) {
							hx_receiver_654.__hx_this.addDouble(value)
						}, func() any {
							return hx_receiver_654.__hx_this
						}, "addDouble")
					}(), Found: true}
				case "addBytes":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes, int, int)] {
						hx_receiver_655 := hx_structural_source_644
						return hxrt.BoundFunction(func(source *haxe__io__Bytes, pos int, len int) {
							hx_receiver_655.__hx_this.addBytes(source, pos, len)
						}, func() any {
							return hx_receiver_655.__hx_this
						}, "addBytes")
					}(), Found: true}
				case "getBytes":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *haxe__io__Bytes] {
						hx_receiver_656 := hx_structural_source_644
						return hxrt.BoundFunction(func() *haxe__io__Bytes {
							return hx_receiver_656.__hx_this.getBytes()
						}, func() any {
							return hx_receiver_656.__hx_this
						}, "getBytes")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "b":
					hx_structural_source_644.b = func(hx_value_645 any) []int {
						if hx_value_645 == nil {
							var hx_zero_646 []int
							return hx_zero_646
						}
						return hx_value_645.([]int)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe__exceptions__NotImplementedException:
		return func(hx_structural_source_657 *haxe__exceptions__NotImplementedException) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_657, func(name string) hxrt.AnonymousField {
				switch name {
				case "posInfos":
					return hxrt.AnonymousField{Value: hx_structural_source_657.posInfos, Found: true}
				case "toString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_658 := hx_structural_source_657
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_658.__hx_this.toString()
						}, func() any {
							return hx_receiver_658.__hx_this
						}, "toString")
					}(), Found: true}
				case "savedMessage":
					return hxrt.AnonymousField{Value: hx_structural_source_657.savedMessage, Found: true}
				case "savedStack":
					return hxrt.AnonymousField{Value: hx_structural_source_657.savedStack, Found: true}
				case "savedPrevious":
					return hxrt.AnonymousField{Value: hx_structural_source_657.savedPrevious, Found: true}
				case "savedNative":
					return hxrt.AnonymousField{Value: hx_structural_source_657.savedNative, Found: true}
				case "get_message":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_661 := hx_structural_source_657
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_661.__hx_this.get_message()
						}, func() any {
							return hx_receiver_661.__hx_this
						}, "get_message")
					}(), Found: true}
				case "get_stack":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_662 := hx_structural_source_657
						return hxrt.BoundFunction(func() any {
							return hx_receiver_662.__hx_this.get_stack()
						}, func() any {
							return hx_receiver_662.__hx_this
						}, "get_stack")
					}(), Found: true}
				case "get_previous":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *haxe__Exception] {
						hx_receiver_663 := hx_structural_source_657
						return hxrt.BoundFunction(func() *haxe__Exception {
							return hx_receiver_663.__hx_this.get_previous()
						}, func() any {
							return hx_receiver_663.__hx_this
						}, "get_previous")
					}(), Found: true}
				case "get_native":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_664 := hx_structural_source_657
						return hxrt.BoundFunction(func() any {
							return hx_receiver_664.__hx_this.get_native()
						}, func() any {
							return hx_receiver_664.__hx_this
						}, "get_native")
					}(), Found: true}
				case "unwrap":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_665 := hx_structural_source_657
						return hxrt.BoundFunction(func() any {
							return hx_receiver_665.__hx_this.unwrap()
						}, func() any {
							return hx_receiver_665.__hx_this
						}, "unwrap")
					}(), Found: true}
				case "details":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_666 := hx_structural_source_657
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_666.__hx_this.details()
						}, func() any {
							return hx_receiver_666.__hx_this
						}, "details")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "savedMessage":
					hx_structural_source_657.savedMessage = func(hx_value_659 any) *string {
						if hx_value_659 == nil {
							var hx_zero_660 *string
							return hx_zero_660
						}
						return hx_value_659.(*string)
					}(incoming)
					return true
				case "savedStack":
					hx_structural_source_657.savedStack = incoming
					return true
				case "savedPrevious":
					hx_structural_source_657.savedPrevious = hx_recover_class_1(incoming)
					return true
				case "savedNative":
					hx_structural_source_657.savedNative = incoming
					return true
				}
				return false
			})
		}(source)
	case *haxe__Exception:
		return func(hx_structural_source_667 *haxe__Exception) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_667, func(name string) hxrt.AnonymousField {
				switch name {
				case "savedMessage":
					return hxrt.AnonymousField{Value: hx_structural_source_667.savedMessage, Found: true}
				case "savedStack":
					return hxrt.AnonymousField{Value: hx_structural_source_667.savedStack, Found: true}
				case "savedPrevious":
					return hxrt.AnonymousField{Value: hx_structural_source_667.savedPrevious, Found: true}
				case "savedNative":
					return hxrt.AnonymousField{Value: hx_structural_source_667.savedNative, Found: true}
				case "get_message":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_670 := hx_structural_source_667
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_670.__hx_this.get_message()
						}, func() any {
							return hx_receiver_670.__hx_this
						}, "get_message")
					}(), Found: true}
				case "get_stack":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_671 := hx_structural_source_667
						return hxrt.BoundFunction(func() any {
							return hx_receiver_671.__hx_this.get_stack()
						}, func() any {
							return hx_receiver_671.__hx_this
						}, "get_stack")
					}(), Found: true}
				case "get_previous":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *haxe__Exception] {
						hx_receiver_672 := hx_structural_source_667
						return hxrt.BoundFunction(func() *haxe__Exception {
							return hx_receiver_672.__hx_this.get_previous()
						}, func() any {
							return hx_receiver_672.__hx_this
						}, "get_previous")
					}(), Found: true}
				case "get_native":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_673 := hx_structural_source_667
						return hxrt.BoundFunction(func() any {
							return hx_receiver_673.__hx_this.get_native()
						}, func() any {
							return hx_receiver_673.__hx_this
						}, "get_native")
					}(), Found: true}
				case "unwrap":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_674 := hx_structural_source_667
						return hxrt.BoundFunction(func() any {
							return hx_receiver_674.__hx_this.unwrap()
						}, func() any {
							return hx_receiver_674.__hx_this
						}, "unwrap")
					}(), Found: true}
				case "toString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_675 := hx_structural_source_667
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_675.__hx_this.toString()
						}, func() any {
							return hx_receiver_675.__hx_this
						}, "toString")
					}(), Found: true}
				case "details":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_676 := hx_structural_source_667
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_676.__hx_this.details()
						}, func() any {
							return hx_receiver_676.__hx_this
						}, "details")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "savedMessage":
					hx_structural_source_667.savedMessage = func(hx_value_668 any) *string {
						if hx_value_668 == nil {
							var hx_zero_669 *string
							return hx_zero_669
						}
						return hx_value_668.(*string)
					}(incoming)
					return true
				case "savedStack":
					hx_structural_source_667.savedStack = incoming
					return true
				case "savedPrevious":
					hx_structural_source_667.savedPrevious = hx_recover_class_1(incoming)
					return true
				case "savedNative":
					hx_structural_source_667.savedNative = incoming
					return true
				}
				return false
			})
		}(source)
	case *haxe__iterators__StringKeyValueIterator:
		return func(hx_structural_source_677 *haxe__iterators__StringKeyValueIterator) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_677, func(name string) hxrt.AnonymousField {
				switch name {
				case "offset":
					return hxrt.AnonymousField{Value: hx_structural_source_677.offset, Found: true}
				case "s":
					return hxrt.AnonymousField{Value: hx_structural_source_677.s, Found: true}
				case "hasNext":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() bool] {
						hx_receiver_682 := hx_structural_source_677
						return hxrt.BoundFunction(func() bool {
							return hx_receiver_682.__hx_this.hasNext()
						}, func() any {
							return hx_receiver_682.__hx_this
						}, "hasNext")
					}(), Found: true}
				case "next":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() hxrt.Anonymous] {
						hx_receiver_683 := hx_structural_source_677
						return hxrt.BoundFunction(func() hxrt.Anonymous {
							return hx_receiver_683.__hx_this.next()
						}, func() any {
							return hx_receiver_683.__hx_this
						}, "next")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "offset":
					hx_structural_source_677.offset = func(hx_value_678 any) int {
						if hx_value_678 == nil {
							var hx_zero_679 int
							return hx_zero_679
						}
						return hx_value_678.(int)
					}(incoming)
					return true
				case "s":
					hx_structural_source_677.s = func(hx_value_680 any) *string {
						if hx_value_680 == nil {
							var hx_zero_681 *string
							return hx_zero_681
						}
						return hx_value_680.(*string)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe__iterators__StringIterator:
		return func(hx_structural_source_684 *haxe__iterators__StringIterator) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_684, func(name string) hxrt.AnonymousField {
				switch name {
				case "offset":
					return hxrt.AnonymousField{Value: hx_structural_source_684.offset, Found: true}
				case "s":
					return hxrt.AnonymousField{Value: hx_structural_source_684.s, Found: true}
				case "hasNext":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() bool] {
						hx_receiver_689 := hx_structural_source_684
						return hxrt.BoundFunction(func() bool {
							return hx_receiver_689.__hx_this.hasNext()
						}, func() any {
							return hx_receiver_689.__hx_this
						}, "hasNext")
					}(), Found: true}
				case "next":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_690 := hx_structural_source_684
						return hxrt.BoundFunction(func() int {
							return hx_receiver_690.__hx_this.next()
						}, func() any {
							return hx_receiver_690.__hx_this
						}, "next")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "offset":
					hx_structural_source_684.offset = func(hx_value_685 any) int {
						if hx_value_685 == nil {
							var hx_zero_686 int
							return hx_zero_686
						}
						return hx_value_685.(int)
					}(incoming)
					return true
				case "s":
					hx_structural_source_684.s = func(hx_value_687 any) *string {
						if hx_value_687 == nil {
							var hx_zero_688 *string
							return hx_zero_688
						}
						return hx_value_687.(*string)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe__iterators__MapKeyValueIterator:
		return func(hx_structural_source_691 *haxe__iterators__MapKeyValueIterator) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_691, func(name string) hxrt.AnonymousField {
				switch name {
				case "map":
					return hxrt.AnonymousField{Value: hx_structural_source_691.map_, Found: true}
				case "keys":
					return hxrt.AnonymousField{Value: hx_structural_source_691.keys, Found: true}
				case "hasNext":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() bool] {
						hx_receiver_694 := hx_structural_source_691
						return hxrt.BoundFunction(func() bool {
							return hx_receiver_694.__hx_this.hasNext()
						}, func() any {
							return hx_receiver_694.__hx_this
						}, "hasNext")
					}(), Found: true}
				case "next":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() hxrt.Anonymous] {
						hx_receiver_695 := hx_structural_source_691
						return hxrt.BoundFunction(func() hxrt.Anonymous {
							return hx_receiver_695.__hx_this.next()
						}, func() any {
							return hx_receiver_695.__hx_this
						}, "next")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "map":
					hx_structural_source_691.map_ = func(hx_value_692 any) haxe__IMap {
						if hx_value_692 == nil {
							var hx_zero_693 haxe__IMap
							return hx_zero_693
						}
						return hx_value_692.(haxe__IMap)
					}(hxrt.ObjectSource(incoming))
					return true
				case "keys":
					hx_structural_source_691.keys = hx_structural_recover(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe__io__BytesOutput:
		return func(hx_structural_source_696 *haxe__io__BytesOutput) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_696, func(name string) hxrt.AnonymousField {
				switch name {
				case "b":
					return hxrt.AnonymousField{Value: hx_structural_source_696.b, Found: true}
				case "get_length":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_697 := hx_structural_source_696
						return hxrt.BoundFunction(func() int {
							return hx_receiver_697.__hx_this.get_length()
						}, func() any {
							return hx_receiver_697.__hx_this
						}, "get_length")
					}(), Found: true}
				case "writeByte":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_698 := hx_structural_source_696
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_698.__hx_this.writeByte(value)
						}, func() any {
							return hx_receiver_698.__hx_this
						}, "writeByte")
					}(), Found: true}
				case "writeBytes":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes, int, int) int] {
						hx_receiver_699 := hx_structural_source_696
						return hxrt.BoundFunction(func(bytes *haxe__io__Bytes, pos int, len int) int {
							return hx_receiver_699.__hx_this.writeBytes(bytes, pos, len)
						}, func() any {
							return hx_receiver_699.__hx_this
						}, "writeBytes")
					}(), Found: true}
				case "getBytes":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *haxe__io__Bytes] {
						hx_receiver_700 := hx_structural_source_696
						return hxrt.BoundFunction(func() *haxe__io__Bytes {
							return hx_receiver_700.__hx_this.getBytes()
						}, func() any {
							return hx_receiver_700.__hx_this
						}, "getBytes")
					}(), Found: true}
				case "bigEndian":
					return hxrt.AnonymousField{Value: hx_structural_source_696.bigEndian, Found: true}
				case "flush":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func()] {
						hx_receiver_701 := hx_structural_source_696
						return hxrt.BoundFunction(func() {
							hx_receiver_701.__hx_this.flush()
						}, func() any {
							return hx_receiver_701.__hx_this
						}, "flush")
					}(), Found: true}
				case "close":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func()] {
						hx_receiver_702 := hx_structural_source_696
						return hxrt.BoundFunction(func() {
							hx_receiver_702.__hx_this.close()
						}, func() any {
							return hx_receiver_702.__hx_this
						}, "close")
					}(), Found: true}
				case "set_bigEndian":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(bool) bool] {
						hx_receiver_703 := hx_structural_source_696
						return hxrt.BoundFunction(func(value bool) bool {
							return hx_receiver_703.__hx_this.set_bigEndian(value)
						}, func() any {
							return hx_receiver_703.__hx_this
						}, "set_bigEndian")
					}(), Found: true}
				case "write":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes)] {
						hx_receiver_704 := hx_structural_source_696
						return hxrt.BoundFunction(func(bytes *haxe__io__Bytes) {
							hx_receiver_704.__hx_this.write(bytes)
						}, func() any {
							return hx_receiver_704.__hx_this
						}, "write")
					}(), Found: true}
				case "writeFullBytes":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes, int, int)] {
						hx_receiver_705 := hx_structural_source_696
						return hxrt.BoundFunction(func(bytes *haxe__io__Bytes, pos int, len int) {
							hx_receiver_705.__hx_this.writeFullBytes(bytes, pos, len)
						}, func() any {
							return hx_receiver_705.__hx_this
						}, "writeFullBytes")
					}(), Found: true}
				case "writeFloat":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(float64)] {
						hx_receiver_706 := hx_structural_source_696
						return hxrt.BoundFunction(func(value float64) {
							hx_receiver_706.__hx_this.writeFloat(value)
						}, func() any {
							return hx_receiver_706.__hx_this
						}, "writeFloat")
					}(), Found: true}
				case "writeDouble":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(float64)] {
						hx_receiver_707 := hx_structural_source_696
						return hxrt.BoundFunction(func(value float64) {
							hx_receiver_707.__hx_this.writeDouble(value)
						}, func() any {
							return hx_receiver_707.__hx_this
						}, "writeDouble")
					}(), Found: true}
				case "writeInt8":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_708 := hx_structural_source_696
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_708.__hx_this.writeInt8(value)
						}, func() any {
							return hx_receiver_708.__hx_this
						}, "writeInt8")
					}(), Found: true}
				case "writeInt16":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_709 := hx_structural_source_696
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_709.__hx_this.writeInt16(value)
						}, func() any {
							return hx_receiver_709.__hx_this
						}, "writeInt16")
					}(), Found: true}
				case "writeUInt16":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_710 := hx_structural_source_696
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_710.__hx_this.writeUInt16(value)
						}, func() any {
							return hx_receiver_710.__hx_this
						}, "writeUInt16")
					}(), Found: true}
				case "writeInt24":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_711 := hx_structural_source_696
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_711.__hx_this.writeInt24(value)
						}, func() any {
							return hx_receiver_711.__hx_this
						}, "writeInt24")
					}(), Found: true}
				case "writeUInt24":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_712 := hx_structural_source_696
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_712.__hx_this.writeUInt24(value)
						}, func() any {
							return hx_receiver_712.__hx_this
						}, "writeUInt24")
					}(), Found: true}
				case "writeInt32":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_713 := hx_structural_source_696
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_713.__hx_this.writeInt32(value)
						}, func() any {
							return hx_receiver_713.__hx_this
						}, "writeInt32")
					}(), Found: true}
				case "prepare":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_714 := hx_structural_source_696
						return hxrt.BoundFunction(func(nbytes int) {
							hx_receiver_714.__hx_this.prepare(nbytes)
						}, func() any {
							return hx_receiver_714.__hx_this
						}, "prepare")
					}(), Found: true}
				case "writeInput":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Input, any)] {
						hx_receiver_715 := hx_structural_source_696
						return hxrt.BoundFunction(func(input *haxe__io__Input, bufsize any) {
							hx_receiver_715.__hx_this.writeInput(input, bufsize)
						}, func() any {
							return hx_receiver_715.__hx_this
						}, "writeInput")
					}(), Found: true}
				case "writeString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string, *haxe__io__Encoding)] {
						hx_receiver_716 := hx_structural_source_696
						return hxrt.BoundFunction(func(value *string, encoding *haxe__io__Encoding) {
							hx_receiver_716.__hx_this.writeString(value, encoding)
						}, func() any {
							return hx_receiver_716.__hx_this
						}, "writeString")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "b":
					hx_structural_source_696.b = hx_recover_class_7(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe__exceptions__PosException:
		return func(hx_structural_source_717 *haxe__exceptions__PosException) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_717, func(name string) hxrt.AnonymousField {
				switch name {
				case "posInfos":
					return hxrt.AnonymousField{Value: hx_structural_source_717.posInfos, Found: true}
				case "toString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_718 := hx_structural_source_717
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_718.__hx_this.toString()
						}, func() any {
							return hx_receiver_718.__hx_this
						}, "toString")
					}(), Found: true}
				case "savedMessage":
					return hxrt.AnonymousField{Value: hx_structural_source_717.savedMessage, Found: true}
				case "savedStack":
					return hxrt.AnonymousField{Value: hx_structural_source_717.savedStack, Found: true}
				case "savedPrevious":
					return hxrt.AnonymousField{Value: hx_structural_source_717.savedPrevious, Found: true}
				case "savedNative":
					return hxrt.AnonymousField{Value: hx_structural_source_717.savedNative, Found: true}
				case "get_message":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_721 := hx_structural_source_717
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_721.__hx_this.get_message()
						}, func() any {
							return hx_receiver_721.__hx_this
						}, "get_message")
					}(), Found: true}
				case "get_stack":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_722 := hx_structural_source_717
						return hxrt.BoundFunction(func() any {
							return hx_receiver_722.__hx_this.get_stack()
						}, func() any {
							return hx_receiver_722.__hx_this
						}, "get_stack")
					}(), Found: true}
				case "get_previous":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *haxe__Exception] {
						hx_receiver_723 := hx_structural_source_717
						return hxrt.BoundFunction(func() *haxe__Exception {
							return hx_receiver_723.__hx_this.get_previous()
						}, func() any {
							return hx_receiver_723.__hx_this
						}, "get_previous")
					}(), Found: true}
				case "get_native":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_724 := hx_structural_source_717
						return hxrt.BoundFunction(func() any {
							return hx_receiver_724.__hx_this.get_native()
						}, func() any {
							return hx_receiver_724.__hx_this
						}, "get_native")
					}(), Found: true}
				case "unwrap":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_725 := hx_structural_source_717
						return hxrt.BoundFunction(func() any {
							return hx_receiver_725.__hx_this.unwrap()
						}, func() any {
							return hx_receiver_725.__hx_this
						}, "unwrap")
					}(), Found: true}
				case "details":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_726 := hx_structural_source_717
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_726.__hx_this.details()
						}, func() any {
							return hx_receiver_726.__hx_this
						}, "details")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "savedMessage":
					hx_structural_source_717.savedMessage = func(hx_value_719 any) *string {
						if hx_value_719 == nil {
							var hx_zero_720 *string
							return hx_zero_720
						}
						return hx_value_719.(*string)
					}(incoming)
					return true
				case "savedStack":
					hx_structural_source_717.savedStack = incoming
					return true
				case "savedPrevious":
					hx_structural_source_717.savedPrevious = hx_recover_class_1(incoming)
					return true
				case "savedNative":
					hx_structural_source_717.savedNative = incoming
					return true
				}
				return false
			})
		}(source)
	case *haxe__ValueException:
		return func(hx_structural_source_727 *haxe__ValueException) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_727, func(name string) hxrt.AnonymousField {
				switch name {
				case "value":
					return hxrt.AnonymousField{Value: hx_structural_source_727.value, Found: true}
				case "unwrap":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_728 := hx_structural_source_727
						return hxrt.BoundFunction(func() any {
							return hx_receiver_728.__hx_this.unwrap()
						}, func() any {
							return hx_receiver_728.__hx_this
						}, "unwrap")
					}(), Found: true}
				case "savedMessage":
					return hxrt.AnonymousField{Value: hx_structural_source_727.savedMessage, Found: true}
				case "savedStack":
					return hxrt.AnonymousField{Value: hx_structural_source_727.savedStack, Found: true}
				case "savedPrevious":
					return hxrt.AnonymousField{Value: hx_structural_source_727.savedPrevious, Found: true}
				case "savedNative":
					return hxrt.AnonymousField{Value: hx_structural_source_727.savedNative, Found: true}
				case "get_message":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_731 := hx_structural_source_727
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_731.__hx_this.get_message()
						}, func() any {
							return hx_receiver_731.__hx_this
						}, "get_message")
					}(), Found: true}
				case "get_stack":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_732 := hx_structural_source_727
						return hxrt.BoundFunction(func() any {
							return hx_receiver_732.__hx_this.get_stack()
						}, func() any {
							return hx_receiver_732.__hx_this
						}, "get_stack")
					}(), Found: true}
				case "get_previous":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *haxe__Exception] {
						hx_receiver_733 := hx_structural_source_727
						return hxrt.BoundFunction(func() *haxe__Exception {
							return hx_receiver_733.__hx_this.get_previous()
						}, func() any {
							return hx_receiver_733.__hx_this
						}, "get_previous")
					}(), Found: true}
				case "get_native":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_734 := hx_structural_source_727
						return hxrt.BoundFunction(func() any {
							return hx_receiver_734.__hx_this.get_native()
						}, func() any {
							return hx_receiver_734.__hx_this
						}, "get_native")
					}(), Found: true}
				case "toString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_735 := hx_structural_source_727
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_735.__hx_this.toString()
						}, func() any {
							return hx_receiver_735.__hx_this
						}, "toString")
					}(), Found: true}
				case "details":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_736 := hx_structural_source_727
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_736.__hx_this.details()
						}, func() any {
							return hx_receiver_736.__hx_this
						}, "details")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "savedMessage":
					hx_structural_source_727.savedMessage = func(hx_value_729 any) *string {
						if hx_value_729 == nil {
							var hx_zero_730 *string
							return hx_zero_730
						}
						return hx_value_729.(*string)
					}(incoming)
					return true
				case "savedStack":
					hx_structural_source_727.savedStack = incoming
					return true
				case "savedPrevious":
					hx_structural_source_727.savedPrevious = hx_recover_class_1(incoming)
					return true
				case "savedNative":
					hx_structural_source_727.savedNative = incoming
					return true
				}
				return false
			})
		}(source)
	default:
		return hxrt.AnonymousFromValue(value)
	}
}

func hx_recover_callable_0(value any) *hxrt.Function[func() bool] {
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
			return func(hx_value_737 any) bool {
				if hx_value_737 == nil {
					var hx_zero_738 bool
					return hx_zero_738
				}
				return hx_value_737.(bool)
			}(any(hxrt.FunctionCall(source)()))
		})
	case *hxrt.Function[func() any]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_739 any) bool {
				if hx_value_739 == nil {
					var hx_zero_740 bool
					return hx_zero_740
				}
				return hx_value_739.(bool)
			}(any(hxrt.FunctionCall(source)()))
		})
	case *hxrt.Function[func() hxrt.Anonymous]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_741 any) bool {
				if hx_value_741 == nil {
					var hx_zero_742 bool
					return hx_zero_742
				}
				return hx_value_741.(bool)
			}(any(hxrt.FunctionCall(source)()))
		})
	case *hxrt.Function[func()]:
		return hxrt.ProjectFunction(source, func() bool {
			hxrt.FunctionCall(source)()
			return func(hx_value_743 any) bool {
				if hx_value_743 == nil {
					var hx_zero_744 bool
					return hx_zero_744
				}
				return hx_value_743.(bool)
			}(any(nil))
		})
	case *hxrt.Function[func() int]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_745 any) bool {
				if hx_value_745 == nil {
					var hx_zero_746 bool
					return hx_zero_746
				}
				return hx_value_745.(bool)
			}(any(hxrt.FunctionCall(source)()))
		})
	case *hxrt.Function[func() float64]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_747 any) bool {
				if hx_value_747 == nil {
					var hx_zero_748 bool
					return hx_zero_748
				}
				return hx_value_747.(bool)
			}(any(hxrt.FunctionCall(source)()))
		})
	case *hxrt.Function[func() *sys__net__Socket]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_749 any) bool {
				if hx_value_749 == nil {
					var hx_zero_750 bool
					return hx_zero_750
				}
				return hx_value_749.(bool)
			}(any(hxrt.FunctionCall(source)()))
		})
	case *hxrt.Function[func() []int]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_751 any) bool {
				if hx_value_751 == nil {
					var hx_zero_752 bool
					return hx_zero_752
				}
				return hx_value_751.(bool)
			}(any(hxrt.FunctionCall(source)()))
		})
	case *hxrt.Function[func() *hxrt.ByteView]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_753 any) bool {
				if hx_value_753 == nil {
					var hx_zero_754 bool
					return hx_zero_754
				}
				return hx_value_753.(bool)
			}(any(hxrt.FunctionCall(source)()))
		})
	case *hxrt.Function[func() haxe__IMap]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_755 any) bool {
				if hx_value_755 == nil {
					var hx_zero_756 bool
					return hx_zero_756
				}
				return hx_value_755.(bool)
			}(any(hxrt.FunctionCall(source)()))
		})
	case *hxrt.Function[func() *haxe__ds__StringMap]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_757 any) bool {
				if hx_value_757 == nil {
					var hx_zero_758 bool
					return hx_zero_758
				}
				return hx_value_757.(bool)
			}(any(hxrt.FunctionCall(source)()))
		})
	case *hxrt.Function[func() *haxe__io__Bytes]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_759 any) bool {
				if hx_value_759 == nil {
					var hx_zero_760 bool
					return hx_zero_760
				}
				return hx_value_759.(bool)
			}(any(hxrt.FunctionCall(source)()))
		})
	case *hxrt.Function[func() *haxe__Exception]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_761 any) bool {
				if hx_value_761 == nil {
					var hx_zero_762 bool
					return hx_zero_762
				}
				return hx_value_761.(bool)
			}(any(hxrt.FunctionCall(source)()))
		})
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}

func hx_recover_callable_1(value any) *hxrt.Function[func() *string] {
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
			return func(hx_value_763 any) *string {
				if hx_value_763 == nil {
					var hx_zero_764 *string
					return hx_zero_764
				}
				return hx_value_763.(*string)
			}(any(hxrt.FunctionCall(source)()))
		})
	case *hxrt.Function[func() *string]:
		return source
	case *hxrt.Function[func() any]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_765 any) *string {
				if hx_value_765 == nil {
					var hx_zero_766 *string
					return hx_zero_766
				}
				return hx_value_765.(*string)
			}(any(hxrt.FunctionCall(source)()))
		})
	case *hxrt.Function[func() hxrt.Anonymous]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_767 any) *string {
				if hx_value_767 == nil {
					var hx_zero_768 *string
					return hx_zero_768
				}
				return hx_value_767.(*string)
			}(any(hxrt.FunctionCall(source)()))
		})
	case *hxrt.Function[func()]:
		return hxrt.ProjectFunction(source, func() *string {
			hxrt.FunctionCall(source)()
			return func(hx_value_769 any) *string {
				if hx_value_769 == nil {
					var hx_zero_770 *string
					return hx_zero_770
				}
				return hx_value_769.(*string)
			}(any(nil))
		})
	case *hxrt.Function[func() int]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_771 any) *string {
				if hx_value_771 == nil {
					var hx_zero_772 *string
					return hx_zero_772
				}
				return hx_value_771.(*string)
			}(any(hxrt.FunctionCall(source)()))
		})
	case *hxrt.Function[func() float64]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_773 any) *string {
				if hx_value_773 == nil {
					var hx_zero_774 *string
					return hx_zero_774
				}
				return hx_value_773.(*string)
			}(any(hxrt.FunctionCall(source)()))
		})
	case *hxrt.Function[func() *sys__net__Socket]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_775 any) *string {
				if hx_value_775 == nil {
					var hx_zero_776 *string
					return hx_zero_776
				}
				return hx_value_775.(*string)
			}(any(hxrt.FunctionCall(source)()))
		})
	case *hxrt.Function[func() []int]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_777 any) *string {
				if hx_value_777 == nil {
					var hx_zero_778 *string
					return hx_zero_778
				}
				return hx_value_777.(*string)
			}(any(hxrt.FunctionCall(source)()))
		})
	case *hxrt.Function[func() *hxrt.ByteView]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_779 any) *string {
				if hx_value_779 == nil {
					var hx_zero_780 *string
					return hx_zero_780
				}
				return hx_value_779.(*string)
			}(any(hxrt.FunctionCall(source)()))
		})
	case *hxrt.Function[func() haxe__IMap]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_781 any) *string {
				if hx_value_781 == nil {
					var hx_zero_782 *string
					return hx_zero_782
				}
				return hx_value_781.(*string)
			}(any(hxrt.FunctionCall(source)()))
		})
	case *hxrt.Function[func() *haxe__ds__StringMap]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_783 any) *string {
				if hx_value_783 == nil {
					var hx_zero_784 *string
					return hx_zero_784
				}
				return hx_value_783.(*string)
			}(any(hxrt.FunctionCall(source)()))
		})
	case *hxrt.Function[func() *haxe__io__Bytes]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_785 any) *string {
				if hx_value_785 == nil {
					var hx_zero_786 *string
					return hx_zero_786
				}
				return hx_value_785.(*string)
			}(any(hxrt.FunctionCall(source)()))
		})
	case *hxrt.Function[func() *haxe__Exception]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_787 any) *string {
				if hx_value_787 == nil {
					var hx_zero_788 *string
					return hx_zero_788
				}
				return hx_value_787.(*string)
			}(any(hxrt.FunctionCall(source)()))
		})
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}

func hx_recover_callable_2(value any) *hxrt.Function[func() any] {
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
	case *hxrt.Function[func() []int]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func() *hxrt.ByteView]:
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
	case *hxrt.Function[func() *haxe__io__Bytes]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func() *haxe__Exception]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}

func hx_recover_callable_3(value any) *hxrt.Function[func(*string)] {
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
	case *hxrt.Function[func(int)]:
		return hxrt.ProjectFunction(source, func(arg_0 *string) {
			hxrt.FunctionCall(source)(func(hx_value_789 any) int {
				if hx_value_789 == nil {
					var hx_zero_790 int
					return hx_zero_790
				}
				return hx_value_789.(int)
			}(any(arg_0)))
		})
	case *hxrt.Function[func(*haxe__io__Bytes)]:
		return hxrt.ProjectFunction(source, func(arg_0 *string) {
			hxrt.FunctionCall(source)(hx_recover_class_6(arg_0))
		})
	case *hxrt.Function[func(bool) bool]:
		return hxrt.ProjectFunction(source, func(arg_0 *string) {
			hxrt.FunctionCall(source)(func(hx_value_791 any) bool {
				if hx_value_791 == nil {
					var hx_zero_792 bool
					return hx_zero_792
				}
				return hx_value_791.(bool)
			}(any(arg_0)))
		})
	case *hxrt.Function[func(float64)]:
		return hxrt.ProjectFunction(source, func(arg_0 *string) {
			hxrt.FunctionCall(source)(func(hx_value_793 any) float64 {
				if hx_value_793 == nil {
					var hx_zero_794 float64
					return hx_zero_794
				}
				return hx_value_793.(float64)
			}(any(arg_0)))
		})
	case *hxrt.Function[func(any) *haxe__io__Bytes]:
		return hxrt.ProjectFunction(source, func(arg_0 *string) {
			hxrt.FunctionCall(source)(arg_0)
		})
	case *hxrt.Function[func(int) *haxe__io__Bytes]:
		return hxrt.ProjectFunction(source, func(arg_0 *string) {
			hxrt.FunctionCall(source)(func(hx_value_795 any) int {
				if hx_value_795 == nil {
					var hx_zero_796 int
					return hx_zero_796
				}
				return hx_value_795.(int)
			}(any(arg_0)))
		})
	case *hxrt.Function[func(int) *string]:
		return hxrt.ProjectFunction(source, func(arg_0 *string) {
			hxrt.FunctionCall(source)(func(hx_value_797 any) int {
				if hx_value_797 == nil {
					var hx_zero_798 int
					return hx_zero_798
				}
				return hx_value_797.(int)
			}(any(arg_0)))
		})
	case *hxrt.Function[func(bool)]:
		return hxrt.ProjectFunction(source, func(arg_0 *string) {
			hxrt.FunctionCall(source)(func(hx_value_799 any) bool {
				if hx_value_799 == nil {
					var hx_zero_800 bool
					return hx_zero_800
				}
				return hx_value_799.(bool)
			}(any(arg_0)))
		})
	case *hxrt.Function[func(*hxrt.SocketHandle)]:
		return hxrt.ProjectFunction(source, func(arg_0 *string) {
			hxrt.FunctionCall(source)(func(hx_value_801 any) *hxrt.SocketHandle {
				if hx_value_801 == nil {
					var hx_zero_802 *hxrt.SocketHandle
					return hx_zero_802
				}
				return hx_value_801.(*hxrt.SocketHandle)
			}(any(arg_0)))
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
			hxrt.FunctionCall(source)(func(hx_value_803 any) *hxrt.HttpExchange {
				if hx_value_803 == nil {
					var hx_zero_804 *hxrt.HttpExchange
					return hx_zero_804
				}
				return hx_value_803.(*hxrt.HttpExchange)
			}(any(arg_0)))
		})
	case *hxrt.Function[func(int) int]:
		return hxrt.ProjectFunction(source, func(arg_0 *string) {
			hxrt.FunctionCall(source)(func(hx_value_805 any) int {
				if hx_value_805 == nil {
					var hx_zero_806 int
					return hx_zero_806
				}
				return hx_value_805.(int)
			}(any(arg_0)))
		})
	case *hxrt.Function[func(*haxe__io__Bytes) int]:
		return hxrt.ProjectFunction(source, func(arg_0 *string) {
			hxrt.FunctionCall(source)(hx_recover_class_6(arg_0))
		})
	case *hxrt.Function[func(int) float64]:
		return hxrt.ProjectFunction(source, func(arg_0 *string) {
			hxrt.FunctionCall(source)(func(hx_value_807 any) int {
				if hx_value_807 == nil {
					var hx_zero_808 int
					return hx_zero_808
				}
				return hx_value_807.(int)
			}(any(arg_0)))
		})
	case *hxrt.Function[func(int) *haxe___Int64_____Int64]:
		return hxrt.ProjectFunction(source, func(arg_0 *string) {
			hxrt.FunctionCall(source)(func(hx_value_809 any) int {
				if hx_value_809 == nil {
					var hx_zero_810 int
					return hx_zero_810
				}
				return hx_value_809.(int)
			}(any(arg_0)))
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
	case *hxrt.Function[func(*haxe___Int64_____Int64)]:
		return hxrt.ProjectFunction(source, func(arg_0 *string) {
			hxrt.FunctionCall(source)(hx_recover_class_8(arg_0))
		})
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}

func hx_recover_callable_4(value any) *hxrt.Function[func(*haxe__io__Bytes)] {
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
			hxrt.FunctionCall(source)(func(hx_value_811 any) *string {
				if hx_value_811 == nil {
					var hx_zero_812 *string
					return hx_zero_812
				}
				return hx_value_811.(*string)
			}(any(arg_0)))
		})
	case *hxrt.Function[func(int)]:
		return hxrt.ProjectFunction(source, func(arg_0 *haxe__io__Bytes) {
			hxrt.FunctionCall(source)(func(hx_value_813 any) int {
				if hx_value_813 == nil {
					var hx_zero_814 int
					return hx_zero_814
				}
				return hx_value_813.(int)
			}(any(arg_0)))
		})
	case *hxrt.Function[func(*haxe__io__Bytes)]:
		return source
	case *hxrt.Function[func(bool) bool]:
		return hxrt.ProjectFunction(source, func(arg_0 *haxe__io__Bytes) {
			hxrt.FunctionCall(source)(func(hx_value_815 any) bool {
				if hx_value_815 == nil {
					var hx_zero_816 bool
					return hx_zero_816
				}
				return hx_value_815.(bool)
			}(any(arg_0)))
		})
	case *hxrt.Function[func(float64)]:
		return hxrt.ProjectFunction(source, func(arg_0 *haxe__io__Bytes) {
			hxrt.FunctionCall(source)(func(hx_value_817 any) float64 {
				if hx_value_817 == nil {
					var hx_zero_818 float64
					return hx_zero_818
				}
				return hx_value_817.(float64)
			}(any(arg_0)))
		})
	case *hxrt.Function[func(any) *haxe__io__Bytes]:
		return hxrt.ProjectFunction(source, func(arg_0 *haxe__io__Bytes) {
			hxrt.FunctionCall(source)(arg_0)
		})
	case *hxrt.Function[func(int) *haxe__io__Bytes]:
		return hxrt.ProjectFunction(source, func(arg_0 *haxe__io__Bytes) {
			hxrt.FunctionCall(source)(func(hx_value_819 any) int {
				if hx_value_819 == nil {
					var hx_zero_820 int
					return hx_zero_820
				}
				return hx_value_819.(int)
			}(any(arg_0)))
		})
	case *hxrt.Function[func(int) *string]:
		return hxrt.ProjectFunction(source, func(arg_0 *haxe__io__Bytes) {
			hxrt.FunctionCall(source)(func(hx_value_821 any) int {
				if hx_value_821 == nil {
					var hx_zero_822 int
					return hx_zero_822
				}
				return hx_value_821.(int)
			}(any(arg_0)))
		})
	case *hxrt.Function[func(bool)]:
		return hxrt.ProjectFunction(source, func(arg_0 *haxe__io__Bytes) {
			hxrt.FunctionCall(source)(func(hx_value_823 any) bool {
				if hx_value_823 == nil {
					var hx_zero_824 bool
					return hx_zero_824
				}
				return hx_value_823.(bool)
			}(any(arg_0)))
		})
	case *hxrt.Function[func(*hxrt.SocketHandle)]:
		return hxrt.ProjectFunction(source, func(arg_0 *haxe__io__Bytes) {
			hxrt.FunctionCall(source)(func(hx_value_825 any) *hxrt.SocketHandle {
				if hx_value_825 == nil {
					var hx_zero_826 *hxrt.SocketHandle
					return hx_zero_826
				}
				return hx_value_825.(*hxrt.SocketHandle)
			}(any(arg_0)))
		})
	case *hxrt.Function[func(any)]:
		return hxrt.ProjectFunction(source, func(arg_0 *haxe__io__Bytes) {
			hxrt.FunctionCall(source)(arg_0)
		})
	case *hxrt.Function[func(*string) *hxrt.Array]:
		return hxrt.ProjectFunction(source, func(arg_0 *haxe__io__Bytes) {
			hxrt.FunctionCall(source)(func(hx_value_827 any) *string {
				if hx_value_827 == nil {
					var hx_zero_828 *string
					return hx_zero_828
				}
				return hx_value_827.(*string)
			}(any(arg_0)))
		})
	case *hxrt.Function[func(*hxrt.HttpExchange)]:
		return hxrt.ProjectFunction(source, func(arg_0 *haxe__io__Bytes) {
			hxrt.FunctionCall(source)(func(hx_value_829 any) *hxrt.HttpExchange {
				if hx_value_829 == nil {
					var hx_zero_830 *hxrt.HttpExchange
					return hx_zero_830
				}
				return hx_value_829.(*hxrt.HttpExchange)
			}(any(arg_0)))
		})
	case *hxrt.Function[func(int) int]:
		return hxrt.ProjectFunction(source, func(arg_0 *haxe__io__Bytes) {
			hxrt.FunctionCall(source)(func(hx_value_831 any) int {
				if hx_value_831 == nil {
					var hx_zero_832 int
					return hx_zero_832
				}
				return hx_value_831.(int)
			}(any(arg_0)))
		})
	case *hxrt.Function[func(*haxe__io__Bytes) int]:
		return hxrt.ProjectFunction(source, func(arg_0 *haxe__io__Bytes) {
			hxrt.FunctionCall(source)(arg_0)
		})
	case *hxrt.Function[func(int) float64]:
		return hxrt.ProjectFunction(source, func(arg_0 *haxe__io__Bytes) {
			hxrt.FunctionCall(source)(func(hx_value_833 any) int {
				if hx_value_833 == nil {
					var hx_zero_834 int
					return hx_zero_834
				}
				return hx_value_833.(int)
			}(any(arg_0)))
		})
	case *hxrt.Function[func(int) *haxe___Int64_____Int64]:
		return hxrt.ProjectFunction(source, func(arg_0 *haxe__io__Bytes) {
			hxrt.FunctionCall(source)(func(hx_value_835 any) int {
				if hx_value_835 == nil {
					var hx_zero_836 int
					return hx_zero_836
				}
				return hx_value_835.(int)
			}(any(arg_0)))
		})
	case *hxrt.Function[func(*string) any]:
		return hxrt.ProjectFunction(source, func(arg_0 *haxe__io__Bytes) {
			hxrt.FunctionCall(source)(func(hx_value_837 any) *string {
				if hx_value_837 == nil {
					var hx_zero_838 *string
					return hx_zero_838
				}
				return hx_value_837.(*string)
			}(any(arg_0)))
		})
	case *hxrt.Function[func(*string) bool]:
		return hxrt.ProjectFunction(source, func(arg_0 *haxe__io__Bytes) {
			hxrt.FunctionCall(source)(func(hx_value_839 any) *string {
				if hx_value_839 == nil {
					var hx_zero_840 *string
					return hx_zero_840
				}
				return hx_value_839.(*string)
			}(any(arg_0)))
		})
	case *hxrt.Function[func(any) any]:
		return hxrt.ProjectFunction(source, func(arg_0 *haxe__io__Bytes) {
			hxrt.FunctionCall(source)(arg_0)
		})
	case *hxrt.Function[func(any) bool]:
		return hxrt.ProjectFunction(source, func(arg_0 *haxe__io__Bytes) {
			hxrt.FunctionCall(source)(arg_0)
		})
	case *hxrt.Function[func(*haxe___Int64_____Int64)]:
		return hxrt.ProjectFunction(source, func(arg_0 *haxe__io__Bytes) {
			hxrt.FunctionCall(source)(hx_recover_class_8(arg_0))
		})
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
}

func hx_recover_callable_5(value any) *hxrt.Function[func(int)] {
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
			hxrt.FunctionCall(source)(func(hx_value_841 any) *string {
				if hx_value_841 == nil {
					var hx_zero_842 *string
					return hx_zero_842
				}
				return hx_value_841.(*string)
			}(any(arg_0)))
		})
	case *hxrt.Function[func(int)]:
		return source
	case *hxrt.Function[func(*haxe__io__Bytes)]:
		return hxrt.ProjectFunction(source, func(arg_0 int) {
			hxrt.FunctionCall(source)(hx_recover_class_6(arg_0))
		})
	case *hxrt.Function[func(bool) bool]:
		return hxrt.ProjectFunction(source, func(arg_0 int) {
			hxrt.FunctionCall(source)(func(hx_value_843 any) bool {
				if hx_value_843 == nil {
					var hx_zero_844 bool
					return hx_zero_844
				}
				return hx_value_843.(bool)
			}(any(arg_0)))
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
			hxrt.FunctionCall(source)(func(hx_value_845 any) bool {
				if hx_value_845 == nil {
					var hx_zero_846 bool
					return hx_zero_846
				}
				return hx_value_845.(bool)
			}(any(arg_0)))
		})
	case *hxrt.Function[func(*hxrt.SocketHandle)]:
		return hxrt.ProjectFunction(source, func(arg_0 int) {
			hxrt.FunctionCall(source)(func(hx_value_847 any) *hxrt.SocketHandle {
				if hx_value_847 == nil {
					var hx_zero_848 *hxrt.SocketHandle
					return hx_zero_848
				}
				return hx_value_847.(*hxrt.SocketHandle)
			}(any(arg_0)))
		})
	case *hxrt.Function[func(any)]:
		return hxrt.ProjectFunction(source, func(arg_0 int) {
			hxrt.FunctionCall(source)(arg_0)
		})
	case *hxrt.Function[func(*string) *hxrt.Array]:
		return hxrt.ProjectFunction(source, func(arg_0 int) {
			hxrt.FunctionCall(source)(func(hx_value_849 any) *string {
				if hx_value_849 == nil {
					var hx_zero_850 *string
					return hx_zero_850
				}
				return hx_value_849.(*string)
			}(any(arg_0)))
		})
	case *hxrt.Function[func(*hxrt.HttpExchange)]:
		return hxrt.ProjectFunction(source, func(arg_0 int) {
			hxrt.FunctionCall(source)(func(hx_value_851 any) *hxrt.HttpExchange {
				if hx_value_851 == nil {
					var hx_zero_852 *hxrt.HttpExchange
					return hx_zero_852
				}
				return hx_value_851.(*hxrt.HttpExchange)
			}(any(arg_0)))
		})
	case *hxrt.Function[func(int) int]:
		return hxrt.ProjectFunction(source, func(arg_0 int) {
			hxrt.FunctionCall(source)(arg_0)
		})
	case *hxrt.Function[func(*haxe__io__Bytes) int]:
		return hxrt.ProjectFunction(source, func(arg_0 int) {
			hxrt.FunctionCall(source)(hx_recover_class_6(arg_0))
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
			hxrt.FunctionCall(source)(func(hx_value_853 any) *string {
				if hx_value_853 == nil {
					var hx_zero_854 *string
					return hx_zero_854
				}
				return hx_value_853.(*string)
			}(any(arg_0)))
		})
	case *hxrt.Function[func(*string) bool]:
		return hxrt.ProjectFunction(source, func(arg_0 int) {
			hxrt.FunctionCall(source)(func(hx_value_855 any) *string {
				if hx_value_855 == nil {
					var hx_zero_856 *string
					return hx_zero_856
				}
				return hx_value_855.(*string)
			}(any(arg_0)))
		})
	case *hxrt.Function[func(any) any]:
		return hxrt.ProjectFunction(source, func(arg_0 int) {
			hxrt.FunctionCall(source)(arg_0)
		})
	case *hxrt.Function[func(any) bool]:
		return hxrt.ProjectFunction(source, func(arg_0 int) {
			hxrt.FunctionCall(source)(arg_0)
		})
	case *hxrt.Function[func(*haxe___Int64_____Int64)]:
		return hxrt.ProjectFunction(source, func(arg_0 int) {
			hxrt.FunctionCall(source)(hx_recover_class_8(arg_0))
		})
	default:
		hxrt.Throw(hxrt.StringFromLiteral("Invalid function value"))
		return nil
	}
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

func hx_recover_class_1(value any) *haxe__Exception {
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

func hx_recover_class_2(value any) *haxe__io__Eof {
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

func hx_recover_class_3(value any) *haxe__io__Input {
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

func hx_recover_class_4(value any) *haxe__ds__StringMap {
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

func hx_recover_class_5(value any) *haxe__ValueException {
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

func hx_recover_class_6(value any) *haxe__io__Bytes {
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

func hx_recover_class_7(value any) *haxe__io__BytesBuffer {
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

func hx_recover_class_8(value any) *haxe___Int64_____Int64 {
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
