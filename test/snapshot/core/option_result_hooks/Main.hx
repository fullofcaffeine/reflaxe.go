enum Maybe<T> {
  None;
  Some(value:T);
}

enum Res<T, E> {
  Ok(value:T);
  Err(error:E);
}

class Main {
  static function unwrapOr(value:Maybe<Int>, fallback:Int):Int {
    return switch (value) {
      case Some(v): v;
      case None: fallback;
    }
  }

  static function render(res:Res<Int, String>):String {
    return switch (res) {
      case Ok(v): Std.string(v);
      case Err(e): e;
    }
  }

  static function main() {
    Sys.println(unwrapOr(Some(7), 0));
    Sys.println(unwrapOr(None, 5));
    Sys.println(render(Ok(9)));
    Sys.println(render(Err("bad")));
  }
}
