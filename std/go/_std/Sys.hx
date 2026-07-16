import haxe.ds.StringMap;
import haxe.io.Eof;
import hxrt.fs.NativeFile;
import hxrt.sys.NativeConsole;
import hxrt.sys.NativeSys;
import hxrt.sys.NativeTerminal;
import sys.io.FileInput;
import sys.io.FileOutput;

/**
	What
	- Owns the complete Haxe 4.3.7 root `Sys` API selected by `haxe.go`.

	Why
	- The mainstream Haxe stdlib implementation cannot be used unchanged on
	  `haxe.go` because it is an extern contract whose target symbols were previously
	  synthesized by `GoCompiler`. Public map construction, aliases, stream wrapping,
	  EOF behavior, and portable fallbacks must be reviewable typed Haxe source.

	How
	- Delegate only native capabilities to narrow typed `hxrt` bridges, build Haxe
	  values and stream wrappers here, and keep unsupported CPU-time accounting an
	  explicit compiler diagnostic instead of substituting wall-clock time. Simple
	  one-step capabilities are inline so direct calls retain Go-shaped output while
	  first-class function references still materialize the source-owned contract.
**/
class Sys {
	/**
		What: Print one Haxe value without a newline.
		Why: The upstream API requires `Dynamic`; that untyped value must remain confined to the display boundary.
		How: Forward it once to the typed native print capability.
	**/
	public static inline function print(v:Dynamic):Void {
		NativeConsole.print(v);
	}

	/** Print one Haxe value followed by the platform newline through the native display boundary. **/
	public static inline function println(v:Dynamic):Void {
		NativeConsole.println(v);
	}

	public static inline function args():Array<String> {
		return NativeSys.args();
	}

	public static inline function getEnv(s:String):String {
		return NativeSys.getEnv(s);
	}

	/**
		What: Set or remove one process environment variable.
		Why: Haxe 4.3.7 eval exposes a non-throwing `Void` contract even for malformed keys.
		How: Use the dedicated portable runtime capability, which intentionally discards the retained native Go error.
	**/
	public static inline function putEnv(s:String, v:Null<String>):Void {
		NativeSys.putEnv(s, v);
	}

	/** Build the public Haxe map from typed runtime entries instead of generated-map internals. **/
	public static function environment():Map<String, String> {
		var environment = new StringMap<String>();
		for (entry in NativeSys.environmentEntries())
			environment.set(entry.key, entry.value);
		return environment;
	}

	public static inline function sleep(seconds:Float):Void {
		NativeSys.sleep(seconds);
	}

	/** Go has no process-global time-locale switch, so the target reports the portable fallback honestly. **/
	public static inline function setTimeLocale(loc:String):Bool {
		var ignored = loc;
		return false;
	}

	public static inline function getCwd():String {
		return NativeSys.getCwd();
	}

	public static inline function setCwd(s:String):Void {
		NativeSys.setCwd(s);
	}

	public static inline function systemName():String {
		return NativeSys.systemName();
	}

	/**
		What: Run a shell command or direct executable and return its exit code.
		Why: A null argument array selects host-shell parsing in the upstream contract; a non-null array must bypass the shell.
		How: Preserve nullability across the typed runtime boundary so `hxrt` can choose the correct launch mode.
	**/
	public static inline function command(cmd:String, ?args:Array<String>):Int {
		return NativeSys.command(cmd, args);
	}

	public static inline function exit(code:Int):Void {
		NativeSys.exit(code);
	}

	public static inline function time():Float {
		return NativeSys.time();
	}

	/**
		What: Preserve the upstream CPU-time member for API typing.
		Why: Go has no portable standard-library process CPU clock; wall time would violate the contract.
		How: Direct use is rejected by the compiler with the established actionable diagnostic.
	**/
	public static function cpuTime():Float {
		throw "Sys.cpuTime is unsupported on haxe.go: process CPU time is not implemented";
	}

	@:deprecated("Use programPath instead")
	public static inline function executablePath():String {
		return programPath();
	}

	public static inline function programPath():String {
		return NativeSys.programPath();
	}

	/**
		What: Read one byte immediately from standard input and optionally echo it once.
		Why: The upstream extern requires interactive character mode; ordinary
		`FileInput.readByte()` remains host-line-buffered on a terminal.
		How: Delegate only terminal control and the native byte read to `hxrt`,
		construct Haxe EOF here, and keep the echo write in staged source.
	**/
	public static function getChar(echo:Bool):Int {
		var value = NativeTerminal.readChar();
		if (value < 0)
			throw new Eof();
		if (echo)
			stdout().writeByte(value);
		return value;
	}

	public static inline function stdin():haxe.io.Input {
		return new FileInput(NativeFile.stdin());
	}

	public static inline function stdout():haxe.io.Output {
		return new FileOutput(NativeFile.stdout());
	}

	public static inline function stderr():haxe.io.Output {
		return new FileOutput(NativeFile.stderr());
	}
}
