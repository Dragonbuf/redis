package adt

import (
	"fmt"
	"testing"
)

func TestList(t *testing.T) {
	list := List{}

	list.Lpush("a")
	list.Lpush("b")
	// fmt.Println(list.Head.Value.Get()) // 2

	fmt.Println(list.Lpop()) // b
	fmt.Println(list.Lpop()) // a

	list.Lpush("world")
	list.Rpush("hello")

	fmt.Println(list.Rpop()) // hello
	fmt.Println(list.Rpop()) // world
	fmt.Println(list.Rpop()) // 0

	t.Error("hget is wrong ", 111)
}
