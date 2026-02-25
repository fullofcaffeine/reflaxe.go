@:goMetal
class LaneWorker {
	private static function unresolvedFail<T>(message:String):go.Result<T> {
		return go.Go.fail(message);
	}

	public static function produce():Void {
		// Intentional: force non-monomorphizable go.Result<T> to exercise metal fallback reporting.
		var laneResult = unresolvedFail("lane");
		trace(laneResult == null);
	}
}
