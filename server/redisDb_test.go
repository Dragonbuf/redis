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
	value1 := "notyouasc"

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

	err = db.HSet(key, filed, value1)
	if err != nil {
		t.Error(err)
	}

	if db.HGet(key, filed) == value {
		t.Error("not equal value " + db.HGet(key, filed))
	}
	if db.HGet(key, filed) != value1 {
		t.Error("not equal value1 " + value1)
	}
}

func TestRedisDb_Del(t *testing.T) {
	db := NewRedisDb()
	db.Set("a", "b")
	db.Del("a")

	if db.Get("a") != "<nil>" {
		t.Error("can not get del key")
	}
}

func TestRedisDb_Hdel(t *testing.T) {
	db := NewRedisDb()
	_ = db.HSet("a", "1", "1")
	_ = db.HSet("a", "2", "2")
	db.Hdel("a", "1")

	if db.HGet("a", "1") != "<nil>" {
		t.Error("can not get hdel key")
	}
}

func TestRedisDb_RPush(t *testing.T) {

	db := NewRedisDb()

	listName := "newList"
	a := "a"
	b := "b"

	db.RPush(listName, a, b)
	if db.RPop(listName) != b {
		t.Error("pop not b")
	}

	if db.RPop(listName) != a {
		t.Error("pop not a")
	}

	if db.RPop(listName) != "<nil>" {
		t.Error("pop not <nil>")
	}
}
