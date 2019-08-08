package server

import "redis/adt"

type RedisServer struct {
	dbNumber int
	redisDb  [16]*RedisDb // 这里写死，只有 16 个数据库
	saveParams
	dirty            int64      // 修改计数器
	time_t           int64      // 上一次执行保存时间
	aofBuffer        adt.Sdshdr // AOF 缓冲区
	clients          adt.List   // 保存所有客户端状态
	stat_peak_memory int64      // 已使用内存峰值
}

type saveParams struct {
	time_t  int // 秒数
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
	return c.redisDb[db]
}

func (c *RedisServer) BGSave() {
	ch := make(chan int, 1)
	<-ch // 后台处理保存
}
func (c *RedisServer) Save() {
	// 立即保存
}

func (c *RedisServer) GetDbNum() int {
	return c.dbNumber
}

func (c *RedisServer) IncrDirty() {
	c.dirty++
}
