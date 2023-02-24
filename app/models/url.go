package models

import (
	"url-shortener/storage"
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

func GetUrl(namespace, segment string) (GetResult, error) {
	var result GetResult
	hasKey, err := storage.HasKey(namespace, formatKey(namespace, segment))
	result.Exists = hasKey

	if err != nil || !hasKey {
		return result, err
	}

	value, err := storage.Get(namespace, formatKey(namespace, segment))
	if err != nil {
		return result, err
	}
	result.Value = value
	return result, err
}
