package go;

class Chan<T> {
	var queue:Array<T>;
	var readIndex:Int;
	var capacity:Int;
	var closed:Bool;

	public function new() {
		queue = [];
		readIndex = 0;
		capacity = 0;
		closed = false;
	}

	@:noCompletion
	public function __hx_setBuffer(buffer:Int):Void {
		capacity = buffer <= 0 ? 0 : buffer;
	}

	public function send(value:T):Void {
		trySend(value);
	}

	public function recv():Null<T> {
		if (readIndex >= queue.length) {
			return null;
		}

		var value = queue[readIndex];
		readIndex++;
		return value;
	}

	public function trySend(value:T):Bool {
		if (closed) {
			return false;
		}
		if (capacity > 0 && unreadCount() >= capacity) {
			return false;
		}
		queue.push(value);
		return true;
	}

	public function tryRecv():Result<T> {
		if (readIndex >= queue.length) {
			return Result.failure("empty");
		}
		var value = queue[readIndex];
		readIndex++;
		return Result.ok(value);
	}

	public function recvOr(defaultValue:T):T {
		var value = recv();
		return value == null ? defaultValue : cast value;
	}

	public function close():Void {
		closed = true;
	}

	inline function unreadCount():Int {
		return queue.length - readIndex;
	}
}
