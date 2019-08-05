package adt

import (
	"reflect"
	"testing"
)

func TestNewRedisObject(t *testing.T) {
	// set 1
	ptr := &Object{1, nil, nil, nil}
	obj := NewRedisObject().SetTypes(REDIS_STRING).SetEncoding(REDIS_ENCODING_INT).SetPtr(ptr)

	if obj.types != REDIS_STRING {
		t.Error("wrong types")
	}

	if obj.encoding != REDIS_ENCODING_INT {
		t.Error("wrong encoding")
	}

	if obj.int != 1 {
		t.Error("wrong int")
	}

	obj.Set(4)

	if obj.types != REDIS_STRING {
		t.Error("wrong types")
	}

	if obj.encoding != REDIS_ENCODING_INT {
		t.Error("wrong encoding")
	}

	if obj.int != 4 {
		t.Error("wrong int")
	}

	str := "you"
	obj.Set(&str)

	if reflect.TypeOf(obj.Get()).String() != "*string" {
		t.Error("wrong type", reflect.TypeOf(obj.Get()).String())
	}

	if obj.types != REDIS_STRING {
		t.Error("wrong types")
	}

	if obj.encoding != REDIS_ENCODING_EMBSTR {
		t.Error("wrong encoding")
	}

	if *obj.Sdshdr.Get() != "you" {
		t.Error("wrong int ", *obj.Sdshdr.Get())
	}
}

func BenchmarkRedisObject_Set(b *testing.B) {

	obj := NewRedisObject()
	for i := 0; i < b.N; i++ {
		obj.Set(i)
	}

}
