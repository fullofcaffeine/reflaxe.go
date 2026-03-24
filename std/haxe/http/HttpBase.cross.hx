package haxe.http;

import haxe.io.Bytes;

private typedef StringKeyValue = {
	var name:String;
	var value:String;
}

/**
	What
	A staged `haxe.http.HttpBase` override for `haxe.go`.

	Why
	Direct `haxe.http.HttpBase` usage is part of the upstream stdlib surface, but
	`haxe.go` previously had no staged owner for it. That left the type available
	to the typer while generated Go missed the emitted class body entirely. The
	real request engine for sys targets still lives in target-owned `sys.Http`;
	this override only restores the portable base-class contract.

	How
	Keep the upstream field layout, callback shape, and default request behavior
	in ordinary Haxe code. `request()` still throws
	`haxe.exceptions.NotImplementedException`, matching the upstream base class,
	while subclasses such as `sys.Http` remain free to override it with real
	target-native behavior.
**/
class HttpBase {
	public var url:String;

	public var responseData(get, never):Null<String>;
	public var responseBytes(default, null):Null<Bytes>;

	var responseAsString:Null<String>;
	var postData:Null<String>;
	var postBytes:Null<Bytes>;
	var headers:Array<StringKeyValue>;
	var params:Array<StringKeyValue>;
	final emptyOnData:(String) -> Void;

	public function new(url:String) {
		this.url = url;
		headers = [];
		params = [];
		emptyOnData = onData;
	}

	public function setHeader(name:String, value:String) {
		for (i in 0...headers.length) {
			if (headers[i].name == name) {
				headers[i] = {name: name, value: value};
				return #if hx3compat this #end;
			}
		}
		headers.push({name: name, value: value});
		#if hx3compat
		return this;
		#end
	}

	public function addHeader(header:String, value:String) {
		headers.push({name: header, value: value});
		#if hx3compat
		return this;
		#end
	}

	public function setParameter(name:String, value:String) {
		for (i in 0...params.length) {
			if (params[i].name == name) {
				params[i] = {name: name, value: value};
				return #if hx3compat this #end;
			}
		}
		params.push({name: name, value: value});
		#if hx3compat
		return this;
		#end
	}

	public function addParameter(name:String, value:String) {
		params.push({name: name, value: value});
		#if hx3compat
		return this;
		#end
	}

	public function setPostData(data:Null<String>) {
		postData = data;
		postBytes = null;
		#if hx3compat
		return this;
		#end
	}

	public function setPostBytes(data:Null<Bytes>) {
		postBytes = data;
		postData = null;
		#if hx3compat
		return this;
		#end
	}

	public function request(?post:Bool):Void {
		throw new haxe.exceptions.NotImplementedException();
	}

	public dynamic function onData(data:String) {}

	public dynamic function onBytes(data:Bytes) {}

	public dynamic function onError(msg:String) {}

	public dynamic function onStatus(status:Int) {}

	function hasOnData():Bool {
		return !Reflect.compareMethods(onData, emptyOnData);
	}

	function success(data:Bytes) {
		responseBytes = data;
		responseAsString = null;
		if (hasOnData()) {
			var s = responseData;
			if (s != null) {
				onData(s);
			}
		}
		onBytes(data);
	}

	function get_responseData():Null<String> {
		if (responseAsString == null && responseBytes != null) {
			responseAsString = responseBytes.getString(0, responseBytes.length, UTF8);
		}
		return responseAsString;
	}
}
