package config

import (
	"errors"
	"os"
	"strconv"
	"strings"
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

type primitive interface {
	string | int | bool
}

func noop(v string) (string, error) {
	return v, nil
}

func loadVariable[T primitive](
	name string,
	parse func(string) (T, error),
	report func(error),
	defaultValue *T,
) T {
	value, ok := os.LookupEnv(name)
	if !ok && defaultValue == nil {
		report(errors.New("ENV: " + name + " is required"))
		return *defaultValue
	}
	if !ok && defaultValue != nil {
		return *defaultValue
	}
	parsed, err := parse(value)
	if err != nil {
		report(errors.New("ENV: " + name + " has invalid value: " + value))
		return *defaultValue
	}
	return parsed
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

func ParseConfig() error {
	r := reporter{
		reportedErrors: []error{},
	}
	setConfig(Config{
		ADDR:           loadVariable("ADDR", noop, r.report, &defaultConfig.ADDR),
		LOG_TRAFFIC:    loadVariable("LOG_TRAFFIC", strconv.ParseBool, r.report, &defaultConfig.LOG_TRAFFIC),
		MAX_URL_LENGTH: loadVariable("MAX_URL_LENGTH", strconv.Atoi, r.report, &defaultConfig.MAX_URL_LENGTH),
	})
	return r.mergeErrors()
}

func GetConfig() *Config {
	return &config
}
