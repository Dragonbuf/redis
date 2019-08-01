package adt

import (
	"testing"
)

func TestNewRedisObject(t *testing.T) {
	// set 1
	ptr := &Object{NewStringObject().SetInt(1), nil, nil}
	obj := NewRedisObject().SetTypes(REDIS_STRING).SetEncoding(REDIS_ENCODING_INT).SetPtr(ptr)

	if obj.types != REDIS_STRING {
		t.Error("wrong types")
	}

	if obj.encoding != REDIS_ENCODING_INT {
		t.Error("wrong encoding")
	}

	if obj.StringObject.int != obj.GetInt() {
		t.Error("wrong int")
	}

	obj.Set(4)

	if obj.types != REDIS_STRING {
		t.Error("wrong types")
	}

	if obj.encoding != REDIS_ENCODING_INT {
		t.Error("wrong encoding")
	}

	if obj.StringObject.int != obj.GetInt() {
		t.Error("wrong int")
	}

	obj.Set("you")

	if obj.types != REDIS_STRING {
		t.Error("wrong types")
	}

	if obj.encoding != REDIS_ENCODING_EMBSTR {
		t.Error("wrong encoding")
	}

	if *obj.StringObject.Get() != "you" {
		t.Error("wrong int ", *obj.StringObject.Get())
	}
}
