package adt

type ZSkipList struct {
	Header *ZSkipListNode
	Tail   *ZSkipListNode
	Level  uint32
	Length uint32
}

type ZSkipListNode struct {
	Level    []ZSkipListLevel
	BackWard *ZSkipListNode
	source   float64
	obj      *Obj
}

type ZSkipListLevel struct {
	Forward *ZSkipListNode
	span    uint32
}

type Obj struct {
}
