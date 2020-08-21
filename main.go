package main

import (
	"fmt"
	_ "net/http/pprof"
	"redis/server"
	"redis/tool"
	"strconv"
)

var Command string
var Key string
var Filed string
var Value string
var CurrentDb int

func main() {
	fmt.Println("[暂时只支持 get set hset hget del hdel expire exit 命令]\n ")
	for {
		fmt.Printf("go-redis 【" + strconv.Itoa(CurrentDb) + "】->  ")
		Command = ""
		Key = ""
		Filed = ""
		Value = ""
		_, _ = fmt.Scanln(&Command, &Key, &Filed, &Value)

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

	if server.Db.ExpireIfNeeded(args[0]) {
		server.Db.Del(args[0])
		server.Db.DelExpire(args[0])
	}

	clientCmd := tool.GetSupportCommand(args)

	if cmd, ok := clientCmd[command]; ok {

		if cmd.ArgsNumber != len(args) {
			fmt.Println("命令需要 ", cmd.ArgsNumber, " 个参数，但是仅提供了 ", len(args))
			return
		}

		// todo 多个 list 时，需要 incr 多次
		server.Server.IncrDirty()
		cmd.Function(cmd.Args)
	} else {
		fmt.Println("not found or support ths command :" + command)
	}
}
