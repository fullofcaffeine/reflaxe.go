import go.Map;
import go.Slice;

class Main {
  static function main() {
    var s = new Slice<Int>();
    s.push(1);
    s.push(2);
    s.push(3);
    s.set(1, 7);
    Sys.println(s.length);
    Sys.println(s.get(1));

    var m = new Map<Int, String>();
    m.set(42, "answer");
    Sys.println(m.exists(42));
    Sys.println(m.get(42));
  }
}
