package main

import "snapshot/hxrt"

func main() {
	a := New_Node(hxrt.StringFromLiteral("a"))
	b := New_Node(hxrt.StringFromLiteral("a"))
	c := New_Node(hxrt.StringFromLiteral("c"))
	d := New_Node(hxrt.StringFromLiteral("d"))
	var this1 *hxrt.AtomicObjectCell
	this1 = hxrt.AtomicObjectNew(a)
	atom := this1
	out(hxrt.StringFromLiteral("load.0"), nodeId(func(hx_value_1 any) *Node {
		if hx_value_1 == nil {
			var hx_zero_2 *Node
			return hx_zero_2
		}
		return hx_value_1.(*Node)
	}(hxrt.AtomicObjectLoad(atom))))
	oldMiss := func(hx_value_3 any) *Node {
		if hx_value_3 == nil {
			var hx_zero_4 *Node
			return hx_zero_4
		}
		return hx_value_3.(*Node)
	}(hxrt.AtomicObjectCompareExchange(atom, b, c))
	out(hxrt.StringFromLiteral("cmp.miss.old"), nodeId(oldMiss))
	out(hxrt.StringFromLiteral("cmp.miss.now"), nodeId(func(hx_value_5 any) *Node {
		if hx_value_5 == nil {
			var hx_zero_6 *Node
			return hx_zero_6
		}
		return hx_value_5.(*Node)
	}(hxrt.AtomicObjectLoad(atom))))
	oldHit := func(hx_value_7 any) *Node {
		if hx_value_7 == nil {
			var hx_zero_8 *Node
			return hx_zero_8
		}
		return hx_value_7.(*Node)
	}(hxrt.AtomicObjectCompareExchange(atom, a, c))
	out(hxrt.StringFromLiteral("cmp.hit.old"), nodeId(oldHit))
	out(hxrt.StringFromLiteral("cmp.hit.now"), nodeId(func(hx_value_9 any) *Node {
		if hx_value_9 == nil {
			var hx_zero_10 *Node
			return hx_zero_10
		}
		return hx_value_9.(*Node)
	}(hxrt.AtomicObjectLoad(atom))))
	oldExchange := func(hx_value_11 any) *Node {
		if hx_value_11 == nil {
			var hx_zero_12 *Node
			return hx_zero_12
		}
		return hx_value_11.(*Node)
	}(hxrt.AtomicObjectExchange(atom, d))
	out(hxrt.StringFromLiteral("xchg.old"), nodeId(oldExchange))
	out(hxrt.StringFromLiteral("xchg.now"), nodeId(func(hx_value_13 any) *Node {
		if hx_value_13 == nil {
			var hx_zero_14 *Node
			return hx_zero_14
		}
		return hx_value_13.(*Node)
	}(hxrt.AtomicObjectLoad(atom))))
	stored := func(hx_value_15 any) *Node {
		if hx_value_15 == nil {
			var hx_zero_16 *Node
			return hx_zero_16
		}
		return hx_value_15.(*Node)
	}(hxrt.AtomicObjectStore(atom, a))
	out(hxrt.StringFromLiteral("store.ret"), nodeId(stored))
	out(hxrt.StringFromLiteral("store.now"), nodeId(func(hx_value_17 any) *Node {
		if hx_value_17 == nil {
			var hx_zero_18 *Node
			return hx_zero_18
		}
		return hx_value_17.(*Node)
	}(hxrt.AtomicObjectLoad(atom))))
	alias := func(hx_value_19 any) *Node {
		if hx_value_19 == nil {
			var hx_zero_20 *Node
			return hx_zero_20
		}
		return hx_value_19.(*Node)
	}(hxrt.AtomicObjectLoad(atom))
	alias.id = hxrt.StringFromLiteral("a_mut")
	out(hxrt.StringFromLiteral("alias.now"), nodeId(func(hx_value_21 any) *Node {
		if hx_value_21 == nil {
			var hx_zero_22 *Node
			return hx_zero_22
		}
		return hx_value_21.(*Node)
	}(hxrt.AtomicObjectLoad(atom))))
	oldAlias := func(hx_value_23 any) *Node {
		if hx_value_23 == nil {
			var hx_zero_24 *Node
			return hx_zero_24
		}
		return hx_value_23.(*Node)
	}(hxrt.AtomicObjectCompareExchange(atom, alias, c))
	out(hxrt.StringFromLiteral("cmp.alias.old"), nodeId(oldAlias))
	out(hxrt.StringFromLiteral("cmp.alias.now"), nodeId(func(hx_value_25 any) *Node {
		if hx_value_25 == nil {
			var hx_zero_26 *Node
			return hx_zero_26
		}
		return hx_value_25.(*Node)
	}(hxrt.AtomicObjectLoad(atom))))
}

func nodeId(value *Node) *string {
	var hx_if_27 *string
	if value == nil {
		hx_if_27 = hxrt.StringFromLiteral("null")
	} else {
		hx_if_27 = value.id
	}
	return hx_if_27
}

func out(label *string, value any) {
	var v any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(label, hxrt.StringFromLiteral("=")), hxrt.StdString(value)))
	hxrt.Println(v)
}

type I_Node interface {
}

type Node struct {
	__hx_this I_Node
	id        *string
}

func New_Node(id *string) *Node {
	self := &Node{}
	self.__hx_this = self
	self.id = id
	return self
}
