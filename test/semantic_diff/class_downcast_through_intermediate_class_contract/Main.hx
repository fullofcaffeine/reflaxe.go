import haxe.exceptions.NotImplementedException;
import haxe.exceptions.PosException;

/**
	An unchecked cast of a base-typed value to an intermediate class: with
	`Leaf extends Mid extends Base`, a `Base` that holds a `Leaf` is a `Mid`, so
	`var mid:Mid = cast value` must succeed. The value keeps its identity, its
	overrides and its fields.

	`ACaster` sorts before every class it casts, and `Main` sorts after them,
	so the casts are lowered both before and after their target classes. The
	last cast targets a standard-library class that joins the program late.
**/
class ACaster {
	public static function toMid(value:Base):Mid {
		return cast value;
	}

	public static function toLeaf(value:Base):Leaf {
		return cast value;
	}
}

class Base {
	public var tag:String;

	public function new() {
		tag = "base";
	}

	public function describe():String {
		return "Base(" + tag + ")";
	}
}

class Mid extends Base {
	public function new() {
		super();
		tag = "mid";
	}

	override public function describe():String {
		return "Mid(" + tag + ")";
	}

	public function midName():String {
		return "midName:" + tag;
	}
}

class Leaf extends Mid {
	public function new() {
		super();
		tag = "leaf";
	}

	override public function describe():String {
		return "Leaf(" + tag + ")";
	}
}

class Tip extends Leaf {
	public function new() {
		super();
		tag = "tip";
	}
}

class Main {
	static function main() {
		var leafAsBase:Base = new Leaf();
		var mid:Mid = cast leafAsBase;
		Sys.println("leaf as Mid tag=" + mid.tag);
		Sys.println("leaf as Mid describe=" + mid.describe());
		Sys.println("leaf as Mid midName=" + mid.midName());
		Sys.println("leaf as Mid isOfType Leaf=" + Std.isOfType(mid, Leaf));
		var midAsBase:Base = mid;
		Sys.println("leaf as Mid keeps identity=" + (midAsBase == leafAsBase));
		mid.tag = "changed";
		Sys.println("write through Mid seen by Base=" + leafAsBase.tag);

		var tipAsBase:Base = new Tip();
		var tipAsMid:Mid = cast tipAsBase;
		var tipAsLeaf:Leaf = cast tipAsMid;
		Sys.println("tip as Mid=" + tipAsMid.describe());
		Sys.println("tip as Leaf from Mid=" + tipAsLeaf.describe());

		var exactMid:Base = new Mid();
		var exact:Mid = cast exactMid;
		Sys.println("exact Mid=" + exact.describe());

		Sys.println("caster toMid=" + ACaster.toMid(new Tip()).describe());
		Sys.println("caster toLeaf=" + ACaster.toLeaf(new Tip()).describe());

		var none:Base = null;
		var noneMid:Mid = cast none;
		Sys.println("null stays null=" + (noneMid == null));

		var error:haxe.Exception = new NotImplementedException("late");
		var pos:PosException = cast error;
		Sys.println("std exception as PosException=" + pos.message);
	}
}
