package main

import (
	"fmt"
	_ "net/http/pprof"
	"os"
	"redis/server"
)

var db *server.RedisDb
var Command string
var Key string
var Filed string
var Value string

func main() {
	fmt.Println("[only support get set exit]\n ")

	db = server.NewRedisDb()

	for {
		fmt.Printf("go-redis ->  ")
		Command = ""
		Key = ""
		Filed = ""
		Value = ""
		_, _ = fmt.Scanln(&Command, &Key, &Filed, &Value)

		switch Command {
		case "set":
			set(Key, Filed)
		case "get":
			get(Key)
		case "hset":
			Hset(Key, Filed, Value)
		case "hget":
			Hget(Key, Filed)
		case "del":
			del(Key)
		case "hdel":
			hdel(Key, Filed)
		case "exit":
			fmt.Println("good bye")
			os.Exit(1)
		default:
			fmt.Println("not found or support ths command :" + Command)
		}
	}

}

func hdel(key, filed string) {
	fmt.Println(db.Hdel(key, filed))
}

func del(key string) {
	fmt.Println(db.Del(key))
}

func set(key, filed string) {
	if len(key) == 0 || len(filed) == 0 {
		fmt.Println("filed or key can not empty")
		return
	}
	db.Set(key, filed)
}

func get(key string) {
	if len(key) == 0 {
		fmt.Println("key can not empty")
		return
	}
	fmt.Println(db.Get(key))
}

func Hset(key, filed, value string) {
	if len(filed) == 0 || len(value) == 0 || len(key) == 0 {
		fmt.Println("filed or value can not empty")
		return
	}
	err := db.HSet(key, filed, value)
	if err != nil {
		fmt.Println(err)
	}
}

func Hget(key, filed string) {
	if len(key) == 0 || len(filed) == 0 {
		fmt.Println("filed or key can not empty")
		return
	}
	fmt.Println(db.HGet(key, filed))
}
