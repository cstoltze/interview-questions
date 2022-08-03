package main
/// Implement an algorithm to find the kth to last element of a singly linked list

import (
	"errors"
)

/// Use a buffer to remember the previous nodes
func (n *Node) KthToLast(k int) (*Node, error) {
	if k <= 0 {
		return nil, errors.New("k must be greater than or equal to 1")
	}
	if k == 1 {
		return n.last(), nil
	}
	return n.kthToLast(0, make([]*Node, k-1))
}

func (n *Node) kthToLast(currentIndex int, previous []*Node) (*Node, error) {
	if n.next == nil {
		kthToLast := previous[currentIndex]
		if kthToLast == nil {
			return nil, errors.New("Invalid Argument: list is shorter than k elements")
		}
		return kthToLast, nil
	}
	previous[currentIndex] = n
	nextIndex := (currentIndex + 1) % len(previous)
	return n.next.kthToLast(nextIndex, previous)
}

func (n *Node) last() *Node {
	if n.next == nil {
		return n
	}
	return n.next.last()
}
