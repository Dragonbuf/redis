package server

import (
	"testing"
)

func TestNewRedisDb(t *testing.T) {
	db := NewRedisDb()
	key := "fuck"
	key2 := "c"
	value := "you"

	db.SetString(&key, &value)
	if db.GetString(&key2) != "<nil>" {
		t.Error("nil key found value")
	}

	if db.GetString(&key) != value {
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

	db.HSetString(&key, &filed, &value)
	db.HSetString(&key, &filed2, &value2)

	if db.HGetString(&key, &filed2) != value2 {
		t.Error("keys2 error")
	}

	if db.HGetString(&key, &filed) != value {
		t.Error("keys2 error")
	}
}
