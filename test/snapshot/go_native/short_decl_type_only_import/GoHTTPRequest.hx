@:go.import("net/http")
@:go.package("http")
@:go.name("Request")
extern class GoHTTPRequest {
	@:go.name("URL")
	public var url(default, null):GoURL;
}
