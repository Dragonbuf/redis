package server

import (
	"errors"
	"fmt"
	"redis/adt"
	"strconv"
	"time"
)

const NilString = "<nil>"

type RedisDb struct {
	dict   *adt.Dict // 数据库空间，保存所有键值对
	expire *adt.Dict // 过期字典，保存键过期时间
}

func NewRedisDb() *RedisDb {
	db := &RedisDb{}
	dict := adt.NewDict()
	db.dict = dict
	return db
}

func (r *RedisDb) Set(key, value string) error {

	k := adt.NewRedisObject().Set(&key)

	// 先找到 redisDb 中实际存值的　hash
	existsRedisObj := r.dict.Hget(k)

	if existsRedisObj != nil {
		if existsRedisObj.GetType() != adt.RedisString {
			errorString := fmt.Sprint("已存在类型为 ", existsRedisObj.GetType(), " 的 (", key, "), 不能设置为 ", adt.RedisString)
			return errors.New(errorString)
		}

		existsRedisObj.Hset(adt.NewRedisObject().Set(&key), adt.NewRedisObject().Set(&value))
	} else {
		r.dict.Hset(adt.NewRedisObject().Set(&key), adt.NewRedisObject().Set(&value))
	}

	return nil
}

func (r *RedisDb) Get(key string) string {
	tarObj := r.dict.Hget(adt.NewRedisObject().Set(&key))
	if tarObj == nil {
		return NilString
	}
	return string(*(*[]byte)(tarObj.Ele))
}

func (r *RedisDb) HSet(key, filed, value string) (err error) {

	k := adt.NewRedisObject().Set(&key)

	// 先找到 redisDb 中实际存值的　hash
	existsRedisObj := r.dict.Hget(k)

	if existsRedisObj != nil {
		if existsRedisObj.GetType() != adt.RedisHash {
			errorString := fmt.Sprint("已存在类型为 ", existsRedisObj.GetType(), " 的 (", key, "), 不能设置为 ", adt.RedisHash)
			return errors.New(errorString)
		}

		existsRedisObj.Hset(&filed, &value)
	} else {
		redisObj := adt.NewRedisObject()
		redisObj.Hset(&filed, &value)

		// 再把 key => dict 存入 r.dict
		r.dict = adt.NewDict().Hset(k, redisObj)
	}

	return nil
}

func (r *RedisDb) HGet(key, filed string) string {

	k := adt.NewRedisObject().Set(&key)
	f := adt.NewRedisObject().Set(&filed)
	existsRedisObj := r.dict.Hget(k)

	if existsRedisObj != nil {
		if existsRedisObj.GetType() != adt.RedisHash {
			return "can not use this get " + existsRedisObj.GetType()
		}

		targetObj := existsRedisObj.Hget(f)

		if targetObj == nil {
			return NilString
		}

		return string(*(*[]byte)(targetObj.Ele))
	}

	return NilString
}

func (r *RedisDb) Del(key string) int {
	return r.dict.Hdel(adt.NewRedisObject().Set(&key))
}

func (r *RedisDb) Hdel(key, filed string) int {

	k := adt.NewRedisObject().Set(&key)
	f := adt.NewRedisObject().Set(&filed)
	existsRedisObj := r.dict.Hget(k)

	if existsRedisObj != nil {
		if existsRedisObj.GetType() != adt.RedisHash {
			return 0
		}
		return existsRedisObj.Hdel(f)
	}

	return 0
}

func (r *RedisDb) RPush(key string, value ...string) {
	obj := adt.NewRedisObject()
	obj.RPush(value)

	r.dict.Hset(adt.NewRedisObject().Set(&key), obj)
}

func (r *RedisDb) RPop(key string) string {
	tarObj := r.dict.Hget(adt.NewRedisObject().Set(&key))
	if tarObj == nil || tarObj.List.IsEmpty() {
		return "<nil>"
	}
	return string(*(*[]byte)(tarObj.List.RPop().Ele))
}
func (r *RedisDb) Expire(key, value string) int {

	if r.expire == nil {
		dict := adt.NewDict()
		r.expire = dict
	}

	second, err := strconv.Atoi(value)
	if err != nil {
		return 0
	}

	i := time.Now().Unix() + int64(second)
	r.expire.Hset(adt.NewRedisObject().Set(&key), adt.NewRedisObject().Set(i))
	return 1
}

func (r *RedisDb) ExpireIfNeeded(key string) bool {

	if r.expire == nil {
		return false
	}

	obj := r.expire.Hget(adt.NewRedisObject().Set(&key))

	if obj == nil {
		return false
	}

	return obj.GetExpireSecond() < time.Now().Unix()
}

func (r *RedisDb) DelExpire(key string) int {
	return r.expire.Hdel(adt.NewRedisObject().Set(&key))
}
