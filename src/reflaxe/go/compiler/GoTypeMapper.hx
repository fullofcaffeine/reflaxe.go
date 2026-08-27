package reflaxe.go.compiler;

#if macro
import haxe.macro.Context;
import haxe.macro.Type;
import haxe.macro.TypeTools;
import reflaxe.go.naming.GoNaming;

typedef GoClassTypeNamer = ClassType->String;
typedef GoEnumTypeNamer = EnumType->String;

class GoTypeMapper {
	/**
		What: Maps one array-like element to its actual Go storage type.
		Why: `Null<Int>`, `Null<Float>`, and `Null<Bool>` must retain Go `nil`
		inside arrays instead of being unwrapped to non-nullable scalar storage.
		How: Keep the existing scalar mapping for ordinary elements and use `any`
		only for nullable primitives, matching local/value storage policy.
	**/
	static function arrayElementStorageGoType(elementType:Type, classTypeName:GoClassTypeNamer, enumTypeName:GoEnumTypeNamer):String {
		return isNullablePrimitiveType(elementType) ? "any" : scalarGoType(elementType, classTypeName, enumTypeName);
	}

	public static function typeToGoType(type:Type, classTypeName:GoClassTypeNamer, enumTypeName:GoEnumTypeNamer):String {
		if (isBytesDataType(type)) {
			return "[]int";
		}
		if (isNativeStringSliceType(type)) {
			return "[]string";
		}
		var nativeSliceElement = nativeSliceElementType(type);
		if (nativeSliceElement != null) {
			return "[]" + arrayElementStorageGoType(nativeSliceElement, classTypeName, enumTypeName);
		}
		var restElement = restElementType(type);
		if (restElement != null) {
			return "[]" + arrayElementStorageGoType(restElement, classTypeName, enumTypeName);
		}
		var vectorElement = vectorElementType(type);
		if (vectorElement != null) {
			return "[]" + arrayElementStorageGoType(vectorElement, classTypeName, enumTypeName);
		}
		var readOnlyElement = readOnlyArrayElementType(type);
		if (readOnlyElement != null) {
			return "*hxrt.Array";
		}

		var followed = Context.follow(type);
		return switch (followed) {
			case TFun(args, returnType):
				goFunctionType(args, returnType, classTypeName, enumTypeName);
			case TInst(classRef, params):
				var classType = classRef.get();
				if (isTypeParameterClass(classType)) {
					typeParameterGoType(classType, classTypeName, enumTypeName);
				} else if (isHaxeExceptionClass(classType)) {
					"*hxrt.ExceptionValue";
				} else if (classType.isInterface) {
					classTypeName(classType);
				} else if (classType.pack.length == 0 && classType.name == "String") {
					"*string";
				} else if (classType.pack.length == 0 && classType.name == "Array" && params.length == 1) {
					"*hxrt.Array";
				} else {
					"*" + classTypeName(classType);
				}
			case TEnum(enumRef, _):
				"*" + enumTypeName(enumRef.get());
			case TAnonymous(_):
				"map[string]any";
			case TAbstract(abstractRef, params):
				var abstractType = abstractRef.get();
				var mapped = abstractGoType(abstractType, params, classTypeName, enumTypeName);
				mapped == null ? "any" : mapped;
			case _:
				"any";
		};
	}

	public static function isStringType(type:Type):Bool {
		var followed = Context.follow(type);
		return switch (followed) {
			case TInst(classRef, _): var classType = classRef.get(); classType.pack.length == 0 && classType.name == "String";
			case TAbstract(abstractRef, _): var abstractType = abstractRef.get(); abstractType.pack.length == 0 && abstractType.name == "String";
			case _:
				false;
		};
	}

	/**
		What: Detects switch subjects whose runtime carrier is a Haxe string.
		Why: User-defined string abstracts retain their nominal type in the typed AST,
		so `isStringType` alone would emit a Go switch over string pointers or `any`.
		How: Follow only abstract backing types, with applied type parameters and a
		cycle guard, until the chain reaches the ordinary Haxe `String` type.
	**/
	public static function isStringSwitchType(type:Type):Bool {
		return isStringSwitchTypeInner(type, new Map<String, Bool>());
	}

	static function isStringSwitchTypeInner(type:Type, seen:Map<String, Bool>):Bool {
		if (isStringType(type)) {
			return true;
		}
		return switch (Context.follow(type)) {
			case TAbstract(abstractRef, params):
				var abstractType = abstractRef.get();
				var identity = abstractType.pack.concat([abstractType.name]).join(".");
				if (seen.exists(identity)) {
					false;
				} else {
					seen.set(identity, true);
					var backingType = TypeTools.applyTypeParameters(abstractType.type, abstractType.params, params);
					isStringSwitchTypeInner(backingType, seen);
				}
			case _:
				false;
		};
	}

	public static function isInterfaceType(type:Type):Bool {
		var followed = Context.follow(type);
		return switch (followed) {
			case TInst(classRef, _):
				classRef.get().isInterface;
			case _:
				false;
		};
	}

	public static function isBoolType(type:Type):Bool {
		var followed = Context.follow(type);
		return switch (followed) {
			case TAbstract(abstractRef, _): var abstractType = abstractRef.get(); abstractType.pack.length == 0 && abstractType.name == "Bool";
			case _:
				false;
		};
	}

	public static function isIntType(type:Type):Bool {
		var followed = Context.follow(type);
		return switch (followed) {
			case TAbstract(abstractRef, _): var abstractType = abstractRef.get(); abstractType.pack.length == 0 && (abstractType.name == "Int"
					|| abstractType.name == "UInt");
			case _:
				false;
		};
	}

	public static function isNullableIntType(type:Type):Bool {
		return switch (type) {
			case TAbstract(abstractRef, params): var abstractType = abstractRef.get(); abstractType.pack.length == 0 && abstractType.name == "Null" && params.length == 1 && isIntType(params[0]);
			case TMono(ref): var resolved = ref.get(); resolved != null && isNullableIntType(resolved);
			case TType(_, _): isNullableIntType(TypeTools.follow(type, true));
			case TLazy(f):
				isNullableIntType(f());
			case _:
				false;
		};
	}

	public static function isNullableFloatType(type:Type):Bool {
		return switch (type) {
			case TAbstract(abstractRef, params): var abstractType = abstractRef.get(); abstractType.pack.length == 0 && abstractType.name == "Null" && params.length == 1 && isFloatType(params[0]);
			case TMono(ref): var resolved = ref.get(); resolved != null && isNullableFloatType(resolved);
			case TType(_, _): isNullableFloatType(TypeTools.follow(type, true));
			case TLazy(f):
				isNullableFloatType(f());
			case _:
				false;
		};
	}

	public static function isNullableBoolType(type:Type):Bool {
		return switch (type) {
			case TAbstract(abstractRef, params): var abstractType = abstractRef.get(); abstractType.pack.length == 0 && abstractType.name == "Null" && params.length == 1 && isBoolType(params[0]);
			case TMono(ref): var resolved = ref.get(); resolved != null && isNullableBoolType(resolved);
			case TType(_, _): isNullableBoolType(TypeTools.follow(type, true));
			case TLazy(f):
				isNullableBoolType(f());
			case _:
				false;
		};
	}

	public static function isNullablePrimitiveType(type:Type):Bool {
		return switch (type) {
			case TAbstract(abstractRef, params): var abstractType = abstractRef.get(); abstractType.pack.length == 0 && abstractType.name == "Null" && params.length == 1 && isPrimitiveValueType(params[0]);
			case TMono(ref): var resolved = ref.get(); resolved != null && isNullablePrimitiveType(resolved);
			case TType(_, _): isNullablePrimitiveType(TypeTools.follow(type, true));
			case TLazy(f):
				isNullablePrimitiveType(f());
			case _:
				false;
		};
	}

	public static function isHaxeInt32Type(type:Type):Bool {
		return switch (type) {
			case TAbstract(abstractRef, _): var abstractType = abstractRef.get(); abstractType.pack.join(".") == "haxe" && abstractType.name == "Int32";
			case TMono(ref): var resolved = ref.get(); resolved != null && isHaxeInt32Type(resolved);
			case TType(_, _):
				isHaxeInt32Type(Context.follow(type));
			case _:
				false;
		};
	}

	public static function isInt32SemanticType(type:Type):Bool {
		return isIntType(type) || isHaxeInt32Type(type);
	}

	public static function isFloatType(type:Type):Bool {
		var followed = Context.follow(type);
		return switch (followed) {
			case TAbstract(abstractRef, _): var abstractType = abstractRef.get(); abstractType.pack.length == 0 && abstractType.name == "Float";
			case _:
				false;
		};
	}

	public static function isStdClassMetaType(type:Type):Bool {
		var followed = Context.follow(type);
		return switch (followed) {
			case TAbstract(abstractRef, _): var abstractType = abstractRef.get(); abstractType.pack.length == 0 && abstractType.name == "Class";
			case _:
				false;
		};
	}

	public static function isStdEnumMetaType(type:Type):Bool {
		var followed = Context.follow(type);
		return switch (followed) {
			case TAbstract(abstractRef, _): var abstractType = abstractRef.get(); abstractType.pack.length == 0 && abstractType.name == "Enum";
			case _:
				false;
		};
	}

	public static function isDefinitelyNonNullableType(type:Type):Bool {
		if (isNullablePrimitiveType(type)) {
			return false;
		}
		return isBoolType(type) || isIntType(type) || isFloatType(type);
	}

	static function isPrimitiveValueType(type:Type):Bool {
		return isBoolType(type) || isIntType(type) || isHaxeInt32Type(type) || isFloatType(type);
	}

	public static function isAnonymousObjectType(type:Type):Bool {
		var followed = Context.follow(type);
		return switch (followed) {
			case TAnonymous(_):
				true;
			case _:
				false;
		};
	}

	public static function isArrayType(type:Type):Bool {
		if (isBytesDataType(type)) {
			return true;
		}
		if (nativeSliceElementType(type) != null) {
			return true;
		}
		if (isNativeStringSliceType(type)) {
			return true;
		}
		if (restElementType(type) != null) {
			return true;
		}
		if (vectorElementType(type) != null) {
			return true;
		}
		if (readOnlyArrayElementType(type) != null) {
			return true;
		}

		return isHaxeArrayType(type);
	}

	/**
		What: Recognizes types that retain Haxe Array identity in generated Go.
		Why: Root `Array<T>` owns shared identity and sparse growth, while
		`ReadOnlyArray<T>` is an aliasing read-only view of that same object. `Rest`,
		`Vector`, and explicit `go.NativeSlice` values remain native slice-shaped
		boundaries instead.
		How: Admit the root nominal Array and the standard ReadOnlyArray abstract, but
		do not generalize this to every indexable or array-like type.
	**/
	public static function isHaxeArrayType(type:Type):Bool {
		if (isBytesDataType(type)) {
			return false;
		}
		if (readOnlyArrayElementType(type) != null) {
			return true;
		}
		var followed = Context.follow(type);
		return switch (followed) {
			case TInst(classRef, _): var classType = classRef.get(); classType.pack.length == 0 && classType.name == "Array";
			case _:
				false;
		};
	}

	public static function arrayElementGoType(type:Type, classTypeName:GoClassTypeNamer, enumTypeName:GoEnumTypeNamer):String {
		if (isNativeStringSliceType(type)) {
			return "string";
		}
		var nativeSliceElement = nativeSliceElementType(type);
		if (nativeSliceElement != null) {
			return arrayElementStorageGoType(nativeSliceElement, classTypeName, enumTypeName);
		}
		var restElement = restElementType(type);
		if (restElement != null) {
			return arrayElementStorageGoType(restElement, classTypeName, enumTypeName);
		}
		var vectorElement = vectorElementType(type);
		if (vectorElement != null) {
			return arrayElementStorageGoType(vectorElement, classTypeName, enumTypeName);
		}
		var readOnlyElement = readOnlyArrayElementType(type);
		if (readOnlyElement != null) {
			return arrayElementStorageGoType(readOnlyElement, classTypeName, enumTypeName);
		}

		var followed = Context.follow(type);
		return switch (followed) {
			case TInst(classRef, params):
				var classType = classRef.get();
				if (classType.pack.length == 0 && classType.name == "Array" && params.length == 1) {
					arrayElementStorageGoType(params[0], classTypeName, enumTypeName);
				} else {
					"any";
				}
			case _:
				"any";
		};
	}

	public static function scalarGoType(type:Type, classTypeName:GoClassTypeNamer, enumTypeName:GoEnumTypeNamer):String {
		if (isBytesDataType(type)) {
			return "[]int";
		}
		if (isNativeStringSliceType(type)) {
			return "[]string";
		}
		var nativeSliceElement = nativeSliceElementType(type);
		if (nativeSliceElement != null) {
			return "[]" + arrayElementStorageGoType(nativeSliceElement, classTypeName, enumTypeName);
		}
		var restElement = restElementType(type);
		if (restElement != null) {
			return "[]" + arrayElementStorageGoType(restElement, classTypeName, enumTypeName);
		}
		var vectorElement = vectorElementType(type);
		if (vectorElement != null) {
			return "[]" + arrayElementStorageGoType(vectorElement, classTypeName, enumTypeName);
		}
		var readOnlyElement = readOnlyArrayElementType(type);
		if (readOnlyElement != null) {
			return "*hxrt.Array";
		}

		var followed = Context.follow(type);
		return switch (followed) {
			case TFun(args, returnType):
				goFunctionType(args, returnType, classTypeName, enumTypeName);
			case TInst(classRef, params):
				var classType = classRef.get();
				if (isTypeParameterClass(classType)) {
					typeParameterGoType(classType, classTypeName, enumTypeName);
				} else if (isHaxeExceptionClass(classType)) {
					"*hxrt.ExceptionValue";
				} else if (classType.isInterface) {
					classTypeName(classType);
				} else if (classType.pack.length == 0 && classType.name == "String") {
					"*string";
				} else if (classType.pack.length == 0 && classType.name == "Array" && params.length == 1) {
					"*hxrt.Array";
				} else {
					"*" + classTypeName(classType);
				}
			case TEnum(enumRef, _):
				"*" + enumTypeName(enumRef.get());
			case TAnonymous(_):
				"map[string]any";
			case TAbstract(abstractRef, params):
				var abstractType = abstractRef.get();
				var mapped = abstractGoType(abstractType, params, classTypeName, enumTypeName);
				mapped == null ? "any" : mapped;
			case _:
				"any";
		};
	}

	static function abstractGoType(abstractType:AbstractType, params:Array<Type>, classTypeName:GoClassTypeNamer, enumTypeName:GoEnumTypeNamer):Null<String> {
		var externCarrier = externBackedAbstractGoType(abstractType, classTypeName, new Map<String, Bool>());
		if (externCarrier != null) {
			return externCarrier;
		}
		if (abstractType.pack.length == 0 && abstractType.name == "Int") {
			return "int";
		}
		if (abstractType.pack.length == 0 && abstractType.name == "UInt") {
			return "int";
		}
		if (isHaxeMapAbstract(abstractType)) {
			return mapAbstractGoType(params);
		}
		if (abstractType.pack.length == 0 && abstractType.name == "Float") {
			return "float64";
		}
		if (abstractType.pack.length == 0 && abstractType.name == "Bool") {
			return "bool";
		}
		if (abstractType.pack.length == 0 && abstractType.name == "String") {
			return "*string";
		}
		if (abstractType.pack.join(".") == "haxe" && abstractType.name == "Int32") {
			return "int";
		}
		if (abstractType.pack.join(".") == "haxe" && abstractType.name == "Int64") {
			return "*haxe___Int64_____Int64";
		}
		if (abstractType.pack.join(".") == "haxe" && abstractType.name == "EnumFlags") {
			return "int";
		}
		if (abstractType.pack.join(".") == "haxe.io"
			&& (abstractType.name == "ArrayBufferView"
				|| abstractType.name == "UInt8Array"
				|| abstractType.name == "UInt16Array"
				|| abstractType.name == "UInt32Array"
				|| abstractType.name == "Int32Array"
				|| abstractType.name == "Float32Array"
				|| abstractType.name == "Float64Array")) {
			return "*haxe__io__ArrayBufferViewImpl";
		}
		return null;
	}

	/**
		What
		- Preserves the native Go type of an abstract backed by typed extern metadata.

		Why
		- Falling every non-primitive abstract back to `any` discards a source-declared
		  opaque carrier and forces generated boxing and type assertions even when the
		  backend already has exact type authority.

		How
		- Follow chains of abstracts until they reach an extern class with
		  `@:go.import`, then render that class through the normal extern type mapper.
		  Cycles and abstracts without typed native authority keep the existing fallback.
	**/
	static function externBackedAbstractGoType(abstractType:AbstractType, classTypeName:GoClassTypeNamer, seen:Map<String, Bool>):Null<String> {
		var qualifiedName = abstractType.pack.concat([abstractType.name]).join(".");
		if (seen.exists(qualifiedName)) {
			return null;
		}
		seen.set(qualifiedName, true);

		return switch (Context.follow(abstractType.type)) {
			case TInst(classRef, _): var classType = classRef.get(); classType.isExtern && hasGoImportMetadata(classType) ? "*" + classTypeName(classType) : null;
			case TAbstract(innerRef, _):
				externBackedAbstractGoType(innerRef.get(), classTypeName, seen);
			case _:
				null;
		};
	}

	static function hasGoImportMetadata(classType:ClassType):Bool {
		for (entry in classType.meta.get()) {
			if (GoMetadataName.GoImport.matches(entry.name)) {
				return true;
			}
		}
		return false;
	}

	public static function goFunctionType(args:Array<{name:String, opt:Bool, t:Type}>, returnType:Type, classTypeName:GoClassTypeNamer,
			enumTypeName:GoEnumTypeNamer):String {
		var params = [
			for (arg in args)
				functionParameterStorageGoType(arg, classTypeName, enumTypeName)
		].join(", ");
		if (isVoidType(returnType)) {
			return "func(" + params + ")";
		}
		return "func(" + params + ") " + scalarGoType(returnType, classTypeName, enumTypeName);
	}

	/**
		What: Maps a declared Haxe function parameter to its callable Go storage type.

		Why: Haxe marks both `?timeout:Int` and `timeout:Int = 10` as optional.
		The generated function literal and every variable or field carrying it must
		therefore share one signature. Mapping only the carrier as `func(any)` makes a
		non-null-default implementation such as `func(int)` fail Go type checking.

		How: Keep pointer/nil-capable parameter mappings unchanged and widen only
		optional scalar value types to `any`. `lowerFunctionParams` uses the same rule
		and asserts the scalar type when the function body reads the parameter.
	**/
	public static function functionParameterStorageGoType(arg:{name:String, opt:Bool, t:Type}, classTypeName:GoClassTypeNamer,
			enumTypeName:GoEnumTypeNamer):String {
		return arg.opt && isPrimitiveValueType(arg.t) ? "any" : scalarGoType(arg.t, classTypeName, enumTypeName);
	}

	/** Resolves the element type of the explicit `go.NativeSlice<T>` boundary. */
	public static function nativeSliceElementType(type:Type):Null<Type> {
		var followed = Context.follow(type);
		return switch (followed) {
			case TInst(classRef, params):
				var classType = classRef.get();
				if (classType.pack.join(".") == "go" && classType.name == "NativeSlice" && params.length == 1) {
					params[0];
				} else {
					null;
				}
			case _:
				null;
		};
	}

	/** True only for the explicit native Go `[]string` boundary. */
	public static function isNativeStringSliceType(type:Type):Bool {
		return switch (Context.follow(type)) {
			case TInst(classRef, _): final classType = classRef.get(); classType.pack.join(".") == "go" && classType.name == "NativeStringSlice";
			case _:
				false;
		};
	}

	/**
		What: Recognizes the upstream byte-storage alias before ordinary alias
		following turns it into root `Array<Int>`.
		Why: The compiler-owned Bytes carrier currently stores a native `[]int`, and
		`BytesData` must remain an aliasing view of that storage rather than a copied
		portable Array.
		How: Match only the declared `haxe.io.BytesData` typedef and preserve it as an
		explicit raw array-like boundary until Bytes ownership migrates separately.
	**/
	public static function isBytesDataType(type:Type):Bool {
		return switch (type) {
			case TType(typeRef, _): var typeDef = typeRef.get(); typeDef.pack.join(".") == "haxe.io" && typeDef.name == "BytesData";
			case TMono(ref): var resolved = ref.get(); resolved != null && isBytesDataType(resolved);
			case TLazy(resolve):
				isBytesDataType(resolve());
			case _:
				false;
		};
	}

	public static function restElementType(type:Type):Null<Type> {
		var followed = Context.follow(type);
		return switch (followed) {
			case TAbstract(abstractRef, params):
				var abstractType = abstractRef.get();
				if (abstractType.pack.join(".") == "haxe" && abstractType.name == "Rest" && params.length == 1) {
					params[0];
				} else {
					null;
				}
			case TType(typeRef, params):
				var typeDef = typeRef.get();
				if (typeDef.pack.join(".") == "haxe._Rest" && typeDef.name == "NativeRest" && params.length == 1) {
					params[0];
				} else {
					null;
				}
			case _:
				null;
		};
	}

	public static function vectorElementType(type:Type):Null<Type> {
		var followed = Context.follow(type);
		return switch (followed) {
			case TAbstract(abstractRef, params):
				var abstractType = abstractRef.get();
				if (abstractType.pack.join(".") == "haxe.ds" && abstractType.name == "Vector" && params.length == 1) {
					params[0];
				} else {
					null;
				}
			case _:
				null;
		};
	}

	public static function readOnlyArrayElementType(type:Type):Null<Type> {
		var followed = Context.follow(type);
		return switch (followed) {
			case TAbstract(abstractRef, params):
				var abstractType = abstractRef.get();
				if (abstractType.pack.join(".") == "haxe.ds" && abstractType.name == "ReadOnlyArray" && params.length == 1) {
					params[0];
				} else {
					null;
				}
			case _:
				null;
		};
	}

	public static function isVoidType(type:Type):Bool {
		var followed = Context.follow(type);
		return switch (followed) {
			case TAbstract(abstractRef, _): var abstractType = abstractRef.get(); abstractType.pack.length == 0 && abstractType.name == "Void";
			case _:
				false;
		};
	}

	static function mapAbstractGoType(params:Array<Type>):String {
		if (params.length != 2) {
			return "any";
		}

		var keyType = Context.follow(params[0]);
		return switch (keyType) {
			case TEnum(_, _):
				"*haxe__ds__EnumValueMap";
			case _:
				if (isStringType(params[0])) {
					"*haxe__ds__StringMap";
				} else if (isIntType(params[0])) {
					"*haxe__ds__IntMap";
				} else {
					"*haxe__ds__ObjectMap";
				}
		};
	}

	static function isHaxeMapAbstract(abstractType:AbstractType):Bool {
		return abstractType.name == "Map" && (abstractType.pack.length == 0 || abstractType.pack.join(".") == "haxe.ds");
	}

	public static function isDynamicCatchType(type:Type):Bool {
		var followed = Context.follow(type);
		return switch (followed) {
			case TDynamic(_):
				true;
			case TAbstract(abstractRef, _): var abstractType = abstractRef.get(); abstractType.pack.length == 0 && abstractType.name == "Dynamic";
			case _:
				false;
		};
	}

	public static function isTypeParameterClass(classType:ClassType):Bool {
		return switch (classType.kind) {
			case KTypeParameter(_):
				true;
			case _:
				false;
		};
	}

	/**
		What
		Maps a constrained Haxe type parameter to a Go type.

		Why
		Most type parameters still erase to `any` on `haxe.go`, but some staged-stdlib
		surfaces depend on method-only constraints such as `K:{function hashCode():Int;}`.
		Erasing those to `any` breaks direct method dispatch and forces incorrect
		reflective fallbacks.

		How
		When the constraint set is a method-only anonymous structure, synthesize a local
		Go interface with matching method signatures. Otherwise, keep the existing `any`
		fallback.
	**/
	static function typeParameterGoType(classType:ClassType, classTypeName:GoClassTypeNamer, enumTypeName:GoEnumTypeNamer):String {
		return switch (classType.kind) {
			case KTypeParameter(constraints):
				var mapped = anonymousMethodConstraintGoType(constraints, classTypeName, enumTypeName);
				mapped == null ? "any" : mapped;
			case _:
				"any";
		};
	}

	static function anonymousMethodConstraintGoType(constraints:Array<Type>, classTypeName:GoClassTypeNamer, enumTypeName:GoEnumTypeNamer):Null<String> {
		if (constraints == null || constraints.length == 0) {
			return null;
		}

		var signatures = new Map<String, String>();
		for (constraint in constraints) {
			switch (Context.follow(constraint)) {
				case TAnonymous(anonRef):
					for (field in anonRef.get().fields) {
						var signature = constraintMethodSignature(field, classTypeName, enumTypeName);
						if (signature == null) {
							return null;
						}
						signatures.set(GoNaming.normalizeIdent(field.name), signature);
					}
				case _:
					return null;
			}
		}

		var ordered = [for (signature in signatures) signature];
		if (ordered.length == 0) {
			return null;
		}
		ordered.sort(Reflect.compare);
		return "interface{" + ordered.join("; ") + "}";
	}

	static function constraintMethodSignature(field:ClassField, classTypeName:GoClassTypeNamer, enumTypeName:GoEnumTypeNamer):Null<String> {
		return switch (Context.follow(field.type)) {
			case TFun(args, returnType):
				var params = [
					for (arg in args)
						functionParameterStorageGoType(arg, classTypeName, enumTypeName)
				].join(", ");
				var returnSuffix = isVoidType(returnType) ? "" : " " + scalarGoType(returnType, classTypeName, enumTypeName);
				GoNaming.normalizeIdent(field.name) + "(" + params + ")" + returnSuffix;
			case _:
				null;
		};
	}

	public static function isHaxeExceptionClass(classType:ClassType):Bool {
		return switch (classType.pack.join(".") + "." + classType.name) {
			case "haxe.Exception", "haxe.ValueException":
				true;
			case _:
				false;
		};
	}

	public static function isHaxeExceptionType(type:Type):Bool {
		var followed = Context.follow(type);
		return switch (followed) {
			case TInst(classRef, _):
				isHaxeExceptionClass(classRef.get());
			case _:
				false;
		};
	}
}
#end
