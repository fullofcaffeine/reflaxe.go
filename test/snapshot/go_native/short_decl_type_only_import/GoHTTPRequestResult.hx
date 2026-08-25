/**
	What: Carries the two native values returned by `http.NewRequest`.
	Why: The fixture needs a request whose URL field belongs to another Go package.
	How: `@:go.tupleReturn` constructs this Haxe-owned carrier from the native call.
**/
class GoHTTPRequestResult {
	public var request(default, null):GoHTTPRequest;
	public var error(default, null):Null<go.Error>;

	public function new(request:GoHTTPRequest, error:Null<go.Error>) {
		this.request = request;
		this.error = error;
	}
}
