#if macro
import haxe.macro.Context;
import haxe.macro.Type;
import haxe.macro.TypedExprTools;
import reflaxe.go.compiler.GoTypeUsageLedger;
import reflaxe.go.compiler.GoTypeUsageLedger.GoAnonymousFieldShape;
import reflaxe.go.compiler.GoTypeUsageLedger.GoFunctionArgumentShape;
import reflaxe.go.compiler.GoTypeUsageLedger.GoImmutableList;
import reflaxe.go.compiler.GoTypeUsageLedger.GoTypeShape;
import reflaxe.go.compiler.GoTypeUsageLedger.GoTypeUsageTargetKind;
import reflaxe.go.compiler.GoTypeUsageLedger.GoUnknownTypeShapeReason;
import reflaxe.helpers.TypeHelper;

using reflaxe.helpers.ModuleTypeHelper;
#end

/**
	Contract for how the type-usage ledger converts macro types into shapes.

	Why
	The ledger keeps the path and id of each enum, typedef, and abstract it has
	read, and reads a class only once per use. It also keeps the JSON of every
	shape without children (a named type without type arguments, a type
	parameter, an unknown marker). These shortcuts must give exactly the shapes,
	recursion decisions, and JSON of the plain conversion and render. Otherwise
	the type-usage report, and every registry decision built on it, would change.

	What
	Proves, with one shared ledger and in both query orders, for every field
	and expression type of two fixture modules:
	- `trailId(...)` equals `TypeHelper.getUniqueId(...)`;
	- the shape renders to the same JSON as `plainShape(...)` rendered by
	  `plainRender(...)`, the render as it was before it kept leaf JSON.
	Also proves that leaves with one path but different kinds keep different
	JSON, that same-named private declarations and type parameters that
	print alike (`pick.V`) keep their own module paths, that nested iterators
	still meet the recursion guard, and that a monomorph Haxe binds after one
	conversion is seen bound by the next conversion through the same ledger.

	How
	After the initialization macros, loads both fixture classes, collects their
	field types and every expression and variable type as the ledger does, and
	compares each conversion with `plainShape(...)`. That function is the
	conversion as it was before the ledger kept declaration facts: it reads every
	declaration through the upstream `TypeHelper` and `ModuleTypeHelper`.
	Leaf JSON is compared before and after it is kept.
**/
@:access(reflaxe.go.compiler.GoTypeUsageLedger)
class GoTypeShapeContract {
	#if macro
	public static function check():Void {
		Context.onAfterInitMacros(run);
	}

	static function require(value:Bool, message:String):Void {
		if (!value)
			Context.fatalError("GoTypeShapeContract: " + message, Context.currentPos());
	}

	static function render(shape:GoTypeShape):String {
		return GoTypeUsageLedger.renderShapeJson(shape);
	}

	static function classType(name:String):ClassType {
		return switch (Context.getType(name)) {
			case TInst(ref, _):
				ref.get();
			case _:
				Context.fatalError("GoTypeShapeContract: " + name + " is not a class", Context.currentPos());
		};
	}

	static function memberFields(owner:ClassType):Array<ClassField> {
		final fields = owner.fields.get().concat(owner.statics.get());
		if (owner.constructor != null)
			fields.push(owner.constructor.get());
		return fields;
	}

	/** Collect the types the ledger converts: field types, expression types, and variable types. */
	static function collectTypes(owner:ClassType, into:Array<Type>):Void {
		for (field in memberFields(owner)) {
			into.push(field.type);
			final body = field.expr();
			if (body == null)
				continue;
			function visit(expr:TypedExpr):Void {
				into.push(expr.t);
				switch (expr.expr) {
					case TVar(variable, _):
						into.push(variable.t);
					case _:
				}
				TypedExprTools.iter(expr, visit);
			}
			visit(body);
		}
	}

	static function fieldType(owner:ClassType, name:String):Type {
		for (field in memberFields(owner))
			if (field.name == name)
				return field.type;
		return Context.fatalError("GoTypeShapeContract: missing fixture field " + name, Context.currentPos());
	}

	/** The type of the local function `pick` in `owner.local()`. */
	static function localPickType(owner:ClassType):Type {
		var found:Null<Type> = null;
		for (field in owner.statics.get()) {
			final body = field.name == "local" ? field.expr() : null;
			if (body == null)
				continue;
			function visit(expr:TypedExpr):Void {
				switch (expr.expr) {
					case TVar(variable, _) if (variable.name == "pick"):
						found = variable.t;
					case _:
				}
				TypedExprTools.iter(expr, visit);
			}
			visit(body);
		}
		require(found != null, "local pick in " + owner.name);
		return (found : Type);
	}

	/** The type parameter class that `pick<V>` takes, as its reference prints. */
	static function pickParameter(pickType:Type):{printed:String, type:Type} {
		return switch (pickType) {
			case TFun([argument], _):
				switch (argument.t) {
					case TInst(ref, _):
						{printed: ref.toString(), type: argument.t};
					case _:
						Context.fatalError("GoTypeShapeContract: pick argument is not a type parameter", Context.currentPos());
				}
			case _:
				Context.fatalError("GoTypeShapeContract: pick is not a one-argument function", Context.currentPos());
		};
	}

	static function run():Void {
		final first = classType("GoTypeShapeContractFixture");
		final other = classType("GoTypeShapeContractOtherFixture");
		final types:Array<Type> = [];
		collectTypes(first, types);
		collectTypes(other, types);
		require(types.length > 40, "fixture types were collected (" + types.length + ")");

		// One ledger for every conversion, in both orders, so kept facts are reused.
		final ledger = new GoTypeUsageLedger();
		for (round in 0...2) {
			final ordered = round == 0 ? types : [for (index in 0...types.length) types[types.length - 1 - index]];
			for (type in ordered) {
				require(ledger.trailId(type) == TypeHelper.getUniqueId(type), "trail id of " + Std.string(type));
				final expected = plainRender(plainShape(type));
				final actual = render(ledger.typeShape(type));
				require(actual == expected, "shape of " + Std.string(type) + ": " + actual + " instead of " + expected);
			}
		}
		require(GoTypeUsageLedger.leafShapeJsonByKey.keys().hasNext(), "leaf JSON was kept");

		// Leaf JSON is kept by kind and path: leaves that share a path keep their own JSON.
		for (path in ["GoTypeShapeContractFixture", "recursive"]) {
			final leaves = [
				GoTypeShape.Nominal(GoTypeUsageTargetKind.Class, path, GoImmutableList.fromArray([])),
				GoTypeShape.Nominal(GoTypeUsageTargetKind.Enum, path, GoImmutableList.fromArray([])),
				GoTypeShape.Nominal(GoTypeUsageTargetKind.Typedef, path, GoImmutableList.fromArray([])),
				GoTypeShape.Nominal(GoTypeUsageTargetKind.Abstract, path, GoImmutableList.fromArray([])),
				GoTypeShape.TypeParameter(path),
				GoTypeShape.UnknownShape(GoUnknownTypeShapeReason.Recursive)
			];
			final seen = new Map<String, Bool>();
			for (round in 0...2)
				for (leaf in leaves) {
					final json = render(leaf);
					require(json == plainRender(leaf), "leaf " + Std.string(leaf) + " in round " + round + ": " + json);
					if (round == 0)
						seen.set(json, true);
				}
			require([for (json in seen.keys()) json].length == leaves.length, "distinct leaf JSON for " + path);
		}

		// Same-named private declarations keep their own module paths.
		for (name in ["hidden", "score", "pair"]) {
			final firstShape = render(ledger.typeShape(fieldType(first, name)));
			final otherShape = render(ledger.typeShape(fieldType(other, name)));
			require(firstShape.indexOf('"path":"GoTypeShapeContractFixture.') != -1, name + " in the first fixture: " + firstShape);
			require(otherShape.indexOf('"path":"GoTypeShapeContractOtherFixture.') != -1, name + " in the other fixture: " + otherShape);
		}

		// Type parameters that print alike keep their own module paths.
		final firstPick = pickParameter(localPickType(first));
		final otherPick = pickParameter(localPickType(other));
		require(firstPick.printed == otherPick.printed, "both local type parameters print as " + firstPick.printed);
		require(render(ledger.typeShape(firstPick.type)) == render(GoTypeShape.TypeParameter("GoTypeShapeContractFixture.V")),
			"first local type parameter path");
		require(render(ledger.typeShape(otherPick.type)) == render(GoTypeShape.TypeParameter("GoTypeShapeContractOtherFixture.V")),
			"other local type parameter path");

		// The followed Iterator typedef still meets the recursion guard.
		require(render(ledger.typeShape(fieldType(first, "iterators"))).indexOf('"path":"recursive"') != -1, "nested Iterator recursion marker");

		// A monomorph is read on every conversion, never kept.
		final mono = Context.makeMonomorph();
		final arrayRef = switch (Context.getType("Array")) {
			case TInst(ref, _): ref;
			case _: Context.fatalError("GoTypeShapeContract: Array is not a class", Context.currentPos());
		};
		final array = TInst(arrayRef, [mono]);
		final callback = TFun([{name: "value", opt: false, t: mono}], array);
		final unbound = render(ledger.typeShape(callback));
		require(unbound.indexOf('"path":"monomorph"') != -1, "unbound monomorph: " + unbound);
		require(unbound == render(plainShape(callback)), "unbound monomorph matches the plain conversion");
		require(Context.unify(mono, Context.getType("Int")), "Haxe binds the monomorph");
		final bound = render(ledger.typeShape(callback));
		require(bound.indexOf('"path":"monomorph"') == -1 && bound.indexOf('"path":"StdTypes.Int"') != -1, "bound monomorph: " + bound);
		require(bound == render(plainShape(callback)), "bound monomorph matches the plain conversion");

		Sys.println("type shape contract passed");
	}

	/**
		The ledger's conversion before it kept declaration facts, kept as the
		oracle for this contract. Do not optimize it.
	**/
	static function plainShape(type:Type, ?trail:Array<String>, depth:Int = 0):GoTypeShape {
		if (type == null) {
			return GoTypeShape.UnknownShape(GoUnknownTypeShapeReason.Missing);
		}
		if (depth >= 64) {
			return GoTypeShape.UnknownShape(GoUnknownTypeShapeReason.DepthLimit);
		}
		var activeTrail = trail == null ? [] : trail;
		var id = TypeHelper.getUniqueId(type);
		if (activeTrail.indexOf(id) != -1) {
			return GoTypeShape.UnknownShape(GoUnknownTypeShapeReason.Recursive);
		}
		var nextTrail = activeTrail.copy();
		nextTrail.push(id);

		return switch (type) {
			case TMono(reference):
				var resolved = reference.get();
				resolved == null ? GoTypeShape.UnknownShape(GoUnknownTypeShapeReason.Monomorph) : plainShape(resolved, nextTrail, depth + 1);
			case TEnum(enumRef, parameters):
				plainNominalShape(GoTypeUsageTargetKind.Enum, TEnumDecl(enumRef).getPath(), parameters, activeTrail, depth);
			case TInst(classRef, parameters):
				var classType = classRef.get();
				switch (classType.kind) {
					case KTypeParameter(_):
						GoTypeShape.TypeParameter(TClassDecl(classRef).getPath());
					case _:
						plainNominalShape(GoTypeUsageTargetKind.Class, TClassDecl(classRef).getPath(), parameters, activeTrail, depth);
				}
			case TType(typeRef, parameters):
				final path = TTypeDecl(typeRef).getPath();
				if (path == "StdTypes.Iterator" && parameters.length == 1) {
					plainShape(Context.follow(type), nextTrail, depth + 1);
				} else {
					plainNominalShape(GoTypeUsageTargetKind.Typedef, path, parameters, activeTrail, depth);
				}
			case TFun(arguments, returnType):
				var outwardArguments = new Array<GoFunctionArgumentShape>();
				for (argument in arguments) {
					outwardArguments.push({
						name: argument.name,
						optional: argument.opt,
						shape: plainShape(argument.t, nextTrail, depth + 1)
					});
				}
				GoTypeShape.Function(GoImmutableList.fromArray(outwardArguments), plainShape(returnType, nextTrail, depth + 1));
			case TAnonymous(anonymousRef):
				var sourceFields = anonymousRef.get().fields.copy();
				sourceFields.sort((a, b) -> GoTypeUsageLedger.compareStrings(a.name, b.name));
				var outwardFields = new Array<GoAnonymousFieldShape>();
				for (field in sourceFields) {
					outwardFields.push({
						name: field.name,
						optional: field.meta != null && field.meta.has(":optional"),
						shape: plainShape(field.type, nextTrail, depth + 1)
					});
				}
				GoTypeShape.Anonymous(GoImmutableList.fromArray(outwardFields));
			case TDynamic(inner):
				GoTypeShape.DynamicShape(inner == null ? null : plainShape(inner, nextTrail, depth + 1));
			case TLazy(resolve):
				plainShape(resolve(), nextTrail, depth + 1);
			case TAbstract(abstractRef, parameters):
				var moduleType:ModuleType = TAbstract(abstractRef);
				plainNominalShape(GoTypeUsageTargetKind.Abstract, moduleType.getPath(), parameters, activeTrail, depth);
		};
	}

	/**
		The ledger's render before it kept leaf JSON, kept as the oracle for this
		contract. Do not optimize it.
	**/
	static function plainRender(shape:GoTypeShape):String {
		return switch (shape) {
			case Nominal(kind, path, parameters):
				plainObject(kind, path, "[" + [for (parameter in parameters) plainRender(parameter)].join(",") + "]", "[]", "null", "[]");
			case TypeParameter(path):
				plainObject("type_parameter", path, "[]", "[]", "null", "[]");
			case Function(arguments, returnType):
				final argumentJson = [
					for (argument in arguments)
						'{"name":"'
						+ GoTypeUsageLedger.jsonEscape(argument.name)
						+ '","optional":'
						+ argument.optional
						+ ',"shape":'
						+ plainRender(argument.shape)
						+ "}"];
				plainObject("function", "", "[]", "[" + argumentJson.join(",") + "]", plainRender(returnType), "[]");
			case Anonymous(fields):
				final fieldJson = [
					for (field in fields)
						'{"name":"'
						+ GoTypeUsageLedger.jsonEscape(field.name)
						+ '","optional":'
						+ field.optional
						+ ',"shape":'
						+ plainRender(field.shape)
						+ "}"];
				plainObject("anonymous", "", "[]", "[]", "null", "[" + fieldJson.join(",") + "]");
			case DynamicShape(inner):
				plainObject("dynamic", "", "[]", "[]", inner == null ? "null" : plainRender(inner), "[]");
			case UnknownShape(reason):
				plainObject("unknown", reason, "[]", "[]", "null", "[]");
		};
	}

	static function plainObject(kind:String, path:String, parameters:String, arguments:String, returnType:String, fields:String):String {
		return '{"kind":"' + GoTypeUsageLedger.jsonEscape(kind) + '","path":"' + GoTypeUsageLedger.jsonEscape(path) + '","parameters":' + parameters
			+ ',"arguments":' + arguments + ',"returnType":' + returnType + ',"fields":' + fields + "}";
	}

	static function plainNominalShape(kind:GoTypeUsageTargetKind, path:String, parameters:Array<Type>, trail:Array<String>, depth:Int):GoTypeShape {
		var outwardParameters = new Array<GoTypeShape>();
		for (parameter in parameters) {
			outwardParameters.push(plainShape(parameter, trail, depth + 1));
		}
		return GoTypeShape.Nominal(kind, path, GoImmutableList.fromArray(outwardParameters));
	}
	#end
}
