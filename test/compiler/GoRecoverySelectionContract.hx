import reflaxe.go.ast.GoPackageName;
import reflaxe.go.ast.GoType;
import reflaxe.go.compiler.GoBuildConstraint;
import reflaxe.go.compiler.GoRecoveryAvailability;

/** Availability follows declaration ownership, not the first consumer of a type. */
class GoRecoverySelectionContract {
	static function require(value:Bool, message:String):Void {
		if (!value)
			throw message;
	}

	static function main():Void {
		var imports = new Map<String, String>();
		imports.set("runtime", "example/runtime");
		imports.set("first", "example/native");
		imports.set("second", "example/native");
		var owner = new GoRecoveryAvailability(alias -> imports.get(alias.value()));
		owner.notePackage("example/runtime", GoBuildConstraint.always());
		owner.noteGeneratedType("Node", "source-one", GoBuildConstraint.parse("alpha"));
		var firstRevision = owner.revision();
		owner.noteGeneratedType("Node", "source-one", GoBuildConstraint.parse("alpha"));
		require(owner.revision() == firstRevision, "identical ownership prevented convergence");
		var conflictingOwner = false;
		try
			owner.noteGeneratedType("Node", "source-one", GoBuildConstraint.parse("beta"))
		catch (_:haxe.Exception)
			conflictingOwner = true;
		require(conflictingOwner && owner.revision() == firstRevision, "a contradictory owner silently replaced its contract");
		owner.noteGeneratedType("Node", "source-two", GoBuildConstraint.parse("beta"));
		var availability = owner.conditionForType(GoType.parse("*runtime.Function[func(*Node) int]"));
		require(availability != null
			&& availability.key() == GoBuildConstraint.parse("alpha || beta").key(), "alternative owners were intersected");
		owner.noteNativeType("example/native", "Item", "native-binding", GoBuildConstraint.parse("gamma"));
		var left = owner.conditionForType(GoType.parse("*first.Item"));
		var right = owner.conditionForType(GoType.parse("*second.Item"));
		require(left != null && right != null && left.key() == right.key(), "import aliases split native ownership");
		owner.beginPass();
		require(owner.conditionForType(GoType.parse("*Pending")) == null, "unresolved ownership became common");
		var refused = false;
		try
			owner.freeze()
		catch (_:haxe.Exception)
			refused = true;
		require(refused, "unresolved ownership crossed freeze");
		owner.noteGeneratedType("Pending", "late-source", GoBuildConstraint.parse("delta"));
		owner.beginPass();
		var pending = owner.conditionForType(GoType.parse("*Pending"));
		require(pending != null && pending.key() == "delta", "late source ownership was lost");
		owner.noteHelper("recover_node", GoType.parse("func(any) *Node"));
		var captured = owner.capture(() -> {
			owner.requireHelper("recover_node");
			owner.requireType(GoType.parse("*first.Item"));
			return 42;
		});
		var complete = owner.conditionFor(captured.dependencies);
		require(captured.value == 42 && complete != null && complete.key() == GoBuildConstraint.parse("(alpha || beta) && gamma").key(),
			"nested helper presence was not included");
		owner.freeze();
		Sys.println("ok");
	}
}
