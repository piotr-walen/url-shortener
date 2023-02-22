package storage

import (
	"context"
	"log"
	"strconv"
	"url-shortener/config"

	"github.com/redis/go-redis/v9"
	"github.com/serialx/hashring"
)

var ctx = context.Background()

func connectSingle(c config.RedisConfig) (*redis.Client, error) {
	addr := c.Name + ":" + strconv.Itoa(c.Port)

	log.Println("Connecting to " + addr + " ...")

	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: c.Password,
		DB:       0,
	})

	log.Println("Checking connection to " + addr + " PING ...")
	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}
	log.Println("... " + pong + " " + addr + " connected successfully ")
	return rdb, err
}

var ring *hashring.HashRing
var serverMap = map[string]*redis.Client{}

func GetRedisClient(key string) *redis.Client {
	serverKey, ok := ring.GetNode(key)
	if !ok {
		return nil
	}
	return serverMap[serverKey]
}

func Connect() error {
	servers := []string{}
	for _, c := range config.GetConfig().RedisConfig {
		servers = append(servers, c.Name)
		rdb, err := connectSingle(c)
		if err != nil {
			return err
		}
		serverMap[c.Name] = rdb
	}
	ring = hashring.New(servers)

	return nil
}
