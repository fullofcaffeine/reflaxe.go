class NonLaneWorker {
	public static function produce():Void {
		// Intentional: force non-monomorphizable go.Result<T> to exercise non-lane fallback attribution.
		var nonLaneResult:Dynamic = go.Go.fail("non-lane");
		trace(nonLaneResult == null);
	}
}
