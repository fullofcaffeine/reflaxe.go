package main

import "snapshot/hxrt"

type hxrt__TypeClassValue struct {
	name *string
}

type hxrt__TypeEnumValue struct {
	name *string
}

func finishServer(server map[string]any) {
	hxrt.TryCatch(func() {
		func(hx_obj_3 map[string]any) *sys__io__Process {
			hx_field_4 := hx_obj_3["process"]
			if hx_field_4 == nil {
				var hx_zero_5 *sys__io__Process
				return hx_zero_5
			}
			return hx_field_4.(*sys__io__Process)
		}(server).__hx_this.close()
	}, func(hx_caught_1 any) {
		hx_tmp := hx_caught_1
		_ = hx_tmp
	})
}

func main() {
	defer hxrt.ThreadWaitForAll()
	successServer := startServer(hxrt.StringFromLiteral("consume"), 1)
	success := run(successServer, hxrt.StringFromLiteral("payload"), 7, 2.0)
	finishServer(successServer)
	var v any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("success="), hxrt.StdString(((((((func(hx_obj_6 map[string]any) int {
		hx_field_7 := hx_obj_6["status"]
		if hx_field_7 == nil {
			var hx_zero_8 int
			return hx_zero_8
		}
		return hx_field_7.(int)
	}(success) == 200) && hxrt.StringEqualStringPtr(func(hx_obj_9 map[string]any) *string {
		hx_field_10 := hx_obj_9["error"]
		if hx_field_10 == nil {
			var hx_zero_11 *string
			return hx_zero_11
		}
		return hx_field_10.(*string)
	}(success), hxrt.StringFromLiteral(""))) && hxrt.StringEqualStringPtr(func(hx_obj_12 map[string]any) *string {
		hx_field_13 := hx_obj_12["data"]
		if hx_field_13 == nil {
			var hx_zero_14 *string
			return hx_zero_14
		}
		return hx_field_13.(*string)
	}(success), hxrt.StringFromLiteral("ok"))) && (func(hx_obj_15 map[string]any) int {
		hx_field_16 := hx_obj_15["reads"]
		if hx_field_16 == nil {
			var hx_zero_17 int
			return hx_zero_17
		}
		return hx_field_16.(int)
	}(success) == 4)) && !func(hx_obj_18 map[string]any) bool {
		hx_field_19 := hx_obj_18["wrongThread"]
		if hx_field_19 == nil {
			var hx_zero_20 bool
			return hx_zero_20
		}
		return hx_field_19.(bool)
	}(success)) && (func(hx_obj_21 map[string]any) int {
		hx_field_22 := hx_obj_21["afterReturnReads"]
		if hx_field_22 == nil {
			var hx_zero_23 int
			return hx_zero_23
		}
		return hx_field_22.(int)
	}(success) == 0)))))
	hxrt.Println(v)
	sourceServer := startServer(hxrt.StringFromLiteral("consume"), 1)
	source := run(sourceServer, hxrt.StringFromLiteral("source-error"), 7, 2.0)
	finishServer(sourceServer)
	var v_1 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("source="), hxrt.StdString((((((func(hx_obj_24 map[string]any) int {
		hx_field_25 := hx_obj_24["status"]
		if hx_field_25 == nil {
			var hx_zero_26 int
			return hx_zero_26
		}
		return hx_field_25.(int)
	}(source) == -1) && hxrt.StringEqualStringPtr(func(hx_obj_27 map[string]any) *string {
		hx_field_28 := hx_obj_27["error"]
		if hx_field_28 == nil {
			var hx_zero_29 *string
			return hx_zero_29
		}
		return hx_field_28.(*string)
	}(source), hxrt.StringFromLiteral("source-exploded"))) && (func(hx_obj_30 map[string]any) int {
		hx_field_31 := hx_obj_30["reads"]
		if hx_field_31 == nil {
			var hx_zero_32 int
			return hx_zero_32
		}
		return hx_field_31.(int)
	}(source) == 2)) && !func(hx_obj_33 map[string]any) bool {
		hx_field_34 := hx_obj_33["wrongThread"]
		if hx_field_34 == nil {
			var hx_zero_35 bool
			return hx_zero_35
		}
		return hx_field_34.(bool)
	}(source)) && (func(hx_obj_36 map[string]any) int {
		hx_field_37 := hx_obj_36["afterReturnReads"]
		if hx_field_37 == nil {
			var hx_zero_38 int
			return hx_zero_38
		}
		return hx_field_37.(int)
	}(source) == 0)))))
	hxrt.Println(v_1)
	eofServer := startServer(hxrt.StringFromLiteral("consume"), 1)
	eof := run(eofServer, hxrt.StringFromLiteral("early-eof"), 7, 2.0)
	finishServer(eofServer)
	var v_2 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("earlyEof="), hxrt.StdString((((((func(hx_obj_39 map[string]any) int {
		hx_field_40 := hx_obj_39["status"]
		if hx_field_40 == nil {
			var hx_zero_41 int
			return hx_zero_41
		}
		return hx_field_40.(int)
	}(eof) == -1) && hxrt.StringEqualStringPtr(func(hx_obj_42 map[string]any) *string {
		hx_field_43 := hx_obj_42["error"]
		if hx_field_43 == nil {
			var hx_zero_44 *string
			return hx_zero_44
		}
		return hx_field_43.(*string)
	}(eof), hxrt.StringFromLiteral("Transfer aborted"))) && (func(hx_obj_45 map[string]any) int {
		hx_field_46 := hx_obj_45["reads"]
		if hx_field_46 == nil {
			var hx_zero_47 int
			return hx_zero_47
		}
		return hx_field_46.(int)
	}(eof) == 2)) && !func(hx_obj_48 map[string]any) bool {
		hx_field_49 := hx_obj_48["wrongThread"]
		if hx_field_49 == nil {
			var hx_zero_50 bool
			return hx_zero_50
		}
		return hx_field_49.(bool)
	}(eof)) && (func(hx_obj_51 map[string]any) int {
		hx_field_52 := hx_obj_51["afterReturnReads"]
		if hx_field_52 == nil {
			var hx_zero_53 int
			return hx_zero_53
		}
		return hx_field_52.(int)
	}(eof) == 0)))))
	hxrt.Println(v_2)
	zeroServer := startServer(hxrt.StringFromLiteral("consume"), 1)
	zero := run(zeroServer, hxrt.StringFromLiteral("zero"), 7, 2.0)
	finishServer(zeroServer)
	var v_3 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("zero="), hxrt.StdString((((((func(hx_obj_54 map[string]any) int {
		hx_field_55 := hx_obj_54["status"]
		if hx_field_55 == nil {
			var hx_zero_56 int
			return hx_zero_56
		}
		return hx_field_55.(int)
	}(zero) == -1) && hxrt.StringEqualStringPtr(func(hx_obj_57 map[string]any) *string {
		hx_field_58 := hx_obj_57["error"]
		if hx_field_58 == nil {
			var hx_zero_59 *string
			return hx_zero_59
		}
		return hx_field_58.(*string)
	}(zero), hxrt.StringFromLiteral("multipart upload made no progress"))) && (func(hx_obj_60 map[string]any) int {
		hx_field_61 := hx_obj_60["reads"]
		if hx_field_61 == nil {
			var hx_zero_62 int
			return hx_zero_62
		}
		return hx_field_61.(int)
	}(zero) == 1)) && !func(hx_obj_63 map[string]any) bool {
		hx_field_64 := hx_obj_63["wrongThread"]
		if hx_field_64 == nil {
			var hx_zero_65 bool
			return hx_zero_65
		}
		return hx_field_64.(bool)
	}(zero)) && (func(hx_obj_66 map[string]any) int {
		hx_field_67 := hx_obj_66["afterReturnReads"]
		if hx_field_67 == nil {
			var hx_zero_68 int
			return hx_zero_68
		}
		return hx_field_67.(int)
	}(zero) == 0)))))
	hxrt.Println(v_3)
	closeServer := startServer(hxrt.StringFromLiteral("close"), 1)
	closed := run(closeServer, hxrt.StringFromLiteral("stream"), 8388608, 2.0)
	finishServer(closeServer)
	var v_4 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("serverClose="), hxrt.StdString(((!hxrt.StringEqualStringPtr(func(hx_obj_69 map[string]any) *string {
		hx_field_70 := hx_obj_69["error"]
		if hx_field_70 == nil {
			var hx_zero_71 *string
			return hx_zero_71
		}
		return hx_field_70.(*string)
	}(closed), hxrt.StringFromLiteral("")) && !func(hx_obj_72 map[string]any) bool {
		hx_field_73 := hx_obj_72["wrongThread"]
		if hx_field_73 == nil {
			var hx_zero_74 bool
			return hx_zero_74
		}
		return hx_field_73.(bool)
	}(closed)) && (func(hx_obj_75 map[string]any) int {
		hx_field_76 := hx_obj_75["afterReturnReads"]
		if hx_field_76 == nil {
			var hx_zero_77 int
			return hx_zero_77
		}
		return hx_field_76.(int)
	}(closed) == 0)))))
	hxrt.Println(v_4)
	timeoutServer := startServer(hxrt.StringFromLiteral("timeout"), 1)
	timed := run(timeoutServer, hxrt.StringFromLiteral("stream"), 16777216, 0.1)
	finishServer(timeoutServer)
	var v_5 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("timeout="), hxrt.StdString(((!hxrt.StringEqualStringPtr(func(hx_obj_78 map[string]any) *string {
		hx_field_79 := hx_obj_78["error"]
		if hx_field_79 == nil {
			var hx_zero_80 *string
			return hx_zero_80
		}
		return hx_field_79.(*string)
	}(timed), hxrt.StringFromLiteral("")) && !func(hx_obj_81 map[string]any) bool {
		hx_field_82 := hx_obj_81["wrongThread"]
		if hx_field_82 == nil {
			var hx_zero_83 bool
			return hx_zero_83
		}
		return hx_field_82.(bool)
	}(timed)) && (func(hx_obj_84 map[string]any) int {
		hx_field_85 := hx_obj_84["afterReturnReads"]
		if hx_field_85 == nil {
			var hx_zero_86 int
			return hx_zero_86
		}
		return hx_field_85.(int)
	}(timed) == 0)))))
	hxrt.Println(v_5)
	repeatedServer := startServer(hxrt.StringFromLiteral("early"), 12)
	repeatedStatus := 0
	repeatedError := 0
	repeatedThread := 0
	repeatedAfterReturn := 0
	_g := 0
	for _g < 12 {
		hx_post_87 := _g
		_g = int(int32((_g + 1)))
		hx_tmp := hx_post_87
		_ = hx_tmp
		early := run(repeatedServer, hxrt.StringFromLiteral("stream"), 8388608, 2.0)
		if func(hx_obj_88 map[string]any) int {
			hx_field_89 := hx_obj_88["status"]
			if hx_field_89 == nil {
				var hx_zero_90 int
				return hx_zero_90
			}
			return hx_field_89.(int)
		}(early) == 413 {
			repeatedStatus = int(int32((repeatedStatus + 1)))
		}
		if hxrt.StringEqualStringPtr(func(hx_obj_91 map[string]any) *string {
			hx_field_92 := hx_obj_91["error"]
			if hx_field_92 == nil {
				var hx_zero_93 *string
				return hx_zero_93
			}
			return hx_field_92.(*string)
		}(early), hxrt.StringFromLiteral("Http Error #413")) {
			repeatedError = int(int32((repeatedError + 1)))
		}
		if !func(hx_obj_94 map[string]any) bool {
			hx_field_95 := hx_obj_94["wrongThread"]
			if hx_field_95 == nil {
				var hx_zero_96 bool
				return hx_zero_96
			}
			return hx_field_95.(bool)
		}(early) {
			repeatedThread = int(int32((repeatedThread + 1)))
		}
		if func(hx_obj_97 map[string]any) int {
			hx_field_98 := hx_obj_97["afterReturnReads"]
			if hx_field_98 == nil {
				var hx_zero_99 int
				return hx_zero_99
			}
			return hx_field_98.(int)
		}(early) == 0 {
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

func run(server map[string]any, mode *string, declaredSize int, timeout float64) map[string]any {
	input := New__Main__TrackingInput(mode, func() *string {
		var hx_if_100 *string
		if hxrt.StringEqualStringPtr(mode, hxrt.StringFromLiteral("payload")) {
			hx_if_100 = hxrt.StringFromLiteral("payload")
		} else {
			hx_if_100 = hxrt.StringFromLiteral("")
		}
		return hx_if_100
	}())
	request := New_sys__Http(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("http://127.0.0.1:"), func(hx_obj_101 map[string]any) int {
		hx_field_102 := hx_obj_101["port"]
		if hx_field_102 == nil {
			var hx_zero_103 int
			return hx_zero_103
		}
		return hx_field_102.(int)
	}(server)), hxrt.StringFromLiteral("/upload")))
	request.cnxTimeout = timeout
	request.__hx_this.fileTransfer(hxrt.StringFromLiteral("asset"), hxrt.StringFromLiteral("demo.bin"), input.haxe__io__Input, declaredSize, hxrt.StringFromLiteral("application/octet-stream"))
	status := -1
	error := hxrt.StringFromLiteral("")
	data := hxrt.StringFromLiteral("")
	request.onStatus = func(value int) {
		status = value
	}
	request.onError = func(value *string) {
		error = value
	}
	request.onData = func(value *string) {
		data = value
	}
	request.__hx_this.request(true)
	input.requestReturned = true
	hxrt.SysSleep(0.01)
	hx_obj_104 := map[string]any{}
	hx_obj_104["status"] = status
	hx_obj_104["error"] = error
	hx_obj_104["data"] = data
	hx_obj_104["reads"] = input.reads
	hx_obj_104["wrongThread"] = input.wrongThread
	hx_obj_104["afterReturnReads"] = input.afterReturnReads
	return hx_obj_104
}

func startServer(mode *string, count int) map[string]any {
	process := New_sys__io__Process(hxrt.StringFromLiteral("python3"), hxrt.NewArray(hxrt.StringFromLiteral("-u"), hxrt.StringFromLiteral("-c"), pythonServerScript(mode, count)), false)
	var port any = Std_parseInt(process.stdout.__hx_this.readLine())
	if port == nil {
		process.__hx_this.close()
		hxrt.Throw(hxrt.StringFromLiteral("failed to read server port"))
	}
	hx_obj_105 := map[string]any{}
	hx_obj_105["process"] = process
	hx_obj_105["port"] = port.(int)
	return hx_obj_105
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

func New__Main__TrackingInput(mode *string, payload *string) *_Main__TrackingInput {
	self := &_Main__TrackingInput{}
	self.haxe__io__Input = New_haxe__io__Input()
	self.haxe__io__Input.__hx_this = self
	self.__hx_this = self
	self.position = 0
	self.requestReturned = false
	self.afterReturnReads = 0
	self.wrongThread = false
	self.reads = 0
	self.mode = mode
	self.payload = haxe__io__Bytes_ofString(payload, nil)
	self.callerMarker = New_sys__thread__Tls()
	func(hx_value_106 any) *string {
		if hx_value_106 == nil {
			var hx_zero_107 *string
			return hx_zero_107
		}
		return hx_value_106.(*string)
	}(self.callerMarker.__hx_this.set_value(hxrt.StringFromLiteral("source-caller")))
	return self
}

func (self *_Main__TrackingInput) readBytes(bytes *haxe__io__Bytes, targetPos int, requested int) int {
	self.reads = int(int32((self.reads + 1)))
	if !hxrt.StringEqualStringPtr(func(hx_value_108 any) *string {
		if hx_value_108 == nil {
			var hx_zero_109 *string
			return hx_zero_109
		}
		return hx_value_108.(*string)
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
			hxrt.Throw(New_haxe__io__Eof())
		}
		bytes.b[targetPos] = 120
		bytes.__hx_rawValid = false
		self.position = int(int32((self.position + 1)))
		return 1
	case *hxrt.StdString(hxrt.StringFromLiteral("source-error")):
		if self.position > 0 {
			hxrt.Throw(hxrt.StringFromLiteral("source-exploded"))
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
			hxrt.Throw(New_haxe__io__Eof())
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
