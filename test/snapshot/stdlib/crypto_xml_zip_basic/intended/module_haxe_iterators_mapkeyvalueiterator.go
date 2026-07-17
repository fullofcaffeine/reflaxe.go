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
	self.keys = func(hx_value_129 any) map[string]any {
		if hx_value_129 == nil {
			var hx_zero_130 map[string]any
			return hx_zero_130
		}
		return hx_value_129.(map[string]any)
	}(map_.keys())
	return self
}

func (self *haxe__iterators__MapKeyValueIterator) hasNext() bool {
	return func(hx_obj_131 map[string]any) func() bool {
		hx_field_132 := hx_obj_131["hasNext"]
		if hx_field_132 == nil {
			var hx_zero_133 func() bool
			return hx_zero_133
		}
		return hx_field_132.(func() bool)
	}(self.keys)()
}

func (self *haxe__iterators__MapKeyValueIterator) next() map[string]any {
	var key any = func(hx_obj_134 map[string]any) func() any {
		hx_field_135 := hx_obj_134["next"]
		if hx_field_135 == nil {
			var hx_zero_136 func() any
			return hx_zero_136
		}
		return hx_field_135.(func() any)
	}(self.keys)()
	hx_obj_137 := map[string]any{}
	hx_obj_137["key"] = key
	hx_obj_137["value"] = self.map_.getIMap(key)
	return hx_obj_137
}
