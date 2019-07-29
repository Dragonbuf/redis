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

func (r *RedisDb) Set(key, value *adt.StringObject) {
	r.dict.Hset(key, value)
}

func (r *RedisDb) SetString(key, value *adt.StringObject) {
	r.dict.HsetString(key, value)
}

func (r *RedisDb) GetString(key *adt.StringObject) string {
	return r.dict.HgetString(key)
}

func (r *RedisDb) HSetString(key, filed, value *adt.StringObject) {
	// todo 先 find ，如果找到了 hash 就使用之前的，找不到在重新生成
	// 或者分两步，首先找到 getOrSetHash 然后在 set
	dictValue := r.dict.Hget(key)

	if dictValue != nil && dictValue.GetType() == adt.REDIS_ENCODING_HT {
		dictValue.HashObject.Hset(filed, value)
	} else {
		r.dict.HsetHash(key, filed, value)
	}
}

func (r *RedisDb) HGetString(key, filed *adt.StringObject) string {
	return r.dict.HgetGetHash(key, filed)
}

func (r *RedisDb) SetHash(key, value *adt.StringObject) {
	r.dict.Hset(key, value)
}

func (r *RedisDb) SetList(key, value *adt.StringObject) {
	r.dict.Hset(key, value)
}

func Test() {

}
