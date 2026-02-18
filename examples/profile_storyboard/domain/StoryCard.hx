package domain;

import haxe.ds.List;

class StoryCard {
  public var id(default, null):Int;
  public var title(default, null):String;
  public var points(default, null):Int;
  public var tags(default, null):List<String>;

  public function new(id:Int, title:String, points:Int, tags:List<String>) {
    this.id = id;
    this.title = title;
    this.points = points;
    this.tags = tags;
  }
}
