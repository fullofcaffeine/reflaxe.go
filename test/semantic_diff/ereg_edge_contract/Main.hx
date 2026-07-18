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
		var captures = ~/(a)(b)?/g;
		safe("replace.captures", function() return Std.string(captures.replace("ab a", "<$1:$2:$$>")));
		safe("replace.haxe.syntax", function() return Std.string(captures.replace("ab", "<$1x:$2y:$$>")));
		safe("map.once", function() return Std.string(once.map("ab ac ad", function(r:EReg) return "[" + r.matched(0) + "]")));
		safe("map.global", function() return Std.string(global.map("ab ac ad", function(r:EReg) return "[" + r.matched(0) + "]")));
		var zeroWidth = ~/x*/g;
		safe("zero.replace", function() return zeroWidth.replace("ab", "|"));
		safe("zero.replace.empty", function() return zeroWidth.replace("", "|"));
		safe("zero.map", function() return zeroWidth.map("ab", function(_) return "|"));
		safe("zero.map.empty", function() return zeroWidth.map("", function(_) return "|"));
		safe("zero.split", function() {
			var parts = zeroWidth.split("ab");
			return parts.length + ":" + parts.map(function(part) return "[" + part + "]").join(",");
		});
		safe("zero.split.empty", function() {
			var parts = zeroWidth.split("");
			return parts.length + ":" + parts.map(function(part) return "[" + part + "]").join(",");
		});
		var adjacentZero = ~/a*/g;
		safe("zero.adjacent.replace", function() return adjacentZero.replace("ab", "|"));
		safe("zero.adjacent.map", function() return adjacentZero.map("ab", function(_) return "|"));
		safe("zero.adjacent.split", function() {
			var parts = adjacentZero.split("ab");
			return parts.length + ":" + parts.map(function(part) return "[" + part + "]").join(",");
		});
		var zeroOnce = ~/x*/;
		safe("zero.once.replace.empty", function() return zeroOnce.replace("", "|"));
		safe("zero.once.map.empty", function() return zeroOnce.map("", function(_) return "|"));
		safe("zero.once.split.empty", function() {
			var parts = zeroOnce.split("");
			return parts.length + ":" + parts.map(function(part) return "[" + part + "]").join(",");
		});

		var multiline = ~/^bar/m;
		var noMultiline = ~/^bar/;
		safe("flag.m.true", function() return Std.string(multiline.match("foo\nbar")));
		safe("flag.m.false", function() return Std.string(noMultiline.match("foo\nbar")));

		var dotAll = ~/a.b/s;
		var noDotAll = ~/a.b/;
		safe("flag.s.true", function() return Std.string(dotAll.match("a\nb")));
		safe("flag.s.false", function() return Std.string(noDotAll.match("a\nb")));

		var insensitive = ~/abc/i;
		safe("flag.i.true", function() return Std.string(insensitive.match("AbC")));
		var unicode = new EReg("é.", "u");
		safe("flag.u.true", function() return Std.string(unicode.match("é🙂")));
		safe("escape.literal", function() {
			var literal = "a+b[c].";
			return Std.string(new EReg(EReg.escape(literal), "").match(literal));
		});

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
