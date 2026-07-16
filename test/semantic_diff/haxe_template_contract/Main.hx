class Main {
	static function main() {
		var basic = new haxe.Template("::name::");
		Sys.println(basic.execute({name: "ok"}));

		var cond = new haxe.Template("::if enabled::yes::else::no::end::");
		Sys.println(cond.execute({enabled: true}));
		Sys.println(cond.execute({enabled: false}));

		var loop = new haxe.Template("::foreach items::::__current__::::end::");
		Sys.println(loop.execute({items: ["a", "b", "c"]}));

		var nested = new haxe.Template("Hello ::user.name::!");
		Sys.println(nested.execute({user: {name: "Go"}}));

		var records = new haxe.Template("::foreach items::::label::=::value::;::end::");
		Sys.println(records.execute({items: [{label: "a", value: 1}, {label: "b", value: 2}]}));

		var stacked = new haxe.Template("::foreach items::::prefix::::__current__::;::end::");
		Sys.println(stacked.execute({prefix: "p", items: [1, 2]}));

		var manualValues = ["x", "y"];
		var manualIndex = 0;
		var manualIterable:Dynamic = {
			iterator: function():Dynamic {
				return {
					hasNext: function():Bool return manualIndex < manualValues.length,
					next: function():Dynamic return manualValues[manualIndex++]
				};
			}
		};
		var manual = new haxe.Template("::foreach items::::__current__::::end::");
		Sys.println(manual.execute({items: manualIterable}));

		var macroTemplate = new haxe.Template("$$upper(name)");
		Sys.println(macroTemplate.execute({name: "go"}, {
			upper: function(resolve:String->Dynamic, value:String):String {
				return Std.string(resolve(value)).toUpperCase();
			}
		}));
	}
}
