@:goMetal
class LaneCollectionsOps {
	public static function eval():String {
		var ch:go.Chan<Int> = go.Go.newChan(1);
		ch.send(7);
		var received = ch.recvOr(-1);

		var slice:go.Slice<Int> = go.Go.newSlice();
		slice.push(received);
		slice.set(0, slice.get(0) + 1);

		var map:go.Map<Int, Int> = go.Go.newMap();
		map.set(1, slice.get(0));
		var value = map.get(1);
		var exists = map.exists(1);

		return Std.string(value) + "|" + (exists ? "1" : "0");
	}
}

class Main {
	static function main() {
		Sys.println(LaneCollectionsOps.eval());
	}
}
