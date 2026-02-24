package go;

class Chan<T> {
	var __hx_native:Dynamic;

	public function new() {
		__hx_native = Go.__chanMake(0);
	}

	@:noCompletion
	public function __hx_setBuffer(buffer:Int):Void {
		__hx_native = Go.__chanMake(buffer);
	}

	public function send(value:T):Void {
		Go.__chanSend(__hx_native, value);
	}

	public function recv():Null<T> {
		return cast Go.__chanRecv(__hx_native);
	}

	public function trySend(value:T):Bool {
		return Go.__chanTrySend(__hx_native, value);
	}

	public function tryRecv():Result<T> {
		return cast Go.__chanTryRecv(__hx_native);
	}

	public function recvOr(defaultValue:T):T {
		return cast Go.__chanRecvOr(__hx_native, defaultValue);
	}

	public function close():Void {
		Go.__chanClose(__hx_native);
	}
}
