package model;

import haxe.ds.List;

class TodoStore {
  var nextId:Int;
  var entries:List<TodoItem>;

  public function new() {
    nextId = 1;
    entries = new List<TodoItem>();
  }

  public function add(title:String, priority:Int):TodoItem {
    var item = new TodoItem(nextId, title, priority);
    nextId++;
    entries.add(item);
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
    item.tags.add(tag);
    return true;
  }

  public function list():List<TodoItem> {
    return entries;
  }

  public function totalCount():Int {
    return entries.length;
  }

  public function openCount():Int {
    var total = 0;
    var count = entries.length;
    var i = 0;
    while (i < count) {
      var value = entries.pop();
      if (value == null) {
        break;
      }
      var item:TodoItem = cast value;
      if (!item.done) {
        total++;
      }
      entries.add(item);
      i++;
    }
    return total;
  }

  public function doneCount():Int {
    var total = 0;
    var count = entries.length;
    var i = 0;
    while (i < count) {
      var value = entries.pop();
      if (value == null) {
        break;
      }
      var item:TodoItem = cast value;
      if (item.done) {
        total++;
      }
      entries.add(item);
      i++;
    }
    return total;
  }

  function findById(id:Int):Null<TodoItem> {
    var found:Null<TodoItem> = null;
    var count = entries.length;
    var i = 0;
    while (i < count) {
      var value = entries.pop();
      if (value == null) {
        break;
      }
      var item:TodoItem = cast value;
      if (item.id == id) {
        found = item;
      }
      entries.add(item);
      i++;
    }
    return found;
  }
}
