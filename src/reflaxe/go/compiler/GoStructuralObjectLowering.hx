package reflaxe.go.compiler;

#if macro
import haxe.macro.Context;
import haxe.macro.Type;
import reflaxe.go.ast.GoAST.GoExpr;
import reflaxe.go.ast.GoAST.GoDecl;
import reflaxe.go.ast.GoAST.GoTypeSwitchCase;
import haxe.macro.TypeTools;
import reflaxe.go.ast.GoAST.GoStmt;
import reflaxe.go.ast.GoAST.GoSwitchCase;
import reflaxe.go.ast.GoCompositeElement;

/** Typed compiler services used to construct source-retaining structural views. */
private typedef GoStructuralObjectConfig = {
	final freshTempName:String->String;
	final sourceGoType:Type->String;
	final readMember:(GoExpr, ClassType, ClassField, Type) -> GoExpr;
	final writeMember:(GoExpr, ClassType, ClassField, Type, GoExpr) -> Null<GoStmt>;
	final noteSource:ClassType->Void;
	final readAnonymous:(GoExpr, String, Type) -> GoExpr;
	final adaptValue:(GoExpr, Type, Type) -> GoExpr;
	final recoverValue:(GoExpr, Type) -> GoExpr;
}

/**
	What: Builds structural object views without copying source fields.
	Why: Haxe structural unification retains object identity and live member access.
	A fresh map of callbacks loses both, including after a narrower projection.
	How: The source is captured once and compiler-owned member plans supply typed
	readers and writers. Runtime storage owns only lifetime and heterogeneous lookup.
**/
class GoStructuralObjectLowering {
	final config:GoStructuralObjectConfig;
	final recoveries:Array<{key:String, target:Type, symbol:String}> = [];

	public function new(config:GoStructuralObjectConfig) {
		this.config = config;
	}

	/** Reset per-compilation admission plans; no runtime registry is introduced. */
	public function reset():Void {
		recoveries.resize(0);
	}

	/** Recover erased values through a finite compiler-owned target-shape resolver. */
	public function recover(value:GoExpr, target:Type):GoExpr {
		var key = TypeTools.toString(target);
		var symbol:Null<String> = null;
		for (plan in recoveries)
			if (plan.key == key) {
				symbol = plan.symbol;
				break;
			}
		if (symbol == null) {
			symbol = "hx_structural_recover_" + recoveries.length;
			recoveries.push({key: key, target: target, symbol: symbol});
		}
		return GoExpr.GoCall(GoExpr.GoIdent(symbol), [value]);
	}

	/**
		Emit only reachable generated carriers which Haxe can unify with the requested
		shape. Their fields use the same typed live plan as direct structural coercion;
		already structural values and null retain the runtime storage recovery path.
	**/
	public function emitRecoveries(classes:Array<ClassType>):Array<GoDecl> {
		var declarations = new Array<GoDecl>();
		var sources = new Array<Type>();
		for (source in classes) {
			if (source.isExtern || source.isInterface)
				continue;
			for (type in Context.getModule(source.module))
				switch (type) {
					case TInst(ref, _) if (ref.get().name == source.name):
						sources.push(type);
					case _:
				}
		}
		var index = 0;
		while (index < recoveries.length) {
			var plan = recoveries[index++];
			var cases = new Array<GoTypeSwitchCase>();
			for (source in sources) {
				if (!Context.unify(source, plan.target))
					continue;
				var view = coerce(GoExpr.GoIdent("source"), source, plan.target);
				if (view != null)
					cases.push({typeName: config.sourceGoType(source), body: [GoStmt.GoReturn(view)]});
			}
			var fallback = [
				GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("hxrt.AnonymousFromValue"), [GoExpr.GoIdent("value")]))
			];
			var body = cases.length == 0 ? fallback : [GoStmt.GoTypeSwitch(GoExpr.GoIdent("value"), "source", cases, fallback)];
			declarations.push(GoDecl.GoFuncDecl(plan.symbol, null, [{name: "value", typeName: "any"}], ["hxrt.Anonymous"], body));
		}
		return declarations;
	}

	/** Admit generated class sources; anonymous width projections need no new view. */
	public function coerce(value:GoExpr, fromType:Type, toType:Type):Null<GoExpr> {
		var target = GoTypeMapper.anonymousFields(toType);
		var anonymousSource = GoTypeMapper.anonymousFields(fromType);
		if (target != null && anonymousSource != null)
			return adaptAnonymous(value, anonymousSource, target);
		var source = switch (Context.follow(fromType)) {
			case TInst(ref, _): ref.get();
			case _: null;
		};
		if (target == null || source == null || source.isExtern)
			return null;
		var sourceName = config.freshTempName("hx_structural_source");
		var sourceExpr = GoExpr.GoIdent(sourceName);
		var reads = new Array<GoSwitchCase>();
		var writes = new Array<GoSwitchCase>();
		for (destination in target) {
			var member = findMember(source, destination.name);
			if (member == null)
				return null;
			if (source.isInterface)
				switch (member.kind) {
					case FMethod(MethNormal) | FMethod(MethInline):
					case _:
						return null;
				}
			var read = config.readMember(sourceExpr, source, member, destination.type);
			reads.push({values: [GoExpr.GoStringLiteral(destination.name)], body: [GoStmt.GoReturn(fieldValue(read))]});
			var write = config.writeMember(sourceExpr, source, member, destination.type, GoExpr.GoIdent("incoming"));
			if (write != null)
				writes.push({values: [GoExpr.GoStringLiteral(destination.name)], body: [write, GoStmt.GoReturn(GoExpr.GoBoolLiteral(true))]});
		}
		config.noteSource(source);
		var read = GoExpr.GoFuncLiteral([{name: "name", typeName: "string"}], ["hxrt.AnonymousField"], [
			GoStmt.GoSwitch(GoExpr.GoIdent("name"), reads, null),
			GoStmt.GoReturn(GoExpr.GoCompositeLiteral("hxrt.AnonymousField", []))
		]);
		var write = GoExpr.GoFuncLiteral([{name: "name", typeName: "string"}, {name: "incoming", typeName: "any"}], ["bool"], [
			GoStmt.GoSwitch(GoExpr.GoIdent("name"), writes, null),
			GoStmt.GoReturn(GoExpr.GoBoolLiteral(false))
		]);
		return GoExpr.GoCall(GoExpr.GoFuncLiteral([{name: sourceName, typeName: config.sourceGoType(fromType)}], ["hxrt.Anonymous"], [
			GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("hxrt.NewAnonymousView"), [sourceExpr, read, write]))
		]), [value]);
	}

	/**
		Change member signatures through live reads, retaining the original source.
		Width-only projections reuse storage. Changed callback signatures adapt each
		read value rather than retaining a snapshot of a mutable source slot.
	**/
	function adaptAnonymous(value:GoExpr, from:Array<ClassField>, to:Array<ClassField>):GoExpr {
		var changed = new Array<{source:ClassField, target:ClassField}>();
		for (destination in to) {
			for (source in from) {
				if (source.name == destination.name && config.sourceGoType(source.type) != config.sourceGoType(destination.type)) {
					changed.push({source: source, target: destination});
					break;
				}
			}
		}
		if (changed.length == 0)
			return value;
		var sourceName = config.freshTempName("hx_structural_origin");
		var source = GoExpr.GoIdent(sourceName);
		var reads = new Array<GoSwitchCase>();
		var writes = new Array<GoSwitchCase>();
		for (member in changed) {
			var key = GoExpr.GoStringLiteral(member.target.name);
			var read = config.readAnonymous(source, member.source.name, member.source.type);
			read = config.adaptValue(read, member.source.type, member.target.type);
			reads.push({values: [key], body: [GoStmt.GoReturn(fieldValue(read))]});
			var incoming = config.recoverValue(GoExpr.GoIdent("incoming"), member.target.type);
			incoming = config.adaptValue(incoming, member.target.type, member.source.type);
			writes.push({
				values: [key],
				body: [
					GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoIdent("hxrt.AnonymousSet"), [source, key, incoming])),
					GoStmt.GoReturn(GoExpr.GoBoolLiteral(true))
				]
			});
		}
		var reader = GoExpr.GoFuncLiteral([{name: "name", typeName: "string"}], ["hxrt.AnonymousField"], [
			GoStmt.GoSwitch(GoExpr.GoIdent("name"), reads, null),
			GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("hxrt.AnonymousLookup"), [source, GoExpr.GoIdent("name")]))
		]);
		var writer = GoExpr.GoFuncLiteral([{name: "name", typeName: "string"}, {name: "incoming", typeName: "any"}], ["bool"], [
			GoStmt.GoSwitch(GoExpr.GoIdent("name"), writes, null),
			GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoIdent("hxrt.AnonymousSet"), [source, GoExpr.GoIdent("name"), GoExpr.GoIdent("incoming")])),
			GoStmt.GoReturn(GoExpr.GoBoolLiteral(true))
		]);
		return GoExpr.GoCall(GoExpr.GoFuncLiteral([{name: sourceName, typeName: "hxrt.Anonymous"}], ["hxrt.Anonymous"], [
			GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("hxrt.NewAnonymousView"), [source, reader, writer]))
		]), [value]);
	}

	function findMember(source:ClassType, name:String):Null<ClassField> {
		for (field in source.fields.get())
			if (field.name == name)
				return field;
		if (source.isInterface)
			for (parent in source.interfaces) {
				var field = findMember(parent.t.get(), name);
				if (field != null)
					return field;
			}
		return source.superClass == null ? null : findMember(source.superClass.t.get(), name);
	}

	function fieldValue(value:GoExpr):GoExpr {
		return GoExpr.GoCompositeLiteral("hxrt.AnonymousField", [
			GoCompositeElement.GoCompositeKeyValue(GoExpr.GoIdent("Value"), value),
			GoCompositeElement.GoCompositeKeyValue(GoExpr.GoIdent("Found"), GoExpr.GoBoolLiteral(true))
		]);
	}
}
#end
