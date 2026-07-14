package go;

/**
	Go-native typed channel facade.

	Why
	Applications that intentionally own Go semantics need real channel blocking,
	buffering, close, and native-panic behavior without raw `__go__` injection.

	What
	`Chan<T>` exposes blocking and non-blocking send/receive operations. `tryRecv`
	distinguishes a temporarily empty channel (`"empty"`) from a drained closed
	channel (`"closed"`). Buffered values remain receivable after close.

	How
	Generated Go uses native typed channels and comma-ok receives. Sending on a
	closed channel, closing a nil/closed channel, and racing send with close retain
	Go's native panic contract; those panics are not Haxe catch values. This facade
	is an explicit Go-native API, not the planned cross-target channel product.
**/
class Chan<T> {
	#if (go || go_output)
	var __hx_native:Go.__ChanHandle<T>;
	#else
	var queue:Array<T>;
	var readIndex:Int;
	var capacity:Int;
	var closed:Bool;
	#end

	/** Create an unbuffered channel. Prefer `Go.newChan(buffer)` for a buffer. */
	public function new() {
		#if (go || go_output)
		__hx_native = Go.__chanMake(0);
		#else
		queue = [];
		readIndex = 0;
		capacity = 0;
		closed = false;
		#end
	}

	@:noCompletion
	public function __hx_setBuffer(buffer:Int):Void {
		#if (go || go_output)
		__hx_native = Go.__chanMake(buffer);
		#else
		capacity = buffer <= 0 ? 0 : buffer;
		#end
	}

	/**
		Block until `value` is accepted. A nil channel blocks forever; a closed
		channel raises Go's native `send on closed channel` panic.
	**/
	public function send(value:T):Void {
		#if (go || go_output)
		Go.__chanSend(__hx_native, value);
		#else
		if (closed) {
			throw "send on closed channel";
		}
		trySend(value);
		#end
	}

	/**
		Block for the next value. Once a closed channel is drained, generated Go
		returns `T`'s Go zero value (represented as `null` when appropriate).
	**/
	public function recv():Null<T> {
		#if (go || go_output)
		return Go.__chanRecv(__hx_native);
		#else
		if (readIndex >= queue.length) {
			return null;
		}

		var value = queue[readIndex];
		readIndex++;
		return value;
		#end
	}

	/**
		Attempt an immediate send. Returns `false` only when sending would block,
		including a nil channel. Sending on a closed channel is still a native panic.
	**/
	public function trySend(value:T):Bool {
		#if (go || go_output)
		return Go.__chanTrySend(__hx_native, value);
		#else
		if (closed) {
			throw "send on closed channel";
		}
		if (capacity > 0 && unreadCount() >= capacity) {
			return false;
		}
		queue.push(value);
		return true;
		#end
	}

	/**
		Attempt an immediate receive. Returns success for a queued value, `"empty"`
		when no value is ready, and `"closed"` once a closed channel is drained.
	**/
	public function tryRecv():Result<T> {
		#if (go || go_output)
		return Go.__chanTryRecv(__hx_native);
		#else
		if (readIndex >= queue.length) {
			if (closed) {
				return Result.failure("closed");
			}
			return Result.failure("empty");
		}
		var value = queue[readIndex];
		readIndex++;
		return Result.ok(value);
		#end
	}

	/**
		Return an immediately available value or `defaultValue`. A drained closed
		channel and a nil/temporarily empty channel both select the default.
	**/
	public function recvOr(defaultValue:T):T {
		#if (go || go_output)
		return Go.__chanRecvOr(__hx_native, defaultValue);
		#else
		var value = recv();
		return value == null ? defaultValue : cast value;
		#end
	}

	/**
		Close the producer side. Only the producer that owns lifecycle should call
		this: closing a nil/already-closed channel is a native Go panic, and send/close
		races must be synchronized by the application.
	**/
	public function close():Void {
		#if (go || go_output)
		Go.__chanClose(__hx_native);
		#else
		if (closed) {
			throw "close of closed channel";
		}
		closed = true;
		#end
	}

	#if !(go || go_output)
	inline function unreadCount():Int {
		return queue.length - readIndex;
	}
	#end
}
