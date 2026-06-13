package app.http;

/**
	What: Minimal request carrier for the example HTTP/1.1 parser.
	Why: Keeps socket parsing separate from incident-domain routing.
	How: `TinyHttpServer` fills one value per accepted connection.
**/
class HttpRequest {
	public var method:String;
	public var path:String;
	public var body:String;

	public function new(method:String, path:String, body:String) {
		this.method = method;
		this.path = path;
		this.body = body;
	}
}
