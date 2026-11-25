/** The root class: no superclass, one interface. */
class GoNominalAncestryContractFixture implements AncestryShape {
	public function new() {}
}

/** A root interface. */
interface AncestryShape {}

/** An interface whose parent is `AncestryShape`. */
interface AncestryNamed extends AncestryShape {}

/** An interface with no parent. */
interface AncestryOther {}

/** A second interface with no parent. */
interface AncestryExtra {}

/** A subclass that adds an interface. */
class AncestryMiddle extends GoNominalAncestryContractFixture implements AncestryNamed {}

/** A subclass with three interfaces; an ancestor also declares the last one. */
class AncestryLeaf extends AncestryMiddle implements AncestryOther implements AncestryExtra implements AncestryShape {}

/** An unrelated class. */
class AncestryLoose {
	public function new() {}
}
