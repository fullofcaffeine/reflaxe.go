package reflaxe.go.analyze;

#if macro
import haxe.macro.Type;
import reflaxe.go.compiler.GoMetadataName;

/**
	Collect scoped raw `__go__` authority declarations.

	Why:
	- Strict mode and strict examples normally reject raw `__go__`.
	- Some low-level abstraction modules still need an explicit escape hatch.

	What:
	- `@:goAllowRaw` marks a module/type as an authority island for raw `__go__`.
	- The marker is resolved to the owning module so a single tagged primary type
	  can authorize the whole module file.

	How:
	- Macro enforcers collect tagged modules after typing.
	- Preflight scanners also honor source files that contain `@:goAllowRaw` so
	  the early text scan does not reject intentionally-authorized modules.
**/
class GoRawInjectionAuthorityAnalyzer {
	public static inline final META_NAME:GoMetadataName = GoMetadataName.GoAllowRaw;
	static final SOURCE_MARKERS = ["@:goAllowRaw", "@goAllowRaw"];

	public static function collect(moduleTypes:Array<ModuleType>):GoRawInjectionAuthoritySnapshot {
		var moduleSet:Map<String, Bool> = [];
		var declarations:Array<GoRawInjectionAuthorityDeclaration> = [];

		for (moduleType in moduleTypes) {
			switch (moduleType) {
				case TClassDecl(classRef):
					var classType = classRef.get();
					var module = moduleNameForClass(classType);
					addTypeDeclarationIfTagged(moduleSet, declarations, module, "class " + classPath(classType), classType.meta, classType.pos);
				case TEnumDecl(enumRef):
					var enumType = enumRef.get();
					var module = moduleNameForEnum(enumType);
					addTypeDeclarationIfTagged(moduleSet, declarations, module, "enum " + enumPath(enumType), enumType.meta, enumType.pos);
				case TTypeDecl(typeRef):
					var typeDecl = typeRef.get();
					var module = moduleNameForTypedef(typeDecl);
					addTypeDeclarationIfTagged(moduleSet, declarations, module, "typedef " + typedefPath(typeDecl), typeDecl.meta, typeDecl.pos);
				case TAbstract(absRef):
					var abstractType = absRef.get();
					var module = moduleNameForAbstract(abstractType);
					addTypeDeclarationIfTagged(moduleSet, declarations, module, "abstract " + abstractPath(abstractType), abstractType.meta, abstractType.pos);
			}
		}

		var modules = [for (module in moduleSet.keys()) module];
		modules.sort(compareStrings);
		declarations.sort(compareDeclarations);
		return {
			modules: modules,
			declarations: declarations
		};
	}

	public static function sourceTextHasRawAuthorityMarker(content:String):Bool {
		if (content == null || content == "") {
			return false;
		}
		for (marker in SOURCE_MARKERS) {
			if (StringTools.contains(content, marker)) {
				return true;
			}
		}
		return false;
	}

	static function addTypeDeclarationIfTagged(moduleSet:Map<String, Bool>, declarations:Array<GoRawInjectionAuthorityDeclaration>, module:String,
			source:String, meta:MetaAccess, pos:haxe.macro.Expr.Position):Void {
		if (meta == null || !metaHasGoAllowRaw(meta)) {
			return;
		}
		var normalized = normalizeModuleLabel(module);
		moduleSet.set(normalized, true);
		declarations.push({
			module: normalized,
			source: source,
			pos: pos
		});
	}

	static function metaHasGoAllowRaw(meta:MetaAccess):Bool {
		for (entry in meta.get()) {
			if (META_NAME.matches(entry.name)) {
				return true;
			}
		}
		return false;
	}

	static inline function normalizeModuleLabel(value:Null<String>):String {
		if (value == null) {
			return "<unknown>";
		}
		var trimmed = StringTools.trim(value);
		return trimmed == "" ? "<unknown>" : trimmed;
	}

	static inline function moduleNameForClass(classType:ClassType):String {
		if (classType.module != null && classType.module.length > 0) {
			return classType.module;
		}
		return pathFromPack(classType.pack, classType.name);
	}

	static inline function moduleNameForAbstract(abstractType:AbstractType):String {
		if (abstractType.module != null && abstractType.module.length > 0) {
			return abstractType.module;
		}
		return pathFromPack(abstractType.pack, abstractType.name);
	}

	static inline function moduleNameForEnum(enumType:EnumType):String {
		if (enumType.module != null && enumType.module.length > 0) {
			return enumType.module;
		}
		return pathFromPack(enumType.pack, enumType.name);
	}

	static inline function moduleNameForTypedef(typedefType:DefType):String {
		if (typedefType.module != null && typedefType.module.length > 0) {
			return typedefType.module;
		}
		return pathFromPack(typedefType.pack, typedefType.name);
	}

	static inline function classPath(classType:ClassType):String {
		return pathFromPack(classType.pack, classType.name);
	}

	static inline function abstractPath(abstractType:AbstractType):String {
		return pathFromPack(abstractType.pack, abstractType.name);
	}

	static inline function enumPath(enumType:EnumType):String {
		return pathFromPack(enumType.pack, enumType.name);
	}

	static inline function typedefPath(typedefType:DefType):String {
		return pathFromPack(typedefType.pack, typedefType.name);
	}

	static inline function pathFromPack(pack:Array<String>, name:String):String {
		return pack == null || pack.length == 0 ? name : pack.join(".") + "." + name;
	}

	static inline function compareStrings(a:String, b:String):Int {
		return a < b ? -1 : (a > b ? 1 : 0);
	}

	static function compareDeclarations(a:GoRawInjectionAuthorityDeclaration, b:GoRawInjectionAuthorityDeclaration):Int {
		var moduleCmp = compareStrings(a.module, b.module);
		if (moduleCmp != 0) {
			return moduleCmp;
		}
		return compareStrings(a.source, b.source);
	}
}

typedef GoRawInjectionAuthorityDeclaration = {
	var module:String;
	var source:String;
	var pos:haxe.macro.Expr.Position;
}

typedef GoRawInjectionAuthoritySnapshot = {
	var modules:Array<String>;
	var declarations:Array<GoRawInjectionAuthorityDeclaration>;
}
#end
