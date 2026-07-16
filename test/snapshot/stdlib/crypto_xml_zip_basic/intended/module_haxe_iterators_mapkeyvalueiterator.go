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
	self.keys = func(hx_value_133 any) map[string]any {
		if hx_value_133 == nil {
			var hx_zero_134 map[string]any
			return hx_zero_134
		}
		return hx_value_133.(map[string]any)
	}(map_.keys())
	return self
}

func (self *haxe__iterators__MapKeyValueIterator) hasNext() bool {
	return func(hx_obj_135 map[string]any) func() bool {
		hx_field_136 := hx_obj_135["hasNext"]
		if hx_field_136 == nil {
			var hx_zero_137 func() bool
			return hx_zero_137
		}
		return hx_field_136.(func() bool)
	}(self.keys)()
}

func (self *haxe__iterators__MapKeyValueIterator) next() map[string]any {
	var key any = func(hx_obj_138 map[string]any) func() any {
		hx_field_139 := hx_obj_138["next"]
		if hx_field_139 == nil {
			var hx_zero_140 func() any
			return hx_zero_140
		}
		return hx_field_139.(func() any)
	}(self.keys)()
	hx_obj_141 := map[string]any{}
	hx_obj_141["key"] = key
	hx_obj_141["value"] = self.map_.getIMap(key)
	return hx_obj_141
}
