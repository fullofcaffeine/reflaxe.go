class Main {
	static function bump(user:{var score:Int; var label:String;}):Void {
		user.score = user.score + 2;
		user.label = user.label + "!";
	}

	static function main() {
		var user = {score: 3, label: "go"};
		bump(user);
		Sys.println(user.score);
		Sys.println(user.label);

		var nested = {inner: {ok: true, count: 2}};
		nested.inner.count = nested.inner.count + 4;
		Sys.println(nested.inner.ok);
		Sys.println(nested.inner.count);
	}
}
