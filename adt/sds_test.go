package adt

import (
	"testing"
)

func TestSds(t *testing.T) {
	sds := NewSdsHdr()
	hello := "hello"
	sds.Set(&hello)

	world := "world"
	sds.Set(&world)

	helloWorld := "helloWorld"
	sds.Set(&helloWorld)
}
