package memory

import (
	"testing"

	"github.com/Meduzz/graph"
)

var (
	node1 = graph.NewNode(1, "")
	node2 = graph.NewNode(2, "")
	node3 = graph.NewNode(3, "")
	node4 = graph.NewNode(4, "")
	node5 = graph.NewNode(5, "")

	likes = "LIKES"

	edge1 = graph.NewEdge(node1, likes, node2)
	edge2 = graph.NewEdge(node3, likes, node4)
	edge3 = graph.NewEdge(node4, likes, node1)
	edge4 = graph.NewEdge(node2, likes, node5)
	edge5 = graph.NewEdge(node5, likes, node3)

	fullGraph = Memory[int]([]*graph.Edge[int]{edge1, edge2, edge3, edge4, edge5})

	Visit         = make(chan *graph.Edge[int], 10)
	SimpleVisitor = func(e *graph.Edge[int]) bool {
		Visit <- e
		degrees := len(fullGraph.Starts(e.End))

		return degrees > 0
	}
)

/*
	1 -> 2
	3 -> 4
	4 -> 1
	2 -> 5
	5 -> 3
*/

func TestAddNodes(t *testing.T) {
	subject := EmptyGraph[int]()
	subject = subject.Add(edge1)

	edges := subject.Edges()

	if len(edges) != 1 {
		t.Errorf("number of edges was not 1 but %d", len(edges))
	}
}

func TestRemoveEdge(t *testing.T) {
	subject := EmptyGraph[int]()
	subject = subject.Add(edge1)
	subject = subject.Add(edge2)

	if len(subject.Edges()) != 2 {
		t.Errorf("number of edges was not 2 but %d", len(subject.Edges()))
	}

	subject = subject.Remove(edge1)

	if len(subject.Edges()) != 1 {
		t.Errorf("number of edges was not 1 but %d", len(subject.Edges()))
	}
}

func TestEdgesEqual(t *testing.T) {
	if !edge1.Equals(edge1) {
		t.Error("edge did not equal it self")
	}

	if edge1.Equals(edge2) {
		t.Error("edge1 was equal to edge2")
	}
}

func TestNodesEqual(t *testing.T) {
	if !node1.Equals(node1) {
		t.Error("node did not equal it self")
	}

	if node1.Equals(node2) {
		t.Error("node1 was equal to node2")
	}
}

func TestBreadthFirst(t *testing.T) {
	graph.BreadthFirst[int](fullGraph, node1, SimpleVisitor)

	if len(Visit) != 5 {
		t.Errorf("visited log was not 1 but %d", len(Visit))
	}

	<-Visit
	<-Visit
	<-Visit
	<-Visit
	<-Visit
}

func TestDepthFirst(t *testing.T) {
	graph.DepthFirst[int](fullGraph, node1, SimpleVisitor)

	if len(Visit) != 5 {
		t.Errorf("visited log was not 5 but %d", len(Visit))
	}

	<-Visit
	<-Visit
	<-Visit
	<-Visit
	<-Visit
}

func TestStarts(t *testing.T) {
	result := fullGraph.Starts(node1)

	if len(result) != 1 {
		t.Errorf("result was not 1 but %d", len(result))
	}

	if !result[0].Equals(edge1) {
		t.Error("first edge was not edge1")
	}
}

func TestDegreenOut(t *testing.T) {
	subject := EmptyGraph[int]()

	rel := "a"

	subject = subject.Add(graph.NewEdge(node1, rel, node2))
	subject = subject.Add(graph.NewEdge(node1, rel, node3))

	count := graph.DegreesOutRelation(subject, node1, rel)

	if count != 2 {
		t.Errorf("count was not 2 but %d", count)
	}
}

func TestDegreesIn(t *testing.T) {
	subject := EmptyGraph[int]()

	rel := "a"

	subject = subject.Add(graph.NewEdge(node2, rel, node1))
	subject = subject.Add(graph.NewEdge(node3, rel, node1))

	count := graph.DegreesInRelation(subject, node1, rel)

	if count != 2 {
		t.Errorf("count was not 2 but %d", count)
	}
}
