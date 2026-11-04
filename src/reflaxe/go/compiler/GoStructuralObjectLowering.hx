package reflaxe.go.compiler;

#if macro
import haxe.macro.Context;
import haxe.macro.Type;
import reflaxe.go.ast.GoAST.GoExpr;
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

	public function new(config:GoStructuralObjectConfig) {
		this.config = config;
	}

	/** Admit generated class sources; anonymous width projections need no new view. */
	public function coerce(value:GoExpr, fromType:Type, toType:Type):Null<GoExpr> {
		var target = switch (Context.follow(toType)) {
			case TAnonymous(ref): ref.get();
			case _: null;
		};
		var source = switch (Context.follow(fromType)) {
			case TInst(ref, _): ref.get();
			case _: null;
		};
		if (target == null || source == null || source.isExtern || source.isInterface)
			return null;
		var sourceName = config.freshTempName("hx_structural_source");
		var sourceExpr = GoExpr.GoIdent(sourceName);
		var reads = new Array<GoSwitchCase>();
		var writes = new Array<GoSwitchCase>();
		for (destination in target.fields) {
			var member = findMember(source, destination.name);
			if (member == null)
				return null;
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

	function findMember(source:ClassType, name:String):Null<ClassField> {
		for (field in source.fields.get())
			if (field.name == name)
				return field;
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
