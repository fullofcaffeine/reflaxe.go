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
	self.keys = func(hx_value_285 any) map[string]any {
		if hx_value_285 == nil {
			var hx_zero_286 map[string]any
			return hx_zero_286
		}
		return hx_value_285.(map[string]any)
	}(map_.keys())
	return self
}

func (self *haxe__iterators__MapKeyValueIterator) hasNext() bool {
	return func(hx_obj_287 map[string]any) func() bool {
		hx_field_288 := hx_obj_287["hasNext"]
		if hx_field_288 == nil {
			var hx_zero_289 func() bool
			return hx_zero_289
		}
		return hx_field_288.(func() bool)
	}(self.keys)()
}

func (self *haxe__iterators__MapKeyValueIterator) next() map[string]any {
	var key any = func(hx_obj_290 map[string]any) func() any {
		hx_field_291 := hx_obj_290["next"]
		if hx_field_291 == nil {
			var hx_zero_292 func() any
			return hx_zero_292
		}
		return hx_field_291.(func() any)
	}(self.keys)()
	hx_obj_293 := map[string]any{}
	hx_obj_293["key"] = key
	hx_obj_293["value"] = self.map_.getIMap(key)
	return hx_obj_293
}
