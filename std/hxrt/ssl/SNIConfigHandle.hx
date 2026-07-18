package hxrt.ssl;

/**
	What: Binds the TLS server's opaque SNI matcher and certificate table.
	Why: Go callbacks and native certificates cannot be represented as portable Haxe
	data, while staged Socket must retain the public configuration workflow.
	How: Typed TLS capabilities create and consume the handle without exposing its state.
**/
@:go.import("hxrt")
@:go.package("hxrt")
@:go.name("SslSocketSNIConfig")
extern class SNIConfigHandle {}
