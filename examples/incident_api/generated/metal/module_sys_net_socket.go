package main

import "examples_incident_api_metal/hxrt"

type I_sys__net__Socket interface {
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
}

type sys__net__Socket struct {
	__hx_this I_sys__net__Socket
	input     haxe__io__Input
	output    haxe__io__Output
	custom    any
	handle    *hxrt.SocketHandle
}

func New_sys__net__Socket() *sys__net__Socket {
	self := &sys__net__Socket{}
	self.__hx_this = self
	self.replaceHandle(hxrt.SocketNewTCP())
	return self
}

func (self *sys__net__Socket) close() {
	hxrt.SocketClose(self.handle)
}

func (self *sys__net__Socket) read() *string {
	return self.input.readAll().toString()
}

func (self *sys__net__Socket) write(content *string) {
	self.output.writeString(content)
}

func (self *sys__net__Socket) connect(host *sys__net__Host, port int) {
	if host == nil {
		hxrt.Throw(hxrt.StringFromLiteral("socket connect requires host"))
	}
	hxrt.SocketConnectTCP(self.handle, host.toString(), port)
}

func (self *sys__net__Socket) listen(connections int) {
	hxrt.SocketListen(self.handle, connections)
}

func (self *sys__net__Socket) shutdown(read bool, write bool) {
	hxrt.SocketShutdown(self.handle, read, write)
}

func (self *sys__net__Socket) bind(host *sys__net__Host, port int) {
	if host == nil {
		hxrt.Throw(hxrt.StringFromLiteral("socket bind requires host"))
	}
	hxrt.SocketBindTCP(self.handle, host.toString(), port)
}

func (self *sys__net__Socket) accept() *sys__net__Socket {
	result := hxrt.SocketAccept(self.handle)
	if result.Status == hxrt.SocketIOBlocked {
		hxrt.Throw(haxe__io__Error_Blocked)
	}
	if (result.Status == hxrt.SocketIOEOF) || (result.Handle == nil) {
		hxrt.Throw(New_haxe__io__Eof())
	}
	accepted := New_sys__net__Socket()
	accepted.replaceHandle(result.Handle)
	return accepted
}

func (self *sys__net__Socket) peer() map[string]any {
	return sys__net__Socket_publicAddress(hxrt.SocketPeer(self.handle))
}

func (self *sys__net__Socket) host() map[string]any {
	return sys__net__Socket_publicAddress(hxrt.SocketHost(self.handle))
}

func (self *sys__net__Socket) setTimeout(timeout float64) {
	hxrt.SocketSetTimeout(self.handle, timeout)
}

func (self *sys__net__Socket) waitForRead() {
	hxrt.SocketWaitForRead(self.handle)
}

func (self *sys__net__Socket) setBlocking(blocking bool) {
	hxrt.SocketSetBlocking(self.handle, blocking)
}

func (self *sys__net__Socket) setFastSend(fastSend bool) {
	hxrt.SocketSetFastSend(self.handle, fastSend)
}

func (self *sys__net__Socket) replaceHandle(next *hxrt.SocketHandle) {
	if (self.handle != nil) && (self.handle != next) {
		hxrt.SocketClose(self.handle)
	}
	self.handle = next
	self.input = New_sys__net__SocketInput(next)
	self.output = New_sys__net__SocketOutput(next)
}

func sys__net__Socket_pick(source *hxrt.Array, indexes []int) *hxrt.Array {
	selected := hxrt.NewArray()
	_g := 0
	_g1 := len(indexes)
	for _g < _g1 {
		hx_post_102 := _g
		_g = int(int32((_g + 1)))
		index := hx_post_102
		sourceIndex := indexes[index]
		if (sourceIndex >= 0) && (sourceIndex < source.Len()) {
			selected.Push(source.Get(sourceIndex))
		}
	}
	return selected
}

func sys__net__Socket_publicAddress(address *hxrt.SocketAddress) map[string]any {
	if address == nil {
		hx_obj_104 := map[string]any{}
		hx_obj_104["host"] = sys__net__Host_fromIPv4(0)
		hx_obj_104["port"] = 0
		return hx_obj_104
	}
	hx_obj_105 := map[string]any{}
	hx_obj_105["host"] = sys__net__Host_fromIPv4(address.Host)
	hx_obj_105["port"] = address.Port
	return hx_obj_105
}

func sys__net__Socket_select_(read *hxrt.Array, write *hxrt.Array, others *hxrt.Array, timeout any) map[string]any {
	if read == nil {
		read = hxrt.NewArray()
	}
	if write == nil {
		write = hxrt.NewArray()
	}
	if others == nil {
		others = hxrt.NewArray()
	}
	readHandles := func() []*hxrt.SocketHandle {
		_g := hxrt.NewArray()
		_g1 := 0
		for _g1 < read.Len() {
			socket := func(hx_value_106 any) *sys__net__Socket {
				if hx_value_106 == nil {
					var hx_zero_107 *sys__net__Socket
					return hx_zero_107
				}
				return hx_value_106.(*sys__net__Socket)
			}(read.Get(_g1))
			_g1 = int(int32((_g1 + 1)))
			_g.Push(socket.handle)
		}
		return func(hx_lambda_raw_109 []any) []*hxrt.SocketHandle {
			hx_lambda_out_110 := make([]*hxrt.SocketHandle, 0, len(hx_lambda_raw_109))
			for _, hx_lambda_item_111 := range hx_lambda_raw_109 {
				hx_lambda_out_110 = append(hx_lambda_out_110, func(hx_value_112 any) *hxrt.SocketHandle {
					if hx_value_112 == nil {
						var hx_zero_113 *hxrt.SocketHandle
						return hx_zero_113
					}
					return hx_value_112.(*hxrt.SocketHandle)
				}(hx_lambda_item_111))
			}
			return hx_lambda_out_110
		}(_g.Values())
	}()
	writeHandles := func() []*hxrt.SocketHandle {
		_g_1 := hxrt.NewArray()
		_g1_1 := 0
		for _g1_1 < write.Len() {
			socket_1 := func(hx_value_114 any) *sys__net__Socket {
				if hx_value_114 == nil {
					var hx_zero_115 *sys__net__Socket
					return hx_zero_115
				}
				return hx_value_114.(*sys__net__Socket)
			}(write.Get(_g1_1))
			_g1_1 = int(int32((_g1_1 + 1)))
			_g_1.Push(socket_1.handle)
		}
		return func(hx_lambda_raw_117 []any) []*hxrt.SocketHandle {
			hx_lambda_out_118 := make([]*hxrt.SocketHandle, 0, len(hx_lambda_raw_117))
			for _, hx_lambda_item_119 := range hx_lambda_raw_117 {
				hx_lambda_out_118 = append(hx_lambda_out_118, func(hx_value_120 any) *hxrt.SocketHandle {
					if hx_value_120 == nil {
						var hx_zero_121 *hxrt.SocketHandle
						return hx_zero_121
					}
					return hx_value_120.(*hxrt.SocketHandle)
				}(hx_lambda_item_119))
			}
			return hx_lambda_out_118
		}(_g_1.Values())
	}()
	otherHandles := func() []*hxrt.SocketHandle {
		_g_2 := hxrt.NewArray()
		_g1_2 := 0
		for _g1_2 < others.Len() {
			socket_2 := func(hx_value_122 any) *sys__net__Socket {
				if hx_value_122 == nil {
					var hx_zero_123 *sys__net__Socket
					return hx_zero_123
				}
				return hx_value_122.(*sys__net__Socket)
			}(others.Get(_g1_2))
			_g1_2 = int(int32((_g1_2 + 1)))
			_g_2.Push(socket_2.handle)
		}
		return func(hx_lambda_raw_125 []any) []*hxrt.SocketHandle {
			hx_lambda_out_126 := make([]*hxrt.SocketHandle, 0, len(hx_lambda_raw_125))
			for _, hx_lambda_item_127 := range hx_lambda_raw_125 {
				hx_lambda_out_126 = append(hx_lambda_out_126, func(hx_value_128 any) *hxrt.SocketHandle {
					if hx_value_128 == nil {
						var hx_zero_129 *hxrt.SocketHandle
						return hx_zero_129
					}
					return hx_value_128.(*hxrt.SocketHandle)
				}(hx_lambda_item_127))
			}
			return hx_lambda_out_126
		}(_g_2.Values())
	}()
	result := hxrt.SocketSelect(readHandles, writeHandles, otherHandles, func() float64 {
		var hx_if_130 float64
		if timeout == nil {
			hx_if_130 = 0.0
		} else {
			hx_if_130 = timeout.(float64)
		}
		return hx_if_130
	}(), (timeout != nil))
	hx_obj_131 := map[string]any{}
	hx_obj_131["read"] = sys__net__Socket_pick(read, result.Read)
	hx_obj_131["write"] = sys__net__Socket_pick(write, result.Write)
	hx_obj_131["others"] = sys__net__Socket_pick(others, result.Others)
	return hx_obj_131
}
