package hxrt

import (
	"crypto/tls"
	"errors"
	"net"
	"strconv"
	"sync"
)

type sslSocketSNIEntry struct {
	match func(*string) bool
	cert  tls.Certificate
}

// SslSocketSNIConfig is the typed opaque SNI certificate-selection handle.
//
// What: Retains Haxe hostname matchers paired with parsed native certificates.
// Why: tls.Config invokes selection during the native handshake, so the table
// cannot live as a portable Haxe collection without crossing an untyped callback.
// How: Staged sys.ssl.Socket builds this handle through one typed capability and
// passes it back when installing a TLS listener.
type SslSocketSNIConfig struct {
	mu      sync.RWMutex
	entries []sslSocketSNIEntry
}

// SslSocketAddSNICertificate returns the typed SNI table containing one more entry.
func SslSocketAddSNICertificate(config *SslSocketSNIConfig, match func(*string) bool, cert *SslCertificate, key *SslKey) *SslSocketSNIConfig {
	if config == nil {
		config = &SslSocketSNIConfig{}
	}
	if match == nil {
		socketThrow(errors.New("sys.ssl.Socket.addSNICertificate callback is nil"))
		return config
	}
	pair, err := sslKeyPair(cert, key)
	if err != nil {
		socketThrow(err)
		return config
	}
	config.mu.Lock()
	config.entries = append(config.entries, sslSocketSNIEntry{match: match, cert: pair})
	config.mu.Unlock()
	return config
}

func sslSocketClientConfig(verifyCert bool, ca *SslCertificate, serverName *string, cert *SslCertificate, key *SslKey) *tls.Config {
	config := &tls.Config{InsecureSkipVerify: !verifyCert}
	if verifyCert {
		if pool := sslCertPool(ca); pool != nil {
			config.RootCAs = pool
		}
	}
	if serverName != nil && *StdString(serverName) != "" {
		config.ServerName = *StdString(serverName)
	}
	if cert != nil || key != nil {
		pair, err := sslKeyPair(cert, key)
		if err != nil {
			socketThrow(err)
			return nil
		}
		config.Certificates = []tls.Certificate{pair}
	}
	return config
}

// SslSocketConnect installs a TLS client connection into an existing typed socket handle.
//
// What: Dials the endpoint's numeric address while verifying and advertising
// its logical hostname by default.
// Why: A resolved IP is transport routing, not the DNS identity named by a
// certificate; conflating them breaks ordinary certificate verification and SNI.
// How: An explicit setHostname value wins, otherwise the endpoint's logical
// host becomes tls.Config.ServerName before the numeric dial begins.
func SslSocketConnect(handle *SocketHandle, endpoint *SocketEndpoint, port int, verifyCert bool, ca *SslCertificate, serverName *string, cert *SslCertificate, key *SslKey) {
	if handle == nil || endpoint == nil || endpoint.NetworkAddress == "" {
		socketThrow(errors.New("socket connect requires host"))
		return
	}
	effectiveServerName := serverName
	if effectiveServerName == nil || *StdString(effectiveServerName) == "" {
		effectiveServerName = StringFromLiteral(endpoint.LogicalHost)
	}
	config := sslSocketClientConfig(verifyCert, ca, effectiveServerName, cert, key)
	if config == nil {
		return
	}
	conn, err := tls.DialWithDialer(handle.dialer(), "tcp4", net.JoinHostPort(endpoint.NetworkAddress, strconv.Itoa(port)), config)
	if err != nil {
		socketThrow(err)
		return
	}
	handle.installConn(conn)
}

func sslSocketServerConfig(cert *SslCertificate, key *SslKey, sni *SslSocketSNIConfig) *tls.Config {
	pair, err := sslKeyPair(cert, key)
	if err != nil {
		socketThrow(err)
		return nil
	}
	config := &tls.Config{Certificates: []tls.Certificate{pair}}
	var entries []sslSocketSNIEntry
	if sni != nil {
		sni.mu.RLock()
		entries = append(entries, sni.entries...)
		sni.mu.RUnlock()
	}
	if len(entries) > 0 {
		config.GetCertificate = func(clientHello *tls.ClientHelloInfo) (*tls.Certificate, error) {
			serverName := StringFromLiteral(clientHello.ServerName)
			for index := range entries {
				entry := &entries[index]
				if entry.match != nil && entry.match(serverName) {
					return &entry.cert, nil
				}
			}
			return &pair, nil
		}
	}
	return config
}

// SslSocketListen installs a TLS listener while retaining its TCP deadline authority.
func SslSocketListen(handle *SocketHandle, host *string, port int, cert *SslCertificate, key *SslKey, sni *SslSocketSNIConfig) {
	if handle == nil || host == nil {
		socketThrow(errors.New("socket bind requires host"))
		return
	}
	config := sslSocketServerConfig(cert, key, sni)
	if config == nil {
		return
	}
	address, err := net.ResolveTCPAddr("tcp4", net.JoinHostPort(*StdString(host), strconv.Itoa(port)))
	if err != nil {
		socketThrow(err)
		return
	}
	tcpListener, err := net.ListenTCP("tcp4", address)
	if err != nil {
		socketThrow(err)
		return
	}
	listener := tls.NewListener(tcpListener, config)
	handle.installListener(listener, tcpListener)
}

// SslSocketHandshake performs the native handshake for the handle's current TLS connection.
func SslSocketHandshake(handle *SocketHandle) {
	if handle == nil {
		return
	}
	conn := handle.snapshotConn()
	if conn == nil {
		return
	}
	if handshaker, ok := conn.(interface{ Handshake() error }); ok {
		if err := handshaker.Handshake(); err != nil {
			socketThrow(err)
		}
	}
}

// SslSocketPeerCertificate returns the typed leaf/chain handle for the TLS peer.
func SslSocketPeerCertificate(handle *SocketHandle) *SslCertificate {
	if handle == nil {
		return nil
	}
	conn := handle.snapshotConn()
	tlsConn, ok := conn.(*tls.Conn)
	if !ok || tlsConn == nil {
		return nil
	}
	if err := tlsConn.Handshake(); err != nil {
		socketThrow(err)
		return nil
	}
	state := tlsConn.ConnectionState()
	if len(state.PeerCertificates) == 0 {
		return nil
	}
	return newSslCertificate(state.PeerCertificates, nil)
}
