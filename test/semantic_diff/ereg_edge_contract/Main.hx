class Main {
	static function safe(label:String, fn:Void->String):Void {
		try {
			Sys.println(label + "=ok:" + fn());
		} catch (e:Dynamic) {
			Sys.println(label + "=err");
		}
	}

	static function main() {
		var once = ~/a./;
		var global = ~/a./g;
		safe("replace.once", function() return Std.string(once.replace("ab ac ad", "X")));
		safe("replace.global", function() return Std.string(global.replace("ab ac ad", "X")));
		safe("map.once", function() return Std.string(once.map("ab ac ad", function(r:EReg) return "[" + r.matched(0) + "]")));
		safe("map.global", function() return Std.string(global.map("ab ac ad", function(r:EReg) return "[" + r.matched(0) + "]")));

		var multiline = ~/^bar/m;
		var noMultiline = ~/^bar/;
		safe("flag.m.true", function() return Std.string(multiline.match("foo\nbar")));
		safe("flag.m.false", function() return Std.string(noMultiline.match("foo\nbar")));

		var dotAll = ~/a.b/s;
		var noDotAll = ~/a.b/;
		safe("flag.s.true", function() return Std.string(dotAll.match("a\nb")));
		safe("flag.s.false", function() return Std.string(noDotAll.match("a\nb")));

		var grouped = ~/(a)(b)?/;
		safe("group.before", function() return Std.string(grouped.matched(0)));
		safe("group.match", function() return Std.string(grouped.match("a")));
		safe("group.g0", function() return Std.string(grouped.matched(0)));
		safe("group.g1", function() return Std.string(grouped.matched(1)));
		safe("group.g2", function() return Std.string(grouped.matched(2)));
		safe("group.g3", function() return Std.string(grouped.matched(3)));

		safe("state.match.ok", function() return Std.string(grouped.match("a")));
		safe("state.match.fail", function() return Std.string(grouped.match("zzz")));
		safe("state.after.match.fail", function() return Std.string(grouped.matched(0)));

		safe("state.matchSub.ok", function() return Std.string(grouped.match("a")));
		safe("state.matchSub.fail", function() return Std.string(grouped.matchSub("zzz", 0)));
		safe("state.after.matchSub.fail", function() return Std.string(grouped.matched(0)));
	}
}
