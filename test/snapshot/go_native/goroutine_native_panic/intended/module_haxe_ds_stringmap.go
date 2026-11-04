package main

import "snapshot/hxrt"

type I_haxe__ds__StringMap interface {
	set(key *string, value any)
	get(key *string) any
	exists(key *string) bool
	remove(key *string) bool
	keys() map[string]any
	iterator() map[string]any
	keyValueIterator() map[string]any
	getIMap(key any) any
	setIMap(key any, value any)
	existsIMap(key any) bool
	removeIMap(key any) bool
	copyIMap() haxe__IMap
	copy() *haxe__ds__StringMap
	toString() *string
	clear()
}

type haxe__ds__StringMap struct {
	__hx_this I_haxe__ds__StringMap
	h         *hxrt.StringMapCell
}

func (self *haxe__ds__StringMap) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_haxe__ds__StringMap() *haxe__ds__StringMap {
	self := &haxe__ds__StringMap{}
	self.__hx_this = self
	self.h = hxrt.StringMapNew()
	return self
}

func (self *haxe__ds__StringMap) set(key *string, value any) {
	hxrt.StringMapSet(self.h, key, value)
}

func (self *haxe__ds__StringMap) get(key *string) any {
	return hxrt.StringMapGet(self.h, key)
}

func (self *haxe__ds__StringMap) exists(key *string) bool {
	return hxrt.StringMapExists(self.h, key)
}

func (self *haxe__ds__StringMap) remove(key *string) bool {
	return hxrt.StringMapRemove(self.h, key)
}

func (self *haxe__ds__StringMap) keys() map[string]any {
	keys := hxrt.StringMapKeys(self.h)
	index := 0
	hx_obj_1 := map[string]any{}
	hx_obj_1["hasNext"] = hxrt.NewFunction(func() bool {
		return (index < len(keys))
	})
	hx_obj_1["next"] = hxrt.NewFunction(func() *string {
		hx_post_2 := index
		index = int(int32((index + 1)))
		return keys[hx_post_2]
	})
	return hx_obj_1
}

func (self *haxe__ds__StringMap) iterator() map[string]any {
	_gthis := self
	keys := hxrt.StringMapKeys(self.h)
	index := 0
	hx_obj_3 := map[string]any{}
	hx_obj_3["hasNext"] = hxrt.NewFunction(func() bool {
		return (index < len(keys))
	})
	hx_obj_3["next"] = hxrt.NewFunction(func() any {
		return hxrt.StringMapGet(_gthis.h, keys[func() int {
			hx_post_4 := index
			index = int(int32((index + 1)))
			return hx_post_4
		}()])
	})
	return hx_obj_3
}

func (self *haxe__ds__StringMap) keyValueIterator() map[string]any {
	_gthis := self
	keys := func(hx_value_5 any) map[string]any {
		if hx_value_5 == nil {
			var hx_zero_6 map[string]any
			return hx_zero_6
		}
		return hx_value_5.(map[string]any)
	}(self.__hx_this.keys())
	hx_obj_7 := map[string]any{}
	hx_obj_7["hasNext"] = hxrt.NewFunction(func() bool {
		return func() bool {
			var hx_callable_11 *hxrt.Function[func() bool] = func(hx_obj_8 map[string]any) *hxrt.Function[func() bool] {
				hx_field_9 := hx_obj_8["hasNext"]
				if hx_field_9 == nil {
					var hx_zero_10 *hxrt.Function[func() bool]
					return hx_zero_10
				}
				return hx_field_9.(*hxrt.Function[func() bool])
			}(keys)
			return hxrt.FunctionCall(hx_callable_11)()
		}()
	})
	hx_obj_7["next"] = hxrt.NewFunction(func() map[string]any {
		key := func() *string {
			var hx_callable_15 *hxrt.Function[func() *string] = func(hx_obj_12 map[string]any) *hxrt.Function[func() *string] {
				hx_field_13 := hx_obj_12["next"]
				if hx_field_13 == nil {
					var hx_zero_14 *hxrt.Function[func() *string]
					return hx_zero_14
				}
				return hx_field_13.(*hxrt.Function[func() *string])
			}(keys)
			return hxrt.FunctionCall(hx_callable_15)()
		}()
		hx_obj_16 := map[string]any{}
		hx_obj_16["key"] = key
		hx_obj_16["value"] = _gthis.__hx_this.get(key)
		return hx_obj_16
	})
	return hx_obj_7
}

func (self *haxe__ds__StringMap) getIMap(key any) any {
	return self.__hx_this.get(hxrt.StdString(func(hx_value_17 any) *string {
		if hx_value_17 == nil {
			var hx_zero_18 *string
			return hx_zero_18
		}
		return hx_value_17.(*string)
	}(key)))
}

func (self *haxe__ds__StringMap) setIMap(key any, value any) {
	self.__hx_this.set(hxrt.StdString(func(hx_value_19 any) *string {
		if hx_value_19 == nil {
			var hx_zero_20 *string
			return hx_zero_20
		}
		return hx_value_19.(*string)
	}(key)), value)
}

func (self *haxe__ds__StringMap) existsIMap(key any) bool {
	return func(hx_value_23 any) bool {
		if hx_value_23 == nil {
			var hx_zero_24 bool
			return hx_zero_24
		}
		return hx_value_23.(bool)
	}(self.__hx_this.exists(hxrt.StdString(func(hx_value_21 any) *string {
		if hx_value_21 == nil {
			var hx_zero_22 *string
			return hx_zero_22
		}
		return hx_value_21.(*string)
	}(key))))
}

func (self *haxe__ds__StringMap) removeIMap(key any) bool {
	return func(hx_value_27 any) bool {
		if hx_value_27 == nil {
			var hx_zero_28 bool
			return hx_zero_28
		}
		return hx_value_27.(bool)
	}(self.__hx_this.remove(hxrt.StdString(func(hx_value_25 any) *string {
		if hx_value_25 == nil {
			var hx_zero_26 *string
			return hx_zero_26
		}
		return hx_value_25.(*string)
	}(key))))
}

func (self *haxe__ds__StringMap) copyIMap() haxe__IMap {
	return func(hx_value_29 any) *haxe__ds__StringMap {
		if hx_value_29 == nil {
			var hx_zero_30 *haxe__ds__StringMap
			return hx_zero_30
		}
		return hx_value_29.(*haxe__ds__StringMap)
	}(self.__hx_this.copy())
}

func (self *haxe__ds__StringMap) copy() *haxe__ds__StringMap {
	copied := New_haxe__ds__StringMap()
	key := func(hx_value_31 any) map[string]any {
		if hx_value_31 == nil {
			var hx_zero_32 map[string]any
			return hx_zero_32
		}
		return hx_value_31.(map[string]any)
	}(self.__hx_this.keys())
	for func() bool {
		var hx_callable_36 *hxrt.Function[func() bool] = func(hx_obj_33 map[string]any) *hxrt.Function[func() bool] {
			hx_field_34 := hx_obj_33["hasNext"]
			if hx_field_34 == nil {
				var hx_zero_35 *hxrt.Function[func() bool]
				return hx_zero_35
			}
			return hx_field_34.(*hxrt.Function[func() bool])
		}(key)
		return hxrt.FunctionCall(hx_callable_36)()
	}() {
		key_1 := func() *string {
			var hx_callable_40 *hxrt.Function[func() *string] = func(hx_obj_37 map[string]any) *hxrt.Function[func() *string] {
				hx_field_38 := hx_obj_37["next"]
				if hx_field_38 == nil {
					var hx_zero_39 *hxrt.Function[func() *string]
					return hx_zero_39
				}
				return hx_field_38.(*hxrt.Function[func() *string])
			}(key)
			return hxrt.FunctionCall(hx_callable_40)()
		}()
		copied.__hx_this.set(key_1, self.__hx_this.get(key_1))
	}
	return copied
}

func (self *haxe__ds__StringMap) toString() *string {
	var out_b *string
	out_b = hxrt.StringFromLiteral("")
	out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral("["))
	iterator := func(hx_value_41 any) map[string]any {
		if hx_value_41 == nil {
			var hx_zero_42 map[string]any
			return hx_zero_42
		}
		return hx_value_41.(map[string]any)
	}(self.__hx_this.keys())
	for func() bool {
		var hx_callable_46 *hxrt.Function[func() bool] = func(hx_obj_43 map[string]any) *hxrt.Function[func() bool] {
			hx_field_44 := hx_obj_43["hasNext"]
			if hx_field_44 == nil {
				var hx_zero_45 *hxrt.Function[func() bool]
				return hx_zero_45
			}
			return hx_field_44.(*hxrt.Function[func() bool])
		}(iterator)
		return hxrt.FunctionCall(hx_callable_46)()
	}() {
		key := func() *string {
			var hx_callable_50 *hxrt.Function[func() *string] = func(hx_obj_47 map[string]any) *hxrt.Function[func() *string] {
				hx_field_48 := hx_obj_47["next"]
				if hx_field_48 == nil {
					var hx_zero_49 *hxrt.Function[func() *string]
					return hx_zero_49
				}
				return hx_field_48.(*hxrt.Function[func() *string])
			}(iterator)
			return hxrt.FunctionCall(hx_callable_50)()
		}()
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral(" => "))
		x := hxrt.StdString(self.__hx_this.get(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(x))
		if func() bool {
			var hx_callable_54 *hxrt.Function[func() bool] = func(hx_obj_51 map[string]any) *hxrt.Function[func() bool] {
				hx_field_52 := hx_obj_51["hasNext"]
				if hx_field_52 == nil {
					var hx_zero_53 *hxrt.Function[func() bool]
					return hx_zero_53
				}
				return hx_field_52.(*hxrt.Function[func() bool])
			}(iterator)
			return hxrt.FunctionCall(hx_callable_54)()
		}() {
			out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral(", "))
		}
	}
	out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral("]"))
	return out_b
}

func (self *haxe__ds__StringMap) clear() {
	hxrt.StringMapClear(self.h)
}

func (self *haxe__ds__StringMap) String() string {
	return *self.__hx_this.toString()
}
