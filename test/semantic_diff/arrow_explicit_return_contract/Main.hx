class Main {
	static function main():Void {
		final effects = new Array<String>();
		final select = (value:Int) -> {
			effects.push("enter:" + value);
			return value < 0 ? mark(effects, "negative", -1) : (value == 0 ? mark(effects, "zero", 0) : mark(effects, "positive", 1));
		};
		Sys.println([select(-2), select(0), select(3)].join(","));
		Sys.println(effects.join(","));

		final early = (value:Int) -> {
			if (value < 0)
				return "early";
			return "tail:" + value;
		};
		Sys.println(early(-1) + "," + early(2));

		final branches = (flag:Bool) -> {
			if (flag) {
				return 7;
			} else {
				return 9;
			}
		};
		Sys.println(branches(true) + branches(false));

		final implicit = (value:Int) -> {
			final doubled = value * 2;
			doubled + 1;
		};
		Sys.println(implicit(4));

		final guarded = (fail:Bool) -> {
			try {
				if (fail)
					throw "failure";
				return 11;
			} catch (message:String) {
				return message.length;
			}
		};
		Sys.println(guarded(false) + "," + guarded(true));
	}

	static function mark(effects:Array<String>, name:String, value:Int):Int {
		effects.push(name);
		return value;
	}
}
