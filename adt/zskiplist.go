package adt

import (
	"math/rand"
	"unsafe"
)

const ZSkipListP = 0.25

var ZSkipListMaxLevel = 64

func ZslRandomLevel() int {
	level := 1
	for float64(rand.Int()&0XFFFF) < (ZSkipListP * 0XFFFF) {
		level += 1
	}

	if level < ZSkipListMaxLevel {
		return level
	}
	return ZSkipListMaxLevel
}

// Redis 只在：实现有序集合键、在集群节点中用作内部数据结构　中使用了跳跃表
type ZSkipList struct {
	Header, Tail *ZSkipListNode // 指向跳跃表的头节点 尾节点
	Length       uint64         //	跳跃表的长度: 跳跃表目前包含节点的数量 (不算表头	)
	Level        int32          // 记录跳跃表内　层数最大节点的层数　（不算表头）
}

type ZSkipListNode struct {
	Ele      unsafe.Pointer // 这里其实是　sds
	source   float64
	BackWard *ZSkipListNode   // 后退指针
	Level    []ZSkipListLevel // 根据 power law（越大的数，出现的概率越小	） 随机生成一个 1-32 直接的值作为 level 的大小
}

type ZSkipListLevel struct {
	Forward *ZSkipListNode // 前进指针
	span    uint32         //跨度
}

func NewZSkipListNode() *ZSkipListNode {
	return &ZSkipListNode{
		Ele:      nil,
		source:   0,
		BackWard: nil,
		Level:    nil,
	}
}

func NewZSkipList() *ZSkipList {

	header := NewZSkipListNode()
	level := ZslRandomLevel()
	zsl := &ZSkipList{
		Header: header,
		Tail:   nil,
		Length: 0,
		Level:  1,
	}

	for j := 0; j < level; j++ {
		zsl.Header.Level[j].Forward = nil
		zsl.Header.Level[j].span = 0
	}

	header.BackWard = nil

	return zsl
}

// source 31 level 3
// create node
func (zsl *ZSkipList) Create() {

}

func (zsl *ZSkipList) Insert() {
	//查找要插入的位置
	// 调整跳跃表的高度
	//插入节点
	//调账 backward
}

// 找到插入的位置
func (zsl *ZSkipList) Find() {
	for i := zsl.Level - 1; i > 0; i-- {
		var rank [64]int32
		if i == zsl.Level-1 {
			rank[i] = i
		} else {
			rank[i] = rank[i+1]
		}
	}
}

func (zsl *ZSkipList) ZRange(start, end int, withScores bool) {

}

func (zsl *ZSkipList) ZCard(start, end int, withScores bool) {

}
