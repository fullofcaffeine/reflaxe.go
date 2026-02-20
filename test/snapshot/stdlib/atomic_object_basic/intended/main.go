package main

import "snapshot/hxrt"

func main() {
	a := New_Node(hxrt.StringFromLiteral("a"))
	_ = a
	b := New_Node(hxrt.StringFromLiteral("a"))
	_ = b
	c := New_Node(hxrt.StringFromLiteral("c"))
	_ = c
	d := New_Node(hxrt.StringFromLiteral("d"))
	_ = d
	var atom any = haxe__atomic___AtomicObject__AtomicObject_Impl___new(a)
	_ = atom
	out(hxrt.StringFromLiteral("load.0"), nodeId(haxe__atomic___AtomicObject__AtomicObject_Impl__load(atom).(*Node)))
	oldMiss := haxe__atomic___AtomicObject__AtomicObject_Impl__compareExchange(atom, b, c).(*Node)
	_ = oldMiss
	out(hxrt.StringFromLiteral("cmp.miss.old"), nodeId(oldMiss))
	out(hxrt.StringFromLiteral("cmp.miss.now"), nodeId(haxe__atomic___AtomicObject__AtomicObject_Impl__load(atom).(*Node)))
	oldHit := haxe__atomic___AtomicObject__AtomicObject_Impl__compareExchange(atom, a, c).(*Node)
	_ = oldHit
	out(hxrt.StringFromLiteral("cmp.hit.old"), nodeId(oldHit))
	out(hxrt.StringFromLiteral("cmp.hit.now"), nodeId(haxe__atomic___AtomicObject__AtomicObject_Impl__load(atom).(*Node)))
	oldExchange := haxe__atomic___AtomicObject__AtomicObject_Impl__exchange(atom, d).(*Node)
	_ = oldExchange
	out(hxrt.StringFromLiteral("xchg.old"), nodeId(oldExchange))
	out(hxrt.StringFromLiteral("xchg.now"), nodeId(haxe__atomic___AtomicObject__AtomicObject_Impl__load(atom).(*Node)))
	stored := haxe__atomic___AtomicObject__AtomicObject_Impl__store(atom, a).(*Node)
	_ = stored
	out(hxrt.StringFromLiteral("store.ret"), nodeId(stored))
	out(hxrt.StringFromLiteral("store.now"), nodeId(haxe__atomic___AtomicObject__AtomicObject_Impl__load(atom).(*Node)))
	alias := haxe__atomic___AtomicObject__AtomicObject_Impl__load(atom).(*Node)
	_ = alias
	alias.id = hxrt.StringFromLiteral("a_mut")
	out(hxrt.StringFromLiteral("alias.now"), nodeId(haxe__atomic___AtomicObject__AtomicObject_Impl__load(atom).(*Node)))
	oldAlias := haxe__atomic___AtomicObject__AtomicObject_Impl__compareExchange(atom, alias, c).(*Node)
	_ = oldAlias
	out(hxrt.StringFromLiteral("cmp.alias.old"), nodeId(oldAlias))
	out(hxrt.StringFromLiteral("cmp.alias.now"), nodeId(haxe__atomic___AtomicObject__AtomicObject_Impl__load(atom).(*Node)))
}

func nodeId(value *Node) *string {
	var hx_if_1 *string
	if hxrt.StringEqualAny(value, nil) {
		hx_if_1 = hxrt.StringFromLiteral("null")
	} else {
		hx_if_1 = value.id
	}
	return hx_if_1
}

func out(label *string, value any) {
	hxrt.Println(hxrt.StringConcatAny(hxrt.StringConcatAny(label, hxrt.StringFromLiteral("=")), hxrt.StdString(value)))
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

func haxe__atomic___AtomicInt__AtomicInt_Impl___new(value int) any {
	return hxrt.AtomicIntNew(value)
}

func haxe__atomic___AtomicInt__AtomicInt_Impl__add(atom any, value int) int {
	return hxrt.AtomicIntAdd(atom, value)
}

func haxe__atomic___AtomicInt__AtomicInt_Impl__sub(atom any, value int) int {
	return hxrt.AtomicIntSub(atom, value)
}

func haxe__atomic___AtomicInt__AtomicInt_Impl__and(atom any, value int) int {
	return hxrt.AtomicIntAnd(atom, value)
}

func haxe__atomic___AtomicInt__AtomicInt_Impl__or(atom any, value int) int {
	return hxrt.AtomicIntOr(atom, value)
}

func haxe__atomic___AtomicInt__AtomicInt_Impl__xor(atom any, value int) int {
	return hxrt.AtomicIntXor(atom, value)
}

func haxe__atomic___AtomicInt__AtomicInt_Impl__compareExchange(atom any, expected int, replacement int) int {
	return hxrt.AtomicIntCompareExchange(atom, expected, replacement)
}

func haxe__atomic___AtomicInt__AtomicInt_Impl__exchange(atom any, value int) int {
	return hxrt.AtomicIntExchange(atom, value)
}

func haxe__atomic___AtomicInt__AtomicInt_Impl__load(atom any) int {
	return hxrt.AtomicIntLoad(atom)
}

func haxe__atomic___AtomicInt__AtomicInt_Impl__store(atom any, value int) int {
	return hxrt.AtomicIntStore(atom, value)
}

func haxe__atomic___AtomicObject__AtomicObject_Impl___new(value any) any {
	return hxrt.AtomicObjectNew(value)
}

func haxe__atomic___AtomicObject__AtomicObject_Impl__load(atom any) any {
	return hxrt.AtomicObjectLoad(atom)
}

func haxe__atomic___AtomicObject__AtomicObject_Impl__store(atom any, value any) any {
	return hxrt.AtomicObjectStore(atom, value)
}

func haxe__atomic___AtomicObject__AtomicObject_Impl__exchange(atom any, value any) any {
	return hxrt.AtomicObjectExchange(atom, value)
}

func haxe__atomic___AtomicObject__AtomicObject_Impl__compareExchange(atom any, expected any, replacement any) any {
	return hxrt.AtomicObjectCompareExchange(atom, expected, replacement)
}
