package adt

type DictEntry struct {
	next *DictEntry
	key  *StringObject
	v    *DictValue // redisObj int64 uint64
}

func NewDictEntry() *DictEntry {
	return &DictEntry{}
}

func (d *DictEntry) setKey(key *StringObject) *DictEntry {
	d.key = key
	return d
}

func (d *DictEntry) setValue(value *DictValue) *DictEntry {
	d.v = value
	return d
}
