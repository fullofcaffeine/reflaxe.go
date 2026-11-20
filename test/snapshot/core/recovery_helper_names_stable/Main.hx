import Shapes;

class Main {
	static function main() {
		Sys.println(Edited.run(new Square(2), new Circle(3), (text:String) -> text.length));
		Sys.println(Stable.run(new Circle(4), (size:Int) -> size * 10));
	}
}
