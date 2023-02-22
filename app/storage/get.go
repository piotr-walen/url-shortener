package storage

func Get(key string) (string, error) {
	return GetRedisClient(key).Get(ctx, key).Result()
}
