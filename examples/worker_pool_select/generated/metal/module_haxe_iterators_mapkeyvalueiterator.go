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
	self.keys = func(hx_value_16 any) map[string]any {
		if hx_value_16 == nil {
			var hx_zero_17 map[string]any
			return hx_zero_17
		}
		return hx_value_16.(map[string]any)
	}(map_.keys())
	return self
}

func (self *haxe__iterators__MapKeyValueIterator) hasNext() bool {
	return func(hx_obj_18 map[string]any) func() bool {
		hx_field_19 := hx_obj_18["hasNext"]
		if hx_field_19 == nil {
			var hx_zero_20 func() bool
			return hx_zero_20
		}
		return hx_field_19.(func() bool)
	}(self.keys)()
}

func (self *haxe__iterators__MapKeyValueIterator) next() map[string]any {
	var key any = func(hx_obj_21 map[string]any) func() any {
		hx_field_22 := hx_obj_21["next"]
		if hx_field_22 == nil {
			var hx_zero_23 func() any
			return hx_zero_23
		}
		return hx_field_22.(func() any)
	}(self.keys)()
	hx_obj_24 := map[string]any{}
	hx_obj_24["key"] = key
	hx_obj_24["value"] = self.map_.getIMap(key)
	return hx_obj_24
}
