package priorityqueue

import (
	"errors"
	"math"
	"math/rand"
	"testing"
)

func testHeap(nItems int, maxHeap bool) *HeapPriorityQueue[float32, int64] {
	// New heap, either min or max
	h := NewHeapPriorityQueue[float32, int64](maxHeap)

	for i := 0; i < nItems; i++ {
		number := rand.Int() % 2000
		numberF := float32(number)
		h.Insert(numberF, int64(number))
	}
	return h
}

func doTestHeap(maxHeap bool, t *testing.T) {
	// New heap, either min or max
	nItems := 2000
	h := testHeap(nItems, maxHeap)

	if h.Size() != nItems {
		t.Errorf("heap size: %d, expected %d\n", h.Size(), nItems)
	}

	var lastPriority int64
	if maxHeap {
		lastPriority = math.MaxInt64
	} else {
		lastPriority = math.MinInt64
	}

	for i := 0; i < nItems; i++ {
		if h.IsEmpty() {
			t.Errorf("ran out of items in heap")
		}
		number, priority, err := h.Pop()
		if err != nil {
			t.Error(err)
		}
		if number != float32(priority) {
			t.Errorf("element contents: %f, expected %f\n", number, float32(priority))
		}
		if maxHeap {
			if priority > lastPriority {
				// max heap returned a larger number than what should have been the largest; it was not in order
				t.Errorf("max heap not in order; previous: %d, current: %d\n", lastPriority, priority)
			}
		} else {
			if priority < lastPriority {
				// min heap returned a smaller number than what should have been the lowest; it was not in order
				t.Errorf("min heap not in order; previous: %d, current: %d\n", lastPriority, priority)
			}
		}
		lastPriority = priority
	}

	_, _, err := h.Pop()
	if !errors.Is(err, ErrEmpty) {
		t.Errorf("expected heap to be empty (ErrEmpty). Instead got %s", err.Error())
	}
}

func TestHeap(t *testing.T) {
	doTestHeap(false, t)
	doTestHeap(true, t)
}

func TestSize(t *testing.T) {
	size := 10
	h := testHeap(size, false)
	if h.Size() != size {
		t.Errorf("heap size: %d, expected %d\n", h.Size(), size)
	}
	size = 64
	h = testHeap(size, false)
	if h.Size() != size {
		t.Errorf("heap size: %d, expected %d\n", h.Size(), size)
	}
}

func TestSetPriority_minHeap(t *testing.T) {
	h := NewHeapPriorityQueue[float32, int64](false)
	for i := 0; i < 5; i++ {
		h.Insert(float32(i), int64(i))
	}

	var itemToChange float32 = 5.0
	var newPriority int64 = -1
	err := h.SetPriority(itemToChange, newPriority)
	if !errors.Is(err, ErrNotFound) {
		// 5.0 is not in the heap, we should get ErrNotFound
		t.Errorf("%f is not in the heap, we should have gotten ErrNotFound\n", itemToChange)
	}

	// Test lowering priority
	itemToChange = 4.0
	err = h.SetPriority(float32(itemToChange), newPriority)
	if err != nil {
		t.Error(err)
	}
	item, priority, err := h.Pop()
	if err != nil {
		t.Error(err)
	}
	if item != itemToChange {
		t.Errorf("Expected %f, got %f\n", itemToChange, item)
	}
	if priority != newPriority {
		t.Errorf("Expected priority %d, got %d\n", newPriority, priority)
	}

	// Test raise priority
	itemToChange = 0.0
	newPriority = 4
	err = h.SetPriority(float32(itemToChange), newPriority)
	if err != nil {
		t.Error(err)
	}

	// pop out the other items, 0.0 is now at the back of the heap
	_, _, _ = h.Pop()
	_, _, _ = h.Pop()
	_, _, _ = h.Pop()
	item, priority, err = h.Pop()
	if err != nil {
		t.Error(err)
	}
	if item != itemToChange {
		t.Errorf("Expected %f, got %f\n", itemToChange, item)
	}
	if priority != newPriority {
		t.Errorf("Expected priority %d, got %d\n", newPriority, priority)
	}
}

func TestSetPriority_maxHeap(t *testing.T) {
	h := NewHeapPriorityQueue[float32, int64](true)
	for i := 0; i < 5; i++ {
		h.Insert(float32(i), int64(i))
	}

	var itemToChange float32 = -1.0
	var newPriority int64 = 5
	err := h.SetPriority(itemToChange, newPriority)
	if !errors.Is(err, ErrNotFound) {
		// -1.0 is not in the heap, we should get ErrNotFound
		t.Errorf("%f is not in the heap, we should have gotten ErrNotFound\n", itemToChange)
	}

	// Test raising priority
	itemToChange = 0.0
	err = h.SetPriority(float32(itemToChange), newPriority)
	if err != nil {
		t.Error(err)
	}
	item, priority, err := h.Pop()
	if err != nil {
		t.Error(err)
	}
	if item != itemToChange {
		t.Errorf("Expected %f, got %f\n", itemToChange, item)
	}
	if priority != newPriority {
		t.Errorf("Expected priority %d, got %d\n", newPriority, priority)
	}

	// Test lowering priority
	itemToChange = 4.0
	newPriority = -1
	err = h.SetPriority(float32(itemToChange), newPriority)
	if err != nil {
		t.Error(err)
	}

	// pop out the other items, 0.0 is now at the back of the heap
	_, _, _ = h.Pop()
	_, _, _ = h.Pop()
	_, _, _ = h.Pop()
	item, priority, err = h.Pop()
	if err != nil {
		t.Error(err)
	}
	if item != itemToChange {
		t.Errorf("Expected %f, got %f\n", itemToChange, item)
	}
	if priority != newPriority {
		t.Errorf("Expected priority %d, got %d\n", newPriority, priority)
	}
}

func TestInterface(t *testing.T) {
	// Ensure that HeapPriorityQueue implements PriorityQueue
	var p PriorityQueue[float32, int64] = &HeapPriorityQueue[float32, int64]{}
	_ = p
}
