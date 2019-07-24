package server

import "redis/adt"

// todo 先实现 dict 不然不好存储　redisDb
type RedisDb struct {
	dict *adt.Dict // 数据库空间，保存所有键值对
}

func NewRedisDb() *RedisDb {
	db := &RedisDb{}
	dict := adt.NewDict()
	db.dict = dict
	return db
}

func (r *RedisDb) Set(key *string, value *adt.DictValue) {
	r.dict.Hset(key, value)
}

func (r *RedisDb) SetString(key, value *string) {
	r.dict.HsetString(key, value)
}

func (r *RedisDb) GetString(key *string) string {
	return r.dict.HgetString(key)
}

func (r *RedisDb) SetHash(key *string, value *adt.DictValue) {
	r.dict.Hset(key, value)
}

func (r *RedisDb) SetList(key *string, value *adt.DictValue) {
	r.dict.Hset(key, value)
}

func Test() {

}
