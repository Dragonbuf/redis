package adt

import (
	"fmt"
	"testing"
)

func TestDict(t *testing.T) {
	dict := NewDict()
	k := "a"
	v := "b"

	sds := NewSdsHdr()
	sds.Set(&v)

	strObj := StringObject{}
	strObj.Sds = sds

	obj := Object{}
	obj.strObj = &strObj
	value := DictValue{obj: &obj}
	value.valueType = DictvalueTypeStringObj

	k2 := "c"
	v2 := "d"

	sds2 := NewSdsHdr()
	sds2.Set(&v2)

	strObj2 := StringObject{}
	strObj2.Sds = sds

	obj2 := Object{}
	obj2.strObj = &strObj
	value2 := DictValue{obj: &obj2}
	value.valueType = DictvalueTypeStringObj

	dict.Hset(&k, &value)
	dict.Hset(&k2, &value2)

	res := dict.Hget(&k)
	if res != &value {
		t.Error("hget is wrong ", res)
	}

	fmt.Println(string(*res.obj.strObj.Sds.buf))

}

func TestHash(t *testing.T) {
	dict := NewDict()
	k := "a"

	hash := dict.GetHash(&k)
	if hash == uint64(0) {
		t.Error("hash is error ", hash)
	}

}
