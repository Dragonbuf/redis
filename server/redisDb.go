package server

// todo 先实现 dict 不然不好存储　redisDb
type RedisDb struct {
	dict map[string]*interface{} // 数据库空间，保存所有键值对
}

func NewRedisDb() *RedisDb {
	return &RedisDb{}
}

func (r *RedisDb) Set(key, value string) {

}

type dict struct {
}

type StringObject struct {
}

type HashObject struct {
}

type ListObject struct {
}

func Test() {

}
