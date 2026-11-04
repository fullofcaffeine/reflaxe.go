package main

import "snapshot/hxrt"

type I_sys__ssl__Socket interface {
	close()
	read() *string
	write(content *string)
	connect(host *sys__net__Host, port int)
	listen(connections int)
	shutdown(read bool, write bool)
	bind(host *sys__net__Host, port int)
	accept() *sys__net__Socket
	peer() hxrt.Anonymous
	host() hxrt.Anonymous
	setTimeout(timeout float64)
	waitForRead()
	setBlocking(blocking bool)
	setFastSend(fastSend bool)
	replaceHandle(next *hxrt.SocketHandle)
	handshake()
	setCA(cert *sys__ssl__Certificate)
	setHostname(name *string)
	setCertificate(cert *sys__ssl__Certificate, key *sys__ssl__Key)
	addSNICertificate(matcher *hxrt.Function[func(*string) bool], cert *sys__ssl__Certificate, key *sys__ssl__Key)
	peerCertificate() *sys__ssl__Certificate
}

type sys__ssl__Socket struct {
	*sys__net__Socket
	__hx_this  I_sys__ssl__Socket
	verifyCert bool
	caCert     *sys__ssl__Certificate
	hostname   *string
	ownCert    *sys__ssl__Certificate
	ownKey     *sys__ssl__Key
	sniConfig  *hxrt.SslSocketSNIConfig
}

func (self *sys__ssl__Socket) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_sys__ssl__Socket() *sys__ssl__Socket {
	return hx_init_sys__ssl__Socket(hx_alloc_sys__ssl__Socket(nil))
}

func hx_alloc_sys__ssl__Socket(owner any) *sys__ssl__Socket {
	self := &sys__ssl__Socket{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_sys__ssl__Socket)
	}
	self.sys__net__Socket = hx_alloc_sys__net__Socket(self.__hx_this)
	return self
}

func hx_init_sys__ssl__Socket(self *sys__ssl__Socket) *sys__ssl__Socket {
	hx_init_sys__net__Socket(self.sys__net__Socket)
	if (sys__ssl__Socket_DEFAULT_VERIFY_CERT == true) && (sys__ssl__Socket_DEFAULT_CA == nil) {
		hxrt.TryCatch(func() {
			sys__ssl__Socket_DEFAULT_CA = sys__ssl__Certificate_loadDefaults()
		}, func(hx_caught_1 any) {
			hx_exception_2 := haxe__Exception_caught(hx_caught_1)
			hx_payload_3 := haxe__Exception_catchValue(hx_caught_1)
			_ = hx_payload_3
			var hx_tmp *haxe__Exception = hx_recover_class_1(hx_exception_2)
			_ = hx_tmp
		})
	}
	self.verifyCert = sys__ssl__Socket_DEFAULT_VERIFY_CERT
	self.caCert = sys__ssl__Socket_DEFAULT_CA
	return self
}

func (self *sys__ssl__Socket) handshake() {
	hxrt.SslSocketHandshake(self.handle)
}

func (self *sys__ssl__Socket) shutdown(read bool, write bool) {
	self.sys__net__Socket.shutdown(read, write)
}

func (self *sys__ssl__Socket) setFastSend(fastSend bool) {
	self.sys__net__Socket.setFastSend(fastSend)
}

func (self *sys__ssl__Socket) setCA(cert *sys__ssl__Certificate) {
	self.caCert = cert
}

func (self *sys__ssl__Socket) setHostname(name *string) {
	self.hostname = name
}

func (self *sys__ssl__Socket) setCertificate(cert *sys__ssl__Certificate, key *sys__ssl__Key) {
	self.ownCert = cert
	self.ownKey = key
}

func (self *sys__ssl__Socket) addSNICertificate(matcher *hxrt.Function[func(*string) bool], cert *sys__ssl__Certificate, key *sys__ssl__Key) {
	if ((matcher == nil) || (cert == nil)) || (key == nil) {
		hxrt.Throw(haxe__Exception_thrown(hxrt.StringFromLiteral("sys.ssl.Socket.addSNICertificate requires callback, certificate, and key")))
	}
	self.sniConfig = hxrt.SslSocketAddSNICertificate(self.sniConfig, hxrt.FunctionNative(matcher), cert.handle, key.handle)
}

func (self *sys__ssl__Socket) connect(host *sys__net__Host, port int) {
	if host == nil {
		hxrt.Throw(haxe__Exception_thrown(hxrt.StringFromLiteral("socket connect requires host")))
	}
	endpoint := hxrt.SocketEndpointNew(host.__hx_this.toString(), host.host)
	hxrt.SslSocketConnect(self.handle, endpoint, port, (self.verifyCert != false), func() *hxrt.SslCertificate {
		var hx_if_4 *hxrt.SslCertificate
		if self.caCert == nil {
			hx_if_4 = nil
		} else {
			hx_if_4 = self.caCert.handle
		}
		return hx_if_4
	}(), self.hostname, func() *hxrt.SslCertificate {
		var hx_if_5 *hxrt.SslCertificate
		if self.ownCert == nil {
			hx_if_5 = nil
		} else {
			hx_if_5 = self.ownCert.handle
		}
		return hx_if_5
	}(), func() *hxrt.SslKey {
		var hx_if_6 *hxrt.SslKey
		if self.ownKey == nil {
			hx_if_6 = nil
		} else {
			hx_if_6 = self.ownKey.handle
		}
		return hx_if_6
	}())
}

func (self *sys__ssl__Socket) bind(host *sys__net__Host, port int) {
	if host == nil {
		hxrt.Throw(haxe__Exception_thrown(hxrt.StringFromLiteral("socket bind requires host")))
	}
	hxrt.SslSocketBind(self.handle, host.__hx_this.toString(), port, func() *hxrt.SslCertificate {
		var hx_if_7 *hxrt.SslCertificate
		if self.ownCert == nil {
			hx_if_7 = nil
		} else {
			hx_if_7 = self.ownCert.handle
		}
		return hx_if_7
	}(), func() *hxrt.SslKey {
		var hx_if_8 *hxrt.SslKey
		if self.ownKey == nil {
			hx_if_8 = nil
		} else {
			hx_if_8 = self.ownKey.handle
		}
		return hx_if_8
	}(), self.sniConfig)
}

func (self *sys__ssl__Socket) accept() *sys__net__Socket {
	result := hxrt.SocketAccept(self.handle)
	if result.Status == hxrt.SocketIOBlocked {
		hxrt.Throw(haxe__Exception_thrown(haxe__io__Error_Blocked))
	}
	if (result.Status == hxrt.SocketIOEOF) || (result.Handle == nil) {
		hxrt.Throw(haxe__Exception_thrown(New_haxe__io__Eof()))
	}
	accepted := New_sys__ssl__Socket()
	accepted.__hx_this.replaceHandle(result.Handle)
	hxrt.SslSocketHandshake(accepted.handle)
	return func(hx_value *sys__ssl__Socket) *sys__net__Socket {
		if hx_value == nil {
			return nil
		}
		return hx_value.sys__net__Socket
	}(accepted)
}

func (self *sys__ssl__Socket) peerCertificate() *sys__ssl__Certificate {
	certificateHandle := hxrt.SslSocketPeerCertificate(self.handle)
	var hx_if_9 *sys__ssl__Certificate
	if certificateHandle == nil {
		hx_if_9 = nil
	} else {
		hx_if_9 = New_sys__ssl__Certificate(certificateHandle)
	}
	return hx_if_9
}

var sys__ssl__Socket_DEFAULT_CA *sys__ssl__Certificate

var sys__ssl__Socket_DEFAULT_VERIFY_CERT bool = true
