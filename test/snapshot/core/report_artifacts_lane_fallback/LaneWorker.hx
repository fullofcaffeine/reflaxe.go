@:goMetal
class LaneWorker {
	public static function produce():Void {
		// Intentional: force non-monomorphizable go.Result<T> to exercise metal fallback reporting.
		var laneResult:Dynamic = go.Go.fail("lane");
		trace(laneResult == null);
	}
}
