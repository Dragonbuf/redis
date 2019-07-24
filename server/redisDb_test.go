package server

import (
	"fmt"
	"testing"
)

func TestNewRedisDb(t *testing.T) {
	db := NewRedisDb()

	key := "hello"
	key2 := "a"
	value := "workd"

	db.SetString(&key, &value)
	db.SetString(&key, &value)
	db.SetString(&key, &value)
	db.SetString(&key, &value)
	db.SetString(&key, &value)
	db.GetString(&key2)
	fmt.Println(db.GetString(&key2))

	if db.GetString(&key2) != "<nil>" {
		t.Error("nil key found value")
	}

	if db.GetString(&key) != value {
		t.Error("value not found")
	}

}
