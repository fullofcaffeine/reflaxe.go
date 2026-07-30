package main

import "snapshot/hxrt"

var CERT_PEM *string = hxrt.StringFromLiteral("-----BEGIN CERTIFICATE-----\nMIIDETCCAfmgAwIBAgIUVftG+IyWIEqIQE+K9ztyXCDYIcswDQYJKoZIhvcNAQEL\nBQAwGDEWMBQGA1UEAwwNZGVmYXVsdC5sb2NhbDAeFw0yNjAzMDcwNjA1NTBaFw0y\nNzAzMDcwNjA1NTBaMBgxFjAUBgNVBAMMDWRlZmF1bHQubG9jYWwwggEiMA0GCSqG\nSIb3DQEBAQUAA4IBDwAwggEKAoIBAQCktatoj06k/dqldeSzjPUnPxeMC/WpprMz\n7tHisw82tHc0Xk18wW/m0Dm+W11kLlq+k5fNuVoQTcqaQDkLF0Zy+Q2K+GWLRCux\n7Ms0ixw6oSIUFnaG8+SByKuEfaW232ZKCWFsSxq0PdVKT3jjcc3ivv6j/kOsbE7j\njELMP9w1askakA/I8CWM0AZyYVZ5ajwYcBBQm1UOzWLeoT+UU7O9VKsEIQHCGt5P\np7U/PDqh5z7KJ+XWIG/jpjZ8IYEo8fTxik+16jN34Ubpnx+I/wUmLK4Aqh2VuuZz\nmmcUFD7JLC+r9ymOUa7DS0bWyUoBDzZyvAO1577p7SXi9/IqJ3NpAgMBAAGjUzBR\nMB0GA1UdDgQWBBR7eOzK13X2mO5Kpc9MZOFGJ+MdlDAfBgNVHSMEGDAWgBR7eOzK\n13X2mO5Kpc9MZOFGJ+MdlDAPBgNVHRMBAf8EBTADAQH/MA0GCSqGSIb3DQEBCwUA\nA4IBAQAxilSFs9ZEEI5RFGVpWUTCH8Iuewc4K0JH2LbGqExgLu5MQJF5xsGoEBRe\nmDB3duaMDwnoDY19hdoClIz/Z5IO0wPEcny3hTb582W8+cRDiCQx0Qz5g2NpsfGH\nwZaGyVzSS2xbt0F6TPRQarXtzV97J067j7bRuMbD4fFYb6iqF1GbaaQsP89sA9b4\nx8yT0TGREdx2Dw29lokd4H16b3O9aYX274qZk9qpgE3oYpCeNsjHCSdTWRDw2+pg\ncKP2FsxcbB4qChjPXhhE5I3zDMTI1/4r4wGbnVIud6TmibElpAU6hwbpo1onyxKw\nVCpI190WWaS+Cz+9vGVNBSsRYjgN\n-----END CERTIFICATE-----\n")

var KEY_PEM *string = hxrt.StringFromLiteral("-----BEGIN PRIVATE KEY-----\nMIIEvwIBADANBgkqhkiG9w0BAQEFAASCBKkwggSlAgEAAoIBAQCktatoj06k/dql\ndeSzjPUnPxeMC/WpprMz7tHisw82tHc0Xk18wW/m0Dm+W11kLlq+k5fNuVoQTcqa\nQDkLF0Zy+Q2K+GWLRCux7Ms0ixw6oSIUFnaG8+SByKuEfaW232ZKCWFsSxq0PdVK\nT3jjcc3ivv6j/kOsbE7jjELMP9w1askakA/I8CWM0AZyYVZ5ajwYcBBQm1UOzWLe\noT+UU7O9VKsEIQHCGt5Pp7U/PDqh5z7KJ+XWIG/jpjZ8IYEo8fTxik+16jN34Ubp\nnx+I/wUmLK4Aqh2VuuZzmmcUFD7JLC+r9ymOUa7DS0bWyUoBDzZyvAO1577p7SXi\n9/IqJ3NpAgMBAAECggEAECDG2jvwGjlOO9XtsUQr7C4iHuE76qMLWJo5vI5GfteQ\ndZUHVuA1Fh+NC+3z2N+uHIukuWz9G+wWGuEPhN3AVPk8oX89qDOiaK90XV6Cwt3s\n0ZU6gW/nz+qHmvdXru60nCrLephnEr+cP0TFZFYMMDf+BK5cz4kid2cQUmItbKBy\nE9k1vHvyyLONK18y72IXsMEA572V+AhtvWVar5qUJiY2+dWtxmWh9dzUvWCUJ/Gj\niqyrJdNxHJfXdteT8QpJEndsX19MWNAMZY+e10XOVzs2Rd4RdnlLwUE37HbPz73V\nC1z2h6tshSOYi9ZwlNdYEjfWbuiQimwgljwc9/NagQKBgQDg4JkcchEOmntEXPz0\n/HTEmAd7Sl75V0ozj7xPTYcaSvdoEtLEbUJbGoVfLKJM17jLuEjpisS1IHa29kkG\n2PACbbIRzOLlCA2Spkzm9R2Rni9qbBCwYbnYbdhV+hshcnr6LwUt8KBAFHBCk2lW\nMfiCsBE1KzBeye1JFxIZ9+dyKQKBgQC7gVSvPDIqPRYP8l8sLbONMHTNhbZpMh23\npdD1i0YpYTIuvrxZMXwY6gt71SJcHIwyh/6HEg8mAIMZutdRzpZc4NWQw98yG9Zo\nkpY7S2uLZ8YnARps+NprdWiARkY3PJjIiK7tieNo+dQtfBeuL3Ox7s/XR0nDUXMP\nxkuAnByfQQKBgQCks9twehsEFyExcOnUhRMA6liQdGgbN1OhcCT78EyDdWS/VQoJ\n0/xFvabxjj9RCK7QhqjgZEKuZpiMaNYTrdAb9zv0zZthJATM5ABvKBgAD1urFnsi\ntHDpk4pfbk9wr+hiVQ32F8dHJ7EREeaUuwTIsyvnRTqoMj0Yy0z2uBtMAQKBgQC2\nEnG+7z7vEP4ZYgrUhVQyp3jkERD9uTJuH892f1UT3VOzXHbcTVbpgmrARkflFbt1\nXeTkF78p8ZlcJLfssiQD8DaxKeHTcICUbrL+xM+bQJuDSGj2o/bEHe/pj1OjU24w\nW7kw45I1X1KPEE6WT3GSuAiOTKTtymtmR/EM44pPgQKBgQC52gWDbpJb1UREu9Sb\naIy1zgqGhSxg1DLChu4vz3t5ZBhWoiw6wG7vQDrtrn9Jl61mVtog0rSTSiIj726W\nLDxSrTdYBhcgkbJf/fw1N6Uk+FfJcVS/zuoJwogkiP6tmircL0qA9sYR4RZo2ubE\nAg6A+vHIw54BclYKSczsXnCe+A==\n-----END PRIVATE KEY-----\n")

func main() {
	defer hxrt.ThreadWaitForAll()
	cert := sys__ssl__Certificate_fromString(CERT_PEM)
	key := sys__ssl__Key_readPEM(KEY_PEM, false, nil)
	server := New_sys__ssl__Socket()
	server.__hx_this.setCertificate(cert, key)
	server.__hx_this.bind(New_sys__net__Host(hxrt.StringFromLiteral("127.0.0.1")), 0)
	server.__hx_this.listen(1)
	bound := server.__hx_this.host()
	acceptedSockets := New_sys__thread__Deque()
	sys__thread__Thread_create(func() {
		acceptedSockets.__hx_this.add(server.__hx_this.accept())
	})
	client := New_sys__ssl__Socket()
	client.verifyCert = false
	client.__hx_this.setCA(cert)
	client.__hx_this.connect(New_sys__net__Host(hxrt.StringFromLiteral("localhost")), func(hx_obj_1 map[string]any) int {
		hx_field_2 := hx_obj_1["port"]
		if hx_field_2 == nil {
			var hx_zero_3 int
			return hx_zero_3
		}
		return hx_field_2.(int)
	}(bound))
	client.__hx_this.handshake()
	accepted := func(hx_value_4 any) *sys__net__Socket {
		if hx_value_4 == nil {
			var hx_zero_5 *sys__net__Socket
			return hx_zero_5
		}
		return hx_value_4.(*sys__net__Socket)
	}(acceptedSockets.__hx_this.pop(true))
	var v any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("server.accept.ssl="), hxrt.StdString(func(hx_value *sys__net__Socket) bool {
		if hx_value == nil {
			return false
		}
		_, ok := hx_value.__hx_this.(*sys__ssl__Socket)
		return ok
	}(accepted))))
	hxrt.Println(v)
	client.output.__hx_this.writeString(hxrt.StringFromLiteral("ping\n"), nil)
	client.output.__hx_this.flush()
	var v_1 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("server.read="), accepted.input.__hx_this.readLine()))
	hxrt.Println(v_1)
	accepted.output.__hx_this.writeString(hxrt.StringFromLiteral("pong\n"), nil)
	accepted.output.__hx_this.flush()
	var v_2 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("client.read="), client.input.__hx_this.readLine()))
	hxrt.Println(v_2)
	readShutdownError := hxrt.StringFromLiteral("missing")
	hxrt.TryCatch(func() {
		client.__hx_this.shutdown(true, false)
	}, func(hx_caught_6 any) {
		error := hxrt.ExceptionCaught(hx_caught_6)
		readShutdownError = hxrt.ExceptionMessage(error)
	})
	hxrt.Println(any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("tls.read.shutdown="), readShutdownError)))
	client.__hx_this.setFastSend(true)
	client.__hx_this.setFastSend(false)
	client.output.__hx_this.writeString(hxrt.StringFromLiteral("final\n"), nil)
	client.output.__hx_this.flush()
	client.__hx_this.shutdown(false, true)
	var v_3 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("server.final="), accepted.input.__hx_this.readLine()))
	hxrt.Println(v_3)
	serverSawWriteEof := false
	hxrt.TryCatch(func() {
		accepted.input.__hx_this.readByte()
	}, func(hx_caught_8 any) {
		switch hx_typed_9 := hx_caught_8.(type) {
		case *haxe__io__Eof:
			hx_tmp := hx_typed_9
			_ = hx_tmp
			serverSawWriteEof = true
		default:
			hxrt.Throw(hx_caught_8)
		}
	})
	var v_4 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("server.write.eof="), hxrt.StdString(serverSawWriteEof)))
	hxrt.Println(v_4)
	accepted.output.__hx_this.writeString(hxrt.StringFromLiteral("after\n"), nil)
	accepted.output.__hx_this.flush()
	var v_5 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("client.after="), client.input.__hx_this.readLine()))
	hxrt.Println(v_5)
	peer := client.__hx_this.peerCertificate()
	var v_6 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("peer.cn="), peer.__hx_this.get_commonName()))
	hxrt.Println(v_6)
	var v_7 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("peer.issuer.cn="), peer.__hx_this.issuer(hxrt.StringFromLiteral("CN"))))
	hxrt.Println(v_7)
	var v_8 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("peer.next="), hxrt.StdString((peer.__hx_this.next() == nil))))
	hxrt.Println(v_8)
	client.__hx_this.close()
	accepted.__hx_this.close()
	server.__hx_this.close()
}
