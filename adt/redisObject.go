package adt

import (
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
	*StringObject
	*HashObject
	*ListObject
}

type Object struct {
	*StringObject
	*HashObject
	*ListObject
}

type ListObject struct {
}

func NewRedisObject() *RedisObject {
	return &RedisObject{}
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

// 这里是 stringObj
func (obj *RedisObject) SetIntStringObj(num int) {
	object := &Object{NewStringObject().SetInt(num), nil, nil}
	obj.SetEncoding(REDIS_ENCODING_INT).SetTypes(REDIS_STRING).SetPtr(object)
}

func (obj *RedisObject) SetRawStringObj(str *string) {

	sds := NewSdsHdr()
	sds.Set(str)

	object := &Object{NewStringObject().SetSds(sds), nil, nil}
	obj.SetEncoding(REDIS_ENCODING_RAW).SetTypes(REDIS_STRING).SetPtr(object)
}

// 字符串长度小于 32 字节，使用 embstr ，申请释放内存只需一次 且保存在一块连续内存，更好利用缓存
func (obj *RedisObject) SetEmbStrStringObj(str *string) {
	sds := NewSdsHdr()
	sds.Set(str)

	object := &Object{NewStringObject().SetSds(sds), nil, nil}
	obj.SetEncoding(REDIS_ENCODING_EMBSTR).SetTypes(REDIS_STRING).SetPtr(object)
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
