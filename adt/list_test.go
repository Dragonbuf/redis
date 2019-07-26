package adt

import (
	"testing"
)

func TestList(t *testing.T) {
	list := NewList()

	a := "a"
	b := "b"
	list.LPush(&a)
	list.LPush(&b)

	if *list.LPop() != b {
		t.Error("list lpop not equals 2")
	}
	if *list.LPop() != a {
		t.Error("list lpop not equals 2")
	}

	world := "world"
	hello := "hello"
	list.LPush(&world)
	list.RPush(&hello)

	if *list.RPop() != hello {
		t.Error("rpop error")
	}

	if *list.RPop() != world {
		t.Error("rpop error")
	}

	if list.RPop() != nil {
		t.Error("rpop error")
	}

}
