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
			var hx_if_5 int
			if func() int {
				_this := self.left
				var hx_if_1 int
				if _this == nil {
					hx_if_1 = 0
				} else {
					hx_if_1 = _this._height
				}
				return hx_if_1
			}() > func() int {
				_this_1 := self.right
				var hx_if_2 int
				if _this_1 == nil {
					hx_if_2 = 0
				} else {
					hx_if_2 = _this_1._height
				}
				return hx_if_2
			}() {
				_this_2 := self.left
				var hx_if_3 int
				if _this_2 == nil {
					hx_if_3 = 0
				} else {
					hx_if_3 = _this_2._height
				}
				hx_if_5 = hx_if_3
			} else {
				_this_3 := self.right
				var hx_if_4 int
				if _this_3 == nil {
					hx_if_4 = 0
				} else {
					hx_if_4 = _this_3._height
				}
				hx_if_5 = hx_if_4
			}
			return hx_if_5
		}()) + hxrt.Int32Wrap(1))))
	} else {
		self._height = h
	}
	return self
}

func (self *haxe__ds__TreeNode) toString() *string {
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
			}(self.left.__hx_this.toString()), hxrt.StringFromLiteral(", "))
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
			}(self.right.__hx_this.toString()))
		}
		return hx_if_11
	}())
}

func (self *haxe__ds__TreeNode) String() string {
	return *self.__hx_this.toString()
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
	self.root = func(hx_value_12 any) *haxe__ds__TreeNode {
		if hx_value_12 == nil {
			var hx_zero_13 *haxe__ds__TreeNode
			return hx_zero_13
		}
		return hx_value_12.(*haxe__ds__TreeNode)
	}(self.__hx_this.setLoop(key, value, self.root))
}

func (self *haxe__ds__BalancedTree) get(key any) any {
	node := self.root
	for node != nil {
		c := func(hx_value_14 any) int {
			if hx_value_14 == nil {
				var hx_zero_15 int
				return hx_zero_15
			}
			return hx_value_14.(int)
		}(self.__hx_this.compare(key, node.key))
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
		self.root = func(hx_value_18 any) *haxe__ds__TreeNode {
			if hx_value_18 == nil {
				var hx_zero_19 *haxe__ds__TreeNode
				return hx_zero_19
			}
			return hx_value_18.(*haxe__ds__TreeNode)
		}(self.__hx_this.removeLoop(key, self.root))
	}, func(hx_caught_16 any) {
		switch hx_typed_17 := hx_caught_16.(type) {
		case *string:
			hx_tmp := hx_typed_17
			_ = hx_tmp
			removed = false
		default:
			hxrt.Throw(hx_caught_16)
		}
	})
	return removed
}

func (self *haxe__ds__BalancedTree) exists(key any) bool {
	node := self.root
	for node != nil {
		c := func(hx_value_20 any) int {
			if hx_value_20 == nil {
				var hx_zero_21 int
				return hx_zero_21
			}
			return hx_value_20.(int)
		}(self.__hx_this.compare(key, node.key))
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
	hx_obj_22 := map[string]any{}
	hx_obj_22["hasNext"] = func() bool {
		return (index < ret.Len())
	}
	hx_obj_22["next"] = func() any {
		hx_post_23 := index
		index = int(int32((index + 1)))
		return ret.Get(hx_post_23)
	}
	return hx_obj_22
}

func (self *haxe__ds__BalancedTree) keys() map[string]any {
	ret := func(hx_value_24 any) *hxrt.Array {
		if hx_value_24 == nil {
			var hx_zero_25 *hxrt.Array
			return hx_zero_25
		}
		return hx_value_24.(*hxrt.Array)
	}(self.__hx_this.keysLoop(self.root, hxrt.NewArray()))
	index := 0
	hx_obj_26 := map[string]any{}
	hx_obj_26["hasNext"] = func() bool {
		return (index < ret.Len())
	}
	hx_obj_26["next"] = func() any {
		hx_post_27 := index
		index = int(int32((index + 1)))
		return ret.Get(hx_post_27)
	}
	return hx_obj_26
}

func (self *haxe__ds__BalancedTree) keyValueIterator() map[string]any {
	_gthis := self
	keyIterator := func(hx_value_28 any) map[string]any {
		if hx_value_28 == nil {
			var hx_zero_29 map[string]any
			return hx_zero_29
		}
		return hx_value_28.(map[string]any)
	}(self.__hx_this.keys())
	hx_obj_30 := map[string]any{}
	hx_obj_30["hasNext"] = func() bool {
		return func(hx_obj_31 map[string]any) func() bool {
			hx_field_32 := hx_obj_31["hasNext"]
			if hx_field_32 == nil {
				var hx_zero_33 func() bool
				return hx_zero_33
			}
			return hx_field_32.(func() bool)
		}(keyIterator)()
	}
	hx_obj_30["next"] = func() map[string]any {
		var key any = func(hx_obj_34 map[string]any) func() any {
			hx_field_35 := hx_obj_34["next"]
			if hx_field_35 == nil {
				var hx_zero_36 func() any
				return hx_zero_36
			}
			return hx_field_35.(func() any)
		}(keyIterator)()
		hx_obj_37 := map[string]any{}
		hx_obj_37["key"] = key
		hx_obj_37["value"] = _gthis.__hx_this.get(key)
		return hx_obj_37
	}
	return hx_obj_30
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
	c := func(hx_value_38 any) int {
		if hx_value_38 == nil {
			var hx_zero_39 int
			return hx_zero_39
		}
		return hx_value_38.(int)
	}(self.__hx_this.compare(k, node.key))
	var hx_if_50 *haxe__ds__TreeNode
	if c == 0 {
		hx_if_50 = New_haxe__ds__TreeNode(node.left, k, v, node.right, func() int {
			var hx_if_40 int
			if node == nil {
				hx_if_40 = 0
			} else {
				hx_if_40 = node._height
			}
			return hx_if_40
		}())
	} else {
		var hx_if_49 *haxe__ds__TreeNode
		if c < 0 {
			nl := func(hx_value_41 any) *haxe__ds__TreeNode {
				if hx_value_41 == nil {
					var hx_zero_42 *haxe__ds__TreeNode
					return hx_zero_42
				}
				return hx_value_41.(*haxe__ds__TreeNode)
			}(self.__hx_this.setLoop(k, v, node.left))
			hx_if_49 = func(hx_value_43 any) *haxe__ds__TreeNode {
				if hx_value_43 == nil {
					var hx_zero_44 *haxe__ds__TreeNode
					return hx_zero_44
				}
				return hx_value_43.(*haxe__ds__TreeNode)
			}(self.__hx_this.balance(nl, node.key, node.value, node.right))
		} else {
			nr := func(hx_value_45 any) *haxe__ds__TreeNode {
				if hx_value_45 == nil {
					var hx_zero_46 *haxe__ds__TreeNode
					return hx_zero_46
				}
				return hx_value_45.(*haxe__ds__TreeNode)
			}(self.__hx_this.setLoop(k, v, node.right))
			hx_if_49 = func(hx_value_47 any) *haxe__ds__TreeNode {
				if hx_value_47 == nil {
					var hx_zero_48 *haxe__ds__TreeNode
					return hx_zero_48
				}
				return hx_value_47.(*haxe__ds__TreeNode)
			}(self.__hx_this.balance(node.left, node.key, node.value, nr))
		}
		hx_if_50 = hx_if_49
	}
	return hx_if_50
}

func (self *haxe__ds__BalancedTree) removeLoop(k any, node *haxe__ds__TreeNode) *haxe__ds__TreeNode {
	if node == nil {
		hxrt.Throw(hxrt.StringFromLiteral("Not_found"))
	}
	c := func(hx_value_51 any) int {
		if hx_value_51 == nil {
			var hx_zero_52 int
			return hx_zero_52
		}
		return hx_value_51.(int)
	}(self.__hx_this.compare(k, node.key))
	var hx_if_64 *haxe__ds__TreeNode
	if c == 0 {
		hx_if_64 = func(hx_value_53 any) *haxe__ds__TreeNode {
			if hx_value_53 == nil {
				var hx_zero_54 *haxe__ds__TreeNode
				return hx_zero_54
			}
			return hx_value_53.(*haxe__ds__TreeNode)
		}(self.__hx_this.merge(node.left, node.right))
	} else {
		var hx_if_63 *haxe__ds__TreeNode
		if c < 0 {
			hx_if_63 = func(hx_value_57 any) *haxe__ds__TreeNode {
				if hx_value_57 == nil {
					var hx_zero_58 *haxe__ds__TreeNode
					return hx_zero_58
				}
				return hx_value_57.(*haxe__ds__TreeNode)
			}(self.__hx_this.balance(func(hx_value_55 any) *haxe__ds__TreeNode {
				if hx_value_55 == nil {
					var hx_zero_56 *haxe__ds__TreeNode
					return hx_zero_56
				}
				return hx_value_55.(*haxe__ds__TreeNode)
			}(self.__hx_this.removeLoop(k, node.left)), node.key, node.value, node.right))
		} else {
			hx_if_63 = func(hx_value_61 any) *haxe__ds__TreeNode {
				if hx_value_61 == nil {
					var hx_zero_62 *haxe__ds__TreeNode
					return hx_zero_62
				}
				return hx_value_61.(*haxe__ds__TreeNode)
			}(self.__hx_this.balance(node.left, node.key, node.value, func(hx_value_59 any) *haxe__ds__TreeNode {
				if hx_value_59 == nil {
					var hx_zero_60 *haxe__ds__TreeNode
					return hx_zero_60
				}
				return hx_value_59.(*haxe__ds__TreeNode)
			}(self.__hx_this.removeLoop(k, node.right))))
		}
		hx_if_64 = hx_if_63
	}
	return hx_if_64
}

func (self *haxe__ds__BalancedTree) keysLoop(node *haxe__ds__TreeNode, acc *hxrt.Array) *hxrt.Array {
	if node != nil {
		acc = func(hx_value_65 any) *hxrt.Array {
			if hx_value_65 == nil {
				var hx_zero_66 *hxrt.Array
				return hx_zero_66
			}
			return hx_value_65.(*hxrt.Array)
		}(self.__hx_this.keysLoop(node.left, acc))
		acc.Push(node.key)
		acc = func(hx_value_68 any) *hxrt.Array {
			if hx_value_68 == nil {
				var hx_zero_69 *hxrt.Array
				return hx_zero_69
			}
			return hx_value_68.(*hxrt.Array)
		}(self.__hx_this.keysLoop(node.right, acc))
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
	t := func(hx_value_70 any) *haxe__ds__TreeNode {
		if hx_value_70 == nil {
			var hx_zero_71 *haxe__ds__TreeNode
			return hx_zero_71
		}
		return hx_value_70.(*haxe__ds__TreeNode)
	}(self.__hx_this.minBinding(t2))
	return func(hx_value_74 any) *haxe__ds__TreeNode {
		if hx_value_74 == nil {
			var hx_zero_75 *haxe__ds__TreeNode
			return hx_zero_75
		}
		return hx_value_74.(*haxe__ds__TreeNode)
	}(self.__hx_this.balance(t1, t.key, t.value, func(hx_value_72 any) *haxe__ds__TreeNode {
		if hx_value_72 == nil {
			var hx_zero_73 *haxe__ds__TreeNode
			return hx_zero_73
		}
		return hx_value_72.(*haxe__ds__TreeNode)
	}(self.__hx_this.removeMinBinding(t2))))
}

func (self *haxe__ds__BalancedTree) minBinding(t *haxe__ds__TreeNode) *haxe__ds__TreeNode {
	var hx_if_80 *haxe__ds__TreeNode
	if t == nil {
		hx_if_80 = func() *haxe__ds__TreeNode {
			hxrt.Throw(hxrt.StringFromLiteral("Not_found"))
			var hx_throw_zero_76 *haxe__ds__TreeNode
			return hx_throw_zero_76
		}()
	} else {
		var hx_if_79 *haxe__ds__TreeNode
		if t.left == nil {
			hx_if_79 = t
		} else {
			hx_if_79 = func(hx_value_77 any) *haxe__ds__TreeNode {
				if hx_value_77 == nil {
					var hx_zero_78 *haxe__ds__TreeNode
					return hx_zero_78
				}
				return hx_value_77.(*haxe__ds__TreeNode)
			}(self.__hx_this.minBinding(t.left))
		}
		hx_if_80 = hx_if_79
	}
	return hx_if_80
}

func (self *haxe__ds__BalancedTree) removeMinBinding(t *haxe__ds__TreeNode) *haxe__ds__TreeNode {
	var hx_if_85 *haxe__ds__TreeNode
	if t.left == nil {
		hx_if_85 = t.right
	} else {
		hx_if_85 = func(hx_value_83 any) *haxe__ds__TreeNode {
			if hx_value_83 == nil {
				var hx_zero_84 *haxe__ds__TreeNode
				return hx_zero_84
			}
			return hx_value_83.(*haxe__ds__TreeNode)
		}(self.__hx_this.balance(func(hx_value_81 any) *haxe__ds__TreeNode {
			if hx_value_81 == nil {
				var hx_zero_82 *haxe__ds__TreeNode
				return hx_zero_82
			}
			return hx_value_81.(*haxe__ds__TreeNode)
		}(self.__hx_this.removeMinBinding(t.left)), t.key, t.value, t.right))
	}
	return hx_if_85
}

func (self *haxe__ds__BalancedTree) balance(l *haxe__ds__TreeNode, k any, v any, r *haxe__ds__TreeNode) *haxe__ds__TreeNode {
	var hx_if_86 int
	if l == nil {
		hx_if_86 = 0
	} else {
		hx_if_86 = l._height
	}
	hl := hx_if_86
	var hx_if_87 int
	if r == nil {
		hx_if_87 = 0
	} else {
		hx_if_87 = r._height
	}
	hr := hx_if_87
	var hx_if_96 *haxe__ds__TreeNode
	if hl > int(int32((hxrt.Int32Wrap(hr) + hxrt.Int32Wrap(2)))) {
		var hx_if_90 *haxe__ds__TreeNode
		if func() int {
			_this := l.left
			var hx_if_88 int
			if _this == nil {
				hx_if_88 = 0
			} else {
				hx_if_88 = _this._height
			}
			return hx_if_88
		}() >= func() int {
			_this_1 := l.right
			var hx_if_89 int
			if _this_1 == nil {
				hx_if_89 = 0
			} else {
				hx_if_89 = _this_1._height
			}
			return hx_if_89
		}() {
			hx_if_90 = New_haxe__ds__TreeNode(l.left, l.key, l.value, New_haxe__ds__TreeNode(l.right, k, v, r, -1), -1)
		} else {
			hx_if_90 = New_haxe__ds__TreeNode(New_haxe__ds__TreeNode(l.left, l.key, l.value, l.right.left, -1), l.right.key, l.right.value, New_haxe__ds__TreeNode(l.right.right, k, v, r, -1), -1)
		}
		hx_if_96 = hx_if_90
	} else {
		var hx_if_95 *haxe__ds__TreeNode
		if hr > int(int32((hxrt.Int32Wrap(hl) + hxrt.Int32Wrap(2)))) {
			var hx_if_93 *haxe__ds__TreeNode
			if func() int {
				_this_2 := r.right
				var hx_if_91 int
				if _this_2 == nil {
					hx_if_91 = 0
				} else {
					hx_if_91 = _this_2._height
				}
				return hx_if_91
			}() > func() int {
				_this_3 := r.left
				var hx_if_92 int
				if _this_3 == nil {
					hx_if_92 = 0
				} else {
					hx_if_92 = _this_3._height
				}
				return hx_if_92
			}() {
				hx_if_93 = New_haxe__ds__TreeNode(New_haxe__ds__TreeNode(l, k, v, r.left, -1), r.key, r.value, r.right, -1)
			} else {
				hx_if_93 = New_haxe__ds__TreeNode(New_haxe__ds__TreeNode(l, k, v, r.left.left, -1), r.left.key, r.left.value, New_haxe__ds__TreeNode(r.left.right, r.key, r.value, r.right, -1), -1)
			}
			hx_if_95 = hx_if_93
		} else {
			hx_if_95 = New_haxe__ds__TreeNode(l, k, v, r, int(int32((hxrt.Int32Wrap(func() int {
				var hx_if_94 int
				if hl > hr {
					hx_if_94 = hl
				} else {
					hx_if_94 = hr
				}
				return hx_if_94
			}()) + hxrt.Int32Wrap(1)))))
		}
		hx_if_96 = hx_if_95
	}
	return hx_if_96
}

func (self *haxe__ds__BalancedTree) compare(k1 any, k2 any) int {
	return Reflect_compare(k1, k2)
}

func (self *haxe__ds__BalancedTree) toString() *string {
	var hx_if_99 *string
	if self.root == nil {
		hx_if_99 = hxrt.StringFromLiteral("[]")
	} else {
		hx_if_99 = hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("["), func(hx_value_97 any) *string {
			if hx_value_97 == nil {
				var hx_zero_98 *string
				return hx_zero_98
			}
			return hx_value_97.(*string)
		}(self.root.__hx_this.toString())), hxrt.StringFromLiteral("]"))
	}
	return hx_if_99
}

func (self *haxe__ds__BalancedTree) clear() {
	self.root = nil
}

func (self *haxe__ds__BalancedTree) getIMap(key any) any {
	return self.__hx_this.get(key)
}

func (self *haxe__ds__BalancedTree) setIMap(key any, value any) {
	self.__hx_this.set(key, value)
}

func (self *haxe__ds__BalancedTree) existsIMap(key any) bool {
	return func(hx_value_100 any) bool {
		if hx_value_100 == nil {
			var hx_zero_101 bool
			return hx_zero_101
		}
		return hx_value_100.(bool)
	}(self.__hx_this.exists(key))
}

func (self *haxe__ds__BalancedTree) removeIMap(key any) bool {
	return func(hx_value_102 any) bool {
		if hx_value_102 == nil {
			var hx_zero_103 bool
			return hx_zero_103
		}
		return hx_value_102.(bool)
	}(self.__hx_this.remove(key))
}

func (self *haxe__ds__BalancedTree) copyIMap() haxe__IMap {
	return func(hx_value_104 any) *haxe__ds__BalancedTree {
		if hx_value_104 == nil {
			var hx_zero_105 *haxe__ds__BalancedTree
			return hx_zero_105
		}
		return hx_value_104.(*haxe__ds__BalancedTree)
	}(self.__hx_this.copy())
}

func (self *haxe__ds__BalancedTree) String() string {
	return *self.__hx_this.toString()
}

func haxe__ds__BalancedTree_iteratorLoop(node *haxe__ds__TreeNode, acc *hxrt.Array) *hxrt.Array {
	if node != nil {
		acc = haxe__ds__BalancedTree_iteratorLoop(node.left, acc)
		acc.Push(node.value)
		acc = haxe__ds__BalancedTree_iteratorLoop(node.right, acc)
	}
	return acc
}
