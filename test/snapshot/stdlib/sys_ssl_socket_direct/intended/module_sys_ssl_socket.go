package main

import "snapshot/hxrt"

type I_sys__ssl__Socket interface {
	handshake()
	setCA(cert *sys__ssl__Certificate)
	setHostname(name *string)
	setCertificate(cert *sys__ssl__Certificate, key *sys__ssl__Key)
	addSNICertificate(cbServernameMatch func(*string) bool, cert *sys__ssl__Certificate, key *sys__ssl__Key)
	connect(host *sys__net__Host, port int)
	bind(host *sys__net__Host, port int)
	accept() *sys__net__Socket
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
	sniConfig  any
}

func New_sys__ssl__Socket() *sys__ssl__Socket {
	self := &sys__ssl__Socket{}
	self.sys__net__Socket = New_sys__net__Socket()
	self.__hx_this = self
	if (sys__ssl__Socket_DEFAULT_VERIFY_CERT == true) && (sys__ssl__Socket_DEFAULT_CA == nil) {
		hxrt.TryCatch(func() {
			sys__ssl__Socket_DEFAULT_CA = sys__ssl__Certificate_loadDefaults()
		}, func(hx_caught_4 any) {
			hx_tmp := hxrt.ExceptionCaught(hx_caught_4)
			_ = hx_tmp
		})
	}
	self.verifyCert = sys__ssl__Socket_DEFAULT_VERIFY_CERT
	self.caCert = sys__ssl__Socket_DEFAULT_CA
	return self
}

func (self *sys__ssl__Socket) handshake() {
	_ = func() int { hxrt.SslSocketHandshake(self.hxrt__socket_conn()); return 0 }()
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

func (self *sys__ssl__Socket) addSNICertificate(cbServernameMatch func(*string) bool, cert *sys__ssl__Certificate, key *sys__ssl__Key) {
	if ((cbServernameMatch == nil) || (cert == nil)) || (key == nil) {
		hxrt.Throw(hxrt.StringFromLiteral("sys.ssl.Socket.addSNICertificate requires callback, certificate, and key"))
	}
	self.sniConfig = hxrt.SslSocketAddSNICertificate(self.sniConfig, cbServernameMatch, cert.handle, key.handle)
}

func (self *sys__ssl__Socket) connect(host *sys__net__Host, port int) {
	if host == nil {
		hxrt.Throw(hxrt.StringFromLiteral("socket connect requires host"))
	}
	resolvedHost := host.toString()
	_ = resolvedHost
	if hxrt.StringEqualStringPtr(resolvedHost, nil) {
		hxrt.Throw(hxrt.StringFromLiteral("socket connect requires host"))
	}
	_ = func() int {
		conn := hxrt.SslSocketConnect(resolvedHost, port, (self.verifyCert != false), func() any {
			var hx_if_6 any
			if self.caCert == nil {
				hx_if_6 = nil
			} else {
				hx_if_6 = self.caCert.handle
			}
			return hx_if_6
		}(), self.hostname, func() any {
			var hx_if_7 any
			if self.ownCert == nil {
				hx_if_7 = nil
			} else {
				hx_if_7 = self.ownCert.handle
			}
			return hx_if_7
		}(), func() any {
			var hx_if_8 any
			if self.ownKey == nil {
				hx_if_8 = nil
			} else {
				hx_if_8 = self.ownKey.handle
			}
			return hx_if_8
		}())
		self.hxrt__socket_setConn(conn)
		return 0
	}()
}

func (self *sys__ssl__Socket) bind(host *sys__net__Host, port int) {
	if host == nil {
		hxrt.Throw(hxrt.StringFromLiteral("socket bind requires host"))
	}
	resolvedHost := host.toString()
	_ = resolvedHost
	if hxrt.StringEqualStringPtr(resolvedHost, nil) {
		hxrt.Throw(hxrt.StringFromLiteral("socket bind requires host"))
	}
	_ = func() int {
		listener := hxrt.SslSocketListen(resolvedHost, port, func() any {
			var hx_if_9 any
			if self.ownCert == nil {
				hx_if_9 = nil
			} else {
				hx_if_9 = self.ownCert.handle
			}
			return hx_if_9
		}(), func() any {
			var hx_if_10 any
			if self.ownKey == nil {
				hx_if_10 = nil
			} else {
				hx_if_10 = self.ownKey.handle
			}
			return hx_if_10
		}(), self.sniConfig)
		if self.listener != nil {
			_ = self.listener.Close()
		}
		self.listener = listener
		self.hxrt__socket_applyListenerDeadline()
		return 0
	}()
}

func (self *sys__ssl__Socket) accept() *sys__net__Socket {
	accepted := self.sys__net__Socket.accept()
	_ = accepted
	_ = func() int { hxrt.SslSocketHandshake(accepted.hxrt__socket_conn()); return 0 }()
	return accepted
}

func (self *sys__ssl__Socket) peerCertificate() *sys__ssl__Certificate {
	var handle any = hxrt.SslSocketPeerCertificate(self.hxrt__socket_conn())
	var hx_if_11 *sys__ssl__Certificate
	if hxrt.AnyEqualsNull(handle) {
		hx_if_11 = nil
	} else {
		hx_if_11 = New_sys__ssl__Certificate(handle)
	}
	return hx_if_11
}

var sys__ssl__Socket_DEFAULT_CA *sys__ssl__Certificate

var sys__ssl__Socket_DEFAULT_VERIFY_CERT bool = true
