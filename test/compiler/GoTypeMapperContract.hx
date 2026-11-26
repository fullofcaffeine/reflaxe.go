#if macro
import haxe.macro.Context;
import haxe.macro.Type;
import reflaxe.go.compiler.GoTypeMapper;
#end

/**
	Contract for how the type mapper reads declarations.

	Why
	Type queries run hundreds of thousands of times in a large program. The type
	mapper therefore identifies a declaration by the text of its reference
	(`Ref.toString()`) and remembers abstract declarations by path, instead of
	rebuilding declarations with `Ref.get()` on every question. Both shortcuts are
	correct only while two promises hold: a reference prints exactly `pack.Name`,
	and an abstract path names one declaration.

	What
	Proves the first promise for root, packaged, private, typedef, enum, and type
	parameter declarations. Proves the mapper still recognizes each declaration it
	checks by path, and that two private abstracts with the same name in different
	modules keep their own Go carriers in either query order. Proves the same for
	the remembered answer of `isIntType`, that the Int, Float and Bool
	predicates follow abstracts to their underlying types (haxe_go-gjqe), and
	that such abstracts use the Go carrier of that type while `Null<T>` of them
	stays nil-capable (haxe_go-5yxq).

	How
	After the initialization macros, loads two fixture classes, compares every
	reachable reference's `toString()` with its declaration's `pack` and `name`,
	and asks the type mapper the questions that use those shortcuts.
**/
class GoTypeMapperContract {
	#if macro
	public static function check():Void {
		Context.onAfterInitMacros(run);
	}

	static function require(value:Bool, message:String):Void {
		if (!value)
			Context.fatalError("GoTypeMapperContract: " + message, Context.currentPos());
	}

	static function requirePath<T:BaseType>(ref:Ref<T>):Void {
		final declaration = ref.get();
		final expected = declaration.pack.length == 0 ? declaration.name : declaration.pack.join(".") + "." + declaration.name;
		require(ref.toString() == expected, "reference prints " + ref.toString() + " for " + expected);
	}

	static function requirePaths(type:Type, depth:Int = 0):Void {
		if (depth > 8)
			return;
		switch (type) {
			case TInst(ref, params):
				requirePath(ref);
				switch (ref.get().kind) {
					case KTypeParameter(constraints):
						for (constraint in constraints)
							requirePaths(constraint, depth + 1);
					case _:
				}
				for (param in params)
					requirePaths(param, depth + 1);
			case TAbstract(ref, params):
				requirePath(ref);
				for (param in params)
					requirePaths(param, depth + 1);
			case TType(ref, params):
				requirePath(ref);
				requirePaths(ref.get().type, depth + 1);
				for (param in params)
					requirePaths(param, depth + 1);
			case TEnum(ref, params):
				requirePath(ref);
				for (param in params)
					requirePaths(param, depth + 1);
			case TFun(args, result):
				for (arg in args)
					requirePaths(arg.t, depth + 1);
				requirePaths(result, depth + 1);
			case TAnonymous(ref):
				for (field in ref.get().fields)
					requirePaths(field.type, depth + 1);
			case TMono(ref):
				final resolved = ref.get();
				if (resolved != null)
					requirePaths(resolved, depth + 1);
			case TLazy(resolve):
				requirePaths(resolve(), depth + 1);
			case TDynamic(_):
		}
	}

	static function fieldTypes(className:String):Map<String, Type> {
		final types = new Map<String, Type>();
		switch (Context.getType(className)) {
			case TInst(ref, _):
				requirePath(ref);
				final classType = ref.get();
				for (param in classType.params)
					requirePaths(param.t);
				for (field in classType.fields.get()) {
					for (param in field.params)
						requirePaths(param.t);
					requirePaths(field.type);
					types.set(field.name, field.type);
				}
			case _:
				require(false, className + " is not a class");
		}
		return types;
	}

	static function goType(type:Type):String {
		return GoTypeMapper.typeToGoType(type, classType -> classType.name, enumType -> enumType.name);
	}

	/** Class facts may be cached by `uniqueClassPath` only for module-level classes. */
	static function checkClassKeysAndFieldNames():Void {
		final classType = switch (Context.getType("GoTypeMapperContractFixture")) {
			case TInst(ref, _): ref.get();
			case _: null;
		};
		require(classType != null, "fixture class");
		require(GoTypeMapper.uniqueClassPath(classType) == "GoTypeMapperContractFixture", "module class key");
		for (param in classType.params)
			switch (param.t) {
				case TInst(ref, _):
					require(GoTypeMapper.uniqueClassPath(ref.get()) == null, "class type parameter " + ref.toString() + " has a key");
				case _:
					require(false, "class type parameter shape");
			}
		var methodParameters = 0;
		var fieldReads = 0;
		for (field in classType.fields.get()) {
			for (param in field.params)
				switch (param.t) {
					case TInst(ref, _):
						methodParameters++;
						require(GoTypeMapper.uniqueClassPath(ref.get()) == null, "method type parameter " + ref.toString() + " has a key");
					case _:
						require(false, "method type parameter shape");
				}
			final body = field.expr();
			if (body == null)
				continue;
			function visit(expr:TypedExpr):Void {
				switch (expr.expr) {
					case TField(_, FInstance(_, _, ref) | FStatic(_, ref)):
						fieldReads++;
						require(GoTypeMapper.referencedFieldName(ref) == ref.get().name, "field reference prints " + ref.toString());
					case _:
				}
				haxe.macro.TypedExprTools.iter(expr, visit);
			}
			visit(body);
		}
		require(methodParameters > 0, "fixture method type parameter");
		require(fieldReads >= 2, "fixture instance and static field reads");
	}

	static function run():Void {
		checkClassKeysAndFieldNames();
		final fixture = fieldTypes("GoTypeMapperContractFixture");
		final other = fieldTypes("GoTypeMapperContractOtherFixture");
		final get = (types:Map<String, Type>, name:String) -> {
			final type = types.get(name);
			require(type != null, "missing fixture field " + name);
			return (type : Type);
		};

		// Same-named private abstracts are separate declarations with separate carriers.
		for (round in 0...2) {
			if (round == 0) {
				require(goType(get(fixture, "id")) == "*string", "string-backed Id");
				require(goType(get(other, "id")) == "int", "int-backed Id");
			} else {
				require(goType(get(other, "id")) == "int", "int-backed Id after reuse");
				require(goType(get(fixture, "id")) == "*string", "string-backed Id after reuse");
			}
		}
		// The underlying Int answer is remembered by path, so it must also keep them apart.
		for (round in 0...2) {
			if (round == 0) {
				require(!GoTypeMapper.isIntType(get(fixture, "id")), "string-backed Id is not an Int value");
				require(GoTypeMapper.isIntType(get(other, "id")), "int-backed Id is an Int value");
			} else {
				require(GoTypeMapper.isIntType(get(other, "id")), "int-backed Id after reuse");
				require(!GoTypeMapper.isIntType(get(fixture, "id")), "string-backed Id after reuse");
			}
		}

		// The scalar predicates follow abstracts to their underlying Int, Float or Bool.
		require(GoTypeMapper.isIntType(get(fixture, "int32")), "haxe.Int32 is an Int value");
		require(GoTypeMapper.isIntType(get(fixture, "nullableInt")), "Null<Int> is an Int value");
		require(GoTypeMapper.isIntType(get(fixture, "level")), "enum abstract over Int is an Int value");
		require(!GoTypeMapper.isFloatType(get(fixture, "level")), "enum abstract over Int is not a Float value");
		require(GoTypeMapper.isFloatType(get(fixture, "meters")), "abstract over Float is a Float value");
		require(!GoTypeMapper.isIntType(get(fixture, "meters")), "abstract over Float is not an Int value");
		require(GoTypeMapper.isFloatType(get(fixture, "distance")), "abstract over an abstract over Float");
		require(GoTypeMapper.isNullableFloatType(get(fixture, "maybeMeters")), "Null<abstract over Float>");
		require(!GoTypeMapper.isNullableFloatType(get(fixture, "meters")), "abstract over Float is not Null<T>");
		require(GoTypeMapper.isIntType(get(fixture, "wrappedInt")), "Wrap<Int> applies its type parameter");
		require(!GoTypeMapper.isIntType(get(fixture, "wrappedText")), "Wrap<String> applies its type parameter");
		require(!GoTypeMapper.isIntType(get(fixture, "maybeCount")), "abstract over Null<Int> keeps its own recovery");
		require(!GoTypeMapper.isIntType(get(fixture, "map")) && !GoTypeMapper.isFloatType(get(fixture, "map")), "Map");
		require(!GoTypeMapper.isIntType(get(fixture, "text")) && !GoTypeMapper.isFloatType(get(fixture, "text")), "String");
		require(GoTypeMapper.isBoolType(get(fixture, "flag")), "abstract over Bool is a Bool value");
		require(GoTypeMapper.isBoolType(get(fixture, "toggle")), "enum abstract over Bool is a Bool value");
		require(!GoTypeMapper.isIntType(get(fixture, "flag")), "abstract over Bool is not an Int value");

		// Such abstracts use the underlying Go carrier; Null<T> of them stays nil-capable (haxe_go-5yxq).
		require(goType(get(fixture, "level")) == "int", "enum abstract over Int carrier");
		require(goType(get(fixture, "meters")) == "float64", "abstract over Float carrier");
		require(goType(get(fixture, "distance")) == "float64", "abstract over an abstract over Float carrier");
		require(goType(get(fixture, "flag")) == "bool", "abstract over Bool carrier");
		require(goType(get(fixture, "toggle")) == "bool", "enum abstract over Bool carrier");
		require(goType(get(fixture, "wrappedInt")) == "int", "Wrap<Int> carrier");
		require(goType(get(fixture, "wrappedText")) == "*string", "Wrap<String> carrier");
		require(goType(get(fixture, "maybeCount")) == "any", "abstract over Null<Int> carrier");
		require(goType(get(fixture, "byLevel")) == "*haxe__ds__IntMap", "Map keyed by an enum abstract over Int");
		require(GoTypeMapper.isNullablePrimitiveType(get(fixture, "maybeMeters")), "Null<abstract over Float> is nil-capable");
		require(GoTypeMapper.isNullablePrimitiveType(get(fixture, "maybeLevel")), "Null<enum abstract over Int> is nil-capable");
		require(!GoTypeMapper.isNullablePrimitiveType(get(fixture, "meters")), "abstract over Float is not nil-capable");
		require(GoTypeMapper.isScalarAbstractType(get(fixture, "level")), "enum abstract over Int is a scalar abstract");
		require(GoTypeMapper.isScalarAbstractType(get(fixture, "int32")), "haxe.Int32 is a scalar abstract");
		require(!GoTypeMapper.isScalarAbstractType(get(fixture, "nullableInt")), "Null<Int> is not a scalar abstract");
		require(!GoTypeMapper.isScalarAbstractType(get(fixture, "maybeCount")), "abstract over Null<Int> is not a scalar abstract");

		require(GoTypeMapper.isAnonymousObjectType(get(fixture, "row")), "typedef structure");
		require(GoTypeMapper.isAnonymousObjectType(get(fixture, "empty")), "empty structure");
		require(GoTypeMapper.isAnonymousObjectType(get(fixture, "structural")), "structural type parameter");
		require(!GoTypeMapper.isAnonymousObjectType(get(fixture, "nominal")), "unconstrained type parameter");
		require(!GoTypeMapper.isAnonymousObjectType(get(fixture, "text")), "class instance");

		require(GoTypeMapper.isStringType(get(fixture, "text")), "String");
		require(goType(get(fixture, "text")) == "*string", "String carrier");
		require(goType(get(fixture, "externText")) == "*string", "extern abstract over String carrier");
		require(goType(get(fixture, "kind")) == "*GoTypeMapperContractKind", "enum carrier");
		require(GoTypeMapper.isNullableIntType(get(fixture, "nullableInt")), "Null<Int>");
		require(!GoTypeMapper.isNullableIntType(get(fixture, "nullableBool")), "Null<Bool> is not Null<Int>");
		require(GoTypeMapper.isNullablePrimitiveType(get(fixture, "nullableBool")), "Null<Bool>");
		require(GoTypeMapper.isHaxeInt32Type(get(fixture, "int32")), "haxe.Int32");
		require(goType(get(fixture, "int32")) == "int", "haxe.Int32 carrier");
		require(goType(get(fixture, "int64")) == "*haxe___Int64_____Int64", "haxe.Int64 carrier");
		require(goType(get(fixture, "map")) == "*haxe__ds__StringMap", "Map<String, Int> carrier");

		final restElement = GoTypeMapper.restElementType(get(fixture, "rest"));
		require(restElement != null && GoTypeMapper.isIntType(restElement), "haxe.Rest element");
		final vectorElement = GoTypeMapper.vectorElementType(get(fixture, "vector"));
		require(vectorElement != null && GoTypeMapper.isStringType(vectorElement), "haxe.ds.Vector element");
		final readOnlyElement = GoTypeMapper.readOnlyArrayElementType(get(fixture, "readOnly"));
		require(readOnlyElement != null && GoTypeMapper.isIntType(readOnlyElement), "haxe.ds.ReadOnlyArray element");
		require(GoTypeMapper.isHaxeArrayType(get(fixture, "readOnly")), "ReadOnlyArray keeps Array identity");
		require(GoTypeMapper.isBytesDataType(get(fixture, "bytes")), "haxe.io.BytesData");
		final sliceElement = GoTypeMapper.nativeSliceElementType(get(fixture, "slice"));
		require(sliceElement != null && GoTypeMapper.isBoolType(sliceElement), "go.NativeSlice element");
		require(goType(get(fixture, "slice")) == "[]bool", "go.NativeSlice carrier");
		require(GoTypeMapper.nativeScalarCarrierName(get(fixture, "slice")) == null, "go.NativeSlice is not a scalar carrier");
		require(GoTypeMapper.isNativeAnyMapType(get(fixture, "anyMap")), "go.NativeAnyMap");
		require(goType(get(fixture, "anyMap")) == "map[string]any", "go.NativeAnyMap carrier");
		require(goType(get(fixture, "anyValue")) == "any", "go.NativeAny carrier");
		require(goType(get(fixture, "anySlice")) == "[]any", "go.NativeSlice<go.NativeAny> carrier");
		for (field => expected in [
			"nativeInt64" => {carrier: "NativeInt64", goType: "int64"},
			"nativeUintptr" => {carrier: "NativeUintptr", goType: "uintptr"},
			"nativeUint32" => {carrier: "NativeUint32", goType: "uint32"},
			"nativeUint16" => {carrier: "NativeUint16", goType: "uint16"},
			"nativeUint" => {carrier: "NativeUint", goType: "uint"}
		]) {
			require(GoTypeMapper.nativeScalarCarrierName(get(fixture, field)) == expected.carrier, "go." + expected.carrier);
			require(GoTypeMapper.nativeScalarCarrierGoType(expected.carrier) == expected.goType, "go." + expected.carrier + " Go type");
			require(goType(get(fixture, field)) == expected.goType, "go." + expected.carrier + " carrier");
		}
		Sys.println("ok");
	}
	#end
}
