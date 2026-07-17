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
	self.keys = func(hx_value_178 any) map[string]any {
		if hx_value_178 == nil {
			var hx_zero_179 map[string]any
			return hx_zero_179
		}
		return hx_value_178.(map[string]any)
	}(map_.keys())
	return self
}

func (self *haxe__iterators__MapKeyValueIterator) hasNext() bool {
	return func(hx_obj_180 map[string]any) func() bool {
		hx_field_181 := hx_obj_180["hasNext"]
		if hx_field_181 == nil {
			var hx_zero_182 func() bool
			return hx_zero_182
		}
		return hx_field_181.(func() bool)
	}(self.keys)()
}

func (self *haxe__iterators__MapKeyValueIterator) next() map[string]any {
	var key any = func(hx_obj_183 map[string]any) func() any {
		hx_field_184 := hx_obj_183["next"]
		if hx_field_184 == nil {
			var hx_zero_185 func() any
			return hx_zero_185
		}
		return hx_field_184.(func() any)
	}(self.keys)()
	hx_obj_186 := map[string]any{}
	hx_obj_186["key"] = key
	hx_obj_186["value"] = self.map_.getIMap(key)
	return hx_obj_186
}
