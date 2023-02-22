package storage

func Set(key, value string) error {
	return GetRedisClient(key).Set(ctx, key, value, 0).Err()
}
