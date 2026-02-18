package profile;

import haxe.ds.List;
import model.TodoItem;

class GopherRuntime implements TodoRuntime {
  public function new() {}

  public function profileId():String {
    return "gopher";
  }

  public function normalizeTitle(title:String):String {
    return title;
  }

  public function normalizeTag(tag:String):String {
    return "go-" + tag;
  }

  public function supportsBatchAdd():Bool {
    return true;
  }

  public function supportsDiagnostics():Bool {
    return false;
  }

  public function diagnostics(items:List<TodoItem>):String {
    return "off";
  }
}
