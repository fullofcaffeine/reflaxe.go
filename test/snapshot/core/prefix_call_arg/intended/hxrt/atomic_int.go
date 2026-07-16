package hxrt

import "sync/atomic"

type AtomicIntCell struct {
	value atomic.Int64
}

func AtomicIntNew(value int) *AtomicIntCell {
	cell := &AtomicIntCell{}
	cell.value.Store(int64(value))
	return cell
}

func AtomicIntLoad(cell *AtomicIntCell) int {
	return int(cell.value.Load())
}

func AtomicIntStore(cell *AtomicIntCell, value int) int {
	cell.value.Store(int64(value))
	return value
}

func AtomicIntExchange(cell *AtomicIntCell, value int) int {
	return int(cell.value.Swap(int64(value)))
}

func AtomicIntCompareExchange(cell *AtomicIntCell, expected int, replacement int) int {
	expectedValue := int64(expected)
	replacementValue := int64(replacement)
	for {
		previous := cell.value.Load()
		if previous != expectedValue {
			return int(previous)
		}
		if cell.value.CompareAndSwap(previous, replacementValue) {
			return int(previous)
		}
	}
}

func AtomicIntAdd(cell *AtomicIntCell, value int) int {
	delta := int64(value)
	return int(cell.value.Add(delta) - delta)
}

func AtomicIntSub(cell *AtomicIntCell, value int) int {
	delta := int64(value)
	return int(cell.value.Add(-delta) + delta)
}

func AtomicIntAnd(cell *AtomicIntCell, value int) int {
	mask := int64(value)
	for {
		previous := cell.value.Load()
		next := previous & mask
		if cell.value.CompareAndSwap(previous, next) {
			return int(previous)
		}
	}
}

func AtomicIntOr(cell *AtomicIntCell, value int) int {
	mask := int64(value)
	for {
		previous := cell.value.Load()
		next := previous | mask
		if cell.value.CompareAndSwap(previous, next) {
			return int(previous)
		}
	}
}

func AtomicIntXor(cell *AtomicIntCell, value int) int {
	mask := int64(value)
	for {
		previous := cell.value.Load()
		next := previous ^ mask
		if cell.value.CompareAndSwap(previous, next) {
			return int(previous)
		}
	}
}
