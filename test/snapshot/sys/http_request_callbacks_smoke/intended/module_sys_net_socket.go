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
	peer() hxrt.Anonymous
	host() hxrt.Anonymous
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

func (self *sys__net__Socket) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_sys__net__Socket() *sys__net__Socket {
	return hx_init_sys__net__Socket(hx_alloc_sys__net__Socket(nil))
}

func hx_alloc_sys__net__Socket(owner any) *sys__net__Socket {
	self := &sys__net__Socket{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_sys__net__Socket)
	}
	return self
}

func hx_init_sys__net__Socket(self *sys__net__Socket) *sys__net__Socket {
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
		hxrt.Throw(haxe__Exception_thrown(hxrt.StringFromLiteral("socket connect requires host")))
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
		hxrt.Throw(haxe__Exception_thrown(hxrt.StringFromLiteral("socket bind requires host")))
	}
	hxrt.SocketBindTCP(self.handle, host.__hx_this.toString(), port)
}

func (self *sys__net__Socket) accept() *sys__net__Socket {
	result := hxrt.SocketAccept(self.handle)
	if result.Status == hxrt.SocketIOBlocked {
		hxrt.Throw(haxe__Exception_thrown(haxe__io__Error_Blocked))
	}
	if (result.Status == hxrt.SocketIOEOF) || (result.Handle == nil) {
		hxrt.Throw(haxe__Exception_thrown(New_haxe__io__Eof()))
	}
	accepted := New_sys__net__Socket()
	accepted.__hx_this.replaceHandle(result.Handle)
	return accepted
}

func (self *sys__net__Socket) peer() hxrt.Anonymous {
	return sys__net__Socket_publicAddress(hxrt.SocketPeer(self.handle))
}

func (self *sys__net__Socket) host() hxrt.Anonymous {
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
	self.input = func(hx_value *sys__net__SocketInput) *haxe__io__Input {
		if hx_value == nil {
			return nil
		}
		return hx_value.haxe__io__Input
	}(New_sys__net__SocketInput(next))
	self.output = func(hx_value *sys__net__SocketOutput) *haxe__io__Output {
		if hx_value == nil {
			return nil
		}
		return hx_value.haxe__io__Output
	}(New_sys__net__SocketOutput(next))
}

func sys__net__Socket_pick(source *hxrt.Array, indexes []int) *hxrt.Array {
	selected := hxrt.NewArray()
	_g := 0
	_g1 := len(indexes)
	for _g < _g1 {
		hx_post_1 := _g
		_g = int(int32((_g + 1)))
		index := hx_post_1
		sourceIndex := indexes[index]
		if (sourceIndex >= 0) && (sourceIndex < source.Len()) {
			selected.Push(source.Get(sourceIndex))
		}
	}
	return selected
}

func sys__net__Socket_publicAddress(address *hxrt.SocketAddress) hxrt.Anonymous {
	if address == nil {
		hx_obj_3 := map[string]any{}
		hx_obj_3["host"] = sys__net__Host_fromIPv4(0)
		hx_obj_3["port"] = 0
		return hxrt.AnonymousFromMap(hx_obj_3)
	}
	hx_obj_4 := map[string]any{}
	hx_obj_4["host"] = sys__net__Host_fromIPv4(address.Host)
	hx_obj_4["port"] = address.Port
	return hxrt.AnonymousFromMap(hx_obj_4)
}

func sys__net__Socket_select_(read *hxrt.Array, write *hxrt.Array, others *hxrt.Array, timeout any) hxrt.Anonymous {
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
			socket := hx_recover_class_0(read.Get(_g1))
			_g1 = int(int32((_g1 + 1)))
			_g.Push(socket.handle)
		}
		return func(hx_lambda_raw_6 []any) []*hxrt.SocketHandle {
			hx_lambda_out_7 := make([]*hxrt.SocketHandle, 0, len(hx_lambda_raw_6))
			for _, hx_lambda_item_8 := range hx_lambda_raw_6 {
				hx_lambda_out_7 = append(hx_lambda_out_7, func(hx_value_9 any) *hxrt.SocketHandle {
					if hx_value_9 == nil {
						var hx_zero_10 *hxrt.SocketHandle
						return hx_zero_10
					}
					return hx_value_9.(*hxrt.SocketHandle)
				}(hx_lambda_item_8))
			}
			return hx_lambda_out_7
		}(_g.Values())
	}()
	writeHandles := func() []*hxrt.SocketHandle {
		_g_1 := hxrt.NewArray()
		_g1_1 := 0
		for _g1_1 < write.Len() {
			socket_1 := hx_recover_class_0(write.Get(_g1_1))
			_g1_1 = int(int32((_g1_1 + 1)))
			_g_1.Push(socket_1.handle)
		}
		return func(hx_lambda_raw_12 []any) []*hxrt.SocketHandle {
			hx_lambda_out_13 := make([]*hxrt.SocketHandle, 0, len(hx_lambda_raw_12))
			for _, hx_lambda_item_14 := range hx_lambda_raw_12 {
				hx_lambda_out_13 = append(hx_lambda_out_13, func(hx_value_15 any) *hxrt.SocketHandle {
					if hx_value_15 == nil {
						var hx_zero_16 *hxrt.SocketHandle
						return hx_zero_16
					}
					return hx_value_15.(*hxrt.SocketHandle)
				}(hx_lambda_item_14))
			}
			return hx_lambda_out_13
		}(_g_1.Values())
	}()
	otherHandles := func() []*hxrt.SocketHandle {
		_g_2 := hxrt.NewArray()
		_g1_2 := 0
		for _g1_2 < others.Len() {
			socket_2 := hx_recover_class_0(others.Get(_g1_2))
			_g1_2 = int(int32((_g1_2 + 1)))
			_g_2.Push(socket_2.handle)
		}
		return func(hx_lambda_raw_18 []any) []*hxrt.SocketHandle {
			hx_lambda_out_19 := make([]*hxrt.SocketHandle, 0, len(hx_lambda_raw_18))
			for _, hx_lambda_item_20 := range hx_lambda_raw_18 {
				hx_lambda_out_19 = append(hx_lambda_out_19, func(hx_value_21 any) *hxrt.SocketHandle {
					if hx_value_21 == nil {
						var hx_zero_22 *hxrt.SocketHandle
						return hx_zero_22
					}
					return hx_value_21.(*hxrt.SocketHandle)
				}(hx_lambda_item_20))
			}
			return hx_lambda_out_19
		}(_g_2.Values())
	}()
	result := hxrt.SocketSelect(readHandles, writeHandles, otherHandles, func() float64 {
		var hx_if_23 float64
		if timeout == nil {
			hx_if_23 = 0.0
		} else {
			hx_if_23 = timeout.(float64)
		}
		return hx_if_23
	}(), (timeout != nil))
	hx_obj_24 := map[string]any{}
	hx_obj_24["read"] = sys__net__Socket_pick(read, result.Read)
	hx_obj_24["write"] = sys__net__Socket_pick(write, result.Write)
	hx_obj_24["others"] = sys__net__Socket_pick(others, result.Others)
	return hxrt.AnonymousFromMap(hx_obj_24)
}
