package adt

import (
	"unsafe"
)

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

// 本来需要根据　buf 长度来返回不同的类型，但是现在　go 不支持　１字节对齐，所以只返回一种
func SdsNewLen(s []byte) unsafe.Pointer {

	strLen := len(s)
	types := SdsReqType(strLen)
	if types == SdsType5 || strLen < 1<<32 {
		types = SdsType32
	}

	// 1 2 4 8 字节，　小于　1 字节暂不计算
	// 计算不同头部所需的空间
	//hdrlen := SdsHdrSize(types)

	switch types {
	case SdsType5:
	case SdsType8:
		sds := Sdshdr8{
			len:   uint8(strLen),
			alloc: uint8(strLen),
			flag:  uint8(1),
			buf:   s,
		}
		sdsPointer := unsafe.Pointer(&sds)
		return unsafe.Pointer(uintptr(sdsPointer) + unsafe.Offsetof(sds.buf))
	case SdsType16:
		sds := Sdshdr16{
			len:   uint16(strLen),
			alloc: uint16(strLen),
			flag:  uint8(2),
			buf:   s,
		}
		sdsPointer := unsafe.Pointer(&sds)
		return unsafe.Pointer(uintptr(sdsPointer) + unsafe.Offsetof(sds.buf))
	case SdsType32:
		sds := Sdshdr32{
			len:   uint32(strLen),
			alloc: uint32(strLen),
			flag:  uint8(3),
			buf:   s,
		}
		sdsPointer := unsafe.Pointer(&sds)
		return unsafe.Pointer(uintptr(sdsPointer) + unsafe.Offsetof(sds.buf))
	}

	sds := Sdshdr64{
		len:   uint64(strLen),
		alloc: uint64(strLen),
		flag:  uint8(4),
		buf:   s,
	}
	sdsPointer := unsafe.Pointer(&sds)
	return unsafe.Pointer(uintptr(sdsPointer) + unsafe.Offsetof(sds.buf))
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

//　根据　buf 长度计算　flag 太复杂，这里不在计算
func PointOffset(types int) int {
	return 0
}

// 因为确定的 sdstype32 所以 -8 找到 flags
func GetFlagsPointByBufPoint(buf unsafe.Pointer) unsafe.Pointer {
	return unsafe.Pointer(uintptr(buf) - uintptr(8))
}

type Sdshdr16 struct {
	len   uint16 // 已使用长度
	alloc uint16 //总长度
	flag  uint8  //低 3 位存类型，高 5 位预留
	buf   []byte
}
type Sdshdr32 struct {
	len   uint32 // 已使用长度
	alloc uint32 //总长度
	flag  uint8  //低 3 位存类型，高 5 位预留
	buf   []byte
}
type Sdshdr64 struct {
	len   uint64 // 已使用长度
	alloc uint64 //总长度
	flag  uint8  //低 3 位存类型，高 5 位预留
	buf   []byte
}

// 在 c 里面　buf 的　size 是　len + 1 因为最后是 \0 作为结束符
type Sdshdr struct {
	len  int
	free int
	buf  *[]byte //这里因为要展示 sds 的细节，所以还是使用 []byte 不是 *string , 但是 Get 返回会统一 *string 方便处理
}
