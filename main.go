package main

import (
	. "fmt"
	_ "net/http/pprof"
	"os"
	"redis/server"
	"redis/tool"
	"strconv"
	"time"
)

var db *server.RedisDb
var Server *server.RedisServer
var Command string
var Key string
var Filed string
var Value string
var CurrentDb int

func testGoroutineWithTimeOut() {

	done := make(chan struct{})
	errChan := make(chan error)

	pool := tool.NewGoroutinePool(10)
	for i := 0; i < 10; i++ {
		pool.Add(1)
		go func() {
			pool.Done()

			println("bg save ing")
			//if err!=nil{
			//	errChan<-errors.New("error")
			//}
		}()
	}

	// wg.Wait()此时也要go出去,防止在wg.Wait()出堵住
	go func() {
		pool.Wait()
		close(done)
	}()

	select {
	// 错误快返回 既可以使用 return nil, err
	case err := <-errChan:
		println(err)
		// 正常结束完成
	case <-done:
		// 超时
	case <-time.After(500 * time.Millisecond):
	}
}

func main() {
	Println("[only support get set hset hget del hdel expire exit]\n ")

	Server = server.NewRedisServer()
	db = Server.Select(0)

	for {
		Printf("go-redis " + strconv.Itoa(CurrentDb) + "->  ")
		Command = ""
		Key = ""
		Filed = ""
		Value = ""
		_, _ = Scanln(&Command, &Key, &Filed, &Value)

		args := []string{Key}

		if Filed != "" {
			args = append(args, Filed)
		}

		if Value != "" {
			args = append(args, Value)
		}

		DoCommand(Command, args)
	}

}

func DoCommand(command string, args []string) {

	if db.ExpireIfNeeded(args[0]) {
		db.Del(args[0])
		db.DelExpire(args[0])
	}

	clientCmd := map[string]struct {
		function   func([]string)
		args       []string
		argsNumber int
	}{
		"set": {
			func(args []string) {
				if len(args) != 2 {
					return
				}

				key := args[0]
				filed := args[1]

				if len(key) == 0 || len(filed) == 0 {
					Println("filed or key can not empty")
					return
				}
				db.Set(key, filed)
			},
			args,
			2,
		},
		"get": {
			func(args []string) {
				if cap(args) != 1 {
					Println("key can not empty")
					return
				}
				Println(db.Get(args[0]))
			},
			args,
			1,
		},
		"expire": {
			func(args []string) {
				if len(args) != 2 {
					Println("filed or key can not empty")
					return
				}
				db.Expire(args[0], args[1])
			},
			args,
			2,
		},
		"hdel": {
			func(args []string) {
				Println(db.Hdel(args[0], args[1]))
			},
			args,
			2,
		},
		"del": {
			func(args []string) {
				Println(db.Del(args[0]))
			},
			args,
			1,
		},
		"hset": {
			func(args []string) {
				if len(args) != 3 {
					Println("filed or value can not empty")
					return
				}
				err := db.HSet(args[0], args[1], args[2])
				if err != nil {
					Println(err)
				}
			},
			args,
			3,
		},
		"hget": {
			func(args []string) {
				if len(args) != 2 {
					Println("filed or key can not empty")
					return
				}
				Println(db.HGet(args[0], args[1]))
			},
			args,
			2,
		},
		"rpush": {
			func(args []string) {
				if len(args) != 3 {
					Println("value or key can not empty")
					return
				}

				if len(args[2]) == 0 {
					db.RPush(args[0], args[1])
				} else {
					db.RPush(args[0], args[1], args[2])
				}
			},
			args,
			3,
		},
		"rpop": {
			func(args []string) {
				if len(args) != 1 {
					Println("key can not empty")
					return
				}

				Println(db.RPop(args[0]))
			},
			args,
			1,
		},
		"select": {
			func(args []string) {
				dbNum, _ := strconv.Atoi(args[0])
				if dbNum > Server.GetDbNum() {
					Println("dbNum error")
					return
				}
				db = Server.Select(dbNum)
				CurrentDb = dbNum
				return
			},
			args,
			1,
		},
		"exit": {
			func(args []string) {
				Println("good bye")
				os.Exit(1)
			},
			args,
			1,
		},
	}

	if cmd, ok := clientCmd[command]; ok {
		// todo 多个 list 时，需要 incr 多次
		Server.IncrDirty()
		cmd.function(cmd.args)
	} else {
		Println("not found or support ths command :" + command)
	}
}
