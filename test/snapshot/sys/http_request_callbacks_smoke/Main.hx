class Main {
	static function main() {
		var http = new haxe.Http("data:text/plain,hello%20from%20haxe.go");
		var dataLog = "";
		var statusLog = -1;
		var byteCount = -1;
		var errLog = "";

		http.onData = function(data) {
			dataLog = data;
		};
		http.onStatus = function(status) {
			statusLog = status;
		};
		http.onBytes = function(bytes) {
			byteCount = bytes.length;
		};
		http.onError = function(msg) {
			errLog = msg;
		};

		http.setHeader("X-Test", "one");
		http.setHeader("X-Test", "two");
		http.addHeader("X-Trace", "ok");
		http.setParameter("a", "1");
		http.addParameter("b", "2");
		http.request();

		Sys.println("data=" + dataLog);
		Sys.println("bytes=" + byteCount);
		Sys.println("status=" + statusLog);
		Sys.println("response=" + http.responseData);
		Sys.println("error=" + errLog);

		var post = new haxe.Http("data:text/plain,ignored");
		post.setPostData("post-body");
		var postData = "";
		post.onData = function(data) {
			postData = data;
		};
		post.request(true);
		Sys.println("post=" + postData);

		var bad = new haxe.Http("://bad");
		var badErr = "";
		bad.onError = function(msg) {
			badErr = msg;
		};
		bad.request();
		Sys.println("bad=" + badErr);

		Sys.println("direct=" + haxe.Http.requestUrl("data:text/plain,direct%20ok"));
	}
}
