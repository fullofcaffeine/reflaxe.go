@:go.import("net/http")
@:go.package("http")
extern class GoHTTP {
	@:go.tupleReturn
	@:go.name("NewRequest")
	public static function newRequest(method:String, target:String, body:Dynamic):GoHTTPRequestResult;
}
