package priorityqueue

import "errors"

var (
	ErrEmpty    = errors.New("priority queue empty")
	ErrNotFound = errors.New("element not found")
)

type PriorityQueue[T any] interface {
	IsEmpty() bool
	Size() int
	Insert(e T, priority int64)
	Pop() (e T, priority int64, err error)
	SetPriority(e T, newPriority int64) error
}
