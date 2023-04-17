package utils

type Node struct {
	value float64
	next  *Node
}

func (node *Node) Value() float64 {
	return node.value
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
