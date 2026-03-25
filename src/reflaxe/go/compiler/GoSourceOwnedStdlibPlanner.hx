package reflaxe.go.compiler;

#if macro
import haxe.macro.Context;
import haxe.macro.PositionTools;
import haxe.macro.Type;

private typedef GoSourceOwnedStdlibPlannerConfig = {
	final availableClassesByName:Map<String, ClassType>;
	final pendingRequiredClassesByName:Map<String, ClassType>;
	final availableEnumsByName:Map<String, EnumType>;
	final pendingRequiredEnumsByName:Map<String, EnumType>;
	final requiredSourceOwnedClassNames:Map<String, Bool>;
	final isCompilerOwnedAuthority:String->Bool;
	final fullClassName:ClassType->String;
	final fullEnumName:EnumType->String;
	final requireStdlibShimGroup:String->Void;
	final markIoSourceOwnedHelperSurfaceRequired:Void->Void;
}

/**
	Planner for source-owned stdlib inclusion and direct-usage routing.

	Why:
	- Portable parity work increasingly relies on staged std modules and direct
	  source-owned std inclusion.
	- Keeping that routing logic inside `GoCompiler` makes ownership decisions hard
	  to review and easy to regress.

	What:
	- Tracks source-owned class/enum availability and pending inclusion queues.
	- Decides when direct stdlib usage should require staged std modules, helper
	  classes, or explicit fatal blockers.

	How:
	- Operates over the compiler's typed class/enum maps through a typed config
	  object.
	- Reuses the existing pending requirement maps so queue behavior remains
	  unchanged while the ownership logic moves out of the monolith.
**/
class GoSourceOwnedStdlibPlanner {
	public final availableClassesByName:Map<String, ClassType>;
	public final pendingRequiredClassesByName:Map<String, ClassType>;
	public final availableEnumsByName:Map<String, EnumType>;
	public final pendingRequiredEnumsByName:Map<String, EnumType>;
	public final requiredSourceOwnedClassNames:Map<String, Bool>;

	final isCompilerOwnedAuthority:String->Bool;
	final fullClassName:ClassType->String;
	final fullEnumName:EnumType->String;
	final requireStdlibShimGroup:String->Void;
	final markIoSourceOwnedHelperSurfaceRequired:Void->Void;

	public function new(config:GoSourceOwnedStdlibPlannerConfig) {
		availableClassesByName = config.availableClassesByName;
		pendingRequiredClassesByName = config.pendingRequiredClassesByName;
		availableEnumsByName = config.availableEnumsByName;
		pendingRequiredEnumsByName = config.pendingRequiredEnumsByName;
		requiredSourceOwnedClassNames = config.requiredSourceOwnedClassNames;
		isCompilerOwnedAuthority = config.isCompilerOwnedAuthority;
		fullClassName = config.fullClassName;
		fullEnumName = config.fullEnumName;
		requireStdlibShimGroup = config.requireStdlibShimGroup;
		markIoSourceOwnedHelperSurfaceRequired = config.markIoSourceOwnedHelperSurfaceRequired;
	}

	public function cacheAvailableClasses(classes:Array<ClassType>):Void {
		clearClassMap(availableClassesByName);
		for (classType in classes) {
			availableClassesByName.set(fullClassName(classType), classType);
		}
	}

	public function cacheAvailableEnums(enums:Array<EnumType>):Void {
		clearEnumMap(availableEnumsByName);
		for (enumType in enums) {
			availableEnumsByName.set(fullEnumName(enumType), enumType);
		}
	}

	public function noteSourceOwnedStdlibUsage(classType:ClassType):Void {
		switch (fullClassName(classType)) {
			case "StringTools":
				requireSourceOwnedStdlibClass("StringTools");
				requireSourceOwnedStdlibClass("haxe.iterators.StringIterator");
				requireSourceOwnedStdlibClass("haxe.iterators.StringKeyValueIterator");
			case "DateTools":
				requireSourceOwnedStdlibClass("DateTools");
			case "haxe.io.Path":
				requireSourceOwnedStdlibClass("haxe.io.Path");
			case "haxe.io.FPHelper":
				requireSourceOwnedStdlibClass("haxe.io.FPHelper");
				requireStdlibShimGroup("stdlib_symbols");
			case "haxe._CallStack.CallStack_Impl_":
				requireSourceOwnedStdlibModule("haxe.CallStack");
				requireSourceOwnedStdlibEnum("haxe.StackItem");
			case "haxe.NativeStackTrace":
				requireSourceOwnedStdlibClass("haxe.NativeStackTrace");
			case "haxe.http.HttpBase":
				requireSourceOwnedStdlibClass("haxe.http.HttpBase");
			case "haxe.rtti.Meta":
				requireSourceOwnedStdlibClass("haxe.rtti.Meta");
			case "haxe.rtti.Rtti":
				requireSourceOwnedStdlibClass("haxe.rtti.Rtti");
			case "haxe.rtti.XmlParser":
				requireSourceOwnedStdlibClass("haxe.rtti.XmlParser");
				requireSourceOwnedStdlibModule("haxe.rtti.CType");
			case "haxe.rtti.TypeApi", "haxe.rtti.CTypeTools":
				requireSourceOwnedStdlibModule("haxe.rtti.CType");
			case "haxe.EntryPoint":
				failDirectEventLoopSurface("haxe.EntryPoint", classType.pos);
			case "haxe.MainLoop":
				failDirectEventLoopSurface("haxe.MainLoop", classType.pos);
			case "haxe.MainEvent":
				failDirectEventLoopSurface("haxe.MainEvent", classType.pos);
			case "haxe.Timer":
				failDirectEventLoopSurface("haxe.Timer", classType.pos);
			case "haxe.Log", "haxe.Resource", "haxe.SysTools":
				requireSourceOwnedStdlibClass(fullClassName(classType));
			case "haxe.ds.ArraySort":
				requireSourceOwnedStdlibClass("haxe.ds.ArraySort");
			case "haxe.ds.BalancedTree":
				requireSourceOwnedStdlibModule("haxe.ds.BalancedTree");
			case "haxe.ds.GenericStack":
				requireSourceOwnedStdlibModule("haxe.ds.GenericStack");
			case "haxe.ds.WeakMap":
				requireSourceOwnedStdlibClass("haxe.ds.WeakMap");
				requireSourceOwnedStdlibClass("haxe.exceptions.PosException");
				requireSourceOwnedStdlibClass("haxe.exceptions.NotImplementedException");
			case "haxe.ds.ListSort":
				requireSourceOwnedStdlibClass("haxe.ds.ListSort");
			case "haxe.Utf8":
				requireSourceOwnedStdlibClass("haxe.Utf8");
			case "haxe.exceptions.PosException":
				requireSourceOwnedStdlibClass("haxe.exceptions.PosException");
			case "haxe.exceptions.ArgumentException":
				requireSourceOwnedStdlibClass("haxe.exceptions.PosException");
				requireSourceOwnedStdlibClass("haxe.exceptions.ArgumentException");
			case "haxe.exceptions.NotImplementedException":
				requireSourceOwnedStdlibClass("haxe.exceptions.PosException");
				requireSourceOwnedStdlibClass("haxe.exceptions.NotImplementedException");
			case "sys.Http":
				requireSourceOwnedStdlibClass("sys.GoHttpHelpers");
			case "haxe.Template":
				requireSourceOwnedStdlibModule("haxe.Template");
				requireStdlibShimGroup("stdlib_symbols");
				requireStdlibShimGroup("template_support");
			case "haxe.ValueException":
				// Direct ValueException usage lowers to the existing hxrt exception
				// carrier rather than emitting a separate source-owned class body.
			case _:
		}
	}

	public function requireIoSourceOwnedHelperClass():Void {
		markIoSourceOwnedHelperSurfaceRequired();
		requireSourceOwnedStdlibClass("haxe.io.GoIoHelpers");
	}

	public function requireSourceOwnedStdlibClass(className:String):Void {
		if (isCompilerOwnedAuthority(className)) {
			return;
		}
		if (!availableClassesByName.exists(className)) {
			var resolved = resolveSourceOwnedStdlibClass(className);
			if (resolved != null) {
				availableClassesByName.set(className, resolved);
			}
		}
		if (!availableClassesByName.exists(className) || pendingRequiredClassesByName.exists(className)) {
			return;
		}
		requiredSourceOwnedClassNames.set(className, true);
		pendingRequiredClassesByName.set(className, availableClassesByName.get(className));
	}

	public function requireSourceOwnedStdlibModule(moduleName:String):Void {
		var resolved = resolveSourceOwnedStdlibModule(moduleName);
		for (moduleType in resolved) {
			if (skipCompilerOwnedSourceModuleType(moduleType)) {
				continue;
			}
			switch (moduleType) {
				case TInst(classRef, _):
					var classType = classRef.get();
					var className = fullClassName(classType);
					availableClassesByName.set(className, classType);
					requiredSourceOwnedClassNames.set(className, true);
					if (!pendingRequiredClassesByName.exists(className)) {
						pendingRequiredClassesByName.set(className, classType);
					}
				case TEnum(enumRef, _):
					var enumType = enumRef.get();
					var enumName = fullEnumName(enumType);
					availableEnumsByName.set(enumName, enumType);
					if (!pendingRequiredEnumsByName.exists(enumName)) {
						pendingRequiredEnumsByName.set(enumName, enumType);
					}
				case _:
			}
		}
	}

	public function requireSourceOwnedStdlibEnum(enumName:String):Void {
		if (!availableEnumsByName.exists(enumName)) {
			var resolved = resolveSourceOwnedStdlibEnum(enumName);
			if (resolved != null) {
				availableEnumsByName.set(enumName, resolved);
			}
		}
		if (!availableEnumsByName.exists(enumName) || pendingRequiredEnumsByName.exists(enumName)) {
			return;
		}
		pendingRequiredEnumsByName.set(enumName, availableEnumsByName.get(enumName));
	}

	public function hasLoadedSourceOwnedStdlibClass(className:String):Bool {
		if (!availableClassesByName.exists(className)) {
			return false;
		}
		var classType = availableClassesByName.get(className);
		if (classType == null || classType.isExtern) {
			return false;
		}
		var location = PositionTools.toLocation(classType.pos);
		if (location == null || location.file == null) {
			return false;
		}
		var file = Std.string(location.file);
		return file != null && (StringTools.contains(file, "/std/") || StringTools.contains(file, "/vendor/"));
	}

	function resolveSourceOwnedStdlibClass(className:String):Null<ClassType> {
		try {
			return switch (Context.getType(className)) {
				case TInst(classRef, _):
					classRef.get();
				case _:
					null;
			};
		} catch (_:Dynamic) {
			return null;
		}
	}

	function resolveSourceOwnedStdlibModule(moduleName:String):Array<Type> {
		try {
			return Context.getModule(moduleName);
		} catch (_:Dynamic) {
			return [];
		}
	}

	function resolveSourceOwnedStdlibEnum(enumName:String):Null<EnumType> {
		try {
			return switch (Context.getType(enumName)) {
				case TEnum(enumRef, _):
					enumRef.get();
				case _:
					null;
			};
		} catch (_:Dynamic) {
			return null;
		}
	}

	function skipCompilerOwnedSourceModuleType(moduleType:Type):Bool {
		return switch (moduleType) {
			case TInst(classRef, _):
				var classType = classRef.get();
				isCompilerOwnedAuthority(fullClassName(classType));
			case _:
				false;
		};
	}

	function failDirectEventLoopSurface(surfaceName:String, pos:haxe.macro.Expr.Position):Void {
		Context.fatalError("Direct "
			+ surfaceName
			+ " usage is explicitly unsupported on haxe.go. "
			+ "A real runtime-backed event-loop contract (sys.thread.EventLoop / sys.thread.Thread) "
			+ "does not exist yet, and the previous source-owned inclusion path generated broken Go. "
			+ "Future runtime work is tracked under haxe.go-14as.19.",
			pos);
	}

	static function clearClassMap(map:Map<String, ClassType>):Void {
		var keys = [for (key in map.keys()) key];
		for (key in keys) {
			map.remove(key);
		}
	}

	static function clearEnumMap(map:Map<String, EnumType>):Void {
		var keys = [for (key in map.keys()) key];
		for (key in keys) {
			map.remove(key);
		}
	}
}
#end
