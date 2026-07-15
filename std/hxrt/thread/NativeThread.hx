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
	@:go.name("ThreadNowSeconds")
	public static function nowSeconds():Float;

	@:go.name("ThreadCurrentId")
	public static function currentId():Int;

	/**
		What
		Allocates one opaque runtime slot for a staged `sys.thread.Tls<T>` instance.

		Why
		Slot identity must not be forgeable or reused while an older value could
		still be observable.

		How
		The runtime returns a process-unique `ThreadLocalHandle`; values themselves
		remain owned by the current logical thread state.
	**/
	@:go.name("ThreadLocalNew")
	public static function localNew():ThreadLocalHandle;

	/**
		What
		Reads the current logical thread's value for one TLS handle.

		Why
		One `sys.thread.Tls<T>` handle can carry any Haxe `T`, so this is an
		unavoidable representation boundary.

		How
		`Dynamic` is localized here; the staged generic wrapper restores the
		declared source type immediately after the call.
	**/
	@:go.name("ThreadLocalGet")
	public static function localGet(handle:ThreadLocalHandle):Dynamic;

	/**
		What
		Writes or clears the current logical thread's value for one TLS handle.

		Why
		The value must share the owning thread state's cleanup boundary instead of
		living in an unbounded staged-Haxe map.

		How
		The runtime stores non-null values in `ThreadState` and removes the slot when
		the staged wrapper passes `null`; `Dynamic` remains confined to this bridge.
	**/
	@:go.name("ThreadLocalSet")
	public static function localSet(handle:ThreadLocalHandle, value:Dynamic):Void;

	@:go.name("ThreadSpawn")
	public static function spawn(job:() -> Void):Int;
	@:go.name("ThreadSpawnWithEventLoop")
	public static function spawnWithEventLoop(job:() -> Void):Int;
	@:go.name("ThreadHasEventLoop")
	public static function hasEventLoop(threadId:Int):Bool;
	@:go.name("ThreadEvents")
	public static function events(threadId:Int):EventLoopHandle;
	@:go.name("ThreadRunWithEventLoop")
	public static function runWithEventLoop(job:() -> Void):Void;

	/**
		Haxe thread messages may contain any Haxe value, so `Dynamic` is intentionally
		localized to the runtime queue boundary and preserved unchanged in both
		directions.
	**/
	@:go.name("ThreadSendMessage")
	public static function sendMessage(threadId:Int, message:Dynamic):Void;

	@:go.name("ThreadReadMessage")
	public static function readMessage(block:Bool):Dynamic;

	@:go.name("ThreadEventLoopNew")
	public static function eventLoopNew():EventLoopHandle;
	@:go.name("ThreadEventLoopPromise")
	public static function eventLoopPromise(handle:EventLoopHandle):Void;
	@:go.name("ThreadEventLoopRun")
	public static function eventLoopRun(handle:EventLoopHandle, event:() -> Void):Void;
	@:go.name("ThreadEventLoopRunPromised")
	public static function eventLoopRunPromised(handle:EventLoopHandle, event:() -> Void):Void;
	@:go.name("ThreadEventLoopRepeat")
	public static function eventLoopRepeat(handle:EventLoopHandle, event:() -> Void, intervalMs:Int):Int;
	@:go.name("ThreadEventLoopCancel")
	public static function eventLoopCancel(handle:EventLoopHandle, eventId:Int):Void;
	@:go.name("ThreadEventLoopProgress")
	public static function eventLoopProgress(handle:EventLoopHandle):EventLoopProgress;
	@:go.name("ThreadEventLoopWait")
	public static function eventLoopWait(handle:EventLoopHandle):Bool;
	@:go.name("ThreadEventLoopWaitTimeout")
	public static function eventLoopWaitTimeout(handle:EventLoopHandle, timeout:Float):Bool;
	@:go.name("ThreadEventLoopLoop")
	public static function eventLoopLoop(handle:EventLoopHandle):Void;

	@:go.name("ThreadLockNew")
	public static function lockNew():LockHandle;
	@:go.name("ThreadLockWait")
	public static function lockWait(handle:LockHandle):Bool;
	@:go.name("ThreadLockWaitTimeout")
	public static function lockWaitTimeout(handle:LockHandle, timeout:Float):Bool;

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
	@:go.name("ThreadSemaphoreRelease")
	public static function semaphoreRelease(handle:SemaphoreHandle):Void;
}
