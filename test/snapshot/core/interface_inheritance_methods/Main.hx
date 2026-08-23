interface RootQuery<T> {
	public function root(value:T):T;
}

interface LeftQuery extends RootQuery<String> {
	public function left():String;
}

interface RightQuery extends RootQuery<String> {
	public function right():String;
}

interface CombinedQuery extends LeftQuery extends RightQuery {
	public function local():String;
}

class QueryService implements CombinedQuery {
	public function new() {}

	public function root(value:String):String {
		return 'root:${value}';
	}

	public function left():String {
		return "left";
	}

	public function right():String {
		return "right";
	}

	public function local():String {
		return "local";
	}
}

class Main {
	static function printQuery(query:CombinedQuery):Void {
		query.root("root");
		query.left();
		query.right();
		query.local();
	}

	static function main() {
		printQuery(new QueryService());
	}
}
