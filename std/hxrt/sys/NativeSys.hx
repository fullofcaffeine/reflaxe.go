package hxrt.sys;

/**
	What
	- Typed bridge to the native process capabilities used by staged root `Sys`.

	Why
	- Process state, environment mutation, clocks, and commands require Go runtime
	  support. The public Haxe API, map construction, aliases, and locale fallback
	  do not belong in compiler shims.

	How
	- Map each operation one-for-one to an exported `runtime/hxrt/sys.go`
	  capability. Console display and standard-file handles use their own runtime
	  slices so selective builds do not acquire OS-process support just to print.
**/
@:go.import("hxrt")
@:go.package("hxrt")
extern class NativeSys {
	@:go.name("SysArgs")
	public static function args():Array<String>;

	@:go.name("SysGetEnv")
	public static function getEnv(key:String):String;

	@:go.name("SysSetEnvironment")
	public static function putEnv(key:String, value:Null<String>):Void;

	@:go.name("SysEnvironmentEntries")
	public static function environmentEntries():Array<SysEnvironmentEntry>;

	@:go.name("SysSleep")
	public static function sleep(seconds:Float):Void;

	@:go.name("SysGetCwd")
	public static function getCwd():String;

	@:go.name("SysChangeCwd")
	public static function setCwd(path:String):Void;

	@:go.name("SysSystemName")
	public static function systemName():String;

	@:go.name("SysCommand")
	public static function command(command:String, args:Array<String>):Int;

	@:go.name("SysExit")
	public static function exit(code:Int):Void;

	@:go.name("SysTime")
	public static function time():Float;

	@:go.name("SysCurrentProgramPath")
	public static function programPath():String;
}
