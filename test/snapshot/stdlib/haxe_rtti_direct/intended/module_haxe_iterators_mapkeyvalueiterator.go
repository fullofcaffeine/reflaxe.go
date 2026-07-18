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
	self.keys = func(hx_value_769 any) map[string]any {
		if hx_value_769 == nil {
			var hx_zero_770 map[string]any
			return hx_zero_770
		}
		return hx_value_769.(map[string]any)
	}(map_.keys())
	return self
}

func (self *haxe__iterators__MapKeyValueIterator) hasNext() bool {
	return func(hx_obj_771 map[string]any) func() bool {
		hx_field_772 := hx_obj_771["hasNext"]
		if hx_field_772 == nil {
			var hx_zero_773 func() bool
			return hx_zero_773
		}
		return hx_field_772.(func() bool)
	}(self.keys)()
}

func (self *haxe__iterators__MapKeyValueIterator) next() map[string]any {
	var key any = func(hx_obj_774 map[string]any) func() any {
		hx_field_775 := hx_obj_774["next"]
		if hx_field_775 == nil {
			var hx_zero_776 func() any
			return hx_zero_776
		}
		return hx_field_775.(func() any)
	}(self.keys)()
	hx_obj_777 := map[string]any{}
	hx_obj_777["key"] = key
	hx_obj_777["value"] = self.map_.getIMap(key)
	return hx_obj_777
}
