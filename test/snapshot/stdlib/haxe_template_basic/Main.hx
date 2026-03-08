class Main {
	static function main() {
		var basic = new haxe.Template("::name::");
		Sys.println(basic.execute({name: "ok"}));

		var cond = new haxe.Template("::if enabled::yes::else::no::end::");
		Sys.println(cond.execute({enabled: true}));
		Sys.println(cond.execute({enabled: false}));

		var loop = new haxe.Template("::foreach items::::__current__::::end::");
		Sys.println(loop.execute({items: ["a", "b", "c"]}));
	}
}
