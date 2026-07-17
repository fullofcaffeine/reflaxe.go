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
	self.keys = func(hx_value_77 any) map[string]any {
		if hx_value_77 == nil {
			var hx_zero_78 map[string]any
			return hx_zero_78
		}
		return hx_value_77.(map[string]any)
	}(map_.keys())
	return self
}

func (self *haxe__iterators__MapKeyValueIterator) hasNext() bool {
	return func(hx_obj_79 map[string]any) func() bool {
		hx_field_80 := hx_obj_79["hasNext"]
		if hx_field_80 == nil {
			var hx_zero_81 func() bool
			return hx_zero_81
		}
		return hx_field_80.(func() bool)
	}(self.keys)()
}

func (self *haxe__iterators__MapKeyValueIterator) next() map[string]any {
	var key any = func(hx_obj_82 map[string]any) func() any {
		hx_field_83 := hx_obj_82["next"]
		if hx_field_83 == nil {
			var hx_zero_84 func() any
			return hx_zero_84
		}
		return hx_field_83.(func() any)
	}(self.keys)()
	hx_obj_85 := map[string]any{}
	hx_obj_85["key"] = key
	hx_obj_85["value"] = self.map_.getIMap(key)
	return hx_obj_85
}
