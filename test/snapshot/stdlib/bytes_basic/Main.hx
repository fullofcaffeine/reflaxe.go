class Main {
  static function main() {
    var bytes = haxe.io.Bytes.ofString("abc");
    bytes.set(1, 122);
    Sys.println(bytes.toString());
    Sys.println(bytes.length);

    var buffer = new haxe.io.BytesBuffer();
    buffer.addString("Hi");
    buffer.addByte(33);
    var out = buffer.getBytes();
    Sys.println(out.toString());
  }
}
