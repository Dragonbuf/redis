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
	hash := d.GetHash(key)
	index := d.GetIndex(hash)
	fmt.Println("debug: hset index is ", index)
	// 如果原先位置已经有数据，拉链法解决　hash 冲突
	if d.ht[0].table[index].key != nil {
		fmt.Println("debug: 出现 hash 冲突，使用拉链法解决 ")
		entry := &DictEntry{key: key, v: value}
		entry.next = d.ht[0].table[index]
		d.ht[0].table[index] = entry
	} else {
		d.ht[0].table[index].key = key
		d.ht[0].table[index].v = value
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
		return d.ht[0].table[index].v
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

// todo 实现 hash 算法, 这里只返回 0
func (d *Dict) GetHash(key *string) (hashVal uint64) {

	for _, v := range *key {
		hashVal = (hashVal << 5) + uint64(v+1)
	}

	return hashVal
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
