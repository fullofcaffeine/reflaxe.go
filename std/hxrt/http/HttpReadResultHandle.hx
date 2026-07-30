package hxrt.http;

/**
	What
	- Opaque typed result of one bounded native HTTP response-body read.

	Why
	- Go readers may return useful bytes and a terminal error together. Exposing
	  only one or the other would either lose progress or hide failure.

	How
	- Map directly to `hxrt.HttpReadResult`; staged source reads the immutable
	  bytes first, then checks the error and clean-EOF accessors.
**/
@:go.import("hxrt")
@:go.package("hxrt")
@:go.name("HttpReadResult")
extern class HttpReadResultHandle {}
