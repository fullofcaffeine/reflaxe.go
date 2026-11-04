import reflaxe.go.compiler.GoBuildConstraint;
import reflaxe.go.ast.GoType;

/** Independent Boolean and type-dependency contracts for selected Go file sets. */
class GoRecoveryAvailabilityContract {
	static function require(value:Bool, message:String):Void {
		if (!value)
			throw message;
	}

	static function main():Void {
		var alpha = GoBuildConstraint.parse("alpha");
		var beta = GoBuildConstraint.parse("beta");
		require(alpha.and(beta).key() == beta.and(alpha).key(), "conjunction order changed availability identity");
		require(alpha.and(alpha).key() == alpha.key(), "duplicate conditions were retained");
		require(alpha.and(alpha.negate()).isNever(), "mutually exclusive declarations can overlap");
		require(alpha.or(alpha.negate()).isAlways(), "complementary probe declarations leave a gap");
		require(alpha.and(GoBuildConstraint.always()).key() == alpha.key(), "common dependency constrained a type");
		require(alpha.or(GoBuildConstraint.never()).key() == alpha.key(), "absent alternative changed availability");
		require(alpha.negate().negate().key() == alpha.key(), "double negation changed availability");
		require(!GoBuildConstraint.parse("true").isAlways(), "the Go tag true was treated as a Boolean constant");
		var both = alpha.and(beta);
		var alternate = alpha.and(beta.negate());
		require(both.render() == "alpha && beta", "conjunction rendering is not deterministic");
		require(alternate.render() == "!beta && alpha", "fallback was not restricted to target availability");
		require(GoBuildConstraint.parse("alpha || beta").and(GoBuildConstraint.parse("gamma")).render() == "(alpha || beta) && gamma",
			"composition lost Boolean precedence");
		var references = GoType.parse("func(*alpha.Node, map[string][]beta.Item, chan<- *gamma.Value, ...*delta.Input) (epsilon.Result, interface{Get(zeta.Key) eta.Value})")
			.namedDependencies();
		var names = [
			for (reference in references)
				(reference.packageName == null ? "" : reference.packageName.value() + ".") + reference.name
		];
		require(names.join(",") == "alpha.Node,beta.Item,delta.Input,epsilon.Result,eta.Value,gamma.Value,zeta.Key", "nested type dependencies were omitted");
		var repeated = GoType.parse("*hxrt.Function[func(*Local) *Local]").namedDependencies();
		require(repeated.length == 2, "physical type dependencies were not deduplicated");
		Sys.println("ok");
	}
}
