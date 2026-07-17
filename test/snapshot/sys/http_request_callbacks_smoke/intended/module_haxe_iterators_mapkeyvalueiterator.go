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
	self.keys = func(hx_value_102 any) map[string]any {
		if hx_value_102 == nil {
			var hx_zero_103 map[string]any
			return hx_zero_103
		}
		return hx_value_102.(map[string]any)
	}(map_.keys())
	return self
}

func (self *haxe__iterators__MapKeyValueIterator) hasNext() bool {
	return func(hx_obj_104 map[string]any) func() bool {
		hx_field_105 := hx_obj_104["hasNext"]
		if hx_field_105 == nil {
			var hx_zero_106 func() bool
			return hx_zero_106
		}
		return hx_field_105.(func() bool)
	}(self.keys)()
}

func (self *haxe__iterators__MapKeyValueIterator) next() map[string]any {
	var key any = func(hx_obj_107 map[string]any) func() any {
		hx_field_108 := hx_obj_107["next"]
		if hx_field_108 == nil {
			var hx_zero_109 func() any
			return hx_zero_109
		}
		return hx_field_108.(func() any)
	}(self.keys)()
	hx_obj_110 := map[string]any{}
	hx_obj_110["key"] = key
	hx_obj_110["value"] = self.map_.getIMap(key)
	return hx_obj_110
}
