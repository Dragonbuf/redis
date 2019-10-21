package adt

import "unsafe"

const (
	SdsType5 = iota
	SdsType8
	SdsType16
	SdsType32
	SdsType64
)

type SdsHdr interface {
	SdsNewLen(len int)
}

type Sdshdr5 struct {
	flag uint8 // 低 3 位存类型，高 5 位存长度
	buf  []byte
}
type Sdshdr8 struct {
	len   uint8 // 已使用长度
	alloc uint8 //总长度
	flag  uint8 //低 3 位存类型，高 5 位预留
	buf   []byte
}

func SdsNewLen(s string) *[]byte {

	strLen := len(s)
	types := SdsReqType(strLen)
	if types == SdsType5 && strLen == 0 {
		types = SdsType8
	}

	// 1 2 4 8 字节，　小于　1 字节暂不计算
	// 计算不同头部所需的空间
	hdrlen := SdsHdrSize(types)
	// todo 可以使用　unsafePoint
	//sh := hdrlen + strLen + 1

	sds := Sdshdr32{
		flag:  0,
		len:   uint32(hdrlen + 1),
		alloc: 0,
		buf:   make([]byte, hdrlen+1), // 解决 c '\0'
	}

	return &sds.buf
}

func SdsFree() {

}

func SdsSetLen() {

}

func SdsHdrSize(types int) int {
	switch types {
	case SdsType5:
		return int(unsafe.Sizeof(Sdshdr5{}))
	case SdsType8:
		return int(unsafe.Sizeof(Sdshdr8{}))
	case SdsType16:
		return int(unsafe.Sizeof(Sdshdr16{}))
	case SdsType32:
		return int(unsafe.Sizeof(Sdshdr32{}))
	case SdsType64:
		return int(unsafe.Sizeof(Sdshdr64{}))
	}
	return 0
}

func SdsReqType(stringSize int) int {
	if stringSize < 1<<5 {
		return SdsType5
	}

	if stringSize < 1<<8 {
		return SdsType8
	}

	if stringSize < 1<<16 {
		return SdsType16
	}

	if stringSize < 1<<32 {
		return SdsType32
	}

	// 不考虑 32 位机器
	return SdsType64
}

type Sdshdr16 struct {
	flag  uint8  //低 3 位存类型，高 5 位预留
	len   uint16 // 已使用长度
	alloc uint16 //总长度
	buf   []byte
}
type Sdshdr32 struct {
	flag  uint8  //低 3 位存类型，高 5 位预留
	len   uint32 // 已使用长度
	alloc uint32 //总长度
	buf   []byte
}
type Sdshdr64 struct {
	flag  uint8  //低 3 位存类型，高 5 位预留
	len   uint64 // 已使用长度
	alloc uint64 //总长度
	buf   []byte
}

// 在 c 里面　buf 的　size 是　len + 1 因为最后是 \0 作为结束符
type Sdshdr struct {
	len  int
	free int
	buf  *[]byte //这里因为要展示 sds 的细节，所以还是使用 []byte 不是 *string , 但是 Get 返回会统一 *string 方便处理
}

func NewSdsHdr() *Sdshdr {
	return &Sdshdr{}
}

func (sds *Sdshdr) SetLen(len int) *Sdshdr {
	sds.len = len
	return sds
}

func (sds *Sdshdr) SetFree(free int) *Sdshdr {
	sds.free = free
	return sds
}
func (sds *Sdshdr) SetBuf(buf *[]byte) *Sdshdr {
	sds.buf = buf
	return sds
}

func (sds *Sdshdr) Set(s *string) int {
	i := len(*s)

	// 如果 sds 本身的 i 长度足够,直接更改
	if sds.HasEnoughLen(i) {
		buf := []byte(*s)
		sds.SetLen(i).SetBuf(&buf)
		return 1
	}

	// 如果空间不够了 ,申请所需空间 2 倍的空间,分别赋值给　free 和 i
	if !sds.HasEnoughLenWithFree(i) {
		buf := []byte(*s)
		sds.SetFree(i).SetLen(i).SetBuf(&buf)

		return 1
	} else {
		// 如果 i + free 足够使用，那么就直接使用 buf 存储
		buf := []byte(*s)
		sds.SetBuf(&buf).SetFree(i - sds.len).SetLen(i)

		return 1
	}
}

func (sds *Sdshdr) Get() *string {
	if sds.IsEmpty() {
		return nil
	}

	str := string(*sds.buf)
	return &str
}

func (sds *Sdshdr) GetLen() int {
	return sds.len
}

func (sds *Sdshdr) IsEmpty() bool {
	return sds.len == 0
}

func (sds *Sdshdr) HasEnoughLen(l int) bool {
	return sds.len >= l
}

func (sds *Sdshdr) HasEnoughLenWithFree(l int) bool {
	return sds.len+sds.free > l
}
