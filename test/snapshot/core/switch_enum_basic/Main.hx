enum Flag {
  Off;
  On;
}

class Main {
  static function toInt(flag:Flag):Int {
    return switch (flag) {
      case Off: 0;
      case On: 1;
    }
  }

  static function main() {
    var current = On;
    switch (current) {
      case Off:
        Sys.println(0);
      case On:
        Sys.println(1);
    }
    Sys.println(toInt(Off));
    Sys.println(toInt(On));
  }
}
