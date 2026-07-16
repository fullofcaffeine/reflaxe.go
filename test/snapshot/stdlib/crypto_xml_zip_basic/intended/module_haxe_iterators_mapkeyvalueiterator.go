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
	self.keys = func(hx_value_143 any) map[string]any {
		if hx_value_143 == nil {
			var hx_zero_144 map[string]any
			return hx_zero_144
		}
		return hx_value_143.(map[string]any)
	}(map_.keys())
	return self
}

func (self *haxe__iterators__MapKeyValueIterator) hasNext() bool {
	return func(hx_obj_145 map[string]any) func() bool {
		hx_field_146 := hx_obj_145["hasNext"]
		if hx_field_146 == nil {
			var hx_zero_147 func() bool
			return hx_zero_147
		}
		return hx_field_146.(func() bool)
	}(self.keys)()
}

func (self *haxe__iterators__MapKeyValueIterator) next() map[string]any {
	var key any = func(hx_obj_148 map[string]any) func() any {
		hx_field_149 := hx_obj_148["next"]
		if hx_field_149 == nil {
			var hx_zero_150 func() any
			return hx_zero_150
		}
		return hx_field_149.(func() any)
	}(self.keys)()
	hx_obj_151 := map[string]any{}
	hx_obj_151["key"] = key
	hx_obj_151["value"] = self.map_.getIMap(key)
	return hx_obj_151
}
