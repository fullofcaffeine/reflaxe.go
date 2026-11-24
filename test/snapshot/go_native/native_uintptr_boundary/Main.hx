import go.NativeSlice;
import go.NativeUintptr;

/** `(*os.File).Fd` returns a Go `uintptr` on every GOOS, including Windows handles. */
@:go.import("os")
@:go.package("os")
@:go.name("File")
extern class NativeFile {
	@:go.name("Fd") public function fd():NativeUintptr;
}

@:go.import("os")
@:go.package("os")
extern class NativeOs {
	@:go.name("Stdout") public static var stdout(default, never):NativeFile;
}

/** `runtime.Func` is the harmless standard-library consumer of a `uintptr` program counter. */
@:go.import("runtime")
@:go.package("runtime")
@:go.name("Func")
extern class NativeFunc {
	@:go.name("Name") public function name():String;
	@:go.name("Entry") public function entry():NativeUintptr;
}

/** A standard-library struct with exported `uintptr` fields. */
@:go.import("runtime")
@:go.package("runtime")
@:go.name("Frame")
@:go.struct
extern class NativeFrame {
	public function new();
	@:go.name("PC") public var pc:NativeUintptr;
	@:go.name("Entry") public var entry:NativeUintptr;
}

@:go.import("runtime")
@:go.package("runtime")
extern class NativeRuntime {
	/** Go: `func Caller(skip int) (pc uintptr, file string, line int, ok bool)`. */
	@:go.name("Caller")
	@:go.tupleReturn
	public static function caller(skip:Int):CallerResult;

	/** Go: `func Callers(skip int, pc []uintptr) int`. */
	@:go.name("Callers") public static function callers(skip:Int, pc:NativeSlice<NativeUintptr>):Int;

	/** Go: `func FuncForPC(pc uintptr) *Func`. */
	@:go.name("FuncForPC") public static function funcForPC(pc:NativeUintptr):Null<NativeFunc>;
}

/** The exact four results of runtime.Caller; `pc` keeps its native `uintptr` width. */
final class CallerResult {
	public final pc:NativeUintptr;
	public final file:String;
	public final line:Int;
	public final ok:Bool;

	public function new(pc:NativeUintptr, file:String, line:Int, ok:Bool) {
		this.pc = pc;
		this.file = file;
		this.line = line;
		this.ok = ok;
	}
}

/** An ordinary Haxe class field that stores the native carrier. */
class DescriptorHolder {
	public var fd:NativeUintptr;

	public function new(fd:NativeUintptr) {
		this.fd = fd;
	}
}

@:goNative
class Main {
	static function check(ok:Bool, label:String):Void {
		if (!ok)
			throw label;
	}

	/** Haxe parameters of the carrier type; `==` compares the native Go values. */
	static function sameValue(left:NativeUintptr, right:NativeUintptr):Bool {
		return left == right;
	}

	static function hasFunction(pc:NativeUintptr):Bool {
		final fn = NativeRuntime.funcForPC(pc);
		return fn != null && fn.name() != "";
	}

	static function main():Void {
		// Local: upstream bd reads term.GetSize(int(os.Stdout.Fd())).
		final stdoutFd:NativeUintptr = NativeOs.stdout.fd();
		final fdNumber = NativeUintptr.toInt(stdoutFd);
		check(sameValue(stdoutFd, NativeOs.stdout.fd()), "parameter");

		// Class field: read, then overwrite with a converted Int.
		final holder = new DescriptorHolder(stdoutFd);
		check(sameValue(holder.fd, stdoutFd), "class field read");
		holder.fd = NativeUintptr.fromInt(fdNumber + 41);
		check(NativeUintptr.toInt(holder.fd) == fdNumber + 41, "class field write");

		// Tuple-carrier field: runtime.Caller returns pc as uintptr.
		final caller = NativeRuntime.caller(0);
		check(caller.ok && caller.line > 0 && caller.file != "", "runtime.Caller");
		check(hasFunction(caller.pc), "tuple field to uintptr parameter");

		// Native struct field and native slice element.
		final frame = new NativeFrame();
		frame.pc = caller.pc;
		frame.entry = NativeRuntime.funcForPC(frame.pc).entry();
		check(sameValue(frame.pc, caller.pc) && NativeUintptr.toInt(frame.entry) != 0, "struct field");
		final pcs = NativeSlice.fromArray([for (_ in 0...4) NativeUintptr.fromInt(0)]);
		final count = NativeRuntime.callers(1, pcs);
		pcs[3] = pcs[0];
		check(count > 0 && hasFunction(pcs[3]), "slice element");

		Sys.println("stdout fd: " + fdNumber);
		Sys.println("round trip: " + NativeUintptr.toInt(NativeUintptr.fromInt(42)));
		Sys.println("native-uintptr:ok");
	}
}
