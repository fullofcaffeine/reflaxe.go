import go.Chan;
import go.Go;

/**
	Why
	Portable builds guard unscoped `go.*` authority by default, while an explicit
	native module boundary must remain usable without changing the global preset.

	What
	Exercises canonical `@:goNative` authority under the portable/guarded policy.

	How
	Uses a concrete typed channel so the native boundary compiles, runs, and keeps
	its authority local to this module.
**/
@:goNative
class Main {
	static function main():Void {
		var channel:Chan<Int> = Go.newChan(1);
		channel.send(11);
		Sys.println(channel.recv());
		channel.close();
	}
}
