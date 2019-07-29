package adt

type DictEntry struct {
	next *DictEntry
	key  *StringObject
	v    *DictValue
}

func NewDictEntry() *DictEntry {
	return &DictEntry{}
}

func (d *DictEntry) setKey(key *StringObject) *DictEntry {
	d.key = key
	return d
}

func (d *DictEntry) setValue(value *StringObject) *DictEntry {
	if d.v == nil {
		val := NewDictValue()
		val.StringObject = value
		d.v = val
	}
	d.v.StringObject = value
	return d
}
