package hxrt

import "sync/atomic"

type AtomicIntCell struct {
	value atomic.Int64
}

func AtomicIntNew(value int) any {
	cell := &AtomicIntCell{}
	cell.value.Store(int64(value))
	return cell
}

func AtomicIntLoad(cell any) int {
	typed := cell.(*AtomicIntCell)
	return int(typed.value.Load())
}

func AtomicIntStore(cell any, value int) int {
	typed := cell.(*AtomicIntCell)
	typed.value.Store(int64(value))
	return value
}

func AtomicIntExchange(cell any, value int) int {
	typed := cell.(*AtomicIntCell)
	return int(typed.value.Swap(int64(value)))
}

func AtomicIntCompareExchange(cell any, expected int, replacement int) int {
	typed := cell.(*AtomicIntCell)
	expectedValue := int64(expected)
	replacementValue := int64(replacement)
	for {
		previous := typed.value.Load()
		if previous != expectedValue {
			return int(previous)
		}
		if typed.value.CompareAndSwap(previous, replacementValue) {
			return int(previous)
		}
	}
}

func AtomicIntAdd(cell any, value int) int {
	typed := cell.(*AtomicIntCell)
	delta := int64(value)
	return int(typed.value.Add(delta) - delta)
}

func AtomicIntSub(cell any, value int) int {
	typed := cell.(*AtomicIntCell)
	delta := int64(value)
	return int(typed.value.Add(-delta) + delta)
}

func AtomicIntAnd(cell any, value int) int {
	typed := cell.(*AtomicIntCell)
	mask := int64(value)
	for {
		previous := typed.value.Load()
		next := previous & mask
		if typed.value.CompareAndSwap(previous, next) {
			return int(previous)
		}
	}
}

func AtomicIntOr(cell any, value int) int {
	typed := cell.(*AtomicIntCell)
	mask := int64(value)
	for {
		previous := typed.value.Load()
		next := previous | mask
		if typed.value.CompareAndSwap(previous, next) {
			return int(previous)
		}
	}
}

func AtomicIntXor(cell any, value int) int {
	typed := cell.(*AtomicIntCell)
	mask := int64(value)
	for {
		previous := typed.value.Load()
		next := previous ^ mask
		if typed.value.CompareAndSwap(previous, next) {
			return int(previous)
		}
	}
}
