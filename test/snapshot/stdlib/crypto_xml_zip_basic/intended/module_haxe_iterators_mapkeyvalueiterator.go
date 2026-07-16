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
	self.keys = func(hx_value_158 any) map[string]any {
		if hx_value_158 == nil {
			var hx_zero_159 map[string]any
			return hx_zero_159
		}
		return hx_value_158.(map[string]any)
	}(map_.keys())
	return self
}

func (self *haxe__iterators__MapKeyValueIterator) hasNext() bool {
	return func(hx_obj_160 map[string]any) func() bool {
		hx_field_161 := hx_obj_160["hasNext"]
		if hx_field_161 == nil {
			var hx_zero_162 func() bool
			return hx_zero_162
		}
		return hx_field_161.(func() bool)
	}(self.keys)()
}

func (self *haxe__iterators__MapKeyValueIterator) next() map[string]any {
	var key any = func(hx_obj_163 map[string]any) func() any {
		hx_field_164 := hx_obj_163["next"]
		if hx_field_164 == nil {
			var hx_zero_165 func() any
			return hx_zero_165
		}
		return hx_field_164.(func() any)
	}(self.keys)()
	hx_obj_166 := map[string]any{}
	hx_obj_166["key"] = key
	hx_obj_166["value"] = self.map_.getIMap(key)
	return hx_obj_166
}
