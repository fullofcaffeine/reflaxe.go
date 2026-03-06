import DateTools;

class Main {
	static function main() {
		var base = Date.fromString("2024-02-29 15:04:05");
		Sys.println(DateTools.format(base, "%D|%F|%R|%T|%r|%a|%A|%b|%h|%B|%C|%d|%e|%H|%k|%I|%l|%m|%M|%p|%S|%u|%w|%y|%Y|%%"));
		Sys.println(DateTools.getMonthDays(base));

		var stamp = DateTools.make({
			ms: 123.0,
			seconds: 5,
			minutes: 4,
			hours: 3,
			days: 2
		});
		var parsed = DateTools.parse(stamp);
		Sys.println(parsed.days);
		Sys.println(parsed.hours);
		Sys.println(parsed.minutes);
		Sys.println(parsed.seconds);
		Sys.println(parsed.ms);

		var shifted = DateTools.delta(base, DateTools.days(1) + DateTools.hours(2) + DateTools.minutes(3) + DateTools.seconds(4));
		Sys.println(DateTools.format(shifted, "%F %T"));
	}
}
