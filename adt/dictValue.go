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

func (d *DictValue) ToString() string {
	if d == nil {
		return "<nil>"
	}

	if d.valueType != REDIS_STRING {
		return "this key can not use get XXX,please use another command"
	}

	return string(*d.buf)
}

func (d *DictValue) GetType() string {
	return d.valueType
}

func (d *DictValue) SetStringValue(value *string) *DictValue {
	sds := NewSdsHdr()
	sds.Set(value)
	return &DictValue{
		StringObject: NewStringObject().SetSds(sds),
		valueType:    DictvalueTypeStringObj}
}
