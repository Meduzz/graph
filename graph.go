package graph

type InmemoryGraph struct {
	graph []*Edge
}

func EmptyGraph() Graph {
	return &InmemoryGraph{
		graph:make([]*Edge, 0),
	}
}

func GraphFromSlice(edges []*Edge) Graph {
	return &InmemoryGraph{
		graph:edges,
	}
}

func (g *InmemoryGraph) Add(edge *Edge) {
	g.graph = append(g.graph, edge)
}

func (g *InmemoryGraph) Remove(edge *Edge) {
	for i, elem := range(g.graph) {
		if elem == edge {
			g.graph = append(g.graph[0:i], g.graph[i + 1: len(g.graph)]...)
		}
	}
}

func (g *InmemoryGraph) RelationsStarting(startNode *Node, relation string) []*Node {
	return endNodes(filter(start(g.graph, startNode), relation))
}

func (g *InmemoryGraph) RelationsEnding(endNode *Node, relation string) []*Node {
	return startNodes(filter(end(g.graph, endNode), relation))
}

func start(edges []*Edge, node *Node) []*Edge {
	matched := make([]*Edge, 0)
	for _, edge := range(edges) {
		if edge.Start.Equals(node) {
			matched = append(matched, edge)
		}
	}

	return matched
}

func end(edges []*Edge, node *Node) []*Edge {
	matched := make([]*Edge, 0)
	for _, edge := range(edges) {
		if edge.End.Equals(node) {
			matched = append(matched, edge)
		}
	}

	return matched
}

func filter(edges []*Edge, relation string) []*Edge {
	matched := make([]*Edge, 0)
	for _, edge := range(edges) {
		if edge.Relation == relation {
			matched = append(matched, edge)
		}
	}

	return matched
}

func endNodes(edges []*Edge) []*Node {
	nodes := make([]*Node, 0)

	for _, edge := range(edges) {
		nodes = append(nodes, edge.End)
	}

	return nodes
}

func startNodes(edges []*Edge) []*Node {
	nodes := make([]*Node, 0)

	for _, edge := range(edges) {
		nodes = append(nodes, edge.Start)
	}

	return nodes
}
