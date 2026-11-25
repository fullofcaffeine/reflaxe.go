import haxe.Constraints.IMap;
import haxe.ds.BalancedTree;
import haxe.exceptions.NotImplementedException;
import haxe.exceptions.PosException;

/**
	What: Tests values against standard-library classes that join the program
	after this class is lowered.

	Why: The compiler lowers classes in name order, so this class is lowered
	before `Main`. `NotImplementedException` and `BalancedTree` join the program
	only while `Main` is lowered, when `Main` creates them. A test built while
	this class was lowered knew neither class, so it returned false.

	How: One method receives a `Dynamic` value and one an interface-typed value.
	The generated Go must test each value against every class of the finished
	program.
**/
class EarlyChecks {
	public static function isException(value:Dynamic):Bool {
		return Std.isOfType(value, haxe.Exception);
	}

	public static function isTree(map:IMap<String, Int>):Bool {
		return Std.isOfType(map, BalancedTree);
	}
}

/**
	What: Tests and converts values typed as `haxe.Exception` while that class is
	still waiting to join the program.

	Why: The parameter type of `isPosException` makes the compiler require
	`haxe.Exception`, but the class joins the class list only after this class
	is lowered. `PosException` is not required at all yet. The test of a base
	value used to compare only the exact `PosException` type, and the
	conversion of a `Dynamic` value used to be a plain Go type assertion, which
	panics on a subclass value.

	How: The methods are lowered in name order, so `toException` is lowered
	after `isPosException` has required `haxe.Exception`.
**/
class EarlyExceptionViews {
	public static function isPosException(value:haxe.Exception):Bool {
		return Std.isOfType(value, PosException);
	}

	public static function toException(value:Dynamic):haxe.Exception {
		var exception:haxe.Exception = value;
		return exception;
	}
}

/**
	What: Converts a `Dynamic` value to `PosException`, which this method is the
	first to require.

	Why: The conversion requires `PosException` while this class is lowered, so
	the class is still waiting to join the class list. The conversion used to be
	a plain Go type assertion, which panics on a `NotImplementedException`.

	How: The class name sorts after `EarlyExceptionViews`, whose type test of
	`PosException` does not require that class.
**/
class EarlyPosExceptionView {
	public static function toPosException(value:Dynamic):PosException {
		var exception:PosException = value;
		return exception;
	}
}

class Main {
	static function main() {
		var error:Dynamic = new NotImplementedException();
		Sys.println("Dynamic NotImplementedException is haxe.Exception=" + EarlyChecks.isException(error));
		Sys.println("Dynamic String is haxe.Exception=" + EarlyChecks.isException("text"));

		var tree = new BalancedTree<String, Int>();
		tree.set("a", 1);
		Sys.println("IMap BalancedTree is BalancedTree=" + EarlyChecks.isTree(tree));

		Sys.println("haxe.Exception NotImplementedException is PosException=" + EarlyExceptionViews.isPosException(error));
		Sys.println("haxe.Exception plain is PosException=" + EarlyExceptionViews.isPosException(new haxe.Exception("plain")));

		var asException = EarlyExceptionViews.toException(error);
		Sys.println("Dynamic as haxe.Exception keeps NotImplementedException=" + Std.isOfType(asException, NotImplementedException));
		var asPosException = EarlyPosExceptionView.toPosException(error);
		Sys.println("Dynamic as PosException keeps NotImplementedException=" + Std.isOfType(asPosException, NotImplementedException));
		Sys.println("Dynamic null as PosException is null=" + (EarlyPosExceptionView.toPosException(null) == null));
	}
}
