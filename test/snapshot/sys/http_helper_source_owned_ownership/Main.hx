@:native("sys__Http")
extern class SysHttpInternals {
	static function hxrt_proxyDescriptor():String;
}

class Main {
	static function main() {
		var http = new haxe.Http("data:text/plain,hello%20leaf");
		var sink = new haxe.io.BytesBuffer();
		http.customRequest(false, cast sink);
		var values = http.getResponseHeaderValues("Content-Type");
		Sys.println("headers=" + ((values != null && values.length > 0) ? values[0] : "null"));
		Sys.println("direct=" + haxe.Http.requestUrl("data:text/plain,direct%20ok"));
		Sys.println("proxy0=" + SysHttpInternals.hxrt_proxyDescriptor());
		haxe.Http.PROXY = cast {
			host: "proxy.local",
			port: 3128,
			auth: {user: "scott", pass: "tiger"}
		};
		Sys.println("proxy1=" + SysHttpInternals.hxrt_proxyDescriptor());
	}
}
