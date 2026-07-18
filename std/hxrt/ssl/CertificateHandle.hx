package hxrt.ssl;

/**
	What: Binds the opaque parsed certificate or trust-pool handle owned by Go TLS.
	Why: Native certificate state has no safe portable Haxe representation and must
	not be weakened to `Dynamic`.
	How: Staged `sys.ssl.Certificate` retains this handle and passes it only to typed hxrt APIs.
**/
@:go.import("hxrt")
@:go.package("hxrt")
@:go.name("SslCertificate")
extern class CertificateHandle {}
