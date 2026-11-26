#if macro
import haxe.macro.Context;
import haxe.macro.Type;
import reflaxe.compiler.TypeUsageTracker;
import reflaxe.compiler.TypeUsageTracker.TypeOrModuleType;
import reflaxe.compiler.TypeUsageTracker.TypeUsageMap;
import reflaxe.go.compiler.GoTypeUsageLedger;
import reflaxe.go.compiler.GoTypeUsageTracker;
import reflaxe.helpers.TypeHelper;

using reflaxe.helpers.ModuleTypeHelper;
#end

/**
	Contract for the compiler's type-usage tracker.

	Why
	`GoTypeUsageTracker` replaces Reflaxe's `TypeUsageTracker` as the source of
	the maps that the type-usage ledger reads, and it keeps the id of each enum,
	typedef, and abstract for the whole compilation. A map that differs from
	Reflaxe's in one entry, one level, or in which of two entries with one id it
	keeps would change the type-usage report and every registry and planner
	decision built on it.

	What
	For every module type of the two fixture modules and of a set of standard
	library modules, tracked by one tracker in both orders, proves that:
	- the map has the same levels as Reflaxe's map for that module type, and each
	  level has the same entries in the same order. An entry is compared by its
	  id, its kind, and its ledger shape, so the two fixture modules' type
	  parameters `pick.V`, which print alike, are told apart;
	- a ledger fed with either map writes the same report;
	- the fixture maps use every level;
	- the tracker kept ids only for enums, typedefs, and abstracts, and kept one
	  for each of the same-named private declarations of the two modules.

	How
	After the initialization macros, loads the modules and calls Reflaxe's
	`TypeUsageTracker.trackTypesInModuleType(...)` (the unchanged vendored
	tracker, which is the oracle) next to `GoTypeUsageTracker.track(...)` for
	each module type.
**/
@:access(reflaxe.go.compiler.GoTypeUsageLedger)
@:access(reflaxe.go.compiler.GoTypeUsageTracker)
class GoTypeUsageTrackerContract {
	#if macro
	static final FIXTURE_MODULES = ["GoTypeUsageTrackerContractFixture", "GoTypeUsageTrackerContractOtherFixture"];

	static final STD_MODULES = [
		"Lambda",
		"StringTools",
		"haxe.EnumTools",
		"haxe.Int64",
		"haxe.ds.ArraySort",
		"haxe.ds.BalancedTree",
		"haxe.ds.List",
		"haxe.ds.Option",
		"haxe.ds.Vector",
		"haxe.format.JsonPrinter",
		"haxe.io.Bytes"
	];

	public static function check():Void {
		Context.onAfterInitMacros(run);
	}

	static function require(value:Bool, message:String):Void {
		if (!value)
			Context.fatalError("GoTypeUsageTrackerContract: " + message, Context.currentPos());
	}

	static function moduleTypes(module:String):Array<ModuleType> {
		final result:Array<ModuleType> = [];
		for (type in Context.getModule(module)) {
			final moduleType = TypeHelper.toModuleType(type);
			require(moduleType != null, "module type of " + Std.string(type));
			if (moduleType != null)
				result.push(moduleType);
		}
		return result;
	}

	/** The levels and entries of one map, in the map's own order. */
	static function render(usage:TypeUsageMap, shapes:GoTypeUsageLedger):String {
		final lines:Array<String> = [];
		for (level in usage.keys()) {
			lines.push("level " + Std.int(level));
			final entries = usage.get(level);
			if (entries != null)
				for (entry in entries)
					lines.push("  " + entryLabel(entry, shapes));
		}
		return lines.join("\n");
	}

	static function entryLabel(entry:TypeOrModuleType, shapes:GoTypeUsageLedger):String {
		return switch (entry) {
			case EModuleType(moduleType):
				"module "
				+ moduleType.getUniqueId()
				+ " "
				+ GoTypeUsageLedger.renderShapeJson(GoTypeUsageLedger.moduleTypeShape(moduleType));
			case EType(type):
				"type " + TypeHelper.getUniqueId(type) + " " + GoTypeUsageLedger.renderShapeJson(shapes.typeShape(type));
		};
	}

	/** The ids of the type parameter classes that print as `printed` in one map. */
	static function parameterIds(usage:TypeUsageMap, printed:String, into:Map<String, Bool>):Void {
		for (level in usage.keys()) {
			final entries = usage.get(level);
			if (entries == null)
				continue;
			for (entry in entries)
				switch (entry) {
					case EModuleType(TClassDecl(classRef)) if (classRef.toString() == printed):
						into.set(TClassDecl(classRef).getUniqueId(), true);
					case _:
				}
		}
	}

	static function run():Void {
		final fixtureTypes:Array<ModuleType> = [];
		for (module in FIXTURE_MODULES)
			for (moduleType in moduleTypes(module))
				fixtureTypes.push(moduleType);
		final all = fixtureTypes.copy();
		for (module in STD_MODULES)
			for (moduleType in moduleTypes(module))
				all.push(moduleType);
		require(fixtureTypes.length >= 14, "fixture module types were loaded (" + fixtureTypes.length + ")");

		// The compiler tracks module types after Haxe has typed every field. Here
		// typing is still lazy: the first `field.expr()` of a field types it, and
		// only a later read of the field sees its typed signature. One discarded
		// pass puts every field into the compiler's state before the comparison.
		for (moduleType in all)
			TypeUsageTracker.trackTypesInModuleType(moduleType);

		// One tracker for every module type and both orders, so kept ids are reused.
		final tracker = new GoTypeUsageTracker();
		final shapes = new GoTypeUsageLedger();
		final expectedLedger = new GoTypeUsageLedger();
		final actualLedger = new GoTypeUsageLedger();
		final levels = new Map<Int, Bool>();
		final pickIds = new Map<String, Bool>();
		var entries = 0;
		for (round in 0...2) {
			final ordered = round == 0 ? all : [for (index in 0...all.length) all[all.length - 1 - index]];
			for (index in 0...ordered.length) {
				final moduleType = ordered[index];
				final expected = TypeUsageTracker.trackTypesInModuleType(moduleType);
				final actual = tracker.track(moduleType);
				final expectedText = render(expected, shapes);
				final actualText = render(actual, shapes);
				require(actualText == expectedText, "map of " + moduleType.getPath() + ":\n" + actualText + "\ninstead of\n" + expectedText);
				if (round == 0) {
					expectedLedger.collect(moduleType, expected);
					actualLedger.collect(moduleType, actual);
					if (index < fixtureTypes.length) {
						for (level in actual.keys()) {
							final values = actual.get(level);
							if (values != null && values.length > 0) {
								levels.set(Std.int(level), true);
								entries += values.length;
							}
						}
						parameterIds(actual, "pick.V", pickIds);
					}
				}
			}
		}
		require(entries > 100, "fixture maps have entries (" + entries + ")");
		for (index in 0...TypeUsageLevel.LevelCount())
			require(levels.exists(1 << index), "fixture maps use level " + (1 << index));
		require([for (id in pickIds.keys()) id].length == 2, "both modules' pick.V were tracked with their own ids");

		final expectedReport = GoTypeUsageLedger.renderJson(expectedLedger.snapshot([], "hxrt"));
		final actualReport = GoTypeUsageLedger.renderJson(actualLedger.snapshot([], "hxrt"));
		require(expectedReport.indexOf('"module": "GoTypeUsageTrackerContractFixture"') != -1, "fixture owners reach the ledger");
		require(actualReport == expectedReport, "ledger report:\n" + actualReport + "\ninstead of\n" + expectedReport);

		// Ids are kept only for enums, typedefs, and abstracts, one per declaration.
		final keys = [for (key in tracker.declarationIds.keys()) key];
		require(keys.length > 10, "the tracker kept declaration ids (" + keys.length + ")");
		for (key in keys)
			require(["E", "T", "A"].indexOf(key.charAt(0)) != -1, "kept the id of " + key);
		for (name in ["THidden", "AScore", "EChoice"]) {
			final kind = name.charAt(0);
			final suffix = "." + name.substr(1);
			final matching = [
				for (key in keys)
					if (key.charAt(0) == kind && StringTools.endsWith(key, suffix)) tracker.declarationIds.get(key)
			];
			require(matching.length == 2 && matching[0] != matching[1], "private "
				+ name.substr(1)
				+ " of both modules kept apart: "
				+ matching);
		}

		Sys.println("type usage tracker contract passed");
	}
	#end
}
