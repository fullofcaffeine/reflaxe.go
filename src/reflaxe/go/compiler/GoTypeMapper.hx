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
		What: Reads the dotted path of a referenced class, abstract, typedef, or
		enum (`pack.Name`, or `Name` in the root package) without decoding it.

		Why: Most type queries only ask "is this `Null`, `Int`, `haxe.ds.Vector`,
		...?", and the compiler asks them hundreds of thousands of times in a large
		program. `Ref.get()` rebuilds the whole declaration as fresh macro values on
		every call: every field, operator, and metadata entry. For `haxe.Int32` or
		`haxe.Int64` one call costs a fraction of a millisecond, and those calls
		dominated declaration lowering time.

		How: The macro `Ref.toString()` prints exactly this path and nothing else,
		so comparing it gives the same answer as comparing `pack` and `name`. That
		includes type parameter classes, whose path is built the same way.
		`GoTypeMapperContract` pins this. Use `Ref.get()` only for checks that need
		more than the path.
	**/
	public static inline function declarationPath<T>(ref:Ref<T>):String {
		return ref.toString();
	}

	/**
		What: Reads the name of a referenced class field without decoding it.
		Why: Decoding a field rebuilds its type, metadata, and expression access.
		How: A field reference's `toString()` prints exactly the field name.
		`GoTypeMapperContract` pins this.
	**/
	public static inline function referencedFieldName(ref:Ref<ClassField>):String {
		return ref.toString();
	}

	/**
		What: Returns a path that names exactly one class declaration in this
		compilation, or null when the class has none.

		Why: Per-compilation caches of facts about a class (for example "does it
		declare instance fields?") need a key that can never mix two declarations
		up. The compiler otherwise decodes the class's whole field list for each
		question.

		How: Module-level classes (ordinary, generic, generic instance, abstract
		implementation, and module-fields classes) have unique type paths: Haxe
		rejects two types with one path, and private types get module-qualified
		paths. Type parameter and constant-expression classes are not module types,
		and their paths repeat (every method `map<S>` owns a `map.S`), so they get
		no key and callers must not cache them.
	**/
	public static function uniqueClassPath(classType:ClassType):Null<String> {
		return switch (classType.kind) {
			case KNormal | KGeneric | KGenericInstance(_, _) | KAbstractImpl(_) | KModuleFields(_):
				classType.pack.length == 0 ? classType.name : classType.pack.join(".") + "." + classType.name;
			case _:
				null;
		};
	}

	/**
		What: Remembers each abstract declaration this compilation has decoded,
		keyed by its path.

		Why: Choosing the Go carrier of an abstract reads its declaration (backing
		type, implementation class, metadata), and `Ref.get()` rebuilds all of it on
		every call. Lowering asks about `haxe.Int32`, `haxe.Int64`, and `Map`
		thousands of times, and their declarations carry many operator fields.

		How: A path names exactly one abstract in a compilation: Haxe rejects two
		types with the same path, private types get a module-qualified path, and
		type parameters are classes, never abstracts. Every caller (type mapping and
		the profile contract analyzer) runs after typing, when declarations no
		longer change. Haxe resets macro statics for every compilation (this map is
		not `@:persistent`), so a compilation server never sees an old entry.
		Callers only read the returned declaration.
	**/
	static final abstractDeclarations = new Map<String, AbstractType>();

	public static function abstractDeclaration(ref:Ref<AbstractType>):AbstractType {
		final path = declarationPath(ref);
		final cached = abstractDeclarations.get(path);
		if (cached != null) {
			return cached;
		}
		final declaration = ref.get();
		abstractDeclarations.set(path, declaration);
		return declaration;
	}

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
		rejectNullableNativeScalarCarrier(type);
		if (isBytesDataType(type)) {
			return "[]int";
		}
		var nativePointerElement = nativePointerElementType(type);
		if (nativePointerElement != null) {
			var nativeElement = nativeScalarPointerElementGoType(nativePointerElement);
			if (nativeElement == null)
				Context.fatalError("go.NativePointer supports only Int, String, and Bool", Context.currentPos());
			return "*" + nativeElement;
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
				"*hxrt.Function[" + goFunctionType(args, returnType, classTypeName, enumTypeName) + "]";
			case TInst(classRef, params):
				var classType = classRef.get();
				final carrierGoType = nativeScalarCarrierClassGoType(classType);
				if (carrierGoType != null) {
					carrierGoType;
				} else if (isTypeParameterClass(classType)) {
					typeParameterGoType(classType, classTypeName, enumTypeName);
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
				"hxrt.Anonymous";
			case TAbstract(abstractRef, params):
				var abstractType = abstractDeclaration(abstractRef);
				var mapped = abstractGoType(abstractType, params, classTypeName, enumTypeName);
				mapped == null ? "any" : mapped;
			case _:
				"any";
		};
	}

	public static function isStringType(type:Type):Bool {
		var followed = Context.follow(type);
		return switch (followed) {
			case TInst(classRef, _): declarationPath(classRef) == "String";
			case TAbstract(abstractRef, _): declarationPath(abstractRef) == "String";
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
				var abstractType = abstractDeclaration(abstractRef);
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
			case TAbstract(abstractRef, _): declarationPath(abstractRef) == "Bool";
			case _:
				false;
		};
	}

	public static function isIntType(type:Type):Bool {
		var followed = Context.follow(type);
		return switch (followed) {
			case TAbstract(abstractRef, _): var path = declarationPath(abstractRef); path == "Int" || path == "UInt";
			case _:
				false;
		};
	}

	public static function isNullableIntType(type:Type):Bool {
		return switch (type) {
			case TAbstract(abstractRef, params): declarationPath(abstractRef) == "Null" && params.length == 1 && isIntType(params[0]);
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
			case TAbstract(abstractRef, params): declarationPath(abstractRef) == "Null" && params.length == 1 && isFloatType(params[0]);
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
			case TAbstract(abstractRef, params): declarationPath(abstractRef) == "Null" && params.length == 1 && isBoolType(params[0]);
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
			case TAbstract(abstractRef, params): declarationPath(abstractRef) == "Null" && params.length == 1 && isPrimitiveValueType(params[0]);
			case TMono(ref): var resolved = ref.get(); resolved != null && isNullablePrimitiveType(resolved);
			case TType(_, _): isNullablePrimitiveType(TypeTools.follow(type, true));
			case TLazy(f):
				isNullablePrimitiveType(f());
			case _:
				false;
		};
	}

	/**
		What: The exact Go scalar type that one opaque `go.Native*` carrier class lowers to.
		Why: Haxe `Int` lowers to Go `int`, and Go never converts implicitly between
		integer types, so each exact Go width needs its own carrier. One table keeps
		the type lowering, the `null` rejections and the explicit `toInt`/`fromInt`
		conversions of every carrier in agreement.
		How: Takes the class name inside package `go`, such as `NativeUint32`, and
		returns its Go type, or null when the class is not a scalar carrier.
	**/
	public static function nativeScalarCarrierGoType(className:String):Null<String> {
		return switch (className) {
			case "NativeInt64": "int64";
			case "NativeByte": "byte";
			case "NativeUintptr": "uintptr";
			case "NativeUint32": "uint32";
			case "NativeUint16": "uint16";
			case "NativeUint": "uint";
			case _: null;
		};
	}

	/** The carrier class name of a `go.Native*` scalar type, such as `NativeUint32`, or null. */
	public static function nativeScalarCarrierName(type:Type):Null<String> {
		return switch (Context.follow(type)) {
			case TInst(ref, _):
				final classType = ref.get();
				nativeScalarCarrierClassGoType(classType) == null ? null : classType.name;
			case _: null;
		};
	}

	/** The Go type of a `go.Native*` scalar carrier class, or null for any other class. */
	public static function nativeScalarCarrierClassGoType(classType:ClassType):Null<String> {
		return classType.pack.length == 1 && classType.pack[0] == "go" ? nativeScalarCarrierGoType(classType.name) : null;
	}

	/** A native scalar carrier is a Go value without nil; `Null<T>` of it needs a separate carrier. */
	static function nullableNativeScalarCarrierName(type:Type):Null<String> {
		return switch (type) {
			case TAbstract(ref, params): declarationPath(ref) == "Null" && params.length == 1 ? nativeScalarCarrierName(params[0]) : null;
			case TMono(ref):
				var resolved = ref.get();
				resolved == null ? null : nullableNativeScalarCarrierName(resolved);
			case TType(_, _): nullableNativeScalarCarrierName(TypeTools.follow(type, true));
			case TLazy(f): nullableNativeScalarCarrierName(f());
			case _: null;
		};
	}

	static function rejectNullableNativeScalarCarrier(type:Type):Void {
		final carrier = nullableNativeScalarCarrierName(type);
		if (carrier != null) {
			Context.fatalError("Null<go." + carrier + "> has no native Go " + nativeScalarCarrierGoType(carrier)
				+ " representation; use a typed optional boundary",
				Context.currentPos());
		}
	}

	public static function isHaxeInt32Type(type:Type):Bool {
		return switch (type) {
			case TAbstract(abstractRef, _): declarationPath(abstractRef) == "haxe.Int32";
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
			case TAbstract(abstractRef, _): declarationPath(abstractRef) == "Float";
			case _:
				false;
		};
	}

	public static function isStdClassMetaType(type:Type):Bool {
		var followed = Context.follow(type);
		return switch (followed) {
			case TAbstract(abstractRef, _): declarationPath(abstractRef) == "Class";
			case _:
				false;
		};
	}

	public static function isStdEnumMetaType(type:Type):Bool {
		var followed = Context.follow(type);
		return switch (followed) {
			case TAbstract(abstractRef, _): declarationPath(abstractRef) == "Enum";
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

	/** Merge only fully structural constraint sets; nominal constraints keep their owner. */
	static function anonymousConstraintFields(constraints:Array<Type>):Null<Array<ClassField>> {
		if (constraints == null || constraints.length == 0)
			return null;
		var fields = new Array<ClassField>();
		var names = new Map<String, Bool>();
		for (constraint in constraints) {
			var current = switch (Context.follow(constraint)) {
				case TAnonymous(ref): ref.get().fields;
				case _: null;
			};
			if (current == null)
				return null;
			for (field in current)
				if (!names.exists(field.name)) {
					names.set(field.name, true);
					fields.push(field);
				}
		}
		// An empty constraint admits nominal objects; it promises no structural slots.
		return fields.length == 0 ? null : fields;
	}

	/**
		What: Reports whether a type uses the structural (anonymous object) storage
		contract. Concrete structures and structural type constraints share it.
		Why: Callers ask this for many values and only need yes or no. Reading a
		structure's fields would rebuild every field declaration on each call.
		How: Every structure qualifies, whatever its fields. A type parameter
		qualifies only when all its constraints are structural and together promise
		at least one field.
	**/
	public static function isAnonymousObjectType(type:Type):Bool {
		return switch (Context.follow(type)) {
			case TAnonymous(_): true;
			case TInst(ref, _): switch (ref.get().kind) {
					case KTypeParameter(constraints): anonymousConstraintFields(constraints) != null;
					case _: false;
				}
			case _: false;
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
			case TInst(classRef, _): declarationPath(classRef) == "Array";
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
		rejectNullableNativeScalarCarrier(type);
		if (isBytesDataType(type)) {
			return "[]int";
		}
		var nativePointerElement = nativePointerElementType(type);
		if (nativePointerElement != null) {
			var nativeElement = nativeScalarPointerElementGoType(nativePointerElement);
			if (nativeElement == null)
				Context.fatalError("go.NativePointer supports only Int, String, and Bool", Context.currentPos());
			return "*" + nativeElement;
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
				"*hxrt.Function[" + goFunctionType(args, returnType, classTypeName, enumTypeName) + "]";
			case TInst(classRef, params):
				var classType = classRef.get();
				final carrierGoType = nativeScalarCarrierClassGoType(classType);
				if (carrierGoType != null) {
					carrierGoType;
				} else if (isTypeParameterClass(classType)) {
					typeParameterGoType(classType, classTypeName, enumTypeName);
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
				"hxrt.Anonymous";
			case TAbstract(abstractRef, params):
				var abstractType = abstractDeclaration(abstractRef);
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
		if (abstractType.meta.has(":coreType")) {
			return null;
		}
		var backingType = TypeTools.applyTypeParameters(abstractType.type, abstractType.params, params);
		return hasOrdinaryStringBacking(backingType, new Map<String, Bool>()) ? "*string" : null;
	}

	/**
		What: Proves that an abstract stores a Haxe string.

		Why: Haxe stores an abstract's value as its underlying type, so an
		abstract over String must share String's pointer carrier in every
		position (parameters, returns, locals, fields, Array elements, Map
		values). The callers that convert values between String and the abstract
		assume that carrier. When the abstract fell back to `any`, a function
		like `function f(n:Name):String return n;` produced Go that did not
		compile. Only `@:coreType` abstracts choose their own carrier.

		The implementation class's `isExtern` flag does not mark a native
		boundary. Haxe's dead code elimination sets it on every implementation
		class that has no fields left (for example
		`abstract Name(String) from String to String {}`) and on every enum
		abstract whose fields are all enum values. It runs even with `-dce no`.
		An `extern abstract` is also only an abstract whose implementation class
		is not generated; its values are still values of the underlying type.
		A typed Go-native carrier comes from `externBackedAbstractGoType`.

		How: Substitute applied parameters along abstract chains, stopping at
		`@:coreType` abstracts and repeated applied types. Abstracts over Int,
		Float, Bool, Dynamic or an unapplied type parameter are not String-backed
		and keep their storage.
	**/
	static function hasOrdinaryStringBacking(type:Type, seen:Map<String, Bool>):Bool {
		return switch (type) {
			case TInst(classRef, _): declarationPath(classRef) == "String";
			case TType(typeRef, params):
				var typedefType = typeRef.get();
				hasOrdinaryStringBacking(TypeTools.applyTypeParameters(typedefType.type, typedefType.params, params), seen);
			case TMono(ref): var resolved = ref.get(); resolved != null && hasOrdinaryStringBacking(resolved, seen);
			case TLazy(resolve):
				hasOrdinaryStringBacking(resolve(), seen);
			case TAbstract(abstractRef, params):
				var abstractType = abstractDeclaration(abstractRef);
				var identity = TypeTools.toString(type);
				if (abstractType.pack.length == 0 && abstractType.name == "String") {
					true;
				} else if (seen.exists(identity)) {
					false;
				} else if (abstractType.pack.length == 0 && abstractType.name == "Null" && params.length == 1) {
					seen.set(identity, true);
					hasOrdinaryStringBacking(params[0], seen);
				} else if (abstractType.meta.has(":coreType")) {
					false;
				} else {
					seen.set(identity, true);
					var backingType = TypeTools.applyTypeParameters(abstractType.type, abstractType.params, params);
					hasOrdinaryStringBacking(backingType, seen);
				}
			case _:
				false;
		};
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
				externBackedAbstractGoType(abstractDeclaration(innerRef), classTypeName, seen);
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
		return "func(" + params + ") " + functionResultStorageGoType(returnType, classTypeName, enumTypeName);
	}

	/**
		What: Selects the result storage shared by declarations and callable carriers.
		Why: Scalar mapping erases Null, but callable results must retain its nil state.
		How: Preserve nullable primitives in the runtime boundary's nil-capable storage.
	**/
	public static function functionResultStorageGoType(type:Type, classTypeName:GoClassTypeNamer, enumTypeName:GoEnumTypeNamer):String {
		return isNullablePrimitiveType(type) ? "any" : scalarGoType(type, classTypeName, enumTypeName);
	}

	/**
		What: Maps a declared Haxe function parameter to its callable Go storage type.

		Why: Haxe marks both `?timeout:Int` and `timeout:Int = 10` as optional.
		The generated function literal and every variable or field carrying it must
		therefore share one signature. Mapping only the carrier as `func(any)` makes a
		non-null-default implementation such as `func(int)` fail Go type checking.

		How: Keep pointer/nil-capable mappings unchanged and widen nullable primitives
		and optional scalar values to `any`. `lowerFunctionParams` uses the same rule
		and asserts the scalar type when the function body reads the parameter.
	**/
	public static function functionParameterStorageGoType(arg:{name:String, opt:Bool, t:Type}, classTypeName:GoClassTypeNamer,
			enumTypeName:GoEnumTypeNamer):String {
		return isNullablePrimitiveType(arg.t)
			|| (arg.opt && isPrimitiveValueType(arg.t)) ? "any" : scalarGoType(arg.t, classTypeName, enumTypeName);
	}

	/** Resolves the element type of the explicit `go.NativeSlice<T>` boundary. */
	public static function nativeSliceElementType(type:Type):Null<Type> {
		var followed = Context.follow(type);
		return switch (followed) {
			case TInst(classRef, params):
				if (declarationPath(classRef) == "go.NativeSlice" && params.length == 1) {
					params[0];
				} else {
					null;
				}
			case _:
				null;
		};
	}

	/**
		What: Resolves the element of an explicit native scalar pointer carrier.
		Why: The Go ABI must use the declared element type at every type position,
			not infer a destination from the expression passed to a variadic call.
		How: Admit only go.NativePointer with one type parameter; other types do
			not participate in this pointer mapping.
	**/
	public static function nativePointerElementType(type:Type):Null<Type> {
		return switch (Context.follow(type)) {
			case TInst(classRef, params): declarationPath(classRef) == "go.NativePointer" && params.length == 1 ? params[0] : null;
			case _:
				null;
		};
	}

	/**
		What: Maps the admitted Haxe scalar elements to native Go element types.
		Why: A broader mapping could silently produce the wrong pointer destination
			for a Go API that mutates through its exact scalar ABI.
		How: Reject nullable and non-scalar elements before Go type emission.
	**/
	public static function nativeScalarPointerElementGoType(type:Type):Null<String> {
		if (isNullablePrimitiveType(type))
			return null;
		return switch (Context.follow(type)) {
			case TAbstract(abstractRef, _):
				switch (declarationPath(abstractRef)) {
					case "Int": "int";
					case "Bool": "bool";
					case _: null;
				};
			case TInst(classRef, _): declarationPath(classRef) == "String" ? "string" : null;
			case _:
				null;
		};
	}

	/** True only for the explicit native Go `[]string` boundary. */
	public static function isNativeStringSliceType(type:Type):Bool {
		return switch (Context.follow(type)) {
			case TInst(classRef, _): declarationPath(classRef) == "go.NativeStringSlice";
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
			case TType(typeRef, _): declarationPath(typeRef) == "haxe.io.BytesData";
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
				if (declarationPath(abstractRef) == "haxe.Rest" && params.length == 1) {
					params[0];
				} else {
					null;
				}
			case TType(typeRef, params):
				if (declarationPath(typeRef) == "haxe._Rest.NativeRest" && params.length == 1) {
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
				if (declarationPath(abstractRef) == "haxe.ds.Vector" && params.length == 1) {
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
				if (declarationPath(abstractRef) == "haxe.ds.ReadOnlyArray" && params.length == 1) {
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
			case TAbstract(abstractRef, _): declarationPath(abstractRef) == "Void";
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
			case TAbstract(abstractRef, _): declarationPath(abstractRef) == "Dynamic";
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
		What: Map structural type constraints to the shared anonymous boundary.
		Why: A constraint admits both generated methods and map-backed callbacks;
		a Go method interface cannot represent both without losing source semantics.
		How: Recognize fully anonymous constraint sets and keep other erasure intact.
	**/
	static function typeParameterGoType(classType:ClassType, classTypeName:GoClassTypeNamer, enumTypeName:GoEnumTypeNamer):String {
		return switch (classType.kind) {
			case KTypeParameter(constraints): anonymousConstraintFields(constraints) == null ? "any" : "hxrt.Anonymous";
			case _: "any";
		};
	}

	public static function isHaxeExceptionClass(classType:ClassType):Bool {
		return switch (classType.pack.join(".") + "." + classType.name) {
			case "haxe.Exception":
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
