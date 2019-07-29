package adt

type HashObject struct {
	*Dict
}

func NewHashObject() *HashObject {
	return &HashObject{}
}

func (obj *HashObject) SetKeyValue(key, value *StringObject) {

}
