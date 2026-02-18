package profile;

import haxe.ds.List;
import model.TodoItem;

class PortableRuntime implements TodoRuntime {
  public function new() {}

  public function profileId():String {
    return "portable";
  }

  public function normalizeTitle(title:String):String {
    return title;
  }

  public function normalizeTag(tag:String):String {
    return tag;
  }

  public function supportsBatchAdd():Bool {
    return false;
  }

  public function supportsDiagnostics():Bool {
    return false;
  }

  public function diagnostics(items:List<TodoItem>):String {
    return "off";
  }
}
