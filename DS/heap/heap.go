package heap

import (
	"errors"
	"mundhrakeshav/dsa/types"
)

type Heap[T any] struct {
	heaps    []T
	lessFunc func(a, b T) bool
}

func New[T types.Ordered]() *Heap[T] {
	less := func(a, b T) bool {
		return a < b
	}
	h, _ := NewAny(less)
	return h
}

// NewAny gives a new heap object. element can be anything, but must provide less function.
func NewAny[T any](less func(a, b T) bool) (*Heap[T], error) {
	if less == nil {
		return nil, errors.New("less func is necessary")
	}
	return &Heap[T]{
		lessFunc: less,
	}, nil
}

func (h *Heap[T]) Top() T {
	return h.heaps[0]
}

func (h *Heap[T]) Empty() bool {
	return len(h.heaps) == 0
}

// Size returns the size of the heap
func (h *Heap[T]) Size() int {
	return len(h.heaps)
}

func (h *Heap[T]) Push(t T) {
	h.heaps = append(h.heaps, t)
	i := len(h.heaps) - 1
	i_parent := ((i - 1) >> 1)
	for i > 0 && h.lessFunc(h.heaps[i_parent], h.heaps[i]) {
		h.swap(i_parent, i)
		i = i_parent
		i_parent = ((i - 1) >> 1)
	}
}

func (h *Heap[T]) swap(i, j int) {
	h.heaps[i], h.heaps[j] = h.heaps[j], h.heaps[i]
}

func (h *Heap[T]) Pop() {

	if len(h.heaps) <= 1 {
		h.heaps = nil
		return
	}
	h.swap(0, len(h.heaps)-1)
	h.heaps = h.heaps[:len(h.heaps)-1]
	h.maxHeapify(0)

}

func (h *Heap[T]) maxHeapify(i int) {
	largest := i;

	lChild, rChild := (largest << 1) + 1, (largest << 1) + 2

	if lChild < len(h.heaps) && h.lessFunc(h.heaps[largest], h.heaps[lChild]) {
		largest = lChild;
	}
	if rChild < len(h.heaps) && h.lessFunc(h.heaps[largest], h.heaps[rChild],) {
		largest = rChild;
	}
	if largest != i {
		h.swap(largest, i)
		h.maxHeapify(largest)
	}

}


