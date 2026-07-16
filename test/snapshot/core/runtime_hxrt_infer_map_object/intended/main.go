package main

import "snapshot/hxrt"

type I_CollectionFeatureKey interface {
}

type CollectionFeatureKey struct {
	__hx_this I_CollectionFeatureKey
}

func New_CollectionFeatureKey() *CollectionFeatureKey {
	self := &CollectionFeatureKey{}
	self.__hx_this = self
	return self
}

func main() {
	key := New_CollectionFeatureKey()
	values := New_haxe__ds__ObjectMap()
	values.set(key, hxrt.StringFromLiteral("one"))
	func(hx_value_1 any) bool {
		if hx_value_1 == nil {
			var hx_zero_2 bool
			return hx_zero_2
		}
		return hx_value_1.(bool)
	}(values.exists(key))
}
