package app.http;

/**
	What: Minimal JSON response carrier for the example HTTP server.
	Why: Makes status/body decisions explicit and testable in Haxe.
	How: The socket server turns this into HTTP/1.1 text.
**/
class HttpResponse {
	public var status:Int;
	public var body:String;

	public function new(status:Int, body:String) {
		this.status = status;
		this.body = body;
	}

	public static function json(status:Int, body:String):HttpResponse {
		return new HttpResponse(status, body);
	}
}
