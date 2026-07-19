package main

import "snapshot/hxrt"

func main() {
	certPem := hxrt.StringFromLiteral("-----BEGIN CERTIFICATE-----\nMIIDETCCAfmgAwIBAgIUVftG+IyWIEqIQE+K9ztyXCDYIcswDQYJKoZIhvcNAQEL\nBQAwGDEWMBQGA1UEAwwNZGVmYXVsdC5sb2NhbDAeFw0yNjAzMDcwNjA1NTBaFw0y\nNzAzMDcwNjA1NTBaMBgxFjAUBgNVBAMMDWRlZmF1bHQubG9jYWwwggEiMA0GCSqG\nSIb3DQEBAQUAA4IBDwAwggEKAoIBAQCktatoj06k/dqldeSzjPUnPxeMC/WpprMz\n7tHisw82tHc0Xk18wW/m0Dm+W11kLlq+k5fNuVoQTcqaQDkLF0Zy+Q2K+GWLRCux\n7Ms0ixw6oSIUFnaG8+SByKuEfaW232ZKCWFsSxq0PdVKT3jjcc3ivv6j/kOsbE7j\njELMP9w1askakA/I8CWM0AZyYVZ5ajwYcBBQm1UOzWLeoT+UU7O9VKsEIQHCGt5P\np7U/PDqh5z7KJ+XWIG/jpjZ8IYEo8fTxik+16jN34Ubpnx+I/wUmLK4Aqh2VuuZz\nmmcUFD7JLC+r9ymOUa7DS0bWyUoBDzZyvAO1577p7SXi9/IqJ3NpAgMBAAGjUzBR\nMB0GA1UdDgQWBBR7eOzK13X2mO5Kpc9MZOFGJ+MdlDAfBgNVHSMEGDAWgBR7eOzK\n13X2mO5Kpc9MZOFGJ+MdlDAPBgNVHRMBAf8EBTADAQH/MA0GCSqGSIb3DQEBCwUA\nA4IBAQAxilSFs9ZEEI5RFGVpWUTCH8Iuewc4K0JH2LbGqExgLu5MQJF5xsGoEBRe\nmDB3duaMDwnoDY19hdoClIz/Z5IO0wPEcny3hTb582W8+cRDiCQx0Qz5g2NpsfGH\nwZaGyVzSS2xbt0F6TPRQarXtzV97J067j7bRuMbD4fFYb6iqF1GbaaQsP89sA9b4\nx8yT0TGREdx2Dw29lokd4H16b3O9aYX274qZk9qpgE3oYpCeNsjHCSdTWRDw2+pg\ncKP2FsxcbB4qChjPXhhE5I3zDMTI1/4r4wGbnVIud6TmibElpAU6hwbpo1onyxKw\nVCpI190WWaS+Cz+9vGVNBSsRYjgN\n-----END CERTIFICATE-----\n")
	keyPem := hxrt.StringFromLiteral("-----BEGIN PRIVATE KEY-----\nMIIEvwIBADANBgkqhkiG9w0BAQEFAASCBKkwggSlAgEAAoIBAQCktatoj06k/dql\ndeSzjPUnPxeMC/WpprMz7tHisw82tHc0Xk18wW/m0Dm+W11kLlq+k5fNuVoQTcqa\nQDkLF0Zy+Q2K+GWLRCux7Ms0ixw6oSIUFnaG8+SByKuEfaW232ZKCWFsSxq0PdVK\nT3jjcc3ivv6j/kOsbE7jjELMP9w1askakA/I8CWM0AZyYVZ5ajwYcBBQm1UOzWLe\noT+UU7O9VKsEIQHCGt5Pp7U/PDqh5z7KJ+XWIG/jpjZ8IYEo8fTxik+16jN34Ubp\nnx+I/wUmLK4Aqh2VuuZzmmcUFD7JLC+r9ymOUa7DS0bWyUoBDzZyvAO1577p7SXi\n9/IqJ3NpAgMBAAECggEAECDG2jvwGjlOO9XtsUQr7C4iHuE76qMLWJo5vI5GfteQ\ndZUHVuA1Fh+NC+3z2N+uHIukuWz9G+wWGuEPhN3AVPk8oX89qDOiaK90XV6Cwt3s\n0ZU6gW/nz+qHmvdXru60nCrLephnEr+cP0TFZFYMMDf+BK5cz4kid2cQUmItbKBy\nE9k1vHvyyLONK18y72IXsMEA572V+AhtvWVar5qUJiY2+dWtxmWh9dzUvWCUJ/Gj\niqyrJdNxHJfXdteT8QpJEndsX19MWNAMZY+e10XOVzs2Rd4RdnlLwUE37HbPz73V\nC1z2h6tshSOYi9ZwlNdYEjfWbuiQimwgljwc9/NagQKBgQDg4JkcchEOmntEXPz0\n/HTEmAd7Sl75V0ozj7xPTYcaSvdoEtLEbUJbGoVfLKJM17jLuEjpisS1IHa29kkG\n2PACbbIRzOLlCA2Spkzm9R2Rni9qbBCwYbnYbdhV+hshcnr6LwUt8KBAFHBCk2lW\nMfiCsBE1KzBeye1JFxIZ9+dyKQKBgQC7gVSvPDIqPRYP8l8sLbONMHTNhbZpMh23\npdD1i0YpYTIuvrxZMXwY6gt71SJcHIwyh/6HEg8mAIMZutdRzpZc4NWQw98yG9Zo\nkpY7S2uLZ8YnARps+NprdWiARkY3PJjIiK7tieNo+dQtfBeuL3Ox7s/XR0nDUXMP\nxkuAnByfQQKBgQCks9twehsEFyExcOnUhRMA6liQdGgbN1OhcCT78EyDdWS/VQoJ\n0/xFvabxjj9RCK7QhqjgZEKuZpiMaNYTrdAb9zv0zZthJATM5ABvKBgAD1urFnsi\ntHDpk4pfbk9wr+hiVQ32F8dHJ7EREeaUuwTIsyvnRTqoMj0Yy0z2uBtMAQKBgQC2\nEnG+7z7vEP4ZYgrUhVQyp3jkERD9uTJuH892f1UT3VOzXHbcTVbpgmrARkflFbt1\nXeTkF78p8ZlcJLfssiQD8DaxKeHTcICUbrL+xM+bQJuDSGj2o/bEHe/pj1OjU24w\nW7kw45I1X1KPEE6WT3GSuAiOTKTtymtmR/EM44pPgQKBgQC52gWDbpJb1UREu9Sb\naIy1zgqGhSxg1DLChu4vz3t5ZBhWoiw6wG7vQDrtrn9Jl61mVtog0rSTSiIj726W\nLDxSrTdYBhcgkbJf/fw1N6Uk+FfJcVS/zuoJwogkiP6tmircL0qA9sYR4RZo2ubE\nAg6A+vHIw54BclYKSczsXnCe+A==\n-----END PRIVATE KEY-----\n")
	cert := sys__ssl__Certificate_fromString(certPem)
	key := sys__ssl__Key_readPEM(keyPem, false, nil)
	payload := haxe__io__Bytes_ofString(hxrt.StringFromLiteral("digest-payload"), nil)
	sig := sys__ssl__Digest_sign(payload, key, any(hxrt.StringFromLiteral("SHA256")))
	var v any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("cert.cn="), cert.__hx_this.get_commonName()))
	hxrt.Println(v)
	var v_1 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("cert.subject.cn="), cert.__hx_this.subject(hxrt.StringFromLiteral("CN"))))
	hxrt.Println(v_1)
	var v_2 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("cert.issuer.cn="), cert.__hx_this.issuer(hxrt.StringFromLiteral("CN"))))
	hxrt.Println(v_2)
	var v_3 any = any(hxrt.StringConcatAny(hxrt.StringFromLiteral("cert.altNames="), cert.__hx_this.get_altNames().Len()))
	hxrt.Println(v_3)
	var v_4 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("cert.next="), hxrt.StdString((cert.__hx_this.next() == nil))))
	hxrt.Println(v_4)
	var v_5 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("digest.sha256="), sys__ssl__Digest_make(payload, any(hxrt.StringFromLiteral("SHA256"))).__hx_this.toHex()))
	hxrt.Println(v_5)
	var v_6 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("digest.sha512="), sys__ssl__Digest_make(payload, any(hxrt.StringFromLiteral("SHA512"))).__hx_this.toHex()))
	hxrt.Println(v_6)
	var v_7 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("digest.verify="), hxrt.StdString(sys__ssl__Digest_verify(payload, sig, key, any(hxrt.StringFromLiteral("SHA256"))))))
	hxrt.Println(v_7)
}
