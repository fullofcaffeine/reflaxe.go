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
	self.keys = func(hx_value_1 any) map[string]any {
		if hx_value_1 == nil {
			var hx_zero_2 map[string]any
			return hx_zero_2
		}
		return hx_value_1.(map[string]any)
	}(map_.keys())
	return self
}

func (self *haxe__iterators__MapKeyValueIterator) hasNext() bool {
	return func(hx_obj_3 map[string]any) func() bool {
		hx_field_4 := hx_obj_3["hasNext"]
		if hx_field_4 == nil {
			var hx_zero_5 func() bool
			return hx_zero_5
		}
		return hx_field_4.(func() bool)
	}(self.keys)()
}

func (self *haxe__iterators__MapKeyValueIterator) next() map[string]any {
	var key any = func(hx_obj_6 map[string]any) func() any {
		hx_field_7 := hx_obj_6["next"]
		if hx_field_7 == nil {
			var hx_zero_8 func() any
			return hx_zero_8
		}
		return hx_field_7.(func() any)
	}(self.keys)()
	hx_obj_9 := map[string]any{}
	hx_obj_9["key"] = key
	hx_obj_9["value"] = self.map_.getIMap(key)
	return hx_obj_9
}
