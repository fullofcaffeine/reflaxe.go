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
		hx_post_35 := _g
		_g = int(int32((_g + 1)))
		index := hx_post_35
		sourceIndex := indexes[index]
		if (sourceIndex >= 0) && (sourceIndex < source.Len()) {
			selected.Push(source.Get(sourceIndex))
		}
	}
	return selected
}

func sys__net__Socket_publicAddress(address *hxrt.SocketAddress) map[string]any {
	if address == nil {
		hx_obj_37 := map[string]any{}
		hx_obj_37["host"] = sys__net__Host_fromIPv4(0)
		hx_obj_37["port"] = 0
		return hx_obj_37
	}
	hx_obj_38 := map[string]any{}
	hx_obj_38["host"] = sys__net__Host_fromIPv4(address.Host)
	hx_obj_38["port"] = address.Port
	return hx_obj_38
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
			socket := func(hx_value_39 any) *sys__net__Socket {
				if hx_value_39 == nil {
					var hx_zero_40 *sys__net__Socket
					return hx_zero_40
				}
				return hx_value_39.(*sys__net__Socket)
			}(read.Get(_g1))
			_g1 = int(int32((_g1 + 1)))
			_g.Push(socket.handle)
		}
		return func(hx_lambda_raw_42 []any) []*hxrt.SocketHandle {
			hx_lambda_out_43 := make([]*hxrt.SocketHandle, 0, len(hx_lambda_raw_42))
			for _, hx_lambda_item_44 := range hx_lambda_raw_42 {
				hx_lambda_out_43 = append(hx_lambda_out_43, func(hx_value_45 any) *hxrt.SocketHandle {
					if hx_value_45 == nil {
						var hx_zero_46 *hxrt.SocketHandle
						return hx_zero_46
					}
					return hx_value_45.(*hxrt.SocketHandle)
				}(hx_lambda_item_44))
			}
			return hx_lambda_out_43
		}(_g.Values())
	}()
	writeHandles := func() []*hxrt.SocketHandle {
		_g_1 := hxrt.NewArray()
		_g1_1 := 0
		for _g1_1 < write.Len() {
			socket_1 := func(hx_value_47 any) *sys__net__Socket {
				if hx_value_47 == nil {
					var hx_zero_48 *sys__net__Socket
					return hx_zero_48
				}
				return hx_value_47.(*sys__net__Socket)
			}(write.Get(_g1_1))
			_g1_1 = int(int32((_g1_1 + 1)))
			_g_1.Push(socket_1.handle)
		}
		return func(hx_lambda_raw_50 []any) []*hxrt.SocketHandle {
			hx_lambda_out_51 := make([]*hxrt.SocketHandle, 0, len(hx_lambda_raw_50))
			for _, hx_lambda_item_52 := range hx_lambda_raw_50 {
				hx_lambda_out_51 = append(hx_lambda_out_51, func(hx_value_53 any) *hxrt.SocketHandle {
					if hx_value_53 == nil {
						var hx_zero_54 *hxrt.SocketHandle
						return hx_zero_54
					}
					return hx_value_53.(*hxrt.SocketHandle)
				}(hx_lambda_item_52))
			}
			return hx_lambda_out_51
		}(_g_1.Values())
	}()
	otherHandles := func() []*hxrt.SocketHandle {
		_g_2 := hxrt.NewArray()
		_g1_2 := 0
		for _g1_2 < others.Len() {
			socket_2 := func(hx_value_55 any) *sys__net__Socket {
				if hx_value_55 == nil {
					var hx_zero_56 *sys__net__Socket
					return hx_zero_56
				}
				return hx_value_55.(*sys__net__Socket)
			}(others.Get(_g1_2))
			_g1_2 = int(int32((_g1_2 + 1)))
			_g_2.Push(socket_2.handle)
		}
		return func(hx_lambda_raw_58 []any) []*hxrt.SocketHandle {
			hx_lambda_out_59 := make([]*hxrt.SocketHandle, 0, len(hx_lambda_raw_58))
			for _, hx_lambda_item_60 := range hx_lambda_raw_58 {
				hx_lambda_out_59 = append(hx_lambda_out_59, func(hx_value_61 any) *hxrt.SocketHandle {
					if hx_value_61 == nil {
						var hx_zero_62 *hxrt.SocketHandle
						return hx_zero_62
					}
					return hx_value_61.(*hxrt.SocketHandle)
				}(hx_lambda_item_60))
			}
			return hx_lambda_out_59
		}(_g_2.Values())
	}()
	result := hxrt.SocketSelect(readHandles, writeHandles, otherHandles, func() float64 {
		var hx_if_63 float64
		if timeout == nil {
			hx_if_63 = 0.0
		} else {
			hx_if_63 = timeout.(float64)
		}
		return hx_if_63
	}(), (timeout != nil))
	hx_obj_64 := map[string]any{}
	hx_obj_64["read"] = sys__net__Socket_pick(read, result.Read)
	hx_obj_64["write"] = sys__net__Socket_pick(write, result.Write)
	hx_obj_64["others"] = sys__net__Socket_pick(others, result.Others)
	return hx_obj_64
}
