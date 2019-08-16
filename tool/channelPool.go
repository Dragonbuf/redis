package tool

import "sync"

type GoroutinePool struct {
	c  chan struct{}
	wg *sync.WaitGroup
}

// 有缓冲 channel
func NewGoroutinePool(maxSize int) *GoroutinePool {
	if maxSize <= 0 {
		panic("too small max size")
	}
	return &GoroutinePool{
		make(chan struct{}, maxSize),
		new(sync.WaitGroup),
	}
}

func (g *GoroutinePool) Add(delta int) {
	g.wg.Add(delta)
	for i := 0; i < delta; i++ {
		g.c <- struct{}{}
	}
}

func (g *GoroutinePool) Done() {
	<-g.c
	g.wg.Done()
}

func (g *GoroutinePool) Wait() {
	g.wg.Wait()
}
