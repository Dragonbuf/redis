package adt

type StringObject struct {
	int
	*Sdshdr
}

func NewStringObject() *StringObject {
	return &StringObject{}
}

func (obj *StringObject) SetInt(num int) *StringObject {
	obj.int = num
	return obj
}

func (obj *StringObject) GetInt() int {
	return obj.int
}

func (obj *StringObject) SetString(str *string) *StringObject {
	obj.Sdshdr = NewSdsHdr()
	obj.Sdshdr.Set(str)
	return obj
}

func (obj *StringObject) SetSds(sds *Sdshdr) *StringObject {
	obj.Sdshdr = sds
	return obj
}
