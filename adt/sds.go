package adt

import "fmt"

// 在 c 里面　buf 的　size 是　len + 1 因为最后是 \0 作为结束符
type Sdshdr struct {
	len  int
	free int
	buf  *[]byte
}

func NewSdsHdr() *Sdshdr {
	return &Sdshdr{}
}

func (sds *Sdshdr) Set(s string) int {
	len := len(s)

	fmt.Printf("debug: 设置字符串 %s, 需要空间 %d \n", s, len)
	// 如果 sds 本身的 len 长度足够,直接更改
	if sds.len >= len {
		fmt.Printf("debug: 空间　len %d 足够使用，直接赋值 \n", sds.len)
		buf := []byte(s)
		sds.buf = &buf
		sds.len = len
		return 1
	}

	// 如果空间不够了 ,申请所需空间 2 倍的空间,赋值给　free 和 len
	if (sds.free + sds.len) <= len {
		fmt.Printf("debug: 空间不够 len + free %d ，　共申请空间 %d \n", sds.free+sds.len, len*2)
		buf := make([]byte, len*2)
		sds.free = len
		sds.len = len
		buf = []byte(s)
		sds.buf = &buf
		return 1
	} else { // 如果 len + free 足够使用，那么就直接使用 buf 存储
		fmt.Printf("debug: 空间 len + free = %d 足够，　使用空间 \n", sds.len+sds.free)

		buf := []byte(s)
		sds.buf = &buf
		sds.free = len - sds.len
		sds.len = len
		return 1
	}

}

func (sds *Sdshdr) Get() string {
	return string(*sds.buf)
}
