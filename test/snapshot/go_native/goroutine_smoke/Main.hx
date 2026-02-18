import go.Chan;
import go.Go;

class Main {
  static function main() {
    var ch:Chan<Int> = Go.newChan();
    Go.spawn(function() {
      ch.send(5);
    });
    Sys.println(ch.recv());
  }
}
