package models

import "url-shortener/storage"

const keyPrefix = "url:"

func prependPrefix(key string) string {
	return keyPrefix + key
}

func AddUrl(key, item string) error {
	hasKey, err := storage.HasKey(prependPrefix(key))
	if err != nil {
		return err
	}
	if !hasKey {
		return storage.Set(prependPrefix(key), item)
	}
	return nil
}

type Url struct {
	Value  string
	Exists bool
}

func GetUrl(key string) (Url, error) {
	var url Url
	hasKey, err := storage.HasKey(prependPrefix(key))
	url.Exists = hasKey

	if err != nil || !hasKey {
		return url, err
	}

	value, err := storage.Get(prependPrefix(key))
	if err != nil {
		return url, err
	}
	url.Value = value
	return url, err
}
