package models

import (
	"url-shortener/storage"

	"github.com/redis/go-redis/v9"
)

const keyPrefix = "url:"

func formatKey(namespace, segment string) string {
	return keyPrefix + namespace + "/" + segment
}

func AddUrl(namespace, segment, targetUrl string) (bool, error) {
	return storage.SetNX(namespace, formatKey(namespace, segment), targetUrl)
}

type GetResult struct {
	Value  string
	Exists bool
}

func GetUrl(namespace, segment string) (*GetResult, error) {
	value, err := storage.Get(namespace, formatKey(namespace, segment))

	if err == redis.Nil {
		return &GetResult{Exists: false}, nil
	}
	if err != nil {
		return nil, err
	}
	return &GetResult{Exists: true, Value: value}, nil
}
