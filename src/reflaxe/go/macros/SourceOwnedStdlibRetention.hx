package reflaxe.go.macros;

#if macro
import haxe.macro.Compiler as MacroCompiler;
#end

/**
	What
	- Retains the small compiler-owned `haxe.io` surface that staged stream
	  subclasses may override or call.

	Why
	- Reflaxe discovers indirect staged-stdlib dependencies while lowering typed
	  code, after Haxe dead-code elimination has already pruned unused base fields.
	  Without type-only retention, a source-owned stream reached indirectly through
	  compiler-owned `Sys` cannot be type-checked even though the methods are part of
	  Haxe's public `Input`, `Output`, and `Bytes` contracts.

	How
	- Attach `@:keep` to exact public fields during compiler initialization, before
	  any of the types are loaded. This preserves typed authority for staged source;
	  it does not emit target behavior or add File-specific compiler lowering.
**/
class SourceOwnedStdlibRetention {
	#if macro
	static final retainedFields = [
		"haxe.io.Bytes.get",
		"haxe.io.Bytes.set",
		"haxe.io.Input.close",
		"haxe.io.Input.readByte",
		"haxe.io.Input.readBytes",
		"haxe.io.Output.close",
		"haxe.io.Output.flush",
		"haxe.io.Output.writeByte",
		"haxe.io.Output.writeBytes"
	];

	public static function init():Void {
		for (fieldPath in retainedFields)
			MacroCompiler.addGlobalMetadata(fieldPath, "@:keep", false, false, true);
	}
	#else
	public static function init():Void {}
	#end
}
