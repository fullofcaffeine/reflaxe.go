class Main {
	static function makeUser(name:String, score:Int) {
		return {name: name, score: score};
	}

	static function main() {
		var user = makeUser("marcelo", 10);
		Sys.println(user.name);

		user.score = user.score + 5;
		Sys.println(user.score);

		var nested = {inner: {flag: true, count: 2}};
		Sys.println(nested.inner.flag);

		nested.inner.count = nested.inner.count + 3;
		Sys.println(nested.inner.count);
	}
}
