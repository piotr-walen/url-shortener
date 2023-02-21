package storage

func Get(key string) (string, error) {
	return rdb.Get(ctx, key).Result()
}
