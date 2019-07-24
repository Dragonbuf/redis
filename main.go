package main

import (
	"fmt"
	"os"
	"redis/server"
)

var (
	Command string
	key     string
	Value   string
)

func main() {
	fmt.Println("[only support get set exit] \nex :\n set fuckyou\n or \n get (will print fuck you) \n")

	db := server.NewRedisDb()

	for {
		fmt.Printf("go-redis ->  ")
		_, _ = fmt.Scanln(&Command, &key, &Value)
		if Command == "set" {
			db.SetString(&key, &Value)
		} else if Command == "get" {
			fmt.Println(db.GetString(&key))
		} else if Command == "exit" {
			fmt.Println("good bye")
			os.Exit(1)
		} else {
			fmt.Println("not found or support ths command :" + Command)
		}
	}

}
