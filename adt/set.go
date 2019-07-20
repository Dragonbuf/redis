package adt

import "reflect"

type IntSet struct {
	encoding uint32 // 编码方式
	length   uint32 //结合包含元素数量
	contents []byte // 保存元素的数组 从小到大排列，不含重复项
}

const (
	INSERT_ENC_INIT16 = "int16"
	INSERT_ENC_INIT32 = "int32"
	INSERT_ENC_INIT64 = "int65"
)

type EncodingType struct {
}

func (i *IntSet) Add(number interface{}) {
	numberType := reflect.TypeOf(number).String()
	if numberType == "int16" {

	} else if numberType == "int32" {

	} else if numberType == "int64" {

	}
}

type EncodingEnum struct {
}
