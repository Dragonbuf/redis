package client

import "redis/adt"

const RedisReplyChunkBytes = 16 * 1024

type RedisClient struct {
	fd            int // 套接字描述符
	name          *adt.RedisObject
	flags         int        // 客户端角色及所处的状态
	queryBuf      adt.Sdshdr // 保存客户端发送的命令请求 超过 1GB 关闭此连接
	argc          int
	argv          []*adt.RedisObject // 命令参数
	cmd           map[adt.Sdshdr]*RedisCommand
	buf           [RedisReplyChunkBytes]byte // 固定缓存区
	bufPos        int
	reply         adt.List // 可变大小缓冲区  使用链表链接多个结构
	authenticated int      // 是否通过身份验证
}

type RedisCommand struct {
	// 命令的实现函数
	// 命令应该给定的参数
	// 命令的总执行次数
	// 总消耗时长
	// 命令的标志
}
