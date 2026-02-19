class Main {
	static function joinParts(parts:Array<String>, delimiter:String):String {
		var out = "";
		for (index in 0...parts.length) {
			if (index > 0) {
				out += delimiter;
			}
			out += parts[index];
		}
		return out;
	}

	static function main() {
		var mapRegex = ~/([a-z]+)([0-9]+)/g;
		var mapped = mapRegex.map("ab12 cd34 ef5", function(expr:EReg) {
			return expr.matched(2) + ":" + expr.matched(1);
		});
		Sys.println("map=" + mapped);

		var split = ~/[,:;]/g.split("a,b:c;d");
		Sys.println("split=" + joinParts(split, "|"));

		var scan = ~/item([0-9]+)/g;
		var haystack = "item1|item22|item333";
		var hits = [];
		var pos = 0;
		while (scan.matchSub(haystack, pos)) {
			var mpos = scan.matchedPos();
			hits.push(scan.matched(1) + "@" + mpos.pos + ":" + mpos.len);
			pos = mpos.pos + mpos.len;
		}
		Sys.println("hits=" + joinParts(hits, ","));

		var once = ~/bar/;
		var didMatch = once.match("foo-bar-baz");
		Sys.println("matched=" + didMatch);
		if (didMatch) {
			Sys.println("left=" + once.matchedLeft());
			Sys.println("right=" + once.matchedRight());
		}

		var collapsed = ~/\s+/g.replace("hi   there   go", "-");
		Sys.println("replace=" + collapsed);
	}
}
