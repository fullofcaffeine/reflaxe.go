package domain;

import haxe.ds.List;

class StoryCard {
	public var id(default, null):Int;
	public var title(default, null):String;
	public var points(default, null):Int;
	public var tags(default, null):List<String>;
	public var state(default, null):String;
	public var owner(default, null):String;

	public function new(id:Int, title:String, points:Int, tags:List<String>, state:String, owner:String) {
		this.id = id;
		this.title = title;
		this.points = points;
		this.tags = tags;
		this.state = state;
		this.owner = owner;
	}
}
