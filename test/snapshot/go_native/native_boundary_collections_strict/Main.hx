/**
	What: Exercises concrete Go collection APIs inside a strict native boundary.
	Why: These APIs have Go-target behavior and therefore cannot use Haxe --interp
		  as a semantic oracle.
	How: A target-runtime snapshot checks the generated program's deterministic
		 output while auto_strict separately enforces that specialization succeeds.
**/
@:goNative
class NativeCollectionOps {
	public static function eval():String {
		var channel:go.Chan<Int> = go.Go.newChan(1);
		channel.send(7);
		var received = channel.recvOr(-1);

		var slice:go.Slice<Int> = go.Go.newSlice();
		slice.push(received);
		slice.set(0, slice.get(0) + 1);

		var map:go.Map<Int, Int> = go.Go.newMap();
		map.set(1, slice.get(0));
		return Std.string(map.get(1)) + "|" + (map.exists(1) ? "1" : "0");
	}
}

/** Target-runtime companion for the strict, concrete go.Result boundary. */
@:goNative
class NativeResultOps {
	public static function eval():String {
		var ok:go.Result<String> = go.Go.ok("done");
		var err:go.Result<String> = go.Go.fail("broken");
		var errValue = err.error();
		var errLabel = errValue == null ? "none" : errValue;
		return ok.unwrap() + "|" + (ok.isOk() ? "1" : "0") + "|" + (err.isErr() ? "1" : "0") + "|" + errLabel;
	}
}

class Main {
	static function main() {
		Sys.println(NativeCollectionOps.eval());
		Sys.println(NativeResultOps.eval());
	}
}
