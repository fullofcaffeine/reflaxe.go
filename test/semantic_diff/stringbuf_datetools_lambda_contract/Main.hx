import DateTools;
import Lambda;

class Main {
	static function main() {
		var buf = new StringBuf();
		buf.add("alpha");
		buf.add("-");
		buf.addSub("bravo-charlie", 0, 5);
		buf.addChar("0".code + 7);
		buf.add("!");
		Sys.println(buf.toString());

		var base = Date.fromString("2024-01-02 03:04:05");
		var shifted = DateTools.delta(base, DateTools.days(2) + DateTools.hours(4) + DateTools.minutes(30));
		Sys.println(DateTools.format(base, "%Y-%m-%d %H:%M:%S"));
		Sys.println(DateTools.format(shifted, "%Y-%m-%d %H:%M:%S"));
		Sys.println(base.getDate());
		Sys.println(shifted.getDate());
		Sys.println(shifted.getHours());

		var values = [1, 2, 3, 4, 5];
		var doubled = Lambda.map(values, function(v:Int) return v * 2);
		var total = 0;
		for (value in doubled) {
			total += value;
		}
		Sys.println(total);
		Sys.println(doubled.length);
	}
}
