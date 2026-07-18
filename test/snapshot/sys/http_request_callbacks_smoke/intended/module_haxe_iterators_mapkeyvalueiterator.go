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
	self.keys = func(hx_value_112 any) map[string]any {
		if hx_value_112 == nil {
			var hx_zero_113 map[string]any
			return hx_zero_113
		}
		return hx_value_112.(map[string]any)
	}(map_.keys())
	return self
}

func (self *haxe__iterators__MapKeyValueIterator) hasNext() bool {
	return func(hx_obj_114 map[string]any) func() bool {
		hx_field_115 := hx_obj_114["hasNext"]
		if hx_field_115 == nil {
			var hx_zero_116 func() bool
			return hx_zero_116
		}
		return hx_field_115.(func() bool)
	}(self.keys)()
}

func (self *haxe__iterators__MapKeyValueIterator) next() map[string]any {
	var key any = func(hx_obj_117 map[string]any) func() any {
		hx_field_118 := hx_obj_117["next"]
		if hx_field_118 == nil {
			var hx_zero_119 func() any
			return hx_zero_119
		}
		return hx_field_118.(func() any)
	}(self.keys)()
	hx_obj_120 := map[string]any{}
	hx_obj_120["key"] = key
	hx_obj_120["value"] = self.map_.getIMap(key)
	return hx_obj_120
}
