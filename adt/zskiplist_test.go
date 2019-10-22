package adt

import (
	"fmt"
	"testing"
	"time"
)

type Server struct {
	Addr    string
	timeout time.Duration
	tls     string
}

type config struct {
	tls string
}

type ZSkip struct {
	Head  string
	Tail  string
	Level int
}

func NewZSip(options ...func(skip *ZSkip)) *ZSkip {
	zs := &ZSkip{}

	for _, option := range options {
		option(zs)
	}

	return zs
}

func CreateHead(name string) func(*ZSkip) {
	return func(skip *ZSkip) {
		skip.Head = name
	}
}

func CreateLevel(l int) func(*ZSkip) {
	return func(skip *ZSkip) {
		skip.Level = l
	}
}

func NewServer(addr string, options ...func(*Server)) (*Server, error) {
	srv := &Server{
		Addr: addr,
	}

	for _, option := range options {
		option(srv)
	}

	return srv, nil
}

func timeout(d time.Duration) func(*Server) {
	return func(srv *Server) {
		srv.timeout = d
	}
}

func tls(c *config) func(*Server) {
	return func(srv *Server) {
		Tls := loadConfig(c)
		srv.tls = Tls
	}
}

func loadConfig(c *config) string {
	return c.tls
}

//使用
func BenchmarkZSkipList_Create(b *testing.B) {
	src, err := NewServer("localhost:8080", timeout(1), tls(&config{tls: "12"}))
	fmt.Println(src, err)

	sk := NewZSip(CreateHead("head"), CreateLevel(666))
	fmt.Println(sk.Level)
}
