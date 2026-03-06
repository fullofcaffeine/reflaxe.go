class Main {
	static function main() {
		var template = new haxe.Template("::name::");
		Sys.println(template.execute({name: "ok"}));
	}
}
