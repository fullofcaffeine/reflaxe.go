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
	modules keep their own Go carriers in either query order.

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

	static function run():Void {
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
				require(goType(get(other, "id")) == "any", "int-backed Id");
			} else {
				require(goType(get(other, "id")) == "any", "int-backed Id after reuse");
				require(goType(get(fixture, "id")) == "*string", "string-backed Id after reuse");
			}
		}

		require(GoTypeMapper.isAnonymousObjectType(get(fixture, "row")), "typedef structure");
		require(GoTypeMapper.isAnonymousObjectType(get(fixture, "empty")), "empty structure");
		require(GoTypeMapper.isAnonymousObjectType(get(fixture, "structural")), "structural type parameter");
		require(!GoTypeMapper.isAnonymousObjectType(get(fixture, "nominal")), "unconstrained type parameter");
		require(!GoTypeMapper.isAnonymousObjectType(get(fixture, "text")), "class instance");

		require(GoTypeMapper.isStringType(get(fixture, "text")), "String");
		require(goType(get(fixture, "text")) == "*string", "String carrier");
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
		require(GoTypeMapper.isNativeInt64Type(get(fixture, "nativeInt64")), "go.NativeInt64");
		require(goType(get(fixture, "nativeInt64")) == "int64", "go.NativeInt64 carrier");
		Sys.println("ok");
	}
	#end
}
