package main

import (
	"fmt"
	"os"
	"redis/server"
)

func main() {
	fmt.Println("[only support get set exit]\n ")

	db := server.NewRedisDb()

	for {
		fmt.Printf("go-redis ->  ")
		Command := ""
		key := ""
		filed := ""
		Value := ""
		_, _ = fmt.Scanln(&Command, &key, &filed, &Value)
		if Command == "set" {
			db.SetString(&key, &filed)
		} else if Command == "get" {
			fmt.Println(db.GetString(&key))
		} else if Command == "hset" {
			if len(filed) == 0 || len(Value) == 0 {
				fmt.Println("filed or value can not empty")
				continue
			}
			db.HSetString(&key, &filed, &Value)
		} else if Command == "hget" {
			if len(key) == 0 || len(filed) == 0 {
				fmt.Println("filed or key can not empty")
				continue
			}
			fmt.Println(db.HGetString(&key, &filed))
		} else if Command == "exit" {
			fmt.Println("good bye")
			os.Exit(1)
		} else {
			fmt.Println("not found or support ths command :" + Command)
		}

	}

}
