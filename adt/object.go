package adt

type Object struct {
	strObj  *StringObject
	hashOjb *HashObject
	listObj *ListObject
}

type StringObject struct {
	Sds *Sdshdr
}

type HashObject struct {
}

type ListObject struct {
}
