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
	self.keys = func(hx_value_108 any) map[string]any {
		if hx_value_108 == nil {
			var hx_zero_109 map[string]any
			return hx_zero_109
		}
		return hx_value_108.(map[string]any)
	}(map_.keys())
	return self
}

func (self *haxe__iterators__MapKeyValueIterator) hasNext() bool {
	return func(hx_obj_110 map[string]any) func() bool {
		hx_field_111 := hx_obj_110["hasNext"]
		if hx_field_111 == nil {
			var hx_zero_112 func() bool
			return hx_zero_112
		}
		return hx_field_111.(func() bool)
	}(self.keys)()
}

func (self *haxe__iterators__MapKeyValueIterator) next() map[string]any {
	var key any = func(hx_obj_113 map[string]any) func() any {
		hx_field_114 := hx_obj_113["next"]
		if hx_field_114 == nil {
			var hx_zero_115 func() any
			return hx_zero_115
		}
		return hx_field_114.(func() any)
	}(self.keys)()
	hx_obj_116 := map[string]any{}
	hx_obj_116["key"] = key
	hx_obj_116["value"] = self.map_.getIMap(key)
	return hx_obj_116
}
