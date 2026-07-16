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
	self.keys = func(hx_value_95 any) map[string]any {
		if hx_value_95 == nil {
			var hx_zero_96 map[string]any
			return hx_zero_96
		}
		return hx_value_95.(map[string]any)
	}(map_.keys())
	return self
}

func (self *haxe__iterators__MapKeyValueIterator) hasNext() bool {
	return func(hx_obj_97 map[string]any) func() bool {
		hx_field_98 := hx_obj_97["hasNext"]
		if hx_field_98 == nil {
			var hx_zero_99 func() bool
			return hx_zero_99
		}
		return hx_field_98.(func() bool)
	}(self.keys)()
}

func (self *haxe__iterators__MapKeyValueIterator) next() map[string]any {
	var key any = func(hx_obj_100 map[string]any) func() any {
		hx_field_101 := hx_obj_100["next"]
		if hx_field_101 == nil {
			var hx_zero_102 func() any
			return hx_zero_102
		}
		return hx_field_101.(func() any)
	}(self.keys)()
	hx_obj_103 := map[string]any{}
	hx_obj_103["key"] = key
	hx_obj_103["value"] = self.map_.getIMap(key)
	return hx_obj_103
}
