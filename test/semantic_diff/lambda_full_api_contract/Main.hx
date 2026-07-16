import haxe.ds.List;

class FullApiIterator {
	final data:Array<Int>;
	var index:Int;

	public function new(data:Array<Int>) {
		this.data = data;
		this.index = 0;
	}

	public function hasNext():Bool {
		return index < data.length;
	}

	public function next():Int {
		return data[index++];
	}
}

class FullApiIterable {
	final data:Array<Int>;

	public function new(data:Array<Int>) {
		this.data = data;
	}

	public function iterator():FullApiIterator {
		return new FullApiIterator(data);
	}
}

class NestedArrayIterator {
	final data:Array<Array<Int>>;
	var index:Int;

	public function new(data:Array<Array<Int>>) {
		this.data = data;
		this.index = 0;
	}

	public function hasNext():Bool {
		return index < data.length;
	}

	public function next():Array<Int> {
		return data[index++];
	}
}

class NestedArrayIterable {
	final data:Array<Array<Int>>;

	public function new(data:Array<Array<Int>>) {
		this.data = data;
	}

	public function iterator():NestedArrayIterator {
		return new NestedArrayIterator(data);
	}
}

class Main {
	static function makeList(values:Array<Int>):List<Int> {
		var out = new List<Int>();
		for (value in values) {
			out.add(value);
		}
		return out;
	}

	static function listString(values:List<Int>):String {
		return [for (value in values) value].join(",");
	}

	static function main() {
		var arrayValues = [1, 2, 3];
		var listValues = makeList([1, 2, 3]);
		var genericValues:Iterable<Int> = new FullApiIterable([1, 2, 3]);

		Sys.println("array.array=" + Lambda.array(arrayValues).join(","));
		Sys.println("array.list=" + Lambda.array(listValues).join(","));
		Sys.println("array.generic=" + Lambda.array(genericValues).join(","));
		Sys.println("list.array=" + listString(Lambda.list(arrayValues)));
		Sys.println("list.generic=" + listString(Lambda.list(genericValues)));

		Sys.println("mapi.array=" + Lambda.mapi(arrayValues, function(index:Int, value:Int):Int return index * 10 + value).join(","));
		Sys.println("mapi.list=" + Lambda.mapi(listValues, function(index:Int, value:Int):Int return index * 10 + value).join(","));
		Sys.println("mapi.generic=" + Lambda.mapi(genericValues, function(index:Int, value:Int):Int return index * 10 + value).join(","));

		var foreachCalls = 0;
		var allBelowThree = Lambda.foreach(genericValues, function(value:Int):Bool {
			foreachCalls++;
			return value < 3;
		});
		Sys.println("foreach.result=" + allBelowThree + ":calls=" + foreachCalls);
		Sys.println("foldi.generic=" + Lambda.foldi(genericValues, function(value:Int, result:Int, index:Int):Int return result + value + index * 10, 0));
		Sys.println("indexOf.present=" + Lambda.indexOf(genericValues, 2));
		Sys.println("indexOf.absent=" + Lambda.indexOf(genericValues, 9));
		Sys.println("find.present=" + Lambda.find(genericValues, function(value:Int):Bool return value > 1));
		Sys.println("find.absent=" + Lambda.find(genericValues, function(value:Int):Bool return value > 9));
		Sys.println("findIndex.present=" + Lambda.findIndex(genericValues, function(value:Int):Bool return value == 3));
		Sys.println("findIndex.absent=" + Lambda.findIndex(genericValues, function(value:Int):Bool return value == 9));
		Sys.println("concat.mixed=" + Lambda.concat(genericValues, listValues).join(","));

		var nestedArrayValues = [[1, 2], [3, 4]];
		var nestedListValues = new List<Array<Int>>();
		nestedListValues.add([5, 6]);
		nestedListValues.add([7]);
		var nestedGenericValues = new NestedArrayIterable([[8], [9, 10]]);
		Sys.println("flatten.array=" + Lambda.flatten(nestedArrayValues).join(","));
		Sys.println("flatten.list=" + Lambda.flatten(nestedListValues).join(","));
		Sys.println("flatten.generic=" + Lambda.flatten(nestedGenericValues).join(","));

		var flatMappedArray = Lambda.flatMap(genericValues, function(value:Int) return [value, value + 10]);
		var flatMappedList = Lambda.flatMap(listValues, function(value:Int) return makeList([value, -value]));
		Sys.println("flatMap.array=" + flatMappedArray.join(","));
		Sys.println("flatMap.list=" + flatMappedList.join(","));
	}
}
