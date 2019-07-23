package adt

import (
	"testing"
)

func TestDict(t *testing.T) {
	dict := NewDict()
	k := "a"
	v := "b"

	k2 := "c"
	v2 := "d"
	dict.Hset((*DictKey)(&k), (*DictValue)(&v))
	dict.Hset((*DictKey)(&k2), (*DictValue)(&v2))

	res := dict.Hget((*DictKey)(&k))
	if res != (*DictValue)(&v) {
		t.Error("hget is wrong ", res)
	}
}

func TestHash(t *testing.T) {
	dict := NewDict()
	k := "a"

	hash := dict.GetHash((*DictKey)(&k))
	if hash == uint64(0) {
		t.Error("hash is error ", hash)
	}

}
