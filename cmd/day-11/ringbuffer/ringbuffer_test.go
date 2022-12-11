package ringbuffer

import (
	"errors"
	"testing"
)

func TestNew(t *testing.T) {
	rb := New[int](4)
	if len(rb.s) != 4 || cap(rb.s) != 4 {
		t.Fail()
	}
	if rb.start != 0 || rb.end != 0 || rb.nItems != 0 {
		t.Fail()
	}
}

func TestEnqueue(t *testing.T) {
	var err error
	rb := New[int](4)
	err = rb.Enqueue(1)
	if err != nil {
		t.Fail()
	}
	err = rb.Enqueue(2)
	if err != nil {
		t.Fail()
	}
	err = rb.Enqueue(3)
	if err != nil {
		t.Fail()
	}
	err = rb.Enqueue(4)
	if err != nil {
		t.Fail()
	}
	err = rb.Enqueue(5) // No capacity
	if !errors.Is(err, ErrBufferFull) {
		t.Fail()
	}
}

func TestDequeue(t *testing.T) {

	var err error
	rb := New[int](4)

	_, err = rb.Dequeue() // Empty, should fail
	if !errors.Is(err, ErrBufferEmpty) {
		t.Fail()
	}

	// Add some items
	err = rb.Enqueue(1)
	if err != nil {
		t.Fail()
	}
	err = rb.Enqueue(2)
	if err != nil {
		t.Fail()
	}

	// Dequeue them in order
	var i int
	i, err = rb.Dequeue()
	if err != nil || i != 1 {
		t.Fail()
	}
	err = rb.Enqueue(3)
	if err != nil {
		t.Fail()
	}
	i, err = rb.Dequeue()
	if err != nil || i != 2 {
		t.Fail()
	}
	i, err = rb.Dequeue()
	if err != nil || i != 3 {
		t.Fail()
	}
	_, err = rb.Dequeue() // Now empty
	if !errors.Is(err, ErrBufferEmpty) {
		t.Fail()
	}
}