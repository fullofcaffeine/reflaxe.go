class Main {
	static function main():Void {
		var list = new List<Int>();
		list.add(4);
		list.push(5);
		list.push(6);
		Sys.println("list.len0=" + Std.string(list.length));
		Sys.println("list.first=" + Std.string(list.first()));
		Sys.println("list.last=" + Std.string(list.last()));
		Sys.println("list.pop0=" + Std.string(list.pop()));
		Sys.println("list.pop1=" + Std.string(list.pop()));
		Sys.println("list.len1=" + Std.string(list.length));
	}
}
