package adt

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
