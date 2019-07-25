package server

import (
	"fmt"
	"testing"
)

func TestNewRedisDb(t *testing.T) {
	db := NewRedisDb()

	key := "fuck"
	key2 := "c"
	value := "you"

	db.SetString(&key, &value)

	fmt.Println(db.GetString(&key2))

	if db.GetString(&key2) != "<nil>" {
		t.Error("nil key found value")
	}

	if db.GetString(&key) != value {
		t.Error("value not found")
	}

}
