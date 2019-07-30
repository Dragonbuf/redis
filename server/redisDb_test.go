package server

import (
	"redis/adt"
	"testing"
)

func TestNewRedisDb(t *testing.T) {
	db := NewRedisDb()
	key := "fuck"
	key2 := "c"
	value := "you"

	db.Set(adt.NewStringObject().SetString(&key), adt.NewDictValue().SetStringObject(adt.NewStringObject().SetString(&value)))
	if db.Get(adt.NewStringObject().SetString(&key2)) != "<nil>" {
		t.Error("nil key found value")
	}

	if db.Get(adt.NewStringObject().SetString(&key)) != value {
		t.Error("value not found")
	}
}

func TestHashSet(t *testing.T) {
	db := NewRedisDb()
	key := "fuck"
	filed := "c"
	value := "you"

	filed2 := "c2"
	value2 := "you2"

	err := db.HSet(adt.NewStringObject().SetString(&key), adt.NewStringObject().SetString(&filed), adt.NewStringObject().SetString(&value))
	if err != nil {
		t.Error(err)
	}

	err = db.HSet(adt.NewStringObject().SetString(&key), adt.NewStringObject().SetString(&filed2), adt.NewStringObject().SetString(&value2))
	if err != nil {
		t.Error(err)
	}

	if db.HGet(adt.NewStringObject().SetString(&key), adt.NewStringObject().SetString(&filed2)) != value2 {
		t.Error("keys2 error")
	}

	if db.HGet(adt.NewStringObject().SetString(&key), adt.NewStringObject().SetString(&filed)) != value {
		t.Error("keys2 error")
	}
}
