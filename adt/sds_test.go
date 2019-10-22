package adt

import (
	"fmt"
	"strconv"
	"testing"
)

type Num struct {
	i string
	j int
}

func TestSds(t *testing.T) {

	newSds8Buf := SdsNewLen([]byte("hello world2, do you know youhello world2, do you know you "))
	flagsPoint := GetFlagsPointByBufPoint(newSds8Buf)
	flagsValue := (*uint8)(flagsPoint)

	fmt.Println(*flagsValue)
}

func BenchmarkSdshdr_Set(b *testing.B) {
	sds := NewSdsHdr()
	for i := 0; i < b.N; i++ {
		str := strconv.Itoa(i)
		sds.Set(&str)
	}
}
