package haxe;

import haxe.Constraints.Function;

using StringTools;

private enum TemplateExpr {
	OpVar(v:String);
	OpExpr(expr:Void->Dynamic);
	OpIf(expr:Void->Dynamic, eif:TemplateExpr, eelse:Null<TemplateExpr>);
	OpStr(str:String);
	OpBlock(items:Array<TemplateExpr>);
	OpForeach(expr:Void->Dynamic, loop:TemplateExpr);
	OpMacro(name:String, params:Array<TemplateExpr>);
}

private typedef Token = {
	var s:Bool;
	var p:String;
	var l:Null<Array<String>>;
}

private typedef ExprToken = {
	var s:Bool;
	var p:String;
}

private class TokenCursor {
	public var tokens:Array<Token>;
	public var index:Int;

	public function new(tokens:Array<Token>) {
		this.tokens = tokens;
		this.index = 0;
	}
}

private class ExprCursor {
	public var tokens:Array<ExprToken>;
	public var index:Int;

	public function new(tokens:Array<ExprToken>) {
		this.tokens = tokens;
		this.index = 0;
	}
}

/**
	What
	A staged `haxe.Template` override for `haxe.go`.

	Why
	The upstream stdlib implementation is a good semantic reference, but it currently
	assumes several target surfaces that `haxe.go` does not expose cleanly through
	source-owned inclusion yet: private `haxe.ds.List` internals, direct string helper
	symbols, and closure shapes that lower too narrowly in Go. That made direct
	`new haxe.Template(...).execute(...)` fail even though the public Template contract
	itself is portable.

	How
	Keep the upstream public API and template syntax model, but adapt the internal
	parser/runtime to use `Array`-backed cursors and explicit helper functions that
	already lower cleanly on `haxe.go`. This keeps ownership in staged std code instead
	of pushing more stdlib semantics into `GoCompiler`.
**/
class Template {
	static var splitter = ~/(::[A-Za-z0-9_ ()&|!+=\/><*."-]+::|\$\$([A-Za-z0-9_-]+)\()/;
	static var expr_splitter = ~/(\(|\)|[ \r\n\t]*"[^"]*"[ \r\n\t]*|[!+=\/><*.&|-]+)/;
	static var expr_trim = ~/^[ ]*([^ ]+)[ ]*$/;
	static var expr_int = ~/^[0-9]+$/;
	static var expr_float = ~/^[+-]?([0-9]+(,[0-9]*)?|,[0-9]+)([Ee][+-]?[0-9]+)?$/;

	public static var globals:Dynamic = {};

	var expr:TemplateExpr;
	var context:Dynamic;
	var macros:Dynamic;
	var stack:Array<Dynamic>;
	var output:String;

	public function new(str:String) {
		var cursor = new TokenCursor(parseTokens(str));
		expr = parseBlock(cursor);
		if (cursor.index < cursor.tokens.length) {
			var token = cursor.tokens[cursor.index];
			throw "Unexpected '" + token.s + "'";
		}
	}

	public function execute(context:Dynamic, ?macros:Dynamic):String {
		this.macros = macros == null ? {} : macros;
		this.context = context;
		stack = [];
		output = "";
		run(expr);
		return output;
	}

	function resolve(v:String):Dynamic {
		if (v == "__current__") {
			return context;
		}
		if (Reflect.isObject(context)) {
			var value = Reflect.getProperty(context, v);
			if (value != null || Reflect.hasField(context, v)) {
				return value;
			}
		}
		for (ctx in stack) {
			var value = Reflect.getProperty(ctx, v);
			if (value != null || Reflect.hasField(ctx, v)) {
				return value;
			}
		}
		return Reflect.field(globals, v);
	}

	function parseTokens(data:String):Array<Token> {
		var tokens:Array<Token> = [];
		while (splitter.match(data)) {
			var p = splitter.matchedPos();
			if (p.pos > 0) {
				tokens.push({p: data.substr(0, p.pos), s: true, l: null});
			}

			if (data.charCodeAt(p.pos) == 58) {
				tokens.push({p: data.substr(p.pos + 2, p.len - 4), s: false, l: null});
				data = splitter.matchedRight();
				continue;
			}

			var parp = p.pos + p.len;
			var npar = 1;
			var params:Array<String> = [];
			var part = "";
			while (true) {
				var c = data.charCodeAt(parp);
				parp++;
				if (c == 40) {
					npar++;
				} else if (c == 41) {
					npar--;
					if (npar <= 0) {
						break;
					}
				} else if (c == null) {
					throw "Unclosed macro parenthesis";
				}
				var chunk = data.substr(parp - 1, 1);
				if (c == 44 && npar == 1) {
					params.push(part);
					part = "";
				} else {
					part += chunk;
				}
			}
			params.push(part);
			tokens.push({p: splitter.matched(2), s: false, l: params});
			data = data.substr(parp, data.length - parp);
		}
		if (data.length > 0) {
			tokens.push({p: data, s: true, l: null});
		}
		return tokens;
	}

	function parseBlock(cursor:TokenCursor):TemplateExpr {
		var items:Array<TemplateExpr> = [];
		while (cursor.index < cursor.tokens.length) {
			var t = cursor.tokens[cursor.index];
			if (!t.s && (t.p == "end" || t.p == "else" || t.p.substr(0, 7) == "elseif ")) {
				break;
			}
			items.push(parse(cursor));
		}
		if (items.length == 1) {
			return items[0];
		}
		return OpBlock(items);
	}

	function parse(cursor:TokenCursor):TemplateExpr {
		var t = popToken(cursor);
		if (t == null) {
			throw "Unexpected <eof>";
		}
		var p = t.p;
		if (t.s) {
			return OpStr(p);
		}
		if (t.l != null) {
			var parsedParams:Array<TemplateExpr> = [];
			for (param in t.l) {
				parsedParams.push(parseBlock(new TokenCursor(parseTokens(param))));
			}
			return OpMacro(p, parsedParams);
		}

		var pos = kwdEnd(p, "if");
		if (pos > 0) {
			p = p.substr(pos, p.length - pos);
			var e = parseExpr(p);
			var eif = parseBlock(cursor);
			var nextToken = peekToken(cursor);
			if (nextToken == null) {
				throw "Unclosed 'if'";
			}
			var eelse:Null<TemplateExpr> = null;
			if (nextToken.p == "end") {
				popToken(cursor);
			} else if (nextToken.p == "else") {
				popToken(cursor);
				eelse = parseBlock(cursor);
				var endToken = popToken(cursor);
				if (endToken == null || endToken.p != "end") {
					throw "Unclosed 'else'";
				}
			} else {
				nextToken.p = nextToken.p.substr(4, nextToken.p.length - 4);
				eelse = parse(cursor);
			}
			return OpIf(e, eif, eelse);
		}

		pos = kwdEnd(p, "foreach");
		if (pos >= 0) {
			p = p.substr(pos, p.length - pos);
			var e = parseExpr(p);
			var efor = parseBlock(cursor);
			var endToken = popToken(cursor);
			if (endToken == null || endToken.p != "end") {
				throw "Unclosed 'foreach'";
			}
			return OpForeach(e, efor);
		}

		if (expr_splitter.match(p)) {
			return OpExpr(parseExpr(p));
		}
		return OpVar(p);
	}

	function parseExpr(data:String):Void->Dynamic {
		var tokens:Array<ExprToken> = [];
		var expr = data;
		while (expr_splitter.match(data)) {
			var p = expr_splitter.matchedPos();
			if (p.pos != 0) {
				tokens.push({p: data.substr(0, p.pos), s: true});
			}
			var token = expr_splitter.matched(0);
			tokens.push({p: token, s: StringTools.contains(token, "\"")});
			data = expr_splitter.matchedRight();
		}
		if (data.length != 0) {
			for (i => c in data) {
				switch (c) {
					case ' '.code:
					case _:
						tokens.push({p: data.substr(i), s: true});
						break;
				}
			}
		}

		var cursor = new ExprCursor(tokens);
		var built:Void->Dynamic;
		try {
			built = makeExpr(cursor);
			if (cursor.index < cursor.tokens.length) {
				throw cursor.tokens[cursor.index].p;
			}
		} catch (s:String) {
			throw "Unexpected '" + s + "' in " + expr;
		}

		var me = this;
		var wrapped:Void->Dynamic = function():Dynamic {
			try {
				return built();
			} catch (exc:Dynamic) {
				// Dynamic catch is intentional: template expressions can call user
				// resolvers/macros that may throw any Haxe value.
				throw "Error : " + Std.string(exc) + " in " + expr;
			}
			return null;
		};
		return wrapped;
	}

	function makeConst(v:String):Void->Dynamic {
		expr_trim.match(v);
		v = expr_trim.matched(1);
		if (v.charCodeAt(0) == 34) {
			var str = v.substr(1, v.length - 2);
			var literal:Void->Dynamic = function():Dynamic return str;
			return literal;
		}
		if (expr_int.match(v)) {
			var i = parseIntLiteral(v);
			var intLiteral:Void->Dynamic = function():Dynamic return i;
			return intLiteral;
		}
		if (expr_float.match(v)) {
			var f = parseFloatLiteral(v);
			var floatLiteral:Void->Dynamic = function():Dynamic return f;
			return floatLiteral;
		}
		var me = this;
		var resolved:Void->Dynamic = function():Dynamic return me.resolve(v);
		return resolved;
	}

	function makePath(e:Void->Dynamic, cursor:ExprCursor):Void->Dynamic {
		var token = peekExprToken(cursor);
		if (token == null || token.p != ".") {
			return e;
		}
		popExprToken(cursor);
		var field = popExprToken(cursor);
		if (field == null || !field.s) {
			throw field == null ? "<eof>" : field.p;
		}
		var name = trimExprToken(field.p);
		return makePath(function():Dynamic {
			return Reflect.field(e(), name);
		}, cursor);
	}

	function makeExpr(cursor:ExprCursor):Void->Dynamic {
		return makePath(makeExpr2(cursor), cursor);
	}

	function skipSpaces(cursor:ExprCursor):Void {
		while (cursor.index < cursor.tokens.length) {
			if (!isSpaceOnly(cursor.tokens[cursor.index].p)) {
				return;
			}
			cursor.index++;
		}
	}

	function makeExpr2(cursor:ExprCursor):Void->Dynamic {
		skipSpaces(cursor);
		var token = popExprToken(cursor);
		skipSpaces(cursor);
		if (token == null) {
			throw "<eof>";
		}
		if (token.s) {
			return makeConst(token.p);
		}

		switch (token.p) {
			case "(":
				skipSpaces(cursor);
				var e1 = makeExpr(cursor);
				skipSpaces(cursor);
				var op = popExprToken(cursor);
				if (op == null || op.s) {
					throw op == null ? "<eof>" : op.p;
				}
				if (op.p == ")") {
					return e1;
				}
				skipSpaces(cursor);
				var e2 = makeExpr(cursor);
				skipSpaces(cursor);
				var close = popExprToken(cursor);
				skipSpaces(cursor);
				if (close == null || close.p != ")") {
					throw close == null ? "<eof>" : close.p;
				}
				return switch (op.p) {
					case "+":
						function():Dynamic return addValues(e1(), e2());
					case "-":
						function():Dynamic return subtractValues(e1(), e2());
					case "*":
						function():Dynamic return multiplyValues(e1(), e2());
					case "/":
						function():Dynamic return divideValues(e1(), e2());
					case ">":
						function():Dynamic return compareValues(e1(), e2()) > 0;
					case "<":
						function():Dynamic return compareValues(e1(), e2()) < 0;
					case ">=":
						function():Dynamic return compareValues(e1(), e2()) >= 0;
					case "<=":
						function():Dynamic return compareValues(e1(), e2()) <= 0;
					case "==":
						function():Dynamic return e1() == e2();
					case "!=":
						function():Dynamic return e1() != e2();
					case "&&":
						function():Dynamic return valueAsBool(e1()) && valueAsBool(e2());
					case "||":
						function():Dynamic return valueAsBool(e1()) || valueAsBool(e2());
					case _:
						throw "Unknown operation " + op.p;
				}
			case "!":
				var inner = makeExpr(cursor);
				return function():Dynamic {
					var value:Dynamic = inner();
					return value == null || value == false;
				};
			case "-":
				var inner = makeExpr(cursor);
				return function():Dynamic return -valueAsFloat(inner());
			case _:
				throw token.p;
		}
	}

	function run(e:TemplateExpr):Void {
		switch (e) {
			case OpVar(v):
				output += Std.string(resolve(v));
			case OpExpr(expr):
				output += Std.string(expr());
			case OpIf(expr, ifExpr, elseExpr):
				var value:Dynamic = expr();
				if (value == null || value == false) {
					if (elseExpr != null) {
						run(elseExpr);
					}
				} else {
					run(ifExpr);
				}
			case OpStr(str):
				output += str;
			case OpBlock(items):
				for (item in items) {
					run(item);
				}
			case OpForeach(expr, loop):
				var value:Dynamic = expr();
				var arrayValues = anyArrayToSlice(value);
				if (arrayValues != null) {
					stack.push(context);
					for (ctx in arrayValues) {
						context = ctx;
						run(loop);
					}
					context = popStackValue();
					return;
				}
				var iterator:Dynamic = null;
				try {
					var iteratorField:Dynamic = Reflect.field(value, "iterator");
					if (iteratorField == null) {
						throw null;
					}
					var candidate:Dynamic = Reflect.callMethod(value, cast iteratorField, []);
					if (!Reflect.hasField(candidate, "hasNext")) {
						throw null;
					}
					iterator = candidate;
				} catch (_:Dynamic) {
					// Dynamic catch is intentional: reflective iterator probing must tolerate
					// arbitrary failures before trying the direct iterator shape.
					try {
						if (value == null || !Reflect.hasField(value, "hasNext")) {
							throw null;
						}
						iterator = value;
					} catch (_:Dynamic) {
						// Dynamic catch is intentional: this is the second reflective probe, and
						// failure becomes the template's public "Cannot iter" error.
						throw "Cannot iter on " + value;
					}
				}
				stack.push(context);
				var iterable:Iterator<Dynamic> = cast iterator;
				for (ctx in iterable) {
					context = ctx;
					run(loop);
				}
				context = popStackValue();
			case OpMacro(name, params):
				var fn:Dynamic = Reflect.field(macros, name);
				var callArgs:Array<Dynamic> = [];
				callArgs.push(resolve);
				for (param in params) {
					switch (param) {
						case OpVar(value):
							callArgs.push(resolve(value));
						case _:
							var previous = output;
							output = "";
							run(param);
							callArgs.push(output);
							output = previous;
					}
				}
				try {
					output += Std.string(Reflect.callMethod(macros, cast fn, callArgs));
				} catch (err:Dynamic) {
					// Dynamic catch is intentional: template macros are user callbacks and
					// may throw arbitrary Haxe values.
					// Dynamic catch is intentional: argument formatting is best-effort only.
					var argsText = try joinDynamicArgs(callArgs) catch (_:Dynamic) "???";
					throw "Macro call " + name + "(" + argsText + ") failed (" + Std.string(err) + ")";
				}
		}
	}

	static inline function peekToken(cursor:TokenCursor):Null<Token> {
		return cursor.index < cursor.tokens.length ? cursor.tokens[cursor.index] : null;
	}

	static inline function popToken(cursor:TokenCursor):Null<Token> {
		if (cursor.index >= cursor.tokens.length) {
			return null;
		}
		return cursor.tokens[cursor.index++];
	}

	static inline function peekExprToken(cursor:ExprCursor):Null<ExprToken> {
		return cursor.index < cursor.tokens.length ? cursor.tokens[cursor.index] : null;
	}

	static inline function popExprToken(cursor:ExprCursor):Null<ExprToken> {
		if (cursor.index >= cursor.tokens.length) {
			return null;
		}
		return cursor.tokens[cursor.index++];
	}

	static function kwdEnd(value:String, keyword:String):Int {
		var pos = -1;
		var length = keyword.length;
		if (value.substr(0, length) == keyword) {
			pos = length;
			for (code in value.substr(length)) {
				switch (code) {
					case ' '.code:
						pos++;
					case _:
						break;
				}
			}
		}
		return pos;
	}

	static function trimExprToken(value:String):String {
		expr_trim.match(value);
		return expr_trim.matched(1);
	}

	static function isSpaceOnly(value:String):Bool {
		for (code in value) {
			if (code != " ".code) {
				return false;
			}
		}
		return true;
	}

	static function parseIntLiteral(value:String):Int {
		var out = 0;
		var index = 0;
		while (index < value.length) {
			var code:Int = value.charCodeAt(index);
			out = out * 10 + (code - "0".code);
			index++;
		}
		return out;
	}

	static function parseFloatLiteral(value:String):Float {
		var normalized = value.replace(",", ".");
		var index = 0;
		var sign = 1.0;
		if (normalized.charCodeAt(index) == "-".code) {
			sign = -1.0;
			index++;
		} else if (normalized.charCodeAt(index) == "+".code) {
			index++;
		}

		var intPart = 0.0;
		while (index < normalized.length) {
			var code:Int = normalized.charCodeAt(index);
			if (code < "0".code || code > "9".code) {
				break;
			}
			intPart = intPart * 10.0 + (code - "0".code);
			index++;
		}

		var fracPart = 0.0;
		var divisor = 1.0;
		if (index < normalized.length && normalized.charCodeAt(index) == ".".code) {
			index++;
			while (index < normalized.length) {
				var code:Int = normalized.charCodeAt(index);
				if (code < "0".code || code > "9".code) {
					break;
				}
				fracPart = fracPart * 10.0 + (code - "0".code);
				divisor *= 10.0;
				index++;
			}
		}

		var result = intPart + fracPart / divisor;
		if (index < normalized.length) {
			var exponentCode = normalized.charCodeAt(index);
			if (exponentCode == "e".code || exponentCode == "E".code) {
				index++;
				var exponentSign = 1;
				if (normalized.charCodeAt(index) == "-".code) {
					exponentSign = -1;
					index++;
				} else if (normalized.charCodeAt(index) == "+".code) {
					index++;
				}
				var exponent = 0;
				while (index < normalized.length) {
					var code:Int = normalized.charCodeAt(index);
					if (code < "0".code || code > "9".code) {
						break;
					}
					exponent = exponent * 10 + (code - "0".code);
					index++;
				}
				while (exponent > 0) {
					result = exponentSign < 0 ? result / 10.0 : result * 10.0;
					exponent--;
				}
			}
		}
		return sign * result;
	}

	static function valueAsBool(value:Dynamic):Bool {
		return !(value == null || value == false);
	}

	static function valueAsFloat(value:Dynamic):Float {
		if (Std.isOfType(value, Int) || Std.isOfType(value, Float)) {
			return cast value;
		}
		if (Std.isOfType(value, String)) {
			return parseFloatLiteral(cast value);
		}
		throw "Expected numeric expression value, got " + Std.string(value);
	}

	static function compareValues(left:Dynamic, right:Dynamic):Int {
		var leftNumeric = Std.isOfType(left, Int) || Std.isOfType(left, Float);
		var rightNumeric = Std.isOfType(right, Int) || Std.isOfType(right, Float);
		if (leftNumeric && rightNumeric) {
			var leftFloat = valueAsFloat(left);
			var rightFloat = valueAsFloat(right);
			return leftFloat < rightFloat ? -1 : leftFloat > rightFloat ? 1 : 0;
		}
		return Reflect.compare(Std.string(left), Std.string(right));
	}

	static function addValues(left:Dynamic, right:Dynamic):Dynamic {
		if (Std.isOfType(left, String) || Std.isOfType(right, String)) {
			return Std.string(left) + Std.string(right);
		}
		return valueAsFloat(left) + valueAsFloat(right);
	}

	static inline function subtractValues(left:Dynamic, right:Dynamic):Float {
		return valueAsFloat(left) - valueAsFloat(right);
	}

	static inline function multiplyValues(left:Dynamic, right:Dynamic):Float {
		return valueAsFloat(left) * valueAsFloat(right);
	}

	static inline function divideValues(left:Dynamic, right:Dynamic):Float {
		return valueAsFloat(left) / valueAsFloat(right);
	}

	function popStackValue():Dynamic {
		var lastIndex = stack.length - 1;
		var value = stack[lastIndex];
		var remaining:Array<Dynamic> = [];
		for (index in 0...lastIndex) {
			remaining.push(stack[index]);
		}
		stack = remaining;
		return value;
	}

	static function joinDynamicArgs(values:Array<Dynamic>):String {
		var out = "";
		for (index in 0...values.length) {
			if (index > 0) {
				out += ",";
			}
			out += Std.string(values[index]);
		}
		return out;
	}

	static function anyArrayToSlice(value:Dynamic):Null<Array<Dynamic>> {
		return untyped __go__("haxe__Template_anyArrayToSlice_runtime({0})", value);
	}
}
