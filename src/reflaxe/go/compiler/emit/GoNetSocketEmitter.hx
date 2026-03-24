package reflaxe.go.compiler.emit;

#if macro
import reflaxe.go.ast.GoAST.GoDecl;
import reflaxe.go.ast.GoAST.GoExpr;
import reflaxe.go.ast.GoAST.GoStmt;

/**
	What:
	Builds the compiler-owned declarations behind `sys.net.Host`,
	`sys.net.Socket`, and the Go socket helper layer.

	Why:
	Socket and deadline behavior are runtime-sensitive and stay compiler-owned
	for now, but the declaration block does not belong inline inside
	`GoCompiler`.

	How:
	Emits the same socket shim declaration set that previously lived inside
	`lowerNetSocketShimDecls()`. This extraction is relocation only; no
	semantic changes belong here.
**/
class GoNetSocketEmitter {
	public static function emit():Array<GoDecl> {
		return [
			GoDecl.GoStructDecl("sys__net__Host", [
				{name: "host", typeName: "*string"},
				{name: "ip", typeName: "int"},
				{
					name: "resolved",
					typeName: "*string"
				}
			]),
			GoDecl.GoFuncDecl("hxrt__host_empty", null, [], ["*sys__net__Host"], [
				GoStmt.GoReturn(GoExpr.GoRaw("&sys__net__Host{host: hxrt.StringFromLiteral(\"\"), ip: 0, resolved: hxrt.StringFromLiteral(\"\")}"))
			]),
			GoDecl.GoFuncDecl("New_sys__net__Host", null, [
				{
					name: "name",
					typeName: "*string"
				}
			], ["*sys__net__Host"], [
				GoStmt.GoRaw("if name == nil {"),
				GoStmt.GoRaw("\thxrt.Throw(hxrt.StringFromLiteral(\"Could not resolve host\"))"),
				GoStmt.GoRaw("\treturn hxrt__host_empty()"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("rawName := *hxrt.StdString(name)"),
				GoStmt.GoRaw("ips, err := net.LookupIP(rawName)"),
				GoStmt.GoRaw("if err != nil || len(ips) == 0 {"),
				GoStmt.GoRaw("\thxrt.Throw(hxrt.StringFromLiteral(\"Could not resolve host\"))"),
				GoStmt.GoRaw("\treturn hxrt__host_empty()"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("selected := ips[0]"),
				GoStmt.GoRaw("for _, candidate := range ips {"),
				GoStmt.GoRaw("\tif v4 := candidate.To4(); v4 != nil {"),
				GoStmt.GoRaw("\t\tselected = v4"),
				GoStmt.GoRaw("\t\tbreak"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("resolved := hxrt.StringFromLiteral(selected.String())"),
				GoStmt.GoReturn(GoExpr.GoRaw("&sys__net__Host{host: name, ip: 0, resolved: resolved}"))
			]),
			GoDecl.GoFuncDecl("toString", {
				name: "self",
				typeName: "*sys__net__Host"
			}, [], ["*string"], [
				GoStmt.GoIf(GoExpr.GoRaw("self == nil || self.resolved == nil"), [
					GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("hxrt.StringFromLiteral"), [GoExpr.GoStringLiteral("")]))
				], null),
				GoStmt.GoReturn(GoExpr.GoSelector(GoExpr.GoIdent("self"), "resolved"))
			]),
			GoDecl.GoFuncDecl("reverse", {
				name: "self",
				typeName: "*sys__net__Host"
			}, [], ["*string"], [
				GoStmt.GoRaw("if self == nil || self.resolved == nil {"),
				GoStmt.GoRaw("\thxrt.Throw(hxrt.StringFromLiteral(\"Could not reverse host\"))"),
				GoStmt.GoRaw("\treturn hxrt.StringFromLiteral(\"\")"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("names, err := net.LookupAddr(*hxrt.StdString(self.resolved))"),
				GoStmt.GoRaw("if err != nil || len(names) == 0 {"),
				GoStmt.GoRaw("\thxrt.Throw(hxrt.StringFromLiteral(\"Could not reverse host\"))"),
				GoStmt.GoRaw("\treturn hxrt.StringFromLiteral(\"\")"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("resolved := strings.TrimSuffix(names[0], \".\")"),
				GoStmt.GoReturn(GoExpr.GoRaw("hxrt.StringFromLiteral(resolved)"))
			]),
			GoDecl.GoFuncDecl("sys__net__Host_localhost", null, [], ["*string"], [
				GoStmt.GoRaw("name, err := os.Hostname()"),
				GoStmt.GoRaw("if err != nil || name == \"\" {"),
				GoStmt.GoRaw("\treturn hxrt.StringFromLiteral(\"localhost\")"),
				GoStmt.GoRaw("}"),
				GoStmt.GoReturn(GoExpr.GoRaw("hxrt.StringFromLiteral(name)"))
			]),
			GoDecl.GoStructDecl("sys__net__SocketInput", [
				{
					name: "reader",
					typeName: "*bufio.Reader"
				},
				{name: "socket", typeName: "*sys__net__Socket"}
			]),
			GoDecl.GoStructDecl("sys__net__SocketOutput", [
				{name: "writer", typeName: "*bufio.Writer"},
				{name: "socket", typeName: "*sys__net__Socket"}
			]),
			GoDecl.GoStructDecl("sys__net__Socket", [
				{name: "input", typeName: "*sys__net__SocketInput"},
				{name: "output", typeName: "*sys__net__SocketOutput"},
				{name: "custom", typeName: "any"},
				{name: "conn", typeName: "net.Conn"},
				{name: "listener", typeName: "net.Listener"},
				{name: "timeout", typeName: "float64"},
				{name: "hasTimeout", typeName: "bool"},
				{name: "blocking", typeName: "bool"},
				{name: "fastSend", typeName: "bool"}
			]),
			GoDecl.GoFuncDecl("New_sys__net__Socket", null, [], ["*sys__net__Socket"], [
				GoStmt.GoReturn(GoExpr.GoRaw("&sys__net__Socket{input: &sys__net__SocketInput{}, output: &sys__net__SocketOutput{}, blocking: true}"))
			]),
			GoDecl.GoFuncDecl("hxrt__socket_deadline", null, [
				{
					name: "timeout",
					typeName: "float64"
				}
			], ["time.Time"], [
				GoStmt.GoRaw("duration := time.Duration(timeout * float64(time.Second))"),
				GoStmt.GoRaw("return time.Now().Add(duration)")
			]),
			GoDecl.GoFuncDecl("hxrt__socket_applyConnDeadline", {
				name: "self",
				typeName: "*sys__net__Socket"
			}, [], [], [
				GoStmt.GoRaw("if self == nil || self.conn == nil {"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if !self.blocking {"),
				GoStmt.GoRaw("\t_ = self.conn.SetDeadline(time.Now())"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if self.hasTimeout {"),
				GoStmt.GoRaw("\t_ = self.conn.SetDeadline(hxrt__socket_deadline(self.timeout))"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("_ = self.conn.SetDeadline(time.Time{})")
			]),
			GoDecl.GoFuncDecl("hxrt__socket_applyListenerDeadline", {
				name: "self",
				typeName: "*sys__net__Socket"
			}, [], [], [
				GoStmt.GoRaw("if self == nil || self.listener == nil {"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("tcpListener, ok := self.listener.(*net.TCPListener)"),
				GoStmt.GoRaw("if !ok {"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if !self.blocking {"),
				GoStmt.GoRaw("\t_ = tcpListener.SetDeadline(time.Now())"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if self.hasTimeout {"),
				GoStmt.GoRaw("\t_ = tcpListener.SetDeadline(hxrt__socket_deadline(self.timeout))"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("_ = tcpListener.SetDeadline(time.Time{})")
			]),
			GoDecl.GoFuncDecl("hxrt__socket_applyFastSend", {
				name: "self",
				typeName: "*sys__net__Socket"
			}, [], [], [
				GoStmt.GoRaw("if self == nil || self.conn == nil {"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("tcpConn, ok := self.conn.(*net.TCPConn)"),
				GoStmt.GoRaw("if !ok {"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if err := tcpConn.SetNoDelay(self.fastSend); err != nil {"),
				GoStmt.GoRaw("\thxrt.Throw(err)"),
				GoStmt.GoRaw("}")
			]),
			GoDecl.GoFuncDecl("hxrt__socket_setConn", {
				name: "self",
				typeName: "*sys__net__Socket"
			}, [{name: "conn", typeName: "net.Conn"}], [], [
				GoStmt.GoIf(GoExpr.GoRaw("self == nil || conn == nil"), [GoStmt.GoReturn(null)], null),
				GoStmt.GoAssign(GoExpr.GoSelector(GoExpr.GoIdent("self"), "conn"), GoExpr.GoIdent("conn")),
				GoStmt.GoAssign(GoExpr.GoSelector(GoExpr.GoIdent("self"), "input"),
					GoExpr.GoRaw("&sys__net__SocketInput{reader: bufio.NewReader(conn), socket: self}")),
				GoStmt.GoAssign(GoExpr.GoSelector(GoExpr.GoIdent("self"), "output"),
					GoExpr.GoRaw("&sys__net__SocketOutput{writer: bufio.NewWriter(conn), socket: self}")),
				GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoIdent("self"), "hxrt__socket_applyFastSend"), [])),
				GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoIdent("self"), "hxrt__socket_applyConnDeadline"), []))
			]),
			GoDecl.GoFuncDecl("hxrt__socket_conn", {
				name: "self",
				typeName: "*sys__net__Socket"
			}, [], ["net.Conn"], [
				GoStmt.GoIf(GoExpr.GoRaw("self == nil"), [GoStmt.GoReturn(GoExpr.GoNil)], null),
				GoStmt.GoReturn(GoExpr.GoSelector(GoExpr.GoIdent("self"), "conn"))
			]),
			GoDecl.GoFuncDecl("close", {
				name: "self",
				typeName: "*sys__net__Socket"
			}, [], [], [
				GoStmt.GoIf(GoExpr.GoRaw("self == nil"), [GoStmt.GoReturn(null)], null),
				GoStmt.GoIf(GoExpr.GoRaw("self.conn != nil"), [
					GoStmt.GoRaw("_ = self.conn.Close()"),
					GoStmt.GoAssign(GoExpr.GoSelector(GoExpr.GoIdent("self"), "conn"), GoExpr.GoNil)
				], null),
				GoStmt.GoIf(GoExpr.GoRaw("self.listener != nil"), [
					GoStmt.GoRaw("_ = self.listener.Close()"),
					GoStmt.GoAssign(GoExpr.GoSelector(GoExpr.GoIdent("self"), "listener"), GoExpr.GoNil)
				], null)
			]),
			GoDecl.GoFuncDecl("connect", {
				name: "self",
				typeName: "*sys__net__Socket"
			}, [
				{
					name: "host",
					typeName: "*sys__net__Host"
				},
				{name: "port", typeName: "int"}
			], [], [
				GoStmt.GoIf(GoExpr.GoRaw("self == nil || host == nil"), [
					GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoIdent("hxrt.Throw"),
						[
							GoExpr.GoCall(GoExpr.GoIdent("hxrt.StringFromLiteral"), [GoExpr.GoStringLiteral("socket connect requires host")])
						])),
					GoStmt.GoReturn(null)
				],
					null),
				GoStmt.GoVarDecl("resolvedHost", "*string", GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoIdent("host"), "toString"), []), true),
				GoStmt.GoIf(GoExpr.GoRaw("resolvedHost == nil"), [
					GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoIdent("hxrt.Throw"), [
						GoExpr.GoCall(GoExpr.GoIdent("hxrt.StringFromLiteral"), [GoExpr.GoStringLiteral("socket connect requires host")])
					])),
					GoStmt.GoReturn(null)
				], null),
				GoStmt.GoVarDecl("address", null, GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoIdent("net"), "JoinHostPort"), [
					GoExpr.GoRaw("*hxrt.StdString(resolvedHost)"),
					GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoIdent("strconv"), "Itoa"), [GoExpr.GoIdent("port")])
				]), true),
				GoStmt.GoRaw("conn, err := net.Dial(\"tcp\", address)"),
				GoStmt.GoIf(GoExpr.GoBinary("!=", GoExpr.GoIdent("err"), GoExpr.GoNil), [
					GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoIdent("hxrt.Throw"), [GoExpr.GoIdent("err")])),
					GoStmt.GoReturn(null)
				],
					null),
				GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoIdent("self"), "hxrt__socket_setConn"), [GoExpr.GoIdent("conn")]))
			]),
			GoDecl.GoFuncDecl("bind", {
				name: "self",
				typeName: "*sys__net__Socket"
			}, [
				{
					name: "host",
					typeName: "*sys__net__Host"
				},
				{name: "port", typeName: "int"}
			], [], [
				GoStmt.GoIf(GoExpr.GoRaw("self == nil || host == nil"), [
					GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoIdent("hxrt.Throw"),
						[
							GoExpr.GoCall(GoExpr.GoIdent("hxrt.StringFromLiteral"), [GoExpr.GoStringLiteral("socket bind requires host")])
						])),
					GoStmt.GoReturn(null)
				],
					null),
				GoStmt.GoVarDecl("resolvedHost", "*string", GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoIdent("host"), "toString"), []), true),
				GoStmt.GoIf(GoExpr.GoRaw("resolvedHost == nil"), [
					GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoIdent("hxrt.Throw"), [
						GoExpr.GoCall(GoExpr.GoIdent("hxrt.StringFromLiteral"), [GoExpr.GoStringLiteral("socket bind requires host")])
					])),
					GoStmt.GoReturn(null)
				], null),
				GoStmt.GoVarDecl("address", null, GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoIdent("net"), "JoinHostPort"), [
					GoExpr.GoRaw("*hxrt.StdString(resolvedHost)"),
					GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoIdent("strconv"), "Itoa"), [GoExpr.GoIdent("port")])
				]), true),
				GoStmt.GoRaw("listener, err := net.Listen(\"tcp\", address)"),
				GoStmt.GoIf(GoExpr.GoBinary("!=", GoExpr.GoIdent("err"), GoExpr.GoNil), [
					GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoIdent("hxrt.Throw"), [GoExpr.GoIdent("err")])),
					GoStmt.GoReturn(null)
				],
					null),
				GoStmt.GoIf(GoExpr.GoRaw("self.listener != nil"), [GoStmt.GoRaw("_ = self.listener.Close()")], null),
				GoStmt.GoAssign(GoExpr.GoSelector(GoExpr.GoIdent("self"), "listener"), GoExpr.GoIdent("listener")),
				GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoIdent("self"), "hxrt__socket_applyListenerDeadline"), []))
			]),
			GoDecl.GoFuncDecl("listen", {
				name: "self",
				typeName: "*sys__net__Socket"
			}, [{name: "connections", typeName: "int"}], [],
				[GoStmt.GoRaw("_ = connections")]),
			GoDecl.GoFuncDecl("accept", {
				name: "self",
				typeName: "*sys__net__Socket"
			}, [], ["*sys__net__Socket"], [
				GoStmt.GoIf(GoExpr.GoRaw("self == nil || self.listener == nil"), [
					GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoIdent("hxrt.Throw"),
						[
							GoExpr.GoCall(GoExpr.GoIdent("hxrt.StringFromLiteral"), [GoExpr.GoStringLiteral("socket accept requires listener")])
						])),
					GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("New_sys__net__Socket"), []))
				],
					null),
				GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoIdent("self"), "hxrt__socket_applyListenerDeadline"), [])),
				GoStmt.GoRaw("conn, err := self.listener.Accept()"),
				GoStmt.GoIf(GoExpr.GoBinary("!=", GoExpr.GoIdent("err"), GoExpr.GoNil), [
					GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoIdent("hxrt.Throw"), [GoExpr.GoIdent("err")])),
					GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("New_sys__net__Socket"), []))
				],
					null),
				GoStmt.GoVarDecl("accepted", null, GoExpr.GoCall(GoExpr.GoIdent("New_sys__net__Socket"), []), true),
				GoStmt.GoAssign(GoExpr.GoSelector(GoExpr.GoIdent("accepted"), "timeout"), GoExpr.GoSelector(GoExpr.GoIdent("self"), "timeout")),
				GoStmt.GoAssign(GoExpr.GoSelector(GoExpr.GoIdent("accepted"), "hasTimeout"), GoExpr.GoSelector(GoExpr.GoIdent("self"), "hasTimeout")),
				GoStmt.GoAssign(GoExpr.GoSelector(GoExpr.GoIdent("accepted"), "blocking"), GoExpr.GoSelector(GoExpr.GoIdent("self"), "blocking")),
				GoStmt.GoAssign(GoExpr.GoSelector(GoExpr.GoIdent("accepted"), "fastSend"), GoExpr.GoSelector(GoExpr.GoIdent("self"), "fastSend")),
				GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoIdent("accepted"), "hxrt__socket_setConn"), [GoExpr.GoIdent("conn")])),
				GoStmt.GoReturn(GoExpr.GoIdent("accepted"))
			]),
			GoDecl.GoFuncDecl("read", {
				name: "self",
				typeName: "*sys__net__Socket"
			}, [], ["*string"], [
				GoStmt.GoIf(GoExpr.GoRaw("self == nil || self.input == nil"), [
					GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("hxrt.StringFromLiteral"), [GoExpr.GoStringLiteral("")]))
				],
					null),
				GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoSelector(GoExpr.GoIdent("self"), "input"), "readLine"), []))
			]),
			GoDecl.GoFuncDecl("write", {
				name: "self",
				typeName: "*sys__net__Socket"
			}, [{name: "content", typeName: "*string"}], [], [
				GoStmt.GoIf(GoExpr.GoRaw("self == nil || self.output == nil"), [GoStmt.GoReturn(null)], null),
				GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoSelector(GoExpr.GoIdent("self"), "output"), "writeString"),
					[GoExpr.GoIdent("content")])),
				GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoSelector(GoExpr.GoIdent("self"), "output"), "flush"), []))
			]),
			GoDecl.GoFuncDecl("shutdown", {
				name: "self",
				typeName: "*sys__net__Socket"
			}, [
				{
					name: "read",
					typeName: "bool"
				},
				{name: "write", typeName: "bool"}
			], [], [
				GoStmt.GoRaw("if self == nil || self.conn == nil || (!read && !write) {"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if tcpConn, ok := self.conn.(*net.TCPConn); ok {"),
				GoStmt.GoRaw("\tif read {"),
				GoStmt.GoRaw("\t\tif err := tcpConn.CloseRead(); err != nil {"),
				GoStmt.GoRaw("\t\t\thxrt.Throw(err)"),
				GoStmt.GoRaw("\t\t}"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tif write {"),
				GoStmt.GoRaw("\t\tif err := tcpConn.CloseWrite(); err != nil {"),
				GoStmt.GoRaw("\t\t\thxrt.Throw(err)"),
				GoStmt.GoRaw("\t\t}"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if err := self.conn.Close(); err != nil {"),
				GoStmt.GoRaw("\thxrt.Throw(err)"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("self.conn = nil")
			]),
			GoDecl.GoFuncDecl("hxrt__socket_addrInfo", null, [
				{
					name: "addr",
					typeName: "net.Addr"
				}
			], ["map[string]any"], [
				GoStmt.GoRaw("if addr == nil {"),
				GoStmt.GoRaw("\treturn map[string]any{\"host\": hxrt__host_empty(), \"port\": 0}"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("rawHost := \"\""),
				GoStmt.GoRaw("rawPort := \"0\""),
				GoStmt.GoRaw("hostPart, portPart, err := net.SplitHostPort(addr.String())"),
				GoStmt.GoRaw("if err == nil {"),
				GoStmt.GoRaw("\trawHost = hostPart"),
				GoStmt.GoRaw("\trawPort = portPart"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("port, _ := strconv.Atoi(rawPort)"),
				GoStmt.GoRaw("if rawHost == \"\" {"),
				GoStmt.GoRaw("\treturn map[string]any{\"host\": hxrt__host_empty(), \"port\": port}"),
				GoStmt.GoRaw("}"),
				GoStmt.GoReturn(GoExpr.GoRaw("map[string]any{\"host\": New_sys__net__Host(hxrt.StringFromLiteral(rawHost)), \"port\": port}"))
			]),
			GoDecl.GoFuncDecl("peer", {
				name: "self",
				typeName: "*sys__net__Socket"
			}, [], ["map[string]any"], [
				GoStmt.GoIf(GoExpr.GoRaw("self == nil || self.conn == nil"), [
					GoStmt.GoReturn(GoExpr.GoRaw("map[string]any{\"host\": hxrt__host_empty(), \"port\": 0}"))
				], null),
				GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("hxrt__socket_addrInfo"), [
					GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoSelector(GoExpr.GoIdent("self"), "conn"), "RemoteAddr"), [])
				]))
			]),
			GoDecl.GoFuncDecl("host", {
				name: "self",
				typeName: "*sys__net__Socket"
			}, [], ["map[string]any"], [
				GoStmt.GoIf(GoExpr.GoRaw("self == nil"), [
					GoStmt.GoReturn(GoExpr.GoRaw("map[string]any{\"host\": hxrt__host_empty(), \"port\": 0}"))
				], null),
				GoStmt.GoIf(GoExpr.GoRaw("self.conn != nil"), [
					GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("hxrt__socket_addrInfo"), [
						GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoSelector(GoExpr.GoIdent("self"), "conn"), "LocalAddr"), [])
					]))
				], null),
				GoStmt.GoIf(GoExpr.GoRaw("self.listener != nil"), [
					GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("hxrt__socket_addrInfo"), [
						GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoSelector(GoExpr.GoIdent("self"), "listener"), "Addr"), [])
					]))
				], null),
				GoStmt.GoReturn(GoExpr.GoRaw("map[string]any{\"host\": hxrt__host_empty(), \"port\": 0}"))
			]),
			GoDecl.GoFuncDecl("setTimeout", {
				name: "self",
				typeName: "*sys__net__Socket"
			}, [{name: "timeout", typeName: "float64"}], [], [
				GoStmt.GoRaw("if self == nil {"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if timeout < 0 {"),
				GoStmt.GoRaw("\tself.hasTimeout = false"),
				GoStmt.GoRaw("\tself.timeout = 0"),
				GoStmt.GoRaw("} else {"),
				GoStmt.GoRaw("\tself.hasTimeout = true"),
				GoStmt.GoRaw("\tself.timeout = timeout"),
				GoStmt.GoRaw("}"),
				GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoIdent("self"), "hxrt__socket_applyConnDeadline"), [])),
				GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoIdent("self"), "hxrt__socket_applyListenerDeadline"), []))
			]),
			GoDecl.GoFuncDecl("waitForRead", {
				name: "self",
				typeName: "*sys__net__Socket"
			}, [], [], [
				GoStmt.GoRaw("if self == nil {"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("_ = sys__net__Socket_select_([]*sys__net__Socket{self}, []*sys__net__Socket{}, []*sys__net__Socket{}, -1)")
			]),
			GoDecl.GoFuncDecl("setBlocking", {
				name: "self",
				typeName: "*sys__net__Socket"
			}, [{name: "b", typeName: "bool"}], [], [
				GoStmt.GoRaw("if self == nil {"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("self.blocking = b"),
				GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoIdent("self"), "hxrt__socket_applyConnDeadline"), [])),
				GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoIdent("self"), "hxrt__socket_applyListenerDeadline"), []))
			]),
			GoDecl.GoFuncDecl("setFastSend", {
				name: "self",
				typeName: "*sys__net__Socket"
			}, [{name: "b", typeName: "bool"}], [], [
				GoStmt.GoRaw("if self == nil {"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("self.fastSend = b"),
				GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoIdent("self"), "hxrt__socket_applyFastSend"), []))
			]),
			GoDecl.GoFuncDecl("sys__net__Socket_select_", null, [
				{
					name: "read",
					typeName: "[]*sys__net__Socket"
				},
				{name: "write", typeName: "[]*sys__net__Socket"},
				{name: "others", typeName: "[]*sys__net__Socket"},
				{name: "timeout", typeName: "...float64"}
			], ["map[string]any"], [
				GoStmt.GoIf(GoExpr.GoBinary("==", GoExpr.GoIdent("read"), GoExpr.GoNil),
					[GoStmt.GoAssign(GoExpr.GoIdent("read"), GoExpr.GoRaw("[]*sys__net__Socket{}"))], null),
				GoStmt.GoIf(GoExpr.GoBinary("==", GoExpr.GoIdent("write"), GoExpr.GoNil),
					[GoStmt.GoAssign(GoExpr.GoIdent("write"), GoExpr.GoRaw("[]*sys__net__Socket{}"))], null),
				GoStmt.GoIf(GoExpr.GoBinary("==", GoExpr.GoIdent("others"), GoExpr.GoNil),
					[GoStmt.GoAssign(GoExpr.GoIdent("others"), GoExpr.GoRaw("[]*sys__net__Socket{}"))], null),
				GoStmt.GoRaw("effectiveTimeout := -1.0"),
				GoStmt.GoRaw("if len(timeout) > 0 {"),
				GoStmt.GoRaw("\teffectiveTimeout = timeout[0]"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("readyRead := make([]*sys__net__Socket, 0, len(read))"),
				GoStmt.GoRaw("readyWrite := make([]*sys__net__Socket, 0, len(write))"),
				GoStmt.GoRaw("readyOther := make([]*sys__net__Socket, 0, len(others))"),
				GoStmt.GoRaw("for _, socket := range read {"),
				GoStmt.GoRaw("\tif socket == nil || socket.conn == nil || socket.input == nil || socket.input.reader == nil {"),
				GoStmt.GoRaw("\t\tcontinue"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\treader := socket.input.reader"),
				GoStmt.GoRaw("\tif reader.Buffered() > 0 {"),
				GoStmt.GoRaw("\t\treadyRead = append(readyRead, socket)"),
				GoStmt.GoRaw("\t\tcontinue"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tif effectiveTimeout >= 0 {"),
				GoStmt.GoRaw("\t\tdeadline := time.Now()"),
				GoStmt.GoRaw("\t\tif effectiveTimeout > 0 {"),
				GoStmt.GoRaw("\t\t\tdeadline = time.Now().Add(time.Duration(effectiveTimeout * float64(time.Second)))"),
				GoStmt.GoRaw("\t\t}"),
				GoStmt.GoRaw("\t\t_ = socket.conn.SetReadDeadline(deadline)"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\t_, err := reader.Peek(1)"),
				GoStmt.GoRaw("\tsocket.hxrt__socket_applyConnDeadline()"),
				GoStmt.GoRaw("\tif err == nil {"),
				GoStmt.GoRaw("\t\treadyRead = append(readyRead, socket)"),
				GoStmt.GoRaw("\t\tcontinue"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tif netErr, ok := err.(net.Error); ok && netErr.Timeout() {"),
				GoStmt.GoRaw("\t\tcontinue"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\treadyOther = append(readyOther, socket)"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("for _, socket := range write {"),
				GoStmt.GoRaw("\tif socket == nil || socket.conn == nil {"),
				GoStmt.GoRaw("\t\tcontinue"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\treadyWrite = append(readyWrite, socket)"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("for _, socket := range others {"),
				GoStmt.GoRaw("\tif socket == nil {"),
				GoStmt.GoRaw("\t\tcontinue"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\treadyOther = append(readyOther, socket)"),
				GoStmt.GoRaw("}"),
				GoStmt.GoReturn(GoExpr.GoRaw("map[string]any{\"read\": readyRead, \"write\": readyWrite, \"others\": readyOther}"))
			]),
			GoDecl.GoFuncDecl("readLine", {
				name: "self",
				typeName: "*sys__net__SocketInput"
			}, [], ["*string"], [
				GoStmt.GoRaw("if self == nil || self.reader == nil {"),
				GoStmt.GoRaw("\treturn hxrt.StringFromLiteral(\"\")"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if self.socket != nil {"),
				GoStmt.GoRaw("\tself.socket.hxrt__socket_applyConnDeadline()"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("line, err := self.reader.ReadString('\\n')"),
				GoStmt.GoRaw("if err != nil && len(line) == 0 {"),
				GoStmt.GoRaw("\treturn hxrt.StringFromLiteral(\"\")"),
				GoStmt.GoRaw("}"),
				GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("hxrt.StringFromLiteral"), [
					GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoIdent("strings"), "TrimRight"), [GoExpr.GoIdent("line"), GoExpr.GoStringLiteral("\r\n")])
				]))
			]),
			GoDecl.GoFuncDecl("writeString", {
				name: "self",
				typeName: "*sys__net__SocketOutput"
			}, [{name: "value", typeName: "*string"}], [], [
				GoStmt.GoRaw("if self == nil || self.writer == nil || value == nil {"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if self.socket != nil {"),
				GoStmt.GoRaw("\tself.socket.hxrt__socket_applyConnDeadline()"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if _, err := self.writer.WriteString(*hxrt.StdString(value)); err != nil {"),
				GoStmt.GoRaw("\thxrt.Throw(err)"),
				GoStmt.GoRaw("}")
			]),
			GoDecl.GoFuncDecl("flush", {
				name: "self",
				typeName: "*sys__net__SocketOutput"
			}, [], [], [
				GoStmt.GoRaw("if self == nil || self.writer == nil {"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if self.socket != nil {"),
				GoStmt.GoRaw("\tself.socket.hxrt__socket_applyConnDeadline()"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if err := self.writer.Flush(); err != nil {"),
				GoStmt.GoRaw("\thxrt.Throw(err)"),
				GoStmt.GoRaw("}")
			])
		];
	}
}
#end
