package reflaxe.go.macros;

#if macro
import haxe.macro.Compiler as MacroCompiler;
import haxe.macro.Context;
import haxe.macro.Expr;
#end

class AutoEmptyConstructor {
	#if macro
	public static inline final DEFINE_INTERFACES = "reflaxe_go_auto_empty_ctor_interfaces";
	static var initialized = false;

	public static function init():Void {
		if (initialized) {
			return;
		}
		initialized = true;

		var raw = Context.definedValue(DEFINE_INTERFACES);
		if (raw == null) {
			return;
		}

		var interfaces = parseInterfaces(raw);
		if (interfaces.length == 0) {
			Context.fatalError('`-D ' + DEFINE_INTERFACES + '` requires one or more fully-qualified interface names.', Context.currentPos());
		}

		for (interfacePath in interfaces) {
			MacroCompiler.addGlobalMetadata(interfacePath, "@:autoBuild(reflaxe.go.macros.AutoEmptyConstructor.build())", false, true, false);
		}
	}

	public static function build():Array<Field> {
		var fields = Context.getBuildFields();
		for (field in fields) {
			if (field.name == "new") {
				return fields;
			}
		}

		fields.push({
			name: "new",
			access: [APublic],
			kind: FFun({
				args: [],
				ret: null,
				expr: macro {}
			}),
			pos: Context.currentPos()
		});

		return fields;
	}

	static function parseInterfaces(raw:String):Array<String> {
		var entries = new Array<String>();
		var seen = new Map<String, Bool>();
		var pathPattern = ~/^[A-Za-z_][A-Za-z0-9_]*(\.[A-Za-z_][A-Za-z0-9_]*)+$/;
		for (entry in raw.split(",")) {
			var trimmed = StringTools.trim(entry);
			if (trimmed == "") {
				continue;
			}
			if (!pathPattern.match(trimmed)) {
				Context.fatalError('Invalid interface path "' + trimmed + '" in -D ' + DEFINE_INTERFACES + '.', Context.currentPos());
			}
			if (seen.exists(trimmed)) {
				continue;
			}
			seen.set(trimmed, true);
			entries.push(trimmed);
		}
		return entries;
	}
	#else
	public static function init():Void {}
	#end
}
