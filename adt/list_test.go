package adt

import (
	"testing"
)

func TestList(t *testing.T) {
	list := NewList()

	a := "a"
	b := "b"
	objA := NewRedisObject().Set(&a)
	objB := NewRedisObject().Set(&b)

	list.LPush(objA)
	list.LPush(objB)

	if *list.LPop().Sdshdr.Get() != b {
		t.Error("list lpop not equals 2")
	}
	if *list.LPop().Sdshdr.Get() != a {
		t.Error("list lpop not equals 2")
	}

	world := "world"
	hello := "hello"

	obj3 := NewRedisObject().Set(&world)
	obj4 := NewRedisObject().Set(&hello)

	list.LPush(obj3)
	list.RPush(obj4)

	if *list.RPop().Sdshdr.Get() != hello {
		t.Error("rpop error")
	}

	if *list.RPop().Sdshdr.Get() != world {
		t.Error("rpop error")
	}

	if list.RPop() != nil {
		t.Error("rpop error")
	}

}
