package ringbuffer

import "errors"

var (
	ErrBufferFull  = errors.New("no room in buffer to enqueue element")
	ErrBufferEmpty = errors.New("buffer empty")
)

func New[T any](capacity int) *RingBuffer[T] {
	return &RingBuffer[T]{
		s:      make([]T, capacity),
		start:  0,
		end:    0,
		nItems: 0,
	}
}

type RingBuffer[T any] struct {
	s      []T
	start  int
	end    int
	nItems int
}

func (s RingBuffer[T]) index(i int) int {
	return i % cap(s.s)
}

func (s *RingBuffer[T]) Enqueue(element T) error {
	if s.nItems == cap(s.s) {
		return ErrBufferFull
	}
	s.s[s.end] = element
	s.end = s.index(s.end + 1)
	s.nItems++
	return nil
}

func (s *RingBuffer[T]) Dequeue() (T, error) {
	var t T
	if s.nItems == 0 {
		return t, ErrBufferEmpty
	}
	t = s.s[s.start]
	s.start = s.index(s.start + 1)
	s.nItems--
	return t, nil
}
