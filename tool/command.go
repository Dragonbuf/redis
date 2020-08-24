// go-redis 所有支持的命令都在这里
package tool

import (
	"fmt"
	"os"
	"redis/server"
	"strconv"
)

type GoRedisArgs struct {
	Command string
	Key     string
	Fields  string
	Value   string
}

type cmd struct {
	Function   func(*GoRedisArgs)
	Args       *GoRedisArgs
	ArgsNumber int
}

func GetSupportCommand(args *GoRedisArgs) map[string]cmd {
	return map[string]cmd{
		"set": {
			func(args *GoRedisArgs) {
				server.Db.Set(args.Key, args.Fields)
			},
			args,
			3,
		},
		"get": {
			func(args *GoRedisArgs) {
				fmt.Println(server.Db.Get(args.Key))
			},
			args,
			2,
		},
		"expire": {
			func(args *GoRedisArgs) {
				server.Db.Expire(args.Key, args.Fields)
			},
			args,
			3,
		},
		"hdel": {
			func(args *GoRedisArgs) {
				fmt.Println(server.Db.Hdel(args.Key, args.Fields))
			},
			args,
			3,
		},
		"del": {
			func(args *GoRedisArgs) {
				fmt.Println(server.Db.Del(args.Key))
			},
			args,
			2,
		},
		"hset": {
			func(args *GoRedisArgs) {
				if err := server.Db.HSet(args.Key, args.Fields, args.Value); err != nil {
					fmt.Println(err)
				}
			},
			args,
			4,
		},
		"hget": {
			func(args *GoRedisArgs) {
				fmt.Println(server.Db.HGet(args.Key, args.Fields))
			},
			args,
			3,
		},
		"rpush": {
			func(args *GoRedisArgs) {
				if args.Size() == 0 {
					server.Db.RPush(args.Key, args.Fields)
				} else {
					server.Db.RPush(args.Key, args.Fields, args.Value)
				}
			},
			args,
			4,
		},
		"rpop": {
			func(args *GoRedisArgs) {
				fmt.Println(server.Db.RPop(args.Key))
			},
			args,
			2,
		},
		"select": {
			func(args *GoRedisArgs) {
				// todo 实现全局函数
				dbNum, _ := strconv.Atoi(args.Key)
				if dbNum > server.Server.GetDbNum() {
					fmt.Println("dbNum 不能超过 ", server.Server.GetDbNum())
					return
				}
				server.Server.Select(dbNum)
				return
			},
			args,
			2,
		},
		"exit": {
			func(args *GoRedisArgs) {
				fmt.Println("good bye")
				os.Exit(1)
			},
			args,
			2,
		},
	}
}

//所有非空的参数
func (g *GoRedisArgs) Size() int {
	if g.Command == "" {
		return 0
	}

	if g.Key == "" {
		return 1
	}

	if g.Fields == "" {
		return 2
	}

	if g.Value == "" {
		return 3
	}

	return 0
}
