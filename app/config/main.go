package config

import (
	"encoding/json"
	"errors"
	"os"
	"strconv"
	"strings"
)

type RedisConfig struct {
	Name     string `json:"name"`
	Port     int    `json:"port"`
	Password string `json:"password"`
}

type variable interface {
	string | int | bool | []RedisConfig
}

func noop(v string) (string, error) {
	return v, nil
}

func parseRedisConfig(s string) ([]RedisConfig, error) {
	redisConfig := []RedisConfig{}
	err := json.Unmarshal([]byte(s), &redisConfig)
	return redisConfig, err
}

func loadVariable[T variable](
	name string,
	required bool,
	parse func(string) (T, error),
	set func(T),
	report func(error),
) {
	value, loaded := os.LookupEnv(name)
	if !loaded && required {
		report(errors.New("ENV: " + name + " is required"))
		return
	}
	if !loaded && !required {
		return
	}
	variable, err := parse(value)
	if err != nil {
		report(errors.New("ENV: " + name + " has invalid value: " + value))
		return
	}
	set(variable)
}

type reporter struct {
	reportedErrors []error
}

func (r *reporter) report(err error) {
	r.reportedErrors = append(r.reportedErrors, err)
}

func (r *reporter) mergeErrors() error {
	var err error
	if len(r.reportedErrors) > 0 {
		errStrings := []string{}
		for _, err := range r.reportedErrors {
			errStrings = append(errStrings, err.Error())
		}
		err = errors.New(strings.Join(errStrings, ", "))
	}
	return err
}

type Config struct {
	MaxUrlLength     int
	Addr             string
	LogTraffic       bool
	LogRedisInstance bool
	RedisConfig      []RedisConfig
}

const (
	DEFAULT_ADDR               = ":8000"
	DEFAULT_MAX_URL_LENGTH     = 6
	DEFAULT_LOG_TRAFFIC        = true
	DEFAULT_LOG_REDIS_INSTANCE = true
)

var c = Config{
	Addr:             DEFAULT_ADDR,
	MaxUrlLength:     DEFAULT_MAX_URL_LENGTH,
	LogTraffic:       DEFAULT_LOG_TRAFFIC,
	LogRedisInstance: DEFAULT_LOG_REDIS_INSTANCE,
}

func ParseConfig() error {
	r := reporter{
		reportedErrors: []error{},
	}
	loadVariable("ADDR", false, noop, func(v string) { c.Addr = v }, r.report)
	loadVariable("LOG_TRAFFIC", false, strconv.ParseBool, func(v bool) { c.LogTraffic = v }, r.report)
	loadVariable("LOG_REDIS_INSTANCE", false, strconv.ParseBool, func(v bool) { c.LogTraffic = v }, r.report)
	loadVariable("MAX_URL_LENGTH", false, strconv.Atoi, func(v int) { c.MaxUrlLength = v }, r.report)
	loadVariable("REDIS_CONFIG", true, parseRedisConfig, func(v []RedisConfig) { c.RedisConfig = v }, r.report)

	return r.mergeErrors()
}

func GetConfig() *Config {
	return &c
}
