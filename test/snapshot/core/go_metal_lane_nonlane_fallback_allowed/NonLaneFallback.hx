class NonLaneFallback {
	public static function run():Void {
		var maybe:Dynamic = go.Go.fail("non-lane");
		trace(maybe == null);
	}
}
