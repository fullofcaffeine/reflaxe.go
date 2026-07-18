package reflaxe.go.compiler.emit;

#if macro
import reflaxe.go.ast.GoAST.GoDecl;
import reflaxe.go.ast.GoAST.GoStmt;

/**
	What:
	Builds the narrow same-package invocation bridge used by staged Serializer and
	Unserializer source.

	Why:
	Generated public Haxe methods intentionally remain package-private Go methods,
	so the separately packaged `hxrt` runtime cannot invoke `hxSerialize`,
	`hxUnserialize`, or structural resolver methods. This is a generated
	representation fact, not token-stream or resolver policy.

	How:
	Emit only interface assertions and anonymous-object function extraction for the
	exact hooks requested by reachable staged serialization classes. No regex,
	token parsing, traversal, caches, type tables, construction, or reflection lives
	here.
**/
class GoSerializationSourceBridgeEmitter {
	public static function emit(includeSerializer:Bool, includeUnserializer:Bool):Array<GoDecl> {
		var declarations = new Array<GoDecl>();
		if (includeSerializer) {
			declarations.push(GoDecl.GoFuncDecl("haxe__GoSerializationBridge_hasSerializeHook", null, [{name: "value", typeName: "any"}], ["bool"], [
				GoStmt.GoRaw("if value == nil { return false }"),
				GoStmt.GoRaw("if _, ok := value.(interface{ hxSerialize(*haxe__Serializer) }); ok { return true }"),
				GoStmt.GoRaw("if _, ok := value.(interface{ HxSerialize(*haxe__Serializer) }); ok { return true }"),
				GoStmt.GoRaw("return false")
			]));
			declarations.push(GoDecl.GoFuncDecl("haxe__GoSerializationBridge_callSerializeHook", null, [
				{name: "value", typeName: "any"},
				{name: "serializer", typeName: "*haxe__Serializer"}
			], ["bool"], [
				GoStmt.GoRaw("if hook, ok := value.(interface{ hxSerialize(*haxe__Serializer) }); ok { hook.hxSerialize(serializer); return true }"),
				GoStmt.GoRaw("if hook, ok := value.(interface{ HxSerialize(*haxe__Serializer) }); ok { hook.HxSerialize(serializer); return true }"),
				GoStmt.GoRaw("return false")
			]));
		}
		if (includeUnserializer) {
			declarations.push(resolverDecl("resolveClass"));
			declarations.push(resolverDecl("resolveEnum"));
			declarations.push(GoDecl.GoFuncDecl("haxe__GoSerializationBridge_callUnserializeHook", null, [
				{name: "value", typeName: "any"},
				{name: "unserializer", typeName: "*haxe__Unserializer"}
			], ["bool"], [
				GoStmt.GoRaw("if hook, ok := value.(interface{ hxUnserialize(*haxe__Unserializer) }); ok { hook.hxUnserialize(unserializer); return true }"),
				GoStmt.GoRaw("if hook, ok := value.(interface{ HxUnserialize(*haxe__Unserializer) }); ok { hook.HxUnserialize(unserializer); return true }"),
				GoStmt.GoRaw("return false")
			]));
		}
		return declarations;
	}

	static function resolverDecl(methodName:String):GoDecl {
		var exportedName = methodName == "resolveClass" ? "ResolveClass" : "ResolveEnum";
		var symbol = "haxe__GoSerializationBridge_" + methodName;
		return GoDecl.GoFuncDecl(symbol, null, [{name: "resolver", typeName: "any"}, {name: "name", typeName: "*string"}], ["any"], [
			GoStmt.GoRaw("if resolver == nil { return nil }"),
			GoStmt.GoRaw("switch current := resolver.(type) {"),
			GoStmt.GoRaw("case interface{ " + methodName + "(*string) any }: return current." + methodName + "(name)"),
			GoStmt.GoRaw("case interface{ " + methodName + "(any) any }: return current." + methodName + "(name)"),
			GoStmt.GoRaw("case interface{ " + exportedName + "(*string) any }: return current." + exportedName + "(name)"),
			GoStmt.GoRaw("case interface{ " + exportedName + "(any) any }: return current." + exportedName + "(name)"),
			GoStmt.GoRaw("case map[string]any:"),
			GoStmt.GoRaw("\tif field := current[\"" + methodName + "\"]; field != nil {"),
			GoStmt.GoRaw("\t\tswitch callback := field.(type) { case func(*string) any: return callback(name); case func(any) any: return callback(name) }"),
			GoStmt.GoRaw("\t}"),
			GoStmt.GoRaw("case map[any]any:"),
			GoStmt.GoRaw("\tif field := current[\"" + methodName + "\"]; field != nil {"),
			GoStmt.GoRaw("\t\tswitch callback := field.(type) { case func(*string) any: return callback(name); case func(any) any: return callback(name) }"),
			GoStmt.GoRaw("\t}"),
			GoStmt.GoRaw("}"),
			GoStmt.GoRaw("return nil")
		]);
	}
}
#end
