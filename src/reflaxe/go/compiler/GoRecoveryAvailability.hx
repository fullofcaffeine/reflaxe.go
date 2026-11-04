package reflaxe.go.compiler;

import reflaxe.go.ast.GoPackageName;
import reflaxe.go.ast.GoType;

/** Requirements of one generated branch, excluding child helpers' implementation cases. */
enum GoRecoveryDependency {
	TypeReference(type:GoType);
	HelperReference(symbol:String);
}

/** A typed builder result together with the storage/helper contracts it requested. */
typedef GoRecoveryCapture<T> = {
	final value:T;
	final dependencies:Array<GoRecoveryDependency>;
}

/**
	Owns declaration availability during one compilation. Physical native identity
	uses the import path, not the alias selected by a consumer. Each declaration
	owner supplies a witness; alternatives combine with OR and required types with
	AND. Missing ownership remains unresolved until the discovery loop closes.
	Helper contracts expose only their signatures, so recursive recovery does not
	inherit every conditional implementation branch of a child helper.
**/
class GoRecoveryAvailability {
	final importPath:GoPackageName->Null<String>;
	final witnesses:Map<String, Map<String, GoBuildConstraint>> = [];
	final packages:Map<String, GoBuildConstraint> = [];
	final helpers:Map<String, GoType> = [];
	final captures:Array<Array<GoRecoveryDependency>> = [];
	final unresolved:Map<String, Bool> = [];
	var factRevision = 0;
	var frozen = false;

	public function new(importPath:GoPackageName->Null<String>) {
		this.importPath = importPath;
	}

	public function reset():Void {
		witnesses.clear();
		packages.clear();
		helpers.clear();
		captures.resize(0);
		unresolved.clear();
		factRevision = 0;
		frozen = false;
	}

	/** Revision changes only when a declaration or helper fact actually changes. */
	public function revision():Int
		return factRevision;

	/** Only explicit whole-package ABI contracts, such as the generated runtime, use this. */
	public function notePackage(path:String, selection:GoBuildConstraint):Void {
		var previous = packages.get(path);
		if (previous != null) {
			if (previous.key() != selection.key())
				throw new haxe.Exception("Runtime package changed its availability contract: " + path);
			return;
		}
		changed();
		packages.set(path, selection);
	}

	public function noteGeneratedType(name:String, owner:String, selection:GoBuildConstraint):Void {
		noteWitness("local:" + name, owner, selection);
	}

	public function noteNativeType(path:String, name:String, owner:String, selection:GoBuildConstraint):Void {
		noteWitness(nativeIdentity(path, name), owner, selection);
	}

	function noteWitness(identity:String, owner:String, selection:GoBuildConstraint):Void {
		var owners = witnesses.get(identity);
		var previous = owners == null ? null : owners.get(owner);
		if (previous != null) {
			if (previous.key() != selection.key())
				throw new haxe.Exception("Recovery declaration owner changed its selection: " + owner);
			return;
		}
		changed();
		if (owners == null) {
			owners = [];
			witnesses.set(identity, owners);
		}
		owners.set(owner, selection);
	}

	public function noteHelper(symbol:String, signature:GoType):Void {
		var previous = helpers.get(symbol);
		if (previous != null) {
			if (!previous.sameRepresentation(signature))
				throw new haxe.Exception("Recovery helper changed its signature: " + symbol);
			return;
		}
		changed();
		helpers.set(symbol, signature);
	}

	/** Recheck unresolved requirements after newly discovered declarations are admitted. */
	public function beginPass():Void {
		if (frozen)
			throw new haxe.Exception("Recovery availability discovery restarted after freeze");
		unresolved.clear();
	}

	public function capture<T>(build:Void->T):GoRecoveryCapture<T> {
		var dependencies = new Array<GoRecoveryDependency>();
		captures.push(dependencies);
		var value = try {
			build();
		} catch (error:haxe.Exception) {
			captures.pop();
			throw error;
		};
		captures.pop();
		return {value: value, dependencies: dependencies};
	}

	public function requireType(type:GoType):Void {
		require(TypeReference(type));
	}

	public function requireHelper(symbol:String):Void {
		require(HelperReference(symbol));
	}

	function require(dependency:GoRecoveryDependency):Void {
		if (captures.length > 0)
			captures[captures.length - 1].push(dependency);
	}

	public function conditionForType(type:GoType):Null<GoBuildConstraint> {
		return conditionFor([TypeReference(type)]);
	}

	/** A condition is returned only after every reachable dependency has an owner. */
	public function conditionFor(dependencies:Array<GoRecoveryDependency>):Null<GoBuildConstraint> {
		var condition = GoBuildConstraint.always();
		var missing = new Map<String, Bool>();
		for (dependency in dependencies) {
			var type = switch (dependency) {
				case TypeReference(value): value;
				case HelperReference(symbol):
					var signature = helpers.get(symbol);
					if (signature == null)
						missing.set("helper:" + symbol, true);
					signature;
			};
			if (type == null)
				continue;
			for (name in type.namedDependencies()) {
				var identity = "local:" + name.name;
				var path:Null<String> = null;
				if (name.packageName != null) {
					path = importPath(name.packageName);
					if (path == null) {
						missing.set("import-alias:" + name.packageName.value(), true);
						continue;
					}
					identity = nativeIdentity(path, name.name);
				}
				var owners = witnesses.get(identity);
				var available:Null<GoBuildConstraint> = null;
				if (owners != null) {
					available = GoBuildConstraint.never();
					for (owner in owners)
						available = available.or(owner);
				} else if (path != null) {
					available = packages.get(path);
				}
				if (available == null)
					missing.set(identity, true);
				else
					condition = condition.and(available);
			}
		}
		// A branch that cannot be selected needs no otherwise unresolved declarations.
		if (condition.isNever())
			return condition;
		var complete = true;
		for (name in missing.keys()) {
			unresolved.set(name, true);
			complete = false;
		}
		return complete ? condition : null;
	}

	/** Refuse unresolved output rather than silently assigning it to common files. */
	public function freeze():Void {
		var names = [for (name in unresolved.keys()) name];
		if (names.length > 0) {
			names.sort((left, right) -> left < right ? -1 : left > right ? 1 : 0);
			throw new haxe.Exception("Unresolved recovery availability: " + names.join(", "));
		}
		frozen = true;
	}

	function changed():Void {
		if (frozen)
			throw new haxe.Exception("Recovery ownership changed after freeze");
		factRevision++;
	}

	static function nativeIdentity(path:String, name:String):String
		return "import:" + path + "#" + name;
}
