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
			var hx_if_13 int
			if func() int {
				_this := self.left
				var hx_if_9 int
				if _this == nil {
					hx_if_9 = 0
				} else {
					hx_if_9 = _this._height
				}
				return hx_if_9
			}() > func() int {
				_this_1 := self.right
				var hx_if_10 int
				if _this_1 == nil {
					hx_if_10 = 0
				} else {
					hx_if_10 = _this_1._height
				}
				return hx_if_10
			}() {
				_this_2 := self.left
				var hx_if_11 int
				if _this_2 == nil {
					hx_if_11 = 0
				} else {
					hx_if_11 = _this_2._height
				}
				hx_if_13 = hx_if_11
			} else {
				_this_3 := self.right
				var hx_if_12 int
				if _this_3 == nil {
					hx_if_12 = 0
				} else {
					hx_if_12 = _this_3._height
				}
				hx_if_13 = hx_if_12
			}
			return hx_if_13
		}()) + hxrt.Int32Wrap(1))))
	} else {
		self._height = h
	}
	return self
}

func (self *haxe__ds__TreeNode) toString() *string {
	return hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(func() *string {
		var hx_if_16 *string
		if self.left == nil {
			hx_if_16 = hxrt.StringFromLiteral("")
		} else {
			hx_if_16 = hxrt.StringConcatStringPtr(func(hx_value_14 any) *string {
				if hx_value_14 == nil {
					var hx_zero_15 *string
					return hx_zero_15
				}
				return hx_value_14.(*string)
			}(self.left.toString()), hxrt.StringFromLiteral(", "))
		}
		return hx_if_16
	}(), hxrt.StdString(self.key)), hxrt.StringFromLiteral(" => ")), hxrt.StdString(self.value)), func() *string {
		var hx_if_19 *string
		if self.right == nil {
			hx_if_19 = hxrt.StringFromLiteral("")
		} else {
			hx_if_19 = hxrt.StringConcatStringPtr(hxrt.StringFromLiteral(", "), func(hx_value_17 any) *string {
				if hx_value_17 == nil {
					var hx_zero_18 *string
					return hx_zero_18
				}
				return hx_value_17.(*string)
			}(self.right.toString()))
		}
		return hx_if_19
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
	keysLoop(node *haxe__ds__TreeNode, acc []any) []any
	merge(t1 *haxe__ds__TreeNode, t2 *haxe__ds__TreeNode) *haxe__ds__TreeNode
	minBinding(t *haxe__ds__TreeNode) *haxe__ds__TreeNode
	removeMinBinding(t *haxe__ds__TreeNode) *haxe__ds__TreeNode
	balance(l *haxe__ds__TreeNode, k any, v any, r *haxe__ds__TreeNode) *haxe__ds__TreeNode
	compare(k1 any, k2 any) int
	toString() *string
	clear()
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
	self.root = func(hx_value_20 any) *haxe__ds__TreeNode {
		if hx_value_20 == nil {
			var hx_zero_21 *haxe__ds__TreeNode
			return hx_zero_21
		}
		return hx_value_20.(*haxe__ds__TreeNode)
	}(self.setLoop(key, value, self.root))
}

func (self *haxe__ds__BalancedTree) get(key any) any {
	node := self.root
	for node != nil {
		c := func(hx_value_22 any) int {
			if hx_value_22 == nil {
				var hx_zero_23 int
				return hx_zero_23
			}
			return hx_value_22.(int)
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
		self.root = func(hx_value_26 any) *haxe__ds__TreeNode {
			if hx_value_26 == nil {
				var hx_zero_27 *haxe__ds__TreeNode
				return hx_zero_27
			}
			return hx_value_26.(*haxe__ds__TreeNode)
		}(self.removeLoop(key, self.root))
	}, func(hx_caught_24 any) {
		switch hx_typed_25 := hx_caught_24.(type) {
		case *string:
			hx_tmp := hx_typed_25
			_ = hx_tmp
			removed = false
		default:
			hxrt.Throw(hx_caught_24)
		}
	})
	return removed
}

func (self *haxe__ds__BalancedTree) exists(key any) bool {
	node := self.root
	for node != nil {
		c := func(hx_value_28 any) int {
			if hx_value_28 == nil {
				var hx_zero_29 int
				return hx_zero_29
			}
			return hx_value_28.(int)
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
	ret := haxe__ds__BalancedTree_iteratorLoop(self.root, []any{})
	index := 0
	hx_obj_30 := map[string]any{}
	hx_obj_30["hasNext"] = func() bool {
		return (index < len(ret))
	}
	hx_obj_30["next"] = func() any {
		hx_post_31 := index
		index = int(int32((index + 1)))
		return ret[hx_post_31]
	}
	return hx_obj_30
}

func (self *haxe__ds__BalancedTree) keys() map[string]any {
	ret := func(hx_value_32 any) []any {
		if hx_value_32 == nil {
			var hx_zero_33 []any
			return hx_zero_33
		}
		return hx_value_32.([]any)
	}(self.keysLoop(self.root, []any{}))
	index := 0
	hx_obj_34 := map[string]any{}
	hx_obj_34["hasNext"] = func() bool {
		return (index < len(ret))
	}
	hx_obj_34["next"] = func() any {
		hx_post_35 := index
		index = int(int32((index + 1)))
		return ret[hx_post_35]
	}
	return hx_obj_34
}

func (self *haxe__ds__BalancedTree) keyValueIterator() map[string]any {
	_gthis := self
	keyIterator := func(hx_value_36 any) map[string]any {
		if hx_value_36 == nil {
			var hx_zero_37 map[string]any
			return hx_zero_37
		}
		return hx_value_36.(map[string]any)
	}(self.keys())
	hx_obj_38 := map[string]any{}
	hx_obj_38["hasNext"] = func() bool {
		return func(hx_obj_39 map[string]any) func() bool {
			hx_field_40 := hx_obj_39["hasNext"]
			if hx_field_40 == nil {
				var hx_zero_41 func() bool
				return hx_zero_41
			}
			return hx_field_40.(func() bool)
		}(keyIterator)()
	}
	hx_obj_38["next"] = func() map[string]any {
		var key any = func(hx_obj_42 map[string]any) func() any {
			hx_field_43 := hx_obj_42["next"]
			if hx_field_43 == nil {
				var hx_zero_44 func() any
				return hx_zero_44
			}
			return hx_field_43.(func() any)
		}(keyIterator)()
		hx_obj_45 := map[string]any{}
		hx_obj_45["key"] = key
		hx_obj_45["value"] = _gthis.get(key)
		return hx_obj_45
	}
	return hx_obj_38
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
	c := func(hx_value_46 any) int {
		if hx_value_46 == nil {
			var hx_zero_47 int
			return hx_zero_47
		}
		return hx_value_46.(int)
	}(self.compare(k, node.key))
	var hx_if_58 *haxe__ds__TreeNode
	if c == 0 {
		hx_if_58 = New_haxe__ds__TreeNode(node.left, k, v, node.right, func() int {
			var hx_if_48 int
			if node == nil {
				hx_if_48 = 0
			} else {
				hx_if_48 = node._height
			}
			return hx_if_48
		}())
	} else {
		var hx_if_57 *haxe__ds__TreeNode
		if c < 0 {
			nl := func(hx_value_49 any) *haxe__ds__TreeNode {
				if hx_value_49 == nil {
					var hx_zero_50 *haxe__ds__TreeNode
					return hx_zero_50
				}
				return hx_value_49.(*haxe__ds__TreeNode)
			}(self.setLoop(k, v, node.left))
			hx_if_57 = func(hx_value_51 any) *haxe__ds__TreeNode {
				if hx_value_51 == nil {
					var hx_zero_52 *haxe__ds__TreeNode
					return hx_zero_52
				}
				return hx_value_51.(*haxe__ds__TreeNode)
			}(self.balance(nl, node.key, node.value, node.right))
		} else {
			nr := func(hx_value_53 any) *haxe__ds__TreeNode {
				if hx_value_53 == nil {
					var hx_zero_54 *haxe__ds__TreeNode
					return hx_zero_54
				}
				return hx_value_53.(*haxe__ds__TreeNode)
			}(self.setLoop(k, v, node.right))
			hx_if_57 = func(hx_value_55 any) *haxe__ds__TreeNode {
				if hx_value_55 == nil {
					var hx_zero_56 *haxe__ds__TreeNode
					return hx_zero_56
				}
				return hx_value_55.(*haxe__ds__TreeNode)
			}(self.balance(node.left, node.key, node.value, nr))
		}
		hx_if_58 = hx_if_57
	}
	return hx_if_58
}

func (self *haxe__ds__BalancedTree) removeLoop(k any, node *haxe__ds__TreeNode) *haxe__ds__TreeNode {
	if node == nil {
		hxrt.Throw(hxrt.StringFromLiteral("Not_found"))
		var hx_throw_zero_59 *haxe__ds__TreeNode
		return hx_throw_zero_59
	}
	c := func(hx_value_60 any) int {
		if hx_value_60 == nil {
			var hx_zero_61 int
			return hx_zero_61
		}
		return hx_value_60.(int)
	}(self.compare(k, node.key))
	var hx_if_73 *haxe__ds__TreeNode
	if c == 0 {
		hx_if_73 = func(hx_value_62 any) *haxe__ds__TreeNode {
			if hx_value_62 == nil {
				var hx_zero_63 *haxe__ds__TreeNode
				return hx_zero_63
			}
			return hx_value_62.(*haxe__ds__TreeNode)
		}(self.merge(node.left, node.right))
	} else {
		var hx_if_72 *haxe__ds__TreeNode
		if c < 0 {
			hx_if_72 = func(hx_value_66 any) *haxe__ds__TreeNode {
				if hx_value_66 == nil {
					var hx_zero_67 *haxe__ds__TreeNode
					return hx_zero_67
				}
				return hx_value_66.(*haxe__ds__TreeNode)
			}(self.balance(func(hx_value_64 any) *haxe__ds__TreeNode {
				if hx_value_64 == nil {
					var hx_zero_65 *haxe__ds__TreeNode
					return hx_zero_65
				}
				return hx_value_64.(*haxe__ds__TreeNode)
			}(self.removeLoop(k, node.left)), node.key, node.value, node.right))
		} else {
			hx_if_72 = func(hx_value_70 any) *haxe__ds__TreeNode {
				if hx_value_70 == nil {
					var hx_zero_71 *haxe__ds__TreeNode
					return hx_zero_71
				}
				return hx_value_70.(*haxe__ds__TreeNode)
			}(self.balance(node.left, node.key, node.value, func(hx_value_68 any) *haxe__ds__TreeNode {
				if hx_value_68 == nil {
					var hx_zero_69 *haxe__ds__TreeNode
					return hx_zero_69
				}
				return hx_value_68.(*haxe__ds__TreeNode)
			}(self.removeLoop(k, node.right))))
		}
		hx_if_73 = hx_if_72
	}
	return hx_if_73
}

func (self *haxe__ds__BalancedTree) keysLoop(node *haxe__ds__TreeNode, acc []any) []any {
	if node != nil {
		acc = func(hx_value_74 any) []any {
			if hx_value_74 == nil {
				var hx_zero_75 []any
				return hx_zero_75
			}
			return hx_value_74.([]any)
		}(self.keysLoop(node.left, acc))
		hx_arr_76 := acc
		hx_arr_76 = append(hx_arr_76, node.key)
		acc = hx_arr_76
		acc = func(hx_value_77 any) []any {
			if hx_value_77 == nil {
				var hx_zero_78 []any
				return hx_zero_78
			}
			return hx_value_77.([]any)
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
	t := func(hx_value_79 any) *haxe__ds__TreeNode {
		if hx_value_79 == nil {
			var hx_zero_80 *haxe__ds__TreeNode
			return hx_zero_80
		}
		return hx_value_79.(*haxe__ds__TreeNode)
	}(self.minBinding(t2))
	return func(hx_value_83 any) *haxe__ds__TreeNode {
		if hx_value_83 == nil {
			var hx_zero_84 *haxe__ds__TreeNode
			return hx_zero_84
		}
		return hx_value_83.(*haxe__ds__TreeNode)
	}(self.balance(t1, t.key, t.value, func(hx_value_81 any) *haxe__ds__TreeNode {
		if hx_value_81 == nil {
			var hx_zero_82 *haxe__ds__TreeNode
			return hx_zero_82
		}
		return hx_value_81.(*haxe__ds__TreeNode)
	}(self.removeMinBinding(t2))))
}

func (self *haxe__ds__BalancedTree) minBinding(t *haxe__ds__TreeNode) *haxe__ds__TreeNode {
	var hx_if_89 *haxe__ds__TreeNode
	if t == nil {
		hx_if_89 = func() *haxe__ds__TreeNode {
			hxrt.Throw(hxrt.StringFromLiteral("Not_found"))
			var hx_throw_zero_85 *haxe__ds__TreeNode
			return hx_throw_zero_85
		}()
	} else {
		var hx_if_88 *haxe__ds__TreeNode
		if t.left == nil {
			hx_if_88 = t
		} else {
			hx_if_88 = func(hx_value_86 any) *haxe__ds__TreeNode {
				if hx_value_86 == nil {
					var hx_zero_87 *haxe__ds__TreeNode
					return hx_zero_87
				}
				return hx_value_86.(*haxe__ds__TreeNode)
			}(self.minBinding(t.left))
		}
		hx_if_89 = hx_if_88
	}
	return hx_if_89
}

func (self *haxe__ds__BalancedTree) removeMinBinding(t *haxe__ds__TreeNode) *haxe__ds__TreeNode {
	var hx_if_94 *haxe__ds__TreeNode
	if t.left == nil {
		hx_if_94 = t.right
	} else {
		hx_if_94 = func(hx_value_92 any) *haxe__ds__TreeNode {
			if hx_value_92 == nil {
				var hx_zero_93 *haxe__ds__TreeNode
				return hx_zero_93
			}
			return hx_value_92.(*haxe__ds__TreeNode)
		}(self.balance(func(hx_value_90 any) *haxe__ds__TreeNode {
			if hx_value_90 == nil {
				var hx_zero_91 *haxe__ds__TreeNode
				return hx_zero_91
			}
			return hx_value_90.(*haxe__ds__TreeNode)
		}(self.removeMinBinding(t.left)), t.key, t.value, t.right))
	}
	return hx_if_94
}

func (self *haxe__ds__BalancedTree) balance(l *haxe__ds__TreeNode, k any, v any, r *haxe__ds__TreeNode) *haxe__ds__TreeNode {
	var hx_if_95 int
	if l == nil {
		hx_if_95 = 0
	} else {
		hx_if_95 = l._height
	}
	hl := hx_if_95
	var hx_if_96 int
	if r == nil {
		hx_if_96 = 0
	} else {
		hx_if_96 = r._height
	}
	hr := hx_if_96
	var hx_if_105 *haxe__ds__TreeNode
	if hl > int(int32((hxrt.Int32Wrap(hr) + hxrt.Int32Wrap(2)))) {
		var hx_if_99 *haxe__ds__TreeNode
		if func() int {
			_this := l.left
			var hx_if_97 int
			if _this == nil {
				hx_if_97 = 0
			} else {
				hx_if_97 = _this._height
			}
			return hx_if_97
		}() >= func() int {
			_this_1 := l.right
			var hx_if_98 int
			if _this_1 == nil {
				hx_if_98 = 0
			} else {
				hx_if_98 = _this_1._height
			}
			return hx_if_98
		}() {
			hx_if_99 = New_haxe__ds__TreeNode(l.left, l.key, l.value, New_haxe__ds__TreeNode(l.right, k, v, r, -1), -1)
		} else {
			hx_if_99 = New_haxe__ds__TreeNode(New_haxe__ds__TreeNode(l.left, l.key, l.value, l.right.left, -1), l.right.key, l.right.value, New_haxe__ds__TreeNode(l.right.right, k, v, r, -1), -1)
		}
		hx_if_105 = hx_if_99
	} else {
		var hx_if_104 *haxe__ds__TreeNode
		if hr > int(int32((hxrt.Int32Wrap(hl) + hxrt.Int32Wrap(2)))) {
			var hx_if_102 *haxe__ds__TreeNode
			if func() int {
				_this_2 := r.right
				var hx_if_100 int
				if _this_2 == nil {
					hx_if_100 = 0
				} else {
					hx_if_100 = _this_2._height
				}
				return hx_if_100
			}() > func() int {
				_this_3 := r.left
				var hx_if_101 int
				if _this_3 == nil {
					hx_if_101 = 0
				} else {
					hx_if_101 = _this_3._height
				}
				return hx_if_101
			}() {
				hx_if_102 = New_haxe__ds__TreeNode(New_haxe__ds__TreeNode(l, k, v, r.left, -1), r.key, r.value, r.right, -1)
			} else {
				hx_if_102 = New_haxe__ds__TreeNode(New_haxe__ds__TreeNode(l, k, v, r.left.left, -1), r.left.key, r.left.value, New_haxe__ds__TreeNode(r.left.right, r.key, r.value, r.right, -1), -1)
			}
			hx_if_104 = hx_if_102
		} else {
			hx_if_104 = New_haxe__ds__TreeNode(l, k, v, r, int(int32((hxrt.Int32Wrap(func() int {
				var hx_if_103 int
				if hl > hr {
					hx_if_103 = hl
				} else {
					hx_if_103 = hr
				}
				return hx_if_103
			}()) + hxrt.Int32Wrap(1)))))
		}
		hx_if_105 = hx_if_104
	}
	return hx_if_105
}

func (self *haxe__ds__BalancedTree) compare(k1 any, k2 any) int {
	return Reflect_compare(k1, k2)
}

func (self *haxe__ds__BalancedTree) toString() *string {
	var hx_if_108 *string
	if self.root == nil {
		hx_if_108 = hxrt.StringFromLiteral("[]")
	} else {
		hx_if_108 = hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("["), func(hx_value_106 any) *string {
			if hx_value_106 == nil {
				var hx_zero_107 *string
				return hx_zero_107
			}
			return hx_value_106.(*string)
		}(self.root.toString())), hxrt.StringFromLiteral("]"))
	}
	return hx_if_108
}

func (self *haxe__ds__BalancedTree) clear() {
	self.root = nil
}

func haxe__ds__BalancedTree_iteratorLoop(node *haxe__ds__TreeNode, acc []any) []any {
	if node != nil {
		acc = haxe__ds__BalancedTree_iteratorLoop(node.left, acc)
		hx_arr_109 := acc
		hx_arr_109 = append(hx_arr_109, node.value)
		acc = hx_arr_109
		acc = haxe__ds__BalancedTree_iteratorLoop(node.right, acc)
	}
	return acc
}
