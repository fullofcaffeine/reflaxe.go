package reflaxe.go.compiler;

#if macro
import haxe.macro.Type.ClassType;
import reflaxe.go.ast.GoType;
import reflaxe.go.naming.GoNaming;

/** One project class and the two facts the scan keeps for it; both are fixed for the class. */
private typedef GoStoredValueClassScanEntry = {
	final source:ClassType;
	final typeName:String;
	final testHasEffects:Bool;
}

/**
	What: Chooses which project classes `GoCompiler.recoverStoredGoValue` must
	test for one target type, in project order.

	Why: Stored-value recovery tests its target against every project class. The
	target matches a class only when it is that class's carrier `*N` or its
	interface `N`. Each test first asks whether the class is a generated
	project class, which reads the class's source position. On the BeadsHX
	bootstrap compile about 1,300 recoveries each tested about 210 classes, and
	almost every test failed. That scan cost about 3 s of a 47 s compile.

	How: A test of a class that fails has two possible side effects. When the
	class has a superclass, `hasGeneratedVirtualCarrier` reaches
	`projectSuperClass`, which can queue that superclass, mark it as an embedded
	superclass, or require it as source-owned standard-library code. When the
	class is a carrier or an interface, `classTypeName` calls `noteStdlibClass`,
	which can require a standard-library shim group. A class without a
	superclass whose name requires no shim group has neither effect, so its test
	only reads state (apart from `declaresInstanceMembers`, which caches a fact of
	the class itself). Such a class is skipped unless its Go type name is the
	target's `N`.

	The caller tests the returned classes in order with its unchanged checks. A
	skipped class could not match and its test would write nothing, so the
	result and every side effect are the same as a test of every class.

	The classes with effects must stay. While classes are lowered, the list
	also holds classes that are queued but not lowered yet, so a test can be the
	first to record a superclass. On the BeadsHX bootstrap compile, a recovery
	inside another class requires `haxe.exceptions.PosException`, the superclass
	of `haxe.exceptions.NotImplementedException`, before that subclass is
	lowered. A scan that also skipped those classes changed the BeadsHX output.

	The scan remembers each class's facts by its position in the list. A
	position that now holds a different class object is read again, so the
	list may grow or change between calls.

	If `hasGeneratedVirtualCarrier` or `classTypeName` gains another write,
	`testHasEffects` must keep every class whose test can make it.
**/
class GoStoredValueClassScan {
	final entries:Array<GoStoredValueClassScanEntry> = [];

	public function new() {}

	/** Forget the remembered facts; a new compilation brings new class objects. */
	public function reset():Void {
		entries.resize(0);
	}

	/**
		Return the classes of `classes` that a recovery to `target` must test, in
		their order in `classes`: each class whose Go type name is the target's
		`N` or `*N`, and each class whose test has side effects.
	**/
	public function testedClasses(classes:Array<ClassType>, target:GoType):Array<ClassType> {
		final name = target.localNominalName();
		final tested = new Array<ClassType>();
		for (index in 0...classes.length) {
			final source = classes[index];
			if (index >= entries.length || entries[index].source != source)
				entries[index] = {source: source, typeName: typeName(source), testHasEffects: testHasEffects(source)};
			final entry = entries[index];
			if (entry.testHasEffects || entry.typeName == name)
				tested.push(source);
		}
		return tested;
	}

	/**
		The Go type name of a non-extern class: the name `classTypeName` returns
		without its side effect. An extern class never matches a stored-value
		target, so its name only decides whether it is tested in vain.
	**/
	public static function typeName(source:ClassType):String {
		return GoNaming.typeSymbol(source.pack, source.name);
	}

	/**
		True when a failing stored-value test of `source` can change compiler
		state: the class has a superclass (`projectSuperClass` may queue and mark
		it) or its name requires a standard-library shim group (`noteStdlibClass`).
	**/
	public static function testHasEffects(source:ClassType):Bool {
		return source.superClass != null || GoStdlibShimClassifier.requiredGroupsForClass(source).length > 0;
	}
}
#end
