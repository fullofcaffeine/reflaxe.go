class Holder<T> {
	final value:T;
	final saved:Void->T;

	public function new(value:T) {
		this.value = value;
		this.saved = () -> value;
	}

	public function reader():Void->T {
		return () -> value;
	}

	public function savedReader():Void->T
		return saved;

	public function echo():T->T
		return value->value;

	public function missing():Void->T
		return null;
}

class Main {
	static function main() {
		final read = new Holder<Int>(19).reader();
		final alias = read;
		Sys.println(read());
		Sys.println(Reflect.compareMethods(read, alias));
		final holder = new Holder<Int>(23);
		Sys.println(holder.echo()(29));
		Sys.println(holder.savedReader()());
		Sys.println(Reflect.compareMethods(holder.savedReader(), holder.savedReader()));
		Sys.println(Reflect.compareMethods(holder.reader(), holder.reader()));
		Sys.println(holder.missing() == null);
		Sys.println(new Holder<String>("text").reader()());
		Sys.println(new Holder<Null<Int>>(null).reader()() == null);
	}
}
