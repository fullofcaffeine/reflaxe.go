package hxrt.http;

import hxrt.io.ByteView;
import hxrt.net.SocketHandle;

/**
	What
	- Typed access to Go URL, HTTP transport, proxy, and live response capabilities.

	Why
	- Native networking and live socket resources cannot be implemented as
	  portable Haxe data. Request selection, payload construction, public header
	  maps, callbacks, and error/status policy remain ordinary staged Haxe.

	How
	- Build one opaque request from strings, scalars, a `ByteView`, and an optional
	  typed socket handle. Start a live exchange, then expose headers and bounded
	  body-read results through typed accessors without depending on generated
	  Haxe object layouts.
**/
@:go.import("hxrt")
@:go.package("hxrt")
extern class NativeHttp {
	/**
		What: Starts one exact-method request description with a progress budget.
		Why: Go defaults for method normalization and whole-client timeouts do not
		match the staged `sys.Http` contract.
		How: Preserve the nullable method token verbatim; negative timeout disables
		native deadlines, zero is immediate, and positive values are per operation.
	**/
	@:go.name("HttpRequestNew")
	public static function newRequest(url:String, post:Bool, method:Null<String>, timeout:Float):HttpRequestHandle;

	/**
		What: Appends one source-ordered parameter in raw and percent-encoded form.
		Why: `HttpBase` preserves repeated entries, while `StringTools.urlEncode`
		is Haxe-visible policy that must not be replaced by Go map/sort/escaping.
		How: Native multipart uses the raw pair; query and form serialization join
		the encoded pair without reparsing or collapsing it.
	**/
	@:go.name("HttpRequestAddParameter")
	public static function addParameter(request:HttpRequestHandle, name:String, value:String, encodedName:String, encodedValue:String):Void;

	@:go.name("HttpRequestAddHeader")
	public static function addHeader(request:HttpRequestHandle, name:String, value:String):Void;

	@:go.name("HttpRequestSetBodyString")
	public static function setBodyString(request:HttpRequestHandle, value:String):Void;

	@:go.name("HttpRequestSetBodyView")
	public static function setBodyView(request:HttpRequestHandle, value:ByteView):Void;

	/**
		What: Installs scalar metadata for one declared-size multipart file.
		Why: `sys.Http.fileTransfer` must retain Haxe `Input` semantics without
		exposing its generated layout, buffering the payload, or letting native
		transport call generated Haxe.
		How: Exchange start creates a pipe-backed native body; staged source pumps
		bounded immutable chunks through the typed sink below.
	**/
	@:go.name("HttpRequestSetMultipartUpload")
	public static function setMultipartUpload(request:HttpRequestHandle, parameter:String, filename:String, mimeType:String, size:Int):Void;

	/**
		What: Selects a native HTTP proxy for this exchange.
		Why: HTTP absolute-target and HTTPS CONNECT behavior require Go transport
		ownership and must not be reconstructed from generated Haxe layouts.
		How: Pass only scalar authority; staged source still owns the public PROXY
		shape and callback policy.
	**/
	@:go.name("HttpRequestSetProxy")
	public static function setProxy(request:HttpRequestHandle, host:String, port:Int, user:Null<String>, pass:Null<String>):Void;

	/**
		What: Supplies a typed socket for plain-HTTP customRequest.
		Why: The current handle does not retain the `sys.ssl.Socket` configuration
		needed to distinguish already-secure HTTPS transport from a plain TCP socket.
		How: Native code consumes and closes the handle for HTTP, and rejects HTTPS
		before transport until a typed secure connector boundary exists.
	**/
	@:go.name("HttpRequestSetSocket")
	public static function setSocket(request:HttpRequestHandle, socket:SocketHandle):Void;

	@:go.name("HttpRequestStartExchange")
	public static function startExchange(request:HttpRequestHandle):HttpExchangeHandle;

	/**
		What: Returns the caller-owned writer for a multipart exchange.
		Why: Staged Haxe must pump its `Input` without exposing the native pipe.
		How: Return `null` for non-multipart exchanges and one opaque sink otherwise.
	**/
	@:go.name("HttpExchangeUploadSink")
	public static function exchangeUploadSink(exchange:HttpExchangeHandle):Null<HttpUploadSinkHandle>;

	/**
		What: Copies one non-empty immutable chunk into the native request body.
		Why: The write must preserve exact declared size and unblock on cancellation.
		How: Return a synchronized terminal message instead of throwing across the
		typed source/runtime boundary.
	**/
	@:go.name("HttpUploadSinkWriteChunk")
	public static function writeUploadChunk(sink:HttpUploadSinkHandle, chunk:ByteView):Null<String>;

	/**
		What: Marks exact-size upload completion.
		Why: Native multipart framing may emit its closing delimiter only after the
		caller supplied every declared byte.
		How: Close the pipe writer normally, or return the deterministic size error.
	**/
	@:go.name("HttpUploadSinkFinish")
	public static function finishUpload(sink:HttpUploadSinkHandle):Null<String>;

	/**
		What: Terminates one upload with its source-owned error.
		Why: Early EOF and `Input` exceptions must release native transport without
		being replaced by the later pipe-close error.
		How: Close the pipe writer with `message`; repeated abort is idempotent.
	**/
	@:go.name("HttpUploadSinkAbort")
	public static function abortUpload(sink:HttpUploadSinkHandle, message:String):Void;

	/**
		What: Waits until native code publishes response headers or terminal error.
		Why: Multipart exchange start returns before the caller pumps its `Input`.
		How: Join only the native transport result; no generated Haxe callback runs
		while waiting.
	**/
	@:go.name("HttpExchangeAwaitResponse")
	public static function awaitResponse(exchange:HttpExchangeHandle):Void;

	@:go.name("HttpExchangeError")
	public static function exchangeError(exchange:HttpExchangeHandle):Null<String>;

	@:go.name("HttpExchangeStatus")
	public static function exchangeStatus(exchange:HttpExchangeHandle):Int;

	@:go.name("HttpExchangeContentLength")
	public static function exchangeContentLength(exchange:HttpExchangeHandle):Int;

	@:go.name("HttpExchangeHeaderCount")
	public static function exchangeHeaderCount(exchange:HttpExchangeHandle):Int;

	@:go.name("HttpExchangeHeaderName")
	public static function exchangeHeaderName(exchange:HttpExchangeHandle, index:Int):String;

	@:go.name("HttpExchangeHeaderValueCount")
	public static function exchangeHeaderValueCount(exchange:HttpExchangeHandle, index:Int):Int;

	@:go.name("HttpExchangeHeaderValue")
	public static function exchangeHeaderValue(exchange:HttpExchangeHandle, headerIndex:Int, valueIndex:Int):String;

	@:go.name("HttpExchangeReadResponseChunk")
	public static function readResponseChunk(exchange:HttpExchangeHandle, maxBytes:Int):HttpReadResultHandle;

	@:go.name("HttpReadResultBody")
	public static function readResultBody(result:HttpReadResultHandle):ByteView;

	@:go.name("HttpReadResultError")
	public static function readResultError(result:HttpReadResultHandle):Null<String>;

	@:go.name("HttpReadResultEOF")
	public static function readResultEof(result:HttpReadResultHandle):Bool;

	@:go.name("HttpExchangeClose")
	public static function closeExchange(exchange:HttpExchangeHandle):Void;

	@:go.name("HttpExchangeCancel")
	public static function cancelExchange(exchange:HttpExchangeHandle):Void;

	@:go.name("HttpProxyDescriptor")
	public static function proxyDescriptor(host:String, port:Int, user:Null<String>, pass:Null<String>):String;
}
