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
	self.keys = func(hx_value_11 any) map[string]any {
		if hx_value_11 == nil {
			var hx_zero_12 map[string]any
			return hx_zero_12
		}
		return hx_value_11.(map[string]any)
	}(map_.keys())
	return self
}

func (self *haxe__iterators__MapKeyValueIterator) hasNext() bool {
	return func(hx_obj_13 map[string]any) func() bool {
		hx_field_14 := hx_obj_13["hasNext"]
		if hx_field_14 == nil {
			var hx_zero_15 func() bool
			return hx_zero_15
		}
		return hx_field_14.(func() bool)
	}(self.keys)()
}

func (self *haxe__iterators__MapKeyValueIterator) next() map[string]any {
	var key any = func(hx_obj_16 map[string]any) func() any {
		hx_field_17 := hx_obj_16["next"]
		if hx_field_17 == nil {
			var hx_zero_18 func() any
			return hx_zero_18
		}
		return hx_field_17.(func() any)
	}(self.keys)()
	hx_obj_19 := map[string]any{}
	hx_obj_19["key"] = key
	hx_obj_19["value"] = self.map_.getIMap(key)
	return hx_obj_19
}
