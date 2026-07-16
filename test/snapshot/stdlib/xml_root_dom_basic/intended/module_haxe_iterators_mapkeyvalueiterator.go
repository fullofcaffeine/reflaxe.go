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
	self.keys = func(hx_value_100 any) map[string]any {
		if hx_value_100 == nil {
			var hx_zero_101 map[string]any
			return hx_zero_101
		}
		return hx_value_100.(map[string]any)
	}(map_.keys())
	return self
}

func (self *haxe__iterators__MapKeyValueIterator) hasNext() bool {
	return func(hx_obj_102 map[string]any) func() bool {
		hx_field_103 := hx_obj_102["hasNext"]
		if hx_field_103 == nil {
			var hx_zero_104 func() bool
			return hx_zero_104
		}
		return hx_field_103.(func() bool)
	}(self.keys)()
}

func (self *haxe__iterators__MapKeyValueIterator) next() map[string]any {
	var key any = func(hx_obj_105 map[string]any) func() any {
		hx_field_106 := hx_obj_105["next"]
		if hx_field_106 == nil {
			var hx_zero_107 func() any
			return hx_zero_107
		}
		return hx_field_106.(func() any)
	}(self.keys)()
	hx_obj_108 := map[string]any{}
	hx_obj_108["key"] = key
	hx_obj_108["value"] = self.map_.getIMap(key)
	return hx_obj_108
}
