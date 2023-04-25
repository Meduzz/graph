package graph

import (
	"github.com/Meduzz/graph/log"

	"github.com/Meduzz/graph/heap"
	"github.com/Meduzz/helper/fp/slice"
)

type (
	Node[T comparable] struct {
		ID   T `json:"id"`
		Type string
	}

	Edge[T comparable] struct {
		Start    *Node[T] `json:"start"`
		Relation string   `json:"relation"`
		End      *Node[T] `json:"end"`
	}

	Graph[T comparable] interface {
		// Add adds an edge to the graph returning the new graph
		Add(edge *Edge[T]) Graph[T]
		// Remove removes an edge from the graph returning the new graph
		Remove(edge *Edge[T]) Graph[T]
		// Starts fetches a list of edges that starts with node start
		Starts(start *Node[T]) []*Edge[T]
		// Ends fetches a list of edges that ends with node end
		Ends(end *Node[T]) []*Edge[T]
		// Edges returns all edges in the graph
		Edges() []*Edge[T]
		// Build indexes? Load data? return errors
		Start() error
	}
)

// Traverse the graph specified in g, in direction specified by param forward from node n.
func Traverse[T comparable](g Graph[T], n *Node[T], forward bool, visitor func(e *Edge[T])) {
	var edges []*Edge[T]

	if forward {
		edges = g.Starts(n)
	} else {
		edges = g.Ends(n)
	}

	slice.ForEach(edges, func(it *Edge[T]) {
		visitor(it)
	})
}

// DegreesOutRelation from node n, count relations of type relation in graph g
func DegreesOutRelation[T comparable](g Graph[T], n *Node[T], relation string) int {
	count := 0

	Traverse(g, n, true, func(e *Edge[T]) {
		if e.Relation == relation {
			count++
		}
	})

	return count
}

// DegreesOutRelation to node n, count relations of type relation in graph g
func DegreesInRelation[T comparable](g Graph[T], n *Node[T], relation string) int {
	count := 0

	Traverse(g, n, false, func(e *Edge[T]) {
		if e.Relation == relation {
			count++
		}
	})

	return count
}

// BreadthFirst traverse the graph breadth first, continuing with any edges that returned true
func BreadthFirst[T comparable](g Graph[T], start *Node[T], visitor func(*Edge[T]) bool) {
	edges := log.NewLog[*Edge[T]]()
	visited := make([]*Node[T], 0)
	visited = append(visited, start)

	starters := g.Starts(start)

	slice.ForEach(starters, func(e *Edge[T]) {
		edges = edges.Append(e)
	})

	if len(edges) == 0 {
		return
	}

	var it *Edge[T]

	for {
		it, edges = edges.Take()
		if visitor(it) {
			next := g.Starts(it.End)

			for _, n := range next {
				if !slice.Contains(visited, n.Start) {
					edges = edges.Append(n)
					visited = append(visited, n.Start)
				}
			}
		}

		if len(edges) == 0 {
			return
		}
	}
}

// DepthFirst traverse the graph depth first, diving into any edges that returned true
func DepthFirst[T comparable](g Graph[T], start *Node[T], visitor func(*Edge[T]) bool) {
	edges := g.Starts(start)
	work := heap.Heap[*Edge[T]](edges)
	visited := make([]*Node[T], 0)
	visited = append(visited, start)

	var e *Edge[T]

	for {
		e, work = work.Pop()

		if visitor(e) {
			added := g.Starts(e.End)

			slice.ForEach(added, func(edge *Edge[T]) {
				if !slice.Contains(visited, edge.Start) {
					work = work.Push(edge)
					visited = append(visited, edge.Start)
				}
			})
		}

		if len(work) == 0 {
			break
		}
	}
}

// NewNode creates a new node
func NewNode[T comparable](id T, typ string) *Node[T] {
	return &Node[T]{
		ID:   id,
		Type: typ,
	}
}

// NewEdge creates a new edge
func NewEdge[T comparable](start *Node[T], relation string, end *Node[T]) *Edge[T] {
	return &Edge[T]{start, relation, end}
}

// Equals can compare 2 nodes
func (n *Node[T]) Equals(other interface{}) bool {
	node, ok := other.(*Node[T])

	if !ok {
		return false
	}

	return n.ID == node.ID && n.Type == node.Type
}

// Equals can compare 2 edges
func (e *Edge[T]) Equals(other interface{}) bool {
	edge, ok := other.(*Edge[T])

	if !ok {
		return false
	}

	return e.Start.Equals(edge.Start) && e.End.Equals(edge.End) && e.Relation == edge.Relation
}
