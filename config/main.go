package config

import (
	"os"
	"strconv"
)

type Config struct {
	MAX_URL_LENGTH int
	ADDR           string
	LOG_TRAFFIC    bool
}

var defaultConfig = Config{
	ADDR:           ":8000",
	MAX_URL_LENGTH: 6,
	LOG_TRAFFIC:    true,
}

var config = Config{}

func setConfig(newConfig Config) {
	config = newConfig
}

func parseStringVariable(name string, defaultValue string) string {
	value, ok := os.LookupEnv(name)
	if !ok {
		return defaultValue
	}
	return value
}

func parseBoolVariable(name string, defaultValue bool) bool {
	value, ok := os.LookupEnv(name)
	if !ok {
		return defaultValue
	}
	parsed, err := strconv.ParseBool(value)
	if err != nil {
		return defaultValue
	}
	return parsed
}

func parseIntVariable(name string, defaultValue int) int {
	value, ok := os.LookupEnv(name)
	if !ok {
		return defaultValue
	}
	parsed, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue
	}
	return parsed
}

func ParseConfig() {
	setConfig(Config{
		ADDR:           parseStringVariable(os.Getenv("ADDR"), defaultConfig.ADDR),
		LOG_TRAFFIC:    parseBoolVariable("LOG_TRAFFIC", defaultConfig.LOG_TRAFFIC),
		MAX_URL_LENGTH: parseIntVariable("MAX_URL_LENGTH", defaultConfig.MAX_URL_LENGTH),
	})
}

func GetConfig() *Config {
	return &config
}
