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
		var hx_if_3 int
		if left == nil {
			hx_if_3 = 0
		} else {
			hx_if_3 = left.height
		}
		leftHeight := hx_if_3
		var hx_if_4 int
		if right == nil {
			hx_if_4 = 0
		} else {
			hx_if_4 = right.height
		}
		rightHeight := hx_if_4
		self.height = int(int32((hxrt.Int32Wrap(func() int {
			var hx_if_5 int
			if leftHeight > rightHeight {
				hx_if_5 = leftHeight
			} else {
				hx_if_5 = rightHeight
			}
			return hx_if_5
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
		var hx_if_8 *string
		if self.left == nil {
			hx_if_8 = hxrt.StringFromLiteral("")
		} else {
			hx_if_8 = hxrt.StringConcatStringPtr(func(hx_value_6 any) *string {
				if hx_value_6 == nil {
					var hx_zero_7 *string
					return hx_zero_7
				}
				return hx_value_6.(*string)
			}(self.left.toString()), hxrt.StringFromLiteral(", "))
		}
		return hx_if_8
	}(), hxrt.StdString(self.key)), hxrt.StringFromLiteral(" => ")), hxrt.StdString(self.value)), func() *string {
		var hx_if_11 *string
		if self.right == nil {
			hx_if_11 = hxrt.StringFromLiteral("")
		} else {
			hx_if_11 = hxrt.StringConcatStringPtr(hxrt.StringFromLiteral(", "), func(hx_value_9 any) *string {
				if hx_value_9 == nil {
					var hx_zero_10 *string
					return hx_zero_10
				}
				return hx_value_9.(*string)
			}(self.right.toString()))
		}
		return hx_if_11
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
	self.root = func(hx_value_12 any) *haxe__ds___EnumValueMap__EnumValueTreeNode {
		if hx_value_12 == nil {
			var hx_zero_13 *haxe__ds___EnumValueMap__EnumValueTreeNode
			return hx_zero_13
		}
		return hx_value_12.(*haxe__ds___EnumValueMap__EnumValueTreeNode)
	}(self.setLoop(key, value, self.root))
}

func (self *haxe__ds__EnumValueMap) get(key any) any {
	node := self.root
	for node != nil {
		result := func(hx_value_14 any) int {
			if hx_value_14 == nil {
				var hx_zero_15 int
				return hx_zero_15
			}
			return hx_value_14.(int)
		}(self.compare(key, node.key))
		if result == 0 {
			return node.value
		}
		var hx_if_16 *haxe__ds___EnumValueMap__EnumValueTreeNode
		if result < 0 {
			hx_if_16 = node.left
		} else {
			hx_if_16 = node.right
		}
		node = hx_if_16
	}
	return nil
}

func (self *haxe__ds__EnumValueMap) exists(key any) bool {
	node := self.root
	for node != nil {
		result := func(hx_value_17 any) int {
			if hx_value_17 == nil {
				var hx_zero_18 int
				return hx_zero_18
			}
			return hx_value_17.(int)
		}(self.compare(key, node.key))
		if result == 0 {
			return true
		}
		var hx_if_19 *haxe__ds___EnumValueMap__EnumValueTreeNode
		if result < 0 {
			hx_if_19 = node.left
		} else {
			hx_if_19 = node.right
		}
		node = hx_if_19
	}
	return false
}

func (self *haxe__ds__EnumValueMap) remove(key any) bool {
	removed := true
	hxrt.TryCatch(func() {
		self.root = func(hx_value_22 any) *haxe__ds___EnumValueMap__EnumValueTreeNode {
			if hx_value_22 == nil {
				var hx_zero_23 *haxe__ds___EnumValueMap__EnumValueTreeNode
				return hx_zero_23
			}
			return hx_value_22.(*haxe__ds___EnumValueMap__EnumValueTreeNode)
		}(self.removeLoop(key, self.root))
	}, func(hx_caught_20 any) {
		switch hx_typed_21 := hx_caught_20.(type) {
		case *string:
			hx_tmp := hx_typed_21
			_ = hx_tmp
			removed = false
		default:
			hxrt.Throw(hx_caught_20)
		}
	})
	return removed
}

func (self *haxe__ds__EnumValueMap) keys() map[string]any {
	values := haxe__ds__EnumValueMap_keysLoop(self.root, hxrt.NewArray())
	index := 0
	hx_obj_24 := map[string]any{}
	hx_obj_24["hasNext"] = func() bool {
		return (index < values.Len())
	}
	hx_obj_24["next"] = func() any {
		hx_post_25 := index
		index = int(int32((index + 1)))
		return values.Get(hx_post_25)
	}
	return hx_obj_24
}

func (self *haxe__ds__EnumValueMap) iterator() map[string]any {
	values := haxe__ds__EnumValueMap_valuesLoop(self.root, hxrt.NewArray())
	index := 0
	hx_obj_26 := map[string]any{}
	hx_obj_26["hasNext"] = func() bool {
		return (index < values.Len())
	}
	hx_obj_26["next"] = func() any {
		hx_post_27 := index
		index = int(int32((index + 1)))
		return values.Get(hx_post_27)
	}
	return hx_obj_26
}

func (self *haxe__ds__EnumValueMap) keyValueIterator() map[string]any {
	_gthis := self
	keys := func(hx_value_28 any) map[string]any {
		if hx_value_28 == nil {
			var hx_zero_29 map[string]any
			return hx_zero_29
		}
		return hx_value_28.(map[string]any)
	}(self.keys())
	hx_obj_30 := map[string]any{}
	hx_obj_30["hasNext"] = func() bool {
		return func(hx_obj_31 map[string]any) func() bool {
			hx_field_32 := hx_obj_31["hasNext"]
			if hx_field_32 == nil {
				var hx_zero_33 func() bool
				return hx_zero_33
			}
			return hx_field_32.(func() bool)
		}(keys)()
	}
	hx_obj_30["next"] = func() map[string]any {
		var key any = func(hx_obj_34 map[string]any) func() any {
			hx_field_35 := hx_obj_34["next"]
			if hx_field_35 == nil {
				var hx_zero_36 func() any
				return hx_zero_36
			}
			return hx_field_35.(func() any)
		}(keys)()
		hx_obj_37 := map[string]any{}
		hx_obj_37["key"] = key
		hx_obj_37["value"] = _gthis.get(key)
		return hx_obj_37
	}
	return hx_obj_30
}

func (self *haxe__ds__EnumValueMap) copy() *haxe__ds__EnumValueMap {
	copied := New_haxe__ds__EnumValueMap()
	copied.root = self.root
	return copied
}

func (self *haxe__ds__EnumValueMap) toString() *string {
	var hx_if_40 *string
	if self.root == nil {
		hx_if_40 = hxrt.StringFromLiteral("[]")
	} else {
		hx_if_40 = hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("["), func(hx_value_38 any) *string {
			if hx_value_38 == nil {
				var hx_zero_39 *string
				return hx_zero_39
			}
			return hx_value_38.(*string)
		}(self.root.toString())), hxrt.StringFromLiteral("]"))
	}
	return hx_if_40
}

func (self *haxe__ds__EnumValueMap) clear() {
	self.root = nil
}

func (self *haxe__ds__EnumValueMap) compare(left any, right any) int {
	result := int(int32((hxrt.Int32Wrap(Type_enumIndex(left)) - hxrt.Int32Wrap(Type_enumIndex(right)))))
	if result != 0 {
		return result
	}
	return func(hx_value_41 any) int {
		if hx_value_41 == nil {
			var hx_zero_42 int
			return hx_zero_42
		}
		return hx_value_41.(int)
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
		hx_post_43 := _g
		_g = int(int32((_g + 1)))
		index := hx_post_43
		result = func(hx_value_44 any) int {
			if hx_value_44 == nil {
				var hx_zero_45 int
				return hx_zero_45
			}
			return hx_value_44.(int)
		}(self.compareArg(left.Get(index), right.Get(index)))
		if result != 0 {
			return result
		}
	}
	return 0
}

func (self *haxe__ds__EnumValueMap) compareArg(left any, right any) int {
	if haxe__ds__EnumValueMap_isEnumValue(left) && haxe__ds__EnumValueMap_isEnumValue(right) {
		return func(hx_value_46 any) int {
			if hx_value_46 == nil {
				var hx_zero_47 int
				return hx_zero_47
			}
			return hx_value_46.(int)
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
		return func(hx_value_52 any) int {
			if hx_value_52 == nil {
				var hx_zero_53 int
				return hx_zero_53
			}
			return hx_value_52.(int)
		}(self.compareArgs(func(hx_value_48 any) *hxrt.Array {
			if hx_value_48 == nil {
				var hx_zero_49 *hxrt.Array
				return hx_zero_49
			}
			return hx_value_48.(*hxrt.Array)
		}(left), func(hx_value_50 any) *hxrt.Array {
			if hx_value_50 == nil {
				var hx_zero_51 *hxrt.Array
				return hx_zero_51
			}
			return hx_value_50.(*hxrt.Array)
		}(right)))
	}
	return Reflect_compare(left, right)
}

func (self *haxe__ds__EnumValueMap) setLoop(key any, value any, node *haxe__ds___EnumValueMap__EnumValueTreeNode) *haxe__ds___EnumValueMap__EnumValueTreeNode {
	if node == nil {
		return New_haxe__ds___EnumValueMap__EnumValueTreeNode(nil, key, value, nil, -1)
	}
	result := func(hx_value_54 any) int {
		if hx_value_54 == nil {
			var hx_zero_55 int
			return hx_zero_55
		}
		return hx_value_54.(int)
	}(self.compare(key, node.key))
	if result == 0 {
		return New_haxe__ds___EnumValueMap__EnumValueTreeNode(node.left, key, value, node.right, node.height)
	}
	if result < 0 {
		return func(hx_value_58 any) *haxe__ds___EnumValueMap__EnumValueTreeNode {
			if hx_value_58 == nil {
				var hx_zero_59 *haxe__ds___EnumValueMap__EnumValueTreeNode
				return hx_zero_59
			}
			return hx_value_58.(*haxe__ds___EnumValueMap__EnumValueTreeNode)
		}(self.balance(func(hx_value_56 any) *haxe__ds___EnumValueMap__EnumValueTreeNode {
			if hx_value_56 == nil {
				var hx_zero_57 *haxe__ds___EnumValueMap__EnumValueTreeNode
				return hx_zero_57
			}
			return hx_value_56.(*haxe__ds___EnumValueMap__EnumValueTreeNode)
		}(self.setLoop(key, value, node.left)), node.key, node.value, node.right))
	}
	return func(hx_value_62 any) *haxe__ds___EnumValueMap__EnumValueTreeNode {
		if hx_value_62 == nil {
			var hx_zero_63 *haxe__ds___EnumValueMap__EnumValueTreeNode
			return hx_zero_63
		}
		return hx_value_62.(*haxe__ds___EnumValueMap__EnumValueTreeNode)
	}(self.balance(node.left, node.key, node.value, func(hx_value_60 any) *haxe__ds___EnumValueMap__EnumValueTreeNode {
		if hx_value_60 == nil {
			var hx_zero_61 *haxe__ds___EnumValueMap__EnumValueTreeNode
			return hx_zero_61
		}
		return hx_value_60.(*haxe__ds___EnumValueMap__EnumValueTreeNode)
	}(self.setLoop(key, value, node.right))))
}

func (self *haxe__ds__EnumValueMap) removeLoop(key any, node *haxe__ds___EnumValueMap__EnumValueTreeNode) *haxe__ds___EnumValueMap__EnumValueTreeNode {
	if node == nil {
		hxrt.Throw(hxrt.StringFromLiteral("Not_found"))
	}
	result := func(hx_value_64 any) int {
		if hx_value_64 == nil {
			var hx_zero_65 int
			return hx_zero_65
		}
		return hx_value_64.(int)
	}(self.compare(key, node.key))
	if result == 0 {
		return func(hx_value_66 any) *haxe__ds___EnumValueMap__EnumValueTreeNode {
			if hx_value_66 == nil {
				var hx_zero_67 *haxe__ds___EnumValueMap__EnumValueTreeNode
				return hx_zero_67
			}
			return hx_value_66.(*haxe__ds___EnumValueMap__EnumValueTreeNode)
		}(self.merge(node.left, node.right))
	}
	if result < 0 {
		return func(hx_value_70 any) *haxe__ds___EnumValueMap__EnumValueTreeNode {
			if hx_value_70 == nil {
				var hx_zero_71 *haxe__ds___EnumValueMap__EnumValueTreeNode
				return hx_zero_71
			}
			return hx_value_70.(*haxe__ds___EnumValueMap__EnumValueTreeNode)
		}(self.balance(func(hx_value_68 any) *haxe__ds___EnumValueMap__EnumValueTreeNode {
			if hx_value_68 == nil {
				var hx_zero_69 *haxe__ds___EnumValueMap__EnumValueTreeNode
				return hx_zero_69
			}
			return hx_value_68.(*haxe__ds___EnumValueMap__EnumValueTreeNode)
		}(self.removeLoop(key, node.left)), node.key, node.value, node.right))
	}
	return func(hx_value_74 any) *haxe__ds___EnumValueMap__EnumValueTreeNode {
		if hx_value_74 == nil {
			var hx_zero_75 *haxe__ds___EnumValueMap__EnumValueTreeNode
			return hx_zero_75
		}
		return hx_value_74.(*haxe__ds___EnumValueMap__EnumValueTreeNode)
	}(self.balance(node.left, node.key, node.value, func(hx_value_72 any) *haxe__ds___EnumValueMap__EnumValueTreeNode {
		if hx_value_72 == nil {
			var hx_zero_73 *haxe__ds___EnumValueMap__EnumValueTreeNode
			return hx_zero_73
		}
		return hx_value_72.(*haxe__ds___EnumValueMap__EnumValueTreeNode)
	}(self.removeLoop(key, node.right))))
}

func (self *haxe__ds__EnumValueMap) merge(left *haxe__ds___EnumValueMap__EnumValueTreeNode, right *haxe__ds___EnumValueMap__EnumValueTreeNode) *haxe__ds___EnumValueMap__EnumValueTreeNode {
	if left == nil {
		return right
	}
	if right == nil {
		return left
	}
	minimum := func(hx_value_76 any) *haxe__ds___EnumValueMap__EnumValueTreeNode {
		if hx_value_76 == nil {
			var hx_zero_77 *haxe__ds___EnumValueMap__EnumValueTreeNode
			return hx_zero_77
		}
		return hx_value_76.(*haxe__ds___EnumValueMap__EnumValueTreeNode)
	}(self.minBinding(right))
	return func(hx_value_80 any) *haxe__ds___EnumValueMap__EnumValueTreeNode {
		if hx_value_80 == nil {
			var hx_zero_81 *haxe__ds___EnumValueMap__EnumValueTreeNode
			return hx_zero_81
		}
		return hx_value_80.(*haxe__ds___EnumValueMap__EnumValueTreeNode)
	}(self.balance(left, minimum.key, minimum.value, func(hx_value_78 any) *haxe__ds___EnumValueMap__EnumValueTreeNode {
		if hx_value_78 == nil {
			var hx_zero_79 *haxe__ds___EnumValueMap__EnumValueTreeNode
			return hx_zero_79
		}
		return hx_value_78.(*haxe__ds___EnumValueMap__EnumValueTreeNode)
	}(self.removeMinBinding(right))))
}

func (self *haxe__ds__EnumValueMap) minBinding(node *haxe__ds___EnumValueMap__EnumValueTreeNode) *haxe__ds___EnumValueMap__EnumValueTreeNode {
	if node == nil {
		hxrt.Throw(hxrt.StringFromLiteral("Not_found"))
	}
	var hx_if_84 *haxe__ds___EnumValueMap__EnumValueTreeNode
	if node.left == nil {
		hx_if_84 = node
	} else {
		hx_if_84 = func(hx_value_82 any) *haxe__ds___EnumValueMap__EnumValueTreeNode {
			if hx_value_82 == nil {
				var hx_zero_83 *haxe__ds___EnumValueMap__EnumValueTreeNode
				return hx_zero_83
			}
			return hx_value_82.(*haxe__ds___EnumValueMap__EnumValueTreeNode)
		}(self.minBinding(node.left))
	}
	return hx_if_84
}

func (self *haxe__ds__EnumValueMap) removeMinBinding(node *haxe__ds___EnumValueMap__EnumValueTreeNode) *haxe__ds___EnumValueMap__EnumValueTreeNode {
	var hx_if_89 *haxe__ds___EnumValueMap__EnumValueTreeNode
	if node.left == nil {
		hx_if_89 = node.right
	} else {
		hx_if_89 = func(hx_value_87 any) *haxe__ds___EnumValueMap__EnumValueTreeNode {
			if hx_value_87 == nil {
				var hx_zero_88 *haxe__ds___EnumValueMap__EnumValueTreeNode
				return hx_zero_88
			}
			return hx_value_87.(*haxe__ds___EnumValueMap__EnumValueTreeNode)
		}(self.balance(func(hx_value_85 any) *haxe__ds___EnumValueMap__EnumValueTreeNode {
			if hx_value_85 == nil {
				var hx_zero_86 *haxe__ds___EnumValueMap__EnumValueTreeNode
				return hx_zero_86
			}
			return hx_value_85.(*haxe__ds___EnumValueMap__EnumValueTreeNode)
		}(self.removeMinBinding(node.left)), node.key, node.value, node.right))
	}
	return hx_if_89
}

func (self *haxe__ds__EnumValueMap) balance(left *haxe__ds___EnumValueMap__EnumValueTreeNode, key any, value any, right *haxe__ds___EnumValueMap__EnumValueTreeNode) *haxe__ds___EnumValueMap__EnumValueTreeNode {
	var hx_if_90 int
	if left == nil {
		hx_if_90 = 0
	} else {
		hx_if_90 = left.height
	}
	leftHeight := hx_if_90
	var hx_if_91 int
	if right == nil {
		hx_if_91 = 0
	} else {
		hx_if_91 = right.height
	}
	rightHeight := hx_if_91
	if leftHeight > int(int32((hxrt.Int32Wrap(rightHeight) + hxrt.Int32Wrap(2)))) {
		var hx_if_92 int
		if left.left == nil {
			hx_if_92 = 0
		} else {
			hx_if_92 = left.left.height
		}
		leftLeftHeight := hx_if_92
		var hx_if_93 int
		if left.right == nil {
			hx_if_93 = 0
		} else {
			hx_if_93 = left.right.height
		}
		leftRightHeight := hx_if_93
		if leftLeftHeight >= leftRightHeight {
			return New_haxe__ds___EnumValueMap__EnumValueTreeNode(left.left, left.key, left.value, New_haxe__ds___EnumValueMap__EnumValueTreeNode(left.right, key, value, right, -1), -1)
		}
		return New_haxe__ds___EnumValueMap__EnumValueTreeNode(New_haxe__ds___EnumValueMap__EnumValueTreeNode(left.left, left.key, left.value, left.right.left, -1), left.right.key, left.right.value, New_haxe__ds___EnumValueMap__EnumValueTreeNode(left.right.right, key, value, right, -1), -1)
	}
	if rightHeight > int(int32((hxrt.Int32Wrap(leftHeight) + hxrt.Int32Wrap(2)))) {
		var hx_if_94 int
		if right.right == nil {
			hx_if_94 = 0
		} else {
			hx_if_94 = right.right.height
		}
		rightRightHeight := hx_if_94
		var hx_if_95 int
		if right.left == nil {
			hx_if_95 = 0
		} else {
			hx_if_95 = right.left.height
		}
		rightLeftHeight := hx_if_95
		if rightRightHeight > rightLeftHeight {
			return New_haxe__ds___EnumValueMap__EnumValueTreeNode(New_haxe__ds___EnumValueMap__EnumValueTreeNode(left, key, value, right.left, -1), right.key, right.value, right.right, -1)
		}
		return New_haxe__ds___EnumValueMap__EnumValueTreeNode(New_haxe__ds___EnumValueMap__EnumValueTreeNode(left, key, value, right.left.left, -1), right.left.key, right.left.value, New_haxe__ds___EnumValueMap__EnumValueTreeNode(right.left.right, right.key, right.value, right.right, -1), -1)
	}
	return New_haxe__ds___EnumValueMap__EnumValueTreeNode(left, key, value, right, int(int32((hxrt.Int32Wrap(func() int {
		var hx_if_96 int
		if leftHeight > rightHeight {
			hx_if_96 = leftHeight
		} else {
			hx_if_96 = rightHeight
		}
		return hx_if_96
	}()) + hxrt.Int32Wrap(1)))))
}

func (self *haxe__ds__EnumValueMap) getIMap(key any) any {
	return self.get(key)
}

func (self *haxe__ds__EnumValueMap) setIMap(key any, value any) {
	self.set(key, value)
}

func (self *haxe__ds__EnumValueMap) existsIMap(key any) bool {
	return func(hx_value_97 any) bool {
		if hx_value_97 == nil {
			var hx_zero_98 bool
			return hx_zero_98
		}
		return hx_value_97.(bool)
	}(self.exists(key))
}

func (self *haxe__ds__EnumValueMap) removeIMap(key any) bool {
	return func(hx_value_99 any) bool {
		if hx_value_99 == nil {
			var hx_zero_100 bool
			return hx_zero_100
		}
		return hx_value_99.(bool)
	}(self.remove(key))
}

func (self *haxe__ds__EnumValueMap) copyIMap() haxe__IMap {
	return func(hx_value_101 any) *haxe__ds__EnumValueMap {
		if hx_value_101 == nil {
			var hx_zero_102 *haxe__ds__EnumValueMap
			return hx_zero_102
		}
		return hx_value_101.(*haxe__ds__EnumValueMap)
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
