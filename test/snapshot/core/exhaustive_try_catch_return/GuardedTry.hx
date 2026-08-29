function readValue(skip:Bool, fail:Bool):String {
	if (skip) {
		return "";
	}
	try {
		if (fail) {
			throw "failed";
		}
		return "value";
	} catch (_:Dynamic) {
		return "fallback";
	}
}
