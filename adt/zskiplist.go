package adt

// Redis 只在：实现有序集合键、在集群节点中用作内部数据结构　中使用了跳跃表
type ZSkipList struct {
	Header *ZSkipListNode // 指向跳跃表的头节点
	Tail   *ZSkipListNode // 指向跳跃表的尾节点
	Level  int32          // 记录跳跃表内　层数最大节点的层数　（不算表头）
	Length uint64         //	跳跃表的长度: 跳跃表目前包含节点的数量 (不算表头	)
}

type ZSkipListNode struct {
	Level    []ZSkipListLevel // 根据 power law（越大的数，出现的概率越小	） 随机生成一个 1-32 直接的值作为 level 的大小
	BackWard *ZSkipListNode   // 后退指针
	source   float64          // 分值    多个节点可以保存相同的分值，但是 obj 必须唯一
	obj      *Obj             // 成员对象
}

type ZSkipListLevel struct {
	Forward *ZSkipListNode // 前进指针
	span    uint32         //跨度
}

type Obj struct {
}

func (z *ZSkipList) Create() {

}

func (z *ZSkipList) ZRange(start, end int, withScores bool) {

}

func (z *ZSkipList) ZCard(start, end int, withScores bool) {

}
