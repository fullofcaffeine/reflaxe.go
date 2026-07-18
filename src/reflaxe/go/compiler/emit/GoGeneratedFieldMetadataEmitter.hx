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
import reflaxe.go.ast.GoUnaryOperator;

private typedef GeneratedFieldEntry = {
	final goTypeName:String;
	final parentGoTypeName:Null<String>;
	final participatesInLookup:Bool;
	final allFields:Array<String>;
	final ownFields:Array<{
		final lookupKey:String;
		final selector:String;
		final typeName:GoType;
	}>;
}

private enum GeneratedFieldOperation {
	Lookup;
	Has;
	Set;
}

/**
	What:
	- Emits same-package adapters for reading, testing, setting, and enumerating
	  generated lowercase Haxe instance fields through staged `Reflect`.

	Why:
	- Go reflection outside the generated program package cannot interface with or
	  set lowercase fields. A runtime registry or unsafe access would duplicate
	  compiler authority and weaken the generated package boundary.

	How:
	- Recover each carrier's canonical `__hx_this` receiver with typed switches,
	  select exact fields by their Haxe lookup keys, and use typed assignments for
	  writes. Each per-class resolver checks only own fields and performs one
	  generated-superclass fallback. Enumeration classifies every generated
	  carrier through the same canonical receiver and returns its inherited,
	  source-visible stored-field names without consulting cross-package reflection.
**/
class GoGeneratedFieldMetadataEmitter {
	public static inline final LOOKUP_SYMBOL = "reflaxe__go___internal__CompilerReflect_generatedField";
	public static inline final HAS_SYMBOL = "reflaxe__go___internal__CompilerReflect_hasGeneratedField";
	public static inline final SET_SYMBOL = "reflaxe__go___internal__CompilerReflect_setGeneratedField";
	public static inline final FIELDS_SYMBOL = "reflaxe__go___internal__CompilerReflect_generatedFields";

	public static function emit(entries:Array<GeneratedFieldEntry>):Array<GoDecl> {
		var lookupEntries = [for (entry in entries) if (entry.participatesInLookup) entry];
		var declarations = [
			centralResolverDecl(lookupEntries, GeneratedFieldOperation.Lookup),
			centralResolverDecl(lookupEntries, GeneratedFieldOperation.Has),
			centralResolverDecl(lookupEntries, GeneratedFieldOperation.Set),
			generatedFieldsDecl(entries)
		];
		for (entry in lookupEntries) {
			declarations.push(classResolverDecl(entry, GeneratedFieldOperation.Lookup));
			declarations.push(classResolverDecl(entry, GeneratedFieldOperation.Has));
			declarations.push(classResolverDecl(entry, GeneratedFieldOperation.Set));
		}
		return declarations;
	}

	/**
		What: Emit nullable generated-carrier classification plus exact field names.

		Why: Staged `Reflect.fields` must distinguish a generated class with zero fields
		from a native object and must not expose embedded Go carrier fields.

		How: Recover the canonical `__hx_this` receiver, type-switch across every
		generated class, and return its precomputed portable field names; `nil` alone
		means the object is not a generated carrier and permits native fallback.
	**/
	static function generatedFieldsDecl(entries:Array<GeneratedFieldEntry>):GoDecl {
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
			body.push(GoStmt.GoTypeSwitch(GoExpr.GoIdent("object"), "value", carrierCases, [GoStmt.GoReturn(GoExpr.GoNil)]));

			var receiverCases = new Array<GoTypeSwitchCase>();
			for (entry in entries) {
				receiverCases.push({
					typeName: GoType.pointer(GoType.named(entry.goTypeName)),
					body: [GoStmt.GoReturn(generatedFieldArray(entry.allFields))]
				});
			}
			body.push(GoStmt.GoTypeSwitch(GoExpr.GoIdent("receiver"), null, receiverCases, [GoStmt.GoReturn(GoExpr.GoNil)]));
		}
		return GoDecl.GoFuncDecl(FIELDS_SYMBOL, null, [{name: "object", typeName: GoType.builtin(GoBuiltinType.AnyType)}], ["*hxrt.Array"], body);
	}

	static function generatedFieldArray(fields:Array<String>):GoExpr {
		return GoExpr.GoCall(GoExpr.GoIdent("hxrt.NewArray"), [
			for (field in fields)
				GoExpr.GoCall(GoExpr.GoIdent("hxrt.StringFromLiteral"), [GoExpr.GoStringLiteral(field)])
		]);
	}

	static function centralResolverDecl(entries:Array<GeneratedFieldEntry>, operation:GeneratedFieldOperation):GoDecl {
		var body = new Array<GoStmt>();
		if (entries.length == 0) {
			body.push(defaultReturn(operation));
		} else {
			body.push(GoStmt.GoVarDecl("key", null,
				GoExpr.GoUnary(GoUnaryOperator.Dereference, GoExpr.GoCall(GoExpr.GoIdent("hxrt.StdString"), [GoExpr.GoIdent("field")])), true));
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
							[defaultReturn(operation)], null),
						GoStmt.GoAssign(GoExpr.GoIdent("receiver"), canonical)
					]
				});
			}
			body.push(GoStmt.GoTypeSwitch(GoExpr.GoIdent("object"), "value", carrierCases, [defaultReturn(operation)]));

			var receiverCases = new Array<GoTypeSwitchCase>();
			for (entry in entries) {
				var args = [GoExpr.GoIdent("value"), GoExpr.GoIdent("key")];
				if (operation == GeneratedFieldOperation.Set) {
					args.push(GoExpr.GoIdent("incoming"));
				}
				receiverCases.push({
					typeName: GoType.pointer(GoType.named(entry.goTypeName)),
					body: [
						GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent(classResolverSymbol(entry.goTypeName, operation)), args))
					]
				});
			}
			body.push(GoStmt.GoTypeSwitch(GoExpr.GoIdent("receiver"), "value", receiverCases, [defaultReturn(operation)]));
		}

		var params = [
			{name: "object", typeName: GoType.builtin(GoBuiltinType.AnyType)},
			{name: "field", typeName: GoType.pointer(GoType.builtin(GoBuiltinType.StringType))}
		];
		if (operation == GeneratedFieldOperation.Set) {
			params.push({name: "incoming", typeName: GoType.builtin(GoBuiltinType.AnyType)});
		}
		return GoDecl.GoFuncDecl(centralSymbol(operation), null, params, [resultType(operation)], body);
	}

	static function classResolverDecl(entry:GeneratedFieldEntry, operation:GeneratedFieldOperation):GoDecl {
		var value = GoExpr.GoIdent("value");
		var body = [
			GoStmt.GoIf(GoExpr.GoBinary(GoBinaryOperator.Equal, value, GoExpr.GoNil), [defaultReturn(operation)], null)
		];
		if (entry.ownFields.length > 0) {
			var fieldCases = new Array<GoSwitchCase>();
			for (field in entry.ownFields) {
				fieldCases.push({
					values: [GoExpr.GoStringLiteral(field.lookupKey)],
					body: fieldCaseBody(value, field.selector, field.typeName, operation)
				});
			}
			body.push(GoStmt.GoSwitch(GoExpr.GoIdent("key"), fieldCases, null));
		}

		if (entry.parentGoTypeName != null) {
			var parent = GoExpr.GoSelector(value, entry.parentGoTypeName);
			body.push(GoStmt.GoIf(GoExpr.GoBinary(GoBinaryOperator.Equal, parent, GoExpr.GoNil), [defaultReturn(operation)], null));
			var args = [parent, GoExpr.GoIdent("key")];
			if (operation == GeneratedFieldOperation.Set) {
				args.push(GoExpr.GoIdent("incoming"));
			}
			body.push(GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent(classResolverSymbol(entry.parentGoTypeName, operation)), args)));
		} else {
			body.push(defaultReturn(operation));
		}

		var params = [
			{name: "value", typeName: GoType.pointer(GoType.named(entry.goTypeName))},
			{name: "key", typeName: GoType.builtin(GoBuiltinType.StringType)}
		];
		if (operation == GeneratedFieldOperation.Set) {
			params.push({name: "incoming", typeName: GoType.builtin(GoBuiltinType.AnyType)});
		}
		return GoDecl.GoFuncDecl(classResolverSymbol(entry.goTypeName, operation), null, params, [resultType(operation)], body);
	}

	static function fieldCaseBody(value:GoExpr, selector:String, typeName:GoType, operation:GeneratedFieldOperation):Array<GoStmt> {
		return switch (operation) {
			case Lookup:
				[GoStmt.GoReturn(GoExpr.GoSelector(value, selector))];
			case Has:
				[GoStmt.GoReturn(GoExpr.GoBoolLiteral(true))];
			case Set:
				var target = GoExpr.GoSelector(value, selector);
				var incoming = GoExpr.GoIdent("incoming");
				[
					GoStmt.GoIf(GoExpr.GoBinary(GoBinaryOperator.Equal, incoming, GoExpr.GoNil), [
						GoStmt.GoVarDecl("zero", typeName, null, false),
						GoStmt.GoAssign(target, GoExpr.GoIdent("zero")),
						GoStmt.GoReturn(GoExpr.GoBoolLiteral(true))
					], null),
					GoStmt.GoTypeSwitch(incoming, "typed", [
						{
							typeName: typeName,
							body: [
								GoStmt.GoAssign(target, GoExpr.GoIdent("typed")),
								GoStmt.GoReturn(GoExpr.GoBoolLiteral(true))
							]
						}
					], [GoStmt.GoReturn(GoExpr.GoBoolLiteral(false))])
				];
		};
	}

	static function centralSymbol(operation:GeneratedFieldOperation):String {
		return switch (operation) {
			case Lookup: LOOKUP_SYMBOL;
			case Has: HAS_SYMBOL;
			case Set: SET_SYMBOL;
		};
	}

	static function classResolverSymbol(goTypeName:String, operation:GeneratedFieldOperation):String {
		var suffix = switch (operation) {
			case Lookup: "lookup";
			case Has: "has";
			case Set: "set";
		};
		return "hxrt__generated_field_" + suffix + "__" + goTypeName;
	}

	static function resultType(operation:GeneratedFieldOperation):GoType {
		return operation == GeneratedFieldOperation.Lookup ? GoType.builtin(GoBuiltinType.AnyType) : GoType.builtin(GoBuiltinType.Bool);
	}

	static function defaultReturn(operation:GeneratedFieldOperation):GoStmt {
		return operation == GeneratedFieldOperation.Lookup ? GoStmt.GoReturn(GoExpr.GoNil) : GoStmt.GoReturn(GoExpr.GoBoolLiteral(false));
	}
}
#end
