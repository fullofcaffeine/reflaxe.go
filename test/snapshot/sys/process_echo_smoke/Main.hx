class Main {
  static function main() {
    var p = new sys.io.Process("echo", ["hi"]);
    var line = p.stdout.readLine();
    Sys.println(line);
    p.close();
  }
}
