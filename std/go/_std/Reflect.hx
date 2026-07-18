import go.NativeSlice;
import haxe.Constraints.Function;
import hxrt.reflect.NativeReflect;
import hxrt.reflect.ReflectFieldLookup;
import reflaxe.go.internal.CompilerReflect;

/**
	What:
	- Implements the complete Haxe 4.3.7 `Reflect` API as staged source for
	  `haxe.go`.

	Why:
	- The mainstream Haxe stdlib implementation cannot be used unchanged on
	  `haxe.go` because it is an extern contract whose dynamic object operations
	  must be supplied by each target.
	- Lookup order, property access, copying, deletion, and var-args adaptation are
	  public Haxe policy. They do not belong in a monolithic compiler-generated Go
	  shim merely because two lookup steps consume compiler-known metadata.
	- `Dynamic` is intentional at this boundary because `Reflect` is Haxe's
	  explicitly untyped API. It is kept out of ordinary compiler and runtime
	  ownership everywhere else.

	How:
	- Ask the same-package `CompilerReflect` adapter only for closed-world facts:
	  class-token RTTI, generated lowercase fields and methods, and enum carriers.
	- Delegate erased Go map/struct/function inspection to the typed
	  `NativeReflect` runtime boundary.
	- Preserve lookup precedence explicitly as class-token metadata, native map or
	  exported struct field, generated Haxe field, generated Haxe method, then
	  exported native Go method.
**/
@:coreApi
class Reflect {
	public static function hasField(o:Dynamic, field:String):Bool {
		if (o == null)
			return false;
		if (CompilerReflect.hasTypeField(o, field))
			return true;
		if (NativeReflect.lookupField(o, field).found)
			return true;
		if (CompilerReflect.hasGeneratedField(o, field))
			return true;
		if (CompilerReflect.generatedMethod(o, field) != null)
			return true;
		return NativeReflect.lookupMethod(o, field).found;
	}

	public static function field(o:Dynamic, field:String):Dynamic {
		if (o == null)
			return null;

		var typeValue = CompilerReflect.typeField(o, field);
		if (typeValue != null || CompilerReflect.hasTypeField(o, field))
			return typeValue;

		var nativeField:ReflectFieldLookup = NativeReflect.lookupField(o, field);
		if (nativeField.found)
			return nativeField.value;

		var generatedField = CompilerReflect.generatedField(o, field);
		if (generatedField != null || CompilerReflect.hasGeneratedField(o, field))
			return generatedField;

		var generatedMethod = CompilerReflect.generatedMethod(o, field);
		if (generatedMethod != null)
			return generatedMethod;

		var nativeMethod = NativeReflect.lookupMethod(o, field);
		return nativeMethod.found ? nativeMethod.value : null;
	}

	public static function setField(o:Dynamic, field:String, value:Dynamic):Void {
		if (o == null)
			throw "Null Access";
		if (NativeReflect.setField(o, field, value))
			return;
		CompilerReflect.setGeneratedField(o, field, value);
	}

	public static function getProperty(o:Dynamic, field:String):Dynamic {
		var getter = Reflect.field(o, "get_" + field);
		return getter == null ? Reflect.field(o, field) : Reflect.callMethod(o, cast getter, []);
	}

	public static function setProperty(o:Dynamic, field:String, value:Dynamic):Void {
		var setter = Reflect.field(o, "set_" + field);
		if (setter == null) {
			Reflect.setField(o, field, value);
			return;
		}
		Reflect.callMethod(o, cast setter, [value]);
	}

	public static function callMethod(o:Dynamic, func:Function, args:Array<Dynamic>):Dynamic {
		return NativeReflect.callMethod(func, NativeSlice.fromArray(args));
	}

	public static function fields(o:Dynamic):Array<String> {
		var generatedFields = CompilerReflect.generatedFields(o);
		if (generatedFields != null)
			return generatedFields;
		return NativeReflect.fields(o).toArray();
	}

	public static inline function isFunction(f:Dynamic):Bool {
		return NativeReflect.isFunction(f);
	}

	public static function compare<T>(a:T, b:T):Int {
		return NativeReflect.compare(a, b);
	}

	public static inline function compareMethods(f1:Dynamic, f2:Dynamic):Bool {
		return NativeReflect.compareMethods(f1, f2);
	}

	public static inline function isObject(v:Dynamic):Bool {
		return NativeReflect.isObject(v);
	}

	public static inline function isEnumValue(v:Dynamic):Bool {
		return CompilerReflect.isEnumValue(v);
	}

	public static inline function deleteField(o:Dynamic, field:String):Bool {
		return NativeReflect.deleteField(o, field);
	}

	public static function copy<T>(o:Null<T>):Null<T> {
		return cast NativeReflect.copy(o);
	}

	@:overload(function(f:Array<Dynamic>->Void):Dynamic {})
	public static function makeVarArgs(f:Array<Dynamic>->Dynamic):Dynamic {
		return NativeReflect.makeVarArgs(f);
	}
}
