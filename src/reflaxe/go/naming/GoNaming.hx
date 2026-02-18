package reflaxe.go.naming;

class GoNaming {
  static final goKeywords = [
    "func", "type", "var", "map", "range", "package", "return", "if", "else", "for", "go", "defer", "select",
    "chan", "switch", "fallthrough", "default", "case"
  ];

  public static function normalizeIdent(name:String):String {
    if (name == "this") {
      return "self";
    }

    var sanitized = new StringBuf();
    for (index in 0...name.length) {
      var ch = name.charCodeAt(index);
      var isLower = ch >= "a".code && ch <= "z".code;
      var isUpper = ch >= "A".code && ch <= "Z".code;
      var isDigit = ch >= "0".code && ch <= "9".code;
      if (isLower || isUpper || isDigit || ch == "_".code) {
        sanitized.addChar(ch);
      } else {
        sanitized.add("_");
      }
    }

    var normalized = sanitized.toString();
    if (normalized == "") {
      normalized = "hx_tmp";
    }

    var hasNonUnderscore = false;
    for (index in 0...normalized.length) {
      if (normalized.charCodeAt(index) != "_".code) {
        hasNonUnderscore = true;
        break;
      }
    }
    if (!hasNonUnderscore) {
      normalized = "hx_tmp";
    }

    var first = normalized.charCodeAt(0);
    var startsWithDigit = first >= "0".code && first <= "9".code;
    if (startsWithDigit) {
      normalized = "hx_" + normalized;
    }

    for (keyword in goKeywords) {
      if (normalized == keyword) {
        return normalized + "_";
      }
    }

    return normalized;
  }

  public static function typeSymbol(pack:Array<String>, typeName:String):String {
    var parts = [for (part in pack) normalizeIdent(part)];
    parts.push(normalizeIdent(typeName));
    return parts.join("__");
  }

  public static function constructorSymbol(pack:Array<String>, typeName:String):String {
    return "New_" + typeSymbol(pack, typeName);
  }

  public static function staticSymbol(pack:Array<String>, typeName:String, fieldName:String, preserveMainStatics:Bool):String {
    if (preserveMainStatics) {
      return normalizeIdent(fieldName);
    }
    return typeSymbol(pack, typeName) + "_" + normalizeIdent(fieldName);
  }
}
