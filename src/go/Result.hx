package go;

class Result<T> {
  public var value(default, null):Null<T>;
  public var errorValue(default, null):Null<Error>;

  public function new(value:Null<T>, errorValue:Null<Error>) {
    this.value = value;
    this.errorValue = errorValue;
  }

  public static function ok<T>(value:T):Result<T> {
    return new Result(value, null);
  }

  public static function failure<T>(message:String):Result<T> {
    return new Result(null, new Error(message));
  }

  public function isOk():Bool {
    return errorValue == null;
  }

  public function isErr():Bool {
    return errorValue != null;
  }

  public function unwrap():T {
    if (errorValue != null) {
      throw errorValue.toString();
    }
    return cast value;
  }

  public function error():Null<String> {
    if (errorValue == null) {
      return null;
    }
    return errorValue.toString();
  }
}
