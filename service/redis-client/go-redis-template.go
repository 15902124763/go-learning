package redis_client

import (
	"github.com/go-redis/redis"
	"go-learning/conf"
	"go-learning/log"
	"strconv"
)

var server = conf.Conf.Redis.Server
var port = conf.Conf.Redis.Port
var password = conf.Conf.Redis.Password
var db = conf.Conf.Redis.DB
var addr = server + ":" + strconv.Itoa(port)

// 获取连接
func GetClient() (c *redis.Client) {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password, // no password set
		DB:       db,       // use default DB
	})

	pong, err := client.Ping().Result()
	log.Infof(pong, err)
	return client
}

func set() {

}
