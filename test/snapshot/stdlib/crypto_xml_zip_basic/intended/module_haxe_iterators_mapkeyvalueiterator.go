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
	self.keys = func(hx_value_157 any) map[string]any {
		if hx_value_157 == nil {
			var hx_zero_158 map[string]any
			return hx_zero_158
		}
		return hx_value_157.(map[string]any)
	}(map_.keys())
	return self
}

func (self *haxe__iterators__MapKeyValueIterator) hasNext() bool {
	return func(hx_obj_159 map[string]any) func() bool {
		hx_field_160 := hx_obj_159["hasNext"]
		if hx_field_160 == nil {
			var hx_zero_161 func() bool
			return hx_zero_161
		}
		return hx_field_160.(func() bool)
	}(self.keys)()
}

func (self *haxe__iterators__MapKeyValueIterator) next() map[string]any {
	var key any = func(hx_obj_162 map[string]any) func() any {
		hx_field_163 := hx_obj_162["next"]
		if hx_field_163 == nil {
			var hx_zero_164 func() any
			return hx_zero_164
		}
		return hx_field_163.(func() any)
	}(self.keys)()
	hx_obj_165 := map[string]any{}
	hx_obj_165["key"] = key
	hx_obj_165["value"] = self.map_.getIMap(key)
	return hx_obj_165
}
