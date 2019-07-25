package main

import (
	"fmt"
	"os"
	"redis/server"
)

func main() {
	fmt.Println("[only support get set exit]  \n")

	db := server.NewRedisDb()

	for {
		fmt.Printf("go-redis ->  ")
		Command := ""
		key := ""
		Value := ""
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
