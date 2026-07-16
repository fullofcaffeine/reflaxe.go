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
	self.keys = func(hx_value_409 any) map[string]any {
		if hx_value_409 == nil {
			var hx_zero_410 map[string]any
			return hx_zero_410
		}
		return hx_value_409.(map[string]any)
	}(map_.keys())
	return self
}

func (self *haxe__iterators__MapKeyValueIterator) hasNext() bool {
	return func(hx_obj_411 map[string]any) func() bool {
		hx_field_412 := hx_obj_411["hasNext"]
		if hx_field_412 == nil {
			var hx_zero_413 func() bool
			return hx_zero_413
		}
		return hx_field_412.(func() bool)
	}(self.keys)()
}

func (self *haxe__iterators__MapKeyValueIterator) next() map[string]any {
	var key any = func(hx_obj_414 map[string]any) func() any {
		hx_field_415 := hx_obj_414["next"]
		if hx_field_415 == nil {
			var hx_zero_416 func() any
			return hx_zero_416
		}
		return hx_field_415.(func() any)
	}(self.keys)()
	hx_obj_417 := map[string]any{}
	hx_obj_417["key"] = key
	hx_obj_417["value"] = self.map_.getIMap(key)
	return hx_obj_417
}
