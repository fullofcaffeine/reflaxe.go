package reflaxe.go.ast;

/**
	What: One structural method signature inside a Go interface type.
	Why: Keeping parameter/results as `GoType` lets import traversal and validation
	reach through method sets.
	How: Prefer `GoType.interfaceMethod`, which copies and validates both lists.
**/
typedef GoTypeInterfaceMethod = {
	final name:String;
	final params:Array<GoType>;
	final results:Array<GoType>;
}

private enum GoTypeKind {
	Builtin(value:GoBuiltinType);
	Named(packageName:Null<GoPackageName>, name:String);
	Pointer(element:GoType);
	Slice(element:GoType);
	ArrayType(length:Int, element:GoType);
	MapType(key:GoType, value:GoType);
	Channel(direction:GoChannelDirection, element:GoType);
	FunctionType(params:Array<GoType>, result:Null<GoType>);
	InterfaceType(methods:Array<GoTypeInterfaceMethod>);
	GenericType(base:GoType, arguments:Array<GoType>);
	MultiResult(results:Array<GoType>);
	Variadic(element:GoType);
}

/**
	What: A structural Go target type, distinct from Haxe source types and portable
	representation-admission policy.

	Why: Rendered type strings hid package use, invalid nesting, variadics, function
	results, and generic arguments from validation and transforms.

	How: Builders create algebraic nodes directly. During incremental migration,
	legacy compiler strings cross one `@:from` parser that either produces the same
	structure or fails before the file printer runs. No unchecked raw-type variant is
	provided.
**/
abstract GoType(GoTypeKind) {
	private inline function new(kind:GoTypeKind) {
		this = kind;
	}

	/** Build one predeclared Go type from the closed builtin set. */
	public static inline function builtin(value:GoBuiltinType):GoType {
		if (value == null) {
			throw "Invalid Go type: builtin cannot be null";
		}
		return new GoType(Builtin(value));
	}

	/** Build an unqualified normalized Go type name. */
	public static function named(name:String):GoType {
		validateIdentifier(name, "type");
		return new GoType(Named(null, name));
	}

	/** Build a package-qualified type without storing `package.Type` as text. */
	public static function qualified(packageName:GoPackageName, name:String):GoType {
		if (packageName == null) {
			throw "Invalid Go type: package qualifier cannot be null";
		}
		validateIdentifier(name, "type");
		return new GoType(Named(packageName, name));
	}

	/** Build a pointer whose element is an ordinary single type. */
	public static function pointer(element:GoType):GoType {
		validateNested(element, "pointer element");
		return new GoType(Pointer(element));
	}

	/** Build a slice whose element is an ordinary single type. */
	public static function slice(element:GoType):GoType {
		validateNested(element, "slice element");
		return new GoType(Slice(element));
	}

	/** Build a fixed-length array, rejecting negative lengths. */
	public static function array(length:Int, element:GoType):GoType {
		if (length < 0) {
			throw "Invalid Go type: array length cannot be negative";
		}
		validateNested(element, "array element");
		return new GoType(ArrayType(length, element));
	}

	/** Build a map and reject structurally non-comparable key types. */
	public static function map(key:GoType, value:GoType):GoType {
		validateNested(key, "map key");
		validateNested(value, "map value");
		if (!isComparableMapKey(key)) {
			throw 'Invalid Go type: map key "' + key.render() + '" is not comparable';
		}
		return new GoType(MapType(key, value));
	}

	/** Build a channel with an explicit send/receive direction contract. */
	public static function channel(direction:GoChannelDirection, element:GoType):GoType {
		if (direction == null) {
			throw "Invalid Go type: channel direction cannot be null";
		}
		validateNested(element, "channel element");
		return new GoType(Channel(direction, element));
	}

	/**
		Build a function type, enforcing final-position variadics and structural
		multi-results.
	**/
	public static function functionType(params:Array<GoType>, results:Array<GoType>):GoType {
		var copiedParams = params == null ? [] : params.copy();
		for (index in 0...copiedParams.length) {
			var param = copiedParams[index];
			if (param == null) {
				throw "Invalid Go type: function parameter cannot be null";
			}
			switch (param.kind()) {
				case Variadic(_):
					if (index != copiedParams.length - 1) {
						throw "Invalid Go type: variadic parameter must be last";
					}
				case MultiResult(_):
					throw "Invalid Go type: multi-result cannot be a function parameter";
				case _:
			}
		}

		var copiedResults = results == null ? [] : results.copy();
		for (result in copiedResults) {
			validateNested(result, "function result");
		}
		var resultType:Null<GoType> = switch (copiedResults.length) {
			case 0: null;
			case 1: copiedResults[0];
			case _: multiResult(copiedResults);
		};
		return new GoType(FunctionType(copiedParams, resultType));
	}

	/** Build one validated method signature for an interface type. */
	public static function interfaceMethod(name:String, params:Array<GoType>, results:Array<GoType>):GoTypeInterfaceMethod {
		validateIdentifier(name, "interface method");
		var signature = functionType(params, results);
		return switch (signature.kind()) {
			case FunctionType(copiedParams, result):
				{
					name: name,
					params: copiedParams,
					results: flattenResult(result)
				};
			case _: throw "Invalid Go type: internal interface method signature";
		};
	}

	/** Build an interface and reject null or duplicate method declarations. */
	public static function interfaceType(methods:Array<GoTypeInterfaceMethod>):GoType {
		var copied = new Array<GoTypeInterfaceMethod>();
		var seen = new Map<String, Bool>();
		for (method in (methods == null ? [] : methods)) {
			if (method == null) {
				throw "Invalid Go type: interface method cannot be null";
			}
			validateIdentifier(method.name, "interface method");
			if (seen.exists(method.name)) {
				throw 'Invalid Go type: duplicate interface method "' + method.name + '"';
			}
			seen.set(method.name, true);
			copied.push(interfaceMethod(method.name, method.params, method.results));
		}
		return new GoType(InterfaceType(copied));
	}

	/** Build Go's empty interface spelling. */
	public static inline function emptyInterface():GoType {
		return new GoType(InterfaceType([]));
	}

	/** Build a generic instantiation over a named base type. */
	public static function generic(base:GoType, arguments:Array<GoType>):GoType {
		if (base == null) {
			throw "Invalid Go type: generic base cannot be null";
		}
		switch (base.kind()) {
			case Named(_, _):
			case _:
				throw 'Invalid Go type: generic base must be named, got "' + base.render() + '"';
		}
		if (arguments == null || arguments.length == 0) {
			throw "Invalid Go type: generic instantiation requires arguments";
		}
		var copied = arguments.copy();
		for (argument in copied) {
			validateNested(argument, "generic argument");
		}
		return new GoType(GenericType(base, copied));
	}

	/** Build the parenthesized result list used by multi-result function types. */
	public static function multiResult(results:Array<GoType>):GoType {
		if (results == null || results.length < 2) {
			throw "Invalid Go type: multi-result requires at least two result types";
		}
		var copied = results.copy();
		for (result in copied) {
			validateNested(result, "multi-result element");
		}
		return new GoType(MultiResult(copied));
	}

	/** Build the variadic marker admitted only as a final function parameter. */
	public static function variadic(element:GoType):GoType {
		validateNested(element, "variadic element");
		return new GoType(Variadic(element));
	}

	/**
		Parse a legacy compiler-produced type spelling through the same structural
		validation used by direct builders. There is intentionally no raw-type node.
	**/
	@:from
	public static function parse(source:String):GoType {
		if (source == null) {
			throw "Invalid Go type: value cannot be null";
		}
		return new GoTypeParser(source).parse();
	}

	/** Render canonical Go type syntax; only printers and diagnostics should need it. */
	public function render():String {
		return switch (this) {
			case Builtin(value): value.token();
			case Named(packageName, name): packageName == null ? name : packageName.value() + "." + name;
			case Pointer(element): "*" + element.render();
			case Slice(element): "[]" + element.render();
			case ArrayType(length, element): "[" + length + "]" + element.render();
			case MapType(key, value): "map[" + key.render() + "]" + value.render();
			case Channel(direction, element):
				switch (direction) {
					case Bidirectional: "chan " + element.render();
					case ReceiveOnly: "<-chan " + element.render();
					case SendOnly: "chan<- " + element.render();
				};
			case FunctionType(params, result):
				var suffix = result == null ? "" : " " + result.render();
				"func(" + [for (param in params) param.render()].join(", ") + ")" + suffix;
			case InterfaceType(methods):
				var rendered = [for (method in methods) renderInterfaceMethod(method)];
				"interface{" + rendered.join("; ") + "}";
			case GenericType(base, arguments): base.render() + "[" + [for (argument in arguments) argument.render()].join(", ") + "]";
			case MultiResult(results): "(" + [for (result in results) result.render()].join(", ") + ")";
			case Variadic(element): "..." + element.render();
		};
	}

	/** Traverse structural type nodes to discover use of an imported package. */
	public function usesPackage(alias:String):Bool {
		return switch (this) {
			case Builtin(_): false;
			case Named(packageName, _): packageName != null && packageName.value() == alias;
			case Pointer(element), Slice(element), ArrayType(_, element), Channel(_, element), Variadic(element): element.usesPackage(alias);
			case MapType(key, value): key.usesPackage(alias) || value.usesPackage(alias);
			case FunctionType(params, result): anyUsesPackage(params, alias) || (result != null && result.usesPackage(alias));
			case InterfaceType(methods):
				var used = false;
				for (method in methods) {
					if (anyUsesPackage(method.params, alias) || anyUsesPackage(method.results, alias)) {
						used = true;
						break;
					}
				}
				used;
			case GenericType(base, arguments): base.usesPackage(alias) || anyUsesPackage(arguments, alias);
			case MultiResult(results): anyUsesPackage(results, alias);
		};
	}

	/**
		What: Report whether this target type can syntactically head a composite literal.
		Why: Builtins, pointers, functions, channels, interfaces, multi-results, and
		variadics cannot be printed before `{...}` as literal types.
		How: Admit named and generic named forms plus the structural array, slice, and
		map forms; later Go type checking resolves named underlying-type constraints.
	**/
	public function supportsCompositeLiteral():Bool {
		return switch (this) {
			case Named(_, _), Slice(_), ArrayType(_, _), MapType(_, _), GenericType(_, _): true;
			case _: false;
		};
	}

	private inline function kind():GoTypeKind {
		return this;
	}

	static function renderInterfaceMethod(method:GoTypeInterfaceMethod):String {
		var suffix = switch (method.results.length) {
			case 0: "";
			case 1: " " + method.results[0].render();
			case _: " (" + [for (result in method.results) result.render()].join(", ") + ")";
		};
		return method.name + "(" + [for (param in method.params) param.render()].join(", ") + ")" + suffix;
	}

	static function flattenResult(result:Null<GoType>):Array<GoType> {
		if (result == null) {
			return [];
		}
		return switch (result.kind()) {
			case MultiResult(results): results.copy();
			case _: [result];
		};
	}

	static function anyUsesPackage(types:Array<GoType>, alias:String):Bool {
		for (type in types) {
			if (type.usesPackage(alias)) {
				return true;
			}
		}
		return false;
	}

	static function validateNested(type:GoType, role:String):Void {
		if (type == null) {
			throw "Invalid Go type: " + role + " cannot be null";
		}
		switch (type.kind()) {
			case MultiResult(_):
				throw "Invalid Go type: " + role + " cannot be a multi-result";
			case Variadic(_):
				throw "Invalid Go type: " + role + " cannot be variadic";
			case _:
		}
	}

	static function isComparableMapKey(type:GoType):Bool {
		return switch (type.kind()) {
			case Slice(_), MapType(_, _), FunctionType(_, _), MultiResult(_), Variadic(_): false;
			case ArrayType(_, element): isComparableMapKey(element);
			case _: true;
		};
	}

	static function validateIdentifier(name:String, role:String):Void {
		if (!GoPackageName.isIdentifier(name)) {
			throw 'Invalid Go type: ' + role + ' identifier "' + name + '"';
		}
	}
}

private class GoTypeParser {
	final source:String;
	var offset:Int = 0;

	public function new(source:String) {
		this.source = source;
	}

	public function parse():GoType {
		skipWhitespace();
		if (offset >= source.length) {
			fail("value cannot be empty");
		}
		var result = parseType();
		skipWhitespace();
		if (offset != source.length) {
			fail('unexpected token "' + source.substr(offset) + '"');
		}
		return result;
	}

	function parseType():GoType {
		skipWhitespace();
		if (consume("...")) {
			return GoType.variadic(parseType());
		}
		if (consume("*")) {
			return GoType.pointer(parseType());
		}
		if (consume("[]")) {
			return GoType.slice(parseType());
		}
		if (peek("[")) {
			return parseArray();
		}
		if (consumeKeyword("map")) {
			expect("[");
			var key = parseType();
			expect("]");
			return GoType.map(key, parseType());
		}
		if (consume("<-")) {
			if (!consumeKeyword("chan")) {
				fail("receive-only type must use <-chan");
			}
			return GoType.channel(GoChannelDirection.ReceiveOnly, parseType());
		}
		if (consumeKeyword("chan")) {
			skipWhitespace();
			var direction = consume("<-") ? GoChannelDirection.SendOnly : GoChannelDirection.Bidirectional;
			return GoType.channel(direction, parseType());
		}
		if (consumeKeyword("func")) {
			return parseFunction();
		}
		if (consumeKeyword("interface")) {
			return parseInterface();
		}
		if (consume("(")) {
			var grouped = parseTypeList(")");
			return switch (grouped.length) {
				case 0: fail("parenthesized type cannot be empty");
				case 1: grouped[0];
				case _: GoType.multiResult(grouped);
			};
		}
		return parseNamed();
	}

	function parseArray():GoType {
		expect("[");
		skipWhitespace();
		var start = offset;
		while (offset < source.length && isDigit(source.charCodeAt(offset))) {
			offset++;
		}
		if (start == offset) {
			fail("array length must be a non-negative integer literal");
		}
		var length = Std.parseInt(source.substr(start, offset - start));
		expect("]");
		return GoType.array(length, parseType());
	}

	function parseFunction():GoType {
		expect("(");
		var params = parseTypeList(")");
		skipWhitespace();
		var results = new Array<GoType>();
		if (peek("(")) {
			consume("(");
			results = parseTypeList(")");
		} else if (!atTypeBoundary()) {
			results.push(parseType());
		}
		return GoType.functionType(params, results);
	}

	function parseInterface():GoType {
		expect("{");
		var methods = new Array<GoTypeInterfaceMethod>();
		skipWhitespace();
		while (!consume("}")) {
			var name = parseIdentifier();
			expect("(");
			var params = parseTypeList(")");
			skipWhitespace();
			var results = new Array<GoType>();
			if (peek("(")) {
				consume("(");
				results = parseTypeList(")");
			} else if (!peek(";") && !peek("}")) {
				results.push(parseType());
			}
			methods.push(GoType.interfaceMethod(name, params, results));
			skipWhitespace();
			consume(";");
			skipWhitespace();
			if (offset >= source.length) {
				fail("unterminated interface type");
			}
		}
		return GoType.interfaceType(methods);
	}

	function parseNamed():GoType {
		var first = parseIdentifier();
		var builtin = GoBuiltinType.fromToken(first);
		var base:GoType;
		if (consume(".")) {
			base = GoType.qualified(GoPackageName.named(first), parseIdentifier());
		} else if (builtin != null) {
			base = GoType.builtin(builtin);
		} else {
			base = GoType.named(first);
		}
		skipWhitespace();
		if (consume("[")) {
			var arguments = parseTypeList("]");
			base = GoType.generic(base, arguments);
		}
		return base;
	}

	function parseTypeList(close:String):Array<GoType> {
		var values = new Array<GoType>();
		skipWhitespace();
		if (consume(close)) {
			return values;
		}
		while (true) {
			values.push(parseType());
			skipWhitespace();
			if (consume(close)) {
				return values;
			}
			expect(",");
		}
	}

	function parseIdentifier():String {
		skipWhitespace();
		var start = offset;
		if (offset >= source.length || !isIdentifierStart(source.charCodeAt(offset))) {
			fail("expected identifier");
		}
		offset++;
		while (offset < source.length && isIdentifierPart(source.charCodeAt(offset))) {
			offset++;
		}
		return source.substr(start, offset - start);
	}

	function atTypeBoundary():Bool {
		skipWhitespace();
		if (offset >= source.length) {
			return true;
		}
		return switch (source.charAt(offset)) {
			case ",", ")", "]", ";", "}": true;
			case _: false;
		};
	}

	function consumeKeyword(keyword:String):Bool {
		skipWhitespace();
		if (!peek(keyword)) {
			return false;
		}
		var after = offset + keyword.length;
		if (after < source.length && isIdentifierPart(source.charCodeAt(after))) {
			return false;
		}
		offset = after;
		return true;
	}

	function expect(token:String):Void {
		skipWhitespace();
		if (!consume(token)) {
			fail('expected "' + token + '"');
		}
	}

	function consume(token:String):Bool {
		skipWhitespace();
		if (!peek(token)) {
			return false;
		}
		offset += token.length;
		return true;
	}

	function peek(token:String):Bool {
		return source.substr(offset, token.length) == token;
	}

	function skipWhitespace():Void {
		while (offset < source.length) {
			var code = source.charCodeAt(offset);
			if (code != 32 && code != 9 && code != 10 && code != 13) {
				break;
			}
			offset++;
		}
	}

	function fail<T>(detail:String):T {
		throw 'Invalid Go type "' + source + '" at offset ' + offset + ": " + detail;
	}

	static inline function isDigit(code:Int):Bool {
		return code >= 48 && code <= 57;
	}

	static inline function isIdentifierStart(code:Int):Bool {
		return code == 95 || (code >= 65 && code <= 90) || (code >= 97 && code <= 122);
	}

	static inline function isIdentifierPart(code:Int):Bool {
		return isIdentifierStart(code) || isDigit(code);
	}
}
