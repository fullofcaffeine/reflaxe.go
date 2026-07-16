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
	self.keys = func(hx_value_5 any) map[string]any {
		if hx_value_5 == nil {
			var hx_zero_6 map[string]any
			return hx_zero_6
		}
		return hx_value_5.(map[string]any)
	}(map_.keys())
	return self
}

func (self *haxe__iterators__MapKeyValueIterator) hasNext() bool {
	return func(hx_obj_7 map[string]any) func() bool {
		hx_field_8 := hx_obj_7["hasNext"]
		if hx_field_8 == nil {
			var hx_zero_9 func() bool
			return hx_zero_9
		}
		return hx_field_8.(func() bool)
	}(self.keys)()
}

func (self *haxe__iterators__MapKeyValueIterator) next() map[string]any {
	var key any = func(hx_obj_10 map[string]any) func() any {
		hx_field_11 := hx_obj_10["next"]
		if hx_field_11 == nil {
			var hx_zero_12 func() any
			return hx_zero_12
		}
		return hx_field_11.(func() any)
	}(self.keys)()
	hx_obj_13 := map[string]any{}
	hx_obj_13["key"] = key
	hx_obj_13["value"] = self.map_.getIMap(key)
	return hx_obj_13
}
