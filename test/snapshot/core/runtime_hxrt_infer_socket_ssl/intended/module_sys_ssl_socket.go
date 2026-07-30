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
	peer() map[string]any
	host() map[string]any
	setTimeout(timeout float64)
	waitForRead()
	setBlocking(blocking bool)
	setFastSend(fastSend bool)
	replaceHandle(next *hxrt.SocketHandle)
	handshake()
	setCA(cert *sys__ssl__Certificate)
	setHostname(name *string)
	setCertificate(cert *sys__ssl__Certificate, key *sys__ssl__Key)
	addSNICertificate(matcher func(*string) bool, cert *sys__ssl__Certificate, key *sys__ssl__Key)
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

func New_sys__ssl__Socket() *sys__ssl__Socket {
	self := &sys__ssl__Socket{}
	self.sys__net__Socket = New_sys__net__Socket()
	self.sys__net__Socket.__hx_this = self
	self.__hx_this = self
	if (sys__ssl__Socket_DEFAULT_VERIFY_CERT == true) && (sys__ssl__Socket_DEFAULT_CA == nil) {
		hxrt.TryCatch(func() {
			sys__ssl__Socket_DEFAULT_CA = sys__ssl__Certificate_loadDefaults()
		}, func(hx_caught_1 any) {
			hx_tmp := hxrt.ExceptionCaught(hx_caught_1)
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

func (self *sys__ssl__Socket) addSNICertificate(matcher func(*string) bool, cert *sys__ssl__Certificate, key *sys__ssl__Key) {
	if ((matcher == nil) || (cert == nil)) || (key == nil) {
		hxrt.Throw(hxrt.StringFromLiteral("sys.ssl.Socket.addSNICertificate requires callback, certificate, and key"))
	}
	self.sniConfig = hxrt.SslSocketAddSNICertificate(self.sniConfig, matcher, cert.handle, key.handle)
}

func (self *sys__ssl__Socket) connect(host *sys__net__Host, port int) {
	if host == nil {
		hxrt.Throw(hxrt.StringFromLiteral("socket connect requires host"))
	}
	endpoint := hxrt.SocketEndpointNew(host.__hx_this.toString(), host.host)
	hxrt.SslSocketConnect(self.handle, endpoint, port, (self.verifyCert != false), func() *hxrt.SslCertificate {
		var hx_if_3 *hxrt.SslCertificate
		if self.caCert == nil {
			hx_if_3 = nil
		} else {
			hx_if_3 = self.caCert.handle
		}
		return hx_if_3
	}(), self.hostname, func() *hxrt.SslCertificate {
		var hx_if_4 *hxrt.SslCertificate
		if self.ownCert == nil {
			hx_if_4 = nil
		} else {
			hx_if_4 = self.ownCert.handle
		}
		return hx_if_4
	}(), func() *hxrt.SslKey {
		var hx_if_5 *hxrt.SslKey
		if self.ownKey == nil {
			hx_if_5 = nil
		} else {
			hx_if_5 = self.ownKey.handle
		}
		return hx_if_5
	}())
}

func (self *sys__ssl__Socket) bind(host *sys__net__Host, port int) {
	if host == nil {
		hxrt.Throw(hxrt.StringFromLiteral("socket bind requires host"))
	}
	hxrt.SslSocketListen(self.handle, host.__hx_this.toString(), port, func() *hxrt.SslCertificate {
		var hx_if_6 *hxrt.SslCertificate
		if self.ownCert == nil {
			hx_if_6 = nil
		} else {
			hx_if_6 = self.ownCert.handle
		}
		return hx_if_6
	}(), func() *hxrt.SslKey {
		var hx_if_7 *hxrt.SslKey
		if self.ownKey == nil {
			hx_if_7 = nil
		} else {
			hx_if_7 = self.ownKey.handle
		}
		return hx_if_7
	}(), self.sniConfig)
}

func (self *sys__ssl__Socket) accept() *sys__net__Socket {
	result := hxrt.SocketAccept(self.handle)
	if result.Status == hxrt.SocketIOBlocked {
		hxrt.Throw(haxe__io__Error_Blocked)
	}
	if (result.Status == hxrt.SocketIOEOF) || (result.Handle == nil) {
		hxrt.Throw(New_haxe__io__Eof())
	}
	accepted := New_sys__ssl__Socket()
	accepted.__hx_this.replaceHandle(result.Handle)
	hxrt.SslSocketHandshake(accepted.handle)
	return accepted.sys__net__Socket
}

func (self *sys__ssl__Socket) peerCertificate() *sys__ssl__Certificate {
	certificateHandle := hxrt.SslSocketPeerCertificate(self.handle)
	var hx_if_8 *sys__ssl__Certificate
	if certificateHandle == nil {
		hx_if_8 = nil
	} else {
		hx_if_8 = New_sys__ssl__Certificate(certificateHandle)
	}
	return hx_if_8
}

var sys__ssl__Socket_DEFAULT_CA *sys__ssl__Certificate

var sys__ssl__Socket_DEFAULT_VERIFY_CERT bool = true
