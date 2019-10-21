package adt

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"strconv"
	"testing"
	"unsafe"
)

type Num struct {
	i string
	j int
}

func TestSds(t *testing.T) {

	var sds8 Sdshdr8
	sds8.len = 0
	sds8.alloc = 0
	sds8.flag = 1
	sds8.buf = []byte("hello")

	bufE := &bytes.Buffer{}
	err := binary.Write(bufE, binary.LittleEndian, sds8)
	fmt.Printf("Error: %v\n", err)
	fmt.Println()
	//fmt.Printf("Sizeof myStruct: %d, Sizeof buf: %d, Len of buf: %d\n", unsafe.Sizeof(sds8), unsafe.Sizeof(buf) )
	// 32
	sds8Pointer := unsafe.Pointer(&sds8)
	buf := unsafe.Pointer(uintptr(sds8Pointer) + unsafe.Offsetof(sds8.buf))

	flagPoint := unsafe.Pointer(uintptr(buf) - uintptr(8))
	flagPoint2 := (*uint8)(flagPoint)
	fmt.Println(flagPoint2)
	*flagPoint2 = 3

	fmt.Println(sds8.flag)

	buf2 := (*[]byte)(buf)

	*buf2 = []byte("world")
	//unsafe.Alignof()

	fmt.Println(string(sds8.buf))
	//if (s == NULL) return;
	//  s_free((char*)s-sdsHdrSize(s[-1]));//此处直接释放内存

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
