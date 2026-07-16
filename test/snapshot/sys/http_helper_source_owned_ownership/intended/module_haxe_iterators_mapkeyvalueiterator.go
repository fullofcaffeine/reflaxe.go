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
	self.keys = func(hx_value_60 any) map[string]any {
		if hx_value_60 == nil {
			var hx_zero_61 map[string]any
			return hx_zero_61
		}
		return hx_value_60.(map[string]any)
	}(map_.keys())
	return self
}

func (self *haxe__iterators__MapKeyValueIterator) hasNext() bool {
	return func(hx_obj_62 map[string]any) func() bool {
		hx_field_63 := hx_obj_62["hasNext"]
		if hx_field_63 == nil {
			var hx_zero_64 func() bool
			return hx_zero_64
		}
		return hx_field_63.(func() bool)
	}(self.keys)()
}

func (self *haxe__iterators__MapKeyValueIterator) next() map[string]any {
	var key any = func(hx_obj_65 map[string]any) func() any {
		hx_field_66 := hx_obj_65["next"]
		if hx_field_66 == nil {
			var hx_zero_67 func() any
			return hx_zero_67
		}
		return hx_field_66.(func() any)
	}(self.keys)()
	hx_obj_68 := map[string]any{}
	hx_obj_68["key"] = key
	hx_obj_68["value"] = self.map_.getIMap(key)
	return hx_obj_68
}
