package model;

class TodoStore {
	var nextId:Int;
	var entries:Array<TodoItem>;

	public function new() {
		nextId = 1;
		entries = [];
	}

	public function add(title:String, priority:Int):TodoItem {
		var item = new TodoItem(nextId, title, priority);
		nextId++;
		entries.push(item);
		return item;
	}

	public function toggle(id:Int):Bool {
		var item = findById(id);
		if (item == null) {
			return false;
		}
		item.done = !item.done;
		return true;
	}

	public function addTag(id:Int, tag:String):Bool {
		var item = findById(id);
		if (item == null) {
			return false;
		}
		item.tags.push(tag);
		return true;
	}

	public function list():Array<TodoItem> {
		return entries;
	}

	public function totalCount():Int {
		return entries.length;
	}

	public function openCount():Int {
		var total = 0;
		for (item in entries) {
			if (!item.done) {
				total++;
			}
		}
		return total;
	}

	public function doneCount():Int {
		var total = 0;
		for (item in entries) {
			if (item.done) {
				total++;
			}
		}
		return total;
	}

	function findById(id:Int):Null<TodoItem> {
		for (item in entries) {
			if (item.id == id) {
				return item;
			}
		}
		return null;
	}
}
