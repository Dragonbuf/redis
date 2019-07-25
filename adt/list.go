package adt

// ListNode ...
type ListNode struct {
	Prev  *ListNode
	Next  *ListNode
	Value *Sdshdr
}

// List ...
type List struct {
	Head *ListNode
	Tail *ListNode
	Len  uint32
}

// Lpush ...
func (list *List) Lpush(value string) int {

	// 将字符串存入sds中
	sdshdr := Sdshdr{}
	sdshdr.Set(&value)

	list.Len++
	if list.Len == 1 {
		curr := &ListNode{nil, nil, &sdshdr}
		list.Head = curr
		list.Tail = curr
	} else {
		list.Head = &ListNode{nil, list.Head, &sdshdr}
	}

	return 1
}

// Lpop ...
func (list *List) Lpop() string {
	pop := list.Head
	if nil == pop {
		return "0"
	}

	list.Head = pop.Next
	list.Len--
	if list.Len == 0 {
		list.Tail = nil
	}

	return pop.Value.Get()
}

// Rpush ...
func (list *List) Rpush(value string) int {
	// 将字符串存入sds中
	sdshdr := Sdshdr{}
	sdshdr.Set(&value)

	list.Len++
	if list.Len == 1 {
		curr := &ListNode{nil, list.Tail, &sdshdr}
		list.Head = curr
		list.Tail = curr
	} else {
		list.Tail = &ListNode{list.Tail, nil, &sdshdr}
	}

	return 1
}

// Rpop ...
func (list *List) Rpop() string {
	pop := list.Tail
	if pop == nil {
		return "0"
	}

	list.Tail = pop.Prev
	list.Len--

	return pop.Value.Get()
}
