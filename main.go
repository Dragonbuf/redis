package main

import (
	"fmt"
	"reflect"
)

func main() {
	var b int16
	b = 1
	fmt.Println(reflect.TypeOf(b).String())

}
