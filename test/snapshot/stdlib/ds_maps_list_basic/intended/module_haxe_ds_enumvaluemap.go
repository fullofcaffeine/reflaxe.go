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

func (self *haxe__ds___EnumValueMap__EnumValueTreeNode) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_haxe__ds___EnumValueMap__EnumValueTreeNode(left *haxe__ds___EnumValueMap__EnumValueTreeNode, key any, value any, right *haxe__ds___EnumValueMap__EnumValueTreeNode, height int) *haxe__ds___EnumValueMap__EnumValueTreeNode {
	return hx_init_haxe__ds___EnumValueMap__EnumValueTreeNode(hx_alloc_haxe__ds___EnumValueMap__EnumValueTreeNode(nil), left, key, value, right, height)
}

func hx_alloc_haxe__ds___EnumValueMap__EnumValueTreeNode(owner any) *haxe__ds___EnumValueMap__EnumValueTreeNode {
	self := &haxe__ds___EnumValueMap__EnumValueTreeNode{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_haxe__ds___EnumValueMap__EnumValueTreeNode)
	}
	return self
}

func hx_init_haxe__ds___EnumValueMap__EnumValueTreeNode(self *haxe__ds___EnumValueMap__EnumValueTreeNode, left *haxe__ds___EnumValueMap__EnumValueTreeNode, key any, value any, right *haxe__ds___EnumValueMap__EnumValueTreeNode, height int) *haxe__ds___EnumValueMap__EnumValueTreeNode {
	self.left = left
	self.key = key
	self.value = value
	self.right = right
	if height == -1 {
		var hx_if_1 int
		if left == nil {
			hx_if_1 = 0
		} else {
			hx_if_1 = left.height
		}
		leftHeight := hx_if_1
		var hx_if_2 int
		if right == nil {
			hx_if_2 = 0
		} else {
			hx_if_2 = right.height
		}
		rightHeight := hx_if_2
		self.height = int((hxrt.Int32Wrap(func() int {
			var hx_if_3 int
			if leftHeight > rightHeight {
				hx_if_3 = leftHeight
			} else {
				hx_if_3 = rightHeight
			}
			return hx_if_3
		}()) + hxrt.Int32Wrap(1)))
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
		var hx_if_6 *string
		if self.left == nil {
			hx_if_6 = hxrt.StringFromLiteral("")
		} else {
			hx_if_6 = hxrt.StringConcatStringPtr(func(hx_value_4 any) *string {
				if hx_value_4 == nil {
					var hx_zero_5 *string
					return hx_zero_5
				}
				return hx_value_4.(*string)
			}(self.left.__hx_this.toString()), hxrt.StringFromLiteral(", "))
		}
		return hx_if_6
	}(), hxrt.StdString(self.key)), hxrt.StringFromLiteral(" => ")), hxrt.StdString(self.value)), func() *string {
		var hx_if_9 *string
		if self.right == nil {
			hx_if_9 = hxrt.StringFromLiteral("")
		} else {
			hx_if_9 = hxrt.StringConcatStringPtr(hxrt.StringFromLiteral(", "), func(hx_value_7 any) *string {
				if hx_value_7 == nil {
					var hx_zero_8 *string
					return hx_zero_8
				}
				return hx_value_7.(*string)
			}(self.right.__hx_this.toString()))
		}
		return hx_if_9
	}())
}

func (self *haxe__ds___EnumValueMap__EnumValueTreeNode) String() string {
	return *self.__hx_this.toString()
}

type I_haxe__ds__EnumValueMap interface {
	set(key any, value any)
	get(key any) any
	exists(key any) bool
	remove(key any) bool
	keys() hxrt.Anonymous
	iterator() hxrt.Anonymous
	keyValueIterator() hxrt.Anonymous
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

func (self *haxe__ds__EnumValueMap) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_haxe__ds__EnumValueMap() *haxe__ds__EnumValueMap {
	return hx_init_haxe__ds__EnumValueMap(hx_alloc_haxe__ds__EnumValueMap(nil))
}

func hx_alloc_haxe__ds__EnumValueMap(owner any) *haxe__ds__EnumValueMap {
	self := &haxe__ds__EnumValueMap{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_haxe__ds__EnumValueMap)
	}
	return self
}

func hx_init_haxe__ds__EnumValueMap(self *haxe__ds__EnumValueMap) *haxe__ds__EnumValueMap {
	return self
}

func (self *haxe__ds__EnumValueMap) set(key any, value any) {
	self.root = hx_recover_class_3(self.__hx_this.setLoop(key, value, self.root))
}

func (self *haxe__ds__EnumValueMap) get(key any) any {
	node := self.root
	for node != nil {
		result := func(hx_value_10 any) int {
			if hx_value_10 == nil {
				var hx_zero_11 int
				return hx_zero_11
			}
			return hx_value_10.(int)
		}(self.__hx_this.compare(key, node.key))
		if result == 0 {
			return node.value
		}
		var hx_if_12 *haxe__ds___EnumValueMap__EnumValueTreeNode
		if result < 0 {
			hx_if_12 = node.left
		} else {
			hx_if_12 = node.right
		}
		node = hx_if_12
	}
	return nil
}

func (self *haxe__ds__EnumValueMap) exists(key any) bool {
	node := self.root
	for node != nil {
		result := func(hx_value_13 any) int {
			if hx_value_13 == nil {
				var hx_zero_14 int
				return hx_zero_14
			}
			return hx_value_13.(int)
		}(self.__hx_this.compare(key, node.key))
		if result == 0 {
			return true
		}
		var hx_if_15 *haxe__ds___EnumValueMap__EnumValueTreeNode
		if result < 0 {
			hx_if_15 = node.left
		} else {
			hx_if_15 = node.right
		}
		node = hx_if_15
	}
	return false
}

func (self *haxe__ds__EnumValueMap) remove(key any) bool {
	removed := true
	hxrt.TryCatch(func() {
		self.root = hx_recover_class_3(self.__hx_this.removeLoop(key, self.root))
	}, func(hx_caught_16 any) {
		hx_exception_17 := haxe__Exception_caught(hx_caught_16)
		hx_payload_18 := haxe__Exception_catchValue(hx_caught_16)
		_ = hx_exception_17
		if func(hx_value any) bool {
			switch hx_value.(type) {
			case *string:
				return true
			default:
				return false
			}
		}(hx_payload_18) {
			var hx_tmp *string = func(hx_value_19 any) *string {
				if hx_value_19 == nil {
					var hx_zero_20 *string
					return hx_zero_20
				}
				return hx_value_19.(*string)
			}(hx_payload_18)
			_ = hx_tmp
			removed = false
		} else {
			hxrt.Throw(hx_caught_16)
		}
	})
	return removed
}

func (self *haxe__ds__EnumValueMap) keys() hxrt.Anonymous {
	values := haxe__ds__EnumValueMap_keysLoop(self.root, hxrt.NewArray())
	index := 0
	hx_obj_21 := map[string]any{}
	hx_obj_21["hasNext"] = hxrt.NewFunction(func() bool {
		return (index < values.Len())
	})
	hx_obj_21["next"] = hxrt.NewFunction(func() any {
		hx_post_22 := index
		index = int(int32((index + 1)))
		return values.Get(hx_post_22)
	})
	return hxrt.AnonymousFromMap(hx_obj_21)
}

func (self *haxe__ds__EnumValueMap) iterator() hxrt.Anonymous {
	values := haxe__ds__EnumValueMap_valuesLoop(self.root, hxrt.NewArray())
	index := 0
	hx_obj_23 := map[string]any{}
	hx_obj_23["hasNext"] = hxrt.NewFunction(func() bool {
		return (index < values.Len())
	})
	hx_obj_23["next"] = hxrt.NewFunction(func() any {
		hx_post_24 := index
		index = int(int32((index + 1)))
		return values.Get(hx_post_24)
	})
	return hxrt.AnonymousFromMap(hx_obj_23)
}

func (self *haxe__ds__EnumValueMap) keyValueIterator() hxrt.Anonymous {
	_gthis := self
	keys := self.__hx_this.keys()
	hx_obj_25 := map[string]any{}
	hx_obj_25["hasNext"] = hxrt.NewFunction(func() bool {
		return func() bool {
			var hx_callable_26 *hxrt.Function[func() bool] = hx_recover_callable_0(hxrt.AnonymousGet(keys, "hasNext"))
			return hxrt.FunctionCall(hx_callable_26)()
		}()
	})
	hx_obj_25["next"] = hxrt.NewFunction(func() hxrt.Anonymous {
		var key any = func() any {
			var hx_callable_27 *hxrt.Function[func() any] = hx_recover_callable_1(hxrt.AnonymousGet(keys, "next"))
			return hxrt.FunctionCall(hx_callable_27)()
		}()
		hx_obj_28 := map[string]any{}
		hx_obj_28["key"] = key
		hx_obj_28["value"] = _gthis.__hx_this.get(key)
		return hxrt.AnonymousFromMap(hx_obj_28)
	})
	return hxrt.AnonymousFromMap(hx_obj_25)
}

func (self *haxe__ds__EnumValueMap) copy() *haxe__ds__EnumValueMap {
	copied := New_haxe__ds__EnumValueMap()
	copied.root = self.root
	return copied
}

func (self *haxe__ds__EnumValueMap) toString() *string {
	var hx_if_31 *string
	if self.root == nil {
		hx_if_31 = hxrt.StringFromLiteral("[]")
	} else {
		hx_if_31 = hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("["), func(hx_value_29 any) *string {
			if hx_value_29 == nil {
				var hx_zero_30 *string
				return hx_zero_30
			}
			return hx_value_29.(*string)
		}(self.root.__hx_this.toString())), hxrt.StringFromLiteral("]"))
	}
	return hx_if_31
}

func (self *haxe__ds__EnumValueMap) clear() {
	self.root = nil
}

func (self *haxe__ds__EnumValueMap) compare(left any, right any) int {
	result := int((hxrt.Int32Wrap(Type_enumIndex(left)) - hxrt.Int32Wrap(Type_enumIndex(right))))
	if result != 0 {
		return result
	}
	return func(hx_value_32 any) int {
		if hx_value_32 == nil {
			var hx_zero_33 int
			return hx_zero_33
		}
		return hx_value_32.(int)
	}(self.__hx_this.compareArgs(Type_enumParameters(left), Type_enumParameters(right)))
}

func (self *haxe__ds__EnumValueMap) compareArgs(left *hxrt.Array, right *hxrt.Array) int {
	result := int((hxrt.Int32Wrap(left.Len()) - hxrt.Int32Wrap(right.Len())))
	if result != 0 {
		return result
	}
	_g := 0
	_g1 := left.Len()
	for _g < _g1 {
		hx_post_34 := _g
		_g = int(int32((_g + 1)))
		index := hx_post_34
		result = func(hx_value_35 any) int {
			if hx_value_35 == nil {
				var hx_zero_36 int
				return hx_zero_36
			}
			return hx_value_35.(int)
		}(self.__hx_this.compareArg(left.Get(index), right.Get(index)))
		if result != 0 {
			return result
		}
	}
	return 0
}

func (self *haxe__ds__EnumValueMap) compareArg(left any, right any) int {
	if haxe__ds__EnumValueMap_isEnumValue(left) && haxe__ds__EnumValueMap_isEnumValue(right) {
		return func(hx_value_37 any) int {
			if hx_value_37 == nil {
				var hx_zero_38 int
				return hx_zero_38
			}
			return hx_value_37.(int)
		}(self.__hx_this.compare(left, right))
	}
	if func(hx_value any) bool {
		switch hx_value.(type) {
		case *hxrt.Array:
			return true
		default:
			return false
		}
	}(left) && func(hx_value any) bool {
		switch hx_value.(type) {
		case *hxrt.Array:
			return true
		default:
			return false
		}
	}(right) {
		return func(hx_value_47 any) int {
			if hx_value_47 == nil {
				var hx_zero_48 int
				return hx_zero_48
			}
			return hx_value_47.(int)
		}(self.__hx_this.compareArgs(func(hx_value_41 any) *hxrt.Array {
			if hx_value_41 == nil {
				var hx_zero_42 *hxrt.Array
				return hx_zero_42
			}
			return hx_value_41.(*hxrt.Array)
		}(any(func(hx_value_39 any) *hxrt.Array {
			if hx_value_39 == nil {
				var hx_zero_40 *hxrt.Array
				return hx_zero_40
			}
			return hx_value_39.(*hxrt.Array)
		}(left))), func(hx_value_45 any) *hxrt.Array {
			if hx_value_45 == nil {
				var hx_zero_46 *hxrt.Array
				return hx_zero_46
			}
			return hx_value_45.(*hxrt.Array)
		}(any(func(hx_value_43 any) *hxrt.Array {
			if hx_value_43 == nil {
				var hx_zero_44 *hxrt.Array
				return hx_zero_44
			}
			return hx_value_43.(*hxrt.Array)
		}(right)))))
	}
	return Reflect_compare(left, right)
}

func (self *haxe__ds__EnumValueMap) setLoop(key any, value any, node *haxe__ds___EnumValueMap__EnumValueTreeNode) *haxe__ds___EnumValueMap__EnumValueTreeNode {
	if node == nil {
		return New_haxe__ds___EnumValueMap__EnumValueTreeNode(nil, key, value, nil, -1)
	}
	result := func(hx_value_49 any) int {
		if hx_value_49 == nil {
			var hx_zero_50 int
			return hx_zero_50
		}
		return hx_value_49.(int)
	}(self.__hx_this.compare(key, node.key))
	if result == 0 {
		return New_haxe__ds___EnumValueMap__EnumValueTreeNode(node.left, key, value, node.right, node.height)
	}
	if result < 0 {
		return hx_recover_class_3(self.__hx_this.balance(hx_recover_class_3(self.__hx_this.setLoop(key, value, node.left)), node.key, node.value, node.right))
	}
	return hx_recover_class_3(self.__hx_this.balance(node.left, node.key, node.value, hx_recover_class_3(self.__hx_this.setLoop(key, value, node.right))))
}

func (self *haxe__ds__EnumValueMap) removeLoop(key any, node *haxe__ds___EnumValueMap__EnumValueTreeNode) *haxe__ds___EnumValueMap__EnumValueTreeNode {
	if node == nil {
		hxrt.Throw(haxe__Exception_thrown(hxrt.StringFromLiteral("Not_found")))
	}
	result := func(hx_value_51 any) int {
		if hx_value_51 == nil {
			var hx_zero_52 int
			return hx_zero_52
		}
		return hx_value_51.(int)
	}(self.__hx_this.compare(key, node.key))
	if result == 0 {
		return hx_recover_class_3(self.__hx_this.merge(node.left, node.right))
	}
	if result < 0 {
		return hx_recover_class_3(self.__hx_this.balance(hx_recover_class_3(self.__hx_this.removeLoop(key, node.left)), node.key, node.value, node.right))
	}
	return hx_recover_class_3(self.__hx_this.balance(node.left, node.key, node.value, hx_recover_class_3(self.__hx_this.removeLoop(key, node.right))))
}

func (self *haxe__ds__EnumValueMap) merge(left *haxe__ds___EnumValueMap__EnumValueTreeNode, right *haxe__ds___EnumValueMap__EnumValueTreeNode) *haxe__ds___EnumValueMap__EnumValueTreeNode {
	if left == nil {
		return right
	}
	if right == nil {
		return left
	}
	minimum := hx_recover_class_3(self.__hx_this.minBinding(right))
	return hx_recover_class_3(self.__hx_this.balance(left, minimum.key, minimum.value, hx_recover_class_3(self.__hx_this.removeMinBinding(right))))
}

func (self *haxe__ds__EnumValueMap) minBinding(node *haxe__ds___EnumValueMap__EnumValueTreeNode) *haxe__ds___EnumValueMap__EnumValueTreeNode {
	if node == nil {
		hxrt.Throw(haxe__Exception_thrown(hxrt.StringFromLiteral("Not_found")))
	}
	var hx_if_53 *haxe__ds___EnumValueMap__EnumValueTreeNode
	if node.left == nil {
		hx_if_53 = node
	} else {
		hx_if_53 = hx_recover_class_3(self.__hx_this.minBinding(node.left))
	}
	return hx_if_53
}

func (self *haxe__ds__EnumValueMap) removeMinBinding(node *haxe__ds___EnumValueMap__EnumValueTreeNode) *haxe__ds___EnumValueMap__EnumValueTreeNode {
	var hx_if_54 *haxe__ds___EnumValueMap__EnumValueTreeNode
	if node.left == nil {
		hx_if_54 = node.right
	} else {
		hx_if_54 = hx_recover_class_3(self.__hx_this.balance(hx_recover_class_3(self.__hx_this.removeMinBinding(node.left)), node.key, node.value, node.right))
	}
	return hx_if_54
}

func (self *haxe__ds__EnumValueMap) balance(left *haxe__ds___EnumValueMap__EnumValueTreeNode, key any, value any, right *haxe__ds___EnumValueMap__EnumValueTreeNode) *haxe__ds___EnumValueMap__EnumValueTreeNode {
	var hx_if_55 int
	if left == nil {
		hx_if_55 = 0
	} else {
		hx_if_55 = left.height
	}
	leftHeight := hx_if_55
	var hx_if_56 int
	if right == nil {
		hx_if_56 = 0
	} else {
		hx_if_56 = right.height
	}
	rightHeight := hx_if_56
	if leftHeight > int((hxrt.Int32Wrap(rightHeight) + hxrt.Int32Wrap(2))) {
		var hx_if_57 int
		if left.left == nil {
			hx_if_57 = 0
		} else {
			hx_if_57 = left.left.height
		}
		leftLeftHeight := hx_if_57
		var hx_if_58 int
		if left.right == nil {
			hx_if_58 = 0
		} else {
			hx_if_58 = left.right.height
		}
		leftRightHeight := hx_if_58
		if leftLeftHeight >= leftRightHeight {
			return New_haxe__ds___EnumValueMap__EnumValueTreeNode(left.left, left.key, left.value, New_haxe__ds___EnumValueMap__EnumValueTreeNode(left.right, key, value, right, -1), -1)
		}
		return New_haxe__ds___EnumValueMap__EnumValueTreeNode(New_haxe__ds___EnumValueMap__EnumValueTreeNode(left.left, left.key, left.value, left.right.left, -1), left.right.key, left.right.value, New_haxe__ds___EnumValueMap__EnumValueTreeNode(left.right.right, key, value, right, -1), -1)
	}
	if rightHeight > int((hxrt.Int32Wrap(leftHeight) + hxrt.Int32Wrap(2))) {
		var hx_if_59 int
		if right.right == nil {
			hx_if_59 = 0
		} else {
			hx_if_59 = right.right.height
		}
		rightRightHeight := hx_if_59
		var hx_if_60 int
		if right.left == nil {
			hx_if_60 = 0
		} else {
			hx_if_60 = right.left.height
		}
		rightLeftHeight := hx_if_60
		if rightRightHeight > rightLeftHeight {
			return New_haxe__ds___EnumValueMap__EnumValueTreeNode(New_haxe__ds___EnumValueMap__EnumValueTreeNode(left, key, value, right.left, -1), right.key, right.value, right.right, -1)
		}
		return New_haxe__ds___EnumValueMap__EnumValueTreeNode(New_haxe__ds___EnumValueMap__EnumValueTreeNode(left, key, value, right.left.left, -1), right.left.key, right.left.value, New_haxe__ds___EnumValueMap__EnumValueTreeNode(right.left.right, right.key, right.value, right.right, -1), -1)
	}
	return New_haxe__ds___EnumValueMap__EnumValueTreeNode(left, key, value, right, int((hxrt.Int32Wrap(func() int {
		var hx_if_61 int
		if leftHeight > rightHeight {
			hx_if_61 = leftHeight
		} else {
			hx_if_61 = rightHeight
		}
		return hx_if_61
	}()) + hxrt.Int32Wrap(1))))
}

func (self *haxe__ds__EnumValueMap) getIMap(key any) any {
	return self.__hx_this.get(key)
}

func (self *haxe__ds__EnumValueMap) setIMap(key any, value any) {
	self.__hx_this.set(key, value)
}

func (self *haxe__ds__EnumValueMap) existsIMap(key any) bool {
	return func(hx_value_62 any) bool {
		if hx_value_62 == nil {
			var hx_zero_63 bool
			return hx_zero_63
		}
		return hx_value_62.(bool)
	}(self.__hx_this.exists(key))
}

func (self *haxe__ds__EnumValueMap) removeIMap(key any) bool {
	return func(hx_value_64 any) bool {
		if hx_value_64 == nil {
			var hx_zero_65 bool
			return hx_zero_65
		}
		return hx_value_64.(bool)
	}(self.__hx_this.remove(key))
}

func (self *haxe__ds__EnumValueMap) copyIMap() haxe__IMap {
	return hx_recover_class_4(self.__hx_this.copy())
}

func (self *haxe__ds__EnumValueMap) String() string {
	return *self.__hx_this.toString()
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
