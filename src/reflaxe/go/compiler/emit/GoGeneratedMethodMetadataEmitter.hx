package reflaxe.go.compiler.emit;

#if macro
import reflaxe.go.ast.GoAST.GoDecl;
import reflaxe.go.ast.GoAST.GoExpr;
import reflaxe.go.ast.GoAST.GoStmt;
import reflaxe.go.ast.GoAST.GoSwitchCase;
import reflaxe.go.ast.GoAST.GoTypeSwitchCase;
import reflaxe.go.ast.GoBinaryOperator;
import reflaxe.go.ast.GoBuiltinType;
import reflaxe.go.ast.GoType;

/**
	What:
	Emits the same-package method metadata used to find already-generated lowercase
	Haxe methods by their source field names.

	Why:
	Go reflection exposes exported methods only, and the separate `hxrt` package
	cannot select unexported methods from the generated program package. Portable
	APIs such as `Reflect.field` and staged `haxe.Template` still need those methods
	as ordinary bound function values.

	How:
	First recover each physical carrier's canonical `__hx_this` receiver. Then use
	one exact concrete-type switch and a per-class resolver containing only that
	class's own methods, with one typed embedded-superclass fallback. The
	`hxrt__generated_method_field` prefix is a compiler-reserved namespace; every
	declaration is deterministic and built from typed Go AST nodes.
**/
class GoGeneratedMethodMetadataEmitter {
	public static inline final LOOKUP_SYMBOL = "hxrt__generated_method_field";

	public static function emit(entries:Array<{
		final goTypeName:String;
		final resolverSymbol:String;
		final parentGoTypeName:Null<String>;
		final parentResolverSymbol:Null<String>;
		final ownMethods:Array<{
			final lookupKey:String;
			final selector:String;
		}>;
	}>):Array<GoDecl> {
		var declarations = [centralLookupDecl(entries)];
		for (entry in entries) {
			declarations.push(classResolverDecl(entry));
		}
		return declarations;
	}

	static function centralLookupDecl(entries:Array<{
		final goTypeName:String;
		final resolverSymbol:String;
		final parentGoTypeName:Null<String>;
		final parentResolverSymbol:Null<String>;
		final ownMethods:Array<{
			final lookupKey:String;
			final selector:String;
		}>;
	}>):GoDecl {
		var body = new Array<GoStmt>();
		if (entries.length == 0) {
			body.push(GoStmt.GoReturn(GoExpr.GoNil));
		} else {
			body.push(GoStmt.GoVarDecl("receiver", GoType.builtin(GoBuiltinType.AnyType), null, false));
			var carrierCases = new Array<GoTypeSwitchCase>();
			for (entry in entries) {
				var value = GoExpr.GoIdent("value");
				var canonical = GoExpr.GoSelector(value, "__hx_this");
				carrierCases.push({
					typeName: GoType.pointer(GoType.named(entry.goTypeName)),
					body: [
						GoStmt.GoIf(GoExpr.GoBinary(GoBinaryOperator.LogicalOr, GoExpr.GoBinary(GoBinaryOperator.Equal, value, GoExpr.GoNil),
							GoExpr.GoBinary(GoBinaryOperator.Equal, canonical, GoExpr.GoNil)),
							[GoStmt.GoReturn(GoExpr.GoNil)], null),
						GoStmt.GoAssign(GoExpr.GoIdent("receiver"), canonical)
					]
				});
			}
			body.push(GoStmt.GoTypeSwitch(GoExpr.GoIdent("obj"), "value", carrierCases, [GoStmt.GoReturn(GoExpr.GoNil)]));

			var receiverCases = new Array<GoTypeSwitchCase>();
			for (entry in entries) {
				receiverCases.push({
					typeName: GoType.pointer(GoType.named(entry.goTypeName)),
					body: [
						GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent(entry.resolverSymbol), [GoExpr.GoIdent("value"), GoExpr.GoIdent("key")]))
					]
				});
			}
			body.push(GoStmt.GoTypeSwitch(GoExpr.GoIdent("receiver"), "value", receiverCases, [GoStmt.GoReturn(GoExpr.GoNil)]));
		}

		return GoDecl.GoFuncDecl(LOOKUP_SYMBOL, null, [
			{name: "obj", typeName: GoType.builtin(GoBuiltinType.AnyType)},
			{name: "key", typeName: GoType.builtin(GoBuiltinType.StringType)}
		], [GoType.builtin(GoBuiltinType.AnyType)], body);
	}

	static function classResolverDecl(entry:{
		final goTypeName:String;
		final resolverSymbol:String;
		final parentGoTypeName:Null<String>;
		final parentResolverSymbol:Null<String>;
		final ownMethods:Array<{
			final lookupKey:String;
			final selector:String;
		}>;
	}):GoDecl {
		var value = GoExpr.GoIdent("value");
		var body = [
			GoStmt.GoIf(GoExpr.GoBinary(GoBinaryOperator.Equal, value, GoExpr.GoNil), [GoStmt.GoReturn(GoExpr.GoNil)], null)
		];
		if (entry.ownMethods.length > 0) {
			var methodCases = new Array<GoSwitchCase>();
			for (method in entry.ownMethods) {
				methodCases.push({
					values: [GoExpr.GoStringLiteral(method.lookupKey)],
					body: [GoStmt.GoReturn(GoExpr.GoSelector(value, method.selector))]
				});
			}
			body.push(GoStmt.GoSwitch(GoExpr.GoIdent("key"), methodCases, null));
		}

		var parentTypeName = entry.parentGoTypeName;
		var parentResolverSymbol = entry.parentResolverSymbol;
		if (parentTypeName != null && parentResolverSymbol != null) {
			var parent = GoExpr.GoSelector(value, parentTypeName);
			body.push(GoStmt.GoIf(GoExpr.GoBinary(GoBinaryOperator.Equal, parent, GoExpr.GoNil), [GoStmt.GoReturn(GoExpr.GoNil)], null));
			body.push(GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent(parentResolverSymbol), [parent, GoExpr.GoIdent("key")])));
		} else {
			body.push(GoStmt.GoReturn(GoExpr.GoNil));
		}

		return GoDecl.GoFuncDecl(entry.resolverSymbol, null, [
			{name: "value", typeName: GoType.pointer(GoType.named(entry.goTypeName))},
			{name: "key", typeName: GoType.builtin(GoBuiltinType.StringType)}
		], [GoType.builtin(GoBuiltinType.AnyType)], body);
	}
}
#end
