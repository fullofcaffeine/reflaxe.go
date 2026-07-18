import hxrt.http.NativeHttp;

class Main {
	static function main() {
		Sys.println(NativeHttp.proxyDescriptor("proxy.local", 3128, null, null));
	}
}
