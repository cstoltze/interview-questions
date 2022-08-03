package main


import (
	"testing"
)

func TestKthToLast(t *testing.T) {
	testCases := []struct{
		node *Node
		k int
		want int
	}{
		{
			node: NodeFromList([]int{1,2,3,4,5}),
			k: 1,
			want: 5,
		},
		{
			node: NodeFromList([]int{1,2,3,4,5}),
			k: 5,
			want: 1,
		},
		{
			node: NodeFromList([]int{1,2,3,4,5}),
			k: 3,
			want: 3,
		},
	}
	for i, testCase := range(testCases) {
		gotNode, err := testCase.node.KthToLast(testCase.k)
		if err != nil {
			t.Errorf("unexpected error: %s", err.Error())
		}
		got := gotNode.data
		if got != testCase.want {
			t.Errorf("incorrect data for test case %d: wanted value %d, got %d", i+1, testCase.want, got)
		}
	}
}

func TestKthToLastErrorConditions(t *testing.T) {
	testCases := []struct{
		node *Node
		k int
		want int
	}{
		{
			node: NodeFromList([]int{1,2,3,4,5}),
			k: 122,
		},
		{
			node: NodeFromList([]int{1,2,3,4,5}),
			k: 0,
		},
		{
			node: NodeFromList([]int{1,2,3,4,5}),
			k: -4,
		},
	}
	for i, testCase := range(testCases) {
		gotNode, err := testCase.node.KthToLast(testCase.k)
		if err == nil {
			t.Errorf("test case %d should have produced an error", i+1)
		}
		if gotNode != nil {
			t.Errorf("unexpected non-nil return")
		}
	}
}
