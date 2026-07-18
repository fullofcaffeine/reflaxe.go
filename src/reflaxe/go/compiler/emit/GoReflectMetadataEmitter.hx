package reflaxe.go.compiler.emit;

#if macro
import reflaxe.go.ast.GoAST.GoDecl;
import reflaxe.go.ast.GoAST.GoExpr;
import reflaxe.go.ast.GoAST.GoStmt;
import reflaxe.go.ast.GoAST.GoTypeSwitchCase;
import reflaxe.go.ast.GoBinaryOperator;
import reflaxe.go.ast.GoBuiltinType;
import reflaxe.go.ast.GoType;
import reflaxe.go.ast.GoUnaryOperator;

/**
	What:
	- Emits the exact closed-world metadata adapters used by staged `Reflect`.

	Why:
	- Class-token RTTI, generated lowercase methods, and enum carriers are known
	  only after the compiler has finalized the reachable program. They must remain
	  separate from ordinary runtime map, struct, and function inspection.

	How:
	- Wrap the narrow RTTI tuple helper with source-callable typed functions, bridge
	  the selective generated-method resolver, and classify enums with an exact
	  typed switch over emitted enum carriers.
**/
class GoReflectMetadataEmitter {
	public static inline final TYPE_FIELD_SYMBOL = "reflaxe__go___internal__CompilerReflect_typeField";
	public static inline final HAS_TYPE_FIELD_SYMBOL = "reflaxe__go___internal__CompilerReflect_hasTypeField";
	public static inline final GENERATED_METHOD_SYMBOL = "reflaxe__go___internal__CompilerReflect_generatedMethod";
	public static inline final IS_ENUM_VALUE_SYMBOL = "reflaxe__go___internal__CompilerReflect_isEnumValue";

	public static function emit(requiresTypeField:Bool, requiresGeneratedMethod:Bool, hasGeneratedMethodPlan:Bool, requiresEnumValue:Bool,
			enumGoTypeNames:Array<String>):Array<GoDecl> {
		var declarations = new Array<GoDecl>();
		if (requiresTypeField) {
			declarations.push(typeFieldDecl());
			declarations.push(hasTypeFieldDecl());
		}
		if (requiresGeneratedMethod) {
			declarations.push(generatedMethodDecl(hasGeneratedMethodPlan));
		}
		if (requiresEnumValue) {
			declarations.push(isEnumValueDecl(enumGoTypeNames));
		}
		return declarations;
	}

	static function typeFieldDecl():GoDecl {
		return GoDecl.GoFuncDecl(TYPE_FIELD_SYMBOL, null, reflectFieldParams(), [GoType.builtin(GoBuiltinType.AnyType)], [
			fieldKeyDecl(),
			GoStmt.GoMultiAssign(["value", "found"], rttiLookupCall(), true),
			GoStmt.GoIf(GoExpr.GoUnary(GoUnaryOperator.LogicalNot, GoExpr.GoIdent("found")), [GoStmt.GoReturn(GoExpr.GoNil)], null),
			GoStmt.GoReturn(GoExpr.GoIdent("value"))
		]);
	}

	static function hasTypeFieldDecl():GoDecl {
		return GoDecl.GoFuncDecl(HAS_TYPE_FIELD_SYMBOL, null, reflectFieldParams(), [GoType.builtin(GoBuiltinType.Bool)], [
			fieldKeyDecl(),
			GoStmt.GoMultiAssign(["_", "found"], rttiLookupCall(), true),
			GoStmt.GoReturn(GoExpr.GoIdent("found"))
		]);
	}

	static function generatedMethodDecl(hasPlan:Bool):GoDecl {
		var body = if (hasPlan) {
			[
				fieldKeyDecl(),
				GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent(GoGeneratedMethodMetadataEmitter.LOOKUP_SYMBOL),
					[GoExpr.GoIdent("object"), GoExpr.GoIdent("key")]))
			];
		} else {
			[GoStmt.GoReturn(GoExpr.GoNil)];
		};
		return GoDecl.GoFuncDecl(GENERATED_METHOD_SYMBOL, null, reflectFieldParams(), [GoType.builtin(GoBuiltinType.AnyType)], body);
	}

	static function isEnumValueDecl(enumGoTypeNames:Array<String>):GoDecl {
		if (enumGoTypeNames.length == 0) {
			return GoDecl.GoFuncDecl(IS_ENUM_VALUE_SYMBOL, null, [{name: "value", typeName: GoType.builtin(GoBuiltinType.AnyType)}],
				[GoType.builtin(GoBuiltinType.Bool)], [GoStmt.GoReturn(GoExpr.GoBoolLiteral(false))]);
		}
		var cases = new Array<GoTypeSwitchCase>();
		for (goTypeName in enumGoTypeNames) {
			cases.push({
				typeName: GoType.pointer(GoType.named(goTypeName)),
				body: [
					GoStmt.GoReturn(GoExpr.GoBinary(GoBinaryOperator.NotEqual, GoExpr.GoIdent("enumValue"), GoExpr.GoNil))
				]
			});
		}
		return GoDecl.GoFuncDecl(IS_ENUM_VALUE_SYMBOL, null, [{name: "value", typeName: GoType.builtin(GoBuiltinType.AnyType)}],
			[GoType.builtin(GoBuiltinType.Bool)], [
				GoStmt.GoTypeSwitch(GoExpr.GoIdent("value"), "enumValue", cases, [GoStmt.GoReturn(GoExpr.GoBoolLiteral(false))])
			]);
	}

	static function reflectFieldParams() {
		return [
			{name: "object", typeName: GoType.builtin(GoBuiltinType.AnyType)},
			{name: "field", typeName: GoType.pointer(GoType.builtin(GoBuiltinType.StringType))}
		];
	}

	static function fieldKeyDecl():GoStmt {
		return GoStmt.GoVarDecl("key", null,
			GoExpr.GoUnary(GoUnaryOperator.Dereference, GoExpr.GoCall(GoExpr.GoIdent("hxrt.StdString"), [GoExpr.GoIdent("field")])), true);
	}

	static function rttiLookupCall():GoExpr {
		return GoExpr.GoCall(GoExpr.GoIdent(GoRttiMetadataEmitter.LOOKUP_SYMBOL), [GoExpr.GoIdent("object"), GoExpr.GoIdent("key")]);
	}
}
#end
