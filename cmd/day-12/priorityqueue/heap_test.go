package priorityqueue

import (
	"errors"
	"math/rand"
	"testing"
)

func TestHeap(t *testing.T) {
	// New min heap
	h := NewHeapPriorityQueue[float32](false)

	nItems := 2000
	for i := 0; i < nItems; i++ {
		number := rand.Int() % 2000
		numberF := float32(number)
		t.Logf("Inserting %f\n", numberF)
		h.Insert(numberF, int64(number))
	}

	var lastNum float32 = float32(-1)
	for i := 0; i < nItems; i++ {
		if h.IsEmpty() {
			t.Fail()
		}
		number, priority, err := h.Pop()
		if err != nil {
			t.Fail()
		}
		if number != float32(priority) {
			t.Fail()
		}
		t.Logf("Received %f\n", number)
		if lastNum > number {
			// min heap returned a smaller number than before; it was not in order
			t.Fail()
		}
		lastNum = number
	}

	_, _, err := h.Pop()
	if !errors.Is(err, ErrEmpty) {
		t.Fail()
	}
}
