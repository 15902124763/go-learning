package redis_client

import (
	"go-learning/conf"
	"strconv"
	"github.com/go-redis/redis"
	"fmt"
	"go-learning/log"
)

var server = conf.Conf.Redis.Server
var port = conf.Conf.Redis.Port
var password = conf.Conf.Redis.Password
var db = conf.Conf.Redis.DB
// 注意int转字符串
var addr  = server + ":" + strconv.Itoa(port)

/*
 基于go-redis
 github地址:https://github.com/go-redis/redis
 redis连接demo，连接数等自行设置
*/
func  GetClient() (clint *redis.Client) {
	fmt.Println(addr)
	fmt.Println(port)
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password, // no password set
		DB:       db,  // use default DB
	})

	pong, err := client.Ping().Result()
	log.Infof(pong, err)
	return client
}

