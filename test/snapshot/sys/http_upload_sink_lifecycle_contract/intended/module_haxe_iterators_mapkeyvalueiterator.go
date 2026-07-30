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
	self.keys = func(hx_value_432 any) map[string]any {
		if hx_value_432 == nil {
			var hx_zero_433 map[string]any
			return hx_zero_433
		}
		return hx_value_432.(map[string]any)
	}(map_.keys())
	return self
}

func (self *haxe__iterators__MapKeyValueIterator) hasNext() bool {
	return func(hx_obj_434 map[string]any) func() bool {
		hx_field_435 := hx_obj_434["hasNext"]
		if hx_field_435 == nil {
			var hx_zero_436 func() bool
			return hx_zero_436
		}
		return hx_field_435.(func() bool)
	}(self.keys)()
}

func (self *haxe__iterators__MapKeyValueIterator) next() map[string]any {
	var key any = func(hx_obj_437 map[string]any) func() any {
		hx_field_438 := hx_obj_437["next"]
		if hx_field_438 == nil {
			var hx_zero_439 func() any
			return hx_zero_439
		}
		return hx_field_438.(func() any)
	}(self.keys)()
	hx_obj_440 := map[string]any{}
	hx_obj_440["key"] = key
	hx_obj_440["value"] = self.map_.getIMap(key)
	return hx_obj_440
}
