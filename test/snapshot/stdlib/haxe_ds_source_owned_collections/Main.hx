class Main {
	static function main() {
		var tree = new haxe.ds.BalancedTree<Int, String>();
		tree.set(2, "two");
		tree.set(1, "one");
		tree.set(3, "three");

		var stack = new haxe.ds.GenericStack<String>();
		stack.add("alpha");
		stack.add("beta");

		Sys.println(tree.get(1));
		Sys.println(stack.toString());
	}
}
