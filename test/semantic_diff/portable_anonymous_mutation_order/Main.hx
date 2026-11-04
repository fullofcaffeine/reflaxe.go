typedef Record = {var count:Int; var text:String; var values:Array<Int>;}

class Main {
	static var record:Record = {count: 2, text: "a", values: [1]};
	static var order = "";

	static function target():Record {
		order += "L";
		return record;
	}

	static function right():Int {
		order += "R";
		return 3;
	}

	static function main() {
		Sys.println(target().count = right());
		Sys.println(order);
		order = "";
		Sys.println(target().count += right());
		Sys.println(order);
		order = "";
		Sys.println(target().count++);
		Sys.println(order);
		order = "";
		Sys.println(++target().count);
		Sys.println(order);
		target().text += "b";
		target().values.push(9);
		Sys.println(record.count);
		Sys.println(record.text);
		Sys.println(record.values.join(","));
	}
}
