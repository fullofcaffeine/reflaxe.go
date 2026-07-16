package main

import "snapshot/hxrt"

type I_Key interface {
	hashCode() int
}

type Key struct {
	__hx_this I_Key
	id        int
}

func New_Key(id int) *Key {
	self := &Key{}
	self.__hx_this = self
	self.id = id
	return self
}

func (self *Key) hashCode() int {
	return self.id
}

func main() {
	var map_valuesByHash *haxe__ds__IntMap
	var map_keysByHash *haxe__ds__IntMap
	map_keysByHash = New_haxe__ds__IntMap()
	map_valuesByHash = New_haxe__ds__IntMap()
	key3 := New_Key(3)
	key4 := New_Key(4)
	hash := key3.hashCode()
	map_keysByHash.set(hash, key3)
	map_valuesByHash.set(hash, hxrt.StringFromLiteral("three"))
	hash_1 := key4.hashCode()
	map_keysByHash.set(hash_1, key4)
	map_valuesByHash.set(hash_1, hxrt.StringFromLiteral("four"))
	var v any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("three="), func(hx_value_1 any) *string {
		if hx_value_1 == nil {
			var hx_zero_2 *string
			return hx_zero_2
		}
		return hx_value_1.(*string)
	}(map_valuesByHash.get(key3.hashCode()))))
	hxrt.Println(v)
	var v_1 any = any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("four="), func(hx_value_3 any) *string {
		if hx_value_3 == nil {
			var hx_zero_4 *string
			return hx_zero_4
		}
		return hx_value_3.(*string)
	}(map_valuesByHash.get(key4.hashCode()))))
	hxrt.Println(v_1)
}
