package sys;

import haxe.ds.StringMap;
import haxe.io.Bytes;
import haxe.io.BytesOutput;
import haxe.io.Input;
import haxe.io.Output;
import hxrt.http.HttpExchangeHandle;
import hxrt.http.HttpUploadSinkHandle;
import hxrt.http.NativeHttp;
import hxrt.io.ByteView;
import sys.net.Socket;

private typedef PendingUpload = {
	var param:String;
	var filename:String;
	var io:Input;
	var size:Int;
	var mimeType:String;
}

/**
	What: Keeps staged source failure separate from native sink failure.
	Why: A source exception must remain the public error that caused abort, while
	an already published early response must not be hidden by pipe closure.
	How: `pumpUpload` returns both facts and `requestWith` applies the documented
	source → response → sink precedence after awaiting native publication.
**/
private typedef UploadPumpResult = {
	var sourceError:Null<String>;
	var sinkError:Null<String>;
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
	  then stream its live opaque exchange through Haxe `Output`, maps, and callbacks.
	- Cross no generated `sys.Http` or `haxe.io.Bytes` layout into the runtime.
**/
class Http extends haxe.http.HttpBase {
	public var noShutdown:Bool;

	/**
		What: Sets the native progress budget in seconds.
		Why: A whole-request deadline would reject healthy long transfers, while
		silently treating zero or negative values as ten seconds changes Haxe policy.
		How: Negative disables native deadlines, zero fails before dialing, and a
		positive value separately bounds direct connect/TLS/header waits, multipart
		sink writes, and response reads. Proxy negotiation and fixed-body write
		admission remain explicitly outside the current portable claim.
	**/
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

	/**
		What: Runs the callback-oriented request through a source-owned BytesOutput.
		Why: `request()` must retain every response byte written before an error,
		while direct `customRequest` must not publish response fields or callbacks.
		How: Temporarily wrap `onError` to snapshot partial bytes, delegate the live
		exchange to `customRequest`, then call `success` only after complete success.
	**/
	public override function request(?post:Bool):Void {
		var output = new BytesOutput();
		var previousOnError = onError;
		var failed = false;
		onError = function(message:String) {
			responseBytes = output.getBytes();
			responseAsString = null;
			failed = true;
			onError = previousOnError;
			onError(message);
		};
		var usePost = (post != null && post == true) || postBytes != null || postData != null || file != null;
		customRequest(usePost, output);
		if (!failed)
			success(output.getBytes());
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
		How: Retain the typed Input and metadata in source. During synchronous
		execution, read bounded immutable chunks on the public caller and pump them
		into the cancellable native sink without exposing the generated Input object
		or buffering the whole file.
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

	/**
		What: Streams one request through the caller-owned `Output`.
		Why: Status, prepare, partial writes, status classification, close, and
		source exceptions are Haxe-visible lifecycle semantics that a fully buffered
		native response cannot preserve.
		How: Read bounded immutable chunks from one typed live exchange, write each
		chunk before interpreting its terminal state, and release or cancel native
		resources before dispatching one source-owned error.
	**/
	function requestWith(post:Bool, api:Output, sock:Null<Socket>, method:Null<String>):Void {
		responseAsString = null;
		responseBytes = null;
		resetResponseHeaders();

		if (StringTools.startsWith(url, "data:")) {
			handleDataRequest(post, api, method);
			return;
		}

		var request = NativeHttp.newRequest(url, post, method, cnxTimeout);
		for (parameter in params)
			NativeHttp.addParameter(request, parameter.name, parameter.value, StringTools.urlEncode(parameter.name), StringTools.urlEncode(parameter.value));
		for (header in headers)
			NativeHttp.addHeader(request, header.name, header.value);

		var upload = file;
		if (upload != null) {
			NativeHttp.setMultipartUpload(request, upload.param, upload.filename, upload.mimeType, upload.size);
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

		var exchange = NativeHttp.startExchange(request);
		var uploadResult:Null<UploadPumpResult> = upload == null ? null : pumpUpload(exchange, upload);
		NativeHttp.awaitResponse(exchange);
		var sourceError = uploadResult == null ? null : uploadResult.sourceError;
		var sinkError = uploadResult == null ? null : uploadResult.sinkError;
		var errorMessage:Null<String> = sourceError;
		var completed = false;
		if (errorMessage == null) {
			var nativeError = NativeHttp.exchangeError(exchange);
			if (nativeError != null) {
				errorMessage = nativeError;
			} else if (NativeHttp.exchangeStatus(exchange) == 0 && sinkError != null) {
				errorMessage = sinkError;
			} else {
				try {
					recordResponseHeaders(exchange);
					var status = NativeHttp.exchangeStatus(exchange);
					onStatus(status);
					var contentLength = NativeHttp.exchangeContentLength(exchange);
					if (contentLength == -2)
						throw "Content-Length exceeds Haxe Int range";
					if (contentLength >= 0)
						api.prepare(contentLength);

					while (true) {
						var read = NativeHttp.readResponseChunk(exchange, 1024);
						var payload = Bytes.__hx_fromNativeView(NativeHttp.readResultBody(read));
						if (payload.length > 0)
							api.writeBytes(payload, 0, payload.length);
						var readError = NativeHttp.readResultError(read);
						if (readError != null)
							throw "Transfer aborted";
						if (NativeHttp.readResultEof(read))
							break;
					}

					var statusError = hxrt_statusError(status);
					if (statusError != null)
						throw statusError;
					api.close();
					completed = true;
				} catch (error:haxe.Exception) {
					errorMessage = error.message;
				}
			}
		}

		if (completed)
			NativeHttp.closeExchange(exchange);
		else
			NativeHttp.cancelExchange(exchange);
		if (errorMessage != null)
			onError(errorMessage);
	}

	/**
		What: Pumps one staged multipart Input into its typed native sink.
		Why: Calling generated Haxe from net/http's transport goroutine breaks the
		request body's cancellation contract and makes captured source state racy.
		How: Read bounded chunks synchronously on the public Haxe caller, preserve
		source failures, and finish or abort the exact-size native pipe explicitly.
		An arbitrary Input blocked inside its own read remains source-owned and
		cannot be forcibly canceled; once it returns, the next sink write observes
		native cancellation promptly.
	**/
	function pumpUpload(exchange:HttpExchangeHandle, upload:PendingUpload):UploadPumpResult {
		var sink:Null<HttpUploadSinkHandle> = NativeHttp.exchangeUploadSink(exchange);
		if (sink == null)
			return {sourceError: null, sinkError: "HTTP upload sink is unavailable"};

		var remaining = upload.size;
		var sourceError:Null<String> = null;
		var sinkError:Null<String> = null;
		while (remaining > 0) {
			var requested = remaining > 32768 ? 32768 : remaining;
			var chunk = Bytes.alloc(requested);
			var count = 0;
			try {
				count = upload.io.readBytes(chunk, 0, requested);
			} catch (_:haxe.io.Eof) {
				sourceError = "Transfer aborted";
			} catch (error:haxe.Exception) {
				sourceError = error.message;
			}
			if (sourceError != null)
				break;
			if (count <= 0) {
				sourceError = "multipart upload made no progress";
				break;
			}
			if (count > requested) {
				sourceError = "multipart upload exceeded the requested chunk size";
				break;
			}
			if (count < requested)
				chunk = chunk.sub(0, count);
			sinkError = NativeHttp.writeUploadChunk(sink, chunk.__hx_nativeView());
			if (sinkError != null)
				break;
			remaining -= count;
		}

		if (sourceError != null) {
			NativeHttp.abortUpload(sink, sourceError);
		} else if (sinkError == null) {
			sinkError = NativeHttp.finishUpload(sink);
		}
		return {sourceError: sourceError, sinkError: sinkError};
	}

	/**
		What: Applies the same Output lifecycle to the deterministic `data:` path.
		Why: Callers should not observe different status/prepare/write/close order
		merely because bytes came from a URL literal instead of a network response.
		How: Assemble the source-owned payload using the same explicit-body and exact
		method policy, then run status, prepare, one bounded write, and close under
		the same single-error dispatch rule.
	**/
	function handleDataRequest(post:Bool, api:Output, method:Null<String>):Void {
		if (method != null && method == "") {
			onError("HTTP method must not be empty");
			return;
		}
		var encoded = url.substr("data:".length);
		var mediaType = "text/plain";
		var comma = firstComma(encoded);
		if (comma >= 0) {
			if (comma > 0)
				mediaType = encoded.substr(0, comma);
			encoded = encoded.substr(comma + 1);
		}
		if (file != null)
			encoded = "multipart file=" + file.filename + ";mime=" + file.mimeType + ";size=" + file.size;
		else if (postBytes != null)
			encoded = postBytes.toString();
		else if (postData != null)
			encoded = postData;
		else if (post)
			encoded = encodedParameters();
		var payloadText = StringTools.urlDecode(encoded);
		var explicitMethod = explicitMethod(method);
		if (explicitMethod != null)
			payloadText = explicitMethod + " " + payloadText;

		var payload = Bytes.ofString(payloadText);
		responseHeaders.set("content-type", mediaType);
		responseHeaders.set("Content-Type", mediaType);
		var errorMessage:Null<String> = null;
		try {
			onStatus(200);
			api.prepare(payload.length);
			if (payload.length > 0)
				api.writeBytes(payload, 0, payload.length);
			api.close();
		} catch (error:haxe.Exception) {
			errorMessage = error.message;
		}
		if (errorMessage != null)
			onError(errorMessage);
	}

	function recordResponseHeaders(exchange:HttpExchangeHandle):Void {
		var count = NativeHttp.exchangeHeaderCount(exchange);
		for (headerIndex in 0...count) {
			var name = NativeHttp.exchangeHeaderName(exchange, headerIndex);
			var normalized = name.toLowerCase();
			var valueCount = NativeHttp.exchangeHeaderValueCount(exchange, headerIndex);
			var values = new Array<String>();
			for (valueIndex in 0...valueCount)
				values.push(NativeHttp.exchangeHeaderValue(exchange, headerIndex, valueIndex));
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

	function encodedParameters():String {
		var encoded = new Array<String>();
		for (parameter in params)
			encoded.push(StringTools.urlEncode(parameter.name) + "=" + StringTools.urlEncode(parameter.value));
		return encoded.join("&");
	}

	static function firstComma(value:String):Int {
		for (index in 0...value.length)
			if (StringTools.fastCodeAt(value, index) == 44)
				return index;
		return -1;
	}

	/**
		What: Preserves a non-empty explicit customRequest method for the data path.
		Why: Method spelling is caller-owned wire policy; uppercasing it would make
		data and network requests disagree and can select another server handler.
		How: Distinguish absence from a supplied token without normalizing its case.
	**/
	static function explicitMethod(method:Null<String>):Null<String> {
		if (method == null || method == "")
			return null;
		return method;
	}

	/**
		What: Classifies every response status after its body has streamed.
		Why: Haxe sys.Http treats statuses below 200 and at least 400 as errors,
		while Go transport policy must not silently redefine public callbacks.
		How: Return the one public error string for rejected ranges and `null` for
		the complete 200...399 success interval.
	**/
	@:noCompletion
	private static function hxrt_statusError(status:Int):Null<String> {
		return status < 200 || status >= 400 ? "Http Error #" + status : null;
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

	/**
		Makes one synchronous GET request and returns its response string.

		Errors are thrown, matching the Haxe 4.3.7 `sys.Http` contract; an error
		message is never returned as though it were response data.
	**/
	public static function requestUrl(url:String):String {
		var request = new Http(url);
		var result = "";
		request.onData = function(data) result = data;
		request.onError = function(message) throw message;
		request.request(false);
		return result;
	}
}
