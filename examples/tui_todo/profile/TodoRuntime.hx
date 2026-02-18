package profile;

import haxe.ds.List;
import model.TodoItem;

interface TodoRuntime {
  public function profileId():String;
  public function normalizeTitle(title:String):String;
  public function normalizeTag(tag:String):String;
  public function supportsBatchAdd():Bool;
  public function supportsDiagnostics():Bool;
  public function diagnostics(items:List<TodoItem>):String;
}
