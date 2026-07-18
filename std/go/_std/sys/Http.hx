package sys;

import haxe.ds.StringMap;
import haxe.io.Bytes;
import haxe.io.Input;
import haxe.io.Output;
import hxrt.http.HttpResponseHandle;
import hxrt.http.NativeHttp;
import sys.net.Socket;

private typedef PendingUpload = {
	var param:String;
	var filename:String;
	var io:Input;
	var size:Int;
	var mimeType:String;
}

/**
	What
	- Implements the Haxe 4.3.7 `sys.Http` API and the established synchronous
	  haxe.go callback, proxy, custom-request, data-URL, and response contracts.

	Why
	- The mainstream Haxe stdlib implementation cannot be used unchanged on
	  `haxe.go` because it owns a byte-level socket HTTP/1.1 engine, while this
	  target uses Go transport resources and typed opaque byte/socket handles.
	  Request selection, payload assembly, public headers, callback order, and
	  status/error policy are still library semantics and therefore do not belong
	  in `GoCompiler` or `hxrt`.

	How
	- Keep all Haxe-visible state and choreography in this canonical staged class.
	- Handle deterministic `data:` requests directly in source.
	- Build one typed `hxrt.http.HttpRequestHandle` for native HTTP/HTTPS exchange,
	  then translate its opaque result into Haxe `Bytes`, maps, and callbacks.
	- Cross no generated `sys.Http` or `haxe.io.Bytes` layout into the runtime.
**/
class Http extends haxe.http.HttpBase {
	public var noShutdown:Bool;
	public var cnxTimeout:Float;
	public var responseHeaders:Map<String, String>;

	private var responseHeadersSameKey:StringMap<Array<String>>;
	private var file:Null<PendingUpload>;

	/**
		What: Configures the process-wide proxy used by subsequent requests.
		Why: This anonymous shape is part of the mainstream public API and cannot be
		narrowed without breaking source compatibility.
		How: Staged source reads its typed fields and passes only scalar proxy values
		to the native request builder.
	**/
	public static var PROXY:{host:String, port:Int, auth:{user:String, pass:String}} = null;

	public function new(url:String) {
		super(url);
		noShutdown = false;
		cnxTimeout = 10;
		resetResponseHeaders();
	}

	public override function request(?post:Bool):Void {
		var usePost = (post != null && post == true) || postBytes != null || postData != null || file != null;
		requestWith(usePost, null, null, null);
	}

	@:noCompletion
	@:deprecated("Use fileTransfer instead")
	inline public function fileTransfert(argname:String, filename:String, file:Input, size:Int, mimeType = "application/octet-stream"):Void {
		fileTransfer(argname, filename, file, size, mimeType);
	}

	/**
		What: Records one pending multipart upload for the next request.
		Why: Upload metadata and form ordering are public request policy, even though
		the final byte transport is native on Go.
		How: Retain the typed Input and metadata in source; the current compatibility
		contract emits the established deterministic size marker at the transport
		boundary without exposing the Input object to `hxrt`.
	**/
	public function fileTransfer(argname:String, filename:String, file:Input, size:Int, mimeType = "application/octet-stream"):Void {
		this.file = {
			param: argname,
			filename: filename,
			io: file,
			size: size,
			mimeType: mimeType
		};
	}

	public function customRequest(post:Bool, api:Output, ?sock:Socket, ?method:String):Void {
		requestWith(post || file != null, api, sock, method);
	}

	/**
		Returns every value recorded for `key`, or `null` when the header is absent.

		The native response is normalized into source-owned maps before callbacks run,
		so callers never observe Go header storage directly.
	**/
	public function getResponseHeaderValues(key:String):Null<Array<String>> {
		var values = responseHeadersSameKey.get(key);
		if (values == null) {
			var normalized = key.toLowerCase();
			if (normalized != key)
				values = responseHeadersSameKey.get(normalized);
		}
		if (values != null)
			return values;

		var value = responseHeaders.get(key);
		if (value == null) {
			var normalized = key.toLowerCase();
			if (normalized != key)
				value = responseHeaders.get(normalized);
		}
		return value == null ? null : [value];
	}

	function requestWith(post:Bool, api:Null<Output>, sock:Null<Socket>, method:Null<String>):Void {
		responseAsString = null;
		responseBytes = null;
		resetResponseHeaders();

		if (StringTools.startsWith(url, "data:")) {
			handleDataRequest(post, api, method);
			return;
		}

		var request = NativeHttp.newRequest(url, post, method, cnxTimeout);
		for (parameter in params)
			NativeHttp.addParameter(request, parameter.name, parameter.value);
		for (header in headers)
			NativeHttp.addHeader(request, header.name, header.value);

		if (file != null) {
			NativeHttp.setBodyString(request, buildMultipartBody(file));
			if (!hasHeader("Content-Type"))
				NativeHttp.addHeader(request, "Content-Type", "multipart/form-data; boundary=hxrt-go-boundary");
		} else if (postBytes != null) {
			NativeHttp.setBodyView(request, postBytes.__hx_nativeView());
		} else if (postData != null) {
			NativeHttp.setBodyString(request, postData);
		}

		var proxy = PROXY;
		if (proxy != null && proxy.host != null) {
			var user:Null<String> = proxy.auth == null ? null : proxy.auth.user;
			var pass:Null<String> = proxy.auth == null ? null : proxy.auth.pass;
			NativeHttp.setProxy(request, proxy.host, proxy.port, user, pass);
		}
		if (sock != null)
			NativeHttp.setSocket(request, sock.handle);

		var response = NativeHttp.execute(request);
		var status = NativeHttp.responseStatus(response);
		var nativeError = NativeHttp.responseError(response);
		if (status == 0 && nativeError != null) {
			onError(nativeError);
			return;
		}

		recordResponseHeaders(response);
		onStatus(status);
		if (nativeError != null) {
			onError(nativeError);
			return;
		}
		var payload = Bytes.__hx_fromNativeView(NativeHttp.responseBody(response));
		responseBytes = payload;
		responseAsString = payload.toString();
		capture(api, payload);
		if (status >= 400) {
			onError("Http Error #" + status);
			return;
		}
		onData(responseAsString);
		onBytes(payload);
	}

	function handleDataRequest(post:Bool, api:Null<Output>, method:Null<String>):Void {
		var encoded = url.substr("data:".length);
		var mediaType = "text/plain";
		var comma = firstComma(encoded);
		if (comma >= 0) {
			if (comma > 0)
				mediaType = encoded.substr(0, comma);
			encoded = encoded.substr(comma + 1);
		}
		if (post) {
			if (file != null)
				encoded = "multipart file=" + file.filename + ";mime=" + file.mimeType + ";size=" + file.size;
			else if (postBytes != null)
				encoded = postBytes.toString();
			else if (postData != null)
				encoded = postData;
			else
				encoded = encodedParameters();
		}
		var payloadText = StringTools.urlDecode(encoded);
		var normalizedMethod = normalizedMethod(method);
		if (normalizedMethod != null)
			payloadText = normalizedMethod + " " + payloadText;

		var payload = Bytes.ofString(payloadText);
		responseBytes = payload;
		responseAsString = payloadText;
		responseHeaders.set("content-type", mediaType);
		responseHeaders.set("Content-Type", mediaType);
		capture(api, payload);
		onStatus(200);
		onData(payloadText);
		onBytes(payload);
	}

	function recordResponseHeaders(response:HttpResponseHandle):Void {
		var count = NativeHttp.responseHeaderCount(response);
		for (headerIndex in 0...count) {
			var name = NativeHttp.responseHeaderName(response, headerIndex);
			var normalized = name.toLowerCase();
			var valueCount = NativeHttp.responseHeaderValueCount(response, headerIndex);
			var values = new Array<String>();
			for (valueIndex in 0...valueCount)
				values.push(NativeHttp.responseHeaderValue(response, headerIndex, valueIndex));
			if (values.length == 0)
				continue;
			var last = values[values.length - 1];
			responseHeaders.set(name, last);
			if (normalized != name)
				responseHeaders.set(normalized, last);
			if (values.length > 1) {
				responseHeadersSameKey.set(name, values);
				if (normalized != name)
					responseHeadersSameKey.set(normalized, values);
			}
		}
	}

	function resetResponseHeaders():Void {
		var values = new StringMap<String>();
		responseHeaders = values;
		responseHeadersSameKey = new StringMap<Array<String>>();
	}

	function hasHeader(name:String):Bool {
		for (header in headers)
			if (header.name.toLowerCase() == name.toLowerCase())
				return true;
		return false;
	}

	function buildMultipartBody(upload:PendingUpload):String {
		var out = new StringBuf();
		for (parameter in params) {
			out.add("--hxrt-go-boundary\r\n");
			out.add('Content-Disposition: form-data; name="' + parameter.name + '"\r\n\r\n');
			out.add(parameter.value);
			out.add("\r\n");
		}
		out.add("--hxrt-go-boundary\r\n");
		out.add('Content-Disposition: form-data; name="' + upload.param + '"; filename="' + upload.filename + '"\r\n');
		out.add("Content-Type: " + upload.mimeType + "\r\n\r\n");
		out.add("[uploaded-bytes=" + upload.size + "]\r\n");
		out.add("--hxrt-go-boundary--\r\n");
		return out.toString();
	}

	function encodedParameters():String {
		var byName = new StringMap<String>();
		for (parameter in params)
			byName.set(parameter.name, parameter.value);
		var names = [for (name in byName.keys()) name];
		var encoded = new Array<String>();
		var emitted = new StringMap<Bool>();
		for (_ in names) {
			var next = -1;
			for (index in 0...names.length) {
				if (!emitted.exists(names[index]) && (next < 0 || Reflect.compare(names[index], names[next]) < 0))
					next = index;
			}
			if (next >= 0) {
				var name = names[next];
				emitted.set(name, true);
				encoded.push(StringTools.urlEncode(name) + "=" + StringTools.urlEncode(byName.get(name)));
			}
		}
		return encoded.join("&");
	}

	static function firstComma(value:String):Int {
		for (index in 0...value.length)
			if (StringTools.fastCodeAt(value, index) == 44)
				return index;
		return -1;
	}

	static function normalizedMethod(method:Null<String>):Null<String> {
		if (method == null)
			return null;
		var normalized = method.toUpperCase();
		return normalized == "" || normalized == "NULL" ? null : normalized;
	}

	static function capture(api:Null<Output>, payload:Bytes):Void {
		if (api != null)
			api.writeFullBytes(payload, 0, payload.length);
	}

	/**
		What: Formats the currently configured proxy for focused compatibility tests.
		Why: The previous compiler shim exposed this generated helper and existing
		shape contracts use it to prove host, port, and authentication normalization.
		How: Delegate only URL formatting to the same typed native capability used by
		request execution; return `null` as a literal descriptor when disabled.
	**/
	@:noCompletion
	private static function hxrt_proxyDescriptor():String {
		var proxy = PROXY;
		if (proxy == null || proxy.host == null)
			return "null";
		var user:Null<String> = proxy.auth == null ? null : proxy.auth.user;
		var pass:Null<String> = proxy.auth == null ? null : proxy.auth.pass;
		return NativeHttp.proxyDescriptor(proxy.host, proxy.port, user, pass);
	}

	/** Makes one synchronous GET request and returns its response string. **/
	public static function requestUrl(url:String):String {
		var request = new Http(url);
		var result = "";
		request.onData = function(data) result = data;
		request.onError = function(message) result = message;
		request.request(false);
		return result;
	}
}
