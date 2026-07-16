package hxrt.date;

/**
	What:
	- Provides typed access to Go wall-clock, calendar, parser, formatter, and
	  timezone operations for the staged `Date` override.

	Why:
	- Host timezone rules and Go's native clock are genuine runtime capabilities,
	  but the public Haxe API and epoch-millisecond carrier belong in source.

	How:
	- Cross only scalar components, epoch milliseconds, strings, and the typed
	  `DateParts` carrier into `runtime/hxrt/date.go`. Native parse errors return
	  through the ordinary Haxe exception boundary.
**/
@:go.import("hxrt")
@:go.package("hxrt")
extern class NativeDate {
	@:go.name("DateLocalTime")
	public static function localTime(year:Int, month:Int, day:Int, hour:Int, min:Int, sec:Int):Float;

	@:go.name("DateNow")
	public static function now():Float;

	@:go.name("DateParse")
	public static function parse(value:String):Float;

	@:go.name("DateLocalParts")
	public static function localParts(milliseconds:Float):DateParts;

	@:go.name("DateUTCParts")
	public static function utcParts(milliseconds:Float):DateParts;

	@:go.name("DateTimezoneOffset")
	public static function timezoneOffset(milliseconds:Float):Int;

	@:go.name("DateFormatLocal")
	public static function formatLocal(milliseconds:Float):String;
}
