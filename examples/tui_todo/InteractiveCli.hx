import app.TodoApp;
import haxe.ds.List;
import haxe.io.Bytes;
import profile.TodoRuntime;

class InteractiveCli {
	static function parsePositiveInt(raw:String):Int {
		if (raw == "") {
			return -1;
		}
		var bytes = Bytes.ofString(raw);
		var value = 0;
		var i = 0;
		while (i < bytes.length) {
			var code = bytes.get(i);
			if (code < 48 || code > 57) {
				return -1;
			}
			value = (value * 10) + (code - 48);
			i++;
		}
		return value;
	}

	static function decodeToken(raw:String):String {
		return StringTools.replace(raw, "_", " ");
	}

	static function printHelp(runtime:TodoRuntime):Void {
		Sys.println("commands:");
		Sys.println("  help");
		Sys.println("  list");
		Sys.println("  summary");
		Sys.println("  diag");
		Sys.println("  add <priority> <title_token>");
		Sys.println("  toggle <id>");
		Sys.println("  tag <id> <tag_token>");
		if (runtime.supportsBatchAdd()) {
			Sys.println("  batch <priority> <title1_token> <title2_token>");
		}
		Sys.println("token note: use '_' instead of spaces (example: Wire_release_artifacts)");
	}

	static function printUsage(runtime:TodoRuntime):Void {
		Sys.println("tui_todo command session (" + runtime.profileId() + ")");
		Sys.println("run scripted contract mode with: --scripted");
		Sys.println("examples:");
		Sys.println("  go run . help");
		Sys.println("  go run . add 2 Write_profile_docs tag 1 docs list");
		if (runtime.supportsBatchAdd()) {
			Sys.println("  go run . batch 3 Ship_generated_go_sync Add_binary_matrix list");
		}
	}

	static function failUsage(message:String):Void {
		Sys.println("error: " + message);
		Sys.println("run `help` for command syntax");
	}

	public static function run(runtime:TodoRuntime):Void {
		var app = new TodoApp(runtime);
		var args = Sys.args();
		if (args.length == 0) {
			printUsage(runtime);
			return;
		}

		var i = 0;
		while (i < args.length) {
			var cmd = args[i];
			if (cmd == "help") {
				printHelp(runtime);
				i++;
				continue;
			}
			if (cmd == "list") {
				Sys.println(app.render());
				i++;
				continue;
			}
			if (cmd == "summary") {
				Sys.println(app.baselineSignature());
				i++;
				continue;
			}
			if (cmd == "diag") {
				Sys.println(app.diagnostics());
				i++;
				continue;
			}
			if (cmd == "add") {
				if (i + 2 >= args.length) {
					failUsage("add requires <priority> <title_token>");
					return;
				}
				var priority = parsePositiveInt(args[i + 1]);
				if (priority < 0) {
					failUsage("invalid priority: " + args[i + 1]);
					return;
				}
				var title = decodeToken(args[i + 2]);
				app.add(title, priority);
				Sys.println("ok add");
				i += 3;
				continue;
			}
			if (cmd == "toggle") {
				if (i + 1 >= args.length) {
					failUsage("toggle requires <id>");
					return;
				}
				var id = parsePositiveInt(args[i + 1]);
				if (id < 0) {
					failUsage("invalid id: " + args[i + 1]);
					return;
				}
				if (app.toggle(id)) {
					Sys.println("ok toggle");
				} else {
					Sys.println("missing id: " + id);
				}
				i += 2;
				continue;
			}
			if (cmd == "tag") {
				if (i + 2 >= args.length) {
					failUsage("tag requires <id> <tag_token>");
					return;
				}
				var id = parsePositiveInt(args[i + 1]);
				if (id < 0) {
					failUsage("invalid id: " + args[i + 1]);
					return;
				}
				var tag = decodeToken(args[i + 2]);
				if (app.tag(id, tag)) {
					Sys.println("ok tag");
				} else {
					Sys.println("missing id: " + id);
				}
				i += 3;
				continue;
			}
			if (cmd == "batch") {
				if (!runtime.supportsBatchAdd()) {
					Sys.println("batch not supported in " + runtime.profileId());
					i++;
					continue;
				}
				if (i + 3 >= args.length) {
					failUsage("batch requires <priority> <title1_token> <title2_token>");
					return;
				}
				var priority = parsePositiveInt(args[i + 1]);
				if (priority < 0) {
					failUsage("invalid priority: " + args[i + 1]);
					return;
				}
				var titles = new List<String>();
				titles.add(decodeToken(args[i + 2]));
				titles.add(decodeToken(args[i + 3]));
				var added = app.addMany(titles, priority);
				Sys.println("ok batch added=" + added);
				i += 4;
				continue;
			}
			failUsage("unknown command: " + cmd);
			return;
		}
	}
}
