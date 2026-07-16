import haxe.ds.ArraySort;
import haxe.ds.List;
import haxe.ds.ListSort;

class Node {
	public var prev:Node;
	public var next:Node;
	public final value:Int;

	public function new(value:Int) {
		this.value = value;
	}
}

class Main {
	static function cmpInt(a:Int, b:Int):Int {
		return a - b;
	}

	static function cmpNode(a:Node, b:Node):Int {
		return a.value - b.value;
	}

	static function main() {
		var values = [4, 1, 3, 2];
		ArraySort.sort(values, cmpInt);
		Sys.println("array=" + values[0] + "," + values[1] + "," + values[2] + "," + values[3]);

		var n4 = new Node(4);
		var n1 = new Node(1);
		var n3 = new Node(3);
		var n2 = new Node(2);
		n4.next = n1;
		n1.prev = n4;
		n1.next = n3;
		n3.prev = n1;
		n3.next = n2;
		n2.prev = n3;
		var head = ListSort.sort(n4, cmpNode);
		Sys.println("list=" + head.value + "," + head.next.value + "," + head.next.next.value + "," + head.next.next.next.value);
		Sys.println("tail=" + head.prev.value);

		var s4 = new Node(4);
		var s1 = new Node(1);
		var s3 = new Node(3);
		var s2 = new Node(2);
		s4.next = s1;
		s1.next = s3;
		s3.next = s2;
		var singleHead = ListSort.sortSingleLinked(s4, cmpNode);
		Sys.println("single="
			+ singleHead.value
			+ ","
			+ singleHead.next.value
			+ ","
			+ singleHead.next.next.value
			+ ","
			+ singleHead.next.next.next.value);
	}
}
