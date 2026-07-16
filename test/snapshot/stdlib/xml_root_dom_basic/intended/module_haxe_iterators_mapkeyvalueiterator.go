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
	self.keys = func(hx_value_110 any) map[string]any {
		if hx_value_110 == nil {
			var hx_zero_111 map[string]any
			return hx_zero_111
		}
		return hx_value_110.(map[string]any)
	}(map_.keys())
	return self
}

func (self *haxe__iterators__MapKeyValueIterator) hasNext() bool {
	return func(hx_obj_112 map[string]any) func() bool {
		hx_field_113 := hx_obj_112["hasNext"]
		if hx_field_113 == nil {
			var hx_zero_114 func() bool
			return hx_zero_114
		}
		return hx_field_113.(func() bool)
	}(self.keys)()
}

func (self *haxe__iterators__MapKeyValueIterator) next() map[string]any {
	var key any = func(hx_obj_115 map[string]any) func() any {
		hx_field_116 := hx_obj_115["next"]
		if hx_field_116 == nil {
			var hx_zero_117 func() any
			return hx_zero_117
		}
		return hx_field_116.(func() any)
	}(self.keys)()
	hx_obj_118 := map[string]any{}
	hx_obj_118["key"] = key
	hx_obj_118["value"] = self.map_.getIMap(key)
	return hx_obj_118
}
