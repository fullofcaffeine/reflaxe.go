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
	out(hxrt.StringFromLiteral("load.0"), nodeId(hx_recover_class_0(hxrt.AtomicObjectLoad(atom))))
	oldMiss := hx_recover_class_0(hxrt.AtomicObjectCompareExchange(atom, b, c))
	out(hxrt.StringFromLiteral("cmp.miss.old"), nodeId(oldMiss))
	out(hxrt.StringFromLiteral("cmp.miss.now"), nodeId(hx_recover_class_0(hxrt.AtomicObjectLoad(atom))))
	oldHit := hx_recover_class_0(hxrt.AtomicObjectCompareExchange(atom, a, c))
	out(hxrt.StringFromLiteral("cmp.hit.old"), nodeId(oldHit))
	out(hxrt.StringFromLiteral("cmp.hit.now"), nodeId(hx_recover_class_0(hxrt.AtomicObjectLoad(atom))))
	oldExchange := hx_recover_class_0(hxrt.AtomicObjectExchange(atom, d))
	out(hxrt.StringFromLiteral("xchg.old"), nodeId(oldExchange))
	out(hxrt.StringFromLiteral("xchg.now"), nodeId(hx_recover_class_0(hxrt.AtomicObjectLoad(atom))))
	stored := hx_recover_class_0(hxrt.AtomicObjectStore(atom, a))
	out(hxrt.StringFromLiteral("store.ret"), nodeId(stored))
	out(hxrt.StringFromLiteral("store.now"), nodeId(hx_recover_class_0(hxrt.AtomicObjectLoad(atom))))
	alias := hx_recover_class_0(hxrt.AtomicObjectLoad(atom))
	alias.id = hxrt.StringFromLiteral("a_mut")
	out(hxrt.StringFromLiteral("alias.now"), nodeId(hx_recover_class_0(hxrt.AtomicObjectLoad(atom))))
	oldAlias := hx_recover_class_0(hxrt.AtomicObjectCompareExchange(atom, alias, c))
	out(hxrt.StringFromLiteral("cmp.alias.old"), nodeId(oldAlias))
	out(hxrt.StringFromLiteral("cmp.alias.now"), nodeId(hx_recover_class_0(hxrt.AtomicObjectLoad(atom))))
}

func nodeId(value *Node) *string {
	var hx_if_1 *string
	if value == nil {
		hx_if_1 = hxrt.StringFromLiteral("null")
	} else {
		hx_if_1 = value.id
	}
	return hx_if_1
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

func (self *Node) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_Node(id *string) *Node {
	return hx_init_Node(hx_alloc_Node(nil), id)
}

func hx_alloc_Node(owner any) *Node {
	self := &Node{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_Node)
	}
	return self
}

func hx_init_Node(self *Node, id *string) *Node {
	self.id = id
	return self
}

func hx_recover_class_0(value any) *Node {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *Node:
		return source
	default:
		return value.(*Node)
	}
}
