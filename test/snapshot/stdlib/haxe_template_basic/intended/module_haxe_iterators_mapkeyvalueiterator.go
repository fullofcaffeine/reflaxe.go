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
	self.keys = func(hx_value_436 any) map[string]any {
		if hx_value_436 == nil {
			var hx_zero_437 map[string]any
			return hx_zero_437
		}
		return hx_value_436.(map[string]any)
	}(map_.keys())
	return self
}

func (self *haxe__iterators__MapKeyValueIterator) hasNext() bool {
	return func(hx_obj_438 map[string]any) func() bool {
		hx_field_439 := hx_obj_438["hasNext"]
		if hx_field_439 == nil {
			var hx_zero_440 func() bool
			return hx_zero_440
		}
		return hx_field_439.(func() bool)
	}(self.keys)()
}

func (self *haxe__iterators__MapKeyValueIterator) next() map[string]any {
	var key any = func(hx_obj_441 map[string]any) func() any {
		hx_field_442 := hx_obj_441["next"]
		if hx_field_442 == nil {
			var hx_zero_443 func() any
			return hx_zero_443
		}
		return hx_field_442.(func() any)
	}(self.keys)()
	hx_obj_444 := map[string]any{}
	hx_obj_444["key"] = key
	hx_obj_444["value"] = self.map_.getIMap(key)
	return hx_obj_444
}
