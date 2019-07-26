package adt

type DictValue struct {
	obj       *Object
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

	return string(*d.obj.strObj.Sds.buf)
}

func (d *DictValue) GetType() string {
	return d.valueType
}

func (d *DictValue) SetStringValue(value *string) *DictValue {
	sds := NewSdsHdr()
	sds.Set(value)
	return &DictValue{
		obj:       &Object{strObj: &StringObject{Sds: sds}},
		valueType: DictvalueTypeStringObj}
}

func (d *DictValue) SetHashValue(filed *string, value *string) *DictValue {
	dict := NewDict()
	dict.HsetString(filed, value)
	return &DictValue{
		obj:       &Object{hashOjb: &HashObject{Dict: dict}},
		valueType: DictvalueTypeHashObj}
}

func (d *DictValue) SetHashObjValue(filed *string, value *string) {
	d.obj.hashOjb.Dict.HsetString(filed, value)
}
