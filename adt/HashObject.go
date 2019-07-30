package adt

type HashObject struct {
	*Dict
}

func NewHashObject() *HashObject {
	return &HashObject{}
}

func (obj *HashObject) SetDict(dict *Dict) *HashObject {
	obj.Dict = dict
	return obj
}

func (obj *HashObject) SetKeyValue(key, value *StringObject) {

}
