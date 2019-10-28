package adt

import "unsafe"

type StringObject struct {
	int
	ele unsafe.Pointer
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
	obj.ele = SdsNewLen([]byte(*str))
	return obj
}
