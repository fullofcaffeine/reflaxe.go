package reflaxe.go.compiler;

#if macro
import haxe.macro.Context;
import reflaxe.go.ast.GoAST.GoDecl;
import reflaxe.go.ast.GoAST.GoExpr;
import reflaxe.go.ast.GoAST.GoParam;
import reflaxe.go.ast.GoAST.GoStmt;
import reflaxe.go.ast.GoAST.GoTypeSwitchCase;
import reflaxe.go.ast.GoBuiltinType;
import reflaxe.go.ast.GoType;
import reflaxe.go.compiler.GoRecoveryDispatch.GoRecoveryCasePlan;
import reflaxe.go.compiler.GoRecoveryDispatch.GoGuardedRecoveryDeclaration;
import reflaxe.go.ast.GoType.GoFunctionSignature;

/** Actual emitted storage, rather than an applied Haxe expression's type. */
private typedef CallableRepresentation = {
	final carrier:GoType;
	final signature:GoFunctionSignature;
}

/** A demand is interned before its nested argument/result conversions expand. */
private typedef CallableRecovery = {
	final representation:CallableRepresentation;
	final symbol:String;
}

/**
	What: Plans typed recovery of callable values read from heterogeneous storage.
	Why: Different Function instantiations cannot be recovered by a Go assertion,
	and projections must start at the normalized source to avoid wrapper chains.
	How: Intern exact emitted representations, record reachable source producers,
	and emit only demanded target helpers with finite typed source dispatch.
	Anonymous storage is opaque here: its fields recover when accessed.
**/
class GoValueRecovery {
	final availability:GoRecoveryAvailability;
	final dispatcher:GoRecoveryDispatch;
	final recoverStored:(GoExpr, GoType) -> GoExpr;
	final sources:Array<CallableRepresentation> = [];
	final demands:Array<CallableRecovery> = [];
	final namedSources:Map<String, Bool> = [];
	final declarations:Map<String, GoType> = [];
	var frozen = false;
	var needsInvocation = false;

	public function new(availability:GoRecoveryAvailability, recoverStored:(GoExpr, GoType) -> GoExpr) {
		this.availability = availability;
		this.dispatcher = new GoRecoveryDispatch(availability);
		this.recoverStored = recoverStored;
	}

	public function reset():Void {
		sources.resize(0);
		demands.resize(0);
		frozen = false;
		needsInvocation = false;
		namedSources.clear();
		declarations.clear();
	}

	/** Register the physical signature from an emitted declaration, including optional ABI. */
	public function noteDeclarations(values:Array<GoDecl>):Void {
		for (value in values)
			switch (value) {
				case GoFuncDecl(name, null, params, results, _):
					var carrier = carrierType([for (param in params) param.typeName], results);
					declarations.set(name, carrier);
					if (namedSources.exists(name))
						noteSource(carrier);
				case _:
			}
	}

	/** Named producers can precede their declaration in source-module order. */
	public function noteNamedSource(name:String):Void {
		namedSources.set(name, true);
		var carrier = declarations.get(name);
		if (carrier != null)
			noteSource(carrier);
	}

	/** A literal producer carries its exact emitted signature in the typed AST. */
	public function noteLiteralSource(value:GoExpr):Void {
		switch (value) {
			case GoFuncLiteral(params, results, _):
				noteSource(carrierType([for (param in params) param.typeName], results));
			case _:
				Context.fatalError("Callable producer requires an emitted function literal", Context.currentPos());
		}
	}

	static function carrierType(params:Array<GoType>, results:Array<GoType>):GoType {
		return GoType.pointer(GoType.generic(GoType.qualified("hxrt", "Function"), [GoType.functionType(params, results)]));
	}

	/** Record a real normalized callable producer; projections are not new origins. */
	public function noteSource(carrier:GoType):Void {
		for (source in sources)
			if (source.carrier.sameRepresentation(carrier))
				return;
		if (frozen)
			Context.fatalError("Callable source discovered after recovery freeze", Context.currentPos());
		sources.push(representation(carrier));
	}

	/** Demand a target before discovering child conversions, preventing recursion. */
	public function recover(value:GoExpr, target:GoType):GoExpr {
		for (demand in demands)
			if (demand.representation.carrier.sameRepresentation(target))
				return demandCall(demand, value);
		if (frozen)
			Context.fatalError("Callable recovery discovered after recovery freeze", Context.currentPos());
		var demand:CallableRecovery = {representation: representation(target), symbol: "hx_recover_callable_" + demands.length};
		demands.push(demand);
		return demandCall(demand, value);
	}

	function demandCall(demand:CallableRecovery, value:GoExpr):GoExpr {
		availability.noteHelper(demand.symbol, GoType.functionType([GoType.parse("any")], [demand.representation.carrier]));
		availability.requireHelper(demand.symbol);
		return GoExpr.GoCall(GoExpr.GoIdent(demand.symbol), [value]);
	}

	/** Known crossings still recover from the normalized origin when signatures differ. */
	public function convert(value:GoExpr, source:GoType, target:GoType):GoExpr {
		if (source.sameRepresentation(target) || target.isBuiltin(GoBuiltinType.AnyType))
			return value;
		if (target.functionCarrierSignature() != null)
			return recover(value, target);
		if (source.isBuiltin(GoBuiltinType.Int) && target.isBuiltin(GoBuiltinType.Float64))
			return GoExpr.GoCall(GoExpr.GoIdent("float64"), [value]);
		return recoverStored(value, target);
	}

	/** Closed compiler dispatch preserves source defaults/results for erased calls. */
	public function requireInvocation():String {
		if (frozen && !needsInvocation)
			Context.fatalError("Erased invocation discovered after freeze", Context.currentPos());
		needsInvocation = true;
		return "reflaxe__go___internal__CompilerReflect_callMethod";
	}

	public function invoke(value:GoExpr, arguments:GoExpr):GoExpr {
		return GoExpr.GoCall(GoExpr.GoIdent(requireInvocation()), [value, arguments]);
	}

	/** Exposes a bounded discovery measure to the compiler's dependency work loop. */
	public function requirementCount():Int {
		return sources.length + demands.length + (needsInvocation ? 1 : 0);
	}

	/**
		Build plans until nested demands stop growing. The caller must also close its
		class, metadata and runtime dependency queues before freezing emission.
	 */
	public function discover():Array<GoGuardedRecoveryDeclaration> {
		var declarations = new Array<GoGuardedRecoveryDeclaration>();
		if (needsInvocation)
			declarations = declarations.concat(lowerInvocation());
		var index = 0;
		while (index < demands.length)
			declarations = declarations.concat(lowerRecovery(demands[index++]));
		return declarations;
	}

	/** Freeze after the compiler-wide discovery loop reaches a fixed point. */
	public function freeze():Void {
		for (name in namedSources.keys())
			if (!declarations.exists(name))
				Context.fatalError("Missing emitted callable declaration: " + name, Context.currentPos());
		frozen = true;
	}

	function representation(carrier:GoType):CallableRepresentation {
		var signature = carrier.functionCarrierSignature();
		if (signature == null)
			Context.fatalError("Expected an emitted Haxe function carrier", Context.currentPos());
		return {carrier: carrier, signature: signature};
	}

	function lowerRecovery(demand:CallableRecovery):Array<GoGuardedRecoveryDeclaration> {
		var target = demand.representation;
		var exact:GoTypeSwitchCase = {typeName: target.carrier, body: [GoStmt.GoReturn(GoExpr.GoIdent("exact"))]};
		var cases = new Array<GoRecoveryCasePlan>();
		for (source in sources) {
			if (source.carrier.sameRepresentation(target.carrier)) {
				cases.push({sourceType: source.carrier, body: [GoStmt.GoReturn(GoExpr.GoIdent("source"))], dependencies: []});
			} else if (admitsArity(source.signature, target.signature)
				&& source.signature.results.length <= 1
				&& target.signature.results.length <= 1) {
				var captured = availability.capture(() -> project(GoExpr.GoIdent("source"), source, target));
				cases.push({sourceType: source.carrier, body: [GoStmt.GoReturn(captured.value)], dependencies: captured.dependencies});
			}
		}
		var failure = [
			GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoIdent("hxrt.Throw"), [
				GoExpr.GoCall(GoExpr.GoIdent("hxrt.StringFromLiteral"), [GoExpr.GoStringLiteral("Invalid function value")])
			])),
			GoStmt.GoReturn(GoExpr.GoNil)
		];
		var body = [
			GoStmt.GoIf(GoExpr.GoBinary("==", GoExpr.GoIdent("value"), GoExpr.GoNil), [GoStmt.GoReturn(GoExpr.GoNil)], null),
			GoStmt.GoTypeSwitch(GoExpr.GoIdent("value"), "exact", [exact], null),
			GoStmt.GoVarDecl("origin", "any", GoExpr.GoCall(GoExpr.GoIdent("hxrt.FunctionOrigin"), [GoExpr.GoIdent("value")]), true),
			GoStmt.GoIf(GoExpr.GoBinary("==", GoExpr.GoIdent("origin"), GoExpr.GoNil), [GoStmt.GoReturn(GoExpr.GoNil)], null)
		];
		// Null-only programs have no producer cases and must not bind an unused variable.
		var plan = dispatcher.lower({
			symbol: demand.symbol,
			params: [{name: "value", typeName: "any"}],
			result: target.carrier,
			prefix: body,
			input: GoExpr.GoIdent("origin"),
			cases: cases,
			fallback: failure
		});
		return plan == null ? [] : plan;
	}

	/** Explicit reflection and Dynamic calls share typed argument recovery. */
	function lowerInvocation():Array<GoGuardedRecoveryDeclaration> {
		var cases = new Array<GoRecoveryCasePlan>();
		var arguments = GoExpr.GoIdent("arguments");
		for (source in sources) {
			if (source.signature.results.length > 1)
				continue;
			var captured = availability.capture(() -> {
				var body = new Array<GoStmt>();
				var actuals = new Array<GoExpr>();
				var variadic = false;
				for (index in 0...source.signature.params.length) {
					var parameter = source.signature.params[index];
					var element = parameter.variadicElement();
					var name = "argument_" + index;
					if (element != null) {
						variadic = true;
						body.push(GoStmt.GoVarDecl(name, GoType.slice(element), GoExpr.GoNil, false));
						body.push(GoStmt.GoRangeStmt("index", "argument", arguments, true, [
							GoStmt.GoIf(GoExpr.GoBinary(">=", GoExpr.GoIdent("index"), GoExpr.GoIntLiteral(index)), [
								GoStmt.GoAssign(GoExpr.GoIdent(name),
									GoExpr.GoCall(GoExpr.GoIdent("append"), [GoExpr.GoIdent(name), recoverStored(GoExpr.GoIdent("argument"), element)]))
							], null)
						]));
						actuals.push(GoExpr.GoIdent(name));
					} else {
						body.push(GoStmt.GoVarDecl(name, "any", GoExpr.GoNil, false));
						body.push(GoStmt.GoIf(GoExpr.GoBinary("<", GoExpr.GoIntLiteral(index), GoExpr.GoCall(GoExpr.GoIdent("len"), [arguments])), [
							GoStmt.GoAssign(GoExpr.GoIdent(name), GoExpr.GoIndex(arguments, GoExpr.GoIntLiteral(index)))
						], null));
						actuals.push(recoverStored(GoExpr.GoIdent(name), parameter));
					}
				}
				var callee = GoExpr.GoCall(GoExpr.GoIdent("hxrt.FunctionCall"), [GoExpr.GoIdent("source")]);
				var call = variadic ? GoExpr.GoVariadicCall(callee, actuals) : GoExpr.GoCall(callee, actuals);
				if (source.signature.results.length == 0) {
					body.push(GoStmt.GoExprStmt(call));
					body.push(GoStmt.GoReturn(GoExpr.GoNil));
				} else
					body.push(GoStmt.GoReturn(call));
				return body;
			});
			cases.push({sourceType: source.carrier, body: captured.value, dependencies: captured.dependencies});
		}
		var fallback = [
			GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("hxrt.ReflectCallMethod"), [GoExpr.GoIdent("value"), arguments]))
		];
		var body = [
			GoStmt.GoVarDecl("origin", "any", GoExpr.GoCall(GoExpr.GoIdent("hxrt.FunctionOrigin"), [GoExpr.GoIdent("value")]), true),
			GoStmt.GoIf(GoExpr.GoBinary("==", GoExpr.GoIdent("origin"), GoExpr.GoNil), [
				GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoIdent("hxrt.Throw"), [
					GoExpr.GoCall(GoExpr.GoIdent("hxrt.StringFromLiteral"), [GoExpr.GoStringLiteral("Cannot call null")])
				])),
				GoStmt.GoReturn(GoExpr.GoNil)
			], null)
		];
		var plan = dispatcher.lower({
			symbol: requireInvocation(),
			params: [{name: "value", typeName: "any"}, {name: "arguments", typeName: "[]any"}],
			result: GoType.parse("any"),
			prefix: body,
			input: GoExpr.GoIdent("origin"),
			cases: cases,
			fallback: fallback
		});
		return plan == null ? [] : plan;
	}

	function admitsArity(source:GoFunctionSignature, target:GoFunctionSignature):Bool {
		for (parameter in target.params)
			if (parameter.variadicElement() != null)
				return false;
		if (source.params.length > 0 && source.params[source.params.length - 1].variadicElement() != null)
			return target.params.length >= source.params.length - 1;
		return source.params.length == target.params.length;
	}

	function project(sourceValue:GoExpr, source:CallableRepresentation, target:CallableRepresentation):GoExpr {
		var params = new Array<GoParam>();
		var arguments = new Array<GoExpr>();
		for (index in 0...target.signature.params.length) {
			var name = "arg_" + index;
			params.push({name: name, typeName: target.signature.params[index]});
			var sourceIndex = index < source.signature.params.length ? index : source.signature.params.length - 1;
			var sourceParameter = source.signature.params[sourceIndex];
			var variadic = sourceParameter.variadicElement();
			arguments.push(convert(GoExpr.GoIdent(name), target.signature.params[index], variadic == null ? sourceParameter : variadic));
		}
		var call = GoExpr.GoCall(GoExpr.GoCall(GoExpr.GoIdent("hxrt.FunctionCall"), [sourceValue]), arguments);
		var body = new Array<GoStmt>();
		if (target.signature.results.length == 0) {
			body.push(GoStmt.GoExprStmt(call));
		} else if (source.signature.results.length == 0) {
			body.push(GoStmt.GoExprStmt(call));
			body.push(GoStmt.GoReturn(recoverStored(GoExpr.GoNil, target.signature.results[0])));
		} else {
			body.push(GoStmt.GoReturn(convert(call, source.signature.results[0], target.signature.results[0])));
		}
		return GoExpr.GoCall(GoExpr.GoIdent("hxrt.ProjectFunction"), [sourceValue, GoExpr.GoFuncLiteral(params, target.signature.results, body)]);
	}
}
#end
