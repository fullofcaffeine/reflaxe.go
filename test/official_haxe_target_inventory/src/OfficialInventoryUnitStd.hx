#if macro
import haxe.macro.Context;
import haxe.macro.Expr;
#end

/**
	What: Expands one pinned official unitstd statement file into a typed case.
	Why: Unitstd files are statement specifications, not importable Haxe classes.
	How: The generated wrapper supplies a literal staged path and upstream
	UnitBuilder.read remains the authority for parsing and assertion rewriting.
**/
class OfficialInventoryUnitStd {
	public static macro function body(path:ExprOf<String>):Expr {
		#if macro
		var value = switch path.expr {
			case EConst(CString(value, _)): value;
			case _: Context.fatalError("official unitstd path must be a literal", path.pos);
		};
		return unit.UnitBuilder.read(value);
		#else
		return null;
		#end
	}
}
