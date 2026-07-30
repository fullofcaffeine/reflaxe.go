//go:build linux

package hxrt

import (
	"net"
	"strconv"
	"testing"
	"time"
)

func TestSocketLinuxBacklogBoundsPendingConnections(t *testing.T) {
	listener := SocketNewTCP()
	defer SocketClose(listener)
	SocketBindTCP(listener, socketTestString(socketTestHost), 0)
	SocketListen(listener, 1)
	bound := SocketHost(listener)
	address := net.JoinHostPort(socketTestHost, strconv.Itoa(bound.Port))

	clients := make([]net.Conn, 0, 16)
	defer func() {
		for _, client := range clients {
			_ = client.Close()
		}
	}()
	for index := 0; index < 16; index++ {
		client, err := net.DialTimeout("tcp4", address, 100*time.Millisecond)
		if err != nil {
			break
		}
		clients = append(clients, client)
	}
	if len(clients) == 0 {
		t.Fatal("backlog listener refused its first pending connection")
	}
	if len(clients) == cap(clients) {
		t.Fatal("backlog=1 admitted every pending connection without an accept")
	}
}
