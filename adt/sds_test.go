package adt

import (
	"testing"
)

func TestSds(t *testing.T) {
	sds := NewSdsHdr()
	sds.Set("hello")
	sds.Set("World")
	sds.Set("helloWorld")
	sds.Set("helloWorld 2")
}
