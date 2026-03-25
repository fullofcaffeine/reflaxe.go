package hxrt.thread;

/**
	What
	Typed bridge to the small `hxrt` thread helper surface used by staged
	`sys.thread` overrides on `haxe.go`.

	Why
	The portable API should stay in ordinary Haxe code. The runtime bridge keeps
	the blocking and thread-local primitives in Go without forcing the staged
	stdlib to fall back to raw `__go__` snippets.

	How
	Each static extern maps directly to one exported `hxrt` function.
**/
@:go.import("hxrt")
@:go.package("hxrt")
extern class NativeThread {
	@:go.name("ThreadCurrentId")
	public static function currentId():Int;

	@:go.name("ThreadLockNew")
	public static function lockNew():LockHandle;
	@:go.name("ThreadLockWait")
	public static function lockWait(handle:LockHandle):Bool;
	@:go.name("ThreadLockWaitTimeout")
	public static function lockWaitTimeout(handle:LockHandle, timeout:Float):Bool;

	/**
		What
		Optional-timeout bridge for staged std wrappers whose omitted primitive
		arguments lower as `any` at the Haxe/Go boundary.

		Why
		The compiler must preserve omitted optional primitive arguments honestly.
		This keeps the `Dynamic` bridge localized to the runtime boundary instead of
		leaking it into the public `sys.thread` surface.

		How
		The Go runtime asserts the timeout payload back to `float64` before
		delegating to the typed timeout helper.
	**/
	@:go.name("ThreadLockWaitTimeoutAny")
	public static function lockWaitTimeoutDynamic(handle:LockHandle, timeout:Dynamic):Bool;

	@:go.name("ThreadLockRelease")
	public static function lockRelease(handle:LockHandle):Void;

	@:go.name("ThreadMutexNew")
	public static function mutexNew():MutexHandle;
	@:go.name("ThreadMutexAcquire")
	public static function mutexAcquire(handle:MutexHandle):Void;
	@:go.name("ThreadMutexTryAcquire")
	public static function mutexTryAcquire(handle:MutexHandle):Bool;
	@:go.name("ThreadMutexRelease")
	public static function mutexRelease(handle:MutexHandle):Void;

	@:go.name("ThreadConditionNew")
	public static function conditionNew():ConditionHandle;
	@:go.name("ThreadConditionAcquire")
	public static function conditionAcquire(handle:ConditionHandle):Void;
	@:go.name("ThreadConditionTryAcquire")
	public static function conditionTryAcquire(handle:ConditionHandle):Bool;
	@:go.name("ThreadConditionRelease")
	public static function conditionRelease(handle:ConditionHandle):Void;
	@:go.name("ThreadConditionWait")
	public static function conditionWait(handle:ConditionHandle):Void;
	@:go.name("ThreadConditionSignal")
	public static function conditionSignal(handle:ConditionHandle):Void;
	@:go.name("ThreadConditionBroadcast")
	public static function conditionBroadcast(handle:ConditionHandle):Void;

	@:go.name("ThreadSemaphoreNew")
	public static function semaphoreNew(value:Int):SemaphoreHandle;
	@:go.name("ThreadSemaphoreAcquire")
	public static function semaphoreAcquire(handle:SemaphoreHandle):Void;
	@:go.name("ThreadSemaphoreTryAcquire")
	public static function semaphoreTryAcquire(handle:SemaphoreHandle):Bool;
	@:go.name("ThreadSemaphoreTryAcquireTimeout")
	public static function semaphoreTryAcquireTimeout(handle:SemaphoreHandle, timeout:Float):Bool;
	@:go.name("ThreadSemaphoreTryAcquireTimeoutAny")
	public static function semaphoreTryAcquireTimeoutDynamic(handle:SemaphoreHandle, timeout:Dynamic):Bool;
	@:go.name("ThreadSemaphoreRelease")
	public static function semaphoreRelease(handle:SemaphoreHandle):Void;
}
