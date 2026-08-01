#if macro
import haxe.macro.Context;
import haxe.macro.Expr;
#end

/** Compile-time selector for the one pinned official unitstd tracer body. */
class OfficialUnitStdAdapter {
	/**
		What: Expands the selected official unitstd statements into the target case.
		Why: `UnitBuilder.generateSpec` imports the entire shared specification module,
		which would add unrelated official surfaces to this narrow smoke.
		How: Delegate parsing and assertion rewriting to the pinned upstream
		`unit.UnitBuilder.read` implementation and require an explicit source path.
	**/
	public static macro function body():Expr {
		#if macro
		var path = Context.definedValue("official_haxe_smoke_unitstd_path");
		if (path == null || path == "") {
			Context.fatalError("missing official_haxe_smoke_unitstd_path", Context.currentPos());
		}
		return unit.UnitBuilder.read(path);
		#else
		return null;
		#end
	}
}
