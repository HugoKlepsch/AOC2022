package priorityqueue

import "errors"

var ErrEmpty = errors.New("priority queue empty")

type PriorityQueue[T any] interface {
	IsEmpty() bool
	Insert(e T, priority int64)
	Pop() (e T, priority int64, err error)
}
