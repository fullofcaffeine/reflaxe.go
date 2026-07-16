package hxrt

import (
	"reflect"
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
	if atomicReferenceEqual(previous, expected) {
		cell.value = replacement
	}
	return previous
}

func atomicReferenceEqual(left any, right any) bool {
	leftValue := reflect.ValueOf(left)
	rightValue := reflect.ValueOf(right)
	if !leftValue.IsValid() || !rightValue.IsValid() {
		return !leftValue.IsValid() && !rightValue.IsValid()
	}
	if leftValue.Type() != rightValue.Type() {
		return false
	}

	switch leftValue.Kind() {
	case reflect.Interface:
		if leftValue.IsNil() || rightValue.IsNil() {
			return leftValue.IsNil() && rightValue.IsNil()
		}
		return atomicReferenceEqual(leftValue.Elem().Interface(), rightValue.Elem().Interface())
	case reflect.Ptr, reflect.UnsafePointer, reflect.Map, reflect.Slice, reflect.Func, reflect.Chan:
		if leftValue.IsNil() || rightValue.IsNil() {
			return leftValue.IsNil() && rightValue.IsNil()
		}
		return leftValue.Pointer() == rightValue.Pointer()
	default:
		if leftValue.Type().Comparable() {
			return left == right
		}
		return false
	}
}
