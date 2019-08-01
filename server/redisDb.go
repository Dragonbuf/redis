package server

import (
	"errors"
	"redis/adt"
)

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

func (r *RedisDb) Set(key *adt.StringObject, value *adt.DictValue) {
	r.dict.Hset(key, value)
}

func (r *RedisDb) Get(key *adt.StringObject) string {
	return r.dict.Hget(key).ToString()
}

func (r *RedisDb) HSet(key, filed, value *adt.StringObject) (err error) {

	// 先找到 redisDb 中实际存值的　hash
	dictValue := r.dict.Hget(key)

	if dictValue != nil {
		if dictValue.GetType() != adt.REDIS_HASH {
			return errors.New("can not use another type")
		}

		dictValue.HashObject.Hset(filed, adt.NewDictValue().SetStringObject(value))
	} else {
		// 先把 filed => value 的 dict 生成
		dict := adt.NewDict().Hset(filed, adt.NewDictValue().SetStringObject(value))

		// 再把 key => dict 存入 r.dict
		r.dict = adt.NewDict().Hset(key, adt.NewDictValue().SetHashObject(adt.NewHashObject().SetDict(dict)))
	}

	return nil
}

func (r *RedisDb) HGet(key, filed *adt.StringObject) string {

	// 先找到 redisDb 中实际存值的　hash
	dictValue := r.dict.Hget(key)

	if dictValue != nil {
		if dictValue.GetType() != adt.REDIS_HASH {
			return "can not use hget if not hash"
		}

		return dictValue.HashObject.Hget(filed).ToString()
	}

	return "not found"
}

func (r *RedisDb) SetList(key, value *adt.StringObject) {

}

func Test() {

}
