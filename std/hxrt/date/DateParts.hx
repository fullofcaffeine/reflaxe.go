package hxrt.date;

/**
	What:
	- Carries one local or UTC calendar decomposition from `hxrt` into staged
	  `Date` source.

	Why:
	- Go's `time.Time` must not leak into the generated Haxe object layout, while
	  seven separate native calls would repeat the same timezone conversion.

	How:
	- Map public fields directly to the representation-neutral `hxrt.DateParts`
	  struct. Months and weekdays already use Haxe's zero-based conventions.
**/
@:go.import("hxrt")
@:go.package("hxrt")
@:go.name("DateParts")
extern class DateParts {
	@:go.name("FullYear")
	public var fullYear:Int;

	@:go.name("Month")
	public var month:Int;

	@:go.name("Date")
	public var date:Int;

	@:go.name("Day")
	public var day:Int;

	@:go.name("Hours")
	public var hours:Int;

	@:go.name("Minutes")
	public var minutes:Int;

	@:go.name("Seconds")
	public var seconds:Int;
}
