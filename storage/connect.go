package storage

import (
	"context"
	"log"
	"url-shortener/config"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()
var rdb *redis.Client = nil

func Connect() error {
	log.Println("Connecting to redis...")

	rdb = redis.NewClient(&redis.Options{
		Addr:     config.GetConfig().REDIS_HOST + ":" + config.GetConfig().REDIS_PORT,
		Password: config.GetConfig().REDIS_PASSWORD,
		DB:       0,
	})

	log.Println("Checking connection to redis PING ...")
	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		return err
	}
	log.Println("... " + pong + " connected successfully ")
	return nil
}
