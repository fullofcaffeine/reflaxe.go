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
	self.keys = func(hx_value_226 any) map[string]any {
		if hx_value_226 == nil {
			var hx_zero_227 map[string]any
			return hx_zero_227
		}
		return hx_value_226.(map[string]any)
	}(map_.keys())
	return self
}

func (self *haxe__iterators__MapKeyValueIterator) hasNext() bool {
	return func(hx_obj_228 map[string]any) func() bool {
		hx_field_229 := hx_obj_228["hasNext"]
		if hx_field_229 == nil {
			var hx_zero_230 func() bool
			return hx_zero_230
		}
		return hx_field_229.(func() bool)
	}(self.keys)()
}

func (self *haxe__iterators__MapKeyValueIterator) next() map[string]any {
	var key any = func(hx_obj_231 map[string]any) func() any {
		hx_field_232 := hx_obj_231["next"]
		if hx_field_232 == nil {
			var hx_zero_233 func() any
			return hx_zero_233
		}
		return hx_field_232.(func() any)
	}(self.keys)()
	hx_obj_234 := map[string]any{}
	hx_obj_234["key"] = key
	hx_obj_234["value"] = self.map_.getIMap(key)
	return hx_obj_234
}
