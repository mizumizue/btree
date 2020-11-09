package btree

type Node struct {
	v      int
	parent *Node
	l      *Node
	r      *Node
}

func (s *Node) hasLeft() bool {
	return s.l != nil
}

func (s *Node) hasRight() bool {
	return s.r != nil
}
