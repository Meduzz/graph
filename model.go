package graph

type (
	Node struct {
		Id interface{}
	}

	Relation struct {
		Type interface{}
	}

	Edge struct {
		Start *Node
		Relation *Relation
		End *Node
	}

	Graph interface {
	}
)
