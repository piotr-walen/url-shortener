package storage

func HasKey(key string) (bool, error) {
	res, err := GetRedisClient(key).Exists(ctx, key).Result()
	return res == 1, err
}
