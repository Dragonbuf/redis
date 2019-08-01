package main

import (
	"fmt"
	"os"
	"redis/adt"
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

			db.Set(adt.NewStringObject().SetString(&key), adt.NewDictValue().SetStringObject(adt.NewStringObject().SetString(&filed)))
		} else if Command == "get" {
			fmt.Println(db.Get(adt.NewStringObject().SetString(&key)))
		} else if Command == "hset" {
			if len(filed) == 0 || len(Value) == 0 {
				fmt.Println("filed or value can not empty")
				continue
			}
			err := db.HSet(adt.NewStringObject().SetString(&key), adt.NewStringObject().SetString(&filed), adt.NewStringObject().SetString(&Value))
			if err != nil {
				fmt.Println(err)
			}
		} else if Command == "hget" {
			if len(key) == 0 || len(filed) == 0 {
				fmt.Println("filed or key can not empty")
				continue
			}
			fmt.Println(db.HGet(adt.NewStringObject().SetString(&key), adt.NewStringObject().SetString(&filed)))
		} else if Command == "exit" {
			fmt.Println("good bye")
			os.Exit(1)
		} else {
			fmt.Println("not found or support ths command :" + Command)
		}

	}

}
