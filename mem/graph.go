package mem

import (
	"github.com/Meduzz/graph"
)

type InmemoryGraph struct {
	graph []*graph.Edge
}

func EmptyGraph() graph.Graph {
	return &InmemoryGraph{
		graph:make([]*graph.Edge, 0),
	}
}

func GraphFromSlice(edges []*graph.Edge) graph.Graph {
	return &InmemoryGraph{
		graph:edges,
	}
}

func (g *InmemoryGraph) Add(edge *graph.Edge) {
	g.graph = append(g.graph, edge)
}

func (g *InmemoryGraph) Remove(edge *graph.Edge) {
	for i, elem := range(g.graph) {
		if elem == edge {
			g.graph = append(g.graph[0:i], g.graph[i + 1: len(g.graph)]...)
		}
	}
}

func (g *InmemoryGraph) RelationsStarting(startNode *graph.Node, relation string) []*graph.Node {
	return endNodes(filter(start(g.graph, startNode), relation))
}

func (g *InmemoryGraph) RelationsEnding(endNode *graph.Node, relation string) []*graph.Node {
	return startNodes(filter(end(g.graph, endNode), relation))
}

func start(edges []*graph.Edge, node *graph.Node) []*graph.Edge {
	matched := make([]*graph.Edge, 0)
	for _, edge := range(edges) {
		if edge.Start.Equals(node) {
			matched = append(matched, edge)
		}
	}

	return matched
}

func end(edges []*graph.Edge, node *graph.Node) []*graph.Edge {
	matched := make([]*graph.Edge, 0)
	for _, edge := range(edges) {
		if edge.End.Equals(node) {
			matched = append(matched, edge)
		}
	}

	return matched
}

func filter(edges []*graph.Edge, relation string) []*graph.Edge {
	matched := make([]*graph.Edge, 0)
	for _, edge := range(edges) {
		if edge.Relation == relation {
			matched = append(matched, edge)
		}
	}

	return matched
}

func endNodes(edges []*graph.Edge) []*graph.Node {
	nodes := make([]*graph.Node, 0)

	for _, edge := range(edges) {
		nodes = append(nodes, edge.End)
	}

	return nodes
}

func startNodes(edges []*graph.Edge) []*graph.Node {
	nodes := make([]*graph.Node, 0)

	for _, edge := range(edges) {
		nodes = append(nodes, edge.Start)
	}

	return nodes
}
