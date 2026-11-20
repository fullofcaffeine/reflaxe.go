package reflaxe.go.compiler.emit;

#if macro
import reflaxe.go.ast.GoAST.GoDecl;
import reflaxe.go.ast.GoAST.GoExpr;
import reflaxe.go.ast.GoAST.GoStmt;
import reflaxe.go.ast.GoBuiltinType;
import reflaxe.go.ast.GoCompositeElement;
import reflaxe.go.ast.GoType;

/**
	What:
	Builds the two small methods that let `hxrt.StdString` print generated enum
	values and class instances the way Haxe does (`Some(1)`, `pack.Point`).

	Why:
	`Std.string` receives erased values. The runtime can see an enum carrier's
	tag and arguments only as private fields, and it cannot see Haxe constructor
	names or class paths at all, so without these methods it printed Go's struct
	dump (`&{1 [1]}`, `&{0x14000110330 1 2}`). Go type names are not a reliable
	source: they are mangled and cannot be mapped back to Haxe paths.

	How:
	- `HxEnumValue() (string, []any)` on every enum carrier returns the
	  constructor name, read from a fixed array by tag, and the arguments.
	- `HxClassName() string` on every class without `toString()` returns the
	  dotted Haxe class path. Each class gets its own method, because a method
	  promoted from an embedded superclass would print the superclass name.
	  Classes with `toString()` keep their `String()` adapter instead.
	The names are exported because `hxrt` is a separate Go package; the runtime
	side lives in `runtime/hxrt/string.go`.
**/
class GoStdStringProtocolEmitter {
	public static function enumValueDecl(goTypeName:String, constructorNames:Array<String>):GoDecl {
		var receiver = GoExpr.GoIdent("self");
		var names = GoExpr.GoCompositeLiteral(GoType.array(constructorNames.length, GoType.builtin(GoBuiltinType.StringType)), [
			for (name in constructorNames) GoCompositeElement.GoCompositeValue(GoExpr.GoStringLiteral(name))
		]);
		return GoDecl.GoFuncDecl("HxEnumValue", {name: "self", typeName: GoType.parse("*" + goTypeName)}, [], [
			GoType.builtin(GoBuiltinType.StringType),
			GoType.slice(GoType.builtin(GoBuiltinType.AnyType))
		], [
			GoStmt.GoReturnValues([
				GoExpr.GoIndex(names, GoExpr.GoSelector(receiver, "tag")),
				GoExpr.GoSelector(receiver, "params")
			])
		]);
	}

	public static function classNameDecl(goTypeName:String, haxeClassPath:String):GoDecl {
		return GoDecl.GoFuncDecl("HxClassName", {name: "self", typeName: GoType.parse("*" + goTypeName)}, [], [GoType.builtin(GoBuiltinType.StringType)],
			[GoStmt.GoReturn(GoExpr.GoStringLiteral(haxeClassPath))]);
	}
}
#end
