package main

import "snapshot/hxrt"

func main() {
	var atom any = haxe__atomic___AtomicInt__AtomicInt_Impl___new(10)
	_ = atom
	out(hxrt.StringFromLiteral("int.load.0"), haxe__atomic___AtomicInt__AtomicInt_Impl__load(atom))
	out(hxrt.StringFromLiteral("int.add.old"), haxe__atomic___AtomicInt__AtomicInt_Impl__add(atom, 5))
	out(hxrt.StringFromLiteral("int.load.1"), haxe__atomic___AtomicInt__AtomicInt_Impl__load(atom))
	out(hxrt.StringFromLiteral("int.sub.old"), haxe__atomic___AtomicInt__AtomicInt_Impl__sub(atom, 2))
	out(hxrt.StringFromLiteral("int.load.2"), haxe__atomic___AtomicInt__AtomicInt_Impl__load(atom))
	out(hxrt.StringFromLiteral("int.and.old"), haxe__atomic___AtomicInt__AtomicInt_Impl__and(atom, 6))
	out(hxrt.StringFromLiteral("int.load.3"), haxe__atomic___AtomicInt__AtomicInt_Impl__load(atom))
	out(hxrt.StringFromLiteral("int.or.old"), haxe__atomic___AtomicInt__AtomicInt_Impl__or(atom, 8))
	out(hxrt.StringFromLiteral("int.load.4"), haxe__atomic___AtomicInt__AtomicInt_Impl__load(atom))
	out(hxrt.StringFromLiteral("int.xor.old"), haxe__atomic___AtomicInt__AtomicInt_Impl__xor(atom, 10))
	out(hxrt.StringFromLiteral("int.load.5"), haxe__atomic___AtomicInt__AtomicInt_Impl__load(atom))
	out(hxrt.StringFromLiteral("int.cmp.miss.old"), haxe__atomic___AtomicInt__AtomicInt_Impl__compareExchange(atom, 7, 100))
	out(hxrt.StringFromLiteral("int.cmp.miss.now"), haxe__atomic___AtomicInt__AtomicInt_Impl__load(atom))
	out(hxrt.StringFromLiteral("int.cmp.hit.old"), haxe__atomic___AtomicInt__AtomicInt_Impl__compareExchange(atom, 6, 11))
	out(hxrt.StringFromLiteral("int.cmp.hit.now"), haxe__atomic___AtomicInt__AtomicInt_Impl__load(atom))
	out(hxrt.StringFromLiteral("int.xchg.old"), haxe__atomic___AtomicInt__AtomicInt_Impl__exchange(atom, 3))
	out(hxrt.StringFromLiteral("int.xchg.now"), haxe__atomic___AtomicInt__AtomicInt_Impl__load(atom))
	out(hxrt.StringFromLiteral("int.store.ret"), haxe__atomic___AtomicInt__AtomicInt_Impl__store(atom, 42))
	out(hxrt.StringFromLiteral("int.store.now"), haxe__atomic___AtomicInt__AtomicInt_Impl__load(atom))
	var this1 any
	_ = this1
	this1 = haxe__atomic___AtomicInt__AtomicInt_Impl___new(0)
	var flag any = this1
	_ = flag
	out(hxrt.StringFromLiteral("bool.load.0"), func() bool {
		v := haxe__atomic___AtomicInt__AtomicInt_Impl__load(flag)
		_ = v
		return (v == 1)
	}())
	out(hxrt.StringFromLiteral("bool.cmp.miss.old"), func() bool {
		v_1 := haxe__atomic___AtomicInt__AtomicInt_Impl__compareExchange(flag, 1, 0)
		_ = v_1
		return (v_1 == 1)
	}())
	out(hxrt.StringFromLiteral("bool.cmp.miss.now"), func() bool {
		v_2 := haxe__atomic___AtomicInt__AtomicInt_Impl__load(flag)
		_ = v_2
		return (v_2 == 1)
	}())
	out(hxrt.StringFromLiteral("bool.cmp.hit.old"), func() bool {
		v_3 := haxe__atomic___AtomicInt__AtomicInt_Impl__compareExchange(flag, 0, 1)
		_ = v_3
		return (v_3 == 1)
	}())
	out(hxrt.StringFromLiteral("bool.cmp.hit.now"), func() bool {
		v_4 := haxe__atomic___AtomicInt__AtomicInt_Impl__load(flag)
		_ = v_4
		return (v_4 == 1)
	}())
	out(hxrt.StringFromLiteral("bool.xchg.old"), func() bool {
		v_5 := haxe__atomic___AtomicInt__AtomicInt_Impl__exchange(flag, 0)
		_ = v_5
		return (v_5 == 1)
	}())
	out(hxrt.StringFromLiteral("bool.xchg.now"), func() bool {
		v_6 := haxe__atomic___AtomicInt__AtomicInt_Impl__load(flag)
		_ = v_6
		return (v_6 == 1)
	}())
	out(hxrt.StringFromLiteral("bool.store.ret"), func() bool {
		v_7 := haxe__atomic___AtomicInt__AtomicInt_Impl__store(flag, 1)
		_ = v_7
		return (v_7 == 1)
	}())
	out(hxrt.StringFromLiteral("bool.store.now"), func() bool {
		v_8 := haxe__atomic___AtomicInt__AtomicInt_Impl__load(flag)
		_ = v_8
		return (v_8 == 1)
	}())
}

func out(label *string, value any) {
	hxrt.Println(hxrt.StringConcatAny(hxrt.StringConcatAny(label, hxrt.StringFromLiteral("=")), hxrt.StdString(value)))
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
