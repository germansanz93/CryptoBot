package utils

type Node struct {
	value uint64
	next  *Node
}

func (node *Node) HasNext() bool {
	return node.next != nil
}

func (node *Node) Next() *Node {
	if node.HasNext() {
		return node.next
	}
	return nil
}
