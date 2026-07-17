package reflaxe.go.compiler.emit;

#if macro
import reflaxe.go.ast.GoAST.GoDecl;
import reflaxe.go.ast.GoAST.GoExpr;
import reflaxe.go.ast.GoAST.GoStmt;
import reflaxe.go.compiler.GoStdlibOwnership;

/**
	What:
	Builds the compiler-owned Go declarations behind `Type.*` reflection and
	runtime class/enum construction support.

	Why:
	This surface is legitimately compiler-owned because it depends on generated
	constructor symbols and backend representation details. Keeping the declaration
	builder in its own module gives parity work an explicit insertion point outside
	the `GoCompiler` monolith.

	How:
	Consumes precomputed class/enum metadata plus small formatting callbacks from
	`GoCompiler`, then emits the same `GoDecl` set that was previously built
	inline. This is a relocation only; no semantics change here.
**/
class GoTypeReflectionEmitter {
	public static function emit(classMetadata:Array<{
		final goTypeName:String;
		final haxeTypeName:String;
		final constructorSymbol:String;
		final constructible:Bool;
		final superHaxeTypeName:Null<String>;
		final staticFieldNames:Array<String>;
		final instanceFieldNames:Array<String>;
	}>, enumMetadata:Array<{
		final goTypeName:String;
		final haxeTypeName:String;
		final constructors:Array<{
			final name:String;
			final index:Int;
			final symbol:String;
			final arity:Int;
		}>;
	}>, goRawQuotedString:String->String,
			goStringArrayCarrierLiteral:Array<String>->String):Array<GoDecl> {
		var allEnumMetadata = enumMetadata.copy();
		allEnumMetadata.push({
			goTypeName: "ValueType",
			haxeTypeName: "ValueType",
			constructors: [
				{
					name: "TNull",
					index: 0,
					symbol: "ValueType_TNull",
					arity: 0
				},
				{
					name: "TInt",
					index: 1,
					symbol: "ValueType_TInt",
					arity: 0
				},
				{
					name: "TFloat",
					index: 2,
					symbol: "ValueType_TFloat",
					arity: 0
				},
				{
					name: "TBool",
					index: 3,
					symbol: "ValueType_TBool",
					arity: 0
				},
				{
					name: "TObject",
					index: 4,
					symbol: "ValueType_TObject",
					arity: 0
				},
				{
					name: "TFunction",
					index: 5,
					symbol: "ValueType_TFunction",
					arity: 0
				},
				{
					name: "TClass",
					index: 6,
					symbol: "ValueType_TClass",
					arity: 1
				},
				{
					name: "TEnum",
					index: 7,
					symbol: "ValueType_TEnum",
					arity: 1
				},
				{
					name: "TUnknown",
					index: 8,
					symbol: "ValueType_TUnknown",
					arity: 0
				}
			]
		});
		allEnumMetadata.sort(function(a, b) return Reflect.compare(a.goTypeName, b.goTypeName));

		var classResolveBody = [
			GoStmt.GoRaw("if name == nil {"),
			GoStmt.GoRaw("\treturn nil"),
			GoStmt.GoRaw("}")
		];
		classResolveBody.push(GoStmt.GoRaw("rawName := *hxrt.StdString(name)"));
		classResolveBody.push(GoStmt.GoRaw("switch rawName {"));
		for (entry in classMetadata) {
			classResolveBody.push(GoStmt.GoRaw("case " + goRawQuotedString(entry.haxeTypeName) + ":"));
			classResolveBody.push(GoStmt.GoRaw("\treturn &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(rawName)}"));
		}
		classResolveBody.push(GoStmt.GoRaw("default:"));
		classResolveBody.push(GoStmt.GoRaw("\treturn nil"));
		classResolveBody.push(GoStmt.GoRaw("}"));

		var enumResolveBody = [
			GoStmt.GoRaw("if name == nil {"),
			GoStmt.GoRaw("\treturn nil"),
			GoStmt.GoRaw("}")
		];
		enumResolveBody.push(GoStmt.GoRaw("rawName := *hxrt.StdString(name)"));
		enumResolveBody.push(GoStmt.GoRaw("switch rawName {"));
		for (entry in allEnumMetadata) {
			enumResolveBody.push(GoStmt.GoRaw("case " + goRawQuotedString(entry.haxeTypeName) + ":"));
			enumResolveBody.push(GoStmt.GoRaw("\treturn &hxrt__TypeEnumValue{name: hxrt.StringFromLiteral(rawName)}"));
		}
		enumResolveBody.push(GoStmt.GoRaw("default:"));
		enumResolveBody.push(GoStmt.GoRaw("\treturn nil"));
		enumResolveBody.push(GoStmt.GoRaw("}"));

		var classCreateBody = [GoStmt.GoRaw("switch className {")];
		for (entry in classMetadata) {
			classCreateBody.push(GoStmt.GoRaw("case " + goRawQuotedString(entry.haxeTypeName) + ":"));
			if (entry.constructible) {
				classCreateBody.push(GoStmt.GoRaw("\treturn hxrt_typeCallAny(" + entry.constructorSymbol + ", args)"));
			} else {
				classCreateBody.push(GoStmt.GoRaw("\treturn nil, false"));
			}
		}
		classCreateBody.push(GoStmt.GoRaw("default:"));
		classCreateBody.push(GoStmt.GoRaw("\treturn nil, false"));
		classCreateBody.push(GoStmt.GoRaw("}"));

		var enumCreateBody = [GoStmt.GoRaw("switch enumName {")];
		for (entry in allEnumMetadata) {
			enumCreateBody.push(GoStmt.GoRaw("case " + goRawQuotedString(entry.haxeTypeName) + ":"));
			enumCreateBody.push(GoStmt.GoRaw("\tif useIndex {"));
			enumCreateBody.push(GoStmt.GoRaw("\t\tswitch constructorIndex {"));
			for (constructor in entry.constructors) {
				enumCreateBody.push(GoStmt.GoRaw("\t\tcase " + constructor.index + ":"));
				enumCreateBody.push(GoStmt.GoRaw("\t\t\tif len(args) != " + constructor.arity + " {"));
				enumCreateBody.push(GoStmt.GoRaw("\t\t\t\treturn nil, false"));
				enumCreateBody.push(GoStmt.GoRaw("\t\t\t}"));
				if (constructor.arity == 0) {
					enumCreateBody.push(GoStmt.GoRaw("\t\t\treturn " + constructor.symbol + ", true"));
				} else {
					enumCreateBody.push(GoStmt.GoRaw("\t\t\treturn hxrt_typeCallAny(" + constructor.symbol + ", args)"));
				}
			}
			enumCreateBody.push(GoStmt.GoRaw("\t\tdefault:"));
			enumCreateBody.push(GoStmt.GoRaw("\t\t\treturn nil, false"));
			enumCreateBody.push(GoStmt.GoRaw("\t\t}"));
			enumCreateBody.push(GoStmt.GoRaw("\t}"));
			enumCreateBody.push(GoStmt.GoRaw("\tswitch constructorName {"));
			for (constructor in entry.constructors) {
				enumCreateBody.push(GoStmt.GoRaw("\tcase " + goRawQuotedString(constructor.name) + ":"));
				enumCreateBody.push(GoStmt.GoRaw("\t\tif len(args) != " + constructor.arity + " {"));
				enumCreateBody.push(GoStmt.GoRaw("\t\t\treturn nil, false"));
				enumCreateBody.push(GoStmt.GoRaw("\t\t}"));
				if (constructor.arity == 0) {
					enumCreateBody.push(GoStmt.GoRaw("\t\treturn " + constructor.symbol + ", true"));
				} else {
					enumCreateBody.push(GoStmt.GoRaw("\t\treturn hxrt_typeCallAny(" + constructor.symbol + ", args)"));
				}
			}
			enumCreateBody.push(GoStmt.GoRaw("\tdefault:"));
			enumCreateBody.push(GoStmt.GoRaw("\t\treturn nil, false"));
			enumCreateBody.push(GoStmt.GoRaw("\t}"));
		}
		enumCreateBody.push(GoStmt.GoRaw("default:"));
		enumCreateBody.push(GoStmt.GoRaw("\treturn nil, false"));
		enumCreateBody.push(GoStmt.GoRaw("}"));

		var enumConstructorBody = new Array<GoStmt>();
		if (allEnumMetadata.length == 0) {
			enumConstructorBody = [
				GoStmt.GoRaw("if hxrt.AnyEqualsNull(e) {"),
				GoStmt.GoRaw("\treturn nil"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("return nil")
			];
		} else {
			enumConstructorBody = [
				GoStmt.GoRaw("if hxrt.AnyEqualsNull(e) {"),
				GoStmt.GoRaw("\treturn nil"),
				GoStmt.GoRaw("}")
			];
			enumConstructorBody.push(GoStmt.GoRaw("switch value := e.(type) {"));
			for (entry in allEnumMetadata) {
				enumConstructorBody.push(GoStmt.GoRaw("case *" + entry.goTypeName + ":"));
				enumConstructorBody.push(GoStmt.GoRaw("\tif value == nil {"));
				enumConstructorBody.push(GoStmt.GoRaw("\t\treturn nil"));
				enumConstructorBody.push(GoStmt.GoRaw("\t}"));
				enumConstructorBody.push(GoStmt.GoRaw("\tswitch value.tag {"));
				for (constructor in entry.constructors) {
					enumConstructorBody.push(GoStmt.GoRaw("\tcase " + constructor.index + ":"));
					enumConstructorBody.push(GoStmt.GoRaw("\t\treturn hxrt.StringFromLiteral(" + goRawQuotedString(constructor.name) + ")"));
				}
				enumConstructorBody.push(GoStmt.GoRaw("\tdefault:"));
				enumConstructorBody.push(GoStmt.GoRaw("\t\treturn nil"));
				enumConstructorBody.push(GoStmt.GoRaw("\t}"));
			}
			enumConstructorBody.push(GoStmt.GoRaw("default:"));
			enumConstructorBody.push(GoStmt.GoRaw("\treturn nil"));
			enumConstructorBody.push(GoStmt.GoRaw("}"));
		}

		var enumIndexBody = new Array<GoStmt>();
		if (allEnumMetadata.length == 0) {
			enumIndexBody = [
				GoStmt.GoRaw("if hxrt.AnyEqualsNull(e) {"),
				GoStmt.GoRaw("\treturn -1"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("return -1")
			];
		} else {
			enumIndexBody = [
				GoStmt.GoRaw("if hxrt.AnyEqualsNull(e) {"),
				GoStmt.GoRaw("\treturn -1"),
				GoStmt.GoRaw("}")
			];
			enumIndexBody.push(GoStmt.GoRaw("switch value := e.(type) {"));
			for (entry in allEnumMetadata) {
				enumIndexBody.push(GoStmt.GoRaw("case *" + entry.goTypeName + ":"));
				enumIndexBody.push(GoStmt.GoRaw("\tif value == nil {"));
				enumIndexBody.push(GoStmt.GoRaw("\t\treturn -1"));
				enumIndexBody.push(GoStmt.GoRaw("\t}"));
				enumIndexBody.push(GoStmt.GoRaw("\treturn value.tag"));
			}
			enumIndexBody.push(GoStmt.GoRaw("default:"));
			enumIndexBody.push(GoStmt.GoRaw("\treturn -1"));
			enumIndexBody.push(GoStmt.GoRaw("}"));
		}

		var enumParametersBody = new Array<GoStmt>();
		if (allEnumMetadata.length == 0) {
			enumParametersBody = [GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("hxrt.NewArray"), []))];
		} else {
			enumParametersBody = [
				GoStmt.GoIf(GoExpr.GoCall(GoExpr.GoIdent("hxrt.AnyEqualsNull"), [GoExpr.GoIdent("e")]),
					[GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("hxrt.NewArray"), []))], null)
			];
			enumParametersBody.push(GoStmt.GoRaw("switch value := e.(type) {"));
			for (entry in allEnumMetadata) {
				enumParametersBody.push(GoStmt.GoRaw("case *" + entry.goTypeName + ":"));
				enumParametersBody.push(GoStmt.GoRaw("\tif value == nil || value.params == nil {"));
				enumParametersBody.push(GoStmt.GoRaw("\t\treturn hxrt.NewArray()"));
				enumParametersBody.push(GoStmt.GoRaw("\t}"));
				enumParametersBody.push(GoStmt.GoRaw("\treturn hxrt.NewArray(value.params...)"));
			}
			enumParametersBody.push(GoStmt.GoRaw("default:"));
			enumParametersBody.push(GoStmt.GoRaw("\treturn hxrt.NewArray()"));
			enumParametersBody.push(GoStmt.GoRaw("}"));
		}

		var getClassBody = [
			GoStmt.GoRaw("if hxrt.AnyEqualsNull(o) {"),
			GoStmt.GoRaw("\treturn nil"),
			GoStmt.GoRaw("}")
		];
		getClassBody.push(GoStmt.GoRaw("switch value := o.(type) {"));
		getClassBody.push(GoStmt.GoRaw("case *hxrt__TypeClassValue:"));
		getClassBody.push(GoStmt.GoRaw("\tif value == nil {"));
		getClassBody.push(GoStmt.GoRaw("\t\treturn nil"));
		getClassBody.push(GoStmt.GoRaw("\t}"));
		getClassBody.push(GoStmt.GoRaw("\treturn value"));
		getClassBody.push(GoStmt.GoRaw("case hxrt__TypeClassValue:"));
		getClassBody.push(GoStmt.GoRaw("\tcopyValue := value"));
		getClassBody.push(GoStmt.GoRaw("\treturn &copyValue"));
		getClassBody.push(GoStmt.GoRaw("case *hxrt.Array:"));
		getClassBody.push(GoStmt.GoRaw("\tif value == nil {"));
		getClassBody.push(GoStmt.GoRaw("\t\treturn nil"));
		getClassBody.push(GoStmt.GoRaw("\t}"));
		getClassBody.push(GoStmt.GoRaw("\treturn &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(\"Array\")}"));
		for (entry in classMetadata) {
			if (!entry.constructible) {
				continue;
			}
			getClassBody.push(GoStmt.GoRaw("case *" + entry.goTypeName + ":"));
			getClassBody.push(GoStmt.GoRaw("\tif value == nil {"));
			getClassBody.push(GoStmt.GoRaw("\t\treturn nil"));
			getClassBody.push(GoStmt.GoRaw("\t}"));
			getClassBody.push(GoStmt.GoRaw("\treturn &hxrt__TypeClassValue{name: hxrt.StringFromLiteral(" + goRawQuotedString(entry.haxeTypeName) + ")}"));
		}
		getClassBody.push(GoStmt.GoRaw("default:"));
		getClassBody.push(GoStmt.GoRaw("\treturn nil"));
		getClassBody.push(GoStmt.GoRaw("}"));

		var getEnumBody = [
			GoStmt.GoRaw("if hxrt.AnyEqualsNull(o) {"),
			GoStmt.GoRaw("\treturn nil"),
			GoStmt.GoRaw("}")
		];
		getEnumBody.push(GoStmt.GoRaw("switch value := o.(type) {"));
		getEnumBody.push(GoStmt.GoRaw("case *hxrt__TypeEnumValue:"));
		getEnumBody.push(GoStmt.GoRaw("\tif value == nil {"));
		getEnumBody.push(GoStmt.GoRaw("\t\treturn nil"));
		getEnumBody.push(GoStmt.GoRaw("\t}"));
		getEnumBody.push(GoStmt.GoRaw("\treturn value"));
		getEnumBody.push(GoStmt.GoRaw("case hxrt__TypeEnumValue:"));
		getEnumBody.push(GoStmt.GoRaw("\tcopyValue := value"));
		getEnumBody.push(GoStmt.GoRaw("\treturn &copyValue"));
		for (entry in allEnumMetadata) {
			getEnumBody.push(GoStmt.GoRaw("case *" + entry.goTypeName + ":"));
			getEnumBody.push(GoStmt.GoRaw("\tif value == nil {"));
			getEnumBody.push(GoStmt.GoRaw("\t\treturn nil"));
			getEnumBody.push(GoStmt.GoRaw("\t}"));
			getEnumBody.push(GoStmt.GoRaw("\treturn &hxrt__TypeEnumValue{name: hxrt.StringFromLiteral(" + goRawQuotedString(entry.haxeTypeName) + ")}"));
		}
		getEnumBody.push(GoStmt.GoRaw("default:"));
		getEnumBody.push(GoStmt.GoRaw("\treturn nil"));
		getEnumBody.push(GoStmt.GoRaw("}"));

		var getSuperClassBody = [
			GoStmt.GoRaw("className, ok := hxrt_typeResolvedClassName(c)"),
			GoStmt.GoRaw("if !ok {"),
			GoStmt.GoRaw("\treturn nil"),
			GoStmt.GoRaw("}"),
			GoStmt.GoRaw("switch className {")
		];
		for (entry in classMetadata) {
			getSuperClassBody.push(GoStmt.GoRaw("case " + goRawQuotedString(entry.haxeTypeName) + ":"));
			if (entry.superHaxeTypeName == null) {
				getSuperClassBody.push(GoStmt.GoRaw("\treturn nil"));
			} else {
				getSuperClassBody.push(GoStmt.GoRaw("\treturn &hxrt__TypeClassValue{name: hxrt.StringFromLiteral("
					+ goRawQuotedString(entry.superHaxeTypeName)
					+ ")}"));
			}
		}
		getSuperClassBody.push(GoStmt.GoRaw("default:"));
		getSuperClassBody.push(GoStmt.GoRaw("\treturn nil"));
		getSuperClassBody.push(GoStmt.GoRaw("}"));

		var getClassFieldsBody = [
			GoStmt.GoRaw("className, ok := hxrt_typeResolvedClassName(c)"),
			GoStmt.GoRaw("if !ok {"),
			GoStmt.GoRaw("\treturn hxrt.NewArray()"),
			GoStmt.GoRaw("}"),
			GoStmt.GoRaw("switch className {")
		];
		for (entry in classMetadata) {
			getClassFieldsBody.push(GoStmt.GoRaw("case " + goRawQuotedString(entry.haxeTypeName) + ":"));
			getClassFieldsBody.push(GoStmt.GoRaw("\treturn " + goStringArrayCarrierLiteral(entry.staticFieldNames)));
		}
		getClassFieldsBody.push(GoStmt.GoRaw("default:"));
		getClassFieldsBody.push(GoStmt.GoRaw("\treturn hxrt.NewArray()"));
		getClassFieldsBody.push(GoStmt.GoRaw("}"));

		var getInstanceFieldsBody = [
			GoStmt.GoRaw("className, ok := hxrt_typeResolvedClassName(c)"),
			GoStmt.GoRaw("if !ok {"),
			GoStmt.GoRaw("\treturn hxrt.NewArray()"),
			GoStmt.GoRaw("}"),
			GoStmt.GoRaw("switch className {")
		];
		for (entry in classMetadata) {
			getInstanceFieldsBody.push(GoStmt.GoRaw("case " + goRawQuotedString(entry.haxeTypeName) + ":"));
			getInstanceFieldsBody.push(GoStmt.GoRaw("\treturn " + goStringArrayCarrierLiteral(entry.instanceFieldNames)));
		}
		getInstanceFieldsBody.push(GoStmt.GoRaw("default:"));
		getInstanceFieldsBody.push(GoStmt.GoRaw("\treturn hxrt.NewArray()"));
		getInstanceFieldsBody.push(GoStmt.GoRaw("}"));

		var getEnumConstructsBody = [
			GoStmt.GoRaw("enumName, ok := hxrt_typeResolvedEnumName(e)"),
			GoStmt.GoRaw("if !ok {"),
			GoStmt.GoRaw("\treturn hxrt.NewArray()"),
			GoStmt.GoRaw("}"),
			GoStmt.GoRaw("switch enumName {")
		];
		for (entry in allEnumMetadata) {
			getEnumConstructsBody.push(GoStmt.GoRaw("case " + goRawQuotedString(entry.haxeTypeName) + ":"));
			var constructorNames = [for (constructor in entry.constructors) constructor.name];
			getEnumConstructsBody.push(GoStmt.GoRaw("\treturn " + goStringArrayCarrierLiteral(constructorNames)));
		}
		getEnumConstructsBody.push(GoStmt.GoRaw("default:"));
		getEnumConstructsBody.push(GoStmt.GoRaw("\treturn hxrt.NewArray()"));
		getEnumConstructsBody.push(GoStmt.GoRaw("}"));

		var classCreateEmptyBody = [GoStmt.GoRaw("switch className {")];
		for (entry in classMetadata) {
			if (!entry.constructible) {
				continue;
			}
			if (!GoStdlibOwnership.canConstructEmptyTypeValue(entry.goTypeName)) {
				continue;
			}
			classCreateEmptyBody.push(GoStmt.GoRaw("case " + goRawQuotedString(entry.haxeTypeName) + ":"));
			classCreateEmptyBody.push(GoStmt.GoRaw("\treturn &" + entry.goTypeName + "{}, true"));
		}
		classCreateEmptyBody.push(GoStmt.GoRaw("default:"));
		classCreateEmptyBody.push(GoStmt.GoRaw("\treturn nil, false"));
		classCreateEmptyBody.push(GoStmt.GoRaw("}"));

		var allEnumsBody = [
			GoStmt.GoRaw("enumName, ok := hxrt_typeResolvedEnumName(e)"),
			GoStmt.GoRaw("if !ok {"),
			GoStmt.GoRaw("\treturn hxrt.NewArray()"),
			GoStmt.GoRaw("}"),
			GoStmt.GoRaw("switch enumName {")
		];
		for (entry in allEnumMetadata) {
			allEnumsBody.push(GoStmt.GoRaw("case " + goRawQuotedString(entry.haxeTypeName) + ":"));
			var zeroAritySymbols = [
				for (constructor in entry.constructors)
					if (constructor.arity == 0) constructor.symbol
			];
			if (zeroAritySymbols.length == 0) {
				allEnumsBody.push(GoStmt.GoRaw("\treturn hxrt.NewArray()"));
			} else {
				allEnumsBody.push(GoStmt.GoRaw("\treturn hxrt.NewArray(" + zeroAritySymbols.join(", ") + ")"));
			}
		}
		allEnumsBody.push(GoStmt.GoRaw("default:"));
		allEnumsBody.push(GoStmt.GoRaw("\treturn hxrt.NewArray()"));
		allEnumsBody.push(GoStmt.GoRaw("}"));

		var typeOfBody = [
			GoStmt.GoRaw("if hxrt.AnyEqualsNull(v) {"),
			GoStmt.GoRaw("\treturn ValueType_TNull"),
			GoStmt.GoRaw("}"),
			GoStmt.GoRaw("if enumValue := Type_getEnum(v); enumValue != nil {"),
			GoStmt.GoRaw("\treturn ValueType_TEnum(enumValue)"),
			GoStmt.GoRaw("}"),
			GoStmt.GoRaw("if classValue := Type_getClass(v); classValue != nil {"),
			GoStmt.GoRaw("\treturn ValueType_TClass(classValue)"),
			GoStmt.GoRaw("}"),
			GoStmt.GoRaw("switch v.(type) {"),
			GoStmt.GoRaw("case bool:"),
			GoStmt.GoRaw("\treturn ValueType_TBool"),
			GoStmt.GoRaw("case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, uintptr:"),
			GoStmt.GoRaw("\treturn ValueType_TInt"),
			GoStmt.GoRaw("case float32, float64:"),
			GoStmt.GoRaw("\treturn ValueType_TFloat"),
			GoStmt.GoRaw("case string, *string:"),
			GoStmt.GoRaw("\treturn ValueType_TClass(&hxrt__TypeClassValue{name: hxrt.StringFromLiteral(\"String\")})"),
			GoStmt.GoRaw("case *hxrt.Array:"),
			GoStmt.GoRaw("\treturn ValueType_TClass(&hxrt__TypeClassValue{name: hxrt.StringFromLiteral(\"Array\")})"),
			GoStmt.GoRaw("}"),
			GoStmt.GoRaw("ref := reflect.ValueOf(v)"),
			GoStmt.GoRaw("if !ref.IsValid() {"),
			GoStmt.GoRaw("\treturn ValueType_TNull"),
			GoStmt.GoRaw("}"),
			GoStmt.GoRaw("switch ref.Kind() {"),
			GoStmt.GoRaw("case reflect.Func:"),
			GoStmt.GoRaw("\treturn ValueType_TFunction"),
			GoStmt.GoRaw("case reflect.Slice, reflect.Array:"),
			GoStmt.GoRaw("\treturn ValueType_TClass(&hxrt__TypeClassValue{name: hxrt.StringFromLiteral(\"Array\")})"),
			GoStmt.GoRaw("case reflect.Map, reflect.Struct, reflect.Interface, reflect.Pointer:"),
			GoStmt.GoRaw("\treturn ValueType_TObject"),
			GoStmt.GoRaw("default:"),
			GoStmt.GoRaw("\treturn ValueType_TUnknown"),
			GoStmt.GoRaw("}")
		];

		return [
			GoDecl.GoStructDecl("ValueType", [{name: "tag", typeName: "int"}, {name: "params", typeName: "[]any"}]),
			GoDecl.GoGlobalVarDecl("ValueType_TNull", "*ValueType", GoExpr.GoRaw("&ValueType{tag: 0, params: []any{}}")),
			GoDecl.GoGlobalVarDecl("ValueType_TInt", "*ValueType", GoExpr.GoRaw("&ValueType{tag: 1, params: []any{}}")),
			GoDecl.GoGlobalVarDecl("ValueType_TFloat", "*ValueType", GoExpr.GoRaw("&ValueType{tag: 2, params: []any{}}")),
			GoDecl.GoGlobalVarDecl("ValueType_TBool", "*ValueType", GoExpr.GoRaw("&ValueType{tag: 3, params: []any{}}")),
			GoDecl.GoGlobalVarDecl("ValueType_TObject", "*ValueType", GoExpr.GoRaw("&ValueType{tag: 4, params: []any{}}")),
			GoDecl.GoGlobalVarDecl("ValueType_TFunction", "*ValueType", GoExpr.GoRaw("&ValueType{tag: 5, params: []any{}}")),
			GoDecl.GoGlobalVarDecl("ValueType_TUnknown", "*ValueType", GoExpr.GoRaw("&ValueType{tag: 8, params: []any{}}")),
			GoDecl.GoFuncDecl("ValueType_TClass", null, [
				{
					name: "c",
					typeName: "any"
				}
			],
				["*ValueType"], [GoStmt.GoReturn(GoExpr.GoRaw("&ValueType{tag: 6, params: []any{c}}"))]),
			GoDecl.GoFuncDecl("ValueType_TEnum", null, [{name: "e", typeName: "any"}], ["*ValueType"],
				[GoStmt.GoReturn(GoExpr.GoRaw("&ValueType{tag: 7, params: []any{e}}"))]),
			GoDecl.GoFuncDecl("hxrt_typeCallAny", null, [{name: "callable", typeName: "any"}, {name: "args", typeName: "[]any"}], ["any", "bool"], [
				GoStmt.GoRaw("result := any(nil)"),
				GoStmt.GoRaw("ok := false"),
				GoStmt.GoRaw("defer func() {"),
				GoStmt.GoRaw("\tif recover() != nil {"),
				GoStmt.GoRaw("\t\tresult = nil"),
				GoStmt.GoRaw("\t\tok = false"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("}()"),
				GoStmt.GoRaw("if callable == nil {"),
				GoStmt.GoRaw("\treturn nil, false"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("fn := reflect.ValueOf(callable)"),
				GoStmt.GoRaw("if !fn.IsValid() || fn.Kind() != reflect.Func {"),
				GoStmt.GoRaw("\treturn nil, false"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("fnType := fn.Type()"),
				GoStmt.GoRaw("if fnType.NumIn() != len(args) {"),
				GoStmt.GoRaw("\treturn nil, false"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("in := make([]reflect.Value, len(args))"),
				GoStmt.GoRaw("for i := 0; i < len(args); i++ {"),
				GoStmt.GoRaw("\tparamType := fnType.In(i)"),
				GoStmt.GoRaw("\targ := args[i]"),
				GoStmt.GoRaw("\tif arg == nil {"),
				GoStmt.GoRaw("\t\tin[i] = reflect.Zero(paramType)"),
				GoStmt.GoRaw("\t\tcontinue"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tv := reflect.ValueOf(arg)"),
				GoStmt.GoRaw("\tif v.IsValid() && v.Type().AssignableTo(paramType) {"),
				GoStmt.GoRaw("\t\tin[i] = v"),
				GoStmt.GoRaw("\t\tcontinue"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tif v.IsValid() && v.Type().ConvertibleTo(paramType) {"),
				GoStmt.GoRaw("\t\tin[i] = v.Convert(paramType)"),
				GoStmt.GoRaw("\t\tcontinue"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tif paramType.Kind() == reflect.Interface && v.IsValid() {"),
				GoStmt.GoRaw("\t\tin[i] = v"),
				GoStmt.GoRaw("\t\tcontinue"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\treturn nil, false"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("out := fn.Call(in)"),
				GoStmt.GoRaw("if len(out) == 0 {"),
				GoStmt.GoRaw("\treturn nil, true"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("first := out[0]"),
				GoStmt.GoRaw("if !first.IsValid() {"),
				GoStmt.GoRaw("\treturn nil, true"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("result = first.Interface()"),
				GoStmt.GoRaw("ok = true"),
				GoStmt.GoRaw("return result, ok")
			]),
			GoDecl.GoFuncDecl("hxrt_typeArrayValues", null, [
				{
					name: "value",
					typeName: "*hxrt.Array"
				}
			], ["[]any"], [
				GoStmt.GoIf(GoExpr.GoBinary("==", GoExpr.GoIdent("value"), GoExpr.GoNil), [GoStmt.GoReturn(GoExpr.GoArrayLiteral("any", []))], null),
				GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoIdent("value"), "Values"), []))
			]),
			GoDecl.GoFuncDecl("hxrt_typeResolvedClassName", null, [
				{
					name: "value",
					typeName: "any"
				}
			], ["string", "bool"], [
				GoStmt.GoRaw("switch current := value.(type) {"),
				GoStmt.GoRaw("case *hxrt__TypeClassValue:"),
				GoStmt.GoRaw("\tif current == nil || current.name == nil {"),
				GoStmt.GoRaw("\t\treturn \"\", false"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\treturn *current.name, true"),
				GoStmt.GoRaw("case hxrt__TypeClassValue:"),
				GoStmt.GoRaw("\tif current.name == nil {"),
				GoStmt.GoRaw("\t\treturn \"\", false"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\treturn *current.name, true"),
				GoStmt.GoRaw("case string:"),
				GoStmt.GoRaw("\treturn current, true"),
				GoStmt.GoRaw("case *string:"),
				GoStmt.GoRaw("\tif current == nil {"),
				GoStmt.GoRaw("\t\treturn \"\", false"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\treturn *current, true"),
				GoStmt.GoRaw("default:"),
				GoStmt.GoRaw("\treturn \"\", false"),
				GoStmt.GoRaw("}")
			]),
			GoDecl.GoFuncDecl("hxrt_typeResolvedEnumName", null, [
				{
					name: "value",
					typeName: "any"
				}
			], ["string", "bool"], [
				GoStmt.GoRaw("switch current := value.(type) {"),
				GoStmt.GoRaw("case *hxrt__TypeEnumValue:"),
				GoStmt.GoRaw("\tif current == nil || current.name == nil {"),
				GoStmt.GoRaw("\t\treturn \"\", false"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\treturn *current.name, true"),
				GoStmt.GoRaw("case hxrt__TypeEnumValue:"),
				GoStmt.GoRaw("\tif current.name == nil {"),
				GoStmt.GoRaw("\t\treturn \"\", false"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\treturn *current.name, true"),
				GoStmt.GoRaw("case string:"),
				GoStmt.GoRaw("\treturn current, true"),
				GoStmt.GoRaw("case *string:"),
				GoStmt.GoRaw("\tif current == nil {"),
				GoStmt.GoRaw("\t\treturn \"\", false"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\treturn *current, true"),
				GoStmt.GoRaw("default:"),
				GoStmt.GoRaw("\treturn \"\", false"),
				GoStmt.GoRaw("}")
			]),
			GoDecl.GoFuncDecl("hxrt_typeCreateClassInstance", null, [
				{
					name: "className",
					typeName: "string"
				},
				{name: "args", typeName: "[]any"}
			],
				["any", "bool"], classCreateBody),
			GoDecl.GoFuncDecl("hxrt_typeCreateClassEmptyInstance", null, [{name: "className", typeName: "string"}], ["any", "bool"], classCreateEmptyBody),
			GoDecl.GoFuncDecl("hxrt_typeCreateEnumInstance", null, [
				{name: "enumName", typeName: "string"},
				{name: "constructorName", typeName: "string"},
				{name: "constructorIndex", typeName: "int"},
				{name: "useIndex", typeName: "bool"},
				{name: "args", typeName: "[]any"}
			],
				["any", "bool"], enumCreateBody),
			GoDecl.GoFuncDecl("Type_getClass", null, [{name: "o", typeName: "any"}], ["any"], getClassBody),
			GoDecl.GoFuncDecl("Type_getEnum", null, [{name: "o", typeName: "any"}], ["any"], getEnumBody),
			GoDecl.GoFuncDecl("Type_getSuperClass", null, [{name: "c", typeName: "any"}], ["any"], getSuperClassBody),
			GoDecl.GoFuncDecl("Type_getClassName", null, [{name: "c", typeName: "any"}], ["*string"], [
				GoStmt.GoRaw("className, ok := hxrt_typeResolvedClassName(c)"),
				GoStmt.GoRaw("if !ok {"),
				GoStmt.GoRaw("\treturn nil"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("return hxrt.StringFromLiteral(className)")
			]),
			GoDecl.GoFuncDecl("Type_getClassFields", null, [
				{
					name: "c",
					typeName: "any"
				}
			],
				["*hxrt.Array"], getClassFieldsBody),
			GoDecl.GoFuncDecl("Type_getInstanceFields", null, [{name: "c", typeName: "any"}], ["*hxrt.Array"], getInstanceFieldsBody),
			GoDecl.GoFuncDecl("Type_getEnumName", null, [{name: "e", typeName: "any"}], ["*string"], [
				GoStmt.GoRaw("enumName, ok := hxrt_typeResolvedEnumName(e)"),
				GoStmt.GoRaw("if !ok {"),
				GoStmt.GoRaw("\treturn nil"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("return hxrt.StringFromLiteral(enumName)")
			]),
			GoDecl.GoFuncDecl("Type_resolveClass", null, [
				{
					name: "name",
					typeName: "*string"
				}
			],
				["any"], classResolveBody),
			GoDecl.GoFuncDecl("Type_resolveEnum", null, [{name: "name", typeName: "*string"}], ["any"], enumResolveBody),
			GoDecl.GoFuncDecl("Type_createInstance", null, [{name: "cl", typeName: "any"}, {name: "args", typeName: "*hxrt.Array"}], ["any"], [
				GoStmt.GoRaw("className, ok := hxrt_typeResolvedClassName(cl)"),
				GoStmt.GoRaw("if !ok {"),
				GoStmt.GoRaw("\treturn nil"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("instance, ok := hxrt_typeCreateClassInstance(className, hxrt_typeArrayValues(args))"),
				GoStmt.GoRaw("if !ok {"),
				GoStmt.GoRaw("\treturn nil"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("return instance")
			]),
			GoDecl.GoFuncDecl("Type_createEmptyInstance", null, [
				{
					name: "cl",
					typeName: "any"
				}
			], ["any"], [
				GoStmt.GoRaw("className, ok := hxrt_typeResolvedClassName(cl)"),
				GoStmt.GoRaw("if !ok {"),
				GoStmt.GoRaw("\treturn nil"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("instance, ok := hxrt_typeCreateClassEmptyInstance(className)"),
				GoStmt.GoRaw("if !ok {"),
				GoStmt.GoRaw("\treturn nil"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("return instance")
			]),
			GoDecl.GoFuncDecl("Type_createEnum", null, [
				{
					name: "e",
					typeName: "any"
				},
				{name: "constr", typeName: "*string"},
				{name: "params", typeName: "*hxrt.Array"}
			], ["any"], [
				GoStmt.GoRaw("enumName, ok := hxrt_typeResolvedEnumName(e)"),
				GoStmt.GoRaw("if !ok {"),
				GoStmt.GoRaw("\treturn nil"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("constructorName := \"\""),
				GoStmt.GoRaw("if constr != nil {"),
				GoStmt.GoRaw("\tconstructorName = *hxrt.StdString(constr)"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("enumValue, ok := hxrt_typeCreateEnumInstance(enumName, constructorName, 0, false, hxrt_typeArrayValues(params))"),
				GoStmt.GoRaw("if !ok {"),
				GoStmt.GoRaw("\treturn nil"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("return enumValue")
			]),
			GoDecl.GoFuncDecl("Type_createEnumIndex", null, [
				{
					name: "e",
					typeName: "any"
				},
				{name: "index", typeName: "int"},
				{name: "params", typeName: "*hxrt.Array"}
			], ["any"], [
				GoStmt.GoRaw("enumName, ok := hxrt_typeResolvedEnumName(e)"),
				GoStmt.GoRaw("if !ok {"),
				GoStmt.GoRaw("\treturn nil"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("enumValue, ok := hxrt_typeCreateEnumInstance(enumName, \"\", index, true, hxrt_typeArrayValues(params))"),
				GoStmt.GoRaw("if !ok {"),
				GoStmt.GoRaw("\treturn nil"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("return enumValue")
			]),
			GoDecl.GoFuncDecl("Type_enumConstructor", null, [
				{
					name: "e",
					typeName: "any"
				}
			],
				["*string"], enumConstructorBody),
			GoDecl.GoFuncDecl("Type_enumIndex", null, [{name: "e", typeName: "any"}], ["int"], enumIndexBody),
			GoDecl.GoFuncDecl("Type_getEnumConstructs", null, [{name: "e", typeName: "any"}], ["*hxrt.Array"], getEnumConstructsBody),
			GoDecl.GoFuncDecl("Type_enumParameters", null, [{name: "e", typeName: "any"}], ["*hxrt.Array"], enumParametersBody),
			GoDecl.GoFuncDecl("Type_allEnums", null, [{name: "e", typeName: "any"}], ["*hxrt.Array"], allEnumsBody),
			GoDecl.GoFuncDecl("Type_typeof", null, [{name: "v", typeName: "any"}], ["any"], typeOfBody),
			GoDecl.GoFuncDecl("Type_enumEq", null, [{name: "a", typeName: "any"}, {name: "b", typeName: "any"}], ["bool"],
				[GoStmt.GoReturn(GoExpr.GoRaw("reflect.DeepEqual(a, b)"))])
		];
	}
}
#end
