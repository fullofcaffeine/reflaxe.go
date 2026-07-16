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
	self.keys = func(hx_value_726 any) map[string]any {
		if hx_value_726 == nil {
			var hx_zero_727 map[string]any
			return hx_zero_727
		}
		return hx_value_726.(map[string]any)
	}(map_.keys())
	return self
}

func (self *haxe__iterators__MapKeyValueIterator) hasNext() bool {
	return func(hx_obj_728 map[string]any) func() bool {
		hx_field_729 := hx_obj_728["hasNext"]
		if hx_field_729 == nil {
			var hx_zero_730 func() bool
			return hx_zero_730
		}
		return hx_field_729.(func() bool)
	}(self.keys)()
}

func (self *haxe__iterators__MapKeyValueIterator) next() map[string]any {
	var key any = func(hx_obj_731 map[string]any) func() any {
		hx_field_732 := hx_obj_731["next"]
		if hx_field_732 == nil {
			var hx_zero_733 func() any
			return hx_zero_733
		}
		return hx_field_732.(func() any)
	}(self.keys)()
	hx_obj_734 := map[string]any{}
	hx_obj_734["key"] = key
	hx_obj_734["value"] = self.map_.getIMap(key)
	return hx_obj_734
}
