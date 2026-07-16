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
func CryptoBase64Encode(values []int, urlSafe bool) *string {
	encoding := base64.RawStdEncoding
	if urlSafe {
		encoding = base64.RawURLEncoding
	}
	return StringFromLiteral(encoding.EncodeToString(cryptoValuesToBytes(values)))
}

// CryptoBase64Decode decodes an unpadded standard or URL-safe Base64 payload.
// Invalid input crosses the ordinary hxrt exception carrier.
func CryptoBase64Decode(value *string, urlSafe bool) []int {
	encoding := base64.RawStdEncoding
	if urlSafe {
		encoding = base64.RawURLEncoding
	}
	decoded, err := encoding.DecodeString(*StdString(value))
	if err != nil {
		Throw(err)
		return []int{}
	}
	return cryptoBytesToValues(decoded)
}

func CryptoMd5String(value *string) *string {
	sum := md5.Sum([]byte(*StdString(value)))
	return StringFromLiteral(hex.EncodeToString(sum[:]))
}

func CryptoMd5Values(values []int) []int {
	sum := md5.Sum(cryptoValuesToBytes(values))
	return cryptoBytesToValues(sum[:])
}

func CryptoSha1String(value *string) *string {
	sum := sha1.Sum([]byte(*StdString(value)))
	return StringFromLiteral(hex.EncodeToString(sum[:]))
}

func CryptoSha1Values(values []int) []int {
	sum := sha1.Sum(cryptoValuesToBytes(values))
	return cryptoBytesToValues(sum[:])
}

func CryptoSha224String(value *string) *string {
	sum := sha256.Sum224([]byte(*StdString(value)))
	return StringFromLiteral(hex.EncodeToString(sum[:]))
}

func CryptoSha224Values(values []int) []int {
	sum := sha256.Sum224(cryptoValuesToBytes(values))
	return cryptoBytesToValues(sum[:])
}

func CryptoSha256String(value *string) *string {
	sum := sha256.Sum256([]byte(*StdString(value)))
	return StringFromLiteral(hex.EncodeToString(sum[:]))
}

func CryptoSha256Values(values []int) []int {
	sum := sha256.Sum256(cryptoValuesToBytes(values))
	return cryptoBytesToValues(sum[:])
}

func cryptoValuesToBytes(values []int) []byte {
	converted := make([]byte, len(values))
	for index, value := range values {
		converted[index] = byte(value)
	}
	return converted
}

func cryptoBytesToValues(values []byte) []int {
	converted := make([]int, len(values))
	for index, value := range values {
		converted[index] = int(value)
	}
	return converted
}
