package storage

func HasKey(key string) (bool, error) {
	res, err := rdb.Exists(ctx, key).Result()
	return res == 1, err
}
