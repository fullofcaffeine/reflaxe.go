package go;

class Chan<T> {
	#if (go || go_output)
	var __hx_native:Go.__ChanHandle<T>;
	#else
	var queue:Array<T>;
	var readIndex:Int;
	var capacity:Int;
	var closed:Bool;
	#end

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

	public function send(value:T):Void {
		#if (go || go_output)
		Go.__chanSend(__hx_native, value);
		#else
		trySend(value);
		#end
	}

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

	public function trySend(value:T):Bool {
		#if (go || go_output)
		return Go.__chanTrySend(__hx_native, value);
		#else
		if (closed) {
			return false;
		}
		if (capacity > 0 && unreadCount() >= capacity) {
			return false;
		}
		queue.push(value);
		return true;
		#end
	}

	public function tryRecv():Result<T> {
		#if (go || go_output)
		return Go.__chanTryRecv(__hx_native);
		#else
		if (readIndex >= queue.length) {
			return Result.failure("empty");
		}
		var value = queue[readIndex];
		readIndex++;
		return Result.ok(value);
		#end
	}

	public function recvOr(defaultValue:T):T {
		#if (go || go_output)
		return Go.__chanRecvOr(__hx_native, defaultValue);
		#else
		var value = recv();
		return value == null ? defaultValue : cast value;
		#end
	}

	public function close():Void {
		#if (go || go_output)
		Go.__chanClose(__hx_native);
		#else
		closed = true;
		#end
	}

	#if !(go || go_output)
	inline function unreadCount():Int {
		return queue.length - readIndex;
	}
	#end
}
