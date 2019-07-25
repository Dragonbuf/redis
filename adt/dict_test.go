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
	k3 := "e"
	v3 := "f"
	k4 := "g"

	dict.HsetString(&k, &v)
	dict.HsetString(&k2, &v2)
	dict.HsetString(&k3, &v3)

	res := dict.HgetString(&k)
	if res != v {
		t.Error("hget k is wrong ", res)
	}

	res = dict.HgetString(&k2)
	if res != v2 {
		t.Error("hget k2 is wrong ", res)
	}

	res = dict.HgetString(&k4)
	if res != "<nil>" {
		t.Error("hget k4 is wrong ", res)
	}

}
