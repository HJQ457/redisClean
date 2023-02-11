package config

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

type RedisConfig struct {
	Host string
	Pass string
}

func RedisClean(r *RedisConfig) {

}

func (r *RedisConfig) Zero() {
	//连接redis
	dial0, coon_err := redis.Dial("tcp", r.Host, redis.DialPassword(r.Pass), redis.DialDatabase(0))
	if coon_err != nil {
		fmt.Printf("redis连接失败，err：%v", coon_err)
	}
	defer dial0.Close()

	//查找redis的key
	enumCache, _ := redis.Strings(dial0.Do("keys", "*enumCache*"))
	viewmodel, _ := redis.Strings(dial0.Do("keys", "*viewmodel*"))
	viewApplication, _ := redis.Strings(dial0.Do("keys", "*viewApplication*"))

	//删除查找出的key
	for i, _ := range enumCache {
		dial0.Do("del", enumCache[i])
	}
	for i, _ := range viewmodel {
		dial0.Do("del", viewmodel[i])
	}
	for i, _ := range viewApplication {
		dial0.Do("del", viewApplication[i])
	}
}

func (r *RedisConfig) Flush() {
	//连接redis
	dial14, coon_err14 := redis.Dial("tcp", r.Host, redis.DialPassword(r.Pass), redis.DialDatabase(14))
	if coon_err14 != nil {
		fmt.Printf("redis连接失败，err：%v", coon_err14)
	}
	defer dial14.Close()

	//清理14号库flushdb
	dial14.Do("flushdb")

	//连接redis
	dial16, coon_err16 := redis.Dial("tcp", r.Host, redis.DialPassword(r.Pass), redis.DialDatabase(16))
	if coon_err16 != nil {
		fmt.Printf("redis连接失败，err：%v", coon_err16)
	}
	defer dial16.Close()

	//清理14号库flushdb
	dial16.Do("flushdb")
}
