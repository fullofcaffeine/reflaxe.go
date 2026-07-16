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
			case "Sys":
				requireSourceOwnedStdlibClass("Sys");
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
			case "haxe.io.ArrayBufferViewImpl", "haxe.io._ArrayBufferView.ArrayBufferView_Impl_":
				requireSourceOwnedStdlibModule("haxe.io.ArrayBufferView");
			case "haxe.io._UInt8Array.UInt8Array_Impl_":
				requireSourceOwnedStdlibModule("haxe.io.UInt8Array");
			case "haxe.io._UInt16Array.UInt16Array_Impl_":
				requireSourceOwnedStdlibModule("haxe.io.UInt16Array");
			case "haxe.io._UInt32Array.UInt32Array_Impl_":
				requireSourceOwnedStdlibModule("haxe.io.UInt32Array");
			case "haxe.io._Int32Array.Int32Array_Impl_":
				requireSourceOwnedStdlibModule("haxe.io.Int32Array");
			case "haxe.io._Float32Array.Float32Array_Impl_":
				requireSourceOwnedStdlibModule("haxe.io.Float32Array");
			case "haxe.io._Float64Array.Float64Array_Impl_":
				requireSourceOwnedStdlibModule("haxe.io.Float64Array");
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
				requireSourceOwnedStdlibClass("haxe.EntryPoint");
				requireSourceOwnedStdlibModule("sys.thread.EventLoop");
				requireSourceOwnedStdlibClass("sys.thread.Thread");
				requireSourceOwnedStdlibClass("sys.thread.Mutex");
				requireSourceOwnedStdlibClass("sys.thread.NoEventLoopException");
			case "haxe.MainLoop", "haxe.MainEvent":
				requireSourceOwnedStdlibModule("haxe.MainLoop");
				requireSourceOwnedStdlibClass("haxe.EntryPoint");
				requireSourceOwnedStdlibClass("haxe.Timer");
				requireSourceOwnedStdlibModule("sys.thread.EventLoop");
				requireSourceOwnedStdlibClass("sys.thread.Thread");
				requireSourceOwnedStdlibClass("sys.thread.Mutex");
				requireSourceOwnedStdlibClass("sys.thread.NoEventLoopException");
			case "haxe.Timer":
				requireSourceOwnedStdlibClass("haxe.Timer");
				requireSourceOwnedStdlibModule("sys.thread.EventLoop");
				requireSourceOwnedStdlibClass("sys.thread.Thread");
				requireSourceOwnedStdlibClass("sys.thread.NoEventLoopException");
			case "haxe.Log", "haxe.Resource", "haxe.SysTools":
				requireSourceOwnedStdlibClass(fullClassName(classType));
			case "sys.db.Connection":
				requireSourceOwnedStdlibClass("sys.db.Connection");
				requireSourceOwnedStdlibClass("sys.db.ResultSet");
				requireSourceOwnedStdlibClass("StringBuf");
			case "sys.db.ResultSet":
				requireSourceOwnedStdlibClass("sys.db.ResultSet");
			case "sys.db.Mysql":
				requireSourceOwnedStdlibClass("sys.db.Mysql");
				requireSourceOwnedStdlibClass("sys.db.Connection");
				requireSourceOwnedStdlibClass("sys.db.ResultSet");
				requireSourceOwnedStdlibClass("StringBuf");
			case "sys.db.Sqlite":
				requireSourceOwnedStdlibClass("sys.db.Sqlite");
				requireSourceOwnedStdlibClass("sys.db.Connection");
				requireSourceOwnedStdlibClass("sys.db.ResultSet");
				requireSourceOwnedStdlibClass("StringBuf");
			case "haxe.ds.ArraySort":
				requireSourceOwnedStdlibClass("haxe.ds.ArraySort");
			case "haxe.ds.BalancedTree":
				requireSourceOwnedStdlibModule("haxe.ds.BalancedTree");
			case "haxe.ds.GenericStack":
				requireSourceOwnedStdlibModule("haxe.ds.GenericStack");
			case "haxe.ds.IntMap", "haxe.ds.ObjectMap", "haxe.ds.StringMap":
				requireSourceOwnedStdlibClass(fullClassName(classType));
				requireSourceOwnedStdlibModule("haxe.Constraints");
				requireSourceOwnedStdlibClass("haxe.iterators.MapKeyValueIterator");
			case "haxe.ds.EnumValueMap":
				requireSourceOwnedStdlibModule("haxe.ds.EnumValueMap");
				requireSourceOwnedStdlibModule("haxe.Constraints");
			case "haxe.ds.List":
				requireSourceOwnedStdlibModule("haxe.ds.List");
			case "haxe.iterators.MapKeyValueIterator":
				requireSourceOwnedStdlibClass("haxe.iterators.MapKeyValueIterator");
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
			case "sys.FileSystem":
				requireSourceOwnedStdlibClass("sys.FileSystem");
			case "sys.io.File":
				requireSourceOwnedStdlibModule("sys.io.File");
				requireSourceOwnedStdlibClass("sys.io.FileInput");
				requireSourceOwnedStdlibClass("sys.io.FileOutput");
				requireSourceOwnedStdlibEnum("sys.io.FileSeek");
			case "sys.io.FileInput":
				requireSourceOwnedStdlibClass("sys.io.FileInput");
				requireSourceOwnedStdlibEnum("sys.io.FileSeek");
			case "sys.io.FileOutput":
				requireSourceOwnedStdlibClass("sys.io.FileOutput");
				requireSourceOwnedStdlibEnum("sys.io.FileSeek");
			case "sys.io.Process":
				requireSourceOwnedStdlibModule("sys.io.Process");
				requireIoSourceOwnedHelperClass();
			case "sys.net.Address":
				requireSourceOwnedStdlibClass("sys.net.Address");
			case "sys.ssl.Certificate":
				requireSourceOwnedStdlibClass("sys.ssl.Certificate");
				requireStdlibShimGroup("stdlib_symbols");
			case "sys.ssl.Digest":
				requireSourceOwnedStdlibClass("sys.ssl.Digest");
				requireSourceOwnedStdlibModule("sys.ssl.DigestAlgorithm");
				requireSourceOwnedStdlibClass("sys.ssl.Key");
				requireStdlibShimGroup("stdlib_symbols");
			case "sys.ssl.DigestAlgorithm":
				requireSourceOwnedStdlibModule("sys.ssl.DigestAlgorithm");
			case "sys.ssl.Key":
				requireSourceOwnedStdlibClass("sys.ssl.Key");
				requireStdlibShimGroup("stdlib_symbols");
			case "sys.ssl.Socket", "sys.ssl._Socket.Socket_Impl_":
				requireSourceOwnedStdlibModule("sys.ssl.Socket");
				requireSourceOwnedStdlibClass("sys.ssl.Certificate");
				requireSourceOwnedStdlibClass("sys.ssl.Key");
				requireStdlibShimGroup("net_socket");
				requireStdlibShimGroup("stdlib_symbols");
			case "sys.thread.Lock":
				requireSourceOwnedStdlibClass("sys.thread.Lock");
			case "sys.thread.Mutex":
				requireSourceOwnedStdlibClass("sys.thread.Mutex");
			case "sys.thread.Condition":
				requireSourceOwnedStdlibClass("sys.thread.Condition");
			case "sys.thread.Semaphore":
				requireSourceOwnedStdlibClass("sys.thread.Semaphore");
			case "sys.thread.Deque":
				requireSourceOwnedStdlibClass("sys.thread.Deque");
			case "sys.thread.Tls":
				requireSourceOwnedStdlibClass("sys.thread.Tls");
			case "sys.thread.NoEventLoopException":
				requireSourceOwnedStdlibClass("sys.thread.NoEventLoopException");
			case "sys.thread.ThreadPoolException":
				requireSourceOwnedStdlibClass("sys.thread.ThreadPoolException");
			case "sys.thread.IThreadPool":
				requireSourceOwnedStdlibClass("sys.thread.IThreadPool");
			case "sys.thread.EventLoop":
				requireSourceOwnedStdlibModule("sys.thread.EventLoop");
			case "sys.thread.Thread":
				requireSourceOwnedStdlibClass("sys.thread.Thread");
				requireSourceOwnedStdlibModule("sys.thread.EventLoop");
				requireSourceOwnedStdlibClass("sys.thread.NoEventLoopException");
			case "sys.thread.ElasticThreadPool":
				requireSourceOwnedStdlibClass("sys.thread.ElasticThreadPool");
				requireSourceOwnedStdlibClass("sys.thread.ElasticThreadPoolWorker");
				requireSourceOwnedStdlibClass("sys.thread.Thread");
				requireSourceOwnedStdlibModule("sys.thread.EventLoop");
				requireSourceOwnedStdlibClass("sys.thread.ThreadPoolException");
			case "sys.thread.FixedThreadPool":
				requireSourceOwnedStdlibClass("sys.thread.FixedThreadPool");
				requireSourceOwnedStdlibClass("sys.thread.FixedThreadPoolWorker");
				requireSourceOwnedStdlibClass("sys.thread.FixedThreadPoolShutdownException");
				requireSourceOwnedStdlibClass("sys.thread.Thread");
				requireSourceOwnedStdlibClass("sys.thread.ThreadPoolException");
			case "haxe.Template":
				requireSourceOwnedStdlibModule("haxe.Template");
				requireStdlibShimGroup("stdlib_symbols");
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
		if (className == "sys.io.FileInput" || className == "sys.io.FileOutput") {
			// Source-owned subclasses still inherit the compiler-owned Input/Output
			// method surface, whose ordinary Haxe algorithms live in GoIoHelpers.
			requireIoSourceOwnedHelperClass();
		}
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
