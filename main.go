package main

import (
	. "fmt"
	_ "net/http/pprof"
	"redis/server"
	"redis/tool"
	"strconv"
	"time"
)

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

	if server.Db.ExpireIfNeeded(args[0]) {
		server.Db.Del(args[0])
		server.Db.DelExpire(args[0])
	}

	clientCmd := tool.GetSupportCommand(args)

	if cmd, ok := clientCmd[command]; ok {
		// todo 多个 list 时，需要 incr 多次
		server.Server.IncrDirty()
		cmd.Function(cmd.Args)
	} else {
		Println("not found or support ths command :" + command)
	}
}
