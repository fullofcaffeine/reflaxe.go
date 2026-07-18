package main

import "snapshot/hxrt"

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
	input     *haxe__io__Input
	output    *haxe__io__Output
	custom    any
	handle    *hxrt.SocketHandle
}

func New_sys__net__Socket() *sys__net__Socket {
	self := &sys__net__Socket{}
	self.__hx_this = self
	self.__hx_this.replaceHandle(hxrt.SocketNewTCP())
	return self
}

func (self *sys__net__Socket) close() {
	hxrt.SocketClose(self.handle)
}

func (self *sys__net__Socket) read() *string {
	return self.input.__hx_this.readAll(nil).__hx_this.toString()
}

func (self *sys__net__Socket) write(content *string) {
	self.output.__hx_this.writeString(content, nil)
}

func (self *sys__net__Socket) connect(host *sys__net__Host, port int) {
	if host == nil {
		hxrt.Throw(hxrt.StringFromLiteral("socket connect requires host"))
	}
	hxrt.SocketConnectTCP(self.handle, host.__hx_this.toString(), port)
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
	hxrt.SocketBindTCP(self.handle, host.__hx_this.toString(), port)
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
	accepted.__hx_this.replaceHandle(result.Handle)
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
	self.input = New_sys__net__SocketInput(next).haxe__io__Input
	self.output = New_sys__net__SocketOutput(next).haxe__io__Output
}

func sys__net__Socket_pick(source *hxrt.Array, indexes []int) *hxrt.Array {
	selected := hxrt.NewArray()
	_g := 0
	_g1 := len(indexes)
	for _g < _g1 {
		hx_post_24 := _g
		_g = int(int32((_g + 1)))
		index := hx_post_24
		sourceIndex := indexes[index]
		if (sourceIndex >= 0) && (sourceIndex < source.Len()) {
			selected.Push(source.Get(sourceIndex))
		}
	}
	return selected
}

func sys__net__Socket_publicAddress(address *hxrt.SocketAddress) map[string]any {
	if address == nil {
		hx_obj_26 := map[string]any{}
		hx_obj_26["host"] = sys__net__Host_fromIPv4(0)
		hx_obj_26["port"] = 0
		return hx_obj_26
	}
	hx_obj_27 := map[string]any{}
	hx_obj_27["host"] = sys__net__Host_fromIPv4(address.Host)
	hx_obj_27["port"] = address.Port
	return hx_obj_27
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
			socket := func(hx_value_28 any) *sys__net__Socket {
				if hx_value_28 == nil {
					var hx_zero_29 *sys__net__Socket
					return hx_zero_29
				}
				return hx_value_28.(*sys__net__Socket)
			}(read.Get(_g1))
			_g1 = int(int32((_g1 + 1)))
			_g.Push(socket.handle)
		}
		return func(hx_lambda_raw_31 []any) []*hxrt.SocketHandle {
			hx_lambda_out_32 := make([]*hxrt.SocketHandle, 0, len(hx_lambda_raw_31))
			for _, hx_lambda_item_33 := range hx_lambda_raw_31 {
				hx_lambda_out_32 = append(hx_lambda_out_32, func(hx_value_34 any) *hxrt.SocketHandle {
					if hx_value_34 == nil {
						var hx_zero_35 *hxrt.SocketHandle
						return hx_zero_35
					}
					return hx_value_34.(*hxrt.SocketHandle)
				}(hx_lambda_item_33))
			}
			return hx_lambda_out_32
		}(_g.Values())
	}()
	writeHandles := func() []*hxrt.SocketHandle {
		_g_1 := hxrt.NewArray()
		_g1_1 := 0
		for _g1_1 < write.Len() {
			socket_1 := func(hx_value_36 any) *sys__net__Socket {
				if hx_value_36 == nil {
					var hx_zero_37 *sys__net__Socket
					return hx_zero_37
				}
				return hx_value_36.(*sys__net__Socket)
			}(write.Get(_g1_1))
			_g1_1 = int(int32((_g1_1 + 1)))
			_g_1.Push(socket_1.handle)
		}
		return func(hx_lambda_raw_39 []any) []*hxrt.SocketHandle {
			hx_lambda_out_40 := make([]*hxrt.SocketHandle, 0, len(hx_lambda_raw_39))
			for _, hx_lambda_item_41 := range hx_lambda_raw_39 {
				hx_lambda_out_40 = append(hx_lambda_out_40, func(hx_value_42 any) *hxrt.SocketHandle {
					if hx_value_42 == nil {
						var hx_zero_43 *hxrt.SocketHandle
						return hx_zero_43
					}
					return hx_value_42.(*hxrt.SocketHandle)
				}(hx_lambda_item_41))
			}
			return hx_lambda_out_40
		}(_g_1.Values())
	}()
	otherHandles := func() []*hxrt.SocketHandle {
		_g_2 := hxrt.NewArray()
		_g1_2 := 0
		for _g1_2 < others.Len() {
			socket_2 := func(hx_value_44 any) *sys__net__Socket {
				if hx_value_44 == nil {
					var hx_zero_45 *sys__net__Socket
					return hx_zero_45
				}
				return hx_value_44.(*sys__net__Socket)
			}(others.Get(_g1_2))
			_g1_2 = int(int32((_g1_2 + 1)))
			_g_2.Push(socket_2.handle)
		}
		return func(hx_lambda_raw_47 []any) []*hxrt.SocketHandle {
			hx_lambda_out_48 := make([]*hxrt.SocketHandle, 0, len(hx_lambda_raw_47))
			for _, hx_lambda_item_49 := range hx_lambda_raw_47 {
				hx_lambda_out_48 = append(hx_lambda_out_48, func(hx_value_50 any) *hxrt.SocketHandle {
					if hx_value_50 == nil {
						var hx_zero_51 *hxrt.SocketHandle
						return hx_zero_51
					}
					return hx_value_50.(*hxrt.SocketHandle)
				}(hx_lambda_item_49))
			}
			return hx_lambda_out_48
		}(_g_2.Values())
	}()
	result := hxrt.SocketSelect(readHandles, writeHandles, otherHandles, func() float64 {
		var hx_if_52 float64
		if timeout == nil {
			hx_if_52 = 0.0
		} else {
			hx_if_52 = timeout.(float64)
		}
		return hx_if_52
	}(), (timeout != nil))
	hx_obj_53 := map[string]any{}
	hx_obj_53["read"] = sys__net__Socket_pick(read, result.Read)
	hx_obj_53["write"] = sys__net__Socket_pick(write, result.Write)
	hx_obj_53["others"] = sys__net__Socket_pick(others, result.Others)
	return hx_obj_53
}
