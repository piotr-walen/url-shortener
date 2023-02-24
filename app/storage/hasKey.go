package storage

func HasKey(node string, key string) (bool, error) {
	res, err := GetRedisClient(node).Exists(ctx, key).Result()
	return res == 1, err
}
