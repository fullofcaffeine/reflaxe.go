package hxrt

import (
	"sync"
)

type AtomicObjectCell struct {
	mu    sync.Mutex
	value any
}

func AtomicObjectNew(value any) *AtomicObjectCell {
	return &AtomicObjectCell{value: value}
}

func AtomicObjectLoad(cell *AtomicObjectCell) any {
	cell.mu.Lock()
	defer cell.mu.Unlock()
	return cell.value
}

func AtomicObjectStore(cell *AtomicObjectCell, value any) any {
	cell.mu.Lock()
	defer cell.mu.Unlock()
	cell.value = value
	return value
}

func AtomicObjectExchange(cell *AtomicObjectCell, value any) any {
	cell.mu.Lock()
	defer cell.mu.Unlock()
	previous := cell.value
	cell.value = value
	return previous
}

func AtomicObjectCompareExchange(cell *AtomicObjectCell, expected any, replacement any) any {
	cell.mu.Lock()
	defer cell.mu.Unlock()
	previous := cell.value
	if referenceEqual(previous, expected) {
		cell.value = replacement
	}
	return previous
}
