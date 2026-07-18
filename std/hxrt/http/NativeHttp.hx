package hxrt.http;

import hxrt.io.ByteView;
import hxrt.net.SocketHandle;

/**
	What
	- Typed access to Go URL, HTTP transport, proxy, and response capabilities.

	Why
	- Native networking and live socket resources cannot be implemented as
	  portable Haxe data. Request selection, payload construction, public header
	  maps, callbacks, and error/status policy remain ordinary staged Haxe.

	How
	- Build one opaque request from strings, scalars, a `ByteView`, and an optional
	  typed socket handle. Execute synchronously, then expose the response through
	  typed scalar and indexed accessors without depending on generated Haxe
	  object layouts.
**/
@:go.import("hxrt")
@:go.package("hxrt")
extern class NativeHttp {
	@:go.name("HttpRequestNew")
	public static function newRequest(url:String, post:Bool, method:Null<String>, timeout:Float):HttpRequestHandle;

	@:go.name("HttpRequestAddParameter")
	public static function addParameter(request:HttpRequestHandle, name:String, value:String):Void;

	@:go.name("HttpRequestAddHeader")
	public static function addHeader(request:HttpRequestHandle, name:String, value:String):Void;

	@:go.name("HttpRequestSetBodyString")
	public static function setBodyString(request:HttpRequestHandle, value:String):Void;

	@:go.name("HttpRequestSetBodyView")
	public static function setBodyView(request:HttpRequestHandle, value:ByteView):Void;

	@:go.name("HttpRequestSetProxy")
	public static function setProxy(request:HttpRequestHandle, host:String, port:Int, user:Null<String>, pass:Null<String>):Void;

	@:go.name("HttpRequestSetSocket")
	public static function setSocket(request:HttpRequestHandle, socket:SocketHandle):Void;

	@:go.name("HttpRequestExecute")
	public static function execute(request:HttpRequestHandle):HttpResponseHandle;

	@:go.name("HttpResponseError")
	public static function responseError(response:HttpResponseHandle):Null<String>;

	@:go.name("HttpResponseStatus")
	public static function responseStatus(response:HttpResponseHandle):Int;

	@:go.name("HttpResponseBody")
	public static function responseBody(response:HttpResponseHandle):ByteView;

	@:go.name("HttpResponseHeaderCount")
	public static function responseHeaderCount(response:HttpResponseHandle):Int;

	@:go.name("HttpResponseHeaderName")
	public static function responseHeaderName(response:HttpResponseHandle, index:Int):String;

	@:go.name("HttpResponseHeaderValueCount")
	public static function responseHeaderValueCount(response:HttpResponseHandle, index:Int):Int;

	@:go.name("HttpResponseHeaderValue")
	public static function responseHeaderValue(response:HttpResponseHandle, headerIndex:Int, valueIndex:Int):String;

	@:go.name("HttpProxyDescriptor")
	public static function proxyDescriptor(host:String, port:Int, user:Null<String>, pass:Null<String>):String;
}
