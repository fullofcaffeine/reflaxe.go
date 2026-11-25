#if macro
import haxe.macro.Context;
import haxe.macro.Type;
import reflaxe.go.ast.GoAST.GoExpr;
import reflaxe.go.ast.GoASTPrinter;
import reflaxe.go.compiler.GoBuildConstraint;
import reflaxe.go.compiler.GoNominalAncestry;
import reflaxe.go.compiler.GoNominalTypeTest;
import reflaxe.go.compiler.GoRecoveryAvailability;
import reflaxe.go.naming.GoNaming;
#end

/**
	Contract for the ancestry walk of nominal type-test discovery.

	Why
	`GoNominalTypeTest.discover` decides which generated classes belong to each
	demanded target with `GoNominalAncestry`, which walks each class's ancestry
	once instead of once per target. Naming a class can record compiler facts,
	so the walk must name the same classes, in the same order, as the recursive
	test it replaced, and give the same answers. A walk that names a class the old
	tests never reached can add a shim group or an extern import to the output.

	What
	For every ordered choice of up to three targets from the fixture, compares the
	answers and the order in which classes are first named with a copy of the
	replaced test, over the same candidates in the same target-major order. Then
	checks the cases that `discover` emits for three targets.

	How
	After the initialization macros, loads the fixture classes. Both walks use a
	naming function that logs each call; both first name the targets, as
	`matches` does when it records a target.
**/
class GoNominalAncestryContract {
	#if macro
	public static function check():Void {
		Context.onAfterInitMacros(run);
	}

	static function require(value:Bool, message:String):Void {
		if (!value)
			Context.fatalError("GoNominalAncestryContract: " + message, Context.currentPos());
	}

	static function load(path:String):ClassType {
		return switch (Context.getType(path)) {
			case TInst(ref, _): ref.get();
			case _: throw "not a class: " + path;
		};
	}

	/** The replaced test, unchanged apart from its naming function. */
	static function oldBelongsTo(source:ClassType, target:ClassType, visited:Map<String, Bool>, className:ClassType->String):Bool {
		var name = className(source);
		if (name == className(target))
			return true;
		if (visited.exists(name))
			return false;
		visited.set(name, true);
		if (source.superClass != null && oldBelongsTo(source.superClass.t.get(), target, visited, className))
			return true;
		for (implemented in source.interfaces)
			if (oldBelongsTo(implemented.t.get(), target, visited, className))
				return true;
		return false;
	}

	/** The names in `log` in the order of their first appearance. */
	static function firstNamed(log:Array<String>):String {
		final seen = new Map<String, Bool>();
		final order = new Array<String>();
		for (name in log)
			if (!seen.exists(name)) {
				seen.set(name, true);
				order.push(name);
			}
		return order.join(",");
	}

	static function compare(targets:Array<ClassType>, candidates:Array<ClassType>):Void {
		final label = [for (target in targets) target.name].join(",");
		final oldLog = new Array<String>();
		final oldName = (source:ClassType) -> {
			oldLog.push(source.name);
			source.name;
		};
		final newLog = new Array<String>();
		final newName = (source:ClassType) -> {
			newLog.push(source.name);
			source.name;
		};
		for (target in targets) {
			oldName(target);
			newName(target);
		}
		final ancestries = [for (candidate in candidates) new GoNominalAncestry(candidate, newName)];
		final oldAnswers = new Array<String>();
		final newAnswers = new Array<String>();
		for (target in targets)
			for (index in 0...candidates.length) {
				oldAnswers.push(candidates[index].name + "<" + target.name + "=" + oldBelongsTo(candidates[index], target, new Map(), oldName));
				newAnswers.push(candidates[index].name + "<" + target.name + "=" + ancestries[index].contains(target.name));
			}
		require(newAnswers.join(" ") == oldAnswers.join(" "), "targets ["
			+ label
			+ "] answer "
			+ newAnswers.join(" ")
			+ ", old test "
			+ oldAnswers.join(" "));
		require(firstNamed(newLog) == firstNamed(oldLog), "targets ["
			+ label
			+ "] name "
			+ firstNamed(newLog)
			+ ", old test "
			+ firstNamed(oldLog));
	}

	/** The concrete case types of each emitted test, as `symbol:Case|Case`. */
	static function emittedCases(text:String):String {
		final out = new Array<String>();
		var current:Null<String> = null;
		var cases = new Array<String>();
		for (line in text.split("\n")) {
			final trimmed = StringTools.trim(line);
			if (StringTools.startsWith(trimmed, "func ")) {
				if (current != null)
					out.push(current + ":" + cases.join("|"));
				current = trimmed.substring(5, trimmed.indexOf("("));
				cases = [];
			} else if (StringTools.startsWith(trimmed, "case *")) {
				cases.push(trimmed.substring(6, trimmed.length - 1));
			}
		}
		if (current != null)
			out.push(current + ":" + cases.join("|"));
		return out.join(" ");
	}

	static function run():Void {
		final module = "GoNominalAncestryContractFixture";
		final root = load(module);
		final shape = load(module + ".AncestryShape");
		final named = load(module + ".AncestryNamed");
		final other = load(module + ".AncestryOther");
		final extra = load(module + ".AncestryExtra");
		final middle = load(module + ".AncestryMiddle");
		final leaf = load(module + ".AncestryLeaf");
		final loose = load(module + ".AncestryLoose");
		final candidates = [leaf, middle, root, loose];
		final all = [root, shape, named, other, extra, middle, leaf, loose];

		for (first in all) {
			compare([first], candidates);
			for (second in all) {
				if (second == first)
					continue;
				compare([first, second], candidates);
				for (third in all)
					if (third != first && third != second)
						compare([first, second, third], candidates);
			}
		}

		// The walk itself: a matched class's ancestors stay reachable, and a class
		// already met answers without walking further. Interfaces follow in the
		// order of `ClassType.interfaces`, which the old test used too.
		final log = new Array<String>();
		final ancestry = new GoNominalAncestry(leaf, source -> {
			log.push(source.name);
			source.name;
		});
		require(ancestry.contains("AncestryMiddle"), "leaf extends middle");
		require(log.join(",") == "AncestryLeaf,AncestryMiddle", "a match stops the walk, named " + log.join(","));
		require(ancestry.contains("AncestryNamed"), "middle's interface stays reachable after middle matched");
		require(ancestry.contains("AncestryLeaf") && ancestry.contains("GoNominalAncestryContractFixture"), "met classes answer");
		require(!ancestry.contains("AncestryLoose"), "an unrelated class is no ancestor");
		require(log.join(",") == "AncestryLeaf,AncestryMiddle,GoNominalAncestryContractFixture,AncestryShape,AncestryNamed,AncestryShape,"
			+ [for (entry in leaf.interfaces) entry.t.get().name].join(","),
			"walk order "
			+ log.join(","));

		// `discover` emits one case per generated class that belongs to each target.
		final className = (source:ClassType) -> GoNaming.typeSymbol(source.pack, source.name);
		final availability = new GoRecoveryAvailability(_ -> null);
		for (source in candidates)
			availability.noteGeneratedType(className(source), "fixture", GoBuildConstraint.always());
		final test = new GoNominalTypeTest(availability, className);
		for (target in [shape, middle, other])
			test.matches(GoExpr.GoIdent("value"), target);
		final text = GoASTPrinter.printFile({packageName: "main", imports: [], decls: [for (plan in test.discover(candidates)) plan.declaration]});
		final cases = emittedCases(text);
		require(cases == "hx_is_nominal_AncestryMiddle:AncestryLeaf|AncestryMiddle hx_is_nominal_AncestryOther:AncestryLeaf "
			+ "hx_is_nominal_AncestryShape:AncestryLeaf|AncestryMiddle|GoNominalAncestryContractFixture",
			"discover emits "
			+ cases);
	}
	#end
}
