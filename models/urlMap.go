package models

type URLMap map[string]string

var urlMap URLMap = URLMap{}

func (*URLMap) add(key, item string) bool {
	_, ok := urlMap[key]
	if !ok {
		urlMap[key] = item
		return true
	}
	return false
}

func (*URLMap) get(key string) (string, bool) {
	value, ok := urlMap[key]
	return value, ok
}

func (*URLMap) hasKey(key string) bool {
	_, ok := urlMap[key]
	return ok
}

func AddUrl(key, item string) bool {
	return urlMap.add(key, item)
}

func GetUrl(key string) (string, bool) {
	return urlMap.get(key)
}

func ExistsUrl(key string) bool {
	return urlMap.hasKey(key)
}
