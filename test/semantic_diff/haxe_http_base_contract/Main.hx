class Main {
	static function main() {
		var method:haxe.http.HttpMethod = haxe.http.HttpMethod.Post;
		var status:haxe.http.HttpStatus = haxe.http.HttpStatus.OK;
		Sys.println("method=" + method);
		Sys.println("status=" + status);

		var base = new haxe.http.HttpBase("http://example.com");
		base.setHeader("X-Test", "1");
		base.addHeader("X-Test", "2");
		base.setParameter("q", "go");
		base.addParameter("lang", "hx");
		base.setPostData("body");
		Sys.println("responseData=" + Std.string(base.responseData));
		try {
			base.request(false);
			Sys.println("request=no_throw");
		} catch (e:Dynamic) {
			Sys.println("requestType=" + Type.getClassName(Type.getClass(e)));
		}
	}
}
