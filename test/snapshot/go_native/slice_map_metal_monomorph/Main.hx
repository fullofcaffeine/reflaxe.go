import go.Go;
import go.Map;
import go.Slice;

class Main {
	static function main() {
		var ints = new Slice<Int>();
		ints.push(3);
		ints.push(5);
		ints.set(1, 8);
		Sys.println(ints.length);
		Sys.println(ints.get(1));
		var intsArray = ints.toArray();
		Sys.println(intsArray[0]);

		var words:Slice<String> = Go.newSlice();
		words.push("go");
		words.push("haxe");
		Sys.println(words.get(0));
		var wordsArray = words.toArray();
		Sys.println(wordsArray.length);

		var scores = new Map<Int, String>();
		scores.set(7, "seven");
		Sys.println(scores.exists(7));
		Sys.println(scores.get(7));
		var missing = scores.get(99);
		Sys.println(missing == null ? "none" : missing);

		var byName:Map<String, Int> = Go.newMap();
		byName.set("alice", 11);
		Sys.println(byName.exists("alice"));
		Sys.println(byName.get("alice"));
	}
}
