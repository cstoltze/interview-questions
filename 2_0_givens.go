package main

/// A node of a singly linked list
type Node struct {
	/// the next node in the linked list
	next *Node
	data int
}

/// The length of a sigly linked list
func (n *Node) Len() int {
	if n.next == nil {
		return 1
	}
	return n.next.Len() + 1
}

func NodeFromList(values []int) *Node {
	var next *Node
	if len(values) > 1 {
		next = NodeFromList(values[1:])
	}
	return &Node{
		data: values[0],
		next: next,
	}
}
