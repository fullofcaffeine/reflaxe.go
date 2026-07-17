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
	self.keys = func(hx_value_96 any) map[string]any {
		if hx_value_96 == nil {
			var hx_zero_97 map[string]any
			return hx_zero_97
		}
		return hx_value_96.(map[string]any)
	}(map_.keys())
	return self
}

func (self *haxe__iterators__MapKeyValueIterator) hasNext() bool {
	return func(hx_obj_98 map[string]any) func() bool {
		hx_field_99 := hx_obj_98["hasNext"]
		if hx_field_99 == nil {
			var hx_zero_100 func() bool
			return hx_zero_100
		}
		return hx_field_99.(func() bool)
	}(self.keys)()
}

func (self *haxe__iterators__MapKeyValueIterator) next() map[string]any {
	var key any = func(hx_obj_101 map[string]any) func() any {
		hx_field_102 := hx_obj_101["next"]
		if hx_field_102 == nil {
			var hx_zero_103 func() any
			return hx_zero_103
		}
		return hx_field_102.(func() any)
	}(self.keys)()
	hx_obj_104 := map[string]any{}
	hx_obj_104["key"] = key
	hx_obj_104["value"] = self.map_.getIMap(key)
	return hx_obj_104
}
