package go;

@:noCompletion
@:dox(hide)
abstract __ChanHandle<T>(Dynamic) from Dynamic to Dynamic {}

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
	public static function __chanMake<T>(buffer:Int):__ChanHandle<T> {
		return null;
	}

	@:noCompletion
	public static function __chanSend<T>(channel:__ChanHandle<T>, value:T):Void {}

	@:noCompletion
	public static function __chanRecv<T>(channel:__ChanHandle<T>):Null<T> {
		return null;
	}

	@:noCompletion
	public static function __chanTrySend<T>(channel:__ChanHandle<T>, value:T):Bool {
		return false;
	}

	@:noCompletion
	public static function __chanRecvOr<T>(channel:__ChanHandle<T>, defaultValue:T):T {
		return defaultValue;
	}

	@:noCompletion
	public static function __chanTryRecv<T>(channel:__ChanHandle<T>):Result<T> {
		return Result.failure("empty");
	}

	@:noCompletion
	public static function __chanClose<T>(channel:__ChanHandle<T>):Void {}
}
