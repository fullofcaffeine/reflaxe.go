class Main {
	static function main() {
		var p = new sys.io.Process("echo", ["hello-sem"]);
		var line = p.stdout.readLine();
		p.close();
		Sys.println(line);

		var p2 = new sys.io.Process("echo", ["a:b:c"]);
		var line2 = p2.stdout.readLine();
		p2.close();
		Sys.println(line2);
	}
}
