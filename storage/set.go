package storage

func Set(key, value string) error {
	return rdb.Set(ctx, key, value, 0).Err()
}
