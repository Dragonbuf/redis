package adt

import (
	"errors"
	"unsafe"
)

// type
const (
	RedisString = "string" // int embstr raw
	RedisList   = "list"   // ziplist linkedlist
	RedisHash   = "hash"   // ziplist ht
	RedisSet    = "set"    // intset ht
	RedisZset   = "zset"   // ziplist skiplist

	// encoding
	RedisEncodingInt        = "int"
	RedisEncodingEmbstr     = "embstr"
	RedisEncodingRaw        = "raw"
	RedisEncodingHt         = "hashtable"
	RedisEncodingLinkedlist = "linkedlist"
	RedisEncodingZiplist    = "ziplist"
	RedisEncodingIntset     = "intset"
	RedisEncodingSkiplist   = "skiplist"
)

type RedisObject struct {
	Ele      unsafe.Pointer
	types    string
	encoding string
	*Object  //这里指向了 object 指针
	lru      int64
}

type Object struct {
	int
	*Sdshdr
	*Dict
	*List
	int64 // 存储过期时间
}

func NewRedisObject() *RedisObject {
	obj := &RedisObject{}
	obj.Object = &Object{}
	return obj
}
func (obj *RedisObject) SetTypes(types string) *RedisObject {
	obj.types = types
	return obj
}
func (obj *RedisObject) SetEncoding(encoding string) *RedisObject {
	obj.encoding = encoding
	return obj
}
func (obj *RedisObject) SetPtr(ptr *Object) *RedisObject {
	obj.Object = ptr
	return obj
}

// 可以设置 int string
func (obj *RedisObject) Set(ptr interface{}) *RedisObject {

	switch ptr.(type) {
	case *string:
		str := ptr.(*string)
		obj.Ele = SdsNewLen([]byte(*str))
		obj.SetEncoding(RedisEncodingRaw).SetTypes(RedisString)
	case int:
		obj.int = ptr.(int)
		obj.SetEncoding(RedisEncodingInt).SetTypes(RedisString)
	case int64:
		obj.int64 = ptr.(int64)
	default:
		panic("error type")
	}

	return obj
}

// 可以设置 int string   todo: 刚开始是　ziplist 后期改成 dict
func (obj *RedisObject) Hset(filed interface{}, value interface{}) {

	obj.SetTypes(RedisHash).SetEncoding(RedisEncodingHt)

	filedObj := NewRedisObject()
	filedObj.Set(filed)

	vObj := NewRedisObject()
	vObj.Set(value)

	if obj.Dict != nil {
		obj.Dict.Hset(filedObj, vObj)
	} else {
		obj.Dict = NewDict().Hset(filedObj, vObj)
	}

}

func (obj *RedisObject) HGet(filed interface{}) (*RedisObject, error) {

	if obj.types != RedisHash {
		return nil, errors.New("type not redis_hash")
	}

	filedObj := NewRedisObject()
	filedObj.Set(filed)

	return obj.Dict.Hget(filedObj), nil
}

func (obj *RedisObject) RPush(strings []string) {
	obj.SetTypes(RedisList).SetEncoding(RedisEncodingLinkedlist)
	obj.List = NewList()

	for _, v := range strings {
		obj.List.RPush(NewRedisObject().Set(&v))
	}
}
func (obj *RedisObject) RPop() *RedisObject {
	if obj.List == nil {
		return nil
	}

	return obj.List.RPop()
}

// 暂时未实现方法
func (obj *RedisObject) Get() interface{} {

	switch obj.types {
	case RedisString:
		switch obj.encoding {
		case RedisEncodingInt:
			return obj.int
		default:
			return obj.Ele
		}
	default:
		return nil
	}

}

func (obj *RedisObject) GetType() string {
	return obj.types
}

func (obj *RedisObject) GetExpireSecond() int64 {
	return obj.int64
}
