package baseAdt

import (
	"testing"
)

// ab+cde+**
func TestTreeNode(t *testing.T) {
	tree := NewTreeNode()
	tree.Element = 2
	left := NewTreeNode()
	left.Element = 1
	right := NewTreeNode()
	right.Element = 3

	tree.Left = left
	tree.Right = right

	res := tree.Find(2)
	if res != 2 {
		t.Error("not find")
	}

	res = tree.FindMin()
	if res != 1 {
		t.Error("find min is wrong")
	}
	res = tree.FindMax()
	if res != 3 {
		t.Error("find max is wrong")
	}

	tree.MakeEmpty()
	if tree.Right.Element != -1 {
		t.Error(tree.Right)
	}
}
