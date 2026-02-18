import go.Chan;

class Main {
  static function main() {
    var ch = new Chan<Int>();
    ch.send(10);
    ch.send(20);
    Sys.println(ch.recv());
    Sys.println(ch.recv());
    Sys.println(ch.recv());
  }
}
