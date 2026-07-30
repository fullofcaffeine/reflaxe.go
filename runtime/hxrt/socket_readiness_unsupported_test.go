//go:build !linux && !darwin

package hxrt

import (
	"net"
	"testing"
)

func socketTestSaturateTCPWriteBufferNative(t *testing.T, _ *net.TCPConn) {
	t.Helper()
	t.Skip("native write-readiness runtime evidence is limited to Linux and Darwin")
}
