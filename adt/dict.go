package adt

const DictvalueTypeStringObj = "stringObj"

// 类型特定函数
type dicType string

// 私有数据　暂时不实现
type privateData struct{}

type Dict struct {
	types       dicType
	privateData privateData
	ht          [2]DictHt
	treHashIdx  int64
}

func NewDict() *Dict {
	return &Dict{
		treHashIdx: -1,
		ht: [2]DictHt{
			{table: make([]*DictEntry, 2), size: 2, sizeMask: 1, used: 0},
			{table: []*DictEntry{}, size: 0, sizeMask: 0, used: 0},
		}}
}

func (d *Dict) HsetString(key, value *string) {

	dictValue := NewDictValue().SetStringValue(value)
	d.Hset(key, dictValue)
}

func (d *Dict) HgetString(key *string) string {
	return d.Hget(key).ToString()
}

func (d *Dict) IsReHashing() bool {
	return d.treHashIdx != -1
}

// key 暂时只支持 string 吧
func (d *Dict) Hset(key *string, value *DictValue) {

	dictHt := NewDictHt()

	// 正在　rehash 插入只插入　ht[1],其他情况只插入　ht[0]
	if d.IsReHashing() {
		dictHt = &d.ht[1]
	} else {
		dictHt = &d.ht[0]
	}

	dictHt.AddDictValue(key, value)

	if dictHt.ShouldReHash() {

		d.BeginReHash()
		var newDictHt *DictHt

		// 如果　ht[1] 没有申请空间
		if d.ht[1].size == 0 {
			newDictHt = NewDictHt()
			newDictHt.InitHtBySize(dictHt.size * 2) // todo 字节对齐申请的空间,这里简单申请 2 倍空间
			d.ht[1] = *newDictHt
		} else {
			newDictHt = &d.ht[1]
		}

		dictHt.MoveTableToNewByIndex(d.treHashIdx, newDictHt)
		d.FinishedCurrentIndexReHash()
		if dictHt.FinishedReHash(d.treHashIdx) {
			d.FinishedAllReHash()
		}
	}

}

func (d *Dict) Hget(key *string) *DictValue {

	dictHt := NewDictHt()
	dictHt = &d.ht[0]
	value := dictHt.findValue(key)

	// 如果正在　rehash 需要分别查询　ht0 ht1 是否存在
	if !d.IsReHashing() || value != nil {
		return value
	}

	dictHt = &d.ht[1]
	return dictHt.findValue(key)
}

func (d *Dict) FinishedCurrentIndexReHash() {
	d.treHashIdx++
}

func (d *Dict) FinishedAllReHash() {
	d.ResetTreHashIdx()
	d.SwapHt()
	d.DestroyHt1()
}

func (d *Dict) ResetTreHashIdx() {
	d.treHashIdx = -1
}

func (d *Dict) BeginReHash() {
	if d.treHashIdx == -1 {
		d.treHashIdx++
	}
}

func (d *Dict) SwapHt() {
	d.ht[0] = d.ht[1]
}

func (d *Dict) DestroyHt1() {
	d.ht[1] = *NewDictHt()
}

// 复制键函数
func (d *Dict) CopyKey() {}

//　复制值函数
func (d *Dict) CopyValue() {}

// 对比键函数
func (d *Dict) Compare() {}

// 销毁键函数
func (d *Dict) DestroyKey() {}

// 销毁键函数
func (d *Dict) DestroyValue() {}
