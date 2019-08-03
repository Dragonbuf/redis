package adt

import (
	"errors"
	"strconv"
)

// type
const (
	REDIS_STRING = "string" // int embstr raw
	REDIS_LIST   = "list"   // ziplist linkedlist
	REDIS_HASH   = "hash"   // ziplist ht
	REDIS_SET    = "set"    // intset ht
	REDIS_ZSET   = "zset"   // ziplist skiplist

	// encoding
	REDIS_ENCODING_INT        = "int"
	REDIS_ENCODING_EMBSTR     = "embstr"
	REDIS_ENCODING_RAW        = "raw"
	REDIS_ENCODING_HT         = "hashtable"
	REDIS_ENCODING_LINKEDLIST = "linkedlist"
	REDIS_ENCODING_ZIPLIST    = "ziplist"
	REDIS_ENCODING_INTSET     = "intset"
	REDIS_ENCODING_SKIPLIST   = "skiplist"
)

type RedisObject struct {
	types    string
	encoding string
	*Object  //这里指向了 object 指针
}

type Object struct {
	int
	*Sdshdr
	*Dict
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
		i := len(*str)
		if i > 32 {
			obj.Sdshdr = NewSdsHdr()
			obj.Sdshdr.Set(str)

			obj.SetEncoding(REDIS_ENCODING_RAW).SetTypes(REDIS_STRING)
		} else { // 字符串长度小于 32 字节，使用 embstr ，申请释放内存只需一次 且保存在一块连续内存，更好利用缓存
			obj.Sdshdr = NewSdsHdr()
			obj.Sdshdr.Set(str)

			obj.SetEncoding(REDIS_ENCODING_EMBSTR).SetTypes(REDIS_STRING)
		}
	case int:
		obj.int = ptr.(int)
		obj.SetEncoding(REDIS_ENCODING_INT).SetTypes(REDIS_STRING)
	default:
		panic("error type")
	}

	return obj
}

// 可以设置 int string   todo: 刚开始是　ziplist 后期改成 dict
func (obj *RedisObject) Hset(filed interface{}, value interface{}) {

	obj.SetTypes(REDIS_HASH).SetEncoding(REDIS_ENCODING_HT)

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

	if obj.types != REDIS_HASH {
		return nil, errors.New("type not redis_hash")
	}

	filedObj := NewRedisObject()
	filedObj.Set(filed)

	return obj.Dict.Hget(filedObj), nil
}

// 暂时未实现方法
func (obj *RedisObject) Get() interface{} {

	switch obj.types {
	case REDIS_STRING:
		switch obj.encoding {
		case REDIS_ENCODING_INT:
			return obj.int
		default:
			return obj.Sdshdr.Get()
		}
	default:
		return nil
	}

}

func (obj *RedisObject) Append(str *string) int {
	if obj.types != REDIS_STRING {
		return -2 // todo 处理错误码，返回信息
	}

	// 如果之前是 int ，需要先转成 string
	if obj.encoding == REDIS_ENCODING_INT {
		*str = strconv.Itoa(obj.int) + *str
		obj.SetEncoding(REDIS_ENCODING_EMBSTR)
	}

	obj.Set(str)
	return obj.GetLen()
}
func (obj *RedisObject) GetType() string {
	return obj.types
}
func (obj *RedisObject) GetEncoding() string {
	return obj.encoding
}
