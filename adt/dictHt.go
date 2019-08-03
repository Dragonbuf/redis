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

func (d *DictHt) InitHtBySize(size uint64) *DictHt {
	d.size = size
	d.used = 0
	d.sizeMask = size - 1

	d.table = make([]*DictEntry, int64(size))
	return d
}

func (d *DictHt) AddDictValue(key *RedisObject, value *RedisObject) {

	// 根据　hash 算法获取 index
	index := d.GetIndex(d.GetHash(key))

	// 查看是否　link 上有重复的　key
	if existsEntry := d.FindSameKey(index, key); existsEntry != nil {
		existsEntry.setValue(value)
		return
	}

	entry := NewDictEntry()
	entry.setKey(key).setValue(value)

	// hash 冲突了
	if d.IsHashConflict(index) {
		entry.next = d.table[index]
		d.table[index] = entry
	} else {
		d.table[index] = entry
	}

	d.IncrUsed()
}

func (d *DictHt) GetHash(key *RedisObject) (hashVal uint64) {

	k := key.Sdshdr.Get()

	for _, v := range *k {
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

func (d *DictHt) HasSameKey(index uint64, key *RedisObject) bool {
	return *d.table[index].key.Sdshdr.Get() == *key.Sdshdr.Get()
}

func (d *DictHt) FindSameKey(index uint64, key *RedisObject) *DictEntry {
	tempFind := d.table[index]

	for tempFind != nil {
		if *tempFind.key.Sdshdr.Get() == *key.Sdshdr.Get() {
			return tempFind
		}
		tempFind = tempFind.next
	}

	return nil
}

func (d *DictHt) ShouldReHash() bool {
	return d.used >= d.size
}

func (d *DictHt) findValue(key *RedisObject) *RedisObject {

	index := d.GetIndex(d.GetHash(key))

	if d.table == nil || d.table[index] == nil {
		return nil
	}

	if !d.IsLinked(index) {
		if *d.table[index].key.Sdshdr.Get() == *key.Sdshdr.Get() {
			return d.table[index].v
		}
		return nil
	}

	// next  有数据
	tempTable := d.table[index]

	for tempTable != nil {
		if *tempTable.key.Sdshdr.Get() == *key.Sdshdr.Get() {
			return tempTable.v
		}
		tempTable = tempTable.next
	}

	return nil
}

func (d *DictHt) delValue(key *RedisObject) int {

	index := d.GetIndex(d.GetHash(key))
	if d.table == nil || d.table[index] == nil {
		return 0
	}

	// next  有数据
	tempTable := d.table[index]
	for tempTable != nil {
		if *tempTable.key.Sdshdr.Get() == *key.Sdshdr.Get() {
			//tempTable.v
			if tempTable.next != nil {
				d.table[index] = tempTable.next
			}

			tempTable.v = nil
			d.used--
		}
		tempTable = tempTable.next
	}

	return 1
}

func (d *DictHt) IsLinked(index uint64) bool {
	return d.table[index].next != nil
}

func (d *DictHt) MoveTableToNewByIndex(i int64, ht *DictHt) int {

	j := 0
	// rehash 当前 treHashIndex 的数据
	for d.table[i] != nil {
		ht.AddDictValue(d.table[i].key, d.table[i].v)
		//ht.IncrUsed()
		j++
		d.table[i] = d.table[i].next
	}
	return j
}

func (d *DictHt) FinishedReHash(i int64) bool {
	return d.size <= uint64(i)
}
func (d *DictHt) IsEmpty() bool {
	return d == nil || d.size <= 0
}

func (d *DictHt) CompareKey(a *RedisObject, b *RedisObject) bool {
	return *a.Sdshdr.Get() == *b.Sdshdr.Get()
}
