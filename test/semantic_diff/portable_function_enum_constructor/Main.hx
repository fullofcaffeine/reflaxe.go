enum Parcel {
	Item(value:Int);
}

class Main {
	static var make:Int->Parcel;

	static function main() {
		make = Parcel.Item;
		final alias = make;
		switch alias(17) {
			case Item(value):
				Sys.println(value);
		}
		Sys.println(Reflect.compareMethods(make, alias));
		Sys.println(Reflect.compareMethods(Parcel.Item, Parcel.Item));
	}
}
