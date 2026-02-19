@:native("sys__Http")
extern class SysHttpInternals {
	static function hxrt_proxyDescriptor():String;
}

class Main {
	static function main() {
		Sys.println("proxy0=" + SysHttpInternals.hxrt_proxyDescriptor());

		haxe.Http.PROXY = cast {
			host: "proxy.local",
			port: 3128,
			auth: {user: "scott", pass: "tiger"}
		};
		Sys.println("proxy1=" + SysHttpInternals.hxrt_proxyDescriptor());

		haxe.Http.PROXY = cast {
			host: "proxy.local",
			port: 80,
			auth: {user: "scott", pass: null}
		};
		Sys.println("proxy2=" + SysHttpInternals.hxrt_proxyDescriptor());

		haxe.Http.PROXY = cast {
			host: "proxy.local:9000",
			port: 3128,
			auth: null
		};
		Sys.println("proxy3=" + SysHttpInternals.hxrt_proxyDescriptor());

		haxe.Http.PROXY = cast {
			host: null,
			port: 3128,
			auth: null
		};
		Sys.println("proxy4=" + SysHttpInternals.hxrt_proxyDescriptor());

		haxe.Http.PROXY = null;
		var http = new haxe.Http("data:text/plain,body");
		var sink = new haxe.io.BytesBuffer();
		http.customRequest(false, cast sink, cast {marker: "sock"}, "PATCH");
		Sys.println("methodSock=" + sink.getBytes().toString());
	}
}
