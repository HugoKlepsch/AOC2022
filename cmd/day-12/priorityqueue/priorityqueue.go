package priorityqueue

import (
	"errors"

	"golang.org/x/exp/constraints"
)

var (
	ErrEmpty    = errors.New("priority queue empty")
	ErrNotFound = errors.New("element not found")
)

type PriorityQueue[ET comparable, PT constraints.Ordered] interface {
	IsEmpty() bool
	Size() int
	Insert(e ET, priority PT)
	Pop() (e ET, priority PT, err error)
	SetPriority(e ET, newPriority PT) error
}
