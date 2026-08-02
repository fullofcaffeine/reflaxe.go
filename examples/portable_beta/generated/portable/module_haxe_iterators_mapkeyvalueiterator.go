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
	self.keys = func(hx_value_47 any) map[string]any {
		if hx_value_47 == nil {
			var hx_zero_48 map[string]any
			return hx_zero_48
		}
		return hx_value_47.(map[string]any)
	}(map_.keys())
	return self
}

func (self *haxe__iterators__MapKeyValueIterator) hasNext() bool {
	return func(hx_obj_49 map[string]any) func() bool {
		hx_field_50 := hx_obj_49["hasNext"]
		if hx_field_50 == nil {
			var hx_zero_51 func() bool
			return hx_zero_51
		}
		return hx_field_50.(func() bool)
	}(self.keys)()
}

func (self *haxe__iterators__MapKeyValueIterator) next() map[string]any {
	var key any = func(hx_obj_52 map[string]any) func() any {
		hx_field_53 := hx_obj_52["next"]
		if hx_field_53 == nil {
			var hx_zero_54 func() any
			return hx_zero_54
		}
		return hx_field_53.(func() any)
	}(self.keys)()
	hx_obj_55 := map[string]any{}
	hx_obj_55["key"] = key
	hx_obj_55["value"] = self.map_.getIMap(key)
	return hx_obj_55
}
