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
			case "Std":
				requireSourceOwnedStdlibClass("Std");
			case "Reflect":
				requireSourceOwnedStdlibClass("Reflect");
			case "EReg":
				requireSourceOwnedStdlibClass("EReg");
				requireSourceOwnedStdlibClass("StringBuf");
			case "haxe.Serializer":
				requireSourceOwnedStdlibClass("haxe.Serializer");
				requireSourceOwnedStdlibClass("StringBuf");
				requireSourceOwnedStdlibModule("haxe.ds.List");
				requireSourceOwnedStdlibClass("haxe.ds.StringMap");
				requireSourceOwnedStdlibClass("haxe.ds.IntMap");
				requireSourceOwnedStdlibClass("haxe.ds.ObjectMap");
			case "haxe.Unserializer":
				requireSourceOwnedStdlibClass("haxe.Unserializer");
				requireSourceOwnedStdlibModule("haxe.ds.List");
				requireSourceOwnedStdlibClass("haxe.ds.StringMap");
				requireSourceOwnedStdlibClass("haxe.ds.IntMap");
				requireSourceOwnedStdlibClass("haxe.ds.ObjectMap");
			case "Sys":
				requireSourceOwnedStdlibClass("Sys");
			case "StringTools":
				requireSourceOwnedStdlibClass("StringTools");
				requireSourceOwnedStdlibClass("haxe.iterators.StringIterator");
				requireSourceOwnedStdlibClass("haxe.iterators.StringKeyValueIterator");
			case "Date":
				requireSourceOwnedStdlibClass("Date");
			case "Math":
				requireSourceOwnedStdlibClass("Math");
			case "DateTools":
				requireSourceOwnedStdlibClass("DateTools");
				requireSourceOwnedStdlibClass("Date");
				requireSourceOwnedStdlibClass("Math");
			case "haxe.io.Path":
				requireSourceOwnedStdlibClass("haxe.io.Path");
			case "haxe.io.Bytes":
				requireBytesSourceClasses();
			case "haxe.io.BytesBuffer":
				requireSourceOwnedStdlibClass("haxe.io.BytesBuffer");
				requireBytesSourceClasses();
			case "haxe.io.Input":
				requireSourceOwnedStdlibClass("haxe.io.Input");
				requireBytesSourceClasses();
				requireSourceOwnedStdlibClass("haxe.io.BytesBuffer");
				requireSourceOwnedStdlibClass("haxe.io.Eof");
				requireSourceOwnedStdlibClass("haxe.exceptions.NotImplementedException");
			case "haxe.io.Output":
				requireSourceOwnedStdlibClass("haxe.io.Output");
				requireBytesSourceClasses();
				requireSourceOwnedStdlibClass("haxe.exceptions.NotImplementedException");
			case "haxe.io.BytesInput":
				requireSourceOwnedStdlibClass("haxe.io.BytesInput");
				requireSourceOwnedStdlibClass("haxe.io.Input");
				requireSourceOwnedStdlibClass("haxe.io.Bytes");
				requireSourceOwnedStdlibClass("haxe.io.Eof");
				requireSourceOwnedStdlibEnum("haxe.io.Error");
			case "haxe.io.BytesOutput":
				requireSourceOwnedStdlibClass("haxe.io.BytesOutput");
				requireSourceOwnedStdlibClass("haxe.io.Output");
				requireSourceOwnedStdlibClass("haxe.io.BytesBuffer");
			case "haxe.io.BufferInput":
				requireSourceOwnedStdlibClass("haxe.io.BufferInput");
				requireSourceOwnedStdlibClass("haxe.io.Input");
				requireSourceOwnedStdlibClass("haxe.io.Bytes");
			case "haxe.io.StringInput":
				requireSourceOwnedStdlibClass("haxe.io.StringInput");
				requireSourceOwnedStdlibClass("haxe.io.BytesInput");
				requireSourceOwnedStdlibClass("haxe.io.Bytes");
			case "haxe.io.Eof":
				requireSourceOwnedStdlibClass("haxe.io.Eof");
			case "haxe.io.Encoding":
				requireSourceOwnedStdlibEnum("haxe.io.Encoding");
			case "haxe.io.Error":
				requireSourceOwnedStdlibEnum("haxe.io.Error");
			case "haxe.io.FPHelper":
				requireSourceOwnedStdlibClass("haxe.io.FPHelper");
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
			case "UnicodeString", "_UnicodeString.UnicodeString_Impl_":
				requireSourceOwnedStdlibModule("UnicodeString");
				requireSourceOwnedStdlibClass("haxe.iterators.StringIteratorUnicode");
				requireSourceOwnedStdlibClass("haxe.iterators.StringKeyValueIteratorUnicode");
			case "haxe.crypto.Base64":
				requireSourceOwnedStdlibClass("haxe.crypto.Base64");
			case "haxe.crypto.Md5", "haxe.crypto.Sha1", "haxe.crypto.Sha224", "haxe.crypto.Sha256":
				requireSourceOwnedStdlibClass(fullClassName(classType));
			case "Xml", "_Xml.XmlType_Impl_":
				requireSourceOwnedStdlibModule("Xml");
			case "haxe.xml.Parser", "haxe.xml.XmlParserException":
				requireSourceOwnedStdlibModule("haxe.xml.Parser");
				requireSourceOwnedStdlibModule("Xml");
				requireSourceOwnedStdlibClass("StringBuf");
			case "haxe.xml.Printer":
				requireSourceOwnedStdlibClass("haxe.xml.Printer");
				requireSourceOwnedStdlibModule("Xml");
				requireSourceOwnedStdlibClass("StringBuf");
			case "haxe.zip.Compress", "haxe.zip.Uncompress":
				requireSourceOwnedStdlibClass(fullClassName(classType));
				requireSourceOwnedStdlibEnum("haxe.zip.FlushMode");
			case "haxe.zip.Tools":
				requireSourceOwnedStdlibClass("haxe.zip.Tools");
				requireSourceOwnedStdlibClass("haxe.zip.Compress");
				requireSourceOwnedStdlibClass("haxe.zip.Uncompress");
				requireSourceOwnedStdlibEnum("haxe.zip.FlushMode");
			case "haxe.exceptions.PosException":
				requireSourceOwnedStdlibClass("haxe.exceptions.PosException");
			case "haxe.exceptions.ArgumentException":
				requireSourceOwnedStdlibClass("haxe.exceptions.PosException");
				requireSourceOwnedStdlibClass("haxe.exceptions.ArgumentException");
			case "haxe.exceptions.NotImplementedException":
				requireSourceOwnedStdlibClass("haxe.exceptions.PosException");
				requireSourceOwnedStdlibClass("haxe.exceptions.NotImplementedException");
			case "sys.Http":
				requireSourceOwnedStdlibClass("sys.Http");
				requireSourceOwnedStdlibClass("haxe.http.HttpBase");
				requireSourceOwnedStdlibClass("haxe.ds.StringMap");
				requireSourceOwnedStdlibClass("haxe.io.Bytes");
				requireSourceOwnedStdlibClass("haxe.io.Input");
				requireSourceOwnedStdlibClass("haxe.io.Output");
				requireSourceOwnedStdlibModule("sys.net.Socket");
				requireSourceOwnedStdlibModule("sys.net._SocketIO");
				requireSourceOwnedStdlibClass("sys.net.Host");
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
				requireBaseIoSourceClasses();
			case "sys.net.Address":
				requireSourceOwnedStdlibClass("sys.net.Address");
			case "sys.net.Host":
				requireSourceOwnedStdlibClass("sys.net.Host");
			case "sys.net.Socket":
				requireSourceOwnedStdlibModule("sys.net.Socket");
				requireSourceOwnedStdlibModule("sys.net._SocketIO");
				requireSourceOwnedStdlibClass("sys.net.Host");
				requireBaseIoSourceClasses();
			case "sys.net.UdpSocket":
				requireSourceOwnedStdlibClass("sys.net.UdpSocket");
				requireSourceOwnedStdlibModule("sys.net.Socket");
				requireSourceOwnedStdlibModule("sys.net._SocketIO");
				requireSourceOwnedStdlibClass("sys.net.Host");
				requireSourceOwnedStdlibClass("sys.net.Address");
				requireBaseIoSourceClasses();
			case "sys.ssl.Certificate":
				requireSourceOwnedStdlibClass("sys.ssl.Certificate");
			case "sys.ssl.Digest":
				requireSourceOwnedStdlibClass("sys.ssl.Digest");
				requireSourceOwnedStdlibModule("sys.ssl.DigestAlgorithm");
				requireSourceOwnedStdlibClass("sys.ssl.Key");
			case "sys.ssl.DigestAlgorithm":
				requireSourceOwnedStdlibModule("sys.ssl.DigestAlgorithm");
			case "sys.ssl.Key":
				requireSourceOwnedStdlibClass("sys.ssl.Key");
			case "sys.ssl.Socket", "sys.ssl._Socket.Socket_Impl_":
				requireSourceOwnedStdlibModule("sys.ssl.Socket");
				requireSourceOwnedStdlibModule("sys.net.Socket");
				requireSourceOwnedStdlibModule("sys.net._SocketIO");
				requireSourceOwnedStdlibClass("sys.net.Host");
				requireSourceOwnedStdlibClass("sys.ssl.Certificate");
				requireSourceOwnedStdlibClass("sys.ssl.Key");
				requireBaseIoSourceClasses();
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
			case "haxe.ValueException":
				// Direct ValueException usage lowers to the existing hxrt exception
				// carrier rather than emitting a separate source-owned class body.
			case _:
		}
	}

	public function requireBaseIoSourceClasses():Void {
		requireSourceOwnedStdlibClass("haxe.io.Input");
		requireSourceOwnedStdlibClass("haxe.io.Output");
		requireBytesSourceClasses();
		requireSourceOwnedStdlibClass("haxe.io.Eof");
		requireSourceOwnedStdlibEnum("haxe.io.Error");
	}

	/**
		What: Enqueue the complete typed source dependency closure for staged
		`haxe.io.Bytes`.

		Why: A compiler-owned consumer can request `Bytes` by name instead of
		encountering a `new Bytes(...)` expression. The public API still mentions
		the concrete `haxe.Int64` carrier, `FPHelper`, encodings, errors, and
		`StringBuf`; omitting those
		declarations can leave otherwise valid selective-runtime output with dangling
		Go type or function references.

		How: Queue only ordinary staged modules and enums. The runtime boundary remains
		the separately inferred typed `hxrt` byte capability.
	**/
	public function requireBytesSourceClasses():Void {
		requireSourceOwnedStdlibClass("haxe.io.Bytes");
		requireSourceOwnedStdlibClass("haxe.io.FPHelper");
		requireSourceOwnedStdlibClass("StringBuf");
		requireSourceOwnedStdlibEnum("haxe.io.Encoding");
		requireSourceOwnedStdlibEnum("haxe.io.Error");
	}

	public function requireSourceOwnedStdlibClass(className:String):Void {
		if (className == "sys.io.FileInput" || className == "sys.io.FileOutput") {
			// Source-owned subclasses inherit the staged Input/Output hierarchy.
			requireBaseIoSourceClasses();
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

	/**
		What: Promote an already-typed, source-backed stdlib class into the ordinary
		source inclusion queue.

		Why: A class can survive only as a nominal type reference after manual DCE.
		Resolving it again by module can restore methods DCE intentionally removed,
		while ignoring it leaves generated Go type names without declarations.

		How: Accept the exact typed `ClassType` carried by the expression/signature,
		verify that it comes from std or vendor source, preserve compiler-owned
		authorities, then cache and enqueue only that declaration.
	**/
	public function requireTypedSourceOwnedStdlibClass(classType:ClassType):Bool {
		var className = fullClassName(classType);
		if (classType.isExtern || isCompilerOwnedAuthority(className)) {
			return false;
		}
		var location = PositionTools.toLocation(classType.pos);
		if (location == null || location.file == null) {
			return false;
		}
		var file = Std.string(location.file);
		if (file == null || (!StringTools.contains(file, "/std/") && !StringTools.contains(file, "/vendor/"))) {
			return false;
		}

		availableClassesByName.set(className, classType);
		requireSourceOwnedStdlibClass(className);
		return requiredSourceOwnedClassNames.exists(className);
	}

	/**
		What: Promote a typed, source-backed stdlib superclass into the normal source
		inclusion queue.

		Why: Manual dead-code elimination can initially select only the user subclass.
		The superclass is still present in the typed inheritance link, but it may not
		yet be in the planner's availability map. Treating that class as absent drops
		the embedded base view and makes otherwise valid Haxe upcasts invalid Go.

		How: Accept only non-extern classes whose declaration comes from staged std or
		vendor source, preserve compiler-owned authorities, cache the already-typed
		class, and reuse the ordinary requirement queue.
	**/
	public function requireSourceOwnedStdlibSuperclass(classType:ClassType):Bool {
		return requireTypedSourceOwnedStdlibClass(classType);
	}

	/**
		What
		Resolves an optional staged class through Haxe's typed macro API.

		Why
		`Context.getType(...)` documents a `String` exception for an unresolved
		name. Catching `Dynamic` here unnecessarily erased that known boundary.

		How
		Accept only a typed class result and translate the documented missing-name
		exception to `null`; unrelated compiler failures are not swallowed.
	**/
	function resolveSourceOwnedStdlibClass(className:String):Null<ClassType> {
		try {
			return switch (Context.getType(className)) {
				case TInst(classRef, _):
					classRef.get();
				case _:
					null;
			};
		} catch (_:String) {
			return null;
		}
	}

	/** Resolves an optional staged module using the same typed `String` failure boundary. */
	function resolveSourceOwnedStdlibModule(moduleName:String):Array<Type> {
		try {
			return Context.getModule(moduleName);
		} catch (_:String) {
			return [];
		}
	}

	/** Resolves an optional staged enum using the same typed `String` failure boundary. */
	function resolveSourceOwnedStdlibEnum(enumName:String):Null<EnumType> {
		try {
			return switch (Context.getType(enumName)) {
				case TEnum(enumRef, _):
					enumRef.get();
				case _:
					null;
			};
		} catch (_:String) {
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
