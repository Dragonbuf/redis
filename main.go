package main

import (
	"fmt"
	_ "net/http/pprof"
	"os"
	"redis/server"
	"strconv"
)

var db *server.RedisDb
var Server *server.RedisServer
var Command string
var Key string
var Filed string
var Value string
var CurrentDb int

func main() {
	fmt.Println("[only support get set hset hget del hdel expire exit]\n ")

	Server = server.NewRedisServer()
	db = Server.Select(0)

	for {
		fmt.Printf("go-redis " + strconv.Itoa(CurrentDb) + "->  ")
		Command = ""
		Key = ""
		Filed = ""
		Value = ""
		_, _ = fmt.Scanln(&Command, &Key, &Filed, &Value)
		DoCommand(Command, Key, Filed, Value)

	}

}

func DoCommand(command, key, filed, value string) {

	if command == "exit" {
		fmt.Println("good bye")
		os.Exit(1)
	}

	if len(command) == 0 || len(key) == 0 {
		fmt.Println("command or key can not empty")
		return
	}

	if command == "select" {
		dbNum, _ := strconv.Atoi(key)
		if dbNum > Server.GetDbNum() {
			fmt.Println("dbNum error")
			return
		}
		db = Server.Select(dbNum)
		CurrentDb = dbNum
		return
	}

	if db.ExpireIfNeeded(key) {
		db.Del(key)
		db.DelExpire(key)
	}

	// todo 改写这里 使用 clientCmd 管理参数及命令

	// todo 多个 list 时，需要 incr 多次
	Server.IncrDirty()

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
	if len(key) == 0 || len(filed) == 0 {
		fmt.Println("filed or key can not empty")
		return
	}
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
