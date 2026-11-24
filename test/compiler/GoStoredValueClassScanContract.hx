#if macro
import haxe.macro.Context;
import haxe.macro.Type;
import reflaxe.go.ast.GoType;
import reflaxe.go.compiler.GoStoredValueClassScan;
import reflaxe.go.naming.GoNaming;
#end

/**
	Contract for the classes that stored-value recovery tests.

	Why
	`GoCompiler.recoverStoredGoValue` tests its target only against the classes
	that `GoStoredValueClassScan` returns. Leaving out a class whose name the
	target can be loses a match. Leaving out a class whose test has side effects
	(a class with a superclass, or one whose name requires a shim group) loses a
	superclass queue entry or a shim group. Either mistake changes the output of
	some program without failing a BeadsHX or snapshot compile that happens to
	record the same facts elsewhere.

	What
	Pins the decision for real typed classes: a carrier target `*N` and an
	interface target `N` keep their class; a subclass and the class `Type` (which
	requires the `type_metadata` shim group) are kept for every target; targets
	that cannot name a class keep only those two. Also pins `localNominalName`
	and that a changed class list is read again.

	How
	After the initialization macros, loads the fixture classes and asks the scan
	about each target. The expected lists are in project order.
**/
class GoStoredValueClassScanContract {
	#if macro
	public static function check():Void {
		Context.onAfterInitMacros(run);
	}

	static function require(value:Bool, message:String):Void {
		if (!value)
			Context.fatalError("GoStoredValueClassScanContract: " + message, Context.currentPos());
	}

	static function load(path:String):ClassType {
		return switch (Context.getType(path)) {
			case TInst(ref, _): ref.get();
			case _: throw "not a class: " + path;
		};
	}

	static function names(classes:Array<ClassType>):String {
		return [for (source in classes) source.name].join(",");
	}

	static function expect(scan:GoStoredValueClassScan, classes:Array<ClassType>, target:String, expected:String):Void {
		final tested = names(scan.testedClasses(classes, GoType.parse(target)));
		require(tested == expected, "target " + target + " tests [" + tested + "], expected [" + expected + "]");
	}

	static function run():Void {
		final module = "GoStoredValueClassScanContractFixture";
		final base = load(module);
		final child = load(module + ".StoredScanChild");
		final shape = load(module + ".StoredScanShape");
		final plain = load(module + ".StoredScanPlain");
		final typeClass = load("Type");

		// The scan compares the names `classTypeName` gives these classes.
		require(GoStoredValueClassScan.typeName(child) == GoNaming.typeSymbol(child.pack, child.name), "type name differs from GoNaming");
		require(GoStoredValueClassScan.typeName(child) == "StoredScanChild", "unexpected child type name " + GoStoredValueClassScan.typeName(child));

		// A superclass and a shim group are the two side effects of a failing test.
		require(GoStoredValueClassScan.testHasEffects(child), "a subclass test can queue and mark its superclass");
		require(GoStoredValueClassScan.testHasEffects(typeClass), "naming Type requires the type_metadata shim group");
		for (source in [base, shape, plain])
			require(!GoStoredValueClassScan.testHasEffects(source), source.name + " has no superclass and no shim group");

		final classes = [plain, base, child, shape, typeClass];
		final scan = new GoStoredValueClassScan();

		// A carrier target keeps its class, and the subclass that embeds it.
		expect(scan, classes, "*GoStoredValueClassScanContractFixture", "GoStoredValueClassScanContractFixture,StoredScanChild,Type");
		expect(scan, classes, "*StoredScanChild", "StoredScanChild,Type");
		// An interface target keeps its interface.
		expect(scan, classes, "StoredScanShape", "StoredScanChild,StoredScanShape,Type");
		// A named target keeps a class of that name; the caller's checks decide.
		expect(scan, classes, "StoredScanPlain", "StoredScanPlain,StoredScanChild,Type");
		expect(scan, classes, "*StoredScanPlain", "StoredScanPlain,StoredScanChild,Type");

		// Targets that cannot name a class keep only the classes with side effects.
		for (target in [
			"int",
			"any",
			"string",
			"*hxrt.Array",
			"hxrt.StoredScanShape",
			"*fmt.StoredScanPlain",
			"**StoredScanPlain",
			"[]*StoredScanPlain",
			"map[string]StoredScanShape",
			"func(StoredScanShape) *StoredScanPlain",
			"*NoSuchClass"
		])
			expect(scan, classes, target, "StoredScanChild,Type");

		// The scan reads a replaced, appended or removed class again.
		final changed = [base, plain, child];
		expect(scan, changed, "*StoredScanChild", "StoredScanChild");
		changed[2] = shape;
		expect(scan, changed, "*StoredScanChild", "");
		changed.push(typeClass);
		expect(scan, changed, "int", "Type");
		expect(scan, [child], "*GoStoredValueClassScanContractFixture", "StoredScanChild");
		scan.reset();
		expect(scan, classes, "*StoredScanChild", "StoredScanChild,Type");

		// `localNominalName` accepts exactly `N` and `*N` in the current package.
		require(GoType.parse("Shape").localNominalName() == "Shape", "N");
		require(GoType.parse("*Shape").localNominalName() == "Shape", "*N");
		for (target in [
			"**Shape",
			"[]Shape",
			"*pkg.Shape",
			"pkg.Shape",
			"int",
			"byte",
			"any",
			"interface{}",
			"map[string]Shape",
			"chan Shape"
		])
			require(GoType.parse(target).localNominalName() == null, target + " is not a local class type");
	}
	#end
}
