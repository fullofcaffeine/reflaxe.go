import SyscallOpen;
import go.NativeUint32;

/** `syscall.Handle` is the named Go type `type Handle uintptr`; it stays a separate value type. */
@:go.import("syscall")
@:go.package("syscall")
@:go.name("Handle")
@:go.valueType
extern class WindowsHandle {}

/** The Windows `syscall` calls, with the same signatures as `golang.org/x/sys/windows`. */
@:go.import("syscall")
@:go.package("syscall")
extern class WindowsSyscall {
	@:go.name("O_RDONLY") public static var oRdonly(default, never):Int;

	/** The Go constant that `GetFileType` returns for a disk file. */
	@:go.name("FILE_TYPE_DISK") public static var fileTypeDisk(default, never):NativeUint32;

	/** Go (Windows): `func Open(path string, mode int, perm uint32) (fd Handle, err error)`. */
	@:go.name("Open")
	@:go.tupleReturn
	@:go.tupleValueResults("0")
	public static function open(path:String, mode:Int, perm:NativeUint32):WindowsOpenResult;

	/** Go (Windows): `func GetFileType(filehandle Handle) (n uint32, err error)`. */
	@:go.name("GetFileType")
	@:go.valueArgs("0")
	@:go.tupleReturn
	public static function getFileType(handle:WindowsHandle):WindowsFileTypeResult;

	/** Go (Windows): `func CloseHandle(handle Handle) (err error)`. */
	@:go.name("CloseHandle")
	@:go.valueArgs("0")
	public static function closeHandle(handle:WindowsHandle):Null<go.NativeError>;
}

/** The `(fd Handle, err error)` results of the Windows `syscall.Open`. */
@:goBuildConstraint("windows")
final class WindowsOpenResult {
	public final fd:WindowsHandle;
	public final err:Null<go.NativeError>;

	public function new(fd:WindowsHandle, err:Null<go.NativeError>) {
		this.fd = fd;
		this.err = err;
	}
}

/** The `(n uint32, err error)` results of `syscall.GetFileType`. */
@:goBuildConstraint("windows")
final class WindowsFileTypeResult {
	public final n:NativeUint32;
	public final err:Null<go.NativeError>;

	public function new(n:NativeUint32, err:Null<go.NativeError>) {
		this.n = n;
		this.err = err;
	}
}

/** Passes a `uint32` variable as the `perm` argument and compares a `uint32` result. */
@:goBuildConstraint("windows")
@:keep
final class WindowsSyscallOpen implements SyscallOpenPort {
	static final installed:Bool = SyscallOpen.install(new WindowsSyscallOpen());

	public function new() {}

	public function openReadOnly(path:String, perm:NativeUint32):Bool {
		final opened = WindowsSyscall.open(path, WindowsSyscall.oRdonly, perm);
		if (opened.err != null)
			return false;
		final fileType = WindowsSyscall.getFileType(opened.fd);
		final closed = WindowsSyscall.closeHandle(opened.fd);
		return fileType.err == null && fileType.n == WindowsSyscall.fileTypeDisk && closed == null;
	}
}
