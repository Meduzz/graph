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
	}
)

func (n *Node) Equals(other *Node) bool {
	return n.Id == other.Id
}

func (e *Edge) Equals(other *Edge) bool {
	return e.Start.Equals(other.Start) && e.End.Equals(other.End) && e.Relation == other.Relation
}