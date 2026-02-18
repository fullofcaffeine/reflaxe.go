import go.Result;

class Main {
  static function main() {
    var ok = Result.ok(7);
    Sys.println(ok.isOk());
    Sys.println(ok.unwrap());

    var err:Result<Int> = Result.failure("boom");
    Sys.println(err.isErr());
    Sys.println(err.error());
  }
}
