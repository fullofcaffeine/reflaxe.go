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
	self.keys = func(hx_value_90 any) map[string]any {
		if hx_value_90 == nil {
			var hx_zero_91 map[string]any
			return hx_zero_91
		}
		return hx_value_90.(map[string]any)
	}(map_.keys())
	return self
}

func (self *haxe__iterators__MapKeyValueIterator) hasNext() bool {
	return func(hx_obj_92 map[string]any) func() bool {
		hx_field_93 := hx_obj_92["hasNext"]
		if hx_field_93 == nil {
			var hx_zero_94 func() bool
			return hx_zero_94
		}
		return hx_field_93.(func() bool)
	}(self.keys)()
}

func (self *haxe__iterators__MapKeyValueIterator) next() map[string]any {
	var key any = func(hx_obj_95 map[string]any) func() any {
		hx_field_96 := hx_obj_95["next"]
		if hx_field_96 == nil {
			var hx_zero_97 func() any
			return hx_zero_97
		}
		return hx_field_96.(func() any)
	}(self.keys)()
	hx_obj_98 := map[string]any{}
	hx_obj_98["key"] = key
	hx_obj_98["value"] = self.map_.getIMap(key)
	return hx_obj_98
}
