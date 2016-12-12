package graph

import (
	"testing"
)

var (
	node1 = &Node{Id:1}
	node2 = &Node{Id:2}
	node3 = &Node{Id:3}
	node4 = &Node{Id:4}

	likes = "LIKES"

	edge1 = &Edge{
		node1,
		likes,
		node2,
	}
	edge2 = &Edge{
		node3,
		likes,
		node4,
	}
	edge3 = &Edge{
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
	g := GraphFromSlice([]*Edge{edge1})

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

	if g.DegreesOut(node1) > 1 {
		t.Error("Node1 had more than 1 outgoing relations.")
	}

	if g.DegreesIn(node1) > 1 {
		t.Error("Node1 had more than 1 incoming relations.")
	}

	if !relationsContains(g.RelationsOut(node1), "LIKES") {
		t.Error("Node1 did not have an outgoing relation of type LIKES.")
	}

	if !relationsContains(g.RelationsIn(node1), "LIKES") {
		t.Error("Node1 did not have an incoming relation of type LIKES.")
	}
}

func edgesContains(haystack []*Edge, needle *Edge) bool {
	ret := false

	for _, elem := range(haystack) {
		if elem.Equals(needle) {
			ret = true
		}
	}

	return ret
}

func nodesContains(haystack []*Node, needle *Node) bool {
	ret := false

	for _, elem := range(haystack) {
		if elem.Equals(needle) {
			ret = true
		}
	}

	return ret
}

func relationsContains(haystack []string, needle string) bool {
	ret := false

	for _, rel := range(haystack) {
		if rel == needle {
			ret = true
		}
	}

	return ret

}