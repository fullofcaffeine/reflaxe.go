import go.NativeUint32;

/** One read-only `syscall.Open` call with an exact Go `uint32` permission argument. */
interface SyscallOpenPort {
	function openReadOnly(path:String, perm:NativeUint32):Bool;
}

/**
	Holds the port that a build-constrained module installs for its GOOS.

	`syscall.Open` returns an `int` descriptor on Unix and a `syscall.Handle` on
	Windows, so each platform declares it in its own module (see UnixSyscallOpen
	and WindowsSyscallOpen). The common module names no platform type.
**/
final class SyscallOpen {
	static var port:Null<SyscallOpenPort> = null;

	public static function install(value:SyscallOpenPort):Bool {
		port = value;
		return true;
	}

	/** True when the open succeeded, or when this GOOS installs no port. */
	public static function openReadOnly(path:String, perm:NativeUint32):Bool {
		final selected = port;
		return selected == null || selected.openReadOnly(path, perm);
	}
}
