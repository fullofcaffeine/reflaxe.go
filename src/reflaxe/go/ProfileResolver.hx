package reflaxe.go;

#if macro
import haxe.macro.Context;
#end

class ProfileResolver {
  public static inline final DEFINE_NAME = "reflaxe_go_profile";
  public static inline final PORTABLE_DEFINE = "reflaxe_go_portable";
  public static inline final IDIOMATIC_DEFINE = "reflaxe_go_idiomatic";
  public static inline final GOPHER_DEFINE = "reflaxe_go_gopher";
  public static inline final METAL_DEFINE = "reflaxe_go_metal";
  public static inline final DEPRECATED_IDIOMATIC = "idiomatic";

  #if macro
  public static function resolve():GoProfile {
    var raw = Context.definedValue(DEFINE_NAME);
    var selected = new Array<{source:String, profile:GoProfile}>();
    var wantsIdiomaticAlias = Context.defined(IDIOMATIC_DEFINE);

    if (raw != null && raw == "") {
      Context.fatalError('`-D ' + DEFINE_NAME + '` requires a value: portable|gopher|metal', Context.currentPos());
    }

    if (wantsIdiomaticAlias) {
      Context.fatalError('`-D ' + IDIOMATIC_DEFINE + '` has been removed. Use `-D ' + DEFINE_NAME + '=gopher`.', Context.currentPos());
    }

    if (raw != null && raw != "") {
      selected.push({
        source: "-D " + DEFINE_NAME + "=" + raw,
        profile: parseProfile(raw)
      });
    }

    if (Context.defined(PORTABLE_DEFINE)) {
      selected.push({source: "-D " + PORTABLE_DEFINE, profile: GoProfile.Portable});
    }
    if (Context.defined(GOPHER_DEFINE)) {
      selected.push({source: "-D " + GOPHER_DEFINE, profile: GoProfile.Gopher});
    }
    if (Context.defined(METAL_DEFINE)) {
      selected.push({source: "-D " + METAL_DEFINE, profile: GoProfile.Metal});
    }

    if (selected.length == 0) {
      return GoProfile.Portable;
    }

    var winner = selected[0];
    for (index in 1...selected.length) {
      var current = selected[index];
      if (current.profile != winner.profile) {
        var sources = [for (entry in selected) entry.source].join(", ");
        Context.fatalError('Conflicting profile defines: ' + sources, Context.currentPos());
      }
    }

    return winner.profile;
  }

  static function parseProfile(raw:String):GoProfile {
    return switch (raw) {
      case "portable": GoProfile.Portable;
      case "gopher": GoProfile.Gopher;
      case "metal": GoProfile.Metal;
      case DEPRECATED_IDIOMATIC:
        Context.fatalError('`-D ' + DEFINE_NAME + '=' + DEPRECATED_IDIOMATIC + '` has been removed. Use `-D ' + DEFINE_NAME + '=gopher`.', Context.currentPos());
        GoProfile.Portable;
      case _:
        Context.fatalError('Invalid profile "' + raw + '" for -D ' + DEFINE_NAME + ' (expected portable|gopher|metal)', Context.currentPos());
        GoProfile.Portable;
    }
  }
  #else
  public static function resolve():GoProfile {
    return GoProfile.Portable;
  }
  #end
}
