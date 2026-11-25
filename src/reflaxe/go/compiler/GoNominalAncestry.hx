package reflaxe.go.compiler;

#if macro
import haxe.macro.Type.ClassType;

/**
	What: The declared ancestry of one class (superclasses and interfaces, by
	generated type name), walked only as far as the questions asked so far need.

	Why: `GoNominalTypeTest.discover` asks, for every demanded target and every
	generated class, whether the class belongs to the target. Each question walked
	the class's ancestry again, rebuilding each declaration with `Ref.get()` and
	naming each class again. The runtime type-token test demands a nominal test
	of every generated class, so on the BeadsHX bootstrap compile each recovery
	pass asked about 150 targets of about 150 classes: about 0.38 s of a 45 s
	compile.

	How: The walk keeps the order of the old recursive test: a class's name is
	checked against the target, then, if its name was not seen before, its
	superclass ancestry and then each interface's ancestry follow. The old test
	returned at the first class with the target's name. `contains` stops there
	too, after it has queued that class's parents, so a later question resumes
	the same walk where the old test, starting again, would reach new classes. A
	class with that name already met answers at once, because the old walk would
	have met it at the same place.

	Naming a class can record compiler facts (a shim group or an extern import),
	and the same class always records the same facts. The walk names each class
	at its first place in the old walk order, and never names a class that the old
	tests would not have reached, so the recorded facts arrive in the same order.
**/
class GoNominalAncestry {
	final className:ClassType->String;
	final met:Map<String, Bool> = [];
	final expanded:Map<String, Bool> = [];
	final pending:Array<ClassType>;

	public function new(source:ClassType, className:ClassType->String) {
		this.className = className;
		pending = [source];
	}

	/** True when the class or one of its declared ancestors has the generated type name `name`. */
	public function contains(name:String):Bool {
		if (met.exists(name))
			return true;
		while (pending.length > 0) {
			final current = pending.pop();
			final currentName = className(current);
			met.set(currentName, true);
			if (!expanded.exists(currentName)) {
				expanded.set(currentName, true);
				// Pop order: the superclass ancestry first, then each interface in order.
				var index = current.interfaces.length;
				while (index-- > 0)
					pending.push(current.interfaces[index].t.get());
				if (current.superClass != null)
					pending.push(current.superClass.t.get());
			}
			if (currentName == name)
				return true;
		}
		return false;
	}
}
#end
