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
	self.keys = func(hx_value_117 any) map[string]any {
		if hx_value_117 == nil {
			var hx_zero_118 map[string]any
			return hx_zero_118
		}
		return hx_value_117.(map[string]any)
	}(map_.keys())
	return self
}

func (self *haxe__iterators__MapKeyValueIterator) hasNext() bool {
	return func(hx_obj_119 map[string]any) func() bool {
		hx_field_120 := hx_obj_119["hasNext"]
		if hx_field_120 == nil {
			var hx_zero_121 func() bool
			return hx_zero_121
		}
		return hx_field_120.(func() bool)
	}(self.keys)()
}

func (self *haxe__iterators__MapKeyValueIterator) next() map[string]any {
	var key any = func(hx_obj_122 map[string]any) func() any {
		hx_field_123 := hx_obj_122["next"]
		if hx_field_123 == nil {
			var hx_zero_124 func() any
			return hx_zero_124
		}
		return hx_field_123.(func() any)
	}(self.keys)()
	hx_obj_125 := map[string]any{}
	hx_obj_125["key"] = key
	hx_obj_125["value"] = self.map_.getIMap(key)
	return hx_obj_125
}
