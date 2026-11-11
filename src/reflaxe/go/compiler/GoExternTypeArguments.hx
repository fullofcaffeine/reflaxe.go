package reflaxe.go.compiler;

#if macro
import haxe.macro.Context;
import haxe.macro.Expr.Position;
import haxe.macro.Type;
import haxe.macro.TypeTools;

/**
	What: Recovers concrete native function type arguments from a typed Haxe use.
	Why: Go cannot infer parameters that occur only in a function's result, and
	ordinary Haxe generic carriers erase those parameters during storage lowering.
	How: Substitute fresh monomorphs into the declaration and unify its complete
	signature with the already checked use. Preserve Null<T> while resolving each
	monomorph; following away its wrapper would change callback storage. Never mutate
	declaration parameters or
	silently substitute an untyped native argument when inference is incomplete.
**/
function resolve(field:ClassField, appliedType:Type, pos:Position):Array<Type> {
	final arguments = [for (_ in field.params) Context.makeMonomorph()];
	final signature = TypeTools.applyTypeParameters(field.type, field.params, arguments);
	// Prefer Go inference when input arguments determine every parameter. Native
	// callback adapters own their physical ABI, including nullable storage.
	final voidType = Context.getType("Void");
	switch ([Context.follow(signature), Context.follow(appliedType)]) {
		case [TFun(inputs, _), TFun(appliedInputs, _)]:
			if (!Context.unify(TFun(inputs, voidType), TFun(appliedInputs, voidType))) {
				Context.fatalError("Cannot resolve native generic inputs for " + field.name, pos);
			}
		case _:
			Context.fatalError("Native generic declaration must be a function: " + field.name, pos);
	}
	var inputsResolveAll = true;
	for (argument in arguments) {
		switch (argument) {
			case TMono(reference) if (reference.get() == null):
				inputsResolveAll = false;
			case _:
		}
		if (!inputsResolveAll) {
			break;
		}
	}
	if (inputsResolveAll) {
		return [];
	}
	if (!Context.unify(signature, appliedType)) {
		Context.fatalError("Cannot resolve native generic arguments for " + field.name, pos);
	}
	return [
		for (argument in arguments)
			switch (argument) {
				case TMono(reference):
					final resolved = reference.get();
					if (resolved == null) {
						Context.fatalError("Native generic argument remains unresolved for " + field.name, pos);
					}
					resolved;
				case resolved:
					resolved;
			}
	];
}
#end
