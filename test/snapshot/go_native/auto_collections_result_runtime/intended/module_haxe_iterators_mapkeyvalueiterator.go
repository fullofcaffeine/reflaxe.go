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
	self.keys = func(hx_value_37 any) map[string]any {
		if hx_value_37 == nil {
			var hx_zero_38 map[string]any
			return hx_zero_38
		}
		return hx_value_37.(map[string]any)
	}(map_.keys())
	return self
}

func (self *haxe__iterators__MapKeyValueIterator) hasNext() bool {
	return func(hx_obj_39 map[string]any) func() bool {
		hx_field_40 := hx_obj_39["hasNext"]
		if hx_field_40 == nil {
			var hx_zero_41 func() bool
			return hx_zero_41
		}
		return hx_field_40.(func() bool)
	}(self.keys)()
}

func (self *haxe__iterators__MapKeyValueIterator) next() map[string]any {
	var key any = func(hx_obj_42 map[string]any) func() any {
		hx_field_43 := hx_obj_42["next"]
		if hx_field_43 == nil {
			var hx_zero_44 func() any
			return hx_zero_44
		}
		return hx_field_43.(func() any)
	}(self.keys)()
	hx_obj_45 := map[string]any{}
	hx_obj_45["key"] = key
	hx_obj_45["value"] = self.map_.getIMap(key)
	return hx_obj_45
}
