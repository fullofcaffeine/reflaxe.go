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
	self.keys = func(hx_value_25 any) map[string]any {
		if hx_value_25 == nil {
			var hx_zero_26 map[string]any
			return hx_zero_26
		}
		return hx_value_25.(map[string]any)
	}(map_.keys())
	return self
}

func (self *haxe__iterators__MapKeyValueIterator) hasNext() bool {
	return func(hx_obj_27 map[string]any) func() bool {
		hx_field_28 := hx_obj_27["hasNext"]
		if hx_field_28 == nil {
			var hx_zero_29 func() bool
			return hx_zero_29
		}
		return hx_field_28.(func() bool)
	}(self.keys)()
}

func (self *haxe__iterators__MapKeyValueIterator) next() map[string]any {
	var key any = func(hx_obj_30 map[string]any) func() any {
		hx_field_31 := hx_obj_30["next"]
		if hx_field_31 == nil {
			var hx_zero_32 func() any
			return hx_zero_32
		}
		return hx_field_31.(func() any)
	}(self.keys)()
	hx_obj_33 := map[string]any{}
	hx_obj_33["key"] = key
	hx_obj_33["value"] = self.map_.getIMap(key)
	return hx_obj_33
}
