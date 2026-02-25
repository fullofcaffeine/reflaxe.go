@:goMetal
class LaneClean {
	public static function run():Void {
		var ok:go.Result<String> = go.Go.ok("clean");
		trace(ok.isOk());
	}
}
