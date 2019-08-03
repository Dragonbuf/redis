package adt

import (
	"strconv"
	"testing"
)

func TestSds(t *testing.T) {

	sds := NewSdsHdr()
	hello := "hello"
	sds.Set(&hello)
	if *sds.Get() != hello {
		t.Error("get hello ")
	}

	world := "world"
	sds.Set(&world)
	if *sds.Get() != world {
		t.Error("get  world")
	}

	helloWorld := "helloWorld"
	sds.Set(&helloWorld)

	if *sds.Get() != helloWorld {
		t.Error("get hello world")
	}
}

func BenchmarkSdshdr_Set(b *testing.B) {
	sds := NewSdsHdr()
	for i := 0; i < b.N; i++ {
		str := strconv.Itoa(i)
		sds.Set(&str)
	}
}
