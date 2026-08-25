class Main {
	static function main():Void {
		final result = GoHTTP.newRequest("GET", "https://example.com", null);
		final url:GoURLAlias = result.request.url;
		GoFmt.println(result.error == null && url != null ? 42 : -1);
	}
}
