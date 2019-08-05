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
	fmt.Println("[only support get set hset hget del hdel expire exit]\n ")

	db = server.NewRedisDb()

	for {
		fmt.Printf("go-redis ->  ")
		Command = ""
		Key = ""
		Filed = ""
		Value = ""
		_, _ = fmt.Scanln(&Command, &Key, &Filed, &Value)
		DoCommand(Command, Key, Filed, Value)

	}

}

func DoCommand(command, key, filed, value string) {

	if command == "quit" {
		fmt.Println("good bye")
		os.Exit(1)
	}

	if len(command) == 0 || len(key) == 0 {
		fmt.Println("command or key can not empty")
		return
	}

	// todo del this key, 现在仅仅判断过期返回过期
	if db.IsExpired(key) {
		fmt.Println(key + " expired 了")
		return
	}

	switch command {
	case "set":
		set(key, filed)
	case "get":
		get(key)
	case "hset":
		Hset(key, filed, value)
	case "hget":
		Hget(key, filed)
	case "rpush":
		Rpush(Key, Filed, Value)
	case "rpop":
		Rpop(Key)

	case "del":
		del(key)
	case "hdel":
		hdel(key, filed)
	case "expire":
		expire(key, filed)
	default:
		fmt.Println("not found or support ths command :" + command)
	}
}

func expire(key, filed string) {
	db.Expire(key, filed)
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

func Rpush(key, filed, value string) {

	if len(key) == 0 || len(filed) == 0 {
		fmt.Println("value or key can not empty")
		return
	}

	if len(value) == 0 {
		db.RPush(key, filed)
	} else {
		db.RPush(key, filed, value)
	}
}

func Rpop(key string) {
	if len(key) == 0 {
		fmt.Println("key can not empty")
		return
	}

	fmt.Println(db.RPop(key))
}
