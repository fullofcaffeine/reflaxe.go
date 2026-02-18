import go.Go;
import go.Result;

class Main {
  static function okValue():Result<Int> {
    return Go.ok(12);
  }

  static function errValue():Result<Int> {
    return Go.fail("bad");
  }

  static function main() {
    var ok = okValue();
    Sys.println(ok.isOk());
    Sys.println(ok.unwrap());

    var err = errValue();
    Sys.println(err.isErr());
    Sys.println(err.error());
  }
}
