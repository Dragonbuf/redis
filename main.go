package main

import (
	"fmt"
	"redis/adt"
)

func main() {
	zskip := adt.ZSkipList{}
	zskip.Level = 1
	fmt.Print(zskip.Level)
}
