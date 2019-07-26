package adt

import (
	"testing"
)

func TestList(t *testing.T) {
	list := List{}

	list.Lpush("a")
	list.Lpush("b")

	if list.Lpop() != "b" {
		t.Error("list lpop not equals 2")
	}
	if list.Lpop() != "a" {
		t.Error("list lpop not equals 2")
	}

	list.Lpush("world")
	list.Rpush("hello")

	if list.Rpop() != "hello" {
		t.Error("rpop error")
	}

	if list.Rpop() != "world" {
		t.Error("rpop error")
	}

	if list.Rpop() != "0" {
		t.Error("rpop error")
	}

}
