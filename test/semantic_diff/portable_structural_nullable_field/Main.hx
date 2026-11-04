class Main {
	static function main() {
		var value:{var count:Null<Int>;} = {count: null};
		Sys.println(value.count);
		Sys.println(Reflect.field(value, "count"));
		value.count = 7;
		Sys.println(value.count);
	}
}
