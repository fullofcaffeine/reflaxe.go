package hxrt

import (
	"sync"
	"testing"
	"unsafe"
)

func TestConcurrentThreadMessagesAndAtomics(t *testing.T) {
	const workers = 64
	const increments = 100

	parentID := ThreadCurrentId()
	intCell := AtomicIntNew(0)
	objectValues := make([]int, workers)
	objectCell := AtomicObjectNew(&objectValues[0])

	var atomicWorkers sync.WaitGroup
	atomicWorkers.Add(workers)
	for worker := 0; worker < workers; worker++ {
		worker := worker
		go func() {
			defer atomicWorkers.Done()
			for iteration := 0; iteration < increments; iteration++ {
				AtomicIntAdd(intCell, 1)
				previous := AtomicObjectLoad(objectCell)
				AtomicObjectCompareExchange(objectCell, previous, &objectValues[worker])
			}
		}()
	}
	atomicWorkers.Wait()

	if got, want := AtomicIntLoad(intCell), workers*increments; got != want {
		t.Fatalf("AtomicIntLoad() = %d, want %d", got, want)
	}
	if got := AtomicObjectLoad(objectCell); got == nil {
		t.Fatal("AtomicObjectLoad() returned nil after concurrent exchanges")
	}

	for worker := 0; worker < workers; worker++ {
		worker := worker
		if threadID := ThreadSpawn(func() {
			ThreadSendMessage(parentID, worker)
		}); threadID == 0 {
			t.Fatal("ThreadSpawn() returned the reserved zero thread ID")
		}
	}

	seen := make(map[int]bool, workers)
	for range workers {
		message, ok := ThreadReadMessage(true).(int)
		if !ok {
			t.Fatalf("ThreadReadMessage() returned a non-int message")
		}
		if seen[message] {
			t.Fatalf("ThreadReadMessage() returned duplicate worker %d", message)
		}
		seen[message] = true
	}
	if len(seen) != workers {
		t.Fatalf("received %d unique messages, want %d", len(seen), workers)
	}
}

func TestCheckptrReflectionPaths(t *testing.T) {
	first := 1
	second := 2
	firstPointer := unsafe.Pointer(&first)
	secondPointer := unsafe.Pointer(&second)
	pointerCell := AtomicObjectNew(firstPointer)

	if got := AtomicObjectCompareExchange(pointerCell, firstPointer, secondPointer); got != firstPointer {
		t.Fatalf("pointer compare/exchange returned %v, want %v", got, firstPointer)
	}
	if got := AtomicObjectLoad(pointerCell); got != secondPointer {
		t.Fatalf("pointer compare/exchange stored %v, want %v", got, secondPointer)
	}

	mapValue := map[string]int{"answer": 42}
	sliceValue := []int{1, 2, 3}
	channelValue := make(chan int)
	functionValue := func() {}
	for name, value := range map[string]any{
		"map":      mapValue,
		"slice":    sliceValue,
		"channel":  channelValue,
		"function": functionValue,
	} {
		if !referenceEqual(value, value) {
			t.Errorf("referenceEqual rejected identical %s reference", name)
		}
	}
	if referenceEqual(mapValue, map[string]int{"answer": 42}) {
		t.Fatal("referenceEqual treated distinct maps as the same reference")
	}
}
