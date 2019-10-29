package adt

import (
	"container/list"
	"net/http"
	"strconv"
	"sync"
	"time"
)

var S *Store

type Store struct {
	Mutex *sync.Mutex
	store map[string]*list.Element
	ll    *list.List
	max   int
}

// 直接返回 []byte 当 write 时，节省转换的时间
func (s *Store) Get(key string) ([]byte, bool) {
	current, exist := s.store[key]
	if exist {
		expire := int64(current.Value.(*Node).expire)
		if expire == 0 || expire > time.Now().Unix() {
			s.ll.MoveToFront(current)
			return current.Value.(*Node).value, true
		}
	}
	return nil, false
}

func (s *Store) Set(key string, value []byte, expire int) {
	current, exist := s.store[key]
	if !exist {
		s.store[key] = s.ll.PushFront(&Node{
			key:    key,
			value:  value,
			expire: expire,
		})

		if s.max != 0 && s.ll.Len() > s.max {
			s.Delete(s.ll.Remove(s.ll.Back()).(*Node).key)
		}
		return
	}

	current.Value.(*Node).value = value
	current.Value.(*Node).expire = expire
	s.ll.MoveToFront(current)
}

func (s *Store) Delete(key string) {
	delete(s.store, key)
}

func (s *Store) Flush() {
	s.store = make(map[string]*list.Element)
	s.ll = list.New()
}

type Node struct {
	key    string
	value  []byte
	expire int
}

func Handler(f func(http.ResponseWriter, *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		S.Mutex.Lock()
		defer S.Mutex.Unlock()
		f(w, r)
	}
}

func Get(w http.ResponseWriter, r *http.Request) {
	value, ok := S.Get(r.URL.Query().Get("key"))
	if !ok {
		_, _ = w.Write([]byte("no such key"))
		return
	}

	w.Header().Set("content-type", "text/plain")
	_, _ = w.Write(value)
}

func Set(w http.ResponseWriter, r *http.Request) {
	expire, _ := strconv.Atoi(r.URL.Query().Get("expire"))
	S.Set(r.URL.Query().Get("key"), []byte(r.URL.Query().Get("value")), expire)
	value, ok := S.Get(r.URL.Query().Get("key"))

	if !ok {
		_, _ = w.Write([]byte("no such key"))
		return
	}

	w.Header().Set("content-type", "text/plain")
	_, _ = w.Write(value)
}

func main() {
	S = &Store{
		Mutex: &sync.Mutex{},
		store: make(map[string]*list.Element),
		ll:    list.New(),
		max:   100,
	}
	http.HandleFunc("/get", Handler(Get))
	http.HandleFunc("/set", Handler(Set))
	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		panic(err)
	}

}
