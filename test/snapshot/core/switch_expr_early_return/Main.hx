typedef Item = {
	final id:String;
}

enum ItemOutcome {
	ItemSuccess(items:Array<Item>);
	ItemFailure(message:String);
}

class Main {
	static function describe(value:Int):String {
		final selected = switch value {
			case 0:
				return "zero";
			case 1:
				"one";
			case _:
				if (value < 0)
					return "negative";
				"many";
		};
		return "selected:" + selected;
	}

	static function selectItem(outcome:ItemOutcome):ItemOutcome {
		final selected = switch outcome {
			case ItemFailure(message):
				return ItemFailure(message);
			case ItemSuccess(items):
				items[0];
		};
		return ItemSuccess([selected]);
	}

	static function main():Void {
		Sys.println(describe(0));
		Sys.println(describe(1));
		Sys.println(describe(-1));
		Sys.println(describe(2));
		switch selectItem(ItemFailure("zero-item")) {
			case ItemFailure(message):
				Sys.println(message);
			case ItemSuccess(_):
		}
		switch selectItem(ItemSuccess([{id: "one"}])) {
			case ItemFailure(_):
			case ItemSuccess(items):
				Sys.println("selected-item:" + items[0].id);
		}
	}
}
