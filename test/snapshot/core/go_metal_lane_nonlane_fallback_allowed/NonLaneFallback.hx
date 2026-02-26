class NonLaneFallback {
	private static function unresolvedFail<T>(message:String):go.Result<T> {
		return go.Go.fail(message);
	}

	public static function run():Void {
		// Intentional: unresolved go.Result<T> should remain allowed outside @:goMetal lanes.
		var nonLaneResult = unresolvedFail("non-lane");
		trace(nonLaneResult == null);
	}
}
