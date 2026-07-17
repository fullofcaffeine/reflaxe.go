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
	self.keys = func(hx_value_24 any) map[string]any {
		if hx_value_24 == nil {
			var hx_zero_25 map[string]any
			return hx_zero_25
		}
		return hx_value_24.(map[string]any)
	}(map_.keys())
	return self
}

func (self *haxe__iterators__MapKeyValueIterator) hasNext() bool {
	return func(hx_obj_26 map[string]any) func() bool {
		hx_field_27 := hx_obj_26["hasNext"]
		if hx_field_27 == nil {
			var hx_zero_28 func() bool
			return hx_zero_28
		}
		return hx_field_27.(func() bool)
	}(self.keys)()
}

func (self *haxe__iterators__MapKeyValueIterator) next() map[string]any {
	var key any = func(hx_obj_29 map[string]any) func() any {
		hx_field_30 := hx_obj_29["next"]
		if hx_field_30 == nil {
			var hx_zero_31 func() any
			return hx_zero_31
		}
		return hx_field_30.(func() any)
	}(self.keys)()
	hx_obj_32 := map[string]any{}
	hx_obj_32["key"] = key
	hx_obj_32["value"] = self.map_.getIMap(key)
	return hx_obj_32
}
