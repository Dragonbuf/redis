package server

import (
	"testing"
)

func TestNewRedisDb(t *testing.T) {
	db := NewRedisDb()
	key := "fuck"
	key2 := "c"
	value := "you"

	if db.Get(key2) != "<nil>" {
		t.Error("nil key found value")
	}

	db.Set(key, value)
	if db.Get(key) != value {
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

	filed3 := "c3"

	err := db.HSet(key, filed, value)
	if err != nil {
		t.Error(err)
	}

	err = db.HSet(key, filed2, value2)
	if err != nil {
		t.Error(err)
	}

	if db.HGet(key, filed) != value {
		t.Error("not equal value " + value)
	}

	if db.HGet(key, filed2) != value2 {
		t.Error("not equal value " + value2)
	}

	if db.HGet(key, filed3) != "<nil>" {
		t.Error("not equal filed3 " + filed3)
	}
}
