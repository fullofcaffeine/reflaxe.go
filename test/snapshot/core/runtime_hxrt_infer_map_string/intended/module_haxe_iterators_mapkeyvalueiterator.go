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
	self.keys = func(hx_value_3 any) map[string]any {
		if hx_value_3 == nil {
			var hx_zero_4 map[string]any
			return hx_zero_4
		}
		return hx_value_3.(map[string]any)
	}(map_.keys())
	return self
}

func (self *haxe__iterators__MapKeyValueIterator) hasNext() bool {
	return func(hx_obj_5 map[string]any) func() bool {
		hx_field_6 := hx_obj_5["hasNext"]
		if hx_field_6 == nil {
			var hx_zero_7 func() bool
			return hx_zero_7
		}
		return hx_field_6.(func() bool)
	}(self.keys)()
}

func (self *haxe__iterators__MapKeyValueIterator) next() map[string]any {
	var key any = func(hx_obj_8 map[string]any) func() any {
		hx_field_9 := hx_obj_8["next"]
		if hx_field_9 == nil {
			var hx_zero_10 func() any
			return hx_zero_10
		}
		return hx_field_9.(func() any)
	}(self.keys)()
	hx_obj_11 := map[string]any{}
	hx_obj_11["key"] = key
	hx_obj_11["value"] = self.map_.getIMap(key)
	return hx_obj_11
}
