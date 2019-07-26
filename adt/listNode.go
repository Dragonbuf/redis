package adt

type ListNode struct {
	Prev  *ListNode
	Next  *ListNode
	Value *Sdshdr
}

func NewListNode() *ListNode {
	return &ListNode{}
}

func (l *ListNode) SetValue(value *Sdshdr) *ListNode {
	l.Value = value
	return l
}

func (l *ListNode) SetPrev(prev *ListNode) *ListNode {
	l.Prev = prev
	return l
}

func (l *ListNode) SetNext(next *ListNode) *ListNode {
	l.Next = next
	return l
}
