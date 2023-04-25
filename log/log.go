package log

import "github.com/Meduzz/helper/fp/slice"

// Log the type used for BreadthFirst
type Log[T any] []T

// TODO create tests

func NewLog[T any]() Log[T] {
	return Log[T](make([]T, 0))
}

func Of[T any](in []T) Log[T] {
	return Log[T](in)
}

func (h Log[T]) Append(item T) Log[T] {
	return append(h, item)
}

func (h Log[T]) Take() (T, Log[T]) {
	return slice.Head(h), slice.Tail(h)
}
