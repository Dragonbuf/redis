package adt

type List struct {
	head *ListNode
	tail *ListNode
	len  uint64
}

func NewList() *List {
	return &List{}
}

func (l *List) SetHead(head *ListNode) *List {
	l.head = head
	return l
}
func (l *List) SetTail(tail *ListNode) *List {
	l.tail = tail
	return l
}

func (l *List) LPush(value *string) int {

	sds := NewSdsHdr()
	sds.Set(value)

	l.IncrLen()

	if l.HasOneNode() {
		node := NewListNode().SetValue(sds)
		l.SetHead(node).SetTail(node)
	} else {
		node := NewListNode().SetNext(l.head).SetValue(sds)
		l.SetHead(node)
	}

	return 1
}

func (l *List) LPop() *string {
	node := l.head
	if node == nil {
		return nil
	}

	l.SetHead(node.Next)
	l.decrLen()

	if l.IsEmpty() {
		l.SetTail(nil)
	}

	str := node.Value.Get()
	return &str
}

func (l *List) RPush(value *string) int {

	sds := NewSdsHdr()
	sds.Set(value)

	l.IncrLen()
	if l.HasOneNode() {
		node := NewListNode().SetValue(sds)
		l.SetHead(node).SetTail(node)
	} else {
		node := NewListNode().SetValue(sds).SetPrev(l.tail)
		l.SetTail(node)
	}

	return 1
}

func (l *List) RPop() *string {
	node := l.tail
	if node == nil {
		return nil
	}

	l.SetTail(node.Prev)
	l.decrLen()

	str := node.Value.Get()
	return &str
}

func (l *List) IncrLen() {
	l.len++
}
func (l *List) decrLen() {
	l.len--
}

func (l *List) HasOneNode() bool {
	return l.len == 1
}

func (l *List) IsEmpty() bool {
	return l.len == 0
}
