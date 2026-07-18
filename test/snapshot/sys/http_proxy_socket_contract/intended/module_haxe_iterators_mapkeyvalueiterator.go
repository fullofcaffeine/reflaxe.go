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
	self.keys = func(hx_value_150 any) map[string]any {
		if hx_value_150 == nil {
			var hx_zero_151 map[string]any
			return hx_zero_151
		}
		return hx_value_150.(map[string]any)
	}(map_.keys())
	return self
}

func (self *haxe__iterators__MapKeyValueIterator) hasNext() bool {
	return func(hx_obj_152 map[string]any) func() bool {
		hx_field_153 := hx_obj_152["hasNext"]
		if hx_field_153 == nil {
			var hx_zero_154 func() bool
			return hx_zero_154
		}
		return hx_field_153.(func() bool)
	}(self.keys)()
}

func (self *haxe__iterators__MapKeyValueIterator) next() map[string]any {
	var key any = func(hx_obj_155 map[string]any) func() any {
		hx_field_156 := hx_obj_155["next"]
		if hx_field_156 == nil {
			var hx_zero_157 func() any
			return hx_zero_157
		}
		return hx_field_156.(func() any)
	}(self.keys)()
	hx_obj_158 := map[string]any{}
	hx_obj_158["key"] = key
	hx_obj_158["value"] = self.map_.getIMap(key)
	return hx_obj_158
}
