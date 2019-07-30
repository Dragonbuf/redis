package adt

import (
	"testing"
)

func TestDict(t *testing.T) {
	dict := NewDict()
	k := "fuck"
	v := "you"
	k1 := "fuck"
	v1 := "you"
	k2 := "c"
	v2 := "d"
	k3 := "e"
	v3 := "f"
	k4 := "g"

	ks := NewStringObject().SetString(&k)
	vs := NewDictValue().SetStringObject(NewStringObject().SetString(&v))

	ks1 := NewStringObject().SetString(&k1)
	vs1 := NewDictValue().SetStringObject(NewStringObject().SetString(&v1))

	ks2 := NewStringObject().SetString(&k2)
	vs2 := NewDictValue().SetStringObject(NewStringObject().SetString(&v2))

	ks3 := NewStringObject().SetString(&k3)
	vs3 := NewDictValue().SetStringObject(NewStringObject().SetString(&v3))

	ks4 := NewStringObject().SetString(&k4)

	dict.HsetString(ks, vs)
	dict.HsetString(ks1, vs1)

	// 完成了一次 rehash ，所以现在 size 是 4
	if used := dict.ht[0].used; used != 1 {
		t.Error("dict ht[0] use error, Used:", used)
	}
	if size := dict.ht[0].size; size != 2 {
		t.Error("dict ht[0] size error, Size:", size)
	}

	res := dict.HgetString(ks2)
	if res == v {
		t.Error("hget k2 is wrong ", res)
	}

	dict.HsetString(ks2, vs2)
	dict.HsetString(ks3, vs3)

	res = dict.HgetString(ks)
	if res != v {
		t.Error("hget k is wrong ", res)
	}

	res = dict.HgetString(ks2)
	if res != v2 {
		t.Error("hget k2 is wrong ", res)
	}

	res = dict.HgetString(ks4)
	if res != "<nil>" {
		t.Error("hget k4 is wrong ", res)
	}

}
