package adt

type DictEntry struct {
	key  *string
	v    *DictValue
	next *DictEntry
}

func NewDictEntry() *DictEntry {
	return &DictEntry{}
}

func (d *DictEntry) setKey(key *string) {
	d.key = key
}

func (d *DictEntry) setValue(value *DictValue) {
	d.v = value
}

func (d *DictEntry) setKeyValue(key *string, value *DictValue) {
	d.setKey(key)
	d.setValue(value)
}
