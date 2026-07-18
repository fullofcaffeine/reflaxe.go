package hxrt

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
)

// CryptoBase64Encode performs only the native byte-codec operation requested by
// staged haxe.crypto.Base64. Padding defaults and URL-safe policy remain in Haxe.
func CryptoBase64Encode(values *ByteView, urlSafe bool) *string {
	encoding := base64.RawStdEncoding
	if urlSafe {
		encoding = base64.RawURLEncoding
	}
	return StringFromLiteral(encoding.EncodeToString(byteViewRaw(values)))
}

// CryptoBase64Decode decodes an unpadded standard or URL-safe Base64 payload.
// Invalid input crosses the ordinary hxrt exception carrier.
func CryptoBase64Decode(value *string, urlSafe bool) *ByteView {
	encoding := base64.RawStdEncoding
	if urlSafe {
		encoding = base64.RawURLEncoding
	}
	decoded, err := encoding.DecodeString(*StdString(value))
	if err != nil {
		Throw(err)
		return &ByteView{raw: []byte{}}
	}
	return &ByteView{raw: decoded}
}

func CryptoMd5String(value *string) *string {
	sum := md5.Sum([]byte(*StdString(value)))
	return StringFromLiteral(hex.EncodeToString(sum[:]))
}

func CryptoMd5Values(values *ByteView) *ByteView {
	sum := md5.Sum(byteViewRaw(values))
	return &ByteView{raw: sum[:]}
}

func CryptoSha1String(value *string) *string {
	sum := sha1.Sum([]byte(*StdString(value)))
	return StringFromLiteral(hex.EncodeToString(sum[:]))
}

func CryptoSha1Values(values *ByteView) *ByteView {
	sum := sha1.Sum(byteViewRaw(values))
	return &ByteView{raw: sum[:]}
}

func CryptoSha224String(value *string) *string {
	sum := sha256.Sum224([]byte(*StdString(value)))
	return StringFromLiteral(hex.EncodeToString(sum[:]))
}

func CryptoSha224Values(values *ByteView) *ByteView {
	sum := sha256.Sum224(byteViewRaw(values))
	return &ByteView{raw: sum[:]}
}

func CryptoSha256String(value *string) *string {
	sum := sha256.Sum256([]byte(*StdString(value)))
	return StringFromLiteral(hex.EncodeToString(sum[:]))
}

func CryptoSha256Values(values *ByteView) *ByteView {
	sum := sha256.Sum256(byteViewRaw(values))
	return &ByteView{raw: sum[:]}
}

func byteViewRaw(view *ByteView) []byte {
	if view == nil {
		return []byte{}
	}
	return view.raw
}
