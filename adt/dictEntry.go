package adt

type DictEntry struct {
	next *DictEntry
	key  *RedisObject
	v    *RedisObject // c 里面是 int64 uint64, *v 这里不方便实现，直接使用了 redisObject
}

func NewDictEntry() *DictEntry {
	return &DictEntry{}
}

func (d *DictEntry) setKey(key *RedisObject) *DictEntry {
	d.key = key
	return d
}

func (d *DictEntry) setValue(value *RedisObject) *DictEntry {
	d.v = value
	return d
}
