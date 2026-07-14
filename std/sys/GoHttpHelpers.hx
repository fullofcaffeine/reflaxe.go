package sys;

import haxe.io.Bytes;

/**
	What
	- Source-owned helper surface for `sys.Http` leaf helpers on `haxe.go`.
	- Owns leaf operations that do not decide request lifecycle or callback timing:
	  `getResponseHeaderValues` lookup glue and API capture fan-out for
	  `customRequest`.

	Why
	- The mainstream Haxe `sys.Http` implementation could not be reused unchanged
	  here because `haxe.go` still owns the core request choreography in
	  `GoCompiler`, including data-URL handling, proxy transport wiring, and the
	  callback/status/error ordering contract.
	- These two leaves still touched generated Go-native storage or interface
	  shapes that are not consistently visible from staged Haxe typing alone, so
	  they had been left in compiler raw blocks longer than necessary.

	How
	- Keep `hxrt__http__requestWith` and proxy URL construction compiler-owned.
	- Move only the leaf helpers here and bridge the hidden generated shapes
	  through narrow framework-owned `__go__` calls.
	- Preserve the public wrapper symbols in `GoCompiler` so the generated API
	  remains stable while ownership shrinks.
**/
@:goAllowRaw
@:keep
class GoHttpHelpers {
	/**
		What
		- Returns all response header values for a header key, preserving the
		  current haxe.go lowercase fallback behavior.

		Why
		- `responseHeadersSameKey` is stored in a generated Go-native map and is
		  not consistently visible from the staged Haxe surface, even though the
		  request choreography already populates it.

		How
		- Read the generated map through a narrow raw bridge, then fall back to
		  `responseHeaders.get(...)` exactly as the compiler shim did before.
	**/
	public static function getResponseHeaderValues(self:Http, key:String):Null<Array<String>> {
		if (self == null) {
			return null;
		}
		var normalized:String = untyped __go__("func() *string { raw := *hxrt.StdString({0}); out := make([]byte, len(raw)); for i := 0; i < len(raw); i++ { c := raw[i]; if c >= 'A' && c <= 'Z' { c += 32 }; out[i] = c }; return hxrt.StringFromLiteral(string(out)) }()",
			key);
		return
			untyped __go__("func() []*string { if {0}.responseHeadersSameKey != nil { if values, ok := {0}.responseHeadersSameKey[*hxrt.StdString({1})]; ok { return values }; if values, ok := {0}.responseHeadersSameKey[*hxrt.StdString({2})]; ok { return values } }; if {0}.responseHeaders == nil { return nil }; single := {0}.responseHeaders.get({1}); if single == nil && *hxrt.StdString({1}) != *hxrt.StdString({2}) { single = {0}.responseHeaders.get({2}) }; if single == nil { return nil }; return []*string{hxrt.StdString(single)} }()",
			self, key, normalized);
	}

	/**
		What
		- Copies HTTP response bytes into the optional API sink supplied to
		  `customRequest`.

		Why
		- The leaf fan-out uses generated Go interface shapes (`add`, `writeBytes`,
		  `writeFullBytes`, `writeString`) that are simpler to preserve through a
		  raw type switch than by re-growing compiler-owned helper blocks.

		How
		- Keep the dispatch local to one framework-owned raw helper and leave the
		  higher-level request lifecycle in the compiler-owned request path.
	**/
	public static function captureApi(api:Dynamic, payload:Bytes):Void {
		if (api == null || payload == null) {
			return;
		}
		untyped __go__("func() int { switch out := {0}.(type) { case *haxe__io__BytesBuffer: out.add({1}); case interface{ add(*haxe__io__Bytes) }: out.add({1}); case interface{ writeBytes(*haxe__io__Bytes, int, int) int }: out.writeBytes({1}, 0, {1}.length); case interface{ writeFullBytes(*haxe__io__Bytes, int, int) }: out.writeFullBytes({1}, 0, {1}.length); case interface{ writeString(*string) }: out.writeString({1}.toString()) }; return 0 }()",
			api, payload);
	}
}
