#if macro
import haxe.macro.Context;
import reflaxe.go.ast.GoAST.GoExpr;
import reflaxe.go.ast.GoAST.GoDecl;
import reflaxe.go.ast.GoASTPrinter;
import reflaxe.go.ast.GoType;
import reflaxe.go.compiler.GoRecoveryDispatch.GoGuardedRecoveryDeclaration;
import reflaxe.go.compiler.GoValueRecovery;
import reflaxe.go.compiler.GoRecoveryAvailability;
import reflaxe.go.compiler.GoBuildConstraint;
#end

/** Compiler-side discovery contracts; generated runtime regressions prove behavior. */
class GoValueRecoveryContract {
	#if macro
	static function newAvailability():GoRecoveryAvailability {
		var availability = new GoRecoveryAvailability(alias -> alias.value() == "hxrt" ? "runtime" : null);
		availability.notePackage("runtime", GoBuildConstraint.always());
		return availability;
	}

	public static function check():Void {
		var owner = new GoValueRecovery(newAvailability(), (value, target) -> GoExpr.GoTypeAssert(value, target));
		var source = GoType.parse("*hxrt.Function[func(int) int]");
		var target = GoType.parse("*hxrt.Function[func(any) any]");
		owner.noteSource(source);
		owner.noteSource(GoType.parse("*hxrt.Function[func(int) int]"));
		owner.recover(GoExpr.GoIdent("input"), target);
		owner.recover(GoExpr.GoIdent("other"), target);
		if (owner.requirementCount() != 2)
			Context.fatalError("Duplicate callable requirements", Context.currentPos());
		var declarations = owner.discover();
		if (declarations.length != 1)
			Context.fatalError("Duplicate recovery helper", Context.currentPos());
		owner.freeze();
		owner.recover(GoExpr.GoIdent("again"), target);
		if (owner.discover().length != 1)
			Context.fatalError("Frozen recovery plan changed", Context.currentPos());

		// The named source is observed before its module declaration. Recovering
		// the signature from applied Haxe typing would lose this physical ABI.
		var named = new GoValueRecovery(newAvailability(), (value, target) -> GoExpr.GoTypeAssert(value, target));
		named.noteNamedSource("optional_source");
		named.noteDeclarations([
			GoDecl.GoFuncDecl("optional_source", null, [{name: "arg", typeName: "int"}], ["int"], [])
		]);
		named.recover(GoExpr.GoIdent("stored"), target);
		if (named.requirementCount() != 2)
			Context.fatalError("Forward named source was lost", Context.currentPos());
		named.discover();
		named.freeze();

		// Recovering a callback argument discovers a second target while the first
		// helper is being built. Both demands must exist before emission freezes.
		var nested = new GoValueRecovery(newAvailability(), (value, target) -> GoExpr.GoTypeAssert(value, target));
		nested.noteSource(GoType.parse("*hxrt.Function[func(*hxrt.Function[func(int) int]) int]"));
		nested.recover(GoExpr.GoIdent("stored"), GoType.parse("*hxrt.Function[func(*hxrt.Function[func(any) any]) any]"));
		if (nested.discover().length != 2)
			Context.fatalError("Nested recovery demand was not closed", Context.currentPos());
		nested.freeze();
		if (nested.discover().length != 2)
			Context.fatalError("Frozen nested recovery changed", Context.currentPos());

		// Helper names and declaration order depend on the demanded set, not on the
		// order in which modules first demanded each target.
		var first = GoType.parse("*hxrt.Function[func(int) int]");
		var second = GoType.parse("*hxrt.Function[func(any) any]");
		var forward = new GoValueRecovery(newAvailability(), (value, target) -> GoExpr.GoTypeAssert(value, target));
		var backward = new GoValueRecovery(newAvailability(), (value, target) -> GoExpr.GoTypeAssert(value, target));
		var forwardCalls = [
			forward.recover(GoExpr.GoIdent("a"), first),
			forward.recover(GoExpr.GoIdent("b"), second)
		];
		var backwardCalls = [
			backward.recover(GoExpr.GoIdent("b"), second),
			backward.recover(GoExpr.GoIdent("a"), first)
		];
		if (Std.string(forwardCalls[0]) != Std.string(backwardCalls[1]) || Std.string(forwardCalls[1]) != Std.string(backwardCalls[0]))
			Context.fatalError("Callable helper name depended on first use", Context.currentPos());
		if (declarationNames(forward.discover()).join(",") != declarationNames(backward.discover()).join(","))
			Context.fatalError("Callable helper placement depended on first use", Context.currentPos());

		// Alias spellings of one representation share one helper.
		var aliases = new GoValueRecovery(newAvailability(), (value, target) -> GoExpr.GoTypeAssert(value, target));
		aliases.recover(GoExpr.GoIdent("a"), GoType.parse("*hxrt.Function[func(byte) int]"));
		aliases.recover(GoExpr.GoIdent("b"), GoType.parse("*hxrt.Function[func(uint8) int]"));
		if (aliases.requirementCount() != 1)
			Context.fatalError("Alias spellings split one callable helper", Context.currentPos());

		checkConfirmingPassReuse();
	}

	static function fail(message:String):Void {
		Context.fatalError("GoValueRecoveryContract: " + message, Context.currentPos());
	}

	/** Prints each helper with its file selection; `Std.string` would cut deep trees short. */
	static function render(values:Array<GoGuardedRecoveryDeclaration>):String {
		final parts = new Array<String>();
		for (value in values) {
			final source = GoASTPrinter.printFile({packageName: "main", imports: [], decls: [value.declaration]});
			parts.push("//go:build " + value.selection.key() + "\n" + source);
		}
		return parts.join("\n");
	}

	/**
		Why
		The compiler's recovery loop runs value discovery once per pass and stops
		when a pass discovers nothing new. The confirming pass may reuse the
		previous pass's helpers only when lowering them again would build exactly
		the same helpers.

		What
		- Unchanged inputs: the next discovery reuses the helpers and asks the
		  compiler to recover no stored value.
		- Between two passes, a new source, a new demand, the erased-invocation
		  flag, a new declaration owner or a changed compiler input: the next
		  discovery lowers the helpers again.
		- During one discovery, a later helper adds a source, sets the invocation
		  flag, or adds an owner of a declaration that an earlier helper read: that
		  result is stale, so the next discovery lowers the helpers again.
		- An owner that never arrives: the confirming discovery still reports it to
		  `freeze`.
		- Without a compiler input revision, discovery never reuses helpers.

		How
		Each changed result must match a discovery by a fresh owner that knew the
		final facts from the start, and must differ from the result before the
		change, so the fixture really needs the extra pass. Results are compared as
		printed Go source with each file selection. The stored-value callback
		changes facts during discovery, as the compiler's callback can.
	**/
	static function checkConfirmingPassReuse():Void {
		final intSource = GoType.parse("*hxrt.Function[func(int) int]");
		final stringSource = GoType.parse("*hxrt.Function[func(string) int]");
		final localSource = GoType.parse("*hxrt.Function[func(*Local) int]");
		final narrowTarget = GoType.parse("*hxrt.Function[func(any) int]");
		final wideTarget = GoType.parse("*hxrt.Function[func(any) any]");

		var calls = 0;
		var external = "compiler-1";
		final steady = new GoValueRecovery(newAvailability(), (value, type) -> {
			calls++;
			return GoExpr.GoTypeAssert(value, type);
		}, null, () -> external);
		steady.noteSource(intSource);
		steady.recover(GoExpr.GoIdent("stored"), wideTarget);
		final first = render(steady.discover());
		final firstCalls = calls;
		if (firstCalls == 0)
			fail("the projection did not recover a stored argument");
		if (render(steady.discover()) != first)
			fail("the confirming discovery changed unchanged helpers");
		if (calls != firstCalls)
			fail("the confirming discovery lowered unchanged helpers again");
		external = "compiler-2";
		steady.discover();
		if (calls == firstCalls)
			fail("a changed compiler input did not lower the helpers again");

		final oneDemand = (owner:GoValueRecovery, _:GoRecoveryAvailability) -> {
			owner.noteSource(intSource);
			owner.recover(GoExpr.GoIdent("stored"), wideTarget);
		};
		final twoDemands = (owner:GoValueRecovery, _:GoRecoveryAvailability) -> {
			owner.noteSource(intSource);
			owner.recover(GoExpr.GoIdent("narrow"), narrowTarget);
			owner.recover(GoExpr.GoIdent("wide"), wideTarget);
		};
		final alphaLocal = (owner:GoValueRecovery, availability:GoRecoveryAvailability) -> {
			availability.noteGeneratedType("Local", "alpha_owner", GoBuildConstraint.parse("alpha"));
			owner.noteSource(localSource);
			owner.recover(GoExpr.GoIdent("narrow"), narrowTarget);
			owner.recover(GoExpr.GoIdent("wide"), wideTarget);
		};
		final commonLocal = (_:GoValueRecovery, availability:GoRecoveryAvailability) -> {
			availability.noteGeneratedType("Local", "common_owner", GoBuildConstraint.always());
		};

		expectChangeBetweenPasses("a source found after discovery", oneDemand, (owner, _) -> owner.noteSource(stringSource));
		expectChangeBetweenPasses("a demand found after discovery", oneDemand, (owner, _) -> owner.recover(GoExpr.GoIdent("later"), narrowTarget));
		expectChangeBetweenPasses("an erased call found after discovery", oneDemand, (owner, _) -> owner.requireInvocation());
		expectChangeBetweenPasses("an owner found after discovery", alphaLocal, commonLocal);

		expectChangeDuringDiscovery("a source found during discovery", twoDemands, (owner, _) -> owner.noteSource(stringSource));
		expectChangeDuringDiscovery("an erased call found during discovery", twoDemands, (owner, _) -> owner.requireInvocation());
		expectChangeDuringDiscovery("an owner added after a helper read its declaration", alphaLocal, commonLocal);

		// A declaration that never gets an owner must still stop the compilation.
		// Each pass starts with no unresolved names, so the confirming discovery
		// has to report the missing owner again before availability freezes.
		final missingAvailability = newAvailability();
		final missing = new GoValueRecovery(missingAvailability, (value, type) -> GoExpr.GoTypeAssert(value, type), null, () -> "compiler");
		missing.noteSource(GoType.parse("*hxrt.Function[func(*Missing) int]"));
		missing.recover(GoExpr.GoIdent("stored"), wideTarget);
		if (missing.discover().length != 0)
			fail("a helper with an unowned source declaration was emitted");
		missingAvailability.beginPass();
		missing.discover();
		final refused = try {
			missingAvailability.freeze();
			false;
		} catch (_:haxe.Exception) {
			true;
		};
		if (!refused)
			fail("a confirming discovery hid an unresolved declaration from freeze");

		// Without a revision of the compiler facts the callback reads, the owner
		// cannot prove them unchanged, so every discovery lowers again.
		var unknownCalls = 0;
		final unknown = new GoValueRecovery(newAvailability(), (value, type) -> {
			unknownCalls++;
			return GoExpr.GoTypeAssert(value, type);
		});
		unknown.noteSource(intSource);
		unknown.recover(GoExpr.GoIdent("stored"), wideTarget);
		unknown.discover();
		final unknownFirst = unknownCalls;
		unknown.discover();
		if (unknownCalls == unknownFirst)
			fail("discovery reused helpers without a compiler input revision");
	}

	/** Discover, apply `change` between two passes, and discover again. */
	static function expectChangeBetweenPasses(name:String, setup:(GoValueRecovery, GoRecoveryAvailability) -> Void,
			change:(GoValueRecovery, GoRecoveryAvailability) -> Void):Void {
		final availability = newAvailability();
		final owner = new GoValueRecovery(availability, (value, type) -> GoExpr.GoTypeAssert(value, type), null, () -> "compiler");
		setup(owner, availability);
		final before = render(owner.discover());
		change(owner, availability);
		expectFinal(name, before, render(owner.discover()), setup, change);
	}

	/** Apply `change` from the second stored-value recovery of a discovery, then discover again. */
	static function expectChangeDuringDiscovery(name:String, setup:(GoValueRecovery, GoRecoveryAvailability) -> Void,
			change:(GoValueRecovery, GoRecoveryAvailability) -> Void):Void {
		final availability = newAvailability();
		var owner:Null<GoValueRecovery> = null;
		var calls = 0;
		final recovery = new GoValueRecovery(availability, (value, type) -> {
			calls++;
			if (calls == 2 && owner != null)
				change(owner, availability);
			return GoExpr.GoTypeAssert(value, type);
		}, null, () -> "compiler");
		owner = recovery;
		setup(recovery, availability);
		final stale = render(recovery.discover());
		if (calls < 2)
			fail(name + ": the fixture no longer recovers two stored values");
		expectFinal(name, stale, render(recovery.discover()), setup, change);
	}

	static function expectFinal(name:String, before:String, after:String, setup:(GoValueRecovery, GoRecoveryAvailability) -> Void,
			change:(GoValueRecovery, GoRecoveryAvailability) -> Void):Void {
		final availability = newAvailability();
		final fresh = new GoValueRecovery(availability, (value, type) -> GoExpr.GoTypeAssert(value, type), null, () -> "compiler");
		setup(fresh, availability);
		change(fresh, availability);
		final expected = render(fresh.discover());
		if (before == expected)
			fail(name + ": the fixture no longer changes the helpers");
		if (after != expected)
			fail(name + ": the next discovery did not lower the helpers again");
	}

	static function declarationNames(values:Array<reflaxe.go.compiler.GoRecoveryDispatch.GoGuardedRecoveryDeclaration>):Array<String> {
		return [
			for (value in values)
				switch (value.declaration) {
					case GoFuncDecl(name, _, _, _, _):
						name;
					case _:
						"";
				}
		];
	}
	#end
}
