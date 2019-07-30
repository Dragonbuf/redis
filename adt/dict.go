package adt

const DictvalueTypeStringObj = "stringObj"
const DictvalueTypeHashObj = "hashObj"

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

func (d *Dict) HsetString(key *StringObject, value *DictValue) {
	d.Hset(key, value)
}

func (d *Dict) HgetString(key *StringObject) string {
	if strObj := d.Hget(key); strObj == nil {
		return "<nil>"
	} else {
		return *strObj.Get()
	}
}

func (d *Dict) HgetGetHash(key, filed *StringObject) string {
	return *d.Hget(key).Get()
}

// key 暂时只支持 string 吧
func (d *Dict) Hset(key *StringObject, value *DictValue) *Dict {

	// 正在　rehash 插入只插入　ht[1],其他情况只插入　ht[0]
	dictHt := d.GetWriteHt()
	dictHt.AddDictValue(key, value)

	// 插入完如果需要 rehash 则开始 rehash
	if dictHt.ShouldReHash() {
		d.BeginReHash()
	}

	return d
}

func (d *Dict) Hget(key *StringObject) *DictValue {

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

func (d *Dict) FinishedAllReHash() {
	d.ResetTreHashIdx()
	d.SwapHt()
	d.DestroyHt1()
}

func (d *Dict) BeginReHash() {
	if !d.IsReHashing() {
		d.treHashIdx++
	}

	writeHt := d.GetWriteHt()
	readHt := d.GetReadHt()

	// 如果　ht[1] 没有申请空间
	if writeHt.IsEmpty() {
		writeHt.InitHtBySize(readHt.size * 2)
	}

	i := readHt.MoveTableToNewByIndex(d.treHashIdx, writeHt)
	d.FinishedCurrentIndexReHash(i)
	if readHt.FinishedReHash(d.treHashIdx) {
		d.FinishedAllReHash()
	}
}

func (d *Dict) IsReHashing() bool {
	return d.treHashIdx != -1
}
func (d *Dict) ResetTreHashIdx() {
	d.treHashIdx = -1
}
func (d *Dict) FinishedCurrentIndexReHash(i int) {
	d.treHashIdx += int64(i)
}

func (d *Dict) SwapHt() {
	d.ht[0] = d.ht[1]
}

func (d *Dict) DestroyHt1() {
	d.ht[1] = *NewDictHt()
	d.ht[1].used = 0
	d.ht[1].size = 0
	d.ht[1].sizeMask = 0
}

func (d *Dict) GetWriteHt() *DictHt {
	if d.IsReHashing() {
		return &d.ht[1]
	}

	return &d.ht[0]
}

func (d *Dict) GetReadHt() *DictHt {
	return &d.ht[0]
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
