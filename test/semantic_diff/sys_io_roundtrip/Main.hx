class Main {
	static function main() {
		var path = "./semdiff_tmp.txt";
		sys.io.File.saveContent(path, "alpha\\nbeta\\n");

		var data = sys.io.File.getContent(path);
		Sys.println(data);

		var p = new sys.io.Process("echo", ["hello-sem"]);
		var line = p.stdout.readLine();
		p.close();
		Sys.println(line);
	}
}
