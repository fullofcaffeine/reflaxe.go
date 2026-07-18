package hxrt.ssl;

/**
	What: Binds one opaque parsed native public or private key.
	Why: Go key implementations vary by algorithm and cannot cross the boundary as
	portable data without losing type and ownership guarantees.
	How: Staged Key retains this handle and typed digest/TLS capabilities consume it.
**/
@:go.import("hxrt")
@:go.package("hxrt")
@:go.name("SslKey")
extern class KeyHandle {}
