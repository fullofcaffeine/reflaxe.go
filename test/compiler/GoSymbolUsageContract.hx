import reflaxe.go.ast.GoAST.GoDecl;
import reflaxe.go.ast.GoAST.GoExpr;
import reflaxe.go.ast.GoAST.GoSelectClause;
import reflaxe.go.ast.GoAST.GoStmt;
import reflaxe.go.ast.GoBinaryOperator;
import reflaxe.go.ast.GoCompositeElement;
import reflaxe.go.ast.GoIncDecOperator;
import reflaxe.go.ast.GoSimpleStmt;
import reflaxe.go.ast.GoSymbolUsage;
import reflaxe.go.ast.GoType;
import reflaxe.go.ast.GoUnaryOperator;

/**
	What: Pins which package names and `package.Member` symbols
	`GoSymbolUsage` finds in Go declarations.

	Why: The compiler drops every candidate import that a file does not use and
	turns on runtime features from the `hxrt.*` symbols a file names. A
	different answer changes generated Go. These cases were checked against
	the earlier per-name walk (one tree walk and one regular expression per
	candidate) before the single walk replaced it.

	How:
	- Fixed cases cover selectors, identifier text, raw Go text (comments and
	  strings inside it count), string literals (they do not), types, short
	  declarations, word boundaries, spacing, and every AST position.
	- A seeded random sweep compares each raw text and identifier with the
	  earlier regular expressions `\bname\s*\.` and `\bpkg\s*\.\s*Member\b`.
	- Keys that are not `name` or `name.Member` are rejected.
**/
class GoSymbolUsageContract {
	static final failures = new Array<String>();

	static function main():Void {
		selectorsAndIdentifiers();
		literalsAndRawText();
		wordBoundaries();
		types();
		everyPositionIsRead();
		declaredNamesAreNotRead();
		randomTextMatchesRegularExpressions();
		invalidKeysAreRejected();
		if (failures.length > 0)
			throw failures.length + " symbol-usage failures:\n" + failures.join("\n");
		Sys.println("ok");
	}

	static function selectorsAndIdentifiers():Void {
		check("package selector", exprs([GoCall(GoSelector(GoIdent("fmt"), "Println"), [GoIdent("x")])]), ["fmt", "fmt.Println"],
			["Println", "x", "fmt.Printf"]);
		// A lone identifier can be a local variable named like a package.
		check("lone identifier", exprs([GoIdent("fmt")]), [], ["fmt"]);
		check("qualified identifier text", exprs([GoCall(GoIdent("hxrt.ThreadWaitForAll"), [])]), ["hxrt", "hxrt.ThreadWaitForAll"], ["ThreadWaitForAll"]);
		check("selector on a call", exprs([GoSelector(GoCall(GoIdent("os"), []), "Exit")]), [], ["os", "os.Exit", "Exit"]);
		check("nested selector", exprs([GoSelector(GoSelector(GoIdent("a"), "b"), "c")]), ["a", "a.b"], ["b", "c", "b.c", "a.c"]);
		check("selector on qualified identifier", exprs([GoSelector(GoIdent("hxrt.Function"), "Call")]), ["hxrt", "hxrt.Function"],
			["Function", "Call", "Function.Call"]);
	}

	static function literalsAndRawText():Void {
		check("string literal", exprs([GoStringLiteral("os.Exit"), GoStringLiteral("strings . Builder")]), [], ["os", "os.Exit", "strings"]);
		check("float literal", exprs([GoFloatLiteral("1.5")]), [], ["1", "1.5"]);
		// Raw Go is read as text: comments and string contents count.
		check("comment in raw statement", body([GoRaw("// see json.Marshal")]), ["json", "json.Marshal"], ["see"]);
		check("string in raw expression", exprs([GoRaw('name := "bytes.Buffer"')]), ["bytes", "bytes.Buffer"], ["name"]);
		check("number in raw text", body([GoRaw("v := 1.5")]), ["1", "1.5"], ["v"]);
		check("chained raw text", body([GoRaw("a.b.c")]), ["a", "b", "a.b", "b.c"], ["c", "a.c"]);
		check("generic raw text", body([GoRaw("hxrt.Function[int]")]), ["hxrt", "hxrt.Function"], ["Function", "int"]);
		check("empty raw text", body([GoRaw(""), GoRaw(".")]), [], ["hxrt"]);
	}

	static function wordBoundaries():Void {
		check("prefixes and suffixes", body([GoRaw("myfmt.A(); fmt2.B(); xfmt .C; _fmt.D")]),
			["myfmt", "fmt2", "xfmt", "_fmt", "myfmt.A", "fmt2.B", "xfmt.C", "_fmt.D"], ["fmt", "fmt.A", "fmt.B", "fmt.C", "fmt.D"]);
		check("longer member", body([GoRaw("hxrt.ArraySortBy(xs)")]), ["hxrt", "hxrt.ArraySortBy"], ["hxrt.ArraySort"]);
		check("spaces around dot", body([GoRaw("hxrt . ArraySort(xs)")]), ["hxrt", "hxrt.ArraySort"], ["ArraySort"]);
		check("line break after dot", body([GoRaw("hxrt.\n\tArraySort(xs)")]), ["hxrt", "hxrt.ArraySort"], []);
		check("two dots", body([GoRaw("hxrt..ArraySort")]), ["hxrt"], ["hxrt.ArraySort", "ArraySort"]);
		check("underscore prefix", body([GoRaw("_hxrt.Function")]), ["_hxrt", "_hxrt.Function"], ["hxrt", "hxrt.Function"]);
	}

	static function types():Void {
		check("global type", [GoGlobalVarDecl("buffer", GoType.parse("*bytes.Buffer"), null)], ["bytes", "bytes.Buffer"], ["Buffer", "buffer"]);
		check("typed declaration", body([GoVarDecl("b", GoType.parse("*bytes.Buffer"), GoIdent("nil"), false)]), ["bytes", "bytes.Buffer"], []);
		// A short declaration prints `b := value`, so its type names no package.
		check("short declaration", body([
			GoVarDecl("b", GoType.parse("*bytes.Buffer"), GoCall(GoIdent("newBuffer"), []), true)
		]), [], ["bytes", "bytes.Buffer"]);
		check("map field type", [
			GoStructDecl("Guard", [{name: "locks", typeName: GoType.parse("map[string]sync.Mutex")}])
		], ["sync", "sync.Mutex"], ["string", "Mutex"]);
		check("generic function type", [
			GoFuncDecl("f", null, [{name: "fn", typeName: GoType.parse("*hxrt.Function[func(*io.Reader) error]")}], [], [])
		], ["hxrt", "hxrt.Function", "io", "io.Reader"], ["error", "Function", "Reader"]);
		check("interface method", [
			GoInterfaceDecl("Closer", [
				{
					name: "Close",
					params: [{name: "ctx", typeName: GoType.parse("context.Context")}],
					results: [GoType.parse("error")]
				}
			])
		], ["context", "context.Context"], ["Close", "ctx", "error"]);
	}

	/** Each AST position names its own package, so a skipped position shows up by name. */
	static function everyPositionIsRead():Void {
		final selectorPackages = new Array<String>();
		final typePackages = new Array<String>();
		function ref(name:String):GoExpr {
			selectorPackages.push(name);
			return GoSelector(GoIdent(name), "X");
		}
		function type(name:String):GoType {
			typePackages.push(name);
			return GoType.parse("*" + name + ".T");
		}
		final statements:Array<GoStmt> = [
			GoVarDecl("v", type("varType"), ref("varValue"), false),
			GoMultiAssign(["a", "b"], ref("multiAssign"), true),
			GoAssign(ref("assignLeft"), ref("assignRight")),
			GoIncDec(ref("incDec"), GoIncDecOperator.Increment),
			GoExprStmt(ref("exprStmt")),
			GoGoStmt(ref("goStmt")),
			GoDeferStmt(ref("deferStmt")),
			GoSendStmt(ref("sendChannel"), ref("sendValue")),
			GoRaw("rawStmt.X()"),
			GoWhile(ref("whileCond"), [GoExprStmt(ref("whileBody"))]),
			GoFor(GoSimpleDeclare("i", ref("forInit")), ref("forCond"), GoSimpleIncDec(ref("forPost"), GoIncDecOperator.Increment),
				[GoExprStmt(ref("forBody"))]),
			GoFor(GoSimpleAssign(ref("simpleAssignLeft"), ref("simpleAssignRight")), null, GoSimpleExpr(ref("simpleExpr")), []),
			GoFor(GoSimpleSend(ref("simpleSendChannel"), ref("simpleSendValue")), null, null, []),
			GoLabeled("outer", GoExprStmt(ref("labeled"))),
			GoRangeStmt("k", "v", ref("rangeSource"), true, [GoExprStmt(ref("rangeBody"))]),
			GoIf(ref("ifCond"), [GoExprStmt(ref("ifThen"))], [GoExprStmt(ref("ifElse"))]),
			GoSwitch(ref("switchValue"), [
				{
					values: [ref("caseValue")],
					body: [GoExprStmt(ref("caseBody"))]
				}
			],
				[GoExprStmt(ref("switchDefault"))]),
			GoTypeSwitch(ref("typeSwitchValue"), "bound", [{typeName: type("typeSwitchCase"), body: [GoExprStmt(ref("typeSwitchBody"))]}],
				[GoExprStmt(ref("typeSwitchDefault"))]),
			GoSelect([
				{clause: GoSelectSend(ref("selectSendChannel"), ref("selectSendValue")), body: [GoExprStmt(ref("selectBody"))]},
				{clause: GoSelectRecv(ref("selectRecv")), body: []},
				{clause: GoSelectRecvAssign(ref("selectAssignTarget"), ref("selectAssignRecv"), true), body: []},
				{clause: GoSelectRecvAssignOk(ref("selectOkTarget"), ref("selectOkFlag"), ref("selectOkRecv"), false), body: []},
				{clause: GoSelectDefault, body: [GoExprStmt(ref("selectDefaultBody"))]}
			]),
			GoReturnValues([ref("returnFirst"), ref("returnSecond")]),
			GoReturn(GoCall(ref("callCallee"), [
				GoInstantiate(ref("instantiateTarget"), [type("instantiateArgument")]),
				GoIndex(ref("indexTarget"), ref("indexValue")),
				GoSlice(ref("sliceTarget"), ref("sliceStart"), ref("sliceEnd")),
				GoCompositeLiteral(type("compositeType"), [
					GoCompositeValue(ref("compositeValue")),
					GoCompositeKeyValue(ref("compositeKey"), ref("compositeKeyedValue")),
					GoCompositeField("Field", ref("compositeField"))
				]),
				GoMakeSlice(type("makeElement"), ref("makeLength"), ref("makeCapacity")),
				GoFuncLiteral([
					{
						name: "p",
						typeName: type("literalParam")
					}
				],
					[type("literalResult")], [GoExprStmt(ref("literalBody"))]),
				GoRaw("rawExpr.X"),
				GoTypeAssert(ref("assertValue"), type("assertType")),
				GoRecvExpr(ref("recvChannel")),
				GoUnary(GoUnaryOperator.LogicalNot, ref("unaryValue")),
				GoBinary(GoBinaryOperator.Add, ref("binaryLeft"), ref("binaryRight")),
				GoVariadicCall(ref("variadicCallee"), [ref("variadicArgument")])
			]))
		];
		final decls:Array<GoDecl> = [
			GoInterfaceDecl("I",
				[
					{name: "M", params: [{name: "p", typeName: type("interfaceParam")}], results: [type("interfaceResult")]}
				]),
			GoStructDecl("S", [{name: "f", typeName: type("structField")}]),
			GoGlobalVarDecl("g", type("globalType"), ref("globalValue")),
			GoFuncDecl("F", {name: "self", typeName: type("receiverType")}, [{name: "p", typeName: type("paramType")}], [type("resultType")], statements)
		];
		final used = ["rawStmt", "rawStmt.X", "rawExpr", "rawExpr.X"];
		for (name in selectorPackages) {
			used.push(name);
			used.push(name + ".X");
		}
		for (name in typePackages) {
			used.push(name);
			used.push(name + ".T");
		}
		check("every position", decls, used, ["X", "T"]);
	}

	/** Declared names, labels and field names are never read, even with a dot inside. */
	static function declaredNamesAreNotRead():Void {
		final decls:Array<GoDecl> = [
			GoInterfaceDecl("interfaceName.X",
				[
					{name: "methodName.X", params: [{name: "methodParam.x", typeName: GoType.parse("int")}], results: []}
				]),
			GoStructDecl("structName.X", [{name: "fieldName.x", typeName: GoType.parse("int")}]),
			GoGlobalVarDecl("globalName.x", GoType.parse("int"), null),
			GoFuncDecl("funcName.X", {name: "receiverName.x", typeName: GoType.parse("*T")}, [{name: "paramName.x", typeName: GoType.parse("int")}], [], [
				GoVarDecl("varName.x", null, GoIntLiteral(1), true),
				GoMultiAssign(["multiName.a", "multiName.b"], GoIdent("pair"), true),
				GoLabeled("labelName.x", GoBreak("breakLabel.x")),
				GoRangeStmt("rangeKey.k", "rangeValue.v", GoIdent("items"), true, [GoContinue]),
				GoTypeSwitch(GoIdent("value"), "bindingName.x", [], null),
				GoFor(GoSimpleDeclare("simpleName.i", GoIntLiteral(0)), null, null, []),
				GoExprStmt(GoCompositeLiteral(GoType.parse("T"), [GoCompositeField("compositeField.x", GoIntLiteral(1))])),
				GoExprStmt(GoSelector(GoIdent("obj"), "selectorField.x")),
				GoExprStmt(GoStringLiteral("stringValue.x")),
				GoReturn(null)
			])
		];
		check("declared names", decls, ["obj"], [
			"interfaceName",
			"methodName",
			"methodParam",
			"structName",
			"fieldName",
			"globalName",
			"funcName",
			"receiverName",
			"paramName",
			"varName",
			"multiName",
			"labelName",
			"breakLabel",
			"rangeKey",
			"rangeValue",
			"bindingName",
			"simpleName",
			"compositeField",
			"selectorField",
			"stringValue",
			"pair",
			"items",
			"value"
		]);
	}

	/**
		Compares raw-text and identifier scanning with the regular expressions the
		earlier walk ran for each name. Seeded random text joins short words with
		dots, spaced dots, line breaks, double dots, quotes, comment slashes, a
		non-ASCII letter, or nothing (so words run together).
	**/
	static function randomTextMatchesRegularExpressions():Void {
		final words = ["a", "b", "ab", "ba", "a_b", "_a", "_", "1", "a1", "h", "x", "\u00e9", "x\u00e9"];
		final separators = [".", ".", ".", " . ", ".\n\t", " ", "..", "", "\"", "/", "(", "\u00e9", "-"];
		final keys = [
			"a", "b", "h", "ab", "ba", "a_b", "_a", "1", "a1", "x", "a.b", "b.a", "a.a", "ab.b", "a.1", "1.a", "h.x", "_.a", "a._", "x.ab"
		];
		final patterns = [for (key in keys) earlierPattern(key)];
		final usedCounts = [for (_ in keys) 0];
		var seed = 20261120;
		function next(bound:Int):Int {
			seed = (seed * 1103515245 + 12345) & 0x7fffffff;
			return (seed >> 8) % bound;
		}
		final texts = 10000;
		for (_ in 0...texts) {
			final text = new StringBuf();
			if (next(3) == 0)
				text.add(separators[next(separators.length)]);
			for (index in 0...next(6)) {
				if (index > 0)
					text.add(separators[next(separators.length)]);
				text.add(words[next(words.length)]);
			}
			if (next(3) == 0)
				text.add(separators[next(separators.length)]);
			final code = text.toString();
			final raw = GoSymbolUsage.collect(body([GoRaw(code)]));
			final ident = GoSymbolUsage.collect(exprs([GoIdent(code)]));
			for (index in 0...keys.length) {
				final expected = patterns[index].match(code);
				if (expected)
					usedCounts[index]++;
				if (raw.uses(keys[index]) != expected || ident.uses(keys[index]) != expected)
					fail('random text ${haxe.Json.stringify(code)}: ${keys[index]} should be ${expected ? "used" : "unused"}');
			}
		}
		// Keep the sweep meaningful: every key must be found in some texts and missed in others.
		for (index in 0...keys.length)
			if (usedCounts[index] < 20 || usedCounts[index] > texts - 20)
				fail("random sweep found " + keys[index] + " in " + usedCounts[index] + " of " + texts + " texts");
	}

	static function invalidKeysAreRejected():Void {
		final usage = GoSymbolUsage.collect(body([GoRaw("a.b.c")]));
		for (key in ["", ".", "a.", ".b", "a.b.c", "a-b", "a b", "é", "a.é", null]) {
			var rejected = false;
			try
				usage.uses(key)
			catch (_:haxe.Exception)
				rejected = true;
			if (!rejected)
				fail("invalid key " + haxe.Json.stringify(key) + " was accepted");
		}
	}

	/** The regular expression the earlier per-name walk ran for one key. */
	static function earlierPattern(key:String):EReg {
		final separator = key.indexOf(".");
		if (separator < 0)
			return new EReg("\\b" + EReg.escape(key) + "\\s*\\.", "");
		return new EReg("\\b" + EReg.escape(key.substr(0, separator)) + "\\s*\\.\\s*" + EReg.escape(key.substr(separator + 1)) + "\\b", "");
	}

	static function check(label:String, decls:Array<GoDecl>, used:Array<String>, unused:Array<String>):Void {
		final usage = GoSymbolUsage.collect(decls);
		for (key in used)
			if (!usage.uses(key))
				fail(label + ": " + key + " should be used");
		for (key in unused)
			if (usage.uses(key))
				fail(label + ": " + key + " should be unused");
	}

	static function body(statements:Array<GoStmt>):Array<GoDecl> {
		return [GoFuncDecl("f", null, [], [], statements)];
	}

	static function exprs(values:Array<GoExpr>):Array<GoDecl> {
		return body([for (value in values) GoExprStmt(value)]);
	}

	static function fail(message:String):Void {
		failures.push(message);
	}
}
