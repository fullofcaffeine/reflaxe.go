package go;

class Go {
	public static function ok<T>(value:T):Result<T> {
		return Result.ok(value);
	}

	public static function fail<T>(message:String):Result<T> {
		return Result.failure(message);
	}

	public static function newChan<T>(buffer:Int = 0):Chan<T> {
		var channel = new Chan<T>();
		if (buffer > 0) {
			channel.__hx_setBuffer(buffer);
		}
		return channel;
	}

	public static function newSlice<T>():Slice<T> {
		return new Slice<T>();
	}

	public static function newMap<K, V>():Map<K, V> {
		return new Map<K, V>();
	}

	public static function spawn(fn:Void->Void):Void {
		__goSpawn(fn);
	}

	@:noCompletion
	public static function __goSpawn(fn:Void->Void):Void {}

	@:noCompletion
	public static function __chanMake(buffer:Int):Dynamic {
		return null;
	}

	@:noCompletion
	public static function __chanSend(channel:Dynamic, value:Dynamic):Void {}

	@:noCompletion
	public static function __chanRecv(channel:Dynamic):Dynamic {
		return null;
	}

	@:noCompletion
	public static function __chanClose(channel:Dynamic):Void {}
}
