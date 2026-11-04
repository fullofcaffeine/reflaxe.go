package reflaxe.go.compiler;

#if macro
import haxe.macro.Context;
import haxe.macro.Type;
import reflaxe.go.ast.GoAST.GoExpr;
import reflaxe.go.ast.GoAST.GoDecl;
import reflaxe.go.ast.GoAST.GoTypeSwitchCase;
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
	var needsRecovery = false;
	var frozen = false;

	public function new(config:GoStructuralObjectConfig) {
		this.config = config;
	}

	/** Reset per-compilation admission plans; no runtime registry is introduced. */
	public function reset():Void {
		needsRecovery = false;
		frozen = false;
	}

	/** Typed consumers share raw structural storage; field reads own recovery. */
	public function recover(value:GoExpr, target:Type):GoExpr {
		return recoverRaw(value);
	}

	/** Recover a heterogeneous object without imposing a target member projection. */
	public function recoverRaw(value:GoExpr):GoExpr {
		if (frozen && !needsRecovery)
			Context.fatalError("Structural recovery discovered after freeze", Context.currentPos());
		needsRecovery = true;
		return GoExpr.GoCall(GoExpr.GoIdent("hx_structural_recover"), [value]);
	}

	public function requirementCount():Int
		return needsRecovery ? 1 : 0;

	public function freeze():Void
		frozen = true;

	/**
		Erased Haxe object transport does not re-check structural unification. Every
		reachable generated class retains its physical stored members, including
		fields outside the first narrow view. Missing fields remain absent.
	**/
	public function emitRecoveries(classes:Array<ClassType>):Array<GoDecl> {
		if (!needsRecovery)
			return [];
		var cases = new Array<GoTypeSwitchCase>();
		for (source in classes) {
			if (source.isExtern || source.isInterface)
				continue;
			for (type in Context.getModule(source.module))
				switch (type) {
					case TInst(ref, _) if (ref.get().name == source.name):
						cases.push({typeName: config.sourceGoType(type), body: [GoStmt.GoReturn(classView(GoExpr.GoIdent("source"), type, source))]});
					case _:
				}
		}
		var fallback = [
			GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("hxrt.AnonymousFromValue"), [GoExpr.GoIdent("value")]))
		];
		var body = cases.length == 0 ? fallback : [GoStmt.GoTypeSwitch(GoExpr.GoIdent("value"), "source", cases, fallback)];
		return [
			GoDecl.GoFuncDecl("hx_structural_recover", null, [{name: "value", typeName: "any"}], ["hxrt.Anonymous"], body)
		];
	}

	/** Admit generated class sources; anonymous width projections need no new view. */
	public function coerce(value:GoExpr, fromType:Type, toType:Type):Null<GoExpr> {
		var target = GoTypeMapper.anonymousFields(toType);
		var anonymousSource = GoTypeMapper.anonymousFields(fromType);
		if (target != null && anonymousSource != null)
			return value;
		if (target != null && config.sourceGoType(fromType) == "any")
			return recoverRaw(value);
		var source = switch (Context.follow(fromType)) {
			case TInst(ref, _): ref.get();
			case _: null;
		};
		if (target == null || source == null || source.isExtern)
			return null;
		return classView(value, fromType, source);
	}

	/** Read and write source declarations, independent of the consumer's shape. */
	function classView(value:GoExpr, fromType:Type, source:ClassType):GoExpr {
		var sourceName = config.freshTempName("hx_structural_source");
		var sourceExpr = GoExpr.GoIdent(sourceName);
		var reads = new Array<GoSwitchCase>();
		var writes = new Array<GoSwitchCase>();
		for (member in sourceMembers(source)) {
			if (source.isInterface)
				switch (member.kind) {
					case FMethod(MethNormal) | FMethod(MethInline):
					case _:
						continue;
				}
			var read = config.readMember(sourceExpr, source, member, member.type);
			reads.push({values: [GoExpr.GoStringLiteral(member.name)], body: [GoStmt.GoReturn(fieldValue(read))]});
			var write = config.writeMember(sourceExpr, source, member, member.type, GoExpr.GoIdent("incoming"));
			if (write != null)
				writes.push({values: [GoExpr.GoStringLiteral(member.name)], body: [write, GoStmt.GoReturn(GoExpr.GoBoolLiteral(true))]});
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

	/** Nearest declarations win while inherited source fields remain visible. */
	function sourceMembers(source:ClassType):Array<ClassField> {
		var fields = new Array<ClassField>();
		var seen = new Map<String, Bool>();
		function collect(type:ClassType):Void {
			for (field in type.fields.get()) {
				if (field.name == "new" || seen.exists(field.name))
					continue;
				seen.set(field.name, true);
				fields.push(field);
			}
			if (type.isInterface)
				for (parent in type.interfaces)
					collect(parent.t.get());
			if (type.superClass != null)
				collect(type.superClass.t.get());
		}
		collect(source);
		return fields;
	}

	function fieldValue(value:GoExpr):GoExpr {
		return GoExpr.GoCompositeLiteral("hxrt.AnonymousField", [
			GoCompositeElement.GoCompositeKeyValue(GoExpr.GoIdent("Value"), value),
			GoCompositeElement.GoCompositeKeyValue(GoExpr.GoIdent("Found"), GoExpr.GoBoolLiteral(true))
		]);
	}
}
#end
