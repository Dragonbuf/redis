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

	if string(*(*[]byte)(list.LPop().Ele)) != b {
		t.Error("list lpop not equals 2")
	}
	if string(*(*[]byte)(list.LPop().Ele)) != a {
		t.Error("list lpop not equals 2")
	}

	world := "world"
	hello := "hello"

	obj3 := NewRedisObject().Set(&world)
	obj4 := NewRedisObject().Set(&hello)

	list.LPush(obj3)
	list.RPush(obj4)

	if string(*(*[]byte)(list.LPop().Ele)) != hello {
		t.Error("rpop error")
	}

	if string(*(*[]byte)(list.LPop().Ele)) != world {
		t.Error("rpop error")
	}

	if list.RPop() != nil {
		t.Error("rpop error")
	}

}
