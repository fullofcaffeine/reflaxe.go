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
	self.keys = func(hx_value_710 any) map[string]any {
		if hx_value_710 == nil {
			var hx_zero_711 map[string]any
			return hx_zero_711
		}
		return hx_value_710.(map[string]any)
	}(map_.keys())
	return self
}

func (self *haxe__iterators__MapKeyValueIterator) hasNext() bool {
	return func(hx_obj_712 map[string]any) func() bool {
		hx_field_713 := hx_obj_712["hasNext"]
		if hx_field_713 == nil {
			var hx_zero_714 func() bool
			return hx_zero_714
		}
		return hx_field_713.(func() bool)
	}(self.keys)()
}

func (self *haxe__iterators__MapKeyValueIterator) next() map[string]any {
	var key any = func(hx_obj_715 map[string]any) func() any {
		hx_field_716 := hx_obj_715["next"]
		if hx_field_716 == nil {
			var hx_zero_717 func() any
			return hx_zero_717
		}
		return hx_field_716.(func() any)
	}(self.keys)()
	hx_obj_718 := map[string]any{}
	hx_obj_718["key"] = key
	hx_obj_718["value"] = self.map_.getIMap(key)
	return hx_obj_718
}
