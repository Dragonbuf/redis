package client

import "redis/server"

type RedisClient struct {
	db *server.RedisDb
}

func (c *RedisClient) Select(db int) {

}
