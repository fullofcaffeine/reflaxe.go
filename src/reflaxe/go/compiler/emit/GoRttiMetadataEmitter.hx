package reflaxe.go.compiler.emit;

#if macro
import reflaxe.go.ast.GoAST.GoDecl;
import reflaxe.go.ast.GoAST.GoExpr;
import reflaxe.go.ast.GoAST.GoStmt;

/**
	What:
	Builds the narrow compiler-owned metadata lookup bridge behind class-token
	`__rtti` / `__meta__` access on Go.

	Why:
	Direct `haxe.rtti` support needs a stable backend contract for metadata fields
	on class tokens, but that contract depends on generated static symbols and the
	backend's class-token representation. That makes it compiler-owned, while still
	keeping the public `haxe.rtti.*` APIs in staged std.

	How:
	Consumes precomputed class metadata symbol bindings from `GoCompiler` and emits
	a small lookup helper. `Reflect.field` / `Reflect.hasField` call this helper
	before falling back to generic map/struct reflection.
**/
class GoRttiMetadataEmitter {
	public static inline final LOOKUP_SYMBOL = "hxrt_typeClassMetadataField";

	public static function emit(classMetadata:Array<{
		final haxeTypeName:String;
		final rttiSymbol:Null<String>;
		final metaSymbol:Null<String>;
	}>, goRawQuotedString:String->String):Array<GoDecl> {
		var body = [
			GoStmt.GoMultiAssign(["classValue", "ok"], GoExpr.GoTypeAssert(GoExpr.GoIdent("value"), "*hxrt__TypeClassValue"), true),
			GoStmt.GoRaw("if !ok || classValue == nil {"),
			GoStmt.GoRaw("\treturn nil, false"),
			GoStmt.GoRaw("}"),
			GoStmt.GoRaw("className := *hxrt.StdString(classValue.name)"),
			GoStmt.GoRaw("switch className {")
		];

		for (entry in classMetadata) {
			body.push(GoStmt.GoRaw("case " + goRawQuotedString(entry.haxeTypeName) + ":"));
			body.push(GoStmt.GoRaw("\tswitch key {"));
			if (entry.rttiSymbol != null) {
				body.push(GoStmt.GoRaw("\tcase \"__rtti\":"));
				body.push(GoStmt.GoRaw("\t\treturn " + entry.rttiSymbol + ", true"));
			}
			if (entry.metaSymbol != null) {
				body.push(GoStmt.GoRaw("\tcase \"__meta__\":"));
				body.push(GoStmt.GoRaw("\t\treturn " + entry.metaSymbol + ", true"));
			}
			body.push(GoStmt.GoRaw("\tdefault:"));
			body.push(GoStmt.GoRaw("\t\treturn nil, false"));
			body.push(GoStmt.GoRaw("\t}"));
		}

		body.push(GoStmt.GoRaw("default:"));
		body.push(GoStmt.GoRaw("\treturn nil, false"));
		body.push(GoStmt.GoRaw("}"));

		return [
			GoDecl.GoFuncDecl(LOOKUP_SYMBOL, null, [{name: "value", typeName: "any"}, {name: "key", typeName: "string"}], ["any", "bool"], body)
		];
	}
}
#end
