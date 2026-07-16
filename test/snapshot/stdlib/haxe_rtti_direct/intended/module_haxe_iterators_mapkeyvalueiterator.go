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
	self.keys = func(hx_value_729 any) map[string]any {
		if hx_value_729 == nil {
			var hx_zero_730 map[string]any
			return hx_zero_730
		}
		return hx_value_729.(map[string]any)
	}(map_.keys())
	return self
}

func (self *haxe__iterators__MapKeyValueIterator) hasNext() bool {
	return func(hx_obj_731 map[string]any) func() bool {
		hx_field_732 := hx_obj_731["hasNext"]
		if hx_field_732 == nil {
			var hx_zero_733 func() bool
			return hx_zero_733
		}
		return hx_field_732.(func() bool)
	}(self.keys)()
}

func (self *haxe__iterators__MapKeyValueIterator) next() map[string]any {
	var key any = func(hx_obj_734 map[string]any) func() any {
		hx_field_735 := hx_obj_734["next"]
		if hx_field_735 == nil {
			var hx_zero_736 func() any
			return hx_zero_736
		}
		return hx_field_735.(func() any)
	}(self.keys)()
	hx_obj_737 := map[string]any{}
	hx_obj_737["key"] = key
	hx_obj_737["value"] = self.map_.getIMap(key)
	return hx_obj_737
}
