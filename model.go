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
