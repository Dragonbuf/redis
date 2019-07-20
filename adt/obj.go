package adt

type RedisObject struct {
	type4    uint8
	encoding uint8
	ptr      *Obj
}
