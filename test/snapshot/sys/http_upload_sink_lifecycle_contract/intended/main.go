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
		hx_recover_class_sys__io__Process(hxrt.AnonymousGet(server, "process")).__hx_this.close()
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

func (self *_Main__TrackingInput) HxClassName() string {
	return "_Main.TrackingInput"
}

func hx_structural_recover(value any) hxrt.Anonymous {
	switch source := value.(type) {
	case *_Main__TrackingInput:
		return func(hx_structural_source_1 *_Main__TrackingInput) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_1, func(name string) hxrt.AnonymousField {
				switch name {
				case "reads":
					return hxrt.AnonymousField{Value: hx_structural_source_1.reads, Found: true}
				case "wrongThread":
					return hxrt.AnonymousField{Value: hx_structural_source_1.wrongThread, Found: true}
				case "afterReturnReads":
					return hxrt.AnonymousField{Value: hx_structural_source_1.afterReturnReads, Found: true}
				case "requestReturned":
					return hxrt.AnonymousField{Value: hx_structural_source_1.requestReturned, Found: true}
				case "mode":
					return hxrt.AnonymousField{Value: hx_structural_source_1.mode, Found: true}
				case "callerMarker":
					return hxrt.AnonymousField{Value: hx_structural_source_1.callerMarker, Found: true}
				case "payload":
					return hxrt.AnonymousField{Value: hx_structural_source_1.payload, Found: true}
				case "position":
					return hxrt.AnonymousField{Value: hx_structural_source_1.position, Found: true}
				case "readBytes":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes, int, int) int] {
						hx_receiver_6 := hx_structural_source_1
						return hxrt.BoundFunction(func(bytes *haxe__io__Bytes, targetPos int, requested int) int {
							return hx_receiver_6.__hx_this.readBytes(bytes, targetPos, requested)
						}, func() any {
							return hx_receiver_6.__hx_this
						}, "readBytes")
					}(), Found: true}
				case "bigEndian":
					return hxrt.AnonymousField{Value: hx_structural_source_1.bigEndian, Found: true}
				case "readByte":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_7 := hx_structural_source_1
						return hxrt.BoundFunction(func() int {
							return hx_receiver_7.__hx_this.readByte()
						}, func() any {
							return hx_receiver_7.__hx_this
						}, "readByte")
					}(), Found: true}
				case "close":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func()] {
						hx_receiver_8 := hx_structural_source_1
						return hxrt.BoundFunction(func() {
							hx_receiver_8.__hx_this.close()
						}, func() any {
							return hx_receiver_8.__hx_this
						}, "close")
					}(), Found: true}
				case "set_bigEndian":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(bool) bool] {
						hx_receiver_9 := hx_structural_source_1
						return hxrt.BoundFunction(func(value bool) bool {
							return hx_receiver_9.__hx_this.set_bigEndian(value)
						}, func() any {
							return hx_receiver_9.__hx_this
						}, "set_bigEndian")
					}(), Found: true}
				case "readAll":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(any) *haxe__io__Bytes] {
						hx_receiver_10 := hx_structural_source_1
						return hxrt.BoundFunction(func(bufsize any) *haxe__io__Bytes {
							return hx_receiver_10.__hx_this.readAll(bufsize)
						}, func() any {
							return hx_receiver_10.__hx_this
						}, "readAll")
					}(), Found: true}
				case "readFullBytes":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes, int, int)] {
						hx_receiver_11 := hx_structural_source_1
						return hxrt.BoundFunction(func(bytes *haxe__io__Bytes, pos int, len int) {
							hx_receiver_11.__hx_this.readFullBytes(bytes, pos, len)
						}, func() any {
							return hx_receiver_11.__hx_this
						}, "readFullBytes")
					}(), Found: true}
				case "read":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int) *haxe__io__Bytes] {
						hx_receiver_12 := hx_structural_source_1
						return hxrt.BoundFunction(func(nbytes int) *haxe__io__Bytes {
							return hx_receiver_12.__hx_this.read(nbytes)
						}, func() any {
							return hx_receiver_12.__hx_this
						}, "read")
					}(), Found: true}
				case "readUntil":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int) *string] {
						hx_receiver_13 := hx_structural_source_1
						return hxrt.BoundFunction(func(end int) *string {
							return hx_receiver_13.__hx_this.readUntil(end)
						}, func() any {
							return hx_receiver_13.__hx_this
						}, "readUntil")
					}(), Found: true}
				case "readLine":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_14 := hx_structural_source_1
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_14.__hx_this.readLine()
						}, func() any {
							return hx_receiver_14.__hx_this
						}, "readLine")
					}(), Found: true}
				case "readFloat":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() float64] {
						hx_receiver_15 := hx_structural_source_1
						return hxrt.BoundFunction(func() float64 {
							return hx_receiver_15.__hx_this.readFloat()
						}, func() any {
							return hx_receiver_15.__hx_this
						}, "readFloat")
					}(), Found: true}
				case "readDouble":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() float64] {
						hx_receiver_16 := hx_structural_source_1
						return hxrt.BoundFunction(func() float64 {
							return hx_receiver_16.__hx_this.readDouble()
						}, func() any {
							return hx_receiver_16.__hx_this
						}, "readDouble")
					}(), Found: true}
				case "readInt8":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_17 := hx_structural_source_1
						return hxrt.BoundFunction(func() int {
							return hx_receiver_17.__hx_this.readInt8()
						}, func() any {
							return hx_receiver_17.__hx_this
						}, "readInt8")
					}(), Found: true}
				case "readInt16":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_18 := hx_structural_source_1
						return hxrt.BoundFunction(func() int {
							return hx_receiver_18.__hx_this.readInt16()
						}, func() any {
							return hx_receiver_18.__hx_this
						}, "readInt16")
					}(), Found: true}
				case "readUInt16":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_19 := hx_structural_source_1
						return hxrt.BoundFunction(func() int {
							return hx_receiver_19.__hx_this.readUInt16()
						}, func() any {
							return hx_receiver_19.__hx_this
						}, "readUInt16")
					}(), Found: true}
				case "readInt24":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_20 := hx_structural_source_1
						return hxrt.BoundFunction(func() int {
							return hx_receiver_20.__hx_this.readInt24()
						}, func() any {
							return hx_receiver_20.__hx_this
						}, "readInt24")
					}(), Found: true}
				case "readUInt24":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_21 := hx_structural_source_1
						return hxrt.BoundFunction(func() int {
							return hx_receiver_21.__hx_this.readUInt24()
						}, func() any {
							return hx_receiver_21.__hx_this
						}, "readUInt24")
					}(), Found: true}
				case "readInt32":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_22 := hx_structural_source_1
						return hxrt.BoundFunction(func() int {
							return hx_receiver_22.__hx_this.readInt32()
						}, func() any {
							return hx_receiver_22.__hx_this
						}, "readInt32")
					}(), Found: true}
				case "readString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, *haxe__io__Encoding) *string] {
						hx_receiver_23 := hx_structural_source_1
						return hxrt.BoundFunction(func(len int, encoding *haxe__io__Encoding) *string {
							return hx_receiver_23.__hx_this.readString(len, encoding)
						}, func() any {
							return hx_receiver_23.__hx_this
						}, "readString")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "requestReturned":
					hx_structural_source_1.requestReturned = func(hx_value_2 any) bool {
						if hx_value_2 == nil {
							var hx_zero_3 bool
							return hx_zero_3
						}
						return hx_value_2.(bool)
					}(incoming)
					return true
				case "position":
					hx_structural_source_1.position = func(hx_value_4 any) int {
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
	case *haxe___Int64_____Int64:
		return func(hx_structural_source_24 *haxe___Int64_____Int64) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_24, func(name string) hxrt.AnonymousField {
				switch name {
				case "high":
					return hxrt.AnonymousField{Value: hx_structural_source_24.high, Found: true}
				case "low":
					return hxrt.AnonymousField{Value: hx_structural_source_24.low, Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "high":
					hx_structural_source_24.high = func(hx_value_25 any) int {
						if hx_value_25 == nil {
							var hx_zero_26 int
							return hx_zero_26
						}
						return hx_value_25.(int)
					}(incoming)
					return true
				case "low":
					hx_structural_source_24.low = func(hx_value_27 any) int {
						if hx_value_27 == nil {
							var hx_zero_28 int
							return hx_zero_28
						}
						return hx_value_27.(int)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *sys__net__SocketOutput:
		return func(hx_structural_source_29 *sys__net__SocketOutput) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_29, func(name string) hxrt.AnonymousField {
				switch name {
				case "handle":
					return hxrt.AnonymousField{Value: hx_structural_source_29.handle, Found: true}
				case "writeByte":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_32 := hx_structural_source_29
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_32.__hx_this.writeByte(value)
						}, func() any {
							return hx_receiver_32.__hx_this
						}, "writeByte")
					}(), Found: true}
				case "writeBytes":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes, int, int) int] {
						hx_receiver_33 := hx_structural_source_29
						return hxrt.BoundFunction(func(bytes *haxe__io__Bytes, pos int, length int) int {
							return hx_receiver_33.__hx_this.writeBytes(bytes, pos, length)
						}, func() any {
							return hx_receiver_33.__hx_this
						}, "writeBytes")
					}(), Found: true}
				case "flush":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func()] {
						hx_receiver_34 := hx_structural_source_29
						return hxrt.BoundFunction(func() {
							hx_receiver_34.__hx_this.flush()
						}, func() any {
							return hx_receiver_34.__hx_this
						}, "flush")
					}(), Found: true}
				case "close":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func()] {
						hx_receiver_35 := hx_structural_source_29
						return hxrt.BoundFunction(func() {
							hx_receiver_35.__hx_this.close()
						}, func() any {
							return hx_receiver_35.__hx_this
						}, "close")
					}(), Found: true}
				case "bigEndian":
					return hxrt.AnonymousField{Value: hx_structural_source_29.bigEndian, Found: true}
				case "set_bigEndian":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(bool) bool] {
						hx_receiver_36 := hx_structural_source_29
						return hxrt.BoundFunction(func(value bool) bool {
							return hx_receiver_36.__hx_this.set_bigEndian(value)
						}, func() any {
							return hx_receiver_36.__hx_this
						}, "set_bigEndian")
					}(), Found: true}
				case "write":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes)] {
						hx_receiver_37 := hx_structural_source_29
						return hxrt.BoundFunction(func(bytes *haxe__io__Bytes) {
							hx_receiver_37.__hx_this.write(bytes)
						}, func() any {
							return hx_receiver_37.__hx_this
						}, "write")
					}(), Found: true}
				case "writeFullBytes":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes, int, int)] {
						hx_receiver_38 := hx_structural_source_29
						return hxrt.BoundFunction(func(bytes *haxe__io__Bytes, pos int, len int) {
							hx_receiver_38.__hx_this.writeFullBytes(bytes, pos, len)
						}, func() any {
							return hx_receiver_38.__hx_this
						}, "writeFullBytes")
					}(), Found: true}
				case "writeFloat":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(float64)] {
						hx_receiver_39 := hx_structural_source_29
						return hxrt.BoundFunction(func(value float64) {
							hx_receiver_39.__hx_this.writeFloat(value)
						}, func() any {
							return hx_receiver_39.__hx_this
						}, "writeFloat")
					}(), Found: true}
				case "writeDouble":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(float64)] {
						hx_receiver_40 := hx_structural_source_29
						return hxrt.BoundFunction(func(value float64) {
							hx_receiver_40.__hx_this.writeDouble(value)
						}, func() any {
							return hx_receiver_40.__hx_this
						}, "writeDouble")
					}(), Found: true}
				case "writeInt8":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_41 := hx_structural_source_29
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_41.__hx_this.writeInt8(value)
						}, func() any {
							return hx_receiver_41.__hx_this
						}, "writeInt8")
					}(), Found: true}
				case "writeInt16":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_42 := hx_structural_source_29
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_42.__hx_this.writeInt16(value)
						}, func() any {
							return hx_receiver_42.__hx_this
						}, "writeInt16")
					}(), Found: true}
				case "writeUInt16":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_43 := hx_structural_source_29
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_43.__hx_this.writeUInt16(value)
						}, func() any {
							return hx_receiver_43.__hx_this
						}, "writeUInt16")
					}(), Found: true}
				case "writeInt24":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_44 := hx_structural_source_29
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_44.__hx_this.writeInt24(value)
						}, func() any {
							return hx_receiver_44.__hx_this
						}, "writeInt24")
					}(), Found: true}
				case "writeUInt24":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_45 := hx_structural_source_29
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_45.__hx_this.writeUInt24(value)
						}, func() any {
							return hx_receiver_45.__hx_this
						}, "writeUInt24")
					}(), Found: true}
				case "writeInt32":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_46 := hx_structural_source_29
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_46.__hx_this.writeInt32(value)
						}, func() any {
							return hx_receiver_46.__hx_this
						}, "writeInt32")
					}(), Found: true}
				case "prepare":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_47 := hx_structural_source_29
						return hxrt.BoundFunction(func(nbytes int) {
							hx_receiver_47.__hx_this.prepare(nbytes)
						}, func() any {
							return hx_receiver_47.__hx_this
						}, "prepare")
					}(), Found: true}
				case "writeInput":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Input, any)] {
						hx_receiver_48 := hx_structural_source_29
						return hxrt.BoundFunction(func(input *haxe__io__Input, bufsize any) {
							hx_receiver_48.__hx_this.writeInput(input, bufsize)
						}, func() any {
							return hx_receiver_48.__hx_this
						}, "writeInput")
					}(), Found: true}
				case "writeString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string, *haxe__io__Encoding)] {
						hx_receiver_49 := hx_structural_source_29
						return hxrt.BoundFunction(func(value *string, encoding *haxe__io__Encoding) {
							hx_receiver_49.__hx_this.writeString(value, encoding)
						}, func() any {
							return hx_receiver_49.__hx_this
						}, "writeString")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "handle":
					hx_structural_source_29.handle = func(hx_value_30 any) *hxrt.SocketHandle {
						if hx_value_30 == nil {
							var hx_zero_31 *hxrt.SocketHandle
							return hx_zero_31
						}
						return hx_value_30.(*hxrt.SocketHandle)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *sys__net__SocketInput:
		return func(hx_structural_source_50 *sys__net__SocketInput) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_50, func(name string) hxrt.AnonymousField {
				switch name {
				case "handle":
					return hxrt.AnonymousField{Value: hx_structural_source_50.handle, Found: true}
				case "readByte":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_53 := hx_structural_source_50
						return hxrt.BoundFunction(func() int {
							return hx_receiver_53.__hx_this.readByte()
						}, func() any {
							return hx_receiver_53.__hx_this
						}, "readByte")
					}(), Found: true}
				case "readBytes":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes, int, int) int] {
						hx_receiver_54 := hx_structural_source_50
						return hxrt.BoundFunction(func(bytes *haxe__io__Bytes, pos int, length int) int {
							return hx_receiver_54.__hx_this.readBytes(bytes, pos, length)
						}, func() any {
							return hx_receiver_54.__hx_this
						}, "readBytes")
					}(), Found: true}
				case "close":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func()] {
						hx_receiver_55 := hx_structural_source_50
						return hxrt.BoundFunction(func() {
							hx_receiver_55.__hx_this.close()
						}, func() any {
							return hx_receiver_55.__hx_this
						}, "close")
					}(), Found: true}
				case "bigEndian":
					return hxrt.AnonymousField{Value: hx_structural_source_50.bigEndian, Found: true}
				case "set_bigEndian":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(bool) bool] {
						hx_receiver_56 := hx_structural_source_50
						return hxrt.BoundFunction(func(value bool) bool {
							return hx_receiver_56.__hx_this.set_bigEndian(value)
						}, func() any {
							return hx_receiver_56.__hx_this
						}, "set_bigEndian")
					}(), Found: true}
				case "readAll":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(any) *haxe__io__Bytes] {
						hx_receiver_57 := hx_structural_source_50
						return hxrt.BoundFunction(func(bufsize any) *haxe__io__Bytes {
							return hx_receiver_57.__hx_this.readAll(bufsize)
						}, func() any {
							return hx_receiver_57.__hx_this
						}, "readAll")
					}(), Found: true}
				case "readFullBytes":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes, int, int)] {
						hx_receiver_58 := hx_structural_source_50
						return hxrt.BoundFunction(func(bytes *haxe__io__Bytes, pos int, len int) {
							hx_receiver_58.__hx_this.readFullBytes(bytes, pos, len)
						}, func() any {
							return hx_receiver_58.__hx_this
						}, "readFullBytes")
					}(), Found: true}
				case "read":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int) *haxe__io__Bytes] {
						hx_receiver_59 := hx_structural_source_50
						return hxrt.BoundFunction(func(nbytes int) *haxe__io__Bytes {
							return hx_receiver_59.__hx_this.read(nbytes)
						}, func() any {
							return hx_receiver_59.__hx_this
						}, "read")
					}(), Found: true}
				case "readUntil":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int) *string] {
						hx_receiver_60 := hx_structural_source_50
						return hxrt.BoundFunction(func(end int) *string {
							return hx_receiver_60.__hx_this.readUntil(end)
						}, func() any {
							return hx_receiver_60.__hx_this
						}, "readUntil")
					}(), Found: true}
				case "readLine":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_61 := hx_structural_source_50
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_61.__hx_this.readLine()
						}, func() any {
							return hx_receiver_61.__hx_this
						}, "readLine")
					}(), Found: true}
				case "readFloat":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() float64] {
						hx_receiver_62 := hx_structural_source_50
						return hxrt.BoundFunction(func() float64 {
							return hx_receiver_62.__hx_this.readFloat()
						}, func() any {
							return hx_receiver_62.__hx_this
						}, "readFloat")
					}(), Found: true}
				case "readDouble":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() float64] {
						hx_receiver_63 := hx_structural_source_50
						return hxrt.BoundFunction(func() float64 {
							return hx_receiver_63.__hx_this.readDouble()
						}, func() any {
							return hx_receiver_63.__hx_this
						}, "readDouble")
					}(), Found: true}
				case "readInt8":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_64 := hx_structural_source_50
						return hxrt.BoundFunction(func() int {
							return hx_receiver_64.__hx_this.readInt8()
						}, func() any {
							return hx_receiver_64.__hx_this
						}, "readInt8")
					}(), Found: true}
				case "readInt16":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_65 := hx_structural_source_50
						return hxrt.BoundFunction(func() int {
							return hx_receiver_65.__hx_this.readInt16()
						}, func() any {
							return hx_receiver_65.__hx_this
						}, "readInt16")
					}(), Found: true}
				case "readUInt16":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_66 := hx_structural_source_50
						return hxrt.BoundFunction(func() int {
							return hx_receiver_66.__hx_this.readUInt16()
						}, func() any {
							return hx_receiver_66.__hx_this
						}, "readUInt16")
					}(), Found: true}
				case "readInt24":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_67 := hx_structural_source_50
						return hxrt.BoundFunction(func() int {
							return hx_receiver_67.__hx_this.readInt24()
						}, func() any {
							return hx_receiver_67.__hx_this
						}, "readInt24")
					}(), Found: true}
				case "readUInt24":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_68 := hx_structural_source_50
						return hxrt.BoundFunction(func() int {
							return hx_receiver_68.__hx_this.readUInt24()
						}, func() any {
							return hx_receiver_68.__hx_this
						}, "readUInt24")
					}(), Found: true}
				case "readInt32":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_69 := hx_structural_source_50
						return hxrt.BoundFunction(func() int {
							return hx_receiver_69.__hx_this.readInt32()
						}, func() any {
							return hx_receiver_69.__hx_this
						}, "readInt32")
					}(), Found: true}
				case "readString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, *haxe__io__Encoding) *string] {
						hx_receiver_70 := hx_structural_source_50
						return hxrt.BoundFunction(func(len int, encoding *haxe__io__Encoding) *string {
							return hx_receiver_70.__hx_this.readString(len, encoding)
						}, func() any {
							return hx_receiver_70.__hx_this
						}, "readString")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "handle":
					hx_structural_source_50.handle = func(hx_value_51 any) *hxrt.SocketHandle {
						if hx_value_51 == nil {
							var hx_zero_52 *hxrt.SocketHandle
							return hx_zero_52
						}
						return hx_value_51.(*hxrt.SocketHandle)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *sys__net__Socket:
		return func(hx_structural_source_71 *sys__net__Socket) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_71, func(name string) hxrt.AnonymousField {
				switch name {
				case "input":
					return hxrt.AnonymousField{Value: hx_structural_source_71.input, Found: true}
				case "output":
					return hxrt.AnonymousField{Value: hx_structural_source_71.output, Found: true}
				case "custom":
					return hxrt.AnonymousField{Value: hx_structural_source_71.custom, Found: true}
				case "handle":
					return hxrt.AnonymousField{Value: hx_structural_source_71.handle, Found: true}
				case "close":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func()] {
						hx_receiver_74 := hx_structural_source_71
						return hxrt.BoundFunction(func() {
							hx_receiver_74.__hx_this.close()
						}, func() any {
							return hx_receiver_74.__hx_this
						}, "close")
					}(), Found: true}
				case "read":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_75 := hx_structural_source_71
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_75.__hx_this.read()
						}, func() any {
							return hx_receiver_75.__hx_this
						}, "read")
					}(), Found: true}
				case "write":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string)] {
						hx_receiver_76 := hx_structural_source_71
						return hxrt.BoundFunction(func(content *string) {
							hx_receiver_76.__hx_this.write(content)
						}, func() any {
							return hx_receiver_76.__hx_this
						}, "write")
					}(), Found: true}
				case "connect":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*sys__net__Host, int)] {
						hx_receiver_77 := hx_structural_source_71
						return hxrt.BoundFunction(func(host *sys__net__Host, port int) {
							hx_receiver_77.__hx_this.connect(host, port)
						}, func() any {
							return hx_receiver_77.__hx_this
						}, "connect")
					}(), Found: true}
				case "listen":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_78 := hx_structural_source_71
						return hxrt.BoundFunction(func(connections int) {
							hx_receiver_78.__hx_this.listen(connections)
						}, func() any {
							return hx_receiver_78.__hx_this
						}, "listen")
					}(), Found: true}
				case "shutdown":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(bool, bool)] {
						hx_receiver_79 := hx_structural_source_71
						return hxrt.BoundFunction(func(read bool, write bool) {
							hx_receiver_79.__hx_this.shutdown(read, write)
						}, func() any {
							return hx_receiver_79.__hx_this
						}, "shutdown")
					}(), Found: true}
				case "bind":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*sys__net__Host, int)] {
						hx_receiver_80 := hx_structural_source_71
						return hxrt.BoundFunction(func(host *sys__net__Host, port int) {
							hx_receiver_80.__hx_this.bind(host, port)
						}, func() any {
							return hx_receiver_80.__hx_this
						}, "bind")
					}(), Found: true}
				case "accept":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *sys__net__Socket] {
						hx_receiver_81 := hx_structural_source_71
						return hxrt.BoundFunction(func() *sys__net__Socket {
							return hx_receiver_81.__hx_this.accept()
						}, func() any {
							return hx_receiver_81.__hx_this
						}, "accept")
					}(), Found: true}
				case "peer":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() hxrt.Anonymous] {
						hx_receiver_82 := hx_structural_source_71
						return hxrt.BoundFunction(func() hxrt.Anonymous {
							return hx_receiver_82.__hx_this.peer()
						}, func() any {
							return hx_receiver_82.__hx_this
						}, "peer")
					}(), Found: true}
				case "host":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() hxrt.Anonymous] {
						hx_receiver_83 := hx_structural_source_71
						return hxrt.BoundFunction(func() hxrt.Anonymous {
							return hx_receiver_83.__hx_this.host()
						}, func() any {
							return hx_receiver_83.__hx_this
						}, "host")
					}(), Found: true}
				case "setTimeout":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(float64)] {
						hx_receiver_84 := hx_structural_source_71
						return hxrt.BoundFunction(func(timeout float64) {
							hx_receiver_84.__hx_this.setTimeout(timeout)
						}, func() any {
							return hx_receiver_84.__hx_this
						}, "setTimeout")
					}(), Found: true}
				case "waitForRead":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func()] {
						hx_receiver_85 := hx_structural_source_71
						return hxrt.BoundFunction(func() {
							hx_receiver_85.__hx_this.waitForRead()
						}, func() any {
							return hx_receiver_85.__hx_this
						}, "waitForRead")
					}(), Found: true}
				case "setBlocking":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(bool)] {
						hx_receiver_86 := hx_structural_source_71
						return hxrt.BoundFunction(func(blocking bool) {
							hx_receiver_86.__hx_this.setBlocking(blocking)
						}, func() any {
							return hx_receiver_86.__hx_this
						}, "setBlocking")
					}(), Found: true}
				case "setFastSend":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(bool)] {
						hx_receiver_87 := hx_structural_source_71
						return hxrt.BoundFunction(func(fastSend bool) {
							hx_receiver_87.__hx_this.setFastSend(fastSend)
						}, func() any {
							return hx_receiver_87.__hx_this
						}, "setFastSend")
					}(), Found: true}
				case "replaceHandle":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*hxrt.SocketHandle)] {
						hx_receiver_88 := hx_structural_source_71
						return hxrt.BoundFunction(func(next *hxrt.SocketHandle) {
							hx_receiver_88.__hx_this.replaceHandle(next)
						}, func() any {
							return hx_receiver_88.__hx_this
						}, "replaceHandle")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "custom":
					hx_structural_source_71.custom = incoming
					return true
				case "handle":
					hx_structural_source_71.handle = func(hx_value_72 any) *hxrt.SocketHandle {
						if hx_value_72 == nil {
							var hx_zero_73 *hxrt.SocketHandle
							return hx_zero_73
						}
						return hx_value_72.(*hxrt.SocketHandle)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *sys__net__Host:
		return func(hx_structural_source_89 *sys__net__Host) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_89, func(name string) hxrt.AnonymousField {
				switch name {
				case "host":
					return hxrt.AnonymousField{Value: hx_structural_source_89.host, Found: true}
				case "ip":
					return hxrt.AnonymousField{Value: hx_structural_source_89.ip, Found: true}
				case "toString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_90 := hx_structural_source_89
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_90.__hx_this.toString()
						}, func() any {
							return hx_receiver_90.__hx_this
						}, "toString")
					}(), Found: true}
				case "reverse":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_91 := hx_structural_source_89
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_91.__hx_this.reverse()
						}, func() any {
							return hx_receiver_91.__hx_this
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
		return func(hx_structural_source_92 *sys__io___Process__ProcessOutput) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_92, func(name string) hxrt.AnonymousField {
				switch name {
				case "handle":
					return hxrt.AnonymousField{Value: hx_structural_source_92.handle, Found: true}
				case "readByte":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_95 := hx_structural_source_92
						return hxrt.BoundFunction(func() int {
							return hx_receiver_95.__hx_this.readByte()
						}, func() any {
							return hx_receiver_95.__hx_this
						}, "readByte")
					}(), Found: true}
				case "readBytes":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes, int, int) int] {
						hx_receiver_96 := hx_structural_source_92
						return hxrt.BoundFunction(func(bytes *haxe__io__Bytes, pos int, length int) int {
							return hx_receiver_96.__hx_this.readBytes(bytes, pos, length)
						}, func() any {
							return hx_receiver_96.__hx_this
						}, "readBytes")
					}(), Found: true}
				case "close":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func()] {
						hx_receiver_97 := hx_structural_source_92
						return hxrt.BoundFunction(func() {
							hx_receiver_97.__hx_this.close()
						}, func() any {
							return hx_receiver_97.__hx_this
						}, "close")
					}(), Found: true}
				case "bigEndian":
					return hxrt.AnonymousField{Value: hx_structural_source_92.bigEndian, Found: true}
				case "set_bigEndian":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(bool) bool] {
						hx_receiver_98 := hx_structural_source_92
						return hxrt.BoundFunction(func(value bool) bool {
							return hx_receiver_98.__hx_this.set_bigEndian(value)
						}, func() any {
							return hx_receiver_98.__hx_this
						}, "set_bigEndian")
					}(), Found: true}
				case "readAll":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(any) *haxe__io__Bytes] {
						hx_receiver_99 := hx_structural_source_92
						return hxrt.BoundFunction(func(bufsize any) *haxe__io__Bytes {
							return hx_receiver_99.__hx_this.readAll(bufsize)
						}, func() any {
							return hx_receiver_99.__hx_this
						}, "readAll")
					}(), Found: true}
				case "readFullBytes":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes, int, int)] {
						hx_receiver_100 := hx_structural_source_92
						return hxrt.BoundFunction(func(bytes *haxe__io__Bytes, pos int, len int) {
							hx_receiver_100.__hx_this.readFullBytes(bytes, pos, len)
						}, func() any {
							return hx_receiver_100.__hx_this
						}, "readFullBytes")
					}(), Found: true}
				case "read":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int) *haxe__io__Bytes] {
						hx_receiver_101 := hx_structural_source_92
						return hxrt.BoundFunction(func(nbytes int) *haxe__io__Bytes {
							return hx_receiver_101.__hx_this.read(nbytes)
						}, func() any {
							return hx_receiver_101.__hx_this
						}, "read")
					}(), Found: true}
				case "readUntil":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int) *string] {
						hx_receiver_102 := hx_structural_source_92
						return hxrt.BoundFunction(func(end int) *string {
							return hx_receiver_102.__hx_this.readUntil(end)
						}, func() any {
							return hx_receiver_102.__hx_this
						}, "readUntil")
					}(), Found: true}
				case "readLine":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_103 := hx_structural_source_92
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_103.__hx_this.readLine()
						}, func() any {
							return hx_receiver_103.__hx_this
						}, "readLine")
					}(), Found: true}
				case "readFloat":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() float64] {
						hx_receiver_104 := hx_structural_source_92
						return hxrt.BoundFunction(func() float64 {
							return hx_receiver_104.__hx_this.readFloat()
						}, func() any {
							return hx_receiver_104.__hx_this
						}, "readFloat")
					}(), Found: true}
				case "readDouble":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() float64] {
						hx_receiver_105 := hx_structural_source_92
						return hxrt.BoundFunction(func() float64 {
							return hx_receiver_105.__hx_this.readDouble()
						}, func() any {
							return hx_receiver_105.__hx_this
						}, "readDouble")
					}(), Found: true}
				case "readInt8":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_106 := hx_structural_source_92
						return hxrt.BoundFunction(func() int {
							return hx_receiver_106.__hx_this.readInt8()
						}, func() any {
							return hx_receiver_106.__hx_this
						}, "readInt8")
					}(), Found: true}
				case "readInt16":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_107 := hx_structural_source_92
						return hxrt.BoundFunction(func() int {
							return hx_receiver_107.__hx_this.readInt16()
						}, func() any {
							return hx_receiver_107.__hx_this
						}, "readInt16")
					}(), Found: true}
				case "readUInt16":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_108 := hx_structural_source_92
						return hxrt.BoundFunction(func() int {
							return hx_receiver_108.__hx_this.readUInt16()
						}, func() any {
							return hx_receiver_108.__hx_this
						}, "readUInt16")
					}(), Found: true}
				case "readInt24":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_109 := hx_structural_source_92
						return hxrt.BoundFunction(func() int {
							return hx_receiver_109.__hx_this.readInt24()
						}, func() any {
							return hx_receiver_109.__hx_this
						}, "readInt24")
					}(), Found: true}
				case "readUInt24":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_110 := hx_structural_source_92
						return hxrt.BoundFunction(func() int {
							return hx_receiver_110.__hx_this.readUInt24()
						}, func() any {
							return hx_receiver_110.__hx_this
						}, "readUInt24")
					}(), Found: true}
				case "readInt32":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_111 := hx_structural_source_92
						return hxrt.BoundFunction(func() int {
							return hx_receiver_111.__hx_this.readInt32()
						}, func() any {
							return hx_receiver_111.__hx_this
						}, "readInt32")
					}(), Found: true}
				case "readString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, *haxe__io__Encoding) *string] {
						hx_receiver_112 := hx_structural_source_92
						return hxrt.BoundFunction(func(len int, encoding *haxe__io__Encoding) *string {
							return hx_receiver_112.__hx_this.readString(len, encoding)
						}, func() any {
							return hx_receiver_112.__hx_this
						}, "readString")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "handle":
					hx_structural_source_92.handle = func(hx_value_93 any) *hxrt.ProcessOutput {
						if hx_value_93 == nil {
							var hx_zero_94 *hxrt.ProcessOutput
							return hx_zero_94
						}
						return hx_value_93.(*hxrt.ProcessOutput)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *sys__io___Process__ProcessInput:
		return func(hx_structural_source_113 *sys__io___Process__ProcessInput) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_113, func(name string) hxrt.AnonymousField {
				switch name {
				case "handle":
					return hxrt.AnonymousField{Value: hx_structural_source_113.handle, Found: true}
				case "writeByte":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_116 := hx_structural_source_113
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_116.__hx_this.writeByte(value)
						}, func() any {
							return hx_receiver_116.__hx_this
						}, "writeByte")
					}(), Found: true}
				case "writeBytes":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes, int, int) int] {
						hx_receiver_117 := hx_structural_source_113
						return hxrt.BoundFunction(func(bytes *haxe__io__Bytes, pos int, length int) int {
							return hx_receiver_117.__hx_this.writeBytes(bytes, pos, length)
						}, func() any {
							return hx_receiver_117.__hx_this
						}, "writeBytes")
					}(), Found: true}
				case "flush":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func()] {
						hx_receiver_118 := hx_structural_source_113
						return hxrt.BoundFunction(func() {
							hx_receiver_118.__hx_this.flush()
						}, func() any {
							return hx_receiver_118.__hx_this
						}, "flush")
					}(), Found: true}
				case "close":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func()] {
						hx_receiver_119 := hx_structural_source_113
						return hxrt.BoundFunction(func() {
							hx_receiver_119.__hx_this.close()
						}, func() any {
							return hx_receiver_119.__hx_this
						}, "close")
					}(), Found: true}
				case "bigEndian":
					return hxrt.AnonymousField{Value: hx_structural_source_113.bigEndian, Found: true}
				case "set_bigEndian":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(bool) bool] {
						hx_receiver_120 := hx_structural_source_113
						return hxrt.BoundFunction(func(value bool) bool {
							return hx_receiver_120.__hx_this.set_bigEndian(value)
						}, func() any {
							return hx_receiver_120.__hx_this
						}, "set_bigEndian")
					}(), Found: true}
				case "write":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes)] {
						hx_receiver_121 := hx_structural_source_113
						return hxrt.BoundFunction(func(bytes *haxe__io__Bytes) {
							hx_receiver_121.__hx_this.write(bytes)
						}, func() any {
							return hx_receiver_121.__hx_this
						}, "write")
					}(), Found: true}
				case "writeFullBytes":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes, int, int)] {
						hx_receiver_122 := hx_structural_source_113
						return hxrt.BoundFunction(func(bytes *haxe__io__Bytes, pos int, len int) {
							hx_receiver_122.__hx_this.writeFullBytes(bytes, pos, len)
						}, func() any {
							return hx_receiver_122.__hx_this
						}, "writeFullBytes")
					}(), Found: true}
				case "writeFloat":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(float64)] {
						hx_receiver_123 := hx_structural_source_113
						return hxrt.BoundFunction(func(value float64) {
							hx_receiver_123.__hx_this.writeFloat(value)
						}, func() any {
							return hx_receiver_123.__hx_this
						}, "writeFloat")
					}(), Found: true}
				case "writeDouble":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(float64)] {
						hx_receiver_124 := hx_structural_source_113
						return hxrt.BoundFunction(func(value float64) {
							hx_receiver_124.__hx_this.writeDouble(value)
						}, func() any {
							return hx_receiver_124.__hx_this
						}, "writeDouble")
					}(), Found: true}
				case "writeInt8":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_125 := hx_structural_source_113
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_125.__hx_this.writeInt8(value)
						}, func() any {
							return hx_receiver_125.__hx_this
						}, "writeInt8")
					}(), Found: true}
				case "writeInt16":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_126 := hx_structural_source_113
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_126.__hx_this.writeInt16(value)
						}, func() any {
							return hx_receiver_126.__hx_this
						}, "writeInt16")
					}(), Found: true}
				case "writeUInt16":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_127 := hx_structural_source_113
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_127.__hx_this.writeUInt16(value)
						}, func() any {
							return hx_receiver_127.__hx_this
						}, "writeUInt16")
					}(), Found: true}
				case "writeInt24":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_128 := hx_structural_source_113
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_128.__hx_this.writeInt24(value)
						}, func() any {
							return hx_receiver_128.__hx_this
						}, "writeInt24")
					}(), Found: true}
				case "writeUInt24":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_129 := hx_structural_source_113
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_129.__hx_this.writeUInt24(value)
						}, func() any {
							return hx_receiver_129.__hx_this
						}, "writeUInt24")
					}(), Found: true}
				case "writeInt32":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_130 := hx_structural_source_113
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_130.__hx_this.writeInt32(value)
						}, func() any {
							return hx_receiver_130.__hx_this
						}, "writeInt32")
					}(), Found: true}
				case "prepare":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_131 := hx_structural_source_113
						return hxrt.BoundFunction(func(nbytes int) {
							hx_receiver_131.__hx_this.prepare(nbytes)
						}, func() any {
							return hx_receiver_131.__hx_this
						}, "prepare")
					}(), Found: true}
				case "writeInput":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Input, any)] {
						hx_receiver_132 := hx_structural_source_113
						return hxrt.BoundFunction(func(input *haxe__io__Input, bufsize any) {
							hx_receiver_132.__hx_this.writeInput(input, bufsize)
						}, func() any {
							return hx_receiver_132.__hx_this
						}, "writeInput")
					}(), Found: true}
				case "writeString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string, *haxe__io__Encoding)] {
						hx_receiver_133 := hx_structural_source_113
						return hxrt.BoundFunction(func(value *string, encoding *haxe__io__Encoding) {
							hx_receiver_133.__hx_this.writeString(value, encoding)
						}, func() any {
							return hx_receiver_133.__hx_this
						}, "writeString")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "handle":
					hx_structural_source_113.handle = func(hx_value_114 any) *hxrt.ProcessInput {
						if hx_value_114 == nil {
							var hx_zero_115 *hxrt.ProcessInput
							return hx_zero_115
						}
						return hx_value_114.(*hxrt.ProcessInput)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *sys__io__Process:
		return func(hx_structural_source_134 *sys__io__Process) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_134, func(name string) hxrt.AnonymousField {
				switch name {
				case "stdout":
					return hxrt.AnonymousField{Value: hx_structural_source_134.stdout, Found: true}
				case "stderr":
					return hxrt.AnonymousField{Value: hx_structural_source_134.stderr, Found: true}
				case "stdin":
					return hxrt.AnonymousField{Value: hx_structural_source_134.stdin, Found: true}
				case "handle":
					return hxrt.AnonymousField{Value: hx_structural_source_134.handle, Found: true}
				case "getPid":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_137 := hx_structural_source_134
						return hxrt.BoundFunction(func() int {
							return hx_receiver_137.__hx_this.getPid()
						}, func() any {
							return hx_receiver_137.__hx_this
						}, "getPid")
					}(), Found: true}
				case "exitCode":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(any) any] {
						hx_receiver_138 := hx_structural_source_134
						return hxrt.BoundFunction(func(block any) any {
							if block == nil {
								block = true
							}
							return hx_receiver_138.__hx_this.exitCode(func(hx_value_139 any) bool {
								if hx_value_139 == nil {
									var hx_zero_140 bool
									return hx_zero_140
								}
								return hx_value_139.(bool)
							}(block))
						}, func() any {
							return hx_receiver_138.__hx_this
						}, "exitCode")
					}(), Found: true}
				case "close":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func()] {
						hx_receiver_141 := hx_structural_source_134
						return hxrt.BoundFunction(func() {
							hx_receiver_141.__hx_this.close()
						}, func() any {
							return hx_receiver_141.__hx_this
						}, "close")
					}(), Found: true}
				case "kill":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func()] {
						hx_receiver_142 := hx_structural_source_134
						return hxrt.BoundFunction(func() {
							hx_receiver_142.__hx_this.kill()
						}, func() any {
							return hx_receiver_142.__hx_this
						}, "kill")
					}(), Found: true}
				case "requireHandle":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *hxrt.Process] {
						hx_receiver_143 := hx_structural_source_134
						return hxrt.BoundFunction(func() *hxrt.Process {
							return hx_receiver_143.__hx_this.requireHandle()
						}, func() any {
							return hx_receiver_143.__hx_this
						}, "requireHandle")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "handle":
					hx_structural_source_134.handle = func(hx_value_135 any) *hxrt.Process {
						if hx_value_135 == nil {
							var hx_zero_136 *hxrt.Process
							return hx_zero_136
						}
						return hx_value_135.(*hxrt.Process)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *sys__Http:
		return func(hx_structural_source_144 *sys__Http) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_144, func(name string) hxrt.AnonymousField {
				switch name {
				case "noShutdown":
					return hxrt.AnonymousField{Value: hx_structural_source_144.noShutdown, Found: true}
				case "cnxTimeout":
					return hxrt.AnonymousField{Value: hx_structural_source_144.cnxTimeout, Found: true}
				case "responseHeaders":
					return hxrt.AnonymousField{Value: hx_structural_source_144.responseHeaders, Found: true}
				case "responseHeadersSameKey":
					return hxrt.AnonymousField{Value: hx_structural_source_144.responseHeadersSameKey, Found: true}
				case "file":
					return hxrt.AnonymousField{Value: hx_structural_source_144.file, Found: true}
				case "request":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(any)] {
						hx_receiver_151 := hx_structural_source_144
						return hxrt.BoundFunction(func(post any) {
							hx_receiver_151.__hx_this.request(post)
						}, func() any {
							return hx_receiver_151.__hx_this
						}, "request")
					}(), Found: true}
				case "fileTransfert":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string, *string, *haxe__io__Input, int, *string)] {
						hx_receiver_152 := hx_structural_source_144
						return hxrt.BoundFunction(func(argname *string, filename *string, file *haxe__io__Input, size int, mimeType *string) {
							if mimeType == nil {
								mimeType = hxrt.StringFromLiteral("application/octet-stream")
							}
							hx_receiver_152.__hx_this.fileTransfert(argname, filename, file, size, mimeType)
						}, func() any {
							return hx_receiver_152.__hx_this
						}, "fileTransfert")
					}(), Found: true}
				case "fileTransfer":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string, *string, *haxe__io__Input, int, *string)] {
						hx_receiver_153 := hx_structural_source_144
						return hxrt.BoundFunction(func(argname *string, filename *string, file *haxe__io__Input, size int, mimeType *string) {
							if mimeType == nil {
								mimeType = hxrt.StringFromLiteral("application/octet-stream")
							}
							hx_receiver_153.__hx_this.fileTransfer(argname, filename, file, size, mimeType)
						}, func() any {
							return hx_receiver_153.__hx_this
						}, "fileTransfer")
					}(), Found: true}
				case "customRequest":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(bool, *haxe__io__Output, *sys__net__Socket, *string)] {
						hx_receiver_154 := hx_structural_source_144
						return hxrt.BoundFunction(func(post bool, api *haxe__io__Output, sock *sys__net__Socket, method *string) {
							hx_receiver_154.__hx_this.customRequest(post, api, sock, method)
						}, func() any {
							return hx_receiver_154.__hx_this
						}, "customRequest")
					}(), Found: true}
				case "getResponseHeaderValues":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string) *hxrt.Array] {
						hx_receiver_155 := hx_structural_source_144
						return hxrt.BoundFunction(func(key *string) *hxrt.Array {
							return hx_receiver_155.__hx_this.getResponseHeaderValues(key)
						}, func() any {
							return hx_receiver_155.__hx_this
						}, "getResponseHeaderValues")
					}(), Found: true}
				case "requestWith":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(bool, *haxe__io__Output, *sys__net__Socket, *string)] {
						hx_receiver_156 := hx_structural_source_144
						return hxrt.BoundFunction(func(post bool, api *haxe__io__Output, sock *sys__net__Socket, method *string) {
							hx_receiver_156.__hx_this.requestWith(post, api, sock, method)
						}, func() any {
							return hx_receiver_156.__hx_this
						}, "requestWith")
					}(), Found: true}
				case "pumpUpload":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*hxrt.HttpExchange, hxrt.Anonymous) hxrt.Anonymous] {
						hx_receiver_157 := hx_structural_source_144
						return hxrt.BoundFunction(func(exchange *hxrt.HttpExchange, upload hxrt.Anonymous) hxrt.Anonymous {
							return hx_receiver_157.__hx_this.pumpUpload(exchange, upload)
						}, func() any {
							return hx_receiver_157.__hx_this
						}, "pumpUpload")
					}(), Found: true}
				case "handleDataRequest":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(bool, *haxe__io__Output, *string)] {
						hx_receiver_158 := hx_structural_source_144
						return hxrt.BoundFunction(func(post bool, api *haxe__io__Output, method *string) {
							hx_receiver_158.__hx_this.handleDataRequest(post, api, method)
						}, func() any {
							return hx_receiver_158.__hx_this
						}, "handleDataRequest")
					}(), Found: true}
				case "recordResponseHeaders":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*hxrt.HttpExchange)] {
						hx_receiver_159 := hx_structural_source_144
						return hxrt.BoundFunction(func(exchange *hxrt.HttpExchange) {
							hx_receiver_159.__hx_this.recordResponseHeaders(exchange)
						}, func() any {
							return hx_receiver_159.__hx_this
						}, "recordResponseHeaders")
					}(), Found: true}
				case "resetResponseHeaders":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func()] {
						hx_receiver_160 := hx_structural_source_144
						return hxrt.BoundFunction(func() {
							hx_receiver_160.__hx_this.resetResponseHeaders()
						}, func() any {
							return hx_receiver_160.__hx_this
						}, "resetResponseHeaders")
					}(), Found: true}
				case "encodedParameters":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_161 := hx_structural_source_144
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_161.__hx_this.encodedParameters()
						}, func() any {
							return hx_receiver_161.__hx_this
						}, "encodedParameters")
					}(), Found: true}
				case "url":
					return hxrt.AnonymousField{Value: hx_structural_source_144.url, Found: true}
				case "responseBytes":
					return hxrt.AnonymousField{Value: hx_structural_source_144.responseBytes, Found: true}
				case "responseAsString":
					return hxrt.AnonymousField{Value: hx_structural_source_144.responseAsString, Found: true}
				case "postData":
					return hxrt.AnonymousField{Value: hx_structural_source_144.postData, Found: true}
				case "postBytes":
					return hxrt.AnonymousField{Value: hx_structural_source_144.postBytes, Found: true}
				case "headers":
					return hxrt.AnonymousField{Value: hx_structural_source_144.headers, Found: true}
				case "params":
					return hxrt.AnonymousField{Value: hx_structural_source_144.params, Found: true}
				case "emptyOnData":
					return hxrt.AnonymousField{Value: hx_structural_source_144.emptyOnData, Found: true}
				case "setHeader":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string, *string)] {
						hx_receiver_172 := hx_structural_source_144
						return hxrt.BoundFunction(func(name *string, value *string) {
							hx_receiver_172.__hx_this.setHeader(name, value)
						}, func() any {
							return hx_receiver_172.__hx_this
						}, "setHeader")
					}(), Found: true}
				case "addHeader":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string, *string)] {
						hx_receiver_173 := hx_structural_source_144
						return hxrt.BoundFunction(func(header *string, value *string) {
							hx_receiver_173.__hx_this.addHeader(header, value)
						}, func() any {
							return hx_receiver_173.__hx_this
						}, "addHeader")
					}(), Found: true}
				case "setParameter":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string, *string)] {
						hx_receiver_174 := hx_structural_source_144
						return hxrt.BoundFunction(func(name *string, value *string) {
							hx_receiver_174.__hx_this.setParameter(name, value)
						}, func() any {
							return hx_receiver_174.__hx_this
						}, "setParameter")
					}(), Found: true}
				case "addParameter":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string, *string)] {
						hx_receiver_175 := hx_structural_source_144
						return hxrt.BoundFunction(func(name *string, value *string) {
							hx_receiver_175.__hx_this.addParameter(name, value)
						}, func() any {
							return hx_receiver_175.__hx_this
						}, "addParameter")
					}(), Found: true}
				case "setPostData":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string)] {
						hx_receiver_176 := hx_structural_source_144
						return hxrt.BoundFunction(func(data *string) {
							hx_receiver_176.__hx_this.setPostData(data)
						}, func() any {
							return hx_receiver_176.__hx_this
						}, "setPostData")
					}(), Found: true}
				case "setPostBytes":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes)] {
						hx_receiver_177 := hx_structural_source_144
						return hxrt.BoundFunction(func(data *haxe__io__Bytes) {
							hx_receiver_177.__hx_this.setPostBytes(data)
						}, func() any {
							return hx_receiver_177.__hx_this
						}, "setPostBytes")
					}(), Found: true}
				case "onData":
					return hxrt.AnonymousField{Value: hxrt.ReadDynamicFunction(hx_structural_source_144.onData), Found: true}
				case "onBytes":
					return hxrt.AnonymousField{Value: hxrt.ReadDynamicFunction(hx_structural_source_144.onBytes), Found: true}
				case "onError":
					return hxrt.AnonymousField{Value: hxrt.ReadDynamicFunction(hx_structural_source_144.onError), Found: true}
				case "onStatus":
					return hxrt.AnonymousField{Value: hxrt.ReadDynamicFunction(hx_structural_source_144.onStatus), Found: true}
				case "hasOnData":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() bool] {
						hx_receiver_178 := hx_structural_source_144
						return hxrt.BoundFunction(func() bool {
							return hx_receiver_178.__hx_this.hasOnData()
						}, func() any {
							return hx_receiver_178.__hx_this
						}, "hasOnData")
					}(), Found: true}
				case "success":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes)] {
						hx_receiver_179 := hx_structural_source_144
						return hxrt.BoundFunction(func(data *haxe__io__Bytes) {
							hx_receiver_179.__hx_this.success(data)
						}, func() any {
							return hx_receiver_179.__hx_this
						}, "success")
					}(), Found: true}
				case "get_responseData":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_180 := hx_structural_source_144
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_180.__hx_this.get_responseData()
						}, func() any {
							return hx_receiver_180.__hx_this
						}, "get_responseData")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "noShutdown":
					hx_structural_source_144.noShutdown = func(hx_value_145 any) bool {
						if hx_value_145 == nil {
							var hx_zero_146 bool
							return hx_zero_146
						}
						return hx_value_145.(bool)
					}(incoming)
					return true
				case "cnxTimeout":
					hx_structural_source_144.cnxTimeout = func(hx_value_147 any) float64 {
						if hx_value_147 == nil {
							var hx_zero_148 float64
							return hx_zero_148
						}
						return hx_value_147.(float64)
					}(incoming)
					return true
				case "responseHeaders":
					hx_structural_source_144.responseHeaders = func(hx_value_149 any) *haxe__ds__StringMap {
						if hx_value_149 == nil {
							var hx_zero_150 *haxe__ds__StringMap
							return hx_zero_150
						}
						return hx_value_149.(*haxe__ds__StringMap)
					}(incoming)
					return true
				case "responseHeadersSameKey":
					hx_structural_source_144.responseHeadersSameKey = hx_recover_class_haxe__ds__StringMap(incoming)
					return true
				case "file":
					hx_structural_source_144.file = hx_structural_recover(incoming)
					return true
				case "url":
					hx_structural_source_144.url = func(hx_value_162 any) *string {
						if hx_value_162 == nil {
							var hx_zero_163 *string
							return hx_zero_163
						}
						return hx_value_162.(*string)
					}(incoming)
					return true
				case "responseAsString":
					hx_structural_source_144.responseAsString = func(hx_value_164 any) *string {
						if hx_value_164 == nil {
							var hx_zero_165 *string
							return hx_zero_165
						}
						return hx_value_164.(*string)
					}(incoming)
					return true
				case "postData":
					hx_structural_source_144.postData = func(hx_value_166 any) *string {
						if hx_value_166 == nil {
							var hx_zero_167 *string
							return hx_zero_167
						}
						return hx_value_166.(*string)
					}(incoming)
					return true
				case "postBytes":
					hx_structural_source_144.postBytes = hx_recover_class_haxe__io__Bytes(incoming)
					return true
				case "headers":
					hx_structural_source_144.headers = func(hx_value_168 any) *hxrt.Array {
						if hx_value_168 == nil {
							var hx_zero_169 *hxrt.Array
							return hx_zero_169
						}
						return hx_value_168.(*hxrt.Array)
					}(incoming)
					return true
				case "params":
					hx_structural_source_144.params = func(hx_value_170 any) *hxrt.Array {
						if hx_value_170 == nil {
							var hx_zero_171 *hxrt.Array
							return hx_zero_171
						}
						return hx_value_170.(*hxrt.Array)
					}(incoming)
					return true
				case "onData":
					hx_structural_source_144.onData = hx_recover_callable_9da3ad96(incoming)
					return true
				case "onBytes":
					hx_structural_source_144.onBytes = hx_recover_callable_c5f1b802(incoming)
					return true
				case "onError":
					hx_structural_source_144.onError = hx_recover_callable_9da3ad96(incoming)
					return true
				case "onStatus":
					hx_structural_source_144.onStatus = hx_recover_callable_d8be58e9(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe__io__Output:
		return func(hx_structural_source_181 *haxe__io__Output) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_181, func(name string) hxrt.AnonymousField {
				switch name {
				case "bigEndian":
					return hxrt.AnonymousField{Value: hx_structural_source_181.bigEndian, Found: true}
				case "writeByte":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_182 := hx_structural_source_181
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_182.__hx_this.writeByte(value)
						}, func() any {
							return hx_receiver_182.__hx_this
						}, "writeByte")
					}(), Found: true}
				case "writeBytes":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes, int, int) int] {
						hx_receiver_183 := hx_structural_source_181
						return hxrt.BoundFunction(func(bytes *haxe__io__Bytes, pos int, len int) int {
							return hx_receiver_183.__hx_this.writeBytes(bytes, pos, len)
						}, func() any {
							return hx_receiver_183.__hx_this
						}, "writeBytes")
					}(), Found: true}
				case "flush":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func()] {
						hx_receiver_184 := hx_structural_source_181
						return hxrt.BoundFunction(func() {
							hx_receiver_184.__hx_this.flush()
						}, func() any {
							return hx_receiver_184.__hx_this
						}, "flush")
					}(), Found: true}
				case "close":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func()] {
						hx_receiver_185 := hx_structural_source_181
						return hxrt.BoundFunction(func() {
							hx_receiver_185.__hx_this.close()
						}, func() any {
							return hx_receiver_185.__hx_this
						}, "close")
					}(), Found: true}
				case "set_bigEndian":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(bool) bool] {
						hx_receiver_186 := hx_structural_source_181
						return hxrt.BoundFunction(func(value bool) bool {
							return hx_receiver_186.__hx_this.set_bigEndian(value)
						}, func() any {
							return hx_receiver_186.__hx_this
						}, "set_bigEndian")
					}(), Found: true}
				case "write":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes)] {
						hx_receiver_187 := hx_structural_source_181
						return hxrt.BoundFunction(func(bytes *haxe__io__Bytes) {
							hx_receiver_187.__hx_this.write(bytes)
						}, func() any {
							return hx_receiver_187.__hx_this
						}, "write")
					}(), Found: true}
				case "writeFullBytes":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes, int, int)] {
						hx_receiver_188 := hx_structural_source_181
						return hxrt.BoundFunction(func(bytes *haxe__io__Bytes, pos int, len int) {
							hx_receiver_188.__hx_this.writeFullBytes(bytes, pos, len)
						}, func() any {
							return hx_receiver_188.__hx_this
						}, "writeFullBytes")
					}(), Found: true}
				case "writeFloat":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(float64)] {
						hx_receiver_189 := hx_structural_source_181
						return hxrt.BoundFunction(func(value float64) {
							hx_receiver_189.__hx_this.writeFloat(value)
						}, func() any {
							return hx_receiver_189.__hx_this
						}, "writeFloat")
					}(), Found: true}
				case "writeDouble":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(float64)] {
						hx_receiver_190 := hx_structural_source_181
						return hxrt.BoundFunction(func(value float64) {
							hx_receiver_190.__hx_this.writeDouble(value)
						}, func() any {
							return hx_receiver_190.__hx_this
						}, "writeDouble")
					}(), Found: true}
				case "writeInt8":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_191 := hx_structural_source_181
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_191.__hx_this.writeInt8(value)
						}, func() any {
							return hx_receiver_191.__hx_this
						}, "writeInt8")
					}(), Found: true}
				case "writeInt16":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_192 := hx_structural_source_181
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_192.__hx_this.writeInt16(value)
						}, func() any {
							return hx_receiver_192.__hx_this
						}, "writeInt16")
					}(), Found: true}
				case "writeUInt16":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_193 := hx_structural_source_181
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_193.__hx_this.writeUInt16(value)
						}, func() any {
							return hx_receiver_193.__hx_this
						}, "writeUInt16")
					}(), Found: true}
				case "writeInt24":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_194 := hx_structural_source_181
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_194.__hx_this.writeInt24(value)
						}, func() any {
							return hx_receiver_194.__hx_this
						}, "writeInt24")
					}(), Found: true}
				case "writeUInt24":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_195 := hx_structural_source_181
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_195.__hx_this.writeUInt24(value)
						}, func() any {
							return hx_receiver_195.__hx_this
						}, "writeUInt24")
					}(), Found: true}
				case "writeInt32":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_196 := hx_structural_source_181
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_196.__hx_this.writeInt32(value)
						}, func() any {
							return hx_receiver_196.__hx_this
						}, "writeInt32")
					}(), Found: true}
				case "prepare":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_197 := hx_structural_source_181
						return hxrt.BoundFunction(func(nbytes int) {
							hx_receiver_197.__hx_this.prepare(nbytes)
						}, func() any {
							return hx_receiver_197.__hx_this
						}, "prepare")
					}(), Found: true}
				case "writeInput":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Input, any)] {
						hx_receiver_198 := hx_structural_source_181
						return hxrt.BoundFunction(func(input *haxe__io__Input, bufsize any) {
							hx_receiver_198.__hx_this.writeInput(input, bufsize)
						}, func() any {
							return hx_receiver_198.__hx_this
						}, "writeInput")
					}(), Found: true}
				case "writeString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string, *haxe__io__Encoding)] {
						hx_receiver_199 := hx_structural_source_181
						return hxrt.BoundFunction(func(value *string, encoding *haxe__io__Encoding) {
							hx_receiver_199.__hx_this.writeString(value, encoding)
						}, func() any {
							return hx_receiver_199.__hx_this
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
		return func(hx_structural_source_200 *haxe__io__Input) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_200, func(name string) hxrt.AnonymousField {
				switch name {
				case "bigEndian":
					return hxrt.AnonymousField{Value: hx_structural_source_200.bigEndian, Found: true}
				case "readByte":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_201 := hx_structural_source_200
						return hxrt.BoundFunction(func() int {
							return hx_receiver_201.__hx_this.readByte()
						}, func() any {
							return hx_receiver_201.__hx_this
						}, "readByte")
					}(), Found: true}
				case "readBytes":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes, int, int) int] {
						hx_receiver_202 := hx_structural_source_200
						return hxrt.BoundFunction(func(bytes *haxe__io__Bytes, pos int, len int) int {
							return hx_receiver_202.__hx_this.readBytes(bytes, pos, len)
						}, func() any {
							return hx_receiver_202.__hx_this
						}, "readBytes")
					}(), Found: true}
				case "close":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func()] {
						hx_receiver_203 := hx_structural_source_200
						return hxrt.BoundFunction(func() {
							hx_receiver_203.__hx_this.close()
						}, func() any {
							return hx_receiver_203.__hx_this
						}, "close")
					}(), Found: true}
				case "set_bigEndian":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(bool) bool] {
						hx_receiver_204 := hx_structural_source_200
						return hxrt.BoundFunction(func(value bool) bool {
							return hx_receiver_204.__hx_this.set_bigEndian(value)
						}, func() any {
							return hx_receiver_204.__hx_this
						}, "set_bigEndian")
					}(), Found: true}
				case "readAll":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(any) *haxe__io__Bytes] {
						hx_receiver_205 := hx_structural_source_200
						return hxrt.BoundFunction(func(bufsize any) *haxe__io__Bytes {
							return hx_receiver_205.__hx_this.readAll(bufsize)
						}, func() any {
							return hx_receiver_205.__hx_this
						}, "readAll")
					}(), Found: true}
				case "readFullBytes":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes, int, int)] {
						hx_receiver_206 := hx_structural_source_200
						return hxrt.BoundFunction(func(bytes *haxe__io__Bytes, pos int, len int) {
							hx_receiver_206.__hx_this.readFullBytes(bytes, pos, len)
						}, func() any {
							return hx_receiver_206.__hx_this
						}, "readFullBytes")
					}(), Found: true}
				case "read":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int) *haxe__io__Bytes] {
						hx_receiver_207 := hx_structural_source_200
						return hxrt.BoundFunction(func(nbytes int) *haxe__io__Bytes {
							return hx_receiver_207.__hx_this.read(nbytes)
						}, func() any {
							return hx_receiver_207.__hx_this
						}, "read")
					}(), Found: true}
				case "readUntil":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int) *string] {
						hx_receiver_208 := hx_structural_source_200
						return hxrt.BoundFunction(func(end int) *string {
							return hx_receiver_208.__hx_this.readUntil(end)
						}, func() any {
							return hx_receiver_208.__hx_this
						}, "readUntil")
					}(), Found: true}
				case "readLine":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_209 := hx_structural_source_200
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_209.__hx_this.readLine()
						}, func() any {
							return hx_receiver_209.__hx_this
						}, "readLine")
					}(), Found: true}
				case "readFloat":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() float64] {
						hx_receiver_210 := hx_structural_source_200
						return hxrt.BoundFunction(func() float64 {
							return hx_receiver_210.__hx_this.readFloat()
						}, func() any {
							return hx_receiver_210.__hx_this
						}, "readFloat")
					}(), Found: true}
				case "readDouble":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() float64] {
						hx_receiver_211 := hx_structural_source_200
						return hxrt.BoundFunction(func() float64 {
							return hx_receiver_211.__hx_this.readDouble()
						}, func() any {
							return hx_receiver_211.__hx_this
						}, "readDouble")
					}(), Found: true}
				case "readInt8":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_212 := hx_structural_source_200
						return hxrt.BoundFunction(func() int {
							return hx_receiver_212.__hx_this.readInt8()
						}, func() any {
							return hx_receiver_212.__hx_this
						}, "readInt8")
					}(), Found: true}
				case "readInt16":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_213 := hx_structural_source_200
						return hxrt.BoundFunction(func() int {
							return hx_receiver_213.__hx_this.readInt16()
						}, func() any {
							return hx_receiver_213.__hx_this
						}, "readInt16")
					}(), Found: true}
				case "readUInt16":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_214 := hx_structural_source_200
						return hxrt.BoundFunction(func() int {
							return hx_receiver_214.__hx_this.readUInt16()
						}, func() any {
							return hx_receiver_214.__hx_this
						}, "readUInt16")
					}(), Found: true}
				case "readInt24":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_215 := hx_structural_source_200
						return hxrt.BoundFunction(func() int {
							return hx_receiver_215.__hx_this.readInt24()
						}, func() any {
							return hx_receiver_215.__hx_this
						}, "readInt24")
					}(), Found: true}
				case "readUInt24":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_216 := hx_structural_source_200
						return hxrt.BoundFunction(func() int {
							return hx_receiver_216.__hx_this.readUInt24()
						}, func() any {
							return hx_receiver_216.__hx_this
						}, "readUInt24")
					}(), Found: true}
				case "readInt32":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_217 := hx_structural_source_200
						return hxrt.BoundFunction(func() int {
							return hx_receiver_217.__hx_this.readInt32()
						}, func() any {
							return hx_receiver_217.__hx_this
						}, "readInt32")
					}(), Found: true}
				case "readString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, *haxe__io__Encoding) *string] {
						hx_receiver_218 := hx_structural_source_200
						return hxrt.BoundFunction(func(len int, encoding *haxe__io__Encoding) *string {
							return hx_receiver_218.__hx_this.readString(len, encoding)
						}, func() any {
							return hx_receiver_218.__hx_this
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
		return func(hx_structural_source_219 *haxe__io__Eof) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_219, func(name string) hxrt.AnonymousField {
				switch name {
				case "toString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_220 := hx_structural_source_219
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_220.__hx_this.toString()
						}, func() any {
							return hx_receiver_220.__hx_this
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
		return func(hx_structural_source_221 *haxe__io__BytesBuffer) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_221, func(name string) hxrt.AnonymousField {
				switch name {
				case "b":
					return hxrt.AnonymousField{Value: hx_structural_source_221.b, Found: true}
				case "get_length":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_224 := hx_structural_source_221
						return hxrt.BoundFunction(func() int {
							return hx_receiver_224.__hx_this.get_length()
						}, func() any {
							return hx_receiver_224.__hx_this
						}, "get_length")
					}(), Found: true}
				case "addByte":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_225 := hx_structural_source_221
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_225.__hx_this.addByte(value)
						}, func() any {
							return hx_receiver_225.__hx_this
						}, "addByte")
					}(), Found: true}
				case "add":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes)] {
						hx_receiver_226 := hx_structural_source_221
						return hxrt.BoundFunction(func(source *haxe__io__Bytes) {
							hx_receiver_226.__hx_this.add(source)
						}, func() any {
							return hx_receiver_226.__hx_this
						}, "add")
					}(), Found: true}
				case "addString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string, *haxe__io__Encoding)] {
						hx_receiver_227 := hx_structural_source_221
						return hxrt.BoundFunction(func(value *string, encoding *haxe__io__Encoding) {
							hx_receiver_227.__hx_this.addString(value, encoding)
						}, func() any {
							return hx_receiver_227.__hx_this
						}, "addString")
					}(), Found: true}
				case "addInt32":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_228 := hx_structural_source_221
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_228.__hx_this.addInt32(value)
						}, func() any {
							return hx_receiver_228.__hx_this
						}, "addInt32")
					}(), Found: true}
				case "addInt64":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe___Int64_____Int64)] {
						hx_receiver_229 := hx_structural_source_221
						return hxrt.BoundFunction(func(value *haxe___Int64_____Int64) {
							hx_receiver_229.__hx_this.addInt64(value)
						}, func() any {
							return hx_receiver_229.__hx_this
						}, "addInt64")
					}(), Found: true}
				case "addFloat":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(float64)] {
						hx_receiver_230 := hx_structural_source_221
						return hxrt.BoundFunction(func(value float64) {
							hx_receiver_230.__hx_this.addFloat(value)
						}, func() any {
							return hx_receiver_230.__hx_this
						}, "addFloat")
					}(), Found: true}
				case "addDouble":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(float64)] {
						hx_receiver_231 := hx_structural_source_221
						return hxrt.BoundFunction(func(value float64) {
							hx_receiver_231.__hx_this.addDouble(value)
						}, func() any {
							return hx_receiver_231.__hx_this
						}, "addDouble")
					}(), Found: true}
				case "addBytes":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes, int, int)] {
						hx_receiver_232 := hx_structural_source_221
						return hxrt.BoundFunction(func(source *haxe__io__Bytes, pos int, len int) {
							hx_receiver_232.__hx_this.addBytes(source, pos, len)
						}, func() any {
							return hx_receiver_232.__hx_this
						}, "addBytes")
					}(), Found: true}
				case "getBytes":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *haxe__io__Bytes] {
						hx_receiver_233 := hx_structural_source_221
						return hxrt.BoundFunction(func() *haxe__io__Bytes {
							return hx_receiver_233.__hx_this.getBytes()
						}, func() any {
							return hx_receiver_233.__hx_this
						}, "getBytes")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "b":
					hx_structural_source_221.b = func(hx_value_222 any) []int {
						if hx_value_222 == nil {
							var hx_zero_223 []int
							return hx_zero_223
						}
						return hx_value_222.([]int)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe__io__Bytes:
		return func(hx_structural_source_234 *haxe__io__Bytes) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_234, func(name string) hxrt.AnonymousField {
				switch name {
				case "length":
					return hxrt.AnonymousField{Value: hx_structural_source_234.length, Found: true}
				case "b":
					return hxrt.AnonymousField{Value: hx_structural_source_234.b, Found: true}
				case "__hx_raw":
					return hxrt.AnonymousField{Value: hx_structural_source_234.__hx_raw, Found: true}
				case "__hx_rawValid":
					return hxrt.AnonymousField{Value: hx_structural_source_234.__hx_rawValid, Found: true}
				case "__hx_dataExposed":
					return hxrt.AnonymousField{Value: hx_structural_source_234.__hx_dataExposed, Found: true}
				case "get":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int) int] {
						hx_receiver_243 := hx_structural_source_234
						return hxrt.BoundFunction(func(pos int) int {
							return hx_receiver_243.__hx_this.get(pos)
						}, func() any {
							return hx_receiver_243.__hx_this
						}, "get")
					}(), Found: true}
				case "set":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, int)] {
						hx_receiver_244 := hx_structural_source_234
						return hxrt.BoundFunction(func(pos int, value int) {
							hx_receiver_244.__hx_this.set(pos, value)
						}, func() any {
							return hx_receiver_244.__hx_this
						}, "set")
					}(), Found: true}
				case "blit":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, *haxe__io__Bytes, int, int)] {
						hx_receiver_245 := hx_structural_source_234
						return hxrt.BoundFunction(func(pos int, src *haxe__io__Bytes, srcpos int, len int) {
							hx_receiver_245.__hx_this.blit(pos, src, srcpos, len)
						}, func() any {
							return hx_receiver_245.__hx_this
						}, "blit")
					}(), Found: true}
				case "fill":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, int, int)] {
						hx_receiver_246 := hx_structural_source_234
						return hxrt.BoundFunction(func(pos int, len int, value int) {
							hx_receiver_246.__hx_this.fill(pos, len, value)
						}, func() any {
							return hx_receiver_246.__hx_this
						}, "fill")
					}(), Found: true}
				case "sub":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, int) *haxe__io__Bytes] {
						hx_receiver_247 := hx_structural_source_234
						return hxrt.BoundFunction(func(pos int, len int) *haxe__io__Bytes {
							return hx_receiver_247.__hx_this.sub(pos, len)
						}, func() any {
							return hx_receiver_247.__hx_this
						}, "sub")
					}(), Found: true}
				case "compare":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes) int] {
						hx_receiver_248 := hx_structural_source_234
						return hxrt.BoundFunction(func(other *haxe__io__Bytes) int {
							return hx_receiver_248.__hx_this.compare(other)
						}, func() any {
							return hx_receiver_248.__hx_this
						}, "compare")
					}(), Found: true}
				case "getDouble":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int) float64] {
						hx_receiver_249 := hx_structural_source_234
						return hxrt.BoundFunction(func(pos int) float64 {
							return hx_receiver_249.__hx_this.getDouble(pos)
						}, func() any {
							return hx_receiver_249.__hx_this
						}, "getDouble")
					}(), Found: true}
				case "getFloat":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int) float64] {
						hx_receiver_250 := hx_structural_source_234
						return hxrt.BoundFunction(func(pos int) float64 {
							return hx_receiver_250.__hx_this.getFloat(pos)
						}, func() any {
							return hx_receiver_250.__hx_this
						}, "getFloat")
					}(), Found: true}
				case "setDouble":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, float64)] {
						hx_receiver_251 := hx_structural_source_234
						return hxrt.BoundFunction(func(pos int, value float64) {
							hx_receiver_251.__hx_this.setDouble(pos, value)
						}, func() any {
							return hx_receiver_251.__hx_this
						}, "setDouble")
					}(), Found: true}
				case "setFloat":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, float64)] {
						hx_receiver_252 := hx_structural_source_234
						return hxrt.BoundFunction(func(pos int, value float64) {
							hx_receiver_252.__hx_this.setFloat(pos, value)
						}, func() any {
							return hx_receiver_252.__hx_this
						}, "setFloat")
					}(), Found: true}
				case "getUInt16":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int) int] {
						hx_receiver_253 := hx_structural_source_234
						return hxrt.BoundFunction(func(pos int) int {
							return hx_receiver_253.__hx_this.getUInt16(pos)
						}, func() any {
							return hx_receiver_253.__hx_this
						}, "getUInt16")
					}(), Found: true}
				case "setUInt16":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, int)] {
						hx_receiver_254 := hx_structural_source_234
						return hxrt.BoundFunction(func(pos int, value int) {
							hx_receiver_254.__hx_this.setUInt16(pos, value)
						}, func() any {
							return hx_receiver_254.__hx_this
						}, "setUInt16")
					}(), Found: true}
				case "getInt32":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int) int] {
						hx_receiver_255 := hx_structural_source_234
						return hxrt.BoundFunction(func(pos int) int {
							return hx_receiver_255.__hx_this.getInt32(pos)
						}, func() any {
							return hx_receiver_255.__hx_this
						}, "getInt32")
					}(), Found: true}
				case "getInt64":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int) *haxe___Int64_____Int64] {
						hx_receiver_256 := hx_structural_source_234
						return hxrt.BoundFunction(func(pos int) *haxe___Int64_____Int64 {
							return hx_receiver_256.__hx_this.getInt64(pos)
						}, func() any {
							return hx_receiver_256.__hx_this
						}, "getInt64")
					}(), Found: true}
				case "setInt32":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, int)] {
						hx_receiver_257 := hx_structural_source_234
						return hxrt.BoundFunction(func(pos int, value int) {
							hx_receiver_257.__hx_this.setInt32(pos, value)
						}, func() any {
							return hx_receiver_257.__hx_this
						}, "setInt32")
					}(), Found: true}
				case "setInt64":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, *haxe___Int64_____Int64)] {
						hx_receiver_258 := hx_structural_source_234
						return hxrt.BoundFunction(func(pos int, value *haxe___Int64_____Int64) {
							hx_receiver_258.__hx_this.setInt64(pos, value)
						}, func() any {
							return hx_receiver_258.__hx_this
						}, "setInt64")
					}(), Found: true}
				case "getString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, int, *haxe__io__Encoding) *string] {
						hx_receiver_259 := hx_structural_source_234
						return hxrt.BoundFunction(func(pos int, len int, encoding *haxe__io__Encoding) *string {
							return hx_receiver_259.__hx_this.getString(pos, len, encoding)
						}, func() any {
							return hx_receiver_259.__hx_this
						}, "getString")
					}(), Found: true}
				case "readString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int, int) *string] {
						hx_receiver_260 := hx_structural_source_234
						return hxrt.BoundFunction(func(pos int, len int) *string {
							return hx_receiver_260.__hx_this.readString(pos, len)
						}, func() any {
							return hx_receiver_260.__hx_this
						}, "readString")
					}(), Found: true}
				case "toString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_261 := hx_structural_source_234
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_261.__hx_this.toString()
						}, func() any {
							return hx_receiver_261.__hx_this
						}, "toString")
					}(), Found: true}
				case "toHex":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_262 := hx_structural_source_234
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_262.__hx_this.toHex()
						}, func() any {
							return hx_receiver_262.__hx_this
						}, "toHex")
					}(), Found: true}
				case "getData":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() []int] {
						hx_receiver_263 := hx_structural_source_234
						return hxrt.BoundFunction(func() []int {
							return hx_receiver_263.__hx_this.getData()
						}, func() any {
							return hx_receiver_263.__hx_this
						}, "getData")
					}(), Found: true}
				case "__hx_nativeView":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *hxrt.ByteView] {
						hx_receiver_264 := hx_structural_source_234
						return hxrt.BoundFunction(func() *hxrt.ByteView {
							return hx_receiver_264.__hx_this.__hx_nativeView()
						}, func() any {
							return hx_receiver_264.__hx_this
						}, "__hx_nativeView")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "b":
					hx_structural_source_234.b = func(hx_value_235 any) []int {
						if hx_value_235 == nil {
							var hx_zero_236 []int
							return hx_zero_236
						}
						return hx_value_235.([]int)
					}(incoming)
					return true
				case "__hx_raw":
					hx_structural_source_234.__hx_raw = func(hx_value_237 any) *hxrt.ByteView {
						if hx_value_237 == nil {
							var hx_zero_238 *hxrt.ByteView
							return hx_zero_238
						}
						return hx_value_237.(*hxrt.ByteView)
					}(incoming)
					return true
				case "__hx_rawValid":
					hx_structural_source_234.__hx_rawValid = func(hx_value_239 any) bool {
						if hx_value_239 == nil {
							var hx_zero_240 bool
							return hx_zero_240
						}
						return hx_value_239.(bool)
					}(incoming)
					return true
				case "__hx_dataExposed":
					hx_structural_source_234.__hx_dataExposed = func(hx_value_241 any) bool {
						if hx_value_241 == nil {
							var hx_zero_242 bool
							return hx_zero_242
						}
						return hx_value_241.(bool)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe__http__HttpBase:
		return func(hx_structural_source_265 *haxe__http__HttpBase) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_265, func(name string) hxrt.AnonymousField {
				switch name {
				case "url":
					return hxrt.AnonymousField{Value: hx_structural_source_265.url, Found: true}
				case "responseBytes":
					return hxrt.AnonymousField{Value: hx_structural_source_265.responseBytes, Found: true}
				case "responseAsString":
					return hxrt.AnonymousField{Value: hx_structural_source_265.responseAsString, Found: true}
				case "postData":
					return hxrt.AnonymousField{Value: hx_structural_source_265.postData, Found: true}
				case "postBytes":
					return hxrt.AnonymousField{Value: hx_structural_source_265.postBytes, Found: true}
				case "headers":
					return hxrt.AnonymousField{Value: hx_structural_source_265.headers, Found: true}
				case "params":
					return hxrt.AnonymousField{Value: hx_structural_source_265.params, Found: true}
				case "emptyOnData":
					return hxrt.AnonymousField{Value: hx_structural_source_265.emptyOnData, Found: true}
				case "setHeader":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string, *string)] {
						hx_receiver_276 := hx_structural_source_265
						return hxrt.BoundFunction(func(name *string, value *string) {
							hx_receiver_276.__hx_this.setHeader(name, value)
						}, func() any {
							return hx_receiver_276.__hx_this
						}, "setHeader")
					}(), Found: true}
				case "addHeader":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string, *string)] {
						hx_receiver_277 := hx_structural_source_265
						return hxrt.BoundFunction(func(header *string, value *string) {
							hx_receiver_277.__hx_this.addHeader(header, value)
						}, func() any {
							return hx_receiver_277.__hx_this
						}, "addHeader")
					}(), Found: true}
				case "setParameter":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string, *string)] {
						hx_receiver_278 := hx_structural_source_265
						return hxrt.BoundFunction(func(name *string, value *string) {
							hx_receiver_278.__hx_this.setParameter(name, value)
						}, func() any {
							return hx_receiver_278.__hx_this
						}, "setParameter")
					}(), Found: true}
				case "addParameter":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string, *string)] {
						hx_receiver_279 := hx_structural_source_265
						return hxrt.BoundFunction(func(name *string, value *string) {
							hx_receiver_279.__hx_this.addParameter(name, value)
						}, func() any {
							return hx_receiver_279.__hx_this
						}, "addParameter")
					}(), Found: true}
				case "setPostData":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string)] {
						hx_receiver_280 := hx_structural_source_265
						return hxrt.BoundFunction(func(data *string) {
							hx_receiver_280.__hx_this.setPostData(data)
						}, func() any {
							return hx_receiver_280.__hx_this
						}, "setPostData")
					}(), Found: true}
				case "setPostBytes":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes)] {
						hx_receiver_281 := hx_structural_source_265
						return hxrt.BoundFunction(func(data *haxe__io__Bytes) {
							hx_receiver_281.__hx_this.setPostBytes(data)
						}, func() any {
							return hx_receiver_281.__hx_this
						}, "setPostBytes")
					}(), Found: true}
				case "request":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(any)] {
						hx_receiver_282 := hx_structural_source_265
						return hxrt.BoundFunction(func(post any) {
							hx_receiver_282.__hx_this.request(post)
						}, func() any {
							return hx_receiver_282.__hx_this
						}, "request")
					}(), Found: true}
				case "onData":
					return hxrt.AnonymousField{Value: hxrt.ReadDynamicFunction(hx_structural_source_265.onData), Found: true}
				case "onBytes":
					return hxrt.AnonymousField{Value: hxrt.ReadDynamicFunction(hx_structural_source_265.onBytes), Found: true}
				case "onError":
					return hxrt.AnonymousField{Value: hxrt.ReadDynamicFunction(hx_structural_source_265.onError), Found: true}
				case "onStatus":
					return hxrt.AnonymousField{Value: hxrt.ReadDynamicFunction(hx_structural_source_265.onStatus), Found: true}
				case "hasOnData":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() bool] {
						hx_receiver_283 := hx_structural_source_265
						return hxrt.BoundFunction(func() bool {
							return hx_receiver_283.__hx_this.hasOnData()
						}, func() any {
							return hx_receiver_283.__hx_this
						}, "hasOnData")
					}(), Found: true}
				case "success":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes)] {
						hx_receiver_284 := hx_structural_source_265
						return hxrt.BoundFunction(func(data *haxe__io__Bytes) {
							hx_receiver_284.__hx_this.success(data)
						}, func() any {
							return hx_receiver_284.__hx_this
						}, "success")
					}(), Found: true}
				case "get_responseData":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_285 := hx_structural_source_265
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_285.__hx_this.get_responseData()
						}, func() any {
							return hx_receiver_285.__hx_this
						}, "get_responseData")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "url":
					hx_structural_source_265.url = func(hx_value_266 any) *string {
						if hx_value_266 == nil {
							var hx_zero_267 *string
							return hx_zero_267
						}
						return hx_value_266.(*string)
					}(incoming)
					return true
				case "responseAsString":
					hx_structural_source_265.responseAsString = func(hx_value_268 any) *string {
						if hx_value_268 == nil {
							var hx_zero_269 *string
							return hx_zero_269
						}
						return hx_value_268.(*string)
					}(incoming)
					return true
				case "postData":
					hx_structural_source_265.postData = func(hx_value_270 any) *string {
						if hx_value_270 == nil {
							var hx_zero_271 *string
							return hx_zero_271
						}
						return hx_value_270.(*string)
					}(incoming)
					return true
				case "postBytes":
					hx_structural_source_265.postBytes = hx_recover_class_haxe__io__Bytes(incoming)
					return true
				case "headers":
					hx_structural_source_265.headers = func(hx_value_272 any) *hxrt.Array {
						if hx_value_272 == nil {
							var hx_zero_273 *hxrt.Array
							return hx_zero_273
						}
						return hx_value_272.(*hxrt.Array)
					}(incoming)
					return true
				case "params":
					hx_structural_source_265.params = func(hx_value_274 any) *hxrt.Array {
						if hx_value_274 == nil {
							var hx_zero_275 *hxrt.Array
							return hx_zero_275
						}
						return hx_value_274.(*hxrt.Array)
					}(incoming)
					return true
				case "onData":
					hx_structural_source_265.onData = hx_recover_callable_9da3ad96(incoming)
					return true
				case "onBytes":
					hx_structural_source_265.onBytes = hx_recover_callable_c5f1b802(incoming)
					return true
				case "onError":
					hx_structural_source_265.onError = hx_recover_callable_9da3ad96(incoming)
					return true
				case "onStatus":
					hx_structural_source_265.onStatus = hx_recover_callable_d8be58e9(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe__exceptions__NotImplementedException:
		return func(hx_structural_source_286 *haxe__exceptions__NotImplementedException) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_286, func(name string) hxrt.AnonymousField {
				switch name {
				case "posInfos":
					return hxrt.AnonymousField{Value: hx_structural_source_286.posInfos, Found: true}
				case "toString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_287 := hx_structural_source_286
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_287.__hx_this.toString()
						}, func() any {
							return hx_receiver_287.__hx_this
						}, "toString")
					}(), Found: true}
				case "savedMessage":
					return hxrt.AnonymousField{Value: hx_structural_source_286.savedMessage, Found: true}
				case "savedStack":
					return hxrt.AnonymousField{Value: hx_structural_source_286.savedStack, Found: true}
				case "savedPrevious":
					return hxrt.AnonymousField{Value: hx_structural_source_286.savedPrevious, Found: true}
				case "savedNative":
					return hxrt.AnonymousField{Value: hx_structural_source_286.savedNative, Found: true}
				case "get_message":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_290 := hx_structural_source_286
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_290.__hx_this.get_message()
						}, func() any {
							return hx_receiver_290.__hx_this
						}, "get_message")
					}(), Found: true}
				case "get_stack":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_291 := hx_structural_source_286
						return hxrt.BoundFunction(func() any {
							return hx_receiver_291.__hx_this.get_stack()
						}, func() any {
							return hx_receiver_291.__hx_this
						}, "get_stack")
					}(), Found: true}
				case "get_previous":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *haxe__Exception] {
						hx_receiver_292 := hx_structural_source_286
						return hxrt.BoundFunction(func() *haxe__Exception {
							return hx_receiver_292.__hx_this.get_previous()
						}, func() any {
							return hx_receiver_292.__hx_this
						}, "get_previous")
					}(), Found: true}
				case "get_native":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_293 := hx_structural_source_286
						return hxrt.BoundFunction(func() any {
							return hx_receiver_293.__hx_this.get_native()
						}, func() any {
							return hx_receiver_293.__hx_this
						}, "get_native")
					}(), Found: true}
				case "unwrap":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_294 := hx_structural_source_286
						return hxrt.BoundFunction(func() any {
							return hx_receiver_294.__hx_this.unwrap()
						}, func() any {
							return hx_receiver_294.__hx_this
						}, "unwrap")
					}(), Found: true}
				case "details":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_295 := hx_structural_source_286
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_295.__hx_this.details()
						}, func() any {
							return hx_receiver_295.__hx_this
						}, "details")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "savedMessage":
					hx_structural_source_286.savedMessage = func(hx_value_288 any) *string {
						if hx_value_288 == nil {
							var hx_zero_289 *string
							return hx_zero_289
						}
						return hx_value_288.(*string)
					}(incoming)
					return true
				case "savedStack":
					hx_structural_source_286.savedStack = incoming
					return true
				case "savedPrevious":
					hx_structural_source_286.savedPrevious = hx_recover_class_haxe__Exception(incoming)
					return true
				case "savedNative":
					hx_structural_source_286.savedNative = incoming
					return true
				}
				return false
			})
		}(source)
	case *haxe__ds__StringMap:
		return func(hx_structural_source_296 *haxe__ds__StringMap) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_296, func(name string) hxrt.AnonymousField {
				switch name {
				case "h":
					return hxrt.AnonymousField{Value: hx_structural_source_296.h, Found: true}
				case "set":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string, any)] {
						hx_receiver_299 := hx_structural_source_296
						return hxrt.BoundFunction(func(key *string, value any) {
							hx_receiver_299.__hx_this.set(key, value)
						}, func() any {
							return hx_receiver_299.__hx_this
						}, "set")
					}(), Found: true}
				case "get":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string) any] {
						hx_receiver_300 := hx_structural_source_296
						return hxrt.BoundFunction(func(key *string) any {
							return hx_receiver_300.__hx_this.get(key)
						}, func() any {
							return hx_receiver_300.__hx_this
						}, "get")
					}(), Found: true}
				case "exists":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string) bool] {
						hx_receiver_301 := hx_structural_source_296
						return hxrt.BoundFunction(func(key *string) bool {
							return hx_receiver_301.__hx_this.exists(key)
						}, func() any {
							return hx_receiver_301.__hx_this
						}, "exists")
					}(), Found: true}
				case "remove":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string) bool] {
						hx_receiver_302 := hx_structural_source_296
						return hxrt.BoundFunction(func(key *string) bool {
							return hx_receiver_302.__hx_this.remove(key)
						}, func() any {
							return hx_receiver_302.__hx_this
						}, "remove")
					}(), Found: true}
				case "keys":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() hxrt.Anonymous] {
						hx_receiver_303 := hx_structural_source_296
						return hxrt.BoundFunction(func() hxrt.Anonymous {
							return hx_receiver_303.__hx_this.keys()
						}, func() any {
							return hx_receiver_303.__hx_this
						}, "keys")
					}(), Found: true}
				case "iterator":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() hxrt.Anonymous] {
						hx_receiver_304 := hx_structural_source_296
						return hxrt.BoundFunction(func() hxrt.Anonymous {
							return hx_receiver_304.__hx_this.iterator()
						}, func() any {
							return hx_receiver_304.__hx_this
						}, "iterator")
					}(), Found: true}
				case "keyValueIterator":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() hxrt.Anonymous] {
						hx_receiver_305 := hx_structural_source_296
						return hxrt.BoundFunction(func() hxrt.Anonymous {
							return hx_receiver_305.__hx_this.keyValueIterator()
						}, func() any {
							return hx_receiver_305.__hx_this
						}, "keyValueIterator")
					}(), Found: true}
				case "getIMap":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(any) any] {
						hx_receiver_306 := hx_structural_source_296
						return hxrt.BoundFunction(func(key any) any {
							return hx_receiver_306.__hx_this.getIMap(key)
						}, func() any {
							return hx_receiver_306.__hx_this
						}, "getIMap")
					}(), Found: true}
				case "setIMap":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(any, any)] {
						hx_receiver_307 := hx_structural_source_296
						return hxrt.BoundFunction(func(key any, value any) {
							hx_receiver_307.__hx_this.setIMap(key, value)
						}, func() any {
							return hx_receiver_307.__hx_this
						}, "setIMap")
					}(), Found: true}
				case "existsIMap":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(any) bool] {
						hx_receiver_308 := hx_structural_source_296
						return hxrt.BoundFunction(func(key any) bool {
							return hx_receiver_308.__hx_this.existsIMap(key)
						}, func() any {
							return hx_receiver_308.__hx_this
						}, "existsIMap")
					}(), Found: true}
				case "removeIMap":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(any) bool] {
						hx_receiver_309 := hx_structural_source_296
						return hxrt.BoundFunction(func(key any) bool {
							return hx_receiver_309.__hx_this.removeIMap(key)
						}, func() any {
							return hx_receiver_309.__hx_this
						}, "removeIMap")
					}(), Found: true}
				case "copyIMap":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() haxe__IMap] {
						hx_receiver_310 := hx_structural_source_296
						return hxrt.BoundFunction(func() haxe__IMap {
							return hx_receiver_310.__hx_this.copyIMap()
						}, func() any {
							return hx_receiver_310.__hx_this
						}, "copyIMap")
					}(), Found: true}
				case "copy":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *haxe__ds__StringMap] {
						hx_receiver_311 := hx_structural_source_296
						return hxrt.BoundFunction(func() *haxe__ds__StringMap {
							return hx_receiver_311.__hx_this.copy()
						}, func() any {
							return hx_receiver_311.__hx_this
						}, "copy")
					}(), Found: true}
				case "toString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_312 := hx_structural_source_296
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_312.__hx_this.toString()
						}, func() any {
							return hx_receiver_312.__hx_this
						}, "toString")
					}(), Found: true}
				case "clear":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func()] {
						hx_receiver_313 := hx_structural_source_296
						return hxrt.BoundFunction(func() {
							hx_receiver_313.__hx_this.clear()
						}, func() any {
							return hx_receiver_313.__hx_this
						}, "clear")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "h":
					hx_structural_source_296.h = func(hx_value_297 any) *hxrt.StringMapCell {
						if hx_value_297 == nil {
							var hx_zero_298 *hxrt.StringMapCell
							return hx_zero_298
						}
						return hx_value_297.(*hxrt.StringMapCell)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe__Exception:
		return func(hx_structural_source_314 *haxe__Exception) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_314, func(name string) hxrt.AnonymousField {
				switch name {
				case "savedMessage":
					return hxrt.AnonymousField{Value: hx_structural_source_314.savedMessage, Found: true}
				case "savedStack":
					return hxrt.AnonymousField{Value: hx_structural_source_314.savedStack, Found: true}
				case "savedPrevious":
					return hxrt.AnonymousField{Value: hx_structural_source_314.savedPrevious, Found: true}
				case "savedNative":
					return hxrt.AnonymousField{Value: hx_structural_source_314.savedNative, Found: true}
				case "get_message":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_317 := hx_structural_source_314
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_317.__hx_this.get_message()
						}, func() any {
							return hx_receiver_317.__hx_this
						}, "get_message")
					}(), Found: true}
				case "get_stack":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_318 := hx_structural_source_314
						return hxrt.BoundFunction(func() any {
							return hx_receiver_318.__hx_this.get_stack()
						}, func() any {
							return hx_receiver_318.__hx_this
						}, "get_stack")
					}(), Found: true}
				case "get_previous":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *haxe__Exception] {
						hx_receiver_319 := hx_structural_source_314
						return hxrt.BoundFunction(func() *haxe__Exception {
							return hx_receiver_319.__hx_this.get_previous()
						}, func() any {
							return hx_receiver_319.__hx_this
						}, "get_previous")
					}(), Found: true}
				case "get_native":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_320 := hx_structural_source_314
						return hxrt.BoundFunction(func() any {
							return hx_receiver_320.__hx_this.get_native()
						}, func() any {
							return hx_receiver_320.__hx_this
						}, "get_native")
					}(), Found: true}
				case "unwrap":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_321 := hx_structural_source_314
						return hxrt.BoundFunction(func() any {
							return hx_receiver_321.__hx_this.unwrap()
						}, func() any {
							return hx_receiver_321.__hx_this
						}, "unwrap")
					}(), Found: true}
				case "toString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_322 := hx_structural_source_314
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_322.__hx_this.toString()
						}, func() any {
							return hx_receiver_322.__hx_this
						}, "toString")
					}(), Found: true}
				case "details":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_323 := hx_structural_source_314
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_323.__hx_this.details()
						}, func() any {
							return hx_receiver_323.__hx_this
						}, "details")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "savedMessage":
					hx_structural_source_314.savedMessage = func(hx_value_315 any) *string {
						if hx_value_315 == nil {
							var hx_zero_316 *string
							return hx_zero_316
						}
						return hx_value_315.(*string)
					}(incoming)
					return true
				case "savedStack":
					hx_structural_source_314.savedStack = incoming
					return true
				case "savedPrevious":
					hx_structural_source_314.savedPrevious = hx_recover_class_haxe__Exception(incoming)
					return true
				case "savedNative":
					hx_structural_source_314.savedNative = incoming
					return true
				}
				return false
			})
		}(source)
	case *sys__thread__Tls:
		return func(hx_structural_source_324 *sys__thread__Tls) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_324, func(name string) hxrt.AnonymousField {
				switch name {
				case "__handle":
					return hxrt.AnonymousField{Value: hx_structural_source_324.__handle, Found: true}
				case "get_value":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_325 := hx_structural_source_324
						return hxrt.BoundFunction(func() any {
							return hx_receiver_325.__hx_this.get_value()
						}, func() any {
							return hx_receiver_325.__hx_this
						}, "get_value")
					}(), Found: true}
				case "set_value":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(any) any] {
						hx_receiver_326 := hx_structural_source_324
						return hxrt.BoundFunction(func(v any) any {
							return hx_receiver_326.__hx_this.set_value(v)
						}, func() any {
							return hx_receiver_326.__hx_this
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
		return func(hx_structural_source_327 *haxe__iterators__StringKeyValueIterator) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_327, func(name string) hxrt.AnonymousField {
				switch name {
				case "offset":
					return hxrt.AnonymousField{Value: hx_structural_source_327.offset, Found: true}
				case "s":
					return hxrt.AnonymousField{Value: hx_structural_source_327.s, Found: true}
				case "hasNext":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() bool] {
						hx_receiver_332 := hx_structural_source_327
						return hxrt.BoundFunction(func() bool {
							return hx_receiver_332.__hx_this.hasNext()
						}, func() any {
							return hx_receiver_332.__hx_this
						}, "hasNext")
					}(), Found: true}
				case "next":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() hxrt.Anonymous] {
						hx_receiver_333 := hx_structural_source_327
						return hxrt.BoundFunction(func() hxrt.Anonymous {
							return hx_receiver_333.__hx_this.next()
						}, func() any {
							return hx_receiver_333.__hx_this
						}, "next")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "offset":
					hx_structural_source_327.offset = func(hx_value_328 any) int {
						if hx_value_328 == nil {
							var hx_zero_329 int
							return hx_zero_329
						}
						return hx_value_328.(int)
					}(incoming)
					return true
				case "s":
					hx_structural_source_327.s = func(hx_value_330 any) *string {
						if hx_value_330 == nil {
							var hx_zero_331 *string
							return hx_zero_331
						}
						return hx_value_330.(*string)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe__iterators__StringIterator:
		return func(hx_structural_source_334 *haxe__iterators__StringIterator) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_334, func(name string) hxrt.AnonymousField {
				switch name {
				case "offset":
					return hxrt.AnonymousField{Value: hx_structural_source_334.offset, Found: true}
				case "s":
					return hxrt.AnonymousField{Value: hx_structural_source_334.s, Found: true}
				case "hasNext":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() bool] {
						hx_receiver_339 := hx_structural_source_334
						return hxrt.BoundFunction(func() bool {
							return hx_receiver_339.__hx_this.hasNext()
						}, func() any {
							return hx_receiver_339.__hx_this
						}, "hasNext")
					}(), Found: true}
				case "next":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_340 := hx_structural_source_334
						return hxrt.BoundFunction(func() int {
							return hx_receiver_340.__hx_this.next()
						}, func() any {
							return hx_receiver_340.__hx_this
						}, "next")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "offset":
					hx_structural_source_334.offset = func(hx_value_335 any) int {
						if hx_value_335 == nil {
							var hx_zero_336 int
							return hx_zero_336
						}
						return hx_value_335.(int)
					}(incoming)
					return true
				case "s":
					hx_structural_source_334.s = func(hx_value_337 any) *string {
						if hx_value_337 == nil {
							var hx_zero_338 *string
							return hx_zero_338
						}
						return hx_value_337.(*string)
					}(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe__iterators__MapKeyValueIterator:
		return func(hx_structural_source_341 *haxe__iterators__MapKeyValueIterator) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_341, func(name string) hxrt.AnonymousField {
				switch name {
				case "map":
					return hxrt.AnonymousField{Value: hx_structural_source_341.map_, Found: true}
				case "keys":
					return hxrt.AnonymousField{Value: hx_structural_source_341.keys, Found: true}
				case "hasNext":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() bool] {
						hx_receiver_344 := hx_structural_source_341
						return hxrt.BoundFunction(func() bool {
							return hx_receiver_344.__hx_this.hasNext()
						}, func() any {
							return hx_receiver_344.__hx_this
						}, "hasNext")
					}(), Found: true}
				case "next":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() hxrt.Anonymous] {
						hx_receiver_345 := hx_structural_source_341
						return hxrt.BoundFunction(func() hxrt.Anonymous {
							return hx_receiver_345.__hx_this.next()
						}, func() any {
							return hx_receiver_345.__hx_this
						}, "next")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "map":
					hx_structural_source_341.map_ = func(hx_value_342 any) haxe__IMap {
						if hx_value_342 == nil {
							var hx_zero_343 haxe__IMap
							return hx_zero_343
						}
						return hx_value_342.(haxe__IMap)
					}(hxrt.ObjectSource(incoming))
					return true
				case "keys":
					hx_structural_source_341.keys = hx_structural_recover(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe__io__BytesOutput:
		return func(hx_structural_source_346 *haxe__io__BytesOutput) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_346, func(name string) hxrt.AnonymousField {
				switch name {
				case "b":
					return hxrt.AnonymousField{Value: hx_structural_source_346.b, Found: true}
				case "get_length":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() int] {
						hx_receiver_347 := hx_structural_source_346
						return hxrt.BoundFunction(func() int {
							return hx_receiver_347.__hx_this.get_length()
						}, func() any {
							return hx_receiver_347.__hx_this
						}, "get_length")
					}(), Found: true}
				case "writeByte":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_348 := hx_structural_source_346
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_348.__hx_this.writeByte(value)
						}, func() any {
							return hx_receiver_348.__hx_this
						}, "writeByte")
					}(), Found: true}
				case "writeBytes":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes, int, int) int] {
						hx_receiver_349 := hx_structural_source_346
						return hxrt.BoundFunction(func(bytes *haxe__io__Bytes, pos int, len int) int {
							return hx_receiver_349.__hx_this.writeBytes(bytes, pos, len)
						}, func() any {
							return hx_receiver_349.__hx_this
						}, "writeBytes")
					}(), Found: true}
				case "getBytes":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *haxe__io__Bytes] {
						hx_receiver_350 := hx_structural_source_346
						return hxrt.BoundFunction(func() *haxe__io__Bytes {
							return hx_receiver_350.__hx_this.getBytes()
						}, func() any {
							return hx_receiver_350.__hx_this
						}, "getBytes")
					}(), Found: true}
				case "bigEndian":
					return hxrt.AnonymousField{Value: hx_structural_source_346.bigEndian, Found: true}
				case "flush":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func()] {
						hx_receiver_351 := hx_structural_source_346
						return hxrt.BoundFunction(func() {
							hx_receiver_351.__hx_this.flush()
						}, func() any {
							return hx_receiver_351.__hx_this
						}, "flush")
					}(), Found: true}
				case "close":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func()] {
						hx_receiver_352 := hx_structural_source_346
						return hxrt.BoundFunction(func() {
							hx_receiver_352.__hx_this.close()
						}, func() any {
							return hx_receiver_352.__hx_this
						}, "close")
					}(), Found: true}
				case "set_bigEndian":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(bool) bool] {
						hx_receiver_353 := hx_structural_source_346
						return hxrt.BoundFunction(func(value bool) bool {
							return hx_receiver_353.__hx_this.set_bigEndian(value)
						}, func() any {
							return hx_receiver_353.__hx_this
						}, "set_bigEndian")
					}(), Found: true}
				case "write":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes)] {
						hx_receiver_354 := hx_structural_source_346
						return hxrt.BoundFunction(func(bytes *haxe__io__Bytes) {
							hx_receiver_354.__hx_this.write(bytes)
						}, func() any {
							return hx_receiver_354.__hx_this
						}, "write")
					}(), Found: true}
				case "writeFullBytes":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Bytes, int, int)] {
						hx_receiver_355 := hx_structural_source_346
						return hxrt.BoundFunction(func(bytes *haxe__io__Bytes, pos int, len int) {
							hx_receiver_355.__hx_this.writeFullBytes(bytes, pos, len)
						}, func() any {
							return hx_receiver_355.__hx_this
						}, "writeFullBytes")
					}(), Found: true}
				case "writeFloat":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(float64)] {
						hx_receiver_356 := hx_structural_source_346
						return hxrt.BoundFunction(func(value float64) {
							hx_receiver_356.__hx_this.writeFloat(value)
						}, func() any {
							return hx_receiver_356.__hx_this
						}, "writeFloat")
					}(), Found: true}
				case "writeDouble":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(float64)] {
						hx_receiver_357 := hx_structural_source_346
						return hxrt.BoundFunction(func(value float64) {
							hx_receiver_357.__hx_this.writeDouble(value)
						}, func() any {
							return hx_receiver_357.__hx_this
						}, "writeDouble")
					}(), Found: true}
				case "writeInt8":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_358 := hx_structural_source_346
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_358.__hx_this.writeInt8(value)
						}, func() any {
							return hx_receiver_358.__hx_this
						}, "writeInt8")
					}(), Found: true}
				case "writeInt16":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_359 := hx_structural_source_346
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_359.__hx_this.writeInt16(value)
						}, func() any {
							return hx_receiver_359.__hx_this
						}, "writeInt16")
					}(), Found: true}
				case "writeUInt16":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_360 := hx_structural_source_346
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_360.__hx_this.writeUInt16(value)
						}, func() any {
							return hx_receiver_360.__hx_this
						}, "writeUInt16")
					}(), Found: true}
				case "writeInt24":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_361 := hx_structural_source_346
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_361.__hx_this.writeInt24(value)
						}, func() any {
							return hx_receiver_361.__hx_this
						}, "writeInt24")
					}(), Found: true}
				case "writeUInt24":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_362 := hx_structural_source_346
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_362.__hx_this.writeUInt24(value)
						}, func() any {
							return hx_receiver_362.__hx_this
						}, "writeUInt24")
					}(), Found: true}
				case "writeInt32":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_363 := hx_structural_source_346
						return hxrt.BoundFunction(func(value int) {
							hx_receiver_363.__hx_this.writeInt32(value)
						}, func() any {
							return hx_receiver_363.__hx_this
						}, "writeInt32")
					}(), Found: true}
				case "prepare":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(int)] {
						hx_receiver_364 := hx_structural_source_346
						return hxrt.BoundFunction(func(nbytes int) {
							hx_receiver_364.__hx_this.prepare(nbytes)
						}, func() any {
							return hx_receiver_364.__hx_this
						}, "prepare")
					}(), Found: true}
				case "writeInput":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*haxe__io__Input, any)] {
						hx_receiver_365 := hx_structural_source_346
						return hxrt.BoundFunction(func(input *haxe__io__Input, bufsize any) {
							hx_receiver_365.__hx_this.writeInput(input, bufsize)
						}, func() any {
							return hx_receiver_365.__hx_this
						}, "writeInput")
					}(), Found: true}
				case "writeString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func(*string, *haxe__io__Encoding)] {
						hx_receiver_366 := hx_structural_source_346
						return hxrt.BoundFunction(func(value *string, encoding *haxe__io__Encoding) {
							hx_receiver_366.__hx_this.writeString(value, encoding)
						}, func() any {
							return hx_receiver_366.__hx_this
						}, "writeString")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "b":
					hx_structural_source_346.b = hx_recover_class_haxe__io__BytesBuffer(incoming)
					return true
				}
				return false
			})
		}(source)
	case *haxe__ValueException:
		return func(hx_structural_source_367 *haxe__ValueException) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_367, func(name string) hxrt.AnonymousField {
				switch name {
				case "value":
					return hxrt.AnonymousField{Value: hx_structural_source_367.value, Found: true}
				case "unwrap":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_368 := hx_structural_source_367
						return hxrt.BoundFunction(func() any {
							return hx_receiver_368.__hx_this.unwrap()
						}, func() any {
							return hx_receiver_368.__hx_this
						}, "unwrap")
					}(), Found: true}
				case "savedMessage":
					return hxrt.AnonymousField{Value: hx_structural_source_367.savedMessage, Found: true}
				case "savedStack":
					return hxrt.AnonymousField{Value: hx_structural_source_367.savedStack, Found: true}
				case "savedPrevious":
					return hxrt.AnonymousField{Value: hx_structural_source_367.savedPrevious, Found: true}
				case "savedNative":
					return hxrt.AnonymousField{Value: hx_structural_source_367.savedNative, Found: true}
				case "get_message":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_371 := hx_structural_source_367
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_371.__hx_this.get_message()
						}, func() any {
							return hx_receiver_371.__hx_this
						}, "get_message")
					}(), Found: true}
				case "get_stack":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_372 := hx_structural_source_367
						return hxrt.BoundFunction(func() any {
							return hx_receiver_372.__hx_this.get_stack()
						}, func() any {
							return hx_receiver_372.__hx_this
						}, "get_stack")
					}(), Found: true}
				case "get_previous":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *haxe__Exception] {
						hx_receiver_373 := hx_structural_source_367
						return hxrt.BoundFunction(func() *haxe__Exception {
							return hx_receiver_373.__hx_this.get_previous()
						}, func() any {
							return hx_receiver_373.__hx_this
						}, "get_previous")
					}(), Found: true}
				case "get_native":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_374 := hx_structural_source_367
						return hxrt.BoundFunction(func() any {
							return hx_receiver_374.__hx_this.get_native()
						}, func() any {
							return hx_receiver_374.__hx_this
						}, "get_native")
					}(), Found: true}
				case "toString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_375 := hx_structural_source_367
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_375.__hx_this.toString()
						}, func() any {
							return hx_receiver_375.__hx_this
						}, "toString")
					}(), Found: true}
				case "details":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_376 := hx_structural_source_367
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_376.__hx_this.details()
						}, func() any {
							return hx_receiver_376.__hx_this
						}, "details")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "savedMessage":
					hx_structural_source_367.savedMessage = func(hx_value_369 any) *string {
						if hx_value_369 == nil {
							var hx_zero_370 *string
							return hx_zero_370
						}
						return hx_value_369.(*string)
					}(incoming)
					return true
				case "savedStack":
					hx_structural_source_367.savedStack = incoming
					return true
				case "savedPrevious":
					hx_structural_source_367.savedPrevious = hx_recover_class_haxe__Exception(incoming)
					return true
				case "savedNative":
					hx_structural_source_367.savedNative = incoming
					return true
				}
				return false
			})
		}(source)
	case *haxe__exceptions__PosException:
		return func(hx_structural_source_377 *haxe__exceptions__PosException) hxrt.Anonymous {
			return hxrt.NewAnonymousView(hx_structural_source_377, func(name string) hxrt.AnonymousField {
				switch name {
				case "posInfos":
					return hxrt.AnonymousField{Value: hx_structural_source_377.posInfos, Found: true}
				case "toString":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_378 := hx_structural_source_377
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_378.__hx_this.toString()
						}, func() any {
							return hx_receiver_378.__hx_this
						}, "toString")
					}(), Found: true}
				case "savedMessage":
					return hxrt.AnonymousField{Value: hx_structural_source_377.savedMessage, Found: true}
				case "savedStack":
					return hxrt.AnonymousField{Value: hx_structural_source_377.savedStack, Found: true}
				case "savedPrevious":
					return hxrt.AnonymousField{Value: hx_structural_source_377.savedPrevious, Found: true}
				case "savedNative":
					return hxrt.AnonymousField{Value: hx_structural_source_377.savedNative, Found: true}
				case "get_message":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_381 := hx_structural_source_377
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_381.__hx_this.get_message()
						}, func() any {
							return hx_receiver_381.__hx_this
						}, "get_message")
					}(), Found: true}
				case "get_stack":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_382 := hx_structural_source_377
						return hxrt.BoundFunction(func() any {
							return hx_receiver_382.__hx_this.get_stack()
						}, func() any {
							return hx_receiver_382.__hx_this
						}, "get_stack")
					}(), Found: true}
				case "get_previous":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *haxe__Exception] {
						hx_receiver_383 := hx_structural_source_377
						return hxrt.BoundFunction(func() *haxe__Exception {
							return hx_receiver_383.__hx_this.get_previous()
						}, func() any {
							return hx_receiver_383.__hx_this
						}, "get_previous")
					}(), Found: true}
				case "get_native":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_384 := hx_structural_source_377
						return hxrt.BoundFunction(func() any {
							return hx_receiver_384.__hx_this.get_native()
						}, func() any {
							return hx_receiver_384.__hx_this
						}, "get_native")
					}(), Found: true}
				case "unwrap":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() any] {
						hx_receiver_385 := hx_structural_source_377
						return hxrt.BoundFunction(func() any {
							return hx_receiver_385.__hx_this.unwrap()
						}, func() any {
							return hx_receiver_385.__hx_this
						}, "unwrap")
					}(), Found: true}
				case "details":
					return hxrt.AnonymousField{Value: func() *hxrt.Function[func() *string] {
						hx_receiver_386 := hx_structural_source_377
						return hxrt.BoundFunction(func() *string {
							return hx_receiver_386.__hx_this.details()
						}, func() any {
							return hx_receiver_386.__hx_this
						}, "details")
					}(), Found: true}
				}
				return hxrt.AnonymousField{}
			}, func(name string, incoming any) bool {
				switch name {
				case "savedMessage":
					hx_structural_source_377.savedMessage = func(hx_value_379 any) *string {
						if hx_value_379 == nil {
							var hx_zero_380 *string
							return hx_zero_380
						}
						return hx_value_379.(*string)
					}(incoming)
					return true
				case "savedStack":
					hx_structural_source_377.savedStack = incoming
					return true
				case "savedPrevious":
					hx_structural_source_377.savedPrevious = hx_recover_class_haxe__Exception(incoming)
					return true
				case "savedNative":
					hx_structural_source_377.savedNative = incoming
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
	case *hxrt.Function[func(*string)]:
		return source
	case *hxrt.Function[func(*haxe__io__Bytes)]:
		return hxrt.ProjectFunction(source, func(arg_0 *string) {
			hxrt.FunctionCall(source)(hx_recover_class_haxe__io__Bytes(arg_0))
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
	case *hxrt.Function[func(any) *haxe__io__Bytes]:
		return hxrt.ProjectFunction(source, func(arg_0 *string) {
			hxrt.FunctionCall(source)(arg_0)
		})
	case *hxrt.Function[func(int) *haxe__io__Bytes]:
		return hxrt.ProjectFunction(source, func(arg_0 *string) {
			hxrt.FunctionCall(source)(func(hx_value_5 any) int {
				if hx_value_5 == nil {
					var hx_zero_6 int
					return hx_zero_6
				}
				return hx_value_5.(int)
			}(arg_0))
		})
	case *hxrt.Function[func(int) *string]:
		return hxrt.ProjectFunction(source, func(arg_0 *string) {
			hxrt.FunctionCall(source)(func(hx_value_7 any) int {
				if hx_value_7 == nil {
					var hx_zero_8 int
					return hx_zero_8
				}
				return hx_value_7.(int)
			}(arg_0))
		})
	case *hxrt.Function[func(float64)]:
		return hxrt.ProjectFunction(source, func(arg_0 *string) {
			hxrt.FunctionCall(source)(func(hx_value_9 any) float64 {
				if hx_value_9 == nil {
					var hx_zero_10 float64
					return hx_zero_10
				}
				return hx_value_9.(float64)
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
	case *hxrt.Function[func() int]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_7 any) *string {
				if hx_value_7 == nil {
					var hx_zero_8 *string
					return hx_zero_8
				}
				return hx_value_7.(*string)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func()]:
		return hxrt.ProjectFunction(source, func() *string {
			hxrt.FunctionCall(source)()
			return func(hx_value_9 any) *string {
				if hx_value_9 == nil {
					var hx_zero_10 *string
					return hx_zero_10
				}
				return hx_value_9.(*string)
			}(nil)
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
	case *hxrt.Function[func() *hxrt.Process]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_15 any) *string {
				if hx_value_15 == nil {
					var hx_zero_16 *string
					return hx_zero_16
				}
				return hx_value_15.(*string)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *haxe__io__Bytes]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_17 any) *string {
				if hx_value_17 == nil {
					var hx_zero_18 *string
					return hx_zero_18
				}
				return hx_value_17.(*string)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() []int]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_19 any) *string {
				if hx_value_19 == nil {
					var hx_zero_20 *string
					return hx_zero_20
				}
				return hx_value_19.(*string)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *hxrt.ByteView]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_21 any) *string {
				if hx_value_21 == nil {
					var hx_zero_22 *string
					return hx_zero_22
				}
				return hx_value_21.(*string)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *haxe__Exception]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_23 any) *string {
				if hx_value_23 == nil {
					var hx_zero_24 *string
					return hx_zero_24
				}
				return hx_value_23.(*string)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() haxe__IMap]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_25 any) *string {
				if hx_value_25 == nil {
					var hx_zero_26 *string
					return hx_zero_26
				}
				return hx_value_25.(*string)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *haxe__ds__StringMap]:
		return hxrt.ProjectFunction(source, func() *string {
			return func(hx_value_27 any) *string {
				if hx_value_27 == nil {
					var hx_zero_28 *string
					return hx_zero_28
				}
				return hx_value_27.(*string)
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
	case *hxrt.Function[func(int)]:
		return hxrt.ProjectFunction(source, func(arg_0 *haxe__io__Bytes) {
			hxrt.FunctionCall(source)(func(hx_value_1 any) int {
				if hx_value_1 == nil {
					var hx_zero_2 int
					return hx_zero_2
				}
				return hx_value_1.(int)
			}(arg_0))
		})
	case *hxrt.Function[func(*string)]:
		return hxrt.ProjectFunction(source, func(arg_0 *haxe__io__Bytes) {
			hxrt.FunctionCall(source)(func(hx_value_3 any) *string {
				if hx_value_3 == nil {
					var hx_zero_4 *string
					return hx_zero_4
				}
				return hx_value_3.(*string)
			}(arg_0))
		})
	case *hxrt.Function[func(*haxe__io__Bytes)]:
		return source
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
	case *hxrt.Function[func(any) *haxe__io__Bytes]:
		return hxrt.ProjectFunction(source, func(arg_0 *haxe__io__Bytes) {
			hxrt.FunctionCall(source)(arg_0)
		})
	case *hxrt.Function[func(int) *haxe__io__Bytes]:
		return hxrt.ProjectFunction(source, func(arg_0 *haxe__io__Bytes) {
			hxrt.FunctionCall(source)(func(hx_value_7 any) int {
				if hx_value_7 == nil {
					var hx_zero_8 int
					return hx_zero_8
				}
				return hx_value_7.(int)
			}(arg_0))
		})
	case *hxrt.Function[func(int) *string]:
		return hxrt.ProjectFunction(source, func(arg_0 *haxe__io__Bytes) {
			hxrt.FunctionCall(source)(func(hx_value_9 any) int {
				if hx_value_9 == nil {
					var hx_zero_10 int
					return hx_zero_10
				}
				return hx_value_9.(int)
			}(arg_0))
		})
	case *hxrt.Function[func(float64)]:
		return hxrt.ProjectFunction(source, func(arg_0 *haxe__io__Bytes) {
			hxrt.FunctionCall(source)(func(hx_value_11 any) float64 {
				if hx_value_11 == nil {
					var hx_zero_12 float64
					return hx_zero_12
				}
				return hx_value_11.(float64)
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
	case *hxrt.Function[func() int]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_7 any) bool {
				if hx_value_7 == nil {
					var hx_zero_8 bool
					return hx_zero_8
				}
				return hx_value_7.(bool)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func()]:
		return hxrt.ProjectFunction(source, func() bool {
			hxrt.FunctionCall(source)()
			return func(hx_value_9 any) bool {
				if hx_value_9 == nil {
					var hx_zero_10 bool
					return hx_zero_10
				}
				return hx_value_9.(bool)
			}(nil)
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
	case *hxrt.Function[func() *hxrt.Process]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_15 any) bool {
				if hx_value_15 == nil {
					var hx_zero_16 bool
					return hx_zero_16
				}
				return hx_value_15.(bool)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *haxe__io__Bytes]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_17 any) bool {
				if hx_value_17 == nil {
					var hx_zero_18 bool
					return hx_zero_18
				}
				return hx_value_17.(bool)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() []int]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_19 any) bool {
				if hx_value_19 == nil {
					var hx_zero_20 bool
					return hx_zero_20
				}
				return hx_value_19.(bool)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *hxrt.ByteView]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_21 any) bool {
				if hx_value_21 == nil {
					var hx_zero_22 bool
					return hx_zero_22
				}
				return hx_value_21.(bool)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *haxe__Exception]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_23 any) bool {
				if hx_value_23 == nil {
					var hx_zero_24 bool
					return hx_zero_24
				}
				return hx_value_23.(bool)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() haxe__IMap]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_25 any) bool {
				if hx_value_25 == nil {
					var hx_zero_26 bool
					return hx_zero_26
				}
				return hx_value_25.(bool)
			}(hxrt.FunctionCall(source)())
		})
	case *hxrt.Function[func() *haxe__ds__StringMap]:
		return hxrt.ProjectFunction(source, func() bool {
			return func(hx_value_27 any) bool {
				if hx_value_27 == nil {
					var hx_zero_28 bool
					return hx_zero_28
				}
				return hx_value_27.(bool)
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
	case *hxrt.Function[func(int)]:
		return source
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

func hx_recover_class_sys__io__Process(value any) *sys__io__Process {
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

func hx_is_nominal__Main__TrackingInput(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *_Main__TrackingInput:
		return (source != nil)
	default:
		return false
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

func hx_is_nominal_haxe___Int64_____Int64(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *haxe___Int64_____Int64:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_haxe__ds__StringMap(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *haxe__ds__StringMap:
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

func hx_is_nominal_haxe__http__HttpBase(value any) bool {
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

func hx_is_nominal_haxe__io__BytesOutput(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *haxe__io__BytesOutput:
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

func hx_is_nominal_haxe__io__Output(value any) bool {
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

func hx_is_nominal_haxe__iterators__MapKeyValueIterator(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *haxe__iterators__MapKeyValueIterator:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_haxe__iterators__StringIterator(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *haxe__iterators__StringIterator:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_haxe__iterators__StringKeyValueIterator(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *haxe__iterators__StringKeyValueIterator:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_sys__Http(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *sys__Http:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_sys__io__Process(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *sys__io__Process:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_sys__io___Process__ProcessInput(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *sys__io___Process__ProcessInput:
		return (source != nil)
	default:
		return false
	}
}

func hx_is_nominal_sys__io___Process__ProcessOutput(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *sys__io___Process__ProcessOutput:
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

func hx_is_nominal_sys__thread__Tls(value any) bool {
	value = hxrt.ObjectSource(value)
	switch source := value.(type) {
	case *sys__thread__Tls:
		return (source != nil)
	default:
		return false
	}
}
