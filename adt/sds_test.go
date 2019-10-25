package adt

import (
	"testing"
)

func TestSds(t *testing.T) {

	newSds8Buf := SdsNewLen([]byte("hello world2, do you know youhello world2, do you know you "))
	flagsPoint := GetFlagsPointByBufPoint(newSds8Buf)
	flagsValue := (*uint8)(flagsPoint)

	if *flagsValue != 3 {
		t.Error("flags value wrong")
	}
}

func BenchmarkSdshdr_Set(b *testing.B) {
	buf := []byte("hello world2, do you know youhello world2, do you know you ")
	for i := 0; i < b.N; i++ {
		SdsNewLen(buf)
	}
}
