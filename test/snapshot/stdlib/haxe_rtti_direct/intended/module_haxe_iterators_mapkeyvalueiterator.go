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
	self.keys = func(hx_value_764 any) map[string]any {
		if hx_value_764 == nil {
			var hx_zero_765 map[string]any
			return hx_zero_765
		}
		return hx_value_764.(map[string]any)
	}(map_.keys())
	return self
}

func (self *haxe__iterators__MapKeyValueIterator) hasNext() bool {
	return func(hx_obj_766 map[string]any) func() bool {
		hx_field_767 := hx_obj_766["hasNext"]
		if hx_field_767 == nil {
			var hx_zero_768 func() bool
			return hx_zero_768
		}
		return hx_field_767.(func() bool)
	}(self.keys)()
}

func (self *haxe__iterators__MapKeyValueIterator) next() map[string]any {
	var key any = func(hx_obj_769 map[string]any) func() any {
		hx_field_770 := hx_obj_769["next"]
		if hx_field_770 == nil {
			var hx_zero_771 func() any
			return hx_zero_771
		}
		return hx_field_770.(func() any)
	}(self.keys)()
	hx_obj_772 := map[string]any{}
	hx_obj_772["key"] = key
	hx_obj_772["value"] = self.map_.getIMap(key)
	return hx_obj_772
}
