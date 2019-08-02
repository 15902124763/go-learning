package redis

import (
	"github.com/go-redis/redis"
	"go-learning/log"
	"go-learning/conf"
)

var server = conf.Conf.Redis.Server
var port = conf.Conf.Redis.Port
var password = conf.Conf.Redis.Password
var db = conf.Conf.Redis.DB
var addr  = server + ":" + string(port)


// 获取连接
func GetClient()  {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password, // no password set
		DB:       db,  // use default DB
	})

	pong, err := client.Ping().Result()
	log.Infof(pong, err)
}
