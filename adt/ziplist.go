package adt

type ZipList struct {
	zlBytes uint32 // 整个压缩列表占用的内存字节数
	zlTail  uint32 // 压缩列表尾节点距离压缩列表起始地址有多少字节
	zlLen   uint16 //压缩列表的节点数量 值等于 65535 时，需要遍历
	entryX  int    // 各个节点
	zlEnd   uint8  //oxFF 压缩列表的末端
}

type EntryX struct {
	previousEntryLength int
	encoding            int
	content             []byte
}
