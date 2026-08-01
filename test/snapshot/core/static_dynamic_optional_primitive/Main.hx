class Main {
	public static dynamic function schedule(callback:Void->Void, ?timeout:Int):Void {
		callback();
	}

	public static dynamic function summarize(enabled:Bool = true, count:Int = 3, ratio:Float = 1.5):Float {
		return enabled ? count + ratio : ratio;
	}

	static function namedDefault(count:Int = 8):Int {
		return count;
	}

	static function invoke(callback:(?count:Int) -> Int):Int {
		return callback();
	}

	static function makeLocalCarrier():(?count:Int) -> Int {
		var source = function(count:Int = 10):Int {
			return count;
		};
		return source;
	}

	static function main():Void {
		schedule(function() trace("scheduled"));
		schedule(function() {}, 10);
		trace(summarize());
		trace(summarize(false, 9, 2.5));

		summarize = function(enabled:Bool = false, count:Int = 4, ratio:Float = 2.0):Float {
			return enabled ? count + ratio : ratio;
		};
		trace(summarize(false, 8, 3.5));
		Sys.println(summarize());

		var local:(?enabled:Bool, ?count:Int) -> Int = null;
		local = function(enabled:Bool = true, count:Int = 6):Int {
			return enabled ? count : -count;
		};
		trace(local(false, 11));

		var initialized:(?enabled:Bool, ?count:Int) -> Int = function(enabled:Bool = true, count:Int = 6):Int {
			return enabled ? count : -count;
		};
		Sys.println(initialized());
		initialized = function(enabled:Bool = false, count:Int = 4):Int {
			return enabled ? count : -count;
		};
		Sys.println(initialized());

		var casted:(?count:Int) -> Int = cast function(count:Int = 2):Int {
			return count;
		};
		Sys.println(casted());

		var namedCarrier:(?count:Int) -> Int = namedDefault;
		Sys.println(namedCarrier());

		var localSource = function(count:Int = 5):Int {
			return count;
		};
		var localAlias:(?count:Int) -> Int = localSource;
		Sys.println(localAlias());
		Sys.println(invoke(localSource));
		Sys.println(invoke(function(count:Int = 12):Int {
			return count;
		}));
		Sys.println(makeLocalCarrier()());

		var localOther = function(count:Int = 6):Int {
			return count;
		};
		var conditional:(?count:Int) -> Int = localAlias() == 5 ? localSource : localOther;
		Sys.println(conditional());
		var switched:(?count:Int) -> Int = switch (localAlias()) {
			case 5: localOther;
			default: localSource;
		};
		Sys.println(switched());
		var blocked:(?count:Int) -> Int = {
			var observed = localAlias();
			if (observed != 5)
				throw "unexpected";
			localSource;
		};
		Sys.println(blocked());
		var caught:(?count:Int) -> Int = try localOther catch (_:Dynamic) localSource;
		Sys.println(caught());

		var callbacks:Array<(?count:Int) -> Int> = [localSource];
		Sys.println(callbacks[0]());
		callbacks[0] = localOther;
		Sys.println(callbacks[0]());
		var holder:{callback:(?count:Int) -> Int} = {callback: localSource};
		Sys.println(holder.callback());
	}
}
