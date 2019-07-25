package adt

type DictHt struct {
	table    []*DictEntry
	size     uint64
	sizeMask uint64
	used     uint64
}

func NewDictHt() *DictHt {
	return &DictHt{}
}

func (d *DictHt) InitHtBySize(size uint64) {
	d.size = size
	d.used = 0
	d.sizeMask = size - 1

	d.table = make([]*DictEntry, int64(size))
}

func (d *DictHt) AddDictValue(key *string, value *DictValue) {

	// 根据　hash 算法获取 index
	index := d.GetIndex(d.GetHash(key))

	entry := NewDictEntry()
	entry.setKeyValue(key, value)

	if d.IsHashConflict(index) {
		entry.next = d.table[index]
		d.table[index] = entry
	} else {
		d.table[index] = entry
	}

	d.IncrUsed()
}

func (d *DictHt) GetHash(key *string) (hashVal uint64) {

	for _, v := range *key {
		hashVal = (hashVal << 5) + uint64(v+1)
	}

	return hashVal
}

func (d *DictHt) GetIndex(hash uint64) uint64 {
	return hash & d.sizeMask
}

func (d *DictHt) IncrUsed() {
	d.used++
}

func (d *DictHt) IsHashConflict(index uint64) bool {
	return d.table[index] != nil
}

func (d *DictHt) ShouldReHash() bool {
	return d.used >= d.sizeMask
}

func (d *DictHt) findValue(key *string) *DictValue {

	index := d.GetIndex(d.GetHash(key))

	if d.table == nil || d.table[index] == nil {
		return nil
	}

	if !d.IsLinked(index) {
		if string(*d.table[index].key) == string(*key) {
			return d.table[index].v
		}
		return nil
	}

	// next  有数据
	tempTable := d.table[index]

	for tempTable != nil {
		if *tempTable.key == *key {
			return tempTable.v
		}
		tempTable = tempTable.next
	}

	return nil
}

func (d *DictHt) IsLinked(index uint64) bool {
	return d.table[index].next != nil
}

func (d *DictHt) MoveTableToNewByIndex(i int64, ht *DictHt) {

	// rehash 当前 treHashIndex 的数据
	for d.table[i] != nil {
		ht.AddDictValue(d.table[i].key, d.table[i].v)
		ht.IncrUsed()
		d.table[i] = d.table[i].next
	}
}

func (d *DictHt) FinishedReHash(i int64) bool {
	return d.size <= uint64(i)
}
func (d *DictHt) IsEmpty() bool {
	return d.size <= 0
}
