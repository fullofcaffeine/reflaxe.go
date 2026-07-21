package reflaxe.go;

import reflaxe.go.compiler.GoCompilerDefine;
#if macro
import haxe.macro.Context;
#end

class ProfileResolver {
	public static inline final DEFINE_NAME:GoCompilerDefine = GoCompilerDefine.DefineProfile;
	public static inline final PORTABLE_DEFINE:GoCompilerDefine = GoCompilerDefine.DefinePortableProfile;
	public static inline final IDIOMATIC_DEFINE:GoCompilerDefine = GoCompilerDefine.DefineIdiomaticProfile;
	public static inline final GOPHER_DEFINE:GoCompilerDefine = GoCompilerDefine.DefineGopherProfile;
	public static inline final METAL_DEFINE:GoCompilerDefine = GoCompilerDefine.DefineMetalProfile;
	public static inline final DEPRECATED_IDIOMATIC = "idiomatic";
	public static inline final REMOVED_GOPHER = "gopher";

	#if macro
	public static function resolve():GoProfile {
		var raw = Context.definedValue(DEFINE_NAME);
		var selected = new Array<{source:String, profile:GoProfile}>();
		var wantsIdiomaticAlias = Context.defined(IDIOMATIC_DEFINE);

		if (raw != null && raw == "") {
			Context.fatalError('`-D ' + DEFINE_NAME + '` requires a value: portable|metal', Context.currentPos());
		}

		if (wantsIdiomaticAlias) {
			Context.fatalError('`-D ' + IDIOMATIC_DEFINE + '` has been removed. Use `-D ' + DEFINE_NAME + '=portable`.', Context.currentPos());
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
			Context.fatalError('`-D ' + GOPHER_DEFINE + '` has been removed. Use `-D ' + DEFINE_NAME + '=portable` or `-D ' + DEFINE_NAME + '=metal`.',
				Context.currentPos());
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
			case REMOVED_GOPHER:
				Context.fatalError('`-D ' + DEFINE_NAME + '=' + raw + '` has been removed. Use `-D ' + DEFINE_NAME + '=portable` or `-D ' + DEFINE_NAME
					+ '=metal`.',
					Context.currentPos());
				GoProfile.Portable;
			case "metal": GoProfile.Metal;
			case DEPRECATED_IDIOMATIC:
				Context.fatalError('`-D ' + DEFINE_NAME + '=' + DEPRECATED_IDIOMATIC + '` has been removed. Use `-D ' + DEFINE_NAME + '=portable`.',
					Context.currentPos());
				GoProfile.Portable;
			case _:
				Context.fatalError('Invalid profile "' + raw + '" for -D ' + DEFINE_NAME + ' (expected portable|metal)', Context.currentPos());
				GoProfile.Portable;
		}
	}
	#else
	public static function resolve():GoProfile {
		return GoProfile.Portable;
	}
	#end
}
