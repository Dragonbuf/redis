package main

import (
	"fmt"
	_ "net/http/pprof"
	"redis/server"
	"redis/tool"
	"strconv"
)

var CurrentDb int

func main() {
	fmt.Println("[暂时只支持 get set hset hget del hdel expire exit 命令]\n ")
	for {
		fmt.Printf("go-redis 【" + strconv.Itoa(CurrentDb) + "】->  ")
		args := &tool.GoRedisArgs{}
		if _, err := fmt.Scanln(&args.Command, &args.Key, &args.Fields, &args.Value); err != nil {
			// 这里有问题，但是暂时忽略
		}
		DoCommand(args)
	}

}

func DoCommand(args *tool.GoRedisArgs) {

	if server.Db.ExpireIfNeeded(args.Command) {
		server.Db.Del(args.Command)
		server.Db.DelExpire(args.Command)
	}

	clientCmd := tool.GetSupportCommand(args)

	if cmd, ok := clientCmd[args.Command]; ok {

		if cmd.ArgsNumber != args.Size() {
			fmt.Println("命令需要 ", cmd.ArgsNumber, " 个参数，但是仅提供了 ", args.Size())
			return
		}

		// todo 多个 list 时，需要 incr 多次
		server.Server.IncrDirty()
		cmd.Function(cmd.Args)
	} else {
		fmt.Println("未找到此命令 :" + args.Command)
	}
}
