package storage

func SetNX(node, key, value string) (bool, error) {
	return GetRedisClient(node).SetNX(ctx, key, value, 0).Result()
}
