class DateTools {
	static var DAY_SHORT_NAMES = ["Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat"];
	static var DAY_NAMES = ["Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"];
	static var MONTH_SHORT_NAMES = [
		"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"
	];
	static var MONTH_NAMES = [
		"January",
		"February",
		"March",
		"April",
		"May",
		"June",
		"July",
		"August",
		"September",
		"October",
		"November",
		"December"
	];

	private static function __format_get(d:Date, e:String):String {
		if (e == "%") {
			return "%";
		}
		if (e == "a") {
			return DAY_SHORT_NAMES[d.getDay()];
		}
		if (e == "A") {
			return DAY_NAMES[d.getDay()];
		}
		if (e == "b" || e == "h") {
			return MONTH_SHORT_NAMES[d.getMonth()];
		}
		if (e == "B") {
			return MONTH_NAMES[d.getMonth()];
		}
		if (e == "C") {
			return StringTools.lpad(Std.string(Math.floor(d.getFullYear() / 100)), "0", 2);
		}
		if (e == "d") {
			return StringTools.lpad(Std.string(d.getDate()), "0", 2);
		}
		if (e == "D") {
			return __format(d, "%m/%d/%y");
		}
		if (e == "e") {
			return Std.string(d.getDate());
		}
		if (e == "F") {
			return __format(d, "%Y-%m-%d");
		}
		if (e == "H" || e == "k") {
			return StringTools.lpad(Std.string(d.getHours()), if (e == "H") "0" else " ", 2);
		}
		if (e == "I" || e == "l") {
			var hour = d.getHours() % 12;
			return StringTools.lpad(Std.string(hour == 0 ? 12 : hour), if (e == "I") "0" else " ", 2);
		}
		if (e == "m") {
			return StringTools.lpad(Std.string(d.getMonth() + 1), "0", 2);
		}
		if (e == "M") {
			return StringTools.lpad(Std.string(d.getMinutes()), "0", 2);
		}
		if (e == "n") {
			return "\n";
		}
		if (e == "p") {
			return if (d.getHours() > 11) "PM" else "AM";
		}
		if (e == "r") {
			return __format(d, "%I:%M:%S %p");
		}
		if (e == "R") {
			return __format(d, "%H:%M");
		}
		if (e == "s") {
			return Std.string(Math.floor(d.getTime() / 1000));
		}
		if (e == "S") {
			return StringTools.lpad(Std.string(d.getSeconds()), "0", 2);
		}
		if (e == "t") {
			return "\t";
		}
		if (e == "T") {
			return __format(d, "%H:%M:%S");
		}
		if (e == "u") {
			var day = d.getDay();
			return if (day == 0) "7" else Std.string(day);
		}
		if (e == "w") {
			return Std.string(d.getDay());
		}
		if (e == "y") {
			return StringTools.lpad(Std.string(d.getFullYear() % 100), "0", 2);
		}
		if (e == "Y") {
			return Std.string(d.getFullYear());
		}
		throw "Date.format %" + e + " not implemented yet.";
	}

	private static function __format(d:Date, f:String):String {
		var result = new StringBuf();
		var pos = 0;
		var length = f.length;

		while (pos < length) {
			if (f.charAt(pos) == "%") {
				result.add(__format_get(d, f.substr(pos + 1, 1)));
				pos += 2;
				continue;
			}
			result.add(f.charAt(pos));
			pos++;
		}
		return result.toString();
	}

	public static function format(d:Date, f:String):String {
		return __format(d, f);
	}

	public static inline function delta(d:Date, t:Float):Date {
		return Date.fromTime(d.getTime() + t);
	}

	static var DAYS_OF_MONTH = [31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31];

	public static function getMonthDays(d:Date):Int {
		var month = d.getMonth();
		var year = d.getFullYear();

		if (month != 1) {
			return DAYS_OF_MONTH[month];
		}

		var isLeap = ((year % 4 == 0) && (year % 100 != 0)) || (year % 400 == 0);
		return isLeap ? 29 : 28;
	}

	public static inline function seconds(n:Float):Float {
		return n * 1000.0;
	}

	public static inline function minutes(n:Float):Float {
		return n * 60.0 * 1000.0;
	}

	public static inline function hours(n:Float):Float {
		return n * 60.0 * 60.0 * 1000.0;
	}

	public static inline function days(n:Float):Float {
		return n * 24.0 * 60.0 * 60.0 * 1000.0;
	}

	public static function parse(t:Float) {
		var s = t / 1000;
		var m = s / 60;
		var h = m / 60;
		return {
			ms: t % 1000,
			seconds: Math.floor(s % 60),
			minutes: Math.floor(m % 60),
			hours: Math.floor(h % 24),
			days: Math.floor(h / 24),
		};
	}

	public static function make(o:{
		ms:Float,
		seconds:Int,
		minutes:Int,
		hours:Int,
		days:Int
	}) {
		return o.ms + 1000.0 * (o.seconds + 60.0 * (o.minutes + 60.0 * (o.hours + 24.0 * o.days)));
	}
}
