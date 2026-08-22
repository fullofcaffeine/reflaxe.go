package main

import "snapshot/hxrt"

type I_sys__net__UdpSocket interface {
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
	setBroadcast(enabled bool)
	sendTo(bytes *haxe__io__Bytes, pos int, length int, address *sys__net__Address) int
	readFrom(bytes *haxe__io__Bytes, pos int, length int, address *sys__net__Address) int
}

type sys__net__UdpSocket struct {
	*sys__net__Socket
	__hx_this I_sys__net__UdpSocket
}

func New_sys__net__UdpSocket() *sys__net__UdpSocket {
	self := &sys__net__UdpSocket{}
	self.sys__net__Socket = New_sys__net__Socket()
	self.sys__net__Socket.__hx_this = self
	self.__hx_this = self
	self.__hx_this.replaceHandle(hxrt.SocketNewUDP())
	return self
}

func (self *sys__net__UdpSocket) bind(host *sys__net__Host, port int) {
	if host == nil {
		hxrt.Throw(hxrt.StringFromLiteral("udp bind requires host"))
	}
	hxrt.SocketUdpBind(self.handle, host.__hx_this.toString(), port)
}

func (self *sys__net__UdpSocket) setBroadcast(enabled bool) {
	hxrt.SocketUdpSetBroadcast(self.handle, enabled)
}

func (self *sys__net__UdpSocket) sendTo(bytes *haxe__io__Bytes, pos int, length int, address *sys__net__Address) int {
	if ((((bytes == nil) || (address == nil)) || (pos < 0)) || (length < 0)) || (int(int32((hxrt.Int32Wrap(pos) + hxrt.Int32Wrap(length)))) > bytes.length) {
		hxrt.Throw(haxe__io__Error_OutsideBounds)
	}
	values := hxrt.NewArray()
	_g := 0
	_g1 := length
	for _g < _g1 {
		hx_post_1 := _g
		_g = int(int32((_g + 1)))
		index := hx_post_1
		values.Push(bytes.b[int(int32((hxrt.Int32Wrap(pos) + hxrt.Int32Wrap(index))))])
	}
	result := hxrt.SocketUdpSendTo(self.handle, func(hx_lambda_raw_3 []any) []int {
		hx_lambda_out_4 := make([]int, 0, len(hx_lambda_raw_3))
		for _, hx_lambda_item_5 := range hx_lambda_raw_3 {
			hx_lambda_out_4 = append(hx_lambda_out_4, func(hx_value_6 any) int {
				if hx_value_6 == nil {
					var hx_zero_7 int
					return hx_zero_7
				}
				return hx_value_6.(int)
			}(hx_lambda_item_5))
		}
		return hx_lambda_out_4
	}(values.Values()), address.host, address.port)
	if result.Status == hxrt.SocketIOBlocked {
		hxrt.Throw(haxe__io__Error_Blocked)
	}
	if result.Status == hxrt.SocketIOEOF {
		hxrt.Throw(New_haxe__io__Eof())
	}
	return result.Count
}

func (self *sys__net__UdpSocket) readFrom(bytes *haxe__io__Bytes, pos int, length int, address *sys__net__Address) int {
	if ((((bytes == nil) || (address == nil)) || (pos < 0)) || (length < 0)) || (int(int32((hxrt.Int32Wrap(pos) + hxrt.Int32Wrap(length)))) > bytes.length) {
		hxrt.Throw(haxe__io__Error_OutsideBounds)
	}
	if length == 0 {
		return 0
	}
	result := hxrt.SocketUdpReadFrom(self.handle, length)
	if result.Status == hxrt.SocketIOBlocked {
		hxrt.Throw(haxe__io__Error_Blocked)
	}
	if result.Status == hxrt.SocketIOEOF {
		hxrt.Throw(New_haxe__io__Eof())
	}
	_g := 0
	_g1 := result.Count
	for _g < _g1 {
		hx_post_8 := _g
		_g = int(int32((_g + 1)))
		index := hx_post_8
		value := result.Values[index]
		bytes.b[int(int32((hxrt.Int32Wrap(pos) + hxrt.Int32Wrap(index))))] = int(int32((hxrt.Int32Wrap(value) & hxrt.Int32Wrap(255))))
		bytes.__hx_rawValid = false
	}
	address.host = result.Host
	address.port = result.Port
	return result.Count
}
