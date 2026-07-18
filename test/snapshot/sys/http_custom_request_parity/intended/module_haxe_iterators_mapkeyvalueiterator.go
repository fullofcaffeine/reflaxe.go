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
	self.keys = func(hx_value_176 any) map[string]any {
		if hx_value_176 == nil {
			var hx_zero_177 map[string]any
			return hx_zero_177
		}
		return hx_value_176.(map[string]any)
	}(map_.keys())
	return self
}

func (self *haxe__iterators__MapKeyValueIterator) hasNext() bool {
	return func(hx_obj_178 map[string]any) func() bool {
		hx_field_179 := hx_obj_178["hasNext"]
		if hx_field_179 == nil {
			var hx_zero_180 func() bool
			return hx_zero_180
		}
		return hx_field_179.(func() bool)
	}(self.keys)()
}

func (self *haxe__iterators__MapKeyValueIterator) next() map[string]any {
	var key any = func(hx_obj_181 map[string]any) func() any {
		hx_field_182 := hx_obj_181["next"]
		if hx_field_182 == nil {
			var hx_zero_183 func() any
			return hx_zero_183
		}
		return hx_field_182.(func() any)
	}(self.keys)()
	hx_obj_184 := map[string]any{}
	hx_obj_184["key"] = key
	hx_obj_184["value"] = self.map_.getIMap(key)
	return hx_obj_184
}
