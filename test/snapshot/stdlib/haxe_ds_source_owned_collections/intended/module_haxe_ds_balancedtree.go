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

func (self *haxe__ds__TreeNode) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_haxe__ds__TreeNode(l *haxe__ds__TreeNode, k any, v any, r *haxe__ds__TreeNode, h int) *haxe__ds__TreeNode {
	return hx_init_haxe__ds__TreeNode(hx_alloc_haxe__ds__TreeNode(nil), l, k, v, r, h)
}

func hx_alloc_haxe__ds__TreeNode(owner any) *haxe__ds__TreeNode {
	self := &haxe__ds__TreeNode{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_haxe__ds__TreeNode)
	}
	return self
}

func hx_init_haxe__ds__TreeNode(self *haxe__ds__TreeNode, l *haxe__ds__TreeNode, k any, v any, r *haxe__ds__TreeNode, h int) *haxe__ds__TreeNode {
	self.left = l
	self.key = k
	self.value = v
	self.right = r
	if h == -1 {
		self._height = int((hxrt.Int32Wrap(func() int {
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
		}()) + hxrt.Int32Wrap(1)))
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
	iterator() hxrt.Anonymous
	keys() hxrt.Anonymous
	keyValueIterator() hxrt.Anonymous
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

func (self *haxe__ds__BalancedTree) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_haxe__ds__BalancedTree() *haxe__ds__BalancedTree {
	return hx_init_haxe__ds__BalancedTree(hx_alloc_haxe__ds__BalancedTree(nil))
}

func hx_alloc_haxe__ds__BalancedTree(owner any) *haxe__ds__BalancedTree {
	self := &haxe__ds__BalancedTree{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_haxe__ds__BalancedTree)
	}
	return self
}

func hx_init_haxe__ds__BalancedTree(self *haxe__ds__BalancedTree) *haxe__ds__BalancedTree {
	return self
}

func (self *haxe__ds__BalancedTree) set(key any, value any) {
	self.root = hx_recover_class_0(self.__hx_this.setLoop(key, value, self.root))
}

func (self *haxe__ds__BalancedTree) get(key any) any {
	node := self.root
	for node != nil {
		c := func(hx_value_12 any) int {
			if hx_value_12 == nil {
				var hx_zero_13 int
				return hx_zero_13
			}
			return hx_value_12.(int)
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
		self.root = hx_recover_class_0(self.__hx_this.removeLoop(key, self.root))
	}, func(hx_caught_14 any) {
		hx_exception_15 := haxe__Exception_caught(hx_caught_14)
		hx_payload_16 := haxe__Exception_catchValue(hx_caught_14)
		_ = hx_exception_15
		if func(hx_value any) bool {
			switch hx_value.(type) {
			case *string:
				return true
			default:
				return false
			}
		}(hx_payload_16) {
			var hx_tmp *string = func(hx_value_17 any) *string {
				if hx_value_17 == nil {
					var hx_zero_18 *string
					return hx_zero_18
				}
				return hx_value_17.(*string)
			}(hx_payload_16)
			_ = hx_tmp
			removed = false
		} else {
			hxrt.Throw(hx_caught_14)
		}
	})
	return removed
}

func (self *haxe__ds__BalancedTree) exists(key any) bool {
	node := self.root
	for node != nil {
		c := func(hx_value_19 any) int {
			if hx_value_19 == nil {
				var hx_zero_20 int
				return hx_zero_20
			}
			return hx_value_19.(int)
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

func (self *haxe__ds__BalancedTree) iterator() hxrt.Anonymous {
	ret := haxe__ds__BalancedTree_iteratorLoop(self.root, hxrt.NewArray())
	index := 0
	hx_obj_21 := map[string]any{}
	hx_obj_21["hasNext"] = hxrt.NewFunction(func() bool {
		return (index < ret.Len())
	})
	hx_obj_21["next"] = hxrt.NewFunction(func() any {
		hx_post_22 := index
		index = int(int32((index + 1)))
		return ret.Get(hx_post_22)
	})
	return hxrt.AnonymousFromMap(hx_obj_21)
}

func (self *haxe__ds__BalancedTree) keys() hxrt.Anonymous {
	ret := func(hx_value_23 any) *hxrt.Array {
		if hx_value_23 == nil {
			var hx_zero_24 *hxrt.Array
			return hx_zero_24
		}
		return hx_value_23.(*hxrt.Array)
	}(self.__hx_this.keysLoop(self.root, hxrt.NewArray()))
	index := 0
	hx_obj_25 := map[string]any{}
	hx_obj_25["hasNext"] = hxrt.NewFunction(func() bool {
		return (index < ret.Len())
	})
	hx_obj_25["next"] = hxrt.NewFunction(func() any {
		hx_post_26 := index
		index = int(int32((index + 1)))
		return ret.Get(hx_post_26)
	})
	return hxrt.AnonymousFromMap(hx_obj_25)
}

func (self *haxe__ds__BalancedTree) keyValueIterator() hxrt.Anonymous {
	_gthis := self
	keyIterator := self.__hx_this.keys()
	hx_obj_27 := map[string]any{}
	hx_obj_27["hasNext"] = hxrt.NewFunction(func() bool {
		return func() bool {
			var hx_callable_28 *hxrt.Function[func() bool] = hx_recover_callable_0(hxrt.AnonymousGet(keyIterator, "hasNext"))
			return hxrt.FunctionCall(hx_callable_28)()
		}()
	})
	hx_obj_27["next"] = hxrt.NewFunction(func() hxrt.Anonymous {
		var key any = func() any {
			var hx_callable_29 *hxrt.Function[func() any] = hx_recover_callable_1(hxrt.AnonymousGet(keyIterator, "next"))
			return hxrt.FunctionCall(hx_callable_29)()
		}()
		hx_obj_30 := map[string]any{}
		hx_obj_30["key"] = key
		hx_obj_30["value"] = _gthis.__hx_this.get(key)
		return hxrt.AnonymousFromMap(hx_obj_30)
	})
	return hxrt.AnonymousFromMap(hx_obj_27)
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
	c := func(hx_value_31 any) int {
		if hx_value_31 == nil {
			var hx_zero_32 int
			return hx_zero_32
		}
		return hx_value_31.(int)
	}(self.__hx_this.compare(k, node.key))
	var hx_if_35 *haxe__ds__TreeNode
	if c == 0 {
		hx_if_35 = New_haxe__ds__TreeNode(node.left, k, v, node.right, func() int {
			var hx_if_33 int
			if node == nil {
				hx_if_33 = 0
			} else {
				hx_if_33 = node._height
			}
			return hx_if_33
		}())
	} else {
		var hx_if_34 *haxe__ds__TreeNode
		if c < 0 {
			nl := hx_recover_class_0(self.__hx_this.setLoop(k, v, node.left))
			hx_if_34 = hx_recover_class_0(self.__hx_this.balance(nl, node.key, node.value, node.right))
		} else {
			nr := hx_recover_class_0(self.__hx_this.setLoop(k, v, node.right))
			hx_if_34 = hx_recover_class_0(self.__hx_this.balance(node.left, node.key, node.value, nr))
		}
		hx_if_35 = hx_if_34
	}
	return hx_if_35
}

func (self *haxe__ds__BalancedTree) removeLoop(k any, node *haxe__ds__TreeNode) *haxe__ds__TreeNode {
	if node == nil {
		hxrt.Throw(haxe__Exception_thrown(hxrt.StringFromLiteral("Not_found")))
	}
	c := func(hx_value_36 any) int {
		if hx_value_36 == nil {
			var hx_zero_37 int
			return hx_zero_37
		}
		return hx_value_36.(int)
	}(self.__hx_this.compare(k, node.key))
	var hx_if_39 *haxe__ds__TreeNode
	if c == 0 {
		hx_if_39 = hx_recover_class_0(self.__hx_this.merge(node.left, node.right))
	} else {
		var hx_if_38 *haxe__ds__TreeNode
		if c < 0 {
			hx_if_38 = hx_recover_class_0(self.__hx_this.balance(hx_recover_class_0(self.__hx_this.removeLoop(k, node.left)), node.key, node.value, node.right))
		} else {
			hx_if_38 = hx_recover_class_0(self.__hx_this.balance(node.left, node.key, node.value, hx_recover_class_0(self.__hx_this.removeLoop(k, node.right))))
		}
		hx_if_39 = hx_if_38
	}
	return hx_if_39
}

func (self *haxe__ds__BalancedTree) keysLoop(node *haxe__ds__TreeNode, acc *hxrt.Array) *hxrt.Array {
	if node != nil {
		acc = func(hx_value_40 any) *hxrt.Array {
			if hx_value_40 == nil {
				var hx_zero_41 *hxrt.Array
				return hx_zero_41
			}
			return hx_value_40.(*hxrt.Array)
		}(self.__hx_this.keysLoop(node.left, acc))
		acc.Push(node.key)
		acc = func(hx_value_43 any) *hxrt.Array {
			if hx_value_43 == nil {
				var hx_zero_44 *hxrt.Array
				return hx_zero_44
			}
			return hx_value_43.(*hxrt.Array)
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
	t := hx_recover_class_0(self.__hx_this.minBinding(t2))
	return hx_recover_class_0(self.__hx_this.balance(t1, t.key, t.value, hx_recover_class_0(self.__hx_this.removeMinBinding(t2))))
}

func (self *haxe__ds__BalancedTree) minBinding(t *haxe__ds__TreeNode) *haxe__ds__TreeNode {
	var hx_if_47 *haxe__ds__TreeNode
	if t == nil {
		hx_if_47 = func() *haxe__ds__TreeNode {
			hxrt.Throw(haxe__Exception_thrown(hxrt.StringFromLiteral("Not_found")))
			var hx_throw_zero_45 *haxe__ds__TreeNode
			return hx_throw_zero_45
		}()
	} else {
		var hx_if_46 *haxe__ds__TreeNode
		if t.left == nil {
			hx_if_46 = t
		} else {
			hx_if_46 = hx_recover_class_0(self.__hx_this.minBinding(t.left))
		}
		hx_if_47 = hx_if_46
	}
	return hx_if_47
}

func (self *haxe__ds__BalancedTree) removeMinBinding(t *haxe__ds__TreeNode) *haxe__ds__TreeNode {
	var hx_if_48 *haxe__ds__TreeNode
	if t.left == nil {
		hx_if_48 = t.right
	} else {
		hx_if_48 = hx_recover_class_0(self.__hx_this.balance(hx_recover_class_0(self.__hx_this.removeMinBinding(t.left)), t.key, t.value, t.right))
	}
	return hx_if_48
}

func (self *haxe__ds__BalancedTree) balance(l *haxe__ds__TreeNode, k any, v any, r *haxe__ds__TreeNode) *haxe__ds__TreeNode {
	var hx_if_49 int
	if l == nil {
		hx_if_49 = 0
	} else {
		hx_if_49 = l._height
	}
	hl := hx_if_49
	var hx_if_50 int
	if r == nil {
		hx_if_50 = 0
	} else {
		hx_if_50 = r._height
	}
	hr := hx_if_50
	var hx_if_59 *haxe__ds__TreeNode
	if hl > int((hxrt.Int32Wrap(hr) + hxrt.Int32Wrap(2))) {
		var hx_if_53 *haxe__ds__TreeNode
		if func() int {
			_this := l.left
			var hx_if_51 int
			if _this == nil {
				hx_if_51 = 0
			} else {
				hx_if_51 = _this._height
			}
			return hx_if_51
		}() >= func() int {
			_this_1 := l.right
			var hx_if_52 int
			if _this_1 == nil {
				hx_if_52 = 0
			} else {
				hx_if_52 = _this_1._height
			}
			return hx_if_52
		}() {
			hx_if_53 = New_haxe__ds__TreeNode(l.left, l.key, l.value, New_haxe__ds__TreeNode(l.right, k, v, r, -1), -1)
		} else {
			hx_if_53 = New_haxe__ds__TreeNode(New_haxe__ds__TreeNode(l.left, l.key, l.value, l.right.left, -1), l.right.key, l.right.value, New_haxe__ds__TreeNode(l.right.right, k, v, r, -1), -1)
		}
		hx_if_59 = hx_if_53
	} else {
		var hx_if_58 *haxe__ds__TreeNode
		if hr > int((hxrt.Int32Wrap(hl) + hxrt.Int32Wrap(2))) {
			var hx_if_56 *haxe__ds__TreeNode
			if func() int {
				_this_2 := r.right
				var hx_if_54 int
				if _this_2 == nil {
					hx_if_54 = 0
				} else {
					hx_if_54 = _this_2._height
				}
				return hx_if_54
			}() > func() int {
				_this_3 := r.left
				var hx_if_55 int
				if _this_3 == nil {
					hx_if_55 = 0
				} else {
					hx_if_55 = _this_3._height
				}
				return hx_if_55
			}() {
				hx_if_56 = New_haxe__ds__TreeNode(New_haxe__ds__TreeNode(l, k, v, r.left, -1), r.key, r.value, r.right, -1)
			} else {
				hx_if_56 = New_haxe__ds__TreeNode(New_haxe__ds__TreeNode(l, k, v, r.left.left, -1), r.left.key, r.left.value, New_haxe__ds__TreeNode(r.left.right, r.key, r.value, r.right, -1), -1)
			}
			hx_if_58 = hx_if_56
		} else {
			hx_if_58 = New_haxe__ds__TreeNode(l, k, v, r, int((hxrt.Int32Wrap(func() int {
				var hx_if_57 int
				if hl > hr {
					hx_if_57 = hl
				} else {
					hx_if_57 = hr
				}
				return hx_if_57
			}()) + hxrt.Int32Wrap(1))))
		}
		hx_if_59 = hx_if_58
	}
	return hx_if_59
}

func (self *haxe__ds__BalancedTree) compare(k1 any, k2 any) int {
	return Reflect_compare(k1, k2)
}

func (self *haxe__ds__BalancedTree) toString() *string {
	var hx_if_62 *string
	if self.root == nil {
		hx_if_62 = hxrt.StringFromLiteral("[]")
	} else {
		hx_if_62 = hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("["), func(hx_value_60 any) *string {
			if hx_value_60 == nil {
				var hx_zero_61 *string
				return hx_zero_61
			}
			return hx_value_60.(*string)
		}(self.root.__hx_this.toString())), hxrt.StringFromLiteral("]"))
	}
	return hx_if_62
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
	return func(hx_value_63 any) bool {
		if hx_value_63 == nil {
			var hx_zero_64 bool
			return hx_zero_64
		}
		return hx_value_63.(bool)
	}(self.__hx_this.exists(key))
}

func (self *haxe__ds__BalancedTree) removeIMap(key any) bool {
	return func(hx_value_65 any) bool {
		if hx_value_65 == nil {
			var hx_zero_66 bool
			return hx_zero_66
		}
		return hx_value_65.(bool)
	}(self.__hx_this.remove(key))
}

func (self *haxe__ds__BalancedTree) copyIMap() haxe__IMap {
	return hx_recover_class_1(self.__hx_this.copy())
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
