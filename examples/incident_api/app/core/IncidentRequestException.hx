package app.core;

/**
	What: Typed request error for invalid incident API input.
	Why: The HTTP boundary may receive malformed JSON or missing fields, but the
	router should not use broad `Dynamic` catches for ordinary control flow.
	How: Parser helpers throw this typed exception with a stable response code, and
	`IncidentApi` maps it to a deterministic JSON error response.
**/
class IncidentRequestException extends haxe.Exception {
	public final code:String;

	public function new(code:String) {
		super(code);
		this.code = code;
	}
}
