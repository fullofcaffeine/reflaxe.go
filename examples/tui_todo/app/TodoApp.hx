package app;

import haxe.ds.List;
import model.TodoItem;
import model.TodoStore;
import profile.TodoRuntime;
import profile.TodoRuntimeMetrics;

class TodoApp {
	var runtime:TodoRuntime;
	var store:TodoStore;

	static function joinStringList(values:List<String>, separator:String):String {
		var out = "";
		var first = true;
		var count = values.length;
		var i = 0;
		while (i < count) {
			var raw = values.pop();
			if (raw == null) {
				break;
			}
			var value:String = cast raw;
			if (!first) {
				out += separator;
			}
			out += value;
			values.add(value);
			first = false;
			i++;
		}
		return out;
	}

	public function new(runtime:TodoRuntime) {
		this.runtime = runtime;
		this.store = new TodoStore();
	}

	public function add(title:String, priority:Int):Int {
		var item = store.add(runtime.normalizeTitle(title), priority);
		return item.id;
	}

	public function addMany(titles:List<String>, priority:Int):Int {
		if (!runtime.supportsBatchAdd()) {
			return 0;
		}

		var added = 0;
		var count = titles.length;
		var i = 0;
		while (i < count) {
			var raw = titles.pop();
			if (raw == null) {
				break;
			}
			var title:String = cast raw;
			add(title, priority);
			titles.add(title);
			added++;
			i++;
		}
		return added;
	}

	public function toggle(id:Int):Bool {
		return store.toggle(id);
	}

	public function tag(id:Int, tag:String):Bool {
		return store.addTag(id, runtime.normalizeTag(tag));
	}

	public function baselineSignature():String {
		return "open=" + openCount() + ",done=" + doneCount() + ",total=" + totalCount();
	}

	public function totalCount():Int {
		return store.totalCount();
	}

	public function openCount():Int {
		return store.openCount();
	}

	public function doneCount():Int {
		return store.doneCount();
	}

	public function diagnostics():String {
		if (!runtime.supportsDiagnostics()) {
			return "off";
		}
		return runtime.diagnostics(buildRuntimeMetrics());
	}

	function buildRuntimeMetrics():TodoRuntimeMetrics {
		var items = store.list();
		var total = items.length;
		var done = 0;
		var p1 = 0;
		var i = 0;
		while (i < total) {
			var value = items.pop();
			if (value == null) {
				break;
			}
			var item:TodoItem = cast value;
			if (item.done) {
				done++;
			}
			if (item.priority == 1) {
				p1++;
			}
			items.add(item);
			i++;
		}
		return new TodoRuntimeMetrics(total, total - done, done, p1);
	}

	public function render():String {
		var out = "== TODO ==";
		var items = store.list();
		var count = items.length;
		var i = 0;
		while (i < count) {
			var raw = items.pop();
			if (raw == null) {
				break;
			}
			var item:TodoItem = cast raw;

			var state = "[ ]";
			if (item.done) {
				state = "[x]";
			}

			var tags = "-";
			if (item.tags.length != 0) {
				tags = joinStringList(item.tags, ",");
			}

			out += "\n" + state + " #" + item.id + " p" + item.priority + " " + item.title + " tags:" + tags;
			items.add(item);
			i++;
		}

		out += "\nsummary " + baselineSignature();
		return out;
	}

	public function items():List<TodoItem> {
		return store.list();
	}
}
