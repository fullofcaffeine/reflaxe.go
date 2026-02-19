class Main {
	static function main() {
		var http = new haxe.Http("data:text/plain,hello%20from%20haxe.go");

		var sink = new haxe.io.BytesBuffer();
		http.customRequest(false, cast sink);
		Sys.println("custom=" + sink.getBytes().toString());

		var values = http.getResponseHeaderValues("Content-Type");
		Sys.println("headers=" + (values == null ? -1 : values.length));
		Sys.println("header0=" + ((values != null && values.length > 0) ? values[0] : "none"));

		var putSink = new haxe.io.BytesBuffer();
		http.customRequest(false, cast putSink, null, "PUT");
		Sys.println("method=" + putSink.getBytes().toString());

		var upload = new haxe.Http("data:text/plain,ignored");
		upload.setParameter("token", "42");
		upload.fileTransfer("asset", "demo.txt", cast null, 4, "text/plain");
		var uploadSink = new haxe.io.BytesBuffer();
		upload.customRequest(true, cast uploadSink);
		Sys.println("upload=" + uploadSink.getBytes().toString());
	}
}
