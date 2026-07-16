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
	self.keys = func(hx_value_62 any) map[string]any {
		if hx_value_62 == nil {
			var hx_zero_63 map[string]any
			return hx_zero_63
		}
		return hx_value_62.(map[string]any)
	}(map_.keys())
	return self
}

func (self *haxe__iterators__MapKeyValueIterator) hasNext() bool {
	return func(hx_obj_64 map[string]any) func() bool {
		hx_field_65 := hx_obj_64["hasNext"]
		if hx_field_65 == nil {
			var hx_zero_66 func() bool
			return hx_zero_66
		}
		return hx_field_65.(func() bool)
	}(self.keys)()
}

func (self *haxe__iterators__MapKeyValueIterator) next() map[string]any {
	var key any = func(hx_obj_67 map[string]any) func() any {
		hx_field_68 := hx_obj_67["next"]
		if hx_field_68 == nil {
			var hx_zero_69 func() any
			return hx_zero_69
		}
		return hx_field_68.(func() any)
	}(self.keys)()
	hx_obj_70 := map[string]any{}
	hx_obj_70["key"] = key
	hx_obj_70["value"] = self.map_.getIMap(key)
	return hx_obj_70
}
