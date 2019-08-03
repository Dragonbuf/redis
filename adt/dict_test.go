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

	ks := NewRedisObject().Set(&k)
	vs := NewRedisObject().Set(&v)

	ks1 := NewRedisObject().Set(&k1)
	vs1 := NewRedisObject().Set(&v1)

	ks2 := NewRedisObject().Set(&k2)
	vs2 := NewRedisObject().Set(&v2)

	ks3 := NewRedisObject().Set(&k3)
	vs3 := NewRedisObject().Set(&v3)

	ks4 := NewRedisObject().Set(&k4)

	dict.Hset(ks, vs)
	dict.Hset(ks1, vs1)

	// 完成了一次 rehash ，所以现在 size 是 4
	if used := dict.ht[0].used; used != 1 {
		t.Error("dict ht[0] use error, Used:", used)
	}
	if size := dict.ht[0].size; size != 2 {
		t.Error("dict ht[0] size error, Size:", size)
	}

	redisObject := dict.Hget(ks2)
	if redisObject != nil {
		t.Error("hget k2 is wrong ", redisObject)
	}

	dict.Hset(ks2, vs2)
	dict.Hset(ks3, vs3)

	redisObject = dict.Hget(ks)
	if *redisObject.Sdshdr.Get() != v {
		t.Error("hget k is wrong ", redisObject)
	}

	redisObject = dict.Hget(ks2)
	if *redisObject.Sdshdr.Get() != v2 {
		t.Error("hget k2 is wrong ", redisObject)
	}

	redisObject = dict.Hget(ks4)
	if redisObject != nil {
		t.Error("hget k4 is wrong ", redisObject)
	}

}

func BenchmarkNewDict(b *testing.B) {

	dict := NewDict()
	k := "fuck"
	ks := NewRedisObject().Set(&k)

	for i := 0; i < b.N; i++ {
		str := NewRedisObject().Set(i)
		dict.Hset(ks, str)
	}

}
