import profile.RuntimeFactory;

class Main {
  static function main() {
    var runtime = RuntimeFactory.create();
    #if example_ci
    Sys.println(Harness.assertContract(runtime));
    #else
    Sys.println(Harness.run(runtime));
    #end
  }
}
