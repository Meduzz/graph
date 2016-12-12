package graph

type (
	Node struct {
		Id interface{}
	}

	Edge struct {
		Start *Node
		Relation string
		End *Node
	}

	Graph interface {
		Add(edge *Edge)
		Remove(edge *Edge)
		RelationsStarting(start *Node, relation string) []*Node
		RelationsEnding(end *Node, relation string) []*Node
		DegreesOut(start *Node) int
		DegreesOutRelation(start *Node, relation string) int
		DegreesIn(end *Node) int
		DegreesInRelation(end *Node, relation string) int
		RelationsOut(startNode *Node) []string
		RelationsIn(endNode *Node) []string
	}
)

func NewNode(id interface{}) *Node {
	return &Node{Id:id}
}

func NewEdge(start *Node, relation string, end *Node) *Edge {
	return &Edge{start, relation, end}
}

func (n *Node) Equals(other *Node) bool {
	return n.Id == other.Id
}

func (e *Edge) Equals(other *Edge) bool {
	return e.Start.Equals(other.Start) && e.End.Equals(other.End) && e.Relation == other.Relation
}