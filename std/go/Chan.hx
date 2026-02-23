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

	public function close():Void {
		Go.__chanClose(__hx_native);
	}
}
