package go;

class Go {
  public static function ok<T>(value:T):Result<T> {
    return Result.ok(value);
  }

  public static function fail<T>(message:String):Result<T> {
    return Result.failure(message);
  }

  public static function newChan<T>():Chan<T> {
    return new Chan<T>();
  }

  public static function newSlice<T>():Slice<T> {
    return new Slice<T>();
  }

  public static function newMap<K, V>():Map<K, V> {
    return new Map<K, V>();
  }
}
