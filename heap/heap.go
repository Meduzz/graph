package heap

import "github.com/Meduzz/helper/fp/slice"

// Heap the type used for DepthFirst
type Heap[T any] []T

func NewHeap[T any]() Heap[T] {
	return Heap[T](make([]T, 0))
}

func Of[T any](in []T) Heap[T] {
	return Heap[T](in)
}

func (h Heap[T]) Push(item T) Heap[T] {
	return append([]T{item}, h...)
}

func (h Heap[T]) Pop() (T, Heap[T]) {
	return slice.Head(h), slice.Tail(h)
}
