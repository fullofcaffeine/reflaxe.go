package hxrt

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"os"
	"path/filepath"
	"strings"
)

type SslKey struct {
	privateKey any
	publicKey  any
}

type SslCertificate struct {
	certs []*x509.Certificate
	pool  *x509.CertPool
	index int
}

func sslDigestName(alg *string) string {
	if alg == nil {
		return "SHA256"
	}
	return strings.ToUpper(*StdString(alg))
}

func sslDigestHash(name string) crypto.Hash {
	switch name {
	case "MD5":
		return crypto.MD5
	case "SHA1":
		return crypto.SHA1
	case "SHA224":
		return crypto.SHA224
	case "SHA256":
		return crypto.SHA256
	case "SHA384":
		return crypto.SHA384
	case "SHA512":
		return crypto.SHA512
	default:
		return 0
	}
}

func sslDigestBytes(data []byte, name string) []byte {
	switch name {
	case "MD5":
		sum := md5.Sum(data)
		return append([]byte(nil), sum[:]...)
	case "SHA1":
		sum := sha1.Sum(data)
		return append([]byte(nil), sum[:]...)
	case "SHA224":
		sum := sha256.Sum224(data)
		return append([]byte(nil), sum[:]...)
	case "SHA256":
		sum := sha256.Sum256(data)
		return append([]byte(nil), sum[:]...)
	case "SHA384":
		sum := sha512.Sum384(data)
		return append([]byte(nil), sum[:]...)
	case "SHA512":
		sum := sha512.Sum512(data)
		return append([]byte(nil), sum[:]...)
	case "RIPEMD160":
		Throw(StringFromLiteral("sys.ssl.Digest RIPEMD160 is not supported on haxe.go yet"))
		return []byte{}
	default:
		Throw(StringFromLiteral("Unsupported sys.ssl.Digest algorithm"))
		return []byte{}
	}
}

func SslDigestMake(data []byte, alg *string) []byte {
	return sslDigestBytes(data, sslDigestName(alg))
}

func parsePrivateDER(der []byte) (any, any, error) {
	if key, err := x509.ParsePKCS1PrivateKey(der); err == nil {
		return key, &key.PublicKey, nil
	}
	if key, err := x509.ParsePKCS8PrivateKey(der); err == nil {
		switch typed := key.(type) {
		case *rsa.PrivateKey:
			return typed, &typed.PublicKey, nil
		case *ecdsa.PrivateKey:
			return typed, &typed.PublicKey, nil
		case ed25519.PrivateKey:
			return typed, typed.Public(), nil
		default:
			return typed, nil, nil
		}
	}
	if key, err := x509.ParseECPrivateKey(der); err == nil {
		return key, &key.PublicKey, nil
	}
	return nil, nil, x509.IncorrectPasswordError
}

func parsePublicDER(der []byte) (any, error) {
	if key, err := x509.ParsePKIXPublicKey(der); err == nil {
		return key, nil
	}
	if key, err := x509.ParsePKCS1PublicKey(der); err == nil {
		return key, nil
	}
	if cert, err := x509.ParseCertificate(der); err == nil {
		return cert.PublicKey, nil
	}
	return nil, x509.IncorrectPasswordError
}

func sslKeyFromBytes(raw []byte, isPublic bool) *SslKey {
	if len(raw) == 0 {
		Throw(StringFromLiteral("Invalid key data"))
		return &SslKey{}
	}
	if block, _ := pem.Decode(raw); block != nil {
		if isPublic {
			key, err := parsePublicDER(block.Bytes)
			if err != nil {
				if priv, pub, perr := parsePrivateDER(block.Bytes); perr == nil {
					return &SslKey{privateKey: priv, publicKey: pub}
				}
				Throw(err)
				return &SslKey{}
			}
			return &SslKey{publicKey: key}
		}
		priv, pub, err := parsePrivateDER(block.Bytes)
		if err != nil {
			Throw(err)
			return &SslKey{}
		}
		return &SslKey{privateKey: priv, publicKey: pub}
	}
	if isPublic {
		key, err := parsePublicDER(raw)
		if err != nil {
			Throw(err)
			return &SslKey{}
		}
		return &SslKey{publicKey: key}
	}
	priv, pub, err := parsePrivateDER(raw)
	if err != nil {
		Throw(err)
		return &SslKey{}
	}
	return &SslKey{privateKey: priv, publicKey: pub}
}

func SslKeyLoadFile(file *string, isPublic bool, _pass *string) any {
	raw, err := os.ReadFile(*StdString(file))
	if err != nil {
		Throw(err)
		return nil
	}
	return sslKeyFromBytes(raw, isPublic)
}

func SslKeyReadPEM(data *string, isPublic bool, _pass *string) any {
	return sslKeyFromBytes([]byte(*StdString(data)), isPublic)
}

func SslKeyReadDER(data []byte, isPublic bool) any {
	return sslKeyFromBytes(data, isPublic)
}

func SslDigestSign(data []byte, handle any, alg *string) []byte {
	key, _ := handle.(*SslKey)
	if key == nil || key.privateKey == nil {
		Throw(StringFromLiteral("sys.ssl.Key private key is not available"))
		return []byte{}
	}
	name := sslDigestName(alg)
	hash := sslDigestHash(name)
	digest := sslDigestBytes(data, name)
	switch typed := key.privateKey.(type) {
	case *rsa.PrivateKey:
		sig, err := rsa.SignPKCS1v15(rand.Reader, typed, hash, digest)
		if err != nil {
			Throw(err)
			return []byte{}
		}
		return sig
	case *ecdsa.PrivateKey:
		sig, err := ecdsa.SignASN1(rand.Reader, typed, digest)
		if err != nil {
			Throw(err)
			return []byte{}
		}
		return sig
	case ed25519.PrivateKey:
		return ed25519.Sign(typed, data)
	default:
		Throw(StringFromLiteral("Unsupported sys.ssl.Key signing type"))
		return []byte{}
	}
}

func SslDigestVerify(data []byte, signature []byte, handle any, alg *string) bool {
	key, _ := handle.(*SslKey)
	if key == nil || key.publicKey == nil {
		return false
	}
	name := sslDigestName(alg)
	hash := sslDigestHash(name)
	digest := sslDigestBytes(data, name)
	switch typed := key.publicKey.(type) {
	case *rsa.PublicKey:
		return rsa.VerifyPKCS1v15(typed, hash, digest, signature) == nil
	case *ecdsa.PublicKey:
		return ecdsa.VerifyASN1(typed, digest, signature)
	case ed25519.PublicKey:
		return ed25519.Verify(typed, data, signature)
	default:
		return false
	}
}

func parseCertificates(raw []byte) []*x509.Certificate {
	certs := make([]*x509.Certificate, 0)
	rest := raw
	for {
		block, remaining := pem.Decode(rest)
		if block == nil {
			break
		}
		rest = remaining
		if block.Type != "CERTIFICATE" {
			continue
		}
		if cert, err := x509.ParseCertificate(block.Bytes); err == nil {
			certs = append(certs, cert)
		}
	}
	if len(certs) > 0 {
		return certs
	}
	parsed, err := x509.ParseCertificates(raw)
	if err == nil {
		return parsed
	}
	Throw(err)
	return []*x509.Certificate{}
}

func newSslCertificate(certs []*x509.Certificate, pool *x509.CertPool) *SslCertificate {
	if pool == nil {
		pool = x509.NewCertPool()
		for _, cert := range certs {
			if cert != nil {
				pool.AddCert(cert)
			}
		}
	}
	return &SslCertificate{certs: certs, pool: pool, index: 0}
}

func SslCertLoadFile(file *string) any {
	raw, err := os.ReadFile(*StdString(file))
	if err != nil {
		Throw(err)
		return nil
	}
	return newSslCertificate(parseCertificates(raw), nil)
}

func SslCertLoadPath(path *string) any {
	pool := x509.NewCertPool()
	certs := make([]*x509.Certificate, 0)
	entries, err := os.ReadDir(*StdString(path))
	if err != nil {
		Throw(err)
		return nil
	}
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		raw, readErr := os.ReadFile(filepath.Join(*StdString(path), entry.Name()))
		if readErr != nil {
			continue
		}
		for _, cert := range parseCertificates(raw) {
			if cert != nil {
				certs = append(certs, cert)
				pool.AddCert(cert)
			}
		}
	}
	return newSslCertificate(certs, pool)
}

func SslCertFromString(str *string) any {
	return newSslCertificate(parseCertificates([]byte(*StdString(str))), nil)
}

func SslCertLoadDefaults() any {
	pool, err := x509.SystemCertPool()
	if err != nil || pool == nil {
		pool = x509.NewCertPool()
	}
	return newSslCertificate([]*x509.Certificate{}, pool)
}

func sslLeafCertificate(handle any) *x509.Certificate {
	cert, _ := handle.(*SslCertificate)
	if cert == nil || cert.index < 0 || cert.index >= len(cert.certs) {
		return nil
	}
	return cert.certs[cert.index]
}

func sslNameField(name pkix.Name, field string) *string {
	switch strings.ToUpper(field) {
	case "CN":
		if name.CommonName == "" {
			return nil
		}
		return StringFromLiteral(name.CommonName)
	case "O":
		if len(name.Organization) == 0 {
			return nil
		}
		return StringFromLiteral(name.Organization[0])
	case "OU":
		if len(name.OrganizationalUnit) == 0 {
			return nil
		}
		return StringFromLiteral(name.OrganizationalUnit[0])
	case "C":
		if len(name.Country) == 0 {
			return nil
		}
		return StringFromLiteral(name.Country[0])
	case "L":
		if len(name.Locality) == 0 {
			return nil
		}
		return StringFromLiteral(name.Locality[0])
	case "ST":
		if len(name.Province) == 0 {
			return nil
		}
		return StringFromLiteral(name.Province[0])
	default:
		return nil
	}
}

func SslCertCommonName(handle any) *string {
	cert := sslLeafCertificate(handle)
	if cert == nil || cert.Subject.CommonName == "" {
		return nil
	}
	return StringFromLiteral(cert.Subject.CommonName)
}

func SslCertAltNames(handle any) []*string {
	cert := sslLeafCertificate(handle)
	if cert == nil {
		return []*string{}
	}
	out := make([]*string, 0, len(cert.DNSNames)+len(cert.IPAddresses)+len(cert.EmailAddresses))
	for _, name := range cert.DNSNames {
		out = append(out, StringFromLiteral(name))
	}
	for _, ip := range cert.IPAddresses {
		out = append(out, StringFromLiteral(ip.String()))
	}
	for _, mail := range cert.EmailAddresses {
		out = append(out, StringFromLiteral(mail))
	}
	return out
}

func SslCertSubject(handle any, field *string) *string {
	cert := sslLeafCertificate(handle)
	if cert == nil {
		return nil
	}
	return sslNameField(cert.Subject, *StdString(field))
}

func SslCertIssuer(handle any, field *string) *string {
	cert := sslLeafCertificate(handle)
	if cert == nil {
		return nil
	}
	return sslNameField(cert.Issuer, *StdString(field))
}

func SslCertNotBeforeMs(handle any) float64 {
	cert := sslLeafCertificate(handle)
	if cert == nil {
		return 0
	}
	return float64(cert.NotBefore.UnixMilli())
}

func SslCertNotAfterMs(handle any) float64 {
	cert := sslLeafCertificate(handle)
	if cert == nil {
		return 0
	}
	return float64(cert.NotAfter.UnixMilli())
}

func SslCertNext(handle any) any {
	cert, _ := handle.(*SslCertificate)
	if cert == nil || cert.index+1 >= len(cert.certs) {
		return nil
	}
	return &SslCertificate{certs: cert.certs, pool: cert.pool, index: cert.index + 1}
}

func SslCertAddPEM(handle any, pemText *string) {
	SslCertAddDER(handle, []byte(*StdString(pemText)))
}

func SslCertAddDER(handle any, der []byte) {
	cert, _ := handle.(*SslCertificate)
	if cert == nil {
		return
	}
	parsed := parseCertificates(der)
	if cert.pool == nil {
		cert.pool = x509.NewCertPool()
	}
	for _, entry := range parsed {
		if entry != nil {
			cert.certs = append(cert.certs, entry)
			cert.pool.AddCert(entry)
		}
	}
}
