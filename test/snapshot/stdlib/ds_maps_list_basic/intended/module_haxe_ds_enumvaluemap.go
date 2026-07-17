package main

import "snapshot/hxrt"

type I_haxe__ds___EnumValueMap__EnumValueTreeNode interface {
	getHeight() int
	toString() *string
}

type haxe__ds___EnumValueMap__EnumValueTreeNode struct {
	__hx_this I_haxe__ds___EnumValueMap__EnumValueTreeNode
	left      *haxe__ds___EnumValueMap__EnumValueTreeNode
	right     *haxe__ds___EnumValueMap__EnumValueTreeNode
	key       any
	value     any
	height    int
}

func New_haxe__ds___EnumValueMap__EnumValueTreeNode(left *haxe__ds___EnumValueMap__EnumValueTreeNode, key any, value any, right *haxe__ds___EnumValueMap__EnumValueTreeNode, height int) *haxe__ds___EnumValueMap__EnumValueTreeNode {
	self := &haxe__ds___EnumValueMap__EnumValueTreeNode{}
	self.__hx_this = self
	self.left = left
	self.key = key
	self.value = value
	self.right = right
	if height == -1 {
		var hx_if_21 int
		if left == nil {
			hx_if_21 = 0
		} else {
			hx_if_21 = left.height
		}
		leftHeight := hx_if_21
		var hx_if_22 int
		if right == nil {
			hx_if_22 = 0
		} else {
			hx_if_22 = right.height
		}
		rightHeight := hx_if_22
		self.height = int(int32((hxrt.Int32Wrap(func() int {
			var hx_if_23 int
			if leftHeight > rightHeight {
				hx_if_23 = leftHeight
			} else {
				hx_if_23 = rightHeight
			}
			return hx_if_23
		}()) + hxrt.Int32Wrap(1))))
	} else {
		self.height = height
	}
	return self
}

func (self *haxe__ds___EnumValueMap__EnumValueTreeNode) getHeight() int {
	return self.height
}

func (self *haxe__ds___EnumValueMap__EnumValueTreeNode) toString() *string {
	return hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(func() *string {
		var hx_if_26 *string
		if self.left == nil {
			hx_if_26 = hxrt.StringFromLiteral("")
		} else {
			hx_if_26 = hxrt.StringConcatStringPtr(func(hx_value_24 any) *string {
				if hx_value_24 == nil {
					var hx_zero_25 *string
					return hx_zero_25
				}
				return hx_value_24.(*string)
			}(self.left.toString()), hxrt.StringFromLiteral(", "))
		}
		return hx_if_26
	}(), hxrt.StdString(self.key)), hxrt.StringFromLiteral(" => ")), hxrt.StdString(self.value)), func() *string {
		var hx_if_29 *string
		if self.right == nil {
			hx_if_29 = hxrt.StringFromLiteral("")
		} else {
			hx_if_29 = hxrt.StringConcatStringPtr(hxrt.StringFromLiteral(", "), func(hx_value_27 any) *string {
				if hx_value_27 == nil {
					var hx_zero_28 *string
					return hx_zero_28
				}
				return hx_value_27.(*string)
			}(self.right.toString()))
		}
		return hx_if_29
	}())
}

type I_haxe__ds__EnumValueMap interface {
	set(key any, value any)
	get(key any) any
	exists(key any) bool
	remove(key any) bool
	keys() map[string]any
	iterator() map[string]any
	keyValueIterator() map[string]any
	copy() *haxe__ds__EnumValueMap
	toString() *string
	clear()
	compare(left any, right any) int
	compareArgs(left *hxrt.Array, right *hxrt.Array) int
	compareArg(left any, right any) int
	setLoop(key any, value any, node *haxe__ds___EnumValueMap__EnumValueTreeNode) *haxe__ds___EnumValueMap__EnumValueTreeNode
	removeLoop(key any, node *haxe__ds___EnumValueMap__EnumValueTreeNode) *haxe__ds___EnumValueMap__EnumValueTreeNode
	merge(left *haxe__ds___EnumValueMap__EnumValueTreeNode, right *haxe__ds___EnumValueMap__EnumValueTreeNode) *haxe__ds___EnumValueMap__EnumValueTreeNode
	minBinding(node *haxe__ds___EnumValueMap__EnumValueTreeNode) *haxe__ds___EnumValueMap__EnumValueTreeNode
	removeMinBinding(node *haxe__ds___EnumValueMap__EnumValueTreeNode) *haxe__ds___EnumValueMap__EnumValueTreeNode
	balance(left *haxe__ds___EnumValueMap__EnumValueTreeNode, key any, value any, right *haxe__ds___EnumValueMap__EnumValueTreeNode) *haxe__ds___EnumValueMap__EnumValueTreeNode
	getIMap(key any) any
	setIMap(key any, value any)
	existsIMap(key any) bool
	removeIMap(key any) bool
	copyIMap() haxe__IMap
}

type haxe__ds__EnumValueMap struct {
	__hx_this I_haxe__ds__EnumValueMap
	root      *haxe__ds___EnumValueMap__EnumValueTreeNode
}

func New_haxe__ds__EnumValueMap() *haxe__ds__EnumValueMap {
	self := &haxe__ds__EnumValueMap{}
	self.__hx_this = self
	return self
}

func (self *haxe__ds__EnumValueMap) set(key any, value any) {
	self.root = func(hx_value_137 any) *haxe__ds___EnumValueMap__EnumValueTreeNode {
		if hx_value_137 == nil {
			var hx_zero_138 *haxe__ds___EnumValueMap__EnumValueTreeNode
			return hx_zero_138
		}
		return hx_value_137.(*haxe__ds___EnumValueMap__EnumValueTreeNode)
	}(self.setLoop(key, value, self.root))
}

func (self *haxe__ds__EnumValueMap) get(key any) any {
	node := self.root
	for node != nil {
		result := func(hx_value_139 any) int {
			if hx_value_139 == nil {
				var hx_zero_140 int
				return hx_zero_140
			}
			return hx_value_139.(int)
		}(self.compare(key, node.key))
		if result == 0 {
			return node.value
		}
		var hx_if_141 *haxe__ds___EnumValueMap__EnumValueTreeNode
		if result < 0 {
			hx_if_141 = node.left
		} else {
			hx_if_141 = node.right
		}
		node = hx_if_141
	}
	return nil
}

func (self *haxe__ds__EnumValueMap) exists(key any) bool {
	node := self.root
	for node != nil {
		result := func(hx_value_142 any) int {
			if hx_value_142 == nil {
				var hx_zero_143 int
				return hx_zero_143
			}
			return hx_value_142.(int)
		}(self.compare(key, node.key))
		if result == 0 {
			return true
		}
		var hx_if_144 *haxe__ds___EnumValueMap__EnumValueTreeNode
		if result < 0 {
			hx_if_144 = node.left
		} else {
			hx_if_144 = node.right
		}
		node = hx_if_144
	}
	return false
}

func (self *haxe__ds__EnumValueMap) remove(key any) bool {
	removed := true
	hxrt.TryCatch(func() {
		self.root = func(hx_value_147 any) *haxe__ds___EnumValueMap__EnumValueTreeNode {
			if hx_value_147 == nil {
				var hx_zero_148 *haxe__ds___EnumValueMap__EnumValueTreeNode
				return hx_zero_148
			}
			return hx_value_147.(*haxe__ds___EnumValueMap__EnumValueTreeNode)
		}(self.removeLoop(key, self.root))
	}, func(hx_caught_145 any) {
		switch hx_typed_146 := hx_caught_145.(type) {
		case *string:
			hx_tmp := hx_typed_146
			_ = hx_tmp
			removed = false
		default:
			hxrt.Throw(hx_caught_145)
		}
	})
	return removed
}

func (self *haxe__ds__EnumValueMap) keys() map[string]any {
	values := haxe__ds__EnumValueMap_keysLoop(self.root, hxrt.NewArray())
	index := 0
	hx_obj_149 := map[string]any{}
	hx_obj_149["hasNext"] = func() bool {
		return (index < values.Len())
	}
	hx_obj_149["next"] = func() any {
		hx_post_150 := index
		index = int(int32((index + 1)))
		return values.Get(hx_post_150)
	}
	return hx_obj_149
}

func (self *haxe__ds__EnumValueMap) iterator() map[string]any {
	values := haxe__ds__EnumValueMap_valuesLoop(self.root, hxrt.NewArray())
	index := 0
	hx_obj_151 := map[string]any{}
	hx_obj_151["hasNext"] = func() bool {
		return (index < values.Len())
	}
	hx_obj_151["next"] = func() any {
		hx_post_152 := index
		index = int(int32((index + 1)))
		return values.Get(hx_post_152)
	}
	return hx_obj_151
}

func (self *haxe__ds__EnumValueMap) keyValueIterator() map[string]any {
	_gthis := self
	keys := func(hx_value_153 any) map[string]any {
		if hx_value_153 == nil {
			var hx_zero_154 map[string]any
			return hx_zero_154
		}
		return hx_value_153.(map[string]any)
	}(self.keys())
	hx_obj_155 := map[string]any{}
	hx_obj_155["hasNext"] = func() bool {
		return func(hx_obj_156 map[string]any) func() bool {
			hx_field_157 := hx_obj_156["hasNext"]
			if hx_field_157 == nil {
				var hx_zero_158 func() bool
				return hx_zero_158
			}
			return hx_field_157.(func() bool)
		}(keys)()
	}
	hx_obj_155["next"] = func() map[string]any {
		var key any = func(hx_obj_159 map[string]any) func() any {
			hx_field_160 := hx_obj_159["next"]
			if hx_field_160 == nil {
				var hx_zero_161 func() any
				return hx_zero_161
			}
			return hx_field_160.(func() any)
		}(keys)()
		hx_obj_162 := map[string]any{}
		hx_obj_162["key"] = key
		hx_obj_162["value"] = _gthis.get(key)
		return hx_obj_162
	}
	return hx_obj_155
}

func (self *haxe__ds__EnumValueMap) copy() *haxe__ds__EnumValueMap {
	copied := New_haxe__ds__EnumValueMap()
	copied.root = self.root
	return copied
}

func (self *haxe__ds__EnumValueMap) toString() *string {
	var hx_if_165 *string
	if self.root == nil {
		hx_if_165 = hxrt.StringFromLiteral("[]")
	} else {
		hx_if_165 = hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("["), func(hx_value_163 any) *string {
			if hx_value_163 == nil {
				var hx_zero_164 *string
				return hx_zero_164
			}
			return hx_value_163.(*string)
		}(self.root.toString())), hxrt.StringFromLiteral("]"))
	}
	return hx_if_165
}

func (self *haxe__ds__EnumValueMap) clear() {
	self.root = nil
}

func (self *haxe__ds__EnumValueMap) compare(left any, right any) int {
	result := int(int32((hxrt.Int32Wrap(Type_enumIndex(left)) - hxrt.Int32Wrap(Type_enumIndex(right)))))
	if result != 0 {
		return result
	}
	return func(hx_value_166 any) int {
		if hx_value_166 == nil {
			var hx_zero_167 int
			return hx_zero_167
		}
		return hx_value_166.(int)
	}(self.compareArgs(Type_enumParameters(left), Type_enumParameters(right)))
}

func (self *haxe__ds__EnumValueMap) compareArgs(left *hxrt.Array, right *hxrt.Array) int {
	result := int(int32((hxrt.Int32Wrap(left.Len()) - hxrt.Int32Wrap(right.Len()))))
	if result != 0 {
		return result
	}
	_g := 0
	_g1 := left.Len()
	for _g < _g1 {
		hx_post_168 := _g
		_g = int(int32((_g + 1)))
		index := hx_post_168
		result = func(hx_value_169 any) int {
			if hx_value_169 == nil {
				var hx_zero_170 int
				return hx_zero_170
			}
			return hx_value_169.(int)
		}(self.compareArg(left.Get(index), right.Get(index)))
		if result != 0 {
			return result
		}
	}
	return 0
}

func (self *haxe__ds__EnumValueMap) compareArg(left any, right any) int {
	if haxe__ds__EnumValueMap_isEnumValue(left) && haxe__ds__EnumValueMap_isEnumValue(right) {
		return func(hx_value_171 any) int {
			if hx_value_171 == nil {
				var hx_zero_172 int
				return hx_zero_172
			}
			return hx_value_171.(int)
		}(self.compare(left, right))
	}
	if func(hx_value any) bool {
		switch hx_value.(type) {
		case *hxrt.Array:
			return true
		default:
			return false
		}
	}(any(left)) && func(hx_value any) bool {
		switch hx_value.(type) {
		case *hxrt.Array:
			return true
		default:
			return false
		}
	}(any(right)) {
		return func(hx_value_177 any) int {
			if hx_value_177 == nil {
				var hx_zero_178 int
				return hx_zero_178
			}
			return hx_value_177.(int)
		}(self.compareArgs(func(hx_value_173 any) *hxrt.Array {
			if hx_value_173 == nil {
				var hx_zero_174 *hxrt.Array
				return hx_zero_174
			}
			return hx_value_173.(*hxrt.Array)
		}(left), func(hx_value_175 any) *hxrt.Array {
			if hx_value_175 == nil {
				var hx_zero_176 *hxrt.Array
				return hx_zero_176
			}
			return hx_value_175.(*hxrt.Array)
		}(right)))
	}
	return Reflect_compare(left, right)
}

func (self *haxe__ds__EnumValueMap) setLoop(key any, value any, node *haxe__ds___EnumValueMap__EnumValueTreeNode) *haxe__ds___EnumValueMap__EnumValueTreeNode {
	if node == nil {
		return New_haxe__ds___EnumValueMap__EnumValueTreeNode(nil, key, value, nil, -1)
	}
	result := func(hx_value_179 any) int {
		if hx_value_179 == nil {
			var hx_zero_180 int
			return hx_zero_180
		}
		return hx_value_179.(int)
	}(self.compare(key, node.key))
	if result == 0 {
		return New_haxe__ds___EnumValueMap__EnumValueTreeNode(node.left, key, value, node.right, node.height)
	}
	if result < 0 {
		return func(hx_value_183 any) *haxe__ds___EnumValueMap__EnumValueTreeNode {
			if hx_value_183 == nil {
				var hx_zero_184 *haxe__ds___EnumValueMap__EnumValueTreeNode
				return hx_zero_184
			}
			return hx_value_183.(*haxe__ds___EnumValueMap__EnumValueTreeNode)
		}(self.balance(func(hx_value_181 any) *haxe__ds___EnumValueMap__EnumValueTreeNode {
			if hx_value_181 == nil {
				var hx_zero_182 *haxe__ds___EnumValueMap__EnumValueTreeNode
				return hx_zero_182
			}
			return hx_value_181.(*haxe__ds___EnumValueMap__EnumValueTreeNode)
		}(self.setLoop(key, value, node.left)), node.key, node.value, node.right))
	}
	return func(hx_value_187 any) *haxe__ds___EnumValueMap__EnumValueTreeNode {
		if hx_value_187 == nil {
			var hx_zero_188 *haxe__ds___EnumValueMap__EnumValueTreeNode
			return hx_zero_188
		}
		return hx_value_187.(*haxe__ds___EnumValueMap__EnumValueTreeNode)
	}(self.balance(node.left, node.key, node.value, func(hx_value_185 any) *haxe__ds___EnumValueMap__EnumValueTreeNode {
		if hx_value_185 == nil {
			var hx_zero_186 *haxe__ds___EnumValueMap__EnumValueTreeNode
			return hx_zero_186
		}
		return hx_value_185.(*haxe__ds___EnumValueMap__EnumValueTreeNode)
	}(self.setLoop(key, value, node.right))))
}

func (self *haxe__ds__EnumValueMap) removeLoop(key any, node *haxe__ds___EnumValueMap__EnumValueTreeNode) *haxe__ds___EnumValueMap__EnumValueTreeNode {
	if node == nil {
		hxrt.Throw(hxrt.StringFromLiteral("Not_found"))
	}
	result := func(hx_value_189 any) int {
		if hx_value_189 == nil {
			var hx_zero_190 int
			return hx_zero_190
		}
		return hx_value_189.(int)
	}(self.compare(key, node.key))
	if result == 0 {
		return func(hx_value_191 any) *haxe__ds___EnumValueMap__EnumValueTreeNode {
			if hx_value_191 == nil {
				var hx_zero_192 *haxe__ds___EnumValueMap__EnumValueTreeNode
				return hx_zero_192
			}
			return hx_value_191.(*haxe__ds___EnumValueMap__EnumValueTreeNode)
		}(self.merge(node.left, node.right))
	}
	if result < 0 {
		return func(hx_value_195 any) *haxe__ds___EnumValueMap__EnumValueTreeNode {
			if hx_value_195 == nil {
				var hx_zero_196 *haxe__ds___EnumValueMap__EnumValueTreeNode
				return hx_zero_196
			}
			return hx_value_195.(*haxe__ds___EnumValueMap__EnumValueTreeNode)
		}(self.balance(func(hx_value_193 any) *haxe__ds___EnumValueMap__EnumValueTreeNode {
			if hx_value_193 == nil {
				var hx_zero_194 *haxe__ds___EnumValueMap__EnumValueTreeNode
				return hx_zero_194
			}
			return hx_value_193.(*haxe__ds___EnumValueMap__EnumValueTreeNode)
		}(self.removeLoop(key, node.left)), node.key, node.value, node.right))
	}
	return func(hx_value_199 any) *haxe__ds___EnumValueMap__EnumValueTreeNode {
		if hx_value_199 == nil {
			var hx_zero_200 *haxe__ds___EnumValueMap__EnumValueTreeNode
			return hx_zero_200
		}
		return hx_value_199.(*haxe__ds___EnumValueMap__EnumValueTreeNode)
	}(self.balance(node.left, node.key, node.value, func(hx_value_197 any) *haxe__ds___EnumValueMap__EnumValueTreeNode {
		if hx_value_197 == nil {
			var hx_zero_198 *haxe__ds___EnumValueMap__EnumValueTreeNode
			return hx_zero_198
		}
		return hx_value_197.(*haxe__ds___EnumValueMap__EnumValueTreeNode)
	}(self.removeLoop(key, node.right))))
}

func (self *haxe__ds__EnumValueMap) merge(left *haxe__ds___EnumValueMap__EnumValueTreeNode, right *haxe__ds___EnumValueMap__EnumValueTreeNode) *haxe__ds___EnumValueMap__EnumValueTreeNode {
	if left == nil {
		return right
	}
	if right == nil {
		return left
	}
	minimum := func(hx_value_201 any) *haxe__ds___EnumValueMap__EnumValueTreeNode {
		if hx_value_201 == nil {
			var hx_zero_202 *haxe__ds___EnumValueMap__EnumValueTreeNode
			return hx_zero_202
		}
		return hx_value_201.(*haxe__ds___EnumValueMap__EnumValueTreeNode)
	}(self.minBinding(right))
	return func(hx_value_205 any) *haxe__ds___EnumValueMap__EnumValueTreeNode {
		if hx_value_205 == nil {
			var hx_zero_206 *haxe__ds___EnumValueMap__EnumValueTreeNode
			return hx_zero_206
		}
		return hx_value_205.(*haxe__ds___EnumValueMap__EnumValueTreeNode)
	}(self.balance(left, minimum.key, minimum.value, func(hx_value_203 any) *haxe__ds___EnumValueMap__EnumValueTreeNode {
		if hx_value_203 == nil {
			var hx_zero_204 *haxe__ds___EnumValueMap__EnumValueTreeNode
			return hx_zero_204
		}
		return hx_value_203.(*haxe__ds___EnumValueMap__EnumValueTreeNode)
	}(self.removeMinBinding(right))))
}

func (self *haxe__ds__EnumValueMap) minBinding(node *haxe__ds___EnumValueMap__EnumValueTreeNode) *haxe__ds___EnumValueMap__EnumValueTreeNode {
	if node == nil {
		hxrt.Throw(hxrt.StringFromLiteral("Not_found"))
	}
	var hx_if_209 *haxe__ds___EnumValueMap__EnumValueTreeNode
	if node.left == nil {
		hx_if_209 = node
	} else {
		hx_if_209 = func(hx_value_207 any) *haxe__ds___EnumValueMap__EnumValueTreeNode {
			if hx_value_207 == nil {
				var hx_zero_208 *haxe__ds___EnumValueMap__EnumValueTreeNode
				return hx_zero_208
			}
			return hx_value_207.(*haxe__ds___EnumValueMap__EnumValueTreeNode)
		}(self.minBinding(node.left))
	}
	return hx_if_209
}

func (self *haxe__ds__EnumValueMap) removeMinBinding(node *haxe__ds___EnumValueMap__EnumValueTreeNode) *haxe__ds___EnumValueMap__EnumValueTreeNode {
	var hx_if_214 *haxe__ds___EnumValueMap__EnumValueTreeNode
	if node.left == nil {
		hx_if_214 = node.right
	} else {
		hx_if_214 = func(hx_value_212 any) *haxe__ds___EnumValueMap__EnumValueTreeNode {
			if hx_value_212 == nil {
				var hx_zero_213 *haxe__ds___EnumValueMap__EnumValueTreeNode
				return hx_zero_213
			}
			return hx_value_212.(*haxe__ds___EnumValueMap__EnumValueTreeNode)
		}(self.balance(func(hx_value_210 any) *haxe__ds___EnumValueMap__EnumValueTreeNode {
			if hx_value_210 == nil {
				var hx_zero_211 *haxe__ds___EnumValueMap__EnumValueTreeNode
				return hx_zero_211
			}
			return hx_value_210.(*haxe__ds___EnumValueMap__EnumValueTreeNode)
		}(self.removeMinBinding(node.left)), node.key, node.value, node.right))
	}
	return hx_if_214
}

func (self *haxe__ds__EnumValueMap) balance(left *haxe__ds___EnumValueMap__EnumValueTreeNode, key any, value any, right *haxe__ds___EnumValueMap__EnumValueTreeNode) *haxe__ds___EnumValueMap__EnumValueTreeNode {
	var hx_if_215 int
	if left == nil {
		hx_if_215 = 0
	} else {
		hx_if_215 = left.height
	}
	leftHeight := hx_if_215
	var hx_if_216 int
	if right == nil {
		hx_if_216 = 0
	} else {
		hx_if_216 = right.height
	}
	rightHeight := hx_if_216
	if leftHeight > int(int32((hxrt.Int32Wrap(rightHeight) + hxrt.Int32Wrap(2)))) {
		var hx_if_217 int
		if left.left == nil {
			hx_if_217 = 0
		} else {
			hx_if_217 = left.left.height
		}
		leftLeftHeight := hx_if_217
		var hx_if_218 int
		if left.right == nil {
			hx_if_218 = 0
		} else {
			hx_if_218 = left.right.height
		}
		leftRightHeight := hx_if_218
		if leftLeftHeight >= leftRightHeight {
			return New_haxe__ds___EnumValueMap__EnumValueTreeNode(left.left, left.key, left.value, New_haxe__ds___EnumValueMap__EnumValueTreeNode(left.right, key, value, right, -1), -1)
		}
		return New_haxe__ds___EnumValueMap__EnumValueTreeNode(New_haxe__ds___EnumValueMap__EnumValueTreeNode(left.left, left.key, left.value, left.right.left, -1), left.right.key, left.right.value, New_haxe__ds___EnumValueMap__EnumValueTreeNode(left.right.right, key, value, right, -1), -1)
	}
	if rightHeight > int(int32((hxrt.Int32Wrap(leftHeight) + hxrt.Int32Wrap(2)))) {
		var hx_if_219 int
		if right.right == nil {
			hx_if_219 = 0
		} else {
			hx_if_219 = right.right.height
		}
		rightRightHeight := hx_if_219
		var hx_if_220 int
		if right.left == nil {
			hx_if_220 = 0
		} else {
			hx_if_220 = right.left.height
		}
		rightLeftHeight := hx_if_220
		if rightRightHeight > rightLeftHeight {
			return New_haxe__ds___EnumValueMap__EnumValueTreeNode(New_haxe__ds___EnumValueMap__EnumValueTreeNode(left, key, value, right.left, -1), right.key, right.value, right.right, -1)
		}
		return New_haxe__ds___EnumValueMap__EnumValueTreeNode(New_haxe__ds___EnumValueMap__EnumValueTreeNode(left, key, value, right.left.left, -1), right.left.key, right.left.value, New_haxe__ds___EnumValueMap__EnumValueTreeNode(right.left.right, right.key, right.value, right.right, -1), -1)
	}
	return New_haxe__ds___EnumValueMap__EnumValueTreeNode(left, key, value, right, int(int32((hxrt.Int32Wrap(func() int {
		var hx_if_221 int
		if leftHeight > rightHeight {
			hx_if_221 = leftHeight
		} else {
			hx_if_221 = rightHeight
		}
		return hx_if_221
	}()) + hxrt.Int32Wrap(1)))))
}

func (self *haxe__ds__EnumValueMap) getIMap(key any) any {
	return self.get(key)
}

func (self *haxe__ds__EnumValueMap) setIMap(key any, value any) {
	self.set(key, value)
}

func (self *haxe__ds__EnumValueMap) existsIMap(key any) bool {
	return func(hx_value_222 any) bool {
		if hx_value_222 == nil {
			var hx_zero_223 bool
			return hx_zero_223
		}
		return hx_value_222.(bool)
	}(self.exists(key))
}

func (self *haxe__ds__EnumValueMap) removeIMap(key any) bool {
	return func(hx_value_224 any) bool {
		if hx_value_224 == nil {
			var hx_zero_225 bool
			return hx_zero_225
		}
		return hx_value_224.(bool)
	}(self.remove(key))
}

func (self *haxe__ds__EnumValueMap) copyIMap() haxe__IMap {
	return func(hx_value_226 any) *haxe__ds__EnumValueMap {
		if hx_value_226 == nil {
			var hx_zero_227 *haxe__ds__EnumValueMap
			return hx_zero_227
		}
		return hx_value_226.(*haxe__ds__EnumValueMap)
	}(self.copy())
}

func haxe__ds__EnumValueMap_isEnumValue(value any) bool {
	return hxrt.IsEnumValue(value)
}

func haxe__ds__EnumValueMap_keysLoop(node *haxe__ds___EnumValueMap__EnumValueTreeNode, out *hxrt.Array) *hxrt.Array {
	if node != nil {
		haxe__ds__EnumValueMap_keysLoop(node.left, out)
		out.Push(node.key)
		haxe__ds__EnumValueMap_keysLoop(node.right, out)
	}
	return out
}

func haxe__ds__EnumValueMap_valuesLoop(node *haxe__ds___EnumValueMap__EnumValueTreeNode, out *hxrt.Array) *hxrt.Array {
	if node != nil {
		haxe__ds__EnumValueMap_valuesLoop(node.left, out)
		out.Push(node.value)
		haxe__ds__EnumValueMap_valuesLoop(node.right, out)
	}
	return out
}
