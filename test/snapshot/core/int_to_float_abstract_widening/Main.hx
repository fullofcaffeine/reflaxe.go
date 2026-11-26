abstract Real(Float) from Float to Float {}
abstract Count(Int) from Int to Int {}

enum abstract Level(Int) to Int {
	var High = 3;
}

class Track {
	public var length:Real;

	public function new() {}
}

/**
	An Int value reaches a Float through an abstract (haxe_go-gjqe).

	Abstracts over Int and Float are stored as Go `any`. An Int that enters an
	abstract over Float is converted when it is written (`float64(three)`,
	`7.0`). A Float read of an abstract over Float uses
	`hxrt.DynamicToFloatSlot`, because an unsafe `cast` or a `Dynamic` value
	can still put an Int there. An abstract over Int that enters a Float slot is
	recovered as an Int and then converted (`float64(hxrt.IntFromNullableAny(count))`).
	Reads whose static type is Float or `Null<Float>` keep their plain
	conversion.
**/
class Main {
	static function halfReal(value:Real):Float
		return (value : Float) / 2;

	static function countAsFloat(count:Count):Float
		return count;

	static function main():Void {
		var seven:Dynamic = 7;
		var fromDynamic:Real = cast seven;
		Sys.println("real from dynamic: " + halfReal(fromDynamic));
		var three = 3;
		var fromLocal:Real = three;
		var fromLiteral:Real = 7;
		Sys.println("real from int: " + halfReal(fromLocal) + " " + halfReal(fromLiteral));
		var track = new Track();
		track.length = three;
		Sys.println("real field: " + halfReal(track.length));
		var count:Count = 5;
		Sys.println("count as float: " + countAsFloat(count) / 2);
		var level:Float = High;
		Sys.println("level as float: " + level / 2);
		var maybe:Null<Float> = 1.5;
		var sure:Float = maybe;
		Sys.println("typed float: " + sure / 2);
	}
}
