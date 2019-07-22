package main

import (
	"fmt"
	"os"
	"redis/adt"
)

var (
	Command string
	key     string
	Value   string
)

func main() {
	fmt.Println("[only support get set exit] \nex :\n set fuckyou\n or \n get (will print fuck you) \n")
	sds := adt.NewSdsHdr()
	for {
		fmt.Printf("go-redis ->  ")
		_, _ = fmt.Scanln(&Command, &Value)
		if Command == "set" {
			sds.Set(Value)
		} else if Command == "get" {
			fmt.Println(sds.Get())
		} else if Command == "exit" {
			fmt.Println("good bye")
			os.Exit(1)
		} else {
			fmt.Println("not found or support ths command :" + Command)
		}
	}

}
