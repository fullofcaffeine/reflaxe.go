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
	self.keys = func(hx_value_148 any) map[string]any {
		if hx_value_148 == nil {
			var hx_zero_149 map[string]any
			return hx_zero_149
		}
		return hx_value_148.(map[string]any)
	}(map_.keys())
	return self
}

func (self *haxe__iterators__MapKeyValueIterator) hasNext() bool {
	return func(hx_obj_150 map[string]any) func() bool {
		hx_field_151 := hx_obj_150["hasNext"]
		if hx_field_151 == nil {
			var hx_zero_152 func() bool
			return hx_zero_152
		}
		return hx_field_151.(func() bool)
	}(self.keys)()
}

func (self *haxe__iterators__MapKeyValueIterator) next() map[string]any {
	var key any = func(hx_obj_153 map[string]any) func() any {
		hx_field_154 := hx_obj_153["next"]
		if hx_field_154 == nil {
			var hx_zero_155 func() any
			return hx_zero_155
		}
		return hx_field_154.(func() any)
	}(self.keys)()
	hx_obj_156 := map[string]any{}
	hx_obj_156["key"] = key
	hx_obj_156["value"] = self.map_.getIMap(key)
	return hx_obj_156
}
