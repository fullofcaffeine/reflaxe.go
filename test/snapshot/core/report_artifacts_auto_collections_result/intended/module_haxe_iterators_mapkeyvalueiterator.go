package main

type I_haxe__iterators__MapKeyValueIterator interface {
	hasNext() bool
	next() map[string]any
}

type haxe__iterators__MapKeyValueIterator struct {
	__hx_this I_haxe__iterators__MapKeyValueIterator
	map_      haxe__IMap
	keys      map[string]any
}

func New_haxe__iterators__MapKeyValueIterator(map_ haxe__IMap) *haxe__iterators__MapKeyValueIterator {
	self := &haxe__iterators__MapKeyValueIterator{}
	self.__hx_this = self
	self.map_ = map_
	self.keys = func(hx_value_9 any) map[string]any {
		if hx_value_9 == nil {
			var hx_zero_10 map[string]any
			return hx_zero_10
		}
		return hx_value_9.(map[string]any)
	}(map_.keys())
	return self
}

func (self *haxe__iterators__MapKeyValueIterator) hasNext() bool {
	return func(hx_obj_11 map[string]any) func() bool {
		hx_field_12 := hx_obj_11["hasNext"]
		if hx_field_12 == nil {
			var hx_zero_13 func() bool
			return hx_zero_13
		}
		return hx_field_12.(func() bool)
	}(self.keys)()
}

func (self *haxe__iterators__MapKeyValueIterator) next() map[string]any {
	var key any = func(hx_obj_14 map[string]any) func() any {
		hx_field_15 := hx_obj_14["next"]
		if hx_field_15 == nil {
			var hx_zero_16 func() any
			return hx_zero_16
		}
		return hx_field_15.(func() any)
	}(self.keys)()
	hx_obj_17 := map[string]any{}
	hx_obj_17["key"] = key
	hx_obj_17["value"] = self.map_.getIMap(key)
	return hx_obj_17
}
