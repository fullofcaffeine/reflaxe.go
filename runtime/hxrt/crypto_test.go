package hxrt

import (
	"encoding/hex"
	"testing"
)

func TestCryptoBase64Codecs(t *testing.T) {
	values := []int{0, 127, 128, 255}
	view := BytesViewFromValues(values)
	if got := *CryptoBase64Encode(view, false); got != "AH+A/w" {
		t.Fatalf("standard encode = %q", got)
	}
	if got := *CryptoBase64Encode(view, true); got != "AH-A_w" {
		t.Fatalf("URL-safe encode = %q", got)
	}
	if got := CryptoBase64Decode(StringFromLiteral("AH+A/w"), false); !equalCryptoViewValues(got, values) {
		t.Fatalf("standard decode = %#v", got)
	}
	if got := CryptoBase64Decode(StringFromLiteral("AH-A_w"), true); !equalCryptoViewValues(got, values) {
		t.Fatalf("URL-safe decode = %#v", got)
	}
}

func equalCryptoViewValues(left *ByteView, right []int) bool {
	leftValues := BytesValuesFromView(left)
	if len(leftValues) != len(right) {
		return false
	}
	for index := range leftValues {
		if leftValues[index] != right[index] {
			return false
		}
	}
	return true
}

func TestCryptoBase64InvalidInputThrows(t *testing.T) {
	deferred := false
	func() {
		defer func() {
			deferred = recover() != nil
		}()
		CryptoBase64Decode(StringFromLiteral("%%%"), false)
	}()
	if !deferred {
		t.Fatal("invalid Base64 input did not cross the hxrt exception carrier")
	}
}

func TestCryptoDigests(t *testing.T) {
	value := StringFromLiteral("ab")
	values := []int{'a', 'b'}
	tests := []struct {
		name       string
		want       string
		stringHash func(*string) *string
		valueHash  func(*ByteView) *ByteView
	}{
		{"md5", "187ef4436122d1cc2f40dc2b92f0eba0", CryptoMd5String, CryptoMd5Values},
		{"sha1", "da23614e02469a0d7c7bd1bdab5c9c474b1904dc", CryptoSha1String, CryptoSha1Values},
		{"sha224", "db3cda86d4429a1d39c148989566b38f7bda0156296bd364ba2f878b", CryptoSha224String, CryptoSha224Values},
		{"sha256", "fb8e20fc2e4c3f248c60c39bd652f3c1347298bb977b8b4d5903b85055620603", CryptoSha256String, CryptoSha256Values},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := *test.stringHash(value); got != test.want {
				t.Fatalf("string digest = %q", got)
			}
			if got := hex.EncodeToString(byteViewRaw(test.valueHash(BytesViewFromValues(values)))); got != test.want {
				t.Fatalf("byte digest = %q", got)
			}
		})
	}
}
