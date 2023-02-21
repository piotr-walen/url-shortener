package config

import (
	"errors"
	"os"
	"strconv"
	"strings"
)

type primitive interface {
	string | int | bool
}

func noop(v string) (string, error) {
	return v, nil
}

func loadVariable[T primitive](
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
	MAX_URL_LENGTH int
	ADDR           string
	LOG_TRAFFIC    bool
	REDIS_HOST     string
	REDIS_PORT     string
	REDIS_PASSWORD string
}

const (
	DEFAULT_ADDR           = ":8000"
	DEFAULT_MAX_URL_LENGTH = 6
	DEFAULT_LOG_TRAFFIC    = true
)

var c = Config{
	ADDR:           DEFAULT_ADDR,
	MAX_URL_LENGTH: DEFAULT_MAX_URL_LENGTH,
	LOG_TRAFFIC:    DEFAULT_LOG_TRAFFIC,
}

func ParseConfig() error {
	r := reporter{
		reportedErrors: []error{},
	}
	loadVariable("ADDR", false, noop, func(v string) { c.ADDR = v }, r.report)
	loadVariable("LOG_TRAFFIC", false, strconv.ParseBool, func(v bool) { c.LOG_TRAFFIC = v }, r.report)
	loadVariable("MAX_URL_LENGTH", false, strconv.Atoi, func(v int) { c.MAX_URL_LENGTH = v }, r.report)
	loadVariable("REDIS_HOST", true, noop, func(v string) { c.REDIS_HOST = v }, r.report)
	loadVariable("REDIS_PORT", true, noop, func(v string) { c.REDIS_PORT = v }, r.report)
	loadVariable("REDIS_PASSWORD", true, noop, func(v string) { c.REDIS_PASSWORD = v }, r.report)

	return r.mergeErrors()
}

func GetConfig() *Config {
	return &c
}
