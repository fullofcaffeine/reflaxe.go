class NonLaneWorker {
	private static function unresolvedFail<T>(message:String):go.Result<T> {
		return go.Go.fail(message);
	}

	public static function produce():Void {
		// Intentional: force non-monomorphizable go.Result<T> to exercise non-lane fallback attribution.
		var nonLaneResult = unresolvedFail("non-lane");
		trace(nonLaneResult == null);
	}
}
