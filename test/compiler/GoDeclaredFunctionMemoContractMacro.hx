#if macro
import haxe.macro.Context;
import haxe.macro.Expr;
#end

/** Build macro of `GoDeclaredFunctionMemoContractFixture`: generated methods that share one position. */
class GoDeclaredFunctionMemoContractMacro {
	#if macro
	/**
		Adds `public static function made(step:Int = value)` and
		`public static function also(step:Int = value + 100)`, both at one position
		that is the same in every class it builds.
	**/
	public static function addMade(value:Int):Array<Field> {
		final fields = Context.getBuildFields();
		final shared = Context.makePosition({file: "GoDeclaredFunctionMemoContractMacro.hx", min: 0, max: 1});
		for (method in [{name: "made", value: value}, {name: "also", value: value + 100}])
			fields.push({
				name: method.name,
				access: [APublic, AStatic],
				pos: shared,
				kind: FFun({
					args: [{name: "step", type: macro :Int, value: macro $v{method.value}}],
					ret: macro :Int,
					expr: macro return step + $v{method.value}
				})
			});
		return fields;
	}
	#end
}
