package storage

func Get(node, key string) (string, error) {
	return GetRedisClient(node).Get(ctx, key).Result()
}
