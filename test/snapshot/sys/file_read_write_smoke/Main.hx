class Main {
  static function main() {
    var path = "./tmp_sys_file_smoke.txt";
    sys.io.File.saveContent(path, "hello");
    var content = sys.io.File.getContent(path);
    Sys.println(content);
  }
}
