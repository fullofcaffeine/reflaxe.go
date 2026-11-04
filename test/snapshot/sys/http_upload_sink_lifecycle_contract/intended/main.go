package main

import "snapshot/hxrt"

type hxrt__TypeClassValue struct {
	name *string
}

type hxrt__TypeEnumValue struct {
	name *string
}

func finishServer(server hxrt.Anonymous) {
	hxrt.TryCatch(func() {
		hx_recover_class_0(hxrt.AnonymousGet(server, "process")).__hx_this.close()
	}, func(hx_caught_1 any) {
		hx_exception_2 := haxe__Exception_caught(hx_caught_1)
		hx_payload_3 := haxe__Exception_catchValue(hx_caught_1)
		_ = hx_exception_2
		var hx_tmp any = hx_payload_3
		_ = hx_tmp
	})
}

func main() {
	defer hxrt.ThreadWaitForAll()
	successServer := startServer(hxrt.StringFromLiteral("consume"), 1)
	success := run(successServer, hxrt.StringFromLiteral("payload"), 7, 2.0)
	finishServer(successServer)
	var v any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("success="), hxrt.StdString(((((((func(hx_value_4 any) int {
		if hx_value_4 == nil {
			var hx_zero_5 int
			return hx_zero_5
		}
		return hx_value_4.(int)
	}(hxrt.AnonymousGet(success, "status")) == 200) && hxrt.StringEqualStringPtr(func(hx_value_6 any) *string {
		if hx_value_6 == nil {
			var hx_zero_7 *string
			return hx_zero_7
		}
		return hx_value_6.(*string)
	}(hxrt.AnonymousGet(success, "error")), hxrt.StringFromLiteral(""))) && hxrt.StringEqualStringPtr(func(hx_value_8 any) *string {
		if hx_value_8 == nil {
			var hx_zero_9 *string
			return hx_zero_9
		}
		return hx_value_8.(*string)
	}(hxrt.AnonymousGet(success, "data")), hxrt.StringFromLiteral("ok"))) && (func(hx_value_10 any) int {
		if hx_value_10 == nil {
			var hx_zero_11 int
			return hx_zero_11
		}
		return hx_value_10.(int)
	}(hxrt.AnonymousGet(success, "reads")) == 4)) && !func(hx_value_12 any) bool {
		if hx_value_12 == nil {
			var hx_zero_13 bool
			return hx_zero_13
		}
		return hx_value_12.(bool)
	}(hxrt.AnonymousGet(success, "wrongThread"))) && (func(hx_value_14 any) int {
		if hx_value_14 == nil {
			var hx_zero_15 int
			return hx_zero_15
		}
		return hx_value_14.(int)
	}(hxrt.AnonymousGet(success, "afterReturnReads")) == 0)))))
	hxrt.Println(v)
	sourceServer := startServer(hxrt.StringFromLiteral("consume"), 1)
	source := run(sourceServer, hxrt.StringFromLiteral("source-error"), 7, 2.0)
	finishServer(sourceServer)
	var v_1 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("source="), hxrt.StdString((((((func(hx_value_16 any) int {
		if hx_value_16 == nil {
			var hx_zero_17 int
			return hx_zero_17
		}
		return hx_value_16.(int)
	}(hxrt.AnonymousGet(source, "status")) == -1) && hxrt.StringEqualStringPtr(func(hx_value_18 any) *string {
		if hx_value_18 == nil {
			var hx_zero_19 *string
			return hx_zero_19
		}
		return hx_value_18.(*string)
	}(hxrt.AnonymousGet(source, "error")), hxrt.StringFromLiteral("source-exploded"))) && (func(hx_value_20 any) int {
		if hx_value_20 == nil {
			var hx_zero_21 int
			return hx_zero_21
		}
		return hx_value_20.(int)
	}(hxrt.AnonymousGet(source, "reads")) == 2)) && !func(hx_value_22 any) bool {
		if hx_value_22 == nil {
			var hx_zero_23 bool
			return hx_zero_23
		}
		return hx_value_22.(bool)
	}(hxrt.AnonymousGet(source, "wrongThread"))) && (func(hx_value_24 any) int {
		if hx_value_24 == nil {
			var hx_zero_25 int
			return hx_zero_25
		}
		return hx_value_24.(int)
	}(hxrt.AnonymousGet(source, "afterReturnReads")) == 0)))))
	hxrt.Println(v_1)
	eofServer := startServer(hxrt.StringFromLiteral("consume"), 1)
	eof := run(eofServer, hxrt.StringFromLiteral("early-eof"), 7, 2.0)
	finishServer(eofServer)
	var v_2 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("earlyEof="), hxrt.StdString((((((func(hx_value_26 any) int {
		if hx_value_26 == nil {
			var hx_zero_27 int
			return hx_zero_27
		}
		return hx_value_26.(int)
	}(hxrt.AnonymousGet(eof, "status")) == -1) && hxrt.StringEqualStringPtr(func(hx_value_28 any) *string {
		if hx_value_28 == nil {
			var hx_zero_29 *string
			return hx_zero_29
		}
		return hx_value_28.(*string)
	}(hxrt.AnonymousGet(eof, "error")), hxrt.StringFromLiteral("Transfer aborted"))) && (func(hx_value_30 any) int {
		if hx_value_30 == nil {
			var hx_zero_31 int
			return hx_zero_31
		}
		return hx_value_30.(int)
	}(hxrt.AnonymousGet(eof, "reads")) == 2)) && !func(hx_value_32 any) bool {
		if hx_value_32 == nil {
			var hx_zero_33 bool
			return hx_zero_33
		}
		return hx_value_32.(bool)
	}(hxrt.AnonymousGet(eof, "wrongThread"))) && (func(hx_value_34 any) int {
		if hx_value_34 == nil {
			var hx_zero_35 int
			return hx_zero_35
		}
		return hx_value_34.(int)
	}(hxrt.AnonymousGet(eof, "afterReturnReads")) == 0)))))
	hxrt.Println(v_2)
	zeroServer := startServer(hxrt.StringFromLiteral("consume"), 1)
	zero := run(zeroServer, hxrt.StringFromLiteral("zero"), 7, 2.0)
	finishServer(zeroServer)
	var v_3 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("zero="), hxrt.StdString((((((func(hx_value_36 any) int {
		if hx_value_36 == nil {
			var hx_zero_37 int
			return hx_zero_37
		}
		return hx_value_36.(int)
	}(hxrt.AnonymousGet(zero, "status")) == -1) && hxrt.StringEqualStringPtr(func(hx_value_38 any) *string {
		if hx_value_38 == nil {
			var hx_zero_39 *string
			return hx_zero_39
		}
		return hx_value_38.(*string)
	}(hxrt.AnonymousGet(zero, "error")), hxrt.StringFromLiteral("multipart upload made no progress"))) && (func(hx_value_40 any) int {
		if hx_value_40 == nil {
			var hx_zero_41 int
			return hx_zero_41
		}
		return hx_value_40.(int)
	}(hxrt.AnonymousGet(zero, "reads")) == 1)) && !func(hx_value_42 any) bool {
		if hx_value_42 == nil {
			var hx_zero_43 bool
			return hx_zero_43
		}
		return hx_value_42.(bool)
	}(hxrt.AnonymousGet(zero, "wrongThread"))) && (func(hx_value_44 any) int {
		if hx_value_44 == nil {
			var hx_zero_45 int
			return hx_zero_45
		}
		return hx_value_44.(int)
	}(hxrt.AnonymousGet(zero, "afterReturnReads")) == 0)))))
	hxrt.Println(v_3)
	closeServer := startServer(hxrt.StringFromLiteral("close"), 1)
	closed := run(closeServer, hxrt.StringFromLiteral("stream"), 8388608, 2.0)
	finishServer(closeServer)
	var v_4 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("serverClose="), hxrt.StdString(((!hxrt.StringEqualStringPtr(func(hx_value_46 any) *string {
		if hx_value_46 == nil {
			var hx_zero_47 *string
			return hx_zero_47
		}
		return hx_value_46.(*string)
	}(hxrt.AnonymousGet(closed, "error")), hxrt.StringFromLiteral("")) && !func(hx_value_48 any) bool {
		if hx_value_48 == nil {
			var hx_zero_49 bool
			return hx_zero_49
		}
		return hx_value_48.(bool)
	}(hxrt.AnonymousGet(closed, "wrongThread"))) && (func(hx_value_50 any) int {
		if hx_value_50 == nil {
			var hx_zero_51 int
			return hx_zero_51
		}
		return hx_value_50.(int)
	}(hxrt.AnonymousGet(closed, "afterReturnReads")) == 0)))))
	hxrt.Println(v_4)
	timeoutServer := startServer(hxrt.StringFromLiteral("timeout"), 1)
	timed := run(timeoutServer, hxrt.StringFromLiteral("stream"), 16777216, 0.1)
	finishServer(timeoutServer)
	var v_5 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("timeout="), hxrt.StdString(((!hxrt.StringEqualStringPtr(func(hx_value_52 any) *string {
		if hx_value_52 == nil {
			var hx_zero_53 *string
			return hx_zero_53
		}
		return hx_value_52.(*string)
	}(hxrt.AnonymousGet(timed, "error")), hxrt.StringFromLiteral("")) && !func(hx_value_54 any) bool {
		if hx_value_54 == nil {
			var hx_zero_55 bool
			return hx_zero_55
		}
		return hx_value_54.(bool)
	}(hxrt.AnonymousGet(timed, "wrongThread"))) && (func(hx_value_56 any) int {
		if hx_value_56 == nil {
			var hx_zero_57 int
			return hx_zero_57
		}
		return hx_value_56.(int)
	}(hxrt.AnonymousGet(timed, "afterReturnReads")) == 0)))))
	hxrt.Println(v_5)
	repeatedServer := startServer(hxrt.StringFromLiteral("early"), 12)
	repeatedStatus := 0
	repeatedError := 0
	repeatedThread := 0
	repeatedAfterReturn := 0
	_g := 0
	for _g < 12 {
		hx_post_58 := _g
		_g = int(int32((_g + 1)))
		hx_tmp := hx_post_58
		_ = hx_tmp
		early := run(repeatedServer, hxrt.StringFromLiteral("stream"), 8388608, 2.0)
		if func(hx_value_59 any) int {
			if hx_value_59 == nil {
				var hx_zero_60 int
				return hx_zero_60
			}
			return hx_value_59.(int)
		}(hxrt.AnonymousGet(early, "status")) == 413 {
			repeatedStatus = int(int32((repeatedStatus + 1)))
		}
		if hxrt.StringEqualStringPtr(func(hx_value_61 any) *string {
			if hx_value_61 == nil {
				var hx_zero_62 *string
				return hx_zero_62
			}
			return hx_value_61.(*string)
		}(hxrt.AnonymousGet(early, "error")), hxrt.StringFromLiteral("Http Error #413")) {
			repeatedError = int(int32((repeatedError + 1)))
		}
		if !func(hx_value_63 any) bool {
			if hx_value_63 == nil {
				var hx_zero_64 bool
				return hx_zero_64
			}
			return hx_value_63.(bool)
		}(hxrt.AnonymousGet(early, "wrongThread")) {
			repeatedThread = int(int32((repeatedThread + 1)))
		}
		if func(hx_value_65 any) int {
			if hx_value_65 == nil {
				var hx_zero_66 int
				return hx_zero_66
			}
			return hx_value_65.(int)
		}(hxrt.AnonymousGet(early, "afterReturnReads")) == 0 {
			repeatedAfterReturn = int(int32((repeatedAfterReturn + 1)))
		}
	}
	finishServer(repeatedServer)
	hxrt.Println(any(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("earlyRepeated=status:"), repeatedStatus), hxrt.StringFromLiteral(";error:")), repeatedError), hxrt.StringFromLiteral(";thread:")), repeatedThread), hxrt.StringFromLiteral(";after:")), repeatedAfterReturn)))
}

func pythonServerScript(mode *string, count int) *string {
	return hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("import http.server\nimport socketserver\nimport time\nMODE = "), func() *string {
		var space *string = nil
		return hxrt.StdString(hxrt.JsonStringify(any(mode), space))
	}()), hxrt.StringFromLiteral("\n")), hxrt.StringFromLiteral("COUNT = ")), count), hxrt.StringFromLiteral("\n")), hxrt.StringFromLiteral("class Handler(http.server.BaseHTTPRequestHandler):\n")), hxrt.StringFromLiteral("    def do_POST(self):\n")), hxrt.StringFromLiteral("        if MODE == 'early':\n")), hxrt.StringFromLiteral("            body = b''\n")), hxrt.StringFromLiteral("            self.send_response(413)\n")), hxrt.StringFromLiteral("            self.send_header('Content-Length', str(len(body)))\n")), hxrt.StringFromLiteral("            self.end_headers()\n")), hxrt.StringFromLiteral("            self.wfile.write(body)\n")), hxrt.StringFromLiteral("            self.wfile.flush()\n")), hxrt.StringFromLiteral("            self.close_connection = True\n")), hxrt.StringFromLiteral("            self.connection.settimeout(0.5)\n")), hxrt.StringFromLiteral("            remaining = int(self.headers.get('Content-Length', '0'))\n")), hxrt.StringFromLiteral("            try:\n")), hxrt.StringFromLiteral("                while remaining > 0:\n")), hxrt.StringFromLiteral("                    chunk = self.rfile.read(min(65536, remaining))\n")), hxrt.StringFromLiteral("                    if not chunk:\n")), hxrt.StringFromLiteral("                        break\n")), hxrt.StringFromLiteral("                    remaining -= len(chunk)\n")), hxrt.StringFromLiteral("            except (TimeoutError, ConnectionResetError):\n")), hxrt.StringFromLiteral("                pass\n")), hxrt.StringFromLiteral("            return\n")), hxrt.StringFromLiteral("        if MODE == 'timeout':\n")), hxrt.StringFromLiteral("            time.sleep(0.5)\n")), hxrt.StringFromLiteral("            return\n")), hxrt.StringFromLiteral("        if MODE == 'close':\n")), hxrt.StringFromLiteral("            self.connection.shutdown(2)\n")), hxrt.StringFromLiteral("            self.connection.close()\n")), hxrt.StringFromLiteral("            return\n")), hxrt.StringFromLiteral("        length = int(self.headers.get('Content-Length', '0'))\n")), hxrt.StringFromLiteral("        self.rfile.read(length)\n")), hxrt.StringFromLiteral("        body = b'ok'\n")), hxrt.StringFromLiteral("        try:\n")), hxrt.StringFromLiteral("            self.send_response(200)\n")), hxrt.StringFromLiteral("            self.send_header('Content-Length', str(len(body)))\n")), hxrt.StringFromLiteral("            self.end_headers()\n")), hxrt.StringFromLiteral("            self.wfile.write(body)\n")), hxrt.StringFromLiteral("        except (BrokenPipeError, ConnectionResetError):\n")), hxrt.StringFromLiteral("            pass\n")), hxrt.StringFromLiteral("    def log_message(self, fmt, *args):\n")), hxrt.StringFromLiteral("        return\n")), hxrt.StringFromLiteral("with socketserver.TCPServer(('127.0.0.1', 0), Handler) as httpd:\n")), hxrt.StringFromLiteral("    print(httpd.server_address[1], flush=True)\n")), hxrt.StringFromLiteral("    for _ in range(COUNT):\n")), hxrt.StringFromLiteral("        httpd.handle_request()\n"))
}

func run(server hxrt.Anonymous, mode *string, declaredSize int, timeout float64) hxrt.Anonymous {
	input := New__Main__TrackingInput(mode, func() *string {
		var hx_if_67 *string
		if hxrt.StringEqualStringPtr(mode, hxrt.StringFromLiteral("payload")) {
			hx_if_67 = hxrt.StringFromLiteral("payload")
		} else {
			hx_if_67 = hxrt.StringFromLiteral("")
		}
		return hx_if_67
	}())
	request := New_sys__Http(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("http://127.0.0.1:"), func(hx_value_68 any) int {
		if hx_value_68 == nil {
			var hx_zero_69 int
			return hx_zero_69
		}
		return hx_value_68.(int)
	}(hxrt.AnonymousGet(server, "port"))), hxrt.StringFromLiteral("/upload")))
	request.cnxTimeout = timeout
	request.__hx_this.fileTransfer(hxrt.StringFromLiteral("asset"), hxrt.StringFromLiteral("demo.bin"), func(hx_value *_Main__TrackingInput) *haxe__io__Input {
		if hx_value == nil {
			return nil
		}
		return hx_value.haxe__io__Input
	}(input), declaredSize, hxrt.StringFromLiteral("application/octet-stream"))
	status := -1
	error := hxrt.StringFromLiteral("")
	data := hxrt.StringFromLiteral("")
	request.onStatus = hxrt.NewFunction(func(value int) {
		status = value
	})
	request.onError = hxrt.NewFunction(func(value *string) {
		error = value
	})
	request.onData = hxrt.NewFunction(func(value *string) {
		data = value
	})
	request.__hx_this.request(true)
	input.requestReturned = true
	hxrt.SysSleep(0.01)
	hx_obj_70 := map[string]any{}
	hx_obj_70["status"] = status
	hx_obj_70["error"] = error
	hx_obj_70["data"] = data
	hx_obj_70["reads"] = input.reads
	hx_obj_70["wrongThread"] = input.wrongThread
	hx_obj_70["afterReturnReads"] = input.afterReturnReads
	return hxrt.AnonymousFromMap(hx_obj_70)
}

func startServer(mode *string, count int) hxrt.Anonymous {
	process := New_sys__io__Process(hxrt.StringFromLiteral("python3"), hxrt.NewArray(hxrt.StringFromLiteral("-u"), hxrt.StringFromLiteral("-c"), pythonServerScript(mode, count)), false)
	var port any = Std_parseInt(process.stdout.__hx_this.readLine())
	if port == nil {
		process.__hx_this.close()
		hxrt.Throw(haxe__Exception_thrown(hxrt.StringFromLiteral("failed to read server port")))
	}
	hx_obj_71 := map[string]any{}
	hx_obj_71["process"] = process
	hx_obj_71["port"] = port.(int)
	return hxrt.AnonymousFromMap(hx_obj_71)
}

type I__Main__TrackingInput interface {
	readByte() int
	readBytes(bytes *haxe__io__Bytes, targetPos int, requested int) int
	close()
	set_bigEndian(value bool) bool
	readAll(bufsize any) *haxe__io__Bytes
	readFullBytes(bytes *haxe__io__Bytes, pos int, len int)
	read(nbytes int) *haxe__io__Bytes
	readUntil(end int) *string
	readLine() *string
	readFloat() float64
	readDouble() float64
	readInt8() int
	readInt16() int
	readUInt16() int
	readInt24() int
	readUInt24() int
	readInt32() int
	readString(len int, encoding *haxe__io__Encoding) *string
}

type _Main__TrackingInput struct {
	*haxe__io__Input
	__hx_this        I__Main__TrackingInput
	reads            int
	wrongThread      bool
	afterReturnReads int
	requestReturned  bool
	mode             *string
	callerMarker     *sys__thread__Tls
	payload          *haxe__io__Bytes
	position         int
}

func (self *_Main__TrackingInput) HxFunctionReceiver() any {
	return self.__hx_this
}

func New__Main__TrackingInput(mode *string, payload *string) *_Main__TrackingInput {
	return hx_init__Main__TrackingInput(hx_alloc__Main__TrackingInput(nil), mode, payload)
}

func hx_alloc__Main__TrackingInput(owner any) *_Main__TrackingInput {
	self := &_Main__TrackingInput{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I__Main__TrackingInput)
	}
	self.haxe__io__Input = hx_alloc_haxe__io__Input(self.__hx_this)
	return self
}

func hx_init__Main__TrackingInput(self *_Main__TrackingInput, mode *string, payload *string) *_Main__TrackingInput {
	hx_init_haxe__io__Input(self.haxe__io__Input)
	self.position = 0
	self.requestReturned = false
	self.afterReturnReads = 0
	self.wrongThread = false
	self.reads = 0
	self.mode = mode
	self.payload = haxe__io__Bytes_ofString(payload, nil)
	self.callerMarker = New_sys__thread__Tls()
	func(hx_value_72 any) *string {
		if hx_value_72 == nil {
			var hx_zero_73 *string
			return hx_zero_73
		}
		return hx_value_72.(*string)
	}(self.callerMarker.__hx_this.set_value(hxrt.StringFromLiteral("source-caller")))
	return self
}

func (self *_Main__TrackingInput) readBytes(bytes *haxe__io__Bytes, targetPos int, requested int) int {
	self.reads = int(int32((self.reads + 1)))
	if !hxrt.StringEqualStringPtr(func(hx_value_74 any) *string {
		if hx_value_74 == nil {
			var hx_zero_75 *string
			return hx_zero_75
		}
		return hx_value_74.(*string)
	}(self.callerMarker.__hx_this.get_value()), hxrt.StringFromLiteral("source-caller")) {
		self.wrongThread = true
	}
	if self.requestReturned {
		self.afterReturnReads = int(int32((self.afterReturnReads + 1)))
	}
	_g := self.mode
	switch *hxrt.StdString(_g) {
	case *hxrt.StdString(hxrt.StringFromLiteral("early-eof")):
		if self.position > 0 {
			hxrt.Throw(haxe__Exception_thrown(New_haxe__io__Eof()))
		}
		bytes.b[targetPos] = 120
		bytes.__hx_rawValid = false
		self.position = int(int32((self.position + 1)))
		return 1
	case *hxrt.StdString(hxrt.StringFromLiteral("source-error")):
		if self.position > 0 {
			hxrt.Throw(haxe__Exception_thrown(hxrt.StringFromLiteral("source-exploded")))
		}
		bytes.b[targetPos] = 120
		bytes.__hx_rawValid = false
		self.position = int(int32((self.position + 1)))
		return 1
	case *hxrt.StdString(hxrt.StringFromLiteral("stream")):
		bytes.__hx_this.fill(targetPos, requested, 97)
		self.position = int((hxrt.Int32Wrap(self.position) + hxrt.Int32Wrap(requested)))
		return requested
	case *hxrt.StdString(hxrt.StringFromLiteral("zero")):
		return 0
	default:
		if self.position >= self.payload.length {
			hxrt.Throw(haxe__Exception_thrown(New_haxe__io__Eof()))
		}
		count := int((hxrt.Int32Wrap(self.payload.length) - hxrt.Int32Wrap(self.position)))
		if count > 2 {
			count = 2
		}
		if count > requested {
			count = requested
		}
		bytes.__hx_this.blit(targetPos, self.payload, self.position, count)
		self.position = int((hxrt.Int32Wrap(self.position) + hxrt.Int32Wrap(count)))
		return count
	}
}

func hx_structural_recover(value any) hxrt.Anonymous {
	switch source := value.(type) {
	case *_Main__TrackingInput:
		return func(hx_structural_source_511 *_Main__TrackingInput) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_511, func(name string) hxrt.AnonymousField {
				switch name {
				case "reads":
					return hxrt.AnonymousField{Value: hx_structural_source_511.reads, Found: true}
				case "wrongThread":
					return hxrt.AnonymousField{Value: hx_structural_source_511.wrongThread, Found: true}
				case "afterReturnReads":
					return hxrt.AnonymousField{Value: hx_structural_source_511.afterReturnReads, Found: true}
				case "requestReturned":
					return hxrt.AnonymousField{Value: hx_structural_source_511.requestReturned, Found: true}
				case "mode":
					return hxrt.AnonymousField{Value: hx_structural_source_511.mode, Found: true}
				case "callerMarker":
					return hxrt.AnonymousField{Value: hx_structural_source_511.callerMarker, Found: true}
				case "payload":
					return hxrt.AnonymousField{Value: hx_structural_source_511.payload, Found: true}
				case "position":
					return hxrt.AnonymousField{Value: hx_structural_source_511.position, Found: true}
				case "readBytes":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes, int, int) int] {
						hx_receiver_516 := hx_structural_source_511
						return hxrt.BoundFunction(func(bytes *haxe__io__Bytes, targetPos int, requested int) int {
							return hx_receiver_516.__hx_this.readBytes(bytes, targetPos, requested)
						}, func() any {
							return hx_receiver_516.__hx_this
						}, "readBytes")
					}(), Found: true}
				case "bigEndian":
					return hxrt.AnonymousField{Value: hx_structural_source_511.bigEndian, Found: true}
				case "readByte":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_517 := hx_structural_source_511
						return hxrt.BoundFunction(func() int {
							return hx_receiver_517.__hx_this.readByte()
						}, func() any {
							return hx_receiver_517.__hx_this
						}, "readByte")
					}(), Found: true}
				case "close":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func()] {
						hx_receiver_518 := hx_structural_source_511
						return hxrt.BoundFunction(func() {
							hx_receiver_518.__hx_this.close()
						}, func() any {
							return hx_receiver_518.__hx_this
						}, "close")
					}(), Found: true}
				case "set_bigEndian":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(bool) bool] {
						hx_receiver_519 := hx_structural_source_511
						return hxrt.BoundFunction(func(value bool) bool {
							return hx_receiver_519.__hx_this.set_bigEndian(value)
						}, func() any {
							return hx_receiver_519.__hx_this
						}, "set_bigEndian")
					}(), Found: true}
				case "readAll":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(any) *haxe__io__Bytes] {
						hx_receiver_520 := hx_structural_source_511
						return hxrt.BoundFunction(func(bufsize any) *haxe__io__Bytes {
							return hx_receiver_520.__hx_this.readAll(bufsize)
						}, func() any {
							return hx_receiver_520.__hx_this
						}, "readAll")
					}(), Found: true}
				case "readFullBytes":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes, int, int)] {
						hx_receiver_521 := hx_structural_source_511
						return hxrt.BoundFunction(func(bytes *haxe__io__Bytes, pos int, len int) {
							hx_receiver_521.__hx_this.readFullBytes(bytes, pos, len)
						}, func() any {
							return hx_receiver_521.__hx_this
						}, "readFullBytes")
					}(), Found: true}
				case "read":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int) *haxe__io__Bytes] {
						hx_receiver_522 := hx_structural_source_511
						return hxrt.BoundFunction(func(nbytes int) *haxe__io__Bytes {
							return hx_receiver_522.__hx_this.read(nbytes)
						}, func() any {
							return hx_receiver_522.__hx_this
						}, "read")
					}(), Found: true}
				case "readUntil":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int) *string] {
						hx_receiver_523 := hx_structural_source_511
						return hxrt.BoundFunction(func(end int) *string {
							return hx_receiver_523.__hx_this.readUntil(end)
						}, func() any {
							return hx_receiver_523.__hx_this
						}, "readUntil")
					}(), Found: true}
				case "readLine":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_524 := hx_structural_source_511
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_524.__hx_this.readLine()
						}, func() any {
							return hx_receiver_524.__hx_this
						}, "readLine")
					}(), Found: true}
				case "readFloat":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() float64] {
						hx_receiver_525 := hx_structural_source_511
						return hxrt.BoundFunction(func() float64 {
							return hx_receiver_525.__hx_this.readFloat()
						}, func() any {
							return hx_receiver_525.__hx_this
						}, "readFloat")
					}(), Found: true}
				case "readDouble":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() float64] {
						hx_receiver_526 := hx_structural_source_511
						return hxrt.BoundFunction(func() float64 {
							return hx_receiver_526.__hx_this.readDouble()
						}, func() any {
							return hx_receiver_526.__hx_this
						}, "readDouble")
					}(), Found: true}
				case "readInt8":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_527 := hx_structural_source_511
						return hxrt.BoundFunction(func() int {
							return hx_receiver_527.__hx_this.readInt8()
						}, func() any {
							return hx_receiver_527.__hx_this
						}, "readInt8")
					}(), Found: true}
				case "readInt16":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_528 := hx_structural_source_511
						return hxrt.BoundFunction(func() int {
							return hx_receiver_528.__hx_this.readInt16()
						}, func() any {
							return hx_receiver_528.__hx_this
						}, "readInt16")
					}(), Found: true}
				case "readUInt16":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_529 := hx_structural_source_511
						return hxrt.BoundFunction(func() int {
							return hx_receiver_529.__hx_this.readUInt16()
						}, func() any {
							return hx_receiver_529.__hx_this
						}, "readUInt16")
					}(), Found: true}
				case "readInt24":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_530 := hx_structural_source_511
						return hxrt.BoundFunction(func() int {
							return hx_receiver_530.__hx_this.readInt24()
						}, func() any {
							return hx_receiver_530.__hx_this
						}, "readInt24")
					}(), Found: true}
				case "readUInt24":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_531 := hx_structural_source_511
						return hxrt.BoundFunction(func() int {
							return hx_receiver_531.__hx_this.readUInt24()
						}, func() any {
							return hx_receiver_531.__hx_this
						}, "readUInt24")
					}(), Found: true}
				case "readInt32":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_532 := hx_structural_source_511
						return hxrt.BoundFunction(func() int {
							return hx_receiver_532.__hx_this.readInt32()
						}, func() any {
							return hx_receiver_532.__hx_this
						}, "readInt32")
					}(), Found: true}
				case "readString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, *haxe__io__Encoding) *string] {
						hx_receiver_533 := hx_structural_source_511
						return hxrt.BoundFunction(func(len int, encoding *haxe__io__Encoding) *string {
							return hx_receiver_533.__hx_this.readString(len, encoding)
						}, func() any {
							return hx_receiver_533.__hx_this
						}, "readString")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "requestReturned":
					hx_structural_source_511.requestReturned = func(hx_value_512 any) bool {
						if hx_value_512 == nil {
							var hx_zero_513 bool
							return hx_zero_513
						}
						return hx_value_512.(bool)
					}(incoming)
					return true
				case "position":
					hx_structural_source_511.position = func(hx_value_514 any) int {
						if hx_value_514 == nil {
							var hx_zero_515 int
							return hx_zero_515
						}
						return hx_value_514.(int)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe___Int64_____Int64:
		return func(hx_structural_source_534 *haxe___Int64_____Int64) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_534, func(name string) hxrt.AnonymousField {
				switch name {
				case "high":
					return hxrt.AnonymousField{Value: hx_structural_source_534.high, Found: true}
				case "low":
					return hxrt.AnonymousField{Value: hx_structural_source_534.low, Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "high":
					hx_structural_source_534.high = func(hx_value_535 any) int {
						if hx_value_535 == nil {
							var hx_zero_536 int
							return hx_zero_536
						}
						return hx_value_535.(int)
					}(incoming)
					return true
				case "low":
					hx_structural_source_534.low = func(hx_value_537 any) int {
						if hx_value_537 == nil {
							var hx_zero_538 int
							return hx_zero_538
						}
						return hx_value_537.(int)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *sys__net__SocketOutput:
		return func(hx_structural_source_539 *sys__net__SocketOutput) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_539, func(name string) hxrt.AnonymousField {
				switch name {
				case "handle":
					return hxrt.AnonymousField{Value: hx_structural_source_539.handle, Found: true}
				case "writeByte":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_542 := hx_structural_source_539
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_542.__hx_this.writeByte(value)
						}, func() any {
							return hx_receiver_542.__hx_this
						}, "writeByte")
					}(), Found: true}
				case "writeBytes":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes, int, int) int] {
						hx_receiver_543 := hx_structural_source_539
						return hxrt.BoundFunction(func(bytes *haxe__io__Bytes, pos int, length int) int {
							return hx_receiver_543.__hx_this.writeBytes(bytes, pos, length)
						}, func() any {
							return hx_receiver_543.__hx_this
						}, "writeBytes")
					}(), Found: true}
				case "flush":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func()] {
						hx_receiver_544 := hx_structural_source_539
						return hxrt.BoundFunction(func() {
							hx_receiver_544.__hx_this.flush()
						}, func() any {
							return hx_receiver_544.__hx_this
						}, "flush")
					}(), Found: true}
				case "close":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func()] {
						hx_receiver_545 := hx_structural_source_539
						return hxrt.BoundFunction(func() {
							hx_receiver_545.__hx_this.close()
						}, func() any {
							return hx_receiver_545.__hx_this
						}, "close")
					}(), Found: true}
				case "bigEndian":
					return hxrt.AnonymousField{Value: hx_structural_source_539.bigEndian, Found: true}
				case "set_bigEndian":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(bool) bool] {
						hx_receiver_546 := hx_structural_source_539
						return hxrt.BoundFunction(func(value bool) bool {
							return hx_receiver_546.__hx_this.set_bigEndian(value)
						}, func() any {
							return hx_receiver_546.__hx_this
						}, "set_bigEndian")
					}(), Found: true}
				case "write":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes)] {
						hx_receiver_547 := hx_structural_source_539
						return hxrt.BoundFunction(func(bytes *haxe__io__Bytes) {
							hx_receiver_547.__hx_this.write(bytes)
						}, func() any {
							return hx_receiver_547.__hx_this
						}, "write")
					}(), Found: true}
				case "writeFullBytes":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes, int, int)] {
						hx_receiver_548 := hx_structural_source_539
						return hxrt.BoundFunction(func(bytes *haxe__io__Bytes, pos int, len int) {
							hx_receiver_548.__hx_this.writeFullBytes(bytes, pos, len)
						}, func() any {
							return hx_receiver_548.__hx_this
						}, "writeFullBytes")
					}(), Found: true}
				case "writeFloat":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(float64)] {
						hx_receiver_549 := hx_structural_source_539
						return hxrt.BoundFunction(func(value float64) {
							hx_receiver_549.__hx_this.writeFloat(value)
						}, func() any {
							return hx_receiver_549.__hx_this
						}, "writeFloat")
					}(), Found: true}
				case "writeDouble":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(float64)] {
						hx_receiver_550 := hx_structural_source_539
						return hxrt.BoundFunction(func(value float64) {
							hx_receiver_550.__hx_this.writeDouble(value)
						}, func() any {
							return hx_receiver_550.__hx_this
						}, "writeDouble")
					}(), Found: true}
				case "writeInt8":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_551 := hx_structural_source_539
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_551.__hx_this.writeInt8(value)
						}, func() any {
							return hx_receiver_551.__hx_this
						}, "writeInt8")
					}(), Found: true}
				case "writeInt16":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_552 := hx_structural_source_539
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_552.__hx_this.writeInt16(value)
						}, func() any {
							return hx_receiver_552.__hx_this
						}, "writeInt16")
					}(), Found: true}
				case "writeUInt16":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_553 := hx_structural_source_539
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_553.__hx_this.writeUInt16(value)
						}, func() any {
							return hx_receiver_553.__hx_this
						}, "writeUInt16")
					}(), Found: true}
				case "writeInt24":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_554 := hx_structural_source_539
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_554.__hx_this.writeInt24(value)
						}, func() any {
							return hx_receiver_554.__hx_this
						}, "writeInt24")
					}(), Found: true}
				case "writeUInt24":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_555 := hx_structural_source_539
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_555.__hx_this.writeUInt24(value)
						}, func() any {
							return hx_receiver_555.__hx_this
						}, "writeUInt24")
					}(), Found: true}
				case "writeInt32":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_556 := hx_structural_source_539
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_556.__hx_this.writeInt32(value)
						}, func() any {
							return hx_receiver_556.__hx_this
						}, "writeInt32")
					}(), Found: true}
				case "prepare":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_557 := hx_structural_source_539
						return hxrt.BoundFunction(func(nbytes int) {
							hx_receiver_557.__hx_this.prepare(nbytes)
						}, func() any {
							return hx_receiver_557.__hx_this
						}, "prepare")
					}(), Found: true}
				case "writeInput":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Input, any)] {
						hx_receiver_558 := hx_structural_source_539
						return hxrt.BoundFunction(func(input *haxe__io__Input, bufsize any) {
							hx_receiver_558.__hx_this.writeInput(input, bufsize)
						}, func() any {
							return hx_receiver_558.__hx_this
						}, "writeInput")
					}(), Found: true}
				case "writeString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string, *haxe__io__Encoding)] {
						hx_receiver_559 := hx_structural_source_539
						return hxrt.BoundFunction(func(value *string, encoding *haxe__io__Encoding) {
							hx_receiver_559.__hx_this.writeString(value, encoding)
						}, func() any {
							return hx_receiver_559.__hx_this
						}, "writeString")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "handle":
					hx_structural_source_539.handle = func(hx_value_540 any) *hxrt.SocketHandle {
						if hx_value_540 == nil {
							var hx_zero_541 *hxrt.SocketHandle
							return hx_zero_541
						}
						return hx_value_540.(*hxrt.SocketHandle)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *sys__net__SocketInput:
		return func(hx_structural_source_560 *sys__net__SocketInput) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_560, func(name string) hxrt.AnonymousField {
				switch name {
				case "handle":
					return hxrt.AnonymousField{Value: hx_structural_source_560.handle, Found: true}
				case "readByte":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_563 := hx_structural_source_560
						return hxrt.BoundFunction(func() int {
							return hx_receiver_563.__hx_this.readByte()
						}, func() any {
							return hx_receiver_563.__hx_this
						}, "readByte")
					}(), Found: true}
				case "readBytes":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes, int, int) int] {
						hx_receiver_564 := hx_structural_source_560
						return hxrt.BoundFunction(func(bytes *haxe__io__Bytes, pos int, length int) int {
							return hx_receiver_564.__hx_this.readBytes(bytes, pos, length)
						}, func() any {
							return hx_receiver_564.__hx_this
						}, "readBytes")
					}(), Found: true}
				case "close":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func()] {
						hx_receiver_565 := hx_structural_source_560
						return hxrt.BoundFunction(func() {
							hx_receiver_565.__hx_this.close()
						}, func() any {
							return hx_receiver_565.__hx_this
						}, "close")
					}(), Found: true}
				case "bigEndian":
					return hxrt.AnonymousField{Value: hx_structural_source_560.bigEndian, Found: true}
				case "set_bigEndian":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(bool) bool] {
						hx_receiver_566 := hx_structural_source_560
						return hxrt.BoundFunction(func(value bool) bool {
							return hx_receiver_566.__hx_this.set_bigEndian(value)
						}, func() any {
							return hx_receiver_566.__hx_this
						}, "set_bigEndian")
					}(), Found: true}
				case "readAll":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(any) *haxe__io__Bytes] {
						hx_receiver_567 := hx_structural_source_560
						return hxrt.BoundFunction(func(bufsize any) *haxe__io__Bytes {
							return hx_receiver_567.__hx_this.readAll(bufsize)
						}, func() any {
							return hx_receiver_567.__hx_this
						}, "readAll")
					}(), Found: true}
				case "readFullBytes":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes, int, int)] {
						hx_receiver_568 := hx_structural_source_560
						return hxrt.BoundFunction(func(bytes *haxe__io__Bytes, pos int, len int) {
							hx_receiver_568.__hx_this.readFullBytes(bytes, pos, len)
						}, func() any {
							return hx_receiver_568.__hx_this
						}, "readFullBytes")
					}(), Found: true}
				case "read":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int) *haxe__io__Bytes] {
						hx_receiver_569 := hx_structural_source_560
						return hxrt.BoundFunction(func(nbytes int) *haxe__io__Bytes {
							return hx_receiver_569.__hx_this.read(nbytes)
						}, func() any {
							return hx_receiver_569.__hx_this
						}, "read")
					}(), Found: true}
				case "readUntil":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int) *string] {
						hx_receiver_570 := hx_structural_source_560
						return hxrt.BoundFunction(func(end int) *string {
							return hx_receiver_570.__hx_this.readUntil(end)
						}, func() any {
							return hx_receiver_570.__hx_this
						}, "readUntil")
					}(), Found: true}
				case "readLine":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_571 := hx_structural_source_560
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_571.__hx_this.readLine()
						}, func() any {
							return hx_receiver_571.__hx_this
						}, "readLine")
					}(), Found: true}
				case "readFloat":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() float64] {
						hx_receiver_572 := hx_structural_source_560
						return hxrt.BoundFunction(func() float64 {
							return hx_receiver_572.__hx_this.readFloat()
						}, func() any {
							return hx_receiver_572.__hx_this
						}, "readFloat")
					}(), Found: true}
				case "readDouble":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() float64] {
						hx_receiver_573 := hx_structural_source_560
						return hxrt.BoundFunction(func() float64 {
							return hx_receiver_573.__hx_this.readDouble()
						}, func() any {
							return hx_receiver_573.__hx_this
						}, "readDouble")
					}(), Found: true}
				case "readInt8":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_574 := hx_structural_source_560
						return hxrt.BoundFunction(func() int {
							return hx_receiver_574.__hx_this.readInt8()
						}, func() any {
							return hx_receiver_574.__hx_this
						}, "readInt8")
					}(), Found: true}
				case "readInt16":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_575 := hx_structural_source_560
						return hxrt.BoundFunction(func() int {
							return hx_receiver_575.__hx_this.readInt16()
						}, func() any {
							return hx_receiver_575.__hx_this
						}, "readInt16")
					}(), Found: true}
				case "readUInt16":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_576 := hx_structural_source_560
						return hxrt.BoundFunction(func() int {
							return hx_receiver_576.__hx_this.readUInt16()
						}, func() any {
							return hx_receiver_576.__hx_this
						}, "readUInt16")
					}(), Found: true}
				case "readInt24":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_577 := hx_structural_source_560
						return hxrt.BoundFunction(func() int {
							return hx_receiver_577.__hx_this.readInt24()
						}, func() any {
							return hx_receiver_577.__hx_this
						}, "readInt24")
					}(), Found: true}
				case "readUInt24":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_578 := hx_structural_source_560
						return hxrt.BoundFunction(func() int {
							return hx_receiver_578.__hx_this.readUInt24()
						}, func() any {
							return hx_receiver_578.__hx_this
						}, "readUInt24")
					}(), Found: true}
				case "readInt32":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_579 := hx_structural_source_560
						return hxrt.BoundFunction(func() int {
							return hx_receiver_579.__hx_this.readInt32()
						}, func() any {
							return hx_receiver_579.__hx_this
						}, "readInt32")
					}(), Found: true}
				case "readString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, *haxe__io__Encoding) *string] {
						hx_receiver_580 := hx_structural_source_560
						return hxrt.BoundFunction(func(len int, encoding *haxe__io__Encoding) *string {
							return hx_receiver_580.__hx_this.readString(len, encoding)
						}, func() any {
							return hx_receiver_580.__hx_this
						}, "readString")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "handle":
					hx_structural_source_560.handle = func(hx_value_561 any) *hxrt.SocketHandle {
						if hx_value_561 == nil {
							var hx_zero_562 *hxrt.SocketHandle
							return hx_zero_562
						}
						return hx_value_561.(*hxrt.SocketHandle)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *sys__net__Socket:
		return func(hx_structural_source_581 *sys__net__Socket) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_581, func(name string) hxrt.AnonymousField {
				switch name {
				case "input":
					return hxrt.AnonymousField{Value: hx_structural_source_581.input, Found: true}
				case "output":
					return hxrt.AnonymousField{Value: hx_structural_source_581.output, Found: true}
				case "custom":
					return hxrt.AnonymousField{Value: hx_structural_source_581.custom, Found: true}
				case "handle":
					return hxrt.AnonymousField{Value: hx_structural_source_581.handle, Found: true}
				case "close":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func()] {
						hx_receiver_584 := hx_structural_source_581
						return hxrt.BoundFunction(func() {
							hx_receiver_584.__hx_this.close()
						}, func() any {
							return hx_receiver_584.__hx_this
						}, "close")
					}(), Found: true}
				case "read":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_585 := hx_structural_source_581
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_585.__hx_this.read()
						}, func() any {
							return hx_receiver_585.__hx_this
						}, "read")
					}(), Found: true}
				case "write":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string)] {
						hx_receiver_586 := hx_structural_source_581
						return hxrt.BoundFunction(func(content *string) {
							hx_receiver_586.__hx_this.write(content)
						}, func() any {
							return hx_receiver_586.__hx_this
						}, "write")
					}(), Found: true}
				case "connect":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*sys__net__Host, int)] {
						hx_receiver_587 := hx_structural_source_581
						return hxrt.BoundFunction(func(host *sys__net__Host, port int) {
							hx_receiver_587.__hx_this.connect(host, port)
						}, func() any {
							return hx_receiver_587.__hx_this
						}, "connect")
					}(), Found: true}
				case "listen":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_588 := hx_structural_source_581
						return hxrt.BoundFunction(func(connections int) {
							hx_receiver_588.__hx_this.listen(connections)
						}, func() any {
							return hx_receiver_588.__hx_this
						}, "listen")
					}(), Found: true}
				case "shutdown":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(bool, bool)] {
						hx_receiver_589 := hx_structural_source_581
						return hxrt.BoundFunction(func(read bool, write bool) {
							hx_receiver_589.__hx_this.shutdown(read, write)
						}, func() any {
							return hx_receiver_589.__hx_this
						}, "shutdown")
					}(), Found: true}
				case "bind":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*sys__net__Host, int)] {
						hx_receiver_590 := hx_structural_source_581
						return hxrt.BoundFunction(func(host *sys__net__Host, port int) {
							hx_receiver_590.__hx_this.bind(host, port)
						}, func() any {
							return hx_receiver_590.__hx_this
						}, "bind")
					}(), Found: true}
				case "accept":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *sys__net__Socket] {
						hx_receiver_591 := hx_structural_source_581
						return hxrt.BoundFunction(func() *sys__net__Socket {
							return hx_receiver_591.__hx_this.accept()
						}, func() any {
							return hx_receiver_591.__hx_this
						}, "accept")
					}(), Found: true}
				case "peer":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() hxrt.Anonymous] {
						hx_receiver_592 := hx_structural_source_581
						return hxrt.BoundFunction(func() hxrt.Anonymous {
							return hx_receiver_592.__hx_this.peer()
						}, func() any {
							return hx_receiver_592.__hx_this
						}, "peer")
					}(), Found: true}
				case "host":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() hxrt.Anonymous] {
						hx_receiver_593 := hx_structural_source_581
						return hxrt.BoundFunction(func() hxrt.Anonymous {
							return hx_receiver_593.__hx_this.host()
						}, func() any {
							return hx_receiver_593.__hx_this
						}, "host")
					}(), Found: true}
				case "setTimeout":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(float64)] {
						hx_receiver_594 := hx_structural_source_581
						return hxrt.BoundFunction(func(timeout float64) {
							hx_receiver_594.__hx_this.setTimeout(timeout)
						}, func() any {
							return hx_receiver_594.__hx_this
						}, "setTimeout")
					}(), Found: true}
				case "waitForRead":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func()] {
						hx_receiver_595 := hx_structural_source_581
						return hxrt.BoundFunction(func() {
							hx_receiver_595.__hx_this.waitForRead()
						}, func() any {
							return hx_receiver_595.__hx_this
						}, "waitForRead")
					}(), Found: true}
				case "setBlocking":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(bool)] {
						hx_receiver_596 := hx_structural_source_581
						return hxrt.BoundFunction(func(blocking bool) {
							hx_receiver_596.__hx_this.setBlocking(blocking)
						}, func() any {
							return hx_receiver_596.__hx_this
						}, "setBlocking")
					}(), Found: true}
				case "setFastSend":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(bool)] {
						hx_receiver_597 := hx_structural_source_581
						return hxrt.BoundFunction(func(fastSend bool) {
							hx_receiver_597.__hx_this.setFastSend(fastSend)
						}, func() any {
							return hx_receiver_597.__hx_this
						}, "setFastSend")
					}(), Found: true}
				case "replaceHandle":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*hxrt.SocketHandle)] {
						hx_receiver_598 := hx_structural_source_581
						return hxrt.BoundFunction(func(next *hxrt.SocketHandle) {
							hx_receiver_598.__hx_this.replaceHandle(next)
						}, func() any {
							return hx_receiver_598.__hx_this
						}, "replaceHandle")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "custom":
					hx_structural_source_581.custom = incoming
					return true
				case "handle":
					hx_structural_source_581.handle = func(hx_value_582 any) *hxrt.SocketHandle {
						if hx_value_582 == nil {
							var hx_zero_583 *hxrt.SocketHandle
							return hx_zero_583
						}
						return hx_value_582.(*hxrt.SocketHandle)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *sys__net__Host:
		return func(hx_structural_source_599 *sys__net__Host) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_599, func(name string) hxrt.AnonymousField {
				switch name {
				case "host":
					return hxrt.AnonymousField{Value: hx_structural_source_599.host, Found: true}
				case "ip":
					return hxrt.AnonymousField{Value: hx_structural_source_599.ip, Found: true}
				case "toString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_600 := hx_structural_source_599
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_600.__hx_this.toString()
						}, func() any {
							return hx_receiver_600.__hx_this
						}, "toString")
					}(), Found: true}
				case "reverse":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_601 := hx_structural_source_599
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_601.__hx_this.reverse()
						}, func() any {
							return hx_receiver_601.__hx_this
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
	case *sys__io___Process__ProcessOutput:
		return func(hx_structural_source_602 *sys__io___Process__ProcessOutput) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_602, func(name string) hxrt.AnonymousField {
				switch name {
				case "handle":
					return hxrt.AnonymousField{Value: hx_structural_source_602.handle, Found: true}
				case "readByte":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_605 := hx_structural_source_602
						return hxrt.BoundFunction(func() int {
							return hx_receiver_605.__hx_this.readByte()
						}, func() any {
							return hx_receiver_605.__hx_this
						}, "readByte")
					}(), Found: true}
				case "readBytes":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes, int, int) int] {
						hx_receiver_606 := hx_structural_source_602
						return hxrt.BoundFunction(func(bytes *haxe__io__Bytes, pos int, length int) int {
							return hx_receiver_606.__hx_this.readBytes(bytes, pos, length)
						}, func() any {
							return hx_receiver_606.__hx_this
						}, "readBytes")
					}(), Found: true}
				case "close":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func()] {
						hx_receiver_607 := hx_structural_source_602
						return hxrt.BoundFunction(func() {
							hx_receiver_607.__hx_this.close()
						}, func() any {
							return hx_receiver_607.__hx_this
						}, "close")
					}(), Found: true}
				case "bigEndian":
					return hxrt.AnonymousField{Value: hx_structural_source_602.bigEndian, Found: true}
				case "set_bigEndian":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(bool) bool] {
						hx_receiver_608 := hx_structural_source_602
						return hxrt.BoundFunction(func(value bool) bool {
							return hx_receiver_608.__hx_this.set_bigEndian(value)
						}, func() any {
							return hx_receiver_608.__hx_this
						}, "set_bigEndian")
					}(), Found: true}
				case "readAll":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(any) *haxe__io__Bytes] {
						hx_receiver_609 := hx_structural_source_602
						return hxrt.BoundFunction(func(bufsize any) *haxe__io__Bytes {
							return hx_receiver_609.__hx_this.readAll(bufsize)
						}, func() any {
							return hx_receiver_609.__hx_this
						}, "readAll")
					}(), Found: true}
				case "readFullBytes":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes, int, int)] {
						hx_receiver_610 := hx_structural_source_602
						return hxrt.BoundFunction(func(bytes *haxe__io__Bytes, pos int, len int) {
							hx_receiver_610.__hx_this.readFullBytes(bytes, pos, len)
						}, func() any {
							return hx_receiver_610.__hx_this
						}, "readFullBytes")
					}(), Found: true}
				case "read":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int) *haxe__io__Bytes] {
						hx_receiver_611 := hx_structural_source_602
						return hxrt.BoundFunction(func(nbytes int) *haxe__io__Bytes {
							return hx_receiver_611.__hx_this.read(nbytes)
						}, func() any {
							return hx_receiver_611.__hx_this
						}, "read")
					}(), Found: true}
				case "readUntil":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int) *string] {
						hx_receiver_612 := hx_structural_source_602
						return hxrt.BoundFunction(func(end int) *string {
							return hx_receiver_612.__hx_this.readUntil(end)
						}, func() any {
							return hx_receiver_612.__hx_this
						}, "readUntil")
					}(), Found: true}
				case "readLine":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_613 := hx_structural_source_602
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_613.__hx_this.readLine()
						}, func() any {
							return hx_receiver_613.__hx_this
						}, "readLine")
					}(), Found: true}
				case "readFloat":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() float64] {
						hx_receiver_614 := hx_structural_source_602
						return hxrt.BoundFunction(func() float64 {
							return hx_receiver_614.__hx_this.readFloat()
						}, func() any {
							return hx_receiver_614.__hx_this
						}, "readFloat")
					}(), Found: true}
				case "readDouble":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() float64] {
						hx_receiver_615 := hx_structural_source_602
						return hxrt.BoundFunction(func() float64 {
							return hx_receiver_615.__hx_this.readDouble()
						}, func() any {
							return hx_receiver_615.__hx_this
						}, "readDouble")
					}(), Found: true}
				case "readInt8":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_616 := hx_structural_source_602
						return hxrt.BoundFunction(func() int {
							return hx_receiver_616.__hx_this.readInt8()
						}, func() any {
							return hx_receiver_616.__hx_this
						}, "readInt8")
					}(), Found: true}
				case "readInt16":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_617 := hx_structural_source_602
						return hxrt.BoundFunction(func() int {
							return hx_receiver_617.__hx_this.readInt16()
						}, func() any {
							return hx_receiver_617.__hx_this
						}, "readInt16")
					}(), Found: true}
				case "readUInt16":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_618 := hx_structural_source_602
						return hxrt.BoundFunction(func() int {
							return hx_receiver_618.__hx_this.readUInt16()
						}, func() any {
							return hx_receiver_618.__hx_this
						}, "readUInt16")
					}(), Found: true}
				case "readInt24":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_619 := hx_structural_source_602
						return hxrt.BoundFunction(func() int {
							return hx_receiver_619.__hx_this.readInt24()
						}, func() any {
							return hx_receiver_619.__hx_this
						}, "readInt24")
					}(), Found: true}
				case "readUInt24":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_620 := hx_structural_source_602
						return hxrt.BoundFunction(func() int {
							return hx_receiver_620.__hx_this.readUInt24()
						}, func() any {
							return hx_receiver_620.__hx_this
						}, "readUInt24")
					}(), Found: true}
				case "readInt32":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_621 := hx_structural_source_602
						return hxrt.BoundFunction(func() int {
							return hx_receiver_621.__hx_this.readInt32()
						}, func() any {
							return hx_receiver_621.__hx_this
						}, "readInt32")
					}(), Found: true}
				case "readString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, *haxe__io__Encoding) *string] {
						hx_receiver_622 := hx_structural_source_602
						return hxrt.BoundFunction(func(len int, encoding *haxe__io__Encoding) *string {
							return hx_receiver_622.__hx_this.readString(len, encoding)
						}, func() any {
							return hx_receiver_622.__hx_this
						}, "readString")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "handle":
					hx_structural_source_602.handle = func(hx_value_603 any) *hxrt.ProcessOutput {
						if hx_value_603 == nil {
							var hx_zero_604 *hxrt.ProcessOutput
							return hx_zero_604
						}
						return hx_value_603.(*hxrt.ProcessOutput)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *sys__io___Process__ProcessInput:
		return func(hx_structural_source_623 *sys__io___Process__ProcessInput) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_623, func(name string) hxrt.AnonymousField {
				switch name {
				case "handle":
					return hxrt.AnonymousField{Value: hx_structural_source_623.handle, Found: true}
				case "writeByte":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_626 := hx_structural_source_623
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_626.__hx_this.writeByte(value)
						}, func() any {
							return hx_receiver_626.__hx_this
						}, "writeByte")
					}(), Found: true}
				case "writeBytes":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes, int, int) int] {
						hx_receiver_627 := hx_structural_source_623
						return hxrt.BoundFunction(func(bytes *haxe__io__Bytes, pos int, length int) int {
							return hx_receiver_627.__hx_this.writeBytes(bytes, pos, length)
						}, func() any {
							return hx_receiver_627.__hx_this
						}, "writeBytes")
					}(), Found: true}
				case "flush":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func()] {
						hx_receiver_628 := hx_structural_source_623
						return hxrt.BoundFunction(func() {
							hx_receiver_628.__hx_this.flush()
						}, func() any {
							return hx_receiver_628.__hx_this
						}, "flush")
					}(), Found: true}
				case "close":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func()] {
						hx_receiver_629 := hx_structural_source_623
						return hxrt.BoundFunction(func() {
							hx_receiver_629.__hx_this.close()
						}, func() any {
							return hx_receiver_629.__hx_this
						}, "close")
					}(), Found: true}
				case "bigEndian":
					return hxrt.AnonymousField{Value: hx_structural_source_623.bigEndian, Found: true}
				case "set_bigEndian":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(bool) bool] {
						hx_receiver_630 := hx_structural_source_623
						return hxrt.BoundFunction(func(value bool) bool {
							return hx_receiver_630.__hx_this.set_bigEndian(value)
						}, func() any {
							return hx_receiver_630.__hx_this
						}, "set_bigEndian")
					}(), Found: true}
				case "write":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes)] {
						hx_receiver_631 := hx_structural_source_623
						return hxrt.BoundFunction(func(bytes *haxe__io__Bytes) {
							hx_receiver_631.__hx_this.write(bytes)
						}, func() any {
							return hx_receiver_631.__hx_this
						}, "write")
					}(), Found: true}
				case "writeFullBytes":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes, int, int)] {
						hx_receiver_632 := hx_structural_source_623
						return hxrt.BoundFunction(func(bytes *haxe__io__Bytes, pos int, len int) {
							hx_receiver_632.__hx_this.writeFullBytes(bytes, pos, len)
						}, func() any {
							return hx_receiver_632.__hx_this
						}, "writeFullBytes")
					}(), Found: true}
				case "writeFloat":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(float64)] {
						hx_receiver_633 := hx_structural_source_623
						return hxrt.BoundFunction(func(value float64) {
							hx_receiver_633.__hx_this.writeFloat(value)
						}, func() any {
							return hx_receiver_633.__hx_this
						}, "writeFloat")
					}(), Found: true}
				case "writeDouble":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(float64)] {
						hx_receiver_634 := hx_structural_source_623
						return hxrt.BoundFunction(func(value float64) {
							hx_receiver_634.__hx_this.writeDouble(value)
						}, func() any {
							return hx_receiver_634.__hx_this
						}, "writeDouble")
					}(), Found: true}
				case "writeInt8":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_635 := hx_structural_source_623
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_635.__hx_this.writeInt8(value)
						}, func() any {
							return hx_receiver_635.__hx_this
						}, "writeInt8")
					}(), Found: true}
				case "writeInt16":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_636 := hx_structural_source_623
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_636.__hx_this.writeInt16(value)
						}, func() any {
							return hx_receiver_636.__hx_this
						}, "writeInt16")
					}(), Found: true}
				case "writeUInt16":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_637 := hx_structural_source_623
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_637.__hx_this.writeUInt16(value)
						}, func() any {
							return hx_receiver_637.__hx_this
						}, "writeUInt16")
					}(), Found: true}
				case "writeInt24":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_638 := hx_structural_source_623
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_638.__hx_this.writeInt24(value)
						}, func() any {
							return hx_receiver_638.__hx_this
						}, "writeInt24")
					}(), Found: true}
				case "writeUInt24":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_639 := hx_structural_source_623
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_639.__hx_this.writeUInt24(value)
						}, func() any {
							return hx_receiver_639.__hx_this
						}, "writeUInt24")
					}(), Found: true}
				case "writeInt32":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_640 := hx_structural_source_623
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_640.__hx_this.writeInt32(value)
						}, func() any {
							return hx_receiver_640.__hx_this
						}, "writeInt32")
					}(), Found: true}
				case "prepare":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_641 := hx_structural_source_623
						return hxrt.BoundFunction(func(nbytes int) {
							hx_receiver_641.__hx_this.prepare(nbytes)
						}, func() any {
							return hx_receiver_641.__hx_this
						}, "prepare")
					}(), Found: true}
				case "writeInput":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Input, any)] {
						hx_receiver_642 := hx_structural_source_623
						return hxrt.BoundFunction(func(input *haxe__io__Input, bufsize any) {
							hx_receiver_642.__hx_this.writeInput(input, bufsize)
						}, func() any {
							return hx_receiver_642.__hx_this
						}, "writeInput")
					}(), Found: true}
				case "writeString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string, *haxe__io__Encoding)] {
						hx_receiver_643 := hx_structural_source_623
						return hxrt.BoundFunction(func(value *string, encoding *haxe__io__Encoding) {
							hx_receiver_643.__hx_this.writeString(value, encoding)
						}, func() any {
							return hx_receiver_643.__hx_this
						}, "writeString")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "handle":
					hx_structural_source_623.handle = func(hx_value_624 any) *hxrt.ProcessInput {
						if hx_value_624 == nil {
							var hx_zero_625 *hxrt.ProcessInput
							return hx_zero_625
						}
						return hx_value_624.(*hxrt.ProcessInput)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *sys__io__Process:
		return func(hx_structural_source_644 *sys__io__Process) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_644, func(name string) hxrt.AnonymousField {
				switch name {
				case "stdout":
					return hxrt.AnonymousField{Value: hx_structural_source_644.stdout, Found: true}
				case "stderr":
					return hxrt.AnonymousField{Value: hx_structural_source_644.stderr, Found: true}
				case "stdin":
					return hxrt.AnonymousField{Value: hx_structural_source_644.stdin, Found: true}
				case "handle":
					return hxrt.AnonymousField{Value: hx_structural_source_644.handle, Found: true}
				case "getPid":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_647 := hx_structural_source_644
						return hxrt.BoundFunction(func() int {
							return hx_receiver_647.__hx_this.getPid()
						}, func() any {
							return hx_receiver_647.__hx_this
						}, "getPid")
					}(), Found: true}
				case "exitCode":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(any) any] {
						hx_receiver_648 := hx_structural_source_644
						return hxrt.BoundFunction(func(block any) any {
							if block == nil {
								block = true
							}
							return hx_receiver_648.__hx_this.exitCode(func(hx_value_649 any) bool {
								if hx_value_649 == nil {
									var hx_zero_650 bool
									return hx_zero_650
								}
								return hx_value_649.(bool)
							}(any(block)))
						}, func() any {
							return hx_receiver_648.__hx_this
						}, "exitCode")
					}(), Found: true}
				case "close":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func()] {
						hx_receiver_651 := hx_structural_source_644
						return hxrt.BoundFunction(func() {
							hx_receiver_651.__hx_this.close()
						}, func() any {
							return hx_receiver_651.__hx_this
						}, "close")
					}(), Found: true}
				case "kill":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func()] {
						hx_receiver_652 := hx_structural_source_644
						return hxrt.BoundFunction(func() {
							hx_receiver_652.__hx_this.kill()
						}, func() any {
							return hx_receiver_652.__hx_this
						}, "kill")
					}(), Found: true}
				case "requireHandle":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *hxrt.Process] {
						hx_receiver_653 := hx_structural_source_644
						return hxrt.BoundFunction(func() *hxrt.Process {
							return hx_receiver_653.__hx_this.requireHandle()
						}, func() any {
							return hx_receiver_653.__hx_this
						}, "requireHandle")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "handle":
					hx_structural_source_644.handle = func(hx_value_645 any) *hxrt.Process {
						if hx_value_645 == nil {
							var hx_zero_646 *hxrt.Process
							return hx_zero_646
						}
						return hx_value_645.(*hxrt.Process)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *sys__Http:
		return func(hx_structural_source_654 *sys__Http) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_654, func(name string) hxrt.AnonymousField {
				switch name {
				case "noShutdown":
					return hxrt.AnonymousField{Value: hx_structural_source_654.noShutdown, Found: true}
				case "cnxTimeout":
					return hxrt.AnonymousField{Value: hx_structural_source_654.cnxTimeout, Found: true}
				case "responseHeaders":
					return hxrt.AnonymousField{Value: hx_structural_source_654.responseHeaders, Found: true}
				case "responseHeadersSameKey":
					return hxrt.AnonymousField{Value: hx_structural_source_654.responseHeadersSameKey, Found: true}
				case "file":
					return hxrt.AnonymousField{Value: hx_structural_source_654.file, Found: true}
				case "request":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(any)] {
						hx_receiver_661 := hx_structural_source_654
						return hxrt.BoundFunction(func(post any) {
							hx_receiver_661.__hx_this.request(post)
						}, func() any {
							return hx_receiver_661.__hx_this
						}, "request")
					}(), Found: true}
				case "fileTransfert":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string, *string, *haxe__io__Input, int, *string)] {
						hx_receiver_662 := hx_structural_source_654
						return hxrt.BoundFunction(func(argname *string, filename *string, file *haxe__io__Input, size int, mimeType *string) {
							if mimeType == nil {
								mimeType = hxrt.StringFromLiteral("application/octet-stream")
							}
							hx_receiver_662.__hx_this.fileTransfert(argname, filename, file, size, mimeType)
						}, func() any {
							return hx_receiver_662.__hx_this
						}, "fileTransfert")
					}(), Found: true}
				case "fileTransfer":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string, *string, *haxe__io__Input, int, *string)] {
						hx_receiver_663 := hx_structural_source_654
						return hxrt.BoundFunction(func(argname *string, filename *string, file *haxe__io__Input, size int, mimeType *string) {
							if mimeType == nil {
								mimeType = hxrt.StringFromLiteral("application/octet-stream")
							}
							hx_receiver_663.__hx_this.fileTransfer(argname, filename, file, size, mimeType)
						}, func() any {
							return hx_receiver_663.__hx_this
						}, "fileTransfer")
					}(), Found: true}
				case "customRequest":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(bool, *haxe__io__Output, *sys__net__Socket, *string)] {
						hx_receiver_664 := hx_structural_source_654
						return hxrt.BoundFunction(func(post bool, api *haxe__io__Output, sock *sys__net__Socket, method *string) {
							hx_receiver_664.__hx_this.customRequest(post, api, sock, method)
						}, func() any {
							return hx_receiver_664.__hx_this
						}, "customRequest")
					}(), Found: true}
				case "getResponseHeaderValues":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string) *hxrt.Array] {
						hx_receiver_665 := hx_structural_source_654
						return hxrt.BoundFunction(func(key *string) *hxrt.Array {
							return hx_receiver_665.__hx_this.getResponseHeaderValues(key)
						}, func() any {
							return hx_receiver_665.__hx_this
						}, "getResponseHeaderValues")
					}(), Found: true}
				case "requestWith":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(bool, *haxe__io__Output, *sys__net__Socket, *string)] {
						hx_receiver_666 := hx_structural_source_654
						return hxrt.BoundFunction(func(post bool, api *haxe__io__Output, sock *sys__net__Socket, method *string) {
							hx_receiver_666.__hx_this.requestWith(post, api, sock, method)
						}, func() any {
							return hx_receiver_666.__hx_this
						}, "requestWith")
					}(), Found: true}
				case "pumpUpload":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*hxrt.HttpExchange, hxrt.Anonymous) hxrt.Anonymous] {
						hx_receiver_667 := hx_structural_source_654
						return hxrt.BoundFunction(func(exchange *hxrt.HttpExchange, upload hxrt.Anonymous) hxrt.Anonymous {
							return hx_receiver_667.__hx_this.pumpUpload(exchange, upload)
						}, func() any {
							return hx_receiver_667.__hx_this
						}, "pumpUpload")
					}(), Found: true}
				case "handleDataRequest":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(bool, *haxe__io__Output, *string)] {
						hx_receiver_668 := hx_structural_source_654
						return hxrt.BoundFunction(func(post bool, api *haxe__io__Output, method *string) {
							hx_receiver_668.__hx_this.handleDataRequest(post, api, method)
						}, func() any {
							return hx_receiver_668.__hx_this
						}, "handleDataRequest")
					}(), Found: true}
				case "recordResponseHeaders":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*hxrt.HttpExchange)] {
						hx_receiver_669 := hx_structural_source_654
						return hxrt.BoundFunction(func(exchange *hxrt.HttpExchange) {
							hx_receiver_669.__hx_this.recordResponseHeaders(exchange)
						}, func() any {
							return hx_receiver_669.__hx_this
						}, "recordResponseHeaders")
					}(), Found: true}
				case "resetResponseHeaders":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func()] {
						hx_receiver_670 := hx_structural_source_654
						return hxrt.BoundFunction(func() {
							hx_receiver_670.__hx_this.resetResponseHeaders()
						}, func() any {
							return hx_receiver_670.__hx_this
						}, "resetResponseHeaders")
					}(), Found: true}
				case "encodedParameters":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_671 := hx_structural_source_654
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_671.__hx_this.encodedParameters()
						}, func() any {
							return hx_receiver_671.__hx_this
						}, "encodedParameters")
					}(), Found: true}
				case "url":
					return hxrt.AnonymousField{Value: hx_structural_source_654.url, Found: true}
				case "responseBytes":
					return hxrt.AnonymousField{Value: hx_structural_source_654.responseBytes, Found: true}
				case "responseAsString":
					return hxrt.AnonymousField{Value: hx_structural_source_654.responseAsString, Found: true}
				case "postData":
					return hxrt.AnonymousField{Value: hx_structural_source_654.postData, Found: true}
				case "postBytes":
					return hxrt.AnonymousField{Value: hx_structural_source_654.postBytes, Found: true}
				case "headers":
					return hxrt.AnonymousField{Value: hx_structural_source_654.headers, Found: true}
				case "params":
					return hxrt.AnonymousField{Value: hx_structural_source_654.params, Found: true}
				case "emptyOnData":
					return hxrt.AnonymousField{Value: hx_structural_source_654.emptyOnData, Found: true}
				case "setHeader":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string, *string)] {
						hx_receiver_682 := hx_structural_source_654
						return hxrt.BoundFunction(func(name *string, value *string) {
							hx_receiver_682.__hx_this.setHeader(name, value)
						}, func() any {
							return hx_receiver_682.__hx_this
						}, "setHeader")
					}(), Found: true}
				case "addHeader":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string, *string)] {
						hx_receiver_683 := hx_structural_source_654
						return hxrt.BoundFunction(func(header *string, value *string) {
							hx_receiver_683.__hx_this.addHeader(header, value)
						}, func() any {
							return hx_receiver_683.__hx_this
						}, "addHeader")
					}(), Found: true}
				case "setParameter":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string, *string)] {
						hx_receiver_684 := hx_structural_source_654
						return hxrt.BoundFunction(func(name *string, value *string) {
							hx_receiver_684.__hx_this.setParameter(name, value)
						}, func() any {
							return hx_receiver_684.__hx_this
						}, "setParameter")
					}(), Found: true}
				case "addParameter":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string, *string)] {
						hx_receiver_685 := hx_structural_source_654
						return hxrt.BoundFunction(func(name *string, value *string) {
							hx_receiver_685.__hx_this.addParameter(name, value)
						}, func() any {
							return hx_receiver_685.__hx_this
						}, "addParameter")
					}(), Found: true}
				case "setPostData":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string)] {
						hx_receiver_686 := hx_structural_source_654
						return hxrt.BoundFunction(func(data *string) {
							hx_receiver_686.__hx_this.setPostData(data)
						}, func() any {
							return hx_receiver_686.__hx_this
						}, "setPostData")
					}(), Found: true}
				case "setPostBytes":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes)] {
						hx_receiver_687 := hx_structural_source_654
						return hxrt.BoundFunction(func(data *haxe__io__Bytes) {
							hx_receiver_687.__hx_this.setPostBytes(data)
						}, func() any {
							return hx_receiver_687.__hx_this
						}, "setPostBytes")
					}(), Found: true}
				case "onData":
					return hxrt.AnonymousField{Value: hxrt.ReadDynamicFunction(hx_structural_source_654.onData), Found: true}
				case "onBytes":
					return hxrt.AnonymousField{Value: hxrt.ReadDynamicFunction(hx_structural_source_654.onBytes), Found: true}
				case "onError":
					return hxrt.AnonymousField{Value: hxrt.ReadDynamicFunction(hx_structural_source_654.onError), Found: true}
				case "onStatus":
					return hxrt.AnonymousField{Value: hxrt.ReadDynamicFunction(hx_structural_source_654.onStatus), Found: true}
				case "hasOnData":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() bool] {
						hx_receiver_688 := hx_structural_source_654
						return hxrt.BoundFunction(func() bool {
							return hx_receiver_688.__hx_this.hasOnData()
						}, func() any {
							return hx_receiver_688.__hx_this
						}, "hasOnData")
					}(), Found: true}
				case "success":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes)] {
						hx_receiver_689 := hx_structural_source_654
						return hxrt.BoundFunction(func(data *haxe__io__Bytes) {
							hx_receiver_689.__hx_this.success(data)
						}, func() any {
							return hx_receiver_689.__hx_this
						}, "success")
					}(), Found: true}
				case "get_responseData":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_690 := hx_structural_source_654
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_690.__hx_this.get_responseData()
						}, func() any {
							return hx_receiver_690.__hx_this
						}, "get_responseData")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "noShutdown":
					hx_structural_source_654.noShutdown = func(hx_value_655 any) bool {
						if hx_value_655 == nil {
							var hx_zero_656 bool
							return hx_zero_656
						}
						return hx_value_655.(bool)
					}(incoming)
					return true
				case "cnxTimeout":
					hx_structural_source_654.cnxTimeout = func(hx_value_657 any) float64 {
						if hx_value_657 == nil {
							var hx_zero_658 float64
							return hx_zero_658
						}
						return hx_value_657.(float64)
					}(incoming)
					return true
				case "responseHeaders":
					hx_structural_source_654.responseHeaders = func(hx_value_659 any) *haxe__ds__StringMap {
						if hx_value_659 == nil {
							var hx_zero_660 *haxe__ds__StringMap
							return hx_zero_660
						}
						return hx_value_659.(*haxe__ds__StringMap)
					}(incoming)
					return true
				case "responseHeadersSameKey":
					hx_structural_source_654.responseHeadersSameKey = hx_recover_class_5(incoming)
					return true
				case "file":
					hx_structural_source_654.file = hx_structural_recover(incoming)
					return true
				case "url":
					hx_structural_source_654.url = func(hx_value_672 any) *string {
						if hx_value_672 == nil {
							var hx_zero_673 *string
							return hx_zero_673
						}
						return hx_value_672.(*string)
					}(incoming)
					return true
				case "responseAsString":
					hx_structural_source_654.responseAsString = func(hx_value_674 any) *string {
						if hx_value_674 == nil {
							var hx_zero_675 *string
							return hx_zero_675
						}
						return hx_value_674.(*string)
					}(incoming)
					return true
				case "postData":
					hx_structural_source_654.postData = func(hx_value_676 any) *string {
						if hx_value_676 == nil {
							var hx_zero_677 *string
							return hx_zero_677
						}
						return hx_value_676.(*string)
					}(incoming)
					return true
				case "postBytes":
					hx_structural_source_654.postBytes = hx_recover_class_7(incoming)
					return true
				case "headers":
					hx_structural_source_654.headers = func(hx_value_678 any) *hxrt.Array {
						if hx_value_678 == nil {
							var hx_zero_679 *hxrt.Array
							return hx_zero_679
						}
						return hx_value_678.(*hxrt.Array)
					}(incoming)
					return true
				case "params":
					hx_structural_source_654.params = func(hx_value_680 any) *hxrt.Array {
						if hx_value_680 == nil {
							var hx_zero_681 *hxrt.Array
							return hx_zero_681
						}
						return hx_value_680.(*hxrt.Array)
					}(incoming)
					return true
				case "onData":
					hx_structural_source_654.onData = hx_recover_callable_3(incoming)
					return true
				case "onBytes":
					hx_structural_source_654.onBytes = hx_recover_callable_4(incoming)
					return true
				case "onError":
					hx_structural_source_654.onError = hx_recover_callable_3(incoming)
					return true
				case "onStatus":
					hx_structural_source_654.onStatus = hx_recover_callable_5(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe__io__Output:
		return func(hx_structural_source_691 *haxe__io__Output) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_691, func(name string) hxrt.AnonymousField {
				switch name {
				case "bigEndian":
					return hxrt.AnonymousField{Value: hx_structural_source_691.bigEndian, Found: true}
				case "writeByte":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_692 := hx_structural_source_691
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_692.__hx_this.writeByte(value)
						}, func() any {
							return hx_receiver_692.__hx_this
						}, "writeByte")
					}(), Found: true}
				case "writeBytes":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes, int, int) int] {
						hx_receiver_693 := hx_structural_source_691
						return hxrt.BoundFunction(func(bytes *haxe__io__Bytes, pos int, len int) int {
							return hx_receiver_693.__hx_this.writeBytes(bytes, pos, len)
						}, func() any {
							return hx_receiver_693.__hx_this
						}, "writeBytes")
					}(), Found: true}
				case "flush":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func()] {
						hx_receiver_694 := hx_structural_source_691
						return hxrt.BoundFunction(func() {
							hx_receiver_694.__hx_this.flush()
						}, func() any {
							return hx_receiver_694.__hx_this
						}, "flush")
					}(), Found: true}
				case "close":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func()] {
						hx_receiver_695 := hx_structural_source_691
						return hxrt.BoundFunction(func() {
							hx_receiver_695.__hx_this.close()
						}, func() any {
							return hx_receiver_695.__hx_this
						}, "close")
					}(), Found: true}
				case "set_bigEndian":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(bool) bool] {
						hx_receiver_696 := hx_structural_source_691
						return hxrt.BoundFunction(func(value bool) bool {
							return hx_receiver_696.__hx_this.set_bigEndian(value)
						}, func() any {
							return hx_receiver_696.__hx_this
						}, "set_bigEndian")
					}(), Found: true}
				case "write":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes)] {
						hx_receiver_697 := hx_structural_source_691
						return hxrt.BoundFunction(func(bytes *haxe__io__Bytes) {
							hx_receiver_697.__hx_this.write(bytes)
						}, func() any {
							return hx_receiver_697.__hx_this
						}, "write")
					}(), Found: true}
				case "writeFullBytes":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes, int, int)] {
						hx_receiver_698 := hx_structural_source_691
						return hxrt.BoundFunction(func(bytes *haxe__io__Bytes, pos int, len int) {
							hx_receiver_698.__hx_this.writeFullBytes(bytes, pos, len)
						}, func() any {
							return hx_receiver_698.__hx_this
						}, "writeFullBytes")
					}(), Found: true}
				case "writeFloat":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(float64)] {
						hx_receiver_699 := hx_structural_source_691
						return hxrt.BoundFunction(func(value float64) {
							hx_receiver_699.__hx_this.writeFloat(value)
						}, func() any {
							return hx_receiver_699.__hx_this
						}, "writeFloat")
					}(), Found: true}
				case "writeDouble":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(float64)] {
						hx_receiver_700 := hx_structural_source_691
						return hxrt.BoundFunction(func(value float64) {
							hx_receiver_700.__hx_this.writeDouble(value)
						}, func() any {
							return hx_receiver_700.__hx_this
						}, "writeDouble")
					}(), Found: true}
				case "writeInt8":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_701 := hx_structural_source_691
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_701.__hx_this.writeInt8(value)
						}, func() any {
							return hx_receiver_701.__hx_this
						}, "writeInt8")
					}(), Found: true}
				case "writeInt16":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_702 := hx_structural_source_691
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_702.__hx_this.writeInt16(value)
						}, func() any {
							return hx_receiver_702.__hx_this
						}, "writeInt16")
					}(), Found: true}
				case "writeUInt16":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_703 := hx_structural_source_691
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_703.__hx_this.writeUInt16(value)
						}, func() any {
							return hx_receiver_703.__hx_this
						}, "writeUInt16")
					}(), Found: true}
				case "writeInt24":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_704 := hx_structural_source_691
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_704.__hx_this.writeInt24(value)
						}, func() any {
							return hx_receiver_704.__hx_this
						}, "writeInt24")
					}(), Found: true}
				case "writeUInt24":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_705 := hx_structural_source_691
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_705.__hx_this.writeUInt24(value)
						}, func() any {
							return hx_receiver_705.__hx_this
						}, "writeUInt24")
					}(), Found: true}
				case "writeInt32":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_706 := hx_structural_source_691
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_706.__hx_this.writeInt32(value)
						}, func() any {
							return hx_receiver_706.__hx_this
						}, "writeInt32")
					}(), Found: true}
				case "prepare":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_707 := hx_structural_source_691
						return hxrt.BoundFunction(func(nbytes int) {
							hx_receiver_707.__hx_this.prepare(nbytes)
						}, func() any {
							return hx_receiver_707.__hx_this
						}, "prepare")
					}(), Found: true}
				case "writeInput":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Input, any)] {
						hx_receiver_708 := hx_structural_source_691
						return hxrt.BoundFunction(func(input *haxe__io__Input, bufsize any) {
							hx_receiver_708.__hx_this.writeInput(input, bufsize)
						}, func() any {
							return hx_receiver_708.__hx_this
						}, "writeInput")
					}(), Found: true}
				case "writeString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string, *haxe__io__Encoding)] {
						hx_receiver_709 := hx_structural_source_691
						return hxrt.BoundFunction(func(value *string, encoding *haxe__io__Encoding) {
							hx_receiver_709.__hx_this.writeString(value, encoding)
						}, func() any {
							return hx_receiver_709.__hx_this
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
		return func(hx_structural_source_710 *haxe__io__Input) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_710, func(name string) hxrt.AnonymousField {
				switch name {
				case "bigEndian":
					return hxrt.AnonymousField{Value: hx_structural_source_710.bigEndian, Found: true}
				case "readByte":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_711 := hx_structural_source_710
						return hxrt.BoundFunction(func() int {
							return hx_receiver_711.__hx_this.readByte()
						}, func() any {
							return hx_receiver_711.__hx_this
						}, "readByte")
					}(), Found: true}
				case "readBytes":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes, int, int) int] {
						hx_receiver_712 := hx_structural_source_710
						return hxrt.BoundFunction(func(bytes *haxe__io__Bytes, pos int, len int) int {
							return hx_receiver_712.__hx_this.readBytes(bytes, pos, len)
						}, func() any {
							return hx_receiver_712.__hx_this
						}, "readBytes")
					}(), Found: true}
				case "close":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func()] {
						hx_receiver_713 := hx_structural_source_710
						return hxrt.BoundFunction(func() {
							hx_receiver_713.__hx_this.close()
						}, func() any {
							return hx_receiver_713.__hx_this
						}, "close")
					}(), Found: true}
				case "set_bigEndian":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(bool) bool] {
						hx_receiver_714 := hx_structural_source_710
						return hxrt.BoundFunction(func(value bool) bool {
							return hx_receiver_714.__hx_this.set_bigEndian(value)
						}, func() any {
							return hx_receiver_714.__hx_this
						}, "set_bigEndian")
					}(), Found: true}
				case "readAll":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(any) *haxe__io__Bytes] {
						hx_receiver_715 := hx_structural_source_710
						return hxrt.BoundFunction(func(bufsize any) *haxe__io__Bytes {
							return hx_receiver_715.__hx_this.readAll(bufsize)
						}, func() any {
							return hx_receiver_715.__hx_this
						}, "readAll")
					}(), Found: true}
				case "readFullBytes":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes, int, int)] {
						hx_receiver_716 := hx_structural_source_710
						return hxrt.BoundFunction(func(bytes *haxe__io__Bytes, pos int, len int) {
							hx_receiver_716.__hx_this.readFullBytes(bytes, pos, len)
						}, func() any {
							return hx_receiver_716.__hx_this
						}, "readFullBytes")
					}(), Found: true}
				case "read":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int) *haxe__io__Bytes] {
						hx_receiver_717 := hx_structural_source_710
						return hxrt.BoundFunction(func(nbytes int) *haxe__io__Bytes {
							return hx_receiver_717.__hx_this.read(nbytes)
						}, func() any {
							return hx_receiver_717.__hx_this
						}, "read")
					}(), Found: true}
				case "readUntil":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int) *string] {
						hx_receiver_718 := hx_structural_source_710
						return hxrt.BoundFunction(func(end int) *string {
							return hx_receiver_718.__hx_this.readUntil(end)
						}, func() any {
							return hx_receiver_718.__hx_this
						}, "readUntil")
					}(), Found: true}
				case "readLine":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_719 := hx_structural_source_710
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_719.__hx_this.readLine()
						}, func() any {
							return hx_receiver_719.__hx_this
						}, "readLine")
					}(), Found: true}
				case "readFloat":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() float64] {
						hx_receiver_720 := hx_structural_source_710
						return hxrt.BoundFunction(func() float64 {
							return hx_receiver_720.__hx_this.readFloat()
						}, func() any {
							return hx_receiver_720.__hx_this
						}, "readFloat")
					}(), Found: true}
				case "readDouble":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() float64] {
						hx_receiver_721 := hx_structural_source_710
						return hxrt.BoundFunction(func() float64 {
							return hx_receiver_721.__hx_this.readDouble()
						}, func() any {
							return hx_receiver_721.__hx_this
						}, "readDouble")
					}(), Found: true}
				case "readInt8":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_722 := hx_structural_source_710
						return hxrt.BoundFunction(func() int {
							return hx_receiver_722.__hx_this.readInt8()
						}, func() any {
							return hx_receiver_722.__hx_this
						}, "readInt8")
					}(), Found: true}
				case "readInt16":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_723 := hx_structural_source_710
						return hxrt.BoundFunction(func() int {
							return hx_receiver_723.__hx_this.readInt16()
						}, func() any {
							return hx_receiver_723.__hx_this
						}, "readInt16")
					}(), Found: true}
				case "readUInt16":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_724 := hx_structural_source_710
						return hxrt.BoundFunction(func() int {
							return hx_receiver_724.__hx_this.readUInt16()
						}, func() any {
							return hx_receiver_724.__hx_this
						}, "readUInt16")
					}(), Found: true}
				case "readInt24":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_725 := hx_structural_source_710
						return hxrt.BoundFunction(func() int {
							return hx_receiver_725.__hx_this.readInt24()
						}, func() any {
							return hx_receiver_725.__hx_this
						}, "readInt24")
					}(), Found: true}
				case "readUInt24":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_726 := hx_structural_source_710
						return hxrt.BoundFunction(func() int {
							return hx_receiver_726.__hx_this.readUInt24()
						}, func() any {
							return hx_receiver_726.__hx_this
						}, "readUInt24")
					}(), Found: true}
				case "readInt32":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_727 := hx_structural_source_710
						return hxrt.BoundFunction(func() int {
							return hx_receiver_727.__hx_this.readInt32()
						}, func() any {
							return hx_receiver_727.__hx_this
						}, "readInt32")
					}(), Found: true}
				case "readString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, *haxe__io__Encoding) *string] {
						hx_receiver_728 := hx_structural_source_710
						return hxrt.BoundFunction(func(len int, encoding *haxe__io__Encoding) *string {
							return hx_receiver_728.__hx_this.readString(len, encoding)
						}, func() any {
							return hx_receiver_728.__hx_this
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
		return func(hx_structural_source_729 *haxe__io__Eof) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_729, func(name string) hxrt.AnonymousField {
				switch name {
				case "toString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_730 := hx_structural_source_729
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_730.__hx_this.toString()
						}, func() any {
							return hx_receiver_730.__hx_this
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
		return func(hx_structural_source_731 *haxe__io__BytesBuffer) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_731, func(name string) hxrt.AnonymousField {
				switch name {
				case "b":
					return hxrt.AnonymousField{Value: hx_structural_source_731.b, Found: true}
				case "get_length":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_734 := hx_structural_source_731
						return hxrt.BoundFunction(func() int {
							return hx_receiver_734.__hx_this.get_length()
						}, func() any {
							return hx_receiver_734.__hx_this
						}, "get_length")
					}(), Found: true}
				case "addByte":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_735 := hx_structural_source_731
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_735.__hx_this.addByte(value)
						}, func() any {
							return hx_receiver_735.__hx_this
						}, "addByte")
					}(), Found: true}
				case "add":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes)] {
						hx_receiver_736 := hx_structural_source_731
						return hxrt.BoundFunction(func(source *haxe__io__Bytes) {
							hx_receiver_736.__hx_this.add(source)
						}, func() any {
							return hx_receiver_736.__hx_this
						}, "add")
					}(), Found: true}
				case "addString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string, *haxe__io__Encoding)] {
						hx_receiver_737 := hx_structural_source_731
						return hxrt.BoundFunction(func(value *string, encoding *haxe__io__Encoding) {
							hx_receiver_737.__hx_this.addString(value, encoding)
						}, func() any {
							return hx_receiver_737.__hx_this
						}, "addString")
					}(), Found: true}
				case "addInt32":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_738 := hx_structural_source_731
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_738.__hx_this.addInt32(value)
						}, func() any {
							return hx_receiver_738.__hx_this
						}, "addInt32")
					}(), Found: true}
				case "addInt64":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe___Int64_____Int64)] {
						hx_receiver_739 := hx_structural_source_731
						return hxrt.BoundFunction(func(value *haxe___Int64_____Int64) {
							hx_receiver_739.__hx_this.addInt64(value)
						}, func() any {
							return hx_receiver_739.__hx_this
						}, "addInt64")
					}(), Found: true}
				case "addFloat":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(float64)] {
						hx_receiver_740 := hx_structural_source_731
						return hxrt.BoundFunction(func(value float64) {
							hx_receiver_740.__hx_this.addFloat(value)
						}, func() any {
							return hx_receiver_740.__hx_this
						}, "addFloat")
					}(), Found: true}
				case "addDouble":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(float64)] {
						hx_receiver_741 := hx_structural_source_731
						return hxrt.BoundFunction(func(value float64) {
							hx_receiver_741.__hx_this.addDouble(value)
						}, func() any {
							return hx_receiver_741.__hx_this
						}, "addDouble")
					}(), Found: true}
				case "addBytes":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes, int, int)] {
						hx_receiver_742 := hx_structural_source_731
						return hxrt.BoundFunction(func(source *haxe__io__Bytes, pos int, len int) {
							hx_receiver_742.__hx_this.addBytes(source, pos, len)
						}, func() any {
							return hx_receiver_742.__hx_this
						}, "addBytes")
					}(), Found: true}
				case "getBytes":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *haxe__io__Bytes] {
						hx_receiver_743 := hx_structural_source_731
						return hxrt.BoundFunction(func() *haxe__io__Bytes {
							return hx_receiver_743.__hx_this.getBytes()
						}, func() any {
							return hx_receiver_743.__hx_this
						}, "getBytes")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "b":
					hx_structural_source_731.b = func(hx_value_732 any) []int {
						if hx_value_732 == nil {
							var hx_zero_733 []int
							return hx_zero_733
						}
						return hx_value_732.([]int)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe__io__Bytes:
		return func(hx_structural_source_744 *haxe__io__Bytes) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_744, func(name string) hxrt.AnonymousField {
				switch name {
				case "length":
					return hxrt.AnonymousField{Value: hx_structural_source_744.length, Found: true}
				case "b":
					return hxrt.AnonymousField{Value: hx_structural_source_744.b, Found: true}
				case "__hx_raw":
					return hxrt.AnonymousField{Value: hx_structural_source_744.__hx_raw, Found: true}
				case "__hx_rawValid":
					return hxrt.AnonymousField{Value: hx_structural_source_744.__hx_rawValid, Found: true}
				case "__hx_dataExposed":
					return hxrt.AnonymousField{Value: hx_structural_source_744.__hx_dataExposed, Found: true}
				case "get":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int) int] {
						hx_receiver_753 := hx_structural_source_744
						return hxrt.BoundFunction(func(pos int) int {
							return hx_receiver_753.__hx_this.get(pos)
						}, func() any {
							return hx_receiver_753.__hx_this
						}, "get")
					}(), Found: true}
				case "set":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, int)] {
						hx_receiver_754 := hx_structural_source_744
						return hxrt.BoundFunction(func(pos int, value int) {
							hx_receiver_754.__hx_this.set(pos, value)
						}, func() any {
							return hx_receiver_754.__hx_this
						}, "set")
					}(), Found: true}
				case "blit":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, *haxe__io__Bytes, int, int)] {
						hx_receiver_755 := hx_structural_source_744
						return hxrt.BoundFunction(func(pos int, src *haxe__io__Bytes, srcpos int, len int) {
							hx_receiver_755.__hx_this.blit(pos, src, srcpos, len)
						}, func() any {
							return hx_receiver_755.__hx_this
						}, "blit")
					}(), Found: true}
				case "fill":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, int, int)] {
						hx_receiver_756 := hx_structural_source_744
						return hxrt.BoundFunction(func(pos int, len int, value int) {
							hx_receiver_756.__hx_this.fill(pos, len, value)
						}, func() any {
							return hx_receiver_756.__hx_this
						}, "fill")
					}(), Found: true}
				case "sub":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, int) *haxe__io__Bytes] {
						hx_receiver_757 := hx_structural_source_744
						return hxrt.BoundFunction(func(pos int, len int) *haxe__io__Bytes {
							return hx_receiver_757.__hx_this.sub(pos, len)
						}, func() any {
							return hx_receiver_757.__hx_this
						}, "sub")
					}(), Found: true}
				case "compare":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes) int] {
						hx_receiver_758 := hx_structural_source_744
						return hxrt.BoundFunction(func(other *haxe__io__Bytes) int {
							return hx_receiver_758.__hx_this.compare(other)
						}, func() any {
							return hx_receiver_758.__hx_this
						}, "compare")
					}(), Found: true}
				case "getDouble":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int) float64] {
						hx_receiver_759 := hx_structural_source_744
						return hxrt.BoundFunction(func(pos int) float64 {
							return hx_receiver_759.__hx_this.getDouble(pos)
						}, func() any {
							return hx_receiver_759.__hx_this
						}, "getDouble")
					}(), Found: true}
				case "getFloat":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int) float64] {
						hx_receiver_760 := hx_structural_source_744
						return hxrt.BoundFunction(func(pos int) float64 {
							return hx_receiver_760.__hx_this.getFloat(pos)
						}, func() any {
							return hx_receiver_760.__hx_this
						}, "getFloat")
					}(), Found: true}
				case "setDouble":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, float64)] {
						hx_receiver_761 := hx_structural_source_744
						return hxrt.BoundFunction(func(pos int, value float64) {
							hx_receiver_761.__hx_this.setDouble(pos, value)
						}, func() any {
							return hx_receiver_761.__hx_this
						}, "setDouble")
					}(), Found: true}
				case "setFloat":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, float64)] {
						hx_receiver_762 := hx_structural_source_744
						return hxrt.BoundFunction(func(pos int, value float64) {
							hx_receiver_762.__hx_this.setFloat(pos, value)
						}, func() any {
							return hx_receiver_762.__hx_this
						}, "setFloat")
					}(), Found: true}
				case "getUInt16":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int) int] {
						hx_receiver_763 := hx_structural_source_744
						return hxrt.BoundFunction(func(pos int) int {
							return hx_receiver_763.__hx_this.getUInt16(pos)
						}, func() any {
							return hx_receiver_763.__hx_this
						}, "getUInt16")
					}(), Found: true}
				case "setUInt16":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, int)] {
						hx_receiver_764 := hx_structural_source_744
						return hxrt.BoundFunction(func(pos int, value int) {
							hx_receiver_764.__hx_this.setUInt16(pos, value)
						}, func() any {
							return hx_receiver_764.__hx_this
						}, "setUInt16")
					}(), Found: true}
				case "getInt32":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int) int] {
						hx_receiver_765 := hx_structural_source_744
						return hxrt.BoundFunction(func(pos int) int {
							return hx_receiver_765.__hx_this.getInt32(pos)
						}, func() any {
							return hx_receiver_765.__hx_this
						}, "getInt32")
					}(), Found: true}
				case "getInt64":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int) *haxe___Int64_____Int64] {
						hx_receiver_766 := hx_structural_source_744
						return hxrt.BoundFunction(func(pos int) *haxe___Int64_____Int64 {
							return hx_receiver_766.__hx_this.getInt64(pos)
						}, func() any {
							return hx_receiver_766.__hx_this
						}, "getInt64")
					}(), Found: true}
				case "setInt32":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, int)] {
						hx_receiver_767 := hx_structural_source_744
						return hxrt.BoundFunction(func(pos int, value int) {
							hx_receiver_767.__hx_this.setInt32(pos, value)
						}, func() any {
							return hx_receiver_767.__hx_this
						}, "setInt32")
					}(), Found: true}
				case "setInt64":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, *haxe___Int64_____Int64)] {
						hx_receiver_768 := hx_structural_source_744
						return hxrt.BoundFunction(func(pos int, value *haxe___Int64_____Int64) {
							hx_receiver_768.__hx_this.setInt64(pos, value)
						}, func() any {
							return hx_receiver_768.__hx_this
						}, "setInt64")
					}(), Found: true}
				case "getString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, int, *haxe__io__Encoding) *string] {
						hx_receiver_769 := hx_structural_source_744
						return hxrt.BoundFunction(func(pos int, len int, encoding *haxe__io__Encoding) *string {
							return hx_receiver_769.__hx_this.getString(pos, len, encoding)
						}, func() any {
							return hx_receiver_769.__hx_this
						}, "getString")
					}(), Found: true}
				case "readString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, int) *string] {
						hx_receiver_770 := hx_structural_source_744
						return hxrt.BoundFunction(func(pos int, len int) *string {
							return hx_receiver_770.__hx_this.readString(pos, len)
						}, func() any {
							return hx_receiver_770.__hx_this
						}, "readString")
					}(), Found: true}
				case "toString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_771 := hx_structural_source_744
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_771.__hx_this.toString()
						}, func() any {
							return hx_receiver_771.__hx_this
						}, "toString")
					}(), Found: true}
				case "toHex":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_772 := hx_structural_source_744
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_772.__hx_this.toHex()
						}, func() any {
							return hx_receiver_772.__hx_this
						}, "toHex")
					}(), Found: true}
				case "getData":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() []int] {
						hx_receiver_773 := hx_structural_source_744
						return hxrt.BoundFunction(func() []int {
							return hx_receiver_773.__hx_this.getData()
						}, func() any {
							return hx_receiver_773.__hx_this
						}, "getData")
					}(), Found: true}
				case "__hx_nativeView":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *hxrt.ByteView] {
						hx_receiver_774 := hx_structural_source_744
						return hxrt.BoundFunction(func() *hxrt.ByteView {
							return hx_receiver_774.__hx_this.__hx_nativeView()
						}, func() any {
							return hx_receiver_774.__hx_this
						}, "__hx_nativeView")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "b":
					hx_structural_source_744.b = func(hx_value_745 any) []int {
						if hx_value_745 == nil {
							var hx_zero_746 []int
							return hx_zero_746
						}
						return hx_value_745.([]int)
					}(incoming)
					return true
				case "__hx_raw":
					hx_structural_source_744.__hx_raw = func(hx_value_747 any) *hxrt.ByteView {
						if hx_value_747 == nil {
							var hx_zero_748 *hxrt.ByteView
							return hx_zero_748
						}
						return hx_value_747.(*hxrt.ByteView)
					}(incoming)
					return true
				case "__hx_rawValid":
					hx_structural_source_744.__hx_rawValid = func(hx_value_749 any) bool {
						if hx_value_749 == nil {
							var hx_zero_750 bool
							return hx_zero_750
						}
						return hx_value_749.(bool)
					}(incoming)
					return true
				case "__hx_dataExposed":
					hx_structural_source_744.__hx_dataExposed = func(hx_value_751 any) bool {
						if hx_value_751 == nil {
							var hx_zero_752 bool
							return hx_zero_752
						}
						return hx_value_751.(bool)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe__http__HttpBase:
		return func(hx_structural_source_775 *haxe__http__HttpBase) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_775, func(name string) hxrt.AnonymousField {
				switch name {
				case "url":
					return hxrt.AnonymousField{Value: hx_structural_source_775.url, Found: true}
				case "responseBytes":
					return hxrt.AnonymousField{Value: hx_structural_source_775.responseBytes, Found: true}
				case "responseAsString":
					return hxrt.AnonymousField{Value: hx_structural_source_775.responseAsString, Found: true}
				case "postData":
					return hxrt.AnonymousField{Value: hx_structural_source_775.postData, Found: true}
				case "postBytes":
					return hxrt.AnonymousField{Value: hx_structural_source_775.postBytes, Found: true}
				case "headers":
					return hxrt.AnonymousField{Value: hx_structural_source_775.headers, Found: true}
				case "params":
					return hxrt.AnonymousField{Value: hx_structural_source_775.params, Found: true}
				case "emptyOnData":
					return hxrt.AnonymousField{Value: hx_structural_source_775.emptyOnData, Found: true}
				case "setHeader":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string, *string)] {
						hx_receiver_786 := hx_structural_source_775
						return hxrt.BoundFunction(func(name *string, value *string) {
							hx_receiver_786.__hx_this.setHeader(name, value)
						}, func() any {
							return hx_receiver_786.__hx_this
						}, "setHeader")
					}(), Found: true}
				case "addHeader":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string, *string)] {
						hx_receiver_787 := hx_structural_source_775
						return hxrt.BoundFunction(func(header *string, value *string) {
							hx_receiver_787.__hx_this.addHeader(header, value)
						}, func() any {
							return hx_receiver_787.__hx_this
						}, "addHeader")
					}(), Found: true}
				case "setParameter":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string, *string)] {
						hx_receiver_788 := hx_structural_source_775
						return hxrt.BoundFunction(func(name *string, value *string) {
							hx_receiver_788.__hx_this.setParameter(name, value)
						}, func() any {
							return hx_receiver_788.__hx_this
						}, "setParameter")
					}(), Found: true}
				case "addParameter":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string, *string)] {
						hx_receiver_789 := hx_structural_source_775
						return hxrt.BoundFunction(func(name *string, value *string) {
							hx_receiver_789.__hx_this.addParameter(name, value)
						}, func() any {
							return hx_receiver_789.__hx_this
						}, "addParameter")
					}(), Found: true}
				case "setPostData":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string)] {
						hx_receiver_790 := hx_structural_source_775
						return hxrt.BoundFunction(func(data *string) {
							hx_receiver_790.__hx_this.setPostData(data)
						}, func() any {
							return hx_receiver_790.__hx_this
						}, "setPostData")
					}(), Found: true}
				case "setPostBytes":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes)] {
						hx_receiver_791 := hx_structural_source_775
						return hxrt.BoundFunction(func(data *haxe__io__Bytes) {
							hx_receiver_791.__hx_this.setPostBytes(data)
						}, func() any {
							return hx_receiver_791.__hx_this
						}, "setPostBytes")
					}(), Found: true}
				case "request":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(any)] {
						hx_receiver_792 := hx_structural_source_775
						return hxrt.BoundFunction(func(post any) {
							hx_receiver_792.__hx_this.request(post)
						}, func() any {
							return hx_receiver_792.__hx_this
						}, "request")
					}(), Found: true}
				case "onData":
					return hxrt.AnonymousField{Value: hxrt.ReadDynamicFunction(hx_structural_source_775.onData), Found: true}
				case "onBytes":
					return hxrt.AnonymousField{Value: hxrt.ReadDynamicFunction(hx_structural_source_775.onBytes), Found: true}
				case "onError":
					return hxrt.AnonymousField{Value: hxrt.ReadDynamicFunction(hx_structural_source_775.onError), Found: true}
				case "onStatus":
					return hxrt.AnonymousField{Value: hxrt.ReadDynamicFunction(hx_structural_source_775.onStatus), Found: true}
				case "hasOnData":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() bool] {
						hx_receiver_793 := hx_structural_source_775
						return hxrt.BoundFunction(func() bool {
							return hx_receiver_793.__hx_this.hasOnData()
						}, func() any {
							return hx_receiver_793.__hx_this
						}, "hasOnData")
					}(), Found: true}
				case "success":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes)] {
						hx_receiver_794 := hx_structural_source_775
						return hxrt.BoundFunction(func(data *haxe__io__Bytes) {
							hx_receiver_794.__hx_this.success(data)
						}, func() any {
							return hx_receiver_794.__hx_this
						}, "success")
					}(), Found: true}
				case "get_responseData":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_795 := hx_structural_source_775
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_795.__hx_this.get_responseData()
						}, func() any {
							return hx_receiver_795.__hx_this
						}, "get_responseData")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "url":
					hx_structural_source_775.url = func(hx_value_776 any) *string {
						if hx_value_776 == nil {
							var hx_zero_777 *string
							return hx_zero_777
						}
						return hx_value_776.(*string)
					}(incoming)
					return true
				case "responseAsString":
					hx_structural_source_775.responseAsString = func(hx_value_778 any) *string {
						if hx_value_778 == nil {
							var hx_zero_779 *string
							return hx_zero_779
						}
						return hx_value_778.(*string)
					}(incoming)
					return true
				case "postData":
					hx_structural_source_775.postData = func(hx_value_780 any) *string {
						if hx_value_780 == nil {
							var hx_zero_781 *string
							return hx_zero_781
						}
						return hx_value_780.(*string)
					}(incoming)
					return true
				case "postBytes":
					hx_structural_source_775.postBytes = hx_recover_class_7(incoming)
					return true
				case "headers":
					hx_structural_source_775.headers = func(hx_value_782 any) *hxrt.Array {
						if hx_value_782 == nil {
							var hx_zero_783 *hxrt.Array
							return hx_zero_783
						}
						return hx_value_782.(*hxrt.Array)
					}(incoming)
					return true
				case "params":
					hx_structural_source_775.params = func(hx_value_784 any) *hxrt.Array {
						if hx_value_784 == nil {
							var hx_zero_785 *hxrt.Array
							return hx_zero_785
						}
						return hx_value_784.(*hxrt.Array)
					}(incoming)
					return true
				case "onData":
					hx_structural_source_775.onData = hx_recover_callable_3(incoming)
					return true
				case "onBytes":
					hx_structural_source_775.onBytes = hx_recover_callable_4(incoming)
					return true
				case "onError":
					hx_structural_source_775.onError = hx_recover_callable_3(incoming)
					return true
				case "onStatus":
					hx_structural_source_775.onStatus = hx_recover_callable_5(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe__exceptions__NotImplementedException:
		return func(hx_structural_source_796 *haxe__exceptions__NotImplementedException) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_796, func(name string) hxrt.AnonymousField {
				switch name {
				case "posInfos":
					return hxrt.AnonymousField{Value: hx_structural_source_796.posInfos, Found: true}
				case "toString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_797 := hx_structural_source_796
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_797.__hx_this.toString()
						}, func() any {
							return hx_receiver_797.__hx_this
						}, "toString")
					}(), Found: true}
				case "savedMessage":
					return hxrt.AnonymousField{Value: hx_structural_source_796.savedMessage, Found: true}
				case "savedStack":
					return hxrt.AnonymousField{Value: hx_structural_source_796.savedStack, Found: true}
				case "savedPrevious":
					return hxrt.AnonymousField{Value: hx_structural_source_796.savedPrevious, Found: true}
				case "savedNative":
					return hxrt.AnonymousField{Value: hx_structural_source_796.savedNative, Found: true}
				case "get_message":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_800 := hx_structural_source_796
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_800.__hx_this.get_message()
						}, func() any {
							return hx_receiver_800.__hx_this
						}, "get_message")
					}(), Found: true}
				case "get_stack":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_801 := hx_structural_source_796
						return hxrt.BoundFunction(func() any {
							return hx_receiver_801.__hx_this.get_stack()
						}, func() any {
							return hx_receiver_801.__hx_this
						}, "get_stack")
					}(), Found: true}
				case "get_previous":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *haxe__Exception] {
						hx_receiver_802 := hx_structural_source_796
						return hxrt.BoundFunction(func() *haxe__Exception {
							return hx_receiver_802.__hx_this.get_previous()
						}, func() any {
							return hx_receiver_802.__hx_this
						}, "get_previous")
					}(), Found: true}
				case "get_native":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_803 := hx_structural_source_796
						return hxrt.BoundFunction(func() any {
							return hx_receiver_803.__hx_this.get_native()
						}, func() any {
							return hx_receiver_803.__hx_this
						}, "get_native")
					}(), Found: true}
				case "unwrap":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_804 := hx_structural_source_796
						return hxrt.BoundFunction(func() any {
							return hx_receiver_804.__hx_this.unwrap()
						}, func() any {
							return hx_receiver_804.__hx_this
						}, "unwrap")
					}(), Found: true}
				case "details":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_805 := hx_structural_source_796
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_805.__hx_this.details()
						}, func() any {
							return hx_receiver_805.__hx_this
						}, "details")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "savedMessage":
					hx_structural_source_796.savedMessage = func(hx_value_798 any) *string {
						if hx_value_798 == nil {
							var hx_zero_799 *string
							return hx_zero_799
						}
						return hx_value_798.(*string)
					}(incoming)
					return true
				case "savedStack":
					hx_structural_source_796.savedStack = incoming
					return true
				case "savedPrevious":
					hx_structural_source_796.savedPrevious = hx_recover_class_2(incoming)
					return true
				case "savedNative":
					hx_structural_source_796.savedNative = incoming
					return true
				}
				return false
			})
		}(source)
	case *haxe__ds__StringMap:
		return func(hx_structural_source_806 *haxe__ds__StringMap) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_806, func(name string) hxrt.AnonymousField {
				switch name {
				case "h":
					return hxrt.AnonymousField{Value: hx_structural_source_806.h, Found: true}
				case "set":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string, any)] {
						hx_receiver_809 := hx_structural_source_806
						return hxrt.BoundFunction(func(key *string, value any) {
							hx_receiver_809.__hx_this.set(key, value)
						}, func() any {
							return hx_receiver_809.__hx_this
						}, "set")
					}(), Found: true}
				case "get":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string) any] {
						hx_receiver_810 := hx_structural_source_806
						return hxrt.BoundFunction(func(key *string) any {
							return hx_receiver_810.__hx_this.get(key)
						}, func() any {
							return hx_receiver_810.__hx_this
						}, "get")
					}(), Found: true}
				case "exists":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string) bool] {
						hx_receiver_811 := hx_structural_source_806
						return hxrt.BoundFunction(func(key *string) bool {
							return hx_receiver_811.__hx_this.exists(key)
						}, func() any {
							return hx_receiver_811.__hx_this
						}, "exists")
					}(), Found: true}
				case "remove":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string) bool] {
						hx_receiver_812 := hx_structural_source_806
						return hxrt.BoundFunction(func(key *string) bool {
							return hx_receiver_812.__hx_this.remove(key)
						}, func() any {
							return hx_receiver_812.__hx_this
						}, "remove")
					}(), Found: true}
				case "keys":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() hxrt.Anonymous] {
						hx_receiver_813 := hx_structural_source_806
						return hxrt.BoundFunction(func() hxrt.Anonymous {
							return hx_receiver_813.__hx_this.keys()
						}, func() any {
							return hx_receiver_813.__hx_this
						}, "keys")
					}(), Found: true}
				case "iterator":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() hxrt.Anonymous] {
						hx_receiver_814 := hx_structural_source_806
						return hxrt.BoundFunction(func() hxrt.Anonymous {
							return hx_receiver_814.__hx_this.iterator()
						}, func() any {
							return hx_receiver_814.__hx_this
						}, "iterator")
					}(), Found: true}
				case "keyValueIterator":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() hxrt.Anonymous] {
						hx_receiver_815 := hx_structural_source_806
						return hxrt.BoundFunction(func() hxrt.Anonymous {
							return hx_receiver_815.__hx_this.keyValueIterator()
						}, func() any {
							return hx_receiver_815.__hx_this
						}, "keyValueIterator")
					}(), Found: true}
				case "getIMap":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(any) any] {
						hx_receiver_816 := hx_structural_source_806
						return hxrt.BoundFunction(func(key any) any {
							return hx_receiver_816.__hx_this.getIMap(key)
						}, func() any {
							return hx_receiver_816.__hx_this
						}, "getIMap")
					}(), Found: true}
				case "setIMap":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(any, any)] {
						hx_receiver_817 := hx_structural_source_806
						return hxrt.BoundFunction(func(key any, value any) {
							hx_receiver_817.__hx_this.setIMap(key, value)
						}, func() any {
							return hx_receiver_817.__hx_this
						}, "setIMap")
					}(), Found: true}
				case "existsIMap":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(any) bool] {
						hx_receiver_818 := hx_structural_source_806
						return hxrt.BoundFunction(func(key any) bool {
							return hx_receiver_818.__hx_this.existsIMap(key)
						}, func() any {
							return hx_receiver_818.__hx_this
						}, "existsIMap")
					}(), Found: true}
				case "removeIMap":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(any) bool] {
						hx_receiver_819 := hx_structural_source_806
						return hxrt.BoundFunction(func(key any) bool {
							return hx_receiver_819.__hx_this.removeIMap(key)
						}, func() any {
							return hx_receiver_819.__hx_this
						}, "removeIMap")
					}(), Found: true}
				case "copyIMap":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() haxe__IMap] {
						hx_receiver_820 := hx_structural_source_806
						return hxrt.BoundFunction(func() haxe__IMap {
							return hx_receiver_820.__hx_this.copyIMap()
						}, func() any {
							return hx_receiver_820.__hx_this
						}, "copyIMap")
					}(), Found: true}
				case "copy":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *haxe__ds__StringMap] {
						hx_receiver_821 := hx_structural_source_806
						return hxrt.BoundFunction(func() *haxe__ds__StringMap {
							return hx_receiver_821.__hx_this.copy()
						}, func() any {
							return hx_receiver_821.__hx_this
						}, "copy")
					}(), Found: true}
				case "toString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_822 := hx_structural_source_806
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_822.__hx_this.toString()
						}, func() any {
							return hx_receiver_822.__hx_this
						}, "toString")
					}(), Found: true}
				case "clear":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func()] {
						hx_receiver_823 := hx_structural_source_806
						return hxrt.BoundFunction(func() {
							hx_receiver_823.__hx_this.clear()
						}, func() any {
							return hx_receiver_823.__hx_this
						}, "clear")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "h":
					hx_structural_source_806.h = func(hx_value_807 any) *hxrt.StringMapCell {
						if hx_value_807 == nil {
							var hx_zero_808 *hxrt.StringMapCell
							return hx_zero_808
						}
						return hx_value_807.(*hxrt.StringMapCell)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe__Exception:
		return func(hx_structural_source_824 *haxe__Exception) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_824, func(name string) hxrt.AnonymousField {
				switch name {
				case "savedMessage":
					return hxrt.AnonymousField{Value: hx_structural_source_824.savedMessage, Found: true}
				case "savedStack":
					return hxrt.AnonymousField{Value: hx_structural_source_824.savedStack, Found: true}
				case "savedPrevious":
					return hxrt.AnonymousField{Value: hx_structural_source_824.savedPrevious, Found: true}
				case "savedNative":
					return hxrt.AnonymousField{Value: hx_structural_source_824.savedNative, Found: true}
				case "get_message":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_827 := hx_structural_source_824
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_827.__hx_this.get_message()
						}, func() any {
							return hx_receiver_827.__hx_this
						}, "get_message")
					}(), Found: true}
				case "get_stack":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_828 := hx_structural_source_824
						return hxrt.BoundFunction(func() any {
							return hx_receiver_828.__hx_this.get_stack()
						}, func() any {
							return hx_receiver_828.__hx_this
						}, "get_stack")
					}(), Found: true}
				case "get_previous":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *haxe__Exception] {
						hx_receiver_829 := hx_structural_source_824
						return hxrt.BoundFunction(func() *haxe__Exception {
							return hx_receiver_829.__hx_this.get_previous()
						}, func() any {
							return hx_receiver_829.__hx_this
						}, "get_previous")
					}(), Found: true}
				case "get_native":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_830 := hx_structural_source_824
						return hxrt.BoundFunction(func() any {
							return hx_receiver_830.__hx_this.get_native()
						}, func() any {
							return hx_receiver_830.__hx_this
						}, "get_native")
					}(), Found: true}
				case "unwrap":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_831 := hx_structural_source_824
						return hxrt.BoundFunction(func() any {
							return hx_receiver_831.__hx_this.unwrap()
						}, func() any {
							return hx_receiver_831.__hx_this
						}, "unwrap")
					}(), Found: true}
				case "toString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_832 := hx_structural_source_824
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_832.__hx_this.toString()
						}, func() any {
							return hx_receiver_832.__hx_this
						}, "toString")
					}(), Found: true}
				case "details":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_833 := hx_structural_source_824
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_833.__hx_this.details()
						}, func() any {
							return hx_receiver_833.__hx_this
						}, "details")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "savedMessage":
					hx_structural_source_824.savedMessage = func(hx_value_825 any) *string {
						if hx_value_825 == nil {
							var hx_zero_826 *string
							return hx_zero_826
						}
						return hx_value_825.(*string)
					}(incoming)
					return true
				case "savedStack":
					hx_structural_source_824.savedStack = incoming
					return true
				case "savedPrevious":
					hx_structural_source_824.savedPrevious = hx_recover_class_2(incoming)
					return true
				case "savedNative":
					hx_structural_source_824.savedNative = incoming
					return true
				}
				return false
			})
		}(source)
	case *sys__thread__Tls:
		return func(hx_structural_source_834 *sys__thread__Tls) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_834, func(name string) hxrt.AnonymousField {
				switch name {
				case "__handle":
					return hxrt.AnonymousField{Value: hx_structural_source_834.__handle, Found: true}
				case "get_value":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_835 := hx_structural_source_834
						return hxrt.BoundFunction(func() any {
							return hx_receiver_835.__hx_this.get_value()
						}, func() any {
							return hx_receiver_835.__hx_this
						}, "get_value")
					}(), Found: true}
				case "set_value":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(any) any] {
						hx_receiver_836 := hx_structural_source_834
						return hxrt.BoundFunction(func(v any) any {
							return hx_receiver_836.__hx_this.set_value(v)
						}, func() any {
							return hx_receiver_836.__hx_this
						}, "set_value")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				}
				return false
			})
		}(source)
	case *haxe__iterators__StringKeyValueIterator:
		return func(hx_structural_source_837 *haxe__iterators__StringKeyValueIterator) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_837, func(name string) hxrt.AnonymousField {
				switch name {
				case "offset":
					return hxrt.AnonymousField{Value: hx_structural_source_837.offset, Found: true}
				case "s":
					return hxrt.AnonymousField{Value: hx_structural_source_837.s, Found: true}
				case "hasNext":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() bool] {
						hx_receiver_842 := hx_structural_source_837
						return hxrt.BoundFunction(func() bool {
							return hx_receiver_842.__hx_this.hasNext()
						}, func() any {
							return hx_receiver_842.__hx_this
						}, "hasNext")
					}(), Found: true}
				case "next":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() hxrt.Anonymous] {
						hx_receiver_843 := hx_structural_source_837
						return hxrt.BoundFunction(func() hxrt.Anonymous {
							return hx_receiver_843.__hx_this.next()
						}, func() any {
							return hx_receiver_843.__hx_this
						}, "next")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "offset":
					hx_structural_source_837.offset = func(hx_value_838 any) int {
						if hx_value_838 == nil {
							var hx_zero_839 int
							return hx_zero_839
						}
						return hx_value_838.(int)
					}(incoming)
					return true
				case "s":
					hx_structural_source_837.s = func(hx_value_840 any) *string {
						if hx_value_840 == nil {
							var hx_zero_841 *string
							return hx_zero_841
						}
						return hx_value_840.(*string)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe__iterators__StringIterator:
		return func(hx_structural_source_844 *haxe__iterators__StringIterator) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_844, func(name string) hxrt.AnonymousField {
				switch name {
				case "offset":
					return hxrt.AnonymousField{Value: hx_structural_source_844.offset, Found: true}
				case "s":
					return hxrt.AnonymousField{Value: hx_structural_source_844.s, Found: true}
				case "hasNext":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() bool] {
						hx_receiver_849 := hx_structural_source_844
						return hxrt.BoundFunction(func() bool {
							return hx_receiver_849.__hx_this.hasNext()
						}, func() any {
							return hx_receiver_849.__hx_this
						}, "hasNext")
					}(), Found: true}
				case "next":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_850 := hx_structural_source_844
						return hxrt.BoundFunction(func() int {
							return hx_receiver_850.__hx_this.next()
						}, func() any {
							return hx_receiver_850.__hx_this
						}, "next")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "offset":
					hx_structural_source_844.offset = func(hx_value_845 any) int {
						if hx_value_845 == nil {
							var hx_zero_846 int
							return hx_zero_846
						}
						return hx_value_845.(int)
					}(incoming)
					return true
				case "s":
					hx_structural_source_844.s = func(hx_value_847 any) *string {
						if hx_value_847 == nil {
							var hx_zero_848 *string
							return hx_zero_848
						}
						return hx_value_847.(*string)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe__iterators__MapKeyValueIterator:
		return func(hx_structural_source_851 *haxe__iterators__MapKeyValueIterator) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_851, func(name string) hxrt.AnonymousField {
				switch name {
				case "map":
					return hxrt.AnonymousField{Value: hx_structural_source_851.map_, Found: true}
				case "keys":
					return hxrt.AnonymousField{Value: hx_structural_source_851.keys, Found: true}
				case "hasNext":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() bool] {
						hx_receiver_854 := hx_structural_source_851
						return hxrt.BoundFunction(func() bool {
							return hx_receiver_854.__hx_this.hasNext()
						}, func() any {
							return hx_receiver_854.__hx_this
						}, "hasNext")
					}(), Found: true}
				case "next":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() hxrt.Anonymous] {
						hx_receiver_855 := hx_structural_source_851
						return hxrt.BoundFunction(func() hxrt.Anonymous {
							return hx_receiver_855.__hx_this.next()
						}, func() any {
							return hx_receiver_855.__hx_this
						}, "next")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "map":
					hx_structural_source_851.map_ = func(hx_value_852 any) haxe__IMap {
						if hx_value_852 == nil {
							var hx_zero_853 haxe__IMap
							return hx_zero_853
						}
						return hx_value_852.(haxe__IMap)
					}(hxrt.ObjectSource(incoming))
					return true
				case "keys":
					hx_structural_source_851.keys = hx_structural_recover(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe__io__BytesOutput:
		return func(hx_structural_source_856 *haxe__io__BytesOutput) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_856, func(name string) hxrt.AnonymousField {
				switch name {
				case "b":
					return hxrt.AnonymousField{Value: hx_structural_source_856.b, Found: true}
				case "get_length":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_857 := hx_structural_source_856
						return hxrt.BoundFunction(func() int {
							return hx_receiver_857.__hx_this.get_length()
						}, func() any {
							return hx_receiver_857.__hx_this
						}, "get_length")
					}(), Found: true}
				case "writeByte":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_858 := hx_structural_source_856
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_858.__hx_this.writeByte(value)
						}, func() any {
							return hx_receiver_858.__hx_this
						}, "writeByte")
					}(), Found: true}
				case "writeBytes":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes, int, int) int] {
						hx_receiver_859 := hx_structural_source_856
						return hxrt.BoundFunction(func(bytes *haxe__io__Bytes, pos int, len int) int {
							return hx_receiver_859.__hx_this.writeBytes(bytes, pos, len)
						}, func() any {
							return hx_receiver_859.__hx_this
						}, "writeBytes")
					}(), Found: true}
				case "getBytes":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *haxe__io__Bytes] {
						hx_receiver_860 := hx_structural_source_856
						return hxrt.BoundFunction(func() *haxe__io__Bytes {
							return hx_receiver_860.__hx_this.getBytes()
						}, func() any {
							return hx_receiver_860.__hx_this
						}, "getBytes")
					}(), Found: true}
				case "bigEndian":
					return hxrt.AnonymousField{Value: hx_structural_source_856.bigEndian, Found: true}
				case "flush":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func()] {
						hx_receiver_861 := hx_structural_source_856
						return hxrt.BoundFunction(func() {
							hx_receiver_861.__hx_this.flush()
						}, func() any {
							return hx_receiver_861.__hx_this
						}, "flush")
					}(), Found: true}
				case "close":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func()] {
						hx_receiver_862 := hx_structural_source_856
						return hxrt.BoundFunction(func() {
							hx_receiver_862.__hx_this.close()
						}, func() any {
							return hx_receiver_862.__hx_this
						}, "close")
					}(), Found: true}
				case "set_bigEndian":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(bool) bool] {
						hx_receiver_863 := hx_structural_source_856
						return hxrt.BoundFunction(func(value bool) bool {
							return hx_receiver_863.__hx_this.set_bigEndian(value)
						}, func() any {
							return hx_receiver_863.__hx_this
						}, "set_bigEndian")
					}(), Found: true}
				case "write":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes)] {
						hx_receiver_864 := hx_structural_source_856
						return hxrt.BoundFunction(func(bytes *haxe__io__Bytes) {
							hx_receiver_864.__hx_this.write(bytes)
						}, func() any {
							return hx_receiver_864.__hx_this
						}, "write")
					}(), Found: true}
				case "writeFullBytes":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes, int, int)] {
						hx_receiver_865 := hx_structural_source_856
						return hxrt.BoundFunction(func(bytes *haxe__io__Bytes, pos int, len int) {
							hx_receiver_865.__hx_this.writeFullBytes(bytes, pos, len)
						}, func() any {
							return hx_receiver_865.__hx_this
						}, "writeFullBytes")
					}(), Found: true}
				case "writeFloat":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(float64)] {
						hx_receiver_866 := hx_structural_source_856
						return hxrt.BoundFunction(func(value float64) {
							hx_receiver_866.__hx_this.writeFloat(value)
						}, func() any {
							return hx_receiver_866.__hx_this
						}, "writeFloat")
					}(), Found: true}
				case "writeDouble":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(float64)] {
						hx_receiver_867 := hx_structural_source_856
						return hxrt.BoundFunction(func(value float64) {
							hx_receiver_867.__hx_this.writeDouble(value)
						}, func() any {
							return hx_receiver_867.__hx_this
						}, "writeDouble")
					}(), Found: true}
				case "writeInt8":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_868 := hx_structural_source_856
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_868.__hx_this.writeInt8(value)
						}, func() any {
							return hx_receiver_868.__hx_this
						}, "writeInt8")
					}(), Found: true}
				case "writeInt16":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_869 := hx_structural_source_856
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_869.__hx_this.writeInt16(value)
						}, func() any {
							return hx_receiver_869.__hx_this
						}, "writeInt16")
					}(), Found: true}
				case "writeUInt16":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_870 := hx_structural_source_856
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_870.__hx_this.writeUInt16(value)
						}, func() any {
							return hx_receiver_870.__hx_this
						}, "writeUInt16")
					}(), Found: true}
				case "writeInt24":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_871 := hx_structural_source_856
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_871.__hx_this.writeInt24(value)
						}, func() any {
							return hx_receiver_871.__hx_this
						}, "writeInt24")
					}(), Found: true}
				case "writeUInt24":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_872 := hx_structural_source_856
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_872.__hx_this.writeUInt24(value)
						}, func() any {
							return hx_receiver_872.__hx_this
						}, "writeUInt24")
					}(), Found: true}
				case "writeInt32":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_873 := hx_structural_source_856
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_873.__hx_this.writeInt32(value)
						}, func() any {
							return hx_receiver_873.__hx_this
						}, "writeInt32")
					}(), Found: true}
				case "prepare":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_874 := hx_structural_source_856
						return hxrt.BoundFunction(func(nbytes int) {
							hx_receiver_874.__hx_this.prepare(nbytes)
						}, func() any {
							return hx_receiver_874.__hx_this
						}, "prepare")
					}(), Found: true}
				case "writeInput":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Input, any)] {
						hx_receiver_875 := hx_structural_source_856
						return hxrt.BoundFunction(func(input *haxe__io__Input, bufsize any) {
							hx_receiver_875.__hx_this.writeInput(input, bufsize)
						}, func() any {
							return hx_receiver_875.__hx_this
						}, "writeInput")
					}(), Found: true}
				case "writeString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string, *haxe__io__Encoding)] {
						hx_receiver_876 := hx_structural_source_856
						return hxrt.BoundFunction(func(value *string, encoding *haxe__io__Encoding) {
							hx_receiver_876.__hx_this.writeString(value, encoding)
						}, func() any {
							return hx_receiver_876.__hx_this
						}, "writeString")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "b":
					hx_structural_source_856.b = hx_recover_class_8(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe__ValueException:
		return func(hx_structural_source_877 *haxe__ValueException) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_877, func(name string) hxrt.AnonymousField {
				switch name {
				case "value":
					return hxrt.AnonymousField{Value: hx_structural_source_877.value, Found: true}
				case "unwrap":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_878 := hx_structural_source_877
						return hxrt.BoundFunction(func() any {
							return hx_receiver_878.__hx_this.unwrap()
						}, func() any {
							return hx_receiver_878.__hx_this
						}, "unwrap")
					}(), Found: true}
				case "savedMessage":
					return hxrt.AnonymousField{Value: hx_structural_source_877.savedMessage, Found: true}
				case "savedStack":
					return hxrt.AnonymousField{Value: hx_structural_source_877.savedStack, Found: true}
				case "savedPrevious":
					return hxrt.AnonymousField{Value: hx_structural_source_877.savedPrevious, Found: true}
				case "savedNative":
					return hxrt.AnonymousField{Value: hx_structural_source_877.savedNative, Found: true}
				case "get_message":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_881 := hx_structural_source_877
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_881.__hx_this.get_message()
						}, func() any {
							return hx_receiver_881.__hx_this
						}, "get_message")
					}(), Found: true}
				case "get_stack":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_882 := hx_structural_source_877
						return hxrt.BoundFunction(func() any {
							return hx_receiver_882.__hx_this.get_stack()
						}, func() any {
							return hx_receiver_882.__hx_this
						}, "get_stack")
					}(), Found: true}
				case "get_previous":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *haxe__Exception] {
						hx_receiver_883 := hx_structural_source_877
						return hxrt.BoundFunction(func() *haxe__Exception {
							return hx_receiver_883.__hx_this.get_previous()
						}, func() any {
							return hx_receiver_883.__hx_this
						}, "get_previous")
					}(), Found: true}
				case "get_native":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_884 := hx_structural_source_877
						return hxrt.BoundFunction(func() any {
							return hx_receiver_884.__hx_this.get_native()
						}, func() any {
							return hx_receiver_884.__hx_this
						}, "get_native")
					}(), Found: true}
				case "toString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_885 := hx_structural_source_877
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_885.__hx_this.toString()
						}, func() any {
							return hx_receiver_885.__hx_this
						}, "toString")
					}(), Found: true}
				case "details":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_886 := hx_structural_source_877
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_886.__hx_this.details()
						}, func() any {
							return hx_receiver_886.__hx_this
						}, "details")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "savedMessage":
					hx_structural_source_877.savedMessage = func(hx_value_879 any) *string {
						if hx_value_879 == nil {
							var hx_zero_880 *string
							return hx_zero_880
						}
						return hx_value_879.(*string)
					}(incoming)
					return true
				case "savedStack":
					hx_structural_source_877.savedStack = incoming
					return true
				case "savedPrevious":
					hx_structural_source_877.savedPrevious = hx_recover_class_2(incoming)
					return true
				case "savedNative":
					hx_structural_source_877.savedNative = incoming
					return true
				}
				return false
			})
		}(source)
	case *haxe__exceptions__PosException:
		return func(hx_structural_source_887 *haxe__exceptions__PosException) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_887, func(name string) hxrt.AnonymousField {
				switch name {
				case "posInfos":
					return hxrt.AnonymousField{Value: hx_structural_source_887.posInfos, Found: true}
				case "toString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_888 := hx_structural_source_887
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_888.__hx_this.toString()
						}, func() any {
							return hx_receiver_888.__hx_this
						}, "toString")
					}(), Found: true}
				case "savedMessage":
					return hxrt.AnonymousField{Value: hx_structural_source_887.savedMessage, Found: true}
				case "savedStack":
					return hxrt.AnonymousField{Value: hx_structural_source_887.savedStack, Found: true}
				case "savedPrevious":
					return hxrt.AnonymousField{Value: hx_structural_source_887.savedPrevious, Found: true}
				case "savedNative":
					return hxrt.AnonymousField{Value: hx_structural_source_887.savedNative, Found: true}
				case "get_message":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_891 := hx_structural_source_887
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_891.__hx_this.get_message()
						}, func() any {
							return hx_receiver_891.__hx_this
						}, "get_message")
					}(), Found: true}
				case "get_stack":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_892 := hx_structural_source_887
						return hxrt.BoundFunction(func() any {
							return hx_receiver_892.__hx_this.get_stack()
						}, func() any {
							return hx_receiver_892.__hx_this
						}, "get_stack")
					}(), Found: true}
				case "get_previous":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *haxe__Exception] {
						hx_receiver_893 := hx_structural_source_887
						return hxrt.BoundFunction(func() *haxe__Exception {
							return hx_receiver_893.__hx_this.get_previous()
						}, func() any {
							return hx_receiver_893.__hx_this
						}, "get_previous")
					}(), Found: true}
				case "get_native":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_894 := hx_structural_source_887
						return hxrt.BoundFunction(func() any {
							return hx_receiver_894.__hx_this.get_native()
						}, func() any {
							return hx_receiver_894.__hx_this
						}, "get_native")
					}(), Found: true}
				case "unwrap":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_895 := hx_structural_source_887
						return hxrt.BoundFunction(func() any {
							return hx_receiver_895.__hx_this.unwrap()
						}, func() any {
							return hx_receiver_895.__hx_this
						}, "unwrap")
					}(), Found: true}
				case "details":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_896 := hx_structural_source_887
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_896.__hx_this.details()
						}, func() any {
							return hx_receiver_896.__hx_this
						}, "details")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "savedMessage":
					hx_structural_source_887.savedMessage = func(hx_value_889 any) *string {
						if hx_value_889 == nil {
							var hx_zero_890 *string
							return hx_zero_890
						}
						return hx_value_889.(*string)
					}(incoming)
					return true
				case "savedStack":
					hx_structural_source_887.savedStack = incoming
					return true
				case "savedPrevious":
					hx_structural_source_887.savedPrevious = hx_recover_class_2(incoming)
					return true
				case "savedNative":
					hx_structural_source_887.savedNative = incoming
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
			return func(hx_value_897 any) bool {
				if hx_value_897 == nil {
					var hx_zero_898 bool
					return hx_zero_898
				}
				return hx_value_897.(bool)
			}(any(hxrt.FunctionCall(source)()))
		})
	case *hxrt.Function[func() any]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_899 any) bool {
				if hx_value_899 == nil {
					var hx_zero_900 bool
					return hx_zero_900
				}
				return hx_value_899.(bool)
			}(any(hxrt.FunctionCall(source)()))
		})
	case *hxrt.Function[func() hxrt.Anonymous]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_901 any) bool {
				if hx_value_901 == nil {
					var hx_zero_902 bool
					return hx_zero_902
				}
				return hx_value_901.(bool)
			}(any(hxrt.FunctionCall(source)()))
		})
	case *hxrt.Function[func() int]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_903 any) bool {
				if hx_value_903 == nil {
					var hx_zero_904 bool
					return hx_zero_904
				}
				return hx_value_903.(bool)
			}(any(hxrt.FunctionCall(source)()))
		})
	case *hxrt.Function[func()]:
		return hxrt.ProjectFunction(source, func() bool {
			hxrt.FunctionCall(source)()
			return func(hx_value_905 any) bool {
				if hx_value_905 == nil {
					var hx_zero_906 bool
					return hx_zero_906
				}
				return hx_value_905.(bool)
			}(any(nil))
		})
	case *hxrt.Function[func() float64]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_907 any) bool {
				if hx_value_907 == nil {
					var hx_zero_908 bool
					return hx_zero_908
				}
				return hx_value_907.(bool)
			}(any(hxrt.FunctionCall(source)()))
		})
	case *hxrt.Function[func() *sys__net__Socket]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_909 any) bool {
				if hx_value_909 == nil {
					var hx_zero_910 bool
					return hx_zero_910
				}
				return hx_value_909.(bool)
			}(any(hxrt.FunctionCall(source)()))
		})
	case *hxrt.Function[func() *hxrt.Process]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_911 any) bool {
				if hx_value_911 == nil {
					var hx_zero_912 bool
					return hx_zero_912
				}
				return hx_value_911.(bool)
			}(any(hxrt.FunctionCall(source)()))
		})
	case *hxrt.Function[func() *haxe__io__Bytes]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_913 any) bool {
				if hx_value_913 == nil {
					var hx_zero_914 bool
					return hx_zero_914
				}
				return hx_value_913.(bool)
			}(any(hxrt.FunctionCall(source)()))
		})
	case *hxrt.Function[func() []int]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_915 any) bool {
				if hx_value_915 == nil {
					var hx_zero_916 bool
					return hx_zero_916
				}
				return hx_value_915.(bool)
			}(any(hxrt.FunctionCall(source)()))
		})
	case *hxrt.Function[func() *hxrt.ByteView]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_917 any) bool {
				if hx_value_917 == nil {
					var hx_zero_918 bool
					return hx_zero_918
				}
				return hx_value_917.(bool)
			}(any(hxrt.FunctionCall(source)()))
		})
	case *hxrt.Function[func() *haxe__Exception]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_919 any) bool {
				if hx_value_919 == nil {
					var hx_zero_920 bool
					return hx_zero_920
				}
				return hx_value_919.(bool)
			}(any(hxrt.FunctionCall(source)()))
		})
	case *hxrt.Function[func() haxe__IMap]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_921 any) bool {
				if hx_value_921 == nil {
					var hx_zero_922 bool
					return hx_zero_922
				}
				return hx_value_921.(bool)
			}(any(hxrt.FunctionCall(source)()))
		})
	case *hxrt.Function[func() *haxe__ds__StringMap]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_923 any) bool {
				if hx_value_923 == nil {
					var hx_zero_924 bool
					return hx_zero_924
				}
				return hx_value_923.(bool)
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
			return func(hx_value_925 any) *string {
				if hx_value_925 == nil {
					var hx_zero_926 *string
					return hx_zero_926
				}
				return hx_value_925.(*string)
			}(any(hxrt.FunctionCall(source)()))
		})
	case *hxrt.Function[func() *string]:
		return source
	case *hxrt.Function[func() any]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_927 any) *string {
				if hx_value_927 == nil {
					var hx_zero_928 *string
					return hx_zero_928
				}
				return hx_value_927.(*string)
			}(any(hxrt.FunctionCall(source)()))
		})
	case *hxrt.Function[func() hxrt.Anonymous]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_929 any) *string {
				if hx_value_929 == nil {
					var hx_zero_930 *string
					return hx_zero_930
				}
				return hx_value_929.(*string)
			}(any(hxrt.FunctionCall(source)()))
		})
	case *hxrt.Function[func() int]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_931 any) *string {
				if hx_value_931 == nil {
					var hx_zero_932 *string
					return hx_zero_932
				}
				return hx_value_931.(*string)
			}(any(hxrt.FunctionCall(source)()))
		})
	case *hxrt.Function[func()]:
		return hxrt.ProjectFunction(source, func() *string {
			hxrt.FunctionCall(source)()
			return func(hx_value_933 any) *string {
				if hx_value_933 == nil {
					var hx_zero_934 *string
					return hx_zero_934
				}
				return hx_value_933.(*string)
			}(any(nil))
		})
	case *hxrt.Function[func() float64]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_935 any) *string {
				if hx_value_935 == nil {
					var hx_zero_936 *string
					return hx_zero_936
				}
				return hx_value_935.(*string)
			}(any(hxrt.FunctionCall(source)()))
		})
	case *hxrt.Function[func() *sys__net__Socket]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_937 any) *string {
				if hx_value_937 == nil {
					var hx_zero_938 *string
					return hx_zero_938
				}
				return hx_value_937.(*string)
			}(any(hxrt.FunctionCall(source)()))
		})
	case *hxrt.Function[func() *hxrt.Process]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_939 any) *string {
				if hx_value_939 == nil {
					var hx_zero_940 *string
					return hx_zero_940
				}
				return hx_value_939.(*string)
			}(any(hxrt.FunctionCall(source)()))
		})
	case *hxrt.Function[func() *haxe__io__Bytes]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_941 any) *string {
				if hx_value_941 == nil {
					var hx_zero_942 *string
					return hx_zero_942
				}
				return hx_value_941.(*string)
			}(any(hxrt.FunctionCall(source)()))
		})
	case *hxrt.Function[func() []int]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_943 any) *string {
				if hx_value_943 == nil {
					var hx_zero_944 *string
					return hx_zero_944
				}
				return hx_value_943.(*string)
			}(any(hxrt.FunctionCall(source)()))
		})
	case *hxrt.Function[func() *hxrt.ByteView]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_945 any) *string {
				if hx_value_945 == nil {
					var hx_zero_946 *string
					return hx_zero_946
				}
				return hx_value_945.(*string)
			}(any(hxrt.FunctionCall(source)()))
		})
	case *hxrt.Function[func() *haxe__Exception]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_947 any) *string {
				if hx_value_947 == nil {
					var hx_zero_948 *string
					return hx_zero_948
				}
				return hx_value_947.(*string)
			}(any(hxrt.FunctionCall(source)()))
		})
	case *hxrt.Function[func() haxe__IMap]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_949 any) *string {
				if hx_value_949 == nil {
					var hx_zero_950 *string
					return hx_zero_950
				}
				return hx_value_949.(*string)
			}(any(hxrt.FunctionCall(source)()))
		})
	case *hxrt.Function[func() *haxe__ds__StringMap]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_951 any) *string {
				if hx_value_951 == nil {
					var hx_zero_952 *string
					return hx_zero_952
				}
				return hx_value_951.(*string)
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
	case *hxrt.Function[func() int]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func()]:
		return hxrt.ProjectFunction(source, func() any {
			hxrt.FunctionCall(source)()
			return nil
		})
	case *hxrt.Function[func() float64]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func() *sys__net__Socket]:
		return hxrt.ProjectFunction(source, func() any {
			return hxrt.FunctionCall(source)()
		})
	case *hxrt.Function[func() *hxrt.Process]:
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
	case *hxrt.Function[func(int)]:
		return hxrt.ProjectFunction(source, func(arg_0 *string) {
			hxrt.FunctionCall(source)(func(hx_value_953 any) int {
				if hx_value_953 == nil {
					var hx_zero_954 int
					return hx_zero_954
				}
				return hx_value_953.(int)
			}(any(arg_0)))
		})
	case *hxrt.Function[func(*string)]:
		return source
	case *hxrt.Function[func(*haxe__io__Bytes)]:
		return hxrt.ProjectFunction(source, func(arg_0 *string) {
			hxrt.FunctionCall(source)(hx_recover_class_7(arg_0))
		})
	case *hxrt.Function[func(bool) bool]:
		return hxrt.ProjectFunction(source, func(arg_0 *string) {
			hxrt.FunctionCall(source)(func(hx_value_955 any) bool {
				if hx_value_955 == nil {
					var hx_zero_956 bool
					return hx_zero_956
				}
				return hx_value_955.(bool)
			}(any(arg_0)))
		})
	case *hxrt.Function[func(any) *haxe__io__Bytes]:
		return hxrt.ProjectFunction(source, func(arg_0 *string) {
			hxrt.FunctionCall(source)(arg_0)
		})
	case *hxrt.Function[func(int) *haxe__io__Bytes]:
		return hxrt.ProjectFunction(source, func(arg_0 *string) {
			hxrt.FunctionCall(source)(func(hx_value_957 any) int {
				if hx_value_957 == nil {
					var hx_zero_958 int
					return hx_zero_958
				}
				return hx_value_957.(int)
			}(any(arg_0)))
		})
	case *hxrt.Function[func(int) *string]:
		return hxrt.ProjectFunction(source, func(arg_0 *string) {
			hxrt.FunctionCall(source)(func(hx_value_959 any) int {
				if hx_value_959 == nil {
					var hx_zero_960 int
					return hx_zero_960
				}
				return hx_value_959.(int)
			}(any(arg_0)))
		})
	case *hxrt.Function[func(float64)]:
		return hxrt.ProjectFunction(source, func(arg_0 *string) {
			hxrt.FunctionCall(source)(func(hx_value_961 any) float64 {
				if hx_value_961 == nil {
					var hx_zero_962 float64
					return hx_zero_962
				}
				return hx_value_961.(float64)
			}(any(arg_0)))
		})
	case *hxrt.Function[func(bool)]:
		return hxrt.ProjectFunction(source, func(arg_0 *string) {
			hxrt.FunctionCall(source)(func(hx_value_963 any) bool {
				if hx_value_963 == nil {
					var hx_zero_964 bool
					return hx_zero_964
				}
				return hx_value_963.(bool)
			}(any(arg_0)))
		})
	case *hxrt.Function[func(*hxrt.SocketHandle)]:
		return hxrt.ProjectFunction(source, func(arg_0 *string) {
			hxrt.FunctionCall(source)(func(hx_value_965 any) *hxrt.SocketHandle {
				if hx_value_965 == nil {
					var hx_zero_966 *hxrt.SocketHandle
					return hx_zero_966
				}
				return hx_value_965.(*hxrt.SocketHandle)
			}(any(arg_0)))
		})
	case *hxrt.Function[func(any) any]:
		return hxrt.ProjectFunction(source, func(arg_0 *string) {
			hxrt.FunctionCall(source)(arg_0)
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
			hxrt.FunctionCall(source)(func(hx_value_967 any) *hxrt.HttpExchange {
				if hx_value_967 == nil {
					var hx_zero_968 *hxrt.HttpExchange
					return hx_zero_968
				}
				return hx_value_967.(*hxrt.HttpExchange)
			}(any(arg_0)))
		})
	case *hxrt.Function[func(*haxe___Int64_____Int64)]:
		return hxrt.ProjectFunction(source, func(arg_0 *string) {
			hxrt.FunctionCall(source)(hx_recover_class_9(arg_0))
		})
	case *hxrt.Function[func(int) int]:
		return hxrt.ProjectFunction(source, func(arg_0 *string) {
			hxrt.FunctionCall(source)(func(hx_value_969 any) int {
				if hx_value_969 == nil {
					var hx_zero_970 int
					return hx_zero_970
				}
				return hx_value_969.(int)
			}(any(arg_0)))
		})
	case *hxrt.Function[func(*haxe__io__Bytes) int]:
		return hxrt.ProjectFunction(source, func(arg_0 *string) {
			hxrt.FunctionCall(source)(hx_recover_class_7(arg_0))
		})
	case *hxrt.Function[func(int) float64]:
		return hxrt.ProjectFunction(source, func(arg_0 *string) {
			hxrt.FunctionCall(source)(func(hx_value_971 any) int {
				if hx_value_971 == nil {
					var hx_zero_972 int
					return hx_zero_972
				}
				return hx_value_971.(int)
			}(any(arg_0)))
		})
	case *hxrt.Function[func(int) *haxe___Int64_____Int64]:
		return hxrt.ProjectFunction(source, func(arg_0 *string) {
			hxrt.FunctionCall(source)(func(hx_value_973 any) int {
				if hx_value_973 == nil {
					var hx_zero_974 int
					return hx_zero_974
				}
				return hx_value_973.(int)
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
	case *hxrt.Function[func(any) bool]:
		return hxrt.ProjectFunction(source, func(arg_0 *string) {
			hxrt.FunctionCall(source)(arg_0)
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
	case *hxrt.Function[func(int)]:
		return hxrt.ProjectFunction(source, func(arg_0 *haxe__io__Bytes) {
			hxrt.FunctionCall(source)(func(hx_value_975 any) int {
				if hx_value_975 == nil {
					var hx_zero_976 int
					return hx_zero_976
				}
				return hx_value_975.(int)
			}(any(arg_0)))
		})
	case *hxrt.Function[func(*string)]:
		return hxrt.ProjectFunction(source, func(arg_0 *haxe__io__Bytes) {
			hxrt.FunctionCall(source)(func(hx_value_977 any) *string {
				if hx_value_977 == nil {
					var hx_zero_978 *string
					return hx_zero_978
				}
				return hx_value_977.(*string)
			}(any(arg_0)))
		})
	case *hxrt.Function[func(*haxe__io__Bytes)]:
		return source
	case *hxrt.Function[func(bool) bool]:
		return hxrt.ProjectFunction(source, func(arg_0 *haxe__io__Bytes) {
			hxrt.FunctionCall(source)(func(hx_value_979 any) bool {
				if hx_value_979 == nil {
					var hx_zero_980 bool
					return hx_zero_980
				}
				return hx_value_979.(bool)
			}(any(arg_0)))
		})
	case *hxrt.Function[func(any) *haxe__io__Bytes]:
		return hxrt.ProjectFunction(source, func(arg_0 *haxe__io__Bytes) {
			hxrt.FunctionCall(source)(arg_0)
		})
	case *hxrt.Function[func(int) *haxe__io__Bytes]:
		return hxrt.ProjectFunction(source, func(arg_0 *haxe__io__Bytes) {
			hxrt.FunctionCall(source)(func(hx_value_981 any) int {
				if hx_value_981 == nil {
					var hx_zero_982 int
					return hx_zero_982
				}
				return hx_value_981.(int)
			}(any(arg_0)))
		})
	case *hxrt.Function[func(int) *string]:
		return hxrt.ProjectFunction(source, func(arg_0 *haxe__io__Bytes) {
			hxrt.FunctionCall(source)(func(hx_value_983 any) int {
				if hx_value_983 == nil {
					var hx_zero_984 int
					return hx_zero_984
				}
				return hx_value_983.(int)
			}(any(arg_0)))
		})
	case *hxrt.Function[func(float64)]:
		return hxrt.ProjectFunction(source, func(arg_0 *haxe__io__Bytes) {
			hxrt.FunctionCall(source)(func(hx_value_985 any) float64 {
				if hx_value_985 == nil {
					var hx_zero_986 float64
					return hx_zero_986
				}
				return hx_value_985.(float64)
			}(any(arg_0)))
		})
	case *hxrt.Function[func(bool)]:
		return hxrt.ProjectFunction(source, func(arg_0 *haxe__io__Bytes) {
			hxrt.FunctionCall(source)(func(hx_value_987 any) bool {
				if hx_value_987 == nil {
					var hx_zero_988 bool
					return hx_zero_988
				}
				return hx_value_987.(bool)
			}(any(arg_0)))
		})
	case *hxrt.Function[func(*hxrt.SocketHandle)]:
		return hxrt.ProjectFunction(source, func(arg_0 *haxe__io__Bytes) {
			hxrt.FunctionCall(source)(func(hx_value_989 any) *hxrt.SocketHandle {
				if hx_value_989 == nil {
					var hx_zero_990 *hxrt.SocketHandle
					return hx_zero_990
				}
				return hx_value_989.(*hxrt.SocketHandle)
			}(any(arg_0)))
		})
	case *hxrt.Function[func(any) any]:
		return hxrt.ProjectFunction(source, func(arg_0 *haxe__io__Bytes) {
			hxrt.FunctionCall(source)(arg_0)
		})
	case *hxrt.Function[func(any)]:
		return hxrt.ProjectFunction(source, func(arg_0 *haxe__io__Bytes) {
			hxrt.FunctionCall(source)(arg_0)
		})
	case *hxrt.Function[func(*string) *hxrt.Array]:
		return hxrt.ProjectFunction(source, func(arg_0 *haxe__io__Bytes) {
			hxrt.FunctionCall(source)(func(hx_value_991 any) *string {
				if hx_value_991 == nil {
					var hx_zero_992 *string
					return hx_zero_992
				}
				return hx_value_991.(*string)
			}(any(arg_0)))
		})
	case *hxrt.Function[func(*hxrt.HttpExchange)]:
		return hxrt.ProjectFunction(source, func(arg_0 *haxe__io__Bytes) {
			hxrt.FunctionCall(source)(func(hx_value_993 any) *hxrt.HttpExchange {
				if hx_value_993 == nil {
					var hx_zero_994 *hxrt.HttpExchange
					return hx_zero_994
				}
				return hx_value_993.(*hxrt.HttpExchange)
			}(any(arg_0)))
		})
	case *hxrt.Function[func(*haxe___Int64_____Int64)]:
		return hxrt.ProjectFunction(source, func(arg_0 *haxe__io__Bytes) {
			hxrt.FunctionCall(source)(hx_recover_class_9(arg_0))
		})
	case *hxrt.Function[func(int) int]:
		return hxrt.ProjectFunction(source, func(arg_0 *haxe__io__Bytes) {
			hxrt.FunctionCall(source)(func(hx_value_995 any) int {
				if hx_value_995 == nil {
					var hx_zero_996 int
					return hx_zero_996
				}
				return hx_value_995.(int)
			}(any(arg_0)))
		})
	case *hxrt.Function[func(*haxe__io__Bytes) int]:
		return hxrt.ProjectFunction(source, func(arg_0 *haxe__io__Bytes) {
			hxrt.FunctionCall(source)(arg_0)
		})
	case *hxrt.Function[func(int) float64]:
		return hxrt.ProjectFunction(source, func(arg_0 *haxe__io__Bytes) {
			hxrt.FunctionCall(source)(func(hx_value_997 any) int {
				if hx_value_997 == nil {
					var hx_zero_998 int
					return hx_zero_998
				}
				return hx_value_997.(int)
			}(any(arg_0)))
		})
	case *hxrt.Function[func(int) *haxe___Int64_____Int64]:
		return hxrt.ProjectFunction(source, func(arg_0 *haxe__io__Bytes) {
			hxrt.FunctionCall(source)(func(hx_value_999 any) int {
				if hx_value_999 == nil {
					var hx_zero_1000 int
					return hx_zero_1000
				}
				return hx_value_999.(int)
			}(any(arg_0)))
		})
	case *hxrt.Function[func(*string) any]:
		return hxrt.ProjectFunction(source, func(arg_0 *haxe__io__Bytes) {
			hxrt.FunctionCall(source)(func(hx_value_1001 any) *string {
				if hx_value_1001 == nil {
					var hx_zero_1002 *string
					return hx_zero_1002
				}
				return hx_value_1001.(*string)
			}(any(arg_0)))
		})
	case *hxrt.Function[func(*string) bool]:
		return hxrt.ProjectFunction(source, func(arg_0 *haxe__io__Bytes) {
			hxrt.FunctionCall(source)(func(hx_value_1003 any) *string {
				if hx_value_1003 == nil {
					var hx_zero_1004 *string
					return hx_zero_1004
				}
				return hx_value_1003.(*string)
			}(any(arg_0)))
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
	case *hxrt.Function[func(int)]:
		return source
	case *hxrt.Function[func(*string)]:
		return hxrt.ProjectFunction(source, func(arg_0 int) {
			hxrt.FunctionCall(source)(func(hx_value_1005 any) *string {
				if hx_value_1005 == nil {
					var hx_zero_1006 *string
					return hx_zero_1006
				}
				return hx_value_1005.(*string)
			}(any(arg_0)))
		})
	case *hxrt.Function[func(*haxe__io__Bytes)]:
		return hxrt.ProjectFunction(source, func(arg_0 int) {
			hxrt.FunctionCall(source)(hx_recover_class_7(arg_0))
		})
	case *hxrt.Function[func(bool) bool]:
		return hxrt.ProjectFunction(source, func(arg_0 int) {
			hxrt.FunctionCall(source)(func(hx_value_1007 any) bool {
				if hx_value_1007 == nil {
					var hx_zero_1008 bool
					return hx_zero_1008
				}
				return hx_value_1007.(bool)
			}(any(arg_0)))
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
	case *hxrt.Function[func(float64)]:
		return hxrt.ProjectFunction(source, func(arg_0 int) {
			hxrt.FunctionCall(source)(float64(arg_0))
		})
	case *hxrt.Function[func(bool)]:
		return hxrt.ProjectFunction(source, func(arg_0 int) {
			hxrt.FunctionCall(source)(func(hx_value_1009 any) bool {
				if hx_value_1009 == nil {
					var hx_zero_1010 bool
					return hx_zero_1010
				}
				return hx_value_1009.(bool)
			}(any(arg_0)))
		})
	case *hxrt.Function[func(*hxrt.SocketHandle)]:
		return hxrt.ProjectFunction(source, func(arg_0 int) {
			hxrt.FunctionCall(source)(func(hx_value_1011 any) *hxrt.SocketHandle {
				if hx_value_1011 == nil {
					var hx_zero_1012 *hxrt.SocketHandle
					return hx_zero_1012
				}
				return hx_value_1011.(*hxrt.SocketHandle)
			}(any(arg_0)))
		})
	case *hxrt.Function[func(any) any]:
		return hxrt.ProjectFunction(source, func(arg_0 int) {
			hxrt.FunctionCall(source)(arg_0)
		})
	case *hxrt.Function[func(any)]:
		return hxrt.ProjectFunction(source, func(arg_0 int) {
			hxrt.FunctionCall(source)(arg_0)
		})
	case *hxrt.Function[func(*string) *hxrt.Array]:
		return hxrt.ProjectFunction(source, func(arg_0 int) {
			hxrt.FunctionCall(source)(func(hx_value_1013 any) *string {
				if hx_value_1013 == nil {
					var hx_zero_1014 *string
					return hx_zero_1014
				}
				return hx_value_1013.(*string)
			}(any(arg_0)))
		})
	case *hxrt.Function[func(*hxrt.HttpExchange)]:
		return hxrt.ProjectFunction(source, func(arg_0 int) {
			hxrt.FunctionCall(source)(func(hx_value_1015 any) *hxrt.HttpExchange {
				if hx_value_1015 == nil {
					var hx_zero_1016 *hxrt.HttpExchange
					return hx_zero_1016
				}
				return hx_value_1015.(*hxrt.HttpExchange)
			}(any(arg_0)))
		})
	case *hxrt.Function[func(*haxe___Int64_____Int64)]:
		return hxrt.ProjectFunction(source, func(arg_0 int) {
			hxrt.FunctionCall(source)(hx_recover_class_9(arg_0))
		})
	case *hxrt.Function[func(int) int]:
		return hxrt.ProjectFunction(source, func(arg_0 int) {
			hxrt.FunctionCall(source)(arg_0)
		})
	case *hxrt.Function[func(*haxe__io__Bytes) int]:
		return hxrt.ProjectFunction(source, func(arg_0 int) {
			hxrt.FunctionCall(source)(hx_recover_class_7(arg_0))
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
			hxrt.FunctionCall(source)(func(hx_value_1017 any) *string {
				if hx_value_1017 == nil {
					var hx_zero_1018 *string
					return hx_zero_1018
				}
				return hx_value_1017.(*string)
			}(any(arg_0)))
		})
	case *hxrt.Function[func(*string) bool]:
		return hxrt.ProjectFunction(source, func(arg_0 int) {
			hxrt.FunctionCall(source)(func(hx_value_1019 any) *string {
				if hx_value_1019 == nil {
					var hx_zero_1020 *string
					return hx_zero_1020
				}
				return hx_value_1019.(*string)
			}(any(arg_0)))
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

func hx_recover_class_0(value any) *sys__io__Process {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *sys__io__Process:
		return source
	default:
		return value.(*sys__io__Process)
	}
}

func hx_recover_class_1(value any) *sys__net__Socket {
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

func hx_recover_class_2(value any) *haxe__Exception {
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

func hx_recover_class_3(value any) *haxe__io__Eof {
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

func hx_recover_class_4(value any) *haxe__io__Input {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *_Main__TrackingInput:
		return source.haxe__io__Input
	case *sys__net__SocketInput:
		return source.haxe__io__Input
	case *sys__io___Process__ProcessOutput:
		return source.haxe__io__Input
	case *haxe__io__Input:
		return source
	default:
		return value.(*haxe__io__Input)
	}
}

func hx_recover_class_5(value any) *haxe__ds__StringMap {
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

func hx_recover_class_6(value any) *haxe__ValueException {
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

func hx_recover_class_7(value any) *haxe__io__Bytes {
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

func hx_recover_class_8(value any) *haxe__io__BytesBuffer {
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

func hx_recover_class_9(value any) *haxe___Int64_____Int64 {
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

func hx_is_nominal_0(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *haxe__io__Eof:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_1(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *haxe__ValueException:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_2(value any) bool {
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

func hx_is_nominal_3(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *_Main__TrackingInput:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_4(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *haxe___Int64_____Int64:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_5(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *haxe__ds__StringMap:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_6(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *haxe__exceptions__NotImplementedException:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_7(value any) bool {
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

func hx_is_nominal_8(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *sys__Http:
		return (source != nil)
	case *haxe__http__HttpBase:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_9(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *haxe__io__Bytes:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_10(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *haxe__io__BytesBuffer:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_11(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *haxe__io__BytesOutput:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_12(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *_Main__TrackingInput:
		return (source != nil)
	case *sys__net__SocketInput:
		return (source != nil)
	case *sys__io___Process__ProcessOutput:
		return (source != nil)
	case *haxe__io__Input:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_13(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *sys__net__SocketOutput:
		return (source != nil)
	case *sys__io___Process__ProcessInput:
		return (source != nil)
	case *haxe__io__Output:
		return (source != nil)
	case *haxe__io__BytesOutput:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_14(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *haxe__iterators__MapKeyValueIterator:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_15(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *haxe__iterators__StringIterator:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_16(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *haxe__iterators__StringKeyValueIterator:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_17(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *sys__Http:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_18(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *sys__io__Process:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_19(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *sys__io___Process__ProcessInput:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_20(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *sys__io___Process__ProcessOutput:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_21(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *sys__net__Host:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_22(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *sys__net__Socket:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_23(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *sys__net__SocketInput:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_24(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *sys__net__SocketOutput:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_25(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *sys__thread__Tls:
		return (source != nil)
	default:
		return false
	}
}
