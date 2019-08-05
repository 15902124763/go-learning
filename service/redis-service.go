package service

import (
	"github.com/go-redis/redis"
	"go-learning/service/redis-client"
	"time"
)

/*
 基于go-redis
 github地址:https://github.com/go-redis/redis
 设置小时demo
*/
func SetHour() (result *redis.StatusCmd) {
	client :=redis_client.GetClient();
	// 注意time.Hour 要乘以一个具体的数值
	set := client.Set("hour", "3个小时", time.Hour*3)
	return set;
}

/*
 基于go-redis
 github地址:https://github.com/go-redis/redis
 设置分钟demo
*/
func SetMini() (result *redis.StatusCmd) {
	client :=redis_client.GetClient();
	// 注意time.Minute 要乘以一个具体的数值
	set := client.Set("minute", "5分钟", time.Minute*5)
	return set;
}

/*
 基于go-redis
 github地址:https://github.com/go-redis/redis
 get的demo
*/
func Get(key string) (*redis.StringCmd) {
	c := redis_client.GetClient();
    return c.Get(key);
}
