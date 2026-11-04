package main

import "snapshot/hxrt"

type I_haxe__ds__ObjectMap interface {
	set(key any, value any)
	get(key any) any
	exists(key any) bool
	remove(key any) bool
	keys() hxrt.Anonymous
	iterator() hxrt.Anonymous
	keyValueIterator() hxrt.Anonymous
	getIMap(key any) any
	setIMap(key any, value any)
	existsIMap(key any) bool
	removeIMap(key any) bool
	copyIMap() haxe__IMap
	copy() *haxe__ds__ObjectMap
	toString() *string
	clear()
}

type haxe__ds__ObjectMap struct {
	__hx_this I_haxe__ds__ObjectMap
	h         *hxrt.ObjectMapCell
}

func (self *haxe__ds__ObjectMap) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_haxe__ds__ObjectMap() *haxe__ds__ObjectMap {
	self := &haxe__ds__ObjectMap{}
	self.__hx_this = self
	self.h = hxrt.ObjectMapNew()
	return self
}

func (self *haxe__ds__ObjectMap) set(key any, value any) {
	hxrt.ObjectMapSet(self.h, key, value)
}

func (self *haxe__ds__ObjectMap) get(key any) any {
	return hxrt.ObjectMapGet(self.h, key)
}

func (self *haxe__ds__ObjectMap) exists(key any) bool {
	return hxrt.ObjectMapExists(self.h, key)
}

func (self *haxe__ds__ObjectMap) remove(key any) bool {
	return hxrt.ObjectMapRemove(self.h, key)
}

func (self *haxe__ds__ObjectMap) keys() hxrt.Anonymous {
	keys := hxrt.ObjectMapKeys(self.h)
	index := 0
	hx_obj_1 := map[string]any{}
	hx_obj_1["hasNext"] = hxrt.NewFunction(func() bool {
		return (index < len(keys))
	})
	hx_obj_1["next"] = hxrt.NewFunction(func() any {
		return keys[func() int {
			hx_post_2 := index
			index = int(int32((index + 1)))
			return hx_post_2
		}()]
	})
	return hxrt.AnonymousFromMap(hx_obj_1)
}

func (self *haxe__ds__ObjectMap) iterator() hxrt.Anonymous {
	_gthis := self
	keys := hxrt.ObjectMapKeys(self.h)
	index := 0
	hx_obj_3 := map[string]any{}
	hx_obj_3["hasNext"] = hxrt.NewFunction(func() bool {
		return (index < len(keys))
	})
	hx_obj_3["next"] = hxrt.NewFunction(func() any {
		return hxrt.ObjectMapGet(_gthis.h, keys[func() int {
			hx_post_4 := index
			index = int(int32((index + 1)))
			return hx_post_4
		}()])
	})
	return hxrt.AnonymousFromMap(hx_obj_3)
}

func (self *haxe__ds__ObjectMap) keyValueIterator() hxrt.Anonymous {
	_gthis := self
	keys := self.__hx_this.keys()
	hx_obj_5 := map[string]any{}
	hx_obj_5["hasNext"] = hxrt.NewFunction(func() bool {
		return func() bool {
			var hx_callable_9 *hxrt.Function[func() bool] = func(hx_obj_6 hxrt.Anonymous) *hxrt.Function[func() bool] {
				hx_field_7 := hxrt.AnonymousGet(hx_obj_6, "hasNext")
				if hx_field_7 == nil {
					var hx_zero_8 *hxrt.Function[func() bool]
					return hx_zero_8
				}
				return hx_field_7.(*hxrt.Function[func() bool])
			}(keys)
			return hxrt.FunctionCall(hx_callable_9)()
		}()
	})
	hx_obj_5["next"] = hxrt.NewFunction(func() hxrt.Anonymous {
		var key any = func() any {
			var hx_callable_13 *hxrt.Function[func() any] = func(hx_obj_10 hxrt.Anonymous) *hxrt.Function[func() any] {
				hx_field_11 := hxrt.AnonymousGet(hx_obj_10, "next")
				if hx_field_11 == nil {
					var hx_zero_12 *hxrt.Function[func() any]
					return hx_zero_12
				}
				return hx_field_11.(*hxrt.Function[func() any])
			}(keys)
			return hxrt.FunctionCall(hx_callable_13)()
		}()
		hx_obj_14 := map[string]any{}
		hx_obj_14["key"] = key
		hx_obj_14["value"] = _gthis.__hx_this.get(key)
		return hxrt.AnonymousFromMap(hx_obj_14)
	})
	return hxrt.AnonymousFromMap(hx_obj_5)
}

func (self *haxe__ds__ObjectMap) getIMap(key any) any {
	return self.__hx_this.get(key)
}

func (self *haxe__ds__ObjectMap) setIMap(key any, value any) {
	self.__hx_this.set(key, value)
}

func (self *haxe__ds__ObjectMap) existsIMap(key any) bool {
	return func(hx_value_15 any) bool {
		if hx_value_15 == nil {
			var hx_zero_16 bool
			return hx_zero_16
		}
		return hx_value_15.(bool)
	}(self.__hx_this.exists(key))
}

func (self *haxe__ds__ObjectMap) removeIMap(key any) bool {
	return func(hx_value_17 any) bool {
		if hx_value_17 == nil {
			var hx_zero_18 bool
			return hx_zero_18
		}
		return hx_value_17.(bool)
	}(self.__hx_this.remove(key))
}

func (self *haxe__ds__ObjectMap) copyIMap() haxe__IMap {
	return func(hx_value_19 any) *haxe__ds__ObjectMap {
		if hx_value_19 == nil {
			var hx_zero_20 *haxe__ds__ObjectMap
			return hx_zero_20
		}
		return hx_value_19.(*haxe__ds__ObjectMap)
	}(self.__hx_this.copy())
}

func (self *haxe__ds__ObjectMap) copy() *haxe__ds__ObjectMap {
	copied := New_haxe__ds__ObjectMap()
	key := self.__hx_this.keys()
	for func() bool {
		var hx_callable_24 *hxrt.Function[func() bool] = func(hx_obj_21 hxrt.Anonymous) *hxrt.Function[func() bool] {
			hx_field_22 := hxrt.AnonymousGet(hx_obj_21, "hasNext")
			if hx_field_22 == nil {
				var hx_zero_23 *hxrt.Function[func() bool]
				return hx_zero_23
			}
			return hx_field_22.(*hxrt.Function[func() bool])
		}(key)
		return hxrt.FunctionCall(hx_callable_24)()
	}() {
		var key_1 any = func() any {
			var hx_callable_28 *hxrt.Function[func() any] = func(hx_obj_25 hxrt.Anonymous) *hxrt.Function[func() any] {
				hx_field_26 := hxrt.AnonymousGet(hx_obj_25, "next")
				if hx_field_26 == nil {
					var hx_zero_27 *hxrt.Function[func() any]
					return hx_zero_27
				}
				return hx_field_26.(*hxrt.Function[func() any])
			}(key)
			return hxrt.FunctionCall(hx_callable_28)()
		}()
		copied.__hx_this.set(key_1, self.__hx_this.get(key_1))
	}
	return copied
}

func (self *haxe__ds__ObjectMap) toString() *string {
	var out_b *string
	out_b = hxrt.StringFromLiteral("")
	out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral("["))
	iterator := self.__hx_this.keys()
	for func() bool {
		var hx_callable_32 *hxrt.Function[func() bool] = func(hx_obj_29 hxrt.Anonymous) *hxrt.Function[func() bool] {
			hx_field_30 := hxrt.AnonymousGet(hx_obj_29, "hasNext")
			if hx_field_30 == nil {
				var hx_zero_31 *hxrt.Function[func() bool]
				return hx_zero_31
			}
			return hx_field_30.(*hxrt.Function[func() bool])
		}(iterator)
		return hxrt.FunctionCall(hx_callable_32)()
	}() {
		var key any = func() any {
			var hx_callable_36 *hxrt.Function[func() any] = func(hx_obj_33 hxrt.Anonymous) *hxrt.Function[func() any] {
				hx_field_34 := hxrt.AnonymousGet(hx_obj_33, "next")
				if hx_field_34 == nil {
					var hx_zero_35 *hxrt.Function[func() any]
					return hx_zero_35
				}
				return hx_field_34.(*hxrt.Function[func() any])
			}(iterator)
			return hxrt.FunctionCall(hx_callable_36)()
		}()
		x := hxrt.StdString(key)
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(x))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral(" => "))
		x_1 := hxrt.StdString(self.__hx_this.get(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(x_1))
		if func() bool {
			var hx_callable_40 *hxrt.Function[func() bool] = func(hx_obj_37 hxrt.Anonymous) *hxrt.Function[func() bool] {
				hx_field_38 := hxrt.AnonymousGet(hx_obj_37, "hasNext")
				if hx_field_38 == nil {
					var hx_zero_39 *hxrt.Function[func() bool]
					return hx_zero_39
				}
				return hx_field_38.(*hxrt.Function[func() bool])
			}(iterator)
			return hxrt.FunctionCall(hx_callable_40)()
		}() {
			out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral(", "))
		}
	}
	out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral("]"))
	return out_b
}

func (self *haxe__ds__ObjectMap) clear() {
	hxrt.ObjectMapClear(self.h)
}

func (self *haxe__ds__ObjectMap) String() string {
	return *self.__hx_this.toString()
}
