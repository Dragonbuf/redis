package adt

type DictValue struct {
	*StringObject
	*HashObject
	*ListObject
	*RedisObject
	int64Obj  int64
	uint64Obj uint64
	valueType string
}

func NewDictValue() *DictValue {
	return &DictValue{}
}

func (d *DictValue) SetType(types string) *DictValue {
	d.valueType = types
	return d
}
func (d *DictValue) SetStringObject(obj *StringObject) *DictValue {
	d.StringObject = obj
	d.valueType = REDIS_STRING
	return d
}
func (d *DictValue) SetHashObject(obj *HashObject) *DictValue {
	d.HashObject = obj
	d.valueType = REDIS_HASH
	return d
}
func (d *DictValue) SetRedisObject(obj *RedisObject) *DictValue {
	d.RedisObject = obj
	return d
}

func (d *DictValue) ToString() string {
	if d == nil {
		return "<nil>"
	}

	if d.valueType != REDIS_STRING {
		return "this key type is [" + d.valueType + "] can not use get ,please use another command"
	}

	return string(*d.buf)
}

func (d *DictValue) GetType() string {
	return d.valueType
}
