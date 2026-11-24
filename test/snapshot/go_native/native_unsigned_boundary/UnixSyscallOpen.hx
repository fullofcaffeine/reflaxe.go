import SyscallOpen;
import go.NativeUint32;

/** The Unix `syscall` calls, with the same signatures as `golang.org/x/sys/unix`. */
@:go.import("syscall")
@:go.package("syscall")
extern class UnixSyscall {
	@:go.name("O_RDONLY") public static var oRdonly(default, never):Int;

	/** Go (Unix): `func Open(path string, mode int, perm uint32) (fd int, err error)`. */
	@:go.name("Open")
	@:go.tupleReturn
	public static function open(path:String, mode:Int, perm:NativeUint32):UnixOpenResult;

	/** Go (Unix): `func Close(fd int) (err error)`. */
	@:go.name("Close") public static function close(fd:Int):Null<go.NativeError>;
}

/** The `(fd int, err error)` results of the Unix `syscall.Open`. */
@:goBuildConstraint("unix")
final class UnixOpenResult {
	public final fd:Int;
	public final err:Null<go.NativeError>;

	public function new(fd:Int, err:Null<go.NativeError>) {
		this.fd = fd;
		this.err = err;
	}
}

/** Passes a `uint32` variable, not a Go constant, as the `perm` argument. */
@:goBuildConstraint("unix")
@:keep
final class UnixSyscallOpen implements SyscallOpenPort {
	static final installed:Bool = SyscallOpen.install(new UnixSyscallOpen());

	public function new() {}

	public function openReadOnly(path:String, perm:NativeUint32):Bool {
		final opened = UnixSyscall.open(path, UnixSyscall.oRdonly, perm);
		if (opened.err != null)
			return false;
		return UnixSyscall.close(opened.fd) == null;
	}
}
