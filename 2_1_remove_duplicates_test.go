package main

import (
	"testing"
)

func TestNodeLen(t *testing.T) {
	node := Node{
		data: 1,
		next: &Node{
			data:1,
			next:nil,
		},
	}
	l := node.Len()
	if 2 != l {
		t.Errorf("length should have been 2, instead got %d", l)
	}
}

func nodeFromArray(arr []int) *Node {
	node := &Node{data: arr[0]}
	current := node
	for _, data := range arr[1:] {
		current.next = &Node{data:data}
		current = current.next
	}
	return node
}

func TestRemoveDuplicates(t *testing.T) {
	node := nodeFromArray([]int{1,1,1,1})
	RemoveDuplicates(node)
	if l := node.Len(); 1 != l {
		t.Errorf("length should have been 1, instead got %d", l)
	}

}

func TestRemoveDuplicatesWithoutTemporaryBuffer(t *testing.T) {
	node := nodeFromArray([]int{2,1,3,1})
	RemoveDuplicatesWithoutTemporaryBuffer(node)
	if l := node.Len(); 3 != l {
		t.Errorf("length should have been 1, instead got %d", l)
	}

}
