package adt

type RedisObject struct {
	types    interface{}
	encoding interface{}
	ptr      *interface{}
}

func NewRedisObject() *RedisObject {
	return &RedisObject{}
}

func (obj *RedisObject) SetTypes(t string) {
	obj.types = t
}

func (obj *RedisObject) SetPtr(ptr *interface{}) {
	obj.ptr = ptr
}
