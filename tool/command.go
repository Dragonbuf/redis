// go-redis 所有支持的命令都在这里
package tool

import (
	"fmt"
	"os"
	"redis/server"
	"strconv"
)

type cmd struct {
	Function   func([]string)
	Args       []string
	ArgsNumber int
}

func GetSupportCommand(args []string) map[string]cmd {
	return map[string]cmd{
		"set": {
			func(args []string) {
				if len(args) != 2 {
					return
				}

				key := args[0]
				filed := args[1]

				if len(key) == 0 || len(filed) == 0 {
					return
				}
				server.Db.Set(key, filed)
			},
			args,
			2,
		},
		"get": {
			func(args []string) {
				if cap(args) != 1 {
					return
				}
				fmt.Println(server.Db.Get(args[0]))
			},
			args,
			1,
		},
		"expire": {
			func(args []string) {
				if len(args) != 2 {
					fmt.Println("filed or key can not empty")
					return
				}
				server.Db.Expire(args[0], args[1])
			},
			args,
			2,
		},
		"hdel": {
			func(args []string) {
				fmt.Println(server.Db.Hdel(args[0], args[1]))
			},
			args,
			2,
		},
		"del": {
			func(args []string) {
				fmt.Println(server.Db.Del(args[0]))
			},
			args,
			1,
		},
		"hset": {
			func(args []string) {
				if len(args) != 3 {
					fmt.Println("filed or value can not empty")
					return
				}
				err := server.Db.HSet(args[0], args[1], args[2])
				if err != nil {
					fmt.Println(err)
				}
			},
			args,
			3,
		},
		"hget": {
			func(args []string) {
				if len(args) != 2 {
					fmt.Println("filed or key can not empty")
					return
				}
				fmt.Println(server.Db.HGet(args[0], args[1]))
			},
			args,
			2,
		},
		"rpush": {
			func(args []string) {
				if len(args) != 3 {
					fmt.Println("value or key can not empty")
					return
				}

				if len(args[2]) == 0 {
					server.Db.RPush(args[0], args[1])
				} else {
					server.Db.RPush(args[0], args[1], args[2])
				}
			},
			args,
			3,
		},
		"rpop": {
			func(args []string) {
				if len(args) != 1 {
					fmt.Println("key can not empty")
					return
				}

				fmt.Println(server.Db.RPop(args[0]))
			},
			args,
			1,
		},
		"select": {
			func(args []string) {
				// todo 实现全局函数
				dbNum, _ := strconv.Atoi(args[0])
				if dbNum > server.Server.GetDbNum() {
					fmt.Println("dbNum error")
					return
				}
				server.Server.Select(dbNum)
				return
			},
			args,
			1,
		},
		"exit": {
			func(args []string) {
				fmt.Println("good bye")
				os.Exit(1)
			},
			args,
			1,
		},
	}
}
