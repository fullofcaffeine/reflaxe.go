package app;

import model.TodoItem;
import model.TodoStore;
import profile.TodoRuntime;
import profile.TodoRuntimeMetrics;

class TodoApp {
	var runtime:TodoRuntime;
	var store:TodoStore;

	static function joinStringList(values:Array<String>, separator:String):String {
		var out = "";
		var first = true;
		for (value in values) {
			if (!first) {
				out += separator;
			}
			out += value;
			first = false;
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

	public function addMany(titles:Array<String>, priority:Int):Int {
		var added = 0;
		for (title in titles) {
			add(title, priority);
			added++;
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
		return runtime.diagnostics(buildRuntimeMetrics());
	}

	function buildRuntimeMetrics():TodoRuntimeMetrics {
		var items = store.list();
		var total = items.length;
		var done = 0;
		var p1 = 0;
		for (item in items) {
			if (item.done) {
				done++;
			}
			if (item.priority == 1) {
				p1++;
			}
		}
		return new TodoRuntimeMetrics(total, total - done, done, p1);
	}

	public function render():String {
		var out = "== TODO ==";
		var items = store.list();
		for (item in items) {
			var state = "[ ]";
			if (item.done) {
				state = "[x]";
			}

			var tags = "-";
			if (item.tags.length != 0) {
				tags = joinStringList(item.tags, ",");
			}

			out += "\n" + state + " #" + item.id + " p" + item.priority + " " + item.title + " tags:" + tags;
		}

		out += "\nsummary " + baselineSignature();
		return out;
	}

	public function items():Array<TodoItem> {
		return store.list();
	}
}
