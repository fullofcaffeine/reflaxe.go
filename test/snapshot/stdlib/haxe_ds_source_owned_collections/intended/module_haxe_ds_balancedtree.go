package main

import "snapshot/hxrt"

type I_haxe__ds__TreeNode interface {
	toString() *string
}

type haxe__ds__TreeNode struct {
	__hx_this I_haxe__ds__TreeNode
	left      *haxe__ds__TreeNode
	right     *haxe__ds__TreeNode
	key       any
	value     any
	_height   int
}

func New_haxe__ds__TreeNode(l *haxe__ds__TreeNode, k any, v any, r *haxe__ds__TreeNode, h int) *haxe__ds__TreeNode {
	self := &haxe__ds__TreeNode{}
	self.__hx_this = self
	self.left = l
	self.key = k
	self.value = v
	self.right = r
	if h == -1 {
		self._height = int(int32((hxrt.Int32Wrap(func() int {
			var hx_if_9 int
			if func() int {
				_this := self.left
				var hx_if_5 int
				if _this == nil {
					hx_if_5 = 0
				} else {
					hx_if_5 = _this._height
				}
				return hx_if_5
			}() > func() int {
				_this_1 := self.right
				var hx_if_6 int
				if _this_1 == nil {
					hx_if_6 = 0
				} else {
					hx_if_6 = _this_1._height
				}
				return hx_if_6
			}() {
				_this_2 := self.left
				var hx_if_7 int
				if _this_2 == nil {
					hx_if_7 = 0
				} else {
					hx_if_7 = _this_2._height
				}
				hx_if_9 = hx_if_7
			} else {
				_this_3 := self.right
				var hx_if_8 int
				if _this_3 == nil {
					hx_if_8 = 0
				} else {
					hx_if_8 = _this_3._height
				}
				hx_if_9 = hx_if_8
			}
			return hx_if_9
		}()) + hxrt.Int32Wrap(1))))
	} else {
		self._height = h
	}
	return self
}

func (self *haxe__ds__TreeNode) toString() *string {
	return hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(func() *string {
		var hx_if_12 *string
		if self.left == nil {
			hx_if_12 = hxrt.StringFromLiteral("")
		} else {
			hx_if_12 = hxrt.StringConcatStringPtr(func(hx_value_10 any) *string {
				if hx_value_10 == nil {
					var hx_zero_11 *string
					return hx_zero_11
				}
				return hx_value_10.(*string)
			}(self.left.toString()), hxrt.StringFromLiteral(", "))
		}
		return hx_if_12
	}(), hxrt.StdString(self.key)), hxrt.StringFromLiteral(" => ")), hxrt.StdString(self.value)), func() *string {
		var hx_if_15 *string
		if self.right == nil {
			hx_if_15 = hxrt.StringFromLiteral("")
		} else {
			hx_if_15 = hxrt.StringConcatStringPtr(hxrt.StringFromLiteral(", "), func(hx_value_13 any) *string {
				if hx_value_13 == nil {
					var hx_zero_14 *string
					return hx_zero_14
				}
				return hx_value_13.(*string)
			}(self.right.toString()))
		}
		return hx_if_15
	}())
}

type I_haxe__ds__BalancedTree interface {
	set(key any, value any)
	get(key any) any
	remove(key any) bool
	exists(key any) bool
	iterator() map[string]any
	keys() map[string]any
	keyValueIterator() map[string]any
	copy() *haxe__ds__BalancedTree
	setLoop(k any, v any, node *haxe__ds__TreeNode) *haxe__ds__TreeNode
	removeLoop(k any, node *haxe__ds__TreeNode) *haxe__ds__TreeNode
	keysLoop(node *haxe__ds__TreeNode, acc *hxrt.Array) *hxrt.Array
	merge(t1 *haxe__ds__TreeNode, t2 *haxe__ds__TreeNode) *haxe__ds__TreeNode
	minBinding(t *haxe__ds__TreeNode) *haxe__ds__TreeNode
	removeMinBinding(t *haxe__ds__TreeNode) *haxe__ds__TreeNode
	balance(l *haxe__ds__TreeNode, k any, v any, r *haxe__ds__TreeNode) *haxe__ds__TreeNode
	compare(k1 any, k2 any) int
	toString() *string
	clear()
	getIMap(key any) any
	setIMap(key any, value any)
	existsIMap(key any) bool
	removeIMap(key any) bool
	copyIMap() haxe__IMap
}

type haxe__ds__BalancedTree struct {
	__hx_this I_haxe__ds__BalancedTree
	root      *haxe__ds__TreeNode
}

func New_haxe__ds__BalancedTree() *haxe__ds__BalancedTree {
	self := &haxe__ds__BalancedTree{}
	self.__hx_this = self
	return self
}

func (self *haxe__ds__BalancedTree) set(key any, value any) {
	self.root = func(hx_value_18 any) *haxe__ds__TreeNode {
		if hx_value_18 == nil {
			var hx_zero_19 *haxe__ds__TreeNode
			return hx_zero_19
		}
		return hx_value_18.(*haxe__ds__TreeNode)
	}(self.setLoop(key, value, self.root))
}

func (self *haxe__ds__BalancedTree) get(key any) any {
	node := self.root
	for node != nil {
		c := func(hx_value_20 any) int {
			if hx_value_20 == nil {
				var hx_zero_21 int
				return hx_zero_21
			}
			return hx_value_20.(int)
		}(self.compare(key, node.key))
		if c == 0 {
			return node.value
		}
		if c < 0 {
			node = node.left
		} else {
			node = node.right
		}
	}
	return nil
}

func (self *haxe__ds__BalancedTree) remove(key any) bool {
	removed := true
	hxrt.TryCatch(func() {
		self.root = func(hx_value_24 any) *haxe__ds__TreeNode {
			if hx_value_24 == nil {
				var hx_zero_25 *haxe__ds__TreeNode
				return hx_zero_25
			}
			return hx_value_24.(*haxe__ds__TreeNode)
		}(self.removeLoop(key, self.root))
	}, func(hx_caught_22 any) {
		switch hx_typed_23 := hx_caught_22.(type) {
		case *string:
			hx_tmp := hx_typed_23
			_ = hx_tmp
			removed = false
		default:
			hxrt.Throw(hx_caught_22)
		}
	})
	return removed
}

func (self *haxe__ds__BalancedTree) exists(key any) bool {
	node := self.root
	for node != nil {
		c := func(hx_value_26 any) int {
			if hx_value_26 == nil {
				var hx_zero_27 int
				return hx_zero_27
			}
			return hx_value_26.(int)
		}(self.compare(key, node.key))
		if c == 0 {
			return true
		} else {
			if c < 0 {
				node = node.left
			} else {
				node = node.right
			}
		}
	}
	return false
}

func (self *haxe__ds__BalancedTree) iterator() map[string]any {
	ret := haxe__ds__BalancedTree_iteratorLoop(self.root, hxrt.NewArray())
	index := 0
	hx_obj_28 := map[string]any{}
	hx_obj_28["hasNext"] = func() bool {
		return (index < ret.Len())
	}
	hx_obj_28["next"] = func() any {
		hx_post_29 := index
		index = int(int32((index + 1)))
		return ret.Get(hx_post_29)
	}
	return hx_obj_28
}

func (self *haxe__ds__BalancedTree) keys() map[string]any {
	ret := func(hx_value_30 any) *hxrt.Array {
		if hx_value_30 == nil {
			var hx_zero_31 *hxrt.Array
			return hx_zero_31
		}
		return hx_value_30.(*hxrt.Array)
	}(self.keysLoop(self.root, hxrt.NewArray()))
	index := 0
	hx_obj_32 := map[string]any{}
	hx_obj_32["hasNext"] = func() bool {
		return (index < ret.Len())
	}
	hx_obj_32["next"] = func() any {
		hx_post_33 := index
		index = int(int32((index + 1)))
		return ret.Get(hx_post_33)
	}
	return hx_obj_32
}

func (self *haxe__ds__BalancedTree) keyValueIterator() map[string]any {
	_gthis := self
	keyIterator := func(hx_value_34 any) map[string]any {
		if hx_value_34 == nil {
			var hx_zero_35 map[string]any
			return hx_zero_35
		}
		return hx_value_34.(map[string]any)
	}(self.keys())
	hx_obj_36 := map[string]any{}
	hx_obj_36["hasNext"] = func() bool {
		return func(hx_obj_37 map[string]any) func() bool {
			hx_field_38 := hx_obj_37["hasNext"]
			if hx_field_38 == nil {
				var hx_zero_39 func() bool
				return hx_zero_39
			}
			return hx_field_38.(func() bool)
		}(keyIterator)()
	}
	hx_obj_36["next"] = func() map[string]any {
		var key any = func(hx_obj_40 map[string]any) func() any {
			hx_field_41 := hx_obj_40["next"]
			if hx_field_41 == nil {
				var hx_zero_42 func() any
				return hx_zero_42
			}
			return hx_field_41.(func() any)
		}(keyIterator)()
		hx_obj_43 := map[string]any{}
		hx_obj_43["key"] = key
		hx_obj_43["value"] = _gthis.get(key)
		return hx_obj_43
	}
	return hx_obj_36
}

func (self *haxe__ds__BalancedTree) copy() *haxe__ds__BalancedTree {
	copied := New_haxe__ds__BalancedTree()
	copied.root = self.root
	return copied
}

func (self *haxe__ds__BalancedTree) setLoop(k any, v any, node *haxe__ds__TreeNode) *haxe__ds__TreeNode {
	if node == nil {
		return New_haxe__ds__TreeNode(nil, k, v, nil, -1)
	}
	c := func(hx_value_44 any) int {
		if hx_value_44 == nil {
			var hx_zero_45 int
			return hx_zero_45
		}
		return hx_value_44.(int)
	}(self.compare(k, node.key))
	var hx_if_56 *haxe__ds__TreeNode
	if c == 0 {
		hx_if_56 = New_haxe__ds__TreeNode(node.left, k, v, node.right, func() int {
			var hx_if_46 int
			if node == nil {
				hx_if_46 = 0
			} else {
				hx_if_46 = node._height
			}
			return hx_if_46
		}())
	} else {
		var hx_if_55 *haxe__ds__TreeNode
		if c < 0 {
			nl := func(hx_value_47 any) *haxe__ds__TreeNode {
				if hx_value_47 == nil {
					var hx_zero_48 *haxe__ds__TreeNode
					return hx_zero_48
				}
				return hx_value_47.(*haxe__ds__TreeNode)
			}(self.setLoop(k, v, node.left))
			hx_if_55 = func(hx_value_49 any) *haxe__ds__TreeNode {
				if hx_value_49 == nil {
					var hx_zero_50 *haxe__ds__TreeNode
					return hx_zero_50
				}
				return hx_value_49.(*haxe__ds__TreeNode)
			}(self.balance(nl, node.key, node.value, node.right))
		} else {
			nr := func(hx_value_51 any) *haxe__ds__TreeNode {
				if hx_value_51 == nil {
					var hx_zero_52 *haxe__ds__TreeNode
					return hx_zero_52
				}
				return hx_value_51.(*haxe__ds__TreeNode)
			}(self.setLoop(k, v, node.right))
			hx_if_55 = func(hx_value_53 any) *haxe__ds__TreeNode {
				if hx_value_53 == nil {
					var hx_zero_54 *haxe__ds__TreeNode
					return hx_zero_54
				}
				return hx_value_53.(*haxe__ds__TreeNode)
			}(self.balance(node.left, node.key, node.value, nr))
		}
		hx_if_56 = hx_if_55
	}
	return hx_if_56
}

func (self *haxe__ds__BalancedTree) removeLoop(k any, node *haxe__ds__TreeNode) *haxe__ds__TreeNode {
	if node == nil {
		hxrt.Throw(hxrt.StringFromLiteral("Not_found"))
	}
	c := func(hx_value_57 any) int {
		if hx_value_57 == nil {
			var hx_zero_58 int
			return hx_zero_58
		}
		return hx_value_57.(int)
	}(self.compare(k, node.key))
	var hx_if_70 *haxe__ds__TreeNode
	if c == 0 {
		hx_if_70 = func(hx_value_59 any) *haxe__ds__TreeNode {
			if hx_value_59 == nil {
				var hx_zero_60 *haxe__ds__TreeNode
				return hx_zero_60
			}
			return hx_value_59.(*haxe__ds__TreeNode)
		}(self.merge(node.left, node.right))
	} else {
		var hx_if_69 *haxe__ds__TreeNode
		if c < 0 {
			hx_if_69 = func(hx_value_63 any) *haxe__ds__TreeNode {
				if hx_value_63 == nil {
					var hx_zero_64 *haxe__ds__TreeNode
					return hx_zero_64
				}
				return hx_value_63.(*haxe__ds__TreeNode)
			}(self.balance(func(hx_value_61 any) *haxe__ds__TreeNode {
				if hx_value_61 == nil {
					var hx_zero_62 *haxe__ds__TreeNode
					return hx_zero_62
				}
				return hx_value_61.(*haxe__ds__TreeNode)
			}(self.removeLoop(k, node.left)), node.key, node.value, node.right))
		} else {
			hx_if_69 = func(hx_value_67 any) *haxe__ds__TreeNode {
				if hx_value_67 == nil {
					var hx_zero_68 *haxe__ds__TreeNode
					return hx_zero_68
				}
				return hx_value_67.(*haxe__ds__TreeNode)
			}(self.balance(node.left, node.key, node.value, func(hx_value_65 any) *haxe__ds__TreeNode {
				if hx_value_65 == nil {
					var hx_zero_66 *haxe__ds__TreeNode
					return hx_zero_66
				}
				return hx_value_65.(*haxe__ds__TreeNode)
			}(self.removeLoop(k, node.right))))
		}
		hx_if_70 = hx_if_69
	}
	return hx_if_70
}

func (self *haxe__ds__BalancedTree) keysLoop(node *haxe__ds__TreeNode, acc *hxrt.Array) *hxrt.Array {
	if node != nil {
		acc = func(hx_value_71 any) *hxrt.Array {
			if hx_value_71 == nil {
				var hx_zero_72 *hxrt.Array
				return hx_zero_72
			}
			return hx_value_71.(*hxrt.Array)
		}(self.keysLoop(node.left, acc))
		acc.Push(node.key)
		acc = func(hx_value_74 any) *hxrt.Array {
			if hx_value_74 == nil {
				var hx_zero_75 *hxrt.Array
				return hx_zero_75
			}
			return hx_value_74.(*hxrt.Array)
		}(self.keysLoop(node.right, acc))
	}
	return acc
}

func (self *haxe__ds__BalancedTree) merge(t1 *haxe__ds__TreeNode, t2 *haxe__ds__TreeNode) *haxe__ds__TreeNode {
	if t1 == nil {
		return t2
	}
	if t2 == nil {
		return t1
	}
	t := func(hx_value_76 any) *haxe__ds__TreeNode {
		if hx_value_76 == nil {
			var hx_zero_77 *haxe__ds__TreeNode
			return hx_zero_77
		}
		return hx_value_76.(*haxe__ds__TreeNode)
	}(self.minBinding(t2))
	return func(hx_value_80 any) *haxe__ds__TreeNode {
		if hx_value_80 == nil {
			var hx_zero_81 *haxe__ds__TreeNode
			return hx_zero_81
		}
		return hx_value_80.(*haxe__ds__TreeNode)
	}(self.balance(t1, t.key, t.value, func(hx_value_78 any) *haxe__ds__TreeNode {
		if hx_value_78 == nil {
			var hx_zero_79 *haxe__ds__TreeNode
			return hx_zero_79
		}
		return hx_value_78.(*haxe__ds__TreeNode)
	}(self.removeMinBinding(t2))))
}

func (self *haxe__ds__BalancedTree) minBinding(t *haxe__ds__TreeNode) *haxe__ds__TreeNode {
	var hx_if_86 *haxe__ds__TreeNode
	if t == nil {
		hx_if_86 = func() *haxe__ds__TreeNode {
			hxrt.Throw(hxrt.StringFromLiteral("Not_found"))
			var hx_throw_zero_82 *haxe__ds__TreeNode
			return hx_throw_zero_82
		}()
	} else {
		var hx_if_85 *haxe__ds__TreeNode
		if t.left == nil {
			hx_if_85 = t
		} else {
			hx_if_85 = func(hx_value_83 any) *haxe__ds__TreeNode {
				if hx_value_83 == nil {
					var hx_zero_84 *haxe__ds__TreeNode
					return hx_zero_84
				}
				return hx_value_83.(*haxe__ds__TreeNode)
			}(self.minBinding(t.left))
		}
		hx_if_86 = hx_if_85
	}
	return hx_if_86
}

func (self *haxe__ds__BalancedTree) removeMinBinding(t *haxe__ds__TreeNode) *haxe__ds__TreeNode {
	var hx_if_91 *haxe__ds__TreeNode
	if t.left == nil {
		hx_if_91 = t.right
	} else {
		hx_if_91 = func(hx_value_89 any) *haxe__ds__TreeNode {
			if hx_value_89 == nil {
				var hx_zero_90 *haxe__ds__TreeNode
				return hx_zero_90
			}
			return hx_value_89.(*haxe__ds__TreeNode)
		}(self.balance(func(hx_value_87 any) *haxe__ds__TreeNode {
			if hx_value_87 == nil {
				var hx_zero_88 *haxe__ds__TreeNode
				return hx_zero_88
			}
			return hx_value_87.(*haxe__ds__TreeNode)
		}(self.removeMinBinding(t.left)), t.key, t.value, t.right))
	}
	return hx_if_91
}

func (self *haxe__ds__BalancedTree) balance(l *haxe__ds__TreeNode, k any, v any, r *haxe__ds__TreeNode) *haxe__ds__TreeNode {
	var hx_if_92 int
	if l == nil {
		hx_if_92 = 0
	} else {
		hx_if_92 = l._height
	}
	hl := hx_if_92
	var hx_if_93 int
	if r == nil {
		hx_if_93 = 0
	} else {
		hx_if_93 = r._height
	}
	hr := hx_if_93
	var hx_if_102 *haxe__ds__TreeNode
	if hl > int(int32((hxrt.Int32Wrap(hr) + hxrt.Int32Wrap(2)))) {
		var hx_if_96 *haxe__ds__TreeNode
		if func() int {
			_this := l.left
			var hx_if_94 int
			if _this == nil {
				hx_if_94 = 0
			} else {
				hx_if_94 = _this._height
			}
			return hx_if_94
		}() >= func() int {
			_this_1 := l.right
			var hx_if_95 int
			if _this_1 == nil {
				hx_if_95 = 0
			} else {
				hx_if_95 = _this_1._height
			}
			return hx_if_95
		}() {
			hx_if_96 = New_haxe__ds__TreeNode(l.left, l.key, l.value, New_haxe__ds__TreeNode(l.right, k, v, r, -1), -1)
		} else {
			hx_if_96 = New_haxe__ds__TreeNode(New_haxe__ds__TreeNode(l.left, l.key, l.value, l.right.left, -1), l.right.key, l.right.value, New_haxe__ds__TreeNode(l.right.right, k, v, r, -1), -1)
		}
		hx_if_102 = hx_if_96
	} else {
		var hx_if_101 *haxe__ds__TreeNode
		if hr > int(int32((hxrt.Int32Wrap(hl) + hxrt.Int32Wrap(2)))) {
			var hx_if_99 *haxe__ds__TreeNode
			if func() int {
				_this_2 := r.right
				var hx_if_97 int
				if _this_2 == nil {
					hx_if_97 = 0
				} else {
					hx_if_97 = _this_2._height
				}
				return hx_if_97
			}() > func() int {
				_this_3 := r.left
				var hx_if_98 int
				if _this_3 == nil {
					hx_if_98 = 0
				} else {
					hx_if_98 = _this_3._height
				}
				return hx_if_98
			}() {
				hx_if_99 = New_haxe__ds__TreeNode(New_haxe__ds__TreeNode(l, k, v, r.left, -1), r.key, r.value, r.right, -1)
			} else {
				hx_if_99 = New_haxe__ds__TreeNode(New_haxe__ds__TreeNode(l, k, v, r.left.left, -1), r.left.key, r.left.value, New_haxe__ds__TreeNode(r.left.right, r.key, r.value, r.right, -1), -1)
			}
			hx_if_101 = hx_if_99
		} else {
			hx_if_101 = New_haxe__ds__TreeNode(l, k, v, r, int(int32((hxrt.Int32Wrap(func() int {
				var hx_if_100 int
				if hl > hr {
					hx_if_100 = hl
				} else {
					hx_if_100 = hr
				}
				return hx_if_100
			}()) + hxrt.Int32Wrap(1)))))
		}
		hx_if_102 = hx_if_101
	}
	return hx_if_102
}

func (self *haxe__ds__BalancedTree) compare(k1 any, k2 any) int {
	return Reflect_compare(k1, k2)
}

func (self *haxe__ds__BalancedTree) toString() *string {
	var hx_if_105 *string
	if self.root == nil {
		hx_if_105 = hxrt.StringFromLiteral("[]")
	} else {
		hx_if_105 = hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("["), func(hx_value_103 any) *string {
			if hx_value_103 == nil {
				var hx_zero_104 *string
				return hx_zero_104
			}
			return hx_value_103.(*string)
		}(self.root.toString())), hxrt.StringFromLiteral("]"))
	}
	return hx_if_105
}

func (self *haxe__ds__BalancedTree) clear() {
	self.root = nil
}

func (self *haxe__ds__BalancedTree) getIMap(key any) any {
	return self.get(key)
}

func (self *haxe__ds__BalancedTree) setIMap(key any, value any) {
	self.set(key, value)
}

func (self *haxe__ds__BalancedTree) existsIMap(key any) bool {
	return func(hx_value_106 any) bool {
		if hx_value_106 == nil {
			var hx_zero_107 bool
			return hx_zero_107
		}
		return hx_value_106.(bool)
	}(self.exists(key))
}

func (self *haxe__ds__BalancedTree) removeIMap(key any) bool {
	return func(hx_value_108 any) bool {
		if hx_value_108 == nil {
			var hx_zero_109 bool
			return hx_zero_109
		}
		return hx_value_108.(bool)
	}(self.remove(key))
}

func (self *haxe__ds__BalancedTree) copyIMap() haxe__IMap {
	return func(hx_value_110 any) *haxe__ds__BalancedTree {
		if hx_value_110 == nil {
			var hx_zero_111 *haxe__ds__BalancedTree
			return hx_zero_111
		}
		return hx_value_110.(*haxe__ds__BalancedTree)
	}(self.copy())
}

func haxe__ds__BalancedTree_iteratorLoop(node *haxe__ds__TreeNode, acc *hxrt.Array) *hxrt.Array {
	if node != nil {
		acc = haxe__ds__BalancedTree_iteratorLoop(node.left, acc)
		acc.Push(node.value)
		acc = haxe__ds__BalancedTree_iteratorLoop(node.right, acc)
	}
	return acc
}
