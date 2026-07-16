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
	self.keys = func(hx_value_73 any) map[string]any {
		if hx_value_73 == nil {
			var hx_zero_74 map[string]any
			return hx_zero_74
		}
		return hx_value_73.(map[string]any)
	}(map_.keys())
	return self
}

func (self *haxe__iterators__MapKeyValueIterator) hasNext() bool {
	return func(hx_obj_75 map[string]any) func() bool {
		hx_field_76 := hx_obj_75["hasNext"]
		if hx_field_76 == nil {
			var hx_zero_77 func() bool
			return hx_zero_77
		}
		return hx_field_76.(func() bool)
	}(self.keys)()
}

func (self *haxe__iterators__MapKeyValueIterator) next() map[string]any {
	var key any = func(hx_obj_78 map[string]any) func() any {
		hx_field_79 := hx_obj_78["next"]
		if hx_field_79 == nil {
			var hx_zero_80 func() any
			return hx_zero_80
		}
		return hx_field_79.(func() any)
	}(self.keys)()
	hx_obj_81 := map[string]any{}
	hx_obj_81["key"] = key
	hx_obj_81["value"] = self.map_.getIMap(key)
	return hx_obj_81
}
