class Main {
	static function main() {
		var tree = new haxe.ds.BalancedTree<Int, String>();
		tree.set(2, "two");
		tree.set(1, "one");
		tree.set(3, "three");
		Sys.println("tree.get2=" + tree.get(2));
		Sys.println("tree.exists4=" + Std.string(tree.exists(4)));
		Sys.println("tree.remove2=" + Std.string(tree.remove(2)));
		Sys.println("tree.exists2=" + Std.string(tree.exists(2)));
		Sys.println("tree.string=" + tree.toString());

		var stack = new haxe.ds.GenericStack<String>();
		stack.add("alpha");
		stack.add("beta");
		stack.add("gamma");
		Sys.println("stack.first=" + stack.first());
		Sys.println("stack.removeBeta=" + Std.string(stack.remove("beta")));
		Sys.println("stack.pop=" + stack.pop());
		Sys.println("stack.string=" + stack.toString());
	}
}
