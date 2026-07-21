package reflaxe.go.analyze;

#if macro
import haxe.macro.Context;
import haxe.macro.Type;
import reflaxe.go.compiler.GoMetadataName;

/**
	Compatibility parser for native-boundary metadata.

	New code should enter through `GoNativeBoundaryAnalyzer`. This module keeps
	the historical name so external macro imports and `@:goMetal` source remain
	compatible while `@:goNative` is the canonical spelling.
**/
class MetalLaneAnalyzer {
	public static function collect(moduleTypes:Array<ModuleType>):MetalLaneSnapshot {
		var moduleSet:Map<String, Bool> = [];
		var declarations:Array<MetalLaneDeclaration> = [];

		for (moduleType in moduleTypes) {
			switch (moduleType) {
				case TClassDecl(classRef):
					var classType = classRef.get();
					var module = moduleNameForClass(classType);
					addTypeDeclarationIfTagged(moduleSet, declarations, module, "class " + classPath(classType), classType.meta, classType.pos);
					collectFieldDeclarations(moduleSet, declarations, module, classType.name, classType.fields.get());
					collectFieldDeclarations(moduleSet, declarations, module, classType.name, classType.statics.get(), true);
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
					if (abstractType.impl != null) {
						var impl = abstractType.impl.get();
						if (impl != null) {
							collectFieldDeclarations(moduleSet, declarations, module, abstractType.name, impl.fields.get());
							collectFieldDeclarations(moduleSet, declarations, module, abstractType.name, impl.statics.get(), true);
						}
					}
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

	static function collectFieldDeclarations(moduleSet:Map<String, Bool>, declarations:Array<MetalLaneDeclaration>, module:String, owner:String,
			fields:Array<ClassField>, isStatic:Bool = false):Void {
		if (fields == null) {
			return;
		}
		var prefix = isStatic ? "static field " : "field ";
		for (field in fields) {
			if (field == null || field.meta == null) {
				continue;
			}
			if (!metaHasNativeBoundary(field.meta)) {
				continue;
			}
			addDeclaration(moduleSet, declarations, module, prefix + owner + "." + field.name, field.pos);
		}
	}

	static function addTypeDeclarationIfTagged(moduleSet:Map<String, Bool>, declarations:Array<MetalLaneDeclaration>, module:String, source:String,
			meta:MetaAccess, pos:haxe.macro.Expr.Position):Void {
		if (meta == null || !metaHasNativeBoundary(meta)) {
			return;
		}
		addDeclaration(moduleSet, declarations, module, source, pos);
	}

	static function addDeclaration(moduleSet:Map<String, Bool>, declarations:Array<MetalLaneDeclaration>, module:String, source:String,
			pos:haxe.macro.Expr.Position):Void {
		var normalized = normalizeModuleLabel(module);
		moduleSet.set(normalized, true);
		declarations.push({
			module: normalized,
			source: source,
			pos: pos
		});
	}

	static function metaHasNativeBoundary(meta:MetaAccess):Bool {
		for (entry in meta.get()) {
			if (GoMetadataName.GoNative.matches(entry.name)) {
				return true;
			}
			if (GoMetadataName.GoMetal.matches(entry.name)) {
				return true;
			}
			if (GoMetadataName.RemovedHaxeMetal.matches(entry.name)) {
				Context.error("Native boundary metadata @:haxeMetal was removed; use @:goNative (@:goMetal remains a compatibility alias).", entry.pos);
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

	static function compareDeclarations(a:MetalLaneDeclaration, b:MetalLaneDeclaration):Int {
		var moduleCmp = compareStrings(a.module, b.module);
		if (moduleCmp != 0) {
			return moduleCmp;
		}
		return compareStrings(a.source, b.source);
	}
}

typedef MetalLaneDeclaration = {
	var module:String;
	var source:String;
	var pos:haxe.macro.Expr.Position;
}

typedef MetalLaneSnapshot = {
	var modules:Array<String>;
	var declarations:Array<MetalLaneDeclaration>;
}
#end
