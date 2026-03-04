package model;

class TodoItem {
	public var id(default, null):Int;
	public var title(default, set):String;
	public var done(default, set):Bool;
	public var priority(default, set):Int;
	public var tags(default, null):Array<String>;

	public function new(id:Int, title:String, priority:Int) {
		this.id = id;
		this.title = title;
		this.done = false;
		this.priority = priority;
		this.tags = [];
	}

	function set_title(value:String):String {
		title = value;
		return value;
	}

	function set_done(value:Bool):Bool {
		done = value;
		return value;
	}

	function set_priority(value:Int):Int {
		priority = value;
		return value;
	}
}
