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
		What: Installs one declared-size multipart file and its bounded chunk reader.
		Why: `sys.Http.fileTransfer` must retain Haxe `Input` semantics without
		exposing its generated layout or buffering the entire payload in `hxrt`.
		How: Native transport asks the callback for at most the supplied byte count
		per read and stops after exactly `size` bytes.
	**/
	@:go.name("HttpRequestSetMultipartUpload")
	public static function setMultipartUpload(request:HttpRequestHandle, parameter:String, filename:String, mimeType:String, size:Int,
		readChunk:Int->Null<ByteView>):Void;

	@:go.name("HttpRequestSetProxy")
	public static function setProxy(request:HttpRequestHandle, host:String, port:Int, user:Null<String>, pass:Null<String>):Void;

	@:go.name("HttpRequestSetSocket")
	public static function setSocket(request:HttpRequestHandle, socket:SocketHandle):Void;

	@:go.name("HttpRequestStartExchange")
	public static function startExchange(request:HttpRequestHandle):HttpExchangeHandle;

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
