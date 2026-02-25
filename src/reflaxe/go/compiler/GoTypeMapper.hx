package reflaxe.go.compiler;

#if macro
import haxe.macro.Context;
import haxe.macro.Type;

typedef GoClassTypeNamer = ClassType->String;
typedef GoEnumTypeNamer = EnumType->String;

class GoTypeMapper {
	public static function typeToGoType(type:Type, classTypeName:GoClassTypeNamer, enumTypeName:GoEnumTypeNamer):String {
		var restElement = restElementType(type);
		if (restElement != null) {
			return "[]" + scalarGoType(restElement, classTypeName, enumTypeName);
		}
		var vectorElement = vectorElementType(type);
		if (vectorElement != null) {
			return "[]" + scalarGoType(vectorElement, classTypeName, enumTypeName);
		}
		var readOnlyElement = readOnlyArrayElementType(type);
		if (readOnlyElement != null) {
			return "[]" + scalarGoType(readOnlyElement, classTypeName, enumTypeName);
		}

		var followed = Context.follow(type);
		return switch (followed) {
			case TFun(args, returnType):
				goFunctionType(args, returnType, classTypeName, enumTypeName);
			case TInst(classRef, params):
				var classType = classRef.get();
				if (isTypeParameterClass(classType)) {
					"any";
				} else if (isHaxeExceptionClass(classType)) {
					"*hxrt.ExceptionValue";
				} else if (isHaxeIoBaseClass(classType)) {
					classTypeName(classType);
				} else if (classType.isInterface) {
					classTypeName(classType);
				} else if (classType.pack.length == 0 && classType.name == "String") {
					"*string";
				} else if (classType.pack.length == 0 && classType.name == "Array" && params.length == 1) {
					"[]" + scalarGoType(params[0], classTypeName, enumTypeName);
				} else {
					"*" + classTypeName(classType);
				}
			case TEnum(enumRef, _):
				"*" + enumTypeName(enumRef.get());
			case TAnonymous(_):
				"map[string]any";
			case TAbstract(abstractRef, _):
				var abstractType = abstractRef.get();
				if (abstractType.pack.length == 0 && abstractType.name == "Int") {
					"int";
				} else if (abstractType.pack.length == 0 && abstractType.name == "Float") {
					"float64";
				} else if (abstractType.pack.length == 0 && abstractType.name == "Bool") {
					"bool";
				} else if (abstractType.pack.length == 0 && abstractType.name == "String") {
					"*string";
				} else if (abstractType.pack.join(".") == "haxe" && abstractType.name == "Int32") {
					"int";
				} else if (abstractType.pack.join(".") == "haxe" && abstractType.name == "Int64") {
					"*haxe___Int64_____Int64";
				} else {
					"any";
				}
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
			case TAbstract(abstractRef, _): var abstractType = abstractRef.get(); abstractType.pack.length == 0 && abstractType.name == "Int";
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
		return isBoolType(type) || isIntType(type) || isFloatType(type);
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
		if (restElementType(type) != null) {
			return true;
		}
		if (vectorElementType(type) != null) {
			return true;
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
		var restElement = restElementType(type);
		if (restElement != null) {
			return scalarGoType(restElement, classTypeName, enumTypeName);
		}
		var vectorElement = vectorElementType(type);
		if (vectorElement != null) {
			return scalarGoType(vectorElement, classTypeName, enumTypeName);
		}
		var readOnlyElement = readOnlyArrayElementType(type);
		if (readOnlyElement != null) {
			return scalarGoType(readOnlyElement, classTypeName, enumTypeName);
		}

		var followed = Context.follow(type);
		return switch (followed) {
			case TInst(classRef, params):
				var classType = classRef.get();
				if (classType.pack.length == 0 && classType.name == "Array" && params.length == 1) {
					scalarGoType(params[0], classTypeName, enumTypeName);
				} else {
					"any";
				}
			case _:
				"any";
		};
	}

	public static function scalarGoType(type:Type, classTypeName:GoClassTypeNamer, enumTypeName:GoEnumTypeNamer):String {
		var restElement = restElementType(type);
		if (restElement != null) {
			return "[]" + scalarGoType(restElement, classTypeName, enumTypeName);
		}
		var vectorElement = vectorElementType(type);
		if (vectorElement != null) {
			return "[]" + scalarGoType(vectorElement, classTypeName, enumTypeName);
		}
		var readOnlyElement = readOnlyArrayElementType(type);
		if (readOnlyElement != null) {
			return "[]" + scalarGoType(readOnlyElement, classTypeName, enumTypeName);
		}

		var followed = Context.follow(type);
		return switch (followed) {
			case TFun(args, returnType):
				goFunctionType(args, returnType, classTypeName, enumTypeName);
			case TInst(classRef, params):
				var classType = classRef.get();
				if (isTypeParameterClass(classType)) {
					"any";
				} else if (isHaxeExceptionClass(classType)) {
					"*hxrt.ExceptionValue";
				} else if (isHaxeIoBaseClass(classType)) {
					classTypeName(classType);
				} else if (classType.isInterface) {
					classTypeName(classType);
				} else if (classType.pack.length == 0 && classType.name == "String") {
					"*string";
				} else if (classType.pack.length == 0 && classType.name == "Array" && params.length == 1) {
					"[]" + scalarGoType(params[0], classTypeName, enumTypeName);
				} else {
					"*" + classTypeName(classType);
				}
			case TEnum(enumRef, _):
				"*" + enumTypeName(enumRef.get());
			case TAnonymous(_):
				"map[string]any";
			case TAbstract(abstractRef, _):
				var abstractType = abstractRef.get();
				if (abstractType.pack.length == 0 && abstractType.name == "Int") {
					"int";
				} else if (abstractType.pack.length == 0 && abstractType.name == "Float") {
					"float64";
				} else if (abstractType.pack.length == 0 && abstractType.name == "Bool") {
					"bool";
				} else if (abstractType.pack.length == 0 && abstractType.name == "String") {
					"*string";
				} else if (abstractType.pack.join(".") == "haxe" && abstractType.name == "Int32") {
					"int";
				} else if (abstractType.pack.join(".") == "haxe" && abstractType.name == "Int64") {
					"*haxe___Int64_____Int64";
				} else {
					"any";
				}
			case _:
				"any";
		};
	}

	public static function goFunctionType(args:Array<{name:String, opt:Bool, t:Type}>, returnType:Type, classTypeName:GoClassTypeNamer,
			enumTypeName:GoEnumTypeNamer):String {
		var params = [for (arg in args) scalarGoType(arg.t, classTypeName, enumTypeName)].join(", ");
		if (isVoidType(returnType)) {
			return "func(" + params + ")";
		}
		return "func(" + params + ") " + scalarGoType(returnType, classTypeName, enumTypeName);
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

	public static function isHaxeExceptionClass(classType:ClassType):Bool {
		return classType.pack.join(".") == "haxe" && classType.name == "Exception";
	}

	public static function isHaxeIoBaseClass(classType:ClassType):Bool {
		return classType.pack.join(".") == "haxe.io" && (classType.name == "Input" || classType.name == "Output");
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
