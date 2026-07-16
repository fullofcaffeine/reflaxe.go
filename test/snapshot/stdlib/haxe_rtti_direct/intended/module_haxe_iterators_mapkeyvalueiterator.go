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
	self.keys = func(hx_value_739 any) map[string]any {
		if hx_value_739 == nil {
			var hx_zero_740 map[string]any
			return hx_zero_740
		}
		return hx_value_739.(map[string]any)
	}(map_.keys())
	return self
}

func (self *haxe__iterators__MapKeyValueIterator) hasNext() bool {
	return func(hx_obj_741 map[string]any) func() bool {
		hx_field_742 := hx_obj_741["hasNext"]
		if hx_field_742 == nil {
			var hx_zero_743 func() bool
			return hx_zero_743
		}
		return hx_field_742.(func() bool)
	}(self.keys)()
}

func (self *haxe__iterators__MapKeyValueIterator) next() map[string]any {
	var key any = func(hx_obj_744 map[string]any) func() any {
		hx_field_745 := hx_obj_744["next"]
		if hx_field_745 == nil {
			var hx_zero_746 func() any
			return hx_zero_746
		}
		return hx_field_745.(func() any)
	}(self.keys)()
	hx_obj_747 := map[string]any{}
	hx_obj_747["key"] = key
	hx_obj_747["value"] = self.map_.getIMap(key)
	return hx_obj_747
}
