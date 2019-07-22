package baseAdt

type TreeNode struct {
	Element int
	Left    *TreeNode
	Right   *TreeNode
}

func NewTreeNode() *TreeNode {
	return &TreeNode{}
}

func (t *TreeNode) MakeEmpty() {
	if t != nil {
		t.Left.MakeEmpty()
		t.Right.MakeEmpty()
		t.Element = -1
	}
}

func (t *TreeNode) Find(element int) int {
	if t == nil {
		return -1
	}

	if element < t.Element {
		return t.Left.Find(element)
	} else if element > t.Element {
		return t.Right.Find(element)
	} else {
		return t.Element
	}
}

func (t *TreeNode) FindMin() int {
	if t == nil {
		return -1
	} else if t.Left == nil {
		return t.Element
	} else {
		return t.Left.FindMin()
	}
}

func (t *TreeNode) FindMax() int {
	if t == nil {
		return -1
	}
	for t.Right != nil {
		t = t.Right
	}

	return t.Element
}
