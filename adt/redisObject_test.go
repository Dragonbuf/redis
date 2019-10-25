package adt

import (
	"reflect"
	"testing"
)

func TestNewRedisObject(t *testing.T) {

	obj := NewRedisObject().SetTypes(RedisString).SetEncoding(RedisEncodingInt).SetPtr(&Object{})

	if obj.types != RedisString {
		t.Error("wrong types")
	}

	if obj.encoding != RedisEncodingInt {
		t.Error("wrong encoding")
	}

	obj.Set(4)

	if obj.types != RedisString {
		t.Error("wrong types")
	}

	if obj.encoding != RedisEncodingInt {
		t.Error("wrong encoding")
	}

	if obj.int != 4 {
		t.Error("wrong int")
	}

	str := "you"
	obj.Set(&str)

	if reflect.TypeOf(obj.Get()).String() != "unsafe.Pointer" {
		t.Error("wrong type", reflect.TypeOf(obj.Get()).String())
	}

	if obj.types != RedisString {
		t.Error("wrong types")
	}

	if string(*(*[]byte)(obj.Ele)) != "you" {
		t.Error("wrong int ", string(*(*[]byte)(obj.Ele)))
	}
}

func BenchmarkRedisObject_Set(b *testing.B) {

	obj := NewRedisObject()
	for i := 0; i < b.N; i++ {
		obj.Set(i)
	}

}
