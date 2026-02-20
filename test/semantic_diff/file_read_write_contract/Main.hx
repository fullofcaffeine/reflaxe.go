class Main {
	static function main() {
		var path = "semantic_file_contract.txt";

		sys.io.File.saveContent(path, "alpha\nbeta");
		var first = sys.io.File.getContent(path);
		Sys.println("first=" + StringTools.replace(first, "\n", "|"));

		sys.io.File.saveContent(path, "gamma");
		var second = sys.io.File.getContent(path);
		Sys.println("second=" + second);
		Sys.println("second.len=" + second.length);
	}
}
