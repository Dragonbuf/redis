package adt

import (
	"fmt"
)

const DictvalueTypeStringObj = "stringObj"

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

type DictEntry struct {
	key  *string
	v    *DictValue
	next *DictEntry
}

type DictValue struct {
	obj       *Object
	int64Obj  int64
	uint64Obj uint64
	valueType string
}

type Object struct {
	strObj  *StringObject
	hashOjb *HashObject
	listObj *ListObject
}

type StringObject struct {
	Sds *Sdshdr
}

type HashObject struct {
}

type ListObject struct {
}

func NewDict() *Dict {
	dict := &Dict{}
	dict.treHashIdx = -1 // 表示没有 rehash

	table := make([]*DictEntry, 2)
	table[0] = &DictEntry{key: nil, v: nil}
	table[1] = &DictEntry{key: nil, v: nil}

	// 初始化 hashTable
	dict.ht[0] = DictHt{table: table, size: 2, sizeMask: 1, used: 0}
	dict.ht[1] = DictHt{table: []*DictEntry{}, size: 0, sizeMask: 0, used: 0}
	return dict
}
func (d *Dict) HsetString(key, value *string) {
	sds := NewSdsHdr()
	sds.Set(value)

	strObj := StringObject{}
	strObj.Sds = sds

	obj := Object{}
	obj.strObj = &strObj
	dictValue := DictValue{obj: &obj}
	dictValue.valueType = DictvalueTypeStringObj

	d.Hset(key, &dictValue)
}

func (d *Dict) HgetString(key *string) string {
	dictValue := d.Hget(key)

	if dictValue == nil {
		return "<nil>"
	}

	if dictValue.valueType != DictvalueTypeStringObj {
		panic("can not user HgetString in dictValue where type is not string")
	}

	return dictValue.obj.strObj.Sds.Get()
}

// 暂时只支持 string 吧
func (d *Dict) Hset(key *string, value *DictValue) {

	// 如果正在 rehash 则所有的数据都新增到 ht[1] 中
	htIndex := 0
	if d.treHashIdx != -1 {
		htIndex = 1
	}

	d.ht[htIndex].used++

	// 开始 rehash
	if d.ht[0].used >= d.ht[0].size {
		d.Rehash()
	}

	hash := d.GetHash(key)
	index := d.GetIndex(hash)

	fmt.Println("debug: hset index is ", index)
	// 如果原先位置已经有数据，拉链法解决　hash 冲突
	if d.ht[htIndex].table[index].key != nil {
		fmt.Println("debug: 出现 hash 冲突，使用拉链法解决 ")
		entry := &DictEntry{key: key, v: value}
		entry.next = d.ht[htIndex].table[index]
		d.ht[htIndex].table[index] = entry
	} else {
		d.ht[htIndex].table[index].key = key
		d.ht[htIndex].table[index].v = value
	}

}

func (d *Dict) Hget(key *string) *DictValue {
	hash := d.GetHash(key)
	index := d.GetIndex(hash)

	fmt.Println("debug: hget [", *key, "] in index ", index)

	if d.ht[0].table[index].key == nil {
		fmt.Println("debug: ht[0] table is nil ")
		return nil
	}

	// 如果 next 没有数据，证明只有单一的 entry
	if d.ht[0].table[index].next == nil {
		if d.ht[0].table[index].key == key {
			return d.ht[0].table[index].v
		}
		return nil
	}

	// next  有数据
	for d.ht[0].table[index] != nil {
		fmt.Println("debug: ht table index key is  ", *d.ht[0].table[index].key)

		if d.ht[0].table[index].key == key {
			fmt.Println("debug: hget found key in index ", index)
			return d.ht[0].table[index].v
		}
		d.ht[0].table[index] = d.ht[0].table[index].next
	}

	fmt.Println("debug: ht[0]  not found this key ")
	return nil
}

func (d *Dict) GetHash(key *string) (hashVal uint64) {

	for _, v := range *key {
		hashVal = (hashVal << 5) + uint64(v+1)
	}

	return hashVal
}

func (d *Dict) Rehash() {
	d.treHashIdx++

	// 刚开始 rehash 时，需要申请 ht[1] 的空间
	if d.treHashIdx == 0 {
		// todo 字节对齐申请的空间
		d.ht[1].size = 2 * d.ht[0].size
		d.ht[1].used = 0
		d.ht[1].sizeMask = d.ht[1].size - 1

		fmt.Println(d.ht[1].size)
		// 申请足够空间
		table := make([]*DictEntry, int64(d.ht[1].size))
		d.ht[1].table = table
	}

	// rehash 当前 treHashIndex 的数据
	for d.ht[0].table[d.treHashIdx] != nil {

		d.ht[1].used++
		key := d.ht[0].table[d.treHashIdx].key
		value := d.ht[0].table[d.treHashIdx].v
		hash := d.GetHash(key)
		index := d.GetIndex(hash)

		fmt.Println("debug: hset index is ", index)
		// 如果原先位置已经有数据，拉链法解决　hash 冲突
		if d.ht[1].table[index] != nil {
			fmt.Println("debug: 出现 hash 冲突，使用拉链法解决 ")
			entry := &DictEntry{key: key, v: value}
			entry.next = d.ht[1].table[index]
			d.ht[1].table[index] = entry
		} else {
			entry := &DictEntry{key: key, v: value}
			entry.key = key
			entry.v = value
			d.ht[1].table[index] = entry
		}

		d.ht[0].table[d.treHashIdx] = d.ht[0].table[d.treHashIdx].next
	}

	// rehash 完毕
	if uint64(d.treHashIdx) == d.ht[0].used {
		d.ht[0] = d.ht[1]
		d.ht[1] = DictHt{}
		d.treHashIdx = -1
	}
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
