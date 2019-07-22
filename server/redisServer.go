package server

type RedisServer struct {
	dbNumber int
	redisDb  *[]RedisDb
}

func (c *RedisServer) Select(db int) {

}
