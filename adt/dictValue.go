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

func (d *DictValue) ToString() string {
	if d == nil {
		return "<nil>"
	}

	if d.valueType != DictvalueTypeStringObj {
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

func (d *DictValue) SetHashValue(filed, value *StringObject) *DictValue {
	dict := NewDict()
	dict.HsetString(filed, value)
	return &DictValue{
		HashObject: &HashObject{Dict: dict},
		valueType:  DictvalueTypeHashObj}
}

func (d *DictValue) SetHashObjValue(filed, value *StringObject) {
	d.Dict.HsetString(filed, value)
}
