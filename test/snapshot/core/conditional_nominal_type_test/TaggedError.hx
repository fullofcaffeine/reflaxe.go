@:goBuildConstraint("exception_probe")
@:keep
class TaggedError extends haxe.Exception implements Tag {
	static final installed = check();

	public function new() {
		super("tagged");
	}

	public function code():Int {
		return 7;
	}

	static function check():Bool {
		var original = new TaggedError();
		var base:haxe.Exception = original;
		var erased:Any = base;
		if (!Std.isOfType(erased, haxe.Exception) || !Std.isOfType(erased, Tag))
			throw "conditional ancestry was not recognized";
		var view:Tag = original;
		if (!Std.isOfType(view, TaggedError))
			throw "interface view lost conditional class identity";
		var absent:TaggedError = null;
		if (Std.isOfType(absent, Tag) || Std.isOfType(absent, haxe.Exception))
			throw "null matched nominal membership";
		try {
			throw base;
		} catch (caught:TaggedError) {
			if (caught != original)
				throw "conditional catch identity changed";
		} catch (_:Dynamic) {
			throw "conditional subclass catch missed";
		}
		var same = try {
			throw base;
			false;
		} catch (caught:TaggedError) {
			caught == original;
		} catch (_:Dynamic) {
			false;
		};
		if (!same)
			throw "conditional expression catch identity changed";
		try {
			try {
				throw original;
			} catch (caught:haxe.Exception) {
				throw caught;
			}
		} catch (caught:Tag) {
			if (caught.code() != 7)
				throw "conditional interface catch changed";
		} catch (_:Dynamic) {
			throw "conditional interface catch missed";
		}
		return Registry.install();
	}
}
