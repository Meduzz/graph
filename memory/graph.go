package memory

import (
	"github.com/Meduzz/graph"
	"github.com/Meduzz/helper/fp/slice"
)

type (
	// Memory is an in memory implementation of graph that could be deserialized from json
	Memory[T comparable] []*graph.Edge[T]
)

// EmptyGraph creates an empty graph
func EmptyGraph[T comparable]() graph.Graph[T] {
	return Memory[T](make([]*graph.Edge[T], 0))
}

// Of creates a graph from existing edges
func Of[T comparable](edges []*graph.Edge[T]) Memory[T] {
	return Memory[T](edges)
}

func (g Memory[T]) Add(edge *graph.Edge[T]) graph.Graph[T] {
	return append(g, edge)
}

func (g Memory[T]) Remove(edge *graph.Edge[T]) graph.Graph[T] {
	result := slice.Filter(g, func(e *graph.Edge[T]) bool {
		return !e.Equals(edge)
	})

	return Memory[T](result)
}

func (g Memory[T]) Starts(node *graph.Node[T]) []*graph.Edge[T] {
	return slice.Filter(g, func(e *graph.Edge[T]) bool {
		return e.Start.Equals(node)
	})
}

func (g Memory[T]) Ends(node *graph.Node[T]) []*graph.Edge[T] {
	return slice.Filter(g, func(e *graph.Edge[T]) bool {
		return e.End.Equals(node)
	})
}

func (g Memory[T]) Edges() []*graph.Edge[T] {
	return g
}

func (g Memory[T]) Start() error {
	return nil
}
