package adt

import "math/rand"

type dictKey string

// 类型特定函数
type dicType string

// 私有数据　暂时不实现
type privateData struct {
}

type Dict struct {
	types       dicType
	privateData privateData
	ht          [2]DictHt
	treHashIdx  int64
}

type DictHt struct {
	table    []*DictEntry
	size     uint64
	sizeMask uint64
	used     uint64
}

//todo 使　v 可能是一个指针　可能是一个　uint64  可能是　int64
type DictEntry struct {
	key  *dictKey
	v    interface{} //
	next *DictEntry
}

type dictVUnion interface {
	val() *Val
	uint64() uint64
	int64() int64
}

type Val struct {
}

func (d *Dict) Hset(key dictKey, value interface{}) {
	hash := d.GetHash(key)
	index := d.GetIndex(hash)

	entry := new(DictEntry)
	entry.key = &key
	entry.v = value

	// 拉链法解决　hash 冲突
	if d.ht[0].table[index] != nil {
		entry.next = d.ht[0].table[index]
	}

	d.ht[0].table[index] = entry.next
}

func (d *Dict) Hget(key dictKey) interface{} {
	hash := d.GetHash(key)
	index := d.GetIndex(hash)

	// 拉链法解决　hash 冲突
	if d.ht[0].table[index] == nil {
		return nil
	}

	for d.ht[0].table[index].next != nil {
		if d.ht[0].table[index].key == &key {
			return d.ht[0].table[index].v
		}
	}

	return nil
}

// todo 实现 hash 算法, 这里使用了随机数
func (d *Dict) GetHash(key dictKey) uint64 {
	return uint64(rand.Int())
}

func (d *Dict) GetIndex(hash uint64) uint64 {
	return hash & d.ht[0].sizeMask
}

// 复制键函数
func (d *Dict) CopyKey(hash uint64) uint64 {
	return hash & d.ht[0].sizeMask
}

//　复制值函数
func (d *Dict) CopyValue(hash uint64) uint64 {
	return hash & d.ht[0].sizeMask
}

// 对比键函数
func (d *Dict) Compare(hash uint64) uint64 {
	return hash & d.ht[0].sizeMask
}

// 销毁键函数
func (d *Dict) DestroyKey(hash uint64) uint64 {
	return hash & d.ht[0].sizeMask
}

// 销毁键函数
func (d *Dict) DestroyValue(hash uint64) uint64 {
	return hash & d.ht[0].sizeMask
}
