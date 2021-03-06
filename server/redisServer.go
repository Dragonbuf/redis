package server

import (
	"redis/adt"
	"unsafe"
)

var Server *RedisServer
var Db *RedisDb

func init() {
	Server = NewRedisServer()
	Db = Server.Select(0)
}

type RedisServer struct {
	dbNumber  int
	currentDb int
	redisDb   [16]*RedisDb // 这里写死，只有 16 个数据库
	saveParams
	dirty          int64          // 修改计数器
	timeT          int64          // 上一次执行保存时间
	aofBuffer      unsafe.Pointer // AOF 缓冲区
	clients        adt.List       // 保存所有客户端状态
	statPeakMemory int64          // 已使用内存峰值
}

type saveParams struct {
	timeT   int // 秒数
	changes int // 修改数
}

func NewRedisServer() *RedisServer {
	var redisDb [16]*RedisDb
	for i := 0; i < 16; i++ {
		redisDb[i] = NewRedisDb()
	}

	server := &RedisServer{}
	server.dbNumber = 16
	server.redisDb = redisDb
	return server
}

func (c *RedisServer) Select(db int) *RedisDb {
	Db = c.redisDb[db]
	c.currentDb = db
	return c.redisDb[db]
}

func (c *RedisServer) BGSave() {
	ch := make(chan int, 1)
	<-ch // 后台处理保存
}
func (c *RedisServer) Save() {
	// 立即保存
}

func (c *RedisServer) GetDbTotal() int {
	return c.dbNumber
}

func (c *RedisServer) GetCurrentDbNumber() int {
	return c.currentDb
}

func (c *RedisServer) IncrDirty() {
	c.dirty++
}
