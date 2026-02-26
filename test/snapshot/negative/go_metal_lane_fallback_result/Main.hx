@:goMetal
class Main {
	private static function unresolvedFail<T>(message:String):go.Result<T> {
		return go.Go.fail(message);
	}

	static function main() {
		// Intentional: unresolved go.Result<T> must trip metal fallback enforcement in lane modules.
		var laneResult = unresolvedFail("lane");
		trace(laneResult == null);
	}
}
