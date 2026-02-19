import app.TodoApp;
import haxe.ds.List;
import haxe.io.Bytes;
import haxe.io.BytesBuffer;
import model.TodoItem;
import profile.TodoRuntime;
import sys.io.File;

class InteractiveCli {
	static inline final STATE_FILE = ".tui_todo_state.txt";

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
		Sys.println("  reset");
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
		Sys.println("state file: " + STATE_FILE + " (current directory)");
	}

	static function printUsage(runtime:TodoRuntime):Void {
		Sys.println("tui_todo command session (" + runtime.profileId() + ")");
		Sys.println("run scripted contract mode with: --scripted");
		Sys.println("commands:");
		Sys.println("  tui_todo reset");
		Sys.println("  tui_todo help");
		Sys.println("  tui_todo add 2 Write_profile_docs tag 1 docs list");
		if (runtime.supportsBatchAdd()) {
			Sys.println("  tui_todo batch 3 Ship_generated_go_sync Add_binary_matrix list");
		}
		Sys.println("generated-source invocation:");
		Sys.println("  go run . <command...>");
		Sys.println("state file: " + STATE_FILE + " (current directory)");
	}

	static function failUsage(message:String):Void {
		Sys.println("error: " + message);
		Sys.println("run `help` for command syntax");
	}

	static function clearState():Void {
		try {
			File.saveContent(STATE_FILE, "");
		} catch (_:Dynamic) {}
	}

	static function encodeField(raw:String):String {
		var out = new BytesBuffer();
		var bytes = Bytes.ofString(raw);
		var i = 0;
		while (i < bytes.length) {
			var code = bytes.get(i);
			if (code == 92) {
				out.addByte(92);
				out.addByte(92);
			} else if (code == 9) {
				out.addByte(92);
				out.addByte(116);
			} else if (code == 10) {
				out.addByte(92);
				out.addByte(110);
			} else if (code == 44) {
				out.addByte(92);
				out.addByte(99);
			} else {
				out.addByte(code);
			}
			i++;
		}
		return out.getBytes().toString();
	}

	static function splitRaw(raw:String, separatorCode:Int):List<String> {
		var out = new List<String>();
		var current = new BytesBuffer();
		var bytes = Bytes.ofString(raw);
		var i = 0;
		while (i < bytes.length) {
			var code = bytes.get(i);
			if (code == separatorCode) {
				out.add(current.getBytes().toString());
				current = new BytesBuffer();
			} else if (code != 13) {
				current.addByte(code);
			}
			i++;
		}
		out.add(current.getBytes().toString());
		return out;
	}

	static function splitEscaped(raw:String, separatorCode:Int):List<String> {
		var out = new List<String>();
		var current = new BytesBuffer();
		var bytes = Bytes.ofString(raw);
		var escaped = false;
		var i = 0;
		while (i < bytes.length) {
			var code = bytes.get(i);
			if (escaped) {
				if (code == 116) {
					current.addByte(9);
				} else if (code == 110) {
					current.addByte(10);
				} else if (code == 99) {
					current.addByte(44);
				} else if (code == 92) {
					current.addByte(92);
				} else {
					current.addByte(code);
				}
				escaped = false;
				i++;
				continue;
			}
			if (code == 92) {
				escaped = true;
				i++;
				continue;
			}
			if (code == separatorCode) {
				out.add(current.getBytes().toString());
				current = new BytesBuffer();
				i++;
				continue;
			}
			current.addByte(code);
			i++;
		}
		out.add(current.getBytes().toString());
		return out;
	}

	static function listIndex(values:List<String>, index:Int):String {
		var count = values.length;
		var i = 0;
		var out = "";
		while (i < count) {
			var value = values.pop();
			if (value == null) {
				break;
			}
			var entry:String = cast value;
			if (i == index) {
				out = entry;
			}
			values.add(entry);
			i++;
		}
		return out;
	}

	static function encodeTags(tags:List<String>):String {
		var out = "";
		var first = true;
		var count = tags.length;
		var i = 0;
		while (i < count) {
			var value = tags.pop();
			if (value == null) {
				break;
			}
			var tag:String = cast value;
			if (!first) {
				out += ",";
			}
			out += encodeField(tag);
			tags.add(tag);
			first = false;
			i++;
		}
		return out;
	}

	static function decodeTags(raw:String):List<String> {
		var out = new List<String>();
		if (raw == "") {
			return out;
		}
		var values = splitEscaped(raw, 44);
		var count = values.length;
		var i = 0;
		while (i < count) {
			var value = values.pop();
			if (value == null) {
				break;
			}
			var tag:String = cast value;
			if (tag != "") {
				out.add(tag);
			}
			values.add(tag);
			i++;
		}
		return out;
	}

	static function saveState(app:TodoApp):Void {
		var items = app.items();
		var out = "";
		var count = items.length;
		var i = 0;
		while (i < count) {
			var raw = items.pop();
			if (raw == null) {
				break;
			}
			var item:TodoItem = cast raw;
			out += encodeField(item.title) + "\t" + item.priority + "\t" + (item.done ? "1" : "0") + "\t" + encodeTags(item.tags) + "\n";
			items.add(item);
			i++;
		}
		File.saveContent(STATE_FILE, out);
	}

	static function loadState(app:TodoApp):Void {
		try {
			var raw = File.getContent(STATE_FILE);
			if (raw == "") {
				return;
			}

			var lines = splitRaw(raw, 10);
			var count = lines.length;
			var i = 0;
			while (i < count) {
				var lineValue = lines.pop();
				if (lineValue == null) {
					break;
				}
				var line:String = cast lineValue;
				if (line == "") {
					lines.add(line);
					i++;
					continue;
				}
				var fields = splitEscaped(line, 9);
				var title = listIndex(fields, 0);
				var priority = parsePositiveInt(listIndex(fields, 1));
				if (priority < 0) {
					priority = 0;
				}
				var done = listIndex(fields, 2) == "1";
				var id = app.add(title, priority);
				if (done) {
					app.toggle(id);
				}
				var tags = decodeTags(listIndex(fields, 3));
				var tagCount = tags.length;
				var j = 0;
				while (j < tagCount) {
					var tagValue = tags.pop();
					if (tagValue == null) {
						break;
					}
					var tag:String = cast tagValue;
					app.tag(id, tag);
					tags.add(tag);
					j++;
				}
				lines.add(line);
				i++;
			}
		} catch (_:Dynamic) {
			return;
		}
	}

	public static function run(runtime:TodoRuntime):Void {
		var app = new TodoApp(runtime);
		loadState(app);
		var args = Sys.args();
		if (args.length == 0) {
			printUsage(runtime);
			return;
		}

		var i = 0;
		while (i < args.length) {
			var cmd = args[i];
			if (cmd == "reset") {
				app = new TodoApp(runtime);
				clearState();
				Sys.println("ok reset");
				i++;
				continue;
			}
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
				saveState(app);
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
					saveState(app);
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
					saveState(app);
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
				if (added > 0) {
					saveState(app);
				}
				Sys.println("ok batch added=" + added);
				i += 4;
				continue;
			}
			failUsage("unknown command: " + cmd);
			return;
		}
	}
}
