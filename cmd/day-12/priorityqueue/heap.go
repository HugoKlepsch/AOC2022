package priorityqueue

import "golang.org/x/exp/constraints"

func NewHeapPriorityQueue[ET comparable, PT constraints.Ordered](maxHeap bool) *HeapPriorityQueue[ET, PT] {
	return &HeapPriorityQueue[ET, PT]{
		maxHeap: maxHeap,
		heap:    make([]node[ET, PT], 0),
	}
}

type node[ET comparable, PT constraints.Ordered] struct {
	e        ET
	priority PT
}

type HeapPriorityQueue[ET comparable, PT constraints.Ordered] struct {
	maxHeap bool
	heap    []node[ET, PT]
	// Invariant: end is the index after the last element in the heap
	end heapIndex
}

func (h *HeapPriorityQueue[ET, PT]) IsEmpty() bool {
	return h.end == 0
}

func (h *HeapPriorityQueue[ET, PT]) Size() int {
	return int(h.end)
}

func (h *HeapPriorityQueue[ET, PT]) Insert(e ET, priority PT) {
	n := node[ET, PT]{
		e:        e,
		priority: priority,
	}
	// Insert at the end of the heap
	heapLen := len(h.heap)
	if int(h.end) == heapLen {
		h.heap = append(h.heap, n)
	} else {
		h.heap[h.end] = n
	}
	nodeIndex := h.end
	h.end++

	// Bubble up
	h.bubbleUp(nodeIndex)
}

func (h *HeapPriorityQueue[ET, PT]) Pop() (ET, PT, error) {
	if h.IsEmpty() {
		var et ET
		var pt PT
		return et, pt, ErrEmpty
	}

	h.swap(0, h.end-1)
	h.end--
	h.bubbleDown(0)

	node := h.heap[h.end]
	return node.e, node.priority, nil
}

func (h *HeapPriorityQueue[ET, PT]) SetPriority(e ET, newPriority PT) error {
	for i, n := range h.heap {
		if n.e == e {
			oldPriority := n.priority
			n.priority = newPriority
			h.heap[i] = n
			if oldPriority > newPriority {
				if h.maxHeap {
					// Smaller priority must bubble down in max heap
					h.bubbleDown(heapIndex(i))
				} else {
					// Smaller priority must bubble up in min heap
					h.bubbleUp(heapIndex(i))
				}
			} else {
				if h.maxHeap {
					// Larger priority must bubble up in max heap
					h.bubbleUp(heapIndex(i))
				} else {
					// Larger priority must bubble down in min heap
					h.bubbleDown(heapIndex(i))
				}
			}
			return nil
		}
	}
	return ErrNotFound
}

func (h *HeapPriorityQueue[ET, PT]) bubbleUp(i heapIndex) {
	// Keep bubbling up the node until it is in heap order
	// (parent is lower than it in a min-heap, higher in a max-heap)
	for {
		p := i.Parent()
		if h.maxHeap {
			if h.heap[p].priority < h.heap[i].priority {
				h.swap(p, i)
				i = p
			} else {
				return
			}
		} else {
			if h.heap[p].priority > h.heap[i].priority {
				h.swap(p, i)
				i = p
			} else {
				return
			}
		}
	}
}

func (h *HeapPriorityQueue[ET, PT]) bubbleDown(i heapIndex) {
	// Keep bubbling the node down until in heap order
	// Children must be higher than current in a min-heap, and lower in a max-heap.
	// If above violated, swap with lower child in a min heap, or higher child in max heap. Continue
	// bubbling down if swapped, else end.
	for {
		var leadingChild, other heapIndex
		cl := i.ChildL()
		cr := i.ChildR()
		if cl >= h.end {
			// Both children don't exist
			return
		} else if cr >= h.end {
			// Only left child exists
			leadingChild = cl
		} else {
			// Both children exist
			// Assume min heap. Find the lowest child
			if h.heap[cl].priority < h.heap[cr].priority {
				leadingChild = cl
				other = cr
			} else {
				leadingChild = cr
				other = cl
			}
			if h.maxHeap {
				leadingChild = other
			}
		}

		// Compare current to leadingChild, and bubble down if necessary
		if h.maxHeap {
			if h.heap[i].priority < h.heap[leadingChild].priority {
				h.swap(i, leadingChild)
				i = leadingChild
			} else {
				return
			}
		} else {
			if h.heap[i].priority > h.heap[leadingChild].priority {
				h.swap(i, leadingChild)
				i = leadingChild
			} else {
				return
			}
		}
	}
}

func (h *HeapPriorityQueue[ET, PT]) swap(i, j heapIndex) {
	v := h.heap[i]
	h.heap[i] = h.heap[j]
	h.heap[j] = v
}

type heapIndex int

func (h heapIndex) Parent() heapIndex {
	return (h - 1) / 2
}

func (h heapIndex) ChildL() heapIndex {
	return (2 * h) + 1
}

func (h heapIndex) ChildR() heapIndex {
	return (2 * h) + 2
}
