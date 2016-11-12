package mem

import (
	"testing"
	"github.com/Meduzz/graph"
)

var (
	node1 = &graph.Node{Id:1}
	node2 = &graph.Node{Id:2}
	node3 = &graph.Node{Id:3}
	node4 = &graph.Node{Id:4}

	likes = "LIKES"

	edge1 = &graph.Edge{
		node1,
		likes,
		node2,
	}
	edge2 = &graph.Edge{
		node3,
		likes,
		node4,
	}
	edge3 = &graph.Edge{
		node4,
		likes,
		node1,
	}
)

func TestInternalAPI(t *testing.T) {
	g := EmptyGraph()
	g.Add(edge1)

	edges := start(g.(*InmemoryGraph).graph, node1)

	if !edgesContains(edges, edge1) {
		t.Error("[Start()] Edges did not contain Edge1.")
	}

	edges = end(g.(*InmemoryGraph).graph, node2)

	if !edgesContains(edges, edge1) {
		t.Error("[End()] Edges did not contain Edge1.")
	}

	edges = filter(g.(*InmemoryGraph).graph, likes)

	if !edgesContains(edges, edge1) {
		t.Error("[End()] Edges did not contain Edge1.")
	}
}

func TestPublicApi(t *testing.T) {
	g := GraphFromSlice([]*graph.Edge{edge1})

	nodes := g.RelationsStarting(node1, likes)

	if len(nodes) != 1 {
		t.Errorf("Nodes does not contain 1 node. (%s)", len(nodes))
	}

	for _, n := range(nodes) {
		nodes = append(nodes, g.RelationsEnding(n, likes)...)
	}

	if len(nodes) != 2 {
		t.Errorf("Nodes does not contain 2 elements. (%s)", len(nodes))
	}

	if !nodesContains(nodes, node2) {
		t.Error("Nodes did not contain to node2.")
	}

	if !nodesContains(nodes, node1) {
		t.Error("Nodes did not contain to node1.")
	}

	g.Add(edge2)
	g.Add(edge3)

	// This should force the graph impl, to puzzle together 2 slices.
	g.Remove(edge2)

	exists := g.RelationsEnding(node2, likes)
	exists = append(exists, g.RelationsStarting(node4, likes)...)

	if len(exists) != 2 {
		t.Errorf("Exists does not match expected length of 2. (%s)", exists)
	}

	if !nodesContains(exists, node1) {
		t.Error("Exists did not contain node1 after remove.")
	}
}

func edgesContains(haystack []*graph.Edge, needle *graph.Edge) bool {
	ret := false

	for _, elem := range(haystack) {
		if elem.Equals(needle) {
			ret = true
		}
	}

	return ret
}

func nodesContains(haystack []*graph.Node, needle *graph.Node) bool {
	ret := false

	for _, elem := range(haystack) {
		if elem.Equals(needle) {
			ret = true
		}
	}

	return ret
}