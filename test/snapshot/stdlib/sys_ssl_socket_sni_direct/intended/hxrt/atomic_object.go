package hxrt

import (
	"reflect"
	"sync"
)

type AtomicObjectCell struct {
	mu    sync.Mutex
	value any
}

func AtomicObjectNew(value any) any {
	return &AtomicObjectCell{value: value}
}

func AtomicObjectLoad(cell any) any {
	typed := cell.(*AtomicObjectCell)
	typed.mu.Lock()
	defer typed.mu.Unlock()
	return typed.value
}

func AtomicObjectStore(cell any, value any) any {
	typed := cell.(*AtomicObjectCell)
	typed.mu.Lock()
	defer typed.mu.Unlock()
	typed.value = value
	return value
}

func AtomicObjectExchange(cell any, value any) any {
	typed := cell.(*AtomicObjectCell)
	typed.mu.Lock()
	defer typed.mu.Unlock()
	previous := typed.value
	typed.value = value
	return previous
}

func AtomicObjectCompareExchange(cell any, expected any, replacement any) any {
	typed := cell.(*AtomicObjectCell)
	typed.mu.Lock()
	defer typed.mu.Unlock()
	previous := typed.value
	if atomicReferenceEqual(previous, expected) {
		typed.value = replacement
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
