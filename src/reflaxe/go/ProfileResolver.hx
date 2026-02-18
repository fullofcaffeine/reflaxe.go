package reflaxe.go;

#if macro
import haxe.macro.Context;
#end

class ProfileResolver {
  public static inline final DEFINE_NAME = "reflaxe_go_profile";

  #if macro
  public static function resolve():GoProfile {
    var raw = Context.definedValue(DEFINE_NAME);
    if (raw == null || raw == "") {
      return GoProfile.Portable;
    }

    return switch (raw) {
      case "portable": GoProfile.Portable;
      case "idiomatic": GoProfile.Idiomatic;
      case "gopher": GoProfile.Gopher;
      case "metal": GoProfile.Metal;
      case _: Context.fatalError('Invalid profile "$raw" for -D $DEFINE_NAME', Context.currentPos());
    }
  }
  #else
  public static function resolve():GoProfile {
    return GoProfile.Portable;
  }
  #end
}
