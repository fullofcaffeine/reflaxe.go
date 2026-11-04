import haxe.CallStack;

class FixedException extends haxe.Exception {
	var frames:CallStack;

	public function new(message:String, frames:CallStack, ?previous:haxe.Exception) {
		super(message, previous);
		this.frames = frames;
	}

	override function get_stack():CallStack
		return frames;
}

class Main {
	static function main():Void {
		var shared:StackItem = Method("Shared", "tail");
		var first:CallStack = [Method("Cause", "start"), shared];
		var second:CallStack = [Method("Outer", "start"), shared];
		Sys.println("subtract=" + CallStack.toString(first.subtract(second)));
		var cause = new FixedException("cause", first);
		var outer = new FixedException("outer", second, cause);
		Sys.println("details=" + outer.details());
		var identical:CallStack = [FilePos(Method("Same", "f"), "sample.hx", 7, 3)];
		var equal:CallStack = [FilePos(Method("Same", "f"), "sample.hx", 7, 3)];
		Sys.println("identical=" + identical.subtract(equal).length);
		var different:CallStack = [Module("elsewhere")];
		Sys.println("distinct=" + first.subtract(different).length);
	}
}
