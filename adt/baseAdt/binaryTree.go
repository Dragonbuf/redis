package baseAdt

type TreeNode struct {
	Element int
	Left    *TreeNode
	Right   *TreeNode
}

func NewTreeNode() *TreeNode {
	return &TreeNode{Element: -1}
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

func (t *TreeNode) FindMin() *TreeNode {
	if t == nil {
		return nil
	} else if t.Left == nil {
		return t
	} else {
		return t.Left.FindMin()
	}
}

func (t *TreeNode) FindMax() *TreeNode {
	if t == nil {
		return nil
	}
	for t.Right != nil {
		t = t.Right
	}

	return t
}

func (t *TreeNode) Insert(element int) *TreeNode {
	if t.Element == -1 {
		t.Element = element
	} else if element < t.Left.Element {
		t.Left = t.Left.Insert(element)
	} else {
		t.Right = t.Right.Insert(element)
	}

	return t
}

func (t *TreeNode) Delete(element int) *TreeNode {
	temp := NewTreeNode()
	if t.Element == -1 {
		panic("tree node is empty")
	} else if element < t.Element {
		t.Left = t.Left.Delete(element)
	} else if element > t.Element {
		t.Right = t.Right.Delete(element)
	} else if t.Left != nil && t.Right != nil {
		temp = t.Right.FindMin()
		t.Element = temp.Element
		t.Right = t.Right.Delete(temp.Element)
	} else {
		temp = t
		if t.Left == nil {
			t = t.Right
		} else if t.Right == nil {
			t = t.Left
		}
	}

	return t
}
