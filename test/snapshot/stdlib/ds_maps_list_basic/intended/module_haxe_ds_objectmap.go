package main

import "snapshot/hxrt"

type I_haxe__ds__ObjectMap interface {
	set(key any, value any)
	get(key any) any
	exists(key any) bool
	remove(key any) bool
	keys() map[string]any
	iterator() map[string]any
	keyValueIterator() map[string]any
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

func (self *haxe__ds__ObjectMap) keys() map[string]any {
	keys := hxrt.ObjectMapKeys(self.h)
	index := 0
	hx_obj_77 := map[string]any{}
	hx_obj_77["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_77["next"] = func() any {
		return keys[func() int {
			hx_post_78 := index
			index = int(int32((index + 1)))
			return hx_post_78
		}()]
	}
	return hx_obj_77
}

func (self *haxe__ds__ObjectMap) iterator() map[string]any {
	_gthis := self
	keys := hxrt.ObjectMapKeys(self.h)
	index := 0
	hx_obj_79 := map[string]any{}
	hx_obj_79["hasNext"] = func() bool {
		return (index < len(keys))
	}
	hx_obj_79["next"] = func() any {
		return hxrt.ObjectMapGet(_gthis.h, keys[func() int {
			hx_post_80 := index
			index = int(int32((index + 1)))
			return hx_post_80
		}()])
	}
	return hx_obj_79
}

func (self *haxe__ds__ObjectMap) keyValueIterator() map[string]any {
	_gthis := self
	keys := func(hx_value_81 any) map[string]any {
		if hx_value_81 == nil {
			var hx_zero_82 map[string]any
			return hx_zero_82
		}
		return hx_value_81.(map[string]any)
	}(self.keys())
	hx_obj_83 := map[string]any{}
	hx_obj_83["hasNext"] = func() bool {
		return func(hx_obj_84 map[string]any) func() bool {
			hx_field_85 := hx_obj_84["hasNext"]
			if hx_field_85 == nil {
				var hx_zero_86 func() bool
				return hx_zero_86
			}
			return hx_field_85.(func() bool)
		}(keys)()
	}
	hx_obj_83["next"] = func() map[string]any {
		var key any = func(hx_obj_87 map[string]any) func() any {
			hx_field_88 := hx_obj_87["next"]
			if hx_field_88 == nil {
				var hx_zero_89 func() any
				return hx_zero_89
			}
			return hx_field_88.(func() any)
		}(keys)()
		hx_obj_90 := map[string]any{}
		hx_obj_90["key"] = key
		hx_obj_90["value"] = _gthis.get(key)
		return hx_obj_90
	}
	return hx_obj_83
}

func (self *haxe__ds__ObjectMap) getIMap(key any) any {
	return self.get(key)
}

func (self *haxe__ds__ObjectMap) setIMap(key any, value any) {
	self.set(key, value)
}

func (self *haxe__ds__ObjectMap) existsIMap(key any) bool {
	return func(hx_value_91 any) bool {
		if hx_value_91 == nil {
			var hx_zero_92 bool
			return hx_zero_92
		}
		return hx_value_91.(bool)
	}(self.exists(key))
}

func (self *haxe__ds__ObjectMap) removeIMap(key any) bool {
	return func(hx_value_93 any) bool {
		if hx_value_93 == nil {
			var hx_zero_94 bool
			return hx_zero_94
		}
		return hx_value_93.(bool)
	}(self.remove(key))
}

func (self *haxe__ds__ObjectMap) copyIMap() haxe__IMap {
	return func(hx_value_95 any) *haxe__ds__ObjectMap {
		if hx_value_95 == nil {
			var hx_zero_96 *haxe__ds__ObjectMap
			return hx_zero_96
		}
		return hx_value_95.(*haxe__ds__ObjectMap)
	}(self.copy())
}

func (self *haxe__ds__ObjectMap) copy() *haxe__ds__ObjectMap {
	copied := New_haxe__ds__ObjectMap()
	key := func(hx_value_97 any) map[string]any {
		if hx_value_97 == nil {
			var hx_zero_98 map[string]any
			return hx_zero_98
		}
		return hx_value_97.(map[string]any)
	}(self.keys())
	for func(hx_obj_99 map[string]any) func() bool {
		hx_field_100 := hx_obj_99["hasNext"]
		if hx_field_100 == nil {
			var hx_zero_101 func() bool
			return hx_zero_101
		}
		return hx_field_100.(func() bool)
	}(key)() {
		var key_1 any = func(hx_obj_102 map[string]any) func() any {
			hx_field_103 := hx_obj_102["next"]
			if hx_field_103 == nil {
				var hx_zero_104 func() any
				return hx_zero_104
			}
			return hx_field_103.(func() any)
		}(key)()
		copied.set(key_1, self.get(key_1))
	}
	return copied
}

func (self *haxe__ds__ObjectMap) toString() *string {
	var out_b *string
	out_b = hxrt.StringFromLiteral("")
	out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral("["))
	iterator := func(hx_value_105 any) map[string]any {
		if hx_value_105 == nil {
			var hx_zero_106 map[string]any
			return hx_zero_106
		}
		return hx_value_105.(map[string]any)
	}(self.keys())
	for func(hx_obj_107 map[string]any) func() bool {
		hx_field_108 := hx_obj_107["hasNext"]
		if hx_field_108 == nil {
			var hx_zero_109 func() bool
			return hx_zero_109
		}
		return hx_field_108.(func() bool)
	}(iterator)() {
		var key any = func(hx_obj_110 map[string]any) func() any {
			hx_field_111 := hx_obj_110["next"]
			if hx_field_111 == nil {
				var hx_zero_112 func() any
				return hx_zero_112
			}
			return hx_field_111.(func() any)
		}(iterator)()
		x := hxrt.StdString(key)
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(x))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral(" => "))
		x_1 := hxrt.StdString(self.get(key))
		out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StdString(x_1))
		if func(hx_obj_113 map[string]any) func() bool {
			hx_field_114 := hx_obj_113["hasNext"]
			if hx_field_114 == nil {
				var hx_zero_115 func() bool
				return hx_zero_115
			}
			return hx_field_114.(func() bool)
		}(iterator)() {
			out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral(", "))
		}
	}
	out_b = hxrt.StringConcatStringPtr(out_b, hxrt.StringFromLiteral("]"))
	return out_b
}

func (self *haxe__ds__ObjectMap) clear() {
	hxrt.ObjectMapClear(self.h)
}
