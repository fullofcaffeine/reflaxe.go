package main

import "snapshot/hxrt"

func main() {
	socket := New_sys__ssl__Socket()
	socket.verifyCert = false
	socket.__hx_this.close()
	hxrt.Println(any(hxrt.StringFromLiteral("tls-socket-ready")))
}
