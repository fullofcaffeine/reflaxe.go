package main

import "snapshot/hxrt"

var DEFAULT_CERT_PEM *string = hxrt.StringFromLiteral("-----BEGIN CERTIFICATE-----\nMIIDETCCAfmgAwIBAgIUVftG+IyWIEqIQE+K9ztyXCDYIcswDQYJKoZIhvcNAQEL\nBQAwGDEWMBQGA1UEAwwNZGVmYXVsdC5sb2NhbDAeFw0yNjAzMDcwNjA1NTBaFw0y\nNzAzMDcwNjA1NTBaMBgxFjAUBgNVBAMMDWRlZmF1bHQubG9jYWwwggEiMA0GCSqG\nSIb3DQEBAQUAA4IBDwAwggEKAoIBAQCktatoj06k/dqldeSzjPUnPxeMC/WpprMz\n7tHisw82tHc0Xk18wW/m0Dm+W11kLlq+k5fNuVoQTcqaQDkLF0Zy+Q2K+GWLRCux\n7Ms0ixw6oSIUFnaG8+SByKuEfaW232ZKCWFsSxq0PdVKT3jjcc3ivv6j/kOsbE7j\njELMP9w1askakA/I8CWM0AZyYVZ5ajwYcBBQm1UOzWLeoT+UU7O9VKsEIQHCGt5P\np7U/PDqh5z7KJ+XWIG/jpjZ8IYEo8fTxik+16jN34Ubpnx+I/wUmLK4Aqh2VuuZz\nmmcUFD7JLC+r9ymOUa7DS0bWyUoBDzZyvAO1577p7SXi9/IqJ3NpAgMBAAGjUzBR\nMB0GA1UdDgQWBBR7eOzK13X2mO5Kpc9MZOFGJ+MdlDAfBgNVHSMEGDAWgBR7eOzK\n13X2mO5Kpc9MZOFGJ+MdlDAPBgNVHRMBAf8EBTADAQH/MA0GCSqGSIb3DQEBCwUA\nA4IBAQAxilSFs9ZEEI5RFGVpWUTCH8Iuewc4K0JH2LbGqExgLu5MQJF5xsGoEBRe\nmDB3duaMDwnoDY19hdoClIz/Z5IO0wPEcny3hTb582W8+cRDiCQx0Qz5g2NpsfGH\nwZaGyVzSS2xbt0F6TPRQarXtzV97J067j7bRuMbD4fFYb6iqF1GbaaQsP89sA9b4\nx8yT0TGREdx2Dw29lokd4H16b3O9aYX274qZk9qpgE3oYpCeNsjHCSdTWRDw2+pg\ncKP2FsxcbB4qChjPXhhE5I3zDMTI1/4r4wGbnVIud6TmibElpAU6hwbpo1onyxKw\nVCpI190WWaS+Cz+9vGVNBSsRYjgN\n-----END CERTIFICATE-----\n")

var DEFAULT_KEY_PEM *string = hxrt.StringFromLiteral("-----BEGIN PRIVATE KEY-----\nMIIEvwIBADANBgkqhkiG9w0BAQEFAASCBKkwggSlAgEAAoIBAQCktatoj06k/dql\ndeSzjPUnPxeMC/WpprMz7tHisw82tHc0Xk18wW/m0Dm+W11kLlq+k5fNuVoQTcqa\nQDkLF0Zy+Q2K+GWLRCux7Ms0ixw6oSIUFnaG8+SByKuEfaW232ZKCWFsSxq0PdVK\nT3jjcc3ivv6j/kOsbE7jjELMP9w1askakA/I8CWM0AZyYVZ5ajwYcBBQm1UOzWLe\noT+UU7O9VKsEIQHCGt5Pp7U/PDqh5z7KJ+XWIG/jpjZ8IYEo8fTxik+16jN34Ubp\nnx+I/wUmLK4Aqh2VuuZzmmcUFD7JLC+r9ymOUa7DS0bWyUoBDzZyvAO1577p7SXi\n9/IqJ3NpAgMBAAECggEAECDG2jvwGjlOO9XtsUQr7C4iHuE76qMLWJo5vI5GfteQ\ndZUHVuA1Fh+NC+3z2N+uHIukuWz9G+wWGuEPhN3AVPk8oX89qDOiaK90XV6Cwt3s\n0ZU6gW/nz+qHmvdXru60nCrLephnEr+cP0TFZFYMMDf+BK5cz4kid2cQUmItbKBy\nE9k1vHvyyLONK18y72IXsMEA572V+AhtvWVar5qUJiY2+dWtxmWh9dzUvWCUJ/Gj\niqyrJdNxHJfXdteT8QpJEndsX19MWNAMZY+e10XOVzs2Rd4RdnlLwUE37HbPz73V\nC1z2h6tshSOYi9ZwlNdYEjfWbuiQimwgljwc9/NagQKBgQDg4JkcchEOmntEXPz0\n/HTEmAd7Sl75V0ozj7xPTYcaSvdoEtLEbUJbGoVfLKJM17jLuEjpisS1IHa29kkG\n2PACbbIRzOLlCA2Spkzm9R2Rni9qbBCwYbnYbdhV+hshcnr6LwUt8KBAFHBCk2lW\nMfiCsBE1KzBeye1JFxIZ9+dyKQKBgQC7gVSvPDIqPRYP8l8sLbONMHTNhbZpMh23\npdD1i0YpYTIuvrxZMXwY6gt71SJcHIwyh/6HEg8mAIMZutdRzpZc4NWQw98yG9Zo\nkpY7S2uLZ8YnARps+NprdWiARkY3PJjIiK7tieNo+dQtfBeuL3Ox7s/XR0nDUXMP\nxkuAnByfQQKBgQCks9twehsEFyExcOnUhRMA6liQdGgbN1OhcCT78EyDdWS/VQoJ\n0/xFvabxjj9RCK7QhqjgZEKuZpiMaNYTrdAb9zv0zZthJATM5ABvKBgAD1urFnsi\ntHDpk4pfbk9wr+hiVQ32F8dHJ7EREeaUuwTIsyvnRTqoMj0Yy0z2uBtMAQKBgQC2\nEnG+7z7vEP4ZYgrUhVQyp3jkERD9uTJuH892f1UT3VOzXHbcTVbpgmrARkflFbt1\nXeTkF78p8ZlcJLfssiQD8DaxKeHTcICUbrL+xM+bQJuDSGj2o/bEHe/pj1OjU24w\nW7kw45I1X1KPEE6WT3GSuAiOTKTtymtmR/EM44pPgQKBgQC52gWDbpJb1UREu9Sb\naIy1zgqGhSxg1DLChu4vz3t5ZBhWoiw6wG7vQDrtrn9Jl61mVtog0rSTSiIj726W\nLDxSrTdYBhcgkbJf/fw1N6Uk+FfJcVS/zuoJwogkiP6tmircL0qA9sYR4RZo2ubE\nAg6A+vHIw54BclYKSczsXnCe+A==\n-----END PRIVATE KEY-----\n")

var SNI_CERT_PEM *string = hxrt.StringFromLiteral("-----BEGIN CERTIFICATE-----\nMIIDHzCCAgegAwIBAgIUMCFv2IwHqj9Wy6ycvAibkSutMKQwDQYJKoZIhvcNAQEL\nBQAwFDESMBAGA1UEAwwJc25pLmxvY2FsMB4XDTI2MDYxMDA1MTUwMloXDTI3MDYx\nMDA1MTUwMlowFDESMBAGA1UEAwwJc25pLmxvY2FsMIIBIjANBgkqhkiG9w0BAQEF\nAAOCAQ8AMIIBCgKCAQEAsQVlch80ZQ4hg3Hb3XMPgk+XPsJuoIdE0rn78PGM/O6W\n1iQocq9hYRosp/aT6itBnpvxh+tCt9D7deoEDBzwU6CkJu7N6CGY0yWWrFkits4f\n8KiTbkWZA3qKfreUOiDkiU3DgM+b9GUyBkgUXxBxW1ql+mFAX6l3eRDK5q9dxxme\nmouBL4yQxMa1bgunsYoZQhV/0RoH0gUuDUTPr12+YEA/XyCJC82p3v50mIcURjnX\nHl26xMRv5O73mJDUqshn3Z0SMmXhHfk6JFsacc2sB7f8FCOAAFsjjhSKHEQ8OTTP\nzcXAw6fcqxF1WNMrsIfZJ4Dzjt2kOYO9KxiLR3+0dwIDAQABo2kwZzAdBgNVHQ4E\nFgQUnbhLMtV250tds1NaZBDZyIvDXfwwHwYDVR0jBBgwFoAUnbhLMtV250tds1Na\nZBDZyIvDXfwwDwYDVR0TAQH/BAUwAwEB/zAUBgNVHREEDTALgglzbmkubG9jYWww\nDQYJKoZIhvcNAQELBQADggEBADHJ3F34+3iZrJNgxuFbr5IOsm4GDCZxd8g0BmWD\n+JxabwGYMk+xGOL/dn7jPwhodHcFbJXrYzT2giGMG9EdLLExqwnpQb7Z+uqga4Ej\nfyQVZZ61Np8wSGJQ7C03Yu6A0jYljW8KWM3UeiX6ee/d8/JXK8Nts8qsKXe6SFJO\nF1vblGLBZyb/DZ4GKD5cp3RZdCQxwlvLmbYKrKDwySJYwWQEx8bD5UUMYPr0Y593\nn6ZbEIWl+17PkP0kQRloeGUOxogZaI9jM/fGalKf52SbipDQV4P8NXqsdW0Im/LA\nCW6HU+6ozncjLj3SyLuwLvd54+YEPhv/cuFgP6/WWqyCx0A=\n-----END CERTIFICATE-----\n")

var SNI_KEY_PEM *string = hxrt.StringFromLiteral("-----BEGIN PRIVATE KEY-----\nMIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQCxBWVyHzRlDiGD\ncdvdcw+CT5c+wm6gh0TSufvw8Yz87pbWJChyr2FhGiyn9pPqK0Gem/GH60K30Pt1\n6gQMHPBToKQm7s3oIZjTJZasWSK2zh/wqJNuRZkDeop+t5Q6IOSJTcOAz5v0ZTIG\nSBRfEHFbWqX6YUBfqXd5EMrmr13HGZ6ai4EvjJDExrVuC6exihlCFX/RGgfSBS4N\nRM+vXb5gQD9fIIkLzane/nSYhxRGOdceXbrExG/k7veYkNSqyGfdnRIyZeEd+Tok\nWxpxzawHt/wUI4AAWyOOFIocRDw5NM/NxcDDp9yrEXVY0yuwh9kngPOO3aQ5g70r\nGItHf7R3AgMBAAECggEANjMyxmhrgG19MWPRL9Kk8v8vjdW2TYxdNDAhxboPsvnS\nUSqs/8BXDoYXGi5TR6WK5+dTYoxT1zgzZf0K1DKgGtrap9kCTorK4gtmQMrh6Brg\niKz0xxSkLv58HSRUTB/6GVgn/e6TD5dUY7v6EMlWC+SLYUgZj7Cxle3gUhVrnyPk\noaEOMDU8w9iiPoCK0E60DReKAmIcXigGqMNlkFjADjfOtr5GDpercL+mvwDmfLmB\ndzYcXuHM3fm1pgoDXN5zrpO7mFSp0VG/O697nZzY18A25AMMNiBUjuk7/u4KJEcA\nMmOxVAU6yjHAKO+YRS98pKsxkvZKNpj69lXtBv6agQKBgQDVnTDnIudJwbaAvfBf\n7ZZ3l46Ojm2s3qDR0Sp+W0tTc5cVhJiJaShzpnYKPbZP6R8EnDrdtldw8sYRRB6R\nPfM3/DX87z793+pKB5WDJNzevO81ecpMhiYgUxlKzzkz7nfDZC8tBQ3YLLhUX5JL\nv7REeUskkmW+xq9e9+ecK/RpVwKBgQDUJWrvJ2tEf2JJ+qdYZBHJXw5+ajIZeUs+\nCh3cNtOrUGDW/y9SKvrRHzUP0PvRle21O4oUimEb1SqdNG2nFB3f4BVUB8X6DcMz\nM16kyNBdk4q2k3xITw2OdErlYZU2WAsLUrOxurDK9ng6YO1fJwz3zrGiNDjbdebY\n1m4NbyN54QKBgQCE0akvfh9LV/wPDoqgSszs7TpBX0PIYeCitShzynYKnGuLgJeL\nkOwLBKyOb5KlGzEjH7TmWFMEMp9+6tkKu/c3j0VOUL/dANXfU9nd4hTHFbiyiliD\nvkGEhcbLIg/SP2sN/YPrvSG/kQbHx2jiWn9OuBBF3BURSt6N8Rx8mUPuHwKBgQCG\nAqj9L8Jz/5/gKaVCkdwmf5SRSJYjP1rHcu6P6FZntpul1IdY+Wt9ZKBJQHOCXppN\nTLIZ7ZwQT+Teb3sA+xUwEcaHUW2/Wqg/FKkpoOz237fVQ29T4hQnM9EH+0+dh5pa\nacC3eb4qR+2EuyvXWry3YWsWkrSD9YOA4FuewuD/IQKBgDUet4ixooKltLLhW6RT\nB+8+u2hKmlctueypoiB9+yiCt/oEiZEHYX9IwlGaax6bpDg6vO77Ay7JwGaggM+K\nhhIsqAuszWLO7bdjdS4JZ1YshP5YusdPs3m9OGeTjnO/+goL+tECrZmzupwg4BTA\nKOG6OBx2h7HDGnJcWJNLuBbJ\n-----END PRIVATE KEY-----\n")

func connectAndPrint(server *sys__ssl__Socket, port int, hostname *string) {
	acceptedSockets := New_sys__thread__Deque()
	sys__thread__Thread_create(func() {
		acceptedSockets.__hx_this.add(server.__hx_this.accept())
	})
	client := New_sys__ssl__Socket()
	client.verifyCert = false
	client.__hx_this.setHostname(hostname)
	client.__hx_this.connect(New_sys__net__Host(hxrt.StringFromLiteral("127.0.0.1")), port)
	client.__hx_this.handshake()
	peer := client.__hx_this.peerCertificate()
	var v any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hostname, hxrt.StringFromLiteral("=")), peer.__hx_this.get_commonName()))
	hxrt.Println(v)
	func(hx_value_1 any) *sys__net__Socket {
		if hx_value_1 == nil {
			var hx_zero_2 *sys__net__Socket
			return hx_zero_2
		}
		return hx_value_1.(*sys__net__Socket)
	}(acceptedSockets.__hx_this.pop(true)).__hx_this.close()
	client.__hx_this.close()
}

func main() {
	defer hxrt.ThreadWaitForAll()
	defaultCert := sys__ssl__Certificate_fromString(DEFAULT_CERT_PEM)
	defaultKey := sys__ssl__Key_readPEM(DEFAULT_KEY_PEM, false, nil)
	sniCert := sys__ssl__Certificate_fromString(SNI_CERT_PEM)
	sniKey := sys__ssl__Key_readPEM(SNI_KEY_PEM, false, nil)
	server := New_sys__ssl__Socket()
	server.__hx_this.setCertificate(defaultCert, defaultKey)
	server.__hx_this.addSNICertificate(func(name *string) bool {
		return hxrt.StringEqualStringPtr(name, hxrt.StringFromLiteral("sni.local"))
	}, sniCert, sniKey)
	server.__hx_this.bind(New_sys__net__Host(hxrt.StringFromLiteral("127.0.0.1")), 0)
	server.__hx_this.listen(2)
	bound := server.__hx_this.host()
	connectAndPrint(server, func(hx_obj_3 map[string]any) int {
		hx_field_4 := hx_obj_3["port"]
		if hx_field_4 == nil {
			var hx_zero_5 int
			return hx_zero_5
		}
		return hx_field_4.(int)
	}(bound), hxrt.StringFromLiteral("default.local"))
	connectAndPrint(server, func(hx_obj_6 map[string]any) int {
		hx_field_7 := hx_obj_6["port"]
		if hx_field_7 == nil {
			var hx_zero_8 int
			return hx_zero_8
		}
		return hx_field_7.(int)
	}(bound), hxrt.StringFromLiteral("sni.local"))
	server.__hx_this.close()
}

type Std struct {
}

type haxe__ds__Option struct {
	tag    int
	params []any
}

var haxe__ds__Option_None *haxe__ds__Option = &haxe__ds__Option{tag: 1, params: []any{}}

func haxe__ds__Option_Some(value any) *haxe__ds__Option {
	return &haxe__ds__Option{tag: 0, params: []any{value}}
}
