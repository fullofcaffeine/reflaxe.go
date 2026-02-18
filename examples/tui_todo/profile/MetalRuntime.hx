package profile;

import haxe.ds.List;
import model.TodoItem;

class MetalRuntime implements TodoRuntime {
  public function new() {}

  public function profileId():String {
    return "metal";
  }

  public function normalizeTitle(title:String):String {
    return title;
  }

  public function normalizeTag(tag:String):String {
    return "metal-" + tag;
  }

  public function supportsBatchAdd():Bool {
    return true;
  }

  public function supportsDiagnostics():Bool {
    return true;
  }

  public function diagnostics(items:List<TodoItem>):String {
    var p1 = 0;
    var completed = 0;
    var count = items.length;
    var i = 0;
    while (i < count) {
      var value = items.pop();
      if (value == null) {
        break;
      }
      var item:TodoItem = cast value;
      if (item.priority == 1) {
        p1++;
      }
      if (item.done) {
        completed++;
      }
      items.add(item);
      i++;
    }
    return "p1=" + p1 + ",completed=" + completed;
  }
}
