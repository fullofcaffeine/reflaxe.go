package main

import "snapshot/hxrt"

func main() {
	var this1 *hxrt.AtomicIntCell
	this1 = hxrt.AtomicIntNew(10)
	atom := this1
	out(hxrt.StringFromLiteral("int.load.0"), hxrt.AtomicIntLoad(atom))
	out(hxrt.StringFromLiteral("int.add.old"), hxrt.AtomicIntAdd(atom, 5))
	out(hxrt.StringFromLiteral("int.load.1"), hxrt.AtomicIntLoad(atom))
	out(hxrt.StringFromLiteral("int.sub.old"), hxrt.AtomicIntSub(atom, 2))
	out(hxrt.StringFromLiteral("int.load.2"), hxrt.AtomicIntLoad(atom))
	out(hxrt.StringFromLiteral("int.and.old"), hxrt.AtomicIntAnd(atom, 6))
	out(hxrt.StringFromLiteral("int.load.3"), hxrt.AtomicIntLoad(atom))
	out(hxrt.StringFromLiteral("int.or.old"), hxrt.AtomicIntOr(atom, 8))
	out(hxrt.StringFromLiteral("int.load.4"), hxrt.AtomicIntLoad(atom))
	out(hxrt.StringFromLiteral("int.xor.old"), hxrt.AtomicIntXor(atom, 10))
	out(hxrt.StringFromLiteral("int.load.5"), hxrt.AtomicIntLoad(atom))
	out(hxrt.StringFromLiteral("int.cmp.miss.old"), hxrt.AtomicIntCompareExchange(atom, 7, 100))
	out(hxrt.StringFromLiteral("int.cmp.miss.now"), hxrt.AtomicIntLoad(atom))
	out(hxrt.StringFromLiteral("int.cmp.hit.old"), hxrt.AtomicIntCompareExchange(atom, 6, 11))
	out(hxrt.StringFromLiteral("int.cmp.hit.now"), hxrt.AtomicIntLoad(atom))
	out(hxrt.StringFromLiteral("int.xchg.old"), hxrt.AtomicIntExchange(atom, 3))
	out(hxrt.StringFromLiteral("int.xchg.now"), hxrt.AtomicIntLoad(atom))
	out(hxrt.StringFromLiteral("int.store.ret"), hxrt.AtomicIntStore(atom, 42))
	out(hxrt.StringFromLiteral("int.store.now"), hxrt.AtomicIntLoad(atom))
	var this1_1 *hxrt.AtomicIntCell
	var this1_2 *hxrt.AtomicIntCell
	this1_2 = hxrt.AtomicIntNew(0)
	this1_1 = this1_2
	flag := this1_1
	out(hxrt.StringFromLiteral("bool.load.0"), func() bool {
		v := hxrt.AtomicIntLoad(flag)
		return (v == 1)
	}())
	out(hxrt.StringFromLiteral("bool.cmp.miss.old"), func() bool {
		v_1 := hxrt.AtomicIntCompareExchange(flag, 1, 0)
		return (v_1 == 1)
	}())
	out(hxrt.StringFromLiteral("bool.cmp.miss.now"), func() bool {
		v_2 := hxrt.AtomicIntLoad(flag)
		return (v_2 == 1)
	}())
	out(hxrt.StringFromLiteral("bool.cmp.hit.old"), func() bool {
		v_3 := hxrt.AtomicIntCompareExchange(flag, 0, 1)
		return (v_3 == 1)
	}())
	out(hxrt.StringFromLiteral("bool.cmp.hit.now"), func() bool {
		v_4 := hxrt.AtomicIntLoad(flag)
		return (v_4 == 1)
	}())
	out(hxrt.StringFromLiteral("bool.xchg.old"), func() bool {
		v_5 := hxrt.AtomicIntExchange(flag, 0)
		return (v_5 == 1)
	}())
	out(hxrt.StringFromLiteral("bool.xchg.now"), func() bool {
		v_6 := hxrt.AtomicIntLoad(flag)
		return (v_6 == 1)
	}())
	out(hxrt.StringFromLiteral("bool.store.ret"), func() bool {
		v_7 := hxrt.AtomicIntStore(flag, 1)
		return (v_7 == 1)
	}())
	out(hxrt.StringFromLiteral("bool.store.now"), func() bool {
		v_8 := hxrt.AtomicIntLoad(flag)
		return (v_8 == 1)
	}())
}

func out(label *string, value any) {
	var v any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(label, hxrt.StringFromLiteral("=")), hxrt.StdString(value)))
	hxrt.Println(v)
}
